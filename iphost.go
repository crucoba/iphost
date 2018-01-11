package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("\nList network interfaces: %s list\n", args[0])
		fmt.Printf("Set an alias: %s set <index> <alias>\n", args[0])
		os.Exit(0)
	}
	switch args[1] {
	case "list":
		listInterfaces()
	case "set":
		setAlias(args[2], args[3])
	default:
		fmt.Printf("\nUnknown command %s\n", args[1])
	}
}
