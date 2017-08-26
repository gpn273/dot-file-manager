package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	FLAG_CONFIG_FILE              string
	FLAG_CONFIG_FILE_FORCE_CREATE bool
	FLAG_SHOW_HELP                bool

	ARG_ACTION string
)

func ParseFlags() {
	flag.StringVar(&FLAG_CONFIG_FILE, "config-file", CONFIG_DEFAULT_FILE_NAME, "DotFile Manager configuration file")
	flag.BoolVar(&FLAG_CONFIG_FILE_FORCE_CREATE, "config-file-force-create", false, "Forcefully create a configuration file")
	flag.BoolVar(&FLAG_SHOW_HELP, "help", false, "Show help information")

	flag.Parse()
}

func ParseArgs() {
	ARG_ACTION = string(os.Args[1])
}

func ShowHelp() {
	fmt.Printf("NAME: %s (%s)\n", APPLICATION_NAME, APPLICATION_VERSION)
	fmt.Println("MAINTAINER: " + APPLICATION_AUTHOR)
	flag.Usage()
}