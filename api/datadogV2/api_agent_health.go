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

// AgentHealthApi service type
type AgentHealthApi datadog.Service

// GetAgentHealthIssue Get an Agent Health issue.
// Get an Agent Health issue and its available remediation in your organization.
// Only open issues are returned. The `remediation` attribute is omitted when
// the issue has no remediation.
//
// Returns `404` if the issue does not exist, is not open, belongs to another
// organization, or you do not have access to its data. This response does not
// reveal whether the issue exists. Returns `403` if you lack the `hosts_read` permission.
func (a *AgentHealthApi) GetAgentHealthIssue(ctx _context.Context, issueId string) (AgentHealthIssueResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AgentHealthIssueResponse
	)

	operationId := "v2.GetAgentHealthIssue"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.AgentHealthApi.GetAgentHealthIssue")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/agenthealth/issues/{issue_id}"
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// ListAgentHealthHostIssuesOptionalParameters holds optional parameters for ListAgentHealthHostIssues.
type ListAgentHealthHostIssuesOptionalParameters struct {
	AgentType *AgentHealthAgentType
}

// NewListAgentHealthHostIssuesOptionalParameters creates an empty struct for parameters.
func NewListAgentHealthHostIssuesOptionalParameters() *ListAgentHealthHostIssuesOptionalParameters {
	this := ListAgentHealthHostIssuesOptionalParameters{}
	return &this
}

// WithAgentType sets the corresponding parameter name and returns the struct.
func (r *ListAgentHealthHostIssuesOptionalParameters) WithAgentType(agentType AgentHealthAgentType) *ListAgentHealthHostIssuesOptionalParameters {
	r.AgentType = &agentType
	return r
}

// ListAgentHealthHostIssues List active Agent Health issues.
// Returns the active Agent Health issues for a node Agent or Cluster Agent resource.
func (a *AgentHealthApi) ListAgentHealthHostIssues(ctx _context.Context, resourceId string, o ...ListAgentHealthHostIssuesOptionalParameters) (AgentHealthHostIssuesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AgentHealthHostIssuesResponse
		optionalParams      ListAgentHealthHostIssuesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListAgentHealthHostIssuesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.AgentHealthApi.ListAgentHealthHostIssues")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/agenthealth/hosts/{resource_id}/issues"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{resource_id}", _neturl.PathEscape(datadog.ParameterToString(resourceId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.AgentType != nil {
		localVarQueryParams.Add("agent_type", datadog.ParameterToString(*optionalParams.AgentType, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
	)

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
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

// ListAgentHealthIssuesOptionalParameters holds optional parameters for ListAgentHealthIssues.
type ListAgentHealthIssuesOptionalParameters struct {
	PageSize           *int64
	PageNumber         *int64
	FilterSeverity     *string
	FilterCategory     *string
	FilterIssueType    *string
	FilterResourceKind *string
	FilterResourceId   *string
}

// NewListAgentHealthIssuesOptionalParameters creates an empty struct for parameters.
func NewListAgentHealthIssuesOptionalParameters() *ListAgentHealthIssuesOptionalParameters {
	this := ListAgentHealthIssuesOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListAgentHealthIssuesOptionalParameters) WithPageSize(pageSize int64) *ListAgentHealthIssuesOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListAgentHealthIssuesOptionalParameters) WithPageNumber(pageNumber int64) *ListAgentHealthIssuesOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithFilterSeverity sets the corresponding parameter name and returns the struct.
func (r *ListAgentHealthIssuesOptionalParameters) WithFilterSeverity(filterSeverity string) *ListAgentHealthIssuesOptionalParameters {
	r.FilterSeverity = &filterSeverity
	return r
}

// WithFilterCategory sets the corresponding parameter name and returns the struct.
func (r *ListAgentHealthIssuesOptionalParameters) WithFilterCategory(filterCategory string) *ListAgentHealthIssuesOptionalParameters {
	r.FilterCategory = &filterCategory
	return r
}

// WithFilterIssueType sets the corresponding parameter name and returns the struct.
func (r *ListAgentHealthIssuesOptionalParameters) WithFilterIssueType(filterIssueType string) *ListAgentHealthIssuesOptionalParameters {
	r.FilterIssueType = &filterIssueType
	return r
}

// WithFilterResourceKind sets the corresponding parameter name and returns the struct.
func (r *ListAgentHealthIssuesOptionalParameters) WithFilterResourceKind(filterResourceKind string) *ListAgentHealthIssuesOptionalParameters {
	r.FilterResourceKind = &filterResourceKind
	return r
}

// WithFilterResourceId sets the corresponding parameter name and returns the struct.
func (r *ListAgentHealthIssuesOptionalParameters) WithFilterResourceId(filterResourceId string) *ListAgentHealthIssuesOptionalParameters {
	r.FilterResourceId = &filterResourceId
	return r
}

// ListAgentHealthIssues List Agent Health issues.
// List Agent Health issues across your organization. Only open issues are returned.
// Returns `403` if you lack the `hosts_read` permission.
//
// Results use a stable order: severity rank (`high`, `medium`, `low`, then other values),
// then `first_seen_at` descending with missing timestamps last, then issue ID ascending.
// Use `page[size]` and the zero-based `page[number]` to paginate. Invalid page values
// return a `400` JSON:API error response. `meta.page.total_filtered_count` is the number
// of issues matching the filters, and `meta.page.total_count` the number of open issues
// in the organization.
//
// Each filter accepts up to 50 comma-separated exact values of at most 256 characters
// each, combined with OR within a filter and AND across filters. Blank entries are ignored; an all-blank filter
// is treated as absent.
func (a *AgentHealthApi) ListAgentHealthIssues(ctx _context.Context, o ...ListAgentHealthIssuesOptionalParameters) (AgentHealthIssuesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AgentHealthIssuesResponse
		optionalParams      ListAgentHealthIssuesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListAgentHealthIssuesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListAgentHealthIssues"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.AgentHealthApi.ListAgentHealthIssues")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/agenthealth/issues"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.FilterSeverity != nil {
		localVarQueryParams.Add("filter[severity]", datadog.ParameterToString(*optionalParams.FilterSeverity, ""))
	}
	if optionalParams.FilterCategory != nil {
		localVarQueryParams.Add("filter[category]", datadog.ParameterToString(*optionalParams.FilterCategory, ""))
	}
	if optionalParams.FilterIssueType != nil {
		localVarQueryParams.Add("filter[issue_type]", datadog.ParameterToString(*optionalParams.FilterIssueType, ""))
	}
	if optionalParams.FilterResourceKind != nil {
		localVarQueryParams.Add("filter[resource_kind]", datadog.ParameterToString(*optionalParams.FilterResourceKind, ""))
	}
	if optionalParams.FilterResourceId != nil {
		localVarQueryParams.Add("filter[resource_id]", datadog.ParameterToString(*optionalParams.FilterResourceId, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 {
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

// NewAgentHealthApi Returns NewAgentHealthApi.
func NewAgentHealthApi(client *datadog.APIClient) *AgentHealthApi {
	return &AgentHealthApi{
		Client: client,
	}
}
