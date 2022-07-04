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
)

// KeyManagementApiService KeyManagementApi service.
type KeyManagementApiService service

type apiCreateAPIKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	body       *APIKeyCreateRequest
}

func (a *KeyManagementApiService) buildCreateAPIKeyRequest(ctx _context.Context, body APIKeyCreateRequest) (apiCreateAPIKeyRequest, error) {
	req := apiCreateAPIKeyRequest{
		ApiService: a,
		ctx:        ctx,
		body:       &body,
	}
	return req, nil
}

// CreateAPIKey Create an API key.
// Create an API key.
func (a *KeyManagementApiService) CreateAPIKey(ctx _context.Context, body APIKeyCreateRequest) (APIKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildCreateAPIKeyRequest(ctx, body)
	if err != nil {
		var localVarReturnValue APIKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.createAPIKeyExecute(req)
}

// createAPIKeyExecute executes the request.
func (a *KeyManagementApiService) createAPIKeyExecute(r apiCreateAPIKeyRequest) (APIKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue APIKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.CreateAPIKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/api_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiCreateCurrentUserApplicationKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	body       *ApplicationKeyCreateRequest
}

func (a *KeyManagementApiService) buildCreateCurrentUserApplicationKeyRequest(ctx _context.Context, body ApplicationKeyCreateRequest) (apiCreateCurrentUserApplicationKeyRequest, error) {
	req := apiCreateCurrentUserApplicationKeyRequest{
		ApiService: a,
		ctx:        ctx,
		body:       &body,
	}
	return req, nil
}

// CreateCurrentUserApplicationKey Create an application key for current user.
// Create an application key for current user
func (a *KeyManagementApiService) CreateCurrentUserApplicationKey(ctx _context.Context, body ApplicationKeyCreateRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildCreateCurrentUserApplicationKeyRequest(ctx, body)
	if err != nil {
		var localVarReturnValue ApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.createCurrentUserApplicationKeyExecute(req)
}

// createCurrentUserApplicationKeyExecute executes the request.
func (a *KeyManagementApiService) createCurrentUserApplicationKeyExecute(r apiCreateCurrentUserApplicationKeyRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue ApplicationKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.CreateCurrentUserApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/current_user/application_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiDeleteAPIKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	apiKeyId   string
}

func (a *KeyManagementApiService) buildDeleteAPIKeyRequest(ctx _context.Context, apiKeyId string) (apiDeleteAPIKeyRequest, error) {
	req := apiDeleteAPIKeyRequest{
		ApiService: a,
		ctx:        ctx,
		apiKeyId:   apiKeyId,
	}
	return req, nil
}

// DeleteAPIKey Delete an API key.
// Delete an API key.
func (a *KeyManagementApiService) DeleteAPIKey(ctx _context.Context, apiKeyId string) (*_nethttp.Response, error) {
	req, err := a.buildDeleteAPIKeyRequest(ctx, apiKeyId)
	if err != nil {
		return nil, err
	}

	return req.ApiService.deleteAPIKeyExecute(req)
}

// deleteAPIKeyExecute executes the request.
func (a *KeyManagementApiService) deleteAPIKeyExecute(r apiDeleteAPIKeyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.DeleteAPIKey")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/api_keys/{api_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"api_key_id"+"}", _neturl.PathEscape(parameterToString(r.apiKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type apiDeleteApplicationKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	appKeyId   string
}

func (a *KeyManagementApiService) buildDeleteApplicationKeyRequest(ctx _context.Context, appKeyId string) (apiDeleteApplicationKeyRequest, error) {
	req := apiDeleteApplicationKeyRequest{
		ApiService: a,
		ctx:        ctx,
		appKeyId:   appKeyId,
	}
	return req, nil
}

// DeleteApplicationKey Delete an application key.
// Delete an application key
func (a *KeyManagementApiService) DeleteApplicationKey(ctx _context.Context, appKeyId string) (*_nethttp.Response, error) {
	req, err := a.buildDeleteApplicationKeyRequest(ctx, appKeyId)
	if err != nil {
		return nil, err
	}

	return req.ApiService.deleteApplicationKeyExecute(req)
}

// deleteApplicationKeyExecute executes the request.
func (a *KeyManagementApiService) deleteApplicationKeyExecute(r apiDeleteApplicationKeyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.DeleteApplicationKey")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(parameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type apiDeleteCurrentUserApplicationKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	appKeyId   string
}

func (a *KeyManagementApiService) buildDeleteCurrentUserApplicationKeyRequest(ctx _context.Context, appKeyId string) (apiDeleteCurrentUserApplicationKeyRequest, error) {
	req := apiDeleteCurrentUserApplicationKeyRequest{
		ApiService: a,
		ctx:        ctx,
		appKeyId:   appKeyId,
	}
	return req, nil
}

// DeleteCurrentUserApplicationKey Delete an application key owned by current user.
// Delete an application key owned by current user
func (a *KeyManagementApiService) DeleteCurrentUserApplicationKey(ctx _context.Context, appKeyId string) (*_nethttp.Response, error) {
	req, err := a.buildDeleteCurrentUserApplicationKeyRequest(ctx, appKeyId)
	if err != nil {
		return nil, err
	}

	return req.ApiService.deleteCurrentUserApplicationKeyExecute(req)
}

// deleteCurrentUserApplicationKeyExecute executes the request.
func (a *KeyManagementApiService) deleteCurrentUserApplicationKeyExecute(r apiDeleteCurrentUserApplicationKeyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.DeleteCurrentUserApplicationKey")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/current_user/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(parameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type apiGetAPIKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	apiKeyId   string
	include    *string
}

// GetAPIKeyOptionalParameters holds optional parameters for GetAPIKey.
type GetAPIKeyOptionalParameters struct {
	Include *string
}

// NewGetAPIKeyOptionalParameters creates an empty struct for parameters.
func NewGetAPIKeyOptionalParameters() *GetAPIKeyOptionalParameters {
	this := GetAPIKeyOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *GetAPIKeyOptionalParameters) WithInclude(include string) *GetAPIKeyOptionalParameters {
	r.Include = &include
	return r
}

func (a *KeyManagementApiService) buildGetAPIKeyRequest(ctx _context.Context, apiKeyId string, o ...GetAPIKeyOptionalParameters) (apiGetAPIKeyRequest, error) {
	req := apiGetAPIKeyRequest{
		ApiService: a,
		ctx:        ctx,
		apiKeyId:   apiKeyId,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetAPIKeyOptionalParameters is allowed")
	}

	if o != nil {
		req.include = o[0].Include
	}
	return req, nil
}

// GetAPIKey Get API key.
// Get an API key.
func (a *KeyManagementApiService) GetAPIKey(ctx _context.Context, apiKeyId string, o ...GetAPIKeyOptionalParameters) (APIKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildGetAPIKeyRequest(ctx, apiKeyId, o...)
	if err != nil {
		var localVarReturnValue APIKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getAPIKeyExecute(req)
}

// getAPIKeyExecute executes the request.
func (a *KeyManagementApiService) getAPIKeyExecute(r apiGetAPIKeyRequest) (APIKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue APIKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.GetAPIKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/api_keys/{api_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"api_key_id"+"}", _neturl.PathEscape(parameterToString(r.apiKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetApplicationKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	appKeyId   string
	include    *string
}

// GetApplicationKeyOptionalParameters holds optional parameters for GetApplicationKey.
type GetApplicationKeyOptionalParameters struct {
	Include *string
}

// NewGetApplicationKeyOptionalParameters creates an empty struct for parameters.
func NewGetApplicationKeyOptionalParameters() *GetApplicationKeyOptionalParameters {
	this := GetApplicationKeyOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *GetApplicationKeyOptionalParameters) WithInclude(include string) *GetApplicationKeyOptionalParameters {
	r.Include = &include
	return r
}

func (a *KeyManagementApiService) buildGetApplicationKeyRequest(ctx _context.Context, appKeyId string, o ...GetApplicationKeyOptionalParameters) (apiGetApplicationKeyRequest, error) {
	req := apiGetApplicationKeyRequest{
		ApiService: a,
		ctx:        ctx,
		appKeyId:   appKeyId,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetApplicationKeyOptionalParameters is allowed")
	}

	if o != nil {
		req.include = o[0].Include
	}
	return req, nil
}

// GetApplicationKey Get an application key.
// Get an application key for your org.
func (a *KeyManagementApiService) GetApplicationKey(ctx _context.Context, appKeyId string, o ...GetApplicationKeyOptionalParameters) (ApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildGetApplicationKeyRequest(ctx, appKeyId, o...)
	if err != nil {
		var localVarReturnValue ApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getApplicationKeyExecute(req)
}

// getApplicationKeyExecute executes the request.
func (a *KeyManagementApiService) getApplicationKeyExecute(r apiGetApplicationKeyRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ApplicationKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.GetApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(parameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetCurrentUserApplicationKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	appKeyId   string
}

func (a *KeyManagementApiService) buildGetCurrentUserApplicationKeyRequest(ctx _context.Context, appKeyId string) (apiGetCurrentUserApplicationKeyRequest, error) {
	req := apiGetCurrentUserApplicationKeyRequest{
		ApiService: a,
		ctx:        ctx,
		appKeyId:   appKeyId,
	}
	return req, nil
}

// GetCurrentUserApplicationKey Get one application key owned by current user.
// Get an application key owned by current user
func (a *KeyManagementApiService) GetCurrentUserApplicationKey(ctx _context.Context, appKeyId string) (ApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildGetCurrentUserApplicationKeyRequest(ctx, appKeyId)
	if err != nil {
		var localVarReturnValue ApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getCurrentUserApplicationKeyExecute(req)
}

// getCurrentUserApplicationKeyExecute executes the request.
func (a *KeyManagementApiService) getCurrentUserApplicationKeyExecute(r apiGetCurrentUserApplicationKeyRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ApplicationKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.GetCurrentUserApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/current_user/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(parameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiListAPIKeysRequest struct {
	ctx                   _context.Context
	ApiService            *KeyManagementApiService
	pageSize              *int64
	pageNumber            *int64
	sort                  *APIKeysSort
	filter                *string
	filterCreatedAtStart  *string
	filterCreatedAtEnd    *string
	filterModifiedAtStart *string
	filterModifiedAtEnd   *string
	include               *string
}

// ListAPIKeysOptionalParameters holds optional parameters for ListAPIKeys.
type ListAPIKeysOptionalParameters struct {
	PageSize              *int64
	PageNumber            *int64
	Sort                  *APIKeysSort
	Filter                *string
	FilterCreatedAtStart  *string
	FilterCreatedAtEnd    *string
	FilterModifiedAtStart *string
	FilterModifiedAtEnd   *string
	Include               *string
}

// NewListAPIKeysOptionalParameters creates an empty struct for parameters.
func NewListAPIKeysOptionalParameters() *ListAPIKeysOptionalParameters {
	this := ListAPIKeysOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListAPIKeysOptionalParameters) WithPageSize(pageSize int64) *ListAPIKeysOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListAPIKeysOptionalParameters) WithPageNumber(pageNumber int64) *ListAPIKeysOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListAPIKeysOptionalParameters) WithSort(sort APIKeysSort) *ListAPIKeysOptionalParameters {
	r.Sort = &sort
	return r
}

// WithFilter sets the corresponding parameter name and returns the struct.
func (r *ListAPIKeysOptionalParameters) WithFilter(filter string) *ListAPIKeysOptionalParameters {
	r.Filter = &filter
	return r
}

// WithFilterCreatedAtStart sets the corresponding parameter name and returns the struct.
func (r *ListAPIKeysOptionalParameters) WithFilterCreatedAtStart(filterCreatedAtStart string) *ListAPIKeysOptionalParameters {
	r.FilterCreatedAtStart = &filterCreatedAtStart
	return r
}

// WithFilterCreatedAtEnd sets the corresponding parameter name and returns the struct.
func (r *ListAPIKeysOptionalParameters) WithFilterCreatedAtEnd(filterCreatedAtEnd string) *ListAPIKeysOptionalParameters {
	r.FilterCreatedAtEnd = &filterCreatedAtEnd
	return r
}

// WithFilterModifiedAtStart sets the corresponding parameter name and returns the struct.
func (r *ListAPIKeysOptionalParameters) WithFilterModifiedAtStart(filterModifiedAtStart string) *ListAPIKeysOptionalParameters {
	r.FilterModifiedAtStart = &filterModifiedAtStart
	return r
}

// WithFilterModifiedAtEnd sets the corresponding parameter name and returns the struct.
func (r *ListAPIKeysOptionalParameters) WithFilterModifiedAtEnd(filterModifiedAtEnd string) *ListAPIKeysOptionalParameters {
	r.FilterModifiedAtEnd = &filterModifiedAtEnd
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListAPIKeysOptionalParameters) WithInclude(include string) *ListAPIKeysOptionalParameters {
	r.Include = &include
	return r
}

func (a *KeyManagementApiService) buildListAPIKeysRequest(ctx _context.Context, o ...ListAPIKeysOptionalParameters) (apiListAPIKeysRequest, error) {
	req := apiListAPIKeysRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type ListAPIKeysOptionalParameters is allowed")
	}

	if o != nil {
		req.pageSize = o[0].PageSize
		req.pageNumber = o[0].PageNumber
		req.sort = o[0].Sort
		req.filter = o[0].Filter
		req.filterCreatedAtStart = o[0].FilterCreatedAtStart
		req.filterCreatedAtEnd = o[0].FilterCreatedAtEnd
		req.filterModifiedAtStart = o[0].FilterModifiedAtStart
		req.filterModifiedAtEnd = o[0].FilterModifiedAtEnd
		req.include = o[0].Include
	}
	return req, nil
}

// ListAPIKeys Get all API keys.
// List all API keys available for your account.
func (a *KeyManagementApiService) ListAPIKeys(ctx _context.Context, o ...ListAPIKeysOptionalParameters) (APIKeysResponse, *_nethttp.Response, error) {
	req, err := a.buildListAPIKeysRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue APIKeysResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listAPIKeysExecute(req)
}

// listAPIKeysExecute executes the request.
func (a *KeyManagementApiService) listAPIKeysExecute(r apiListAPIKeysRequest) (APIKeysResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue APIKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.ListAPIKeys")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/api_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.filterCreatedAtStart != nil {
		localVarQueryParams.Add("filter[created_at][start]", parameterToString(*r.filterCreatedAtStart, ""))
	}
	if r.filterCreatedAtEnd != nil {
		localVarQueryParams.Add("filter[created_at][end]", parameterToString(*r.filterCreatedAtEnd, ""))
	}
	if r.filterModifiedAtStart != nil {
		localVarQueryParams.Add("filter[modified_at][start]", parameterToString(*r.filterModifiedAtStart, ""))
	}
	if r.filterModifiedAtEnd != nil {
		localVarQueryParams.Add("filter[modified_at][end]", parameterToString(*r.filterModifiedAtEnd, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiListApplicationKeysRequest struct {
	ctx                  _context.Context
	ApiService           *KeyManagementApiService
	pageSize             *int64
	pageNumber           *int64
	sort                 *ApplicationKeysSort
	filter               *string
	filterCreatedAtStart *string
	filterCreatedAtEnd   *string
}

// ListApplicationKeysOptionalParameters holds optional parameters for ListApplicationKeys.
type ListApplicationKeysOptionalParameters struct {
	PageSize             *int64
	PageNumber           *int64
	Sort                 *ApplicationKeysSort
	Filter               *string
	FilterCreatedAtStart *string
	FilterCreatedAtEnd   *string
}

// NewListApplicationKeysOptionalParameters creates an empty struct for parameters.
func NewListApplicationKeysOptionalParameters() *ListApplicationKeysOptionalParameters {
	this := ListApplicationKeysOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListApplicationKeysOptionalParameters) WithPageSize(pageSize int64) *ListApplicationKeysOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListApplicationKeysOptionalParameters) WithPageNumber(pageNumber int64) *ListApplicationKeysOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListApplicationKeysOptionalParameters) WithSort(sort ApplicationKeysSort) *ListApplicationKeysOptionalParameters {
	r.Sort = &sort
	return r
}

// WithFilter sets the corresponding parameter name and returns the struct.
func (r *ListApplicationKeysOptionalParameters) WithFilter(filter string) *ListApplicationKeysOptionalParameters {
	r.Filter = &filter
	return r
}

// WithFilterCreatedAtStart sets the corresponding parameter name and returns the struct.
func (r *ListApplicationKeysOptionalParameters) WithFilterCreatedAtStart(filterCreatedAtStart string) *ListApplicationKeysOptionalParameters {
	r.FilterCreatedAtStart = &filterCreatedAtStart
	return r
}

// WithFilterCreatedAtEnd sets the corresponding parameter name and returns the struct.
func (r *ListApplicationKeysOptionalParameters) WithFilterCreatedAtEnd(filterCreatedAtEnd string) *ListApplicationKeysOptionalParameters {
	r.FilterCreatedAtEnd = &filterCreatedAtEnd
	return r
}

func (a *KeyManagementApiService) buildListApplicationKeysRequest(ctx _context.Context, o ...ListApplicationKeysOptionalParameters) (apiListApplicationKeysRequest, error) {
	req := apiListApplicationKeysRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type ListApplicationKeysOptionalParameters is allowed")
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

// ListApplicationKeys Get all application keys.
// List all application keys available for your org
func (a *KeyManagementApiService) ListApplicationKeys(ctx _context.Context, o ...ListApplicationKeysOptionalParameters) (ListApplicationKeysResponse, *_nethttp.Response, error) {
	req, err := a.buildListApplicationKeysRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue ListApplicationKeysResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listApplicationKeysExecute(req)
}

// listApplicationKeysExecute executes the request.
func (a *KeyManagementApiService) listApplicationKeysExecute(r apiListApplicationKeysRequest) (ListApplicationKeysResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListApplicationKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.ListApplicationKeys")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/application_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.filterCreatedAtStart != nil {
		localVarQueryParams.Add("filter[created_at][start]", parameterToString(*r.filterCreatedAtStart, ""))
	}
	if r.filterCreatedAtEnd != nil {
		localVarQueryParams.Add("filter[created_at][end]", parameterToString(*r.filterCreatedAtEnd, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiListCurrentUserApplicationKeysRequest struct {
	ctx                  _context.Context
	ApiService           *KeyManagementApiService
	pageSize             *int64
	pageNumber           *int64
	sort                 *ApplicationKeysSort
	filter               *string
	filterCreatedAtStart *string
	filterCreatedAtEnd   *string
}

// ListCurrentUserApplicationKeysOptionalParameters holds optional parameters for ListCurrentUserApplicationKeys.
type ListCurrentUserApplicationKeysOptionalParameters struct {
	PageSize             *int64
	PageNumber           *int64
	Sort                 *ApplicationKeysSort
	Filter               *string
	FilterCreatedAtStart *string
	FilterCreatedAtEnd   *string
}

// NewListCurrentUserApplicationKeysOptionalParameters creates an empty struct for parameters.
func NewListCurrentUserApplicationKeysOptionalParameters() *ListCurrentUserApplicationKeysOptionalParameters {
	this := ListCurrentUserApplicationKeysOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListCurrentUserApplicationKeysOptionalParameters) WithPageSize(pageSize int64) *ListCurrentUserApplicationKeysOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListCurrentUserApplicationKeysOptionalParameters) WithPageNumber(pageNumber int64) *ListCurrentUserApplicationKeysOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListCurrentUserApplicationKeysOptionalParameters) WithSort(sort ApplicationKeysSort) *ListCurrentUserApplicationKeysOptionalParameters {
	r.Sort = &sort
	return r
}

// WithFilter sets the corresponding parameter name and returns the struct.
func (r *ListCurrentUserApplicationKeysOptionalParameters) WithFilter(filter string) *ListCurrentUserApplicationKeysOptionalParameters {
	r.Filter = &filter
	return r
}

// WithFilterCreatedAtStart sets the corresponding parameter name and returns the struct.
func (r *ListCurrentUserApplicationKeysOptionalParameters) WithFilterCreatedAtStart(filterCreatedAtStart string) *ListCurrentUserApplicationKeysOptionalParameters {
	r.FilterCreatedAtStart = &filterCreatedAtStart
	return r
}

// WithFilterCreatedAtEnd sets the corresponding parameter name and returns the struct.
func (r *ListCurrentUserApplicationKeysOptionalParameters) WithFilterCreatedAtEnd(filterCreatedAtEnd string) *ListCurrentUserApplicationKeysOptionalParameters {
	r.FilterCreatedAtEnd = &filterCreatedAtEnd
	return r
}

func (a *KeyManagementApiService) buildListCurrentUserApplicationKeysRequest(ctx _context.Context, o ...ListCurrentUserApplicationKeysOptionalParameters) (apiListCurrentUserApplicationKeysRequest, error) {
	req := apiListCurrentUserApplicationKeysRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type ListCurrentUserApplicationKeysOptionalParameters is allowed")
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

// ListCurrentUserApplicationKeys Get all application keys owned by current user.
// List all application keys available for current user
func (a *KeyManagementApiService) ListCurrentUserApplicationKeys(ctx _context.Context, o ...ListCurrentUserApplicationKeysOptionalParameters) (ListApplicationKeysResponse, *_nethttp.Response, error) {
	req, err := a.buildListCurrentUserApplicationKeysRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue ListApplicationKeysResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listCurrentUserApplicationKeysExecute(req)
}

// listCurrentUserApplicationKeysExecute executes the request.
func (a *KeyManagementApiService) listCurrentUserApplicationKeysExecute(r apiListCurrentUserApplicationKeysRequest) (ListApplicationKeysResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListApplicationKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.ListCurrentUserApplicationKeys")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/current_user/application_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.filterCreatedAtStart != nil {
		localVarQueryParams.Add("filter[created_at][start]", parameterToString(*r.filterCreatedAtStart, ""))
	}
	if r.filterCreatedAtEnd != nil {
		localVarQueryParams.Add("filter[created_at][end]", parameterToString(*r.filterCreatedAtEnd, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiUpdateAPIKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	apiKeyId   string
	body       *APIKeyUpdateRequest
}

func (a *KeyManagementApiService) buildUpdateAPIKeyRequest(ctx _context.Context, apiKeyId string, body APIKeyUpdateRequest) (apiUpdateAPIKeyRequest, error) {
	req := apiUpdateAPIKeyRequest{
		ApiService: a,
		ctx:        ctx,
		apiKeyId:   apiKeyId,
		body:       &body,
	}
	return req, nil
}

// UpdateAPIKey Edit an API key.
// Update an API key.
func (a *KeyManagementApiService) UpdateAPIKey(ctx _context.Context, apiKeyId string, body APIKeyUpdateRequest) (APIKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildUpdateAPIKeyRequest(ctx, apiKeyId, body)
	if err != nil {
		var localVarReturnValue APIKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.updateAPIKeyExecute(req)
}

// updateAPIKeyExecute executes the request.
func (a *KeyManagementApiService) updateAPIKeyExecute(r apiUpdateAPIKeyRequest) (APIKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue APIKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.UpdateAPIKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/api_keys/{api_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"api_key_id"+"}", _neturl.PathEscape(parameterToString(r.apiKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiUpdateApplicationKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	appKeyId   string
	body       *ApplicationKeyUpdateRequest
}

func (a *KeyManagementApiService) buildUpdateApplicationKeyRequest(ctx _context.Context, appKeyId string, body ApplicationKeyUpdateRequest) (apiUpdateApplicationKeyRequest, error) {
	req := apiUpdateApplicationKeyRequest{
		ApiService: a,
		ctx:        ctx,
		appKeyId:   appKeyId,
		body:       &body,
	}
	return req, nil
}

// UpdateApplicationKey Edit an application key.
// Edit an application key
func (a *KeyManagementApiService) UpdateApplicationKey(ctx _context.Context, appKeyId string, body ApplicationKeyUpdateRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildUpdateApplicationKeyRequest(ctx, appKeyId, body)
	if err != nil {
		var localVarReturnValue ApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.updateApplicationKeyExecute(req)
}

// updateApplicationKeyExecute executes the request.
func (a *KeyManagementApiService) updateApplicationKeyExecute(r apiUpdateApplicationKeyRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue ApplicationKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.UpdateApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(parameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiUpdateCurrentUserApplicationKeyRequest struct {
	ctx        _context.Context
	ApiService *KeyManagementApiService
	appKeyId   string
	body       *ApplicationKeyUpdateRequest
}

func (a *KeyManagementApiService) buildUpdateCurrentUserApplicationKeyRequest(ctx _context.Context, appKeyId string, body ApplicationKeyUpdateRequest) (apiUpdateCurrentUserApplicationKeyRequest, error) {
	req := apiUpdateCurrentUserApplicationKeyRequest{
		ApiService: a,
		ctx:        ctx,
		appKeyId:   appKeyId,
		body:       &body,
	}
	return req, nil
}

// UpdateCurrentUserApplicationKey Edit an application key owned by current user.
// Edit an application key owned by current user
func (a *KeyManagementApiService) UpdateCurrentUserApplicationKey(ctx _context.Context, appKeyId string, body ApplicationKeyUpdateRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildUpdateCurrentUserApplicationKeyRequest(ctx, appKeyId, body)
	if err != nil {
		var localVarReturnValue ApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.updateCurrentUserApplicationKeyExecute(req)
}

// updateCurrentUserApplicationKeyExecute executes the request.
func (a *KeyManagementApiService) updateCurrentUserApplicationKeyExecute(r apiUpdateCurrentUserApplicationKeyRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue ApplicationKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagementApiService.UpdateCurrentUserApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/current_user/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(parameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
