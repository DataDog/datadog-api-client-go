# LogsAggregateBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **map[string]string** | The key, value pairs for each group by | [optional] 
**Computes** | Pointer to [**map[string]LogsAggregateBucketValue**](LogsAggregateBucketValue.md) | A map of the metric name -&gt; value for regular compute or list of values for a timeseries | [optional] 

## Methods

### NewLogsAggregateBucket

`func NewLogsAggregateBucket() *LogsAggregateBucket`

NewLogsAggregateBucket instantiates a new LogsAggregateBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsAggregateBucketWithDefaults

`func NewLogsAggregateBucketWithDefaults() *LogsAggregateBucket`

NewLogsAggregateBucketWithDefaults instantiates a new LogsAggregateBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *LogsAggregateBucket) GetBy() map[string]string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *LogsAggregateBucket) GetByOk() (*map[string]string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *LogsAggregateBucket) SetBy(v map[string]string)`

SetBy sets By field to given value.

### HasBy

`func (o *LogsAggregateBucket) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetComputes

`func (o *LogsAggregateBucket) GetComputes() map[string]LogsAggregateBucketValue`

GetComputes returns the Computes field if non-nil, zero value otherwise.

### GetComputesOk

`func (o *LogsAggregateBucket) GetComputesOk() (*map[string]LogsAggregateBucketValue, bool)`

GetComputesOk returns a tuple with the Computes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputes

`func (o *LogsAggregateBucket) SetComputes(v map[string]LogsAggregateBucketValue)`

SetComputes sets Computes field to given value.

### HasComputes

`func (o *LogsAggregateBucket) HasComputes() bool`

HasComputes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


