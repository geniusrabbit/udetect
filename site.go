package udetect

import (
	"strings"
)

// Site information
type Site struct {
	ExtID         string   `json:"eid,omitempty"`          // External ID
	Domain        string   `json:"domain,omitempty"`       //
	Cat           []string `json:"cat,omitempty"`          // Array of categories
	PrivacyPolicy int      `json:"pivacypolicy,omitempty"` // Default: 1 ("1": has a privacy policy)
	Keywords      string   `json:"keywords,omitempty"`     // Comma separated list of keywords about the site.
	Page          string   `json:"page,omitempty"`         // URL of the page
	Ref           string   `json:"ref,omitempty"`          // Referrer URL
	Search        string   `json:"search,omitempty"`       // Search string that caused naviation
	Mobile        int      `json:"mobile,omitempty"`       // Mobile ("1": site is mobile optimised)
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
