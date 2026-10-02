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

// CIVisibilityLogsApi service type
type CIVisibilityLogsApi datadog.Service

// SubmitCILogOptionalParameters holds optional parameters for SubmitCILog.
type SubmitCILogOptionalParameters struct {
	ContentEncoding *CILogContentEncoding
}

// NewSubmitCILogOptionalParameters creates an empty struct for parameters.
func NewSubmitCILogOptionalParameters() *SubmitCILogOptionalParameters {
	this := SubmitCILogOptionalParameters{}
	return &this
}

// WithContentEncoding sets the corresponding parameter name and returns the struct.
func (r *SubmitCILogOptionalParameters) WithContentEncoding(contentEncoding CILogContentEncoding) *SubmitCILogOptionalParameters {
	r.ContentEncoding = &contentEncoding
	return r
}

// SubmitCILog Send CI job logs.
// Send log lines for a CI job over HTTP. See the [CI Visibility Pipelines
// API](https://docs.datadoghq.com/api/latest/ci-visibility-pipelines/send-pipeline-event/) for submitting the
// associated pipeline and job events.
//
// A request can contain one log object or an array of up to 1,000 log objects. The maximum uncompressed request
// body size is 5.1 MiB.
//
// You can stream log lines while a CI job runs or send them after it finishes. After you submit the completed job
// event, 20 seconds without a new log line marks the job's logs as complete. Lines sent after that may not appear.
//
// A job can have up to 128 additional attributes and 256 tags. Additional attributes are top-level fields with
// string, number, boolean, or null values. Nested objects and arrays are rejected. Additional attributes and
// `ddtags` apply to all log lines in the job. If an additional attribute has different values on different lines,
// the first value received is used. Tags supplied on different lines are combined. A job can contain up to
// 2,000,000 log records or 1 GiB of message bytes in total.
//
// To reduce request size, send gzip-compressed JSON with the `Content-Encoding: gzip` header. Retry requests after
// a 408, 429, 500, or 503 response.
func (a *CIVisibilityLogsApi) SubmitCILog(ctx _context.Context, body []CILogItem, o ...SubmitCILogOptionalParameters) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue interface{}
		optionalParams      SubmitCILogOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type SubmitCILogOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.CIVisibilityLogsApi.SubmitCILog")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/cilogs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if len(body) > 1000 {
		return localVarReturnValue, nil, datadog.ReportError("body must have less than 1000 elements")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	if optionalParams.ContentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = datadog.ParameterToString(*optionalParams.ContentEncoding, "")
	}

	// body params
	localVarPostBody = &body
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 408 || localVarHTTPResponse.StatusCode == 413 || localVarHTTPResponse.StatusCode == 429 {
			var v CILogIntakeErrors
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 {
			var v CILogErrors
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

// NewCIVisibilityLogsApi Returns NewCIVisibilityLogsApi.
func NewCIVisibilityLogsApi(client *datadog.APIClient) *CIVisibilityLogsApi {
	return &CIVisibilityLogsApi{
		Client: client,
	}
}
