package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, ctx := connectDb()

	env := &env{client, ctx}
	defer env.closeClient()

	api := r.Group("/api")
	{
		api.GET("/planet", env.getAllPlanets)
		api.GET("/planet/search", env.searchPlanet)
		api.POST("/planet", env.addPlanet)
		api.DELETE("/planet/:id", env.deletePlanet)

	}

	r.Run(":3000")
}
