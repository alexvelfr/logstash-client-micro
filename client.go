package logstashclientmicro

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	serviceName string
	url         string
	httpClient  *http.Client
}

func (c *client) LogError(ctx context.Context, msg Message) error {
	return c.logError(ctx, msg)
}

func (c *client) logError(ctx context.Context, msg Message) error {
	var errStr string
	if msg.Error != nil {
		errStr = msg.Error.Error()
	}
	m := message{
		Message:      msg,
		Microservice: c.serviceName,
		Error:        errStr,
	}
	btReq, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewReader(btReq))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}
	_, err = c.httpClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

// NewClient return client instance by conf
// serviceName - name of microservice
// url - logstash url with http input
// insecureTLS - flag, can use self-sign certs
func NewClient(serviceName, uri string, insecureTLS bool) Client {
	var httpClient *http.Client
	if insecureTLS {
		httpClient = newInsecureHTTPClient()
	} else {
		httpClient = newHTTPClient()
	}
	base, err := url.Parse(uri)
	if err != nil {
		return nil
	}
	return &client{
		serviceName: serviceName,
		url:         base.String(),
		httpClient:  httpClient,
	}
}

func newInsecureHTTPClient() *http.Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{
		Transport: tr,
		Timeout:   120 * time.Second,
	}
	return client
}
func newHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	return client
}
