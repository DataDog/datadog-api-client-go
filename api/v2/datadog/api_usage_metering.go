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
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
)

// UsageMeteringApiService service type
type UsageMeteringApiService common.Service

type apiGetCostByOrgRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
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

func (a *UsageMeteringApiService) buildGetCostByOrgRequest(ctx _context.Context, startMonth time.Time, o ...GetCostByOrgOptionalParameters) (apiGetCostByOrgRequest, error) {
	req := apiGetCostByOrgRequest{
		ApiService: a,
		ctx:        ctx,
		startMonth: &startMonth,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type GetCostByOrgOptionalParameters is allowed")
	}

	if o != nil {
		req.endMonth = o[0].EndMonth
	}
	return req, nil
}

// GetCostByOrg Get cost across multi-org account.
// Get cost across multi-org account. Cost by org data for a given month becomes available no later than the 16th of the following month.
func (a *UsageMeteringApiService) GetCostByOrg(ctx _context.Context, startMonth time.Time, o ...GetCostByOrgOptionalParameters) (CostByOrgResponse, *_nethttp.Response, error) {
	req, err := a.buildGetCostByOrgRequest(ctx, startMonth, o...)
	if err != nil {
		var localVarReturnValue CostByOrgResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getCostByOrgExecute(req)
}

// getCostByOrgExecute executes the request.
func (a *UsageMeteringApiService) getCostByOrgExecute(r apiGetCostByOrgRequest) (CostByOrgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CostByOrgResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApiService.GetCostByOrg")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/cost_by_org"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startMonth == nil {
		return localVarReturnValue, nil, common.ReportError("startMonth is required and must be specified")
	}
	localVarQueryParams.Add("start_month", common.ParameterToString(*r.startMonth, ""))
	if r.endMonth != nil {
		localVarQueryParams.Add("end_month", common.ParameterToString(*r.endMonth, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

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
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetEstimatedCostByOrgRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
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

func (a *UsageMeteringApiService) buildGetEstimatedCostByOrgRequest(ctx _context.Context, o ...GetEstimatedCostByOrgOptionalParameters) (apiGetEstimatedCostByOrgRequest, error) {
	req := apiGetEstimatedCostByOrgRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type GetEstimatedCostByOrgOptionalParameters is allowed")
	}

	if o != nil {
		req.startMonth = o[0].StartMonth
		req.endMonth = o[0].EndMonth
		req.startDate = o[0].StartDate
		req.endDate = o[0].EndDate
	}
	return req, nil
}

// GetEstimatedCostByOrg Get estimated cost across multi-org account.
// Get estimated cost across multi-org account.
func (a *UsageMeteringApiService) GetEstimatedCostByOrg(ctx _context.Context, o ...GetEstimatedCostByOrgOptionalParameters) (CostByOrgResponse, *_nethttp.Response, error) {
	req, err := a.buildGetEstimatedCostByOrgRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue CostByOrgResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getEstimatedCostByOrgExecute(req)
}

// getEstimatedCostByOrgExecute executes the request.
func (a *UsageMeteringApiService) getEstimatedCostByOrgExecute(r apiGetEstimatedCostByOrgRequest) (CostByOrgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CostByOrgResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApiService.GetEstimatedCostByOrg")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/estimated_cost_by_org"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startMonth != nil {
		localVarQueryParams.Add("start_month", common.ParameterToString(*r.startMonth, ""))
	}
	if r.endMonth != nil {
		localVarQueryParams.Add("end_month", common.ParameterToString(*r.endMonth, ""))
	}
	if r.startDate != nil {
		localVarQueryParams.Add("start_date", common.ParameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("end_date", common.ParameterToString(*r.endDate, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

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
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetUsageApplicationSecurityMonitoringRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
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

func (a *UsageMeteringApiService) buildGetUsageApplicationSecurityMonitoringRequest(ctx _context.Context, startHr time.Time, o ...GetUsageApplicationSecurityMonitoringOptionalParameters) (apiGetUsageApplicationSecurityMonitoringRequest, error) {
	req := apiGetUsageApplicationSecurityMonitoringRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type GetUsageApplicationSecurityMonitoringOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

// GetUsageApplicationSecurityMonitoring Get hourly usage for Application Security.
// Get hourly usage for Application Security .
func (a *UsageMeteringApiService) GetUsageApplicationSecurityMonitoring(ctx _context.Context, startHr time.Time, o ...GetUsageApplicationSecurityMonitoringOptionalParameters) (UsageApplicationSecurityMonitoringResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageApplicationSecurityMonitoringRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageApplicationSecurityMonitoringResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageApplicationSecurityMonitoringExecute(req)
}

// getUsageApplicationSecurityMonitoringExecute executes the request.
func (a *UsageMeteringApiService) getUsageApplicationSecurityMonitoringExecute(r apiGetUsageApplicationSecurityMonitoringRequest) (UsageApplicationSecurityMonitoringResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageApplicationSecurityMonitoringResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApiService.GetUsageApplicationSecurityMonitoring")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/application_security"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, common.ReportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", common.ParameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", common.ParameterToString(*r.endHr, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

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
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetUsageLambdaTracedInvocationsRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
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

func (a *UsageMeteringApiService) buildGetUsageLambdaTracedInvocationsRequest(ctx _context.Context, startHr time.Time, o ...GetUsageLambdaTracedInvocationsOptionalParameters) (apiGetUsageLambdaTracedInvocationsRequest, error) {
	req := apiGetUsageLambdaTracedInvocationsRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type GetUsageLambdaTracedInvocationsOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

// GetUsageLambdaTracedInvocations Get hourly usage for Lambda Traced Invocations.
// Get hourly usage for Lambda Traced Invocations.
func (a *UsageMeteringApiService) GetUsageLambdaTracedInvocations(ctx _context.Context, startHr time.Time, o ...GetUsageLambdaTracedInvocationsOptionalParameters) (UsageLambdaTracedInvocationsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageLambdaTracedInvocationsRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageLambdaTracedInvocationsResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageLambdaTracedInvocationsExecute(req)
}

// getUsageLambdaTracedInvocationsExecute executes the request.
func (a *UsageMeteringApiService) getUsageLambdaTracedInvocationsExecute(r apiGetUsageLambdaTracedInvocationsRequest) (UsageLambdaTracedInvocationsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageLambdaTracedInvocationsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApiService.GetUsageLambdaTracedInvocations")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/lambda_traced_invocations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, common.ReportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", common.ParameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", common.ParameterToString(*r.endHr, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

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
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetUsageObservabilityPipelinesRequest struct {
	ctx        _context.Context
	ApiService *UsageMeteringApiService
	startHr    *time.Time
	endHr      *time.Time
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

func (a *UsageMeteringApiService) buildGetUsageObservabilityPipelinesRequest(ctx _context.Context, startHr time.Time, o ...GetUsageObservabilityPipelinesOptionalParameters) (apiGetUsageObservabilityPipelinesRequest, error) {
	req := apiGetUsageObservabilityPipelinesRequest{
		ApiService: a,
		ctx:        ctx,
		startHr:    &startHr,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type GetUsageObservabilityPipelinesOptionalParameters is allowed")
	}

	if o != nil {
		req.endHr = o[0].EndHr
	}
	return req, nil
}

// GetUsageObservabilityPipelines Get hourly usage for Observability Pipelines.
// Get hourly usage for Observability Pipelines.
func (a *UsageMeteringApiService) GetUsageObservabilityPipelines(ctx _context.Context, startHr time.Time, o ...GetUsageObservabilityPipelinesOptionalParameters) (UsageObservabilityPipelinesResponse, *_nethttp.Response, error) {
	req, err := a.buildGetUsageObservabilityPipelinesRequest(ctx, startHr, o...)
	if err != nil {
		var localVarReturnValue UsageObservabilityPipelinesResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getUsageObservabilityPipelinesExecute(req)
}

// getUsageObservabilityPipelinesExecute executes the request.
func (a *UsageMeteringApiService) getUsageObservabilityPipelinesExecute(r apiGetUsageObservabilityPipelinesRequest) (UsageObservabilityPipelinesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UsageObservabilityPipelinesResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.UsageMeteringApiService.GetUsageObservabilityPipelines")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/usage/observability_pipelines"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.startHr == nil {
		return localVarReturnValue, nil, common.ReportError("startHr is required and must be specified")
	}
	localVarQueryParams.Add("start_hr", common.ParameterToString(*r.startHr, ""))
	if r.endHr != nil {
		localVarQueryParams.Add("end_hr", common.ParameterToString(*r.endHr, ""))
	}
	localVarHeaderParams["Accept"] = "application/json;datetime-format=rfc3339"

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
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// UsageMeteringApi Returns new UsageMeteringApi service.
func UsageMeteringApi(client *common.APIClient) *UsageMeteringApiService {
	return &UsageMeteringApiService{
		Client: client,
	}
}
