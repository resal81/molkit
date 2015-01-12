package gmx

import (
	"bytes"
	"io"
	"os"

	"github.com/resal81/molkit/ff"
)

func WriteTopToFile(top *ff.TopSystem, forcefield *ff.ForceField, fname string) error {
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	return writetop(top, forcefield, file)
}

func WriteTopToString(top *ff.TopSystem, forcefield *ff.ForceField) (string, error) {
	var out bytes.Buffer

	err := writetop(top, forcefield, &out)
	if err != nil {
		return "", err
	}

	return out.String(), nil

}

func writetop(top *ff.TopSystem, forcefield *ff.ForceField, writer io.Writer) error {
	writer.Write([]byte("hello world"))
	return nil
}
