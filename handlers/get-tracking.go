package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/models"
	"github.com/gin-gonic/gin"
)

//GetTrakings retorna a lista de rastreios que têm novas informações
func GetTrakings(c *gin.Context) {
	t := models.TrackingInfo{}
	t.TrackingCode = "teste"

	c.JSON(http.StatusOK, t)

	// if err := c.ShouldBindWith(&t, binding.JSON); err == nil {
	// 	c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }
}
