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
	"reflect"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/google/uuid"
)

// RumTeamsOwnershipApi service type
type RumTeamsOwnershipApi datadog.Service

// CreateTeamsOwnershipMapping Create a teams ownership mapping.
// Create a teams ownership mapping for your organization.
// Returns the teams ownership mapping object from the request body when the request is successful.
func (a *RumTeamsOwnershipApi) CreateTeamsOwnershipMapping(ctx _context.Context, body TeamsOwnershipMappingCreateRequest) (TeamsOwnershipMappingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue TeamsOwnershipMappingResponse
	)

	operationId := "v2.CreateTeamsOwnershipMapping"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumTeamsOwnershipApi.CreateTeamsOwnershipMapping")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/config/teams-ownership/mappings"

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

// CreateTeamsOwnershipMappingsBatch Bulk create and remove teams ownership mappings.
// Add and remove teams ownership mappings for your organization in a single atomic request, following
// the JSON:API [atomic operations extension](https://jsonapi.org/ext/atomic/).
// Operations are applied together: if any operation is invalid, none of the operations are applied.
// Add operations are processed before remove operations, so results may not appear in the same
// order as the request.
func (a *RumTeamsOwnershipApi) CreateTeamsOwnershipMappingsBatch(ctx _context.Context, body TeamsOwnershipMappingBatchRequest) (TeamsOwnershipMappingBatchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue TeamsOwnershipMappingBatchResponse
	)

	operationId := "v2.CreateTeamsOwnershipMappingsBatch"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumTeamsOwnershipApi.CreateTeamsOwnershipMappingsBatch")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/config/teams-ownership/mappings/operations"

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
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

