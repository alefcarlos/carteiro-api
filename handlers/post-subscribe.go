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
	if _err := c.ShouldBindWith(&tracking, binding.JSON); _err == nil {
		//Adicionar novo objeto na listagem
		if _itemAdded, _err := repo.AddTracking(tracking); _err != nil {
			utils.SendConflict(c, _err.Error())
		} else {
			utils.SendSuccess(c, gin.H{"ID": _itemAdded.ID})
		}

	} else {
		utils.SendBadRequest(c, _err.Error())
	}
}
