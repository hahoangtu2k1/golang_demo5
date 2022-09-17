package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func demoGin() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"GetAll": "done",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Post": "Done",
		})
	})
	r.GET("/query", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		add := c.Query("add")
		c.JSON(http.StatusOK, gin.H{
			"Ten":    name,
			"Tuoi":   age,
			"DiaChi": add,
		})
	})
	r.GET("/params/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"Ten":  name,
			"Tuoi": age,
		})
	})
	r.POST("/posts", func(c *gin.Context) {
		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, gin.H{
				"information": string(value),
			})
		}

	})
	r.Run()
}
