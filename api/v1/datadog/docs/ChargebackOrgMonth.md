# ChargebackOrgMonth

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Charges** | Pointer to [**[]ChargebackSummaryCharge**](ChargebackSummaryCharge.md) | Charges by product and charge type. | [optional] 
**Date** | Pointer to **time.Time** | The month for the usage. | [optional] 
**OrgId** | Pointer to **int64** | ID of the sub-org. | [optional] 
**TotalCost** | Pointer to **float64** | Total cost for all products for the sub-org within the month. | [optional] 

## Methods

### NewChargebackOrgMonth

`func NewChargebackOrgMonth() *ChargebackOrgMonth`

NewChargebackOrgMonth instantiates a new ChargebackOrgMonth object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewChargebackOrgMonthWithDefaults

`func NewChargebackOrgMonthWithDefaults() *ChargebackOrgMonth`

NewChargebackOrgMonthWithDefaults instantiates a new ChargebackOrgMonth object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCharges

`func (o *ChargebackOrgMonth) GetCharges() []ChargebackSummaryCharge`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *ChargebackOrgMonth) GetChargesOk() (*[]ChargebackSummaryCharge, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *ChargebackOrgMonth) SetCharges(v []ChargebackSummaryCharge)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *ChargebackOrgMonth) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetDate

`func (o *ChargebackOrgMonth) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ChargebackOrgMonth) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ChargebackOrgMonth) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ChargebackOrgMonth) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetOrgId

`func (o *ChargebackOrgMonth) GetOrgId() int64`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ChargebackOrgMonth) GetOrgIdOk() (*int64, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ChargebackOrgMonth) SetOrgId(v int64)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ChargebackOrgMonth) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTotalCost

`func (o *ChargebackOrgMonth) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *ChargebackOrgMonth) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *ChargebackOrgMonth) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *ChargebackOrgMonth) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


