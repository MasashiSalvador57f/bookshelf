package main

import (
	"bookshelf/controllers"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

// main ...
func main() {

	router := gin.Default()
	router.GET("/:id", func(c *gin.Context) {
		ctrl := controllers.NewUser()
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if id <= 0 {
			c.JSON(400, map[string]interface{}{})
			return
		}
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, map[string]interface{}{})
			return
		}
		c.JSON(200, result)
	})

	router.POST("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{"foo": "post"})
	})

	router.PUT("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{"foo": "put"})
	})

	router.DELETE("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{"foo": "delete"})
	})
	router.Run(":8080")

}
