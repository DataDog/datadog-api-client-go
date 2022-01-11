# MonthlyUsageAttributionBody

## Properties

| Name                | Type                                                                             | Description                                                                                                                                                                                             | Notes      |
| ------------------- | -------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Month**           | Pointer to **time.Time**                                                         | Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM].                                                                                                                                          | [optional] |
| **OrgName**         | Pointer to **string**                                                            | The name of the organization.                                                                                                                                                                           | [optional] |
| **PublicId**        | Pointer to **string**                                                            | The organization public ID.                                                                                                                                                                             | [optional] |
| **TagConfigSource** | Pointer to **string**                                                            | The source of the usage attribution tag configuration and the selected tags in the format &#x60;&lt;source_org_name&gt;:&lt;selected tag 1&gt;///&lt;selected tag 2&gt;///&lt;selected tag 3&gt;&#x60;. | [optional] |
| **Tags**            | Pointer to **map[string][]string**                                               | Usage Summary by tag name.                                                                                                                                                                              | [optional] |
| **UpdatedAt**       | Pointer to **time.Time**                                                         | Datetime of the most recent update to the usage values.                                                                                                                                                 | [optional] |
| **Values**          | Pointer to [**MonthlyUsageAttributionValues**](MonthlyUsageAttributionValues.md) |                                                                                                                                                                                                         | [optional] |

## Methods

### NewMonthlyUsageAttributionBody

`func NewMonthlyUsageAttributionBody() *MonthlyUsageAttributionBody`

NewMonthlyUsageAttributionBody instantiates a new MonthlyUsageAttributionBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonthlyUsageAttributionBodyWithDefaults

`func NewMonthlyUsageAttributionBodyWithDefaults() *MonthlyUsageAttributionBody`

NewMonthlyUsageAttributionBodyWithDefaults instantiates a new MonthlyUsageAttributionBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetMonth

`func (o *MonthlyUsageAttributionBody) GetMonth() time.Time`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *MonthlyUsageAttributionBody) GetMonthOk() (*time.Time, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *MonthlyUsageAttributionBody) SetMonth(v time.Time)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *MonthlyUsageAttributionBody) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetOrgName

`func (o *MonthlyUsageAttributionBody) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *MonthlyUsageAttributionBody) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *MonthlyUsageAttributionBody) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *MonthlyUsageAttributionBody) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *MonthlyUsageAttributionBody) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *MonthlyUsageAttributionBody) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *MonthlyUsageAttributionBody) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *MonthlyUsageAttributionBody) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetTagConfigSource

`func (o *MonthlyUsageAttributionBody) GetTagConfigSource() string`

GetTagConfigSource returns the TagConfigSource field if non-nil, zero value otherwise.

### GetTagConfigSourceOk

`func (o *MonthlyUsageAttributionBody) GetTagConfigSourceOk() (*string, bool)`

GetTagConfigSourceOk returns a tuple with the TagConfigSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagConfigSource

`func (o *MonthlyUsageAttributionBody) SetTagConfigSource(v string)`

SetTagConfigSource sets TagConfigSource field to given value.

### HasTagConfigSource

`func (o *MonthlyUsageAttributionBody) HasTagConfigSource() bool`

HasTagConfigSource returns a boolean if a field has been set.

### GetTags

`func (o *MonthlyUsageAttributionBody) GetTags() map[string][]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MonthlyUsageAttributionBody) GetTagsOk() (*map[string][]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MonthlyUsageAttributionBody) SetTags(v map[string][]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MonthlyUsageAttributionBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MonthlyUsageAttributionBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MonthlyUsageAttributionBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MonthlyUsageAttributionBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MonthlyUsageAttributionBody) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValues

`func (o *MonthlyUsageAttributionBody) GetValues() MonthlyUsageAttributionValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MonthlyUsageAttributionBody) GetValuesOk() (*MonthlyUsageAttributionValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MonthlyUsageAttributionBody) SetValues(v MonthlyUsageAttributionValues)`

SetValues sets Values field to given value.

### HasValues

`func (o *MonthlyUsageAttributionBody) HasValues() bool`

HasValues returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
