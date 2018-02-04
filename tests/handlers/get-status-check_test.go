package handlers

import (
	"net/http"
	"testing"

	"github.com/alefcarlos/carteiro-api/tests"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestStatusWithOk(t *testing.T) {
	r := gofight.New()

	r.GET("/status").
		SetDebug(true).
		Run(tests.GetGinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
