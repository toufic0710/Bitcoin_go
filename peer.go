package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

// Peer represents a node in the blockchain network
type Peer struct {
	Address    string
	Port       int
	Blockchain *Blockchain
}

// Message struct to encapsulate messages sent over the network
type Message struct {
	Type string
	Data string
}

// NewPeer initializes a new peer with a blockchain instance
func NewPeer(address string, port int, blockchain *Blockchain) *Peer {
	return &Peer{Address: address, Port: port, Blockchain: blockchain}
}

// StartListening initiates the peer to accept incoming connections
func (p *Peer) StartListening() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", p.Address, p.Port))
	if err != nil {
		log.Fatalf("Unable to start listener on %s,%d, %v", p.Address, p.Port, err)
	}
	defer listener.Close()

	log.Printf("Listening for peers on %s:%d", p.Address, p.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %v", err)
			continue
		}
		go p.handleConnection(conn)
	}
}

// handleConnection manages incoming connections
func (p *Peer) handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("Connected to %s", conn.RemoteAddr().String())

	request, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Printf("Error reading request: %v", err)
		return
	}

	var msg Message
	err = json.Unmarshal(request, &msg)
	if err != nil {
		log.Printf("Error unmarshalling request: %v", err)
		return
	}

	switch msg.Type {
	case "AddBlock":
		p.Blockchain.AddBlock([]string{msg.Data})
		log.Println("Added block with data:", msg.Data)
	case "AddData":
		p.Blockchain.AddData(msg.Data)
	case "MineBlock":
		p.Blockchain.MinePendingData()
	default:
		log.Println("Unknown message type received")
	}
}

// ConnectToPeer attempts to connect to another peer and sends a message
func (p *Peer) ConnectToPeer(address string, message Message) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("Error connecting to peer %s: %v", address, err)
		return
	}
	defer conn.Close()

	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		log.Printf("Error sending message to %s: %v", address, err)
		return
	}

	log.Printf("Message sent to %s", address)
}

/*// StartPeerNetwork initializes the peer network, listens for connections, and connects to known peers
func StartPeerNetwork(address string, port int, blockchain *Blockchain) {
	peer := NewPeer(address, port, blockchain)
	go peer.StartListening()

	// Example: Connect to another peer (update with actual address and port)
	otherPeerAddress := "localhost:3000"
	if strconv.Itoa(port) != "3000" { // Prevent connecting to itself if running the example peer
		message := Message{Type: "AddBlock", Data: "Example block data from peer"}
		go peer.ConnectToPeer(otherPeerAddress, message)
	}
}*/
