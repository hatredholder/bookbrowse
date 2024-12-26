package cmd

import (
	"fmt"
	"os"

	"github.com/hatredholder/mediabrowse/internal/templates"
	"github.com/hatredholder/mediabrowse/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "mediabrowse allows you to browse movies & books in your terminal",
	Long: `mediabrowse - search movies & books within your terminal

Examples:
  $ mediabrowse movie "Memento"
  $ mediabrowse book "The Old Man and the Sea" --template markdown
  `,
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		if showTemplates, _ := cmd.Flags().GetBool("templates"); showTemplates == true {
			fmt.Println("Currently available templates:", utils.FindAvailableTmpls())
			os.Exit(0)
		}
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

	// set local flags
	rootCmd.Flags().BoolP("templates", "", false, "show available templates")

	// set custom usage template
	rootCmd.Root().SetUsageTemplate(templates.UsageTmpl)

	// disable help command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	// disable completion command
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}
