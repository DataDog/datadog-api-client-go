# SyntheticsConfigVariable

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Example** | Pointer to **string** | Example for the variable. | [optional] 
**Id** | **string** | ID of the variable for global variables. | 
**Name** | **string** | Name of the variable. | 
**Pattern** | Pointer to **string** | Pattern of the variable. | [optional] 
**Type** | [**SyntheticsConfigVariableType**](SyntheticsConfigVariableType.md) |  | 

## Methods

### NewSyntheticsConfigVariable

`func NewSyntheticsConfigVariable(id string, name string, type_ SyntheticsConfigVariableType) *SyntheticsConfigVariable`

NewSyntheticsConfigVariable instantiates a new SyntheticsConfigVariable object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsConfigVariableWithDefaults

`func NewSyntheticsConfigVariableWithDefaults() *SyntheticsConfigVariable`

NewSyntheticsConfigVariableWithDefaults instantiates a new SyntheticsConfigVariable object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetExample

`func (o *SyntheticsConfigVariable) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *SyntheticsConfigVariable) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *SyntheticsConfigVariable) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *SyntheticsConfigVariable) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetId

`func (o *SyntheticsConfigVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticsConfigVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticsConfigVariable) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SyntheticsConfigVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsConfigVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsConfigVariable) SetName(v string)`

SetName sets Name field to given value.


### GetPattern

`func (o *SyntheticsConfigVariable) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SyntheticsConfigVariable) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SyntheticsConfigVariable) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *SyntheticsConfigVariable) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsConfigVariable) GetType() SyntheticsConfigVariableType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsConfigVariable) GetTypeOk() (*SyntheticsConfigVariableType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsConfigVariable) SetType(v SyntheticsConfigVariableType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


