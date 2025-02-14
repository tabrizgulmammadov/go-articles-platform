package db

import (
	"context"
	"fmt"
	"github.com/tabrizgulmammadov/go-articles-platform/internal/store"
	"log"
	"math/rand"
)

var usernames = []string{
	"SkyWalker92",
	"CosmicPanda_23",
	"QuantumLeap55",
	"NightOwl247",
	"StarDust88",
	"ThunderBolt21",
	"MoonRider99",
	"PixelPioneer",
	"CyberNinja42",
	"OceanBreeze77",
	"MountainGoat15",
	"DesertFox404",
	"IceDragon66",
	"SunflowerSage",
	"RainbowRider83",
	"CloudNine95",
	"ElectricEagle",
	"WildWolf359",
	"MetalMaster44",
	"ForestRunner",
	"LavaLamp123",
	"AquaMarine87",
	"SilverArrow29",
	"GoldenPhoenix",
	"CrystalClear55",
	"MapleLeaf365",
	"RedPanda808",
	"BlueJay424",
	"GreenDragon71",
	"PurpleHaze99",
	"CoralReef33",
	"SandCastle77",
	"SnowLeopard21",
}

var titles = []string{
	"The Future of Artificial Intelligence: A Deep Dive",
	"10 Essential Tips for Remote Work Success",
	"Understanding Blockchain Technology Basics",
	"How to Master Time Management in 2025",
	"The Complete Guide to Cloud Computing",
	"Sustainable Living: Small Changes, Big Impact",
	"Digital Privacy in the Modern Age",
	"Building Resilience: Lessons from Global Leaders",
	"The Rise of Virtual Reality in Education",
	"Healthy Habits for Programming Professionals",
	"Machine Learning: From Theory to Practice",
	"The Art of Effective Communication",
	"Cybersecurity Best Practices for 2025",
	"Understanding Quantum Computing Basics",
	"The Impact of Social Media on Society",
	"Mindfulness and Productivity at Work",
	"Green Technology Innovations",
	"The Future of Remote Healthcare",
	"Data Science: Getting Started Guide",
	"Building Strong Team Culture Virtually",
	"Innovation in Renewable Energy",
	"The Psychology of Decision Making",
	"Web Development Trends for 2025",
	"Understanding 5G Technology",
	"The Evolution of E-commerce",
	"Mental Health in the Digital Age",
	"Space Exploration: Latest Discoveries",
	"The Power of Habit Formation",
	"Cryptocurrency Investment Basics",
	"Urban Farming: A Modern Solution",
	"The Future of Transportation",
	"Understanding Machine Vision",
	"Digital Marketing Strategies That Work",
	"The Science of Sleep and Productivity",
	"Next-Generation Mobile Technology",
	"Climate Change: Taking Action Now",
	"The Rise of No-Code Development",
	"Understanding Neural Networks",
	"The Future of Work: 2025 and Beyond",
	"Sustainable Architecture Trends",
	"The Impact of IoT on Daily Life",
	"Modern Database Technologies",
	"The Art of Public Speaking",
	"Understanding Cloud Security",
	"The Future of Online Education",
	"Renewable Energy Solutions",
	"The Psychology of User Experience",
	"Smart Cities: The Future is Now",
	"Understanding DevOps Culture",
	"The Evolution of Mobile Apps",
}

var contents = []string{
	"AI technology continues to evolve rapidly, transforming industries and daily life. From chatbots to autonomous systems, we explore the latest developments.",
	"Remote work requires new skills and habits. Learn essential strategies for staying productive and maintaining work-life balance in a virtual environment.",
	"Blockchain is more than just cryptocurrency. Discover the fundamental concepts and potential applications across different sectors.",
	"Effective time management is crucial in today's fast-paced world. Here are proven techniques to maximize your productivity and reduce stress.",
	"Cloud computing has revolutionized how businesses operate. Understanding its basics is essential for modern professionals.",
	"Small daily choices can have a significant environmental impact. Learn practical steps for more sustainable living.",
	"Protecting your digital privacy is more important than ever. Learn key strategies for safeguarding your personal information online.",
	"Building resilience is crucial for long-term success. Learn from global leaders about overcoming challenges and adapting to change.",
	"Virtual reality is transforming education, offering immersive learning experiences that were previously impossible.",
	"Maintaining physical and mental health is crucial for programmers. These habits will help you stay healthy while coding.",
	"Machine learning basics explained in simple terms, with practical examples and applications.",
	"Good communication is key to success in both personal and professional life. Master these essential skills.",
	"Stay protected against cyber threats with these updated security practices for the digital age.",
	"Quantum computing explained simply: what it is, how it works, and why it matters for our future.",
	"Examining the profound effects of social media on modern society and personal relationships.",
	"Incorporate mindfulness practices into your workday for better focus and productivity.",
	"Exploring the latest innovations in sustainable technology and their environmental impact.",
	"Telemedicine is changing healthcare delivery. Learn about the latest developments and benefits.",
	"Starting your journey in data science? Here's what you need to know to begin.",
	"Building team culture in remote environments requires new approaches. Discover effective strategies.",
	"Recent breakthroughs in renewable energy are making sustainable power more accessible.",
	"Understanding how we make decisions can help us make better choices. Learn the science behind it.",
	"Stay current with the latest trends shaping web development and user experience.",
	"5G technology is more than just faster internet. Learn about its transformative potential.",
	"E-commerce continues to evolve. Understand the latest trends and future directions.",
	"Digital wellness is crucial in today's connected world. Learn strategies for maintaining balance.",
	"Recent space discoveries are changing our understanding of the universe. Learn more.",
	"Habits shape our lives. Discover science-backed methods for building better habits.",
	"Understanding cryptocurrency basics before investing is crucial. Here's what you need to know.",
	"Urban farming could be the future of food production. Learn about innovative approaches.",
	"Transportation is being revolutionized by new technologies. Explore what's coming next.",
	"Machine vision is transforming industries. Learn how computers are learning to see.",
	"Effective digital marketing requires understanding modern consumer behavior and tools.",
	"Quality sleep is crucial for productivity. Learn science-backed tips for better rest.",
	"Mobile technology continues to advance. Discover what's on the horizon.",
	"Climate action is crucial now. Learn practical steps for making a difference.",
	"No-code platforms are democratizing software development. Learn how to leverage them.",
	"Neural networks explained simply: how they work and why they matter.",
	"The workplace is evolving rapidly. Prepare for the changes ahead.",
	"Sustainable architecture is shaping our cities. Explore innovative building designs.",
	"IoT devices are transforming daily life. Understand the benefits and challenges.",
	"Modern databases are evolving. Learn about new approaches to data management.",
	"Public speaking is a crucial skill. Master these essential techniques.",
	"Cloud security requires special attention. Learn key principles and best practices.",
	"Online education is here to stay. Discover how to make the most of digital learning.",
	"Renewable energy is becoming more accessible. Learn about the latest solutions.",
	"User experience design shapes how we interact with technology. Learn the basics.",
	"Smart cities are becoming reality. Explore how technology is transforming urban life.",
	"DevOps culture is changing how we build software. Understand the key principles.",
	"Mobile apps continue to evolve. Learn about the latest development trends.",
}

