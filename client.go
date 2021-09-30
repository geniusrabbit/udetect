package udetect

import (
	"context"
	"net"
	"strings"

	"github.com/google/uuid"

	"github.com/sspserver/udetect/protocol"
)

// Request object
type Request struct {
	UDID            string      `json:"udid,omitempty"` // Advertisement Device ID (IDFA, AAID), Identifier for Advertising (IFA)
	UID             uuid.UUID   `json:"uid,omitempty"`
	SessID          uuid.UUID   `json:"sessid,omitempty"`
	IP              string      `json:"ip,omitempty"`
	UA              string      `json:"ua,omitempty"`
	URL             string      `json:"url,omitempty"`
	Ref             string      `json:"r,omitempty"`     // Referer
	DNT             int8        `json:"dnt,omitempty"`   // "1": Do not track
	LMT             int8        `json:"lmt,omitempty"`   // "1": Limit Ad Tracking
	Adblock         int8        `json:"ab,omitempty"`    // "1": AdBlock is ON
	PrivateBrowsing int8        `json:"pb,omitempty"`    // "1": Private Browsing mode ON
	JS              int8        `json:"js,omitempty"`    //
	Languages       []string    `json:"langs,omitempty"` //
	PrimaryLanguage string      `json:"lang,omitempty"`  // Browser language (en-US)
	FlashVer        string      `json:"flver,omitempty"` // Flash version
	Width           int         `json:"w,omitempty"`     // Window in pixels
	Height          int         `json:"h,omitempty"`     // Window in pixels
	Extensions      []Extension `json:"extensions,omitempty"`
}

// Response object
type Response struct {
	User   *User   `json:"user"`
	Device *Device `json:"device"`
	Geo    *Geo    `json:"geo"`
}

// Client implementation fordata requests
type Client struct {
	transport Transport
}

// NewClient for udetect API
func NewClient(tr Transport) *Client {
	return &Client{transport: tr}
}

// Detect user data information
func (c *Client) Detect(ctx context.Context, req *Request) (*Response, error) {
	uid, err := protocol.UUIDFrom(req.UID)
	if err != nil {
		return nil, err
	}
	sessid, err := protocol.UUIDFrom(req.SessID)
	if err != nil {
		return nil, err
	}
	nReq := &protocol.Request{
		Udid:            req.UDID,
		Uid:             uid,
		Sessid:          sessid,
		Ip:              req.IP,
		Ua:              req.UA,
		Url:             req.URL,
		Ref:             req.Ref,
		DNT:             req.DNT == 1,
		LMT:             req.LMT == 1,
		Adblock:         req.Adblock == 1,
		PrivateBrowsing: req.PrivateBrowsing == 1,
		Js:              req.JS == 1,
		Languages:       req.Languages,
		PrimaryLanguage: req.PrimaryLanguage,
		FlashVer:        req.FlashVer,
		Width:           int32(req.Width),
		Height:          int32(req.Height),
	}
	resp, err := c.transport.Detect(ctx, nReq)
	if err != nil {
		return nil, err
	}
	var userUUID uuid.UUID
	if resp.User.GetUuid() != nil {
		userUUID, err = uuid.Parse(resp.User.GetUuid().String())
		if err != nil {
			return nil, err
		}
	} else {
		userUUID = uuid.New()
	}
	return &Response{
		User: &User{
			UUID:          userUUID,
			SessionID:     resp.User.GetSessid().String(),
			FingerPrintID: resp.User.GetFingerprint(),
			AgeStart:      int(resp.User.GetAgeStart()),
			AgeEnd:        int(resp.User.GetAgeEnd()),
			Keywords:      strings.Join(resp.User.GetKeywords(), ","),
			Interests:     nil,
			Sex:           nil,
		},
		Device: &Device{
			ID:    uint(resp.Device.GetId()),
			Make:  resp.Device.GetMake(),
			Model: resp.Device.GetModel(),
			OS: &OS{
				ID:      uint(resp.Device.Os.GetId()),
				Name:    resp.Device.Os.GetName(),
				Version: resp.Device.Os.GetVersion(),
			},
			Browser: &Browser{
				ID:              resp.Device.Browser.GetId(),
				Name:            resp.Device.Browser.GetName(),
				Version:         resp.Device.Browser.GetVersion(),
				DNT:             int(req.DNT),
				LMT:             int(req.LMT),
				Adblock:         int(req.Adblock),
				PrivateBrowsing: int(req.PrivateBrowsing),
				UA:              req.UA,
				Ref:             req.Ref,
				JS:              int(req.JS),
				Languages:       req.Languages,
				PrimaryLanguage: req.PrimaryLanguage,
				FlashVer:        req.FlashVer,
				Width:           req.Width,
				Height:          req.Height,
				IsRobot:         int(resp.Device.Browser.GetIsRobot()),
				Extensions:      req.Extensions,
			},
			IFA:        req.UDID,
			ConnType:   int(resp.Device.GetConnectiontype()),
			DeviceType: DeviceType(resp.Device.GetDevicetype()),
			Height:     int(resp.Device.GetHeight()),
			Width:      int(resp.Device.GetWidth()),
			PPI:        int(resp.Device.GetPpi()),
			PxRatio:    float64(resp.Device.GetPxRatio()),
			HwVer:      resp.Device.GetHwver(),
		},
		Geo: &Geo{
			ID:            uint(resp.Geo.GetId()),
			IP:            net.ParseIP(resp.Geo.GetIp()),
			Lat:           float64(resp.Geo.GetLat()),
			Lon:           float64(resp.Geo.GetLon()),
			Country:       resp.Geo.GetCountry(),
			Region:        resp.Geo.GetRegion(),
			RegionFIPS104: resp.Geo.GetRegionFIPS104(),
			Metro:         resp.Geo.GetMetro(),
			City:          resp.Geo.GetCity(),
			Zip:           resp.Geo.GetZip(),
			UTCOffset:     int(resp.Geo.GetId()),
			Carrier: &Carrier{
				ID:   uint(resp.Geo.GetCarrier().GetId()),
				Name: resp.Geo.GetCarrier().GetName(),
				Code: resp.Geo.GetCarrier().GetCode(),
			},
		},
	}, nil
}

// Close client connection
func (c *Client) Close() error {
	if c.transport != nil {
		return c.transport.Close()
	}
	return nil
}
