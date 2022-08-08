package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/williamfzc/sidebike/pkg/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "test",
	Long:  `test`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server cmd")
		server := &server.Server{
			Port: 9410,
		}
		server.Execute()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
