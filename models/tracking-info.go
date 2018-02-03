package models

//TrackingInfo Informações do rastreio
type TrackingInfo struct {
	ID              int    `json:"id"`
	TrackingCode    string `json:"code" binding:"required"`
	LastType        string `json:"type"`
	LastStatus      int    `json:"status"`
	LastDescription string `json:"description"`
}
