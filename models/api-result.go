package models

//APIResultError modelo padrão para respostas de erro da API
type APIResultError struct {
	Message string `json:"message"`
}

//APIResultOk modelo padrão para reposta de sucesso da API
type APIResultOk struct {
	Result interface{} `json:"result"`
}
