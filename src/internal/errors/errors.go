package errors

import "errors"

/*
ErrTrackingCodeDuplicated é um erro que é apresentado quando o tracking code já existe
*/
var ErrTrackingCodeDuplicated = errors.New("Esse número de rastreio já têm uma notificação ativa")

//ErrIDNotFound é um erro que é apresentando quando o ID passado não existe
var ErrIDNotFound = errors.New("O ID informado não foi encontrado")
