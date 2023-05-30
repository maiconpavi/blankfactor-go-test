package view

import (
	"github.com/gin-gonic/gin"
	"github.com/maiconpavi/blankfactor-go-test/handler"
)

func EventView(router *gin.Engine) {
	group := router.Group("/event")
	{
		group.GET("/list", handler.EventList)
		group.GET("/list-overlap-pairs", handler.EventListOverlapPairs)
		group.POST("", handler.EventPost)
		group.GET("/:id", handler.EventGet)
		group.PUT("/:id", handler.EventPut)
		group.DELETE("/:id", handler.EventDelete)
	}
}
