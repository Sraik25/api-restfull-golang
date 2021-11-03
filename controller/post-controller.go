package controller

import (
	"encoding/json"
	"net/http"

	"example/entity"
	"example/errors"
	"example/service"
)

type controller struct {}

var (
	// posts []entity.Post

	postService service.PostService 
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

// Obtiene todos los posts
func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the post"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

// Agrega un post
func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
		return
	}

	err1 := postService.Validate(&post)

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)

	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
