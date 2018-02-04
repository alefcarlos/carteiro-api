package models

//TrackingUpdateInfo modelo exclusivo para atualização das informações de rastreio
type TrackingUpdateInfo struct {
	LastType        string `json:"type"`
	LastStatus      int    `json:"status"`
	LastDescription string `json:"description"`
}
