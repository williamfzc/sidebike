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
			return
		}

		// config format
		var agentConfig = &agent.Config{}
		err = viper.Unmarshal(agentConfig)
		if err != nil {
			return
		}

		// agent name
		labelFromEnv, ok := os.LookupEnv("SIDEBIKE_AGENT_NAME")
		if ok {
			agentConfig.MachineLabel = labelFromEnv
		}

		agentInst := agent.CreateAgent(agentConfig)
		// save config back
		appliedConfig, err := json.MarshalIndent(&agentInst.Config, "", "  ")
		if err != nil {
			// ignore
		}
		_ = os.WriteFile(ConfigFileName, appliedConfig, 0644)

		// start up
		agentInst.Run()
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
