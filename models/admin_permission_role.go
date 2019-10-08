package models

type AdminPermissionRole struct {
	PermissionId uint `orm:"column(permission_id)"`
	RoleId       uint `orm:"column(role_id)"`
}
