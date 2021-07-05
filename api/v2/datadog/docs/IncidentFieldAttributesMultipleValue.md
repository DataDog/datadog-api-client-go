# IncidentFieldAttributesMultipleValue

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Type** | Pointer to [**IncidentFieldAttributesValueType**](IncidentFieldAttributesValueType.md) |  | [optional] [default to INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT]
**Value** | Pointer to **[]string** | The multiple values selected for this field. | [optional] 

## Methods

### NewIncidentFieldAttributesMultipleValue

`func NewIncidentFieldAttributesMultipleValue() *IncidentFieldAttributesMultipleValue`

NewIncidentFieldAttributesMultipleValue instantiates a new IncidentFieldAttributesMultipleValue object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentFieldAttributesMultipleValueWithDefaults

`func NewIncidentFieldAttributesMultipleValueWithDefaults() *IncidentFieldAttributesMultipleValue`

NewIncidentFieldAttributesMultipleValueWithDefaults instantiates a new IncidentFieldAttributesMultipleValue object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetType

`func (o *IncidentFieldAttributesMultipleValue) GetType() IncidentFieldAttributesValueType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentFieldAttributesMultipleValue) GetTypeOk() (*IncidentFieldAttributesValueType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentFieldAttributesMultipleValue) SetType(v IncidentFieldAttributesValueType)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentFieldAttributesMultipleValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *IncidentFieldAttributesMultipleValue) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IncidentFieldAttributesMultipleValue) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IncidentFieldAttributesMultipleValue) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IncidentFieldAttributesMultipleValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


