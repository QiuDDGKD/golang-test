package ping

import (
	"github.com/gin-gonic/gin"
)

type helpDoc struct {
}

type QueryParam struct {
	Name   *string `json:"name" form:"name" binding:"required"`
	Age    *int    `json:"age" form:"age" binding:"required"`
	IsGood *bool   `json:"is_good" form:"is_good"`
}

func (doc *helpDoc) print() {
	println("help doc")
}

func Get(c *gin.Context) {
	var query QueryParam
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(500, gin.H{
			"code":    200,
			"message": "param error",
		})
		return
	}
	println(query.Name)
	println(query.Age)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "good",
	})
}

func Post(c *gin.Context) {
	var query QueryParam
	if err := c.ShouldBind(&query); err != nil {
		println(err.Error())
		c.JSON(500, gin.H{
			"code":    200,
			"message": "param error",
		})
		return
	}
	println("name: ", query.Name)
	println("age: ", query.Age)
	println("is_good: ", query.IsGood)
	if query.IsGood == nil {
		println("is")
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "param error",
	})
}
