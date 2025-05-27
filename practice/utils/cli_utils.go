package utils

import (
	"flag"
	"fmt"
)

func ParseFlags() (string, int) {
	var dbHost string
	var port int

	flag.StringVar(&dbHost, "dbhost", "localhost", "Database host")
	flag.IntVar(&port, "port", 8080, "Server port")

	flag.Parse()

	fmt.Printf("Using DB host: %s\n", dbHost)
	fmt.Printf("Server port: %d\n", port)

	return dbHost, port
}
