package http

import (
	"net"
	"net/http"
	"time"
)

// NewDefaultHTTPClient provides http.Client
func NewDefaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 1 * time.Minute,
			}).DialContext,
			IdleConnTimeout:       1 * time.Minute,
			ResponseHeaderTimeout: 5 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
		},
		Timeout: 5 * time.Second,
	}
}
