# SyntheticsTestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptSelfSigned** | Pointer to **bool** | For browser test, whether or not the test should allow self signed certificate. | [optional] 
**AllowInsecure** | Pointer to **bool** | Allows loading insecure content for an HTTP request. | [optional] 
**DeviceIds** | Pointer to [**[]SyntheticsDeviceID**](SyntheticsDeviceID.md) | Array with the different device IDs used to run the test. | [optional] 
**FollowRedirects** | Pointer to **bool** | For API SSL test, whether or not the test should follow redirects. | [optional] 
**MinFailureDuration** | Pointer to **int64** | Minimum amount of time before declaring the test has failed. | [optional] 
**MinLocationFailed** | Pointer to **int64** | Minimum amount of locations that are allowed to fail for the test. | [optional] 
**Retry** | Pointer to [**SyntheticsTestOptionsRetry**](SyntheticsTestOptions_retry.md) |  | [optional] 
**TickEvery** | Pointer to [**SyntheticsTickInterval**](SyntheticsTickInterval.md) |  | [optional] 

## Methods

### NewSyntheticsTestOptions

`func NewSyntheticsTestOptions() *SyntheticsTestOptions`

NewSyntheticsTestOptions instantiates a new SyntheticsTestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsTestOptionsWithDefaults

`func NewSyntheticsTestOptionsWithDefaults() *SyntheticsTestOptions`

NewSyntheticsTestOptionsWithDefaults instantiates a new SyntheticsTestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptSelfSigned

`func (o *SyntheticsTestOptions) GetAcceptSelfSigned() bool`

GetAcceptSelfSigned returns the AcceptSelfSigned field if non-nil, zero value otherwise.

### GetAcceptSelfSignedOk

`func (o *SyntheticsTestOptions) GetAcceptSelfSignedOk() (*bool, bool)`

GetAcceptSelfSignedOk returns a tuple with the AcceptSelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptSelfSigned

`func (o *SyntheticsTestOptions) SetAcceptSelfSigned(v bool)`

SetAcceptSelfSigned sets AcceptSelfSigned field to given value.

### HasAcceptSelfSigned

`func (o *SyntheticsTestOptions) HasAcceptSelfSigned() bool`

HasAcceptSelfSigned returns a boolean if a field has been set.

### GetAllowInsecure

`func (o *SyntheticsTestOptions) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *SyntheticsTestOptions) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *SyntheticsTestOptions) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *SyntheticsTestOptions) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetDeviceIds

`func (o *SyntheticsTestOptions) GetDeviceIds() []SyntheticsDeviceID`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *SyntheticsTestOptions) GetDeviceIdsOk() (*[]SyntheticsDeviceID, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *SyntheticsTestOptions) SetDeviceIds(v []SyntheticsDeviceID)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *SyntheticsTestOptions) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetFollowRedirects

`func (o *SyntheticsTestOptions) GetFollowRedirects() bool`

GetFollowRedirects returns the FollowRedirects field if non-nil, zero value otherwise.

### GetFollowRedirectsOk

`func (o *SyntheticsTestOptions) GetFollowRedirectsOk() (*bool, bool)`

GetFollowRedirectsOk returns a tuple with the FollowRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRedirects

`func (o *SyntheticsTestOptions) SetFollowRedirects(v bool)`

SetFollowRedirects sets FollowRedirects field to given value.

### HasFollowRedirects

`func (o *SyntheticsTestOptions) HasFollowRedirects() bool`

HasFollowRedirects returns a boolean if a field has been set.

### GetMinFailureDuration

`func (o *SyntheticsTestOptions) GetMinFailureDuration() int64`

GetMinFailureDuration returns the MinFailureDuration field if non-nil, zero value otherwise.

### GetMinFailureDurationOk

`func (o *SyntheticsTestOptions) GetMinFailureDurationOk() (*int64, bool)`

GetMinFailureDurationOk returns a tuple with the MinFailureDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFailureDuration

`func (o *SyntheticsTestOptions) SetMinFailureDuration(v int64)`

SetMinFailureDuration sets MinFailureDuration field to given value.

### HasMinFailureDuration

`func (o *SyntheticsTestOptions) HasMinFailureDuration() bool`

HasMinFailureDuration returns a boolean if a field has been set.

### GetMinLocationFailed

`func (o *SyntheticsTestOptions) GetMinLocationFailed() int64`

GetMinLocationFailed returns the MinLocationFailed field if non-nil, zero value otherwise.

### GetMinLocationFailedOk

`func (o *SyntheticsTestOptions) GetMinLocationFailedOk() (*int64, bool)`

GetMinLocationFailedOk returns a tuple with the MinLocationFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLocationFailed

`func (o *SyntheticsTestOptions) SetMinLocationFailed(v int64)`

SetMinLocationFailed sets MinLocationFailed field to given value.

### HasMinLocationFailed

`func (o *SyntheticsTestOptions) HasMinLocationFailed() bool`

HasMinLocationFailed returns a boolean if a field has been set.

### GetRetry

`func (o *SyntheticsTestOptions) GetRetry() SyntheticsTestOptionsRetry`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *SyntheticsTestOptions) GetRetryOk() (*SyntheticsTestOptionsRetry, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *SyntheticsTestOptions) SetRetry(v SyntheticsTestOptionsRetry)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *SyntheticsTestOptions) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetTickEvery

`func (o *SyntheticsTestOptions) GetTickEvery() SyntheticsTickInterval`

GetTickEvery returns the TickEvery field if non-nil, zero value otherwise.

### GetTickEveryOk

`func (o *SyntheticsTestOptions) GetTickEveryOk() (*SyntheticsTickInterval, bool)`

GetTickEveryOk returns a tuple with the TickEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickEvery

`func (o *SyntheticsTestOptions) SetTickEvery(v SyntheticsTickInterval)`

SetTickEvery sets TickEvery field to given value.

### HasTickEvery

`func (o *SyntheticsTestOptions) HasTickEvery() bool`

HasTickEvery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


