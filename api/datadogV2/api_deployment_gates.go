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
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/google/uuid"
)

// DeploymentGatesApi service type
type DeploymentGatesApi datadog.Service

// CreateDeploymentGate Create deployment gate.
// Endpoint to create a deployment gate.
func (a *DeploymentGatesApi) CreateDeploymentGate(ctx _context.Context, body CreateDeploymentGateParams) (DeploymentGateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue DeploymentGateResponse
	)

	operationId := "v2.CreateDeploymentGate"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.CreateDeploymentGate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates"

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// CreateDeploymentRule Create deployment rule.
// Endpoint to create a deployment rule. A gate for the rule must already exist.
func (a *DeploymentGatesApi) CreateDeploymentRule(ctx _context.Context, gateId string, body CreateDeploymentRuleParams) (DeploymentRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue DeploymentRuleResponse
	)

	operationId := "v2.CreateDeploymentRule"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.CreateDeploymentRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/{gate_id}/rules"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{gate_id}", _neturl.PathEscape(datadog.ParameterToString(gateId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// DeleteDeploymentGate Delete deployment gate.
// Endpoint to delete a deployment gate. Rules associated with the gate are also deleted.
func (a *DeploymentGatesApi) DeleteDeploymentGate(ctx _context.Context, id string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	operationId := "v2.DeleteDeploymentGate"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.DeleteDeploymentGate")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/{id}"
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v HTTPCDGatesNotFoundResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// DeleteDeploymentRule Delete deployment rule.
// Endpoint to delete a deployment rule.
func (a *DeploymentGatesApi) DeleteDeploymentRule(ctx _context.Context, gateId string, id string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	operationId := "v2.DeleteDeploymentRule"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.DeleteDeploymentRule")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/{gate_id}/rules/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{gate_id}", _neturl.PathEscape(datadog.ParameterToString(gateId, "")))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v HTTPCDGatesNotFoundResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// GetDeploymentGate Get deployment gate.
// Endpoint to get a deployment gate.
func (a *DeploymentGatesApi) GetDeploymentGate(ctx _context.Context, id string) (DeploymentGateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DeploymentGateResponse
	)

	operationId := "v2.GetDeploymentGate"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.GetDeploymentGate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/{id}"
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v HTTPCDGatesNotFoundResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// GetDeploymentGateRules Get rules for a deployment gate.
// Endpoint to get rules for a deployment gate.
func (a *DeploymentGatesApi) GetDeploymentGateRules(ctx _context.Context, gateId string) (DeploymentGateRulesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DeploymentGateRulesResponse
	)

	operationId := "v2.GetDeploymentGateRules"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.GetDeploymentGateRules")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/{gate_id}/rules"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{gate_id}", _neturl.PathEscape(datadog.ParameterToString(gateId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// GetDeploymentGatesEvaluationResult Get a deployment gate evaluation result.
// Retrieves the result of a deployment gate evaluation by its evaluation ID.
// If the evaluation is still in progress, `data.attributes.gate_status` will be `in_progress`;
// continue polling until it returns `pass` or `fail`.
// Polling every 10-20 seconds is recommended.
// The endpoint may return a 404 if called too soon after triggering; retry after a few seconds.
func (a *DeploymentGatesApi) GetDeploymentGatesEvaluationResult(ctx _context.Context, id uuid.UUID) (DeploymentGatesEvaluationResultResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DeploymentGatesEvaluationResultResponse
	)

	operationId := "v2.GetDeploymentGatesEvaluationResult"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.GetDeploymentGatesEvaluationResult")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployments/gates/evaluation/{id}"
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v HTTPCDGatesNotFoundResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// GetDeploymentRule Get deployment rule.
// Endpoint to get a deployment rule.
func (a *DeploymentGatesApi) GetDeploymentRule(ctx _context.Context, gateId string, id string) (DeploymentRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DeploymentRuleResponse
	)

	operationId := "v2.GetDeploymentRule"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.GetDeploymentRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/{gate_id}/rules/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{gate_id}", _neturl.PathEscape(datadog.ParameterToString(gateId, "")))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v HTTPCDRulesNotFoundResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// ListDeploymentGateEvaluationsOptionalParameters holds optional parameters for ListDeploymentGateEvaluations.
type ListDeploymentGateEvaluationsOptionalParameters struct {
	FilterFrom         *time.Time
	FilterTo           *time.Time
	FilterService      *[]string
	FilterEnv          *[]string
	FilterIdentifier   *[]string
	FilterStatus       *[]DeploymentGatesEvaluationResultResponseAttributesGateStatus
	FilterDryRun       *bool
	FilterEvaluationId *uuid.UUID
	FilterGateId       *uuid.UUID
	FilterVersion      *[]string
	PageSize           *int64
	PageCursor         *string
}

// NewListDeploymentGateEvaluationsOptionalParameters creates an empty struct for parameters.
func NewListDeploymentGateEvaluationsOptionalParameters() *ListDeploymentGateEvaluationsOptionalParameters {
	this := ListDeploymentGateEvaluationsOptionalParameters{}
	return &this
}

// WithFilterFrom sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterFrom(filterFrom time.Time) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterFrom = &filterFrom
	return r
}

// WithFilterTo sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterTo(filterTo time.Time) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterTo = &filterTo
	return r
}

// WithFilterService sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterService(filterService []string) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterService = &filterService
	return r
}

