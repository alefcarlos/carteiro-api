package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/models"
	"github.com/gin-gonic/gin"
)

//PostNewSubscribe permite adicionar um novo rastreio para monitoramento
func PostNewSubscribe(c *gin.Context) {
	var tracking models.TrackingInfo

	//Obter o objeto a partir do body
	if _err := c.ShouldBindJSON(&tracking); _err == nil {
		//Adicionar novo objeto na listagem
		c.JSON(http.StatusBadRequest, gin.H{"message": _err.Error()})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": _err.Error()})
	}
}
