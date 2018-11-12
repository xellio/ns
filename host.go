package ns

import (
	"time"

	"github.com/jinzhu/gorm"
)

//
// Host holds all information about detected hosts
//  <host starttime="..." endtime="...">....</host>
//
type Host struct {
	gorm.Model
	Starttime int    `xml:"starttime,attr"`
	Endtime   int    `xml:"endtime,attr"`
	Status    Status `xml:"status"`
	StatusID  int
	Addresses []Address  `xml:"address"`
	Hostnames []Hostname `xml:"hostnames>hostname"`
	Ports     []Port     `xml:"ports>port"`
	OS        OS         `xml:"os"`
	OSID      int
	Times     Times `xml:"times"`
	TimesID   int
	Trace     Trace `xml:"trace"`
	TraceID   int
	LastSeen  time.Time
}

//
// Status holds information about its Hosts status
//  <status state="up" reason="localhost-response" reason_ttl="0"/>
//
type Status struct {
	gorm.Model
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL int    `xml:"reason_ttl,attr"`
}

//
// Times ...
//  <times srtt="27" rttvar="3" to="100000"/>
//
type Times struct {
	gorm.Model
	SRTT   int `xml:"srtt,attr"`
	RTTVar int `xml:"rttvar,attr"`
	TO     int `xml:"to,attr"`
}

//
// Address ...
//  <address addr="127.0.0.1" addrtype="ipv4"/>
//
type Address struct {
	gorm.Model
	HostID   int
	Addr     string `xml:"addr,attr" gorm:"unique_index:idx_addr_type_vendor"`
	AddrType string `xml:"addrtype,attr" gorm:"unique_index:idx_addr_type_vendor"`
	Vendor   string `xml:"vendor,attr" gorm:"unique_index:idx_addr_type_vendor"`
}

//
// Hostname ...
//
//
type Hostname struct {
	gorm.Model
	HostID int
	Name   string `xml:"name,attr"`
	Type   string `xml:"type,attr"`
}

//
// CPE ...
//  <cpe>cpe:/a:igor_sysoev:nginx:1.13.12</cpe>
//
type CPE struct {
	gorm.Model
	Value string `xml:",chardata"`
}

//
// BeforeCreate sets the created_at, updated_at and last_seen date on the Host struct.
// This function is triggered by gorm
//
func (h *Host) BeforeCreate(scope *gorm.Scope) error {
	var err error

	err = scope.SetColumn("created_at", time.Now())
	if err != nil {
		return err
	}
	err = scope.SetColumn("updated_at", time.Now())
	if err != nil {
		return err
	}
	err = scope.SetColumn("last_seen", time.Now())
	if err != nil {
		return err
	}
	return nil
}

//
// BeforeUpdate updates updated_at and last_seen date on the Host struct.
// This function is triggered by gorm
//
func (h *Host) BeforeUpdate(scope *gorm.Scope) error {
	var err error
	err = scope.SetColumn("updated_at", time.Now())
	if err != nil {
		return err
	}
	err = scope.SetColumn("last_seen", time.Now())
	if err != nil {
		return err
	}
	return nil
}
