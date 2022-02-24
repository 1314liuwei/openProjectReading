// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"focus-single/internal/service/internal/dao/internal"
)

// fileDao is the data access object for table gf_file.
// You can define custom methods on it to extend its functionality as you wish.
type fileDao struct {
	*internal.FileDao
}

var (
	// File is globally public accessible object for table gf_file operations.
	File = fileDao{
		internal.NewFileDao(),
	}
)

// Fill with you ideas below.
