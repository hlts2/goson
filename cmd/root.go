package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gson",
	Short: "A CLI tool to view JSON",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
