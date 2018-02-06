package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin/binding"

	"github.com/alefcarlos/carteiro-api/errors"
	"github.com/alefcarlos/carteiro-api/models"
	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/gin-gonic/gin"
)

//PutTrackingInfo atualiza as informações de status de um determinado tracking
func PutTrackingInfo(c *gin.Context) {
	info := models.TrackingUpdateInfo{}

	//Obtém o registro através do id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.SendBadRequest(c, "Informe um Id válido.")
		return
	}

	//Tentar realizar o bind do JSON
	if err := c.ShouldBindWith(&info, binding.JSON); err != nil {
		utils.SendBadRequest(c, err.Error())
		return
	}

	if err = repo.UpdateTrackingInfo(id, info); err != nil {
		switch err {
		case errors.ErrIDNotFound:
			utils.SendNotFound(c, err.Error())
		default:
			utils.SendBadRequest(c, err.Error())
		}

	} else {
		utils.SendSuccess(c, "Informações atualizadas com sucesso.")
	}
}
