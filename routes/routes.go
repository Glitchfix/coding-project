package routes

import (
	controllers "github.com/Glitchfix/coding-project/controllers"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/Glitchfix/coding-project/docs"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/availability/:user_id", controllers.GetAvailability)
	router.POST("/availability", controllers.CreateAvailability)
	router.GET("/overlap/:user_id1/:user_id2", controllers.FindOverlap)
}
