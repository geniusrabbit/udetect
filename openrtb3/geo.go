package openrtb3

import (
	"github.com/bsm/openrtb/v3"
	"github.com/geniusrabbit/udetect"
)

// GeoFrom returns OpenRTB Geo type
func GeoFrom(g *udetect.Geo) *openrtb.Geo {
	return &openrtb.Geo{
		Latitude:      g.Lat,           // Latitude from -90 to 90
		Longitude:     g.Lon,           // Longitude from -180 to 180
		Type:          0,               // Indicate the source of the geo data
		Accuracy:      0,               // Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services
		LastFix:       0,               // Number of seconds since this geolocation fix was established.
		IPService:     0,               // Service or provider used to determine geolocation from IP address if applicable
		Country:       g.Country,       // Country using ISO 3166-1 Alpha 3
		Region:        g.Region,        // Region using ISO 3166-2
		RegionFIPS104: g.RegionFIPS104, // Region of a country using FIPS 10-4
		Metro:         g.Metro,         //
		City:          g.City,          //
		ZIP:           g.ZIP,           //
		UTCOffset:     g.UTCOffset,     // Local time as the number +/- of minutes from UTC
		Ext:           nil,             //
	}
}
