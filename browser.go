package udetect

import "github.com/demdxx/langlib"

// Browser base information structure
type Browser struct {
	ID              uint64         `json:"id,omitempty"`   // Internal system ID
	Name            string         `json:"name,omitempty"` //
	Version         string         `json:"ver,omitempty"`  //
	DNT             int8           `json:"dnt,omitempty"`  // "1": Do not track
	LMT             int8           `json:"lmt,omitempty"`  // "1": Limit Ad Tracking
	AdBlock         int8           `json:"ab,omitempty"`   // "1": AdBlock is ON
	PrivateBrowsing int8           `json:"pb,omitempty"`   // "1": Private Browsing mode ON
	IsRobot         int8           `json:"rb,omitempty"`
	JS              int8           `json:"js,omitempty"`    //
	UA              string         `json:"ua,omitempty"`    // User agent
	Ref             string         `json:"r,omitempty"`     // Referer
	Languages       []langlib.Code `json:"langs,omitempty"` // ISO-639-1
	PrimaryLanguage langlib.Code   `json:"lang,omitempty"`  // ISO-639-1
	FlashVer        string         `json:"flver,omitempty"` // Flash version
	Width           int            `json:"w,omitempty"`     // Window in pixels
	Height          int            `json:"h,omitempty"`     // Window in pixels
	Extensions      []Extension    `json:"extensions,omitempty"`
}

// Extension of some Browser/OS
type Extension struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"ver,omitempty"`
}

// BrowserDefault value
var BrowserDefault Browser

// Clone a new Browser
func (b *Browser) Clone() *Browser {
	return &Browser{
		ID:              b.ID,
		Name:            b.Name,
		Version:         b.Version,
		DNT:             b.DNT,
		LMT:             b.LMT,
		AdBlock:         b.AdBlock,
		PrivateBrowsing: b.PrivateBrowsing,
		IsRobot:         b.IsRobot,
		JS:              b.JS,
		UA:              b.UA,
		Ref:             b.Ref,
		Languages:       b.Languages,
		PrimaryLanguage: b.PrimaryLanguage,
		FlashVer:        b.FlashVer,
		Width:           b.Width,
		Height:          b.Height,
		Extensions:      b.Extensions,
	}
}

// LanguageISO2 returns the catalog ISO-639-1 code, or "" when the code is undefined.
func LanguageISO2(code langlib.Code) string {
	if code.IsUndefined() {
		return ""
	}
	return code.ISO2()
}

// LanguagesISO2 maps codes to ISO-639-1 strings. Undefined codes become "".
// An empty input returns nil.
func LanguagesISO2(codes []langlib.Code) []string {
	if len(codes) == 0 {
		return nil
	}
	out := make([]string, len(codes))
	for i, code := range codes {
		out[i] = LanguageISO2(code)
	}
	return out
}
