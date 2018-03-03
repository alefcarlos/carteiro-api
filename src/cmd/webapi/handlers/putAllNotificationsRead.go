package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/src/internal/repo"
	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/utils"
)

//PutAllNotificationsRead atualiza todas as notificações para lida
func PutAllNotificationsRead(w http.ResponseWriter, r *http.Request) {
	if _, err := repo.PutAllNotificationsRead(); err != nil {
		utils.SendJSONBadRequest(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, "sucesso garoto ;)")
	}
}
