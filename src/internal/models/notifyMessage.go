package models

//NotifyModel informações de notificação
type NotifyModel struct {

	//Address é a informação vinda do BotFramework que representa a identificação do usuário numa determinada conversão e num determinado canal
	Address BotFrameworkAddressInfo `json:"address" bson:"address"`

	Message string `json:"message" bson:"message"`

	IsRead bool `json:"isRead" bson:"isRead"`
}
