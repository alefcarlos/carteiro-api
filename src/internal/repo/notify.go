package repo

import (
	"log"

	"github.com/alefcarlos/carteiro-api/src/internal/models"

	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/apiModels"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AddNotifyAll adiciona uma mensagem simples para todos que têm registro
func AddNotifyAll(message *apiModels.NotifyMessage) (err error) {
	session := GetMongoSession()
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
	session := GetMongoSession()
	defer session.Close()
	collection := getNotifyCollection(session)

	err = collection.Find(bson.M{"isRead": false}).All(&result)
	if err != nil {
		log.Printf("Não foi possível obter todas as notificações %s", err.Error())
	}
	return
}

//PutAllNotificationsRead atualiza todos os registros de notificações para lido
func PutAllNotificationsRead() (changeInfo *mgo.ChangeInfo, err error) {
	session := GetMongoSession()
	defer session.Close()
	collection := getNotifyCollection(session)

	query := bson.M{}
	change := bson.M{"$set": bson.M{"isRead": true}}
	changeInfo, err = collection.UpdateAll(query, change)

	return
}

func getNotifyCollection(session *mgo.Session) *mgo.Collection {
	return session.DB("carteiro-db").C("notify")
}
