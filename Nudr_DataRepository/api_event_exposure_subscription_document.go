/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudr_DataRepository

import (
	"free5gc/lib/openapi"
	"free5gc/lib/openapi/models"

	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type EventExposureSubscriptionDocumentApiService service

/*
EventExposureSubscriptionDocumentApiService Deletes a eeSubscription for a group of UEs or any UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueGroupId
 * @param subsId Unique ID of the subscription to remove
*/

func (a *EventExposureSubscriptionDocumentApiService) RemoveEeGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", fmt.Sprintf("%v", ueGroupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", fmt.Sprintf("%v", subsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}
	_ = apiError

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return localVarHTTPResponse, nil
	default:
		return localVarHTTPResponse, openapi.ReportError("%d is not a valid status code in RemoveEeGroupSubscriptions", localVarHTTPResponse.StatusCode)
	}
}

/*
EventExposureSubscriptionDocumentApiService Deletes a eeSubscription
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueId
 * @param subsId Unique ID of the subscription to remove
*/

func (a *EventExposureSubscriptionDocumentApiService) RemoveeeSubscriptions(ctx context.Context, ueId string, subsId string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", fmt.Sprintf("%v", ueId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", fmt.Sprintf("%v", subsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}
	_ = apiError

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return localVarHTTPResponse, nil
	default:
		return localVarHTTPResponse, openapi.ReportError("%d is not a valid status code in RemoveeeSubscriptions", localVarHTTPResponse.StatusCode)
	}
}

/*
EventExposureSubscriptionDocumentApiService Stores an individual ee subscription of a group of UEs or any UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueGroupId
 * @param subsId
 * @param optional nil or *UpdateEeGroupSubscriptionsParamOpts - Optional Parameters:
 * @param "EeSubscription" (optional.Interface of EeSubscription) -
*/

type UpdateEeGroupSubscriptionsParamOpts struct {
	EeSubscription optional.Interface
}

func (a *EventExposureSubscriptionDocumentApiService) UpdateEeGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string, localVarOptionals *UpdateEeGroupSubscriptionsParamOpts) (*http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", fmt.Sprintf("%v", ueGroupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", fmt.Sprintf("%v", subsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	if localVarOptionals != nil && localVarOptionals.EeSubscription.IsSet() {
		localVarOptionalEeSubscription, localVarOptionalEeSubscriptionok := localVarOptionals.EeSubscription.Value().(models.EeSubscription)
		if !localVarOptionalEeSubscriptionok {
			return nil, openapi.ReportError("eeSubscription should be EeSubscription")
		}
		localVarPostBody = &localVarOptionalEeSubscription
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return localVarHTTPResponse, nil
	default:
		var v models.ProblemDetails
		err = openapi.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	}
}

/*
EventExposureSubscriptionDocumentApiService Stores an individual ee subscriptions of a UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueId
 * @param subsId
 * @param optional nil or *UpdateEesubscriptionsParamOpts - Optional Parameters:
 * @param "EeSubscription" (optional.Interface of EeSubscription) -
*/

type UpdateEesubscriptionsParamOpts struct {
	EeSubscription optional.Interface
}

func (a *EventExposureSubscriptionDocumentApiService) UpdateEesubscriptions(ctx context.Context, ueId string, subsId string, localVarOptionals *UpdateEesubscriptionsParamOpts) (*http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", fmt.Sprintf("%v", ueId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", fmt.Sprintf("%v", subsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	if localVarOptionals != nil && localVarOptionals.EeSubscription.IsSet() {
		localVarOptionalEeSubscription, localVarOptionalEeSubscriptionok := localVarOptionals.EeSubscription.Value().(models.EeSubscription)
		if !localVarOptionalEeSubscriptionok {
			return nil, openapi.ReportError("eeSubscription should be EeSubscription")
		}
		localVarPostBody = &localVarOptionalEeSubscription
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return localVarHTTPResponse, nil
	default:
		var v models.ProblemDetails
		err = openapi.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	}
}
