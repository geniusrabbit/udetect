package udetect

// App information
type App struct {
	ExtID         string   `json:"eid,omitempty"`          // External ID
	Keywords      string   `json:"keywords,omitempty"`     // Comma separated list of keywords about the site.
	Cat           []string `json:"cat,omitempty"`          // Array of categories
	Bundle        string   `json:"bundle,omitempty"`       // App bundle or package name
	StoreURL      string   `json:"storeurl,omitempty"`     // App store URL for an installed app
	Ver           string   `json:"ver,omitempty"`          // App version
	Paid          int      `json:"paid,omitempty"`         // "1": Paid, "2": Free
	PrivacyPolicy int      `json:"pivacypolicy,omitempty"` // Default: 1 ("1": has a privacy policy)
}

// AppDefault object
var AppDefault App

// DomainPrepared value
func (a *App) DomainPrepared() []string {
	return PrepareDomain(a.Bundle)
}
