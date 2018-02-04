package handlers

import (
	"net/http"
	"testing"

	"github.com/alefcarlos/carteiro-api/models"

	"github.com/alefcarlos/carteiro-api/tests"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

//TestPostNewSubscribeWithBadRequest verifica se ao passar um json inválido, a api retorna 400
func TestPostNewSubscribeWithBadRequest(t *testing.T) {
	r := gofight.New()

	r.POST("/subscribe/").
		SetDebug(true).
		SetJSON(gofight.D{}).
		Run(tests.GetGinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}

func TestPostNewSubscribeWithOk(t *testing.T) {
	r := gofight.New()

	_tracking := models.TrackingInfo{
		TrackingCode: "alef",
		Address: models.BotFrameworkAddressInfo{
			ChannelID:  "channelId",
			ServiceURL: "serviceUrl",
			Bot: models.BotFrameworkIdentityInfo{
				ID:   "bot-id",
				Name: "bot-name",
			},
			User: models.BotFrameworkIdentityInfo{
				ID:   "user-id",
				Name: "user-name",
			},
		},
	}

	r.POST("/subscribe/").
		SetDebug(true).
		// SetJSON(_tracking). como passar aqui ?
		Run(tests.GetGinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
