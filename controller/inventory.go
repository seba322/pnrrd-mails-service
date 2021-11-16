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
// @ID get-inventory
// @Tags inventory
// @Produce json
// @Param institution query string true "Id de institución"
// @Param hierarchy query string true "Tipo de jararquía de inventario solicitado"
// @Param hierarchy_id query string false "Id de la jerarquía solicitada"
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

		inventory, err := inventoryModel.GetInventory(params.Institution, params.Hierarchy, params.Hierarchy_id)

		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo obtener el inventario para esta consulta", err))
		}

		c.JSON(http.StatusOK, inventory)

	}

}

// UpsertInventory : Ingresar actualizaciones de inventario
// @Summary Ingresar actualizaciones de inventario, para una institución
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
