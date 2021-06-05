package userModel

import "time"

type User struct {
	ID              int       `json:"id" gorm:"column:id"`
	UserID          string    `json:"user_id" gorm:"column:user_id"`
	UserNameZh      string    `json:"user_name_zh" binding:"username"`
	UserNameEn      string    `json:"user_name_en"`
	UserPassword    string    `json:"password"`
	UserMobilePhone string    `json:"user_mobile_phone"`
	UserTelePhone   string    `json:"user_telephone" gorm:"column:user_telephone"`
	UserMail        string    `json:"user_mail"`
	CreateAt        time.Time `json:"create_at"`
	Remark          string    `json:"remark"`
}

func (u *User) TableName() string {
	return "t_users"
}

func New(attr ...UserAttrFunc) *User {
	u := &User{}
	UserAttrFuncs(attr).Apply(u)
	return u
}

func (s *User) Mutate(attr ...UserAttrFunc) *User {
	UserAttrFuncs(attr).Apply(s)
	return s
}
