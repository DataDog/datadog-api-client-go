# SyntheticsVariableParser

## Properties

| Name      | Type                                                                            | Description                                                                 | Notes      |
| --------- | ------------------------------------------------------------------------------- | --------------------------------------------------------------------------- | ---------- |
| **Type**  | [**SyntheticsGlobalVariableParserType**](SyntheticsGlobalVariableParserType.md) |                                                                             |
| **Value** | Pointer to **string**                                                           | Regex or JSON path used for the parser. Not used with type &#x60;raw&#x60;. | [optional] |

## Methods

### NewSyntheticsVariableParser

`func NewSyntheticsVariableParser(type_ SyntheticsGlobalVariableParserType) *SyntheticsVariableParser`

NewSyntheticsVariableParser instantiates a new SyntheticsVariableParser object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsVariableParserWithDefaults

`func NewSyntheticsVariableParserWithDefaults() *SyntheticsVariableParser`

NewSyntheticsVariableParserWithDefaults instantiates a new SyntheticsVariableParser object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetType

`func (o *SyntheticsVariableParser) GetType() SyntheticsGlobalVariableParserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsVariableParser) GetTypeOk() (*SyntheticsGlobalVariableParserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsVariableParser) SetType(v SyntheticsGlobalVariableParserType)`

SetType sets Type field to given value.

### GetValue

`func (o *SyntheticsVariableParser) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SyntheticsVariableParser) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SyntheticsVariableParser) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SyntheticsVariableParser) HasValue() bool`

HasValue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
