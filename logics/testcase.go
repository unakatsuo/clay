package logics

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/models"
	"strconv"
	tplpkg "text/template"
)

func GetTestCommands(db *gorm.DB, queryFields string) ([]interface{}, error) {

	testCommands := []*models.TestCommand{}

	if err := db.Select(queryFields).Find(&testCommands).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(testCommands))
	for i, data := range testCommands {
		result[i] = data
	}

	return result, nil

}

func GetTestCommand(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testCommand := &models.TestCommand{}

	if err := db.Select(queryFields).First(testCommand, id).Error; err != nil {
		return nil, err
	}

	return testCommand, nil

}

func CreateTestCommand(db *gorm.DB, data interface{}) (interface{}, error) {

	testCommand := data.(*models.TestCommand)

	if err := db.Create(&testCommand).Error; err != nil {
		return nil, err
	}

	return testCommand, nil
}

func UpdateTestCommand(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testCommand := data.(*models.TestCommand)
	testCommand.ID, _ = strconv.Atoi(id)

	if err := db.Save(testCommand).Error; err != nil {
		return nil, err
	}

	return testCommand, nil
}

func DeleteTestCommand(db *gorm.DB, id string) error {

	testCommand := &models.TestCommand{}

	if err := db.First(&testCommand, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testCommand).Error; err != nil {
		return err
	}

	return nil

}

func GetTestPatterns(db *gorm.DB, queryFields string) ([]interface{}, error) {

	testPatterns := []*models.TestPattern{}

	if err := db.Select(queryFields).Find(&testPatterns).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(testPatterns))
	for i, data := range testPatterns {
		result[i] = data
	}

	return result, nil

}

func GetTestPattern(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testPattern := &models.TestPattern{}

	if err := db.Select(queryFields).First(testPattern, id).Error; err != nil {
		return nil, err
	}

	return testPattern, nil

}

func CreateTestPattern(db *gorm.DB, data interface{}) (interface{}, error) {

	testPattern := data.(*models.TestPattern)

	if err := db.Create(&testPattern).Error; err != nil {
		return nil, err
	}

	return testPattern, nil
}

func UpdateTestPattern(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testPattern := data.(*models.TestPattern)
	testPattern.ID, _ = strconv.Atoi(id)

	if err := db.Save(testPattern).Error; err != nil {
		return nil, err
	}

	return testPattern, nil
}

func DeleteTestPattern(db *gorm.DB, id string) error {

	testPattern := &models.TestPattern{}

	if err := db.First(&testPattern, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testPattern).Error; err != nil {
		return err
	}

	return nil

}

func GetTestCases(db *gorm.DB, queryFields string) ([]interface{}, error) {

	testCases := []*models.TestCase{}

	if err := db.Select(queryFields).Find(&testCases).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(testCases))
	for i, data := range testCases {
		result[i] = data
	}

	return result, nil

}

func GetTestCase(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testCase := &models.TestCase{}

	if err := db.Select(queryFields).First(testCase, id).Error; err != nil {
		return nil, err
	}

	return testCase, nil

}

func CreateTestCase(db *gorm.DB, data interface{}) (interface{}, error) {

	testCase := data.(*models.TestCase)

	if err := db.Create(&testCase).Error; err != nil {
		return nil, err
	}

	return testCase, nil
}

func UpdateTestCase(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testCase := data.(*models.TestCase)
	testCase.ID, _ = strconv.Atoi(id)

	if err := db.Save(testCase).Error; err != nil {
		return nil, err
	}

	return testCase, nil
}

func DeleteTestCase(db *gorm.DB, id string) error {

	testCase := &models.TestCase{}

	if err := db.First(&testCase, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testCase).Error; err != nil {
		return err
	}

	return nil

}

func ApplyTestCase(db *gorm.DB, id string, _ string) (interface{}, error) {
	result, err := generateTestScripts(db, id)
	if err != nil {
		return "", err
	}
	a := []*models.TestCommand(result)
	c := bytes.Buffer{}
	for _, b := range a {
		c.WriteString("#--------------------------------------\n")
		c.WriteString(fmt.Sprintf("# %s\n", b.ServiceName))
		c.WriteString("# --- server script ---\n")
		c.WriteString(fmt.Sprintf("%s\n", b.SerevrScriptTemplate))
		c.WriteString("# --- client script ---\n")
		c.WriteString(fmt.Sprintf("%s\n", b.ClientScriptTemplate))
		c.WriteString("#--------------------------------------\n")
	}
	return c.String(), nil
}

func generateTestScripts(db *gorm.DB, id string) ([]*models.TestCommand, error) {

	result := []*models.TestCommand{}

	testCase := &models.TestCase{}
	if err := db.Preload("TestPatterns").
		Preload("TestPatterns.TestCommand").First(&testCase, id).Error; err != nil {
		return result, err
	}

	testCommandMap := make(map[string]*models.TestCommand)
	for _, testPattern := range testCase.TestPatterns {
		testCommandMap[testPattern.TestCommand.ServiceName] = testPattern.TestCommand
	}

	requirements := []*models.Requirement{}
	if err := db.Preload("Service").
		Preload("Service.Connections").
		Preload("SourcePort").
		Preload("SourcePort.Node").
		Preload("DestinationPort").
		Preload("DestinationPort.Node").Select("*").Find(&requirements).Error; err != nil {
		return result, err
	}

	for _, requirement := range requirements {
		testCommand := testCommandMap[requirement.Service.Name]
		serverScript := testCommand.SerevrScriptTemplate
		clientScript := testCommand.ClientScriptTemplate

		var docServerScript bytes.Buffer
		tplServerScript, _ := tplpkg.New("template_server_script").Parse(serverScript)
		tplServerScript.Execute(&docServerScript, requirement)
		serverScript = docServerScript.String()

		var docClientScript bytes.Buffer
		tplClientScript, _ := tplpkg.New("template_server_script").Parse(clientScript)
		tplClientScript.Execute(&docClientScript, requirement)
		clientScript = docClientScript.String()

		newTestCommand := &models.TestCommand{
			ServiceName:          fmt.Sprintf("%s_to_%s_%s", requirement.SourcePort.Ipv4Address.String, requirement.DestinationPort.Ipv4Address.String, requirement.Service.Name),
			SerevrScriptTemplate: serverScript,
			ClientScriptTemplate: clientScript,
		}
		result = append(result, newTestCommand)
	}

	return result, nil

}
