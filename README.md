# Simplified Blockchain

This project is a simplified blockchain implementation designed to demonstrate fundamental blockchain technology concepts including block creation, proof of work (PoW), and a basic peer-to-peer (P2P) network for sharing data and blocks.

## Features

- **Dynamic block creation** with proof of work.
- **Data addition** to the blockchain via command-line interface.
- **Simple P2P communication** between network nodes to share data and blocks.

## Prerequisites

- **GoLang 1.18**

## Installation

1. **Clone the Repository**

   Open a terminal and run the following command to clone this repository:

   ```sh
   git clone https://github.com/toufic0710/Bitcoin_go.git
   cd Bitcion_go
   ```


2. **Compile the Project**

Compile the project to generate the executable:


```
go build -o blockchainApp
```


## Usage

### Start a Node

To start a blockchain node on your machine, use the following command. Replace `<port>` with the port you want your node to listen on.


```./blockchainApp -port=<port>```

Example:

```./blockchainApp -port=3000```


### Connect a Node to Another

To connect a new node to the network, specify the address of an existing node using the `-connect` option when starting a new node.

```
./blockchainApp -port=<new_port> -connect=localhost:<existing_port>
```

Example:


```
./blockchainApp -port=3001 -connect=localhost:3000
```


### Add Data to the Blockchain

Once the node is started, you can add data to the blockchain using the `add` command followed by your data.

```
add <your_data_here>

```

### Mine a New Block

To mine a new block with pending data, use the `mine` command.

```
mine
```


## Architecture

The project consists of several main files:

- `blockchain.go`: Contains the blockchain logic, including block creation and mining.
- `peer.go`: Manages P2P communication between nodes.
- `main.go`: The program's entry point, handling the user interface and orchestrating blockchain operations.
- `proofOfWork.go`: Contains the proof of work logic.





