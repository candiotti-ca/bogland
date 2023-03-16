package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Plantounette struct {
	Name  string `json:"name" binding:"required"`
	Order string `json:"order"`
}

func setupRouter() *gin.Engine {
	plants := make(map[int]Plantounette, 0)
	r := gin.Default()

	r.GET("plants", func(c *gin.Context) {
		c.JSON(http.StatusOK, plants)
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo": "bar", // user := c.MustGet(gin.AuthUserKey).(string)
	}))

	authorized.POST("plants", func(c *gin.Context) {
		plantounette := Plantounette{}
		if c.Bind(&plantounette) == nil {
			plants[len(plants)] = plantounette
			c.Status(http.StatusCreated)
		} else {
			c.Status(http.StatusInternalServerError)
		}
	})

	authorized.PUT("plants/:id", func(c *gin.Context) {
		plantId, err := strconv.Atoi(c.Params.ByName("id"))

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if _, ok := plants[plantId]; !ok {
			c.Status(http.StatusNotFound)
			return
		}

		plantounette := Plantounette{}
		if c.Bind(&plantounette) == nil {
			plants[plantId] = plantounette
			c.Status(http.StatusNoContent)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{})
	})

	authorized.DELETE("plants/:id", func(c *gin.Context) {
		plantId, err := strconv.Atoi(c.Params.ByName("id"))

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if _, ok := plants[plantId]; !ok {
			c.Status(http.StatusNotFound)
			return
		}

		delete(plants, plantId)
		c.Status(http.StatusNoContent)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
