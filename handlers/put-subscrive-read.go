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

	if _err == nil && repo.UpdateTrackingRead(_id) {
		utils.SendSuccess(c, "Atualizado com sucesso.")
	} else {
		utils.SendNotFound(c, "O ID não foi encontrado.")
	}
}
