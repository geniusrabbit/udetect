package api

import (
	"context"
	"math/rand"

	"github.com/sspserver/udetect/protocol"
)

// Detector API implementation
type Detector struct {
	protocol.UnimplementedDetectorServer
}

// Detect method must implement user information detection and return the whole information about him
func (d *Detector) Detect(ctx context.Context, req *protocol.Request) (*protocol.Response, error) {
	age := rand.Int31n(100)
	return &protocol.Response{
		User: &protocol.User{
			Uuid:        &protocol.UUID{Value: []byte{}},
			Sessid:      &protocol.UUID{Value: []byte{}},
			Fingerprint: "finger",
			AgeStart:    age,
			AgeEnd:      age + 10,
			Keywords:    []string{"sport", "footbol", "cars"}[:rand.Int31n(3)],
		},
		Device: &protocol.Device{
			Id:             1,
			Make:           "test",
			Model:          "test",
			Connectiontype: protocol.ConnType_Ethernet,
			Os: &protocol.OS{
				Id:      1,
				Name:    "My OS",
				Version: "1.2.4-beta",
			},
			Browser: &protocol.Browser{
				Id:      1,
				Name:    "My Browser",
				Version: "1.0.1",
				IsRobot: 1,
			},
		},
		Geo: &protocol.GeoLocation{
			Id: 1,
			Carrier: &protocol.Carrier{
				Id:   1,
				Name: "test",
				Code: "us-att",
			},
			Country: "US",
			Region:  "California",
		},
	}, nil
}
