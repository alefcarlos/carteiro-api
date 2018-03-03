package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/src/internal/repo"
	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/utils"
)

//GetAvailableTrackings retorna a lista de rastreios que têm novas informações
func GetAvailableTrackings(w http.ResponseWriter, r *http.Request) {
	if result, err := repo.GetTrackingsToNotify(); err != nil {
		utils.SendJSONBadRequest(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, result)
	}
}
