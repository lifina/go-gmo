package payment

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/cenkalti/backoff"
)

// Client ... gmo pg payment API client
type Client struct {
	client   *http.Client
	SiteID   string
	SitePass string
	ShopID   string
	ShopPass string
	APIHost  string
}

// NewClient ... new client
func NewClient(
	siteID,
	sitePass,
	shopID,
	shopPass string,
	sandBox bool) (*Client, error) {
	if !(siteID != "" && sitePass != "" && shopID != "" && shopPass != "") {
		return nil, errors.New("Invalid parameters")
	}

	var apiHost string
	if sandBox {
		apiHost = apiHostSandbox
	} else {
		apiHost = apiHostProduction
	}

	return &Client{
		SiteID:   siteID,
		SitePass: sitePass,
		ShopID:   shopID,
		ShopPass: shopPass,
		APIHost:  apiHost,
	}, nil
}

type baseRequestBody struct {
	SiteID   string `json:"SiteID"`
	SitePass string `json:"SitePass"`
	ShopID   string `json:"ShopID"`
	ShopPass string `json:"ShopPass"`
}

func (c *Client) do(
	path string,
	body io.Reader,
	respBody interface{},
) (*http.Response, error) {

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.APIHost, path),
		body,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	var resp *http.Response
	backoffCfg := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 4)
	err = backoff.Retry(func() (err error) {
		resp, err = client.Do(req)
		if err != nil {
			return err
		}
		return nil
	}, backoffCfg)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if respBody != nil {
		if w, ok := respBody.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(respBody)
			if err != nil && err != io.EOF {
				return nil, err
			}
		}
	}

	return resp, nil
}