# IncidentFieldAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**IncidentFieldAttributesValueType**](IncidentFieldAttributesValueType.md) |  | [default to INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT]
**Value** | Pointer to **[]string** | The multiple values selected for this field. | [optional] 

## Methods

### NewIncidentFieldAttributes

`func NewIncidentFieldAttributes(type_ IncidentFieldAttributesValueType, ) *IncidentFieldAttributes`

NewIncidentFieldAttributes instantiates a new IncidentFieldAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentFieldAttributesWithDefaults

`func NewIncidentFieldAttributesWithDefaults() *IncidentFieldAttributes`

NewIncidentFieldAttributesWithDefaults instantiates a new IncidentFieldAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IncidentFieldAttributes) GetType() IncidentFieldAttributesValueType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentFieldAttributes) GetTypeOk() (*IncidentFieldAttributesValueType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentFieldAttributes) SetType(v IncidentFieldAttributesValueType)`

SetType sets Type field to given value.


### GetValue

`func (o *IncidentFieldAttributes) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IncidentFieldAttributes) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IncidentFieldAttributes) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IncidentFieldAttributes) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


