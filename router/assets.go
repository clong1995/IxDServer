package router

import "net/http"

func init() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("assets"))))
}
