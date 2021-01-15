package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shingravirei/swgo/handler"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "World",
		})
	})

	api := r.Group("/api")
	{
		api.GET("/planet", handler.GetAllPlanets)
		api.GET("/planet/search", handler.SearchPlanet)
		api.POST("/planet", handler.AddPlanet)
		api.DELETE("/planet/:id", handler.DeletePlanet)

	}

	r.Run(":3000")
}
