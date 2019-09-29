package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	//增加
	http.HandleFunc("/user/add", handler.UserAdd)
	//删除
	http.HandleFunc("/user/delete", handler.UserDelete)
	//列表
	http.HandleFunc("/user/list", handler.UserList)
}
