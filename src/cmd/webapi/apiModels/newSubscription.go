package apiModels

import "github.com/alefcarlos/carteiro-api/src/internal/models"

//NewSubscription Modelo utilizado na requisição POST /subscribe
type NewSubscription struct {
	//Address é a informação vinda do BotFramework que representa a identificação do usuário numa determinada conversão e num determinado canal
	Address models.BotFrameworkAddressInfo `json:"address"`

	Code string `json:"code"`
}
