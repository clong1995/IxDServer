package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/service"
	"net/http"
)

func DepartmentList(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		//读取head
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//业务
		rows, err := service.DepartmentList(uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}
