// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

// Config is the golang structure for table config.
type Config struct {
	Id           int    `orm:"id,primary"    json:"id"`           // 主键
	Name         string `orm:"name"          json:"name"`         // 名称
	Key          string `orm:"key"           json:"key"`          // 键
	Value        string `orm:"value"         json:"value"`        // 值
	Code         string `orm:"code"          json:"code"`         // 编码
	DataType     int    `orm:"data_type"     json:"dataType"`     // 数据类型//radio/1,KV,2,字典,3,数组
	ParentId     int    `orm:"parent_id"     json:"parentId"`     // 类型
	ParentKey    string `orm:"parent_key"    json:"parentKey"`    //
	Sort         int    `orm:"sort"          json:"sort"`         // 排序号
	ProjectId    int    `orm:"project_id"    json:"projectId"`    // 项目ID
	CopyStatus   int    `orm:"copy_status"   json:"copyStatus"`   // 拷贝状态 1 拷贝  2  不拷贝
	ChangeStatus int    `orm:"change_status" json:"changeStatus"` // 1 不可更改 2 可以更改
	Enable       int    `orm:"enable"        json:"enable"`       // 是否启用//radio/1,启用,2,禁用
	UpdateTime   string `orm:"update_time"   json:"updateTime"`   // 更新时间
	UpdateId     int    `orm:"update_id"     json:"updateId"`     // 更新人
	CreateTime   string `orm:"create_time"   json:"createTime"`   // 创建时间
	CreateId     int    `orm:"create_id"     json:"createId"`     // 创建者
}

// Department is the golang structure for table department.
type Department struct {
	Id         int    `orm:"id,primary"  json:"id"`         // 主键
	ParentId   int    `orm:"parent_id"   json:"parentId"`   // 上级机构
	Name       string `orm:"name,unique" json:"name"`       // 部门/11111
	Code       string `orm:"code"        json:"code"`       // 机构编码
	Sort       int    `orm:"sort"        json:"sort"`       // 序号
	Linkman    string `orm:"linkman"     json:"linkman"`    // 联系人
	LinkmanNo  string `orm:"linkman_no"  json:"linkmanNo"`  // 联系人电话
	Remark     string `orm:"remark"      json:"remark"`     // 机构描述
	Enable     int    `orm:"enable"      json:"enable"`     // 是否启用//radio/1,启用,2,禁用
	UpdateTime string `orm:"update_time" json:"updateTime"` // 更新时间
	UpdateId   int    `orm:"update_id"   json:"updateId"`   // 更新人
	CreateTime string `orm:"create_time" json:"createTime"` // 创建时间
	CreateId   int    `orm:"create_id"   json:"createId"`   // 创建者
}

// Log is the golang structure for table log.
type Log struct {
	Id         int    `orm:"id,primary"  json:"id"`         // 主键
	LogType    int    `orm:"log_type"    json:"logType"`    // 类型
	OperObject string `orm:"oper_object" json:"operObject"` // 操作对象
	OperTable  string `orm:"oper_table"  json:"operTable"`  // 操作表
	OperId     int    `orm:"oper_id"     json:"operId"`     // 操作主键
	OperType   string `orm:"oper_type"   json:"operType"`   // 操作类型
	OperRemark string `orm:"oper_remark" json:"operRemark"` // 操作备注
	Enable     int    `orm:"enable"      json:"enable"`     // 是否启用//radio/1,启用,2,禁用
	UpdateTime string `orm:"update_time" json:"updateTime"` // 更新时间
	UpdateId   int    `orm:"update_id"   json:"updateId"`   // 更新人
	CreateTime string `orm:"create_time" json:"createTime"` // 创建时间
	CreateId   int    `orm:"create_id"   json:"createId"`   // 创建者
}