// DeleteTeamsOwnershipMapping Delete a teams ownership mapping.
// Delete a specific teams ownership mapping from your organization.
func (a *RumTeamsOwnershipApi) DeleteTeamsOwnershipMapping(ctx _context.Context, id string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	operationId := "v2.DeleteTeamsOwnershipMapping"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumTeamsOwnershipApi.DeleteTeamsOwnershipMapping")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/config/teams-ownership/mappings/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// GetTeamsOwnershipMapping Get a teams ownership mapping.
// Get a specific teams ownership mapping from your organization.
func (a *RumTeamsOwnershipApi) GetTeamsOwnershipMapping(ctx _context.Context, id string) (TeamsOwnershipMappingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue TeamsOwnershipMappingResponse
	)

	operationId := "v2.GetTeamsOwnershipMapping"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumTeamsOwnershipApi.GetTeamsOwnershipMapping")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/config/teams-ownership/mappings/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// ListTeamsOwnershipMappingsOptionalParameters holds optional parameters for ListTeamsOwnershipMappings.
type ListTeamsOwnershipMappingsOptionalParameters struct {
	FilterViewName      *[]string
	FilterTeamHandle    *[]string
	FilterApplicationId *[]uuid.UUID
	FilterService       *[]string
}

// NewListTeamsOwnershipMappingsOptionalParameters creates an empty struct for parameters.
func NewListTeamsOwnershipMappingsOptionalParameters() *ListTeamsOwnershipMappingsOptionalParameters {
	this := ListTeamsOwnershipMappingsOptionalParameters{}
	return &this
}

// WithFilterViewName sets the corresponding parameter name and returns the struct.
func (r *ListTeamsOwnershipMappingsOptionalParameters) WithFilterViewName(filterViewName []string) *ListTeamsOwnershipMappingsOptionalParameters {
	r.FilterViewName = &filterViewName
	return r
}

// WithFilterTeamHandle sets the corresponding parameter name and returns the struct.
func (r *ListTeamsOwnershipMappingsOptionalParameters) WithFilterTeamHandle(filterTeamHandle []string) *ListTeamsOwnershipMappingsOptionalParameters {
	r.FilterTeamHandle = &filterTeamHandle
	return r
}

// WithFilterApplicationId sets the corresponding parameter name and returns the struct.
func (r *ListTeamsOwnershipMappingsOptionalParameters) WithFilterApplicationId(filterApplicationId []uuid.UUID) *ListTeamsOwnershipMappingsOptionalParameters {
	r.FilterApplicationId = &filterApplicationId
	return r
}

// WithFilterService sets the corresponding parameter name and returns the struct.
func (r *ListTeamsOwnershipMappingsOptionalParameters) WithFilterService(filterService []string) *ListTeamsOwnershipMappingsOptionalParameters {
	r.FilterService = &filterService
	return r
}

// ListTeamsOwnershipMappings List teams ownership mappings.
// Get the list of teams ownership mappings for your organization, optionally filtered.
func (a *RumTeamsOwnershipApi) ListTeamsOwnershipMappings(ctx _context.Context, o ...ListTeamsOwnershipMappingsOptionalParameters) (TeamsOwnershipMappingsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue TeamsOwnershipMappingsResponse
		optionalParams      ListTeamsOwnershipMappingsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListTeamsOwnershipMappingsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListTeamsOwnershipMappings"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumTeamsOwnershipApi.ListTeamsOwnershipMappings")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/config/teams-ownership/mappings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterViewName != nil {
		t := *optionalParams.FilterViewName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[view_name]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[view_name]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterTeamHandle != nil {
		t := *optionalParams.FilterTeamHandle
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[team_handle]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[team_handle]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterApplicationId != nil {
		t := *optionalParams.FilterApplicationId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[application_id]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[application_id]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterService != nil {
		t := *optionalParams.FilterService
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[service]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[service]", datadog.ParameterToString(t, "multi"))
		}
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

// ListTeamsOwnershipRulesOptionalParameters holds optional parameters for ListTeamsOwnershipRules.
type ListTeamsOwnershipRulesOptionalParameters struct {
	FilterViewName      *[]string
	FilterTeamHandle    *[]string
	FilterApplicationId *[]uuid.UUID
	FilterService       *[]string
}

// NewListTeamsOwnershipRulesOptionalParameters creates an empty struct for parameters.
func NewListTeamsOwnershipRulesOptionalParameters() *ListTeamsOwnershipRulesOptionalParameters {
	this := ListTeamsOwnershipRulesOptionalParameters{}
	return &this
}

// WithFilterViewName sets the corresponding parameter name and returns the struct.
func (r *ListTeamsOwnershipRulesOptionalParameters) WithFilterViewName(filterViewName []string) *ListTeamsOwnershipRulesOptionalParameters {
	r.FilterViewName = &filterViewName
	return r
}

// WithFilterTeamHandle sets the corresponding parameter name and returns the struct.
func (r *ListTeamsOwnershipRulesOptionalParameters) WithFilterTeamHandle(filterTeamHandle []string) *ListTeamsOwnershipRulesOptionalParameters {
	r.FilterTeamHandle = &filterTeamHandle
	return r
}

// WithFilterApplicationId sets the corresponding parameter name and returns the struct.
func (r *ListTeamsOwnershipRulesOptionalParameters) WithFilterApplicationId(filterApplicationId []uuid.UUID) *ListTeamsOwnershipRulesOptionalParameters {
	r.FilterApplicationId = &filterApplicationId
	return r
}

// WithFilterService sets the corresponding parameter name and returns the struct.
func (r *ListTeamsOwnershipRulesOptionalParameters) WithFilterService(filterService []string) *ListTeamsOwnershipRulesOptionalParameters {
	r.FilterService = &filterService
	return r
}

// ListTeamsOwnershipRules List teams ownership rules.
// Get the list of teams ownership rules for your organization, optionally filtered.
// Rules group the underlying mappings by `view_name`, `application_id`, `service`, and `match_type`,
// collapsing every team that owns the same view into a single entry.
func (a *RumTeamsOwnershipApi) ListTeamsOwnershipRules(ctx _context.Context, o ...ListTeamsOwnershipRulesOptionalParameters) (TeamsOwnershipRulesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue TeamsOwnershipRulesResponse
		optionalParams      ListTeamsOwnershipRulesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListTeamsOwnershipRulesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListTeamsOwnershipRules"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumTeamsOwnershipApi.ListTeamsOwnershipRules")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/config/teams-ownership/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterViewName != nil {
		t := *optionalParams.FilterViewName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[view_name]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[view_name]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterTeamHandle != nil {
		t := *optionalParams.FilterTeamHandle
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[team_handle]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[team_handle]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterApplicationId != nil {
		t := *optionalParams.FilterApplicationId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[application_id]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[application_id]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterService != nil {
		t := *optionalParams.FilterService
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[service]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[service]", datadog.ParameterToString(t, "multi"))
		}
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

// NewRumTeamsOwnershipApi Returns NewRumTeamsOwnershipApi.
func NewRumTeamsOwnershipApi(client *datadog.APIClient) *RumTeamsOwnershipApi {
	return &RumTeamsOwnershipApi{
		Client: client,
	}
}
