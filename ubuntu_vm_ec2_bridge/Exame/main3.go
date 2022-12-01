package main

import (
	"Exame/algorithms"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Call with path to config. Ex: go run main.go .\\nodesConfig\\nodeA.json")
	}
	configPath := os.Args[1]
	configData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Config file ", configPath, " doesn't exist.")
	}
	node := algorithms.NewDijkstraNode(configData)
	go node.Start()

	for {
	}

}
