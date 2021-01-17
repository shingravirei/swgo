package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type env struct {
	client *mongo.Client
	ctx    context.Context
}

func (e *env) getAllPlanetsHandler(c *gin.Context) {
	planets := e.findAllPlanets()

	c.JSON(200, planets)
}

func (e *env) searchPlanetHandler(c *gin.Context) {
	type query struct {
		Name string `form:"name" json:"name"`
	}

	var q query

	c.Bind(&q)

	planet := e.searchAPlanet(q.Name)

	c.JSON(200, planet)
}

func (e *env) addPlanetHandler(c *gin.Context) {
	type NewPlanet struct {
		Name          string
		Climate       string
		Terrain       string
		NumberOfFilms int
	}

	var newPlanet NewPlanet

	c.ShouldBindJSON(&newPlanet)
	newPlanet.NumberOfFilms = getPlanetMovieCount(newPlanet.Name)

	e.insertPlanet(&newPlanet)

	c.Status(201)
}

func (e *env) deletePlanetHandler(c *gin.Context) {
	type uri struct {
		ID string `uri:"id" binding:"required"`
	}

	var q uri

	if err := c.ShouldBindUri(&q); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	result := e.deleteOnePlanet(q.ID)

	if result.DeletedCount == 0 {
		c.Status(404)
	} else {
		c.Status(204)
	}

}

func (e *env) closeClient() {
	if err := e.client.Disconnect(e.ctx); err != nil {
		panic(err)
	}
}

func (e *env) insertPlanet(planet interface{}) {
	collection := e.client.Database("sw2").Collection("planet")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	doc, _ := bson.Marshal(planet)

	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}

}

func (e *env) findAllPlanets() []bson.M {
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

func (e *env) deleteOnePlanet(id string) *mongo.DeleteResult {
	collection := e.client.Database("sw2").Collection("planet")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": idPrimitive}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (e *env) searchAPlanet(name string) planet {
	collection := e.client.Database("sw2").Collection("planet")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.M{"name": name}

	var p planet

	err := collection.FindOne(ctx, filter).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	return p
}
