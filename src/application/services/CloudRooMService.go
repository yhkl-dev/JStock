package services

import (
	"JStock/src/application/assembler"

	"gorm.io/gorm"
)

type CloudRoomService struct {
	AssCloudRoomReq *assembler.CloudRoomRequest
	DB              *gorm.DB `inject:"-"`
}
