package repo

import (
	"github.com/alefcarlos/carteiro-api/models"
)

//Trackings lista de itens a serem monitorados
var Trackings = []models.TrackingInfo{}

//AddTracking Adicionar um novo registro para monitoramento
func AddTracking(item models.TrackingInfo) []models.TrackingInfo {
	_result := append(Trackings, item)
	Trackings = _result
	return Trackings
}
