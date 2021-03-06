// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gfast/app/system/dao/internal"
	"gfast/app/system/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// sysAuthRuleDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysAuthRuleDao struct {
	internal.SysAuthRuleDao
}

var (
	// SysAuthRule is globally public accessible object for table sys_auth_rule operations.
	SysAuthRule = sysAuthRuleDao{
		internal.SysAuthRule,
	}
)

func (d *sysAuthRuleDao) Scope(f func(m gmvc.M) gmvc.M) *sysAuthRuleDao {
	nd := *d
	if m := f(d.M); m != nil {
		nd.M = m
	}
	return &nd
}

// Fill with you ideas below.
func (d *sysAuthRuleDao) GetMenuList() (list []*model.SysAuthRuleInfoRes, err error) {

	err = d.Fields(model.SysAuthRuleInfoRes{}).Order("weigh desc,id asc").Scan(&list)
	return
}

func (d *sysAuthRuleDao) CheckMenuNameUnique(name string, id int) bool {

	m := d.Where("name=?", name)
	if id != 0 {
		m = m.And("id!=?", id)
	}
	c, err := m.Count()
	if err != nil {
		g.Log().Error(err)
		return false
	}
	return c == 0
}

//检查菜单路由地址是否已经存在
func (d *sysAuthRuleDao) CheckMenuPathUnique(path string, id int) bool {

	model := d.Where("path=?", path).Where("menu_type<>?", 2)
	if id != 0 {
		model = model.And("id!=?", id)
	}
	c, err := model.Count()
	if err != nil {
		g.Log().Error(err)
		return false
	}
	return c == 0
}
