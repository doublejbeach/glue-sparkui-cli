package cmd

import (
	"github.com/spf13/cobra"
)

var buildImageCmd = &cobra.Command{
	Use:   "image-build",
	Short: "Build Docker image for local Spark UI viewer",
	Run: func(cmd *cobra.Command, args []string) {
	},
	SilenceUsage: true,
}

func init() {
	sparkuiRoot.AddCommand(buildImageCmd)
}
