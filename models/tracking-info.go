package models

//TrackingInfo Informações do rastreio
type TrackingInfo struct {
	ID              int    `json:"id"`
	TrackingCode    string `json:"code" binding:"required"`
	LastType        string `json:"type"`
	LastStatus      int    `json:"status"`
	LastDescription string `json:"description"`

	//Address é a informação vinda do BotFramework que representa a identificação do usuário numa determinada conversão e num determinado canal
	Address BotFrameworkAddressInfo `json:"address"`

	IsSeen bool `json:"isSeen"`
}

//BotFrameworkAddressInfo informações do endereço do usuário vindo do bot https://docs.botframework.com/en-us/node/builder/chat-reference/interfaces/_botbuilder_d_.iaddress.html
type BotFrameworkAddressInfo struct {
	Bot        BotFrameworkIdentityInfo `json:"bot"`
	ChannelID  string                   `json:"channelId" binding:"required"`
	User       BotFrameworkIdentityInfo `json:"user"`
	ServiceURL string                   `json:"serviceUrl" binding:"required"`
}

//BotFrameworkIdentityInfo informações da identificação do BotFramework https://docs.botframework.com/en-us/node/builder/chat-reference/interfaces/_botbuilder_d_.iidentity.html
type BotFrameworkIdentityInfo struct {
	ID      string `json:"id" binding:"required"`
	IsGroup bool   `json:"isGruop"`
	Name    string `json:"name" binding:"required"`
}

//IsDelivered valida se o rastreio já foi entregue
func (item *TrackingInfo) IsDelivered() bool {
	return (item.LastType == "BDI" || item.LastType == "BDE" || item.LastType == "BDR") && (item.LastStatus == 0 || item.LastStatus == 1)
}
