package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/Unchana19/go-learn/internal/store"
)

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(20)
	for _, u := range users {
		if err := store.Users.Create(ctx, u); err != nil {
			log.Println("Error creating user: ", err)
			return
		}
	}

	posts := generatePosts(50, users)
	for _, p := range posts {
		if err := store.Posts.Create(ctx, p); err != nil {
			log.Println("Error creating post: ", err)
			return
		}
	}

	comments := generateComments(100, posts, users)
	for _, c := range comments {
		if err := store.Comments.Create(ctx, c); err != nil {
			log.Println("Error creating comment: ", err)
			return
		}
	}

	log.Println("Database seeded successfully")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: fmt.Sprintf("user_%d", i),
			Email:    fmt.Sprintf("user_%d@example.com", i),
			Password: "password",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		posts[i] = &store.Post{
			Title:   fmt.Sprintf("Post %d", i),
			Content: fmt.Sprintf("Content for post %d", i),
			UserID:  users[rand.Intn(len(users))].ID,
			Tags:    []string{fmt.Sprintf("tag_%d", rand.Intn(10)), fmt.Sprintf("tag_%d", rand.Intn(10))},
		}
	}

	return posts
}

func generateComments(num int, posts []*store.Post, users []*store.User) []*store.Comment {
	comments := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		comments[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: fmt.Sprintf("Comment %d", i),
		}
	}

	return comments
}
