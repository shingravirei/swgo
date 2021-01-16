package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type env struct {
	client *mongo.Client
	ctx    context.Context
}

func (e *env) getAllPlanets(c *gin.Context) {
	planets := e.findPlanets()

	c.JSON(200, planets)
}

func (e *env) searchPlanet(c *gin.Context) {
	c.JSON(200, gin.H{
		"search": "ok",
	})
}

func (e *env) addPlanet(c *gin.Context) {
	c.JSON(200, gin.H{
		"added": "a planet",
	})
}

func (e *env) deletePlanet(c *gin.Context) {
	type query struct {
		ID uint `uri:"id" binding:"required"`
	}

	var q query

	if err := c.ShouldBindUri(&q); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(200, q)
}

func (e *env) closeClient() {
	if err := e.client.Disconnect(e.ctx); err != nil {
		panic(err)
	}
}

func (e *env) newPlanet() {
	collection := e.client.Database("sw2").Collection("planet")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})

}

func (e *env) findPlanets() []bson.M {
	collection := e.client.Database("sw2").Collection("planet")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	var results []bson.M

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}
