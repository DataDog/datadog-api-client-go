/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	_context "context"
	_fmt "fmt"
	_ioutil "io/ioutil"
	_log "log"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ServiceLevelObjectiveCorrectionsApiService ServiceLevelObjectiveCorrectionsApi service
type ServiceLevelObjectiveCorrectionsApiService service

type ApiCreateSLOCorrectionRequest struct {
	ctx        _context.Context
	ApiService *ServiceLevelObjectiveCorrectionsApiService
	body       *SLOCorrectionCreateRequest
}

func (r ApiCreateSLOCorrectionRequest) Body(body SLOCorrectionCreateRequest) ApiCreateSLOCorrectionRequest {
	r.body = &body
	return r
}

func (r ApiCreateSLOCorrectionRequest) Execute() (SLOCorrectionResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateSLOCorrectionExecute(r)
}

/*
 * CreateSLOCorrection Create an SLO correction
 * Create an SLO Correction
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateSLOCorrectionRequest
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) CreateSLOCorrection(ctx _context.Context) ApiCreateSLOCorrectionRequest {
	return ApiCreateSLOCorrectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SLOCorrectionResponse
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) CreateSLOCorrectionExecute(r ApiCreateSLOCorrectionRequest) (SLOCorrectionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SLOCorrectionResponse
	)

	operationId := "CreateSLOCorrection"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceLevelObjectiveCorrectionsApiService.CreateSLOCorrection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/slo/correction"

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

	// Set Operation-ID header for telemetry
	localVarHeaderParams["DD-OPERATION-ID"] = "CreateSLOCorrection"

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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
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

type ApiDeleteSLOCorrectionRequest struct {
	ctx             _context.Context
	ApiService      *ServiceLevelObjectiveCorrectionsApiService
	sloCorrectionId string
}

func (r ApiDeleteSLOCorrectionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteSLOCorrectionExecute(r)
}

/*
 * DeleteSLOCorrection Delete an SLO Correction
 * Permanently delete the specified SLO Correction object
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sloCorrectionId The ID of the SLO correction object
 * @return ApiDeleteSLOCorrectionRequest
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) DeleteSLOCorrection(ctx _context.Context, sloCorrectionId string) ApiDeleteSLOCorrectionRequest {
	return ApiDeleteSLOCorrectionRequest{
		ApiService:      a,
		ctx:             ctx,
		sloCorrectionId: sloCorrectionId,
	}
}

/*
 * Execute executes the request
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) DeleteSLOCorrectionExecute(r ApiDeleteSLOCorrectionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	operationId := "DeleteSLOCorrection"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceLevelObjectiveCorrectionsApiService.DeleteSLOCorrection")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/slo/correction/{slo_correction_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"slo_correction_id"+"}", _neturl.PathEscape(parameterToString(r.sloCorrectionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

	// Set Operation-ID header for telemetry
	localVarHeaderParams["DD-OPERATION-ID"] = "DeleteSLOCorrection"

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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetSLOCorrectionRequest struct {
	ctx             _context.Context
	ApiService      *ServiceLevelObjectiveCorrectionsApiService
	sloCorrectionId string
}

func (r ApiGetSLOCorrectionRequest) Execute() (SLOCorrectionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetSLOCorrectionExecute(r)
}

/*
 * GetSLOCorrection Get an SLO correction for an SLO
 * Get an SLO Correction
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sloCorrectionId The ID of the SLO correction object
 * @return ApiGetSLOCorrectionRequest
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) GetSLOCorrection(ctx _context.Context, sloCorrectionId string) ApiGetSLOCorrectionRequest {
	return ApiGetSLOCorrectionRequest{
		ApiService:      a,
		ctx:             ctx,
		sloCorrectionId: sloCorrectionId,
	}
}

/*
 * Execute executes the request
 * @return SLOCorrectionResponse
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) GetSLOCorrectionExecute(r ApiGetSLOCorrectionRequest) (SLOCorrectionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SLOCorrectionResponse
	)

	operationId := "GetSLOCorrection"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceLevelObjectiveCorrectionsApiService.GetSLOCorrection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/slo/correction/{slo_correction_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"slo_correction_id"+"}", _neturl.PathEscape(parameterToString(r.sloCorrectionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

	// Set Operation-ID header for telemetry
	localVarHeaderParams["DD-OPERATION-ID"] = "GetSLOCorrection"

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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
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

type ApiListSLOCorrectionRequest struct {
	ctx        _context.Context
	ApiService *ServiceLevelObjectiveCorrectionsApiService
}

func (r ApiListSLOCorrectionRequest) Execute() (SLOCorrectionListResponse, *_nethttp.Response, error) {
	return r.ApiService.ListSLOCorrectionExecute(r)
}

/*
 * ListSLOCorrection Get all SLO corrections
 * Get all Service Level Objective Corrections
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListSLOCorrectionRequest
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) ListSLOCorrection(ctx _context.Context) ApiListSLOCorrectionRequest {
	return ApiListSLOCorrectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SLOCorrectionListResponse
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) ListSLOCorrectionExecute(r ApiListSLOCorrectionRequest) (SLOCorrectionListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SLOCorrectionListResponse
	)

	operationId := "ListSLOCorrection"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceLevelObjectiveCorrectionsApiService.ListSLOCorrection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/slo/correction"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

	// Set Operation-ID header for telemetry
	localVarHeaderParams["DD-OPERATION-ID"] = "ListSLOCorrection"

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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
				newErr.error = err.Error()
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

type ApiUpdateSLOCorrectionRequest struct {
	ctx             _context.Context
	ApiService      *ServiceLevelObjectiveCorrectionsApiService
	sloCorrectionId string
	body            *SLOCorrectionUpdateRequest
}

func (r ApiUpdateSLOCorrectionRequest) Body(body SLOCorrectionUpdateRequest) ApiUpdateSLOCorrectionRequest {
	r.body = &body
	return r
}

func (r ApiUpdateSLOCorrectionRequest) Execute() (SLOCorrectionResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateSLOCorrectionExecute(r)
}

/*
 * UpdateSLOCorrection Update an SLO Correction
 * Update the specified SLO correction object object
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sloCorrectionId The ID of the SLO correction object
 * @return ApiUpdateSLOCorrectionRequest
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) UpdateSLOCorrection(ctx _context.Context, sloCorrectionId string) ApiUpdateSLOCorrectionRequest {
	return ApiUpdateSLOCorrectionRequest{
		ApiService:      a,
		ctx:             ctx,
		sloCorrectionId: sloCorrectionId,
	}
}

/*
 * Execute executes the request
 * @return SLOCorrectionResponse
 */
func (a *ServiceLevelObjectiveCorrectionsApiService) UpdateSLOCorrectionExecute(r ApiUpdateSLOCorrectionRequest) (SLOCorrectionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SLOCorrectionResponse
	)

	operationId := "UpdateSLOCorrection"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceLevelObjectiveCorrectionsApiService.UpdateSLOCorrection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/slo/correction/{slo_correction_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"slo_correction_id"+"}", _neturl.PathEscape(parameterToString(r.sloCorrectionId, "")), -1)

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

	// Set Operation-ID header for telemetry
	localVarHeaderParams["DD-OPERATION-ID"] = "UpdateSLOCorrection"

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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
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
