package handlers

import (
	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/alefcarlos/carteiro-api/models"
	"github.com/gin-gonic/gin"
)

//PostNewSubscribe permite adicionar um novo rastreio para monitoramento
func PostNewSubscribe(c *gin.Context) {
	var tracking models.TrackingInfo

	//Obter o objeto a partir do body
	if _err := c.ShouldBindJSON(&tracking); _err == nil {
		//Adicionar novo objeto na listagem
		if err := repo.AddTracking(tracking); err != nil {
			utils.SendConflict(c, "O código já foi informado para monitoramento")
		} else {
			utils.SendSuccess(c, "Novo monitoramento cadastrado com sucesso!")
		}

	} else {
		utils.SendBadRequest(c, _err.Error())
	}
}
