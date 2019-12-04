/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// SnapshotsApiService SnapshotsApi service
type SnapshotsApiService service

// GetGraphSnapshotOpts Optional parameters for the method 'GetGraphSnapshot'
type GetGraphSnapshotOpts struct {
	EventQuery optional.String
	GraphDef   optional.String
	Title      optional.String
}

/*
GetGraphSnapshot Take graph snapshots
### Overview Take graph snapshots ### Arguments * **&#x60;metric_query&#x60;** [*required*]: The metric query. * **&#x60;start&#x60;** [*required*]: The POSIX timestamp of the start of the query. * **&#x60;end&#x60;** [*required*]: The POSIX timestamp of the end of the query. * **&#x60;event_query&#x60;** [*optional*, *default* &#x3D; **None**]: A query that adds event bands to the graph. * **&#x60;graph_def&#x60;** [*optional*, *default* &#x3D; **None**]: A JSON document defining the graph.   graph_def can be used instead of metric_query. The JSON document uses the   [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar)   and should be formatted to a single line then URLEncoded.  * **&#x60;title&#x60;** [*optional*, *default* &#x3D; **None**]: A title for the graph.   If no title is specified, the graph doesn’t have a title.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param metricQuery The metric query.
 * @param start The POSIX timestamp of the start of the query.
 * @param end The POSIX timestamp of the end of the query.
 * @param optional nil or *GetGraphSnapshotOpts - Optional Parameters:
 * @param "EventQuery" (optional.String) -  A query that adds event bands to the graph.
 * @param "GraphDef" (optional.String) -  A JSON document defining the graph. graph_def can be used instead of metric_query. The JSON document uses the [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar) and should be formatted to a single line then URLEncoded.
 * @param "Title" (optional.String) -  A title for the graph. If no title is specified, the graph doesn’t have a title.
@return GraphSnapshot
*/
func (a *SnapshotsApiService) GetGraphSnapshot(ctx _context.Context, metricQuery string, start int64, end int64, localVarOptionals *GetGraphSnapshotOpts) (GraphSnapshot, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GraphSnapshot
	)

	localBasePath, err := a.client.cfg.ServerUrlWithContext(ctx, "SnapshotsApiService.GetGraphSnapshot")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/graph/snapshot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("metric_query", parameterToString(metricQuery, ""))
	localVarQueryParams.Add("start", parameterToString(start, ""))
	localVarQueryParams.Add("end", parameterToString(end, ""))
	if localVarOptionals != nil && localVarOptionals.EventQuery.IsSet() {
		localVarQueryParams.Add("event_query", parameterToString(localVarOptionals.EventQuery.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GraphDef.IsSet() {
		localVarQueryParams.Add("graph_def", parameterToString(localVarOptionals.GraphDef.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Title.IsSet() {
		localVarQueryParams.Add("title", parameterToString(localVarOptionals.Title.Value(), ""))
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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["api_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["application_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("application_key", key)
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
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
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error400
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
