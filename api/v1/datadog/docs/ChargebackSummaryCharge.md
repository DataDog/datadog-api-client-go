# ChargebackSummaryCharge

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**ChargeType** | Pointer to **string** | The subscription type for the usage (&#x60;committed&#x60;, &#x60;on_demand&#x60;, or &#x60;committed_remaining&#x60;). | [optional] 
**Cost** | Pointer to **float64** | Cost of usage for product attributable to sub-org. | [optional] 
**PricePerUnit** | Pointer to **float64** | Price in dollars per unit of usage. | [optional] 
**ProductName** | Pointer to **string** | Name of product. | [optional] 
**Units** | Pointer to **float64** | Number of billed units attributable to sub-org. | [optional] 

## Methods

### NewChargebackSummaryCharge

`func NewChargebackSummaryCharge() *ChargebackSummaryCharge`

NewChargebackSummaryCharge instantiates a new ChargebackSummaryCharge object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewChargebackSummaryChargeWithDefaults

`func NewChargebackSummaryChargeWithDefaults() *ChargebackSummaryCharge`

NewChargebackSummaryChargeWithDefaults instantiates a new ChargebackSummaryCharge object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetChargeType

`func (o *ChargebackSummaryCharge) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *ChargebackSummaryCharge) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *ChargebackSummaryCharge) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *ChargebackSummaryCharge) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetCost

`func (o *ChargebackSummaryCharge) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ChargebackSummaryCharge) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ChargebackSummaryCharge) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ChargebackSummaryCharge) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *ChargebackSummaryCharge) GetPricePerUnit() float64`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *ChargebackSummaryCharge) GetPricePerUnitOk() (*float64, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *ChargebackSummaryCharge) SetPricePerUnit(v float64)`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *ChargebackSummaryCharge) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

### GetProductName

`func (o *ChargebackSummaryCharge) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ChargebackSummaryCharge) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ChargebackSummaryCharge) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ChargebackSummaryCharge) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetUnits

`func (o *ChargebackSummaryCharge) GetUnits() float64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ChargebackSummaryCharge) GetUnitsOk() (*float64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ChargebackSummaryCharge) SetUnits(v float64)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ChargebackSummaryCharge) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


