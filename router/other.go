package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	//增加
	http.HandleFunc("/other/weather", handler.OtherWeather)
}
