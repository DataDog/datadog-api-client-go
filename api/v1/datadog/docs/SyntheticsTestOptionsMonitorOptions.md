# SyntheticsTestOptionsMonitorOptions

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**RenotifyInterval** | Pointer to **int64** | Time interval before renotifying if the test is still failing (in minutes). | [optional] 

## Methods

### NewSyntheticsTestOptionsMonitorOptions

`func NewSyntheticsTestOptionsMonitorOptions() *SyntheticsTestOptionsMonitorOptions`

NewSyntheticsTestOptionsMonitorOptions instantiates a new SyntheticsTestOptionsMonitorOptions object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsTestOptionsMonitorOptionsWithDefaults

`func NewSyntheticsTestOptionsMonitorOptionsWithDefaults() *SyntheticsTestOptionsMonitorOptions`

NewSyntheticsTestOptionsMonitorOptionsWithDefaults instantiates a new SyntheticsTestOptionsMonitorOptions object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRenotifyInterval

`func (o *SyntheticsTestOptionsMonitorOptions) GetRenotifyInterval() int64`

GetRenotifyInterval returns the RenotifyInterval field if non-nil, zero value otherwise.

### GetRenotifyIntervalOk

`func (o *SyntheticsTestOptionsMonitorOptions) GetRenotifyIntervalOk() (*int64, bool)`

GetRenotifyIntervalOk returns a tuple with the RenotifyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyInterval

`func (o *SyntheticsTestOptionsMonitorOptions) SetRenotifyInterval(v int64)`

SetRenotifyInterval sets RenotifyInterval field to given value.

### HasRenotifyInterval

`func (o *SyntheticsTestOptionsMonitorOptions) HasRenotifyInterval() bool`

HasRenotifyInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


