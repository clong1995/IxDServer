package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	//增加
	http.HandleFunc("/bug/add", handler.BugAdd)
	//列表
	http.HandleFunc("/bug/list", handler.BugList)
	//info
	http.HandleFunc("/bug/info", handler.BugInfo)
}
