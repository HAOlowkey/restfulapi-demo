package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	vers bool
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "restful-api",
	Short: "restful-api crud dmeo",
	Long:  `restful-api ...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println("0.0.1")
			return nil
		}
		return errors.New("no flags find")
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the restful-api version")
}
