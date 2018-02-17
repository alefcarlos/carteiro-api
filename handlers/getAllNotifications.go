package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
)

//GetAllNotifications retorna todas as notificações
func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	if result, err := repo.GetNotifications(); err != nil {
		utils.SendJSONBadRequest(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, result)
	}

}
