package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qb0C80aE/clay/logics"
	"github.com/qb0C80aE/clay/models"
)

func GetProtocols(c *gin.Context) {
	processMultiGet(c, models.ProtocolModel, logics.GetProtocols, OutputJsonError, OutputMultiJsonResult)
}

func GetProtocol(c *gin.Context) {
	processSingleGet(c, models.ProtocolModel, logics.GetProtocol, OutputJsonError, OutputSingleJsonResult)
}

func CreateProtocol(c *gin.Context) {
	container := &models.Protocol{}
	processCreate(c, container, logics.CreateProtocol, OutputJsonError, OutputSingleJsonResult)
}

func UpdateProtocol(c *gin.Context) {
	container := &models.Protocol{}
	processUpdate(c, container, logics.UpdateProtocol, OutputJsonError, OutputSingleJsonResult)
}

func DeleteProtocol(c *gin.Context) {
	processDelete(c, logics.DeleteProtocol, OutputJsonError, OutputNothing)
}

func GetServices(c *gin.Context) {
	processMultiGet(c, models.ServiceModel, logics.GetServices, OutputJsonError, OutputMultiJsonResult)
}

func GetService(c *gin.Context) {
	processSingleGet(c, models.ServiceModel, logics.GetService, OutputJsonError, OutputSingleJsonResult)
}

func CreateService(c *gin.Context) {
	container := &models.Service{}
	processCreate(c, container, logics.CreateService, OutputJsonError, OutputSingleJsonResult)
}

func UpdateService(c *gin.Context) {
	container := &models.Service{}
	processUpdate(c, container, logics.UpdateService, OutputJsonError, OutputSingleJsonResult)
}

func DeleteService(c *gin.Context) {
	processDelete(c, logics.DeleteService, OutputJsonError, OutputNothing)
}

func GetConnections(c *gin.Context) {
	processMultiGet(c, models.ConnectionModel, logics.GetConnections, OutputJsonError, OutputMultiJsonResult)
}

func GetConnection(c *gin.Context) {
	processSingleGet(c, models.ConnectionModel, logics.GetConnection, OutputJsonError, OutputSingleJsonResult)
}

func CreateConnection(c *gin.Context) {
	container := &models.Connection{}
	processCreate(c, container, logics.CreateConnection, OutputJsonError, OutputSingleJsonResult)
}

func UpdateConnection(c *gin.Context) {
	container := &models.Connection{}
	processUpdate(c, container, logics.UpdateConnection, OutputJsonError, OutputSingleJsonResult)
}

func DeleteConnection(c *gin.Context) {
	processDelete(c, logics.DeleteConnection, OutputJsonError, OutputNothing)
}

func GetRequirements(c *gin.Context) {
	processMultiGet(c, models.RequirementModel, logics.GetRequirements, OutputJsonError, OutputMultiJsonResult)
}

func GetRequirement(c *gin.Context) {
	processSingleGet(c, models.RequirementModel, logics.GetRequirement, OutputJsonError, OutputSingleJsonResult)
}

func CreateRequirement(c *gin.Context) {
	container := &models.Requirement{}
	processCreate(c, container, logics.CreateRequirement, OutputJsonError, OutputSingleJsonResult)
}

func UpdateRequirement(c *gin.Context) {
	container := &models.Requirement{}
	processUpdate(c, container, logics.UpdateRequirement, OutputJsonError, OutputSingleJsonResult)
}

func DeleteRequirement(c *gin.Context) {
	processDelete(c, logics.DeleteRequirement, OutputJsonError, OutputNothing)
}
