package logics

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/models"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func updateDesignFile(db *gorm.DB, id string, environment *models.Environment) error {
	design, err := GetDesign(db, id, "*")

	if err != nil {
		return err
	}

	jsonString, err := json.MarshalIndent(design, "", "    ")
	if ioutil.WriteFile(fmt.Sprintf("%s/%s", environment.GitRepositoryURI, environment.DesignFileName), jsonString, os.ModePerm); err != nil {
		return err
	}

	var out bytes.Buffer
	cmd := exec.Command("git", "add", environment.DesignFileName)
	cmd.Dir = environment.GitRepositoryURI
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return errors.New(out.String())
	}

	return nil
}

func updateTemplateFile(db *gorm.DB, id string, environment *models.Environment) error {
	template, err := ApplyTemplate(db, id, "*")

	if err != nil {
		return err
	}

	if ioutil.WriteFile(fmt.Sprintf("%s/%s", environment.GitRepositoryURI, environment.TemplateFileName), ([]byte)(template.(string)), os.ModePerm); err != nil {
		return err
	}

	var out bytes.Buffer
	cmd := exec.Command("git", "add", environment.TemplateFileName)
	cmd.Dir = environment.GitRepositoryURI
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return errors.New(out.String())
	}

	return nil
}

func updateTestCaseFile(db *gorm.DB, id string, environment *models.Environment) error {
	testCommands, err := generateTestScripts(db, id)

	if err != nil {
		return err
	}

	var out bytes.Buffer
	cmd := exec.Command("rm", "-rf", environment.TestCaseDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return errors.New(out.String())
	}

	cmd = exec.Command("mkdir", environment.TestCaseDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return errors.New(out.String())
	}

	for _, testCommand := range testCommands {
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_server.sh",
			environment.GitRepositoryURI,
			environment.TestCaseDirectoryName,
			testCommand.ServiceName),
			([]byte)(testCommand.SerevrScriptTemplate),
			os.ModePerm); err != nil {
			return err
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_client.sh",
			environment.GitRepositoryURI,
			environment.TestCaseDirectoryName,
			testCommand.ServiceName),
			([]byte)(testCommand.ClientScriptTemplate),
			os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func commit(environment *models.Environment, message string) error {
	var out bytes.Buffer
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = environment.GitRepositoryURI
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return errors.New(out.String())
	}
	cmd = exec.Command("git", "commit", "-m", message)
	cmd.Dir = environment.GitRepositoryURI
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return errors.New(out.String())
	}
	return nil
}

func GetEnvironments(db *gorm.DB, queryFields string) ([]interface{}, error) {

	environments := []*models.Environment{}

	if err := db.Select(queryFields).Find(&environments).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(environments))
	for i, data := range environments {
		result[i] = data
	}

	return result, nil

}

func GetEnvironment(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	environment := &models.Environment{}

	if err := db.Select(queryFields).First(environment, id).Error; err != nil {
		return nil, err
	}

	return environment, nil

}

func CreateEnvironment(db *gorm.DB, data interface{}) (interface{}, error) {

	environment := data.(*models.Environment)

	if err := db.Create(environment).Error; err != nil {
		return nil, err
	}

	return environment, nil

}

func UpdateEnvironment(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	environment := data.(*models.Environment)
	environment.ID, _ = strconv.Atoi(id)

	if err := db.Save(&environment).Error; err != nil {
		return nil, err
	}

	return environment, nil

}

func DeleteEnvironment(db *gorm.DB, id string) error {

	environment := &models.Environment{}

	if err := db.First(&environment, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&environment).Error; err != nil {
		return err
	}

	return nil

}

func ApplyEnvironment(db *gorm.DB, id string, _ interface{}) (interface{}, error) {
	data, err := GetEnvironment(db, id, "*")
	if err != nil {
		return "", err
	}

	environment := data.(*models.Environment)

	message := fmt.Sprintf("Automatic commit at %s", time.Now().String())
	if err := updateDesignFile(db, id, environment); err != nil {
		return "", err
	}
	if err := updateTemplateFile(db, id, environment); err != nil {
		return "", err
	}
	if err := updateTestCaseFile(db, id, environment); err != nil {
		return "", err
	}
	if err := commit(environment, message); err != nil {
		return "", err
	}
	return "", nil
}
