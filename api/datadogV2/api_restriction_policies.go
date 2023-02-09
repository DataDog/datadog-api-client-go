// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionPoliciesApi service type
type RestrictionPoliciesApi datadog.Service

type apiGetRestrictionPolicyRequest struct {
	ctx        _context.Context
	resourceId string
}

func (a *RestrictionPoliciesApi) buildGetRestrictionPolicyRequest(ctx _context.Context, resourceId string) (apiGetRestrictionPolicyRequest, error) {
	req := apiGetRestrictionPolicyRequest{
		ctx:        ctx,
		resourceId: resourceId,
	}
	return req, nil
}

// GetRestrictionPolicy Get a restriction policy.
// Retrieves the restriction policy associated with a specified resource.
func (a *RestrictionPoliciesApi) GetRestrictionPolicy(ctx _context.Context, resourceId string) (GetRestrictionPolicyResponse, *_nethttp.Response, error) {
	req, err := a.buildGetRestrictionPolicyRequest(ctx, resourceId)
	if err != nil {
		var localVarReturnValue GetRestrictionPolicyResponse
		return localVarReturnValue, nil, err
	}

	return a.getRestrictionPolicyExecute(req)
}

// getRestrictionPolicyExecute executes the request.
func (a *RestrictionPoliciesApi) getRestrictionPolicyExecute(r apiGetRestrictionPolicyRequest) (GetRestrictionPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue GetRestrictionPolicyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.RestrictionPoliciesApi.GetRestrictionPolicy")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/restriction_policy/{resource_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"resource_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.resourceId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
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

// NewRestrictionPoliciesApi Returns NewRestrictionPoliciesApi.
func NewRestrictionPoliciesApi(client *datadog.APIClient) *RestrictionPoliciesApi {
	return &RestrictionPoliciesApi{
		Client: client,
	}
}
