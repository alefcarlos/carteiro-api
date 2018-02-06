package handlers

import (
	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"
	"github.com/gin-gonic/gin/binding"

	"github.com/alefcarlos/carteiro-api/models"
	"github.com/gin-gonic/gin"
)

//PostNewSubscribe permite adicionar um novo rastreio para monitoramento
func PostNewSubscribe(c *gin.Context) {
	var tracking models.TrackingInfo

	//Obter o objeto a partir do body
	if err := c.ShouldBindWith(&tracking, binding.JSON); err == nil {
		//Adicionar novo objeto na listagem
		if itemAdded, err := repo.AddTracking(tracking); err != nil {
			utils.SendConflict(c, err.Error())
		} else {
			utils.SendSuccess(c, gin.H{"ID": itemAdded.ID})
		}

	} else {
		utils.SendBadRequest(c, err.Error())
	}
}
