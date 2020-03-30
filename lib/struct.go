package lib

type InfoVersion struct {
	Uuid string `json:"uuid" binding:"required"`
}

type InfoVersionResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Code    string             `json:"code"`
	Result  *ResultInfoVersion `json:"result"`
}

type ResultInfoVersion struct {
	Version          string `json:"versi"`
	Operating_system string `json:"os"`
}
