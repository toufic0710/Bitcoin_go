package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "Port to listen on")
	connectTo := flag.String("connect", "", "Address of the peer to connect to")
	flag.Parse()

	blockchain := NewBlockchain(1)
	address := "localhost"
	peer := NewPeer(address, *port, blockchain)

	go peer.StartListening()

	if *connectTo != "" {
		fmt.Println("Attempting to connect to peer:", *connectTo)
		peer.ConnectToPeer(*connectTo, Message{Type: "NewPeer", Data: "Hi from " + address})
	} else {
		fmt.Println("No peer to connect to, waiting for connections...")
	}

	fmt.Println("Type 'add <data>' to add new data to the blockchain, or 'mine' to mine a new block with pending data.")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Split(input, " ")

		switch args[0] {
		case "add":
			if len(args) < 2 {
				fmt.Println("Usage: add <data>")
				continue
			}
			data := strings.Join(args[1:], " ")
			blockchain.AddData(data)
			fmt.Println("Data added to pending transactions.")
		case "mine":
			blockchain.MinePendingData()
			fmt.Println("New block mined.")
		default:
			fmt.Println("Unknown command. Try 'add <data>' or 'mine'.")
		}
	}
}
