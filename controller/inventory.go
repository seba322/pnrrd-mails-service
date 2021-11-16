package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/citiaps/pnrrd-formulario-jerarquias/model"
	"github.com/citiaps/pnrrd-formulario-jerarquias/util"
	"github.com/gin-gonic/gin"
)

type InventoryController struct {
}

func (inventoryController *InventoryController) Routes(base *gin.RouterGroup, authNormal *jwt.GinJWTMiddleware) *gin.RouterGroup {

	inventoryRouter := base.Group("/inventories")
	{
		inventoryRouter.GET("", inventoryController.GetInventory())
		inventoryRouter.PUT("", inventoryController.UpsertInventory())
	}

	return inventoryRouter
}

// GetInventory : Obtener inventario de una institucion, para una jerarquia especifica
// @Summary Obtener inventario de una institución, para una jerarquía especifica
// @Description Con este servicio se puede obtener el inventario.
// @Description Existen 2 tipos de inventario, de información ( tag INFORMATION ) y recursos ( tag RESOURCE). El primero enfocado en las capacidades y el segundo en  información institucional.
// @Description Siempre que se quiera obtener la información institucional en parámetro hierarchy debe ir “NACIONAL”.
// @ID get-inventory
// @Tags inventory
// @Produce json
// @Param institution query string true "Id de institución"
// @Param hierarchy query string true "Tipo de jararquía de inventario solicitado"
// @Param hierarchy_id query string false "Id de la jerarquía solicitada, solo requerido si no es jerarquia nacional"
// @Param type query string true "tipo de inventario solicitado (puede ser INFORMATION o RESOURCE)"
// @Success 200 {object} model.Inventory
// @Failure 400 {object} util.Error
// @Router /inventories [get]
func (inventoryController *InventoryController) GetInventory() func(c *gin.Context) {
	return func(c *gin.Context) {
		params := InventoryParams{}

		err := c.Bind(&params)

		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("Parametors requeridos faltantes", err))
			return
		}

		if params.TypeInv == InformationTypeForm {
			params.Hierarchy = model.NACIONAL_HIERARCHY
		}

		inventory, err := inventoryModel.GetInventory(params.Institution, params.Hierarchy, params.Hierarchy_id, params.TypeInv)

		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo obtener el inventario para esta consulta", err))
		}

		c.JSON(http.StatusOK, inventory)

	}

}

// UpsertInventory : Ingresar actualizaciones de inventario
// @Summary Ingresar actualizaciones de inventario, para una institución
// @Description El body es un arreglo del objeto de respuesta que se muestra mas abajo
// @Description Cada respuesta debe ir en el arreglo del body (da lo mismo el orden)
// @Description Es importantel agregar el index de la capacidad que se esta declarando
// @ID update-inventory
// @Tags inventory
// @Produce json
// @Accept  json
// @Param inventory body []model.Inventory true "Actualizar inventario"
// @Success 200 {object} model.Inventory
// @Failure 400 {object} util.Error
// @Router /inventories [put]
func (inventoryController *InventoryController) UpsertInventory() func(c *gin.Context) {
	return func(c *gin.Context) {
		// type InventoryBody struct{
		// 	Data []model.Inventory `form:"data" json:"data"`
		// }
		var inventories []model.Inventory

		err := c.Bind(&inventories)

		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo decodificar json", err))
			return
		}

		err = inventoryModel.UpsertManyInventory(inventories)

		if err != nil {
			c.JSON(http.StatusBadRequest, util.GetError("No se pudo insertar el inventario", err))
			return
		}

		c.JSON(http.StatusOK, inventories)

	}

}
