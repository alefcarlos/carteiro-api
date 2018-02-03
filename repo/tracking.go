package repo

import (
	"github.com/alefcarlos/carteiro-api/errors"
	"github.com/alefcarlos/carteiro-api/models"
)

//Trackings lista de itens a serem monitorados
var Trackings = []models.TrackingInfo{}

//AddTracking Adicionar um novo registro para monitoramento
func AddTracking(item models.TrackingInfo) error {
	//Verificar se já existe códio registrado
	_alreadyExists := Filter(Trackings, func(t models.TrackingInfo) bool {
		return item.TrackingCode == t.TrackingCode
	})

	//Verificar se já não exste o mesmo elemento
	if len(_alreadyExists) > 0 {
		return errors.ErrTrackingCodeDuplicated
	}

	//Setar informação de identificação
	item.ID = len(Trackings) + 1
	_result := append(Trackings, item)
	Trackings = _result

	return nil
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
