package hcaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Secret string
	Client http.Client
	Host   string
}

type SiteVerifyResponse struct {
	Success            bool        `json:"success"`
	ChallengeTimeStamp time.Time   `json:"challenge_ts"`
	HostName           string      `json:"hostname"`
	Credit             bool        `json:"credit,omitempty"`
	ErrorCodes         []ErrorCode `json:"error-codes" bson:"error_codes"`
	Score              float32     `json:"score"`
}

func NewClient(secret string, host ...string) *Client {
	var h string
	if host != nil {
		h = host[0]
	} else {
		h = "https://hcaptcha.com/siteverify"
	}
	return &Client{
		Secret: secret,
		Client: http.Client{},
		Host:   h,
	}
}

func (c *Client) SendRequest(r string, ipAddr ...string) (*SiteVerifyResponse, error) {
	v := url.Values{"secret": {c.Secret}, "response": {r}}
	if ipAddr != nil {
		v.Set("remoteip", ipAddr[0])
	}

	resp, err := c.Client.PostForm(c.Host, v)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := SiteVerifyResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
