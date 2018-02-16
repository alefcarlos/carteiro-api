package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alefcarlos/carteiro-api/apiModels"
	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
)

//PostNewTracking permite adicionar um novo rastreio para monitoramento
func PostNewTracking(w http.ResponseWriter, r *http.Request) {
	var subscription apiModels.NewSubscription

	b, err := utils.ReadBody(r)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	err = json.Unmarshal(b, &subscription)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	if itemAdded, err := repo.AddTracking(subscription); err != nil {
		utils.SendJSONConfict(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, itemAdded)
	}
}
