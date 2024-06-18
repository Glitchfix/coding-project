package main

import (
	"github.com/Glitchfix/coding-project/config"
	"github.com/Glitchfix/coding-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.SetupDatabase()
	routes.SetupRoutes(r)

	r.Run()
}
