package handler

import "github.com/gin-gonic/gin"

// GetAllPlanets returns all current planets from the DB
func GetAllPlanets(c *gin.Context) {
	c.JSON(200, gin.H{
		"all": "planets",
	})
}

// SearchPlanet does as it name says
func SearchPlanet(c *gin.Context) {
	c.JSON(200, gin.H{
		"search": "ok",
	})
}

// AddPlanet adds a planet
func AddPlanet(c *gin.Context) {
	c.JSON(200, gin.H{
		"added": "a planet",
	})
}

// DeletePlanet delets a planet
func DeletePlanet(c *gin.Context) {
	type Planet struct {
		ID uint `uri:"id" binding:"required"`
	}

	var planet Planet

	if err := c.ShouldBindUri(&planet); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(200, planet)
}
