package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "echopair",
		Short: "echopair is a echo client and server in a single binary",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
