# SyntheticsGlobalVariableParseTestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | When type is &#x60;http_header&#x60;, name of the header to use to extract the value. | [optional] 
**Parser** | [**SyntheticsVariableParser**](SyntheticsVariableParser.md) |  | 
**Type** | [**SyntheticsGlobalVariableParseTestOptionsType**](SyntheticsGlobalVariableParseTestOptionsType.md) |  | 

## Methods

### NewSyntheticsGlobalVariableParseTestOptions

`func NewSyntheticsGlobalVariableParseTestOptions(parser SyntheticsVariableParser, type_ SyntheticsGlobalVariableParseTestOptionsType, ) *SyntheticsGlobalVariableParseTestOptions`

NewSyntheticsGlobalVariableParseTestOptions instantiates a new SyntheticsGlobalVariableParseTestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsGlobalVariableParseTestOptionsWithDefaults

`func NewSyntheticsGlobalVariableParseTestOptionsWithDefaults() *SyntheticsGlobalVariableParseTestOptions`

NewSyntheticsGlobalVariableParseTestOptionsWithDefaults instantiates a new SyntheticsGlobalVariableParseTestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *SyntheticsGlobalVariableParseTestOptions) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SyntheticsGlobalVariableParseTestOptions) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SyntheticsGlobalVariableParseTestOptions) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *SyntheticsGlobalVariableParseTestOptions) HasField() bool`

HasField returns a boolean if a field has been set.

### GetParser

`func (o *SyntheticsGlobalVariableParseTestOptions) GetParser() SyntheticsVariableParser`

GetParser returns the Parser field if non-nil, zero value otherwise.

### GetParserOk

`func (o *SyntheticsGlobalVariableParseTestOptions) GetParserOk() (*SyntheticsVariableParser, bool)`

GetParserOk returns a tuple with the Parser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParser

`func (o *SyntheticsGlobalVariableParseTestOptions) SetParser(v SyntheticsVariableParser)`

SetParser sets Parser field to given value.


### GetType

`func (o *SyntheticsGlobalVariableParseTestOptions) GetType() SyntheticsGlobalVariableParseTestOptionsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsGlobalVariableParseTestOptions) GetTypeOk() (*SyntheticsGlobalVariableParseTestOptionsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsGlobalVariableParseTestOptions) SetType(v SyntheticsGlobalVariableParseTestOptionsType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


