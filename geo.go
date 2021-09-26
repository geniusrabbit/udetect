package udetect

import (
	"net"
)

// UndefinedCountryCode 2 chars
const UndefinedCountryCode = "**"

// Geo base information structure
type Geo struct {
	ID            uint     `json:"id,omitempty"`            // Internal geo ID
	IP            net.IP   `json:"ip,omitempty"`            // IPv4/6
	Carrier       *Carrier `json:"carrier,omitempty"`       // Carrier or ISP derived from the IP address
	Lat           float64  `json:"lat,omitempty"`           // Latitude from -90 to 90
	Lon           float64  `json:"lon,omitempty"`           // Longitude from -180 to 180
	Country       string   `json:"country,omitempty"`       // Country using ISO 3166-1 Alpha 2
	Region        string   `json:"region,omitempty"`        // Region using ISO 3166-2
	RegionFIPS104 string   `json:"regionFIPS104,omitempty"` // Region of a country using FIPS 10-4
	Metro         string   `json:"metro,omitempty"`         //
	City          string   `json:"city,omitempty"`          //
	Zip           string   `json:"zip,omitempty"`           //
	UTCOffset     int      `json:"utcoffset,omitempty"`     // Local time as the number +/- of minutes from UTC
}

// GeoDefault value
var GeoDefault = Geo{Country: UndefinedCountryCode, Carrier: &CarrierDefault}

// IsIPv6 format
func (g *Geo) IsIPv6() bool {
	return g.IP != nil && g.IP.To4() == nil
}

// IPv4String string value
func (g *Geo) IPv4String() string {
	if g == nil || g.IP == nil {
		return ""
	}
	if g.IsIPv6() {
		return ""
	}
	return g.IP.String()
}

// IPv6String string value
func (g *Geo) IPv6String() string {
	if g == nil || g.IP == nil {
		return ""
	}
	if !g.IsIPv6() {
		return ""
	}
	return g.IP.String()
}
