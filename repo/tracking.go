package repo

import (
	"github.com/alefcarlos/carteiro-api/errors"
	"github.com/alefcarlos/carteiro-api/models"
)

//Trackings lista de itens a serem monitorados
var trackings = []models.TrackingInfo{}

//AddTracking Adicionar um novo registro para monitoramento
//Retorna o item adicionado, só que agora com o ID
func AddTracking(item models.TrackingInfo) (models.TrackingInfo, error) {
	//Verificar se já existe códio registrado
	_alreadyExists := Filter(trackings, func(t models.TrackingInfo) bool {
		return item.TrackingCode == t.TrackingCode
	})

	//Verificar se já não exste o mesmo elemento
	if len(_alreadyExists) > 0 {
		return models.TrackingInfo{}, errors.ErrTrackingCodeDuplicated
		// return nil, errors.ErrTrackingCodeDuplicated
	}

	//Setar informação de identificação
	item.ID = len(trackings) + 1
	_result := append(trackings, item)
	trackings = _result

	return item, nil
}

//GetTrackings Obtém todos os monitoramentos que ainda não foram lidos
func GetTrackings() []models.TrackingInfo {
	//Filtar somente o que ainda não foram entregues
	_notRead := Filter(trackings, func(t models.TrackingInfo) bool {
		return !t.IsSeen
	})

	return _notRead
}

//UpdateTrackingRead atualiza um registro como Lido
func UpdateTrackingRead(id int) error {
	_trackingIndex := Index(trackings, func(t models.TrackingInfo) bool {
		return t.ID == id
	})

	if _trackingIndex >= 0 {
		trackings[_trackingIndex].IsSeen = true
		return nil
	}

	return errors.ErrIDNotFound
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

//Index retorna o index do item de acordo com fn()
func Index(s []models.TrackingInfo, fn func(models.TrackingInfo) bool) int {
	index := -1

	for i, v := range s {
		if fn(v) {
			index = i
			break
		}
	}

	return index
}
