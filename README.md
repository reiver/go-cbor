# go-cbor

Package **cbor** implements encoding and decoding of CBOR (Concise Binary Object Representation) as defined by [IETF RFC-8949](https://datatracker.ietf.org/doc/html/rfc8949), for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-cbor

[![GoDoc](https://godoc.org/github.com/reiver/go-cbor?status.svg)](https://godoc.org/github.com/reiver/go-cbor)

## JSON versus CBOR

**CBOR** is sometimes compared to **JSON**.

Sometimes CBOR is called a **binary JSON**.

There is some truth to that, but — there are also **a lot** of differences between **CBOR** and **JSON**.

**CBOR** and **JSON** are similar in that — they are both ways of represents hierarchical key-value pairs.

But, **CBOR** has different **basic-types** than **JSON**.

## Import

To import package **cbor** use `import` code like the following:
```
import "github.com/reiver/go-cbor"
```

## Installation

To install package **cbor** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-cbor
```

## Author

Package **cbor** was written by [Charles Iliya Krempeaux](http://reiver.link)
