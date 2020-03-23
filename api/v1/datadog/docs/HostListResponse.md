# HostListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostList** | Pointer to [**[]Host**](Host.md) | TODO. | [optional] 
**TotalMatching** | Pointer to **int64** | TODO. | [optional] 
**TotalReturned** | Pointer to **int64** | TODO. | [optional] 

## Methods

### NewHostListResponse

`func NewHostListResponse() *HostListResponse`

NewHostListResponse instantiates a new HostListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostListResponseWithDefaults

`func NewHostListResponseWithDefaults() *HostListResponse`

NewHostListResponseWithDefaults instantiates a new HostListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostList

`func (o *HostListResponse) GetHostList() []Host`

GetHostList returns the HostList field if non-nil, zero value otherwise.

### GetHostListOk

`func (o *HostListResponse) GetHostListOk() ([]Host, bool)`

GetHostListOk returns a tuple with the HostList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostList

`func (o *HostListResponse) HasHostList() bool`

HasHostList returns a boolean if a field has been set.

### SetHostList

`func (o *HostListResponse) SetHostList(v []Host)`

SetHostList gets a reference to the given []Host and assigns it to the HostList field.

### GetTotalMatching

`func (o *HostListResponse) GetTotalMatching() int64`

GetTotalMatching returns the TotalMatching field if non-nil, zero value otherwise.

### GetTotalMatchingOk

`func (o *HostListResponse) GetTotalMatchingOk() (int64, bool)`

GetTotalMatchingOk returns a tuple with the TotalMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalMatching

`func (o *HostListResponse) HasTotalMatching() bool`

HasTotalMatching returns a boolean if a field has been set.

### SetTotalMatching

`func (o *HostListResponse) SetTotalMatching(v int64)`

SetTotalMatching gets a reference to the given int64 and assigns it to the TotalMatching field.

### GetTotalReturned

`func (o *HostListResponse) GetTotalReturned() int64`

GetTotalReturned returns the TotalReturned field if non-nil, zero value otherwise.

### GetTotalReturnedOk

`func (o *HostListResponse) GetTotalReturnedOk() (int64, bool)`

GetTotalReturnedOk returns a tuple with the TotalReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalReturned

`func (o *HostListResponse) HasTotalReturned() bool`

HasTotalReturned returns a boolean if a field has been set.

### SetTotalReturned

`func (o *HostListResponse) SetTotalReturned(v int64)`

SetTotalReturned gets a reference to the given int64 and assigns it to the TotalReturned field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


