# MonitorsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CheckCanDeleteMonitor**](MonitorsApi.md#CheckCanDeleteMonitor) | **Get** /api/v1/monitor/can_delete | Check if a monitor can be deleted
[**CreateMonitor**](MonitorsApi.md#CreateMonitor) | **Post** /api/v1/monitor | Create a monitor
[**DeleteMonitor**](MonitorsApi.md#DeleteMonitor) | **Delete** /api/v1/monitor/{monitor_id} | Delete a monitor
[**GetMonitor**](MonitorsApi.md#GetMonitor) | **Get** /api/v1/monitor/{monitor_id} | Get a monitor&#39;s details
[**ListMonitors**](MonitorsApi.md#ListMonitors) | **Get** /api/v1/monitor | Get all monitor details
[**SearchMonitorGroups**](MonitorsApi.md#SearchMonitorGroups) | **Get** /api/v1/monitor/groups/search | Monitors group search
[**SearchMonitors**](MonitorsApi.md#SearchMonitors) | **Get** /api/v1/monitor/search | Monitors search
[**UpdateMonitor**](MonitorsApi.md#UpdateMonitor) | **Put** /api/v1/monitor/{monitor_id} | Edit a monitor
[**ValidateMonitor**](MonitorsApi.md#ValidateMonitor) | **Post** /api/v1/monitor/validate | Validate a monitor



## CheckCanDeleteMonitor

> CheckCanDeleteMonitorResponse CheckCanDeleteMonitor(ctx, monitorIds)

Check if the given monitors can be deleted.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    monitorIds := []int64{int64(123)} // []int64 | The IDs of the monitor to check.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsApi.CheckCanDeleteMonitor(ctx, monitorIds)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CheckCanDeleteMonitor`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCanDeleteMonitor`: CheckCanDeleteMonitorResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MonitorsApi.CheckCanDeleteMonitor:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**monitorIds** | **[]int64** | The IDs of the monitor to check. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**CheckCanDeleteMonitorResponse**](CheckCanDeleteMonitorResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitor

> Monitor CreateMonitor(ctx, body)

Create a monitor using the specified options.

#### Monitor Types

The type of monitor chosen from:

- anomaly: `query alert`
- APM: `query alert` or `trace-analytics alert`
- composite: `composite`
- custom: `service check`
- event: `event alert`
- forecast: `query alert`
- host: `service check`
- integration: `query alert` or `service check`
- live process: `process alert`
- logs: `log alert`
- metric: `metric alert`
- network: `service check`
- outlier: `query alert`
- process: `service check`
- rum: `rum alert`
- SLO: `slo alert`
- watchdog: `event alert`
- event-v2: `event-v2 alert`
- audit: `audit alert`

#### Query Types

**Metric Alert Query**

Example: `time_aggr(time_window):space_aggr:metric{tags} [by {key}] operator #`

- `time_aggr`: avg, sum, max, min, change, or pct_change
- `time_window`: `last_#m` (with `#` between 1 and 10080 depending on the monitor type) or `last_#h`(with `#` between 1 and 168 depending on the monitor type) or `last_1d`, or `last_1w`
- `space_aggr`: avg, sum, min, or max
- `tags`: one or more tags (comma-separated), or *
- `key`: a 'key' in key:value tag syntax; defines a separate alert for each tag in the group (multi-alert)
- `operator`: <, <=, >, >=, ==, or !=
- `#`: an integer or decimal number used to set the threshold

If you are using the `_change_` or `_pct_change_` time aggregator, instead use `change_aggr(time_aggr(time_window),
timeshift):space_aggr:metric{tags} [by {key}] operator #` with:

- `change_aggr` change, pct_change
- `time_aggr` avg, sum, max, min [Learn more](https://docs.datadoghq.com/monitors/monitor_types/#define-the-conditions)
- `time_window` last\_#m (between 1 and 2880 depending on the monitor type), last\_#h (between 1 and 48 depending on the monitor type), or last_#d (1 or 2)
- `timeshift` #m_ago (5, 10, 15, or 30), #h_ago (1, 2, or 4), or 1d_ago

Use this to create an outlier monitor using the following query:
`avg(last_30m):outliers(avg:system.cpu.user{role:es-events-data} by {host}, 'dbscan', 7) > 0`

**Service Check Query**

Example: `"check".over(tags).last(count).by(group).count_by_status()`

- **`check`** name of the check, e.g. `datadog.agent.up`
- **`tags`** one or more quoted tags (comma-separated), or "*". e.g.: `.over("env:prod", "role:db")`; **`over`** cannot be blank.
- **`count`** must be at greater than or equal to your max threshold (defined in the `options`). It is limited to 100.
For example, if you've specified to notify on 1 critical, 3 ok, and 2 warn statuses, `count` should be at least 3.
- **`group`** must be specified for check monitors. Per-check grouping is already explicitly known for some service checks.
For example, Postgres integration monitors are tagged by `db`, `host`, and `port`, and Network monitors by `host`, `instance`, and `url`. See [Service Checks](https://docs.datadoghq.com/api/latest/service-checks/) documentation for more information.

**Event Alert Query**

Example: `events('sources:nagios status:error,warning priority:normal tags: "string query"').rollup("count").last("1h")"`

- **`event`**, the event query string:
- **`string_query`** free text query to match against event title and text.
- **`sources`** event sources (comma-separated).
- **`status`** event statuses (comma-separated). Valid options: error, warn, and info.
- **`priority`** event priorities (comma-separated). Valid options: low, normal, all.
- **`host`** event reporting host (comma-separated).
- **`tags`** event tags (comma-separated).
- **`excluded_tags`** excluded event tags (comma-separated).
- **`rollup`** the stats roll-up method. `count` is the only supported method now.
- **`last`** the timeframe to roll up the counts. Examples: 45m, 4h. Supported timeframes: m, h and d. This value should not exceed 48 hours.

**NOTE** Only available on US1 and EU.

**Event V2 Alert Query**

Example: `events(query).rollup(rollup_method[, measure]).last(time_window) operator #`

- **`query`** The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
- **`rollup_method`** The stats roll-up method - supports `count`, `avg` and `cardinality`.
- **`measure`** For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
- **`time_window`** #m (between 1 and 2880), #h (between 1 and 48).
- **`operator`** `<`, `<=`, `>`, `>=`, `==`, or `!=`.
- **`#`** an integer or decimal number used to set the threshold.

**NOTE** Only available on US1-FED, US3, and in closed beta on EU and US1.

**Process Alert Query**

Example: `processes(search).over(tags).rollup('count').last(timeframe) operator #`

- **`search`** free text search string for querying processes.
Matching processes match results on the [Live Processes](https://docs.datadoghq.com/infrastructure/process/?tab=linuxwindows) page.
- **`tags`** one or more tags (comma-separated)
- **`timeframe`** the timeframe to roll up the counts. Examples: 10m, 4h. Supported timeframes: s, m, h and d
- **`operator`** <, <=, >, >=, ==, or !=
- **`#`** an integer or decimal number used to set the threshold

**Logs Alert Query**

Example: `logs(query).index(index_name).rollup(rollup_method[, measure]).last(time_window) operator #`

- **`query`** The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
- **`index_name`** For multi-index organizations, the log index in which the request is performed.
- **`rollup_method`** The stats roll-up method - supports `count`, `avg` and `cardinality`.
- **`measure`** For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
- **`time_window`** #m (between 1 and 2880), #h (between 1 and 48).
- **`operator`** `<`, `<=`, `>`, `>=`, `==`, or `!=`.
- **`#`** an integer or decimal number used to set the threshold.

**Composite Query**

Example: `12345 && 67890`, where `12345` and `67890` are the IDs of non-composite monitors

* **`name`** [*required*, *default* = **dynamic, based on query**]: The name of the alert.
* **`message`** [*required*, *default* = **dynamic, based on query**]: A message to include with notifications for this monitor.
Email notifications can be sent to specific users by using the same '@username' notation as events.
* **`tags`** [*optional*, *default* = **empty list**]: A list of tags to associate with your monitor.
When getting all monitor details via the API, use the `monitor_tags` argument to filter results by these tags.
It is only available via the API and isn't visible or editable in the Datadog UI.

**SLO Alert Query**

Example: `error_budget("slo_id").over("time_window") operator #`

- **`slo_id`**: The alphanumeric SLO ID of the SLO you are configuring the alert for.
- **`time_window`**: The time window of the SLO target you wish to alert on. Valid options: `7d`, `30d`, `90d`.
- **`operator`**: `>=` or `>`

**Audit Alert Query**

Example: `audits(query).rollup(rollup_method[, measure]).last(time_window) operator #`

- **`query`** The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
- **`rollup_method`** The stats roll-up method - supports `count`, `avg` and `cardinality`.
- **`measure`** For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
- **`time_window`** #m (between 1 and 2880), #h (between 1 and 48).
- **`operator`** `<`, `<=`, `>`, `>=`, `==`, or `!=`.
- **`#`** an integer or decimal number used to set the threshold.

**NOTE** Only available on US1-FED and in closed beta on EU, US3, and US1.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewMonitor("avg(last_5m):sum:system.net.bytes_rcvd{host:host0} > 100", datadog.MonitorType("composite")) // Monitor | Create a monitor request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsApi.CreateMonitor(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitor`: Monitor
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MonitorsApi.CreateMonitor:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**Monitor**](Monitor.md) | Create a monitor request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**Monitor**](Monitor.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitor

> DeletedMonitor DeleteMonitor(ctx, monitorId, datadog.DeleteMonitorOptionalParameters{})

Delete the specified monitor

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    monitorId := int64(789) // int64 | The ID of the monitor.
    force := "false" // string | Delete the monitor even if it's referenced by other resources (e.g. SLO, composite monitor). (optional)
    optionalParams := datadog.DeleteMonitorOptionalParameters{
        Force: &force,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsApi.DeleteMonitor(ctx, monitorId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.DeleteMonitor`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMonitor`: DeletedMonitor
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MonitorsApi.DeleteMonitor:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**monitorId** | **int64** | The ID of the monitor. | 


### Optional Parameters


Other parameters are passed through a pointer to a DeleteMonitorOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**force** | **string** | Delete the monitor even if it&#39;s referenced by other resources (e.g. SLO, composite monitor). | 

### Return type

[**DeletedMonitor**](DeletedMonitor.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitor

> Monitor GetMonitor(ctx, monitorId, datadog.GetMonitorOptionalParameters{})

Get details about the specified monitor from your organization.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    monitorId := int64(789) // int64 | The ID of the monitor
    groupStates := "groupStates_example" // string | When specified, shows additional information about the group states. Choose one or more from `all`, `alert`, `warn`, and `no data`. (optional)
    optionalParams := datadog.GetMonitorOptionalParameters{
        GroupStates: &groupStates,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsApi.GetMonitor(ctx, monitorId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitor`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitor`: Monitor
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MonitorsApi.GetMonitor:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**monitorId** | **int64** | The ID of the monitor | 


### Optional Parameters


Other parameters are passed through a pointer to a GetMonitorOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**groupStates** | **string** | When specified, shows additional information about the group states. Choose one or more from &#x60;all&#x60;, &#x60;alert&#x60;, &#x60;warn&#x60;, and &#x60;no data&#x60;. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMonitors

> []Monitor ListMonitors(ctx, datadog.ListMonitorsOptionalParameters{})

Get details about the specified monitor from your organization.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    groupStates := "alert" // string | When specified, shows additional information about the group states. Choose one or more from `all`, `alert`, `warn`, and `no data`. (optional)
    name := "name_example" // string | A string to filter monitors by name. (optional)
    tags := "host:host0" // string | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. For example, `host:host0`. (optional)
    monitorTags := "service:my-app" // string | A comma separated list indicating what service and/or custom tags, if any, should be used to filter the list of monitors. Tags created in the Datadog UI automatically have the service key prepended. For example, `service:my-app`. (optional)
    withDowntimes := true // bool | If this argument is set to true, then the returned data includes all current downtimes for each monitor. (optional)
    idOffset := int64(789) // int64 | Monitor ID offset. (optional)
    page := int64(0) // int64 | The page to start paginating from. If this argument is not specified, the request returns all monitors without pagination. (optional)
    pageSize := int32(20) // int32 | The number of monitors to return per page. If the page argument is not specified, the default behavior returns all monitors without a `page_size` limit. However, if page is specified and `page_size` is not, the argument defaults to 100. (optional)
    optionalParams := datadog.ListMonitorsOptionalParameters{
        GroupStates: &groupStates,
        Name: &name,
        Tags: &tags,
        MonitorTags: &monitorTags,
        WithDowntimes: &withDowntimes,
        IdOffset: &idOffset,
        Page: &page,
        PageSize: &pageSize,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsApi.ListMonitors(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ListMonitors`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMonitors`: []Monitor
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MonitorsApi.ListMonitors:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListMonitorsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**groupStates** | **string** | When specified, shows additional information about the group states. Choose one or more from &#x60;all&#x60;, &#x60;alert&#x60;, &#x60;warn&#x60;, and &#x60;no data&#x60;. | 
**name** | **string** | A string to filter monitors by name. | 
**tags** | **string** | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. For example, &#x60;host:host0&#x60;. | 
**monitorTags** | **string** | A comma separated list indicating what service and/or custom tags, if any, should be used to filter the list of monitors. Tags created in the Datadog UI automatically have the service key prepended. For example, &#x60;service:my-app&#x60;. | 
**withDowntimes** | **bool** | If this argument is set to true, then the returned data includes all current downtimes for each monitor. | 
**idOffset** | **int64** | Monitor ID offset. | 
**page** | **int64** | The page to start paginating from. If this argument is not specified, the request returns all monitors without pagination. | 
**pageSize** | **int32** | The number of monitors to return per page. If the page argument is not specified, the default behavior returns all monitors without a &#x60;page_size&#x60; limit. However, if page is specified and &#x60;page_size&#x60; is not, the argument defaults to 100. | 

### Return type

[**[]Monitor**](Monitor.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMonitorGroups

> MonitorGroupSearchResponse SearchMonitorGroups(ctx, datadog.SearchMonitorGroupsOptionalParameters{})

Search and filter your monitor groups details.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    query := "query_example" // string | After entering a search query in your [Manage Monitor page][1] use the query parameter value in the URL of the page as value for this parameter. Consult the dedicated [manage monitor documentation][2] page to learn more.  The query can contain any number of space-separated monitor attributes, for instance `query=\"type:metric status:alert\"`.  [1]: https://app.datadoghq.com/monitors/manage [2]: /monitors/manage_monitor/#find-the-monitors (optional)
    page := int64(789) // int64 | Page to start paginating from. (optional) (default to 0)
    perPage := int64(789) // int64 | Number of monitors to return per page. (optional) (default to 30)
    sort := "sort_example" // string | String for sort order, composed of field and sort order separate by a comma, e.g. `name,asc`. Supported sort directions: `asc`, `desc`. Supported fields:  * `name` * `status` * `tags` (optional)
    optionalParams := datadog.SearchMonitorGroupsOptionalParameters{
        Query: &query,
        Page: &page,
        PerPage: &perPage,
        Sort: &sort,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsApi.SearchMonitorGroups(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.SearchMonitorGroups`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMonitorGroups`: MonitorGroupSearchResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MonitorsApi.SearchMonitorGroups:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a SearchMonitorGroupsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**query** | **string** | After entering a search query in your [Manage Monitor page][1] use the query parameter value in the URL of the page as value for this parameter. Consult the dedicated [manage monitor documentation][2] page to learn more.  The query can contain any number of space-separated monitor attributes, for instance &#x60;query&#x3D;\&quot;type:metric status:alert\&quot;&#x60;.  [1]: https://app.datadoghq.com/monitors/manage [2]: /monitors/manage_monitor/#find-the-monitors | 
**page** | **int64** | Page to start paginating from. | [default to 0]
**perPage** | **int64** | Number of monitors to return per page. | [default to 30]
**sort** | **string** | String for sort order, composed of field and sort order separate by a comma, e.g. &#x60;name,asc&#x60;. Supported sort directions: &#x60;asc&#x60;, &#x60;desc&#x60;. Supported fields:  * &#x60;name&#x60; * &#x60;status&#x60; * &#x60;tags&#x60; | 

### Return type

[**MonitorGroupSearchResponse**](MonitorGroupSearchResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMonitors

> MonitorSearchResponse SearchMonitors(ctx, datadog.SearchMonitorsOptionalParameters{})

Search and filter your monitors details.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    query := "query_example" // string | After entering a search query in your [Manage Monitor page][1] use the query parameter value in the URL of the page as value for this parameter. Consult the dedicated [manage monitor documentation][2] page to learn more.  The query can contain any number of space-separated monitor attributes, for instance `query=\"type:metric status:alert\"`.  [1]: https://app.datadoghq.com/monitors/manage [2]: /monitors/manage_monitor/#find-the-monitors (optional)
    page := int64(789) // int64 | Page to start paginating from. (optional) (default to 0)
    perPage := int64(789) // int64 | Number of monitors to return per page. (optional) (default to 30)
    sort := "sort_example" // string | String for sort order, composed of field and sort order separate by a comma, e.g. `name,asc`. Supported sort directions: `asc`, `desc`. Supported fields:  * `name` * `status` * `tags` (optional)
    optionalParams := datadog.SearchMonitorsOptionalParameters{
        Query: &query,
        Page: &page,
        PerPage: &perPage,
        Sort: &sort,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsApi.SearchMonitors(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.SearchMonitors`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMonitors`: MonitorSearchResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MonitorsApi.SearchMonitors:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a SearchMonitorsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**query** | **string** | After entering a search query in your [Manage Monitor page][1] use the query parameter value in the URL of the page as value for this parameter. Consult the dedicated [manage monitor documentation][2] page to learn more.  The query can contain any number of space-separated monitor attributes, for instance &#x60;query&#x3D;\&quot;type:metric status:alert\&quot;&#x60;.  [1]: https://app.datadoghq.com/monitors/manage [2]: /monitors/manage_monitor/#find-the-monitors | 
**page** | **int64** | Page to start paginating from. | [default to 0]
**perPage** | **int64** | Number of monitors to return per page. | [default to 30]
**sort** | **string** | String for sort order, composed of field and sort order separate by a comma, e.g. &#x60;name,asc&#x60;. Supported sort directions: &#x60;asc&#x60;, &#x60;desc&#x60;. Supported fields:  * &#x60;name&#x60; * &#x60;status&#x60; * &#x60;tags&#x60; | 

### Return type

[**MonitorSearchResponse**](MonitorSearchResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitor

> Monitor UpdateMonitor(ctx, monitorId, body)

Edit the specified monitor.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    monitorId := int64(789) // int64 | The ID of the monitor.
    body := *datadog.NewMonitorUpdateRequest() // MonitorUpdateRequest | Edit a monitor request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsApi.UpdateMonitor(ctx, monitorId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.UpdateMonitor`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitor`: Monitor
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MonitorsApi.UpdateMonitor:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**monitorId** | **int64** | The ID of the monitor. |  |
**body** | [**MonitorUpdateRequest**](MonitorUpdateRequest.md) | Edit a monitor request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**Monitor**](Monitor.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMonitor

> interface{} ValidateMonitor(ctx, body)

Validate the monitor provided in the request.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewMonitor("avg(last_5m):sum:system.net.bytes_rcvd{host:host0} > 100", datadog.MonitorType("composite")) // Monitor | Monitor request object

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsApi.ValidateMonitor(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ValidateMonitor`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateMonitor`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MonitorsApi.ValidateMonitor:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**Monitor**](Monitor.md) | Monitor request object | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

**interface{}**

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

