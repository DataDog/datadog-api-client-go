# UsageFargateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | Pointer to [**[]UsageFargateHour**](UsageFargateHour.md) | Array with the number of hourly Fargate tasks recorded for a given organization. | [optional] 

## Methods

### NewUsageFargateResponse

`func NewUsageFargateResponse() *UsageFargateResponse`

NewUsageFargateResponse instantiates a new UsageFargateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageFargateResponseWithDefaults

`func NewUsageFargateResponseWithDefaults() *UsageFargateResponse`

NewUsageFargateResponseWithDefaults instantiates a new UsageFargateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *UsageFargateResponse) GetUsage() []UsageFargateHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageFargateResponse) GetUsageOk() (*[]UsageFargateHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageFargateResponse) SetUsage(v []UsageFargateHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageFargateResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


