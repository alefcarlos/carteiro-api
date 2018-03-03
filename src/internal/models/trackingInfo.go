package models

import "gopkg.in/mgo.v2/bson"

//SubscribeInfo Informações de notificação de um clinte
type SubscribeInfo struct {
	ID bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`

	//Address é a informação vinda do BotFramework que representa a identificação do usuário numa determinada conversão e num determinado canal
	Address BotFrameworkAddressInfo `json:"address" bson:"address"`

	//Trackings lista de rastreios que devemos monitorar
	Trackings []Tracking `json:"trackings" bson:"trackings"`
}

//Tracking Informação de um rastreio
type Tracking struct {
	TrackingCode    string `json:"code" binding:"required" bson:"code"`
	LastType        string `json:"type" bson:"type"`
	LastStatus      int    `json:"status" bson:"status"`
	LastDescription string `json:"description" bson:"description"`
	LastDestination string `json:"destination" bson:"destination"`

	//MustNotify indica se devemos notficar o usuário
	MustNotify bool `json:"mustNotify" bson:"mustNotify"`

	//IsRead indica se o usuário já leu essa notificação
	IsRead bool `json:"isRead" bson:"isRead"`
}

//IsDelivered valida se um rastreio já foi entregue
func (item *Tracking) IsDelivered() bool {
	return (item.LastType == "BDI" || item.LastType == "BDE" || item.LastType == "BDR") && (item.LastStatus == 0 || item.LastStatus == 1)
}
