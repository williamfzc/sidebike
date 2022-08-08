package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "sidebike",
	Short: "test",
	Long: `test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root cmd")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
