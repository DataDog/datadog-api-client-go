# UsageSyntheticsAPIResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Usage** | Pointer to [**[]UsageSyntheticsAPIHour**](UsageSyntheticsAPIHour.md) | Get hourly usage for Synthetics API tests. | [optional] 

## Methods

### NewUsageSyntheticsAPIResponse

`func NewUsageSyntheticsAPIResponse() *UsageSyntheticsAPIResponse`

NewUsageSyntheticsAPIResponse instantiates a new UsageSyntheticsAPIResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSyntheticsAPIResponseWithDefaults

`func NewUsageSyntheticsAPIResponseWithDefaults() *UsageSyntheticsAPIResponse`

NewUsageSyntheticsAPIResponseWithDefaults instantiates a new UsageSyntheticsAPIResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetUsage

`func (o *UsageSyntheticsAPIResponse) GetUsage() []UsageSyntheticsAPIHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageSyntheticsAPIResponse) GetUsageOk() (*[]UsageSyntheticsAPIHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageSyntheticsAPIResponse) SetUsage(v []UsageSyntheticsAPIHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageSyntheticsAPIResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


