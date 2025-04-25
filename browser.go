package udetect

// Browser base information structure
type Browser struct {
	ID              uint64      `json:"id,omitempty"`   // Internal system ID
	Name            string      `json:"name,omitempty"` //
	Version         string      `json:"ver,omitempty"`  //
	DNT             int8        `json:"dnt,omitempty"`  // "1": Do not track
	LMT             int8        `json:"lmt,omitempty"`  // "1": Limit Ad Tracking
	Adblock         int8        `json:"ab,omitempty"`   // "1": AdBlock is ON
	PrivateBrowsing int8        `json:"pb,omitempty"`   // "1": Private Browsing mode ON
	IsRobot         int8        `json:"rb,omitempty"`
	JS              int8        `json:"js,omitempty"`    //
	UA              string      `json:"ua,omitempty"`    // User agent
	Ref             string      `json:"r,omitempty"`     // Referer
	Languages       []string    `json:"langs,omitempty"` //
	PrimaryLanguage string      `json:"lang,omitempty"`  // Browser language (en-US)
	FlashVer        string      `json:"flver,omitempty"` // Flash version
	Width           int         `json:"w,omitempty"`     // Window in pixels
	Height          int         `json:"h,omitempty"`     // Window in pixels
	Extensions      []Extension `json:"extensions,omitempty"`
}

// Extension of some Browser/OS
type Extension struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"ver,omitempty"`
}

// BrowserDefault value
var BrowserDefault Browser
