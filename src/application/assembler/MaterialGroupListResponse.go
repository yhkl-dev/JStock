package assembler

import (
	"JStock/src/application/dto"
	materialmodel "JStock/src/domain/models/materialModel"
)

type MaterialGroupListResponse struct{}

func (s *MaterialGroupListResponse) D2M_MaterialGroupList(agg interface{}) *dto.MaterialGroupListResponse {
	groupList := make(dto.MaterialGroupListResponse, 0)
	resGroupList := agg.([]*materialmodel.MaterialGroup)
	for _, m := range resGroupList {
		info := &dto.MaterialGroupMainResponse{}
		info.ID = m.ID
		info.GroupName = m.GroupName
		info.Comment = m.Comment
		groupList = append(groupList, *info)
	}
	return &groupList
}
