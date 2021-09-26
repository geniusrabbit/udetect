package openrtb

import (
	"github.com/bsm/openrtb"

	"github.com/sspserver/udetect"
)

// OpenRTBDeviceType returns constant as OpenRTB type
func OpenRTBDeviceType(dt udetect.DeviceType) int {
	switch dt {
	case udetect.DeviceTypeMobile:
		return openrtb.DeviceTypeMobile
	case udetect.DeviceTypePC:
		return openrtb.DeviceTypePC
	case udetect.DeviceTypeTV:
		return openrtb.DeviceTypeTV
	case udetect.DeviceTypePhone:
		return openrtb.DeviceTypePhone
	case udetect.DeviceTypeTablet:
		return openrtb.DeviceTypeTablet
	case udetect.DeviceTypeConnected:
		return openrtb.DeviceTypeConnected
	case udetect.DeviceTypeSetTopBox:
		return openrtb.DeviceTypeSetTopBox
	case udetect.DeviceTypeWatch, udetect.DeviceTypeGlasses:
	}
	return openrtb.DeviceTypeUnknown
}

// DeviceFrom open RTB Device type
func DeviceFrom(d *udetect.Device, geo *udetect.Geo) *openrtb.Device {
	if d == nil {
		return nil
	}
	var (
		browser = d.Browser
		os      = d.OS
		carrier *udetect.Carrier
		ipV4    = geo.IPv4String()
	)
	if browser == nil {
		browser = &udetect.BrowserDefault
	}
	if os == nil {
		os = &udetect.OSDefault
	}
	if geo == nil {
		geo = &udetect.GeoDefault
	}
	if carrier = geo.Carrier; carrier == nil {
		carrier = &udetect.CarrierDefault
	}

	// IP by default
	if ipV4 == "" && geo.IPv6String() == "" {
		ipV4 = "0.0.0.0"
	}

	return &openrtb.Device{
		UA:         browser.UA,                      // User agent
		Geo:        GeoFrom(geo),                    // Location of the device assumed to be the user’s current location
		DNT:        browser.DNT,                     // "1": Do not track
		LMT:        browser.LMT,                     // "1": Limit Ad Tracking
		IP:         ipV4,                            // IPv4
		IPv6:       geo.IPv6String(),                // IPv6
		DeviceType: OpenRTBDeviceType(d.DeviceType), // The general type of d.
		Make:       d.Make,                          // Device make
		Model:      d.Model,                         // Device model
		OS:         os.Name,                         // Device OS
		OSVer:      os.Version,                      // Device OS version
		HwVer:      d.HwVer,                         // Hardware version of the device (e.g., "5S" for iPhone 5S).
		H:          d.Height,                        // Physical height of the screen in pixels.
		W:          d.Width,                         // Physical width of the screen in pixels.
		PPI:        d.PPI,                           // Screen size as pixels per linear inch.
		PxRatio:    d.PxRatio,                       // The ratio of physical pixels to device independent pixels.
		JS:         browser.JS,                      // Javascript status ("0": Disabled, "1": Enabled)
		GeoFetch:   0,                               // Indicates if the geolocation API will be available to JavaScript code running in the banner,
		FlashVer:   browser.FlashVer,                // Flash version
		Language:   browser.PrimaryLanguage,         // Browser language
		Carrier:    carrier.Name,                    // Carrier or ISP derived from the IP address
		MCCMNC:     "",                              // Mobile carrier as the concatenated MCC-MNC code (e.g., "310-005" identifies Verizon Wireless CDMA in the USA).
		ConnType:   d.ConnType,                      // Network connection type.
		IFA:        d.IFA,                           // Native identifier for advertisers
		IDSHA1:     "",                              // SHA1 hashed device ID
		IDMD5:      "",                              // MD5 hashed device ID
		PIDSHA1:    "",                              // SHA1 hashed platform device ID
		PIDMD5:     "",                              // MD5 hashed platform device ID
		MacSHA1:    "",                              // SHA1 hashed device ID; IMEI when available, else MEID or ESN
		MacMD5:     "",                              // MD5 hashed device ID; IMEI when available, else MEID or ESN
	}
}
