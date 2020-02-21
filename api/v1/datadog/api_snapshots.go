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
)

// Linger please
var (
	_ _context.Context
)

// SnapshotsApiService SnapshotsApi service
type SnapshotsApiService service

type apiGetGraphSnapshotRequest struct {
	ctx         _context.Context
	apiService  *SnapshotsApiService
	metricQuery *string
	start       *int64
	end         *int64
	eventQuery  *string
	graphDef    *string
	title       *string
}

func (r apiGetGraphSnapshotRequest) MetricQuery(metricQuery string) apiGetGraphSnapshotRequest {
	r.metricQuery = &metricQuery
	return r
}

func (r apiGetGraphSnapshotRequest) Start(start int64) apiGetGraphSnapshotRequest {
	r.start = &start
	return r
}

func (r apiGetGraphSnapshotRequest) End(end int64) apiGetGraphSnapshotRequest {
	r.end = &end
	return r
}

func (r apiGetGraphSnapshotRequest) EventQuery(eventQuery string) apiGetGraphSnapshotRequest {
	r.eventQuery = &eventQuery
	return r
}

func (r apiGetGraphSnapshotRequest) GraphDef(graphDef string) apiGetGraphSnapshotRequest {
	r.graphDef = &graphDef
	return r
}

func (r apiGetGraphSnapshotRequest) Title(title string) apiGetGraphSnapshotRequest {
	r.title = &title
	return r
}

/*
GetGraphSnapshot Take graph snapshots
### Overview
Take graph snapshots
### Arguments
* **`metric_query`** [*required*]: The metric query.
* **`start`** [*required*]: The POSIX timestamp of the start of the query.
* **`end`** [*required*]: The POSIX timestamp of the end of the query.
* **`event_query`** [*optional*, *default* = **None**]: A query that adds event bands to the graph.
* **`graph_def`** [*optional*, *default* = **None**]: A JSON document defining the graph.
  graph_def can be used instead of metric_query. The JSON document uses the
  [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar)
  and should be formatted to a single line then URLEncoded.

* **`title`** [*optional*, *default* = **None**]: A title for the graph.
  If no title is specified, the graph doesn’t have a title.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiGetGraphSnapshotRequest
*/
func (a *SnapshotsApiService) GetGraphSnapshot(ctx _context.Context) apiGetGraphSnapshotRequest {
	return apiGetGraphSnapshotRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return GraphSnapshot
*/
func (r apiGetGraphSnapshotRequest) Execute() (GraphSnapshot, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GraphSnapshot
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.GetGraphSnapshot")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/graph/snapshot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.metricQuery == nil {
		return localVarReturnValue, nil, reportError("metricQuery is required and must be specified")
	}

	if r.start == nil {
		return localVarReturnValue, nil, reportError("start is required and must be specified")
	}

	if r.end == nil {
		return localVarReturnValue, nil, reportError("end is required and must be specified")
	}

	localVarQueryParams.Add("metric_query", parameterToString(*r.metricQuery, ""))
	localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	localVarQueryParams.Add("end", parameterToString(*r.end, ""))
	if r.eventQuery != nil {
		localVarQueryParams.Add("event_query", parameterToString(*r.eventQuery, ""))
	}
	if r.graphDef != nil {
		localVarQueryParams.Add("graph_def", parameterToString(*r.graphDef, ""))
	}
	if r.title != nil {
		localVarQueryParams.Add("title", parameterToString(*r.title, ""))
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
			var v GraphSnapshot
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
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
