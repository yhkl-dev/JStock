package assembler

import (
	"JStock/src/application/dto"
	cloudroom "JStock/src/domain/models/cloudRoom"
)

type CloudRoomRequest struct{}

func (s *CloudRoomRequest) D2M_CloudRoomModel(dto *dto.CloudRoomMainRequest) *cloudroom.CloudRoomModel {
	return cloudroom.NewCloudRoomModel()
}
