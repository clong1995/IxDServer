package service

import (
	. "IxDServer/config"
	"io/ioutil"
	"net/http"
)

func OtherApiWeather(city string) (interface{}, error) {
	urlStr := "https://restapi.amap.com/v3/weather/weatherInfo?city=" + city + "&key=" + CONF.GaoDeKey
	res, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}
