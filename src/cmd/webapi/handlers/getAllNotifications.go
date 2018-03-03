package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/src/internal/repo"
	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/utils"
)

//GetAllNotifications retorna todas as notificações
func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	if result, err := repo.GetNotifications(); err != nil {
		utils.SendJSONBadRequest(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, result)
	}

}
