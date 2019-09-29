package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	//qiniu的key
	http.HandleFunc("/qiniu/key", handler.QiniuKey)
	//qiniu的fileInfo
	http.HandleFunc("/qiniu/fileInfo", handler.QiniuFileInfo)
}
