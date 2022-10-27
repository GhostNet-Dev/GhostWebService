package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/GhostNet-Dev/GhostWebService/pkg/webserver"
)

const (
	defaultConfigFilename = "global"
	envPrefix             = "GWS" // GhostWebServer
)

// StartCmdTest test command binding
var StartCmdTest = &cobra.Command{
	Use:   "test",
	Short: "test library",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi~ i'm Ghost")
	},
}

// RootCmd root command binding
var RootCmd = &cobra.Command{
	Use:   "GhostWebService",
	Short: "GhostWebServer in MasterNode",
	Long:  `GhostNet Tiny Web Server for Distributed Web Service`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
		return initializeConfig(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Working with OutOrStdout/OutOrStderr allows us to unit test our command easier
		//out := cmd.OutOrStdout()
		var server webserver.HttpServer
		server.RootPath = rootPath
		server.StartServer(host, port)
	},
}

var (
	host     string
	port     string
	rootPath string
)

func init() {
	RootCmd.Flags().StringVarP(&host, "ip", "i", "127.0.0.1", "Host Address")
	RootCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port Number")
	RootCmd.Flags().StringVarP(&rootPath, "rootPath", "", "", "Home Directory Path")
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()
	v.SetConfigName(defaultConfigFilename)
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	bindFlags(cmd, v)
	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}

		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
