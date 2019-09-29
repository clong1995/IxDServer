package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/template"
	"IxDServer/service"
	"net/http"
)

//增加场景
func TemplateAdd(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(template.Add)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.TemplateAdd(uid, p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

//获取场景列表
func TemplateGetListByUser(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数

		//业务
		row, err := service.TemplateGetListByUser(uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, row)
	}
}

//获取一个场景

//获取所有菜单的列表
func TemplateGetById(w http.ResponseWriter, r *http.Request) {
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
		p := new(template.GetById)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		rows, err := service.TemplateGetById(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}
