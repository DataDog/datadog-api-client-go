# LogsMetricFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | The search query - following the log search syntax. | [optional] [default to "*"]

## Methods

### NewLogsMetricFilter

`func NewLogsMetricFilter() *LogsMetricFilter`

NewLogsMetricFilter instantiates a new LogsMetricFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsMetricFilterWithDefaults

`func NewLogsMetricFilterWithDefaults() *LogsMetricFilter`

NewLogsMetricFilterWithDefaults instantiates a new LogsMetricFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *LogsMetricFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogsMetricFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogsMetricFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *LogsMetricFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


