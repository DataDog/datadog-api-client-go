# UsageSDSResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Usage** | Pointer to [**[]UsageSDSHour**](UsageSDSHour.md) | Get hourly usage for Sensitive Data Scanner. | [optional] 

## Methods

### NewUsageSDSResponse

`func NewUsageSDSResponse() *UsageSDSResponse`

NewUsageSDSResponse instantiates a new UsageSDSResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSDSResponseWithDefaults

`func NewUsageSDSResponseWithDefaults() *UsageSDSResponse`

NewUsageSDSResponseWithDefaults instantiates a new UsageSDSResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetUsage

`func (o *UsageSDSResponse) GetUsage() []UsageSDSHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageSDSResponse) GetUsageOk() (*[]UsageSDSHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageSDSResponse) SetUsage(v []UsageSDSHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageSDSResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


