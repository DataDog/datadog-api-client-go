# ServiceLevelObjectiveQuery

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Denominator** | **string** | A Datadog metric query for total (valid) events. | 
**Numerator** | **string** | A Datadog metric query for good events. | 

## Methods

### NewServiceLevelObjectiveQuery

`func NewServiceLevelObjectiveQuery(denominator string, numerator string, ) *ServiceLevelObjectiveQuery`

NewServiceLevelObjectiveQuery instantiates a new ServiceLevelObjectiveQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLevelObjectiveQueryWithDefaults

`func NewServiceLevelObjectiveQueryWithDefaults() *ServiceLevelObjectiveQuery`

NewServiceLevelObjectiveQueryWithDefaults instantiates a new ServiceLevelObjectiveQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenominator

`func (o *ServiceLevelObjectiveQuery) GetDenominator() string`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *ServiceLevelObjectiveQuery) GetDenominatorOk() (*string, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *ServiceLevelObjectiveQuery) SetDenominator(v string)`

SetDenominator sets Denominator field to given value.


### GetNumerator

`func (o *ServiceLevelObjectiveQuery) GetNumerator() string`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *ServiceLevelObjectiveQuery) GetNumeratorOk() (*string, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *ServiceLevelObjectiveQuery) SetNumerator(v string)`

SetNumerator sets Numerator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