// WithFilterEnv sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterEnv(filterEnv []string) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterEnv = &filterEnv
	return r
}

// WithFilterIdentifier sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterIdentifier(filterIdentifier []string) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterIdentifier = &filterIdentifier
	return r
}

// WithFilterStatus sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterStatus(filterStatus []DeploymentGatesEvaluationResultResponseAttributesGateStatus) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterStatus = &filterStatus
	return r
}

// WithFilterDryRun sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterDryRun(filterDryRun bool) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterDryRun = &filterDryRun
	return r
}

// WithFilterEvaluationId sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterEvaluationId(filterEvaluationId uuid.UUID) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterEvaluationId = &filterEvaluationId
	return r
}

// WithFilterGateId sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterGateId(filterGateId uuid.UUID) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterGateId = &filterGateId
	return r
}

// WithFilterVersion sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithFilterVersion(filterVersion []string) *ListDeploymentGateEvaluationsOptionalParameters {
	r.FilterVersion = &filterVersion
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithPageSize(pageSize int64) *ListDeploymentGateEvaluationsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageCursor sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGateEvaluationsOptionalParameters) WithPageCursor(pageCursor string) *ListDeploymentGateEvaluationsOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}

// ListDeploymentGateEvaluations List deployment gate evaluations.
// Returns deployment gate evaluations started in a maximum 30-day window (the default is the previous 24 hours).
// Results are ordered by start time, newest first.
// In-progress state is near-real-time and mutable. Finished state is eventually consistent.
func (a *DeploymentGatesApi) ListDeploymentGateEvaluations(ctx _context.Context, o ...ListDeploymentGateEvaluationsOptionalParameters) (DeploymentGateEvaluationsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DeploymentGateEvaluationsResponse
		optionalParams      ListDeploymentGateEvaluationsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListDeploymentGateEvaluationsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListDeploymentGateEvaluations"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.ListDeploymentGateEvaluations")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/evaluations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterFrom != nil {
		localVarQueryParams.Add("filter[from]", datadog.ParameterToString(*optionalParams.FilterFrom, ""))
	}
	if optionalParams.FilterTo != nil {
		localVarQueryParams.Add("filter[to]", datadog.ParameterToString(*optionalParams.FilterTo, ""))
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
	if optionalParams.FilterEnv != nil {
		t := *optionalParams.FilterEnv
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[env]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[env]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterIdentifier != nil {
		t := *optionalParams.FilterIdentifier
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[identifier]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[identifier]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterStatus != nil {
		t := *optionalParams.FilterStatus
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[status]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[status]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterDryRun != nil {
		localVarQueryParams.Add("filter[dry_run]", datadog.ParameterToString(*optionalParams.FilterDryRun, ""))
	}
	if optionalParams.FilterEvaluationId != nil {
		localVarQueryParams.Add("filter[evaluation_id]", datadog.ParameterToString(*optionalParams.FilterEvaluationId, ""))
	}
	if optionalParams.FilterGateId != nil {
		localVarQueryParams.Add("filter[gate_id]", datadog.ParameterToString(*optionalParams.FilterGateId, ""))
	}
	if optionalParams.FilterVersion != nil {
		t := *optionalParams.FilterVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[version]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[version]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageCursor != nil {
		localVarQueryParams.Add("page[cursor]", datadog.ParameterToString(*optionalParams.PageCursor, ""))
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
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

// ListDeploymentGateEvaluationsWithPagination provides a paginated version of ListDeploymentGateEvaluations returning a channel with all items.
func (a *DeploymentGatesApi) ListDeploymentGateEvaluationsWithPagination(ctx _context.Context, o ...ListDeploymentGateEvaluationsOptionalParameters) (<-chan datadog.PaginationResult[DeploymentGateEvaluationData], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(20)
	if len(o) == 0 {
		o = append(o, ListDeploymentGateEvaluationsOptionalParameters{})
	}
	if o[0].PageSize != nil {
		pageSize_ = *o[0].PageSize
	}
	o[0].PageSize = &pageSize_

	items := make(chan datadog.PaginationResult[DeploymentGateEvaluationData], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListDeploymentGateEvaluations(ctx, o...)
			if err != nil {
				var returnItem DeploymentGateEvaluationData
				items <- datadog.PaginationResult[DeploymentGateEvaluationData]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[DeploymentGateEvaluationData]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) == 0 {
				break
			}
			cursorMeta, ok := resp.GetMetaOk()
			if !ok {
				break
			}
			cursorMetaPage, ok := cursorMeta.GetPageOk()
			if !ok {
				break
			}
			cursorMetaPageNextCursor, ok := cursorMetaPage.GetNextCursorOk()
			if !ok {
				break
			}

			o[0].PageCursor = cursorMetaPageNextCursor
		}
		close(items)
	}()
	return items, cancel
}

// ListDeploymentGatesOptionalParameters holds optional parameters for ListDeploymentGates.
type ListDeploymentGatesOptionalParameters struct {
	FilterService    *string
	FilterEnv        *string
	FilterIdentifier *string
	FilterDryRun     *bool
	PageCursor       *string
	PageSize         *int64
}

// NewListDeploymentGatesOptionalParameters creates an empty struct for parameters.
func NewListDeploymentGatesOptionalParameters() *ListDeploymentGatesOptionalParameters {
	this := ListDeploymentGatesOptionalParameters{}
	return &this
}

// WithFilterService sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGatesOptionalParameters) WithFilterService(filterService string) *ListDeploymentGatesOptionalParameters {
	r.FilterService = &filterService
	return r
}

// WithFilterEnv sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGatesOptionalParameters) WithFilterEnv(filterEnv string) *ListDeploymentGatesOptionalParameters {
	r.FilterEnv = &filterEnv
	return r
}

// WithFilterIdentifier sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGatesOptionalParameters) WithFilterIdentifier(filterIdentifier string) *ListDeploymentGatesOptionalParameters {
	r.FilterIdentifier = &filterIdentifier
	return r
}

// WithFilterDryRun sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGatesOptionalParameters) WithFilterDryRun(filterDryRun bool) *ListDeploymentGatesOptionalParameters {
	r.FilterDryRun = &filterDryRun
	return r
}

// WithPageCursor sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGatesOptionalParameters) WithPageCursor(pageCursor string) *ListDeploymentGatesOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentGatesOptionalParameters) WithPageSize(pageSize int64) *ListDeploymentGatesOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// ListDeploymentGates Get all deployment gates.
// Returns a paginated list of all deployment gates for the organization.
// Use `page[cursor]` and `page[size]` query parameters to paginate through results.
func (a *DeploymentGatesApi) ListDeploymentGates(ctx _context.Context, o ...ListDeploymentGatesOptionalParameters) (DeploymentGatesListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DeploymentGatesListResponse
		optionalParams      ListDeploymentGatesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListDeploymentGatesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListDeploymentGates"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.ListDeploymentGates")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterService != nil {
		localVarQueryParams.Add("filter[service]", datadog.ParameterToString(*optionalParams.FilterService, ""))
	}
	if optionalParams.FilterEnv != nil {
		localVarQueryParams.Add("filter[env]", datadog.ParameterToString(*optionalParams.FilterEnv, ""))
	}
	if optionalParams.FilterIdentifier != nil {
		localVarQueryParams.Add("filter[identifier]", datadog.ParameterToString(*optionalParams.FilterIdentifier, ""))
	}
	if optionalParams.FilterDryRun != nil {
		localVarQueryParams.Add("filter[dry_run]", datadog.ParameterToString(*optionalParams.FilterDryRun, ""))
	}
	if optionalParams.PageCursor != nil {
		localVarQueryParams.Add("page[cursor]", datadog.ParameterToString(*optionalParams.PageCursor, ""))
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
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// ListDeploymentRuleEvaluationsOptionalParameters holds optional parameters for ListDeploymentRuleEvaluations.
type ListDeploymentRuleEvaluationsOptionalParameters struct {
	FilterFrom             *time.Time
	FilterTo               *time.Time
	FilterGateEvaluationId *uuid.UUID
	FilterEvaluationId     *uuid.UUID
	FilterGateId           *uuid.UUID
	FilterRuleId           *uuid.UUID
	FilterService          *[]string
	FilterEnv              *[]string
	FilterIdentifier       *[]string
	FilterVersion          *[]string
	FilterStatus           *[]DeploymentGatesEvaluationResultResponseAttributesGateStatus
	FilterType             *[]DeploymentGateRuleEvaluationType
	FilterDryRun           *bool
	FilterGateDryRun       *bool
	FilterName             *[]string
	PageSize               *int64
	PageCursor             *string
}

// NewListDeploymentRuleEvaluationsOptionalParameters creates an empty struct for parameters.
func NewListDeploymentRuleEvaluationsOptionalParameters() *ListDeploymentRuleEvaluationsOptionalParameters {
	this := ListDeploymentRuleEvaluationsOptionalParameters{}
	return &this
}

// WithFilterFrom sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterFrom(filterFrom time.Time) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterFrom = &filterFrom
	return r
}

// WithFilterTo sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterTo(filterTo time.Time) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterTo = &filterTo
	return r
}

