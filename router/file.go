package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	//新建文件夹
	http.HandleFunc("/file/addFolder", handler.FileAddFolder)
	//添加文件
	http.HandleFunc("/file/addFile", handler.FileAddFile)
	//顶级目录
	http.HandleFunc("/file/listTopFolder", handler.FileListTopFolder)
	//文件列表
	http.HandleFunc("/file/list", handler.FileList)
	http.HandleFunc("/file/deleteList", handler.FileDeleteList)
	//任务列表
	http.HandleFunc("/file/taskList", handler.FileTaskList)
	//删除文件
	http.HandleFunc("/file/delete", handler.FileDelete)
	//删除文件
	http.HandleFunc("/file/uploadFinish", handler.FileUploadFinish)
}
