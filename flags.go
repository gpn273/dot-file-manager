package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	FLAG_SHOW_HELP = flag.Bool("help", false, "Show help information")
	FLAG_VERBOSE = flag.Bool("verbose", false, "Display verbose information")

	ARG_ACTION string
)

func ParseFlags() {
	flag.Parse()
}

func GetFriendlyArgName() string {
	return strings.ToLower(ARG_ACTION)
}

func ParseArgs() {
	if len(os.Args) < 2 {
		ShowHelp()
		os.Exit(1)
	} else {
		ARG_ACTION = string(os.Args[1])
		if isEmpty(ARG_ACTION) {
			ShowHelp()
			os.Exit(1)
		}
	}
}

func ShowHelp() {
	fmt.Printf("NAME: %s (%s)\n", APPLICATION_NAME, APPLICATION_VERSION)
	fmt.Println("MAINTAINER: " + APPLICATION_AUTHOR)
	fmt.Println("\nUsage: dot-file-manager [update | create-config | purge-backups | create-global-installation]\n")
	flag.Usage()
}