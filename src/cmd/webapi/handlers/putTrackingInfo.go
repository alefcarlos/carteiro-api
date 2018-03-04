package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/apiModels"
	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/utils"
	"github.com/alefcarlos/carteiro-api/src/internal/errors"
	"github.com/alefcarlos/carteiro-api/src/internal/repo"
)

//PutTrackingInfo atualiza as informações de status de um determinado tracking
func PutTrackingInfo(w http.ResponseWriter, r *http.Request) {
	info := apiModels.TrackingUpdateInfo{}

	b, err := utils.ReadBody(r)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	err = json.Unmarshal(b, &info)
	if err != nil {
		utils.SendJSONBadRequest(w, err.Error())
		return
	}

	if err = repo.UpdateTrackingInfo(info); err != nil {
		switch err {
		case errors.ErrIDNotFound:
			utils.SendJSONNotFound(w, err.Error())
		default:
			utils.SendJSONBadRequest(w, err.Error())
		}

	} else {
		utils.SendJSONSuccess(w, "Informações atualizadas com sucesso.")
	}
}
