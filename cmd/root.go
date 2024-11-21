package cmd

import (
	"fmt"
	"os"

	"github.com/hatredholder/bookbrowse/internal"
	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bookbrowse",
	Short: "bookbrowse allows you to browse books through hardcover.app in your terminal",
	Long: `
bookbrowse - Search books within your terminal

Usage
  $ bookbrowse [options] [book name]

Example:
  $ bookbrowse --markdown "The Old Man and the Sea"

Options:
  -h --help              Display usage details
  -m --markdown          Format output into Markdown
  -d --fulldescription   Display full description
  -r --tenrating         Return rating out of ten (for ex. 5.6/10)
  `,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: print error message on > 1 argument

		hits := api.Query(args[0])

		fmt.Println(internal.Format(internal.Chooser(hits)))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
