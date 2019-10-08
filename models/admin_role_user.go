package models

type AdminRoleUser struct {
	AdminUserId uint `orm:"column(admin_user_id)"`
	RoleId      uint `orm:"column(role_id)"`
}
