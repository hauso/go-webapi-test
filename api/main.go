package main

import (
	"webapi/src/controller"
	"webapi/src/repository"
	"webapi/src/routes"
	"webapi/src/service"
	"webapi/infrastructure"
	"webapi/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()

	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRoute := routes.NewPostRoute(postController, router)
	postRoute.Setup()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	db.DB.AutoMigrate(&models.Post{}, &models.User{})

	router.Gin.Run(":8000")
}
