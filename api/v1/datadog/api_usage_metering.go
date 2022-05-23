/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"bytes"
	_context "context"
	_fmt "fmt"
	_ioutil "io/ioutil"
	_log "log"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
	"strings"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// UsageMeteringApiService UsageMeteringApi service
type UsageMeteringApiService service

type apiGetDailyCustomReportsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	pageSize   *int64
	pageNumber *int64
	sortDir    *UsageSortDirection
	sort       *UsageSort
}

type GetDailyCustomReportsOptionalParameters struct {
	PageSize   *int64
	PageNumber *int64
	SortDir    *UsageSortDirection
	Sort       *UsageSort
}

func NewGetDailyCustomReportsOptionalParameters() *GetDailyCustomReportsOptionalParameters {
	this := GetDailyCustomReportsOptionalParameters{}
	return &this
}
func (r *GetDailyCustomReportsOptionalParameters) WithPageSize(pageSize int64) *GetDailyCustomReportsOptionalParameters {
	r.PageSize = &pageSize
	return r
}
func (r *GetDailyCustomReportsOptionalParameters) WithPageNumber(pageNumber int64) *GetDailyCustomReportsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}
func (r *GetDailyCustomReportsOptionalParameters) WithSortDir(sortDir UsageSortDirection) *GetDailyCustomReportsOptionalParameters {
	r.SortDir = &sortDir
	return r
}
func (r *GetDailyCustomReportsOptionalParameters) WithSort(sort UsageSort) *GetDailyCustomReportsOptionalParameters {
	r.Sort = &sort
	return r
}

func (a *UsageMeteringApiService) buildGetDailyCustomReportsRequest(ctx _context.Context, o ...GetDailyCustomReportsOptionalParameters) (apiGetDailyCustomReportsRequest, error) {
	req := apiGetDailyCustomReportsRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetDailyCustomReportsOptionalParameters is allowed")
	}

	if o != nil {
		req.pageSize = o[0].PageSize
		req.pageNumber = o[0].PageNumber
		req.sortDir = o[0].SortDir
		req.sort = o[0].Sort
	}
	return req, nil
}

/*
 * GetDailyCustomReports Get the list of available daily custom reports
 * Get daily custom reports.
 */
