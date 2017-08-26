package main

import (
	"flag"
	"path"
	"fmt"
)

var (
	flag_config_file string
	flag_config_file_force_create bool
	flag_show_help bool
)

func ParseFlags() {
	flag.StringVar(&flag_config_file, "config-file", path.Join(CONFIG_DEFAULT_FILE_LOCATION, CONFIG_DEFAULT_FILE_NAME), "DotFile Manager configuration file")
	flag.BoolVar(&flag_config_file_force_create, "config-file-force-create", false, "Forcefully create a configuration file")
	flag.BoolVar(&flag_show_help, "help", false, "Show help information")

	flag.Parse()
}

func ShowHelp() {
	fmt.Printf("NAME: %s (%s)\n", APPLICATION_NAME, APPLICATION_VERSION)
	fmt.Println("MAINTAINER: " + APPLICATION_AUTHOR)
	flag.Usage()
}