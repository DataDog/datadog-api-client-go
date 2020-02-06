# SyntheticsTestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptSelfSigned** | Pointer to **bool** |  | [optional] 
**DeviceIds** | Pointer to [**[]SyntheticsDeviceId**](SyntheticsDeviceID.md) |  | [optional] 
**FollowRedirects** | Pointer to **bool** |  | [optional] 
**MinLocationFailed** | Pointer to **int64** |  | [optional] 
**Retry** | Pointer to [**SyntheticsTestOptionsRetry**](SyntheticsTestOptions_retry.md) |  | [optional] 
**TickEvery** | Pointer to **int64** |  | [optional] 

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

`func (o *SyntheticsTestOptions) GetAcceptSelfSignedOk() (bool, bool)`

GetAcceptSelfSignedOk returns a tuple with the AcceptSelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptSelfSigned

`func (o *SyntheticsTestOptions) HasAcceptSelfSigned() bool`

HasAcceptSelfSigned returns a boolean if a field has been set.

### SetAcceptSelfSigned

`func (o *SyntheticsTestOptions) SetAcceptSelfSigned(v bool)`

SetAcceptSelfSigned gets a reference to the given bool and assigns it to the AcceptSelfSigned field.

### GetDeviceIds

`func (o *SyntheticsTestOptions) GetDeviceIds() []SyntheticsDeviceId`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *SyntheticsTestOptions) GetDeviceIdsOk() ([]SyntheticsDeviceId, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceIds

`func (o *SyntheticsTestOptions) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### SetDeviceIds

`func (o *SyntheticsTestOptions) SetDeviceIds(v []SyntheticsDeviceId)`

SetDeviceIds gets a reference to the given []SyntheticsDeviceId and assigns it to the DeviceIds field.

### GetFollowRedirects

`func (o *SyntheticsTestOptions) GetFollowRedirects() bool`

GetFollowRedirects returns the FollowRedirects field if non-nil, zero value otherwise.

### GetFollowRedirectsOk

`func (o *SyntheticsTestOptions) GetFollowRedirectsOk() (bool, bool)`

GetFollowRedirectsOk returns a tuple with the FollowRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFollowRedirects

`func (o *SyntheticsTestOptions) HasFollowRedirects() bool`

HasFollowRedirects returns a boolean if a field has been set.

### SetFollowRedirects

`func (o *SyntheticsTestOptions) SetFollowRedirects(v bool)`

SetFollowRedirects gets a reference to the given bool and assigns it to the FollowRedirects field.

### GetMinLocationFailed

`func (o *SyntheticsTestOptions) GetMinLocationFailed() int64`

GetMinLocationFailed returns the MinLocationFailed field if non-nil, zero value otherwise.

### GetMinLocationFailedOk

`func (o *SyntheticsTestOptions) GetMinLocationFailedOk() (int64, bool)`

GetMinLocationFailedOk returns a tuple with the MinLocationFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinLocationFailed

`func (o *SyntheticsTestOptions) HasMinLocationFailed() bool`

HasMinLocationFailed returns a boolean if a field has been set.

### SetMinLocationFailed

`func (o *SyntheticsTestOptions) SetMinLocationFailed(v int64)`

SetMinLocationFailed gets a reference to the given int64 and assigns it to the MinLocationFailed field.

### GetRetry

`func (o *SyntheticsTestOptions) GetRetry() SyntheticsTestOptionsRetry`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *SyntheticsTestOptions) GetRetryOk() (SyntheticsTestOptionsRetry, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRetry

`func (o *SyntheticsTestOptions) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### SetRetry

`func (o *SyntheticsTestOptions) SetRetry(v SyntheticsTestOptionsRetry)`

SetRetry gets a reference to the given SyntheticsTestOptionsRetry and assigns it to the Retry field.

### GetTickEvery

`func (o *SyntheticsTestOptions) GetTickEvery() int64`

GetTickEvery returns the TickEvery field if non-nil, zero value otherwise.

### GetTickEveryOk

`func (o *SyntheticsTestOptions) GetTickEveryOk() (int64, bool)`

GetTickEveryOk returns a tuple with the TickEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTickEvery

`func (o *SyntheticsTestOptions) HasTickEvery() bool`

HasTickEvery returns a boolean if a field has been set.

### SetTickEvery

`func (o *SyntheticsTestOptions) SetTickEvery(v int64)`

SetTickEvery gets a reference to the given int64 and assigns it to the TickEvery field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


