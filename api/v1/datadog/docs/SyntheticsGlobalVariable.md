# SyntheticsGlobalVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Creation timestamp of the global variable. | [optional] 
**CreatedBy** | Pointer to **int32** | The ID of the user who created the global variable. | [optional] 
**Description** | Pointer to **string** | Description of the global variable. | [optional] 
**Id** | Pointer to **string** | Unique identifier of the global variable. | [optional] 
**ModifiedAt** | Pointer to [**time.Time**](time.Time.md) | Modification timestamp of the global variable. | [optional] 
**Name** | Pointer to **string** | Name of the global variable. | [optional] 
**Tags** | Pointer to **[]string** | Tags of the global variable. | [optional] 
**Value** | Pointer to [**SyntheticsGlobalVariableValue**](SyntheticsGlobalVariable_value.md) |  | [optional] 

## Methods

### NewSyntheticsGlobalVariable

`func NewSyntheticsGlobalVariable() *SyntheticsGlobalVariable`

NewSyntheticsGlobalVariable instantiates a new SyntheticsGlobalVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsGlobalVariableWithDefaults

`func NewSyntheticsGlobalVariableWithDefaults() *SyntheticsGlobalVariable`

NewSyntheticsGlobalVariableWithDefaults instantiates a new SyntheticsGlobalVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SyntheticsGlobalVariable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SyntheticsGlobalVariable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SyntheticsGlobalVariable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SyntheticsGlobalVariable) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SyntheticsGlobalVariable) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SyntheticsGlobalVariable) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SyntheticsGlobalVariable) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SyntheticsGlobalVariable) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *SyntheticsGlobalVariable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyntheticsGlobalVariable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyntheticsGlobalVariable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyntheticsGlobalVariable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SyntheticsGlobalVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticsGlobalVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticsGlobalVariable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyntheticsGlobalVariable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SyntheticsGlobalVariable) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SyntheticsGlobalVariable) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SyntheticsGlobalVariable) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SyntheticsGlobalVariable) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsGlobalVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsGlobalVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsGlobalVariable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyntheticsGlobalVariable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *SyntheticsGlobalVariable) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticsGlobalVariable) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyntheticsGlobalVariable) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SyntheticsGlobalVariable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetValue

`func (o *SyntheticsGlobalVariable) GetValue() SyntheticsGlobalVariableValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SyntheticsGlobalVariable) GetValueOk() (*SyntheticsGlobalVariableValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SyntheticsGlobalVariable) SetValue(v SyntheticsGlobalVariableValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *SyntheticsGlobalVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


