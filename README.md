
## Go

```

go test -coverprofile=cover.out

go test -coverprofile=cover.out -coverpkg github.com/resal81/molkit/ff,github.com/resal81/molkit/molio/gmx 


go tool cover -func=cover.out
go tool cover -html=cover.out

```

## Hierarchy

```
Collection                      : independent systems; e.g. entries in a SDF file

System                          : a multi-chain complex; e.g. a simulation box
    |__ Polymer                 : chain
        |   |__ Fragment        : residue
        |           |__ Atom    :
        |
        |__ Bonds, Angles, Dihedrals, Impropers


- Bond specifies the connection between atoms. A Bond doesn't have direction.



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

```go
sys.DeepCopy()  // returns a deep copy; useful when building topologies

// reading files - implement using io.Reader
molio.ReadPDBFile(fname)
molio.ReadPSFFile(fname)

molio.ReadGroTopFile(sys *System, fname string)

// selection
fnBB := selection.ByAtomName(true, 'CA', 'C', 'N', 'O')
sel  := selection.Select(pdb, fnBB, ...)

// writing files
molio.WritePDB(io.Writer)
molio.WritePSF(io.Writer)


// forcefield
ff := forcefield.NewFF(FF_CHARMM)
ff.ReadFile('...') // based on extension
ff.ReadFile('...')

ff.Apply(sys System)


```

