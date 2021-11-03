package main

import (
	"fmt"
	"net/http"

	"example/controller"
	router "example/http"
	"example/repository"
	"example/service"
)

var (
	PostRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(PostRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/", postController.AddPost)

	httpRouter.SERVER(port)
}
