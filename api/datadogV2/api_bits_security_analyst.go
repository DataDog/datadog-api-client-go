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

// BitsSecurityAnalystApi service type
type BitsSecurityAnalystApi datadog.Service

// CountSecurityAgentInvestigationsOptionalParameters holds optional parameters for CountSecurityAgentInvestigations.
type CountSecurityAgentInvestigationsOptionalParameters struct {
	FilterTitle          *string
	FilterFrom           *string
	FilterTo             *string
	FilterVerdict        *string
	FilterRuleId         *string
	FilterRuleType       *string
	FilterSignalSeverity *string
	FilterSignalSource   *string
	FilterTags           *string
	FilterEntity         *string
}

// NewCountSecurityAgentInvestigationsOptionalParameters creates an empty struct for parameters.
func NewCountSecurityAgentInvestigationsOptionalParameters() *CountSecurityAgentInvestigationsOptionalParameters {
	this := CountSecurityAgentInvestigationsOptionalParameters{}
	return &this
}

// WithFilterTitle sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterTitle(filterTitle string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterTitle = &filterTitle
	return r
}

// WithFilterFrom sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterFrom(filterFrom string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterFrom = &filterFrom
	return r
}

// WithFilterTo sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterTo(filterTo string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterTo = &filterTo
	return r
}

// WithFilterVerdict sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterVerdict(filterVerdict string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterVerdict = &filterVerdict
	return r
}

// WithFilterRuleId sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterRuleId(filterRuleId string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterRuleId = &filterRuleId
	return r
}

// WithFilterRuleType sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterRuleType(filterRuleType string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterRuleType = &filterRuleType
	return r
}

// WithFilterSignalSeverity sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterSignalSeverity(filterSignalSeverity string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterSignalSeverity = &filterSignalSeverity
	return r
}

// WithFilterSignalSource sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterSignalSource(filterSignalSource string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterSignalSource = &filterSignalSource
	return r
}

// WithFilterTags sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterTags(filterTags string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterTags = &filterTags
	return r
}

// WithFilterEntity sets the corresponding parameter name and returns the struct.
func (r *CountSecurityAgentInvestigationsOptionalParameters) WithFilterEntity(filterEntity string) *CountSecurityAgentInvestigationsOptionalParameters {
	r.FilterEntity = &filterEntity
	return r
}

