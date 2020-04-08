# SyntheticsBrowserError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | TODO. | 
**Name** | Pointer to **string** | TODO. | 
**StatusCode** | Pointer to **int64** | TODO. | [optional] 
**Type** | Pointer to [**SyntheticsBrowserErrorType**](SyntheticsBrowserErrorType.md) |  | 

## Methods

### NewSyntheticsBrowserError

`func NewSyntheticsBrowserError(description string, name string, type_ SyntheticsBrowserErrorType, ) *SyntheticsBrowserError`

NewSyntheticsBrowserError instantiates a new SyntheticsBrowserError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsBrowserErrorWithDefaults

`func NewSyntheticsBrowserErrorWithDefaults() *SyntheticsBrowserError`

NewSyntheticsBrowserErrorWithDefaults instantiates a new SyntheticsBrowserError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SyntheticsBrowserError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyntheticsBrowserError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyntheticsBrowserError) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *SyntheticsBrowserError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsBrowserError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsBrowserError) SetName(v string)`

SetName sets Name field to given value.


### GetStatusCode

`func (o *SyntheticsBrowserError) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SyntheticsBrowserError) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SyntheticsBrowserError) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *SyntheticsBrowserError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsBrowserError) GetType() SyntheticsBrowserErrorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsBrowserError) GetTypeOk() (*SyntheticsBrowserErrorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsBrowserError) SetType(v SyntheticsBrowserErrorType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


