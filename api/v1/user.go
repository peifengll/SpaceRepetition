package v1

import "github.com/peifengll/SpaceRepetition/internal/model"

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" example:"alan"`
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
}
type GetProfileResponseData struct {
	UserId   string `json:"userId"`
	Nickname string `json:"nickname" example:"alan"`
}
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}

type UserUpdateReq struct {
	Gender           int64   `gorm:"column:gender;type:bigint;comment:性别" json:"gender"` // 性别
	Age              int64   `gorm:"column:age;type:bigint" json:"age"`
	Username         string  `gorm:"column:username;type:longtext;comment:登录的用户名" json:"username"` // 登录的用户名
	Introduction     string  `gorm:"column:introduction;type:longtext" json:"introduction"`
	MaxInterval      int64   `gorm:"column:max_interval;type:int;comment:最大复习间隔，超过这个就算复习完成（推荐365） 默认365" json:"maxInterval"`                  // 最大复习间隔，超过这个就算复习完成（推荐365） 默认365
	Weights          string  `gorm:"column:weights;type:varchar(80);comment:复习计算需要的参数" json:"weights"`                                        // 复习计算需要的参数
	RequestRetention float32 `gorm:"column:request_retention;type:float;comment:想要下次见到一张卡片时，回忆起的概率（推荐70%~90%,默认90%）" json:"requestRetention"` // 想要下次见到一张卡片时，回忆起的概率（推荐70%~90%,默认90%）
}

type CalDeckNum struct {
	UserId string `gorm:"column:user_id"`
	Num    int64  `gorm:"column:num"`
}

type UserWithNum struct {
	*model.User
	Num int64
}
