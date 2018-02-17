package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
)

//PutAllNotificationsRead atualiza todas as notificações para lida
func PutAllNotificationsRead(w http.ResponseWriter, r *http.Request) {
	if _, err := repo.PutAllNotificationsRead(); err != nil {
		utils.SendJSONBadRequest(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, "sucesso garoto ;)")
	}
}
