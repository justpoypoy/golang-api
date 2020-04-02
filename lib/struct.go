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

type LogJSON struct {
	ClientIP   string  `json:"clientip"`
	TimeStamp  string  `json:"timestamp"`
	Method     string  `json:"method"`
	Path       string  `json:"path"`
	StatusCode int     `json:"statuscode"`
	Latency    float64 `json:"latency"`
	Request    string  `json:"request"`
	Response   string  `json:"response"`
}
