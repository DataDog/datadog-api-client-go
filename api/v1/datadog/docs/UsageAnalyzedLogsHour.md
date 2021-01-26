# UsageAnalyzedLogsHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyzedLogs** | Pointer to **int64** | Contains the number of analyzed logs. | [optional] 
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 

## Methods

### NewUsageAnalyzedLogsHour

`func NewUsageAnalyzedLogsHour() *UsageAnalyzedLogsHour`

NewUsageAnalyzedLogsHour instantiates a new UsageAnalyzedLogsHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAnalyzedLogsHourWithDefaults

`func NewUsageAnalyzedLogsHourWithDefaults() *UsageAnalyzedLogsHour`

NewUsageAnalyzedLogsHourWithDefaults instantiates a new UsageAnalyzedLogsHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzedLogs

`func (o *UsageAnalyzedLogsHour) GetAnalyzedLogs() int64`

GetAnalyzedLogs returns the AnalyzedLogs field if non-nil, zero value otherwise.

### GetAnalyzedLogsOk

`func (o *UsageAnalyzedLogsHour) GetAnalyzedLogsOk() (*int64, bool)`

GetAnalyzedLogsOk returns a tuple with the AnalyzedLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedLogs

`func (o *UsageAnalyzedLogsHour) SetAnalyzedLogs(v int64)`

SetAnalyzedLogs sets AnalyzedLogs field to given value.

### HasAnalyzedLogs

`func (o *UsageAnalyzedLogsHour) HasAnalyzedLogs() bool`

HasAnalyzedLogs returns a boolean if a field has been set.

### GetHour

`func (o *UsageAnalyzedLogsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageAnalyzedLogsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageAnalyzedLogsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageAnalyzedLogsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