// Menu is the golang structure for table menu.
type Menu struct {
	Id         int    `orm:"id,primary"  json:"id"`         // 主键
	ParentId   int    `orm:"parent_id"   json:"parentId"`   // 父id
	Name       string `orm:"name"        json:"name"`       // 名称/11111
	Icon       string `orm:"icon"        json:"icon"`       // 菜单图标
	Urlkey     string `orm:"urlkey"      json:"urlkey"`     // 菜单key
	Url        string `orm:"url"         json:"url"`        // 链接地址
	Perms      string `orm:"perms"       json:"perms"`      // 授权(多个用逗号分隔，如：user:list,user:create)
	Status     int    `orm:"status"      json:"status"`     // 状态//radio/2,隐藏,1,显示
	Type       int    `orm:"type"        json:"type"`       // 类型//select/1,目录,2,菜单,3,按钮
	Sort       int    `orm:"sort"        json:"sort"`       // 排序
	Level      int    `orm:"level"       json:"level"`      // 级别
	Enable     int    `orm:"enable"      json:"enable"`     // 是否启用//radio/1,启用,2,禁用
	UpdateTime string `orm:"update_time" json:"updateTime"` // 更新时间
	UpdateId   int    `orm:"update_id"   json:"updateId"`   // 更新人
	CreateTime string `orm:"create_time" json:"createTime"` // 创建时间
	CreateId   int    `orm:"create_id"   json:"createId"`   // 创建者
}

// Role is the golang structure for table role.
type Role struct {
	Id         int    `orm:"id,primary"  json:"id"`         // 主键
	Name       string `orm:"name"        json:"name"`       // 名称/11111/
	Status     int    `orm:"status"      json:"status"`     // 状态//radio/2,隐藏,1,显示
	Sort       int    `orm:"sort"        json:"sort"`       // 排序
	Remark     string `orm:"remark"      json:"remark"`     // 说明//textarea
	Enable     int    `orm:"enable"      json:"enable"`     // 是否启用//radio/1,启用,2,禁用
	UpdateTime string `orm:"update_time" json:"updateTime"` // 更新时间
	UpdateId   int    `orm:"update_id"   json:"updateId"`   // 更新人
	CreateTime string `orm:"create_time" json:"createTime"` // 创建时间
	CreateId   int    `orm:"create_id"   json:"createId"`   // 创建者
}

// RoleMenu is the golang structure for table role_menu.
type RoleMenu struct {
	Id     int `orm:"id,primary" json:"id"`     // 主键
	RoleId int `orm:"role_id"    json:"roleId"` // 角色id
	MenuId int `orm:"menu_id"    json:"menuId"` // 菜单id
}

// User is the golang structure for table user.
type User struct {
	Id           int    `orm:"id,primary"      json:"id"`           // 主键
	Uuid         string `orm:"uuid"            json:"uuid"`         // UUID
	Username     string `orm:"username,unique" json:"username"`     // 登录名/11111
	Password     string `orm:"password"        json:"password"`     // 密码
	Salt         string `orm:"salt"            json:"salt"`         // 密码盐
	RealName     string `orm:"real_name"       json:"realName"`     // 真实姓名
	DepartId     int    `orm:"depart_id"       json:"departId"`     // 部门/11111/dict
	UserType     int    `orm:"user_type"       json:"userType"`     // 类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户
	Status       int    `orm:"status"          json:"status"`       // 状态
	Thirdid      string `orm:"thirdid"         json:"thirdid"`      // 第三方ID
	Endtime      string `orm:"endtime"         json:"endtime"`      // 结束时间
	Email        string `orm:"email"           json:"email"`        // email
	Tel          string `orm:"tel"             json:"tel"`          // 手机号
	Address      string `orm:"address"         json:"address"`      // 地址
	TitleUrl     string `orm:"title_url"       json:"titleUrl"`     // 头像地址
	Remark       string `orm:"remark"          json:"remark"`       // 说明
	Theme        string `orm:"theme"           json:"theme"`        // 主题
	BackSiteId   int    `orm:"back_site_id"    json:"backSiteId"`   // 后台选择站点ID
	CreateSiteId int    `orm:"create_site_id"  json:"createSiteId"` // 创建站点ID
	ProjectId    int    `orm:"project_id"      json:"projectId"`    // 项目ID
	ProjectName  string `orm:"project_name"    json:"projectName"`  // 项目名称
	Enable       int    `orm:"enable"          json:"enable"`       // 是否启用//radio/1,启用,2,禁用
	UpdateTime   string `orm:"update_time"     json:"updateTime"`   // 更新时间
	UpdateId     int    `orm:"update_id"       json:"updateId"`     // 更新人
	CreateTime   string `orm:"create_time"     json:"createTime"`   // 创建时间
	CreateId     int    `orm:"create_id"       json:"createId"`     // 创建者
}

// UserRole is the golang structure for table user_role.
type UserRole struct {
	Id     int `orm:"id,primary" json:"id"`     // 主键
	UserId int `orm:"user_id"    json:"userId"` // 用户id
	RoleId int `orm:"role_id"    json:"roleId"` // 角色id
}
