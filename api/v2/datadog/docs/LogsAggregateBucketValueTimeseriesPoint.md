# LogsAggregateBucketValueTimeseriesPoint

## Properties

| Name      | Type                   | Description                   | Notes      |
| --------- | ---------------------- | ----------------------------- | ---------- |
| **Time**  | Pointer to **string**  | The time value for this point | [optional] |
| **Value** | Pointer to **float64** | The value for this point      | [optional] |

## Methods

### NewLogsAggregateBucketValueTimeseriesPoint

`func NewLogsAggregateBucketValueTimeseriesPoint() *LogsAggregateBucketValueTimeseriesPoint`

NewLogsAggregateBucketValueTimeseriesPoint instantiates a new LogsAggregateBucketValueTimeseriesPoint object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsAggregateBucketValueTimeseriesPointWithDefaults

`func NewLogsAggregateBucketValueTimeseriesPointWithDefaults() *LogsAggregateBucketValueTimeseriesPoint`

NewLogsAggregateBucketValueTimeseriesPointWithDefaults instantiates a new LogsAggregateBucketValueTimeseriesPoint object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTime

`func (o *LogsAggregateBucketValueTimeseriesPoint) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LogsAggregateBucketValueTimeseriesPoint) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LogsAggregateBucketValueTimeseriesPoint) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *LogsAggregateBucketValueTimeseriesPoint) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValue

`func (o *LogsAggregateBucketValueTimeseriesPoint) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LogsAggregateBucketValueTimeseriesPoint) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LogsAggregateBucketValueTimeseriesPoint) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *LogsAggregateBucketValueTimeseriesPoint) HasValue() bool`

HasValue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
