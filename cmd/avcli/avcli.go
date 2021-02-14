package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/deanhigh/alphavantage"
)

func dumpFundamental(symbol string) error {
	avc, err := alphavantage.NewClient()
	if err != nil {
		return err
	}

	co, err := avc.FundamentalService.GetCompanyOverview(symbol)
	if err != nil {
		return err
	}

	err = alphavantage.DumpJSON(os.Stdout, co)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fundamentalCommand := flag.NewFlagSet("fundamental", flag.ExitOnError)

	symbol := fundamentalCommand.String("symbol", "", "Set to the location of the config file")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("Subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case fundamentalCommand.Name():
		fundamentalCommand.Parse(os.Args[2:])

		if fundamentalCommand.Parsed() {
			dumpFundamental(*symbol)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
