package handlers

import (
	"net/http"
	"testing"

	"github.com/alefcarlos/carteiro-api/handlers"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

//TestCheckStatus verifica se o resultado da /status foi 200 - OK
func TestGinHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/status").
		SetDebug(true).
		Run(handlers.GetStatusGinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
