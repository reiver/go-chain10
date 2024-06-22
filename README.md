# go-chain10

Package **chain10** provides the exponent used to convert from the _smallest demonation_ to the _default denomination_ for different blockchain-networks, for the Go programming language.

For example, for Ethereum, the exponent is 18 because: 1 ETH = 10^18 wei

This package provides those **exponent**s as **constants**, which are easily usable by programs written in the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-chain10

[![GoDoc](https://godoc.org/github.com/reiver/go-chain10?status.svg)](https://godoc.org/github.com/reiver/go-chain10)

## Example

You can get the exponent based on the chain-id with code such as:

```golang
import "github.com/reiver/go-chain10"

// ...

var chainid uint64 = 1 // Ethereum Mainnet

var exponent uint64 = chain10.Exponent(chainid)
```

## Import

To import package **chain10** use `import` code like the follownig:
```
import "github.com/reiver/go-chain10"
```

## Installation

To install package **chain10** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-chain10
```

## Author

Package **chain10** was written by [Charles Iliya Krempeaux](http://reiver.link)
