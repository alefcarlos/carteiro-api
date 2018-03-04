package repo

import (
	"log"

	"github.com/alefcarlos/carteiro-api/src/internal/errors"
	"github.com/alefcarlos/carteiro-api/src/internal/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GetSubscription obtém uma inscrição a partir do endereço do cliente
func GetSubscription(address models.BotFrameworkAddressInfo) (subscription *models.SubscribeInfo, err error) {
	session := GetMongoSession()
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
	session := GetMongoSession()
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
	session := GetMongoSession()
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
