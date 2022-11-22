package webserver

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

// GhostWeb for Web Service
var (
	RootPath string
	urlMap   map[string]func() (interface{}, string)
)

func init() {
	urlMap = make(map[string]func() (interface{}, string))
}

// RegisterCallMap regist callback handler with file when parsing url
func RegisterCallMap(url string, handler func() (interface{}, string)) {
	urlMap[url] = handler
}

func executeResponse(w http.ResponseWriter, filename string, param map[string]interface{}) {
	var fullPath bytes.Buffer
	fullPath.WriteString(RootPath)
	fullPath.WriteString(filename)

	if t, err := template.ParseFiles(fullPath.String()); err != nil {
		log.Fatal(err)
	} else {
		t.Execute(w, param)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		executeResponse(w, "404.html", nil)
	}
}

func processingHandler(w http.ResponseWriter, r *http.Request) {
	handler, ok := urlMap[r.URL.Path]

	if ok == false {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	internalParam, filename := handler()
	//var internalParam, filename = urlMap[r.URL.Path].(func() (interface{}, string))()

	var param map[string]interface{}
	param = make(map[string]interface{})
	param["UrlParam"] = r.URL.Query() // Ref: https://pkg.go.dev/net/url#URL
	param["InternalParam"] = internalParam

	executeResponse(w, filename, param)
}

// StartServer web site 첫 화면과 그것을 구성하는 파일을 Service 할 수 있는 handler를 등록한다.
func StartServer(host string, port string) {
	http.HandleFunc("/", processingHandler) // indexHandler에 httpServer를 붙여야한다..why?
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
