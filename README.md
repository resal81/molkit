[![Build Status](https://travis-ci.org/resal81/molkit.svg?branch=master)](https://travis-ci.org/resal81/molkit)
[![codecov.io](http://codecov.io/github/resal81/molkit/coverage.svg?branch=master)](http://codecov.io/github/resal81/molkit?branch=master)

# Molkit: A molecular manipulation toolkit written in Go

This is a work-in-progress kit for manipulating molecular structures such as 
topologies.


## Installation
- Make sure you have Go installed. More information [here](https://golang.org/doc/install).
- Then:

```bash
go get -U github.com/resal81/molkit
```


## Hierarchy
```
AtomType    ->  Atom    -> Fragment -> Chain -> Complex
Element     ->  Atom


```


