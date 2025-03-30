package cmd

import "github.com/spf13/cobra"

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate to run spark-history server in local",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	sparkuiRoot.AddCommand(validateCmd)
}
