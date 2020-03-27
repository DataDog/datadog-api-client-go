# SyntheticsGetTestLatestResultsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTs** | Pointer to **float64** |  | 
**ProbeDc** | Pointer to **[]string** |  | [optional] 
**ToTs** | Pointer to **float64** |  | 

## Methods

### NewSyntheticsGetTestLatestResultsPayload

`func NewSyntheticsGetTestLatestResultsPayload(fromTs float64, toTs float64, ) *SyntheticsGetTestLatestResultsPayload`

NewSyntheticsGetTestLatestResultsPayload instantiates a new SyntheticsGetTestLatestResultsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsGetTestLatestResultsPayloadWithDefaults

`func NewSyntheticsGetTestLatestResultsPayloadWithDefaults() *SyntheticsGetTestLatestResultsPayload`

NewSyntheticsGetTestLatestResultsPayloadWithDefaults instantiates a new SyntheticsGetTestLatestResultsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromTs

`func (o *SyntheticsGetTestLatestResultsPayload) GetFromTs() float64`

GetFromTs returns the FromTs field if non-nil, zero value otherwise.

### GetFromTsOk

`func (o *SyntheticsGetTestLatestResultsPayload) GetFromTsOk() (*float64, bool)`

GetFromTsOk returns a tuple with the FromTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTs

`func (o *SyntheticsGetTestLatestResultsPayload) SetFromTs(v float64)`

SetFromTs sets FromTs field to given value.


### GetProbeDc

`func (o *SyntheticsGetTestLatestResultsPayload) GetProbeDc() []string`

GetProbeDc returns the ProbeDc field if non-nil, zero value otherwise.

### GetProbeDcOk

`func (o *SyntheticsGetTestLatestResultsPayload) GetProbeDcOk() (*[]string, bool)`

GetProbeDcOk returns a tuple with the ProbeDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeDc

`func (o *SyntheticsGetTestLatestResultsPayload) SetProbeDc(v []string)`

SetProbeDc sets ProbeDc field to given value.

### HasProbeDc

`func (o *SyntheticsGetTestLatestResultsPayload) HasProbeDc() bool`

HasProbeDc returns a boolean if a field has been set.

### GetToTs

`func (o *SyntheticsGetTestLatestResultsPayload) GetToTs() float64`

GetToTs returns the ToTs field if non-nil, zero value otherwise.

### GetToTsOk

`func (o *SyntheticsGetTestLatestResultsPayload) GetToTsOk() (*float64, bool)`

GetToTsOk returns a tuple with the ToTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTs

`func (o *SyntheticsGetTestLatestResultsPayload) SetToTs(v float64)`

SetToTs sets ToTs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


