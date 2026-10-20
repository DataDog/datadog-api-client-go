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

// AgentHealthApi service type
type AgentHealthApi datadog.Service

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

// NewAgentHealthApi Returns NewAgentHealthApi.
func NewAgentHealthApi(client *datadog.APIClient) *AgentHealthApi {
	return &AgentHealthApi{
		Client: client,
	}
}
