package algorithms

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net"
	"sort"
)

type DijkstraNode struct {
	Name         string                  // Nome do nó
	Edges        map[string]int64        // Chave nome, valor peso da aresta
	Table        map[string]int64        // Resultados que esse nó possui por enquanto
	MyConnection *net.UDPConn            // Conexão que escuto
	Connections  map[string]*net.UDPConn // Conexão com nós adjacentes
	LastWave     int64                   // Para identificar rodada atual
	Clock        int64                   // Para ignorar mensagens antigas
	Origin       bool                    // Para apenas o processo que chamou Dijkstra consumir mensagem de END
}

type Message struct {
	OriginName  string           `json:"originName"`
	Destination string           `json:"destination"`
	Clock       int64            `json:"clock"`
	Candidates  map[string]int64 `json:"candidates"`
	Table       map[string]int64 `json:"table"`
	Type        string           `json:"type"`
	Wave        int64            `json:"wave"`
}

type dijkstraNodeConfig struct {
	Name        string           `json:"Name"`
	MyAddress   string           `json:"MyAddress"`
	MyPort      string           `json:"MyPort"`
	Edges       map[string]int64 `json:"Edges"`
	Connections map[string]struct {
		Address string `json:"Address"`
		Port    string `json:"Port"`
	} `json:"Connections"`
}

func NewDijkstraNode(configData []byte) (node *DijkstraNode) {
	var config dijkstraNodeConfig
	err := json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatal("Config file is invalid")
	}

	node = new(DijkstraNode)
	node.Edges = make(map[string]int64)
	node.Table = make(map[string]int64)
	node.Connections = make(map[string]*net.UDPConn)

	node.Name = config.Name
	node.Clock = 0

	for dest, edgeCost := range config.Edges {
		node.Edges[dest] = edgeCost
	}

	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+config.MyPort)
	if err != nil {
		log.Fatal("Couldn't resolve UDP address ", ":"+config.MyPort, " - ", err.Error())
	}
	node.MyConnection, err = net.ListenUDP("udp", ServerAddr) // Minha conexão pra ouvir
	if err != nil {
		log.Fatal("Couldn't listen to UDP address ", err.Error())
	}
	log.Println("Currently listening to", node.MyConnection.LocalAddr().String())

	for dest, connInfo := range config.Connections {
		ServerAddr, err := net.ResolveUDPAddr("udp", connInfo.Address+":"+connInfo.Port)
		if err != nil {
			log.Fatal("Error connecting to node ", dest, " - ", err.Error())
		}
		conn, err := net.DialUDP("udp", nil, ServerAddr)
		if err != nil {
			log.Fatal("Error dialing node ", dest, " - ", err.Error())
		}
		node.Connections[dest] = conn
		log.Println("Connected to node", dest, "successfully.")
	}

	log.Println("Node", node.Name, "running")

	return node
}

func (n *DijkstraNode) sendTable(destName string, candidates map[string]int64, wave int64, originName string) {
	//log.Println("I'm", n.Name, "sending STEP to", destName)
	//fmt.Printf("%+v\n", n.Table)
	//log.Println(candidates)
	n.Clock++ // Aumenta clock (convenção antes)
	msg := Message{
		OriginName:  originName,
		Destination: destName,
		Clock:       n.Clock,
		Candidates:  candidates,
		Table:       n.Table,
		Type:        "STEP",
		Wave:        wave,
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("Error marshalling message:", err.Error())
	}
	if _, ok := n.Connections[destName]; ok {
		// Possuímos conexão direta com o atual mínimo
		_, err = n.Connections[destName].Write(jsonData)
		if err != nil {
			log.Fatal(err.Error())
		}
		//log.Println("Msg tag")
	} else {
		// Não possuímos conexão direta com o atual mínimo.
		// Espalhar mensagem nesse caso
		for _, conn := range n.Connections {
			//log.Println(destName, "is not on my list. Spreading message to", key, "I'm", n.Name)
			_, err = conn.Write(jsonData)
			if err != nil {
				log.Fatal(err.Error())
			}
			//log.Println("Msg tag")
		}
	}
}

func (n *DijkstraNode) finishDijkstra(msg *Message) {
	for key, val := range msg.Table {
		n.Table[key] = val
	}
	n.printTable()
}

func (n *DijkstraNode) printTable() {
	type TmpStruct struct {
		Key string
		Val int64
	}
	tmpList := make([]*TmpStruct, len(n.Table))
	i := 0
	for key, val := range n.Table {
		tmpList[i] = &TmpStruct{
			Key: key,
			Val: val,
		}
		i += 1
	}
	sort.Slice(tmpList, func(i, j int) bool {
		return tmpList[i].Key < tmpList[j].Key
	})
	for _, val := range tmpList {
		log.Println("Node", val.Key, "- Dist:", val.Val)
	}
}

func (n *DijkstraNode) Start() {
	n.listen()
}

