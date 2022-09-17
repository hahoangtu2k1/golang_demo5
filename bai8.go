package main

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

type Name struct {
	Name string
}

func Bai8() {
	r := gin.Default()
	r.GET("/names", getAll)
	r.POST("/names", postName)
	r.GET("/names/:index", getNameByLocation)
	r.POST("/names/:index", postNameByLocation)
	r.DELETE("/names/:index", deleteNameByLocation)
	r.Run()

}
func getAll(c *gin.Context) {
	// d := []string{}
	total, err := connRedis.LLen(ctx, "userName").Result()
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	} else {
		for i := 0; i < int(total); i++ {
			data, err := connRedis.LIndex(ctx, "userName", int64(i)).Result()
			if err != nil {
				c.JSON(404, gin.H{
					"error": err,
				})
			} else {
				c.JSON(200, gin.H{
					"name": string(data),
				})
			}
		}
	}
}
func postName(c *gin.Context) {
	name := Name{
		Name: "New Name",
	}
	update, err := json.Marshal(name)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	} else {
		_, err := connRedis.LPush(ctx, "userName", update).Result()
		if err != nil {
			c.JSON(404, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"update done": string(update),
			})
		}
	}
}
func getNameByLocation(c *gin.Context) {
	param := c.Param("index")
	conv, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(404, gin.H{
			"index is not an integer": err,
		})
	} else {
		getIndex, err := connRedis.LIndex(ctx, "userName", int64(conv)).Result()
		if err != nil {
			c.JSON(404, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"name": string(getIndex),
			})
		}

	}
}
func postNameByLocation(c *gin.Context) {
	param := c.Param("index")
	conv, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	} else {
		data, err := connRedis.LIndex(ctx, "userName", int64(conv)).Result()
		if err != nil {
			c.JSON(404, gin.H{
				"error": err,
			})
		} else {
			_, err := connRedis.LSet(ctx, "userName", int64(conv), "update done").Result()
			if err != nil {
				c.JSON(404, gin.H{
					"error": err,
				})
			} else {
				c.JSON(200, gin.H{
					"update successfully": string(data),
				})
			}
		}
	}
}
func deleteNameByLocation(c *gin.Context) {
	param := c.Param("index")
	conv, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	} else {
		data, err := connRedis.LIndex(ctx, "userName", int64(conv)).Result()
		if err != nil {
			c.JSON(404, gin.H{
				"error": err,
			})
		} else {
			_, err := connRedis.LRem(ctx, "userName", int64(conv), data).Result()
			if err != nil {
				c.JSON(404, gin.H{
					"error": err,
				})
			} else {
				c.JSON(200, gin.H{
					"Delete Done": string(data),
				})
			}
		}

	}
}
