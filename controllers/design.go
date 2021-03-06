package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qb0C80aE/clay/db"
	"github.com/qb0C80aE/clay/extension"
	"github.com/qb0C80aE/clay/logics"
	"github.com/qb0C80aE/clay/models"
)

func init() {
	url := "designs/present"
	extension.RegisterUniqueEndpoint(url, url)

	extension.RegisterRoute(extension.MethodGet, url, GetDesign)
	extension.RegisterRoute(extension.MethodPut, url, UpdateDesign)
	extension.RegisterRoute(extension.MethodDelete, url, DeleteDesign)
}

func GetDesign(c *gin.Context) {
	ProcessSingleGet(c, models.DesignModel, logics.GetDesign, OutputJsonError, OutputSingleJsonResult)
}

func UpdateDesign(c *gin.Context) {
	db.DBInstance(c).Exec("pragma foreign_keys = off;")
	ProcessUpdate(c, &models.Design{}, logics.UpdateDesign, OutputJsonError, OutputSingleJsonResult)
	db.DBInstance(c).Exec("pragma foreign_keys = on;")
}

func DeleteDesign(c *gin.Context) {
	db.DBInstance(c).Exec("pragma foreign_keys = off;")
	ProcessDelete(c, logics.DeleteDesign, OutputJsonError, OutputNothing)
	db.DBInstance(c).Exec("pragma foreign_keys = on;")
}
