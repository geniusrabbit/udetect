package openrtb3

import (
	"github.com/bsm/openrtb/v3"
	"github.com/demdxx/gocast/v2"

	"github.com/geniusrabbit/udetect"
)

// SiteFrom returns OpenRTB Site type
func SiteFrom(s *udetect.Site) *openrtb.Site {
	if s == nil {
		return nil
	}
	return &openrtb.Site{
		Inventory: openrtb.Inventory{
			ID:            s.ExtID,                                      // External ID
			Keywords:      s.Keywords,                                   // Comma separated list of keywords about the site.
			Categories:    gocast.Slice[openrtb.ContentCategory](s.Cat), // Array of IAB content categories
			Domain:        s.Domain,                                     //
			PrivacyPolicy: intRef(s.PrivacyPolicy),                      // Default: 1 ("1": has a privacy policy)
		},
		Page:     s.Page,     // URL of the page
		Referrer: s.Referrer, // Referrer URL
		Search:   s.Search,   // Search string that caused naviation
		Mobile:   s.Mobile,   // Mobile ("1": site is mobile optimised)
	}
}
