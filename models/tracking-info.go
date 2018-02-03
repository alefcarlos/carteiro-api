package models

//TrackingInfo Informações do rastreio
type TrackingInfo struct {
	ID              int    `json:"id"`
	TrackingCode    string `json:"code" binding:"required"`
	LastType        string `json:"type"`
	LastStatus      int    `json:"status"`
	LastDescription string `json:"description"`

	//Address é a informação vinda do BotFramework que representa a identificação do usuário
	Address BotFrameworkAddressInfo `json:"address" binding:"required"`
}

//BotFrameworkAddressInfo informações do endereço do usuário vindo do bot https://docs.botframework.com/en-us/node/builder/chat-reference/interfaces/_botbuilder_d_.iaddress.html
type BotFrameworkAddressInfo struct {
	Bot          BotFrameworkIdentityInfo `json:"bot"`
	ChannelID    string                   `json:"channelId"`
	Conversation BotFrameworkIdentityInfo `json:"conversation"`
	User         BotFrameworkIdentityInfo `json:"user"`
	ServiceURL   string                   `json:"serviceUrl"`
}

//BotFrameworkIdentityInfo informações da identificação do BotFramework https://docs.botframework.com/en-us/node/builder/chat-reference/interfaces/_botbuilder_d_.iidentity.html
type BotFrameworkIdentityInfo struct {
	ID      string `json:"id"`
	IsGroup bool   `json:"isGruop"`
	Name    string `json:"name"`
}
