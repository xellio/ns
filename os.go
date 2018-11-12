package ns

import "github.com/jinzhu/gorm"

//
// OS ...
//  <os>...</os>
//
type OS struct {
	gorm.Model
	Portused        []Portused    `xml:"portused"`
	OSMatch         []OSMatch     `xml:"osmatch"`
	OSFingerprint   OSFingerprint `xml:"osfingerprint"`
	OSFingerprintID int
}

//
// Portused ...
//  <portused state="open" proto="tcp" portid="22"/>
//
type Portused struct {
	gorm.Model
	OSID     int
	State    string `xml:"state,attr"`
	Protocol string `xml:"proto,attr"`
	PortID   int    `xml:"portid,attr"`
}

//
// OSMatch ...
//  <osmatch name="Linux 3.12 - 4.10" accuracy="100" line="63515">...</osmatch>
//
type OSMatch struct {
	gorm.Model
	OSID     int
	Name     string    `xml:"name,attr"`
	Accuracy int       `xml:"accuracy,attr"`
	Line     int       `xml:"line,attr"`
	OSClass  []OSClass `xml:"osclass"`
}

//
// OSClass ...
//  <osclass type="general purpose" vendor="Linux" osfamily="Linux" osgen="4.X" accuracy="100">...</osclass>
//
type OSClass struct {
	gorm.Model
	OSMatchID int
	Type      string `xml:"type,attr"`
	Vendor    string `xml:"vendor,attr"`
	OSFamily  string `xml:"osfamily,attr"`
	Accuracy  int    `xml:"accuracy,attr"`
	CPE       []CPE  `xml:"cpe"`
}

//
// OSFingerprint ...
//
type OSFingerprint struct {
	gorm.Model
	Fingerprint string `xml:"fingerprint,attr" sql:"type:text"`
}
