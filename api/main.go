package main

import (
	"msgapi/src/controller"
	"msgapi/src/repository"
	"msgapi/src/routes"
	"msgapi/src/service"
	"msgapi/infrastructure"
	"msgapi/models"
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
