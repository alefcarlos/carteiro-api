package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/alefcarlos/carteiro-api/models"
)

//GetSubscription obtém informações de uma inscrição
func GetSubscription(w http.ResponseWriter, r *http.Request) {
	var addressInfo models.BotFrameworkAddressInfo

	b, err := utils.ReadBody(r)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	err = json.Unmarshal(b, &addressInfo)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	if item, err := repo.GetSubscription(addressInfo); err != nil {
		utils.SendJSONBadRequest(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, item)
	}
}
