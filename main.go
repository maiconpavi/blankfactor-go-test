package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/maiconpavi/blankfactor-go-test/docs"
	"github.com/maiconpavi/blankfactor-go-test/view"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Blankfactor Test API
// @version 1.0
// @description This API is a test for Blankfactor

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	log.Println("Starting Gin")
	if os.Getenv("PRODUCTION_FLAG") != "" {
		gin.SetMode(gin.ReleaseMode)
	}
	docs.SwaggerInfo.BasePath = "/"
	log.Printf("Gin cold start, at: %s", docs.SwaggerInfo.Host)

	engine := gin.New()
	engine.Use(gin.Logger(), gin.CustomRecovery(HandleRecovery))

	view.EventView(engine)

	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	engine.Run(":8080")
}

func HandleRecovery(c *gin.Context, err any) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.(error).Error()})
}
