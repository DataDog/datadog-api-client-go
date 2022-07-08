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

type TagsApiService common.Service

type apiCreateHostTagsRequest struct {
	ctx        _context.Context
	ApiService *TagsApiService
	hostName   string
	body       *HostTags
	source     *string
}

// CreateHostTagsOptionalParameters holds optional parameters for CreateHostTags.
type CreateHostTagsOptionalParameters struct {
	Source *string
}

// NewCreateHostTagsOptionalParameters creates an empty struct for parameters.
func NewCreateHostTagsOptionalParameters() *CreateHostTagsOptionalParameters {
	this := CreateHostTagsOptionalParameters{}
	return &this
}

// WithSource sets the corresponding parameter name and returns the struct.
func (r *CreateHostTagsOptionalParameters) WithSource(source string) *CreateHostTagsOptionalParameters {
	r.Source = &source
	return r
}

func (a *TagsApiService) buildCreateHostTagsRequest(ctx _context.Context, hostName string, body HostTags, o ...CreateHostTagsOptionalParameters) (apiCreateHostTagsRequest, error) {
	req := apiCreateHostTagsRequest{
		ApiService: a,
		ctx:        ctx,
		hostName:   hostName,
		body:       &body,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type CreateHostTagsOptionalParameters is allowed")
	}

	if o != nil {
		req.source = o[0].Source
	}
	return req, nil
}

// CreateHostTags Add tags to a host.
// This endpoint allows you to add new tags to a host,
// optionally specifying where these tags come from.
func (a *TagsApiService) CreateHostTags(ctx _context.Context, hostName string, body HostTags, o ...CreateHostTagsOptionalParameters) (HostTags, *_nethttp.Response, error) {
	req, err := a.buildCreateHostTagsRequest(ctx, hostName, body, o...)
	if err != nil {
		var localVarReturnValue HostTags
		return localVarReturnValue, nil, err
	}

	return req.ApiService.createHostTagsExecute(req)
}

