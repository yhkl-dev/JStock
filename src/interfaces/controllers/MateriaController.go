package controllers

import (
	"JStock/src/application/services"
)


type MaterialController struct {
	MaterialService *services.MaterialService `inject:"-"`
}