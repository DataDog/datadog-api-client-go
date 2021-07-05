# SyntheticsPrivateLocation

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Description** | **string** | Description of the private location. | 
**Id** | Pointer to **string** | Unique identifier of the private location. | [optional] [readonly] 
**Name** | **string** | Name of the private location. | 
**Secrets** | Pointer to [**SyntheticsPrivateLocationSecrets**](SyntheticsPrivateLocationSecrets.md) |  | [optional] 
**Tags** | **[]string** | Array of tags attached to the private location. | 

## Methods

### NewSyntheticsPrivateLocation

`func NewSyntheticsPrivateLocation(description string, name string, tags []string) *SyntheticsPrivateLocation`

NewSyntheticsPrivateLocation instantiates a new SyntheticsPrivateLocation object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsPrivateLocationWithDefaults

`func NewSyntheticsPrivateLocationWithDefaults() *SyntheticsPrivateLocation`

NewSyntheticsPrivateLocationWithDefaults instantiates a new SyntheticsPrivateLocation object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDescription

`func (o *SyntheticsPrivateLocation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyntheticsPrivateLocation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyntheticsPrivateLocation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *SyntheticsPrivateLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticsPrivateLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticsPrivateLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyntheticsPrivateLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsPrivateLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsPrivateLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsPrivateLocation) SetName(v string)`

SetName sets Name field to given value.


### GetSecrets

`func (o *SyntheticsPrivateLocation) GetSecrets() SyntheticsPrivateLocationSecrets`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *SyntheticsPrivateLocation) GetSecretsOk() (*SyntheticsPrivateLocationSecrets, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *SyntheticsPrivateLocation) SetSecrets(v SyntheticsPrivateLocationSecrets)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *SyntheticsPrivateLocation) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetTags

`func (o *SyntheticsPrivateLocation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticsPrivateLocation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyntheticsPrivateLocation) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


