package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	defaultConfigFile = "/opt/elasticbeanstalk/deploy/configuration/containerconfiguration"

	name       = flag.String("name", "", "name of ENV variable to get")
	configFile = flag.String("config-file", defaultConfigFile, "full path to config where to look for ENV")
)

func main() {
	flag.Parse()

	if *name == "" {
		log.Fatal("ERROR: missing --name arg")
	}

	if value := readEnvFromUserEnv(); value != "" {
		fmt.Println(value)
		os.Exit(0)
	}

	value, err := readEnvFromFile()
	if err != nil {
		log.Fatalf("ERROR: reading name:%s from file:%s, ", *name, configFile, err)
	}

	fmt.Println(value)
}

func readEnvFromUserEnv() string {
	return os.Getenv(*name)
}

func readEnvFromFile() (string, error) {
	configFileRaw, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return "", err
	}

	var c ebConfig
	if err := json.Unmarshal(configFileRaw, &c); err != nil {
		return "", err
	}

	for _, fullEnv := range c.Optionsettings.Environment {
		split := strings.Split(fullEnv, "=")
		if split[0] == *name {
			return split[1], nil
		}
	}

	return "", nil
}
