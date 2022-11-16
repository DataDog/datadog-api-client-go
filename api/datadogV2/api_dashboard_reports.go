// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"bytes"
	_context "context"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardReportsApi service type
type DashboardReportsApi datadog.Service

type apiCreateDashboardReportConfigRequest struct {
	ctx         _context.Context
	dashboardId string
	body        *DashboardReportCreateRequest
}

func (a *DashboardReportsApi) buildCreateDashboardReportConfigRequest(ctx _context.Context, dashboardId string, body DashboardReportCreateRequest) (apiCreateDashboardReportConfigRequest, error) {
	req := apiCreateDashboardReportConfigRequest{
		ctx:         ctx,
		dashboardId: dashboardId,
		body:        &body,
	}
	return req, nil
}

// CreateDashboardReportConfig Create a new Dashboard Report.
// New dashboard report configuration for a given dashboard. This creates a new report email schedule, frequency, timeframe, and more.
func (a *DashboardReportsApi) CreateDashboardReportConfig(ctx _context.Context, dashboardId string, body DashboardReportCreateRequest) (DashboardReportResponse, *_nethttp.Response, error) {
	req, err := a.buildCreateDashboardReportConfigRequest(ctx, dashboardId, body)
	if err != nil {
		var localVarReturnValue DashboardReportResponse
		return localVarReturnValue, nil, err
	}

	return a.createDashboardReportConfigExecute(req)
}

