package errors

import "errors"

/*
ErrTrackingCodeDuplicated é um erro que é apresentado quando o tracking code já existe
*/
var ErrTrackingCodeDuplicated = errors.New("Esse número de rastreio já têm uma notificação ativa")
