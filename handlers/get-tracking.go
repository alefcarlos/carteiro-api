package handlers

import (
	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/gin-gonic/gin"
)

//GetTrakings retorna a lista de rastreios que têm novas informações
func GetTrakings(c *gin.Context) {
	utils.SendSuccess(c, repo.Trackings)
}
