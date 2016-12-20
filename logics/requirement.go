package logics

import (
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/models"
	"strconv"
)

func GetProtocols(db *gorm.DB, queryFields string) ([]interface{}, error) {

	protocols := []*models.Protocol{}

	if err := db.Select(queryFields).Find(&protocols).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(protocols))
	for i, data := range protocols {
		result[i] = data
	}

	return result, nil

}

func GetProtocol(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	protocol := &models.Protocol{}

	if err := db.Select(queryFields).First(protocol, id).Error; err != nil {
		return nil, err
	}

	return protocol, nil

}

func CreateProtocol(db *gorm.DB, data interface{}) (interface{}, error) {

	protocol := data.(*models.Protocol)

	if err := db.Create(&protocol).Error; err != nil {
		return nil, err
	}

	return protocol, nil
}

func UpdateProtocol(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	protocol := data.(*models.Protocol)
	protocol.ID, _ = strconv.Atoi(id)

	if err := db.Save(protocol).Error; err != nil {
		return nil, err
	}

	return protocol, nil
}

func DeleteProtocol(db *gorm.DB, id string) error {

	protocol := &models.Protocol{}

	if err := db.First(&protocol, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&protocol).Error; err != nil {
		return err
	}

	return nil

}

func GetServices(db *gorm.DB, queryFields string) ([]interface{}, error) {

	services := []*models.Service{}

	if err := db.Select(queryFields).Find(&services).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(services))
	for i, data := range services {
		result[i] = data
	}

	return result, nil

}

func GetService(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	service := &models.Service{}

	if err := db.Select(queryFields).First(service, id).Error; err != nil {
		return nil, err
	}

	return service, nil

}

func CreateService(db *gorm.DB, data interface{}) (interface{}, error) {

	service := data.(*models.Service)

	if err := db.Create(&service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func UpdateService(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	service := data.(*models.Service)
	service.ID, _ = strconv.Atoi(id)

	if err := db.Save(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func DeleteService(db *gorm.DB, id string) error {

	service := &models.Service{}

	if err := db.First(&service, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&service).Error; err != nil {
		return err
	}

	return nil

}

func GetConnections(db *gorm.DB, queryFields string) ([]interface{}, error) {

	connections := []*models.Connection{}

	if err := db.Select(queryFields).Find(&connections).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(connections))
	for i, data := range connections {
		result[i] = data
	}

	return result, nil

}

func GetConnection(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	connection := &models.Connection{}

	if err := db.Select(queryFields).First(connection, id).Error; err != nil {
		return nil, err
	}

	return connection, nil

}

func CreateConnection(db *gorm.DB, data interface{}) (interface{}, error) {

	connection := data.(*models.Connection)

	if err := db.Create(&connection).Error; err != nil {
		return nil, err
	}

	return connection, nil
}

func UpdateConnection(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	connection := data.(*models.Connection)
	connection.ID, _ = strconv.Atoi(id)

	if err := db.Save(connection).Error; err != nil {
		return nil, err
	}

	return connection, nil
}

func DeleteConnection(db *gorm.DB, id string) error {

	connection := &models.Connection{}

	if err := db.First(&connection, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&connection).Error; err != nil {
		return err
	}

	return nil

}

func GetRequirements(db *gorm.DB, queryFields string) ([]interface{}, error) {

	requirements := []*models.Requirement{}

	if err := db.Select(queryFields).Find(&requirements).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(requirements))
	for i, data := range requirements {
		result[i] = data
	}

	return result, nil

}

func GetRequirement(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	requirement := &models.Requirement{}

	if err := db.Select(queryFields).First(requirement, id).Error; err != nil {
		return nil, err
	}

	return requirement, nil

}

func CreateRequirement(db *gorm.DB, data interface{}) (interface{}, error) {

	requirement := data.(*models.Requirement)

	if err := db.Create(&requirement).Error; err != nil {
		return nil, err
	}

	return requirement, nil
}

func UpdateRequirement(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	requirement := data.(*models.Requirement)
	requirement.ID, _ = strconv.Atoi(id)

	if err := db.Save(requirement).Error; err != nil {
		return nil, err
	}

	return requirement, nil
}

func DeleteRequirement(db *gorm.DB, id string) error {

	requirement := &models.Requirement{}

	if err := db.First(&requirement, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&requirement).Error; err != nil {
		return err
	}

	return nil

}
