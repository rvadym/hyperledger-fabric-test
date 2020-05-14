# Simple queue smart contract for Hyperledger Fabric

## Requirements

- [Hyperledger Fabric requirements](https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html)
- Make

## Installation

```bash
make install
```

## Start development/testing setup

Open console and start docker containers:

```bash
make build
make start
```

Open one more console and SSH into development container:

```make
make ssh
```

Open one more console and SSH into chaincode container:

```make
make ssh-chaincode
```

Open one more console and SSH into cli container:

```make
make ssh-cli
```

## Unit testing

Inside development container:

```bash
make test
```

## Install and init chaincode

Inside chaincode container:

```bash
go build -o simplequeue
CORE_CHAINCODE_ID_NAME=mycc:0 CORE_PEER_TLS_ENABLED=false ./simplequeue -peer.address peer:7052
```

Inside cli container:

```bash
peer chaincode install -p ../fabric -n mycc -v 0
```

You must see:

```bash
Installed remotely: response:<status:200 payload:"OK" >
```

Init:

```bash
peer chaincode instantiate -n mycc -v 0 -c '{"Args":["init"]}' -C myc
```

## CouchDB using Fauxton UI

Open `http://localhost:5984/` <br />
L: admin <br />
P: test <br />
DB: myc_mycc <br />

## Chaincode usage

Inside cli container:

Add items to the queue:
```bash
peer chaincode invoke -n mycc -c '{"Args":["enqueue","Item #1"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["enqueue","Item #2"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["enqueue","Item #3"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["enqueue","Item #4"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["enqueue","Item #5"]}' -C myc
```

Take items from the queue (can be executed few times):
```bash
peer chaincode invoke -n mycc -c '{"Args":["dequeue"]}' -C myc
```

