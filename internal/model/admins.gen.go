// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAdmin = "admins"

// Admin mapped from table <admins>
type Admin struct {
	ID         int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Username   string `gorm:"column:username;type:varchar(50);not null" json:"username"`
	Password   string `gorm:"column:password;type:varchar(100);not null" json:"password"`
	Email      string `gorm:"column:email;type:varchar(100);not null" json:"email"`
	Name       string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Phone      string `gorm:"column:phone;type:varchar(20);not null" json:"phone"`
	Privileges int64  `gorm:"column:privileges;type:int;not null;comment:0是终极权限  还可撤销别人的一级权限，但不能赋为0权限，可删除其他人的管理员账号 , 1级是完整权限 ，可更改2级为1级，2级是部分" json:"privileges"` // 0是终极权限  还可撤销别人的一级权限，但不能赋为0权限，可删除其他人的管理员账号 , 1级是完整权限 ，可更改2级为1级，2级是部分
}

// TableName Admin's table name
func (*Admin) TableName() string {
	return TableNameAdmin
}