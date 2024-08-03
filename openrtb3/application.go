package openrtb3

import (
	"github.com/bsm/openrtb/v3"
	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/udetect"
)

// ApplicationFrom returnds openrtb App type
func ApplicationFrom(a *udetect.App) *openrtb.App {
	if a == nil {
		return nil
	}
	return &openrtb.App{
		Inventory: openrtb.Inventory{
			ID:            a.ExtID,                                      // External ID
			Keywords:      a.Keywords,                                   // Comma separated list of keywords about the site.
			Categories:    gocast.Slice[openrtb.ContentCategory](a.Cat), // Array of IAB content categories
			PrivacyPolicy: intRef(a.PrivacyPolicy),                      // Default: 1 ("1": has a privacy policy)
		},
		Bundle:   a.Bundle,   // App bundle or package name
		StoreURL: a.StoreURL, // App store URL for an installed app
		Version:  a.Ver,      // App version
		Paid:     a.Paid,     // "1": Paid, "2": Free
	}
}
