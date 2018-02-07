package handlers

import (
	"net/http"
	"testing"

	"github.com/alefcarlos/carteiro-api/router"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestStatusWithOk(t *testing.T) {
	r := gofight.New()

	r.GET("/status").
		Run(router.GetRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
