package usermodel

import "time"

func WithUserNameZh(name string) UserModelMainAttrFunc {
	return func(model *UserModelMain) {
		model.UserInfo.UserNameZh = name
	}
}

func WithUserNameEn(name string) UserModelMainAttrFunc {
	return func(model *UserModelMain) {
		model.UserInfo.UserNameEn = name
	}
}

func WithUserMobilePhone(phone string) UserModelMainAttrFunc {
	return func(model *UserModelMain) {
		model.UserInfo.UserMobilePhone = phone
	}
}

func WithUserTelePhone(phone string) UserModelMainAttrFunc {
	return func(model *UserModelMain) {
		model.UserInfo.UserTelePhone = phone
	}
}

func WithUserEmail(email string) UserModelMainAttrFunc {
	return func(model *UserModelMain) {
		model.UserInfo.UserEmail = email
	}
}

func WithUserRemark(remark string) UserModelMainAttrFunc {
	return func(model *UserModelMain) {
		model.UserInfo.Remark = remark
	}
}

type VUserInfo struct {
	UserNameZh      string    `json:"user_name_zh" gorm:"column:user_name_zh"`
	UserNameEn      string    `json:"user_name_en" gorm:"column:user_name_en"`
	UserMobilePhone string    `json:"user_mobile_phone" gorm:"column:user_mobile_phone"`
	UserTelePhone   string    `json:"user_telephone" gorm:"column:user_telephone"`
	UserEmail       string    `json:"user_email" gorm:"column:user_email"`
	Remark          string    `json:"remark" gorm:"column:remark"`
	CreateAt        time.Time `json:"create_at" gorm:"column:=create_at"`
}

func NewVUserInfo() *VUserInfo {
	return &VUserInfo{}
}
