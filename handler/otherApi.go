package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/otherApi"
	"IxDServer/service"
	"net/http"
)

//TODO 获取所有菜单的列表，後面可能根權限有關
func OtherWeather(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		p := new(otherApi.Weather)
		err := p.Format(w, r)
		if err != nil {
			return
		}

		//读取head
		token := r.Header.Get("Authorization")
		_, err = service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//业务
		row, err := service.OtherApiWeather(p.City)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, row)
	}
}
