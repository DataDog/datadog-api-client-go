# UsageCWSResponse

## Properties

| Name      | Type                                             | Description                                   | Notes      |
| --------- | ------------------------------------------------ | --------------------------------------------- | ---------- |
| **Usage** | Pointer to [**[]UsageCWSHour**](UsageCWSHour.md) | Get hourly usage for Cloud Workload Security. | [optional] |

## Methods

### NewUsageCWSResponse

`func NewUsageCWSResponse() *UsageCWSResponse`

NewUsageCWSResponse instantiates a new UsageCWSResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageCWSResponseWithDefaults

`func NewUsageCWSResponseWithDefaults() *UsageCWSResponse`

NewUsageCWSResponseWithDefaults instantiates a new UsageCWSResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetUsage

`func (o *UsageCWSResponse) GetUsage() []UsageCWSHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageCWSResponse) GetUsageOk() (*[]UsageCWSHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageCWSResponse) SetUsage(v []UsageCWSHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageCWSResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
