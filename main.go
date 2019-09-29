package main

import (
	"IxDServer/server/http"
	"log"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	//启动http
	http.StartHttp()
}
