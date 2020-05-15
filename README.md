# Simple queue smart contract for Hyperledger Fabric

## Requirements

- [Hyperledger Fabric requirements](https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html)
- Make

## Installation

```bash
make install
make help # to see list of available commands
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

Get few items filtered by content:
```bash
peer chaincode query -n mycc -c '{"Args":["search","bla"]}' -C myc
```

Get item by id:
```bash
peer chaincode query -n mycc -c '{"Args":["get","febb94de-45ca-4cdf-ac73-eb7062bf5999"]}' -C myc
```

Update item content by id:
```bash
peer chaincode invoke -n mycc -c '{"Args":["update","26f83189-bfae-4612-b1c5-608d35af675c", "bla-1234567"]}' -C myc
```

Update few items content by ids:
```bash
peer chaincode invoke -n mycc -c '{"Args":["batchUpdate","c6fec7de-c29d-4e43-98d6-23c1077ab9ce", "bla-090909", "196e69c8-b716-4048-b930-77d65674d466", "Something else"]}' -C myc
```

Move item in the queue (positive value to move up and negative to move down):
```bash
peer chaincode invoke -n mycc -c '{"Args":["move", "ba1fe7d8-65ae-4125-b5b3-1ffbfcfe2062", "-2"]}' -C myc
```

Delete item by id:
```bash
peer chaincode invoke -n mycc -c '{"Args":["delete","5ffca407-8f6d-467a-80c0-b3194e775728"]}' -C myc
```

