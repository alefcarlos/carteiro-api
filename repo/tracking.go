package repo

import (
	"github.com/alefcarlos/carteiro-api/models"
)

//Trackings lista de itens a serem monitorados
var Trackings = []models.TrackingInfo{
	models.TrackingInfo{ID: 1, TrackingCode: "TESTE", LastType: "BDI", LastStatus: 1, LastDescription: "Em Rota"},
}
