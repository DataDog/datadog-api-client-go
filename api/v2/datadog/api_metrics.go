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

// MetricsApiService MetricsApi service.
type MetricsApiService service

type apiCreateBulkTagsMetricsConfigurationRequest struct {
	ctx        _context.Context
	ApiService *MetricsApiService
	body       *MetricBulkTagConfigCreateRequest
}

func (a *MetricsApiService) buildCreateBulkTagsMetricsConfigurationRequest(ctx _context.Context, body MetricBulkTagConfigCreateRequest) (apiCreateBulkTagsMetricsConfigurationRequest, error) {
	req := apiCreateBulkTagsMetricsConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
		body:       &body,
	}
	return req, nil
}

// CreateBulkTagsMetricsConfiguration Configure tags for multiple metrics.
// Create and define a list of queryable tag keys for a set of existing count, gauge, rate, and distribution metrics.
// Metrics are selected by passing a metric name prefix. Use the Delete method of this API path to remove tag configurations.
// Results can be sent to a set of account email addresses, just like the same operation in the Datadog web app.
// If multiple calls include the same metric, the last configuration applied (not by submit order) is used, do not
// expect deterministic ordering of concurrent calls.
// Can only be used with application keys of users with the `Manage Tags for Metrics` permission.
func (a *MetricsApiService) CreateBulkTagsMetricsConfiguration(ctx _context.Context, body MetricBulkTagConfigCreateRequest) (MetricBulkTagConfigResponse, *_nethttp.Response, error) {
	req, err := a.buildCreateBulkTagsMetricsConfigurationRequest(ctx, body)
	if err != nil {
		var localVarReturnValue MetricBulkTagConfigResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.createBulkTagsMetricsConfigurationExecute(req)
}

// createBulkTagsMetricsConfigurationExecute executes the request.
func (a *MetricsApiService) createBulkTagsMetricsConfigurationExecute(r apiCreateBulkTagsMetricsConfigurationRequest) (MetricBulkTagConfigResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue MetricBulkTagConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.CreateBulkTagsMetricsConfiguration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics/config/bulk-tags"

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

type apiCreateTagConfigurationRequest struct {
	ctx        _context.Context
	ApiService *MetricsApiService
	metricName string
	body       *MetricTagConfigurationCreateRequest
}

func (a *MetricsApiService) buildCreateTagConfigurationRequest(ctx _context.Context, metricName string, body MetricTagConfigurationCreateRequest) (apiCreateTagConfigurationRequest, error) {
	req := apiCreateTagConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
		metricName: metricName,
		body:       &body,
	}
	return req, nil
}

// CreateTagConfiguration Create a tag configuration.
// Create and define a list of queryable tag keys for an existing count/gauge/rate/distribution metric.
// Optionally, include percentile aggregations on any distribution metric or configure custom aggregations
// on any count, rate, or gauge metric.
// Can only be used with application keys of users with the `Manage Tags for Metrics` permission.
func (a *MetricsApiService) CreateTagConfiguration(ctx _context.Context, metricName string, body MetricTagConfigurationCreateRequest) (MetricTagConfigurationResponse, *_nethttp.Response, error) {
	req, err := a.buildCreateTagConfigurationRequest(ctx, metricName, body)
	if err != nil {
		var localVarReturnValue MetricTagConfigurationResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.createTagConfigurationExecute(req)
}

// createTagConfigurationExecute executes the request.
func (a *MetricsApiService) createTagConfigurationExecute(r apiCreateTagConfigurationRequest) (MetricTagConfigurationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue MetricTagConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.CreateTagConfiguration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics/{metric_name}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"metric_name"+"}", _neturl.PathEscape(parameterToString(r.metricName, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 409 {
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

type apiDeleteBulkTagsMetricsConfigurationRequest struct {
	ctx        _context.Context
	ApiService *MetricsApiService
	body       *MetricBulkTagConfigDeleteRequest
}

func (a *MetricsApiService) buildDeleteBulkTagsMetricsConfigurationRequest(ctx _context.Context, body MetricBulkTagConfigDeleteRequest) (apiDeleteBulkTagsMetricsConfigurationRequest, error) {
	req := apiDeleteBulkTagsMetricsConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
		body:       &body,
	}
	return req, nil
}

// DeleteBulkTagsMetricsConfiguration Configure tags for multiple metrics.
// Delete all custom lists of queryable tag keys for a set of existing count, gauge, rate, and distribution metrics.
// Metrics are selected by passing a metric name prefix.
// Results can be sent to a set of account email addresses, just like the same operation in the Datadog web app.
// Can only be used with application keys of users with the `Manage Tags for Metrics` permission.
func (a *MetricsApiService) DeleteBulkTagsMetricsConfiguration(ctx _context.Context, body MetricBulkTagConfigDeleteRequest) (MetricBulkTagConfigResponse, *_nethttp.Response, error) {
	req, err := a.buildDeleteBulkTagsMetricsConfigurationRequest(ctx, body)
	if err != nil {
		var localVarReturnValue MetricBulkTagConfigResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.deleteBulkTagsMetricsConfigurationExecute(req)
}

// deleteBulkTagsMetricsConfigurationExecute executes the request.
func (a *MetricsApiService) deleteBulkTagsMetricsConfigurationExecute(r apiDeleteBulkTagsMetricsConfigurationRequest) (MetricBulkTagConfigResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodDelete
		localVarPostBody    interface{}
		localVarReturnValue MetricBulkTagConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.DeleteBulkTagsMetricsConfiguration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics/config/bulk-tags"

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

type apiDeleteTagConfigurationRequest struct {
	ctx        _context.Context
	ApiService *MetricsApiService
	metricName string
}

func (a *MetricsApiService) buildDeleteTagConfigurationRequest(ctx _context.Context, metricName string) (apiDeleteTagConfigurationRequest, error) {
	req := apiDeleteTagConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
		metricName: metricName,
	}
	return req, nil
}

// DeleteTagConfiguration Delete a tag configuration.
// Deletes a metric's tag configuration. Can only be used with application
// keys from users with the `Manage Tags for Metrics` permission.
func (a *MetricsApiService) DeleteTagConfiguration(ctx _context.Context, metricName string) (*_nethttp.Response, error) {
	req, err := a.buildDeleteTagConfigurationRequest(ctx, metricName)
	if err != nil {
		return nil, err
	}

	return req.ApiService.deleteTagConfigurationExecute(req)
}

// deleteTagConfigurationExecute executes the request.
func (a *MetricsApiService) deleteTagConfigurationExecute(r apiDeleteTagConfigurationRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.DeleteTagConfiguration")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics/{metric_name}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"metric_name"+"}", _neturl.PathEscape(parameterToString(r.metricName, "")), -1)

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

type apiEstimateMetricsOutputSeriesRequest struct {
	ctx                   _context.Context
	ApiService            *MetricsApiService
	metricName            string
	filterGroups          *string
	filterHoursAgo        *int32
	filterNumAggregations *int32
	filterPct             *bool
	filterTimespanH       *int32
}

// EstimateMetricsOutputSeriesOptionalParameters holds optional parameters for EstimateMetricsOutputSeries.
type EstimateMetricsOutputSeriesOptionalParameters struct {
	FilterGroups          *string
	FilterHoursAgo        *int32
	FilterNumAggregations *int32
	FilterPct             *bool
	FilterTimespanH       *int32
}

// NewEstimateMetricsOutputSeriesOptionalParameters creates an empty struct for parameters.
func NewEstimateMetricsOutputSeriesOptionalParameters() *EstimateMetricsOutputSeriesOptionalParameters {
	this := EstimateMetricsOutputSeriesOptionalParameters{}
	return &this
}

// WithFilterGroups sets the corresponding parameter name and returns the struct.
func (r *EstimateMetricsOutputSeriesOptionalParameters) WithFilterGroups(filterGroups string) *EstimateMetricsOutputSeriesOptionalParameters {
	r.FilterGroups = &filterGroups
	return r
}

// WithFilterHoursAgo sets the corresponding parameter name and returns the struct.
func (r *EstimateMetricsOutputSeriesOptionalParameters) WithFilterHoursAgo(filterHoursAgo int32) *EstimateMetricsOutputSeriesOptionalParameters {
	r.FilterHoursAgo = &filterHoursAgo
	return r
}

// WithFilterNumAggregations sets the corresponding parameter name and returns the struct.
func (r *EstimateMetricsOutputSeriesOptionalParameters) WithFilterNumAggregations(filterNumAggregations int32) *EstimateMetricsOutputSeriesOptionalParameters {
	r.FilterNumAggregations = &filterNumAggregations
	return r
}

// WithFilterPct sets the corresponding parameter name and returns the struct.
func (r *EstimateMetricsOutputSeriesOptionalParameters) WithFilterPct(filterPct bool) *EstimateMetricsOutputSeriesOptionalParameters {
	r.FilterPct = &filterPct
	return r
}

// WithFilterTimespanH sets the corresponding parameter name and returns the struct.
func (r *EstimateMetricsOutputSeriesOptionalParameters) WithFilterTimespanH(filterTimespanH int32) *EstimateMetricsOutputSeriesOptionalParameters {
	r.FilterTimespanH = &filterTimespanH
	return r
}

func (a *MetricsApiService) buildEstimateMetricsOutputSeriesRequest(ctx _context.Context, metricName string, o ...EstimateMetricsOutputSeriesOptionalParameters) (apiEstimateMetricsOutputSeriesRequest, error) {
	req := apiEstimateMetricsOutputSeriesRequest{
		ApiService: a,
		ctx:        ctx,
		metricName: metricName,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type EstimateMetricsOutputSeriesOptionalParameters is allowed")
	}

	if o != nil {
		req.filterGroups = o[0].FilterGroups
		req.filterHoursAgo = o[0].FilterHoursAgo
		req.filterNumAggregations = o[0].FilterNumAggregations
		req.filterPct = o[0].FilterPct
		req.filterTimespanH = o[0].FilterTimespanH
	}
	return req, nil
}

// EstimateMetricsOutputSeries Tag Configuration Cardinality Estimator.
// Returns the estimated cardinality for a metric with a given tag, percentile and number of aggregations configuration using Metrics without Limits&trade;.
func (a *MetricsApiService) EstimateMetricsOutputSeries(ctx _context.Context, metricName string, o ...EstimateMetricsOutputSeriesOptionalParameters) (MetricEstimateResponse, *_nethttp.Response, error) {
	req, err := a.buildEstimateMetricsOutputSeriesRequest(ctx, metricName, o...)
	if err != nil {
		var localVarReturnValue MetricEstimateResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.estimateMetricsOutputSeriesExecute(req)
}

// estimateMetricsOutputSeriesExecute executes the request.
func (a *MetricsApiService) estimateMetricsOutputSeriesExecute(r apiEstimateMetricsOutputSeriesRequest) (MetricEstimateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MetricEstimateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.EstimateMetricsOutputSeries")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics/{metric_name}/estimate"
	localVarPath = strings.Replace(localVarPath, "{"+"metric_name"+"}", _neturl.PathEscape(parameterToString(r.metricName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.filterGroups != nil {
		localVarQueryParams.Add("filter[groups]", parameterToString(*r.filterGroups, ""))
	}
	if r.filterHoursAgo != nil {
		localVarQueryParams.Add("filter[hours_ago]", parameterToString(*r.filterHoursAgo, ""))
	}
	if r.filterNumAggregations != nil {
		localVarQueryParams.Add("filter[num_aggregations]", parameterToString(*r.filterNumAggregations, ""))
	}
	if r.filterPct != nil {
		localVarQueryParams.Add("filter[pct]", parameterToString(*r.filterPct, ""))
	}
	if r.filterTimespanH != nil {
		localVarQueryParams.Add("filter[timespan_h]", parameterToString(*r.filterTimespanH, ""))
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

type apiListTagConfigurationByNameRequest struct {
	ctx        _context.Context
	ApiService *MetricsApiService
	metricName string
}

func (a *MetricsApiService) buildListTagConfigurationByNameRequest(ctx _context.Context, metricName string) (apiListTagConfigurationByNameRequest, error) {
	req := apiListTagConfigurationByNameRequest{
		ApiService: a,
		ctx:        ctx,
		metricName: metricName,
	}
	return req, nil
}

// ListTagConfigurationByName List tag configuration by name.
// Returns the tag configuration for the given metric name.
func (a *MetricsApiService) ListTagConfigurationByName(ctx _context.Context, metricName string) (MetricTagConfigurationResponse, *_nethttp.Response, error) {
	req, err := a.buildListTagConfigurationByNameRequest(ctx, metricName)
	if err != nil {
		var localVarReturnValue MetricTagConfigurationResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listTagConfigurationByNameExecute(req)
}

// listTagConfigurationByNameExecute executes the request.
func (a *MetricsApiService) listTagConfigurationByNameExecute(r apiListTagConfigurationByNameRequest) (MetricTagConfigurationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MetricTagConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.ListTagConfigurationByName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics/{metric_name}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"metric_name"+"}", _neturl.PathEscape(parameterToString(r.metricName, "")), -1)

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

type apiListTagConfigurationsRequest struct {
	ctx                      _context.Context
	ApiService               *MetricsApiService
	filterConfigured         *bool
	filterTagsConfigured     *string
	filterMetricType         *MetricTagConfigurationMetricTypes
	filterIncludePercentiles *bool
	filterTags               *string
	windowSeconds            *int64
}

// ListTagConfigurationsOptionalParameters holds optional parameters for ListTagConfigurations.
type ListTagConfigurationsOptionalParameters struct {
	FilterConfigured         *bool
	FilterTagsConfigured     *string
	FilterMetricType         *MetricTagConfigurationMetricTypes
	FilterIncludePercentiles *bool
	FilterTags               *string
	WindowSeconds            *int64
}

// NewListTagConfigurationsOptionalParameters creates an empty struct for parameters.
func NewListTagConfigurationsOptionalParameters() *ListTagConfigurationsOptionalParameters {
	this := ListTagConfigurationsOptionalParameters{}
	return &this
}

// WithFilterConfigured sets the corresponding parameter name and returns the struct.
func (r *ListTagConfigurationsOptionalParameters) WithFilterConfigured(filterConfigured bool) *ListTagConfigurationsOptionalParameters {
	r.FilterConfigured = &filterConfigured
	return r
}

// WithFilterTagsConfigured sets the corresponding parameter name and returns the struct.
func (r *ListTagConfigurationsOptionalParameters) WithFilterTagsConfigured(filterTagsConfigured string) *ListTagConfigurationsOptionalParameters {
	r.FilterTagsConfigured = &filterTagsConfigured
	return r
}

// WithFilterMetricType sets the corresponding parameter name and returns the struct.
func (r *ListTagConfigurationsOptionalParameters) WithFilterMetricType(filterMetricType MetricTagConfigurationMetricTypes) *ListTagConfigurationsOptionalParameters {
	r.FilterMetricType = &filterMetricType
	return r
}

// WithFilterIncludePercentiles sets the corresponding parameter name and returns the struct.
func (r *ListTagConfigurationsOptionalParameters) WithFilterIncludePercentiles(filterIncludePercentiles bool) *ListTagConfigurationsOptionalParameters {
	r.FilterIncludePercentiles = &filterIncludePercentiles
	return r
}

// WithFilterTags sets the corresponding parameter name and returns the struct.
func (r *ListTagConfigurationsOptionalParameters) WithFilterTags(filterTags string) *ListTagConfigurationsOptionalParameters {
	r.FilterTags = &filterTags
	return r
}

// WithWindowSeconds sets the corresponding parameter name and returns the struct.
func (r *ListTagConfigurationsOptionalParameters) WithWindowSeconds(windowSeconds int64) *ListTagConfigurationsOptionalParameters {
	r.WindowSeconds = &windowSeconds
	return r
}

func (a *MetricsApiService) buildListTagConfigurationsRequest(ctx _context.Context, o ...ListTagConfigurationsOptionalParameters) (apiListTagConfigurationsRequest, error) {
	req := apiListTagConfigurationsRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type ListTagConfigurationsOptionalParameters is allowed")
	}

	if o != nil {
		req.filterConfigured = o[0].FilterConfigured
		req.filterTagsConfigured = o[0].FilterTagsConfigured
		req.filterMetricType = o[0].FilterMetricType
		req.filterIncludePercentiles = o[0].FilterIncludePercentiles
		req.filterTags = o[0].FilterTags
		req.windowSeconds = o[0].WindowSeconds
	}
	return req, nil
}

// ListTagConfigurations List tag configurations.
// Returns all configured count/gauge/rate/distribution metric names
// (with additional filters if specified).
func (a *MetricsApiService) ListTagConfigurations(ctx _context.Context, o ...ListTagConfigurationsOptionalParameters) (MetricsAndMetricTagConfigurationsResponse, *_nethttp.Response, error) {
	req, err := a.buildListTagConfigurationsRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue MetricsAndMetricTagConfigurationsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listTagConfigurationsExecute(req)
}

// listTagConfigurationsExecute executes the request.
func (a *MetricsApiService) listTagConfigurationsExecute(r apiListTagConfigurationsRequest) (MetricsAndMetricTagConfigurationsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MetricsAndMetricTagConfigurationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.ListTagConfigurations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.filterConfigured != nil {
		localVarQueryParams.Add("filter[configured]", parameterToString(*r.filterConfigured, ""))
	}
	if r.filterTagsConfigured != nil {
		localVarQueryParams.Add("filter[tags_configured]", parameterToString(*r.filterTagsConfigured, ""))
	}
	if r.filterMetricType != nil {
		localVarQueryParams.Add("filter[metric_type]", parameterToString(*r.filterMetricType, ""))
	}
	if r.filterIncludePercentiles != nil {
		localVarQueryParams.Add("filter[include_percentiles]", parameterToString(*r.filterIncludePercentiles, ""))
	}
	if r.filterTags != nil {
		localVarQueryParams.Add("filter[tags]", parameterToString(*r.filterTags, ""))
	}
	if r.windowSeconds != nil {
		localVarQueryParams.Add("window[seconds]", parameterToString(*r.windowSeconds, ""))
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

type apiListTagsByMetricNameRequest struct {
	ctx        _context.Context
	ApiService *MetricsApiService
	metricName string
}

func (a *MetricsApiService) buildListTagsByMetricNameRequest(ctx _context.Context, metricName string) (apiListTagsByMetricNameRequest, error) {
	req := apiListTagsByMetricNameRequest{
		ApiService: a,
		ctx:        ctx,
		metricName: metricName,
	}
	return req, nil
}

// ListTagsByMetricName List tags by metric name.
// View indexed tag key-value pairs for a given metric name.
func (a *MetricsApiService) ListTagsByMetricName(ctx _context.Context, metricName string) (MetricAllTagsResponse, *_nethttp.Response, error) {
	req, err := a.buildListTagsByMetricNameRequest(ctx, metricName)
	if err != nil {
		var localVarReturnValue MetricAllTagsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listTagsByMetricNameExecute(req)
}

// listTagsByMetricNameExecute executes the request.
func (a *MetricsApiService) listTagsByMetricNameExecute(r apiListTagsByMetricNameRequest) (MetricAllTagsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MetricAllTagsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.ListTagsByMetricName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics/{metric_name}/all-tags"
	localVarPath = strings.Replace(localVarPath, "{"+"metric_name"+"}", _neturl.PathEscape(parameterToString(r.metricName, "")), -1)

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

type apiListVolumesByMetricNameRequest struct {
	ctx        _context.Context
	ApiService *MetricsApiService
	metricName string
}

func (a *MetricsApiService) buildListVolumesByMetricNameRequest(ctx _context.Context, metricName string) (apiListVolumesByMetricNameRequest, error) {
	req := apiListVolumesByMetricNameRequest{
		ApiService: a,
		ctx:        ctx,
		metricName: metricName,
	}
	return req, nil
}

// ListVolumesByMetricName List distinct metric volumes by metric name.
// View distinct metrics volumes for the given metric name.
//
// Custom metrics generated in-app from other products will return `null` for ingested volumes.
func (a *MetricsApiService) ListVolumesByMetricName(ctx _context.Context, metricName string) (MetricVolumesResponse, *_nethttp.Response, error) {
	req, err := a.buildListVolumesByMetricNameRequest(ctx, metricName)
	if err != nil {
		var localVarReturnValue MetricVolumesResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listVolumesByMetricNameExecute(req)
}

// listVolumesByMetricNameExecute executes the request.
func (a *MetricsApiService) listVolumesByMetricNameExecute(r apiListVolumesByMetricNameRequest) (MetricVolumesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MetricVolumesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.ListVolumesByMetricName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics/{metric_name}/volumes"
	localVarPath = strings.Replace(localVarPath, "{"+"metric_name"+"}", _neturl.PathEscape(parameterToString(r.metricName, "")), -1)

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

type apiSubmitMetricsRequest struct {
	ctx             _context.Context
	ApiService      *MetricsApiService
	body            *MetricPayload
	contentEncoding *MetricContentEncoding
}

// SubmitMetricsOptionalParameters holds optional parameters for SubmitMetrics.
type SubmitMetricsOptionalParameters struct {
	ContentEncoding *MetricContentEncoding
}

// NewSubmitMetricsOptionalParameters creates an empty struct for parameters.
func NewSubmitMetricsOptionalParameters() *SubmitMetricsOptionalParameters {
	this := SubmitMetricsOptionalParameters{}
	return &this
}

// WithContentEncoding sets the corresponding parameter name and returns the struct.
func (r *SubmitMetricsOptionalParameters) WithContentEncoding(contentEncoding MetricContentEncoding) *SubmitMetricsOptionalParameters {
	r.ContentEncoding = &contentEncoding
	return r
}

func (a *MetricsApiService) buildSubmitMetricsRequest(ctx _context.Context, body MetricPayload, o ...SubmitMetricsOptionalParameters) (apiSubmitMetricsRequest, error) {
	req := apiSubmitMetricsRequest{
		ApiService: a,
		ctx:        ctx,
		body:       &body,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type SubmitMetricsOptionalParameters is allowed")
	}

	if o != nil {
		req.contentEncoding = o[0].ContentEncoding
	}
	return req, nil
}

// SubmitMetrics Submit metrics.
// The metrics end-point allows you to post time-series data that can be graphed on Datadog’s dashboards.
// The maximum payload size is 500 kilobytes (512000 bytes). Compressed payloads must have a decompressed size of less than 5 megabytes (5242880 bytes).
//
// If you’re submitting metrics directly to the Datadog API without using DogStatsD, expect:
//
// - 64 bits for the timestamp
// - 64 bits for the value
// - 20 bytes for the metric names
// - 50 bytes for the timeseries
// - The full payload is approximately 100 bytes.
func (a *MetricsApiService) SubmitMetrics(ctx _context.Context, body MetricPayload, o ...SubmitMetricsOptionalParameters) (IntakePayloadAccepted, *_nethttp.Response, error) {
	req, err := a.buildSubmitMetricsRequest(ctx, body, o...)
	if err != nil {
		var localVarReturnValue IntakePayloadAccepted
		return localVarReturnValue, nil, err
	}

	return req.ApiService.submitMetricsExecute(req)
}

// submitMetricsExecute executes the request.
func (a *MetricsApiService) submitMetricsExecute(r apiSubmitMetricsRequest) (IntakePayloadAccepted, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue IntakePayloadAccepted
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.SubmitMetrics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/series"

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

	if r.contentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = parameterToString(*r.contentEncoding, "")
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
		if localVarHTTPResponse.StatusCode == 408 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
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

type apiUpdateTagConfigurationRequest struct {
	ctx        _context.Context
	ApiService *MetricsApiService
	metricName string
	body       *MetricTagConfigurationUpdateRequest
}

func (a *MetricsApiService) buildUpdateTagConfigurationRequest(ctx _context.Context, metricName string, body MetricTagConfigurationUpdateRequest) (apiUpdateTagConfigurationRequest, error) {
	req := apiUpdateTagConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
		metricName: metricName,
		body:       &body,
	}
	return req, nil
}

// UpdateTagConfiguration Update a tag configuration.
// Update the tag configuration of a metric or percentile aggregations of a distribution metric or custom aggregations
// of a count, rate, or gauge metric.
// Can only be used with application keys from users with the `Manage Tags for Metrics` permission.
func (a *MetricsApiService) UpdateTagConfiguration(ctx _context.Context, metricName string, body MetricTagConfigurationUpdateRequest) (MetricTagConfigurationResponse, *_nethttp.Response, error) {
	req, err := a.buildUpdateTagConfigurationRequest(ctx, metricName, body)
	if err != nil {
		var localVarReturnValue MetricTagConfigurationResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.updateTagConfigurationExecute(req)
}

// updateTagConfigurationExecute executes the request.
func (a *MetricsApiService) updateTagConfigurationExecute(r apiUpdateTagConfigurationRequest) (MetricTagConfigurationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue MetricTagConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.UpdateTagConfiguration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/metrics/{metric_name}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"metric_name"+"}", _neturl.PathEscape(parameterToString(r.metricName, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 422 {
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
