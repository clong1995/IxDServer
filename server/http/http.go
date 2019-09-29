package http

import (
	. "IxDServer/config"
	_ "IxDServer/router"
	"log"
	"net/http"
)

func StartHttp() {
	err := http.ListenAndServe(CONF.HttpAddr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
