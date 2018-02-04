package handlers

import (
	"strconv"

	"github.com/alefcarlos/carteiro-api/errors"
	"github.com/alefcarlos/carteiro-api/models"
	"github.com/alefcarlos/carteiro-api/repo"
	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/gin-gonic/gin"
)

//PutTrackingInfo atualiza as informações de status de um determinado tracking
func PutTrackingInfo(c *gin.Context) {
	_info := models.TrackingUpdateInfo{}

	//Obtém o registro através do id
	_id, _err := strconv.Atoi(c.Param("id"))

	if _err != nil {
		utils.SendBadRequest(c, "Informe um Id válido.")
		return
	}

	//Tentar realizar o bind do JSON
	if _err := c.ShouldBindJSON(&_info); _err != nil {
		utils.SendBadRequest(c, _err.Error())
		return
	}

	if _err = repo.UpdateTrackingInfo(_id, _info); _err != nil {
		switch _err {
		case errors.ErrIDNotFound:
			utils.SendNotFound(c, _err.Error())
		default:
			utils.SendBadRequest(c, _err.Error())
		}

	} else {
		utils.SendSuccess(c, "Informações atualizadas com sucesso.")
	}
}
