package webserver

import (
	"html/template"
	"log"
	"net/http"
)

// GhostWeb for Web Service
var (
	RootPath string
	param map[string]interface{}
	urlMap map[string]interface{}
)

func init() {
	param = make(map[string]interface{})
	urlMap = make(map[string]interface{})
}

// RegisterCallMap regist callback handler when parsing url 
func RegisterCallMap(url string, handler interface{}) {
	urlMap[url] = handler
}

func processingHandler(w http.ResponseWriter, r *http.Request) {
	var internalParam, filename = urlMap[r.URL.Path].(func() (interface{}, string))()

	param["UrlParam"] = r.URL.Query() 	// Ref: https://pkg.go.dev/net/url#URL
	param["InternalParam"] = internalParam

	if t, err := template.ParseFiles(filename); err != nil {
		log.Fatal(err)
	} else {
		t.Execute(w, param)
	}
}

// StartServer web site 첫 화면과 그것을 구성하는 파일을 Service 할 수 있는 handler를 등록한다.
func StartServer(host string, port string) {
	http.HandleFunc("/", processingHandler) // indexHandler에 httpServer를 붙여야한다..why?
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
