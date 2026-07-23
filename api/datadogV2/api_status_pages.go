// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/google/uuid"
)

// StatusPagesApi service type
type StatusPagesApi datadog.Service

// CreateBackfilledDegradationOptionalParameters holds optional parameters for CreateBackfilledDegradation.
type CreateBackfilledDegradationOptionalParameters struct {
	Include *string
}

// NewCreateBackfilledDegradationOptionalParameters creates an empty struct for parameters.
func NewCreateBackfilledDegradationOptionalParameters() *CreateBackfilledDegradationOptionalParameters {
	this := CreateBackfilledDegradationOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *CreateBackfilledDegradationOptionalParameters) WithInclude(include string) *CreateBackfilledDegradationOptionalParameters {
	r.Include = &include
	return r
}

// CreateBackfilledDegradation Create backfilled degradation.
// Creates a backfilled degradation with predefined updates.
func (a *StatusPagesApi) CreateBackfilledDegradation(ctx _context.Context, pageId uuid.UUID, body CreateBackfilledDegradationRequest, o ...CreateBackfilledDegradationOptionalParameters) (Degradation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Degradation
		optionalParams      CreateBackfilledDegradationOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type CreateBackfilledDegradationOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.CreateBackfilledDegradation")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradations/backfill"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// CreateBackfilledMaintenanceOptionalParameters holds optional parameters for CreateBackfilledMaintenance.
type CreateBackfilledMaintenanceOptionalParameters struct {
	Include *string
}

// NewCreateBackfilledMaintenanceOptionalParameters creates an empty struct for parameters.
func NewCreateBackfilledMaintenanceOptionalParameters() *CreateBackfilledMaintenanceOptionalParameters {
	this := CreateBackfilledMaintenanceOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *CreateBackfilledMaintenanceOptionalParameters) WithInclude(include string) *CreateBackfilledMaintenanceOptionalParameters {
	r.Include = &include
	return r
}

// CreateBackfilledMaintenance Create backfilled maintenance.
// Creates a backfilled maintenance with predefined updates.
func (a *StatusPagesApi) CreateBackfilledMaintenance(ctx _context.Context, pageId uuid.UUID, body CreateBackfilledMaintenanceRequest, o ...CreateBackfilledMaintenanceOptionalParameters) (Maintenance, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Maintenance
		optionalParams      CreateBackfilledMaintenanceOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type CreateBackfilledMaintenanceOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.CreateBackfilledMaintenance")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/maintenances/backfill"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// CreateComponentOptionalParameters holds optional parameters for CreateComponent.
type CreateComponentOptionalParameters struct {
	Include *string
}

// NewCreateComponentOptionalParameters creates an empty struct for parameters.
func NewCreateComponentOptionalParameters() *CreateComponentOptionalParameters {
	this := CreateComponentOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *CreateComponentOptionalParameters) WithInclude(include string) *CreateComponentOptionalParameters {
	r.Include = &include
	return r
}

// CreateComponent Create component.
// Creates a new component.
func (a *StatusPagesApi) CreateComponent(ctx _context.Context, pageId uuid.UUID, body CreateComponentRequest, o ...CreateComponentOptionalParameters) (StatusPagesComponent, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue StatusPagesComponent
		optionalParams      CreateComponentOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type CreateComponentOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.CreateComponent")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/components"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// CreateDegradationOptionalParameters holds optional parameters for CreateDegradation.
type CreateDegradationOptionalParameters struct {
	NotifySubscribers *bool
	Include           *string
}

// NewCreateDegradationOptionalParameters creates an empty struct for parameters.
func NewCreateDegradationOptionalParameters() *CreateDegradationOptionalParameters {
	this := CreateDegradationOptionalParameters{}
	return &this
}

// WithNotifySubscribers sets the corresponding parameter name and returns the struct.
func (r *CreateDegradationOptionalParameters) WithNotifySubscribers(notifySubscribers bool) *CreateDegradationOptionalParameters {
	r.NotifySubscribers = &notifySubscribers
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *CreateDegradationOptionalParameters) WithInclude(include string) *CreateDegradationOptionalParameters {
	r.Include = &include
	return r
}

// CreateDegradation Create degradation.
// Creates a new degradation.
func (a *StatusPagesApi) CreateDegradation(ctx _context.Context, pageId uuid.UUID, body CreateDegradationRequest, o ...CreateDegradationOptionalParameters) (Degradation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Degradation
		optionalParams      CreateDegradationOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type CreateDegradationOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.CreateDegradation")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradations"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.NotifySubscribers != nil {
		localVarQueryParams.Add("notify_subscribers", datadog.ParameterToString(*optionalParams.NotifySubscribers, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// CreateDegradationTemplateOptionalParameters holds optional parameters for CreateDegradationTemplate.
type CreateDegradationTemplateOptionalParameters struct {
	Include *string
}

// NewCreateDegradationTemplateOptionalParameters creates an empty struct for parameters.
func NewCreateDegradationTemplateOptionalParameters() *CreateDegradationTemplateOptionalParameters {
	this := CreateDegradationTemplateOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *CreateDegradationTemplateOptionalParameters) WithInclude(include string) *CreateDegradationTemplateOptionalParameters {
	r.Include = &include
	return r
}

// CreateDegradationTemplate Create degradation template.
// Creates a new degradation template.
func (a *StatusPagesApi) CreateDegradationTemplate(ctx _context.Context, pageId uuid.UUID, body CreateDegradationTemplateRequest, o ...CreateDegradationTemplateOptionalParameters) (DegradationTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue DegradationTemplate
		optionalParams      CreateDegradationTemplateOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type CreateDegradationTemplateOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.CreateDegradationTemplate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradation_templates"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// CreateMaintenanceOptionalParameters holds optional parameters for CreateMaintenance.
type CreateMaintenanceOptionalParameters struct {
	NotifySubscribers *bool
	Include           *string
}

// NewCreateMaintenanceOptionalParameters creates an empty struct for parameters.
func NewCreateMaintenanceOptionalParameters() *CreateMaintenanceOptionalParameters {
	this := CreateMaintenanceOptionalParameters{}
	return &this
}

// WithNotifySubscribers sets the corresponding parameter name and returns the struct.
func (r *CreateMaintenanceOptionalParameters) WithNotifySubscribers(notifySubscribers bool) *CreateMaintenanceOptionalParameters {
	r.NotifySubscribers = &notifySubscribers
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *CreateMaintenanceOptionalParameters) WithInclude(include string) *CreateMaintenanceOptionalParameters {
	r.Include = &include
	return r
}

// CreateMaintenance Schedule maintenance.
// Schedules a new maintenance.
func (a *StatusPagesApi) CreateMaintenance(ctx _context.Context, pageId uuid.UUID, body CreateMaintenanceRequest, o ...CreateMaintenanceOptionalParameters) (Maintenance, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Maintenance
		optionalParams      CreateMaintenanceOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type CreateMaintenanceOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.CreateMaintenance")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/maintenances"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.NotifySubscribers != nil {
		localVarQueryParams.Add("notify_subscribers", datadog.ParameterToString(*optionalParams.NotifySubscribers, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// CreateMaintenanceTemplateOptionalParameters holds optional parameters for CreateMaintenanceTemplate.
type CreateMaintenanceTemplateOptionalParameters struct {
	Include *string
}

// NewCreateMaintenanceTemplateOptionalParameters creates an empty struct for parameters.
func NewCreateMaintenanceTemplateOptionalParameters() *CreateMaintenanceTemplateOptionalParameters {
	this := CreateMaintenanceTemplateOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *CreateMaintenanceTemplateOptionalParameters) WithInclude(include string) *CreateMaintenanceTemplateOptionalParameters {
	r.Include = &include
	return r
}

// CreateMaintenanceTemplate Create maintenance template.
// Creates a new maintenance template.
func (a *StatusPagesApi) CreateMaintenanceTemplate(ctx _context.Context, pageId uuid.UUID, body CreateMaintenanceTemplateRequest, o ...CreateMaintenanceTemplateOptionalParameters) (MaintenanceTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue MaintenanceTemplate
		optionalParams      CreateMaintenanceTemplateOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type CreateMaintenanceTemplateOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.CreateMaintenanceTemplate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/maintenance_templates"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// CreateStatusPageOptionalParameters holds optional parameters for CreateStatusPage.
type CreateStatusPageOptionalParameters struct {
	Include *string
}

// NewCreateStatusPageOptionalParameters creates an empty struct for parameters.
func NewCreateStatusPageOptionalParameters() *CreateStatusPageOptionalParameters {
	this := CreateStatusPageOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *CreateStatusPageOptionalParameters) WithInclude(include string) *CreateStatusPageOptionalParameters {
	r.Include = &include
	return r
}

// CreateStatusPage Create status page.
// Creates a new status page in an unpublished state. Use the dedicated [publish](#publish-status-page) status page endpoint to publish the page after creation.
func (a *StatusPagesApi) CreateStatusPage(ctx _context.Context, body CreateStatusPageRequest, o ...CreateStatusPageOptionalParameters) (StatusPage, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue StatusPage
		optionalParams      CreateStatusPageOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type CreateStatusPageOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.CreateStatusPage")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// DeleteComponent Delete component.
// Deletes a component by its ID.
func (a *StatusPagesApi) DeleteComponent(ctx _context.Context, pageId uuid.UUID, componentId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.DeleteComponent")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/components/{component_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{component_id}", _neturl.PathEscape(datadog.ParameterToString(componentId, "")))

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

// DeleteDegradation Delete degradation.
// Deletes a degradation by its ID.
func (a *StatusPagesApi) DeleteDegradation(ctx _context.Context, pageId uuid.UUID, degradationId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.DeleteDegradation")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradations/{degradation_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{degradation_id}", _neturl.PathEscape(datadog.ParameterToString(degradationId, "")))

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

// DeleteDegradationTemplate Delete degradation template.
// Deletes a degradation template by its ID (soft delete).
func (a *StatusPagesApi) DeleteDegradationTemplate(ctx _context.Context, pageId uuid.UUID, templateId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.DeleteDegradationTemplate")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradation_templates/{template_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{template_id}", _neturl.PathEscape(datadog.ParameterToString(templateId, "")))

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

// DeleteMaintenanceTemplate Delete maintenance template.
// Deletes a maintenance template by its ID (soft delete).
func (a *StatusPagesApi) DeleteMaintenanceTemplate(ctx _context.Context, pageId uuid.UUID, templateId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.DeleteMaintenanceTemplate")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/maintenance_templates/{template_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{template_id}", _neturl.PathEscape(datadog.ParameterToString(templateId, "")))

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

// DeleteStatusPage Delete status page.
// Deletes a status page by its ID.
func (a *StatusPagesApi) DeleteStatusPage(ctx _context.Context, pageId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.DeleteStatusPage")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

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

// EditDegradationUpdateOptionalParameters holds optional parameters for EditDegradationUpdate.
type EditDegradationUpdateOptionalParameters struct {
	Include *string
}

// NewEditDegradationUpdateOptionalParameters creates an empty struct for parameters.
func NewEditDegradationUpdateOptionalParameters() *EditDegradationUpdateOptionalParameters {
	this := EditDegradationUpdateOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *EditDegradationUpdateOptionalParameters) WithInclude(include string) *EditDegradationUpdateOptionalParameters {
	r.Include = &include
	return r
}

// EditDegradationUpdate Edit degradation update.
// Edits a specific degradation update.
func (a *StatusPagesApi) EditDegradationUpdate(ctx _context.Context, degradationId uuid.UUID, pageId uuid.UUID, updateId uuid.UUID, body PatchDegradationUpdateRequest, o ...EditDegradationUpdateOptionalParameters) (DegradationUpdate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue DegradationUpdate
		optionalParams      EditDegradationUpdateOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type EditDegradationUpdateOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.EditDegradationUpdate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradations/{degradation_id}/updates/{update_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{degradation_id}", _neturl.PathEscape(datadog.ParameterToString(degradationId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{update_id}", _neturl.PathEscape(datadog.ParameterToString(updateId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// GetComponentOptionalParameters holds optional parameters for GetComponent.
type GetComponentOptionalParameters struct {
	Include *string
}

// NewGetComponentOptionalParameters creates an empty struct for parameters.
func NewGetComponentOptionalParameters() *GetComponentOptionalParameters {
	this := GetComponentOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *GetComponentOptionalParameters) WithInclude(include string) *GetComponentOptionalParameters {
	r.Include = &include
	return r
}

// GetComponent Get component.
// Retrieves a specific component by its ID.
func (a *StatusPagesApi) GetComponent(ctx _context.Context, pageId uuid.UUID, componentId uuid.UUID, o ...GetComponentOptionalParameters) (StatusPagesComponent, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue StatusPagesComponent
		optionalParams      GetComponentOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetComponentOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.GetComponent")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/components/{component_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{component_id}", _neturl.PathEscape(datadog.ParameterToString(componentId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// GetDegradationOptionalParameters holds optional parameters for GetDegradation.
type GetDegradationOptionalParameters struct {
	Include *string
}

// NewGetDegradationOptionalParameters creates an empty struct for parameters.
func NewGetDegradationOptionalParameters() *GetDegradationOptionalParameters {
	this := GetDegradationOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *GetDegradationOptionalParameters) WithInclude(include string) *GetDegradationOptionalParameters {
	r.Include = &include
	return r
}

// GetDegradation Get degradation.
// Retrieves a specific degradation by its ID.
func (a *StatusPagesApi) GetDegradation(ctx _context.Context, pageId uuid.UUID, degradationId uuid.UUID, o ...GetDegradationOptionalParameters) (Degradation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue Degradation
		optionalParams      GetDegradationOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetDegradationOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.GetDegradation")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradations/{degradation_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{degradation_id}", _neturl.PathEscape(datadog.ParameterToString(degradationId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// GetDegradationTemplateOptionalParameters holds optional parameters for GetDegradationTemplate.
type GetDegradationTemplateOptionalParameters struct {
	Include *string
}

// NewGetDegradationTemplateOptionalParameters creates an empty struct for parameters.
func NewGetDegradationTemplateOptionalParameters() *GetDegradationTemplateOptionalParameters {
	this := GetDegradationTemplateOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *GetDegradationTemplateOptionalParameters) WithInclude(include string) *GetDegradationTemplateOptionalParameters {
	r.Include = &include
	return r
}

// GetDegradationTemplate Get degradation template.
// Retrieves a specific degradation template by its ID.
func (a *StatusPagesApi) GetDegradationTemplate(ctx _context.Context, pageId uuid.UUID, templateId uuid.UUID, o ...GetDegradationTemplateOptionalParameters) (DegradationTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DegradationTemplate
		optionalParams      GetDegradationTemplateOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetDegradationTemplateOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.GetDegradationTemplate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradation_templates/{template_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{template_id}", _neturl.PathEscape(datadog.ParameterToString(templateId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// GetMaintenanceOptionalParameters holds optional parameters for GetMaintenance.
type GetMaintenanceOptionalParameters struct {
	Include *string
}

// NewGetMaintenanceOptionalParameters creates an empty struct for parameters.
func NewGetMaintenanceOptionalParameters() *GetMaintenanceOptionalParameters {
	this := GetMaintenanceOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *GetMaintenanceOptionalParameters) WithInclude(include string) *GetMaintenanceOptionalParameters {
	r.Include = &include
	return r
}

// GetMaintenance Get maintenance.
// Retrieves a specific maintenance by its ID.
func (a *StatusPagesApi) GetMaintenance(ctx _context.Context, pageId uuid.UUID, maintenanceId uuid.UUID, o ...GetMaintenanceOptionalParameters) (Maintenance, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue Maintenance
		optionalParams      GetMaintenanceOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetMaintenanceOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.GetMaintenance")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/maintenances/{maintenance_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{maintenance_id}", _neturl.PathEscape(datadog.ParameterToString(maintenanceId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// GetMaintenanceTemplateOptionalParameters holds optional parameters for GetMaintenanceTemplate.
type GetMaintenanceTemplateOptionalParameters struct {
	Include *string
}

// NewGetMaintenanceTemplateOptionalParameters creates an empty struct for parameters.
func NewGetMaintenanceTemplateOptionalParameters() *GetMaintenanceTemplateOptionalParameters {
	this := GetMaintenanceTemplateOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *GetMaintenanceTemplateOptionalParameters) WithInclude(include string) *GetMaintenanceTemplateOptionalParameters {
	r.Include = &include
	return r
}

// GetMaintenanceTemplate Get maintenance template.
// Retrieves a specific maintenance template by its ID.
func (a *StatusPagesApi) GetMaintenanceTemplate(ctx _context.Context, pageId uuid.UUID, templateId uuid.UUID, o ...GetMaintenanceTemplateOptionalParameters) (MaintenanceTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MaintenanceTemplate
		optionalParams      GetMaintenanceTemplateOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetMaintenanceTemplateOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.GetMaintenanceTemplate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/maintenance_templates/{template_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{template_id}", _neturl.PathEscape(datadog.ParameterToString(templateId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// GetStatusPageOptionalParameters holds optional parameters for GetStatusPage.
type GetStatusPageOptionalParameters struct {
	Include *string
}

// NewGetStatusPageOptionalParameters creates an empty struct for parameters.
func NewGetStatusPageOptionalParameters() *GetStatusPageOptionalParameters {
	this := GetStatusPageOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *GetStatusPageOptionalParameters) WithInclude(include string) *GetStatusPageOptionalParameters {
	r.Include = &include
	return r
}

// GetStatusPage Get status page.
// Retrieves a specific status page by its ID.
func (a *StatusPagesApi) GetStatusPage(ctx _context.Context, pageId uuid.UUID, o ...GetStatusPageOptionalParameters) (StatusPage, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue StatusPage
		optionalParams      GetStatusPageOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetStatusPageOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.GetStatusPage")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// ListComponentsOptionalParameters holds optional parameters for ListComponents.
type ListComponentsOptionalParameters struct {
	Include *string
}

// NewListComponentsOptionalParameters creates an empty struct for parameters.
func NewListComponentsOptionalParameters() *ListComponentsOptionalParameters {
	this := ListComponentsOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListComponentsOptionalParameters) WithInclude(include string) *ListComponentsOptionalParameters {
	r.Include = &include
	return r
}

// ListComponents List components.
// Lists all components for a status page.
func (a *StatusPagesApi) ListComponents(ctx _context.Context, pageId uuid.UUID, o ...ListComponentsOptionalParameters) (StatusPagesComponentArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue StatusPagesComponentArray
		optionalParams      ListComponentsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListComponentsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.ListComponents")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/components"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// ListDegradationTemplatesOptionalParameters holds optional parameters for ListDegradationTemplates.
type ListDegradationTemplatesOptionalParameters struct {
	Include *string
}

// NewListDegradationTemplatesOptionalParameters creates an empty struct for parameters.
func NewListDegradationTemplatesOptionalParameters() *ListDegradationTemplatesOptionalParameters {
	this := ListDegradationTemplatesOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListDegradationTemplatesOptionalParameters) WithInclude(include string) *ListDegradationTemplatesOptionalParameters {
	r.Include = &include
	return r
}

// ListDegradationTemplates List degradation templates.
// Lists all degradation templates for a status page.
func (a *StatusPagesApi) ListDegradationTemplates(ctx _context.Context, pageId uuid.UUID, o ...ListDegradationTemplatesOptionalParameters) (DegradationTemplateArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DegradationTemplateArray
		optionalParams      ListDegradationTemplatesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListDegradationTemplatesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.ListDegradationTemplates")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradation_templates"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// ListDegradationsOptionalParameters holds optional parameters for ListDegradations.
type ListDegradationsOptionalParameters struct {
	FilterPageId   *string
	PageOffset     *int64
	PageLimit      *int64
	Include        *string
	FilterStatus   *string
	Sort           *string
	FilterSourceId *string
}

// NewListDegradationsOptionalParameters creates an empty struct for parameters.
func NewListDegradationsOptionalParameters() *ListDegradationsOptionalParameters {
	this := ListDegradationsOptionalParameters{}
	return &this
}

// WithFilterPageId sets the corresponding parameter name and returns the struct.
func (r *ListDegradationsOptionalParameters) WithFilterPageId(filterPageId string) *ListDegradationsOptionalParameters {
	r.FilterPageId = &filterPageId
	return r
}

// WithPageOffset sets the corresponding parameter name and returns the struct.
func (r *ListDegradationsOptionalParameters) WithPageOffset(pageOffset int64) *ListDegradationsOptionalParameters {
	r.PageOffset = &pageOffset
	return r
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListDegradationsOptionalParameters) WithPageLimit(pageLimit int64) *ListDegradationsOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListDegradationsOptionalParameters) WithInclude(include string) *ListDegradationsOptionalParameters {
	r.Include = &include
	return r
}

// WithFilterStatus sets the corresponding parameter name and returns the struct.
func (r *ListDegradationsOptionalParameters) WithFilterStatus(filterStatus string) *ListDegradationsOptionalParameters {
	r.FilterStatus = &filterStatus
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListDegradationsOptionalParameters) WithSort(sort string) *ListDegradationsOptionalParameters {
	r.Sort = &sort
	return r
}

// WithFilterSourceId sets the corresponding parameter name and returns the struct.
func (r *ListDegradationsOptionalParameters) WithFilterSourceId(filterSourceId string) *ListDegradationsOptionalParameters {
	r.FilterSourceId = &filterSourceId
	return r
}

// ListDegradations List degradations.
// Lists all degradations for the organization. Optionally filter by status and page.
func (a *StatusPagesApi) ListDegradations(ctx _context.Context, o ...ListDegradationsOptionalParameters) (DegradationArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DegradationArray
		optionalParams      ListDegradationsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListDegradationsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.ListDegradations")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/degradations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterPageId != nil {
		localVarQueryParams.Add("filter[page_id]", datadog.ParameterToString(*optionalParams.FilterPageId, ""))
	}
	if optionalParams.PageOffset != nil {
		localVarQueryParams.Add("page[offset]", datadog.ParameterToString(*optionalParams.PageOffset, ""))
	}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
	if optionalParams.FilterStatus != nil {
		localVarQueryParams.Add("filter[status]", datadog.ParameterToString(*optionalParams.FilterStatus, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.FilterSourceId != nil {
		localVarQueryParams.Add("filter[source_id]", datadog.ParameterToString(*optionalParams.FilterSourceId, ""))
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

// ListMaintenanceTemplatesOptionalParameters holds optional parameters for ListMaintenanceTemplates.
type ListMaintenanceTemplatesOptionalParameters struct {
	Include *string
}

// NewListMaintenanceTemplatesOptionalParameters creates an empty struct for parameters.
func NewListMaintenanceTemplatesOptionalParameters() *ListMaintenanceTemplatesOptionalParameters {
	this := ListMaintenanceTemplatesOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListMaintenanceTemplatesOptionalParameters) WithInclude(include string) *ListMaintenanceTemplatesOptionalParameters {
	r.Include = &include
	return r
}

// ListMaintenanceTemplates List maintenance templates.
// Lists all maintenance templates for a status page.
func (a *StatusPagesApi) ListMaintenanceTemplates(ctx _context.Context, pageId uuid.UUID, o ...ListMaintenanceTemplatesOptionalParameters) (MaintenanceTemplateArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MaintenanceTemplateArray
		optionalParams      ListMaintenanceTemplatesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListMaintenanceTemplatesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.ListMaintenanceTemplates")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/maintenance_templates"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// ListMaintenancesOptionalParameters holds optional parameters for ListMaintenances.
type ListMaintenancesOptionalParameters struct {
	FilterPageId *string
	PageOffset   *int64
	PageLimit    *int64
	Include      *string
	FilterStatus *string
	Sort         *string
}

// NewListMaintenancesOptionalParameters creates an empty struct for parameters.
func NewListMaintenancesOptionalParameters() *ListMaintenancesOptionalParameters {
	this := ListMaintenancesOptionalParameters{}
	return &this
}

// WithFilterPageId sets the corresponding parameter name and returns the struct.
func (r *ListMaintenancesOptionalParameters) WithFilterPageId(filterPageId string) *ListMaintenancesOptionalParameters {
	r.FilterPageId = &filterPageId
	return r
}

// WithPageOffset sets the corresponding parameter name and returns the struct.
func (r *ListMaintenancesOptionalParameters) WithPageOffset(pageOffset int64) *ListMaintenancesOptionalParameters {
	r.PageOffset = &pageOffset
	return r
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListMaintenancesOptionalParameters) WithPageLimit(pageLimit int64) *ListMaintenancesOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListMaintenancesOptionalParameters) WithInclude(include string) *ListMaintenancesOptionalParameters {
	r.Include = &include
	return r
}

// WithFilterStatus sets the corresponding parameter name and returns the struct.
func (r *ListMaintenancesOptionalParameters) WithFilterStatus(filterStatus string) *ListMaintenancesOptionalParameters {
	r.FilterStatus = &filterStatus
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListMaintenancesOptionalParameters) WithSort(sort string) *ListMaintenancesOptionalParameters {
	r.Sort = &sort
	return r
}

// ListMaintenances List maintenances.
// Lists all maintenances for the organization. Optionally filter by status and page.
func (a *StatusPagesApi) ListMaintenances(ctx _context.Context, o ...ListMaintenancesOptionalParameters) (MaintenanceArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MaintenanceArray
		optionalParams      ListMaintenancesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListMaintenancesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.ListMaintenances")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/maintenances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterPageId != nil {
		localVarQueryParams.Add("filter[page_id]", datadog.ParameterToString(*optionalParams.FilterPageId, ""))
	}
	if optionalParams.PageOffset != nil {
		localVarQueryParams.Add("page[offset]", datadog.ParameterToString(*optionalParams.PageOffset, ""))
	}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
	if optionalParams.FilterStatus != nil {
		localVarQueryParams.Add("filter[status]", datadog.ParameterToString(*optionalParams.FilterStatus, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
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

// ListStatusPagesOptionalParameters holds optional parameters for ListStatusPages.
type ListStatusPagesOptionalParameters struct {
	PageOffset         *int64
	PageLimit          *int64
	FilterDomainPrefix *string
	Include            *string
}

// NewListStatusPagesOptionalParameters creates an empty struct for parameters.
func NewListStatusPagesOptionalParameters() *ListStatusPagesOptionalParameters {
	this := ListStatusPagesOptionalParameters{}
	return &this
}

// WithPageOffset sets the corresponding parameter name and returns the struct.
func (r *ListStatusPagesOptionalParameters) WithPageOffset(pageOffset int64) *ListStatusPagesOptionalParameters {
	r.PageOffset = &pageOffset
	return r
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListStatusPagesOptionalParameters) WithPageLimit(pageLimit int64) *ListStatusPagesOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// WithFilterDomainPrefix sets the corresponding parameter name and returns the struct.
func (r *ListStatusPagesOptionalParameters) WithFilterDomainPrefix(filterDomainPrefix string) *ListStatusPagesOptionalParameters {
	r.FilterDomainPrefix = &filterDomainPrefix
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListStatusPagesOptionalParameters) WithInclude(include string) *ListStatusPagesOptionalParameters {
	r.Include = &include
	return r
}

// ListStatusPages List status pages.
// Lists all status pages for the organization.
func (a *StatusPagesApi) ListStatusPages(ctx _context.Context, o ...ListStatusPagesOptionalParameters) (StatusPageArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue StatusPageArray
		optionalParams      ListStatusPagesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListStatusPagesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.ListStatusPages")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageOffset != nil {
		localVarQueryParams.Add("page[offset]", datadog.ParameterToString(*optionalParams.PageOffset, ""))
	}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
	}
	if optionalParams.FilterDomainPrefix != nil {
		localVarQueryParams.Add("filter[domain_prefix]", datadog.ParameterToString(*optionalParams.FilterDomainPrefix, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// PublishStatusPage Publish status page.
// Publishes a status page. For pages of type `public`, makes the status page available on the public internet and requires the `status_pages_public_page_publish` permission. For pages of type `internal`, makes the status page available under the `status-pages/$domain_prefix/view` route within the Datadog organization and requires the `status_pages_internal_page_publish` permission. The `status_pages_settings_write` permission is temporarily honored as we migrate publishing functionality from the update status page endpoint to the publish status page endpoint.
func (a *StatusPagesApi) PublishStatusPage(ctx _context.Context, pageId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.PublishStatusPage")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/publish"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

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

// SoftDeleteDegradationUpdate Soft delete degradation update.
// Soft-deletes a degradation update.
func (a *StatusPagesApi) SoftDeleteDegradationUpdate(ctx _context.Context, degradationId uuid.UUID, pageId uuid.UUID, updateId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.SoftDeleteDegradationUpdate")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradations/{degradation_id}/updates/{update_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{degradation_id}", _neturl.PathEscape(datadog.ParameterToString(degradationId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{update_id}", _neturl.PathEscape(datadog.ParameterToString(updateId, "")))

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

// UnpublishStatusPage Unpublish status page.
// Unpublishes a status page. For pages of type `public`, removes the status page from the public internet and requires the `status_pages_public_page_publish` permission. For pages of type `internal`, removes the `status-pages/$domain_prefix/view` route from the Datadog organization and requires the `status_pages_internal_page_publish` permission. The `status_pages_settings_write` permission is temporarily honored as we migrate unpublishing functionality from the update status page endpoint to the unpublish status page endpoint.
func (a *StatusPagesApi) UnpublishStatusPage(ctx _context.Context, pageId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.UnpublishStatusPage")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/unpublish"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

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

// UpdateComponentOptionalParameters holds optional parameters for UpdateComponent.
type UpdateComponentOptionalParameters struct {
	Include *string
}

// NewUpdateComponentOptionalParameters creates an empty struct for parameters.
func NewUpdateComponentOptionalParameters() *UpdateComponentOptionalParameters {
	this := UpdateComponentOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *UpdateComponentOptionalParameters) WithInclude(include string) *UpdateComponentOptionalParameters {
	r.Include = &include
	return r
}

// UpdateComponent Update component.
// Updates an existing component's attributes.
func (a *StatusPagesApi) UpdateComponent(ctx _context.Context, pageId uuid.UUID, componentId uuid.UUID, body PatchComponentRequest, o ...UpdateComponentOptionalParameters) (StatusPagesComponent, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue StatusPagesComponent
		optionalParams      UpdateComponentOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type UpdateComponentOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.UpdateComponent")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/components/{component_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{component_id}", _neturl.PathEscape(datadog.ParameterToString(componentId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// UpdateDegradationOptionalParameters holds optional parameters for UpdateDegradation.
type UpdateDegradationOptionalParameters struct {
	NotifySubscribers *bool
	Include           *string
}

// NewUpdateDegradationOptionalParameters creates an empty struct for parameters.
func NewUpdateDegradationOptionalParameters() *UpdateDegradationOptionalParameters {
	this := UpdateDegradationOptionalParameters{}
	return &this
}

// WithNotifySubscribers sets the corresponding parameter name and returns the struct.
func (r *UpdateDegradationOptionalParameters) WithNotifySubscribers(notifySubscribers bool) *UpdateDegradationOptionalParameters {
	r.NotifySubscribers = &notifySubscribers
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *UpdateDegradationOptionalParameters) WithInclude(include string) *UpdateDegradationOptionalParameters {
	r.Include = &include
	return r
}

// UpdateDegradation Update degradation.
// Updates an existing degradation's attributes.
func (a *StatusPagesApi) UpdateDegradation(ctx _context.Context, pageId uuid.UUID, degradationId uuid.UUID, body PatchDegradationRequest, o ...UpdateDegradationOptionalParameters) (Degradation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue Degradation
		optionalParams      UpdateDegradationOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type UpdateDegradationOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.UpdateDegradation")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradations/{degradation_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{degradation_id}", _neturl.PathEscape(datadog.ParameterToString(degradationId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.NotifySubscribers != nil {
		localVarQueryParams.Add("notify_subscribers", datadog.ParameterToString(*optionalParams.NotifySubscribers, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// UpdateDegradationTemplateOptionalParameters holds optional parameters for UpdateDegradationTemplate.
type UpdateDegradationTemplateOptionalParameters struct {
	Include *string
}

// NewUpdateDegradationTemplateOptionalParameters creates an empty struct for parameters.
func NewUpdateDegradationTemplateOptionalParameters() *UpdateDegradationTemplateOptionalParameters {
	this := UpdateDegradationTemplateOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *UpdateDegradationTemplateOptionalParameters) WithInclude(include string) *UpdateDegradationTemplateOptionalParameters {
	r.Include = &include
	return r
}

// UpdateDegradationTemplate Update degradation template.
// Updates an existing degradation template's attributes.
func (a *StatusPagesApi) UpdateDegradationTemplate(ctx _context.Context, templateId uuid.UUID, pageId uuid.UUID, body PatchDegradationTemplateRequest, o ...UpdateDegradationTemplateOptionalParameters) (DegradationTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue DegradationTemplate
		optionalParams      UpdateDegradationTemplateOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type UpdateDegradationTemplateOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.UpdateDegradationTemplate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/degradation_templates/{template_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{template_id}", _neturl.PathEscape(datadog.ParameterToString(templateId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// UpdateMaintenanceOptionalParameters holds optional parameters for UpdateMaintenance.
type UpdateMaintenanceOptionalParameters struct {
	NotifySubscribers *bool
	Include           *string
}

// NewUpdateMaintenanceOptionalParameters creates an empty struct for parameters.
func NewUpdateMaintenanceOptionalParameters() *UpdateMaintenanceOptionalParameters {
	this := UpdateMaintenanceOptionalParameters{}
	return &this
}

// WithNotifySubscribers sets the corresponding parameter name and returns the struct.
func (r *UpdateMaintenanceOptionalParameters) WithNotifySubscribers(notifySubscribers bool) *UpdateMaintenanceOptionalParameters {
	r.NotifySubscribers = &notifySubscribers
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *UpdateMaintenanceOptionalParameters) WithInclude(include string) *UpdateMaintenanceOptionalParameters {
	r.Include = &include
	return r
}

// UpdateMaintenance Update maintenance.
// Updates an existing maintenance's attributes.
func (a *StatusPagesApi) UpdateMaintenance(ctx _context.Context, pageId uuid.UUID, maintenanceId uuid.UUID, body PatchMaintenanceRequest, o ...UpdateMaintenanceOptionalParameters) (Maintenance, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue Maintenance
		optionalParams      UpdateMaintenanceOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type UpdateMaintenanceOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.UpdateMaintenance")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/maintenances/{maintenance_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{maintenance_id}", _neturl.PathEscape(datadog.ParameterToString(maintenanceId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.NotifySubscribers != nil {
		localVarQueryParams.Add("notify_subscribers", datadog.ParameterToString(*optionalParams.NotifySubscribers, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// UpdateMaintenanceTemplateOptionalParameters holds optional parameters for UpdateMaintenanceTemplate.
type UpdateMaintenanceTemplateOptionalParameters struct {
	Include *string
}

// NewUpdateMaintenanceTemplateOptionalParameters creates an empty struct for parameters.
func NewUpdateMaintenanceTemplateOptionalParameters() *UpdateMaintenanceTemplateOptionalParameters {
	this := UpdateMaintenanceTemplateOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *UpdateMaintenanceTemplateOptionalParameters) WithInclude(include string) *UpdateMaintenanceTemplateOptionalParameters {
	r.Include = &include
	return r
}

// UpdateMaintenanceTemplate Update maintenance template.
// Updates an existing maintenance template's attributes.
func (a *StatusPagesApi) UpdateMaintenanceTemplate(ctx _context.Context, pageId uuid.UUID, templateId uuid.UUID, body PatchMaintenanceTemplateRequest, o ...UpdateMaintenanceTemplateOptionalParameters) (MaintenanceTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue MaintenanceTemplate
		optionalParams      UpdateMaintenanceTemplateOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type UpdateMaintenanceTemplateOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.UpdateMaintenanceTemplate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}/maintenance_templates/{template_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{template_id}", _neturl.PathEscape(datadog.ParameterToString(templateId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// UpdateStatusPageOptionalParameters holds optional parameters for UpdateStatusPage.
type UpdateStatusPageOptionalParameters struct {
	DeleteSubscribers *bool
	Include           *string
}

// NewUpdateStatusPageOptionalParameters creates an empty struct for parameters.
func NewUpdateStatusPageOptionalParameters() *UpdateStatusPageOptionalParameters {
	this := UpdateStatusPageOptionalParameters{}
	return &this
}

// WithDeleteSubscribers sets the corresponding parameter name and returns the struct.
func (r *UpdateStatusPageOptionalParameters) WithDeleteSubscribers(deleteSubscribers bool) *UpdateStatusPageOptionalParameters {
	r.DeleteSubscribers = &deleteSubscribers
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *UpdateStatusPageOptionalParameters) WithInclude(include string) *UpdateStatusPageOptionalParameters {
	r.Include = &include
	return r
}

// UpdateStatusPage Update status page.
// Updates an existing status page's attributes. **Note**: Publishing and unpublishing via the `enabled` property will be deprecated on this endpoint. Use the dedicated [publish](#publish-status-page) and [unpublish](#unpublish-status-page) status page endpoints instead.
func (a *StatusPagesApi) UpdateStatusPage(ctx _context.Context, pageId uuid.UUID, body PatchStatusPageRequest, o ...UpdateStatusPageOptionalParameters) (StatusPage, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue StatusPage
		optionalParams      UpdateStatusPageOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type UpdateStatusPageOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.StatusPagesApi.UpdateStatusPage")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/statuspages/{page_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{page_id}", _neturl.PathEscape(datadog.ParameterToString(pageId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.DeleteSubscribers != nil {
		localVarQueryParams.Add("delete_subscribers", datadog.ParameterToString(*optionalParams.DeleteSubscribers, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
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

// NewStatusPagesApi Returns NewStatusPagesApi.
func NewStatusPagesApi(client *datadog.APIClient) *StatusPagesApi {
	return &StatusPagesApi{
		Client: client,
	}
}
