package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
)

//GetAvailableTrackings retorna a lista de rastreios que têm novas informações
func GetAvailableTrackings(w http.ResponseWriter, r *http.Request) {
	if result, err := repo.GetTrackingsToNotify(); err != nil {
		utils.SendJSONBadRequest(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, result)
	}
}
