package req

type GaodeWeatherRequest struct {
	IP     string `json:"ip" form:"ip"`
	Adcode string `json:"adcode" form:"adcode"`
}
