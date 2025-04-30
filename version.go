package main

import "fmt"

var (
	version = "dev"
    commit  = "none"
    date    = "unknown"
)

func printVersion() {
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Commit: %s\n", commit)
	fmt.Printf("Date: %s\n", date)
}
