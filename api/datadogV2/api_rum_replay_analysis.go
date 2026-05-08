// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_fmt "fmt"
	_log "log"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumReplayAnalysisApi service type
type RumReplayAnalysisApi datadog.Service

// GetReplayAnalysisIssue Get replay analysis issue.
// Retrieve details of a specific RUM replay analysis issue by its identifier.
func (a *RumReplayAnalysisApi) GetReplayAnalysisIssue(ctx _context.Context, issueId string) (ReplayAnalysisIssueResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ReplayAnalysisIssueResponse
	)

	operationId := "v2.GetReplayAnalysisIssue"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumReplayAnalysisApi.GetReplayAnalysisIssue")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/replay/analysis/issues/{issue_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{issue_id}", _neturl.PathEscape(datadog.ParameterToString(issueId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 404 {
			var v JSONAPIErrorResponse
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

// ListReplayAnalysisIssueSessionsOptionalParameters holds optional parameters for ListReplayAnalysisIssueSessions.
type ListReplayAnalysisIssueSessionsOptionalParameters struct {
	Sort       *string
	PageNumber *int32
	PageSize   *int32
}

// NewListReplayAnalysisIssueSessionsOptionalParameters creates an empty struct for parameters.
func NewListReplayAnalysisIssueSessionsOptionalParameters() *ListReplayAnalysisIssueSessionsOptionalParameters {
	this := ListReplayAnalysisIssueSessionsOptionalParameters{}
	return &this
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssueSessionsOptionalParameters) WithSort(sort string) *ListReplayAnalysisIssueSessionsOptionalParameters {
	r.Sort = &sort
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssueSessionsOptionalParameters) WithPageNumber(pageNumber int32) *ListReplayAnalysisIssueSessionsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssueSessionsOptionalParameters) WithPageSize(pageSize int32) *ListReplayAnalysisIssueSessionsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// ListReplayAnalysisIssueSessions List replay analysis issue sessions.
// Retrieve a paginated list of sessions related to a specific RUM replay analysis issue.
func (a *RumReplayAnalysisApi) ListReplayAnalysisIssueSessions(ctx _context.Context, issueId string, o ...ListReplayAnalysisIssueSessionsOptionalParameters) (ReplayAnalysisIssueSessionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ReplayAnalysisIssueSessionsResponse
		optionalParams      ListReplayAnalysisIssueSessionsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListReplayAnalysisIssueSessionsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListReplayAnalysisIssueSessions"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumReplayAnalysisApi.ListReplayAnalysisIssueSessions")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/replay/analysis/issues/{issue_id}/sessions"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{issue_id}", _neturl.PathEscape(datadog.ParameterToString(issueId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 404 {
			var v JSONAPIErrorResponse
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

// ListReplayAnalysisIssuesOptionalParameters holds optional parameters for ListReplayAnalysisIssues.
type ListReplayAnalysisIssuesOptionalParameters struct {
	FilterApplicationId *string
	FilterSeverity      *string
	FilterViewName      *string
	FilterIssueCategory *string
	Sort                *string
	PageNumber          *int32
	PageSize            *int32
}

// NewListReplayAnalysisIssuesOptionalParameters creates an empty struct for parameters.
func NewListReplayAnalysisIssuesOptionalParameters() *ListReplayAnalysisIssuesOptionalParameters {
	this := ListReplayAnalysisIssuesOptionalParameters{}
	return &this
}

// WithFilterApplicationId sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssuesOptionalParameters) WithFilterApplicationId(filterApplicationId string) *ListReplayAnalysisIssuesOptionalParameters {
	r.FilterApplicationId = &filterApplicationId
	return r
}

// WithFilterSeverity sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssuesOptionalParameters) WithFilterSeverity(filterSeverity string) *ListReplayAnalysisIssuesOptionalParameters {
	r.FilterSeverity = &filterSeverity
	return r
}

// WithFilterViewName sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssuesOptionalParameters) WithFilterViewName(filterViewName string) *ListReplayAnalysisIssuesOptionalParameters {
	r.FilterViewName = &filterViewName
	return r
}

// WithFilterIssueCategory sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssuesOptionalParameters) WithFilterIssueCategory(filterIssueCategory string) *ListReplayAnalysisIssuesOptionalParameters {
	r.FilterIssueCategory = &filterIssueCategory
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssuesOptionalParameters) WithSort(sort string) *ListReplayAnalysisIssuesOptionalParameters {
	r.Sort = &sort
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssuesOptionalParameters) WithPageNumber(pageNumber int32) *ListReplayAnalysisIssuesOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListReplayAnalysisIssuesOptionalParameters) WithPageSize(pageSize int32) *ListReplayAnalysisIssuesOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// ListReplayAnalysisIssues List replay analysis issues.
// Retrieve a paginated list of RUM replay analysis issues, optionally filtered by application, severity, view name, or issue category.
func (a *RumReplayAnalysisApi) ListReplayAnalysisIssues(ctx _context.Context, o ...ListReplayAnalysisIssuesOptionalParameters) (ReplayAnalysisIssuesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ReplayAnalysisIssuesResponse
		optionalParams      ListReplayAnalysisIssuesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListReplayAnalysisIssuesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListReplayAnalysisIssues"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumReplayAnalysisApi.ListReplayAnalysisIssues")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/replay/analysis/issues"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterApplicationId != nil {
		localVarQueryParams.Add("filter[application_id]", datadog.ParameterToString(*optionalParams.FilterApplicationId, ""))
	}
	if optionalParams.FilterSeverity != nil {
		localVarQueryParams.Add("filter[severity]", datadog.ParameterToString(*optionalParams.FilterSeverity, ""))
	}
	if optionalParams.FilterViewName != nil {
		localVarQueryParams.Add("filter[view_name]", datadog.ParameterToString(*optionalParams.FilterViewName, ""))
	}
	if optionalParams.FilterIssueCategory != nil {
		localVarQueryParams.Add("filter[issue_category]", datadog.ParameterToString(*optionalParams.FilterIssueCategory, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v JSONAPIErrorResponse
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

// NewRumReplayAnalysisApi Returns NewRumReplayAnalysisApi.
func NewRumReplayAnalysisApi(client *datadog.APIClient) *RumReplayAnalysisApi {
	return &RumReplayAnalysisApi{
		Client: client,
	}
}
