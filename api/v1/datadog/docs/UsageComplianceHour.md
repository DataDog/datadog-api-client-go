# UsageComplianceHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComplianceContainerCount** | Pointer to **int64** | The total number of compliance container hours from the start of the given hour&#39;s month until the given hour. | [optional] 
**ComplianceHostCount** | Pointer to **int64** | The total number of compliance hosts hours from the start of the given hour&#39;s month until the given hour. | [optional] 
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 

## Methods

### NewUsageComplianceHour

`func NewUsageComplianceHour() *UsageComplianceHour`

NewUsageComplianceHour instantiates a new UsageComplianceHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageComplianceHourWithDefaults

`func NewUsageComplianceHourWithDefaults() *UsageComplianceHour`

NewUsageComplianceHourWithDefaults instantiates a new UsageComplianceHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplianceContainerCount

`func (o *UsageComplianceHour) GetComplianceContainerCount() int64`

GetComplianceContainerCount returns the ComplianceContainerCount field if non-nil, zero value otherwise.

### GetComplianceContainerCountOk

`func (o *UsageComplianceHour) GetComplianceContainerCountOk() (*int64, bool)`

GetComplianceContainerCountOk returns a tuple with the ComplianceContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceContainerCount

`func (o *UsageComplianceHour) SetComplianceContainerCount(v int64)`

SetComplianceContainerCount sets ComplianceContainerCount field to given value.

### HasComplianceContainerCount

`func (o *UsageComplianceHour) HasComplianceContainerCount() bool`

HasComplianceContainerCount returns a boolean if a field has been set.

### GetComplianceHostCount

`func (o *UsageComplianceHour) GetComplianceHostCount() int64`

GetComplianceHostCount returns the ComplianceHostCount field if non-nil, zero value otherwise.

### GetComplianceHostCountOk

`func (o *UsageComplianceHour) GetComplianceHostCountOk() (*int64, bool)`

GetComplianceHostCountOk returns a tuple with the ComplianceHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceHostCount

`func (o *UsageComplianceHour) SetComplianceHostCount(v int64)`

SetComplianceHostCount sets ComplianceHostCount field to given value.

### HasComplianceHostCount

`func (o *UsageComplianceHour) HasComplianceHostCount() bool`

HasComplianceHostCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageComplianceHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageComplianceHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageComplianceHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageComplianceHour) HasHour() bool`

HasHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


