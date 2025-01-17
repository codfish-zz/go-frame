// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package oracle implements gdb.Driver, which supports operations for database Oracle.
//
// Note:
// 1. It does not support Save/Replace features.
// 2. It does not support LastInsertId.
package oracle

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

// Driver is the driver for oracle database.
type Driver struct {
	*gdb.Core
}

const (
	quoteChar = `"`
)

func init() {
	if err := gdb.Register(`oracle`, New()); err != nil {
		panic(err)
	}
}

// formatSqlTmp formats sql template string into one line.
func formatSqlTmp(sqlTmp string) string {
	var err error
	// format sql template string.
	sqlTmp, err = gregex.ReplaceString(`[\n\r\s]+`, " ", gstr.Trim(sqlTmp))
	if err != nil {
		panic(err)
	}
	sqlTmp, err = gregex.ReplaceString(`\s{2,}`, " ", gstr.Trim(sqlTmp))
	if err != nil {
		panic(err)
	}
	return sqlTmp
}

// New create and returns a driver that implements gdb.Driver, which supports operations for Oracle.
func New() gdb.Driver {
	return &Driver{}
}

// New creates and returns a database object for oracle.
// It implements the interface of gdb.Driver for extra database driver installation.
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// GetChars returns the security char for this type of database.
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
