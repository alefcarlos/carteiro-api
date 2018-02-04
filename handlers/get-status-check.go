package handlers

import (
	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/gin-gonic/gin"
)

//GetStatus retorna a lista de rastreios que têm novas informações
func GetStatus(c *gin.Context) {
	utils.SendSuccess(c, "tudo na paz, irmão!")
}

// GetStatusGinEngine is gin router.
func GetStatusGinEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/status", GetStatus)

	return r
}
