package utils

import uuid "github.com/satori/go.uuid"

func GenerateMaterialID() string {
	return uuid.NewV1().String()
}
