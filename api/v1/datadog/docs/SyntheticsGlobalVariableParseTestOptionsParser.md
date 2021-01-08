# SyntheticsGlobalVariableParseTestOptionsParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SyntheticsGlobalVariableParserType**](SyntheticsGlobalVariableParserType.md) |  | 
**Value** | Pointer to **string** | Regex or JSON path used for the parser. Not used with type &#x60;raw&#x60;. | [optional] 

## Methods

### NewSyntheticsGlobalVariableParseTestOptionsParser

`func NewSyntheticsGlobalVariableParseTestOptionsParser(type_ SyntheticsGlobalVariableParserType, ) *SyntheticsGlobalVariableParseTestOptionsParser`

NewSyntheticsGlobalVariableParseTestOptionsParser instantiates a new SyntheticsGlobalVariableParseTestOptionsParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsGlobalVariableParseTestOptionsParserWithDefaults

`func NewSyntheticsGlobalVariableParseTestOptionsParserWithDefaults() *SyntheticsGlobalVariableParseTestOptionsParser`

NewSyntheticsGlobalVariableParseTestOptionsParserWithDefaults instantiates a new SyntheticsGlobalVariableParseTestOptionsParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SyntheticsGlobalVariableParseTestOptionsParser) GetType() SyntheticsGlobalVariableParserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsGlobalVariableParseTestOptionsParser) GetTypeOk() (*SyntheticsGlobalVariableParserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsGlobalVariableParseTestOptionsParser) SetType(v SyntheticsGlobalVariableParserType)`

SetType sets Type field to given value.


### GetValue

`func (o *SyntheticsGlobalVariableParseTestOptionsParser) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SyntheticsGlobalVariableParseTestOptionsParser) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SyntheticsGlobalVariableParseTestOptionsParser) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SyntheticsGlobalVariableParseTestOptionsParser) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


