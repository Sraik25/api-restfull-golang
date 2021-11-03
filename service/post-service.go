package service

import (
	"errors"
	"example/entity"
	"example/repository"
	"math/rand"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	// posts []entity.Post

	repo repository.PostRepository
)

func NewPostService(rep repository.PostRepository) PostService {
	repo = rep
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}

	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = int(rand.Int63())
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
