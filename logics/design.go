package logics

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/mohae/deepcopy"
	"github.com/qb0C80aE/clay/models"
)

func GetDesign(db *gorm.DB, _ string, queryFields string) (interface{}, error) {

	requirements := []*models.Requirement{}
	if err := db.Select(queryFields).Find(&requirements).Error; err != nil {
		return nil, err
	}

	connections := []*models.Connection{}
	if err := db.Select(queryFields).Find(&connections).Error; err != nil {
		return nil, err
	}

	services := []*models.Service{}
	if err := db.Select(queryFields).Find(&services).Error; err != nil {
		return nil, err
	}

	protocols := []*models.Protocol{}
	if err := db.Select(queryFields).Find(&protocols).Error; err != nil {
		return nil, err
	}

	nodePvs := []*models.NodePv{}
	if err := db.Select(queryFields).Find(&nodePvs).Error; err != nil {
		return nil, err
	}

	nodeTypes := []*models.NodeType{}
	if err := db.Select(queryFields).Find(&nodeTypes).Error; err != nil {
		return nil, err
	}

	nodes := []*models.Node{}
	if err := db.Preload("Ports").Preload("NodeGroups").Select(queryFields).Find(&nodes).Error; err != nil {
		return nil, err
	}

	nodeGroups := []*models.NodeGroup{}
	if err := db.Preload("Nodes").Select(queryFields).Find(&nodeGroups).Error; err != nil {
		return nil, err
	}

	ports := []*models.Port{}
	if err := db.Select(queryFields).Find(&ports).Error; err != nil {
		return nil, err
	}

	design := &models.Design{}
	design.NodePvs = nodePvs
	design.NodeTypes = nodeTypes
	design.Nodes = nodes
	design.NodeGroups = nodeGroups
	design.Ports = ports
	design.Protocols = protocols
	design.Services = services
	design.Connections = connections
	design.Requirements = requirements

	return design, nil

}

func UpdateDesign(db *gorm.DB, _ string, data interface{}) (interface{}, error) {

	design := data.(*models.Design)
	originalDesign := deepcopy.Copy(design).(*models.Design)

	if err := db.Exec("delete from requirements;").Error; err != nil {
		return nil, err
	}

	if err := db.Exec("delete from connections;").Error; err != nil {
		return nil, err
	}

	if err := db.Exec("delete from services;").Error; err != nil {
		return nil, err
	}

	if err := db.Exec("delete from protocols;").Error; err != nil {
		return nil, err
	}

	if err := db.Exec("delete from nodes;").Error; err != nil {
		return nil, err
	}

	if err := db.Exec("delete from node_types;").Error; err != nil {
		return nil, err
	}

	if err := db.Exec("delete from node_pvs;").Error; err != nil {
		return nil, err
	}

	if err := db.Exec("delete from node_groups;").Error; err != nil {
		return nil, err
	}

	if err := db.Exec("delete from node_group_association;").Error; err != nil {
		return nil, err
	}

	if err := db.Exec("delete from ports;").Error; err != nil {
		return nil, err
	}

	nodeTypes := design.NodeTypes
	for _, nodeType := range nodeTypes {
		if err := db.Create(nodeType).Error; err != nil {
			return nil, err
		}
	}

	nodePvs := design.NodePvs
	for _, nodePv := range nodePvs {
		if err := db.Create(nodePv).Error; err != nil {
			return nil, err
		}
	}

	nodes := design.Nodes
	for _, node := range nodes {
		node.Ports = nil
		node.NodeGroups = nil
		if err := db.Create(node).Error; err != nil {
			return nil, err
		}
	}

	nodeGroups := design.NodeGroups
	for _, nodeGroup := range nodeGroups {
		nodeGroup.Nodes = nil
		if err := db.Create(nodeGroup).Error; err != nil {
			return nil, err
		}
	}

	ports := design.Ports
	for _, port := range ports {
		port.DestinationPortID = sql.NullInt64{Int64: 0, Valid: false}
		if err := db.Create(port).Error; err != nil {
			return nil, err
		}
	}

	nodeGroups = originalDesign.NodeGroups
	for _, nodeGroup := range nodeGroups {
		for _, node := range nodeGroup.Nodes {
			relatedNode := &models.Node{}
			if err := db.Preload("NodeGroups").First(&relatedNode, node.ID).Association("NodeGroups").Append(nodeGroup).Error; err != nil {
				return nil, err
			}
		}
	}

	ports = originalDesign.Ports
	for _, port := range ports {
		if err := db.Save(port).Error; err != nil {
			return nil, err
		}
	}

	protocols := design.Protocols
	for _, protocol := range protocols {
		if err := db.Create(protocol).Error; err != nil {
			return nil, err
		}
	}

	services := design.Services
	for _, service := range services {
		if err := db.Create(service).Error; err != nil {
			return nil, err
		}
	}

	connections := design.Connections
	for _, connection := range connections {
		if err := db.Create(connection).Error; err != nil {
			return nil, err
		}
	}

	requirements := design.Requirements
	for _, requirement := range requirements {
		if err := db.Create(requirement).Error; err != nil {
			return nil, err
		}
	}

	return design, nil

}

func DeleteDesign(db *gorm.DB, _ string) error {

	if err := db.Exec("delete from requirements;").Error; err != nil {
		return err
	}

	if err := db.Exec("delete from connections;").Error; err != nil {
		return err
	}

	if err := db.Exec("delete from services;").Error; err != nil {
		return err
	}

	if err := db.Exec("delete from protocols;").Error; err != nil {
		return err
	}

	if err := db.Exec("delete from nodes;").Error; err != nil {
		return err
	}

	if err := db.Exec("delete from node_types;").Error; err != nil {
		return err
	}

	if err := db.Exec("delete from node_pvs;").Error; err != nil {
		return err
	}

	if err := db.Exec("delete from node_groups;").Error; err != nil {
		return err
	}

	if err := db.Exec("delete from node_group_association;").Error; err != nil {
		return err
	}

	if err := db.Exec("delete from ports;").Error; err != nil {
		return err
	}

	return nil

}
