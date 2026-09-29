package udetect

import (
	"strings"
)

// Site information
type Site struct {
	ExtID         string `json:"eid,omitempty"`          // External ID
	Domain        string `json:"domain,omitempty"`       //
	R0Cat         []uint `json:"r0cat,omitempty"`        // r0 category ids
	PrivacyPolicy int    `json:"pivacypolicy,omitempty"` // Default: 1 ("1": has a privacy policy)
	Keywords      string `json:"keywords,omitempty"`     // Comma separated list of keywords about the site.
	Page          string `json:"page,omitempty"`         // URL of the page
	Referrer      string `json:"ref,omitempty"`          // Referrer URL
	Search        string `json:"search,omitempty"`       // Search string that caused naviation
	Mobile        int    `json:"mobile,omitempty"`       // Mobile ("1": site is mobile optimised)
}

// Content10Codes returns Content Taxonomy 1.0 codes for the r0 ids. An id without that code is skipped.
func (s *Site) Content10Codes() []string {
	if s == nil {
		return nil
	}
	return content10Codes(s.R0Cat)
}

// SiteDefault info
var SiteDefault Site

// DomainPrepared value
func (s *Site) DomainPrepared() []string {
	if s == nil {
		return nil
	}
	return PrepareDomain(s.Domain)
}

// PrepareDomain parts
func PrepareDomain(domain string) (list []string) {
	domain = strings.ToLower(domain)
	if domain == "" {
		return []string{"*."}
	}

	list = make([]string, 0, 5)
	list = append(list, domain)
	if strings.HasPrefix(domain, "www.") {
		list = append(list, "*."+domain)
		domain = domain[4:]
	}

	list = append(list, "*."+domain)
	arr := strings.Split(domain, ".")
	for i := 1; i < len(arr); i++ {
		list = append(list, "*."+strings.Join(arr[i:], "."))
	}

	return list
}
