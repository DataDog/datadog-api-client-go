// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"bytes"
	_context "context"
	_fmt "fmt"
	_ioutil "io/ioutil"
	_log "log"
	_nethttp "net/http"
	_neturl "net/url"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageMeteringApi service type
type UsageMeteringApi datadog.Service

type apiGetCostByOrgRequest struct {
	ctx        _context.Context
	startMonth *time.Time
	endMonth   *time.Time
}

// GetCostByOrgOptionalParameters holds optional parameters for GetCostByOrg.
type GetCostByOrgOptionalParameters struct {
	EndMonth *time.Time
}

// NewGetCostByOrgOptionalParameters creates an empty struct for parameters.
func NewGetCostByOrgOptionalParameters() *GetCostByOrgOptionalParameters {
	this := GetCostByOrgOptionalParameters{}
	return &this
}

// WithEndMonth sets the corresponding parameter name and returns the struct.
func (r *GetCostByOrgOptionalParameters) WithEndMonth(endMonth time.Time) *GetCostByOrgOptionalParameters {
	r.EndMonth = &endMonth
	return r
}

func (a *UsageMeteringApi) buildGetCostByOrgRequest(ctx _context.Context, startMonth time.Time, o ...GetCostByOrgOptionalParameters) (apiGetCostByOrgRequest, error) {
	req := apiGetCostByOrgRequest{
		ctx:        ctx,
		startMonth: &startMonth,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type GetCostByOrgOptionalParameters is allowed")
	}

	if o != nil {
		req.endMonth = o[0].EndMonth
	}
	return req, nil
}

// GetCostByOrg Get cost across multi-org account.
// Get cost across multi-org account. Cost by org data for a given month becomes available no later than the 16th of the following month.
func (a *UsageMeteringApi) GetCostByOrg(ctx _context.Context, startMonth time.Time, o ...GetCostByOrgOptionalParameters) (CostByOrgResponse, *_nethttp.Response, error) {
	req, err := a.buildGetCostByOrgRequest(ctx, startMonth, o...)
	if err != nil {
		var localVarReturnValue CostByOrgResponse
		return localVarReturnValue, nil, err
	}

	return a.getCostByOrgExecute(req)
}

