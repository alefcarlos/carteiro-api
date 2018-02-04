package handlers

import (
	"strconv"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
	"github.com/gin-gonic/gin"
)

//PutSubscribeRead atualiza um registro como Lido
func PutSubscribeRead(c *gin.Context) {
	//Obtém o registro através do id
	_id, _err := strconv.Atoi(c.Param("id"))

	if _err != nil {
		utils.SendBadRequest(c, "Informe um Id válido.")
		return
	}

	if _err = repo.UpdateTrackingRead(_id); _err == nil {
		utils.SendSuccess(c, "Atualizado com sucesso.")
	} else {
		utils.SendNotFound(c, _err.Error())
	}
}
