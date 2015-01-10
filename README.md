


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

FFAtom
    bonds

FFResidue
    atoms  
    linkerAtoms

FFCap
    atoms
    linkerAtom


ForceField
    |__ AtomParams
    |__ BondParams
    |__ AngleParams
    |__ DihedralParams
    |__ ImproperParams



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

