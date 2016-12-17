package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tmaesaka/cellar/api"
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
	config := api.NewConfig()

	flag.IntVar(&config.Port, "port", 8084, "TCP port number to listen on (default: 8084)")
	flag.StringVar(&config.DataDir, "datadir", ".", "Path to the Cellar data directory")
	flag.BoolVar(&config.Verbose, "v", false, "Set cellard to be verbose")
	flag.BoolVar(&config.VeryVerbose, "vv", false, "Set cellard to be very verbose")
	flag.Parse()

	fmt.Fprintf(os.Stderr, "Starting cellard... listening on port %d\n", config.Port)
}
