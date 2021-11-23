package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/citiaps/pnrrd-formulario-jerarquias/util"
	"github.com/gin-gonic/gin"
)

type HierarchyController struct {
}

func (hierarchyController *HierarchyController) Routes(base *gin.RouterGroup, authNormal *jwt.GinJWTMiddleware) *gin.RouterGroup {

	hierarchyRouter := base.Group("/hierarchies", hierarchyController.GetAll())
	{
		hierarchyRouter.GET("")
	}

	return hierarchyRouter

}

// GetAll : Obtener lista de jararquias
// @Summary Obtener la lista de jerarquias, agrupadas en Regiones-Provincias-Comunas
// @ID get-hierarchy
// @Tags hierarchy
// @Produce json
// @Success 200 {object} model.Hierarchy
// @Failure 400 {object} util.Error
// @Router /hierarchies [get]
func (hierarchyController *HierarchyController) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {

		hierarchies, err := hierarchyModel.FindAll()

		if err != nil {
			c.JSON(http.StatusNotFound, util.GetError("No se pudo obtener las jearquías", err))
		}

		c.JSON(http.StatusOK, hierarchies)
	}
}
