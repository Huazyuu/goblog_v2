package ThirdApiService

import (
	"backend/global"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

type GaodeIPResponse struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Adcode    string `json:"adcode"`
	Rectangle string `json:"rectangle"`
}
type GaodeWeatherResponse struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Lives    []struct {
		Province         string `json:"province"`
		City             string `json:"city"`
		Adcode           string `json:"adcode"`
		Weather          string `json:"weather"`
		Temperature      string `json:"temperature"`
		Winddirection    string `json:"winddirection"`
		Windpower        string `json:"windpower"`
		Humidity         string `json:"humidity"`
		Reporttime       string `json:"reporttime"`
		TemperatureFloat string `json:"temperature_float"`
		HumidityFloat    string `json:"humidity_float"`
	} `json:"lives"`
}

type GaodeService struct {
}

func (g GaodeService) GaodeWeatherService(adcode string) (resp GaodeWeatherResponse, err error) {
	response, err := http.Get(fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s", global.Config.Gaode.Key, adcode))
	if err != nil {
		logrus.Errorf(err.Error())
		return resp, errors.New("获取天气失败")
	}
	if response.Status == strconv.Itoa(http.StatusNotFound) {
		return resp, errors.New("获取天气失败,地区adcode输入错误")
	}
	byteData, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf(err.Error())
		return resp, err
	}
	err = json.Unmarshal(byteData, &resp)
	if err != nil {
		logrus.Errorf(err.Error())
		return resp, errors.New("解析天气失败")
	}
	return resp, nil
}

func (g GaodeService) GetPositionFromIP(ip string) (resp GaodeIPResponse, err error) {
	response, err := http.Get(fmt.Sprintf("https://restapi.amap.com/v3/ip?key=%s&ip=%s", global.Config.Gaode.Key, ip))
	if err != nil {
		logrus.Errorf(err.Error())
		return resp, errors.New("获取定位失败")
	}
	byteData, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf(err.Error())
		return resp, err
	}
	err = json.Unmarshal(byteData, &resp)
	if err != nil {
		logrus.Errorf(err.Error())
		return resp, errors.New("解析定位失败")
	}
	return resp, nil
}
