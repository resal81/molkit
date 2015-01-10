


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
- 

sys.DeepCopy()  // returns a deep copy; useful when building topologies

molio.WritePDB(io.Writer)
molio.WritePSF(io.Writer)

molio.ReadPDB(io.Reader)
molio.ReadPSF(io.Reader)


ff := forcefield.NewFF(FF_CHARMM)
ff.ReadFile('...') // based on extension
ff.ReadFile('...')




AtomParam
    GetA()
    GetEpsilon()
    ...
BondParam
AngleParam
DihedralParam
ImproperParam

```

## Example PDB read

```
// 
pdb := ParsePDB(source io.Reader)

fnBB := selection.ByAtomName('CA', 'C', 'N', 'O')
sel  := selection.Select(pdb, fnBB)

```


## Example CHARMM FF read

```



```

## Example solute insertion using PDB/TOP file

```

```
