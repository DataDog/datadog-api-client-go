# IncidentFieldAttributesSingleValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**IncidentFieldAttributesSingleValueType**](IncidentFieldAttributesSingleValueType.md) |  | [optional] [default to INCIDENTFIELDATTRIBUTESSINGLEVALUETYPE_DROPDOWN]
**Value** | Pointer to **string** | The single value selected for this field. | [optional] 

## Methods

### NewIncidentFieldAttributesSingleValue

`func NewIncidentFieldAttributesSingleValue() *IncidentFieldAttributesSingleValue`

NewIncidentFieldAttributesSingleValue instantiates a new IncidentFieldAttributesSingleValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentFieldAttributesSingleValueWithDefaults

`func NewIncidentFieldAttributesSingleValueWithDefaults() *IncidentFieldAttributesSingleValue`

NewIncidentFieldAttributesSingleValueWithDefaults instantiates a new IncidentFieldAttributesSingleValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IncidentFieldAttributesSingleValue) GetType() IncidentFieldAttributesSingleValueType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentFieldAttributesSingleValue) GetTypeOk() (*IncidentFieldAttributesSingleValueType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentFieldAttributesSingleValue) SetType(v IncidentFieldAttributesSingleValueType)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentFieldAttributesSingleValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *IncidentFieldAttributesSingleValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IncidentFieldAttributesSingleValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IncidentFieldAttributesSingleValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IncidentFieldAttributesSingleValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


