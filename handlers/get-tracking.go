package handlers

import (
	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/gin-gonic/gin"
)

//GetAvailableTrackings retorna a lista de rastreios que têm novas informações
func GetAvailableTrackings(c *gin.Context) {
	utils.SendSuccess(c, repo.GetTrackings())
}
