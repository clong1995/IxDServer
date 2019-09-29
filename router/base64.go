package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	http.HandleFunc("/base64/add", handler.Base64Add)
	http.HandleFunc("/base64/get", handler.Base64Get)
}
