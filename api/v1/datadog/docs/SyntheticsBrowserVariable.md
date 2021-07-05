# SyntheticsBrowserVariable

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Example** | Pointer to **string** | Example for the variable. | [optional] 
**Id** | Pointer to **string** | ID for the variable. | [optional] 
**Name** | **string** | Name of the variable. | 
**Pattern** | Pointer to **string** | Pattern of the variable. | [optional] 
**Type** | [**SyntheticsBrowserVariableType**](SyntheticsBrowserVariableType.md) |  | 

## Methods

### NewSyntheticsBrowserVariable

`func NewSyntheticsBrowserVariable(name string, type_ SyntheticsBrowserVariableType) *SyntheticsBrowserVariable`

NewSyntheticsBrowserVariable instantiates a new SyntheticsBrowserVariable object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsBrowserVariableWithDefaults

`func NewSyntheticsBrowserVariableWithDefaults() *SyntheticsBrowserVariable`

NewSyntheticsBrowserVariableWithDefaults instantiates a new SyntheticsBrowserVariable object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetExample

`func (o *SyntheticsBrowserVariable) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *SyntheticsBrowserVariable) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *SyntheticsBrowserVariable) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *SyntheticsBrowserVariable) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetId

`func (o *SyntheticsBrowserVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticsBrowserVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticsBrowserVariable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyntheticsBrowserVariable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsBrowserVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsBrowserVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsBrowserVariable) SetName(v string)`

SetName sets Name field to given value.


### GetPattern

`func (o *SyntheticsBrowserVariable) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SyntheticsBrowserVariable) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SyntheticsBrowserVariable) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *SyntheticsBrowserVariable) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsBrowserVariable) GetType() SyntheticsBrowserVariableType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsBrowserVariable) GetTypeOk() (*SyntheticsBrowserVariableType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsBrowserVariable) SetType(v SyntheticsBrowserVariableType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


