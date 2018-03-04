package repo

import (
	"log"

	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/apiModels"
	"github.com/alefcarlos/carteiro-api/src/internal/errors"
	"github.com/alefcarlos/carteiro-api/src/internal/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AddTracking Adicionar um novo registro para monitoramento
//Retorna o item adicionado, só que agora com o ID
func AddTracking(entity apiModels.NewSubscription) (*models.Tracking, error) {
	session := GetMongoSession()
	defer session.Close()
	collection := getSubscriptionCollection(session)

	//Validar se já existe essa inscrição
	sub, err := GetSubscription(entity.Address)
	if err != nil && err != mgo.ErrNotFound {
		log.Printf("Erro ao tentar obter a inscrição %s", err.Error())
		return nil, err
	}

	//Verificar se o código de rastreio já existe
	indexTracking := IndexTracking(sub.Trackings, func(tracking models.Tracking) bool {
		return tracking.TrackingCode == entity.Code
	})

	//Se existir, retorna erro de duplicidade
	if indexTracking > -1 {
		log.Printf("Houve tentativa de inserir novo tracking repetido %s\n", entity.Code)
		return nil, errors.ErrTrackingCodeDuplicated
	}

	//Não existe então adicionar
	newTracking := models.Tracking{
		TrackingCode: entity.Code,
	}

	//Adicionar um item no array
	if err = addTracking(collection, sub.ID, &newTracking); err != nil {
		return nil, err
	}

	return &newTracking, nil
}

//GetTrackingsToNotify lista todos os rastreios que devem ser notificados aos usuários
func GetTrackingsToNotify() (result []*models.SubscribeInfo, err error) {
	session := GetMongoSession()
	defer session.Close()
	collection := getSubscriptionCollection(session)

	err = collection.Find(bson.M{"trackings.mustNotify": true}).All(&result)
	if err != nil {
		log.Printf("Não foi possível obter todas as trackings a serem notificadas%s", err.Error())
	}

	return
}

//SetTrackingsRead atualiza todos os registros de tracking para lido
func SetTrackingsRead() (changeInfo *mgo.ChangeInfo, err error) {
	session := GetMongoSession()
	defer session.Close()
	collection := getSubscriptionCollection(session)

	query := bson.M{"trackings.mustNotify": true}
	change := bson.M{"$set": bson.M{"trackings.$.isRead": true, "trackings.$.mustNotify": false}}
	changeInfo, err = collection.UpdateAll(query, change)

	return
}

//UpdateTrackingInfo atualiza as informações de rastreio de um registro
func UpdateTrackingInfo(info apiModels.TrackingUpdateInfo) (err error) {
	session := GetMongoSession()
	defer session.Close()
	collection := getSubscriptionCollection(session)

	//Validar se já existe essa inscrição
	sub, err := GetSubscription(info.Address)
	if err != nil && err != mgo.ErrNotFound {
		log.Printf("Erro ao tentar obter a inscrição %s", err.Error())
		return
	}

	trackingIndex := IndexTracking(sub.Trackings, func(t models.Tracking) bool {
		return t.TrackingCode == info.Code
	})

	if trackingIndex == -1 {
		err = errors.ErrIDNotFound
	}

	// 	trackings[trackingIndex].LastDescription = info.LastDescription
	// 	trackings[trackingIndex].LastStatus = info.LastStatus
	// 	trackings[trackingIndex].LastType = info.LastType
	// 	trackings[trackingIndex].MustNotify = true
	// 	trackings[trackingIndex].IsRead = false
	// 	trackings[trackingIndex].LastDestination = info.LastDestination
	// 	return nil
	// }
	query := bson.M{"trackings.code": info.Code}
	change := bson.M{"$set": bson.M{"trackings.$.isRead": false, "trackings.$.mustNotify": true}}
	err = collection.Update(query, change)

	return
}

// FilterTracking returns a new slice holding only
// the elements of s that satisfy fn()
func FilterTracking(s []models.Tracking, fn func(models.Tracking) bool) []*models.Tracking {
	var p []*models.Tracking // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, &v)
		}
	}

	return p
}

//IndexTracking retorna o index do item de acordo com fn()
func IndexTracking(s []models.Tracking, fn func(models.Tracking) bool) int {
	index := -1

	for i, v := range s {
		if fn(v) {
			index = i
			break
		}
	}

	return index
}

func addTracking(collection *mgo.Collection, subscriptionID bson.ObjectId, entity *models.Tracking) (err error) {
	change := bson.M{"$push": bson.M{"trackings": entity}}
	err = collection.Update(bson.M{"_id": subscriptionID}, change)
	return
}
