/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ClickOnMe/checker"
	"os"

	"github.com/spf13/cobra"
)

// Base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ClickOnMe",
	Short: "Simple tool to test if a webpage is vulnerable to ClickJacking",
	Long: `ClickOnMe @Vexfrow
	
	ClickOnMe is a tool that verify if a WebPage (or a set of WebPages) is vulnerable to the ClickJacking vulnerability
	It can either be used by providing a unique URL with the "--url" option
	or a list of URLs by writing them in a file (one per line) and using the "--inputFile" option `,
}

func Execute() {
	err := rootCmd.Execute()
	checker.StartChecking()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&checker.FileURLs, "inputFile", "i", "", "Input File containing URLs")
	rootCmd.Flags().StringVarP(&checker.OutputFile, "outputFile", "o", "", "Output File containing results")
	rootCmd.Flags().StringVarP(&checker.URL, "url", "u", "", "URL that must be checked")
	rootCmd.Flags().IntVarP(&checker.RateLimit, "rate", "r", 25, "Rate limit for the number of request that must be send every second")
}
