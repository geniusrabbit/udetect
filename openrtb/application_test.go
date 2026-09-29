package openrtb

import (
	"testing"

	"github.com/geniusrabbit/udetect"
)

func TestApplicationFromContent10(t *testing.T) {
	app := ApplicationFrom(&udetect.App{R0Cat: []uint{40}})
	if app == nil {
		t.Fatal("expected app")
	}
	if len(app.Cat) != 2 || app.Cat[0] != "IAB1-1" || app.Cat[1] != "IAB9-11" {
		t.Fatalf("cat=%v", app.Cat)
	}
}

func TestSiteFromContent10(t *testing.T) {
	site := SiteFrom(&udetect.Site{R0Cat: []uint{40}})
	if site == nil {
		t.Fatal("expected site")
	}
	if len(site.Cat) != 2 || site.Cat[0] != "IAB1-1" || site.Cat[1] != "IAB9-11" {
		t.Fatalf("cat=%v", site.Cat)
	}
}
