# LogsMetricResponseGroupBy

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Path** | Pointer to **string** | The path to the value the log-based metric will be aggregated over. | [optional] 
**TagName** | Pointer to **string** | Eventual name of the tag that gets created. By default, the path attribute is used as the tag name. | [optional] 

## Methods

### NewLogsMetricResponseGroupBy

`func NewLogsMetricResponseGroupBy() *LogsMetricResponseGroupBy`

NewLogsMetricResponseGroupBy instantiates a new LogsMetricResponseGroupBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsMetricResponseGroupByWithDefaults

`func NewLogsMetricResponseGroupByWithDefaults() *LogsMetricResponseGroupBy`

NewLogsMetricResponseGroupByWithDefaults instantiates a new LogsMetricResponseGroupBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *LogsMetricResponseGroupBy) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogsMetricResponseGroupBy) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogsMetricResponseGroupBy) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogsMetricResponseGroupBy) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTagName

`func (o *LogsMetricResponseGroupBy) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *LogsMetricResponseGroupBy) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *LogsMetricResponseGroupBy) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *LogsMetricResponseGroupBy) HasTagName() bool`

HasTagName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


