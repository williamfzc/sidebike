package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/williamfzc/sidebike/pkg/agent"
	"os"
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

		// agent name
		labelFromEnv, ok := os.LookupEnv("SIDEBIKE_AGENT_NAME")
		if ok {
			agentConfig.MachineLabel = labelFromEnv
		}

		// start up
		agentInst := agent.CreateAgent(agentConfig)
		agentInst.Run()
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
