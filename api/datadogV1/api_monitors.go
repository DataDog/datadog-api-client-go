// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorsApi service type
type MonitorsApi datadog.Service

// CheckCanDeleteMonitor Check if a monitor can be deleted.
// Check if the given monitors can be deleted.
func (a *MonitorsApi) CheckCanDeleteMonitor(ctx _context.Context, monitorIds []int64) (CheckCanDeleteMonitorResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CheckCanDeleteMonitorResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.CheckCanDeleteMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/can_delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("monitor_ids", datadog.ParameterToString(monitorIds, "csv"))
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v CheckCanDeleteMonitorResponse
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

// CreateMonitor Create a monitor.
// Create a monitor using the specified options.
//
// #### Monitor Types
//
// The type of monitor chosen from:
//
// - anomaly: `query alert`
// - APM: `query alert` or `trace-analytics alert`
// - composite: `composite`
// - custom: `service check`
// - forecast: `query alert`
// - host: `service check`
// - integration: `query alert` or `service check`
// - live process: `process alert`
// - logs: `log alert`
// - metric: `query alert`
// - network: `service check`
// - outlier: `query alert`
// - process: `service check`
// - rum: `rum alert`
// - SLO: `slo alert`
// - watchdog: `event-v2 alert`
// - event-v2: `event-v2 alert`
// - audit: `audit alert`
// - error-tracking: `error-tracking alert`
// - database-monitoring: `database-monitoring alert`
// - network-performance: `network-performance alert`
// - cloud cost: `cost alert`
//
// **Notes**:
// - Synthetic monitors are created through the Synthetics API. See the [Synthetics API](https://docs.datadoghq.com/api/latest/synthetics/) documentation for more information.
// - Log monitors require an unscoped App Key.
//
// #### Query Types
//
// ##### Metric Alert Query
//
// Example: `time_aggr(time_window):space_aggr:metric{tags} [by {key}] operator #`
//
// - `time_aggr`: avg, sum, max, min, change, or pct_change
// - `time_window`: `last_#m` (with `#` between 1 and 10080 depending on the monitor type) or `last_#h`(with `#` between 1 and 168 depending on the monitor type) or `last_1d`, or `last_1w`
// - `space_aggr`: avg, sum, min, or max
// - `tags`: one or more tags (comma-separated), or *
// - `key`: a 'key' in key:value tag syntax; defines a separate alert for each tag in the group (multi-alert)
// - `operator`: <, <=, >, >=, ==, or !=
// - `#`: an integer or decimal number used to set the threshold
//
// If you are using the `_change_` or `_pct_change_` time aggregator, instead use `change_aggr(time_aggr(time_window),
// timeshift):space_aggr:metric{tags} [by {key}] operator #` with:
//
// - `change_aggr` change, pct_change
// - `time_aggr` avg, sum, max, min [Learn more](https://docs.datadoghq.com/monitors/create/types/#define-the-conditions)
// - `time_window` last\_#m (between 1 and 2880 depending on the monitor type), last\_#h (between 1 and 48 depending on the monitor type), or last_#d (1 or 2)
// - `timeshift` #m_ago (5, 10, 15, or 30), #h_ago (1, 2, or 4), or 1d_ago
//
// Use this to create an outlier monitor using the following query:
// `avg(last_30m):outliers(avg:system.cpu.user{role:es-events-data} by {host}, 'dbscan', 7) > 0`
//
// ##### Service Check Query
//
// Example: `"check".over(tags).last(count).by(group).count_by_status()`
//
// - `check` name of the check, for example `datadog.agent.up`
// - `tags` one or more quoted tags (comma-separated), or "*". for example: `.over("env:prod", "role:db")`; `over` cannot be blank.
// - `count` must be at greater than or equal to your max threshold (defined in the `options`). It is limited to 100.
// For example, if you've specified to notify on 1 critical, 3 ok, and 2 warn statuses, `count` should be at least 3.
// - `group` must be specified for check monitors. Per-check grouping is already explicitly known for some service checks.
// For example, Postgres integration monitors are tagged by `db`, `host`, and `port`, and Network monitors by `host`, `instance`, and `url`. See [Service Checks](https://docs.datadoghq.com/api/latest/service-checks/) documentation for more information.
//
// ##### Event Alert Query
//
// **Note:** The Event Alert Query has been replaced by the Event V2 Alert Query. For more information, see the [Event Migration guide](https://docs.datadoghq.com/service_management/events/guides/migrating_to_new_events_features/).
//
// ##### Event V2 Alert Query
//
// Example: `events(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg` and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// ##### Process Alert Query
//
// Example: `processes(search).over(tags).rollup('count').last(timeframe) operator #`
//
// - `search` free text search string for querying processes.
// Matching processes match results on the [Live Processes](https://docs.datadoghq.com/infrastructure/process/?tab=linuxwindows) page.
// - `tags` one or more tags (comma-separated)
// - `timeframe` the timeframe to roll up the counts. Examples: 10m, 4h. Supported timeframes: s, m, h and d
// - `operator` <, <=, >, >=, ==, or !=
// - `#` an integer or decimal number used to set the threshold
//
// ##### Logs Alert Query
//
// Example: `logs(query).index(index_name).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `index_name` For multi-index organizations, the log index in which the request is performed.
// - `rollup_method` The stats roll-up method - supports `count`, `avg` and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// ##### Composite Query
//
// Example: `12345 && 67890`, where `12345` and `67890` are the IDs of non-composite monitors
//
// * `name` [*required*, *default* = **dynamic, based on query**]: The name of the alert.
// * `message` [*required*, *default* = **dynamic, based on query**]: A message to include with notifications for this monitor.
// Email notifications can be sent to specific users by using the same '@username' notation as events.
// * `tags` [*optional*, *default* = **empty list**]: A list of tags to associate with your monitor.
// When getting all monitor details via the API, use the `monitor_tags` argument to filter results by these tags.
// It is only available via the API and isn't visible or editable in the Datadog UI.
//
// ##### SLO Alert Query
//
// Example: `error_budget("slo_id").over("time_window") operator #`
//
// - `slo_id`: The alphanumeric SLO ID of the SLO you are configuring the alert for.
// - `time_window`: The time window of the SLO target you wish to alert on. Valid options: `7d`, `30d`, `90d`.
// - `operator`: `>=` or `>`
//
// ##### Audit Alert Query
//
// Example: `audits(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg` and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// ##### CI Pipelines Alert Query
//
// Example: `ci-pipelines(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg`, and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// ##### CI Tests Alert Query
//
// Example: `ci-tests(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg`, and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// ##### Error Tracking Alert Query
//
// "New issue" example: `error-tracking(query).source(issue_source).new().rollup(rollup_method[, measure]).by(group_by).last(time_window) operator #`
// "High impact issue" example: `error-tracking(query).source(issue_source).impact().rollup(rollup_method[, measure]).by(group_by).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `issue_source` The issue source - supports `all`, `browser`, `mobile` and `backend` and defaults to `all` if omitted.
// - `rollup_method` The stats roll-up method - supports `count`, `avg`, and `cardinality` and defaults to `count` if omitted.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `group by` Comma-separated list of attributes to group by - should contain at least `issue.id`.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// **Database Monitoring Alert Query**
//
// Example: `database-monitoring(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg`, and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// **Network Performance Alert Query**
//
// Example: `network-performance(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg`, and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// **Cost Alert Query**
//
// Example: `formula(query).timeframe_type(time_window).function(parameter) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `timeframe_type` The timeframe type to evaluate the cost
//   - for `forecast` supports `current`
//   - for `change`, `anomaly`, `threshold` supports `last`
//
// - `time_window` - supports daily roll-up e.g. `7d`
// - `function` - [optional, defaults to `threshold` monitor if omitted] supports `change`, `anomaly`, `forecast`
// - `parameter` Specify the parameter of the type
//   - for `change`:
//   - supports `relative`, `absolute`
//   - [optional] supports `#`, where `#` is an integer or decimal number used to set the threshold
//   - for `anomaly`:
//   - supports `direction=both`, `direction=above`, `direction=below`
//   - [optional] supports `threshold=#`, where `#` is an integer or decimal number used to set the threshold
//
// - `operator`
//   - for `threshold` supports `<`, `<=`, `>`, `>=`, `==`, or `!=`
//   - for `change` supports `>`, `<`
//   - for `anomaly` supports `>=`
//   - for `forecast` supports `>`
//
// - `#` an integer or decimal number used to set the threshold.
func (a *MonitorsApi) CreateMonitor(ctx _context.Context, body Monitor) (Monitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Monitor
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.CreateMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor"

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

// DeleteMonitorOptionalParameters holds optional parameters for DeleteMonitor.
type DeleteMonitorOptionalParameters struct {
	Force *string
}

// NewDeleteMonitorOptionalParameters creates an empty struct for parameters.
func NewDeleteMonitorOptionalParameters() *DeleteMonitorOptionalParameters {
	this := DeleteMonitorOptionalParameters{}
	return &this
}

// WithForce sets the corresponding parameter name and returns the struct.
func (r *DeleteMonitorOptionalParameters) WithForce(force string) *DeleteMonitorOptionalParameters {
	r.Force = &force
	return r
}

// DeleteMonitor Delete a monitor.
// Delete the specified monitor
func (a *MonitorsApi) DeleteMonitor(ctx _context.Context, monitorId int64, o ...DeleteMonitorOptionalParameters) (DeletedMonitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodDelete
		localVarPostBody    interface{}
		localVarReturnValue DeletedMonitor
		optionalParams      DeleteMonitorOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type DeleteMonitorOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.DeleteMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/{monitor_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{monitor_id}", _neturl.PathEscape(datadog.ParameterToString(monitorId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Force != nil {
		localVarQueryParams.Add("force", datadog.ParameterToString(*optionalParams.Force, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// GetMonitorOptionalParameters holds optional parameters for GetMonitor.
type GetMonitorOptionalParameters struct {
	GroupStates   *string
	WithDowntimes *bool
}

// NewGetMonitorOptionalParameters creates an empty struct for parameters.
func NewGetMonitorOptionalParameters() *GetMonitorOptionalParameters {
	this := GetMonitorOptionalParameters{}
	return &this
}

// WithGroupStates sets the corresponding parameter name and returns the struct.
func (r *GetMonitorOptionalParameters) WithGroupStates(groupStates string) *GetMonitorOptionalParameters {
	r.GroupStates = &groupStates
	return r
}

// WithWithDowntimes sets the corresponding parameter name and returns the struct.
func (r *GetMonitorOptionalParameters) WithWithDowntimes(withDowntimes bool) *GetMonitorOptionalParameters {
	r.WithDowntimes = &withDowntimes
	return r
}

// GetMonitor Get a monitor's details.
// Get details about the specified monitor from your organization.
func (a *MonitorsApi) GetMonitor(ctx _context.Context, monitorId int64, o ...GetMonitorOptionalParameters) (Monitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue Monitor
		optionalParams      GetMonitorOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetMonitorOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.GetMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/{monitor_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{monitor_id}", _neturl.PathEscape(datadog.ParameterToString(monitorId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.GroupStates != nil {
		localVarQueryParams.Add("group_states", datadog.ParameterToString(*optionalParams.GroupStates, ""))
	}
	if optionalParams.WithDowntimes != nil {
		localVarQueryParams.Add("with_downtimes", datadog.ParameterToString(*optionalParams.WithDowntimes, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// ListMonitorsOptionalParameters holds optional parameters for ListMonitors.
type ListMonitorsOptionalParameters struct {
	GroupStates   *string
	Name          *string
	Tags          *string
	MonitorTags   *string
	WithDowntimes *bool
	IdOffset      *int64
	Page          *int64
	PageSize      *int32
}

// NewListMonitorsOptionalParameters creates an empty struct for parameters.
func NewListMonitorsOptionalParameters() *ListMonitorsOptionalParameters {
	this := ListMonitorsOptionalParameters{}
	return &this
}

// WithGroupStates sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithGroupStates(groupStates string) *ListMonitorsOptionalParameters {
	r.GroupStates = &groupStates
	return r
}

// WithName sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithName(name string) *ListMonitorsOptionalParameters {
	r.Name = &name
	return r
}

// WithTags sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithTags(tags string) *ListMonitorsOptionalParameters {
	r.Tags = &tags
	return r
}

// WithMonitorTags sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithMonitorTags(monitorTags string) *ListMonitorsOptionalParameters {
	r.MonitorTags = &monitorTags
	return r
}

// WithWithDowntimes sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithWithDowntimes(withDowntimes bool) *ListMonitorsOptionalParameters {
	r.WithDowntimes = &withDowntimes
	return r
}

// WithIdOffset sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithIdOffset(idOffset int64) *ListMonitorsOptionalParameters {
	r.IdOffset = &idOffset
	return r
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithPage(page int64) *ListMonitorsOptionalParameters {
	r.Page = &page
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithPageSize(pageSize int32) *ListMonitorsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// ListMonitors Get all monitors.
// Get all monitors from your organization.
func (a *MonitorsApi) ListMonitors(ctx _context.Context, o ...ListMonitorsOptionalParameters) ([]Monitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []Monitor
		optionalParams      ListMonitorsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListMonitorsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.ListMonitors")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.GroupStates != nil {
		localVarQueryParams.Add("group_states", datadog.ParameterToString(*optionalParams.GroupStates, ""))
	}
	if optionalParams.Name != nil {
		localVarQueryParams.Add("name", datadog.ParameterToString(*optionalParams.Name, ""))
	}
	if optionalParams.Tags != nil {
		localVarQueryParams.Add("tags", datadog.ParameterToString(*optionalParams.Tags, ""))
	}
	if optionalParams.MonitorTags != nil {
		localVarQueryParams.Add("monitor_tags", datadog.ParameterToString(*optionalParams.MonitorTags, ""))
	}
	if optionalParams.WithDowntimes != nil {
		localVarQueryParams.Add("with_downtimes", datadog.ParameterToString(*optionalParams.WithDowntimes, ""))
	}
	if optionalParams.IdOffset != nil {
		localVarQueryParams.Add("id_offset", datadog.ParameterToString(*optionalParams.IdOffset, ""))
	}
	if optionalParams.Page != nil {
		localVarQueryParams.Add("page", datadog.ParameterToString(*optionalParams.Page, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page_size", datadog.ParameterToString(*optionalParams.PageSize, ""))
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

// ListMonitorsWithPagination provides a paginated version of ListMonitors returning a channel with all items.
func (a *MonitorsApi) ListMonitorsWithPagination(ctx _context.Context, o ...ListMonitorsOptionalParameters) (<-chan datadog.PaginationResult[Monitor], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(100)
	if len(o) == 0 {
		o = append(o, ListMonitorsOptionalParameters{})
	}
	if o[0].PageSize != nil {
		pageSize_ = *o[0].PageSize
	}
	o[0].PageSize = &pageSize_
	page_ := int64(0)
	o[0].Page = &page_

	items := make(chan datadog.PaginationResult[Monitor], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListMonitors(ctx, o...)
			if err != nil {
				var returnItem Monitor
				items <- datadog.PaginationResult[Monitor]{Item: returnItem, Error: err}
				break
			}
			results := resp

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[Monitor]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) < int(pageSize_) {
				break
			}
			pageOffset_ := *o[0].Page + 1
			o[0].Page = &pageOffset_
		}
		close(items)
	}()
	return items, cancel
}

// SearchMonitorGroupsOptionalParameters holds optional parameters for SearchMonitorGroups.
type SearchMonitorGroupsOptionalParameters struct {
	Query   *string
	Page    *int64
	PerPage *int64
	Sort    *string
}

// NewSearchMonitorGroupsOptionalParameters creates an empty struct for parameters.
func NewSearchMonitorGroupsOptionalParameters() *SearchMonitorGroupsOptionalParameters {
	this := SearchMonitorGroupsOptionalParameters{}
	return &this
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorGroupsOptionalParameters) WithQuery(query string) *SearchMonitorGroupsOptionalParameters {
	r.Query = &query
	return r
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorGroupsOptionalParameters) WithPage(page int64) *SearchMonitorGroupsOptionalParameters {
	r.Page = &page
	return r
}

// WithPerPage sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorGroupsOptionalParameters) WithPerPage(perPage int64) *SearchMonitorGroupsOptionalParameters {
	r.PerPage = &perPage
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorGroupsOptionalParameters) WithSort(sort string) *SearchMonitorGroupsOptionalParameters {
	r.Sort = &sort
	return r
}

// SearchMonitorGroups Monitors group search.
// Search and filter your monitor groups details.
func (a *MonitorsApi) SearchMonitorGroups(ctx _context.Context, o ...SearchMonitorGroupsOptionalParameters) (MonitorGroupSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MonitorGroupSearchResponse
		optionalParams      SearchMonitorGroupsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type SearchMonitorGroupsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.SearchMonitorGroups")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/groups/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", datadog.ParameterToString(*optionalParams.Query, ""))
	}
	if optionalParams.Page != nil {
		localVarQueryParams.Add("page", datadog.ParameterToString(*optionalParams.Page, ""))
	}
	if optionalParams.PerPage != nil {
		localVarQueryParams.Add("per_page", datadog.ParameterToString(*optionalParams.PerPage, ""))
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

// SearchMonitorsOptionalParameters holds optional parameters for SearchMonitors.
type SearchMonitorsOptionalParameters struct {
	Query   *string
	Page    *int64
	PerPage *int64
	Sort    *string
}

// NewSearchMonitorsOptionalParameters creates an empty struct for parameters.
func NewSearchMonitorsOptionalParameters() *SearchMonitorsOptionalParameters {
	this := SearchMonitorsOptionalParameters{}
	return &this
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorsOptionalParameters) WithQuery(query string) *SearchMonitorsOptionalParameters {
	r.Query = &query
	return r
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorsOptionalParameters) WithPage(page int64) *SearchMonitorsOptionalParameters {
	r.Page = &page
	return r
}

// WithPerPage sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorsOptionalParameters) WithPerPage(perPage int64) *SearchMonitorsOptionalParameters {
	r.PerPage = &perPage
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorsOptionalParameters) WithSort(sort string) *SearchMonitorsOptionalParameters {
	r.Sort = &sort
	return r
}

// SearchMonitors Monitors search.
// Search and filter your monitors details.
func (a *MonitorsApi) SearchMonitors(ctx _context.Context, o ...SearchMonitorsOptionalParameters) (MonitorSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MonitorSearchResponse
		optionalParams      SearchMonitorsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type SearchMonitorsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.SearchMonitors")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", datadog.ParameterToString(*optionalParams.Query, ""))
	}
	if optionalParams.Page != nil {
		localVarQueryParams.Add("page", datadog.ParameterToString(*optionalParams.Page, ""))
	}
	if optionalParams.PerPage != nil {
		localVarQueryParams.Add("per_page", datadog.ParameterToString(*optionalParams.PerPage, ""))
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

// UpdateMonitor Edit a monitor.
// Edit the specified monitor.
func (a *MonitorsApi) UpdateMonitor(ctx _context.Context, monitorId int64, body MonitorUpdateRequest) (Monitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue Monitor
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.UpdateMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/{monitor_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{monitor_id}", _neturl.PathEscape(datadog.ParameterToString(monitorId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// ValidateExistingMonitor Validate an existing monitor.
// Validate the monitor provided in the request.
func (a *MonitorsApi) ValidateExistingMonitor(ctx _context.Context, monitorId int64, body Monitor) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.ValidateExistingMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/{monitor_id}/validate"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{monitor_id}", _neturl.PathEscape(datadog.ParameterToString(monitorId, "")))

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

// ValidateMonitor Validate a monitor.
// Validate the monitor provided in the request.
//
// **Note**: Log monitors require an unscoped App Key.
func (a *MonitorsApi) ValidateMonitor(ctx _context.Context, body Monitor) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v1.MonitorsApi.ValidateMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/validate"

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

// NewMonitorsApi Returns NewMonitorsApi.
func NewMonitorsApi(client *datadog.APIClient) *MonitorsApi {
	return &MonitorsApi{
		Client: client,
	}
}
