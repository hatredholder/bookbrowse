package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hatredholder/bookbrowse/internal"
	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/hatredholder/bookbrowse/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bookbrowse [book name]",
	Short: "bookbrowse allows you to browse books through hardcover.app in your terminal",
	Long: `bookbrowse - search books within your terminal

Example:
  $ bookbrowse "The Old Man and the Sea"
  `,
	Version: "v0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		if os.Getenv("HARDCOVER_API_KEY") == "" {
			fmt.Println("error: HARDCOVER_API_KEY environment variable must be set")
			fmt.Println("get an API key at https://hardcover.app/account/api")
			os.Exit(0)
		}

		// Check if config exists
		if _, err := os.Stat(utils.GetConfigDir()); os.IsNotExist(err) {
			// Init config
			if err := os.Mkdir(utils.GetConfigDir(), os.ModePerm); err != nil {
				log.Fatal(err)
			}
			// Copy templates
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			if err := os.CopyFS(filepath.Join(utils.GetTeplatesDir(), "predefined"), os.DirFS(filepath.Join(wd, "templates"))); err != nil {
				log.Fatal(err)
			}
		}

		hits := api.Query(args)
		book := internal.Chooser(hits)
		result := internal.Format(book, cmd.Flags())

		fmt.Print(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("template", "t", "default", "template for format output")
}
