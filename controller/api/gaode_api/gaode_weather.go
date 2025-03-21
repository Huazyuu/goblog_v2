package gaode_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/service/ThirdApiService"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

func (GaodeApi) WeatherInfoByIPView(c *gin.Context) {
	var cr req.GaodeWeatherRequest
	_ = c.ShouldBindQuery(&cr)

	if global.Config.Gaode.Enable {
		var gaodeService ThirdApiService.GaodeService

		var ip string
		if cr.IP == "" {
			// 没传就用本地请求的ip
			ip = c.ClientIP()
		}
		ip = cr.IP
		if !utils.IsPublicIPAddr(ip) {
			ip = "114.247.50.2"
		}

		Position, err := gaodeService.GetPositionFromIP(ip)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}

		cityCode := Position.Adcode

		Weather, err := gaodeService.GaodeWeatherService(cityCode)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}

		if len(Weather.Lives) > 0 {
			res.OkWithData(Weather.Lives[0], c)
			return
		} else {
			res.FailWithMessage("请输入正确ip", c)
		}
		return
	}
	res.FailWithMessage("接口未开放", c)

}
func (GaodeApi) WeatherInfoByAdcodeView(c *gin.Context) {
	var cr req.GaodeWeatherRequest
	_ = c.ShouldBindQuery(&cr)

	if global.Config.Gaode.Enable {
		var gaodeService ThirdApiService.GaodeService

		cityCode := cr.Adcode
		global.Log.Info("adcode:", cityCode)
		Weather, err := gaodeService.GaodeWeatherService(cityCode)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}

		if len(Weather.Lives) > 0 {
			res.OkWithData(Weather.Lives[0], c)
			return
		} else {
			res.FailWithMessage("请输入当地adcode", c)
		}
		return
	}
	res.FailWithMessage("接口未开放", c)
}
