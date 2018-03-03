package apiModels

//NotifyMessage mensagem que devemos notificar a todos os usuários que têm
//notificação inscrita
type NotifyMessage struct {
	Message string `json:"message"`
}
