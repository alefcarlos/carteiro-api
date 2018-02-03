package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//APIResultError modelo padrão para respostas de erro da API
type APIResultError struct {
	Message string `json:"message"`
}

//APIResultOk modelo padrão para reposta de sucesso da API
type APIResultOk struct {
	Result interface{} `json:"result"`
}

//SendSuccess auxiliar para respostas de sucesso
func SendSuccess(c *gin.Context, result interface{}) {
	_result := APIResultOk{
		Result: result,
	}

	c.JSON(http.StatusOK, _result)
}

//SendBadRequest auxiliar para respostas de erro
func SendBadRequest(c *gin.Context, message string) {
	_result := APIResultError{
		Message: message,
	}

	c.JSON(http.StatusBadRequest, _result)
}

//SendConflict axuliar para respotas de conflito
func SendConflict(c *gin.Context, message string) {
	_result := APIResultError{
		Message: message,
	}

	c.JSON(http.StatusConflict, _result)
}
