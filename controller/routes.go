package controller

import (
	"github.com/citiaps/pnrrd-formulario-jerarquias/middleware"
	"github.com/citiaps/pnrrd-formulario-jerarquias/model"
	"github.com/gin-gonic/gin"
)

// PaginationParams : Parametros de paginacion
type PaginationParams struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}
type FormParams struct {
	FormType  string `form:"type" binding:"required"`
	Hierarchy string `form:"hierarchy"`
}

type InventoryParams struct {
	Institution  string `form:"institution" binding:"required"`
	Hierarchy    string `form:"hierarchy" binding:"required"`
	Hierarchy_id string `form:"hierarchy_id"`
}

const (
	RolAdmin    = "ADMIN"
	RolOnemi    = "ONEMI"
	RolRegional = "REGIONAL"
	RolNacional = "NACIONAL"
)

const (
	GeneralHierarchyForm = "GENERAL"
	InformationTypeForm  = "INFORMATION"
	ResourceTypeForm     = "RESOURCE"
)

// Controllers

var formController FormController
var inventoryController InventoryController

// Models
var userModel model.User
var formModel model.Form
var inventoryModel model.Inventory

func Routes(base *gin.RouterGroup) {
	// Middleware
	authNormal := middleware.LoadJWTAuth()

	// authenticationController.Routes(base, authNormal)
	formController.Routes(base, authNormal)
	inventoryController.Routes(base, authNormal)

}
