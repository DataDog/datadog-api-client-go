# LogsExclusionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | Default query is &#39;*&#39;, meaning all logs flowing in the index would be excluded. Scope down exclusion filter to only a subset of logs with a log query. | [optional] 
**SampleRate** | Pointer to **float64** |  | 

## Methods

### NewLogsExclusionFilter

`func NewLogsExclusionFilter(sampleRate float64, ) *LogsExclusionFilter`

NewLogsExclusionFilter instantiates a new LogsExclusionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsExclusionFilterWithDefaults

`func NewLogsExclusionFilterWithDefaults() *LogsExclusionFilter`

NewLogsExclusionFilterWithDefaults instantiates a new LogsExclusionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *LogsExclusionFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogsExclusionFilter) GetQueryOk() (string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *LogsExclusionFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *LogsExclusionFilter) SetQuery(v string)`

SetQuery gets a reference to the given string and assigns it to the Query field.

### GetSampleRate

`func (o *LogsExclusionFilter) GetSampleRate() float64`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *LogsExclusionFilter) GetSampleRateOk() (float64, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleRate

`func (o *LogsExclusionFilter) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.

### SetSampleRate

`func (o *LogsExclusionFilter) SetSampleRate(v float64)`

SetSampleRate gets a reference to the given float64 and assigns it to the SampleRate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


