package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/models"
	"github.com/gin-gonic/gin"
)

//NotImplementedYet apenas para testes
func NotImplementedYet(c *gin.Context) {
	_result := models.APIResultError{
		Message: "Calma, gafanhoto, esse método ainda não foi implementado!",
	}

	c.JSON(http.StatusBadRequest, _result)
}
