package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// func GetStatus(c *gin.Context) {
// 	utils.SendSuccess(c, "tudo na paz, irmão!")
// }

//GetStatus retorna a lista de rastreios que têm novas informações
func GetStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
