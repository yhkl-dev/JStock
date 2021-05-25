package material

import "fmt"

type Material struct {
	ID             int
	MaterialNumber string
	Plant          string
	MaterialName   string
}

func (m *Material) String() string {
	return fmt.Sprintf("<%s, %s>", m.MaterialName, m.MaterialNumber)
}
