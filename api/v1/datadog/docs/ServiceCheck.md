# ServiceCheck

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Check** | **string** | The check. | 
**HostName** | **string** | The host name correlated with the check. | 
**Message** | Pointer to **string** | Message containing check status. | [optional] 
**Status** | [**ServiceCheckStatus**](ServiceCheckStatus.md) |  | 
**Tags** | **[]string** | Tags related to a check. | 
**Timestamp** | Pointer to **int64** | Time of check. | [optional] 

## Methods

### NewServiceCheck

`func NewServiceCheck(check string, hostName string, status ServiceCheckStatus, tags []string) *ServiceCheck`

NewServiceCheck instantiates a new ServiceCheck object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewServiceCheckWithDefaults

`func NewServiceCheckWithDefaults() *ServiceCheck`

NewServiceCheckWithDefaults instantiates a new ServiceCheck object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCheck

`func (o *ServiceCheck) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *ServiceCheck) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *ServiceCheck) SetCheck(v string)`

SetCheck sets Check field to given value.


### GetHostName

`func (o *ServiceCheck) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ServiceCheck) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ServiceCheck) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetMessage

`func (o *ServiceCheck) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceCheck) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceCheck) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServiceCheck) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceCheck) GetStatus() ServiceCheckStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceCheck) GetStatusOk() (*ServiceCheckStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceCheck) SetStatus(v ServiceCheckStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *ServiceCheck) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceCheck) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceCheck) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTimestamp

`func (o *ServiceCheck) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServiceCheck) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServiceCheck) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServiceCheck) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