// WithFilterGateEvaluationId sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterGateEvaluationId(filterGateEvaluationId uuid.UUID) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterGateEvaluationId = &filterGateEvaluationId
	return r
}

// WithFilterEvaluationId sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterEvaluationId(filterEvaluationId uuid.UUID) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterEvaluationId = &filterEvaluationId
	return r
}

// WithFilterGateId sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterGateId(filterGateId uuid.UUID) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterGateId = &filterGateId
	return r
}

// WithFilterRuleId sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterRuleId(filterRuleId uuid.UUID) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterRuleId = &filterRuleId
	return r
}

// WithFilterService sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterService(filterService []string) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterService = &filterService
	return r
}

// WithFilterEnv sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterEnv(filterEnv []string) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterEnv = &filterEnv
	return r
}

// WithFilterIdentifier sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterIdentifier(filterIdentifier []string) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterIdentifier = &filterIdentifier
	return r
}

// WithFilterVersion sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterVersion(filterVersion []string) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterVersion = &filterVersion
	return r
}

// WithFilterStatus sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterStatus(filterStatus []DeploymentGatesEvaluationResultResponseAttributesGateStatus) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterStatus = &filterStatus
	return r
}

// WithFilterType sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterType(filterType []DeploymentGateRuleEvaluationType) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterType = &filterType
	return r
}

