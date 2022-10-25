package main

import (
	"os"
	"github.com/GhostNet-Dev/GhostWebService/cmd/server/commands"
)


func main() {	
	startCmd := cmd.RootCmd
	startCmd.AddCommand(cmd.StartCmdTest)
	if err := startCmd.Execute(); err != nil {
		os.Exit(1)
	}
}