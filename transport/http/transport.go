package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sspserver/udetect/protocol"
)

const (
	macroIP        = "{ip}"
	macroUserAgent = "{ua}"
)

// Transport for user specific detection
type Transport struct {
	url             string
	method          string
	isPreparableURL bool
	client          *http.Client
}

// NewTransport returns transport interface with HTTP protocol support
func NewTransport(url string, options ...Option) *Transport {
	cli := &Transport{
		url:             url,
		method:          http.MethodPost,
		isPreparableURL: strings.Contains(url, macroIP) || strings.Contains(url, macroUserAgent),
		client:          &http.Client{},
	}
	for _, opt := range options {
		opt(cli)
	}
	return cli
}

// Detect information from request
func (tr *Transport) Detect(req *protocol.Request) (resp *protocol.Response, err error) {
	var (
		httpRequest  *http.Request
		httpResponse *http.Response
		body         *bytes.Buffer
	)
	if tr.method == http.MethodPost {
		var data bytes.Buffer
		if err = json.NewEncoder(&data).Encode(req); err != nil {
			return nil, err
		}
		body = &data
	}
	httpRequest, err = http.NewRequest(tr.method, tr.preparedURL(req), body)
	if err != nil {
		return nil, err
	}
	httpResponse, err = tr.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	resp = &protocol.Response{}
	err = json.NewDecoder(httpResponse.Body).Decode(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (tr *Transport) preparedURL(req *protocol.Request) string {
	if !tr.isPreparableURL {
		return tr.url
	}
	replacer := strings.NewReplacer(macroIP, req.Ip, macroUserAgent, req.Ua)
	return replacer.Replace(tr.url)
}
