package main

import (
	"Exame/algorithms"
	"bufio"
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

	readChan := make(chan string)
	go readInput(readChan)

	for {
		<-readChan
		node.Dijkstra()
	}

}

func readInput(ch chan string) {
	// Rotina que “escuta” o stdin
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal("Error listening to keyboard: ", err.Error())
		}
		ch <- string(text)
	}
}
