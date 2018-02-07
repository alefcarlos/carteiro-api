package router

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/handlers"
	"github.com/julienschmidt/httprouter"
)

//GetRouter Obtém rotas
func GetRouter() http.Handler {
	r := httprouter.New()
	r.GET("/status", handlers.GetStatus)

	r.GET("/tracking", handlers.GetAvailableTrackings)
	r.PUT("/tracking/:id", handlers.PutTrackingInfo)
	r.PUT("/tracking/:id/seen", handlers.PutTrackingSeen)

	r.POST("/subscribe", handlers.PostNewSubscribe)

	r.POST("/notify/all", handlers.NotImplementedYet)

	return r
}
