


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

AtomParam
    GetA()
    GetEpsilon()
    ...
BondParam
AngleParam
DihedralParam
ImproperParam


```

```go
sys.DeepCopy()  // returns a deep copy; useful when building topologies

// reading files
molio.ReadPDB(io.Reader)
molio.ReadPSF(io.Reader)

molio.ReadPDBGroTop(io.Reader, io.Reader)

// selection
fnBB := selection.ByAtomName('CA', 'C', 'N', 'O')
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

