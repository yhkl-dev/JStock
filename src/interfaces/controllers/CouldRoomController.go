package controllers

import "JStock/src/application/services"

type CloudRoomController struct {
	MaterialService *services.MaterialService `inject:"-"`
}
