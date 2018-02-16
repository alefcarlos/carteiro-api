package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alefcarlos/carteiro-api/apiModels"
	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
)

//PostNotifyAll envia uma mensagem de notificação para todos os usuários inscritos
func PostNotifyAll(w http.ResponseWriter, r *http.Request) {
	var message apiModels.NotifyMessage

	b, err := utils.ReadBody(r)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	err = json.Unmarshal(b, &message)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	if err := repo.AddNotifyAll(&message); err != nil {
		utils.SendJSONConfict(w, err.Error())
	} else {
		utils.SendJSONSuccess(w, "Tudo certo, mano")
	}
}
