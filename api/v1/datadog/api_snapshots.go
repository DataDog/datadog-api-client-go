// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/api/common"
)

// SnapshotsApiService service type
type SnapshotsApiService common.Service

type apiGetGraphSnapshotRequest struct {
	ctx         _context.Context
	ApiService  *SnapshotsApiService
	start       *int64
	end         *int64
	metricQuery *string
	eventQuery  *string
	graphDef    *string
	title       *string
	height      *int64
	width       *int64
}

// GetGraphSnapshotOptionalParameters holds optional parameters for GetGraphSnapshot.
type GetGraphSnapshotOptionalParameters struct {
	MetricQuery *string
	EventQuery  *string
	GraphDef    *string
	Title       *string
	Height      *int64
	Width       *int64
}

// NewGetGraphSnapshotOptionalParameters creates an empty struct for parameters.
func NewGetGraphSnapshotOptionalParameters() *GetGraphSnapshotOptionalParameters {
	this := GetGraphSnapshotOptionalParameters{}
	return &this
}

// WithMetricQuery sets the corresponding parameter name and returns the struct.
func (r *GetGraphSnapshotOptionalParameters) WithMetricQuery(metricQuery string) *GetGraphSnapshotOptionalParameters {
	r.MetricQuery = &metricQuery
	return r
}

// WithEventQuery sets the corresponding parameter name and returns the struct.
func (r *GetGraphSnapshotOptionalParameters) WithEventQuery(eventQuery string) *GetGraphSnapshotOptionalParameters {
	r.EventQuery = &eventQuery
	return r
}

// WithGraphDef sets the corresponding parameter name and returns the struct.
func (r *GetGraphSnapshotOptionalParameters) WithGraphDef(graphDef string) *GetGraphSnapshotOptionalParameters {
	r.GraphDef = &graphDef
	return r
}

// WithTitle sets the corresponding parameter name and returns the struct.
func (r *GetGraphSnapshotOptionalParameters) WithTitle(title string) *GetGraphSnapshotOptionalParameters {
	r.Title = &title
	return r
}

// WithHeight sets the corresponding parameter name and returns the struct.
func (r *GetGraphSnapshotOptionalParameters) WithHeight(height int64) *GetGraphSnapshotOptionalParameters {
	r.Height = &height
	return r
}

// WithWidth sets the corresponding parameter name and returns the struct.
func (r *GetGraphSnapshotOptionalParameters) WithWidth(width int64) *GetGraphSnapshotOptionalParameters {
	r.Width = &width
	return r
}

func (a *SnapshotsApiService) buildGetGraphSnapshotRequest(ctx _context.Context, start int64, end int64, o ...GetGraphSnapshotOptionalParameters) (apiGetGraphSnapshotRequest, error) {
	req := apiGetGraphSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
		start:      &start,
		end:        &end,
	}

	if len(o) > 1 {
		return req, common.ReportError("only one argument of type GetGraphSnapshotOptionalParameters is allowed")
	}

	if o != nil {
		req.metricQuery = o[0].MetricQuery
		req.eventQuery = o[0].EventQuery
		req.graphDef = o[0].GraphDef
		req.title = o[0].Title
		req.height = o[0].Height
		req.width = o[0].Width
	}
	return req, nil
}

// GetGraphSnapshot Take graph snapshots.
// Take graph snapshots.
// **Note**: When a snapshot is created, there is some delay before it is available.
func (a *SnapshotsApiService) GetGraphSnapshot(ctx _context.Context, start int64, end int64, o ...GetGraphSnapshotOptionalParameters) (GraphSnapshot, *_nethttp.Response, error) {
	req, err := a.buildGetGraphSnapshotRequest(ctx, start, end, o...)
	if err != nil {
		var localVarReturnValue GraphSnapshot
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getGraphSnapshotExecute(req)
}

// getGraphSnapshotExecute executes the request.
func (a *SnapshotsApiService) getGraphSnapshotExecute(r apiGetGraphSnapshotRequest) (GraphSnapshot, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue GraphSnapshot
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.SnapshotsApiService.GetGraphSnapshot")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/graph/snapshot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.start == nil {
		return localVarReturnValue, nil, common.ReportError("start is required and must be specified")
	}
	if r.end == nil {
		return localVarReturnValue, nil, common.ReportError("end is required and must be specified")
	}
	localVarQueryParams.Add("start", common.ParameterToString(*r.start, ""))
	localVarQueryParams.Add("end", common.ParameterToString(*r.end, ""))
	if r.metricQuery != nil {
		localVarQueryParams.Add("metric_query", common.ParameterToString(*r.metricQuery, ""))
	}
	if r.eventQuery != nil {
		localVarQueryParams.Add("event_query", common.ParameterToString(*r.eventQuery, ""))
	}
	if r.graphDef != nil {
		localVarQueryParams.Add("graph_def", common.ParameterToString(*r.graphDef, ""))
	}
	if r.title != nil {
		localVarQueryParams.Add("title", common.ParameterToString(*r.title, ""))
	}
	if r.height != nil {
		localVarQueryParams.Add("height", common.ParameterToString(*r.height, ""))
	}
	if r.width != nil {
		localVarQueryParams.Add("width", common.ParameterToString(*r.width, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
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
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// SnapshotsApi Returns new SnapshotsApi service.
func SnapshotsApi(client *common.APIClient) *SnapshotsApiService {
	return &SnapshotsApiService{
		Client: client,
	}
}
