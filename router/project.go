package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	http.HandleFunc("/project/add", handler.ProjectAdd)
	http.HandleFunc("/project/delete", handler.ProjectDelete)
	/*http.HandleFunc("/project/update", handler.ProjectUpdate)*/
	http.HandleFunc("/project/list", handler.ProjectList)
}
