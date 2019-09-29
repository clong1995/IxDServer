package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	//增加
	http.HandleFunc("/menu/add", handler.MenuAdd)
	http.HandleFunc("/menu/addMulti", handler.MenuAddMulti)
	//获取当前用户的菜单
	http.HandleFunc("/menu/getList", handler.MenuGetList)
	//TODO 根据用户获取菜单，后期加子菜单
	http.HandleFunc("/menu/getListByUser", handler.MenuGetListByUser)
	//根据用户获取菜单包含状态
	http.HandleFunc("/menu/getListByUserStatus", handler.MenuGetListByUserStatus)
}
