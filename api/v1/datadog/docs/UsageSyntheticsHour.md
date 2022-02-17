# UsageSyntheticsHour

## Properties

| Name                | Type                     | Description                                      | Notes      |
| ------------------- | ------------------------ | ------------------------------------------------ | ---------- |
| **CheckCallsCount** | Pointer to **int64**     | Contains the number of Synthetics API tests run. | [optional] |
| **Hour**            | Pointer to **time.Time** | The hour for the usage.                          | [optional] |
| **OrgName**         | Pointer to **string**    | The organization name.                           | [optional] |
| **PublicId**        | Pointer to **string**    | The organization public ID.                      | [optional] |

## Methods

### NewUsageSyntheticsHour

`func NewUsageSyntheticsHour() *UsageSyntheticsHour`

NewUsageSyntheticsHour instantiates a new UsageSyntheticsHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSyntheticsHourWithDefaults

`func NewUsageSyntheticsHourWithDefaults() *UsageSyntheticsHour`

NewUsageSyntheticsHourWithDefaults instantiates a new UsageSyntheticsHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCheckCallsCount

`func (o *UsageSyntheticsHour) GetCheckCallsCount() int64`

GetCheckCallsCount returns the CheckCallsCount field if non-nil, zero value otherwise.

### GetCheckCallsCountOk

`func (o *UsageSyntheticsHour) GetCheckCallsCountOk() (*int64, bool)`

GetCheckCallsCountOk returns a tuple with the CheckCallsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCallsCount

`func (o *UsageSyntheticsHour) SetCheckCallsCount(v int64)`

SetCheckCallsCount sets CheckCallsCount field to given value.

### HasCheckCallsCount

`func (o *UsageSyntheticsHour) HasCheckCallsCount() bool`

HasCheckCallsCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageSyntheticsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageSyntheticsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageSyntheticsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageSyntheticsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageSyntheticsHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageSyntheticsHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageSyntheticsHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageSyntheticsHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageSyntheticsHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageSyntheticsHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageSyntheticsHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageSyntheticsHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
