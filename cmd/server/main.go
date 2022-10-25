package main

import (
	"os"
	"github.com/GhostNet-Dev/GhostWebService/cmd/commands"
)


func main() {	
	cmd := cmd.NewRootCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}