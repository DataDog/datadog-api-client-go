# UsageOnlineArchiveResponse

## Properties

| Name      | Type                                                                 | Description                               | Notes      |
| --------- | -------------------------------------------------------------------- | ----------------------------------------- | ---------- |
| **Usage** | Pointer to [**[]UsageOnlineArchiveHour**](UsageOnlineArchiveHour.md) | Response containing Online Archive usage. | [optional] |

## Methods

### NewUsageOnlineArchiveResponse

`func NewUsageOnlineArchiveResponse() *UsageOnlineArchiveResponse`

NewUsageOnlineArchiveResponse instantiates a new UsageOnlineArchiveResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageOnlineArchiveResponseWithDefaults

`func NewUsageOnlineArchiveResponseWithDefaults() *UsageOnlineArchiveResponse`

NewUsageOnlineArchiveResponseWithDefaults instantiates a new UsageOnlineArchiveResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetUsage

`func (o *UsageOnlineArchiveResponse) GetUsage() []UsageOnlineArchiveHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageOnlineArchiveResponse) GetUsageOk() (*[]UsageOnlineArchiveHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageOnlineArchiveResponse) SetUsage(v []UsageOnlineArchiveHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageOnlineArchiveResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
