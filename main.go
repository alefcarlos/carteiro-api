package main

import (
	"github.com/alefcarlos/carteiro-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "tudo ok, porra",
		})
	})

	r.GET("/trackings", handlers.GetTrakings)

	r.POST("/subscribe", handlers.PostNewSubscribe)
	r.PUT("/subscribe", handlers.NotImplementedYet)

	r.Run(":3000")
}
