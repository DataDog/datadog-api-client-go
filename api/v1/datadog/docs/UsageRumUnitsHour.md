# UsageRumUnitsHour

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**BrowserRumLiteSessionCount** | Pointer to **int64** | Number of browser RUM lite sessions. | [optional] 
**BrowserRumReplaySessionCount** | Pointer to **int64** | Number of browser RUM replay sessions. | [optional] 
**BrowserRumUnits** | Pointer to **int64** | The number of browser RUM units. | [optional] 
**MobileRumLiteSessionCount** | Pointer to **int64** | Number of mobile RUM lite sessions. | [optional] 
**MobileRumUnits** | Pointer to **int64** | The number of mobile RUM units. | [optional] 
**OrgName** | Pointer to **string** | The organization name. | [optional] 
**PublicId** | Pointer to **string** | The organization public ID. | [optional] 
**RumUnits** | Pointer to **NullableInt64** | Total RUM units across mobile and browser RUM. | [optional] 

## Methods

### NewUsageRumUnitsHour

`func NewUsageRumUnitsHour() *UsageRumUnitsHour`

NewUsageRumUnitsHour instantiates a new UsageRumUnitsHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageRumUnitsHourWithDefaults

`func NewUsageRumUnitsHourWithDefaults() *UsageRumUnitsHour`

NewUsageRumUnitsHourWithDefaults instantiates a new UsageRumUnitsHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBrowserRumLiteSessionCount

`func (o *UsageRumUnitsHour) GetBrowserRumLiteSessionCount() int64`

GetBrowserRumLiteSessionCount returns the BrowserRumLiteSessionCount field if non-nil, zero value otherwise.

### GetBrowserRumLiteSessionCountOk

`func (o *UsageRumUnitsHour) GetBrowserRumLiteSessionCountOk() (*int64, bool)`

GetBrowserRumLiteSessionCountOk returns a tuple with the BrowserRumLiteSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserRumLiteSessionCount

`func (o *UsageRumUnitsHour) SetBrowserRumLiteSessionCount(v int64)`

SetBrowserRumLiteSessionCount sets BrowserRumLiteSessionCount field to given value.

### HasBrowserRumLiteSessionCount

`func (o *UsageRumUnitsHour) HasBrowserRumLiteSessionCount() bool`

HasBrowserRumLiteSessionCount returns a boolean if a field has been set.

### GetBrowserRumReplaySessionCount

`func (o *UsageRumUnitsHour) GetBrowserRumReplaySessionCount() int64`

GetBrowserRumReplaySessionCount returns the BrowserRumReplaySessionCount field if non-nil, zero value otherwise.

### GetBrowserRumReplaySessionCountOk

`func (o *UsageRumUnitsHour) GetBrowserRumReplaySessionCountOk() (*int64, bool)`

GetBrowserRumReplaySessionCountOk returns a tuple with the BrowserRumReplaySessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserRumReplaySessionCount

`func (o *UsageRumUnitsHour) SetBrowserRumReplaySessionCount(v int64)`

SetBrowserRumReplaySessionCount sets BrowserRumReplaySessionCount field to given value.

### HasBrowserRumReplaySessionCount

`func (o *UsageRumUnitsHour) HasBrowserRumReplaySessionCount() bool`

HasBrowserRumReplaySessionCount returns a boolean if a field has been set.

### GetBrowserRumUnits

`func (o *UsageRumUnitsHour) GetBrowserRumUnits() int64`

GetBrowserRumUnits returns the BrowserRumUnits field if non-nil, zero value otherwise.

### GetBrowserRumUnitsOk

`func (o *UsageRumUnitsHour) GetBrowserRumUnitsOk() (*int64, bool)`

GetBrowserRumUnitsOk returns a tuple with the BrowserRumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserRumUnits

`func (o *UsageRumUnitsHour) SetBrowserRumUnits(v int64)`

SetBrowserRumUnits sets BrowserRumUnits field to given value.

### HasBrowserRumUnits

`func (o *UsageRumUnitsHour) HasBrowserRumUnits() bool`

HasBrowserRumUnits returns a boolean if a field has been set.

### GetMobileRumLiteSessionCount

`func (o *UsageRumUnitsHour) GetMobileRumLiteSessionCount() int64`

GetMobileRumLiteSessionCount returns the MobileRumLiteSessionCount field if non-nil, zero value otherwise.

### GetMobileRumLiteSessionCountOk

`func (o *UsageRumUnitsHour) GetMobileRumLiteSessionCountOk() (*int64, bool)`

GetMobileRumLiteSessionCountOk returns a tuple with the MobileRumLiteSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumLiteSessionCount

`func (o *UsageRumUnitsHour) SetMobileRumLiteSessionCount(v int64)`

SetMobileRumLiteSessionCount sets MobileRumLiteSessionCount field to given value.

### HasMobileRumLiteSessionCount

`func (o *UsageRumUnitsHour) HasMobileRumLiteSessionCount() bool`

HasMobileRumLiteSessionCount returns a boolean if a field has been set.

### GetMobileRumUnits

`func (o *UsageRumUnitsHour) GetMobileRumUnits() int64`

GetMobileRumUnits returns the MobileRumUnits field if non-nil, zero value otherwise.

### GetMobileRumUnitsOk

`func (o *UsageRumUnitsHour) GetMobileRumUnitsOk() (*int64, bool)`

GetMobileRumUnitsOk returns a tuple with the MobileRumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumUnits

`func (o *UsageRumUnitsHour) SetMobileRumUnits(v int64)`

SetMobileRumUnits sets MobileRumUnits field to given value.

### HasMobileRumUnits

`func (o *UsageRumUnitsHour) HasMobileRumUnits() bool`

HasMobileRumUnits returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageRumUnitsHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageRumUnitsHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageRumUnitsHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageRumUnitsHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageRumUnitsHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageRumUnitsHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageRumUnitsHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageRumUnitsHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetRumUnits

`func (o *UsageRumUnitsHour) GetRumUnits() int64`

GetRumUnits returns the RumUnits field if non-nil, zero value otherwise.

### GetRumUnitsOk

`func (o *UsageRumUnitsHour) GetRumUnitsOk() (*int64, bool)`

GetRumUnitsOk returns a tuple with the RumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumUnits

`func (o *UsageRumUnitsHour) SetRumUnits(v int64)`

SetRumUnits sets RumUnits field to given value.

### HasRumUnits

`func (o *UsageRumUnitsHour) HasRumUnits() bool`

HasRumUnits returns a boolean if a field has been set.

### SetRumUnitsNil

`func (o *UsageRumUnitsHour) SetRumUnitsNil(b bool)`

 SetRumUnitsNil sets the value for RumUnits to be an explicit nil

### UnsetRumUnits
`func (o *UsageRumUnitsHour) UnsetRumUnits()`

UnsetRumUnits ensures that no value is present for RumUnits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