func (n *DijkstraNode) sendFinishMessage(lastMsg *Message) {
	n.Clock++ // Aumenta clock (convenção antes)
	msg := Message{
		Table:       n.Table,
		Type:        "END",
		Wave:        lastMsg.Wave,
		OriginName:  lastMsg.OriginName,
		Destination: lastMsg.OriginName,
		Clock:       n.Clock,
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("Error marshalling message:", err.Error())
	}

	if _, ok := n.Connections[lastMsg.OriginName]; ok {
		// Possuímos conexão direta com a origem
		_, err = n.Connections[lastMsg.OriginName].Write(jsonData)
		if err != nil {
			log.Fatal(err.Error())
		}
		//log.Println("Msg tag")
	} else {
		// Não possuímos conexão direta com o atual mínimo.
		// Espalhar mensagem nesse caso
		for _, conn := range n.Connections {
			//log.Println(lastMsg.OriginName, "is not on my list. Spreading message to", key, "I'm", n.Name)
			_, err = conn.Write(jsonData)
			if err != nil {
				log.Fatal(err.Error())
			}
			//log.Println("Msg tag")
		}
	}
}

func (n *DijkstraNode) step(msg *Message) {

	for key, val := range msg.Table {
		if _, ok := n.Table[key]; !ok { // Tabela recebida contem nó que não conhecia
			n.Table[key] = val
		} else {
			if val < n.Table[key] { // Tabela recebida contém nó que eu conheço, mas com peso menor
				// (Nunca chega aqui, pois cada nó processa uma única vez)
				n.Table[key] = val
			}
		}
	}

	myCost := n.Table[n.Name]

	for key, edgeCost := range n.Edges {
		if _, ok := n.Table[key]; !ok { // Novo nó
			n.Table[key] = myCost + edgeCost
			msg.Candidates[key] = n.Table[key] // Adicionamos novo candidato
		} else {
			if myCost+edgeCost < n.Table[key] { // Nó já existe. Atualiza se temos caminho melhor
				n.Table[key] = myCost + edgeCost
				if _, ok := msg.Candidates[key]; ok {
					msg.Candidates[key] = myCost + edgeCost
				}
			}
		}
	}

	if len(msg.Candidates) > 0 {
		// Determina o menor candidato para delegar atividade
		var minKey string
		var minVal int64 = -1
		for key, val := range msg.Candidates {
			if key == n.Name {
				continue
			}
			if minVal == -1 || val < minVal {
				minKey = key
				minVal = val
			}
		}

		// Remove o mínimo, pois já iremos processá-lo
		delete(msg.Candidates, minKey)
		n.sendTable(minKey, msg.Candidates, msg.Wave, msg.OriginName)

	} else {
		// Terminamos. Enviar resposta pra quem chamou o algoritmo
		//log.Println("I'm", n.Name, "and i'm sending END message")
		n.sendFinishMessage(msg)
	}
}

func (n *DijkstraNode) listen() {
	for {
		buf := make([]byte, 100*1024)
		size, _, err := n.MyConnection.ReadFromUDP(buf) // Ler mensagem
		var msg Message
		err = json.Unmarshal(buf[0:size], &msg) // "Castar" para nossa struct Message
		if err != nil {
			log.Fatal("Error casting finish message to struct:", err.Error())
		}

		if msg.Wave != n.LastWave { // Nova rodada
			n.Table = make(map[string]int64) // Limpa o map
			n.LastWave = msg.Wave
			n.Clock = 0
		}

		if msg.Clock <= n.Clock {
			//log.Println("I'm", n.Name, "and received old message. Discarding")
			continue
		}
		n.Clock = msg.Clock // Atualiza próprio clock

		switch msg.Type {
		case "STEP":
			if msg.Destination == n.Name {
				// Eu sou o destinatário
				//log.Println("I'm", n.Name, "and received message for me")
				n.step(&msg)
			} else {
				// Não sou o destinatário, devo redirecionar se o clock da mensagem for maior
				//log.Println("I'm", n.Name, "and received message not for me. Redirecting")
				for _, conn := range n.Connections {
					_, err = conn.Write(buf[0:size])
					if err != nil {
						log.Fatal(err.Error())
					}
					//log.Println("Msg tag")
				}
			}
		case "END":
			if msg.Destination == n.Name {
				// Sou quem chamou o algoritmo e recebi END
				//log.Println("I'm", n.Name, "and received END message")
				n.finishDijkstra(&msg)
			} else {
				// Não sou quem chamou o algoritmo e recebi END. Redirecionar
				//log.Println("I'm", n.Name, "and received END message not for me. Redirecting")
				for _, conn := range n.Connections {
					_, err = conn.Write(buf[0:size])
					if err != nil {
						log.Fatal(err.Error())
					}
					//log.Println("Msg tag")
				}
			}
		}
	}
}

func (n *DijkstraNode) Dijkstra() {
	n.Table = make(map[string]int64) // Limpa o map
	n.Table[n.Name] = 0              // Inicializa próprio peso
	n.Clock = 0                      // Reseta clock
	n.Origin = true                  // Destinatário do resultado

	// Inicializa tabela a partir de arestas saindo
	for key, val := range n.Edges {
		n.Table[key] = val
	}

	// Determina o menor valor na tabela para delegar atividade
	candidates := make(map[string]int64)
	var minKey string
	var minVal int64 = -1
	for key, val := range n.Table {
		if key == n.Name {
			continue
		}
		candidates[key] = val
		if minVal == -1 || val < minVal {
			minKey = key
			minVal = val
		}
	}

	// Se não há arestas de saída, encerra aqui
	if len(candidates) == 0 {
		n.printTable()
		return
	}

	// Remove o mínimo, pois já iremos processá-lo
	delete(candidates, minKey)
	nBig, _ := rand.Int(rand.Reader, big.NewInt(10000))
	wave := nBig.Int64()
	//log.Println("New wave", wave)
	n.LastWave = wave
	n.sendTable(minKey, candidates, wave, n.Name)
}
