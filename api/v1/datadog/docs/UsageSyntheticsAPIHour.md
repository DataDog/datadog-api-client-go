# UsageSyntheticsAPIHour

## Properties

| Name                | Type                     | Description                                      | Notes      |
| ------------------- | ------------------------ | ------------------------------------------------ | ---------- |
| **CheckCallsCount** | Pointer to **int64**     | Contains the number of Synthetics API tests run. | [optional] |
| **Hour**            | Pointer to **time.Time** | The hour for the usage.                          | [optional] |
| **OrgName**         | Pointer to **string**    | The organization name.                           | [optional] |
| **PublicId**        | Pointer to **string**    | The organization public ID.                      | [optional] |

## Methods

### NewUsageSyntheticsAPIHour

`func NewUsageSyntheticsAPIHour() *UsageSyntheticsAPIHour`

NewUsageSyntheticsAPIHour instantiates a new UsageSyntheticsAPIHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSyntheticsAPIHourWithDefaults

`func NewUsageSyntheticsAPIHourWithDefaults() *UsageSyntheticsAPIHour`

NewUsageSyntheticsAPIHourWithDefaults instantiates a new UsageSyntheticsAPIHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCheckCallsCount

`func (o *UsageSyntheticsAPIHour) GetCheckCallsCount() int64`

GetCheckCallsCount returns the CheckCallsCount field if non-nil, zero value otherwise.

### GetCheckCallsCountOk

`func (o *UsageSyntheticsAPIHour) GetCheckCallsCountOk() (*int64, bool)`

GetCheckCallsCountOk returns a tuple with the CheckCallsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCallsCount

`func (o *UsageSyntheticsAPIHour) SetCheckCallsCount(v int64)`

SetCheckCallsCount sets CheckCallsCount field to given value.

### HasCheckCallsCount

`func (o *UsageSyntheticsAPIHour) HasCheckCallsCount() bool`

HasCheckCallsCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageSyntheticsAPIHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageSyntheticsAPIHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageSyntheticsAPIHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageSyntheticsAPIHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageSyntheticsAPIHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageSyntheticsAPIHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageSyntheticsAPIHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageSyntheticsAPIHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageSyntheticsAPIHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageSyntheticsAPIHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageSyntheticsAPIHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageSyntheticsAPIHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
