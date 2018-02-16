package repo

import (
	"log"

	"github.com/alefcarlos/carteiro-api/models"

	"github.com/alefcarlos/carteiro-api/apiModels"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AddNotifyAll adiciona uma mensagem simples para todos que têm registro
func AddNotifyAll(message *apiModels.NotifyMessage) (err error) {
	session := MongoSession.Copy()
	defer session.Close()
	collection := getNotifyCollection(session)

	subs, err := GetSubscriptions()

	if err != nil {
		return
	}

	for _, value := range subs {
		entity := models.NotifyModel{
			Address: value.Address,
			Message: message.Message,
		}

		err = collection.Insert(entity)
		if err != nil {
			log.Printf("Não foi possível inserir uma nova mensagem de notificação %s", err.Error())
		}
	}

	err = nil
	return
}

//GetNotifications Obtém todas as notificações que não foram lidas
func GetNotifications() (result []models.NotifyModel, err error) {
	session := MongoSession.Copy()
	defer session.Close()
	collection := getNotifyCollection(session)

	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Printf("Não foi possível obter todas as notificações %s", err.Error())
	}
	return
}

func getNotifyCollection(session *mgo.Session) *mgo.Collection {
	return session.DB("carteiro-db").C("notify")
}
