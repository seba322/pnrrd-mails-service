package controller

import (
	"github.com/citiaps/template-go-rest/middleware"
	"github.com/citiaps/template-go-rest/model"
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

var dogController DogController
var formController FormController

// Models
var dogModel model.Dog
var userModel model.User
var formModel model.Form

func Routes(base *gin.RouterGroup) {
	// Middleware
	authNormal := middleware.LoadJWTAuth()

	// authenticationController.Routes(base, authNormal)
	dogController.Routes(base, authNormal)

	formController.Routes(base, authNormal)

}
