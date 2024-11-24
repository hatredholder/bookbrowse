package cmd

import (
	"fmt"
	"os"

	"github.com/hatredholder/bookbrowse/internal"
	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bookbrowse [book name]",
	Short: "bookbrowse allows you to browse books through hardcover.app in your terminal",
	Long: `bookbrowse - Search books within your terminal

Example:
  $ bookbrowse "The Old Man and the Sea" -m
  `,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: print error message on > 1 argument

		hits := api.Query(args[0])

		fmt.Println(internal.Format(internal.Chooser(hits), cmd.Flags()))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("markdown", "m", false, "Format output into markdown")
	rootCmd.Flags().BoolP("fulldescription", "d", false, "Display full description")
}
