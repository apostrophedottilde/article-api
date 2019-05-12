package main

import (
	"github.com/apostrophedottilde/blog-article-api/app/article/repository/mongo"
	"github.com/apostrophedottilde/blog-article-api/app/article/repository/mongo/wrapper"
	"os"
	"os/signal"
	"syscall"

	"github.com/apostrophedottilde/blog-article-api/app/article/controller"
	"github.com/apostrophedottilde/blog-article-api/app/article/transport/http/adapter"
	"github.com/apostrophedottilde/blog-article-api/app/article/usecase"
)

func main() {

	dalWrapper := wrapper.NewMongoWrapper()
	userRepo := mongo.NewRepository(dalWrapper)
	useCase := usecase.New(userRepo)
	userController := controller.New(useCase)
	adapter := adapter.New(userController)

	adapter.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
}
