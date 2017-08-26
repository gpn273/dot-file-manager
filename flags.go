package main

import (
	"flag"
	"path"
)

var (
	flag_config_file string
	flag_config_file_force_create bool
)

func ParseFlags() {
	flag.StringVar(&flag_config_file, "config-file", path.Join(CONFIG_DEFAULT_FILE_LOCATION, CONFIG_DEFAULT_FILE_NAME), "DotFile Manager configuration file")
	flag.BoolVar(&flag_config_file_force_create, "config-file-force-create", false, "Forcefully create a configuration file")

	flag.Parse()
}