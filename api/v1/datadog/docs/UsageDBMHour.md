# UsageDBMHour

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**DbmHostCount** | Pointer to **int64** | The total number of Database Monitoring host hours from the start of the given hour’s month until the given hour. | [optional] 
**DbmQueriesCount** | Pointer to **int64** | The total number of normalized Database Monitoring queries from the start of the given hour’s month until the given hour. | [optional] 
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 
**OrgName** | Pointer to **string** | The organization name. | [optional] 
**PublicId** | Pointer to **string** | The organization public ID. | [optional] 

## Methods

### NewUsageDBMHour

`func NewUsageDBMHour() *UsageDBMHour`

NewUsageDBMHour instantiates a new UsageDBMHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageDBMHourWithDefaults

`func NewUsageDBMHourWithDefaults() *UsageDBMHour`

NewUsageDBMHourWithDefaults instantiates a new UsageDBMHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDbmHostCount

`func (o *UsageDBMHour) GetDbmHostCount() int64`

GetDbmHostCount returns the DbmHostCount field if non-nil, zero value otherwise.

### GetDbmHostCountOk

`func (o *UsageDBMHour) GetDbmHostCountOk() (*int64, bool)`

GetDbmHostCountOk returns a tuple with the DbmHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmHostCount

`func (o *UsageDBMHour) SetDbmHostCount(v int64)`

SetDbmHostCount sets DbmHostCount field to given value.

### HasDbmHostCount

`func (o *UsageDBMHour) HasDbmHostCount() bool`

HasDbmHostCount returns a boolean if a field has been set.

### GetDbmQueriesCount

`func (o *UsageDBMHour) GetDbmQueriesCount() int64`

GetDbmQueriesCount returns the DbmQueriesCount field if non-nil, zero value otherwise.

### GetDbmQueriesCountOk

`func (o *UsageDBMHour) GetDbmQueriesCountOk() (*int64, bool)`

GetDbmQueriesCountOk returns a tuple with the DbmQueriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmQueriesCount

`func (o *UsageDBMHour) SetDbmQueriesCount(v int64)`

SetDbmQueriesCount sets DbmQueriesCount field to given value.

### HasDbmQueriesCount

`func (o *UsageDBMHour) HasDbmQueriesCount() bool`

HasDbmQueriesCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageDBMHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageDBMHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageDBMHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageDBMHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageDBMHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageDBMHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageDBMHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageDBMHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageDBMHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageDBMHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageDBMHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageDBMHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


