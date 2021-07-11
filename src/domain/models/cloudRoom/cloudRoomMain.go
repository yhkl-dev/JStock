package cloudroom

import (
	"JStock/src/domain/models/repos"
	"time"
)

type CloudRoomModel struct {
	ID          int                   `json:"id" gorm:"column:id"`
	CloudRoomID string                `json:"cloud_root_id" gorm:"column:cloud_root_id"`
	RoomName    string                `json:"room_name" gorm:"column:room_name"`
	Owner       string                `json:"owner" gorm:"column:owner"`
	CreateAt    time.Time             `json:"create_at" gorm:"column:create_at"`
	Repo        repos.ICloudRoomModel `gorm:"-"`
}

func NewCloudRoomModel() *CloudRoomModel {
	return &CloudRoomModel{}
}

func (*CloudRoomModel) Name() string {
	return "CloudRoomModel"
}

func (s *CloudRoomModel) New() error {
	return s.Repo.New(s)
}

func (s *CloudRoomModel) Load() error {
	return s.Repo.Load(s)
}

func (s *CloudRoomModel) ListCloudRoom(page, pageSize int) (interface{}, error) {
	return s.Repo.ListCloudRoom(s, page, pageSize)
}
