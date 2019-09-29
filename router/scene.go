package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	http.HandleFunc("/scene/add", handler.SceneAdd)
	http.HandleFunc("/scene/getById", handler.SceneGetById)
	/*http.HandleFunc("/project/delete", handler.ProjectDelete)
	http.HandleFunc("/project/update", handler.ProjectUpdate)*/
	//http.HandleFunc("/scene/getById", handler.SceneGetById)
	http.HandleFunc("/scene/getListByProject", handler.SceneGetListByProject)
}
