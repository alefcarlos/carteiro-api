package apiModels

//TrackingUpdateInfo modelo exclusivo para atualização das informações de rastreio
type TrackingUpdateInfo struct {
	LastType        string `json:"type"`
	LastStatus      int    `json:"status"`
	LastDescription string `json:"description"`
	LastDestination string `json:"destination"`
}
