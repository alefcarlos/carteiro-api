package router

import (
	"context"
	"net/http"

	"github.com/alefcarlos/carteiro-api/handlers"
	"github.com/alefcarlos/carteiro-api/middlewares"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

//Método auxiliar para fazer um wrap no httprouter.Handle,
//pois httprouter utiliza o terceiro parâmetro para os parâmetros da requisição
//Então temos de inserir o valor direto no contexto
//https://gist.github.com/nmerouze/5ed810218c661b40f5c4
//https://www.nicolasmerouze.com/guide-routers-golang/ : ntegrate httprouter to our framework
func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

//GetRouter Obtém rotas
func GetRouter() http.Handler {

	commomHandlers := alice.New(middlewares.Logger, middlewares.Recover)
	r := httprouter.New()

	r.GET("/status", wrapHandler(commomHandlers.ThenFunc(handlers.GetStatus)))

	// r.GET("/tracking/notify", wrapHandler(commomHandlers.ThenFunc(handlers.GetAvailableTrackings)))

	// r.PUT("/tracking/:id", wrapHandler(commomHandlers.ThenFunc(handlers.PutTrackingInfo)))
	// r.PUT("/tracking/:id/seen", wrapHandler(commomHandlers.ThenFunc(handlers.PutTrackingSeen)))

	r.GET("/subscriptions", wrapHandler(commomHandlers.ThenFunc(handlers.GetAllSubscriptions)))
	r.POST("/subscriptions/get", wrapHandler(commomHandlers.ThenFunc(handlers.GetSubscription)))
	r.POST("/subscriptions", wrapHandler(commomHandlers.ThenFunc(handlers.PostNewSubscribe)))

	// r.POST("/notify/all", wrapHandler(commomHandlers.ThenFunc(handlers.NotImplementedYet)))

	return r
}
