package main

import (
	"Exame/algorithms"
	"bufio"
	"log"
	"os"
	"os/exec"
)

func main() {
	configPath := ".\\nodeGen\\node0.json"
	configData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Config file ", configPath, " doesn't exist.")
	}
	node := algorithms.NewDijkstraNode(configData)
	go node.Start()

	readChan := make(chan string)
	go readInput(readChan)

	files, err := os.ReadDir(".\\nodeGen")
	if err != nil {
		log.Fatal(err)
	}

	var cmdList []*exec.Cmd

	for i, f := range files {
		if i == 0 {
			continue
		}
		cmd := exec.Command("go", "run", ".\\main3.go", ".\\nodeGen\\"+f.Name())
		cmdList = append(cmdList, cmd)
	}

	for _, c := range cmdList {
		go func(cm *exec.Cmd) {
			cm.Stdout = os.Stdout
			cm.Stderr = os.Stderr
			err := cm.Run()
			log.Println(err.Error())
		}(c)
	}

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
