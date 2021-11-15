package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/citiaps/template-go-rest/model"
	"github.com/citiaps/template-go-rest/util"
	"github.com/gin-gonic/gin"
)

type FormController struct {
}

func (formController *FormController) Routes(base *gin.RouterGroup, authNormal *jwt.GinJWTMiddleware) *gin.RouterGroup {

	formRouter := base.Group("/forms")
	{
		formRouter.GET("", formController.GetForm())
		formRouter.POST("", formController.CreateForm())
	}

	return formRouter
}

// GetForm : Obtener un formulario para jerarquía y tipo especifico
// @Summary Obtener formulario general o de recursos para una jerarquía especifica
// @ID get-form
// @Tags forms
// @Produce json
// @Param type query string true "Tipo de formulario, que puede ser INFORMATION o RESOURCE"
// @Param hierarchy query string false "Jerarquía del  formulario, por defecto solo se maneja general"
// @Success 200 {object} model.Form
// @Failure 400 {object} util.Error
// @Router /forms [get]

func (formController *FormController) GetForm() func(c *gin.Context) {
	return func(c *gin.Context) {

		params := FormParams{}
		// params.FormType="INFORMATION"
		params.FormType = GeneralHierarchyForm
		err := c.Bind(&params)

		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo encontrar el parametro type", err))
			return
		}

		form, err := formModel.GetForm(params.FormType, params.Hierarchy)

		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo obtener el fomulario para esta consulta", err))
		}

		c.JSON(http.StatusOK, form)

	}
}

// CreateForm: Crear un nuevo template de formulario
// @Summary crear un nuevo formulario
// @ID create-form
// @Tags forms
// @Accept  json
// @Produce json
// @Param form body model.Form true "Crear Formulario"
// @Success 200 {object} model.Form
// @Failure 400 {object} util.Error
// @Router /forms [post]
func (formController *FormController) CreateForm() func(c *gin.Context) {
	return func(c *gin.Context) {
		var form model.Form
		err := c.Bind(&form)

		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo decodificar json", err))
			return
		}
		err = formModel.Create(&form)

		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo insertar el formulario", err))
			return
		}

		c.JSON(http.StatusOK, form)

	}
}
