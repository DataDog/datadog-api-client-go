# SyntheticsTestOptionsRetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The amount of location that needs to fail for the test to be retried. | [optional] 
**Interval** | Pointer to **float64** | The interval over which the amount of location needed to fail for the test to be retried. | [optional] 

## Methods

### NewSyntheticsTestOptionsRetry

`func NewSyntheticsTestOptionsRetry() *SyntheticsTestOptionsRetry`

NewSyntheticsTestOptionsRetry instantiates a new SyntheticsTestOptionsRetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsTestOptionsRetryWithDefaults

`func NewSyntheticsTestOptionsRetryWithDefaults() *SyntheticsTestOptionsRetry`

NewSyntheticsTestOptionsRetryWithDefaults instantiates a new SyntheticsTestOptionsRetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SyntheticsTestOptionsRetry) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SyntheticsTestOptionsRetry) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SyntheticsTestOptionsRetry) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SyntheticsTestOptionsRetry) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetInterval

`func (o *SyntheticsTestOptionsRetry) GetInterval() float64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SyntheticsTestOptionsRetry) GetIntervalOk() (*float64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SyntheticsTestOptionsRetry) SetInterval(v float64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *SyntheticsTestOptionsRetry) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