// WithFilterDryRun sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterDryRun(filterDryRun bool) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterDryRun = &filterDryRun
	return r
}

// WithFilterGateDryRun sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterGateDryRun(filterGateDryRun bool) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterGateDryRun = &filterGateDryRun
	return r
}

// WithFilterName sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithFilterName(filterName []string) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.FilterName = &filterName
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithPageSize(pageSize int64) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageCursor sets the corresponding parameter name and returns the struct.
func (r *ListDeploymentRuleEvaluationsOptionalParameters) WithPageCursor(pageCursor string) *ListDeploymentRuleEvaluationsOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}

// ListDeploymentRuleEvaluations List deployment gate rule evaluations.
// Returns rule evaluations whose gate evaluation started in a maximum 30-day window (the default is the previous 24 hours).
// Filter by gate, rule, gate evaluation, or rule evaluation ID; omit IDs for cross-evaluation searches.
// Results are ordered by start time, newest first.
// In-progress state is near-real-time and mutable. Finished state is eventually consistent.
// Gate-level and rule-level dry-run states are independent.
// Pagination is deterministic but not snapshot isolated; clients should deduplicate by rule evaluation ID.
func (a *DeploymentGatesApi) ListDeploymentRuleEvaluations(ctx _context.Context, o ...ListDeploymentRuleEvaluationsOptionalParameters) (DeploymentGateRuleEvaluationsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DeploymentGateRuleEvaluationsResponse
		optionalParams      ListDeploymentRuleEvaluationsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListDeploymentRuleEvaluationsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListDeploymentRuleEvaluations"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.ListDeploymentRuleEvaluations")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/evaluations/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterFrom != nil {
		localVarQueryParams.Add("filter[from]", datadog.ParameterToString(*optionalParams.FilterFrom, ""))
	}
	if optionalParams.FilterTo != nil {
		localVarQueryParams.Add("filter[to]", datadog.ParameterToString(*optionalParams.FilterTo, ""))
	}
	if optionalParams.FilterGateEvaluationId != nil {
		localVarQueryParams.Add("filter[gate_evaluation_id]", datadog.ParameterToString(*optionalParams.FilterGateEvaluationId, ""))
	}
	if optionalParams.FilterEvaluationId != nil {
		localVarQueryParams.Add("filter[evaluation_id]", datadog.ParameterToString(*optionalParams.FilterEvaluationId, ""))
	}
	if optionalParams.FilterGateId != nil {
		localVarQueryParams.Add("filter[gate_id]", datadog.ParameterToString(*optionalParams.FilterGateId, ""))
	}
	if optionalParams.FilterRuleId != nil {
		localVarQueryParams.Add("filter[rule_id]", datadog.ParameterToString(*optionalParams.FilterRuleId, ""))
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
	if optionalParams.FilterEnv != nil {
		t := *optionalParams.FilterEnv
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[env]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[env]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterIdentifier != nil {
		t := *optionalParams.FilterIdentifier
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[identifier]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[identifier]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterVersion != nil {
		t := *optionalParams.FilterVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[version]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[version]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterStatus != nil {
		t := *optionalParams.FilterStatus
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[status]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[status]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterType != nil {
		t := *optionalParams.FilterType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[type]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[type]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterDryRun != nil {
		localVarQueryParams.Add("filter[dry_run]", datadog.ParameterToString(*optionalParams.FilterDryRun, ""))
	}
	if optionalParams.FilterGateDryRun != nil {
		localVarQueryParams.Add("filter[gate_dry_run]", datadog.ParameterToString(*optionalParams.FilterGateDryRun, ""))
	}
	if optionalParams.FilterName != nil {
		t := *optionalParams.FilterName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[name]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[name]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageCursor != nil {
		localVarQueryParams.Add("page[cursor]", datadog.ParameterToString(*optionalParams.PageCursor, ""))
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
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

// ListDeploymentRuleEvaluationsWithPagination provides a paginated version of ListDeploymentRuleEvaluations returning a channel with all items.
func (a *DeploymentGatesApi) ListDeploymentRuleEvaluationsWithPagination(ctx _context.Context, o ...ListDeploymentRuleEvaluationsOptionalParameters) (<-chan datadog.PaginationResult[DeploymentGateRuleEvaluationData], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(50)
	if len(o) == 0 {
		o = append(o, ListDeploymentRuleEvaluationsOptionalParameters{})
	}
	if o[0].PageSize != nil {
		pageSize_ = *o[0].PageSize
	}
	o[0].PageSize = &pageSize_

	items := make(chan datadog.PaginationResult[DeploymentGateRuleEvaluationData], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListDeploymentRuleEvaluations(ctx, o...)
			if err != nil {
				var returnItem DeploymentGateRuleEvaluationData
				items <- datadog.PaginationResult[DeploymentGateRuleEvaluationData]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[DeploymentGateRuleEvaluationData]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) == 0 {
				break
			}
			cursorMeta, ok := resp.GetMetaOk()
			if !ok {
				break
			}
			cursorMetaPage, ok := cursorMeta.GetPageOk()
			if !ok {
				break
			}
			cursorMetaPageNextCursor, ok := cursorMetaPage.GetNextCursorOk()
			if !ok {
				break
			}

			o[0].PageCursor = cursorMetaPageNextCursor
		}
		close(items)
	}()
	return items, cancel
}

// TriggerDeploymentGatesEvaluation Trigger a deployment gate evaluation.
// Triggers an asynchronous deployment gate evaluation for the given service and environment.
// Returns an evaluation ID that can be used to poll for the result via the
// `GET /api/v2/deployments/gates/evaluation/{id}` endpoint.
//
// When the `configuration` attribute is provided, rules are evaluated inline from that configuration
// and no pre-configured gate is required. When `configuration` is omitted, rules are resolved from the
// gate pre-configured for the given service and environment through the Datadog UI, API, or Terraform.
func (a *DeploymentGatesApi) TriggerDeploymentGatesEvaluation(ctx _context.Context, body DeploymentGatesEvaluationRequest) (DeploymentGatesEvaluationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue DeploymentGatesEvaluationResponse
	)

	operationId := "v2.TriggerDeploymentGatesEvaluation"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.TriggerDeploymentGatesEvaluation")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployments/gates/evaluation"

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v HTTPCDGatesNotFoundResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// UpdateDeploymentGate Update deployment gate.
// Endpoint to update a deployment gate.
func (a *DeploymentGatesApi) UpdateDeploymentGate(ctx _context.Context, id string, body UpdateDeploymentGateParams) (DeploymentGateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue DeploymentGateResponse
	)

	operationId := "v2.UpdateDeploymentGate"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.UpdateDeploymentGate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v HTTPCDGatesNotFoundResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// UpdateDeploymentRule Update deployment rule.
// Endpoint to update a deployment rule.
func (a *DeploymentGatesApi) UpdateDeploymentRule(ctx _context.Context, gateId string, id string, body UpdateDeploymentRuleParams) (DeploymentRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue DeploymentRuleResponse
	)

	operationId := "v2.UpdateDeploymentRule"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DeploymentGatesApi.UpdateDeploymentRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/deployment_gates/{gate_id}/rules/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{gate_id}", _neturl.PathEscape(datadog.ParameterToString(gateId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v HTTPCDGatesBadRequestResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v HTTPCDRulesNotFoundResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPCIAppErrors
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

// NewDeploymentGatesApi Returns NewDeploymentGatesApi.
func NewDeploymentGatesApi(client *datadog.APIClient) *DeploymentGatesApi {
	return &DeploymentGatesApi{
		Client: client,
	}
}
