# SyntheticsParsingOptions

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Field** | Pointer to **string** | When type is &#x60;http_header&#x60;, name of the header to use to extract the value. | [optional] 
**Name** | Pointer to **string** | Name of the variable to extract. | [optional] 
**Parser** | Pointer to [**SyntheticsVariableParser**](SyntheticsVariableParser.md) |  | [optional] 
**Type** | Pointer to [**SyntheticsGlobalVariableParseTestOptionsType**](SyntheticsGlobalVariableParseTestOptionsType.md) |  | [optional] 

## Methods

### NewSyntheticsParsingOptions

`func NewSyntheticsParsingOptions() *SyntheticsParsingOptions`

NewSyntheticsParsingOptions instantiates a new SyntheticsParsingOptions object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsParsingOptionsWithDefaults

`func NewSyntheticsParsingOptionsWithDefaults() *SyntheticsParsingOptions`

NewSyntheticsParsingOptionsWithDefaults instantiates a new SyntheticsParsingOptions object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetField

`func (o *SyntheticsParsingOptions) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SyntheticsParsingOptions) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SyntheticsParsingOptions) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *SyntheticsParsingOptions) HasField() bool`

HasField returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsParsingOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsParsingOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsParsingOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyntheticsParsingOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParser

`func (o *SyntheticsParsingOptions) GetParser() SyntheticsVariableParser`

GetParser returns the Parser field if non-nil, zero value otherwise.

### GetParserOk

`func (o *SyntheticsParsingOptions) GetParserOk() (*SyntheticsVariableParser, bool)`

GetParserOk returns a tuple with the Parser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParser

`func (o *SyntheticsParsingOptions) SetParser(v SyntheticsVariableParser)`

SetParser sets Parser field to given value.

### HasParser

`func (o *SyntheticsParsingOptions) HasParser() bool`

HasParser returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsParsingOptions) GetType() SyntheticsGlobalVariableParseTestOptionsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsParsingOptions) GetTypeOk() (*SyntheticsGlobalVariableParseTestOptionsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsParsingOptions) SetType(v SyntheticsGlobalVariableParseTestOptionsType)`

SetType sets Type field to given value.

### HasType

`func (o *SyntheticsParsingOptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


