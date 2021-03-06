// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gfast/app/system/dao/internal"
)

// toolsGenTableColumnDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type toolsGenTableColumnDao struct {
	*internal.ToolsGenTableColumnDao
}

var (
	// ToolsGenTableColumn is globally public accessible object for table tools_gen_table_column operations.
	ToolsGenTableColumn toolsGenTableColumnDao
)

func init() {
	ToolsGenTableColumn = toolsGenTableColumnDao{
		internal.NewToolsGenTableColumnDao(),
	}
}

// Fill with you ideas below.
