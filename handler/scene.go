package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/scene"
	"IxDServer/service"
	"net/http"
)

//增加场景
func SceneAdd(w http.ResponseWriter, r *http.Request) {
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
		p := new(scene.Add)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		scene, err := service.SceneAdd(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, scene)
	}
}

//获取场景列表
func SceneGetListByProject(w http.ResponseWriter, r *http.Request) {
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
		p := new(scene.GetListByProject)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		row, err := service.SceneGetListByProject(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, row)
	}
}

//获取一个场景

//获取所有菜单的列表
func SceneGetById(w http.ResponseWriter, r *http.Request) {
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
		p := new(scene.GetById)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		rows, err := service.SceneGetById(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}
