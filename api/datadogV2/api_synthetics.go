// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsApi service type
type SyntheticsApi datadog.Service

type apiGetOnDemandConcurrencyCapRequest struct {
	ctx _context.Context
}

func (a *SyntheticsApi) buildGetOnDemandConcurrencyCapRequest(ctx _context.Context) (apiGetOnDemandConcurrencyCapRequest, error) {
	req := apiGetOnDemandConcurrencyCapRequest{
		ctx: ctx,
	}
	return req, nil
}

// GetOnDemandConcurrencyCap Get the on-demand concurrency cap.
// Get the on-demand concurrency cap.
func (a *SyntheticsApi) GetOnDemandConcurrencyCap(ctx _context.Context) (OnDemandConcurrencyCapResponse, *_nethttp.Response, error) {
	req, err := a.buildGetOnDemandConcurrencyCapRequest(ctx)
	if err != nil {
		var localVarReturnValue OnDemandConcurrencyCapResponse
		return localVarReturnValue, nil, err
	}

	return a.getOnDemandConcurrencyCapExecute(req)
}

// getOnDemandConcurrencyCapExecute executes the request.
func (a *SyntheticsApi) getOnDemandConcurrencyCapExecute(r apiGetOnDemandConcurrencyCapRequest) (OnDemandConcurrencyCapResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue OnDemandConcurrencyCapResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.SyntheticsApi.GetOnDemandConcurrencyCap")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/synthetics/settings/on_demand_concurrency_cap"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiSetOnDemandConcurrencyCapRequest struct {
	ctx  _context.Context
	body *OnDemandConcurrencyCapAttributes
}

func (a *SyntheticsApi) buildSetOnDemandConcurrencyCapRequest(ctx _context.Context, body OnDemandConcurrencyCapAttributes) (apiSetOnDemandConcurrencyCapRequest, error) {
	req := apiSetOnDemandConcurrencyCapRequest{
		ctx:  ctx,
		body: &body,
	}
	return req, nil
}

// SetOnDemandConcurrencyCap Save new value for on-demand concurrency cap.
// Save new value for on-demand concurrency cap.
func (a *SyntheticsApi) SetOnDemandConcurrencyCap(ctx _context.Context, body OnDemandConcurrencyCapAttributes) (OnDemandConcurrencyCapResponse, *_nethttp.Response, error) {
	req, err := a.buildSetOnDemandConcurrencyCapRequest(ctx, body)
	if err != nil {
		var localVarReturnValue OnDemandConcurrencyCapResponse
		return localVarReturnValue, nil, err
	}

	return a.setOnDemandConcurrencyCapExecute(req)
}

// setOnDemandConcurrencyCapExecute executes the request.
func (a *SyntheticsApi) setOnDemandConcurrencyCapExecute(r apiSetOnDemandConcurrencyCapRequest) (OnDemandConcurrencyCapResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue OnDemandConcurrencyCapResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.SyntheticsApi.SetOnDemandConcurrencyCap")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/synthetics/settings/on_demand_concurrency_cap"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewSyntheticsApi Returns NewSyntheticsApi.
func NewSyntheticsApi(client *datadog.APIClient) *SyntheticsApi {
	return &SyntheticsApi{
		Client: client,
	}
}
