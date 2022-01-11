# SyntheticsListGlobalVariablesResponse

## Properties

| Name          | Type                                                                     | Description                          | Notes      |
| ------------- | ------------------------------------------------------------------------ | ------------------------------------ | ---------- |
| **Variables** | Pointer to [**[]SyntheticsGlobalVariable**](SyntheticsGlobalVariable.md) | Array of Synthetic global variables. | [optional] |

## Methods

### NewSyntheticsListGlobalVariablesResponse

`func NewSyntheticsListGlobalVariablesResponse() *SyntheticsListGlobalVariablesResponse`

NewSyntheticsListGlobalVariablesResponse instantiates a new SyntheticsListGlobalVariablesResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsListGlobalVariablesResponseWithDefaults

`func NewSyntheticsListGlobalVariablesResponseWithDefaults() *SyntheticsListGlobalVariablesResponse`

NewSyntheticsListGlobalVariablesResponseWithDefaults instantiates a new SyntheticsListGlobalVariablesResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetVariables

`func (o *SyntheticsListGlobalVariablesResponse) GetVariables() []SyntheticsGlobalVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SyntheticsListGlobalVariablesResponse) GetVariablesOk() (*[]SyntheticsGlobalVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SyntheticsListGlobalVariablesResponse) SetVariables(v []SyntheticsGlobalVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SyntheticsListGlobalVariablesResponse) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
