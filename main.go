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

	_trackingGroup := r.Group("/tracking")
	{
		_trackingGroup.GET("/", handlers.GetAvailableTrackings)
		_trackingGroup.PUT("/:id", handlers.PutTrackingInfo)
		_trackingGroup.PUT("/:id/seen", handlers.PutTrackingSeen)
	}

	_subscribeGroup := r.Group("/subscribe")
	{
		_subscribeGroup.POST("/", handlers.PostNewSubscribe)
	}

	r.Run(":3000")
}
