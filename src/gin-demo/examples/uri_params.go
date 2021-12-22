package main

import (
	"github.com/gin-gonic/gin"
)

type Student struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/:name/:id", func(c *gin.Context) {
		var student Student

		if err := c.ShouldBindUri(&student); err != nil {
			c.JSON(400, gin.H{"msg": err})
			// return
		}

		c.JSON(200, gin.H{"name": student.Name, "uuid": student.ID})
	})

	r.Run(":8088")
}
