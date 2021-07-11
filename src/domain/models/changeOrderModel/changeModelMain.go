package changemodel

import "time"

type ChangeForm struct {
	ID          int                 `json:"id"`
	Source      *NetWorkChangeModel `gorm:""`
	Destination *NetWorkChangeModel `gorm:""`
	Purpose     string              `json:"purpose"`
	Status      int                 `json:"status"`
}

type NetWorkChangeModel struct {
	ID             int       `json:"id"`
	Owner          string    `json:"owner"`
	AppName        string    `json:"app_name"`
	CloudRoomID    int       `json:"cloud_room_id"`
	ResourceTypeID int       `json:"resource_type"`
	ResourceIP     string    `json:"resource_ip"`
	Protocal       string    `json:"protocal"`
	Port           uint      `json:"port"`
	CreateAt       time.Time `json:"create_at"`
}

type CloudRoom struct {
	ID          int    `json:"id"`
	RoomName    string `json:"room_name"`
	CloudRoomID string `json:"cloud_root_id"`
	Owner       string `json:"owner"`
	CreateAt    string `json:"create_at"`
}

type ResourceTypeID struct {
	ID          int    `json:"id"`
	TypeName    string `json:"type_name"`
	TypeComment string `json:"type_comment"`
}

type MainModel struct {
	ID       int    `json:"id"`
	Sender   string `json:"sender"`
	CC       string `json:"cc"`
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}

type MainContentTemplate struct {
	ID              int       `json:"id"`
	ContentTemplate string    `json:"content_tmeplate"`
	CreatAt         time.Time `json:"create_at"`
}