func (a *UsageMeteringApiService) GetDailyCustomReports(ctx _context.Context, o ...GetDailyCustomReportsOptionalParameters) (UsageCustomReportsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetDailyCustomReportsRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue UsageCustomReportsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getDailyCustomReportsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageCustomReportsResponse
 */
func (a *UsageMeteringApiService) getDailyCustomReportsExecute(r apiGetDailyCustomReportsRequest) (UsageCustomReportsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageCustomReportsResponse
	)

	operationId := "GetDailyCustomReports"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetDailyCustomReports")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/daily_custom_reports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.sortDir != nil {
		localVarQueryParams.Add("sort_dir", parameterToString(*r.sortDir, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetHourlyUsageAttributionRequest struct {
	ctx              _context.Context
	ApiService       *UsageMeteringApiService
	startHr          *time.Time
	usageType        *HourlyUsageAttributionUsageType
	endHr            *time.Time
	nextRecordId     *string
	tagBreakdownKeys *string
}

type GetHourlyUsageAttributionOptionalParameters struct {
	EndHr            *time.Time
	NextRecordId     *string
	TagBreakdownKeys *string
}

func NewGetHourlyUsageAttributionOptionalParameters() *GetHourlyUsageAttributionOptionalParameters {
	this := GetHourlyUsageAttributionOptionalParameters{}
	return &this
}
func (r *GetHourlyUsageAttributionOptionalParameters) WithEndHr(endHr time.Time) *GetHourlyUsageAttributionOptionalParameters {
	r.EndHr = &endHr
	return r
}
func (r *GetHourlyUsageAttributionOptionalParameters) WithNextRecordId(nextRecordId string) *GetHourlyUsageAttributionOptionalParameters {
	r.NextRecordId = &nextRecordId
	return r
}
func (r *GetHourlyUsageAttributionOptionalParameters) WithTagBreakdownKeys(tagBreakdownKeys string) *GetHourlyUsageAttributionOptionalParameters {
	r.TagBreakdownKeys = &tagBreakdownKeys
	return r
}

func (a *UsageMeteringApiService) buildGetHourlyUsageAttributionRequest(ctx _context.Context, startHr time.Time, usageType HourlyUsageAttributionUsageType, o ...GetHourlyUsageAttributionOptionalParameters) (apiGetHourlyUsageAttributionRequest, error) {
	req := apiGetHourlyUsageAttributionRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
		usageType:  &usageType,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetHourlyUsageAttributionOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
		req.nextRecordId = o[0].NextRecordId
		req.tagBreakdownKeys = o[0].TagBreakdownKeys
	}
	return req, nil
}

/*
 * GetHourlyUsageAttribution Get hourly usage attribution
 * Get hourly usage attribution.
 *
 * This API endpoint is paginated. To make sure you receive all records, check if the value of `next_record_id` is
 * set in the response. If it is, make another request and pass `next_record_id` as a parameter.
 * Pseudo code example:
 *
 * ```
 * response := GetHourlyUsageAttribution(start_month)
 * cursor := response.metadata.pagination.next_record_id
 * WHILE cursor != null BEGIN
 *   sleep(5 seconds)  # Avoid running into rate limit
 *   response := GetHourlyUsageAttribution(start_month, next_record_id=cursor)
 *   cursor := response.metadata.pagination.next_record_id
 * END
 * ```
 */
func (a *UsageMeteringApiService) GetHourlyUsageAttribution(ctx _context.Context, startHr time.Time, usageType HourlyUsageAttributionUsageType, o ...GetHourlyUsageAttributionOptionalParameters) (HourlyUsageAttributionResponse, *_nethttp.Response, error) {
	req, err := a.buildGetHourlyUsageAttributionRequest(ctx, startHr, usageType, o...)
	if err != nil {
		var localVarReturnValue HourlyUsageAttributionResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getHourlyUsageAttributionExecute(req)
}

/*
 * Execute executes the request
 * @return HourlyUsageAttributionResponse
 */
func (a *UsageMeteringApiService) getHourlyUsageAttributionExecute(r apiGetHourlyUsageAttributionRequest) (HourlyUsageAttributionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue HourlyUsageAttributionResponse
	)

	operationId := "GetHourlyUsageAttribution"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetHourlyUsageAttribution")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/hourly-attribution"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	if r.usageType == nil {
		return localVarReturnValue, nil, reportError("usageType is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	localVarQueryParams.Add("usage_type", parameterToString(*r.usageType, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}
	if r.nextRecordId != nil {
		localVarQueryParams.Add("next_record_id", parameterToString(*r.nextRecordId, ""))
	}
	if r.tagBreakdownKeys != nil {
		localVarQueryParams.Add("tag_breakdown_keys", parameterToString(*r.tagBreakdownKeys, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetIncidentManagementRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetIncidentManagementOptionalParameters struct {
	EndHr *time.Time
}

func NewGetIncidentManagementOptionalParameters() *GetIncidentManagementOptionalParameters {
	this := GetIncidentManagementOptionalParameters{}
	return &this
}
func (r *GetIncidentManagementOptionalParameters) WithEndHr(endHr time.Time) *GetIncidentManagementOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetIncidentManagementRequest(ctx _context.Context, startHr time.Time, o ...GetIncidentManagementOptionalParameters) (apiGetIncidentManagementRequest, error) {
	req := apiGetIncidentManagementRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetIncidentManagementOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetIncidentManagement Get hourly usage for incident management
 * Get hourly usage for incident management.
 */
func (a *UsageMeteringApiService) GetIncidentManagement(ctx _context.Context, startHr time.Time, o ...GetIncidentManagementOptionalParameters) (UsageIncidentManagementResponse, *_nethttp.Response, error) {
	req, err := a.buildGetIncidentManagementRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageIncidentManagementResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getIncidentManagementExecute(req)
}

/*
 * Execute executes the request
 * @return UsageIncidentManagementResponse
 */
func (a *UsageMeteringApiService) getIncidentManagementExecute(r apiGetIncidentManagementRequest) (UsageIncidentManagementResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageIncidentManagementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetIncidentManagement")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/incident-management"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetIngestedSpansRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetIngestedSpansOptionalParameters struct {
	EndHr *time.Time
}

func NewGetIngestedSpansOptionalParameters() *GetIngestedSpansOptionalParameters {
	this := GetIngestedSpansOptionalParameters{}
	return &this
}
func (r *GetIngestedSpansOptionalParameters) WithEndHr(endHr time.Time) *GetIngestedSpansOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetIngestedSpansRequest(ctx _context.Context, startHr time.Time, o ...GetIngestedSpansOptionalParameters) (apiGetIngestedSpansRequest, error) {
	req := apiGetIngestedSpansRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetIngestedSpansOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetIngestedSpans Get hourly usage for ingested spans
 * Get hourly usage for ingested spans.
 */
func (a *UsageMeteringApiService) GetIngestedSpans(ctx _context.Context, startHr time.Time, o ...GetIngestedSpansOptionalParameters) (UsageIngestedSpansResponse, *_nethttp.Response, error) {
	req, err := a.buildGetIngestedSpansRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageIngestedSpansResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getIngestedSpansExecute(req)
}

/*
 * Execute executes the request
 * @return UsageIngestedSpansResponse
 */
func (a *UsageMeteringApiService) getIngestedSpansExecute(r apiGetIngestedSpansRequest) (UsageIngestedSpansResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageIngestedSpansResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetIngestedSpans")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/ingested-spans"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetMonthlyCustomReportsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	pageSize   *int64
	pageNumber *int64
	sortDir    *UsageSortDirection
	sort       *UsageSort
}

type GetMonthlyCustomReportsOptionalParameters struct {
	PageSize   *int64
	PageNumber *int64
	SortDir    *UsageSortDirection
	Sort       *UsageSort
}

func NewGetMonthlyCustomReportsOptionalParameters() *GetMonthlyCustomReportsOptionalParameters {
	this := GetMonthlyCustomReportsOptionalParameters{}
	return &this
}
func (r *GetMonthlyCustomReportsOptionalParameters) WithPageSize(pageSize int64) *GetMonthlyCustomReportsOptionalParameters {
	r.PageSize = &pageSize
	return r
}
func (r *GetMonthlyCustomReportsOptionalParameters) WithPageNumber(pageNumber int64) *GetMonthlyCustomReportsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}
func (r *GetMonthlyCustomReportsOptionalParameters) WithSortDir(sortDir UsageSortDirection) *GetMonthlyCustomReportsOptionalParameters {
	r.SortDir = &sortDir
	return r
}
func (r *GetMonthlyCustomReportsOptionalParameters) WithSort(sort UsageSort) *GetMonthlyCustomReportsOptionalParameters {
	r.Sort = &sort
	return r
}

func (a *UsageMeteringApiService) buildGetMonthlyCustomReportsRequest(ctx _context.Context, o ...GetMonthlyCustomReportsOptionalParameters) (apiGetMonthlyCustomReportsRequest, error) {
	req := apiGetMonthlyCustomReportsRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetMonthlyCustomReportsOptionalParameters is allowed")
	}

	if o != nil {
		req.pageSize = o[0].PageSize
		req.pageNumber = o[0].PageNumber
		req.sortDir = o[0].SortDir
		req.sort = o[0].Sort
	}
	return req, nil
}

/*
 * GetMonthlyCustomReports Get the list of available monthly custom reports
 * Get monthly custom reports.
 */
func (a *UsageMeteringApiService) GetMonthlyCustomReports(ctx _context.Context, o ...GetMonthlyCustomReportsOptionalParameters) (UsageCustomReportsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetMonthlyCustomReportsRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue UsageCustomReportsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getMonthlyCustomReportsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageCustomReportsResponse
 */
func (a *UsageMeteringApiService) getMonthlyCustomReportsExecute(r apiGetMonthlyCustomReportsRequest) (UsageCustomReportsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageCustomReportsResponse
	)

	operationId := "GetMonthlyCustomReports"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetMonthlyCustomReports")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monthly_custom_reports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.sortDir != nil {
		localVarQueryParams.Add("sort_dir", parameterToString(*r.sortDir, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetMonthlyUsageAttributionRequest struct {
	ctx              _context.Context
	ApiService       *UsageMeteringApiService
	startMonth       *time.Time
	fields           *MonthlyUsageAttributionSupportedMetrics
	endMonth         *time.Time
	sortDirection    *UsageSortDirection
	sortName         *MonthlyUsageAttributionSupportedMetrics
	tagBreakdownKeys *string
	nextRecordId     *string
}

type GetMonthlyUsageAttributionOptionalParameters struct {
	EndMonth         *time.Time
	SortDirection    *UsageSortDirection
	SortName         *MonthlyUsageAttributionSupportedMetrics
	TagBreakdownKeys *string
	NextRecordId     *string
}

func NewGetMonthlyUsageAttributionOptionalParameters() *GetMonthlyUsageAttributionOptionalParameters {
	this := GetMonthlyUsageAttributionOptionalParameters{}
	return &this
}
func (r *GetMonthlyUsageAttributionOptionalParameters) WithEndMonth(endMonth time.Time) *GetMonthlyUsageAttributionOptionalParameters {
	r.EndMonth = &endMonth
	return r
}
func (r *GetMonthlyUsageAttributionOptionalParameters) WithSortDirection(sortDirection UsageSortDirection) *GetMonthlyUsageAttributionOptionalParameters {
	r.SortDirection = &sortDirection
	return r
}
func (r *GetMonthlyUsageAttributionOptionalParameters) WithSortName(sortName MonthlyUsageAttributionSupportedMetrics) *GetMonthlyUsageAttributionOptionalParameters {
	r.SortName = &sortName
	return r
}
func (r *GetMonthlyUsageAttributionOptionalParameters) WithTagBreakdownKeys(tagBreakdownKeys string) *GetMonthlyUsageAttributionOptionalParameters {
	r.TagBreakdownKeys = &tagBreakdownKeys
	return r
}
func (r *GetMonthlyUsageAttributionOptionalParameters) WithNextRecordId(nextRecordId string) *GetMonthlyUsageAttributionOptionalParameters {
	r.NextRecordId = &nextRecordId
	return r
}

func (a *UsageMeteringApiService) buildGetMonthlyUsageAttributionRequest(ctx _context.Context, startMonth time.Time, fields MonthlyUsageAttributionSupportedMetrics, o ...GetMonthlyUsageAttributionOptionalParameters) (apiGetMonthlyUsageAttributionRequest, error) {
	req := apiGetMonthlyUsageAttributionRequest{
		ApiService: a,
		ctx:        ctx,
		startMonth: &startMonth,
		fields:     &fields,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetMonthlyUsageAttributionOptionalParameters is allowed")
	}

	if o != nil {
		req.endMonth = o[0].EndMonth
		req.sortDirection = o[0].SortDirection
		req.sortName = o[0].SortName
		req.tagBreakdownKeys = o[0].TagBreakdownKeys
		req.nextRecordId = o[0].NextRecordId
	}
	return req, nil
}

/*
 * GetMonthlyUsageAttribution Get monthly usage attribution
 * Get monthly usage attribution.
 *
 * This API endpoint is paginated. To make sure you receive all records, check if the value of `next_record_id` is
 * set in the response. If it is, make another request and pass `next_record_id` as a parameter.
 * Pseudo code example:
 *
 * ```
 * response := GetMonthlyUsageAttribution(start_month)
 * cursor := response.metadata.pagination.next_record_id
 * WHILE cursor != null BEGIN
 *   sleep(5 seconds)  # Avoid running into rate limit
 *   response := GetMonthlyUsageAttribution(start_month, next_record_id=cursor)
 *   cursor := response.metadata.pagination.next_record_id
 * END
 * ```
 */
func (a *UsageMeteringApiService) GetMonthlyUsageAttribution(ctx _context.Context, startMonth time.Time, fields MonthlyUsageAttributionSupportedMetrics, o ...GetMonthlyUsageAttributionOptionalParameters) (MonthlyUsageAttributionResponse, *_nethttp.Response, error) {
	req, err := a.buildGetMonthlyUsageAttributionRequest(ctx, startMonth, fields, o...)
	if err != nil {
		var localVarReturnValue MonthlyUsageAttributionResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getMonthlyUsageAttributionExecute(req)
}

/*
 * Execute executes the request
 * @return MonthlyUsageAttributionResponse
 */
func (a *UsageMeteringApiService) getMonthlyUsageAttributionExecute(r apiGetMonthlyUsageAttributionRequest) (MonthlyUsageAttributionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MonthlyUsageAttributionResponse
	)

	operationId := "GetMonthlyUsageAttribution"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetMonthlyUsageAttribution")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/monthly-attribution"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startMonth == nil {
		return localVarReturnValue, nil, reportError("startMonth is required and must be specified")
	}
	if r.fields == nil {
		return localVarReturnValue, nil, reportError("fields is required and must be specified")
	}
	localVarQueryParams.Add("start_month", parameterToString(*r.startMonth, ""))
	localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	if r.endMonth != nil {
		localVarQueryParams.Add("end_month", parameterToString(*r.endMonth, ""))
	}
	if r.sortDirection != nil {
		localVarQueryParams.Add("sort_direction", parameterToString(*r.sortDirection, ""))
	}
	if r.sortName != nil {
		localVarQueryParams.Add("sort_name", parameterToString(*r.sortName, ""))
	}
	if r.tagBreakdownKeys != nil {
		localVarQueryParams.Add("tag_breakdown_keys", parameterToString(*r.tagBreakdownKeys, ""))
	}
	if r.nextRecordId != nil {
		localVarQueryParams.Add("next_record_id", parameterToString(*r.nextRecordId, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetSpecifiedDailyCustomReportsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	reportId   string
}

func (a *UsageMeteringApiService) buildGetSpecifiedDailyCustomReportsRequest(ctx _context.Context, reportId string) (apiGetSpecifiedDailyCustomReportsRequest, error) {
	req := apiGetSpecifiedDailyCustomReportsRequest{
		ApiService: a,
		ctx:        ctx,
		reportId:   reportId,
	}
	return req, nil
}

/*
 * GetSpecifiedDailyCustomReports Get specified daily custom reports
 * Get specified daily custom reports.
 */
func (a *UsageMeteringApiService) GetSpecifiedDailyCustomReports(ctx _context.Context, reportId string) (UsageSpecifiedCustomReportsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetSpecifiedDailyCustomReportsRequest(ctx, reportId)
	if err != nil {
		var localVarReturnValue UsageSpecifiedCustomReportsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getSpecifiedDailyCustomReportsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageSpecifiedCustomReportsResponse
 */
func (a *UsageMeteringApiService) getSpecifiedDailyCustomReportsExecute(r apiGetSpecifiedDailyCustomReportsRequest) (UsageSpecifiedCustomReportsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageSpecifiedCustomReportsResponse
	)

	operationId := "GetSpecifiedDailyCustomReports"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetSpecifiedDailyCustomReports")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/daily_custom_reports/{report_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"report_id"+"}", _neturl.PathEscape(parameterToString(r.reportId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetSpecifiedMonthlyCustomReportsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	reportId   string
}

func (a *UsageMeteringApiService) buildGetSpecifiedMonthlyCustomReportsRequest(ctx _context.Context, reportId string) (apiGetSpecifiedMonthlyCustomReportsRequest, error) {
	req := apiGetSpecifiedMonthlyCustomReportsRequest{
		ApiService: a,
		ctx:        ctx,
		reportId:   reportId,
	}
	return req, nil
}

/*
 * GetSpecifiedMonthlyCustomReports Get specified monthly custom reports
 * Get specified monthly custom reports.
 */
func (a *UsageMeteringApiService) GetSpecifiedMonthlyCustomReports(ctx _context.Context, reportId string) (UsageSpecifiedCustomReportsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetSpecifiedMonthlyCustomReportsRequest(ctx, reportId)
	if err != nil {
		var localVarReturnValue UsageSpecifiedCustomReportsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getSpecifiedMonthlyCustomReportsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageSpecifiedCustomReportsResponse
 */
func (a *UsageMeteringApiService) getSpecifiedMonthlyCustomReportsExecute(r apiGetSpecifiedMonthlyCustomReportsRequest) (UsageSpecifiedCustomReportsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageSpecifiedCustomReportsResponse
	)

	operationId := "GetSpecifiedMonthlyCustomReports"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetSpecifiedMonthlyCustomReports")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monthly_custom_reports/{report_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"report_id"+"}", _neturl.PathEscape(parameterToString(r.reportId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageAnalyzedLogsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageAnalyzedLogsOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageAnalyzedLogsOptionalParameters() *GetUsageAnalyzedLogsOptionalParameters {
	this := GetUsageAnalyzedLogsOptionalParameters{}
	return &this
}
func (r *GetUsageAnalyzedLogsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageAnalyzedLogsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageAnalyzedLogsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageAnalyzedLogsOptionalParameters) (apiGetUsageAnalyzedLogsRequest, error) {
	req := apiGetUsageAnalyzedLogsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageAnalyzedLogsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageAnalyzedLogs Get hourly usage for analyzed logs
 * Get hourly usage for analyzed logs (Security Monitoring).
 */
func (a *UsageMeteringApiService) GetUsageAnalyzedLogs(ctx _context.Context, startHr time.Time, o ...GetUsageAnalyzedLogsOptionalParameters) (UsageAnalyzedLogsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageAnalyzedLogsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageAnalyzedLogsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageAnalyzedLogsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageAnalyzedLogsResponse
 */
func (a *UsageMeteringApiService) getUsageAnalyzedLogsExecute(r apiGetUsageAnalyzedLogsRequest) (UsageAnalyzedLogsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageAnalyzedLogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageAnalyzedLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/analyzed_logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageAttributionRequest struct {
	ctx                _context.Context
	ApiService         *UsageMeteringApiService
	startMonth         *time.Time
	fields             *UsageAttributionSupportedMetrics
	endMonth           *time.Time
	sortDirection      *UsageSortDirection
	sortName           *UsageAttributionSort
	includeDescendants *bool
	offset             *int64
	limit              *int64
}

type GetUsageAttributionOptionalParameters struct {
	EndMonth           *time.Time
	SortDirection      *UsageSortDirection
	SortName           *UsageAttributionSort
	IncludeDescendants *bool
	Offset             *int64
	Limit              *int64
}

func NewGetUsageAttributionOptionalParameters() *GetUsageAttributionOptionalParameters {
	this := GetUsageAttributionOptionalParameters{}
	return &this
}
func (r *GetUsageAttributionOptionalParameters) WithEndMonth(endMonth time.Time) *GetUsageAttributionOptionalParameters {
	r.EndMonth = &endMonth
	return r
}
func (r *GetUsageAttributionOptionalParameters) WithSortDirection(sortDirection UsageSortDirection) *GetUsageAttributionOptionalParameters {
	r.SortDirection = &sortDirection
	return r
}
func (r *GetUsageAttributionOptionalParameters) WithSortName(sortName UsageAttributionSort) *GetUsageAttributionOptionalParameters {
	r.SortName = &sortName
	return r
}
func (r *GetUsageAttributionOptionalParameters) WithIncludeDescendants(includeDescendants bool) *GetUsageAttributionOptionalParameters {
	r.IncludeDescendants = &includeDescendants
	return r
}
func (r *GetUsageAttributionOptionalParameters) WithOffset(offset int64) *GetUsageAttributionOptionalParameters {
	r.Offset = &offset
	return r
}
func (r *GetUsageAttributionOptionalParameters) WithLimit(limit int64) *GetUsageAttributionOptionalParameters {
	r.Limit = &limit
	return r
}

func (a *UsageMeteringApiService) buildGetUsageAttributionRequest(ctx _context.Context, startMonth time.Time, fields UsageAttributionSupportedMetrics, o ...GetUsageAttributionOptionalParameters) (apiGetUsageAttributionRequest, error) {
	req := apiGetUsageAttributionRequest{
		ApiService: a,
		ctx:        ctx,
		startMonth: &startMonth,
		fields:     &fields,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageAttributionOptionalParameters is allowed")
	}

	if o != nil {
		req.endMonth = o[0].EndMonth
		req.sortDirection = o[0].SortDirection
		req.sortName = o[0].SortName
		req.includeDescendants = o[0].IncludeDescendants
		req.offset = o[0].Offset
		req.limit = o[0].Limit
	}
	return req, nil
}

/*
 * GetUsageAttribution Get usage attribution
 * Get usage attribution.
 */
func (a *UsageMeteringApiService) GetUsageAttribution(ctx _context.Context, startMonth time.Time, fields UsageAttributionSupportedMetrics, o ...GetUsageAttributionOptionalParameters) (UsageAttributionResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageAttributionRequest(ctx, startMonth, fields, o...)
	if err != nil {
		var localVarReturnValue UsageAttributionResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageAttributionExecute(req)
}

/*
 * Execute executes the request
 * @return UsageAttributionResponse
 */
func (a *UsageMeteringApiService) getUsageAttributionExecute(r apiGetUsageAttributionRequest) (UsageAttributionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageAttributionResponse
	)

	operationId := "GetUsageAttribution"
	if r.ApiService.client.cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, GenericOpenAPIError{error: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageAttribution")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/attribution"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startMonth == nil {
		return localVarReturnValue, nil, reportError("startMonth is required and must be specified")
	}
	if r.fields == nil {
		return localVarReturnValue, nil, reportError("fields is required and must be specified")
	}
	localVarQueryParams.Add("start_month", parameterToString(*r.startMonth, ""))
	localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	if r.endMonth != nil {
		localVarQueryParams.Add("end_month", parameterToString(*r.endMonth, ""))
	}
	if r.sortDirection != nil {
		localVarQueryParams.Add("sort_direction", parameterToString(*r.sortDirection, ""))
	}
	if r.sortName != nil {
		localVarQueryParams.Add("sort_name", parameterToString(*r.sortName, ""))
	}
	if r.includeDescendants != nil {
		localVarQueryParams.Add("include_descendants", parameterToString(*r.includeDescendants, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageAuditLogsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageAuditLogsOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageAuditLogsOptionalParameters() *GetUsageAuditLogsOptionalParameters {
	this := GetUsageAuditLogsOptionalParameters{}
	return &this
}
func (r *GetUsageAuditLogsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageAuditLogsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageAuditLogsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageAuditLogsOptionalParameters) (apiGetUsageAuditLogsRequest, error) {
	req := apiGetUsageAuditLogsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageAuditLogsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageAuditLogs Get hourly usage for audit logs
 * Get hourly usage for audit logs.
 */
func (a *UsageMeteringApiService) GetUsageAuditLogs(ctx _context.Context, startHr time.Time, o ...GetUsageAuditLogsOptionalParameters) (UsageAuditLogsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageAuditLogsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageAuditLogsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageAuditLogsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageAuditLogsResponse
 */
func (a *UsageMeteringApiService) getUsageAuditLogsExecute(r apiGetUsageAuditLogsRequest) (UsageAuditLogsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageAuditLogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/audit_logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageBillableSummaryRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	month      *time.Time
}

type GetUsageBillableSummaryOptionalParameters struct {
	Month *time.Time
}

func NewGetUsageBillableSummaryOptionalParameters() *GetUsageBillableSummaryOptionalParameters {
	this := GetUsageBillableSummaryOptionalParameters{}
	return &this
}
func (r *GetUsageBillableSummaryOptionalParameters) WithMonth(month time.Time) *GetUsageBillableSummaryOptionalParameters {
	r.Month = &month
	return r
}

func (a *UsageMeteringApiService) buildGetUsageBillableSummaryRequest(ctx _context.Context, o ...GetUsageBillableSummaryOptionalParameters) (apiGetUsageBillableSummaryRequest, error) {
	req := apiGetUsageBillableSummaryRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageBillableSummaryOptionalParameters is allowed")
	}

	if o != nil {
		req.month = o[0].Month
	}
	return req, nil
}

/*
 * GetUsageBillableSummary Get billable usage across your account
 * Get billable usage across your account.
 */
func (a *UsageMeteringApiService) GetUsageBillableSummary(ctx _context.Context, o ...GetUsageBillableSummaryOptionalParameters) (UsageBillableSummaryResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageBillableSummaryRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue UsageBillableSummaryResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageBillableSummaryExecute(req)
}

/*
 * Execute executes the request
 * @return UsageBillableSummaryResponse
 */
func (a *UsageMeteringApiService) getUsageBillableSummaryExecute(r apiGetUsageBillableSummaryRequest) (UsageBillableSummaryResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageBillableSummaryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageBillableSummary")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/billable-summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.month != nil {
		localVarQueryParams.Add("month", parameterToString(*r.month, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageCIAppRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageCIAppOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageCIAppOptionalParameters() *GetUsageCIAppOptionalParameters {
	this := GetUsageCIAppOptionalParameters{}
	return &this
}
func (r *GetUsageCIAppOptionalParameters) WithEndHr(endHr time.Time) *GetUsageCIAppOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageCIAppRequest(ctx _context.Context, startHr time.Time, o ...GetUsageCIAppOptionalParameters) (apiGetUsageCIAppRequest, error) {
	req := apiGetUsageCIAppRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageCIAppOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageCIApp Get hourly usage for CI Visibility
 * Get hourly usage for CI Visibility (Tests, Pipeline, and Spans).
 */
func (a *UsageMeteringApiService) GetUsageCIApp(ctx _context.Context, startHr time.Time, o ...GetUsageCIAppOptionalParameters) (UsageCIVisibilityResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageCIAppRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageCIVisibilityResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageCIAppExecute(req)
}

/*
 * Execute executes the request
 * @return UsageCIVisibilityResponse
 */
func (a *UsageMeteringApiService) getUsageCIAppExecute(r apiGetUsageCIAppRequest) (UsageCIVisibilityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageCIVisibilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageCIApp")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/ci-app"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageCWSRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageCWSOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageCWSOptionalParameters() *GetUsageCWSOptionalParameters {
	this := GetUsageCWSOptionalParameters{}
	return &this
}
func (r *GetUsageCWSOptionalParameters) WithEndHr(endHr time.Time) *GetUsageCWSOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageCWSRequest(ctx _context.Context, startHr time.Time, o ...GetUsageCWSOptionalParameters) (apiGetUsageCWSRequest, error) {
	req := apiGetUsageCWSRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageCWSOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageCWS Get hourly usage for Cloud Workload Security
 * Get hourly usage for Cloud Workload Security.
 */
func (a *UsageMeteringApiService) GetUsageCWS(ctx _context.Context, startHr time.Time, o ...GetUsageCWSOptionalParameters) (UsageCWSResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageCWSRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageCWSResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageCWSExecute(req)
}

/*
 * Execute executes the request
 * @return UsageCWSResponse
 */
func (a *UsageMeteringApiService) getUsageCWSExecute(r apiGetUsageCWSRequest) (UsageCWSResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageCWSResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageCWS")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/cws"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageCloudSecurityPostureManagementRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageCloudSecurityPostureManagementOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageCloudSecurityPostureManagementOptionalParameters() *GetUsageCloudSecurityPostureManagementOptionalParameters {
	this := GetUsageCloudSecurityPostureManagementOptionalParameters{}
	return &this
}
func (r *GetUsageCloudSecurityPostureManagementOptionalParameters) WithEndHr(endHr time.Time) *GetUsageCloudSecurityPostureManagementOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageCloudSecurityPostureManagementRequest(ctx _context.Context, startHr time.Time, o ...GetUsageCloudSecurityPostureManagementOptionalParameters) (apiGetUsageCloudSecurityPostureManagementRequest, error) {
	req := apiGetUsageCloudSecurityPostureManagementRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageCloudSecurityPostureManagementOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageCloudSecurityPostureManagement Get hourly usage for CSPM
 * Get hourly usage for Cloud Security Posture Management (CSPM).
 */
func (a *UsageMeteringApiService) GetUsageCloudSecurityPostureManagement(ctx _context.Context, startHr time.Time, o ...GetUsageCloudSecurityPostureManagementOptionalParameters) (UsageCloudSecurityPostureManagementResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageCloudSecurityPostureManagementRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageCloudSecurityPostureManagementResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageCloudSecurityPostureManagementExecute(req)
}

/*
 * Execute executes the request
 * @return UsageCloudSecurityPostureManagementResponse
 */
func (a *UsageMeteringApiService) getUsageCloudSecurityPostureManagementExecute(r apiGetUsageCloudSecurityPostureManagementRequest) (UsageCloudSecurityPostureManagementResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageCloudSecurityPostureManagementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageCloudSecurityPostureManagement")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/cspm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageDBMRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageDBMOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageDBMOptionalParameters() *GetUsageDBMOptionalParameters {
	this := GetUsageDBMOptionalParameters{}
	return &this
}
func (r *GetUsageDBMOptionalParameters) WithEndHr(endHr time.Time) *GetUsageDBMOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageDBMRequest(ctx _context.Context, startHr time.Time, o ...GetUsageDBMOptionalParameters) (apiGetUsageDBMRequest, error) {
	req := apiGetUsageDBMRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageDBMOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageDBM Get hourly usage for Database Monitoring
 * Get hourly usage for Database Monitoring
 */
func (a *UsageMeteringApiService) GetUsageDBM(ctx _context.Context, startHr time.Time, o ...GetUsageDBMOptionalParameters) (UsageDBMResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageDBMRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageDBMResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageDBMExecute(req)
}

/*
 * Execute executes the request
 * @return UsageDBMResponse
 */
func (a *UsageMeteringApiService) getUsageDBMExecute(r apiGetUsageDBMRequest) (UsageDBMResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageDBMResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageDBM")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/dbm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageFargateRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageFargateOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageFargateOptionalParameters() *GetUsageFargateOptionalParameters {
	this := GetUsageFargateOptionalParameters{}
	return &this
}
func (r *GetUsageFargateOptionalParameters) WithEndHr(endHr time.Time) *GetUsageFargateOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageFargateRequest(ctx _context.Context, startHr time.Time, o ...GetUsageFargateOptionalParameters) (apiGetUsageFargateRequest, error) {
	req := apiGetUsageFargateRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageFargateOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageFargate Get hourly usage for Fargate
 * Get hourly usage for [Fargate](https://docs.datadoghq.com/integrations/ecs_fargate/).
 */
func (a *UsageMeteringApiService) GetUsageFargate(ctx _context.Context, startHr time.Time, o ...GetUsageFargateOptionalParameters) (UsageFargateResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageFargateRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageFargateResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageFargateExecute(req)
}

/*
 * Execute executes the request
 * @return UsageFargateResponse
 */
func (a *UsageMeteringApiService) getUsageFargateExecute(r apiGetUsageFargateRequest) (UsageFargateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageFargateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageFargate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/fargate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageHostsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageHostsOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageHostsOptionalParameters() *GetUsageHostsOptionalParameters {
	this := GetUsageHostsOptionalParameters{}
	return &this
}
func (r *GetUsageHostsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageHostsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageHostsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageHostsOptionalParameters) (apiGetUsageHostsRequest, error) {
	req := apiGetUsageHostsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageHostsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageHosts Get hourly usage for hosts and containers
 * Get hourly usage for hosts and containers.
 */
func (a *UsageMeteringApiService) GetUsageHosts(ctx _context.Context, startHr time.Time, o ...GetUsageHostsOptionalParameters) (UsageHostsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageHostsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageHostsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageHostsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageHostsResponse
 */
func (a *UsageMeteringApiService) getUsageHostsExecute(r apiGetUsageHostsRequest) (UsageHostsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageHostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageHosts")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/hosts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageIndexedSpansRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageIndexedSpansOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageIndexedSpansOptionalParameters() *GetUsageIndexedSpansOptionalParameters {
	this := GetUsageIndexedSpansOptionalParameters{}
	return &this
}
func (r *GetUsageIndexedSpansOptionalParameters) WithEndHr(endHr time.Time) *GetUsageIndexedSpansOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageIndexedSpansRequest(ctx _context.Context, startHr time.Time, o ...GetUsageIndexedSpansOptionalParameters) (apiGetUsageIndexedSpansRequest, error) {
	req := apiGetUsageIndexedSpansRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageIndexedSpansOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageIndexedSpans Get hourly usage for indexed spans
 * Get hourly usage for indexed spans.
 */
func (a *UsageMeteringApiService) GetUsageIndexedSpans(ctx _context.Context, startHr time.Time, o ...GetUsageIndexedSpansOptionalParameters) (UsageIndexedSpansResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageIndexedSpansRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageIndexedSpansResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageIndexedSpansExecute(req)
}

/*
 * Execute executes the request
 * @return UsageIndexedSpansResponse
 */
func (a *UsageMeteringApiService) getUsageIndexedSpansExecute(r apiGetUsageIndexedSpansRequest) (UsageIndexedSpansResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageIndexedSpansResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageIndexedSpans")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/indexed-spans"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageInternetOfThingsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageInternetOfThingsOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageInternetOfThingsOptionalParameters() *GetUsageInternetOfThingsOptionalParameters {
	this := GetUsageInternetOfThingsOptionalParameters{}
	return &this
}
func (r *GetUsageInternetOfThingsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageInternetOfThingsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageInternetOfThingsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageInternetOfThingsOptionalParameters) (apiGetUsageInternetOfThingsRequest, error) {
	req := apiGetUsageInternetOfThingsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageInternetOfThingsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageInternetOfThings Get hourly usage for IoT
 * Get hourly usage for IoT.
 */
func (a *UsageMeteringApiService) GetUsageInternetOfThings(ctx _context.Context, startHr time.Time, o ...GetUsageInternetOfThingsOptionalParameters) (UsageIoTResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageInternetOfThingsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageIoTResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageInternetOfThingsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageIoTResponse
 */
func (a *UsageMeteringApiService) getUsageInternetOfThingsExecute(r apiGetUsageInternetOfThingsRequest) (UsageIoTResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageIoTResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageInternetOfThings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/iot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageLambdaRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageLambdaOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageLambdaOptionalParameters() *GetUsageLambdaOptionalParameters {
	this := GetUsageLambdaOptionalParameters{}
	return &this
}
func (r *GetUsageLambdaOptionalParameters) WithEndHr(endHr time.Time) *GetUsageLambdaOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageLambdaRequest(ctx _context.Context, startHr time.Time, o ...GetUsageLambdaOptionalParameters) (apiGetUsageLambdaRequest, error) {
	req := apiGetUsageLambdaRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageLambdaOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageLambda Get hourly usage for Lambda
 * Get hourly usage for lambda.
 */
func (a *UsageMeteringApiService) GetUsageLambda(ctx _context.Context, startHr time.Time, o ...GetUsageLambdaOptionalParameters) (UsageLambdaResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageLambdaRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageLambdaResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageLambdaExecute(req)
}

/*
 * Execute executes the request
 * @return UsageLambdaResponse
 */
func (a *UsageMeteringApiService) getUsageLambdaExecute(r apiGetUsageLambdaRequest) (UsageLambdaResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageLambdaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageLambda")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/aws_lambda"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageLogsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageLogsOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageLogsOptionalParameters() *GetUsageLogsOptionalParameters {
	this := GetUsageLogsOptionalParameters{}
	return &this
}
func (r *GetUsageLogsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageLogsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageLogsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageLogsOptionalParameters) (apiGetUsageLogsRequest, error) {
	req := apiGetUsageLogsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageLogsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageLogs Get hourly usage for Logs
 * Get hourly usage for logs.
 */
func (a *UsageMeteringApiService) GetUsageLogs(ctx _context.Context, startHr time.Time, o ...GetUsageLogsOptionalParameters) (UsageLogsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageLogsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageLogsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageLogsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageLogsResponse
 */
func (a *UsageMeteringApiService) getUsageLogsExecute(r apiGetUsageLogsRequest) (UsageLogsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageLogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageLogsByIndexRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
	indexName  *[]string
}

type GetUsageLogsByIndexOptionalParameters struct {
	EndHr     *time.Time
	IndexName *[]string
}

func NewGetUsageLogsByIndexOptionalParameters() *GetUsageLogsByIndexOptionalParameters {
	this := GetUsageLogsByIndexOptionalParameters{}
	return &this
}
func (r *GetUsageLogsByIndexOptionalParameters) WithEndHr(endHr time.Time) *GetUsageLogsByIndexOptionalParameters {
	r.EndHr = &endHr
	return r
}
func (r *GetUsageLogsByIndexOptionalParameters) WithIndexName(indexName []string) *GetUsageLogsByIndexOptionalParameters {
	r.IndexName = &indexName
	return r
}

func (a *UsageMeteringApiService) buildGetUsageLogsByIndexRequest(ctx _context.Context, startHr time.Time, o ...GetUsageLogsByIndexOptionalParameters) (apiGetUsageLogsByIndexRequest, error) {
	req := apiGetUsageLogsByIndexRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageLogsByIndexOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
		req.indexName = o[0].IndexName
	}
	return req, nil
}

/*
 * GetUsageLogsByIndex Get hourly usage for Logs by Index
 * Get hourly usage for logs by index.
 */
func (a *UsageMeteringApiService) GetUsageLogsByIndex(ctx _context.Context, startHr time.Time, o ...GetUsageLogsByIndexOptionalParameters) (UsageLogsByIndexResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageLogsByIndexRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageLogsByIndexResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageLogsByIndexExecute(req)
}

/*
 * Execute executes the request
 * @return UsageLogsByIndexResponse
 */
func (a *UsageMeteringApiService) getUsageLogsByIndexExecute(r apiGetUsageLogsByIndexRequest) (UsageLogsByIndexResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageLogsByIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageLogsByIndex")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/logs_by_index"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}
	if r.indexName != nil {
		t := *r.indexName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("index_name", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("index_name", parameterToString(t, "multi"))
		}
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageLogsByRetentionRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageLogsByRetentionOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageLogsByRetentionOptionalParameters() *GetUsageLogsByRetentionOptionalParameters {
	this := GetUsageLogsByRetentionOptionalParameters{}
	return &this
}
func (r *GetUsageLogsByRetentionOptionalParameters) WithEndHr(endHr time.Time) *GetUsageLogsByRetentionOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageLogsByRetentionRequest(ctx _context.Context, startHr time.Time, o ...GetUsageLogsByRetentionOptionalParameters) (apiGetUsageLogsByRetentionRequest, error) {
	req := apiGetUsageLogsByRetentionRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageLogsByRetentionOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageLogsByRetention Get hourly logs usage by retention
 * Get hourly usage for indexed logs by retention period.
 */
func (a *UsageMeteringApiService) GetUsageLogsByRetention(ctx _context.Context, startHr time.Time, o ...GetUsageLogsByRetentionOptionalParameters) (UsageLogsByRetentionResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageLogsByRetentionRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageLogsByRetentionResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageLogsByRetentionExecute(req)
}

/*
 * Execute executes the request
 * @return UsageLogsByRetentionResponse
 */
func (a *UsageMeteringApiService) getUsageLogsByRetentionExecute(r apiGetUsageLogsByRetentionRequest) (UsageLogsByRetentionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageLogsByRetentionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageLogsByRetention")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/logs-by-retention"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageNetworkFlowsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageNetworkFlowsOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageNetworkFlowsOptionalParameters() *GetUsageNetworkFlowsOptionalParameters {
	this := GetUsageNetworkFlowsOptionalParameters{}
	return &this
}
func (r *GetUsageNetworkFlowsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageNetworkFlowsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageNetworkFlowsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageNetworkFlowsOptionalParameters) (apiGetUsageNetworkFlowsRequest, error) {
	req := apiGetUsageNetworkFlowsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageNetworkFlowsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageNetworkFlows Get hourly usage for Network Flows
 * Get hourly usage for network flows.
 */
func (a *UsageMeteringApiService) GetUsageNetworkFlows(ctx _context.Context, startHr time.Time, o ...GetUsageNetworkFlowsOptionalParameters) (UsageNetworkFlowsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageNetworkFlowsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageNetworkFlowsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageNetworkFlowsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageNetworkFlowsResponse
 */
func (a *UsageMeteringApiService) getUsageNetworkFlowsExecute(r apiGetUsageNetworkFlowsRequest) (UsageNetworkFlowsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageNetworkFlowsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageNetworkFlows")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/network_flows"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageNetworkHostsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageNetworkHostsOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageNetworkHostsOptionalParameters() *GetUsageNetworkHostsOptionalParameters {
	this := GetUsageNetworkHostsOptionalParameters{}
	return &this
}
func (r *GetUsageNetworkHostsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageNetworkHostsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageNetworkHostsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageNetworkHostsOptionalParameters) (apiGetUsageNetworkHostsRequest, error) {
	req := apiGetUsageNetworkHostsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageNetworkHostsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageNetworkHosts Get hourly usage for Network Hosts
 * Get hourly usage for network hosts.
 */
func (a *UsageMeteringApiService) GetUsageNetworkHosts(ctx _context.Context, startHr time.Time, o ...GetUsageNetworkHostsOptionalParameters) (UsageNetworkHostsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageNetworkHostsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageNetworkHostsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageNetworkHostsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageNetworkHostsResponse
 */
func (a *UsageMeteringApiService) getUsageNetworkHostsExecute(r apiGetUsageNetworkHostsRequest) (UsageNetworkHostsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageNetworkHostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageNetworkHosts")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/network_hosts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageOnlineArchiveRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageOnlineArchiveOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageOnlineArchiveOptionalParameters() *GetUsageOnlineArchiveOptionalParameters {
	this := GetUsageOnlineArchiveOptionalParameters{}
	return &this
}
func (r *GetUsageOnlineArchiveOptionalParameters) WithEndHr(endHr time.Time) *GetUsageOnlineArchiveOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageOnlineArchiveRequest(ctx _context.Context, startHr time.Time, o ...GetUsageOnlineArchiveOptionalParameters) (apiGetUsageOnlineArchiveRequest, error) {
	req := apiGetUsageOnlineArchiveRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageOnlineArchiveOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageOnlineArchive Get hourly usage for Online Archive
 * Get hourly usage for Online Archive.
 */
func (a *UsageMeteringApiService) GetUsageOnlineArchive(ctx _context.Context, startHr time.Time, o ...GetUsageOnlineArchiveOptionalParameters) (UsageOnlineArchiveResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageOnlineArchiveRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageOnlineArchiveResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageOnlineArchiveExecute(req)
}

/*
 * Execute executes the request
 * @return UsageOnlineArchiveResponse
 */
func (a *UsageMeteringApiService) getUsageOnlineArchiveExecute(r apiGetUsageOnlineArchiveRequest) (UsageOnlineArchiveResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageOnlineArchiveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageOnlineArchive")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/online-archive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageProfilingRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageProfilingOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageProfilingOptionalParameters() *GetUsageProfilingOptionalParameters {
	this := GetUsageProfilingOptionalParameters{}
	return &this
}
func (r *GetUsageProfilingOptionalParameters) WithEndHr(endHr time.Time) *GetUsageProfilingOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageProfilingRequest(ctx _context.Context, startHr time.Time, o ...GetUsageProfilingOptionalParameters) (apiGetUsageProfilingRequest, error) {
	req := apiGetUsageProfilingRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageProfilingOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageProfiling Get hourly usage for profiled hosts
 * Get hourly usage for profiled hosts.
 */
func (a *UsageMeteringApiService) GetUsageProfiling(ctx _context.Context, startHr time.Time, o ...GetUsageProfilingOptionalParameters) (UsageProfilingResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageProfilingRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageProfilingResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageProfilingExecute(req)
}

/*
 * Execute executes the request
 * @return UsageProfilingResponse
 */
func (a *UsageMeteringApiService) getUsageProfilingExecute(r apiGetUsageProfilingRequest) (UsageProfilingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageProfilingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageProfiling")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/profiling"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageRumSessionsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
	type_      *string
}

type GetUsageRumSessionsOptionalParameters struct {
	EndHr *time.Time
	Type  *string
}

func NewGetUsageRumSessionsOptionalParameters() *GetUsageRumSessionsOptionalParameters {
	this := GetUsageRumSessionsOptionalParameters{}
	return &this
}
func (r *GetUsageRumSessionsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageRumSessionsOptionalParameters {
	r.EndHr = &endHr
	return r
}
func (r *GetUsageRumSessionsOptionalParameters) WithType(type_ string) *GetUsageRumSessionsOptionalParameters {
	r.Type = &type_
	return r
}

func (a *UsageMeteringApiService) buildGetUsageRumSessionsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageRumSessionsOptionalParameters) (apiGetUsageRumSessionsRequest, error) {
	req := apiGetUsageRumSessionsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageRumSessionsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
		req.type_ = o[0].Type
	}
	return req, nil
}

/*
 * GetUsageRumSessions Get hourly usage for RUM Sessions
 * Get hourly usage for [RUM](https://docs.datadoghq.com/real_user_monitoring/) Sessions.
 */
func (a *UsageMeteringApiService) GetUsageRumSessions(ctx _context.Context, startHr time.Time, o ...GetUsageRumSessionsOptionalParameters) (UsageRumSessionsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageRumSessionsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageRumSessionsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageRumSessionsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageRumSessionsResponse
 */
func (a *UsageMeteringApiService) getUsageRumSessionsExecute(r apiGetUsageRumSessionsRequest) (UsageRumSessionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageRumSessionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageRumSessions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/rum_sessions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageRumUnitsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageRumUnitsOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageRumUnitsOptionalParameters() *GetUsageRumUnitsOptionalParameters {
	this := GetUsageRumUnitsOptionalParameters{}
	return &this
}
func (r *GetUsageRumUnitsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageRumUnitsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageRumUnitsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageRumUnitsOptionalParameters) (apiGetUsageRumUnitsRequest, error) {
	req := apiGetUsageRumUnitsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageRumUnitsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageRumUnits Get hourly usage for RUM Units
 * Get hourly usage for [RUM](https://docs.datadoghq.com/real_user_monitoring/) Units.
 */
func (a *UsageMeteringApiService) GetUsageRumUnits(ctx _context.Context, startHr time.Time, o ...GetUsageRumUnitsOptionalParameters) (UsageRumUnitsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageRumUnitsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageRumUnitsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageRumUnitsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageRumUnitsResponse
 */
func (a *UsageMeteringApiService) getUsageRumUnitsExecute(r apiGetUsageRumUnitsRequest) (UsageRumUnitsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageRumUnitsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageRumUnits")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/rum"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageSDSRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageSDSOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageSDSOptionalParameters() *GetUsageSDSOptionalParameters {
	this := GetUsageSDSOptionalParameters{}
	return &this
}
func (r *GetUsageSDSOptionalParameters) WithEndHr(endHr time.Time) *GetUsageSDSOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageSDSRequest(ctx _context.Context, startHr time.Time, o ...GetUsageSDSOptionalParameters) (apiGetUsageSDSRequest, error) {
	req := apiGetUsageSDSRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageSDSOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageSDS Get hourly usage for Sensitive Data Scanner
 * Get hourly usage for Sensitive Data Scanner.
 */
func (a *UsageMeteringApiService) GetUsageSDS(ctx _context.Context, startHr time.Time, o ...GetUsageSDSOptionalParameters) (UsageSDSResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageSDSRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageSDSResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageSDSExecute(req)
}

/*
 * Execute executes the request
 * @return UsageSDSResponse
 */
func (a *UsageMeteringApiService) getUsageSDSExecute(r apiGetUsageSDSRequest) (UsageSDSResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageSDSResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageSDS")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/sds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageSNMPRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageSNMPOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageSNMPOptionalParameters() *GetUsageSNMPOptionalParameters {
	this := GetUsageSNMPOptionalParameters{}
	return &this
}
func (r *GetUsageSNMPOptionalParameters) WithEndHr(endHr time.Time) *GetUsageSNMPOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageSNMPRequest(ctx _context.Context, startHr time.Time, o ...GetUsageSNMPOptionalParameters) (apiGetUsageSNMPRequest, error) {
	req := apiGetUsageSNMPRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageSNMPOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageSNMP Get hourly usage for SNMP devices
 * Get hourly usage for SNMP devices.
 */
func (a *UsageMeteringApiService) GetUsageSNMP(ctx _context.Context, startHr time.Time, o ...GetUsageSNMPOptionalParameters) (UsageSNMPResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageSNMPRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageSNMPResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageSNMPExecute(req)
}

/*
 * Execute executes the request
 * @return UsageSNMPResponse
 */
func (a *UsageMeteringApiService) getUsageSNMPExecute(r apiGetUsageSNMPRequest) (UsageSNMPResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageSNMPResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageSNMP")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/snmp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageSummaryRequest struct {
	ctx               _context.Context
	ApiService        *UsageMeteringApiService
	startMonth        *time.Time
	endMonth          *time.Time
	includeOrgDetails *bool
}

type GetUsageSummaryOptionalParameters struct {
	EndMonth          *time.Time
	IncludeOrgDetails *bool
}

func NewGetUsageSummaryOptionalParameters() *GetUsageSummaryOptionalParameters {
	this := GetUsageSummaryOptionalParameters{}
	return &this
}
func (r *GetUsageSummaryOptionalParameters) WithEndMonth(endMonth time.Time) *GetUsageSummaryOptionalParameters {
	r.EndMonth = &endMonth
	return r
}
func (r *GetUsageSummaryOptionalParameters) WithIncludeOrgDetails(includeOrgDetails bool) *GetUsageSummaryOptionalParameters {
	r.IncludeOrgDetails = &includeOrgDetails
	return r
}

func (a *UsageMeteringApiService) buildGetUsageSummaryRequest(ctx _context.Context, startMonth time.Time, o ...GetUsageSummaryOptionalParameters) (apiGetUsageSummaryRequest, error) {
	req := apiGetUsageSummaryRequest{
		ApiService: a,
		ctx:        ctx,
		startMonth: &startMonth,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageSummaryOptionalParameters is allowed")
	}

	if o != nil {
		req.endMonth = o[0].EndMonth
		req.includeOrgDetails = o[0].IncludeOrgDetails
	}
	return req, nil
}

/*
 * GetUsageSummary Get usage across your multi-org account
 * Get usage across your multi-org account. You must have the multi-org feature enabled.
 */
func (a *UsageMeteringApiService) GetUsageSummary(ctx _context.Context, startMonth time.Time, o ...GetUsageSummaryOptionalParameters) (UsageSummaryResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageSummaryRequest(ctx, startMonth, o...)
	if err != nil {
		var localVarReturnValue UsageSummaryResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageSummaryExecute(req)
}

/*
 * Execute executes the request
 * @return UsageSummaryResponse
 */
func (a *UsageMeteringApiService) getUsageSummaryExecute(r apiGetUsageSummaryRequest) (UsageSummaryResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageSummaryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageSummary")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startMonth == nil {
		return localVarReturnValue, nil, reportError("startMonth is required and must be specified")
	}
	localVarQueryParams.Add("start_month", parameterToString(*r.startMonth, ""))
	if r.endMonth != nil {
		localVarQueryParams.Add("end_month", parameterToString(*r.endMonth, ""))
	}
	if r.includeOrgDetails != nil {
		localVarQueryParams.Add("include_org_details", parameterToString(*r.includeOrgDetails, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageSyntheticsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageSyntheticsOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageSyntheticsOptionalParameters() *GetUsageSyntheticsOptionalParameters {
	this := GetUsageSyntheticsOptionalParameters{}
	return &this
}
func (r *GetUsageSyntheticsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageSyntheticsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageSyntheticsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageSyntheticsOptionalParameters) (apiGetUsageSyntheticsRequest, error) {
	req := apiGetUsageSyntheticsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageSyntheticsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageSynthetics Get hourly usage for Synthetics Checks
 * Get hourly usage for [Synthetics checks](https://docs.datadoghq.com/synthetics/).
 */
func (a *UsageMeteringApiService) GetUsageSynthetics(ctx _context.Context, startHr time.Time, o ...GetUsageSyntheticsOptionalParameters) (UsageSyntheticsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageSyntheticsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageSyntheticsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageSyntheticsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageSyntheticsResponse
 */
func (a *UsageMeteringApiService) getUsageSyntheticsExecute(r apiGetUsageSyntheticsRequest) (UsageSyntheticsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageSyntheticsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageSynthetics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/synthetics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageSyntheticsAPIRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageSyntheticsAPIOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageSyntheticsAPIOptionalParameters() *GetUsageSyntheticsAPIOptionalParameters {
	this := GetUsageSyntheticsAPIOptionalParameters{}
	return &this
}
func (r *GetUsageSyntheticsAPIOptionalParameters) WithEndHr(endHr time.Time) *GetUsageSyntheticsAPIOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageSyntheticsAPIRequest(ctx _context.Context, startHr time.Time, o ...GetUsageSyntheticsAPIOptionalParameters) (apiGetUsageSyntheticsAPIRequest, error) {
	req := apiGetUsageSyntheticsAPIRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageSyntheticsAPIOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageSyntheticsAPI Get hourly usage for Synthetics API Checks
 * Get hourly usage for [synthetics API checks](https://docs.datadoghq.com/synthetics/).
 */
func (a *UsageMeteringApiService) GetUsageSyntheticsAPI(ctx _context.Context, startHr time.Time, o ...GetUsageSyntheticsAPIOptionalParameters) (UsageSyntheticsAPIResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageSyntheticsAPIRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageSyntheticsAPIResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageSyntheticsAPIExecute(req)
}

/*
 * Execute executes the request
 * @return UsageSyntheticsAPIResponse
 */
func (a *UsageMeteringApiService) getUsageSyntheticsAPIExecute(r apiGetUsageSyntheticsAPIRequest) (UsageSyntheticsAPIResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageSyntheticsAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageSyntheticsAPI")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/synthetics_api"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageSyntheticsBrowserRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageSyntheticsBrowserOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageSyntheticsBrowserOptionalParameters() *GetUsageSyntheticsBrowserOptionalParameters {
	this := GetUsageSyntheticsBrowserOptionalParameters{}
	return &this
}
func (r *GetUsageSyntheticsBrowserOptionalParameters) WithEndHr(endHr time.Time) *GetUsageSyntheticsBrowserOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageSyntheticsBrowserRequest(ctx _context.Context, startHr time.Time, o ...GetUsageSyntheticsBrowserOptionalParameters) (apiGetUsageSyntheticsBrowserRequest, error) {
	req := apiGetUsageSyntheticsBrowserRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageSyntheticsBrowserOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageSyntheticsBrowser Get hourly usage for Synthetics Browser Checks
 * Get hourly usage for synthetics browser checks.
 */
func (a *UsageMeteringApiService) GetUsageSyntheticsBrowser(ctx _context.Context, startHr time.Time, o ...GetUsageSyntheticsBrowserOptionalParameters) (UsageSyntheticsBrowserResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageSyntheticsBrowserRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageSyntheticsBrowserResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageSyntheticsBrowserExecute(req)
}

/*
 * Execute executes the request
 * @return UsageSyntheticsBrowserResponse
 */
func (a *UsageMeteringApiService) getUsageSyntheticsBrowserExecute(r apiGetUsageSyntheticsBrowserRequest) (UsageSyntheticsBrowserResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageSyntheticsBrowserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageSyntheticsBrowser")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/synthetics_browser"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageTimeseriesRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
}

type GetUsageTimeseriesOptionalParameters struct {
	EndHr *time.Time
}

func NewGetUsageTimeseriesOptionalParameters() *GetUsageTimeseriesOptionalParameters {
	this := GetUsageTimeseriesOptionalParameters{}
	return &this
}
func (r *GetUsageTimeseriesOptionalParameters) WithEndHr(endHr time.Time) *GetUsageTimeseriesOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApiService) buildGetUsageTimeseriesRequest(ctx _context.Context, startHr time.Time, o ...GetUsageTimeseriesOptionalParameters) (apiGetUsageTimeseriesRequest, error) {
	req := apiGetUsageTimeseriesRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageTimeseriesOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

/*
 * GetUsageTimeseries Get hourly usage for custom metrics
 * Get hourly usage for [custom metrics](https://docs.datadoghq.com/developers/metrics/custom_metrics/).
 */
func (a *UsageMeteringApiService) GetUsageTimeseries(ctx _context.Context, startHr time.Time, o ...GetUsageTimeseriesOptionalParameters) (UsageTimeseriesResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageTimeseriesRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageTimeseriesResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageTimeseriesExecute(req)
}

/*
 * Execute executes the request
 * @return UsageTimeseriesResponse
 */
func (a *UsageMeteringApiService) getUsageTimeseriesExecute(r apiGetUsageTimeseriesRequest) (UsageTimeseriesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageTimeseriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageTimeseries")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/timeseries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, reportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", parameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", parameterToString(*r.endHr, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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

type apiGetUsageTopAvgMetricsRequest struct {
	ctx          _context.Context
	ApiService   *UsageMeteringApiService
	month        *time.Time
	day          *time.Time
	names        *[]string
	limit        *int32
	nextRecordId *string
}

type GetUsageTopAvgMetricsOptionalParameters struct {
	Month        *time.Time
	Day          *time.Time
	Names        *[]string
	Limit        *int32
	NextRecordId *string
}

func NewGetUsageTopAvgMetricsOptionalParameters() *GetUsageTopAvgMetricsOptionalParameters {
	this := GetUsageTopAvgMetricsOptionalParameters{}
	return &this
}
func (r *GetUsageTopAvgMetricsOptionalParameters) WithMonth(month time.Time) *GetUsageTopAvgMetricsOptionalParameters {
	r.Month = &month
	return r
}
func (r *GetUsageTopAvgMetricsOptionalParameters) WithDay(day time.Time) *GetUsageTopAvgMetricsOptionalParameters {
	r.Day = &day
	return r
}
func (r *GetUsageTopAvgMetricsOptionalParameters) WithNames(names []string) *GetUsageTopAvgMetricsOptionalParameters {
	r.Names = &names
	return r
}
func (r *GetUsageTopAvgMetricsOptionalParameters) WithLimit(limit int32) *GetUsageTopAvgMetricsOptionalParameters {
	r.Limit = &limit
	return r
}
func (r *GetUsageTopAvgMetricsOptionalParameters) WithNextRecordId(nextRecordId string) *GetUsageTopAvgMetricsOptionalParameters {
	r.NextRecordId = &nextRecordId
	return r
}

func (a *UsageMeteringApiService) buildGetUsageTopAvgMetricsRequest(ctx _context.Context, o ...GetUsageTopAvgMetricsOptionalParameters) (apiGetUsageTopAvgMetricsRequest, error) {
	req := apiGetUsageTopAvgMetricsRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type GetUsageTopAvgMetricsOptionalParameters is allowed")
	}

	if o != nil {
		req.month = o[0].Month
		req.day = o[0].Day
		req.names = o[0].Names
		req.limit = o[0].Limit
		req.nextRecordId = o[0].NextRecordId
	}
	return req, nil
}

/*
 * GetUsageTopAvgMetrics Get all custom metrics by hourly average
 * Get all [custom metrics](https://docs.datadoghq.com/developers/metrics/custom_metrics/) by hourly average. Use the month parameter to get a month-to-date data resolution or use the day parameter to get a daily resolution. One of the two is required, and only one of the two is allowed.
 */
func (a *UsageMeteringApiService) GetUsageTopAvgMetrics(ctx _context.Context, o ...GetUsageTopAvgMetricsOptionalParameters) (UsageTopAvgMetricsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageTopAvgMetricsRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue UsageTopAvgMetricsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageTopAvgMetricsExecute(req)
}

/*
 * Execute executes the request
 * @return UsageTopAvgMetricsResponse
 */
func (a *UsageMeteringApiService) getUsageTopAvgMetricsExecute(r apiGetUsageTopAvgMetricsRequest) (UsageTopAvgMetricsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageTopAvgMetricsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageMeteringApiService.GetUsageTopAvgMetrics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/usage/top_avg_metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.month != nil {
		localVarQueryParams.Add("month", parameterToString(*r.month, ""))
	}
	if r.day != nil {
		localVarQueryParams.Add("day", parameterToString(*r.day, ""))
	}
	if r.names != nil {
		t := *r.names
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("names", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("names", parameterToString(t, "multi"))
		}
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.nextRecordId != nil {
		localVarQueryParams.Add("next_record_id", parameterToString(*r.nextRecordId, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;datetime-format=rfc3339"}

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
		if localVarHTTPResponse.StatusCode == 429 {
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
