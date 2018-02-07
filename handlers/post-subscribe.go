package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/alefcarlos/carteiro-api/models"
)

//PostNewSubscribe permite adicionar um novo rastreio para monitoramento
func PostNewSubscribe(w http.ResponseWriter, r *http.Request) {
	var tracking models.TrackingInfo

	b, err := utils.ReadBody(r)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	err = json.Unmarshal(b, &tracking)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	if itemAdded, err := repo.AddTracking(tracking); err != nil {
		utils.SendJSONConfict(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, itemAdded.ID)
	}
}
