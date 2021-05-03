# SyntheticsGlobalVariable

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Description** | **string** | Description of the global variable. | 
**Id** | Pointer to **string** | Unique identifier of the global variable. | [optional] [readonly] 
**Name** | **string** | Name of the global variable. | 
**ParseTestOptions** | Pointer to [**SyntheticsGlobalVariableParseTestOptions**](SyntheticsGlobalVariableParseTestOptions.md) |  | [optional] 
**ParseTestPublicId** | Pointer to **string** | A Synthetic test ID to use as a test to generate the variable value. | [optional] 
**Tags** | **[]string** | Tags of the global variable. | 
**Value** | [**SyntheticsGlobalVariableValue**](SyntheticsGlobalVariableValue.md) |  | 

## Methods

### NewSyntheticsGlobalVariable

`func NewSyntheticsGlobalVariable(description string, name string, tags []string, value SyntheticsGlobalVariableValue, ) *SyntheticsGlobalVariable`

NewSyntheticsGlobalVariable instantiates a new SyntheticsGlobalVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsGlobalVariableWithDefaults

`func NewSyntheticsGlobalVariableWithDefaults() *SyntheticsGlobalVariable`

NewSyntheticsGlobalVariableWithDefaults instantiates a new SyntheticsGlobalVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetParseTestOptions

`func (o *SyntheticsGlobalVariable) GetParseTestOptions() SyntheticsGlobalVariableParseTestOptions`

GetParseTestOptions returns the ParseTestOptions field if non-nil, zero value otherwise.

### GetParseTestOptionsOk

`func (o *SyntheticsGlobalVariable) GetParseTestOptionsOk() (*SyntheticsGlobalVariableParseTestOptions, bool)`

GetParseTestOptionsOk returns a tuple with the ParseTestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseTestOptions

`func (o *SyntheticsGlobalVariable) SetParseTestOptions(v SyntheticsGlobalVariableParseTestOptions)`

SetParseTestOptions sets ParseTestOptions field to given value.

### HasParseTestOptions

`func (o *SyntheticsGlobalVariable) HasParseTestOptions() bool`

HasParseTestOptions returns a boolean if a field has been set.

### GetParseTestPublicId

`func (o *SyntheticsGlobalVariable) GetParseTestPublicId() string`

GetParseTestPublicId returns the ParseTestPublicId field if non-nil, zero value otherwise.

### GetParseTestPublicIdOk

`func (o *SyntheticsGlobalVariable) GetParseTestPublicIdOk() (*string, bool)`

GetParseTestPublicIdOk returns a tuple with the ParseTestPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseTestPublicId

`func (o *SyntheticsGlobalVariable) SetParseTestPublicId(v string)`

SetParseTestPublicId sets ParseTestPublicId field to given value.

### HasParseTestPublicId

`func (o *SyntheticsGlobalVariable) HasParseTestPublicId() bool`

HasParseTestPublicId returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


