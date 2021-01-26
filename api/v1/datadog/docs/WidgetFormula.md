# WidgetFormula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Expression alias. | [optional] 
**Formula** | **string** | String expression built from queries, formulas, and functions. | 
**Limit** | Pointer to [**WidgetFormulaLimit**](WidgetFormula_limit.md) |  | [optional] 

## Methods

### NewWidgetFormula

`func NewWidgetFormula(formula string, ) *WidgetFormula`

NewWidgetFormula instantiates a new WidgetFormula object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetFormulaWithDefaults

`func NewWidgetFormulaWithDefaults() *WidgetFormula`

NewWidgetFormulaWithDefaults instantiates a new WidgetFormula object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *WidgetFormula) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *WidgetFormula) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *WidgetFormula) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *WidgetFormula) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetFormula

`func (o *WidgetFormula) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *WidgetFormula) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *WidgetFormula) SetFormula(v string)`

SetFormula sets Formula field to given value.


### GetLimit

`func (o *WidgetFormula) GetLimit() WidgetFormulaLimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WidgetFormula) GetLimitOk() (*WidgetFormulaLimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WidgetFormula) SetLimit(v WidgetFormulaLimit)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *WidgetFormula) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


