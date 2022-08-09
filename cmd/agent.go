package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/williamfzc/sidebike/pkg/agent"
)

var agentCmd = &cobra.Command{
	Use:    "agent",
	Short:  "test",
	Long:   `test`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		// config file
		viper.AddConfigPath(".")
		viper.SetConfigFile("sidebike.json")
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

		// start up
		agentInst := agent.CreateAgent(agentConfig)
		agentInst.Run()
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
