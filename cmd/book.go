package cmd

import (
	"fmt"
	"os"

	"github.com/hatredholder/mediabrowse/internal"
	"github.com/hatredholder/mediabrowse/internal/api"
	"github.com/spf13/cobra"
)

var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "search books through Hardcover",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		if os.Getenv("HARDCOVER_API_KEY") == "" {
			fmt.Println("HARDCOVER_API_KEY environment variable must be set")
			fmt.Println("get an API key at https://hardcover.app/account/api")
			os.Exit(0)
		}

		hits := api.Query(args)
		book := internal.Chooser(hits)
		result := internal.Format(book, cmd.Flags())

		fmt.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(bookCmd)
}
