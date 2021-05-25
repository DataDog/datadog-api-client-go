# UsageAttributionBody

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Month** | Pointer to **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM]. | [optional] 
**OrgName** | Pointer to **string** | The name of the organization. | [optional] 
**PublicId** | Pointer to **string** | The organization public ID. | [optional] 
**TagConfigSource** | Pointer to **string** | The source of the usage attribution tag configuration and the selected tags in the format &#x60;&lt;source_org_name&gt;:&lt;selected tag 1&gt;-&lt;selected tag 2&gt;-&lt;selected tag 3&gt;&#x60;. | [optional] 
**Tags** | Pointer to **map[string][]string** | Usage Summary by tag name. | [optional] 
**UpdatedAt** | Pointer to **string** | Shows the the most recent hour in the current months for all organizations for which all usages were calculated. | [optional] 
**Values** | Pointer to [**UsageAttributionValues**](UsageAttributionValues.md) |  | [optional] 

## Methods

### NewUsageAttributionBody

`func NewUsageAttributionBody() *UsageAttributionBody`

NewUsageAttributionBody instantiates a new UsageAttributionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAttributionBodyWithDefaults

`func NewUsageAttributionBodyWithDefaults() *UsageAttributionBody`

NewUsageAttributionBodyWithDefaults instantiates a new UsageAttributionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *UsageAttributionBody) GetMonth() time.Time`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *UsageAttributionBody) GetMonthOk() (*time.Time, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *UsageAttributionBody) SetMonth(v time.Time)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *UsageAttributionBody) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageAttributionBody) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageAttributionBody) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageAttributionBody) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageAttributionBody) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageAttributionBody) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageAttributionBody) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageAttributionBody) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageAttributionBody) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetTagConfigSource

`func (o *UsageAttributionBody) GetTagConfigSource() string`

GetTagConfigSource returns the TagConfigSource field if non-nil, zero value otherwise.

### GetTagConfigSourceOk

`func (o *UsageAttributionBody) GetTagConfigSourceOk() (*string, bool)`

GetTagConfigSourceOk returns a tuple with the TagConfigSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagConfigSource

`func (o *UsageAttributionBody) SetTagConfigSource(v string)`

SetTagConfigSource sets TagConfigSource field to given value.

### HasTagConfigSource

`func (o *UsageAttributionBody) HasTagConfigSource() bool`

HasTagConfigSource returns a boolean if a field has been set.

### GetTags

`func (o *UsageAttributionBody) GetTags() map[string][]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UsageAttributionBody) GetTagsOk() (*map[string][]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UsageAttributionBody) SetTags(v map[string][]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UsageAttributionBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UsageAttributionBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UsageAttributionBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UsageAttributionBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UsageAttributionBody) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValues

`func (o *UsageAttributionBody) GetValues() UsageAttributionValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *UsageAttributionBody) GetValuesOk() (*UsageAttributionValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *UsageAttributionBody) SetValues(v UsageAttributionValues)`

SetValues sets Values field to given value.

### HasValues

`func (o *UsageAttributionBody) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


