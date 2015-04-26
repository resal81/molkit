[![Build Status](https://travis-ci.org/resal81/molkit.svg?branch=master)](https://travis-ci.org/resal81/molkit)
[![Coverage Status](https://coveralls.io/repos/resal81/molkit/badge.svg)](https://coveralls.io/r/resal81/molkit)

# Molkit: A molecular manipulation toolkit written in Go

This is a work-in-progress kit for manipulating molecular structures such as 
topologies.


## Installation
- Make sure you have Go installed. More information [here](https://golang.org/doc/install).
- Then:

```bash
go get -U github.com/resal81/molkit
```


## Details

### Structure Hierarchy

```
Collection                      : independent systems; e.g. entries in a SDF file

System                          : a multi-chain complex; e.g. a simulation box
    |__ Polymer                 : chain
        |   |__ Fragment        : residue
        |           |__ Atom    :
        |
        |__ Bonds, Angles, Dihedrals, Impropers


Bond : specifies the connection between atoms. A Bond doesn't have direction.



TopPolymer
    |__ TopFragment
    |       |__ TopAtom
    |__ TopPairs
    |__ TopExclusions
    |__ TopSETTLE
    |__ TopPositionRestraints

TopolDB
    |__ TopFragment
            |__ TopAtom
            |__ InternalCoordinate

ParamsDB
    |__ ParamAtomTypes
    |__ ParamBondTypes
    |__ ParamAngleTypes
    |__ ParamDihedralTypes
    |__ ParamConstraintTypes
    |__ ParamNonBondedTypes

ForceField
    |__ TopolDB
    |__ ParamsDB
```

### Testing

```
go test -coverprofile=cover.out

go test -coverprofile=cover.out -coverpkg github.com/resal81/molkit/ff,github.com/resal81/molkit/molio/gmx 


go tool cover -func=cover.out
go tool cover -html=cover.out


go test -cpuprofile=cpu.out 
go tool pprof -pdf [executable] cpu.out > out.pdf
```



