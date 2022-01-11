# WidgetFormula

## Properties

| Name                   | Type                                                                       | Description                                                    | Notes      |
| ---------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------- | ---------- |
| **Alias**              | Pointer to **string**                                                      | Expression alias.                                              | [optional] |
| **CellDisplayMode**    | Pointer to [**TableWidgetCellDisplayMode**](TableWidgetCellDisplayMode.md) |                                                                | [optional] |
| **ConditionalFormats** | Pointer to [**[]WidgetConditionalFormat**](WidgetConditionalFormat.md)     | List of conditional formats.                                   | [optional] |
| **Formula**            | **string**                                                                 | String expression built from queries, formulas, and functions. |
| **Limit**              | Pointer to [**WidgetFormulaLimit**](WidgetFormulaLimit.md)                 |                                                                | [optional] |

## Methods

### NewWidgetFormula

`func NewWidgetFormula(formula string) *WidgetFormula`

NewWidgetFormula instantiates a new WidgetFormula object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewWidgetFormulaWithDefaults

`func NewWidgetFormulaWithDefaults() *WidgetFormula`

NewWidgetFormulaWithDefaults instantiates a new WidgetFormula object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

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

### GetCellDisplayMode

`func (o *WidgetFormula) GetCellDisplayMode() TableWidgetCellDisplayMode`

GetCellDisplayMode returns the CellDisplayMode field if non-nil, zero value otherwise.

### GetCellDisplayModeOk

`func (o *WidgetFormula) GetCellDisplayModeOk() (*TableWidgetCellDisplayMode, bool)`

GetCellDisplayModeOk returns a tuple with the CellDisplayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellDisplayMode

`func (o *WidgetFormula) SetCellDisplayMode(v TableWidgetCellDisplayMode)`

SetCellDisplayMode sets CellDisplayMode field to given value.

### HasCellDisplayMode

`func (o *WidgetFormula) HasCellDisplayMode() bool`

HasCellDisplayMode returns a boolean if a field has been set.

### GetConditionalFormats

`func (o *WidgetFormula) GetConditionalFormats() []WidgetConditionalFormat`

GetConditionalFormats returns the ConditionalFormats field if non-nil, zero value otherwise.

### GetConditionalFormatsOk

`func (o *WidgetFormula) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool)`

GetConditionalFormatsOk returns a tuple with the ConditionalFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalFormats

`func (o *WidgetFormula) SetConditionalFormats(v []WidgetConditionalFormat)`

SetConditionalFormats sets ConditionalFormats field to given value.

### HasConditionalFormats

`func (o *WidgetFormula) HasConditionalFormats() bool`

HasConditionalFormats returns a boolean if a field has been set.

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
