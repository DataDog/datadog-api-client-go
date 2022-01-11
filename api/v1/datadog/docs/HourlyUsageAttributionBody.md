# HourlyUsageAttributionBody

## Properties

| Name                | Type                                                                                 | Description                                                                                                                                                                                                  | Notes      |
| ------------------- | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ---------- |
| **Hour**            | Pointer to **time.Time**                                                             | The hour for the usage.                                                                                                                                                                                      | [optional] |
| **OrgName**         | Pointer to **string**                                                                | The name of the organization.                                                                                                                                                                                | [optional] |
| **PublicId**        | Pointer to **string**                                                                | The organization public ID.                                                                                                                                                                                  | [optional] |
| **TagConfigSource** | Pointer to **string**                                                                | The source of the usage attribution tag configuration and the selected tags in the format of &#x60;&lt;source_org_name&gt;:::&lt;selected tag 1&gt;///&lt;selected tag 2&gt;///&lt;selected tag 3&gt;&#x60;. | [optional] |
| **Tags**            | Pointer to **map[string][]string**                                                   | Usage Summary by tag name.                                                                                                                                                                                   | [optional] |
| **TotalUsageSum**   | Pointer to **float64**                                                               | Total product usage for the given tags within the hour.                                                                                                                                                      | [optional] |
| **UpdatedAt**       | Pointer to **string**                                                                | Shows the most recent hour in the current month for all organizations where usages are calculated.                                                                                                           | [optional] |
| **UsageType**       | Pointer to [**HourlyUsageAttributionUsageType**](HourlyUsageAttributionUsageType.md) |                                                                                                                                                                                                              | [optional] |

## Methods

### NewHourlyUsageAttributionBody

`func NewHourlyUsageAttributionBody() *HourlyUsageAttributionBody`

NewHourlyUsageAttributionBody instantiates a new HourlyUsageAttributionBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHourlyUsageAttributionBodyWithDefaults

`func NewHourlyUsageAttributionBodyWithDefaults() *HourlyUsageAttributionBody`

NewHourlyUsageAttributionBodyWithDefaults instantiates a new HourlyUsageAttributionBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *HourlyUsageAttributionBody) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *HourlyUsageAttributionBody) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *HourlyUsageAttributionBody) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *HourlyUsageAttributionBody) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetOrgName

`func (o *HourlyUsageAttributionBody) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *HourlyUsageAttributionBody) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *HourlyUsageAttributionBody) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *HourlyUsageAttributionBody) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *HourlyUsageAttributionBody) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *HourlyUsageAttributionBody) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *HourlyUsageAttributionBody) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *HourlyUsageAttributionBody) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetTagConfigSource

`func (o *HourlyUsageAttributionBody) GetTagConfigSource() string`

GetTagConfigSource returns the TagConfigSource field if non-nil, zero value otherwise.

### GetTagConfigSourceOk

`func (o *HourlyUsageAttributionBody) GetTagConfigSourceOk() (*string, bool)`

GetTagConfigSourceOk returns a tuple with the TagConfigSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagConfigSource

`func (o *HourlyUsageAttributionBody) SetTagConfigSource(v string)`

SetTagConfigSource sets TagConfigSource field to given value.

### HasTagConfigSource

`func (o *HourlyUsageAttributionBody) HasTagConfigSource() bool`

HasTagConfigSource returns a boolean if a field has been set.

### GetTags

`func (o *HourlyUsageAttributionBody) GetTags() map[string][]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HourlyUsageAttributionBody) GetTagsOk() (*map[string][]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HourlyUsageAttributionBody) SetTags(v map[string][]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HourlyUsageAttributionBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTotalUsageSum

`func (o *HourlyUsageAttributionBody) GetTotalUsageSum() float64`

GetTotalUsageSum returns the TotalUsageSum field if non-nil, zero value otherwise.

### GetTotalUsageSumOk

`func (o *HourlyUsageAttributionBody) GetTotalUsageSumOk() (*float64, bool)`

GetTotalUsageSumOk returns a tuple with the TotalUsageSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageSum

`func (o *HourlyUsageAttributionBody) SetTotalUsageSum(v float64)`

SetTotalUsageSum sets TotalUsageSum field to given value.

### HasTotalUsageSum

`func (o *HourlyUsageAttributionBody) HasTotalUsageSum() bool`

HasTotalUsageSum returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HourlyUsageAttributionBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HourlyUsageAttributionBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HourlyUsageAttributionBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HourlyUsageAttributionBody) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsageType

`func (o *HourlyUsageAttributionBody) GetUsageType() HourlyUsageAttributionUsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *HourlyUsageAttributionBody) GetUsageTypeOk() (*HourlyUsageAttributionUsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *HourlyUsageAttributionBody) SetUsageType(v HourlyUsageAttributionUsageType)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *HourlyUsageAttributionBody) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
