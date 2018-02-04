package repo

import (
	"github.com/alefcarlos/carteiro-api/errors"
	"github.com/alefcarlos/carteiro-api/models"
)

//Trackings lista de itens a serem monitorados
var trackings = []models.TrackingInfo{
	models.TrackingInfo{ID: 1, LastStatus: 2, LastType: "a"},
	models.TrackingInfo{ID: 1, LastStatus: 1, LastType: "BDI"},
}

//AddTracking Adicionar um novo registro para monitoramento
func AddTracking(item models.TrackingInfo) error {
	//Verificar se já existe códio registrado
	_alreadyExists := Filter(trackings, func(t models.TrackingInfo) bool {
		return item.TrackingCode == t.TrackingCode
	})

	//Verificar se já não exste o mesmo elemento
	if len(_alreadyExists) > 0 {
		return errors.ErrTrackingCodeDuplicated
	}

	//Setar informação de identificação
	item.ID = len(trackings) + 1
	_result := append(trackings, item)
	trackings = _result

	return nil
}

//GetTrackings Obtém todos os rastreios que ainda não foram entegues
func GetTrackings() []models.TrackingInfo {
	//Filtar somente o que ainda não foram entregues
	_notDelivered := Filter(trackings, func(t models.TrackingInfo) bool {
		return !t.IsDelivered()
	})

	return _notDelivered
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []models.TrackingInfo, fn func(models.TrackingInfo) bool) []models.TrackingInfo {
	var p []models.TrackingInfo // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
