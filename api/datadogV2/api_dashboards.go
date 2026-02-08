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

// DashboardsApi service type
type DashboardsApi datadog.Service

// SearchDashboardsOptionalParameters holds optional parameters for SearchDashboards.
type SearchDashboardsOptionalParameters struct {
	Query   *string
	Sort    *string
	Include *string
	Page    *int32
	Limit   *int32
}

// NewSearchDashboardsOptionalParameters creates an empty struct for parameters.
func NewSearchDashboardsOptionalParameters() *SearchDashboardsOptionalParameters {
	this := SearchDashboardsOptionalParameters{}
	return &this
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *SearchDashboardsOptionalParameters) WithQuery(query string) *SearchDashboardsOptionalParameters {
	r.Query = &query
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *SearchDashboardsOptionalParameters) WithSort(sort string) *SearchDashboardsOptionalParameters {
	r.Sort = &sort
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *SearchDashboardsOptionalParameters) WithInclude(include string) *SearchDashboardsOptionalParameters {
	r.Include = &include
	return r
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *SearchDashboardsOptionalParameters) WithPage(page int32) *SearchDashboardsOptionalParameters {
	r.Page = &page
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *SearchDashboardsOptionalParameters) WithLimit(limit int32) *SearchDashboardsOptionalParameters {
	r.Limit = &limit
	return r
}

// SearchDashboards Search dashboards.
// Search for dashboards using a query string.
func (a *DashboardsApi) SearchDashboards(ctx _context.Context, o ...SearchDashboardsOptionalParameters) (DashboardSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DashboardSearchResponse
		optionalParams      SearchDashboardsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type SearchDashboardsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.SearchDashboards"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DashboardsApi.SearchDashboards")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dashboards/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", datadog.ParameterToString(*optionalParams.Query, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
	if optionalParams.Page != nil {
		localVarQueryParams.Add("page", datadog.ParameterToString(*optionalParams.Page, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", datadog.ParameterToString(*optionalParams.Limit, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 503 {
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

// NewDashboardsApi Returns NewDashboardsApi.
func NewDashboardsApi(client *datadog.APIClient) *DashboardsApi {
	return &DashboardsApi{
		Client: client,
	}
}
