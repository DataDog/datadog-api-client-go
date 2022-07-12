// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/api/common"
)

// IPRangesApi service type
type IPRangesApi common.Service

type apiGetIPRangesRequest struct {
	ctx _context.Context
	Api *IPRangesApi
}

func (a *IPRangesApi) buildGetIPRangesRequest(ctx _context.Context) (apiGetIPRangesRequest, error) {
	req := apiGetIPRangesRequest{
		Api: a,
		ctx: ctx,
	}
	return req, nil
}

// GetIPRanges List IP Ranges.
// Get information about Datadog IP ranges.
func (a *IPRangesApi) GetIPRanges(ctx _context.Context) (IPRanges, *_nethttp.Response, error) {
	req, err := a.buildGetIPRangesRequest(ctx)
	if err != nil {
		var localVarReturnValue IPRanges
		return localVarReturnValue, nil, err
	}

	return req.Api.getIPRangesExecute(req)
}

// getIPRangesExecute executes the request.
func (a *IPRangesApi) getIPRangesExecute(r apiGetIPRangesRequest) (IPRanges, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue IPRanges
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.IPRangesApi.GetIPRanges")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
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
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewIPRangesApi Returns NewIPRangesApi.
func NewIPRangesApi(client *common.APIClient) *IPRangesApi {
	return &IPRangesApi{
		Client: client,
	}
}
