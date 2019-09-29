package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/project"
	"IxDServer/service"
	"net/http"
)

func ProjectAdd(w http.ResponseWriter, r *http.Request) {
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
		p := new(project.Add)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.ProjectAdd(uid, p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

func ProjectList(w http.ResponseWriter, r *http.Request) {
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

		//业务
		row, err := service.ProjectList(uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, row)
	}
}

//删除设备
func ProjectDelete(w http.ResponseWriter, r *http.Request) {
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
		p := new(project.Delete)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.ProjectDelete(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}