var tags = []string{
	"technology",
	"programming",
	"ai",
	"machine-learning",
	"web-development",
	"cloud-computing",
	"cybersecurity",
	"data-science",
	"blockchain",
	"career",
	"productivity",
	"software",
	"innovation",
	"startup",
	"database",
	"devops",
	"mobile",
	"networking",
	"javascript",
	"python",
	"golang",
	"react",
	"docker",
	"kubernetes",
	"aws",
	"azure",
	"microservices",
	"api",
	"testing",
	"design",
	"ui-ux",
	"architecture",
	"algorithms",
	"security",
	"linux",
	"git",
	"agile",
	"backend",
	"frontend",
	"fullstack",
	"performance",
	"analytics",
	"coding",
	"development",
	"engineering",
	"infrastructure",
	"automation",
	"best-practices",
	"tutorial",
	"guide",
}

var comments = []string{
	"Great article! This really helped me understand the basics.",
	"Would love to see a follow-up post on advanced topics in this area.",
	"The examples you provided were very clear and practical.",
	"I've been using this approach for a while and can confirm it works well.",
	"Thanks for sharing! This solved a problem I've been struggling with.",
	"Interesting perspective. Have you considered the impact on legacy systems?",
	"Very well explained. Even beginners can follow this easily.",
	"This is exactly what I was looking for. Bookmarked for future reference!",
	"Could you elaborate more on the security implications?",
	"The diagrams really helped me visualize the concept.",
	"I implemented this solution and it improved our performance by 50%!",
	"Nice write-up! Looking forward to more content like this.",
	"This is a game-changer for our development process.",
	"Would be great to see some code examples in other languages too.",
	"The troubleshooting section was particularly helpful.",
	"I appreciate the detailed explanations of each step.",
	"This matches my experience in production environments.",
	"Great introduction to a complex topic.",
	"The best explanation I've found on this subject so far.",
	"Very timely article. We're just starting to implement this.",
	"Have you considered writing about related topics?",
	"This should be required reading for junior developers.",
	"Finally, someone explained this in a way that makes sense!",
	"Excellent resource for teams transitioning to this technology.",
	"The comparison table was particularly insightful.",
	"I learned something new today. Thanks for sharing!",
	"Would love to see more real-world use cases.",
	"This helped our team resolve a long-standing issue.",
	"Clear, concise, and practical - exactly what we needed.",
	"Looking forward to implementing these best practices.",
	"The performance tips are especially valuable.",
	"This clarified several misconceptions I had.",
	"Great balance of theory and practical application.",
	"These insights will definitely improve our workflow.",
	"Solid advice backed by real experience.",
	"This approach saved us hours of debugging.",
	"Really appreciate the depth of this analysis.",
	"Perfect timing! We were just discussing this issue.",
	"The step-by-step guide was very helpful.",
	"This answers many questions our team had.",
	"Excellent overview of a complex subject.",
	"The benchmarks were particularly interesting.",
	"This will be very useful for our upcoming project.",
	"Great explanation of the trade-offs involved.",
	"Would love to hear more about edge cases.",
	"This has changed how I think about the problem.",
	"The security considerations were eye-opening.",
	"Very thorough coverage of the topic.",
	"This matches current industry best practices.",
	"Can't wait to try this approach in our system.",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Printf("error creating user: %v", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Printf("error creating post: %v", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Printf("error creating comment: %v", err)
			return
		}
	}

	log.Println("Seeding completed")
}

func generateUsers(count int) []*store.User {
	users := make([]*store.User, count)
	for i := 0; i < count; i++ {
		username := usernames[rand.Intn(len(usernames))]
		users[i] = &store.User{
			Username: username + fmt.Sprintf("%d", i),
			Email:    username + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123456",
		}
	}

	return users
}

func generatePosts(count int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, count)
	for i := 0; i < count; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(count int, users []*store.User, posts []*store.Post) []*store.Comment {
	generatedComments := make([]*store.Comment, count)
	for i := 0; i < count; i++ {
		generatedComments[i] = &store.Comment{
			UserID:  users[rand.Intn(len(users))].ID,
			PostID:  posts[rand.Intn(len(posts))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return generatedComments
}
