package departmentmodel

import "time"

type Department struct {
	ID             int
	ParentID       int
	DepartmentName string
	Description    string
	CreateAt       time.Time
}
