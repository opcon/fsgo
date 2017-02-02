package fs

import (
	"errors"
	"github.com/fatih/structs"
)

func (f *FieldSystem) RdbeNum() int {
	return len(f.Fscom.RdbeTsysData[:])
}

func (f *FieldSystem) RdbeMap(index int) (m map[string]interface{}, err error) {
	if index < 0 || index >= f.RdbeNum() {
		return nil, errors.New("invalid RDBE id")
	}

	i := f.Fscom.RdbeTsysData[index].Iping
	if i < 0 || int(i) >= len(f.Fscom.RdbeTsysData[index].Data) {
		return nil, errors.New("no data available")
	}

	return structs.Map(f.Fscom.RdbeTsysData[index].Data[i]), nil
}

// Returns a function which returns if RDBE info has been updated
// since last check
func (f *FieldSystem) RdbeUpdatedFn(index int) (func() bool, error) {
	if index < 0 || index >= f.RdbeNum() {
		return nil, errors.New("invalid RDBE index")
	}
	i := -1
	return func() bool {
		j := int(f.Fscom.RdbeTsysData[index].Iping)
		if j < 0 || j >= len(f.Fscom.RdbeTsysData[index].Data) || i == j {
			i = j // For case when j < 0
			return false
		}
		i = j
		return true
	}, nil

}
