package sapmapping

import (
	"fmt"
	"time"
)

type SAP struct {
	ID            int
	LocalVairable string
	SapVariable   string
	Comments      string
	Created       time.Time
}

func (s *SAP) String() string {
	return fmt.Sprintf("<%s: %s>", s.LocalVairable, s.SapVariable)
}