// CountSecurityAgentInvestigations Count Bits Security Analyst investigations.
// Count completed Bits Security Analyst investigations by the values of one field. Buckets are sorted by
// ascending count. With `field=entity` or `field=signal_source`, an investigation counts in the bucket of
// each of its entities or sources.
func (a *BitsSecurityAnalystApi) CountSecurityAgentInvestigations(ctx _context.Context, field SecurityAgentInvestigationCountField, o ...CountSecurityAgentInvestigationsOptionalParameters) (SecurityAgentInvestigationCountsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityAgentInvestigationCountsResponse
		optionalParams      CountSecurityAgentInvestigationsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type CountSecurityAgentInvestigationsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.CountSecurityAgentInvestigations"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.BitsSecurityAnalystApi.CountSecurityAgentInvestigations")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/security_agent/investigations/count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("field", datadog.ParameterToString(field, ""))
	if optionalParams.FilterTitle != nil {
		localVarQueryParams.Add("filter[title]", datadog.ParameterToString(*optionalParams.FilterTitle, ""))
	}
	if optionalParams.FilterFrom != nil {
		localVarQueryParams.Add("filter[from]", datadog.ParameterToString(*optionalParams.FilterFrom, ""))
	}
	if optionalParams.FilterTo != nil {
		localVarQueryParams.Add("filter[to]", datadog.ParameterToString(*optionalParams.FilterTo, ""))
	}
	if optionalParams.FilterVerdict != nil {
		localVarQueryParams.Add("filter[verdict]", datadog.ParameterToString(*optionalParams.FilterVerdict, ""))
	}
	if optionalParams.FilterRuleId != nil {
		localVarQueryParams.Add("filter[rule_id]", datadog.ParameterToString(*optionalParams.FilterRuleId, ""))
	}
	if optionalParams.FilterRuleType != nil {
		localVarQueryParams.Add("filter[rule_type]", datadog.ParameterToString(*optionalParams.FilterRuleType, ""))
	}
	if optionalParams.FilterSignalSeverity != nil {
		localVarQueryParams.Add("filter[signal_severity]", datadog.ParameterToString(*optionalParams.FilterSignalSeverity, ""))
	}
	if optionalParams.FilterSignalSource != nil {
		localVarQueryParams.Add("filter[signal_source]", datadog.ParameterToString(*optionalParams.FilterSignalSource, ""))
	}
	if optionalParams.FilterTags != nil {
		localVarQueryParams.Add("filter[tags]", datadog.ParameterToString(*optionalParams.FilterTags, ""))
	}
	if optionalParams.FilterEntity != nil {
		localVarQueryParams.Add("filter[entity]", datadog.ParameterToString(*optionalParams.FilterEntity, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 {
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

// GetSecurityAgentInvestigationOptionalParameters holds optional parameters for GetSecurityAgentInvestigation.
type GetSecurityAgentInvestigationOptionalParameters struct {
	FieldsInvestigations *[]SecurityAgentInvestigationField
}

// NewGetSecurityAgentInvestigationOptionalParameters creates an empty struct for parameters.
func NewGetSecurityAgentInvestigationOptionalParameters() *GetSecurityAgentInvestigationOptionalParameters {
	this := GetSecurityAgentInvestigationOptionalParameters{}
	return &this
}

// WithFieldsInvestigations sets the corresponding parameter name and returns the struct.
func (r *GetSecurityAgentInvestigationOptionalParameters) WithFieldsInvestigations(fieldsInvestigations []SecurityAgentInvestigationField) *GetSecurityAgentInvestigationOptionalParameters {
	r.FieldsInvestigations = &fieldsInvestigations
	return r
}

// GetSecurityAgentInvestigation Get a Bits Security Analyst investigation.
// Get a Bits Security Analyst investigation by its ID, with its verdict, conclusion, evidence, steps, and
// response suggestions. While the investigation is running, the response has a `Retry-After` header.
func (a *BitsSecurityAnalystApi) GetSecurityAgentInvestigation(ctx _context.Context, investigationId string, o ...GetSecurityAgentInvestigationOptionalParameters) (SecurityAgentInvestigationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityAgentInvestigationResponse
		optionalParams      GetSecurityAgentInvestigationOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetSecurityAgentInvestigationOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.GetSecurityAgentInvestigation"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.BitsSecurityAnalystApi.GetSecurityAgentInvestigation")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/security_agent/investigations/{investigation_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{investigation_id}", _neturl.PathEscape(datadog.ParameterToString(investigationId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FieldsInvestigations != nil {
		localVarQueryParams.Add("fields[investigations]", datadog.ParameterToString(*optionalParams.FieldsInvestigations, "csv"))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// ListSecurityAgentInvestigationsOptionalParameters holds optional parameters for ListSecurityAgentInvestigations.
type ListSecurityAgentInvestigationsOptionalParameters struct {
	FilterSignalId       *string
	FilterTitle          *string
	FilterFrom           *string
	FilterTo             *string
	FilterVerdict        *string
	FilterState          *string
	FilterRuleId         *string
	FilterRuleType       *string
	FilterSignalSeverity *string
	FilterSignalSource   *string
	FilterTags           *string
	FilterEntity         *string
	Sort                 *string
	PageNumber           *int64
	PageSize             *int64
}

// NewListSecurityAgentInvestigationsOptionalParameters creates an empty struct for parameters.
func NewListSecurityAgentInvestigationsOptionalParameters() *ListSecurityAgentInvestigationsOptionalParameters {
	this := ListSecurityAgentInvestigationsOptionalParameters{}
	return &this
}

// WithFilterSignalId sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterSignalId(filterSignalId string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterSignalId = &filterSignalId
	return r
}

// WithFilterTitle sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterTitle(filterTitle string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterTitle = &filterTitle
	return r
}

// WithFilterFrom sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterFrom(filterFrom string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterFrom = &filterFrom
	return r
}

// WithFilterTo sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterTo(filterTo string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterTo = &filterTo
	return r
}

// WithFilterVerdict sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterVerdict(filterVerdict string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterVerdict = &filterVerdict
	return r
}

// WithFilterState sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterState(filterState string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterState = &filterState
	return r
}

// WithFilterRuleId sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterRuleId(filterRuleId string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterRuleId = &filterRuleId
	return r
}

