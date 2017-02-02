package fs

import (
	"errors"
	"github.com/fatih/structs"
)

func (fs *Fscom) RdbeNum() int {
	return len(fs.RdbeTsysData[:])
}

func (fs *Fscom) RdbeMap(index int) (m map[string]interface{}, err error) {
	if index < 0 || index >= fs.RdbeNum() {
		return nil, errors.New("invalid RDBE id")
	}

	i := fs.RdbeTsysData[index].Iping
	if i < 0 || int(i) >= len(fs.RdbeTsysData[index].Data) {
		return nil, errors.New("no data available")
	}

	return structs.Map(fs.RdbeTsysData[index].Data[i]), nil
}

// Returns a function which returns if RDBE info has been updated
// since last check
func (fs *Fscom) RdbeUpdatedFn(index int) (func() bool, error) {
	if index < 0 || index >= fs.RdbeNum() {
		return nil, errors.New("invalid RDBE index")
	}
	i := -1
	return func() bool {
		j := int(fs.RdbeTsysData[index].Iping)
		if j < 0 || j >= len(fs.RdbeTsysData[index].Data) || i == j {
			return false
		}
		i = j
		return true
	}, nil

}
