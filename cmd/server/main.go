package main

import (
	"os"
	"github.com/GhostNet-Dev/GhostWebService/cmd/server/commands"
)


func main() {	
	cmd := cmd.NewRootCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}