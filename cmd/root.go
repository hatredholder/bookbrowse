package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mediabrowse [book/movie] [title] [flags]",
	Short: "mediabrowse allows you to browse movies & books in your terminal",
	Long: `mediabrowse - search movies & books within your terminal

Example:
  $ mediabrowse book "The Old Man and the Sea"
  `,
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// set global flags
	rootCmd.PersistentFlags().StringP("template", "t", "default", "template for format output")
	rootCmd.PersistentFlags().Lookup("template").DefValue = ""

	// disable completion command
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	// disable help command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
