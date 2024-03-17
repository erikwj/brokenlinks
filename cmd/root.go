/*
Copyright © 2024 @erikwj
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/erikwj/brokenlinks/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "brokenlinks",
	Short: "A cli to validate a markdown tree for broken links",
	Long: `A cli to validate a markdown tree for broken links

	Currently support for:
	- image links in png, svg, or gif format
	- web links [manually for now]
	- file links in same directory
	`,
	// Execution
	Run: func(cmd *cobra.Command, args []string) {
		directory := dir
		extension := ext
		f := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == extension {
				if debug {
					fmt.Fprintf(cmd.OutOrStdout(), "# Validating %s \n", path)
				}

				if err := internal.ValidateLinks(path, extension, only_errors); err != nil {
					fmt.Printf("# Error validating links in file %s: %v\n", path, err)
				}
			}
			return nil
		}
		err := filepath.Walk(directory, f)

		if err != nil {
			fmt.Printf("# Error walking the path %s: %v\n", directory, err)
			os.Exit(1)
		}
	},
}

var (
	dir         string
	ext         string
	debug       bool
	only_errors bool
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	RootCmd.PersistentFlags().StringVar(&ext, "ext", ".md", "File extension to be filtered on")
	RootCmd.PersistentFlags().StringVar(&dir, "dir", "", "Required: directory to be checked")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Optional: print file names that are being checked; default: false")
	RootCmd.PersistentFlags().BoolVar(&only_errors, "only_errors", false, "Optional: print only errors, no weblinks; default: false")

}
