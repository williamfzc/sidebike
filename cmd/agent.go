package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/williamfzc/sidebike/pkg/agent"
)

const ConfigFileName = "sidebike.json"

var agentCmd = &cobra.Command{
	Use:    "agent",
	Short:  "test",
	Long:   `test`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		// config file
		viper.AddConfigPath(".")
		viper.SetConfigFile(ConfigFileName)
		err := viper.ReadInConfig()
		if err != nil {
			// no user config found, ignore
			logger.Infof("no config found, use the default one")
		}

		// config format
		var agentConfig = &agent.Config{}
		err = viper.Unmarshal(agentConfig)
		if err != nil {
			logger.Errorf("error in parsing config file: %s", err)
			return
		}

		// agent name
		labelFromEnv, ok := os.LookupEnv("SIDEBIKE_AGENT_NAME")
		if ok {
			logger.Infof("read custom agent name from env: %s", labelFromEnv)
			agentConfig.MachineLabel = labelFromEnv
		}

		agentInst := agent.CreateAgent(agentConfig)
		// save config back
		appliedConfig, err := json.MarshalIndent(&agentInst.Config, "", "  ")
		if err != nil {
			// ignore
			logger.Warnf("error when save config back: %s", err)
		}
		_ = os.WriteFile(ConfigFileName, appliedConfig, 0644)

		// start up
		agentInst.Run()
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
