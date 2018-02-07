package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/alefcarlos/carteiro-api/errors"
	"github.com/alefcarlos/carteiro-api/models"
	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
)

//PutTrackingInfo atualiza as informações de status de um determinado tracking
func PutTrackingInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	info := models.TrackingUpdateInfo{}

	//Obtém o registro através do id
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		utils.SendJSONBadRequest(w, "Informe um Id válido.")
		return
	}

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

	if err = repo.UpdateTrackingInfo(id, info); err != nil {
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
