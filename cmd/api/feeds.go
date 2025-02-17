package main

import (
	"github.com/tabrizgulmammadov/go-articles-platform/internal/store"
	"net/http"
)

// GetUserFeed godoc
//
//	@Summary		Get user feed
//	@Description	Returns a paginated list of posts for a specific user feed.
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int							false	"Number of posts to retrieve (default: 20)"
//	@Param			offset	query		int							false	"Offset for pagination (default: 0)"
//	@Param			sort	query		string						false	"Sorting order (asc/desc, default: desc)"
//	@Success		200		{array}		[]store.PostWithMetadata	"List of posts in user feed"
//	@Failure		400		{object}	error						"Invalid query parameters"
//	@Failure		500		{object}	error						"Internal server error"
//	@Router			/users/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	feeds, err := app.store.Posts.GetUserFeed(ctx, int64(1), fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feeds); err != nil {
		app.internalServerError(w, r, err)
	}
}
