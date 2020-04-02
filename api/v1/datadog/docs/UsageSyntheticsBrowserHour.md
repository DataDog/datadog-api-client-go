# UsageSyntheticsBrowserHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserCheckCallsCount** | Pointer to **int64** | Contains the number of Synthetics Browser tests run. | [optional] 
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 

## Methods

### NewUsageSyntheticsBrowserHour

`func NewUsageSyntheticsBrowserHour() *UsageSyntheticsBrowserHour`

NewUsageSyntheticsBrowserHour instantiates a new UsageSyntheticsBrowserHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSyntheticsBrowserHourWithDefaults

`func NewUsageSyntheticsBrowserHourWithDefaults() *UsageSyntheticsBrowserHour`

NewUsageSyntheticsBrowserHourWithDefaults instantiates a new UsageSyntheticsBrowserHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserCheckCallsCount

`func (o *UsageSyntheticsBrowserHour) GetBrowserCheckCallsCount() int64`

GetBrowserCheckCallsCount returns the BrowserCheckCallsCount field if non-nil, zero value otherwise.

### GetBrowserCheckCallsCountOk

`func (o *UsageSyntheticsBrowserHour) GetBrowserCheckCallsCountOk() (*int64, bool)`

GetBrowserCheckCallsCountOk returns a tuple with the BrowserCheckCallsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCheckCallsCount

`func (o *UsageSyntheticsBrowserHour) SetBrowserCheckCallsCount(v int64)`

SetBrowserCheckCallsCount sets BrowserCheckCallsCount field to given value.

### HasBrowserCheckCallsCount

`func (o *UsageSyntheticsBrowserHour) HasBrowserCheckCallsCount() bool`

HasBrowserCheckCallsCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageSyntheticsBrowserHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageSyntheticsBrowserHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageSyntheticsBrowserHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageSyntheticsBrowserHour) HasHour() bool`

HasHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


