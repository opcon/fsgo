package fs

import (
	"errors"
	"github.com/fatih/structs"
)

func (fs *Fscom) RdbeNum() int {
	return len(fs.Rdbe_tsys_data[:])
}

func (fs *Fscom) RdbeMap(index int) (m map[string]interface{}, err error) {
	if index < 0 || index >= fs.NumRdbes {
		return nil, errors.New("invalid RDBE id")
	}

	i := fs.Rdbe_tsys_data[index].Iping
	if i < 0 || i >= len(fs.Rdbe_tsys_data[index].Data) {
		return nil, errors.New("invalid iping index")
	}

	return structs.Map(fs.Rdbe_tsys_data[index].Data[i]), nil
}

func (fs *Fscom) RdbeUpdatedFn(index int) (func() bool, error) {
	if index < 0 || index >= fs.NumRdbes {
		return nil, errors.New("invalid RDBE index")
	}

	i := fs.Rdbe_tsys_data[index].Iping

	return func() bool {
		j := fs.Rdbe_tsys_data[index].Iping
		if j < 0 || j >= len(fs.Rdbe_tsys_data[index].Data) || i == j {
			return false
		}
		i = j
		return true
	}
}
