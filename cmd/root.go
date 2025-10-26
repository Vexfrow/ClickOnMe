package cmd

import (
	"ClickOnMe/checker"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ClickOnMe",
	Short: "Simple tool to test if a webpage is vulnerable to ClickJacking",
	Long: `


	ClickOnMe is a tool that verify if a WebPage (or a set of WebPages) is potentially vulnerable to the ClickJacking vulnerability
	It can either be used by providing a unique URL with the "--url" option
	or a list of URLs by writing them in a file (one per line) and using the "--inputFile" option

	`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print(cmd.Long)
		checker.StartChecking()
	},
}

func Execute() {

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n\n\n", rootCmd.Long)
		fmt.Printf("Usage :\n  ./ClickOnMe <[-i PathToInputFile] or [-u URL]> [-o PathToOutputFile] [-r nbRequestPerSecond]\n\n")
		fmt.Printf("Flag :\n  -h, --help			Print Help\n  -c, --color			Toggle off colors (Colors only work if outputs are printed on the console)\n\n")
		fmt.Printf("Options :\n  -i, --inputFile		Path to the file containing URLs\n  -o, --outputFile	Path to the file that will contain results\n  -u, --URL			URL that must be checked\n  -r, --rate			Number of request per second (default: 25 r/s)\n\n")
		os.Exit(0)
	})
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&checker.ToggleColor, "color", "c", true, "Toggle off colors (Colors only work if outputs are printed on the console)")
	rootCmd.Flags().StringVarP(&checker.FileURLs, "inputFile", "i", "", "Input File containing URLs")
	rootCmd.Flags().StringVarP(&checker.OutputFile, "outputFile", "o", "", "Output File that will contain results")
	rootCmd.Flags().StringVarP(&checker.URL, "url", "u", "", "URL that must be checked")
	rootCmd.Flags().IntVarP(&checker.RateLimit, "rate", "r", 25, "Rate limit for the number of request that must be send every second (default : 25)")
}
