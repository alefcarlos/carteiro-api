package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/alefcarlos/carteiro-api/repo"
)

//GetAllSubscriptions retorna a lista de rastreios que têm novas informações
func GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	if result, err := repo.GetSubscriptions(); err != nil {
		utils.SendJSONBadRequest(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, result)
	}

}
