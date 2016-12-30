package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tmaesaka/cellar/api"
	"github.com/tmaesaka/cellar/config"
)

func usage() {
	fmt.Printf("cellard (dev)\n")
	fmt.Printf("  --port <num>      %s\n", flag.Lookup("port").Usage)
	fmt.Printf("  --datadir <path>  %s\n", flag.Lookup("datadir").Usage)
	fmt.Printf("  -v                %s\n", flag.Lookup("v").Usage)
	fmt.Printf("  -vv               %s\n", flag.Lookup("vv").Usage)
	fmt.Printf("  -h, --help        Print this help message\n")

	os.Exit(0)
}

func init() {
	flag.Usage = usage
}

func main() {
	config := config.NewApiConfig()

	flag.IntVar(&config.Port, "port", 8084, "TCP port number to listen on (default: 8084)")
	flag.StringVar(&config.DataDir, "datadir", ".", "Path to the Cellar data directory")
	flag.BoolVar(&config.Verbose, "v", false, "Set cellard to be verbose")
	flag.BoolVar(&config.VeryVerbose, "vv", false, "Set cellard to be very verbose")
	flag.Parse()

	if err := api.Run(config); err != nil {
		// TODO(toru): Use err.Error() once api.Run() starts doing stuff.
		fmt.Fprintf(os.Stderr, "Failed to start cellard\n")
		os.Exit(1)
	}
}
