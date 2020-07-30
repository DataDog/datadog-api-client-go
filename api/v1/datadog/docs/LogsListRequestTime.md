# LogsListRequestTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**time.Time**](time.Time.md) | Minimum timestamp for requested logs. | 
**Timezone** | Pointer to **string** | Timezone can be specified both as an offset (e.g. \&quot;UTC+03:00\&quot;) or a regional zone (e.g. \&quot;Europe/Paris\&quot;). | [optional] 
**To** | [**time.Time**](time.Time.md) | Maximum timestamp for requested logs. | 

## Methods

### NewLogsListRequestTime

`func NewLogsListRequestTime(from time.Time, to time.Time, ) *LogsListRequestTime`

NewLogsListRequestTime instantiates a new LogsListRequestTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsListRequestTimeWithDefaults

`func NewLogsListRequestTimeWithDefaults() *LogsListRequestTime`

NewLogsListRequestTimeWithDefaults instantiates a new LogsListRequestTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *LogsListRequestTime) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LogsListRequestTime) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LogsListRequestTime) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTimezone

`func (o *LogsListRequestTime) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *LogsListRequestTime) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *LogsListRequestTime) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *LogsListRequestTime) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTo

`func (o *LogsListRequestTime) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LogsListRequestTime) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LogsListRequestTime) SetTo(v time.Time)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


