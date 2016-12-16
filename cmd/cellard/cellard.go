package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("cellard (dev)\n")
	fmt.Printf("  --port <num>  %s\n", flag.Lookup("port").Usage)
	fmt.Printf("  -h, --help    Print this help message\n")

	os.Exit(0)
}

func init() {
	flag.Usage = usage
}

func main() {
	port := flag.Int("port", 8084, "TCP port number to listen on (default: 8084)")
	flag.Parse()

	fmt.Fprintf(os.Stderr, "Starting cellard... listening on port %d\n", *port)
}
