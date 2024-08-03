package udetect

// Device base information structure
type Device struct {
	ID         uint       `json:"id,omitempty"`             // Device ID
	Make       string     `json:"make,omitempty"`           // Device make
	Model      string     `json:"model,omitempty"`          // Device model
	OS         *OS        `json:"os,omitempty"`             // Device OS
	Browser    *Browser   `json:"browser,omitempty"`        // Device OS version
	ConnType   int        `json:"connectiontype,omitempty"` //
	DeviceType DeviceType `json:"devicetype,omitempty"`     //
	IFA        string     `json:"ifa,omitempty"`            // Native identifier for advertisers
	Height     int        `json:"h,omitempty"`              // Physical height of the screen in pixels.
	Width      int        `json:"w,omitempty"`              // Physical width of the screen in pixels.
	PPI        int        `json:"ppi,omitempty"`            // Screen size as pixels per linear inch.
	PxRatio    float64    `json:"pxratio,omitempty"`        // The ratio of physical pixels to device independent pixels.
	HwVer      string     `json:"hwv,omitempty"`            // Hardware version of the device (e.g., "5S" for iPhone 5S).
}

// DeviceDefault value
var DeviceDefault = Device{Browser: &BrowserDefault, OS: &OSDefault}
