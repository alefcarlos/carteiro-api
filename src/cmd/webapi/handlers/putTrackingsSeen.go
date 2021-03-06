package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/src/internal/repo"
	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/utils"
)

//PutTrackingsSeen todoas as notificações de tracking
func PutTrackingsSeen(w http.ResponseWriter, r *http.Request) {
	if _, err := repo.SetTrackingsRead(); err == nil {
		utils.SendJSONSuccess(w, "Atualizado com sucesso.")
	} else {
		utils.SendJSONNotFound(w, err.Error())
	}
}
