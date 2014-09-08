# Eth Mods

Ethan: Since I don't know about Casey's git sequence below, everything for now is in branch eris. So to play
```
    go get github.com/eris-ltd/eth-go-mods
    cd $GOPATH/src/github.com/eris-ltd/eth-go-mods
    git checkout eris
```

Doug stuff is in `ethchain/doug.go`, and there's a couple functions suffixed with `Eris`. The old `deCerver/chain` is now in `ethtest` (see ethtest/README.md for details), which is also the bones of decerver interface to eth-go.

Run a simple node with `go run ethtest/tests/main.go -t run`

--------------------------------------------------
This library should contain all of the modifications we make to `eth-go`. Because of how `go` works the libs we build here should be a one for one trade with the `eth-go` libs.

To build on EI's mods to the Core Ethereum-Go libraries:

* git remote add origin git@github.com:eris-ltd/eth-go-mods.git
* git remote add ethereum git@github.com:ethereum/eth-go.git
* git pull origin master
* git pull -s recursive -X ours ethereum develop:eth-dev

Do your work then push to `origin`. You don't have to use `origin` as the remote name for Eris's version of eth-go, but feel free to.

The eth-dev branch of this repo should always track Jeff's dev branch from the main `eth-go` repo. If you need to pull across changes he made then you'll just have to figure out how to do that...! :)

--------------------------------------------------

Ethereum Go is split up in several sub packages Please refer to each
individual package for more information.
  1. [eth](https://github.com/ethereum/eth-go)
  2. [ethchain](https://github.com/ethereum/eth-go/tree/master/ethchain)
  3. [ethwire](https://github.com/ethereum/eth-go/tree/master/ethwire)
  4. [ethdb](https://github.com/ethereum/eth-go/tree/master/ethdb)
  5. [ethutil](https://github.com/ethereum/eth-go/tree/master/ethutil)
  6. [ethpipe](https://github.com/ethereum/eth-go/tree/master/ethpipe)
  7. [ethvm](https://github.com/ethereum/eth-go/tree/master/ethvm)
  8. [ethtrie](https://github.com/ethereum/eth-go/tree/master/ethtrie)
  9. [ethreact](https://github.com/ethereum/eth-go/tree/master/ethreact)
  10. [ethlog](https://github.com/ethereum/eth-go/tree/master/ethlog)

The [eth](https://github.com/ethereum/eth-go) is the top-level package of the Ethereum protocol. It functions as the Ethereum bootstrapping and peer communication layer. The [ethchain](https://github.com/ethereum/eth-go/tree/master/ethchain) contains the Ethereum blockchain, block manager, transaction and transaction handlers. The [ethwire](https://github.com/ethereum/eth-go/tree/master/ethwire) contains the Ethereum [wire protocol](http://wiki.ethereum.org/index.php/Wire_Protocol) which can be used to hook in to the Ethereum network. [ethutil](https://github.com/ethereum/eth-go/tree/master/ethutil) contains utility functions which are not Ethereum specific. The utility package contains the [patricia trie](http://wiki.ethereum.org/index.php/Patricia_Tree), [RLP Encoding](http://wiki.ethereum.org/index.php/RLP) and hex encoding helpers. The [ethdb](https://github.com/ethereum/eth-go/tree/master/ethdb) package contains the LevelDB interface and memory DB interface.
<<<<<<< HEAD

Ethereum Go Development package (C) Jeffrey Wilcke
=======
>>>>>>> 3cf7d73... update readme

Ethereum Go Development package (C) Jeffrey Wilcke
