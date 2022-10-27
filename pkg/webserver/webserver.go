package webserver

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// HttpServer for Web Service
type HttpServer struct {
	RootPath string
	param map[string]interface{}
}

func (httpServer *HttpServer) indexHandler(w http.ResponseWriter, r *http.Request) {
	// Ref: https://pkg.go.dev/net/url#URL
	httpServer.param["UrlParam"] = r.URL.Query()
	var filename string
	if r.URL.Path == "/" {
		filename = "index.html"
	} else {
		filename = strings.TrimLeft(r.URL.Path, "/")
	}

	if t, err := template.ParseFiles(filename); err != nil {
		log.Fatal(err)
	} else {
		t.Execute(w, httpServer.param)
	}
}

// StartServer web site 첫 화면과 그것을 구성하는 파일을 Service 할 수 있는 handler를 등록한다.
func (httpServer *HttpServer) StartServer(host string, port string) {
	httpServer.param = make(map[string]interface{})

	http.HandleFunc("/", httpServer.indexHandler) // indexHandler에 httpServer를 붙여야한다..why?
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
