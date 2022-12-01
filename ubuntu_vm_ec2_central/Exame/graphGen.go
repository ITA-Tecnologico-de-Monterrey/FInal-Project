package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	n := 20           // Number of nodes
	edgeProb := 0.2   // Chance of existing edge between 2 nodes
	maxEdgeCost := 10 // Maximum possible edge cost

	rand.Seed(time.Now().UnixNano())
	adjMatrix := make([][]int64, n)
	for i := 0; i < n; i++ {
		adjMatrix[i] = make([]int64, n)
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				continue
			}
			if rand.Float64() <= edgeProb {
				cost := rand.Int63n(int64(maxEdgeCost-1)) + 1
				adjMatrix[i][j] = cost
				adjMatrix[j][i] = cost
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d, ", adjMatrix[i][j])
		}
		fmt.Printf("\n")
	}

	type Connections struct {
		Address string `json:"Address"`
		Port    string `json:"Port"`
	}

	type DijkstraNodeConfig struct {
		Name        string                 `json:"Name"`
		MyAddress   string                 `json:"MyAddress"`
		MyPort      string                 `json:"MyPort"`
		Edges       map[string]int64       `json:"Edges"`
		Connections map[string]Connections `json:"Connections"`
	}

	basePort := 10000
	for i := 0; i < n; i++ {
		iStr := strconv.FormatInt(int64(i), 10)
		iPort := strconv.FormatInt(int64(basePort+i), 10)
		conf := DijkstraNodeConfig{
			Name:        iStr,
			MyAddress:   "localhost",
			MyPort:      iPort,
			Edges:       make(map[string]int64),
			Connections: make(map[string]Connections),
		}
		for j := 0; j < n; j++ {
			if adjMatrix[i][j] == 0 {
				continue
			}
			jStr := strconv.FormatInt(int64(j), 10)
			jPort := strconv.FormatInt(int64(basePort+j), 10)
			conf.Edges[jStr] = adjMatrix[i][j]
			conf.Connections[jStr] = Connections{
				Address: "localhost",
				Port:    jPort,
			}
		}
		jsonData, err := json.Marshal(conf)
		if err != nil {
			log.Fatal(err.Error())
		}
		f, err := os.Create(".\\nodeGen\\node" + iStr + ".json")
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = f.WriteString(string(jsonData))
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	log.Println("Done")

}