// WithFilterRuleType sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterRuleType(filterRuleType string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterRuleType = &filterRuleType
	return r
}

// WithFilterSignalSeverity sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterSignalSeverity(filterSignalSeverity string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterSignalSeverity = &filterSignalSeverity
	return r
}

// WithFilterSignalSource sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterSignalSource(filterSignalSource string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterSignalSource = &filterSignalSource
	return r
}

// WithFilterTags sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterTags(filterTags string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterTags = &filterTags
	return r
}

// WithFilterEntity sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithFilterEntity(filterEntity string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.FilterEntity = &filterEntity
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithSort(sort string) *ListSecurityAgentInvestigationsOptionalParameters {
	r.Sort = &sort
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithPageNumber(pageNumber int64) *ListSecurityAgentInvestigationsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListSecurityAgentInvestigationsOptionalParameters) WithPageSize(pageSize int64) *ListSecurityAgentInvestigationsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// ListSecurityAgentInvestigations List Bits Security Analyst investigations.
// List Bits Security Analyst investigations.
//
// Without `filter[signal_id]`, returns the latest investigation of each signal, filtered, sorted, and paginated.
// With `filter[signal_id]`, returns every investigation of that signal, newest first. When the signal has no
// investigation, the first page carries `meta.reason`, which says why.
func (a *BitsSecurityAnalystApi) ListSecurityAgentInvestigations(ctx _context.Context, o ...ListSecurityAgentInvestigationsOptionalParameters) (SecurityAgentInvestigationsListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityAgentInvestigationsListResponse
		optionalParams      ListSecurityAgentInvestigationsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListSecurityAgentInvestigationsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListSecurityAgentInvestigations"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.BitsSecurityAnalystApi.ListSecurityAgentInvestigations")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/security_agent/investigations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterSignalId != nil {
		localVarQueryParams.Add("filter[signal_id]", datadog.ParameterToString(*optionalParams.FilterSignalId, ""))
	}
	if optionalParams.FilterTitle != nil {
		localVarQueryParams.Add("filter[title]", datadog.ParameterToString(*optionalParams.FilterTitle, ""))
	}
	if optionalParams.FilterFrom != nil {
		localVarQueryParams.Add("filter[from]", datadog.ParameterToString(*optionalParams.FilterFrom, ""))
	}
	if optionalParams.FilterTo != nil {
		localVarQueryParams.Add("filter[to]", datadog.ParameterToString(*optionalParams.FilterTo, ""))
	}
	if optionalParams.FilterVerdict != nil {
		localVarQueryParams.Add("filter[verdict]", datadog.ParameterToString(*optionalParams.FilterVerdict, ""))
	}
	if optionalParams.FilterState != nil {
		localVarQueryParams.Add("filter[state]", datadog.ParameterToString(*optionalParams.FilterState, ""))
	}
	if optionalParams.FilterRuleId != nil {
		localVarQueryParams.Add("filter[rule_id]", datadog.ParameterToString(*optionalParams.FilterRuleId, ""))
	}
	if optionalParams.FilterRuleType != nil {
		localVarQueryParams.Add("filter[rule_type]", datadog.ParameterToString(*optionalParams.FilterRuleType, ""))
	}
	if optionalParams.FilterSignalSeverity != nil {
		localVarQueryParams.Add("filter[signal_severity]", datadog.ParameterToString(*optionalParams.FilterSignalSeverity, ""))
	}
	if optionalParams.FilterSignalSource != nil {
		localVarQueryParams.Add("filter[signal_source]", datadog.ParameterToString(*optionalParams.FilterSignalSource, ""))
	}
	if optionalParams.FilterTags != nil {
		localVarQueryParams.Add("filter[tags]", datadog.ParameterToString(*optionalParams.FilterTags, ""))
	}
	if optionalParams.FilterEntity != nil {
		localVarQueryParams.Add("filter[entity]", datadog.ParameterToString(*optionalParams.FilterEntity, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 {
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

// NewBitsSecurityAnalystApi Returns NewBitsSecurityAnalystApi.
func NewBitsSecurityAnalystApi(client *datadog.APIClient) *BitsSecurityAnalystApi {
	return &BitsSecurityAnalystApi{
		Client: client,
	}
}
