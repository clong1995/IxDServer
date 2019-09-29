package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/qiniu"
	"IxDServer/service"
	"net/http"
)

//七牛token
func QiniuKey(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	token := r.Header.Get("Authorization")
	_, err := service.AuthUnToken(token)
	if err != nil {
		network.ErrStrCode(w, err.Error(), AUTH)
		return
	}
	//业务
	row := service.QiniuKey()
	network.Succ(w, row)
}

func QiniuFileInfo(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		//认证
		//读取head
		network.FbdReq(w)
	}
	//从头部取出token
	token := r.Header.Get("Authorization")
	//校验token
	_, err := service.AuthUnToken(token)
	if err != nil {
		network.ErrStrCode(w, err.Error(), AUTH)
		return
	}
	//参数
	p := new(qiniu.FileInfo)
	err = p.Format(w, r)
	if err != nil {
		return
	}
	//业务
	row, err := service.QiniuFileInfo(p)
	if err != nil {
		network.ErrStr(w, err.Error())
		return
	}
	network.Succ(w, row)
}
