package main

import (
	"fmt"
	"os"
)

func getArgs() map[string]string {
	allArgs := os.Args[1:]
	if len(allArgs) < 2 {
		fmt.Println("Error: Not enough arguments supplied")
		os.Exit(1)
	}
	argsMap := make(map[string]string)
	argsMap["config"] = allArgs[0]
	argsMap["directoryToSort"] = allArgs[1]

	return argsMap
}

func main() {
	args := getArgs()
	config := LoadConfig(args["config"])
	buckets := LoadBucketsFromConfig(config.Buckets, args["directoryToSort"])
	SortFilesIntoBuckets(&buckets, args["directoryToSort"], true)
	fmt.Printf("Completed, %s is sorted", args["directoryToSort"])
}
