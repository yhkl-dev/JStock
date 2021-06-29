package usermodel

import (
	"time"
)

func WithCreateAt(t time.Time) UserModelMainAttrFunc {
	return func(model *UserModelMain) {
		model.CreateAt.CreateAt = t
	}
}

type VUserCreatAt struct {
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
}

func (s *VUserCreatAt) GetFormat() string {
	return s.String()
}

func (s *VUserCreatAt) String() string {
	return s.CreateAt.Format("2006-01-02 15:04:05")
}

func NewVUserCreatAt() *VUserCreatAt {
	return &VUserCreatAt{}
}
