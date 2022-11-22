package ghostweb

import (
	"fmt"

	web "github.com/GhostNet-Dev/GhostWebService/pkg/webserver"
)

func indexView() (interface{}, string) {
	return nil, "index.html"
}

func init() {
	web.RegisterCallMap("/", indexView)
}

// StartGhostWeb ghost web을 위한 website를 시작하는 모듈
func StartGhostWeb(rootPath string, host string, port string) {
	fmt.Println("GhostWeb Start.")
	web.RootPath = rootPath
	web.StartServer(host, port)
}
