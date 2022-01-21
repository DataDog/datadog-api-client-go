# UsageCloudSecurityPostureManagementHour

## Properties

| Name                    | Type                           | Description                                                                                   | Notes      |
| ----------------------- | ------------------------------ | --------------------------------------------------------------------------------------------- | ---------- |
| **AasHostCount**        | Pointer to **NullableFloat64** | The number of Cloud Security Posture Management Azure app services hosts during a given hour. | [optional] |
| **AzureHostCount**      | Pointer to **NullableFloat64** | The number of Cloud Security Posture Management Azure hosts during a given hour.              | [optional] |
| **ComplianceHostCount** | Pointer to **NullableFloat64** | The number of Cloud Security Posture Management hosts during a given hour.                    | [optional] |
| **ContainerCount**      | Pointer to **NullableFloat64** | The total number of Cloud Security Posture Management containers during a given hour.         | [optional] |
| **HostCount**           | Pointer to **NullableFloat64** | The total number of Cloud Security Posture Management hosts during a given hour.              | [optional] |
| **Hour**                | Pointer to **time.Time**       | The hour for the usage.                                                                       | [optional] |
| **OrgName**             | Pointer to **string**          | The organization name.                                                                        | [optional] |
| **PublicId**            | Pointer to **string**          | The organization public ID.                                                                   | [optional] |

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

### GetAasHostCount

`func (o *UsageCloudSecurityPostureManagementHour) GetAasHostCount() float64`

GetAasHostCount returns the AasHostCount field if non-nil, zero value otherwise.

### GetAasHostCountOk

`func (o *UsageCloudSecurityPostureManagementHour) GetAasHostCountOk() (*float64, bool)`

GetAasHostCountOk returns a tuple with the AasHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAasHostCount

`func (o *UsageCloudSecurityPostureManagementHour) SetAasHostCount(v float64)`

SetAasHostCount sets AasHostCount field to given value.

### HasAasHostCount

`func (o *UsageCloudSecurityPostureManagementHour) HasAasHostCount() bool`

HasAasHostCount returns a boolean if a field has been set.

### SetAasHostCountNil

`func (o *UsageCloudSecurityPostureManagementHour) SetAasHostCountNil(b bool)`

SetAasHostCountNil sets the value for AasHostCount to be an explicit nil

### UnsetAasHostCount

`func (o *UsageCloudSecurityPostureManagementHour) UnsetAasHostCount()`

UnsetAasHostCount ensures that no value is present for AasHostCount, not even an explicit nil

### GetAzureHostCount

`func (o *UsageCloudSecurityPostureManagementHour) GetAzureHostCount() float64`

GetAzureHostCount returns the AzureHostCount field if non-nil, zero value otherwise.

### GetAzureHostCountOk

`func (o *UsageCloudSecurityPostureManagementHour) GetAzureHostCountOk() (*float64, bool)`

GetAzureHostCountOk returns a tuple with the AzureHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureHostCount

`func (o *UsageCloudSecurityPostureManagementHour) SetAzureHostCount(v float64)`

SetAzureHostCount sets AzureHostCount field to given value.

### HasAzureHostCount

`func (o *UsageCloudSecurityPostureManagementHour) HasAzureHostCount() bool`

HasAzureHostCount returns a boolean if a field has been set.

### SetAzureHostCountNil

`func (o *UsageCloudSecurityPostureManagementHour) SetAzureHostCountNil(b bool)`

SetAzureHostCountNil sets the value for AzureHostCount to be an explicit nil

### UnsetAzureHostCount

`func (o *UsageCloudSecurityPostureManagementHour) UnsetAzureHostCount()`

UnsetAzureHostCount ensures that no value is present for AzureHostCount, not even an explicit nil

### GetComplianceHostCount

`func (o *UsageCloudSecurityPostureManagementHour) GetComplianceHostCount() float64`

GetComplianceHostCount returns the ComplianceHostCount field if non-nil, zero value otherwise.

### GetComplianceHostCountOk

`func (o *UsageCloudSecurityPostureManagementHour) GetComplianceHostCountOk() (*float64, bool)`

GetComplianceHostCountOk returns a tuple with the ComplianceHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceHostCount

`func (o *UsageCloudSecurityPostureManagementHour) SetComplianceHostCount(v float64)`

SetComplianceHostCount sets ComplianceHostCount field to given value.

### HasComplianceHostCount

`func (o *UsageCloudSecurityPostureManagementHour) HasComplianceHostCount() bool`

HasComplianceHostCount returns a boolean if a field has been set.

### SetComplianceHostCountNil

`func (o *UsageCloudSecurityPostureManagementHour) SetComplianceHostCountNil(b bool)`

SetComplianceHostCountNil sets the value for ComplianceHostCount to be an explicit nil

### UnsetComplianceHostCount

`func (o *UsageCloudSecurityPostureManagementHour) UnsetComplianceHostCount()`

UnsetComplianceHostCount ensures that no value is present for ComplianceHostCount, not even an explicit nil

### GetContainerCount

`func (o *UsageCloudSecurityPostureManagementHour) GetContainerCount() float64`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *UsageCloudSecurityPostureManagementHour) GetContainerCountOk() (*float64, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *UsageCloudSecurityPostureManagementHour) SetContainerCount(v float64)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *UsageCloudSecurityPostureManagementHour) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### SetContainerCountNil

`func (o *UsageCloudSecurityPostureManagementHour) SetContainerCountNil(b bool)`

SetContainerCountNil sets the value for ContainerCount to be an explicit nil

### UnsetContainerCount

`func (o *UsageCloudSecurityPostureManagementHour) UnsetContainerCount()`

UnsetContainerCount ensures that no value is present for ContainerCount, not even an explicit nil

### GetHostCount

`func (o *UsageCloudSecurityPostureManagementHour) GetHostCount() float64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *UsageCloudSecurityPostureManagementHour) GetHostCountOk() (*float64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *UsageCloudSecurityPostureManagementHour) SetHostCount(v float64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *UsageCloudSecurityPostureManagementHour) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### SetHostCountNil

`func (o *UsageCloudSecurityPostureManagementHour) SetHostCountNil(b bool)`

SetHostCountNil sets the value for HostCount to be an explicit nil

### UnsetHostCount

`func (o *UsageCloudSecurityPostureManagementHour) UnsetHostCount()`

UnsetHostCount ensures that no value is present for HostCount, not even an explicit nil

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
