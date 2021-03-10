# SyntheticsAPIStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | Pointer to [**[]SyntheticsAssertion**](SyntheticsAssertion.md) | Array of assertions used for the test. | [optional] [default to []]
**ExtractedValues** | Pointer to [**[]SyntheticsParsingOptions**](SyntheticsParsingOptions.md) | Array of values to parse and save as variables from the response. | [optional] 
**Name** | Pointer to **string** | The name of the step. | [optional] 
**Request** | Pointer to [**SyntheticsTestRequest**](SyntheticsTestRequest.md) |  | [optional] 
**Subtype** | Pointer to [**SyntheticsAPIStepSubtype**](SyntheticsAPIStepSubtype.md) |  | [optional] 

## Methods

### NewSyntheticsAPIStep

`func NewSyntheticsAPIStep() *SyntheticsAPIStep`

NewSyntheticsAPIStep instantiates a new SyntheticsAPIStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsAPIStepWithDefaults

`func NewSyntheticsAPIStepWithDefaults() *SyntheticsAPIStep`

NewSyntheticsAPIStepWithDefaults instantiates a new SyntheticsAPIStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *SyntheticsAPIStep) GetAssertions() []SyntheticsAssertion`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *SyntheticsAPIStep) GetAssertionsOk() (*[]SyntheticsAssertion, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *SyntheticsAPIStep) SetAssertions(v []SyntheticsAssertion)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *SyntheticsAPIStep) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.

### GetExtractedValues

`func (o *SyntheticsAPIStep) GetExtractedValues() []SyntheticsParsingOptions`

GetExtractedValues returns the ExtractedValues field if non-nil, zero value otherwise.

### GetExtractedValuesOk

`func (o *SyntheticsAPIStep) GetExtractedValuesOk() (*[]SyntheticsParsingOptions, bool)`

GetExtractedValuesOk returns a tuple with the ExtractedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedValues

`func (o *SyntheticsAPIStep) SetExtractedValues(v []SyntheticsParsingOptions)`

SetExtractedValues sets ExtractedValues field to given value.

### HasExtractedValues

`func (o *SyntheticsAPIStep) HasExtractedValues() bool`

HasExtractedValues returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsAPIStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsAPIStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsAPIStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyntheticsAPIStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequest

`func (o *SyntheticsAPIStep) GetRequest() SyntheticsTestRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *SyntheticsAPIStep) GetRequestOk() (*SyntheticsTestRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *SyntheticsAPIStep) SetRequest(v SyntheticsTestRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *SyntheticsAPIStep) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSubtype

`func (o *SyntheticsAPIStep) GetSubtype() SyntheticsAPIStepSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SyntheticsAPIStep) GetSubtypeOk() (*SyntheticsAPIStepSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SyntheticsAPIStep) SetSubtype(v SyntheticsAPIStepSubtype)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *SyntheticsAPIStep) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


