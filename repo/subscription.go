package repo

import (
	"log"

	"github.com/alefcarlos/carteiro-api/errors"
	"github.com/alefcarlos/carteiro-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GetSubscription obtém uma inscrição a partir do endereço do cliente
func GetSubscription(address models.BotFrameworkAddressInfo) (subscription *models.SubscribeInfo, err error) {
	session := MongoSession.Copy()
	defer session.Close()
	collection := getSubscriptionCollection(session)

	err = collection.Find(bson.M{"address.channelId": address.ChannelID, "address.conversation.id": address.Conversation.ID, "address.user.id": address.User.ID}).One(&subscription)
	if err != nil {
		log.Printf("Não foi possível obter todas as inscrições %s", err.Error())
	}
	return
}

//GetSubscriptions Obtém todos os monitoramentos
func GetSubscriptions() (result []models.SubscribeInfo, err error) {
	session := MongoSession.Copy()
	defer session.Close()
	collection := getSubscriptionCollection(session)

	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Printf("Não foi possível obter todas as inscrições %s", err.Error())
	}
	return
}

//AddSubscription Adiciona uma nova inscrição de monitoramento
func AddSubscription(entity models.SubscribeInfo) (model *models.SubscribeInfo, err error) {
	session := MongoSession.Copy()
	defer session.Close()
	collection := getSubscriptionCollection(session)

	subscription, err := GetSubscription(entity.Address)

	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}

	if subscription != nil {
		return nil, errors.ErrTrackingCodeDuplicated
	}

	err = collection.Insert(entity)
	if err != nil {
		log.Printf("Não foi possível inserir um nova inscrição %s", err.Error())
	}
	model = &entity
	return
}

func getSubscriptionCollection(session *mgo.Session) *mgo.Collection {
	return session.DB("carteiro-db").C("subscriptions")
}

// //GetTrackingsToNotify lista todos os rastreios que devem ser notificados aos usuários
// func GetTrackingsToNotify() []models.TrackingInfo {
// 	//Filtar somente o que ainda não foram entregues
// 	_found := Filter(trackings, func(t models.TrackingInfo) bool {
// 		return t.MustNotify
// 	})

// 	return _found
// }

// //UpdateTrackingRead atualiza um registro como Lido
// func UpdateTrackingRead(id int) error {
// 	trackingIndex := Index(trackings, func(t models.TrackingInfo) bool {
// 		return t.ID == id
// 	})

// 	if trackingIndex >= 0 {
// 		trackings[trackingIndex].IsRead = true
// 		trackings[trackingIndex].MustNotify = false
// 		return nil
// 	}

// 	return errors.ErrIDNotFound
// }

// //UpdateTrackingInfo atualiza as informações de rastreio de um registro
// func UpdateTrackingInfo(id int, info models.TrackingUpdateInfo) error {

// 	trackingIndex := Index(trackings, func(t models.TrackingInfo) bool {
// 		return t.ID == id
// 	})

// 	if trackingIndex == -1 {
// 		return errors.ErrIDNotFound
// 	}

// 	trackings[trackingIndex].LastDescription = info.LastDescription
// 	trackings[trackingIndex].LastStatus = info.LastStatus
// 	trackings[trackingIndex].LastType = info.LastType
// 	trackings[trackingIndex].MustNotify = true
// 	trackings[trackingIndex].IsRead = false
// 	trackings[trackingIndex].LastDestination = info.LastDestination
// 	return nil
// }

// // Filter returns a new slice holding only
// // the elements of s that satisfy fn()
// func Filter(s []models.TrackingInfo, fn func(models.TrackingInfo) bool) []models.TrackingInfo {
// 	var p []models.TrackingInfo // == nil
// 	for _, v := range s {
// 		if fn(v) {
// 			p = append(p, v)
// 		}
// 	}
// 	return p
// }

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

//IndexSubscription retorna o index do item de acordo com fn()
func IndexSubscription(s []models.SubscribeInfo, fn func(models.SubscribeInfo) bool) int {
	index := -1

	for i, v := range s {
		if fn(v) {
			index = i
			break
		}
	}

	return index
}
