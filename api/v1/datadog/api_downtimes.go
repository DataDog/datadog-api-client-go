/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DowntimesApiService DowntimesApi service
type DowntimesApiService service

type apiCancelDowntimeRequest struct {
	ctx        _context.Context
	apiService *DowntimesApiService
	downtimeId int64
}

/*
CancelDowntime Cancel a downtime
Cancel a Downtime
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param downtimeId ID of the downtime to cancel
@return apiCancelDowntimeRequest
*/
func (a *DowntimesApiService) CancelDowntime(ctx _context.Context, downtimeId int64) apiCancelDowntimeRequest {
	return apiCancelDowntimeRequest{
		apiService: a,
		ctx:        ctx,
		downtimeId: downtimeId,
	}
}

/*
Execute executes the request

*/
func (r apiCancelDowntimeRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "DowntimesApiService.CancelDowntime")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/downtime/{downtime_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"downtime_id"+"}", _neturl.QueryEscape(parameterToString(r.downtimeId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["apiKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["appKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
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

type apiCancelDowntimesByScopeRequest struct {
	ctx        _context.Context
	apiService *DowntimesApiService
	body       *CancelDowntimesByScopeRequest
}

func (r apiCancelDowntimesByScopeRequest) Body(body CancelDowntimesByScopeRequest) apiCancelDowntimesByScopeRequest {
	r.body = &body
	return r
}

/*
CancelDowntimesByScope Cancel downtimes by scope
### Overview
DELETE all Downtimes that match the scope of X
### Arguments
* **`scope`** [*required*]: Cancel all downtimes with the given scope(s),
  e.g.: `env:prod`, `role:db,role:db-slave`
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiCancelDowntimesByScopeRequest
*/
func (a *DowntimesApiService) CancelDowntimesByScope(ctx _context.Context) apiCancelDowntimesByScopeRequest {
	return apiCancelDowntimesByScopeRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return CanceledDowntimesIds
*/
func (r apiCancelDowntimesByScopeRequest) Execute() (CanceledDowntimesIds, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CanceledDowntimesIds
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "DowntimesApiService.CancelDowntimesByScope")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/downtime/cancel/by_scope"

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
			if auth, ok := auth["apiKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["appKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
		if localVarHTTPResponse.StatusCode == 200 {
			var v CanceledDowntimesIds
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiCreateDowntimeRequest struct {
	ctx        _context.Context
	apiService *DowntimesApiService
	body       *Downtime
}

func (r apiCreateDowntimeRequest) Body(body Downtime) apiCreateDowntimeRequest {
	r.body = &body
	return r
}

/*
CreateDowntime Schedule a downtime
* **`scope`** [*required*]: The scope(s) to which the downtime applies, e.g. `host:app2`.
  Provide multiple scopes as a comma-separated list, e.g. `env:dev,env:prod`. The
  resulting downtime applies to sources that matches ALL provided scopes (i.e.
  `env:dev` **AND** `env:prod`), NOT any of them.

* **`monitor_tags`** [*optional*, *default*=**no monitor tag filter**]: A comma-separated
  list of monitor tags, i.e. tags that are applied directly to monitors, *not* tags
  that are used in monitor queries (which are filtered by the `scope` parameter), to
  which the downtime applies. The resulting downtime applies to monitors that match
  ALL provided monitor tags (i.e. `service:postgres` **AND** `team:frontend`), NOT any of them.

* **`monitor_id`** [*optional*, *default*=**None**]: A single monitor to which the downtime
  applies. If not provided, the downtime applies to all monitors.

* **`start`** [*optional*, *default*=**None**]: POSIX timestamp to start the downtime.
  If not provided, the downtime starts the moment it is created.

* **`end`** [*optional*, *default*=**None**]: POSIX timestamp to end the downtime.
  If not provided, the downtime is in effect indefinitely (i.e. until you cancel it).

* **`message`** [*optional*, *default*=**None**]: A message to include with notifications
  for this downtime. Email notifications can be sent to specific users by using
   the same '@username' notation as events

* **`timezone`** [*optional*, *default* = **UTC**]: The timezone for the downtime.
* **`recurrence`** [*optional*, *default*=**None**]: An object defining the recurrence of the
  downtime with a variety of parameters:

  * **`type`** the type of recurrence. Choose from: `days`, `weeks`, `months`, `years`.

  * **`period`** how often to repeat as an integer. For example to repeat every 3 days,
    select a type of `days` and a period of `3`.

  * **`week_days`** (optional) a list of week days to repeat on. Choose from: `Mon`,
    `Tue`, `Wed`, `Thu`, `Fri`, `Sat` or `Sun`. Only applicable when `type` is `weeks`.
    **First letter must be capitalized.**
  * **`until_occurrences`** (optional) how many times the downtime is rescheduled.
    **`until_occurrences` and `until_date`** are mutually exclusive

  * **`until_date`** (optional) the date at which the recurrence should end
    as a POSIX timestmap. **`until_occurrences` and `until_date`** are mutually exclusive
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiCreateDowntimeRequest
*/
func (a *DowntimesApiService) CreateDowntime(ctx _context.Context) apiCreateDowntimeRequest {
	return apiCreateDowntimeRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return Downtime
*/
func (r apiCreateDowntimeRequest) Execute() (Downtime, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Downtime
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "DowntimesApiService.CreateDowntime")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/downtime"

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
			if auth, ok := auth["apiKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["appKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
		if localVarHTTPResponse.StatusCode == 200 {
			var v Downtime
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetAllDowntimesRequest struct {
	ctx         _context.Context
	apiService  *DowntimesApiService
	currentOnly *bool
}

func (r apiGetAllDowntimesRequest) CurrentOnly(currentOnly bool) apiGetAllDowntimesRequest {
	r.currentOnly = &currentOnly
	return r
}

/*
GetAllDowntimes Get all downtimes
### Overview
Get All Scheduled Downtimes
### Arguments
* **`current_only`** [*optional*, *default* = **False**]: Only return downtimes that are active when the request is made.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiGetAllDowntimesRequest
*/
func (a *DowntimesApiService) GetAllDowntimes(ctx _context.Context) apiGetAllDowntimesRequest {
	return apiGetAllDowntimesRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return []Downtime
*/
func (r apiGetAllDowntimesRequest) Execute() ([]Downtime, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Downtime
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "DowntimesApiService.GetAllDowntimes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/downtime"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.currentOnly != nil {
		localVarQueryParams.Add("current_only", parameterToString(*r.currentOnly, ""))
	}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["apiKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["appKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
		if localVarHTTPResponse.StatusCode == 200 {
			var v []Downtime
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetDowntimeRequest struct {
	ctx        _context.Context
	apiService *DowntimesApiService
	downtimeId int64
}

/*
GetDowntime Get a downtime
### Overview
Get Downtime Detail by downtime_id
### Arguments
This endpoint takes no JSON arguments."
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param downtimeId ID of the downtime to fetch
@return apiGetDowntimeRequest
*/
func (a *DowntimesApiService) GetDowntime(ctx _context.Context, downtimeId int64) apiGetDowntimeRequest {
	return apiGetDowntimeRequest{
		apiService: a,
		ctx:        ctx,
		downtimeId: downtimeId,
	}
}

/*
Execute executes the request
 @return Downtime
*/
func (r apiGetDowntimeRequest) Execute() (Downtime, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Downtime
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "DowntimesApiService.GetDowntime")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/downtime/{downtime_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"downtime_id"+"}", _neturl.QueryEscape(parameterToString(r.downtimeId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["apiKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["appKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
		if localVarHTTPResponse.StatusCode == 200 {
			var v Downtime
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiUpdateDowntimeRequest struct {
	ctx        _context.Context
	apiService *DowntimesApiService
	downtimeId int64
	body       *Downtime
}

func (r apiUpdateDowntimeRequest) Body(body Downtime) apiUpdateDowntimeRequest {
	r.body = &body
	return r
}

/*
UpdateDowntime Update a downtime
### Overview
Update a single Downtime by downtime_id.
### Arguments
* **`id`** [*required*]: The integer id of the downtime to be updated
* **`scope`** [*required*]: The scope to which the downtime applies, e.g. 'host:app2'.
  Provide multiple scopes as a comma-separated list, e.g. 'env:dev,env:prod'.
  The resulting downtime applies to sources that matches ALL provided scopes
  (i.e. env:dev AND env:prod), NOT any of them.

* **`monitor_tags`** [*optional*, *default*=**no monitor tag filter**]: A comma-separated
  list of monitor tags, i.e. tags that are applied directly to monitors, *not* tags that
  are used in monitor queries (which are filtered by the `scope` parameter), to which
  the downtime applies. The resulting downtime applies to monitors that match ALL provided
  monitor tags (i.e. `service:postgres` **AND** `team:frontend`), NOT any of them.

* **`monitor_id`** [*optional*, *default*=**None**]: A single monitor to which the downtime
  applies. If not provided, the downtime applies to all monitors.

* **`start`** [*optional*, *default* = **original start**]: POSIX timestamp to start
  the downtime.

* **`end`** [*optional*, *default* = **original end**]: POSIX timestamp to end the downtime.
  If not provided, the downtime is in effect indefinitely (i.e. until you cancel it).

* **`message`** [*required*, *default* = **original message**]: A message to include with
  notifications for this downtime. Email notifications can be sent to specific users by
  using the same '@username' notation as events

* **`timezone`** [*optional*, default = **original timezone** ]: The timezone for the downtime.
* **`recurrence`** [*optional*, *default* = **original recurrence**]: An object defining the
  recurrence of the downtime with a variety of parameters:

    * **`type`** the type of recurrence. Choose from: `days`, `weeks`, `months`, `years`.

    * **`period`** how often to repeat as an integer. For example to repeat every 3 days,
      select a type of `days` and a period of `3`.

    * **`week_days`** (optional) a list of week days to repeat on. Choose from: `Mon`, `Tue`,
      `Wed`, `Thu`, `Fri`, `Sat` or `Sun`. Only applicable when `type` is `weeks`.
      **First letter must be capitalized.**

    * **`until_occurrences`** (optional) how many times the downtime is rescheduled.
      **`until_occurrences` and `until_date`** are mutually exclusive

    * **`until_date`** (optional) the date at which the recurrence should end as a POSIX
      timestmap. **`until_occurrences` and `until_date`** are mutually exclusive
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param downtimeId ID of the downtime to update
@return apiUpdateDowntimeRequest
*/
func (a *DowntimesApiService) UpdateDowntime(ctx _context.Context, downtimeId int64) apiUpdateDowntimeRequest {
	return apiUpdateDowntimeRequest{
		apiService: a,
		ctx:        ctx,
		downtimeId: downtimeId,
	}
}

/*
Execute executes the request
 @return Downtime
*/
func (r apiUpdateDowntimeRequest) Execute() (Downtime, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Downtime
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "DowntimesApiService.UpdateDowntime")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/downtime/{downtime_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"downtime_id"+"}", _neturl.QueryEscape(parameterToString(r.downtimeId, "")), -1)

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
			if auth, ok := auth["apiKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["appKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
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
		if localVarHTTPResponse.StatusCode == 200 {
			var v Downtime
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
