/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"fmt"

	"goku/internal"
)


var inputFile string
var outputFormat string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goku",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		if outputFormat != "json" && outputFormat != "yaml"{
			fmt.Printf("invalid output format, supported formats are json and yaml: %s", outputFormat)
			os.Exit(1)
		} 
		err := internal.Convert(inputFile, outputFormat)
		if err != nil {
			fmt.Printf("conversion failed: %s", err)
			os.Exit(1)
		}
		

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goku.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "input file path")
	rootCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "output format - json or yaml")
	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
}


