package handlers

import (
	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/gin-gonic/gin"
)

//NotImplementedYet apenas para testes
func NotImplementedYet(c *gin.Context) {
	utils.SendBadRequest(c, "Calma, gafanhoto, esse método ainda não foi implementado!")
}
