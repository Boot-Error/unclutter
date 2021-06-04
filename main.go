package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func getArgs() map[string]string {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}
	directoryToSort := flag.String("dir", currentDir, "directory to sort, default to current working directory")

	flag.Parse()
	argsMap := make(map[string]string)
	argsMap["directoryToSort"] = *directoryToSort

	return argsMap
}

func main() {
	args := getArgs()
	configFilePath := SetupConfigLocally()
	config := LoadConfig(configFilePath)
	buckets := LoadBucketsFromConfig(config.Buckets, args["directoryToSort"])
	SortFilesIntoBuckets(&buckets, args["directoryToSort"], true)
	fmt.Printf("Completed, %s is sorted", args["directoryToSort"])
}
