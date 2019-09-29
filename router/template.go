package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	http.HandleFunc("/template/add", handler.TemplateAdd)
	http.HandleFunc("/template/getListByUser", handler.TemplateGetListByUser)
	http.HandleFunc("/template/getById", handler.TemplateGetById)
}
