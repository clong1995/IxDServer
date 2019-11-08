package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	//增加
	http.HandleFunc("/user/add", handler.UserAdd)
	//添加公司用户
	http.HandleFunc("/user/addCompanyUser", handler.UserAddCompanyUser)
	http.HandleFunc("/user/addDepartmentUser", handler.UserAddDepartmentUser)
	http.HandleFunc("/user/addSuperAdmin", handler.UserAddSuperAdmin)
	//删除
	http.HandleFunc("/user/delete", handler.UserDelete)
	//列表
	http.HandleFunc("/user/list", handler.UserList)
	//列表
	http.HandleFunc("/user/info", handler.UserInfo)
}
