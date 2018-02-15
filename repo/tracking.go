package repo

import (
	"log"

	"github.com/alefcarlos/carteiro-api/apiModels"
	"github.com/alefcarlos/carteiro-api/errors"
	"github.com/alefcarlos/carteiro-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AddTracking Adicionar um novo registro para monitoramento
//Retorna o item adicionado, só que agora com o ID
func AddTracking(entity apiModels.NewSubscription) (*models.Tracking, error) {
	session := MongoSession.Copy()
	defer session.Close()
	collection := session.DB("carteiro-db").C("subscriptions")

	//Validar se já existe essa inscrição
	sub, err := GetSubscription(entity.Address)
	if err != nil && err != mgo.ErrNotFound {
		log.Printf("Erro ao tentar obter a inscrição %s", err.Error())
	}

	//Se não existir a inscrição, então devemos inserir e atualiza variável indexSub
	if err == mgo.ErrNotFound {
		sub = models.SubscribeInfo{
			Address: entity.Address,
		}

		if err = AddSubscription(sub); err != nil {
			return nil, err
		}
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

func addTracking(collection *mgo.Collection, subscriptionID bson.ObjectId, entity *models.Tracking) (err error) {
	change := bson.M{"$push": bson.M{"trackings": entity}}
	err = collection.Update(bson.M{"_id": subscriptionID}, change)
	return
}
