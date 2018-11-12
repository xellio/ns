package ns

import "github.com/jinzhu/gorm"

//
// Port ...
//  <port protocol="tcp" portid="22">...</port>
//
type Port struct {
	gorm.Model
	HostID    int
	Protocol  string `xml:"protocol,attr"`
	PortID    int    `xml:"portid,attr"`
	State     State  `xml:"state"`
	StateID   int
	Service   Service `xml:"service"`
	ServiceID int
	Scripts   []Script `xml:"script"`
}

//
// State ...
//  <state state="open" reason="syn-ack" reason_ttl="64"/>
//
type State struct {
	gorm.Model
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL int    `xml:"reason_ttl,attr"`
}

//
// Service ...
//  <service name="http" product="nginx" version="1.13.12" method="probed" conf="10">...</service>
//
type Service struct {
	gorm.Model
	Name    string `xml:"name,attr"`
	Product string `xml:"product,attr"`
	Version string `xml:"version,attr"`
	Method  string `xml:"method,attr"`
	Conf    int    `xml:"conf,attr"`
	CPE     []*CPE `xml:"cpe"`
}

//
// Script ...
//  <script id="ssh-hostkey" output="&#xa;  .....">...</script>
//
type Script struct {
	gorm.Model
	PortID   int
	ScriptID string    `xml:"id,attr"`
	Output   string    `xml:"output,attr" sql:"type:text"`
	Elements []Element `xml:"elem"`
}

//
// Element ...
//  <elem key="type">ssh-rsa</elem>
//
type Element struct {
	gorm.Model
	ScriptID int
	Key      string `xml:"key,attr"`
	Elem     string `xml:",chardata" sql:"type:text"`
}
