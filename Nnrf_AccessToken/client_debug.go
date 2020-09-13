//+build debug

/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_AccessToken

import (
	"crypto/tls"
	"bitbucket.org/free5gc-team/http2_util"
	"net/http"

	"golang.org/x/net/http2"
)

// APIClient manages communication with the NRF OAuth2 API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AccessTokenRequestApi *AccessTokenRequestApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.httpClient == nil {
		cfg.httpClient = http.DefaultClient
		cfg.httpClient.Transport = &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Rand:               http2_util.ZeroSource{},
			},
		}
	}
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccessTokenRequestApi = (*AccessTokenRequestApiService)(&c.common)

	return c
}
