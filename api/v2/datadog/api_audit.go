/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// AuditApiService AuditApi service
type AuditApiService service

type apiListAuditLogsRequest struct {
	ctx         _context.Context
	ApiService  *AuditApiService
	filterQuery *string
	filterFrom  *time.Time
	filterTo    *time.Time
	sort        *AuditLogsSort
	pageCursor  *string
	pageLimit   *int32
}

type ListAuditLogsOptionalParameters struct {
	FilterQuery *string
	FilterFrom  *time.Time
	FilterTo    *time.Time
	Sort        *AuditLogsSort
	PageCursor  *string
	PageLimit   *int32
}

func NewListAuditLogsOptionalParameters() *ListAuditLogsOptionalParameters {
	this := ListAuditLogsOptionalParameters{}
	return &this
}
func (r *ListAuditLogsOptionalParameters) WithFilterQuery(filterQuery string) *ListAuditLogsOptionalParameters {
	r.FilterQuery = &filterQuery
	return r
}
func (r *ListAuditLogsOptionalParameters) WithFilterFrom(filterFrom time.Time) *ListAuditLogsOptionalParameters {
	r.FilterFrom = &filterFrom
	return r
}
func (r *ListAuditLogsOptionalParameters) WithFilterTo(filterTo time.Time) *ListAuditLogsOptionalParameters {
	r.FilterTo = &filterTo
	return r
}
func (r *ListAuditLogsOptionalParameters) WithSort(sort AuditLogsSort) *ListAuditLogsOptionalParameters {
	r.Sort = &sort
	return r
}
func (r *ListAuditLogsOptionalParameters) WithPageCursor(pageCursor string) *ListAuditLogsOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}
func (r *ListAuditLogsOptionalParameters) WithPageLimit(pageLimit int32) *ListAuditLogsOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

/*
 * ListAuditLogs Get a list of Audit Logs events
 * List endpoint returns events that match a Audit Logs search query.
 * [Results are paginated][1].
 *
 * Use this endpoint to see your latest Audit Logs events.
 *
 * [1]: https://docs.datadoghq.com/logs/guide/collect-multiple-logs-with-pagination
 */
func (a *AuditApiService) ListAuditLogs(ctx _context.Context, o ...ListAuditLogsOptionalParameters) (AuditLogsEventsResponse, *_nethttp.Response, error) {
	req := apiListAuditLogsRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		var localVarReturnValue AuditLogsEventsResponse
		return localVarReturnValue, nil, reportError("only one argument of type ListAuditLogsOptionalParameters is allowed")
	}

	if o != nil {
		req.filterQuery = o[0].FilterQuery
		req.filterFrom = o[0].FilterFrom
		req.filterTo = o[0].FilterTo
		req.sort = o[0].Sort
		req.pageCursor = o[0].PageCursor
		req.pageLimit = o[0].PageLimit
	}

	return req.ApiService.listAuditLogsExecute(req)
}

/*
 * Execute executes the request
 * @return AuditLogsEventsResponse
 */
func (a *AuditApiService) listAuditLogsExecute(r apiListAuditLogsRequest) (AuditLogsEventsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AuditLogsEventsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditApiService.ListAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/audit/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.filterQuery != nil {
		localVarQueryParams.Add("filter[query]", parameterToString(*r.filterQuery, ""))
	}
	if r.filterFrom != nil {
		localVarQueryParams.Add("filter[from]", parameterToString(*r.filterFrom, ""))
	}
	if r.filterTo != nil {
		localVarQueryParams.Add("filter[to]", parameterToString(*r.filterTo, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.pageCursor != nil {
		localVarQueryParams.Add("page[cursor]", parameterToString(*r.pageCursor, ""))
	}
	if r.pageLimit != nil {
		localVarQueryParams.Add("page[limit]", parameterToString(*r.pageLimit, ""))
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

type apiSearchAuditLogsRequest struct {
	ctx        _context.Context
	ApiService *AuditApiService
	body       *AuditLogsSearchEventsRequest
}

type SearchAuditLogsOptionalParameters struct {
	Body *AuditLogsSearchEventsRequest
}

func NewSearchAuditLogsOptionalParameters() *SearchAuditLogsOptionalParameters {
	this := SearchAuditLogsOptionalParameters{}
	return &this
}
func (r *SearchAuditLogsOptionalParameters) WithBody(body AuditLogsSearchEventsRequest) *SearchAuditLogsOptionalParameters {
	r.Body = &body
	return r
}

/*
 * SearchAuditLogs Search Audit Logs events
 * List endpoint returns Audit Logs events that match an Audit search query.
 * [Results are paginated][1].
 *
 * Use this endpoint to build complex Audit Logs events filtering and search.
 *
 * [1]: https://docs.datadoghq.com/logs/guide/collect-multiple-logs-with-pagination
 */
func (a *AuditApiService) SearchAuditLogs(ctx _context.Context, o ...SearchAuditLogsOptionalParameters) (AuditLogsEventsResponse, *_nethttp.Response, error) {
	req := apiSearchAuditLogsRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		var localVarReturnValue AuditLogsEventsResponse
		return localVarReturnValue, nil, reportError("only one argument of type SearchAuditLogsOptionalParameters is allowed")
	}

	if o != nil {
		req.body = o[0].Body
	}

	return req.ApiService.searchAuditLogsExecute(req)
}

/*
 * Execute executes the request
 * @return AuditLogsEventsResponse
 */
func (a *AuditApiService) searchAuditLogsExecute(r apiSearchAuditLogsRequest) (AuditLogsEventsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AuditLogsEventsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditApiService.SearchAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/audit/events/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
