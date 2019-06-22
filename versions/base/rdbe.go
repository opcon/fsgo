// Copyright 2019 NVI Inc. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

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
		err = errors.New("invalid RDBE id")
		return
	}

	i := int(f.Fscom.RdbeTsysData[index].Iping)
	if i < 0 || i >= len(f.Fscom.RdbeTsysData[index].Data) {
		err = errors.New("no data available")
		return
	}

	data := f.Fscom.RdbeTsysData[index].Data[i]
	m = structs.Map(data)
	m["Epoch"] = cstr(data.Epoch[:])

	m["PcalAmp"] = data.PcalAmp[:]
	m["PcalPhase"] = data.PcalPhase[:]

	tsysmap := make([][]float32, len(data.Tsys))
	for i := range data.Tsys {
		tsysmap[i] = data.Tsys[i][:]
	}
	m["Tsys"] = tsysmap

	return m, nil
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
