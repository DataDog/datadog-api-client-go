// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMReplayViewershipApi service type
type RUMReplayViewershipApi datadog.Service

// CreateRUMReplaySessionWatch Create RUM replay session watch.
// Record a session watch.
func (a *RUMReplayViewershipApi) CreateRUMReplaySessionWatch(ctx _context.Context, sessionId string, body Watch) (Watch, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Watch
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMReplayViewershipApi.CreateRUMReplaySessionWatch")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/replay/sessions/{session_id}/watches"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{session_id}", _neturl.PathEscape(datadog.ParameterToString(sessionId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
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

// DeleteRUMReplaySessionWatch Delete RUM replay session watch.
// Delete session watch history.
func (a *RUMReplayViewershipApi) DeleteRUMReplaySessionWatch(ctx _context.Context, sessionId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMReplayViewershipApi.DeleteRUMReplaySessionWatch")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/replay/sessions/{session_id}/watches"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{session_id}", _neturl.PathEscape(datadog.ParameterToString(sessionId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return nil, err
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
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 429 {
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

// ListRUMReplaySessionWatchersOptionalParameters holds optional parameters for ListRUMReplaySessionWatchers.
type ListRUMReplaySessionWatchersOptionalParameters struct {
	PageSize   *int64
	PageNumber *int64
}

// NewListRUMReplaySessionWatchersOptionalParameters creates an empty struct for parameters.
func NewListRUMReplaySessionWatchersOptionalParameters() *ListRUMReplaySessionWatchersOptionalParameters {
	this := ListRUMReplaySessionWatchersOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListRUMReplaySessionWatchersOptionalParameters) WithPageSize(pageSize int64) *ListRUMReplaySessionWatchersOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListRUMReplaySessionWatchersOptionalParameters) WithPageNumber(pageNumber int64) *ListRUMReplaySessionWatchersOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// ListRUMReplaySessionWatchers List RUM replay session watchers.
// List session watchers.
func (a *RUMReplayViewershipApi) ListRUMReplaySessionWatchers(ctx _context.Context, sessionId string, o ...ListRUMReplaySessionWatchersOptionalParameters) (WatcherArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue WatcherArray
		optionalParams      ListRUMReplaySessionWatchersOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListRUMReplaySessionWatchersOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMReplayViewershipApi.ListRUMReplaySessionWatchers")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/replay/sessions/{session_id}/watchers"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{session_id}", _neturl.PathEscape(datadog.ParameterToString(sessionId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
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

// ListRUMReplayViewershipHistorySessionsOptionalParameters holds optional parameters for ListRUMReplayViewershipHistorySessions.
type ListRUMReplayViewershipHistorySessionsOptionalParameters struct {
	FilterWatchedAtStart *int64
	PageNumber           *int64
	FilterCreatedBy      *string
	FilterWatchedAtEnd   *int64
	FilterSessionIds     *string
	PageSize             *int64
	FilterApplicationId  *string
}

// NewListRUMReplayViewershipHistorySessionsOptionalParameters creates an empty struct for parameters.
func NewListRUMReplayViewershipHistorySessionsOptionalParameters() *ListRUMReplayViewershipHistorySessionsOptionalParameters {
	this := ListRUMReplayViewershipHistorySessionsOptionalParameters{}
	return &this
}

// WithFilterWatchedAtStart sets the corresponding parameter name and returns the struct.
func (r *ListRUMReplayViewershipHistorySessionsOptionalParameters) WithFilterWatchedAtStart(filterWatchedAtStart int64) *ListRUMReplayViewershipHistorySessionsOptionalParameters {
	r.FilterWatchedAtStart = &filterWatchedAtStart
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListRUMReplayViewershipHistorySessionsOptionalParameters) WithPageNumber(pageNumber int64) *ListRUMReplayViewershipHistorySessionsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithFilterCreatedBy sets the corresponding parameter name and returns the struct.
func (r *ListRUMReplayViewershipHistorySessionsOptionalParameters) WithFilterCreatedBy(filterCreatedBy string) *ListRUMReplayViewershipHistorySessionsOptionalParameters {
	r.FilterCreatedBy = &filterCreatedBy
	return r
}

// WithFilterWatchedAtEnd sets the corresponding parameter name and returns the struct.
func (r *ListRUMReplayViewershipHistorySessionsOptionalParameters) WithFilterWatchedAtEnd(filterWatchedAtEnd int64) *ListRUMReplayViewershipHistorySessionsOptionalParameters {
	r.FilterWatchedAtEnd = &filterWatchedAtEnd
	return r
}

// WithFilterSessionIds sets the corresponding parameter name and returns the struct.
func (r *ListRUMReplayViewershipHistorySessionsOptionalParameters) WithFilterSessionIds(filterSessionIds string) *ListRUMReplayViewershipHistorySessionsOptionalParameters {
	r.FilterSessionIds = &filterSessionIds
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListRUMReplayViewershipHistorySessionsOptionalParameters) WithPageSize(pageSize int64) *ListRUMReplayViewershipHistorySessionsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithFilterApplicationId sets the corresponding parameter name and returns the struct.
func (r *ListRUMReplayViewershipHistorySessionsOptionalParameters) WithFilterApplicationId(filterApplicationId string) *ListRUMReplayViewershipHistorySessionsOptionalParameters {
	r.FilterApplicationId = &filterApplicationId
	return r
}

// ListRUMReplayViewershipHistorySessions List RUM replay viewership history sessions.
// List watched sessions.
func (a *RUMReplayViewershipApi) ListRUMReplayViewershipHistorySessions(ctx _context.Context, o ...ListRUMReplayViewershipHistorySessionsOptionalParameters) (ViewershipHistorySessionArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ViewershipHistorySessionArray
		optionalParams      ListRUMReplayViewershipHistorySessionsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListRUMReplayViewershipHistorySessionsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMReplayViewershipApi.ListRUMReplayViewershipHistorySessions")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/replay/viewership-history/sessions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterWatchedAtStart != nil {
		localVarQueryParams.Add("filter[watched_at][start]", datadog.ParameterToString(*optionalParams.FilterWatchedAtStart, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.FilterCreatedBy != nil {
		localVarQueryParams.Add("filter[created_by]", datadog.ParameterToString(*optionalParams.FilterCreatedBy, ""))
	}
	if optionalParams.FilterWatchedAtEnd != nil {
		localVarQueryParams.Add("filter[watched_at][end]", datadog.ParameterToString(*optionalParams.FilterWatchedAtEnd, ""))
	}
	if optionalParams.FilterSessionIds != nil {
		localVarQueryParams.Add("filter[session_ids]", datadog.ParameterToString(*optionalParams.FilterSessionIds, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.FilterApplicationId != nil {
		localVarQueryParams.Add("filter[application_id]", datadog.ParameterToString(*optionalParams.FilterApplicationId, ""))
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

// NewRUMReplayViewershipApi Returns NewRUMReplayViewershipApi.
func NewRUMReplayViewershipApi(client *datadog.APIClient) *RUMReplayViewershipApi {
	return &RUMReplayViewershipApi{
		Client: client,
	}
}
