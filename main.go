package main

import (
	"fmt"
	"log"
	"os"

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
		api.GET("/planet", env.getAllPlanetsHandler)
		api.GET("/planet/search", env.searchPlanetHandler)
		api.POST("/planet", env.addPlanetHandler)
		api.DELETE("/planet/:id", env.deletePlanetHandler)

	}

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
