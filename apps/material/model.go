package material

import "fmt"

type Material struct {
	Id             int
	MaterialNumber string
	Plant          string
	MaterialName   string
}

func (m *Material) String() string {
	return fmt.Sprintf("<%s, %s>", m.MaterialName, m.MaterialNumber)
}
