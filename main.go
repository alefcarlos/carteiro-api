package main

import (
	"github.com/alefcarlos/carteiro-api/handlers"
	"github.com/gin-gonic/gin"
)

// GetMainEngine obtém as configurações do server
func GetMainEngine() *gin.Engine {
	r := gin.Default()

	r.GET("/status", handlers.GetStatus)

	trackingGroup := r.Group("/tracking")
	{
		trackingGroup.GET("/", handlers.GetAvailableTrackings)
		trackingGroup.PUT("/:id", handlers.PutTrackingInfo)
		trackingGroup.PUT("/:id/seen", handlers.PutTrackingSeen)
	}

	subscribeGroup := r.Group("/subscribe")
	{
		subscribeGroup.POST("/", handlers.PostNewSubscribe)
	}

	notify := r.Group("/notify")
	{
		notify.POST("/all", handlers.NotImplementedYet)
	}

	return r
}

func main() {
	GetMainEngine().Run(":8080")
}
