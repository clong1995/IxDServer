package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	http.HandleFunc("/department/list", handler.DepartmentList)
}
