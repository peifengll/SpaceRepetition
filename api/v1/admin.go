package v1

// 注册信息
type AdminRegisterReq struct {
	Username   string `gorm:"column:username;type:varchar(50);not null" json:"username"`
	Password   string `gorm:"column:password;type:varchar(100);not null" json:"password"`
	Email      string `gorm:"column:email;type:varchar(100);not null" json:"email"`
	Name       string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Phone      string `gorm:"column:phone;type:varchar(20);not null" json:"phone"`
	Privileges int64  `gorm:"column:privileges;type:int;not null;" json:"privileges"`
}

// 更改权限
type UpdatePrivilegesReq struct {
	ID         int64 ` json:"id"`         // 要改的人的id
	SelfId     int64 `json:"self_id"`     // 改东西的人的id
	Privileges int64 ` json:"privileges"` // 要改成什么权限
}

type LoginAdminReq struct {
	Username string ` json:"username"`
	Password string ` json:"password"`
}

type LoginAdminResponse struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Privileges int64  `json:"privileges"`
}

type DelAdminReq struct {
	ID     int64 `json:"id"`
	SelfId int64 `json:"self_id"` // 改东西的人的id
}
