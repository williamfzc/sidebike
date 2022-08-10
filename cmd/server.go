package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/williamfzc/sidebike/pkg/server"
	"os"
)

const ServerConfigFileName = "sidebike-server.json"

var serverCmd = &cobra.Command{
	Use:    "server",
	Short:  "test",
	Long:   `test`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		// config file
		viper.AddConfigPath(".")
		viper.SetConfigFile(ServerConfigFileName)
		err := viper.ReadInConfig()
		if err != nil {
			// no user config found, ignore
			logger.Infof("no config found, use the default one")
		}

		var serverConfig = &server.Config{}
		err = viper.Unmarshal(serverConfig)
		if err != nil {
			logger.Errorf("error in parsing config file: %s", err)
			return
		}

		serverInst := server.CreateNewServer(serverConfig)

		// save config back
		appliedConfig, err := json.MarshalIndent(&serverInst.Config, "", "  ")
		if err != nil {
			// ignore
			logger.Warnf("error when save config back: %s", err)
		}
		_ = os.WriteFile(ServerConfigFileName, appliedConfig, 0644)

		serverInst.Execute()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
