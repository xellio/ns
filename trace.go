package ns

import "github.com/jinzhu/gorm"

//
// Trace ...
//  <trace>...</trace>
//
type Trace struct {
	gorm.Model
	Hops []*Hop `xml:"hop"`
}

//
// Hop ...
//  <hop ttl="1" ipaddr="127.0.0.1" rtt="0.34"/>
//
type Hop struct {
	gorm.Model
	TraceID int
	TTL     int     `xml:"ttl,attr"`
	IPAddr  string  `xml:"ipaddr,attr"`
	RTT     float32 `xml:"rtt,attr"`
	Host    string  `xml:"host,attr"`
}
