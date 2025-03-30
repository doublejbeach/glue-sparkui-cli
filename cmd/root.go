package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var sparkuiRoot = &cobra.Command{
	Use:   "glue-sparkui",
	Short: "Local CLI SparkUI on your Glue Job.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				return
			}
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
	SilenceUsage: true,
}

func init() {
	sparkuiRoot.CompletionOptions.DisableDefaultCmd = true
	sparkuiRoot.SetHelpCommand(&cobra.Command{Hidden: true})
}

func Execute() {
	if err := sparkuiRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
