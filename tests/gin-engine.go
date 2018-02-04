package tests

import (
	"github.com/alefcarlos/carteiro-api/handlers"
	"github.com/gin-gonic/gin"
)

// GetGinEngine is gin router.
func GetGinEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

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

	return r
}
