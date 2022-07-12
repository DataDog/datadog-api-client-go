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
	"strings"

	"github.com/DataDog/datadog-api-client-go/api/common"
)

// ServiceAccountsApi service type
type ServiceAccountsApi common.Service

type apiCreateServiceAccountApplicationKeyRequest struct {
	ctx              _context.Context
	Api              *ServiceAccountsApi
	serviceAccountId string
	body             *ApplicationKeyCreateRequest
}

func (a *ServiceAccountsApi) buildCreateServiceAccountApplicationKeyRequest(ctx _context.Context, serviceAccountId string, body ApplicationKeyCreateRequest) (apiCreateServiceAccountApplicationKeyRequest, error) {
	req := apiCreateServiceAccountApplicationKeyRequest{
		Api:              a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
		body:             &body,
	}
	return req, nil
}

// CreateServiceAccountApplicationKey Create an application key for this service account.
// Create an application key for this service account.
func (a *ServiceAccountsApi) CreateServiceAccountApplicationKey(ctx _context.Context, serviceAccountId string, body ApplicationKeyCreateRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildCreateServiceAccountApplicationKeyRequest(ctx, serviceAccountId, body)
	if err != nil {
		var localVarReturnValue ApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.Api.createServiceAccountApplicationKeyExecute(req)
}

// createServiceAccountApplicationKeyExecute executes the request.
func (a *ServiceAccountsApi) createServiceAccountApplicationKeyExecute(r apiCreateServiceAccountApplicationKeyRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue ApplicationKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.ServiceAccountsApi.CreateServiceAccountApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(common.ParameterToString(r.serviceAccountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, common.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type apiDeleteServiceAccountApplicationKeyRequest struct {
	ctx              _context.Context
	Api              *ServiceAccountsApi
	serviceAccountId string
	appKeyId         string
}

func (a *ServiceAccountsApi) buildDeleteServiceAccountApplicationKeyRequest(ctx _context.Context, serviceAccountId string, appKeyId string) (apiDeleteServiceAccountApplicationKeyRequest, error) {
	req := apiDeleteServiceAccountApplicationKeyRequest{
		Api:              a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
		appKeyId:         appKeyId,
	}
	return req, nil
}

// DeleteServiceAccountApplicationKey Delete an application key for this service account.
// Delete an application key owned by this service account.
func (a *ServiceAccountsApi) DeleteServiceAccountApplicationKey(ctx _context.Context, serviceAccountId string, appKeyId string) (*_nethttp.Response, error) {
	req, err := a.buildDeleteServiceAccountApplicationKeyRequest(ctx, serviceAccountId, appKeyId)
	if err != nil {
		return nil, err
	}

	return req.Api.deleteServiceAccountApplicationKeyExecute(req)
}

// deleteServiceAccountApplicationKeyExecute executes the request.
func (a *ServiceAccountsApi) deleteServiceAccountApplicationKeyExecute(r apiDeleteServiceAccountApplicationKeyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.ServiceAccountsApi.DeleteServiceAccountApplicationKey")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(common.ParameterToString(r.serviceAccountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(common.ParameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type apiGetServiceAccountApplicationKeyRequest struct {
	ctx              _context.Context
	Api              *ServiceAccountsApi
	serviceAccountId string
	appKeyId         string
}

func (a *ServiceAccountsApi) buildGetServiceAccountApplicationKeyRequest(ctx _context.Context, serviceAccountId string, appKeyId string) (apiGetServiceAccountApplicationKeyRequest, error) {
	req := apiGetServiceAccountApplicationKeyRequest{
		Api:              a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
		appKeyId:         appKeyId,
	}
	return req, nil
}

// GetServiceAccountApplicationKey Get one application key for this service account.
// Get an application key owned by this service account.
func (a *ServiceAccountsApi) GetServiceAccountApplicationKey(ctx _context.Context, serviceAccountId string, appKeyId string) (PartialApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildGetServiceAccountApplicationKeyRequest(ctx, serviceAccountId, appKeyId)
	if err != nil {
		var localVarReturnValue PartialApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.Api.getServiceAccountApplicationKeyExecute(req)
}

// getServiceAccountApplicationKeyExecute executes the request.
func (a *ServiceAccountsApi) getServiceAccountApplicationKeyExecute(r apiGetServiceAccountApplicationKeyRequest) (PartialApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PartialApplicationKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.ServiceAccountsApi.GetServiceAccountApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(common.ParameterToString(r.serviceAccountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(common.ParameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type apiListServiceAccountApplicationKeysRequest struct {
	ctx                  _context.Context
	Api                  *ServiceAccountsApi
	serviceAccountId     string
	pageSize             *int64
	pageNumber           *int64
	sort                 *ApplicationKeysSort
	filter               *string
	filterCreatedAtStart *string
	filterCreatedAtEnd   *string
}

// ListServiceAccountApplicationKeysOptionalParameters holds optional parameters for ListServiceAccountApplicationKeys.
type ListServiceAccountApplicationKeysOptionalParameters struct {
	PageSize             *int64
	PageNumber           *int64
	Sort                 *ApplicationKeysSort
	Filter               *string
	FilterCreatedAtStart *string
	FilterCreatedAtEnd   *string
}

// NewListServiceAccountApplicationKeysOptionalParameters creates an empty struct for parameters.
func NewListServiceAccountApplicationKeysOptionalParameters() *ListServiceAccountApplicationKeysOptionalParameters {
	this := ListServiceAccountApplicationKeysOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithPageSize(pageSize int64) *ListServiceAccountApplicationKeysOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithPageNumber(pageNumber int64) *ListServiceAccountApplicationKeysOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithSort(sort ApplicationKeysSort) *ListServiceAccountApplicationKeysOptionalParameters {
	r.Sort = &sort
	return r
}

// WithFilter sets the corresponding parameter name and returns the struct.
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithFilter(filter string) *ListServiceAccountApplicationKeysOptionalParameters {
	r.Filter = &filter
	return r
}

// WithFilterCreatedAtStart sets the corresponding parameter name and returns the struct.
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithFilterCreatedAtStart(filterCreatedAtStart string) *ListServiceAccountApplicationKeysOptionalParameters {
	r.FilterCreatedAtStart = &filterCreatedAtStart
	return r
}

// WithFilterCreatedAtEnd sets the corresponding parameter name and returns the struct.
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithFilterCreatedAtEnd(filterCreatedAtEnd string) *ListServiceAccountApplicationKeysOptionalParameters {
	r.FilterCreatedAtEnd = &filterCreatedAtEnd
	return r
}

func (a *ServiceAccountsApi) buildListServiceAccountApplicationKeysRequest(ctx _context.Context, serviceAccountId string, o ...ListServiceAccountApplicationKeysOptionalParameters) (apiListServiceAccountApplicationKeysRequest, error) {
	req := apiListServiceAccountApplicationKeysRequest{
		Api:              a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type ListServiceAccountApplicationKeysOptionalParameters is allowed")
	}

	if o != nil {
		req.pageSize = o[0].PageSize
		req.pageNumber = o[0].PageNumber
		req.sort = o[0].Sort
		req.filter = o[0].Filter
		req.filterCreatedAtStart = o[0].FilterCreatedAtStart
		req.filterCreatedAtEnd = o[0].FilterCreatedAtEnd
	}
	return req, nil
}

// ListServiceAccountApplicationKeys List application keys for this service account.
// List all application keys available for this service account.
func (a *ServiceAccountsApi) ListServiceAccountApplicationKeys(ctx _context.Context, serviceAccountId string, o ...ListServiceAccountApplicationKeysOptionalParameters) (ListApplicationKeysResponse, *_nethttp.Response, error) {
	req, err := a.buildListServiceAccountApplicationKeysRequest(ctx, serviceAccountId, o...)
	if err != nil {
		var localVarReturnValue ListApplicationKeysResponse
		return localVarReturnValue, nil, err
	}

	return req.Api.listServiceAccountApplicationKeysExecute(req)
}

// listServiceAccountApplicationKeysExecute executes the request.
func (a *ServiceAccountsApi) listServiceAccountApplicationKeysExecute(r apiListServiceAccountApplicationKeysRequest) (ListApplicationKeysResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListApplicationKeysResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.ServiceAccountsApi.ListServiceAccountApplicationKeys")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(common.ParameterToString(r.serviceAccountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", common.ParameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", common.ParameterToString(*r.pageNumber, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", common.ParameterToString(*r.sort, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", common.ParameterToString(*r.filter, ""))
	}
	if r.filterCreatedAtStart != nil {
		localVarQueryParams.Add("filter[created_at][start]", common.ParameterToString(*r.filterCreatedAtStart, ""))
	}
	if r.filterCreatedAtEnd != nil {
		localVarQueryParams.Add("filter[created_at][end]", common.ParameterToString(*r.filterCreatedAtEnd, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type apiUpdateServiceAccountApplicationKeyRequest struct {
	ctx              _context.Context
	Api              *ServiceAccountsApi
	serviceAccountId string
	appKeyId         string
	body             *ApplicationKeyUpdateRequest
}

func (a *ServiceAccountsApi) buildUpdateServiceAccountApplicationKeyRequest(ctx _context.Context, serviceAccountId string, appKeyId string, body ApplicationKeyUpdateRequest) (apiUpdateServiceAccountApplicationKeyRequest, error) {
	req := apiUpdateServiceAccountApplicationKeyRequest{
		Api:              a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
		appKeyId:         appKeyId,
		body:             &body,
	}
	return req, nil
}

// UpdateServiceAccountApplicationKey Edit an application key for this service account.
// Edit an application key owned by this service account.
func (a *ServiceAccountsApi) UpdateServiceAccountApplicationKey(ctx _context.Context, serviceAccountId string, appKeyId string, body ApplicationKeyUpdateRequest) (PartialApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildUpdateServiceAccountApplicationKeyRequest(ctx, serviceAccountId, appKeyId, body)
	if err != nil {
		var localVarReturnValue PartialApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.Api.updateServiceAccountApplicationKeyExecute(req)
}

// updateServiceAccountApplicationKeyExecute executes the request.
func (a *ServiceAccountsApi) updateServiceAccountApplicationKeyExecute(r apiUpdateServiceAccountApplicationKeyRequest) (PartialApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue PartialApplicationKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.ServiceAccountsApi.UpdateServiceAccountApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(common.ParameterToString(r.serviceAccountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(common.ParameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, common.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

// NewServiceAccountsApi Returns NewServiceAccountsApi.
func NewServiceAccountsApi(client *common.APIClient) *ServiceAccountsApi {
	return &ServiceAccountsApi{
		Client: client,
	}
}
