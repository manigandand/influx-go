package influxdb

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func newTransport() *http.Transport {
	return &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
}

func defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout:   time.Second * 20,
		Transport: newTransport(),
	}
}

// HTTPClientWithTLSConfig returns an *http.Client with sane timeouts and the provided TLSClientConfig.
func HTTPClientWithTLSConfig(conf *tls.Config) *http.Client {
	return &http.Client{
		Timeout: time.Second * 20,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig:     conf,
		},
	}
}
