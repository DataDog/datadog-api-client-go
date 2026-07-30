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

// ThreatIntelligenceApi service type
type ThreatIntelligenceApi datadog.Service

// IngestStixThreatIntelOptionalParameters holds optional parameters for IngestStixThreatIntel.
type IngestStixThreatIntelOptionalParameters struct {
	ContentEncoding *STIXContentEncoding
}

// NewIngestStixThreatIntelOptionalParameters creates an empty struct for parameters.
func NewIngestStixThreatIntelOptionalParameters() *IngestStixThreatIntelOptionalParameters {
	this := IngestStixThreatIntelOptionalParameters{}
	return &this
}

// WithContentEncoding sets the corresponding parameter name and returns the struct.
func (r *IngestStixThreatIntelOptionalParameters) WithContentEncoding(contentEncoding STIXContentEncoding) *IngestStixThreatIntelOptionalParameters {
	r.ContentEncoding = &contentEncoding
	return r
}

// IngestStixThreatIntel Ingest STIX threat intelligence.
// Ingest a STIX 2.1 bundle containing threat intelligence indicators. Only indicator objects are supported. Supported indicator patterns contain IPv4 addresses, IPv6 addresses, domain names, or SHA-256 file hashes.
//
// Unsupported objects and patterns increment the `unsupported` counter. Patterns that cannot be parsed increment the `invalid` counter. Processing is best effort, so valid supported indicators in the same bundle are still added.
//
// A successful response means ingestion has completed. Reference-table materialization and enrichment happen asynchronously. Requests are limited to 50 MB and 10 requests per second per API key. Gzip-compressed request bodies are supported.
func (a *ThreatIntelligenceApi) IngestStixThreatIntel(ctx _context.Context, tiVendor string, body STIXBundleRequest, o ...IngestStixThreatIntelOptionalParameters) (STIXIngestResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue STIXIngestResponse
		optionalParams      IngestStixThreatIntelOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type IngestStixThreatIntelOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.IngestStixThreatIntel"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ThreatIntelligenceApi.IngestStixThreatIntel")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/threat-intel/stix"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if datadog.Strlen(tiVendor) < 1 {
		return localVarReturnValue, nil, datadog.ReportError("tiVendor must have at least 1 elements")
	}
	if datadog.Strlen(tiVendor) > 10 {
		return localVarReturnValue, nil, datadog.ReportError("tiVendor must have less than 10 elements")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	localVarHeaderParams["ti_vendor"] = datadog.ParameterToString(tiVendor, "")

	if optionalParams.ContentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = datadog.ParameterToString(*optionalParams.ContentEncoding, "")
	}

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 413 || localVarHTTPResponse.StatusCode == 429 || localVarHTTPResponse.StatusCode == 502 || localVarHTTPResponse.StatusCode == 503 {
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

// NewThreatIntelligenceApi Returns NewThreatIntelligenceApi.
func NewThreatIntelligenceApi(client *datadog.APIClient) *ThreatIntelligenceApi {
	return &ThreatIntelligenceApi{
		Client: client,
	}
}
