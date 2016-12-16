package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	port := flag.Int("port", 8084, "TCP port number to listen on (default: 8084)")
	flag.Parse()

	fmt.Fprintf(os.Stderr, "Starting cellard... listening on port %d\n", *port)
}
