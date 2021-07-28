# ChargebackSummaryResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CostAllocation** | Pointer to [**[]ChargebackOrgMonth**](ChargebackOrgMonth.md) | Get cost allocated by sub-org. | [optional] 

## Methods

### NewChargebackSummaryResponse

`func NewChargebackSummaryResponse() *ChargebackSummaryResponse`

NewChargebackSummaryResponse instantiates a new ChargebackSummaryResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewChargebackSummaryResponseWithDefaults

`func NewChargebackSummaryResponseWithDefaults() *ChargebackSummaryResponse`

NewChargebackSummaryResponseWithDefaults instantiates a new ChargebackSummaryResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCostAllocation

`func (o *ChargebackSummaryResponse) GetCostAllocation() []ChargebackOrgMonth`

GetCostAllocation returns the CostAllocation field if non-nil, zero value otherwise.

### GetCostAllocationOk

`func (o *ChargebackSummaryResponse) GetCostAllocationOk() (*[]ChargebackOrgMonth, bool)`

GetCostAllocationOk returns a tuple with the CostAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAllocation

`func (o *ChargebackSummaryResponse) SetCostAllocation(v []ChargebackOrgMonth)`

SetCostAllocation sets CostAllocation field to given value.

### HasCostAllocation

`func (o *ChargebackSummaryResponse) HasCostAllocation() bool`

HasCostAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


