package gobuy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"text/template"
	"time"
)

const (
	scheme      = "https"
	baseUrlTmpl = "{{.Scheme}}://{{.ShopDomain}}/"
)

type BuyClient struct {
	client     *http.Client
	shopDomain string
	appName    string
	apiKey     string
	appId      string
}

// Client creates a new BuyClient with the given options
func Client(shopDomain, appName, apiKey, appId string) *BuyClient {
	return &BuyClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		shopDomain: shopDomain,
		appName:    appName,
		apiKey:     apiKey,
		appId:      appId,
	}
}

type header map[string]string

func (h *header) Add(key string, value string) {
	(*h)[key] = value
}

func (h *header) Set(key string, value string) {
	(*h)[key] = value
}

func (h *header) Get(key string) string {
	return (*h)[key]
}

func (h *header) Del(key string) {
	delete((*h), key)
}

// basePathParams are the default URL path paramaters that can be used
type basePathParams struct {
	Scheme     string
	ShopDomain string
	AppId      string
	AppName    string
}

type requestOptions struct {
	method      string
	urlTemplate string
	body        io.Reader
	bodyObj     interface{}
	pathParams  interface{}
	queryParams map[string]string
}

func (b *BuyClient) buildRequest(opts requestOptions, h header) (*http.Request, error) {
	var path string = opts.urlTemplate
	// Apply the path params to the path template
	if opts.pathParams != nil {
		t, err := template.New("REQ Path Template").Parse(opts.urlTemplate)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBufferString("")
		err = t.Execute(buf, opts.pathParams)
		if err != nil {
			return nil, err
		}
		path = buf.String()
	}

	// Apply the base Path params to the url
	t, err := template.New("REQ URL Template").Parse(baseUrlTmpl + path)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBufferString("")
	err = t.Execute(buf, basePathParams{Scheme: scheme, ShopDomain: b.shopDomain, AppId: b.appId, AppName: b.appName})
	if err != nil {
		return nil, err
	}
	url := buf.String()

	if opts.bodyObj != nil {
		b, err := json.Marshal(opts.bodyObj)
		if err != nil {
			return nil, err
		}
		opts.body = bytes.NewBuffer(b)
	}

	// Create the new request
	req, err := http.NewRequest(opts.method, url, opts.body)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range opts.queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	for k, v := range h {
		req.Header.Add(k, v)
	}
	return req, nil
}

func (b *BuyClient) buildShopifyRequest(opts requestOptions, extraHeaders ...header) (*http.Request, error) {
	defaultHeaders := make(header)
	defaultHeaders.Add("Authorization", formatBasicAuthorization(b.apiKey))
	defaultHeaders.Add("Content-Type", "application/json")

	if len(extraHeaders) > 0 {
		for _, h := range extraHeaders {
			for k, v := range h {
				defaultHeaders.Add(k, v)
			}
		}
	}
	return b.buildRequest(opts, defaultHeaders)
}

func (b *BuyClient) sendShopifyRequest(opts requestOptions, extraHeaders ...header) (*shopifyResponse, error) {
	req, err := b.buildShopifyRequest(opts, extraHeaders...)
	if err != nil {
		return nil, err
	}
	return b.send(req)
}

func (b BuyClient) send(req *http.Request) (*shopifyResponse, error) {
	rsp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	r := &shopifyResponse{rsp}
	if r.StatusCode != 200 {
		return nil, r.GetError()
	}
	return r, nil
}

type shopifyResponse struct {
	*http.Response
}

func (r *shopifyResponse) JsonDecode(v interface{}) error {
	d := json.NewDecoder(r.Body)
	return d.Decode(v)
}

func (r *shopifyResponse) GetError() error {
	d := json.NewDecoder(r.Body)
	type ShopifyErr struct {
		Error string `json:"error"`
	}
	shopErr := &ShopifyErr{Error: "Could not decode response"}
	err := d.Decode(&shopErr)
	if err != nil {
		return err
	}
	return errors.New(fmt.Sprintf("%d: %s", r.StatusCode, shopErr.Error))
}
