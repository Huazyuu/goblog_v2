# 项目需要调用的高德api

> [概述-Web服务 API|高德地图API](https://lbs.amap.com/api/webservice/summary)

IP://restapi.amap.com/v3/ip?key=您的key&ip=114.247.50.2

```json
{
  "status": "1",
  "info": "OK",
  "infocode": "10000",
  "province": "北京市",
  "city": "北京市",
  "adcode": "110000",
  "rectangle": "116.0119343,39.66127144;116.7829835,40.2164962"
}
```

Weather://restapi.amap.com/v3/weather/weatherInfo?key=您的 key&city=110101

```json
{
  "status": "1",
  "count": "1",
  "info": "OK",
  "infocode": "10000",
  "lives": [
    "0"
    :
    {
      "province": "北京",
      "city": "东城区",
      "adcode": "110101",
      "weather": "晴",
      "temperature": "17",
      "winddirection": "西北",
      "windpower": "≤3",
      "humidity": "15",
      "reporttime": "2025-03-20 22:02:22",
      "temperature_float": "17.0",
      "humidity_float": "15.0"
    }
  ]
}
```



