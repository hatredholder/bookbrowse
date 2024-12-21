package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var movieCmd = &cobra.Command{
	Use:   "movie",
	Short: "search movies through omdbapi.com",
	Long: `search movies through omdbapi.com

Example:
  $ mediabrowse movie "Memento"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		if os.Getenv("OMDB_API_KEY") == "" {
			fmt.Println("OMDB_API_KEY environment variable must be set")
			fmt.Println("get an API key at https://www.omdbapi.com/apikey.aspx")
			os.Exit(0)
		}

		// TODO: build movie search logic
		fmt.Println("movie searching is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(movieCmd)
}