// createDashboardReportConfigExecute executes the request.
func (a *DashboardReportsApi) createDashboardReportConfigExecute(r apiCreateDashboardReportConfigRequest) (DashboardReportResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue DashboardReportResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.DashboardReportsApi.CreateDashboardReportConfig")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dashboards/{dashboard_id}/reports"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboard_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.dashboardId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
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

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 422 || localVarHTTPResponse.StatusCode == 429 {
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

type apiDeleteDashboardReportConfigRequest struct {
	ctx         _context.Context
	dashboardId string
	reportUuid  string
}

func (a *DashboardReportsApi) buildDeleteDashboardReportConfigRequest(ctx _context.Context, dashboardId string, reportUuid string) (apiDeleteDashboardReportConfigRequest, error) {
	req := apiDeleteDashboardReportConfigRequest{
		ctx:         ctx,
		dashboardId: dashboardId,
		reportUuid:  reportUuid,
	}
	return req, nil
}

// DeleteDashboardReportConfig Delete a Dashboard Report.
// Delete a dashboard report configuration, disabling the sending of scheduled emails for this report in the future. This operation cannot be undone. To pause the sending of emails for this report without deleting it, deactivate the report with a `PATCH` request.
func (a *DashboardReportsApi) DeleteDashboardReportConfig(ctx _context.Context, dashboardId string, reportUuid string) (*_nethttp.Response, error) {
	req, err := a.buildDeleteDashboardReportConfigRequest(ctx, dashboardId, reportUuid)
	if err != nil {
		return nil, err
	}

	return a.deleteDashboardReportConfigExecute(req)
}

// deleteDashboardReportConfigExecute executes the request.
func (a *DashboardReportsApi) deleteDashboardReportConfigExecute(r apiDeleteDashboardReportConfigRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.DashboardReportsApi.DeleteDashboardReportConfig")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dashboards/{dashboard_id}/reports/{report_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboard_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.dashboardId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"report_uuid"+"}", _neturl.PathEscape(datadog.ParameterToString(r.reportUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

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
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 422 || localVarHTTPResponse.StatusCode == 429 {
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

type apiGetDashboardReportConfigRequest struct {
	ctx         _context.Context
	dashboardId string
	reportUuid  string
}

func (a *DashboardReportsApi) buildGetDashboardReportConfigRequest(ctx _context.Context, dashboardId string, reportUuid string) (apiGetDashboardReportConfigRequest, error) {
	req := apiGetDashboardReportConfigRequest{
		ctx:         ctx,
		dashboardId: dashboardId,
		reportUuid:  reportUuid,
	}
	return req, nil
}

// GetDashboardReportConfig Get a Dashboard Report.
// Fetch a single Dashboard Report configuration. This includes the same information provided when fetching a dashboard's list of currently configured reports, but only for singular reports.
func (a *DashboardReportsApi) GetDashboardReportConfig(ctx _context.Context, dashboardId string, reportUuid string) (DashboardReportResponse, *_nethttp.Response, error) {
	req, err := a.buildGetDashboardReportConfigRequest(ctx, dashboardId, reportUuid)
	if err != nil {
		var localVarReturnValue DashboardReportResponse
		return localVarReturnValue, nil, err
	}

	return a.getDashboardReportConfigExecute(req)
}

// getDashboardReportConfigExecute executes the request.
func (a *DashboardReportsApi) getDashboardReportConfigExecute(r apiGetDashboardReportConfigRequest) (DashboardReportResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DashboardReportResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.DashboardReportsApi.GetDashboardReportConfig")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dashboards/{dashboard_id}/reports/{report_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboard_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.dashboardId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"report_uuid"+"}", _neturl.PathEscape(datadog.ParameterToString(r.reportUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

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

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 422 || localVarHTTPResponse.StatusCode == 429 {
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

type apiGetDashboardReportConfigsListRequest struct {
	ctx         _context.Context
	dashboardId string
}

func (a *DashboardReportsApi) buildGetDashboardReportConfigsListRequest(ctx _context.Context, dashboardId string) (apiGetDashboardReportConfigsListRequest, error) {
	req := apiGetDashboardReportConfigsListRequest{
		ctx:         ctx,
		dashboardId: dashboardId,
	}
	return req, nil
}

// GetDashboardReportConfigsList Get Dashboard Reports for a Dashboard.
// List of the dashboard reports configured on a given dashboard. This list includes report configurations that are both enabled and disabled, but does not include reports that have been deleted for the given dashboard.
func (a *DashboardReportsApi) GetDashboardReportConfigsList(ctx _context.Context, dashboardId string) (DashboardReportsResponse, *_nethttp.Response, error) {
	req, err := a.buildGetDashboardReportConfigsListRequest(ctx, dashboardId)
	if err != nil {
		var localVarReturnValue DashboardReportsResponse
		return localVarReturnValue, nil, err
	}

	return a.getDashboardReportConfigsListExecute(req)
}

// getDashboardReportConfigsListExecute executes the request.
func (a *DashboardReportsApi) getDashboardReportConfigsListExecute(r apiGetDashboardReportConfigsListRequest) (DashboardReportsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DashboardReportsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.DashboardReportsApi.GetDashboardReportConfigsList")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dashboards/{dashboard_id}/reports"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboard_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.dashboardId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

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

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 422 || localVarHTTPResponse.StatusCode == 429 {
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

type apiUpdateDashboardReportConfigRequest struct {
	ctx         _context.Context
	dashboardId string
	reportUuid  string
	body        *DashboardReportUpdateRequest
}

func (a *DashboardReportsApi) buildUpdateDashboardReportConfigRequest(ctx _context.Context, dashboardId string, reportUuid string, body DashboardReportUpdateRequest) (apiUpdateDashboardReportConfigRequest, error) {
	req := apiUpdateDashboardReportConfigRequest{
		ctx:         ctx,
		dashboardId: dashboardId,
		reportUuid:  reportUuid,
		body:        &body,
	}
	return req, nil
}

// UpdateDashboardReportConfig Update Dashboard Report.
// Update a Dashboard Report configuration, including the schedule and email settings.  Changes to the schedule happen immediately, but it may take up to five minutes for your report to be sent if the scheduled time is close to the time of the update request.
func (a *DashboardReportsApi) UpdateDashboardReportConfig(ctx _context.Context, dashboardId string, reportUuid string, body DashboardReportUpdateRequest) (DashboardReportResponse, *_nethttp.Response, error) {
	req, err := a.buildUpdateDashboardReportConfigRequest(ctx, dashboardId, reportUuid, body)
	if err != nil {
		var localVarReturnValue DashboardReportResponse
		return localVarReturnValue, nil, err
	}

	return a.updateDashboardReportConfigExecute(req)
}

// updateDashboardReportConfigExecute executes the request.
func (a *DashboardReportsApi) updateDashboardReportConfigExecute(r apiUpdateDashboardReportConfigRequest) (DashboardReportResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue DashboardReportResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v2.DashboardReportsApi.UpdateDashboardReportConfig")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dashboards/{dashboard_id}/reports/{report_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboard_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.dashboardId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"report_uuid"+"}", _neturl.PathEscape(datadog.ParameterToString(r.reportUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
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

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 422 || localVarHTTPResponse.StatusCode == 429 {
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

// NewDashboardReportsApi Returns NewDashboardReportsApi.
func NewDashboardReportsApi(client *datadog.APIClient) *DashboardReportsApi {
	return &DashboardReportsApi{
		Client: client,
	}
}
