package udetect

// App information
type App struct {
	ExtID         string `json:"eid,omitempty"`          // External ID
	Keywords      string `json:"keywords,omitempty"`     // Comma separated list of keywords about the site.
	R0Cat         []uint `json:"r0cat,omitempty"`        // r0 category ids
	Bundle        string `json:"bundle,omitempty"`       // App bundle or package name
	StoreURL      string `json:"storeurl,omitempty"`     // App store URL for an installed app
	Ver           string `json:"ver,omitempty"`          // App version
	Paid          int    `json:"paid,omitempty"`         // "1": Paid, "2": Free
	PrivacyPolicy int    `json:"pivacypolicy,omitempty"` // Default: 1 ("1": has a privacy policy)
}

// Content10Codes returns Content Taxonomy 1.0 codes for the r0 ids. An id without that code is skipped.
func (a *App) Content10Codes() []string {
	if a == nil {
		return nil
	}
	return content10Codes(a.R0Cat)
}

// AppDefault object
var AppDefault App

// DomainPrepared value
func (a *App) DomainPrepared() []string {
	return PrepareDomain(a.Bundle)
}
