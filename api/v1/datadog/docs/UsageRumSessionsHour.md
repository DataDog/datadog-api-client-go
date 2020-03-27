# UsageRumSessionsHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 
**SessionCount** | Pointer to **int64** | Contains the number of RUM Sessions. | [optional] 

## Methods

### NewUsageRumSessionsHour

`func NewUsageRumSessionsHour() *UsageRumSessionsHour`

NewUsageRumSessionsHour instantiates a new UsageRumSessionsHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageRumSessionsHourWithDefaults

`func NewUsageRumSessionsHourWithDefaults() *UsageRumSessionsHour`

NewUsageRumSessionsHourWithDefaults instantiates a new UsageRumSessionsHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *UsageRumSessionsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageRumSessionsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageRumSessionsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageRumSessionsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetSessionCount

`func (o *UsageRumSessionsHour) GetSessionCount() int64`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *UsageRumSessionsHour) GetSessionCountOk() (*int64, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *UsageRumSessionsHour) SetSessionCount(v int64)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *UsageRumSessionsHour) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


