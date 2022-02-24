// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"focus-single/internal/service/internal/dao/internal"
)

// userDao is the data access object for table gf_user.
// You can define custom methods on it to extend its functionality as you wish.
type userDao struct {
	*internal.UserDao
}

var (
	// User is globally public accessible object for table gf_user operations.
	User = userDao{
		internal.NewUserDao(),
	}
)

// Fill with you ideas below.
