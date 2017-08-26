package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	FLAG_SHOW_HELP                bool
	FLAG_VERBOSE bool

	ARG_ACTION string
)

func ParseFlags() {
	flag.BoolVar(&FLAG_SHOW_HELP, "help", false, "Show help information")
	flag.BoolVar(&FLAG_VERBOSE, "verbose", false, "Display verbose information")

	flag.Parse()
}

func ParseArgs() {
	ARG_ACTION = string(os.Args[1])
	if isEmpty(ARG_ACTION) {
		ShowHelp()
		os.Exit(1)
	}
}

func ShowHelp() {
	fmt.Printf("NAME: %s (%s)\n", APPLICATION_NAME, APPLICATION_VERSION)
	fmt.Println("MAINTAINER: " + APPLICATION_AUTHOR)
	flag.Usage()
}