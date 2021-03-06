// Copyright 2014 Rana Ian. All rights reserved.
// Use of this source code is governed by The MIT License
// found in the accompanying LICENSE file.

package ora

/*
#include <stdlib.h>
#include <oci.h>
#include "version.h"
*/
import "C"
import (
	"unsafe"
)

type defRaw struct {
	ociDef
	ociRaw     *C.OCIRaw
	isNullable bool
	buf        []byte
	columnSize int
}

func (def *defRaw) define(position int, columnSize int, isNullable bool, rset *Rset) error {
	def.rset = rset
	def.isNullable = isNullable
	def.columnSize = columnSize
	def.buf = make([]byte, fetchArrLen*columnSize)

	return def.ociDef.defineByPos(position, unsafe.Pointer(&def.buf[0]), columnSize, C.SQLT_BIN)
}

func (def *defRaw) value(offset int) (value interface{}, err error) {
	if def.isNullable {
		bytesValue := Raw{IsNull: def.nullInds[offset] < 0}
		if !bytesValue.IsNull {
			bytesValue.Value = def.buf[offset*def.columnSize : (offset+1)*def.columnSize]
		}
		value = bytesValue
	} else {
		if def.nullInds[offset] > -1 {
			value = def.buf[offset*def.columnSize : (offset+1)*def.columnSize]
		}
	}
	return value, err
}

func (def *defRaw) alloc() error {
	return nil
}

func (def *defRaw) free() {
}

func (def *defRaw) close() (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = errR(value)
		}
	}()

	rset := def.rset
	def.rset = nil
	def.ocidef = nil
	def.ociRaw = nil
	def.buf = nil
	def.arrHlp.close()
	rset.putDef(defIdxRaw, def)
	return nil
}
