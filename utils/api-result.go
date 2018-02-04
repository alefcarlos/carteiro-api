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

//SendSuccess auxiliar para respostas de 200 sucesso
func SendSuccess(c *gin.Context, result interface{}) {
	_result := APIResultOk{
		Result: result,
	}

	c.JSON(http.StatusOK, _result)
}

//SendBadRequest auxiliar para respostas de 400 erro
func SendBadRequest(c *gin.Context, message string) {
	_result := APIResultError{
		Message: message,
	}

	c.JSON(http.StatusBadRequest, _result)
}

//SendConflict axuliar para respotas de 409 conflito
func SendConflict(c *gin.Context, message string) {
	_result := APIResultError{
		Message: message,
	}

	c.JSON(http.StatusConflict, _result)
}

//SendNotFound axuliar para respotas 404 - NotFound
func SendNotFound(c *gin.Context, message string) {
	_result := APIResultError{
		Message: message,
	}

	c.JSON(http.StatusConflict, _result)
}
