package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/base64"
	"IxDServer/service"
	"net/http"
)

//增加场景
func Base64Add(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		token := r.Header.Get("Authorization")
		_, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(base64.Add)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		id, err := service.Base64Add(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, id)
	}
}

//获取所有菜单的列表
func Base64Get(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {

		//读取head
		token := r.Header.Get("Authorization")
		_, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(base64.Get)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		rows, err := service.Base64Get(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}
