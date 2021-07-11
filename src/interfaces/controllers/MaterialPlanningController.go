package controllers

type MaterialPlaningController struct {
	MaterialPlaningService *services.MaterialPlaningService `inject:"-"`
}
