package model

import "admin-permission/config"

type base struct {
	Id        int   `gorm:"column:id;primaryKey;autoIncrement;type:int"`
	CreatedAt int64 `gorm:"column:created_at;type:int;size:10;autoCreateTime"`
	UpdatedAt int64 `gorm:"column:updated_at;type:int;size:10;autoUpdateTime"`
}

// 管理员主表
type AdminModel struct {
	base
	Account       string `gorm:"column:account;type:string;size:20;index;unique;not null;comment:账号"`
	Password      string `gorm:"column:password;type:string;size:64;not null;comment:密码"`
	Name          string `gorm:"column:name;type:string;size:10;comment:真实姓名"`
	Mobile        string `gorm:"column:mobile;type:string;size:11;comment:手机号码"`
	Email         string `gorm:"column:email;type:string;size:30;comment:电子邮箱"`
	Group         int    `gorm:"column:group;type:int;size:5;default:0;index;comment:用户组"`
	Status        int    `gorm:"column:status;type:tinyInt;size:5;default:1;comment:状态，0-冻结，1-可用，-1-删除"`
	LastLoginIp   int    `gorm:"column:last_login_ip;type:int;size:10;comment:最后登录IP"`
	LastLoginTime int    `gorm:"column:last_login_time;type:int;size:10;comment:最后登录时间"`
}

func (AdminModel) TableName() string {
	return config.GetDbPrefix() + "admin"
}

// 路由表
type RouteModel struct {
	base
	Url string `gorm:"column:url;type:string;size:64;unique"`
}

func (RouteModel) TableName() string {
	return config.GetDbPrefix() + "route"
}

// 管理员权限表
type AdminPrivilegeModel struct {
	base
	AdminId   int    `gorm:"column:admin_id;type:int;size:5;not null;comment:管理员ID"`
	Privilege string `gorm:"column:privilege;type:text;comment:权限表"`
}

func (AdminPrivilegeModel) TableName() string {
	return config.GetDbPrefix() + "admin_privileges"
}

// 用户组表
type GroupModel struct {
	base
	Name          string `gorm:"column:name;type:string;size:20;not null;comment:用户组名称"`
	Mark          string `gorm:"column:mark;type:string;size:20;not null;comment:用户组标识，唯一"`
	Parent        int    `gorm:"column:parent;type:int;size:5;default:0;not null;comment:上级用户组"`
	DataPrivilege int    `gorm:"column:data_privilege;type:tinyInt;size:1;default:0;comment:数据权限，0-所有数据，1-该组数据，2-仅自己的数据"`
	Privilege     string `gorm:"column:privilege;type:text;comment:权限表"`
}

func (GroupModel) TableName() string {
	return config.GetDbPrefix() + "groups"
}

// 菜单表
type MenuModel struct {
	base
	Title  string `gorm:"column:title;type:string;size:15;comment:菜单名称"`
	PageId int    `gorm:"column:page_id;type:int;size:5;comment:页面id"`
	Parent int    `gorm:"column:parent;type:int;size:5;comment:父菜单ID"`
	Level  int    `gorm:"column:level;type:int;size:1;comment:菜单等级"`
	Sort   int    `gorm:"column:sort;type:int;size:5;comment:排序"`
	Status int    `gorm:"column:status;type:int;size:1;comment:状态"`
}

func (MenuModel) TableName() string {
	return config.GetDbPrefix() + "menu"
}

// 页面表
type PageModel struct {
	base
	Name    string `gorm:"column:name;type:string;size:20;not null;comment:页面名称"`
	Mark    string `gorm:"column:mark;type:string;size:20;unique;comment:页面标识，唯一"`
	PageUrl string `gorm:"column:page_url;type:string;size:64;comment:页面URL"`
	ApiUrl  string `gorm:"column:api_url;type:string;size:64;comment:api url"`
	Type    int    `gorm:"column:type;type:tinyInt;size:1;comment:页面类型，1-页面，2-功能"`
	Parent  int    `gorm:"column:parent;type:int;size:5;default:0;comment:父页面"`
}

func (PageModel) TableName() string {
	return config.GetDbPrefix() + "pages"
}
