# UsageSyntheticsBrowserHour

## Properties

| Name                       | Type                     | Description                                          | Notes      |
| -------------------------- | ------------------------ | ---------------------------------------------------- | ---------- |
| **BrowserCheckCallsCount** | Pointer to **int64**     | Contains the number of Synthetics Browser tests run. | [optional] |
| **Hour**                   | Pointer to **time.Time** | The hour for the usage.                              | [optional] |
| **OrgName**                | Pointer to **string**    | The organization name.                               | [optional] |
| **PublicId**               | Pointer to **string**    | The organization public ID.                          | [optional] |

## Methods

### NewUsageSyntheticsBrowserHour

`func NewUsageSyntheticsBrowserHour() *UsageSyntheticsBrowserHour`

NewUsageSyntheticsBrowserHour instantiates a new UsageSyntheticsBrowserHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSyntheticsBrowserHourWithDefaults

`func NewUsageSyntheticsBrowserHourWithDefaults() *UsageSyntheticsBrowserHour`

NewUsageSyntheticsBrowserHourWithDefaults instantiates a new UsageSyntheticsBrowserHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

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

### GetOrgName

`func (o *UsageSyntheticsBrowserHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageSyntheticsBrowserHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageSyntheticsBrowserHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageSyntheticsBrowserHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageSyntheticsBrowserHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageSyntheticsBrowserHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageSyntheticsBrowserHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageSyntheticsBrowserHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
