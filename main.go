package main

import (
	"Project-Akhir/controllers"
	"Project-Akhir/database"
	"Project-Akhir/middleware"
	"Project-Akhir/repositori"
	"Project-Akhir/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ReadDB()

	userRepository := repositori.NewUSerRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(*userService)

	socialMediaRepository := repositori.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(*socialMediaRepository)
	socialMediaController := controllers.NewSocialMediaController(*socialMediaService)

	photoRepository := repositori.NewPhotoRepository(db)
	photoService := services.NewPhotoService(*photoRepository)
	photoController := controllers.NewPhotoController(*photoService)

	commetRepository := repositori.NewCommentRepository(db)
	commentService := services.NewCommentService(*commetRepository, *photoRepository)
	commentController := controllers.NewCommentControoler(*commentService)

	x := gin.Default()

	x.POST("/user/register", userController.Register)
	x.POST("/user/login", userController.Login)

	sosmedGroup := x.Group("/social_media", middleware.AuthMiddleware)
	sosmedGroup.POST("/", socialMediaController.CreateSocialMedia)
	sosmedGroup.GET("/", socialMediaController.GetAllSocialMedia)
	sosmedGroup.GET("/:id", socialMediaController.GetByIdSocialMedia)
	sosmedGroup.PUT("/:id", socialMediaController.UpdateSocialMedia)
	sosmedGroup.DELETE("/:id", socialMediaController.DeleteSocialMedia)

	photoGroup := x.Group("/photo", middleware.AuthMiddleware)
	photoGroup.POST("/", photoController.CreatePhoto)
	photoGroup.GET("/", photoController.GetAllPhoto)
	photoGroup.GET("/:id", photoController.GetPhotoById)
	photoGroup.PUT("/:id", photoController.UpdatePhoto)
	photoGroup.DELETE("/:id", photoController.DeletePhoto)

	commentGroup := x.Group("/comment", middleware.AuthMiddleware)
	commentGroup.POST("/:id_photo", commentController.CreateCommentByPhotoID)
	commentGroup.GET("/", commentController.GetAllCommnet)
	commentGroup.GET("/:id", commentController.GetCommentsByID)
	commentGroup.PUT("/:id", commentController.UpdateComment)
	commentGroup.DELETE("/:id", commentController.DeleteComment)

	x.Run(":8088")

}
