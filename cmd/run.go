package cmd

import "github.com/spf13/cobra"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run spark history server for Glue Job",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	sparkuiRoot.AddCommand(runCmd)
}
