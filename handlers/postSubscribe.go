package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alefcarlos/carteiro-api/apiModels"
	"github.com/alefcarlos/carteiro-api/models"
	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
)

//PostNewSubscribe adicionar uma nova inscrição de notificação
func PostNewSubscribe(w http.ResponseWriter, r *http.Request) {
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

	info := models.SubscribeInfo{
		Address: subscription.Address,
	}

	if itemAdded, err := repo.AddSubscription(info); err != nil {
		utils.SendJSONConfict(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, itemAdded)
	}
}
