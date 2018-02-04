package main

import (
	"github.com/alefcarlos/carteiro-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/status", handlers.GetStatus)

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

	_notify := r.Group("/notify")
	{
		_notify.POST("/all", handlers.NotImplementedYet)
	}

	r.Run(":3000")
}
