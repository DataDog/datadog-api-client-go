// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// DashboardsApi service type
type DashboardsApi datadog.Service

// ListDashboardsOptionalParameters holds optional parameters for ListDashboards.
type ListDashboardsOptionalParameters struct {
	FilterShared  *bool
	FilterDeleted *bool
	Count         *int64
	Start         *int64
}

// NewListDashboardsOptionalParameters creates an empty struct for parameters.
func NewListDashboardsOptionalParameters() *ListDashboardsOptionalParameters {
	this := ListDashboardsOptionalParameters{}
	return &this
}

// WithFilterShared sets the corresponding parameter name and returns the struct.
func (r *ListDashboardsOptionalParameters) WithFilterShared(filterShared bool) *ListDashboardsOptionalParameters {
	r.FilterShared = &filterShared
	return r
}

// WithFilterDeleted sets the corresponding parameter name and returns the struct.
func (r *ListDashboardsOptionalParameters) WithFilterDeleted(filterDeleted bool) *ListDashboardsOptionalParameters {
	r.FilterDeleted = &filterDeleted
	return r
}

// WithCount sets the corresponding parameter name and returns the struct.
func (r *ListDashboardsOptionalParameters) WithCount(count int64) *ListDashboardsOptionalParameters {
	r.Count = &count
	return r
}

// WithStart sets the corresponding parameter name and returns the struct.
func (r *ListDashboardsOptionalParameters) WithStart(start int64) *ListDashboardsOptionalParameters {
	r.Start = &start
	return r
}

// ListDashboards Get all dashboards.
// Get all dashboards.
//
// **Note**: This query will only return custom created or cloned dashboards.
// This query will not return preset dashboards.
func (a *DashboardsApi) ListDashboards(ctx _context.Context, o ...ListDashboardsOptionalParameters) (DashboardSummary, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DashboardSummary
		optionalParams      ListDashboardsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListDashboardsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.DashboardsApi.ListDashboards")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/dashboard"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["DD-API-Version"] = "2027-01-01"
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterShared != nil {
		localVarQueryParams.Add("filter[shared]", datadog.ParameterToString(*optionalParams.FilterShared, ""))
	}
	if optionalParams.FilterDeleted != nil {
		localVarQueryParams.Add("filter[deleted]", datadog.ParameterToString(*optionalParams.FilterDeleted, ""))
	}
	if optionalParams.Count != nil {
		localVarQueryParams.Add("count", datadog.ParameterToString(*optionalParams.Count, ""))
	}
	if optionalParams.Start != nil {
		localVarQueryParams.Add("start", datadog.ParameterToString(*optionalParams.Start, ""))
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v datadogV1.APIErrorResponse
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

// ListDashboardsWithPagination provides a paginated version of ListDashboards returning a channel with all items.
func (a *DashboardsApi) ListDashboardsWithPagination(ctx _context.Context, o ...ListDashboardsOptionalParameters) (<-chan datadog.PaginationResult[DashboardSummaryDefinition], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(100)
	if len(o) == 0 {
		o = append(o, ListDashboardsOptionalParameters{})
	}
	if o[0].Count != nil {
		pageSize_ = *o[0].Count
	}
	o[0].Count = &pageSize_

	items := make(chan datadog.PaginationResult[DashboardSummaryDefinition], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListDashboards(ctx, o...)
			if err != nil {
				var returnItem DashboardSummaryDefinition
				items <- datadog.PaginationResult[DashboardSummaryDefinition]{Item: returnItem, Error: err}
				break
			}
			respDashboards, ok := resp.GetDashboardsOk()
			if !ok {
				break
			}
			results := *respDashboards

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[DashboardSummaryDefinition]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) < int(pageSize_) {
				break
			}
			if o[0].Start == nil {
				o[0].Start = &pageSize_
			} else {
				pageOffset_ := *o[0].Start + pageSize_
				o[0].Start = &pageOffset_
			}
		}
		close(items)
	}()
	return items, cancel
}

// NewDashboardsApi Returns NewDashboardsApi.
func NewDashboardsApi(client *datadog.APIClient) *DashboardsApi {
	return &DashboardsApi{
		Client: client,
	}
}
