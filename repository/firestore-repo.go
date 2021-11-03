package repository

import (
	"context"
	"example/entity"
	"log"

	"cloud.google.com/go/firestore"
)

type repo struct{}

// newFirestoreRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "pragmatic-reviews"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of post: %v", err)
			return nil, err
		}

		post := entity.Post{
			Id:    doc.Data()["Id"].(int),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}
	return posts, nil

}
