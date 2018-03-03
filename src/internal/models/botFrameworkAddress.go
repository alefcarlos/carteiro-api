package models

//BotFrameworkAddressInfo informações do endereço do usuário vindo do bot https://docs.botframework.com/en-us/node/builder/chat-reference/interfaces/_botbuilder_d_.iaddress.html
type BotFrameworkAddressInfo struct {
	Bot          BotFrameworkIdentityInfo `json:"bot" bson:"bot"`
	Conversation BotFrameworkIdentityInfo `json:"conversation" bson:"conversation"`
	ChannelID    string                   `json:"channelId" binding:"required" bson:"channelId"`
	User         BotFrameworkIdentityInfo `json:"user" bson:"user"`
	ServiceURL   string                   `json:"serviceUrl" binding:"required" bson:"serviceUrl"`
}

//BotFrameworkIdentityInfo informações da identificação do BotFramework https://docs.botframework.com/en-us/node/builder/chat-reference/interfaces/_botbuilder_d_.iidentity.html
type BotFrameworkIdentityInfo struct {
	ID      string `json:"id" binding:"required" bson:"id"`
	IsGroup bool   `json:"isGroup" bson:"isGroup"`
	Name    string `json:"name" binding:"required" bson:"name"`
}
