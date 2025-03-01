/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ntsctsf_TimeSynchronization

import (
    "crypto/tls"
    "golang.org/x/net/http2"
    "net/http"
)

// APIClient manages communication with the Ntsctsf_TimeSynchronization Service API API v1.0.2
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
    cfg    *Configuration
    common service // Reuse a single struct instead of allocating one for each service on the heap.

    // API Services
    	IndividualTimeSynchronizationExposureConfigurationDocumentApi *IndividualTimeSynchronizationExposureConfigurationDocumentApiService
    	IndividualTimeSynchronizationExposureSubscriptionDocumentApi *IndividualTimeSynchronizationExposureSubscriptionDocumentApiService
    	TimeSynchronizationExposureSubscriptionsCollectionApi *TimeSynchronizationExposureSubscriptionsCollectionApiService
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
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
    }

    c := &APIClient{}
    c.cfg = cfg
    c.common.client = c

    // API Services
    c.IndividualTimeSynchronizationExposureConfigurationDocumentApi = (*IndividualTimeSynchronizationExposureConfigurationDocumentApiService)(&c.common)
    c.IndividualTimeSynchronizationExposureSubscriptionDocumentApi = (*IndividualTimeSynchronizationExposureSubscriptionDocumentApiService)(&c.common)
    c.TimeSynchronizationExposureSubscriptionsCollectionApi = (*TimeSynchronizationExposureSubscriptionsCollectionApiService)(&c.common)

    return c
}
