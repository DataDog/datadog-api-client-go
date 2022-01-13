# SyntheticsAPIStep

## Properties

| Name                | Type                                                                       | Description                                                                                                                                           | Notes                      |
| ------------------- | -------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------- |
| **AllowFailure**    | Pointer to **bool**                                                        | Determines whether or not to continue with test if this step fails.                                                                                   | [optional]                 |
| **Assertions**      | Pointer to [**[]SyntheticsAssertion**](SyntheticsAssertion.md)             | Array of assertions used for the test.                                                                                                                | [optional] [default to []] |
| **ExtractedValues** | Pointer to [**[]SyntheticsParsingOptions**](SyntheticsParsingOptions.md)   | Array of values to parse and save as variables from the response.                                                                                     | [optional]                 |
| **IsCritical**      | Pointer to **bool**                                                        | Determines whether or not to consider the entire test as failed if this step fails. Can be used only if &#x60;allowFailure&#x60; is &#x60;true&#x60;. | [optional]                 |
| **Name**            | Pointer to **string**                                                      | The name of the step.                                                                                                                                 | [optional]                 |
| **Request**         | Pointer to [**SyntheticsTestRequest**](SyntheticsTestRequest.md)           |                                                                                                                                                       | [optional]                 |
| **Retry**           | Pointer to [**SyntheticsTestOptionsRetry**](SyntheticsTestOptionsRetry.md) |                                                                                                                                                       | [optional]                 |
| **Subtype**         | Pointer to [**SyntheticsAPIStepSubtype**](SyntheticsAPIStepSubtype.md)     |                                                                                                                                                       | [optional]                 |

## Methods

### NewSyntheticsAPIStep

`func NewSyntheticsAPIStep() *SyntheticsAPIStep`

NewSyntheticsAPIStep instantiates a new SyntheticsAPIStep object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsAPIStepWithDefaults

`func NewSyntheticsAPIStepWithDefaults() *SyntheticsAPIStep`

NewSyntheticsAPIStepWithDefaults instantiates a new SyntheticsAPIStep object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAllowFailure

`func (o *SyntheticsAPIStep) GetAllowFailure() bool`

GetAllowFailure returns the AllowFailure field if non-nil, zero value otherwise.

### GetAllowFailureOk

`func (o *SyntheticsAPIStep) GetAllowFailureOk() (*bool, bool)`

GetAllowFailureOk returns a tuple with the AllowFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFailure

`func (o *SyntheticsAPIStep) SetAllowFailure(v bool)`

SetAllowFailure sets AllowFailure field to given value.

### HasAllowFailure

`func (o *SyntheticsAPIStep) HasAllowFailure() bool`

HasAllowFailure returns a boolean if a field has been set.

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

### GetIsCritical

`func (o *SyntheticsAPIStep) GetIsCritical() bool`

GetIsCritical returns the IsCritical field if non-nil, zero value otherwise.

### GetIsCriticalOk

`func (o *SyntheticsAPIStep) GetIsCriticalOk() (*bool, bool)`

GetIsCriticalOk returns a tuple with the IsCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCritical

`func (o *SyntheticsAPIStep) SetIsCritical(v bool)`

SetIsCritical sets IsCritical field to given value.

### HasIsCritical

`func (o *SyntheticsAPIStep) HasIsCritical() bool`

HasIsCritical returns a boolean if a field has been set.

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

### GetRetry

`func (o *SyntheticsAPIStep) GetRetry() SyntheticsTestOptionsRetry`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *SyntheticsAPIStep) GetRetryOk() (*SyntheticsTestOptionsRetry, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *SyntheticsAPIStep) SetRetry(v SyntheticsTestOptionsRetry)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *SyntheticsAPIStep) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

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
