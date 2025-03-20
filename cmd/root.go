package cmd

import (
	"github.com/spf13/cobra"
)
var (
	Region string
	rootCmd = &cobra.Command{
		Use:   "corgi",
		Short: "A scheduler for Golang tasks",
		Long: `A scheduler for Golang tasks. You can manage tasks.`,
	}
)
func Execute() error {
	return rootCmd.Execute()
}
func init() {
}
