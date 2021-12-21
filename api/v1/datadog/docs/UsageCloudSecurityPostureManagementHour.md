# UsageCloudSecurityPostureManagementHour

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**ContainerCount** | Pointer to **int64** | The total number of Cloud Security Posture Management containers during a given hour. | [optional] 
**HostCount** | Pointer to **int64** | The total number of Cloud Security Posture Management hosts during a given hour. | [optional] 
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 
**OrgName** | Pointer to **string** | The organization name. | [optional] 
**PublicId** | Pointer to **string** | The organization public ID. | [optional] 

## Methods

### NewUsageCloudSecurityPostureManagementHour

`func NewUsageCloudSecurityPostureManagementHour() *UsageCloudSecurityPostureManagementHour`

NewUsageCloudSecurityPostureManagementHour instantiates a new UsageCloudSecurityPostureManagementHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageCloudSecurityPostureManagementHourWithDefaults

`func NewUsageCloudSecurityPostureManagementHourWithDefaults() *UsageCloudSecurityPostureManagementHour`

NewUsageCloudSecurityPostureManagementHourWithDefaults instantiates a new UsageCloudSecurityPostureManagementHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetContainerCount

`func (o *UsageCloudSecurityPostureManagementHour) GetContainerCount() int64`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *UsageCloudSecurityPostureManagementHour) GetContainerCountOk() (*int64, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *UsageCloudSecurityPostureManagementHour) SetContainerCount(v int64)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *UsageCloudSecurityPostureManagementHour) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetHostCount

`func (o *UsageCloudSecurityPostureManagementHour) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *UsageCloudSecurityPostureManagementHour) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *UsageCloudSecurityPostureManagementHour) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *UsageCloudSecurityPostureManagementHour) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageCloudSecurityPostureManagementHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageCloudSecurityPostureManagementHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageCloudSecurityPostureManagementHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageCloudSecurityPostureManagementHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageCloudSecurityPostureManagementHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageCloudSecurityPostureManagementHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageCloudSecurityPostureManagementHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageCloudSecurityPostureManagementHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageCloudSecurityPostureManagementHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageCloudSecurityPostureManagementHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageCloudSecurityPostureManagementHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageCloudSecurityPostureManagementHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


