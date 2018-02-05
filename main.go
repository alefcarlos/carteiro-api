package main

import (
	"log"
	"net/http"

	"github.com/alefcarlos/carteiro-api/handlers"
	"github.com/julienschmidt/httprouter"
)

//GetMainEngine obtém as configurações do server
// func GetMainEngine() *gin.Engine {
// 	r := gin.Default()

// 	r.GET("/status", handlers.GetStatus)

// 	_trackingGroup := r.Group("/tracking")
// 	{
// 		_trackingGroup.GET("/", handlers.GetAvailableTrackings)
// 		_trackingGroup.PUT("/:id", handlers.PutTrackingInfo)
// 		_trackingGroup.PUT("/:id/seen", handlers.PutTrackingSeen)
// 	}

// 	_subscribeGroup := r.Group("/subscribe")
// 	{
// 		_subscribeGroup.POST("/", handlers.PostNewSubscribe)
// 	}

// 	_notify := r.Group("/notify")
// 	{
// 		_notify.POST("/all", handlers.NotImplementedYet)
// 	}

// 	return r
// }

func main() {
	router := httprouter.New()

	router.GET("/status", handlers.GetStatus)

	log.Fatal(http.ListenAndServe(":8080", router))
}
