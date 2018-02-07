package handlers

import (
	"net/http"
	"strconv"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
	"github.com/julienschmidt/httprouter"
)

//PutTrackingSeen atualiza um registro como Lido
func PutTrackingSeen(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//Obtém o registro através do id
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		utils.SendJSONBadRequest(w, "Informe um Id válido.")
		return
	}

	if err = repo.UpdateTrackingRead(id); err == nil {
		utils.SendJSONSuccess(w, "Atualizado com sucesso.")
	} else {
		utils.SendJSONNotFound(w, err.Error())
	}
}
