# SyntheticsStepDetailWarnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | TODO. | 
**Type** | Pointer to [**SyntheticsWarningType**](SyntheticsWarningType.md) |  | 

## Methods

### NewSyntheticsStepDetailWarnings

`func NewSyntheticsStepDetailWarnings(message string, type_ SyntheticsWarningType, ) *SyntheticsStepDetailWarnings`

NewSyntheticsStepDetailWarnings instantiates a new SyntheticsStepDetailWarnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsStepDetailWarningsWithDefaults

`func NewSyntheticsStepDetailWarningsWithDefaults() *SyntheticsStepDetailWarnings`

NewSyntheticsStepDetailWarningsWithDefaults instantiates a new SyntheticsStepDetailWarnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SyntheticsStepDetailWarnings) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyntheticsStepDetailWarnings) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyntheticsStepDetailWarnings) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *SyntheticsStepDetailWarnings) GetType() SyntheticsWarningType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsStepDetailWarnings) GetTypeOk() (*SyntheticsWarningType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsStepDetailWarnings) SetType(v SyntheticsWarningType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


