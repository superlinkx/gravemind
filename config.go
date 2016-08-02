package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

func getArgs(params *Arguments) error {
	flag.StringVar(&params.Config, "config", "/etc/gravemind/gravemind.json", "Location of config file (default is /etc/gravemind/gravemind.json)")

	flag.Parse()

	return nil
}

func readConfig(config string, server *Server) error {
	file, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return err
	}

	if err := json.Unmarshal(file, &server); err != nil {
		fmt.Printf("JSON Unmarshalling error: %v\n", err)
		return err
	}

	return nil
}
