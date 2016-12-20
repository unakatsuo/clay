package models

type Design struct {
	Nodes        []*Node        `json:"nodes"`
	NodeTypes    []*NodeType    `json:"node_types"`
	NodePvs      []*NodePv      `json:"node_pvs"`
	Ports        []*Port        `json:"ports"`
	NodeGroups   []*NodeGroup   `json:"node_groups"`
	Protocols    []*Protocol    `json:"protocols"`
	Connections  []*Connection  `json:"connections"`
	Services     []*Service     `json:"services"`
	Requirements []*Requirement `json:"requirements"`
}

var DesignModel = &Design{}
