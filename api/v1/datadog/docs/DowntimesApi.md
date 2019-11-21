# \DowntimesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDowntime**](DowntimesApi.md#CancelDowntime) | **Delete** /api/v1/downtime/{downtime_id} | Cancel a downtime
[**CancelDowntimesByScope**](DowntimesApi.md#CancelDowntimesByScope) | **Post** /api/v1/downtime/cancel/by_scope | Cancel downtimes by scope
[**CreateDowntime**](DowntimesApi.md#CreateDowntime) | **Post** /api/v1/downtime | Schedule a downtime
[**GetAllDowntimes**](DowntimesApi.md#GetAllDowntimes) | **Get** /api/v1/downtime | Get all downtimes
[**GetDowntime**](DowntimesApi.md#GetDowntime) | **Get** /api/v1/downtime/{downtime_id} | Get a downtime
[**UpdateDowntime**](DowntimesApi.md#UpdateDowntime) | **Put** /api/v1/downtime/{downtime_id} | Update a downtime



## CancelDowntime

> CancelDowntime(ctx, downtimeId)

Cancel a downtime

Cancel a Downtime

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtimeId** | **int64**| ID of the downtime to cancel | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelDowntimesByScope

> CanceledDowntimesIds CancelDowntimesByScope(ctx, cancelDowntimesByScopeRequest)

Cancel downtimes by scope

### Overview
DELETE all Downtimes that match the scope of X
### Arguments
* **`scope`** [*required*]: Cancel all downtimes with the given scope(s),
  e.g.: `env:prod`, `role:db,role:db-slave`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cancelDowntimesByScopeRequest** | [**CancelDowntimesByScopeRequest**](CancelDowntimesByScopeRequest.md)| Scope to cancel downtimes for | 

### Return type

[**CanceledDowntimesIds**](CanceledDowntimesIds.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDowntime

> Downtime CreateDowntime(ctx, downtime)

Schedule a downtime

* **`scope`** [*required*]: The scope(s) to which the downtime applies, e.g. `host:app2`.
  Provide multiple scopes as a comma-separated list, e.g. `env:dev,env:prod`. The
  resulting downtime applies to sources that matches ALL provided scopes (i.e.
  `env:dev` **AND** `env:prod`), NOT any of them.

* **`monitor_tags`** [*optional*, *default*=**no monitor tag filter**]: A comma-separated
  list of monitor tags, i.e. tags that are applied directly to monitors, *not* tags
  that are used in monitor queries (which are filtered by the `scope` parameter), to
  which the downtime applies. The resulting downtime applies to monitors that match
  ALL provided monitor tags (i.e. `service:postgres` **AND** `team:frontend`), NOT any of them.

* **`monitor_id`** [*optional*, *default*=**None**]: A single monitor to which the downtime
  applies. If not provided, the downtime applies to all monitors.

* **`start`** [*optional*, *default*=**None**]: POSIX timestamp to start the downtime.
  If not provided, the downtime starts the moment it is created.

* **`end`** [*optional*, *default*=**None**]: POSIX timestamp to end the downtime.
  If not provided, the downtime is in effect indefinitely (i.e. until you cancel it).

* **`message`** [*optional*, *default*=**None**]: A message to include with notifications
  for this downtime. Email notifications can be sent to specific users by using
   the same '@username' notation as events

* **`timezone`** [*optional*, *default* = **UTC**]: The timezone for the downtime.
* **`recurrence`** [*optional*, *default*=**None**]: An object defining the recurrence of the
  downtime with a variety of parameters:

  * **`type`** the type of recurrence. Choose from: `days`, `weeks`, `months`, `years`.

  * **`period`** how often to repeat as an integer. For example to repeat every 3 days,
    select a type of `days` and a period of `3`.

  * **`week_days`** (optional) a list of week days to repeat on. Choose from: `Mon`,
    `Tue`, `Wed`, `Thu`, `Fri`, `Sat` or `Sun`. Only applicable when `type` is `weeks`.
    **First letter must be capitalized.**
  * **`until_occurrences`** (optional) how many times the downtime is rescheduled.
    **`until_occurrences` and `until_date`** are mutually exclusive

  * **`until_date`** (optional) the date at which the recurrence should end
    as a POSIX timestmap. **`until_occurrences` and `until_date`** are mutually exclusive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtime** | [**Downtime**](Downtime.md)| Downtime request object | 

### Return type

[**Downtime**](Downtime.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDowntimes

> []Downtime GetAllDowntimes(ctx, optional)

Get all downtimes

### Overview
Get All Scheduled Downtimes
### Arguments
* **`current_only`** [*optional*, *default* = **False**]: Only return downtimes
  that are active when the request is made.'

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllDowntimesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllDowntimesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currentOnly** | **optional.Bool**|  | 

### Return type

[**[]Downtime**](Downtime.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDowntime

> Downtime GetDowntime(ctx, downtimeId)

Get a downtime

### Overview
Get Downtime Detail by downtime_id
### Arguments
This endpoint takes no JSON arguments."

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtimeId** | **int64**| ID of the downtime to fetch | 

### Return type

[**Downtime**](Downtime.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDowntime

> Downtime UpdateDowntime(ctx, downtimeId, downtime)

Update a downtime

### Overview
Update a single Downtime by downtime_id.
### Arguments
* **`id`** [*required*]: The integer id of the downtime to be updated
* **`scope`** [*required*]: The scope to which the downtime applies, e.g. 'host:app2'.
  Provide multiple scopes as a comma-separated list, e.g. 'env:dev,env:prod'.
  The resulting downtime applies to sources that matches ALL provided scopes
  (i.e. env:dev AND env:prod), NOT any of them.

* **`monitor_tags`** [*optional*, *default*=**no monitor tag filter**]: A comma-separated
  list of monitor tags, i.e. tags that are applied directly to monitors, *not* tags that
  are used in monitor queries (which are filtered by the `scope` parameter), to which
  the downtime applies. The resulting downtime applies to monitors that match ALL provided
  monitor tags (i.e. `service:postgres` **AND** `team:frontend`), NOT any of them.

* **`monitor_id`** [*optional*, *default*=**None**]: A single monitor to which the downtime
  applies. If not provided, the downtime applies to all monitors.

* **`start`** [*optional*, *default* = **original start**]: POSIX timestamp to start
  the downtime.

* **`end`** [*optional*, *default* = **original end**]: POSIX timestamp to end the downtime.
  If not provided, the downtime is in effect indefinitely (i.e. until you cancel it).

* **`message`** [*required*, *default* = **original message**]: A message to include with
  notifications for this downtime. Email notifications can be sent to specific users by
  using the same '@username' notation as events

* **`timezone`** [*optional*, default = **original timezone** ]: The timezone for the downtime.
* **`recurrence`** [*optional*, *default* = **original recurrence**]: An object defining the
  recurrence of the downtime with a variety of parameters:

    * **`type`** the type of recurrence. Choose from: `days`, `weeks`, `months`, `years`.

    * **`period`** how often to repeat as an integer. For example to repeat every 3 days,
      select a type of `days` and a period of `3`.

    * **`week_days`** (optional) a list of week days to repeat on. Choose from: `Mon`, `Tue`,
      `Wed`, `Thu`, `Fri`, `Sat` or `Sun`. Only applicable when `type` is `weeks`.
      **First letter must be capitalized.**

    * **`until_occurrences`** (optional) how many times the downtime is rescheduled.
      **`until_occurrences` and `until_date`** are mutually exclusive

    * **`until_date`** (optional) the date at which the recurrence should end as a POSIX
      timestmap. **`until_occurrences` and `until_date`** are mutually exclusive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtimeId** | **int64**| ID of the downtime to update | 
**downtime** | [**Downtime**](Downtime.md)| Downtime request object | 

### Return type

[**Downtime**](Downtime.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

