package handlers

import (
	"strconv"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
	"github.com/gin-gonic/gin"
)

//PutTrackingSeen atualiza um registro como Lido
func PutTrackingSeen(c *gin.Context) {
	//Obtém o registro através do id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.SendBadRequest(c, "Informe um Id válido.")
		return
	}

	if err = repo.UpdateTrackingRead(id); err == nil {
		utils.SendSuccess(c, "Atualizado com sucesso.")
	} else {
		utils.SendNotFound(c, err.Error())
	}
}
