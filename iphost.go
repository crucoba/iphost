package main

import (
	"os"
)

func main() {
	args := os.Args
	switch args[1] {
	case "list":
		listInterfaces()
	case "set":
		setAlias(args[2], args[3])
	}
}