// createHostTagsExecute executes the request.
func (a *TagsApiService) createHostTagsExecute(r apiCreateHostTagsRequest) (HostTags, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue HostTags
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.TagsApiService.CreateHostTags")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/hosts/{host_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"host_name"+"}", _neturl.PathEscape(common.ParameterToString(r.hostName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, common.ReportError("body is required and must be specified")
	}
	if r.source != nil {
		localVarQueryParams.Add("source", common.ParameterToString(*r.source, ""))
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := common.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

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

type apiDeleteHostTagsRequest struct {
	ctx        _context.Context
	ApiService *TagsApiService
	hostName   string
	source     *string
}

// DeleteHostTagsOptionalParameters holds optional parameters for DeleteHostTags.
type DeleteHostTagsOptionalParameters struct {
	Source *string
}

// NewDeleteHostTagsOptionalParameters creates an empty struct for parameters.
func NewDeleteHostTagsOptionalParameters() *DeleteHostTagsOptionalParameters {
	this := DeleteHostTagsOptionalParameters{}
	return &this
}

// WithSource sets the corresponding parameter name and returns the struct.
func (r *DeleteHostTagsOptionalParameters) WithSource(source string) *DeleteHostTagsOptionalParameters {
	r.Source = &source
	return r
}

func (a *TagsApiService) buildDeleteHostTagsRequest(ctx _context.Context, hostName string, o ...DeleteHostTagsOptionalParameters) (apiDeleteHostTagsRequest, error) {
	req := apiDeleteHostTagsRequest{
		ApiService: a,
		ctx:        ctx,
		hostName:   hostName,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type DeleteHostTagsOptionalParameters is allowed")
	}

	if o != nil {
		req.source = o[0].Source
	}
	return req, nil
}

// DeleteHostTags Remove host tags.
// This endpoint allows you to remove all user-assigned tags
// for a single host.
func (a *TagsApiService) DeleteHostTags(ctx _context.Context, hostName string, o ...DeleteHostTagsOptionalParameters) (*_nethttp.Response, error) {
	req, err := a.buildDeleteHostTagsRequest(ctx, hostName, o...)
	if err != nil {
		return nil, err
	}

	return req.ApiService.deleteHostTagsExecute(req)
}

// deleteHostTagsExecute executes the request.
func (a *TagsApiService) deleteHostTagsExecute(r apiDeleteHostTagsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.TagsApiService.DeleteHostTags")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/hosts/{host_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"host_name"+"}", _neturl.PathEscape(common.ParameterToString(r.hostName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.source != nil {
		localVarQueryParams.Add("source", common.ParameterToString(*r.source, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type apiGetHostTagsRequest struct {
	ctx        _context.Context
	ApiService *TagsApiService
	hostName   string
	source     *string
}

// GetHostTagsOptionalParameters holds optional parameters for GetHostTags.
type GetHostTagsOptionalParameters struct {
	Source *string
}

// NewGetHostTagsOptionalParameters creates an empty struct for parameters.
func NewGetHostTagsOptionalParameters() *GetHostTagsOptionalParameters {
	this := GetHostTagsOptionalParameters{}
	return &this
}

// WithSource sets the corresponding parameter name and returns the struct.
func (r *GetHostTagsOptionalParameters) WithSource(source string) *GetHostTagsOptionalParameters {
	r.Source = &source
	return r
}

func (a *TagsApiService) buildGetHostTagsRequest(ctx _context.Context, hostName string, o ...GetHostTagsOptionalParameters) (apiGetHostTagsRequest, error) {
	req := apiGetHostTagsRequest{
		ApiService: a,
		ctx:        ctx,
		hostName:   hostName,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type GetHostTagsOptionalParameters is allowed")
	}

	if o != nil {
		req.source = o[0].Source
	}
	return req, nil
}

// GetHostTags Get host tags.
// Return the list of tags that apply to a given host.
func (a *TagsApiService) GetHostTags(ctx _context.Context, hostName string, o ...GetHostTagsOptionalParameters) (HostTags, *_nethttp.Response, error) {
	req, err := a.buildGetHostTagsRequest(ctx, hostName, o...)
	if err != nil {
		var localVarReturnValue HostTags
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getHostTagsExecute(req)
}

// getHostTagsExecute executes the request.
func (a *TagsApiService) getHostTagsExecute(r apiGetHostTagsRequest) (HostTags, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue HostTags
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.TagsApiService.GetHostTags")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/hosts/{host_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"host_name"+"}", _neturl.PathEscape(common.ParameterToString(r.hostName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.source != nil {
		localVarQueryParams.Add("source", common.ParameterToString(*r.source, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type apiListHostTagsRequest struct {
	ctx        _context.Context
	ApiService *TagsApiService
	source     *string
}

// ListHostTagsOptionalParameters holds optional parameters for ListHostTags.
type ListHostTagsOptionalParameters struct {
	Source *string
}

// NewListHostTagsOptionalParameters creates an empty struct for parameters.
func NewListHostTagsOptionalParameters() *ListHostTagsOptionalParameters {
	this := ListHostTagsOptionalParameters{}
	return &this
}

// WithSource sets the corresponding parameter name and returns the struct.
func (r *ListHostTagsOptionalParameters) WithSource(source string) *ListHostTagsOptionalParameters {
	r.Source = &source
	return r
}

func (a *TagsApiService) buildListHostTagsRequest(ctx _context.Context, o ...ListHostTagsOptionalParameters) (apiListHostTagsRequest, error) {
	req := apiListHostTagsRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type ListHostTagsOptionalParameters is allowed")
	}

	if o != nil {
		req.source = o[0].Source
	}
	return req, nil
}

// ListHostTags Get Tags.
// Return a mapping of tags to hosts for your whole infrastructure.
func (a *TagsApiService) ListHostTags(ctx _context.Context, o ...ListHostTagsOptionalParameters) (TagToHosts, *_nethttp.Response, error) {
	req, err := a.buildListHostTagsRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue TagToHosts
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listHostTagsExecute(req)
}

// listHostTagsExecute executes the request.
func (a *TagsApiService) listHostTagsExecute(r apiListHostTagsRequest) (TagToHosts, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue TagToHosts
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.TagsApiService.ListHostTags")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/hosts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.source != nil {
		localVarQueryParams.Add("source", common.ParameterToString(*r.source, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type apiUpdateHostTagsRequest struct {
	ctx        _context.Context
	ApiService *TagsApiService
	hostName   string
	body       *HostTags
	source     *string
}

// UpdateHostTagsOptionalParameters holds optional parameters for UpdateHostTags.
type UpdateHostTagsOptionalParameters struct {
	Source *string
}

// NewUpdateHostTagsOptionalParameters creates an empty struct for parameters.
func NewUpdateHostTagsOptionalParameters() *UpdateHostTagsOptionalParameters {
	this := UpdateHostTagsOptionalParameters{}
	return &this
}

// WithSource sets the corresponding parameter name and returns the struct.
func (r *UpdateHostTagsOptionalParameters) WithSource(source string) *UpdateHostTagsOptionalParameters {
	r.Source = &source
	return r
}

func (a *TagsApiService) buildUpdateHostTagsRequest(ctx _context.Context, hostName string, body HostTags, o ...UpdateHostTagsOptionalParameters) (apiUpdateHostTagsRequest, error) {
	req := apiUpdateHostTagsRequest{
		ApiService: a,
		ctx:        ctx,
		hostName:   hostName,
		body:       &body,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type UpdateHostTagsOptionalParameters is allowed")
	}

	if o != nil {
		req.source = o[0].Source
	}
	return req, nil
}

// UpdateHostTags Update host tags.
// This endpoint allows you to update/replace all tags in
// an integration source with those supplied in the request.
func (a *TagsApiService) UpdateHostTags(ctx _context.Context, hostName string, body HostTags, o ...UpdateHostTagsOptionalParameters) (HostTags, *_nethttp.Response, error) {
	req, err := a.buildUpdateHostTagsRequest(ctx, hostName, body, o...)
	if err != nil {
		var localVarReturnValue HostTags
		return localVarReturnValue, nil, err
	}

	return req.ApiService.updateHostTagsExecute(req)
}

// updateHostTagsExecute executes the request.
func (a *TagsApiService) updateHostTagsExecute(r apiUpdateHostTagsRequest) (HostTags, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue HostTags
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.TagsApiService.UpdateHostTags")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/hosts/{host_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"host_name"+"}", _neturl.PathEscape(common.ParameterToString(r.hostName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, common.ReportError("body is required and must be specified")
	}
	if r.source != nil {
		localVarQueryParams.Add("source", common.ParameterToString(*r.source, ""))
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := common.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

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

// TagsApi Returns new TagsApi service.
func TagsApi(client *common.APIClient) *TagsApiService {
	return &TagsApiService{
		Client: client,
	}
}