// getCostByOrgExecute executes the request.
func (a *UsageMeteringApi) getCostByOrgExecute(r apiGetCostByOrgRequest) (CostByOrgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CostByOrgResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApi.GetCostByOrg")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/cost_by_org"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startMonth == nil {
		return localVarReturnValue, nil, datadog.ReportError("startMonth is required and must be specified")
	}
	localVarQueryParams.Add("start_month", datadog.ParameterToString(*r.startMonth, ""))
	if r.endMonth != nil {
		localVarQueryParams.Add("end_month", datadog.ParameterToString(*r.endMonth, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		newErr := datadog.GenericOpenAPIError{
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
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetEstimatedCostByOrgRequest struct {
	ctx        _context.Context
	view       *string
	startMonth *time.Time
	endMonth   *time.Time
	startDate  *time.Time
	endDate    *time.Time
}

// GetEstimatedCostByOrgOptionalParameters holds optional parameters for GetEstimatedCostByOrg.
type GetEstimatedCostByOrgOptionalParameters struct {
	StartMonth *time.Time
	EndMonth   *time.Time
	StartDate  *time.Time
	EndDate    *time.Time
}

// NewGetEstimatedCostByOrgOptionalParameters creates an empty struct for parameters.
func NewGetEstimatedCostByOrgOptionalParameters() *GetEstimatedCostByOrgOptionalParameters {
	this := GetEstimatedCostByOrgOptionalParameters{}
	return &this
}

// WithStartMonth sets the corresponding parameter name and returns the struct.
func (r *GetEstimatedCostByOrgOptionalParameters) WithStartMonth(startMonth time.Time) *GetEstimatedCostByOrgOptionalParameters {
	r.StartMonth = &startMonth
	return r
}

// WithEndMonth sets the corresponding parameter name and returns the struct.
func (r *GetEstimatedCostByOrgOptionalParameters) WithEndMonth(endMonth time.Time) *GetEstimatedCostByOrgOptionalParameters {
	r.EndMonth = &endMonth
	return r
}

// WithStartDate sets the corresponding parameter name and returns the struct.
func (r *GetEstimatedCostByOrgOptionalParameters) WithStartDate(startDate time.Time) *GetEstimatedCostByOrgOptionalParameters {
	r.StartDate = &startDate
	return r
}

// WithEndDate sets the corresponding parameter name and returns the struct.
func (r *GetEstimatedCostByOrgOptionalParameters) WithEndDate(endDate time.Time) *GetEstimatedCostByOrgOptionalParameters {
	r.EndDate = &endDate
	return r
}

func (a *UsageMeteringApi) buildGetEstimatedCostByOrgRequest(ctx _context.Context, view string, o ...GetEstimatedCostByOrgOptionalParameters) (apiGetEstimatedCostByOrgRequest, error) {
	req := apiGetEstimatedCostByOrgRequest{
		ctx:  ctx,
		view: &view,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type GetEstimatedCostByOrgOptionalParameters is allowed")
	}

	if o != nil {
		req.startMonth = o[0].StartMonth
		req.endMonth = o[0].EndMonth
		req.startDate = o[0].StartDate
		req.endDate = o[0].EndDate
	}
	return req, nil
}

// GetEstimatedCostByOrg Get estimated cost across your account.
// Get estimated cost across multi-org and single root-org accounts.
// Estimated cost data is only available for the current month and previous month. To access historical costs prior to this, use the /cost_by_org endpoint.
func (a *UsageMeteringApi) GetEstimatedCostByOrg(ctx _context.Context, view string, o ...GetEstimatedCostByOrgOptionalParameters) (CostByOrgResponse, *_nethttp.Response, error) {
	req, err := a.buildGetEstimatedCostByOrgRequest(ctx, view, o...)
	if err != nil {
		var localVarReturnValue CostByOrgResponse
		return localVarReturnValue, nil, err
	}

	return a.getEstimatedCostByOrgExecute(req)
}

// getEstimatedCostByOrgExecute executes the request.
func (a *UsageMeteringApi) getEstimatedCostByOrgExecute(r apiGetEstimatedCostByOrgRequest) (CostByOrgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CostByOrgResponse
	)

	operationId := "v2.GetEstimatedCostByOrg"
	if a.Client.Cfg.IsUnstableOperationEnabled(operationId) {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	} else {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApi.GetEstimatedCostByOrg")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/estimated_cost"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.view == nil {
		return localVarReturnValue, nil, datadog.ReportError("view is required and must be specified")
	}
	localVarQueryParams.Add("view", datadog.ParameterToString(*r.view, ""))
	if r.startMonth != nil {
		localVarQueryParams.Add("start_month", datadog.ParameterToString(*r.startMonth, ""))
	}
	if r.endMonth != nil {
		localVarQueryParams.Add("end_month", datadog.ParameterToString(*r.endMonth, ""))
	}
	if r.startDate != nil {
		localVarQueryParams.Add("start_date", datadog.ParameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("end_date", datadog.ParameterToString(*r.endDate, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		newErr := datadog.GenericOpenAPIError{
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
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetHourlyUsageRequest struct {
	ctx                      _context.Context
	filterTimestampStart     *time.Time
	filterProductFamilies    *string
	filterTimestampEnd       *time.Time
	filterIncludeDescendants *bool
	filterVersions           *string
	pageLimit                *int32
	pageNextRecordId         *string
}

// GetHourlyUsageOptionalParameters holds optional parameters for GetHourlyUsage.
type GetHourlyUsageOptionalParameters struct {
	FilterTimestampEnd       *time.Time
	FilterIncludeDescendants *bool
	FilterVersions           *string
	PageLimit                *int32
	PageNextRecordId         *string
}

// NewGetHourlyUsageOptionalParameters creates an empty struct for parameters.
func NewGetHourlyUsageOptionalParameters() *GetHourlyUsageOptionalParameters {
	this := GetHourlyUsageOptionalParameters{}
	return &this
}

// WithFilterTimestampEnd sets the corresponding parameter name and returns the struct.
func (r *GetHourlyUsageOptionalParameters) WithFilterTimestampEnd(filterTimestampEnd time.Time) *GetHourlyUsageOptionalParameters {
	r.FilterTimestampEnd = &filterTimestampEnd
	return r
}

// WithFilterIncludeDescendants sets the corresponding parameter name and returns the struct.
func (r *GetHourlyUsageOptionalParameters) WithFilterIncludeDescendants(filterIncludeDescendants bool) *GetHourlyUsageOptionalParameters {
	r.FilterIncludeDescendants = &filterIncludeDescendants
	return r
}

// WithFilterVersions sets the corresponding parameter name and returns the struct.
func (r *GetHourlyUsageOptionalParameters) WithFilterVersions(filterVersions string) *GetHourlyUsageOptionalParameters {
	r.FilterVersions = &filterVersions
	return r
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *GetHourlyUsageOptionalParameters) WithPageLimit(pageLimit int32) *GetHourlyUsageOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// WithPageNextRecordId sets the corresponding parameter name and returns the struct.
func (r *GetHourlyUsageOptionalParameters) WithPageNextRecordId(pageNextRecordId string) *GetHourlyUsageOptionalParameters {
	r.PageNextRecordId = &pageNextRecordId
	return r
}

func (a *UsageMeteringApi) buildGetHourlyUsageRequest(ctx _context.Context, filterTimestampStart time.Time, filterProductFamilies string, o ...GetHourlyUsageOptionalParameters) (apiGetHourlyUsageRequest, error) {
	req := apiGetHourlyUsageRequest{
		ctx:                   ctx,
		filterTimestampStart:  &filterTimestampStart,
		filterProductFamilies: &filterProductFamilies,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type GetHourlyUsageOptionalParameters is allowed")
	}

	if o != nil {
		req.filterTimestampEnd = o[0].FilterTimestampEnd
		req.filterIncludeDescendants = o[0].FilterIncludeDescendants
		req.filterVersions = o[0].FilterVersions
		req.pageLimit = o[0].PageLimit
		req.pageNextRecordId = o[0].PageNextRecordId
	}
	return req, nil
}

// GetHourlyUsage Get hourly usage by product family.
// Get hourly usage by product family
func (a *UsageMeteringApi) GetHourlyUsage(ctx _context.Context, filterTimestampStart time.Time, filterProductFamilies string, o ...GetHourlyUsageOptionalParameters) (HourlyUsageResponse, *_nethttp.Response, error) {
	req, err := a.buildGetHourlyUsageRequest(ctx, filterTimestampStart, filterProductFamilies, o...)
	if err != nil {
		var localVarReturnValue HourlyUsageResponse
		return localVarReturnValue, nil, err
	}

	return a.getHourlyUsageExecute(req)
}

// getHourlyUsageExecute executes the request.
func (a *UsageMeteringApi) getHourlyUsageExecute(r apiGetHourlyUsageRequest) (HourlyUsageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue HourlyUsageResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApi.GetHourlyUsage")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/hourly_usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.filterTimestampStart == nil {
		return localVarReturnValue, nil, datadog.ReportError("filterTimestampStart is required and must be specified")
	}
	if r.filterProductFamilies == nil {
		return localVarReturnValue, nil, datadog.ReportError("filterProductFamilies is required and must be specified")
	}
	localVarQueryParams.Add("filter[timestamp][start]", datadog.ParameterToString(*r.filterTimestampStart, ""))
	localVarQueryParams.Add("filter[product_families]", datadog.ParameterToString(*r.filterProductFamilies, ""))
	if r.filterTimestampEnd != nil {
		localVarQueryParams.Add("filter[timestamp][end]", datadog.ParameterToString(*r.filterTimestampEnd, ""))
	}
	if r.filterIncludeDescendants != nil {
		localVarQueryParams.Add("filter[include_descendants]", datadog.ParameterToString(*r.filterIncludeDescendants, ""))
	}
	if r.filterVersions != nil {
		localVarQueryParams.Add("filter[versions]", datadog.ParameterToString(*r.filterVersions, ""))
	}
	if r.pageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*r.pageLimit, ""))
	}
	if r.pageNextRecordId != nil {
		localVarQueryParams.Add("page[next_record_id]", datadog.ParameterToString(*r.pageNextRecordId, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		newErr := datadog.GenericOpenAPIError{
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
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetUsageApplicationSecurityMonitoringRequest struct {
	ctx     _context.Context
	startHr *time.Time
	endHr   *time.Time
}

// GetUsageApplicationSecurityMonitoringOptionalParameters holds optional parameters for GetUsageApplicationSecurityMonitoring.
type GetUsageApplicationSecurityMonitoringOptionalParameters struct {
	EndHr *time.Time
}

// NewGetUsageApplicationSecurityMonitoringOptionalParameters creates an empty struct for parameters.
func NewGetUsageApplicationSecurityMonitoringOptionalParameters() *GetUsageApplicationSecurityMonitoringOptionalParameters {
	this := GetUsageApplicationSecurityMonitoringOptionalParameters{}
	return &this
}

// WithEndHr sets the corresponding parameter name and returns the struct.
func (r *GetUsageApplicationSecurityMonitoringOptionalParameters) WithEndHr(endHr time.Time) *GetUsageApplicationSecurityMonitoringOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApi) buildGetUsageApplicationSecurityMonitoringRequest(ctx _context.Context, startHr time.Time, o ...GetUsageApplicationSecurityMonitoringOptionalParameters) (apiGetUsageApplicationSecurityMonitoringRequest, error) {
	req := apiGetUsageApplicationSecurityMonitoringRequest{
		ctx:     ctx,
		startHr: &startHr,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type GetUsageApplicationSecurityMonitoringOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

// GetUsageApplicationSecurityMonitoring Get hourly usage for application security.
// Get hourly usage for application security .
// **Note:** hourly usage data for all products is now available in the [Get hourly usage by product family API](https://docs.datadoghq.com/api/latest/usage-metering/#get-hourly-usage-by-product-family)
func (a *UsageMeteringApi) GetUsageApplicationSecurityMonitoring(ctx _context.Context, startHr time.Time, o ...GetUsageApplicationSecurityMonitoringOptionalParameters) (UsageApplicationSecurityMonitoringResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageApplicationSecurityMonitoringRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageApplicationSecurityMonitoringResponse
		return localVarReturnValue, nil, err
	}

	return a.getUsageApplicationSecurityMonitoringExecute(req)
}

// getUsageApplicationSecurityMonitoringExecute executes the request.
func (a *UsageMeteringApi) getUsageApplicationSecurityMonitoringExecute(r apiGetUsageApplicationSecurityMonitoringRequest) (UsageApplicationSecurityMonitoringResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageApplicationSecurityMonitoringResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApi.GetUsageApplicationSecurityMonitoring")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/application_security"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, datadog.ReportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", datadog.ParameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", datadog.ParameterToString(*r.endHr, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		newErr := datadog.GenericOpenAPIError{
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
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetUsageLambdaTracedInvocationsRequest struct {
	ctx     _context.Context
	startHr *time.Time
	endHr   *time.Time
}

// GetUsageLambdaTracedInvocationsOptionalParameters holds optional parameters for GetUsageLambdaTracedInvocations.
type GetUsageLambdaTracedInvocationsOptionalParameters struct {
	EndHr *time.Time
}

// NewGetUsageLambdaTracedInvocationsOptionalParameters creates an empty struct for parameters.
func NewGetUsageLambdaTracedInvocationsOptionalParameters() *GetUsageLambdaTracedInvocationsOptionalParameters {
	this := GetUsageLambdaTracedInvocationsOptionalParameters{}
	return &this
}

// WithEndHr sets the corresponding parameter name and returns the struct.
func (r *GetUsageLambdaTracedInvocationsOptionalParameters) WithEndHr(endHr time.Time) *GetUsageLambdaTracedInvocationsOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApi) buildGetUsageLambdaTracedInvocationsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageLambdaTracedInvocationsOptionalParameters) (apiGetUsageLambdaTracedInvocationsRequest, error) {
	req := apiGetUsageLambdaTracedInvocationsRequest{
		ctx:     ctx,
		startHr: &startHr,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type GetUsageLambdaTracedInvocationsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

// GetUsageLambdaTracedInvocations Get hourly usage for lambda traced invocations.
// Get hourly usage for lambda traced invocations.
// **Note:** hourly usage data for all products is now available in the [Get hourly usage by product family API](https://docs.datadoghq.com/api/latest/usage-metering/#get-hourly-usage-by-product-family)
func (a *UsageMeteringApi) GetUsageLambdaTracedInvocations(ctx _context.Context, startHr time.Time, o ...GetUsageLambdaTracedInvocationsOptionalParameters) (UsageLambdaTracedInvocationsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageLambdaTracedInvocationsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageLambdaTracedInvocationsResponse
		return localVarReturnValue, nil, err
	}

	return a.getUsageLambdaTracedInvocationsExecute(req)
}

// getUsageLambdaTracedInvocationsExecute executes the request.
func (a *UsageMeteringApi) getUsageLambdaTracedInvocationsExecute(r apiGetUsageLambdaTracedInvocationsRequest) (UsageLambdaTracedInvocationsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageLambdaTracedInvocationsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApi.GetUsageLambdaTracedInvocations")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/lambda_traced_invocations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, datadog.ReportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", datadog.ParameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", datadog.ParameterToString(*r.endHr, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		newErr := datadog.GenericOpenAPIError{
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
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetUsageObservabilityPipelinesRequest struct {
	ctx     _context.Context
	startHr *time.Time
	endHr   *time.Time
}

// GetUsageObservabilityPipelinesOptionalParameters holds optional parameters for GetUsageObservabilityPipelines.
type GetUsageObservabilityPipelinesOptionalParameters struct {
	EndHr *time.Time
}

// NewGetUsageObservabilityPipelinesOptionalParameters creates an empty struct for parameters.
func NewGetUsageObservabilityPipelinesOptionalParameters() *GetUsageObservabilityPipelinesOptionalParameters {
	this := GetUsageObservabilityPipelinesOptionalParameters{}
	return &this
}

// WithEndHr sets the corresponding parameter name and returns the struct.
func (r *GetUsageObservabilityPipelinesOptionalParameters) WithEndHr(endHr time.Time) *GetUsageObservabilityPipelinesOptionalParameters {
	r.EndHr = &endHr
	return r
}

func (a *UsageMeteringApi) buildGetUsageObservabilityPipelinesRequest(ctx _context.Context, startHr time.Time, o ...GetUsageObservabilityPipelinesOptionalParameters) (apiGetUsageObservabilityPipelinesRequest, error) {
	req := apiGetUsageObservabilityPipelinesRequest{
		ctx:     ctx,
		startHr: &startHr,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type GetUsageObservabilityPipelinesOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

// GetUsageObservabilityPipelines Get hourly usage for observability pipelines.
// Get hourly usage for observability pipelines.
// **Note:** hourly usage data for all products is now available in the [Get hourly usage by product family API](https://docs.datadoghq.com/api/latest/usage-metering/#get-hourly-usage-by-product-family)
func (a *UsageMeteringApi) GetUsageObservabilityPipelines(ctx _context.Context, startHr time.Time, o ...GetUsageObservabilityPipelinesOptionalParameters) (UsageObservabilityPipelinesResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageObservabilityPipelinesRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageObservabilityPipelinesResponse
		return localVarReturnValue, nil, err
	}

	return a.getUsageObservabilityPipelinesExecute(req)
}

// getUsageObservabilityPipelinesExecute executes the request.
func (a *UsageMeteringApi) getUsageObservabilityPipelinesExecute(r apiGetUsageObservabilityPipelinesRequest) (UsageObservabilityPipelinesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageObservabilityPipelinesResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApi.GetUsageObservabilityPipelines")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/observability_pipelines"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, datadog.ReportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", datadog.ParameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", datadog.ParameterToString(*r.endHr, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
		newErr := datadog.GenericOpenAPIError{
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
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewUsageMeteringApi Returns NewUsageMeteringApi.
func NewUsageMeteringApi(client *datadog.APIClient) *UsageMeteringApi {
	return &UsageMeteringApi{
		Client: client,
	}
}
