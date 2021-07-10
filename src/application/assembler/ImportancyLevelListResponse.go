package assembler

import "JStock/src/application/dto"

type ImportancyLevelListResponse struct{}

func (s *ImportancyLevelListResponse) D2M_LevelList(agg interface{}) *dto.ImportancyLevelListResponse {
	m := make(dto.ImportancyLevelListResponse, 0)
	return &m
}
