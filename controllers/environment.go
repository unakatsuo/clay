package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qb0C80aE/clay/logics"
	"github.com/qb0C80aE/clay/models"
)

func GetEnvironments(c *gin.Context) {
	processMultiGet(c, models.EnvironmentModel, logics.GetEnvironments, OutputJsonError, OutputMultiJsonResult)
}

func GetEnvironment(c *gin.Context) {
	processSingleGet(c, models.EnvironmentModel, logics.GetEnvironment, OutputJsonError, OutputSingleJsonResult)
}

func CreateEnvironment(c *gin.Context) {
	container := &models.Environment{}
	processCreate(c, container, logics.CreateEnvironment, OutputJsonError, OutputSingleJsonResult)
}

func UpdateEnvironment(c *gin.Context) {
	container := &models.Environment{}
	processUpdate(c, container, logics.UpdateEnvironment, OutputJsonError, OutputSingleJsonResult)
}

func DeleteEnvironment(c *gin.Context) {
	processDelete(c, logics.DeleteEnvironment, OutputJsonError, OutputNothing)
}

func ApplyEnvironment(c *gin.Context) {
	container := &models.Environment{}
	processUpdate(c, container, logics.ApplyEnvironment, OutputJsonError, OutputNothing)
}