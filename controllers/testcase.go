package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qb0C80aE/clay/logics"
	"github.com/qb0C80aE/clay/models"
)

func GetTestCommands(c *gin.Context) {
	processMultiGet(c, models.TestCommandModel, logics.GetTestCommands, OutputJsonError, OutputMultiJsonResult)
}

func GetTestCommand(c *gin.Context) {
	processSingleGet(c, models.TestCommandModel, logics.GetTestCommand, OutputJsonError, OutputSingleJsonResult)
}

func CreateTestCommand(c *gin.Context) {
	container := &models.TestCommand{}
	processCreate(c, container, logics.CreateTestCommand, OutputJsonError, OutputSingleJsonResult)
}

func UpdateTestCommand(c *gin.Context) {
	container := &models.TestCommand{}
	processUpdate(c, container, logics.UpdateTestCommand, OutputJsonError, OutputSingleJsonResult)
}

func DeleteTestCommand(c *gin.Context) {
	processDelete(c, logics.DeleteTestCommand, OutputJsonError, OutputNothing)
}

func GetTestPatterns(c *gin.Context) {
	processMultiGet(c, models.TestPatternModel, logics.GetTestPatterns, OutputJsonError, OutputMultiJsonResult)
}

func GetTestPattern(c *gin.Context) {
	processSingleGet(c, models.TestPatternModel, logics.GetTestPattern, OutputJsonError, OutputSingleJsonResult)
}

func CreateTestPattern(c *gin.Context) {
	container := &models.TestPattern{}
	processCreate(c, container, logics.CreateTestPattern, OutputJsonError, OutputSingleJsonResult)
}

func UpdateTestPattern(c *gin.Context) {
	container := &models.TestPattern{}
	processUpdate(c, container, logics.UpdateTestPattern, OutputJsonError, OutputSingleJsonResult)
}

func DeleteTestPattern(c *gin.Context) {
	processDelete(c, logics.DeleteTestPattern, OutputJsonError, OutputNothing)
}

func GetTestCases(c *gin.Context) {
	processMultiGet(c, models.TestCaseModel, logics.GetTestCases, OutputJsonError, OutputMultiJsonResult)
}

func GetTestCase(c *gin.Context) {
	processSingleGet(c, models.TestCaseModel, logics.GetTestCase, OutputJsonError, OutputSingleJsonResult)
}

func CreateTestCase(c *gin.Context) {
	container := &models.TestCase{}
	processCreate(c, container, logics.CreateTestCase, OutputJsonError, OutputSingleJsonResult)
}

func UpdateTestCase(c *gin.Context) {
	container := &models.TestCase{}
	processUpdate(c, container, logics.UpdateTestCase, OutputJsonError, OutputSingleJsonResult)
}

func DeleteTestCase(c *gin.Context) {
	processDelete(c, logics.DeleteTestCase, OutputJsonError, OutputNothing)
}

func ApplyTestCase(c *gin.Context) {
	container := &models.TestCase{}
	processSingleGet(c, container, logics.ApplyTestCase, OutputJsonError, OutputTextResult)
}