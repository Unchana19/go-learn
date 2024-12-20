package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/Unchana19/social/internal/store"
)

var usernames = []string{
	"user1",
	"user2",
	"user3",
	"user4",
	"user5",
	"user6",
	"user7",
	"user8",
	"user9",
	"user10",
	"user11",
	"user12",
	"user13",
	"user14",
	"user15",
	"user16",
	"user17",
	"user18",
	"user19",
	"user20",
	"user21",
	"user22",
	"user23",
	"user24",
	"user25",
}

var titles = []string{
	"Understanding Golang Basics",
	"Mastering Concurrency in Go",
	"Introduction to Web Development with Go",
	"Building APIs with Go and Gin",
	"Understanding Pointers in Go",
	"File Handling in Go Made Easy",
	"Working with Goroutines and Channels",
	"Exploring Go's Standard Library",
	"Error Handling in Go Best Practices",
	"Optimizing Go Applications for Performance",
	"Go vs Python: A Developer's Perspective",
	"Getting Started with Go Modules",
	"Unit Testing in Go",
	"Creating RESTful Services with Go",
	"Effective Debugging in Go",
	"Implementing Microservices in Go",
	"Understanding Interfaces in Go",
	"Managing Dependencies in Go Projects",
	"Deploying Go Applications to the Cloud",
	"Real-World Applications of Go",
}

var contents = []string{
	"This article explores the basics of the Go programming language.",
	"Learn how to master concurrency using goroutines and channels in Go.",
	"A beginner-friendly guide to web development with the Go programming language.",
	"Step-by-step tutorial on building robust APIs with Go and Gin framework.",
	"A deep dive into pointers and their applications in Go programming.",
	"Understand file handling operations such as reading and writing in Go.",
	"Explore how goroutines and channels make Go a powerful concurrent language.",
	"An overview of Go's standard library and its most commonly used packages.",
	"Best practices for error handling and creating reliable Go applications.",
	"Tips and tricks for optimizing the performance of Go applications.",
	"A comparison of Go and Python for various use cases in software development.",
	"A complete guide to Go modules for dependency management in your projects.",
	"Learn how to write and execute unit tests for your Go programs.",
	"Build a RESTful service using Go with real-world examples and best practices.",
	"Discover effective debugging tools and techniques for Go developers.",
	"A practical guide to implementing microservices architecture using Go.",
	"Understand the power of interfaces and polymorphism in Go.",
	"Master dependency management in Go projects with hands-on examples.",
	"Deploy Go applications to cloud platforms like AWS and Google Cloud.",
	"Explore the real-world use cases of Go in industry and open-source projects.",
}

var tags = []string{
	"go",
	"programming",
	"golang",
	"web-development",
	"api",
	"concurrency",
	"goroutines",
	"channels",
	"error-handling",
	"performance",
	"microservices",
	"cloud",
	"file-handling",
	"unit-testing",
	"interfaces",
	"debugging",
	"rest-api",
	"standard-library",
	"deployment",
	"best-practices",
}

var comments = []string{
	"Great article! Very informative.",
	"I learned a lot about Go from this post.",
	"Can you provide more examples for beginners?",
	"This explanation is so clear and concise. Thanks!",
	"Looking forward to more posts like this.",
	"I had trouble understanding pointers, but this helped a lot.",
	"How do I implement this in a real-world project?",
	"This content is exactly what I was looking for!",
	"Could you elaborate on error handling in Go?",
	"Fantastic tutorial. Keep up the good work!",
	"The section on concurrency is a game changer for me.",
	"This cleared up all my doubts about Go modules.",
	"Are there any best practices for optimizing performance?",
	"Great job explaining interfaces in Go.",
	"Could you share some resources for further reading?",
	"This is by far the best guide I’ve found on this topic.",
	"Thanks for sharing your knowledge!",
	"Any tips for debugging Go applications?",
	"I’m new to Go, and this was super helpful!",
	"Do you plan to write about advanced Go topics in the future?",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
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

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
