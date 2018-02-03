package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/models"

	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/gin-gonic/gin"
)

//GetTrakings retorna a lista de rastreios que têm novas informações
func GetTrakings(c *gin.Context) {
	_result := models.APIResultOk{
		Result: repo.Trackings,
	}

	c.JSON(http.StatusOK, _result)
}
