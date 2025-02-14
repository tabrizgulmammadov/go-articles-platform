package store

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
	Comments  []Comment `json:"comments"`
	User      User      `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentsCount int `json:"comments_count"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
		SELECT id, content, title, user_id, tags, created_at, updated_at, version
		  FROM posts
		 WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var post Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Content,
		&post.Title,
		&post.UserID,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostStore) GetUserFeed(ctx context.Context, id int64, fq PaginatedFeedQuery) ([]PostWithMetadata, error) {
	baseQuery := `
		SELECT p.id, p.user_id, p.title, p.content, p.tags, p.created_at, p.version,
		       u.username,
		       COUNT(c.id) AS comments_count
		  FROM posts p
	 LEFT JOIN comments c ON c.post_id = p.id
	 LEFT JOIN users u ON u.id = p.user_id
	      JOIN followers f ON f.follower_id = p.user_id
		 WHERE (f.user_id = $1 OR p.user_id = $1)`

	args := []interface{}{id} // $1 is always the user ID
	if fq.Search != "" {
		baseQuery += " AND (p.title ILIKE '%' || $2 || '%' OR p.content ILIKE '%' || $2 || '%')"
		args = append(args, fq.Search)
	}
	if len(fq.Tags) > 0 {
		baseQuery += " AND (p.tags @> $" + strconv.Itoa(len(args)+1) + "::varchar[])"
		args = append(args, pq.Array(fq.Tags))
	} else {
		baseQuery += " AND (p.tags @> $" + strconv.Itoa(len(args)+1) + "::varchar[] OR $" + strconv.Itoa(len(args)+1) + "::varchar[] = '{}'::varchar[])"
		args = append(args, pq.Array([]string{}))
	}

	baseQuery += " GROUP BY p.id, u.username ORDER BY p.created_at " + fq.Sort + " LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, fq.Limit, fq.Offset)

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feed []PostWithMetadata
	for rows.Next() {
		var p PostWithMetadata
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Title,
			&p.Content,
			pq.Array(&p.Tags),
			&p.CreatedAt,
			&p.Version,
			&p.User.Username,
			&p.CommentsCount)
		if err != nil {
			return nil, err
		}

		feed = append(feed, p)
	}

	return feed, nil
}

func (s *PostStore) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	query := `
		UPDATE posts
		   SET title = $1, content = $2, version = version + 1
		 WHERE id = $3 AND version = $4
	 RETURNING version
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
		post.Version,
	).Scan(&post.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}
