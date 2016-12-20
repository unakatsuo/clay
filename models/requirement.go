package models

type Protocol struct {
	ID   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"not null;unique"`
}

type Service struct {
	ID          int           `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name        string        `json:"name" gorm:"not null;unique"`
	Connections []*Connection `json:"connections"`
}

type Connection struct {
	ID         int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ServiceID  int       `json:"service_id" gorm:"not null" sql:"type:integer references services(id)"`
	ProtocolID int       `json:"protocol_id" gorm:"not null" sql:"type:integer references protocols(id)"`
	Protocol   *Protocol `json:"protocol"`
	PortNumber int       `json:"port_number" gorm:"not null"`
}

type Requirement struct {
	ID                int      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	SourcePortID      int      `json:"source_port_id" gorm:"not null" sql:"type:integer references ports(id)"`
	SourcePort        *Port    `json:"source_port"`
	DestinationPortID int      `json:"destination_port_id" gorm:"not null" sql:"type:integer references ports(id)"`
	DestinationPort   *Port    `json:"destination_port"`
	ServiceID         int      `json:"service_id" gorm:"not null" sql:"type:integer references services(id)"`
	Service           *Service `json:"service"`
	Accessibility     bool     `json:"accessibility"`
}

var ProtocolModel = &Protocol{}
var ServiceModel = &Service{}
var ConnectionModel = &Connection{}
var RequirementModel = &Requirement{}
