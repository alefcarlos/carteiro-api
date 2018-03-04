package apiModels

import "github.com/alefcarlos/carteiro-api/src/internal/models"

//TrackingUpdateInfo modelo exclusivo para atualização das informações de rastreio
type TrackingUpdateInfo struct {
	//Address é a informação vinda do BotFramework que representa a identificação do usuário numa determinada conversão e num determinado canal
	Address models.BotFrameworkAddressInfo `json:"address"`

	Code            string `json:"code"`
	LastType        string `json:"type"`
	LastStatus      int    `json:"status"`
	LastDescription string `json:"description"`
	LastDestination string `json:"destination"`
}
