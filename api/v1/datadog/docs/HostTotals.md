# HostTotals

## Properties

| Name            | Type                 | Description                                                    | Notes      |
| --------------- | -------------------- | -------------------------------------------------------------- | ---------- |
| **TotalActive** | Pointer to **int64** | Total number of active host (UP and ???) reporting to Datadog. | [optional] |
| **TotalUp**     | Pointer to **int64** | Number of host that are UP and reporting to Datadog.           | [optional] |

## Methods

### NewHostTotals

`func NewHostTotals() *HostTotals`

NewHostTotals instantiates a new HostTotals object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHostTotalsWithDefaults

`func NewHostTotalsWithDefaults() *HostTotals`

NewHostTotalsWithDefaults instantiates a new HostTotals object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTotalActive

`func (o *HostTotals) GetTotalActive() int64`

GetTotalActive returns the TotalActive field if non-nil, zero value otherwise.

### GetTotalActiveOk

`func (o *HostTotals) GetTotalActiveOk() (*int64, bool)`

GetTotalActiveOk returns a tuple with the TotalActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActive

`func (o *HostTotals) SetTotalActive(v int64)`

SetTotalActive sets TotalActive field to given value.

### HasTotalActive

`func (o *HostTotals) HasTotalActive() bool`

HasTotalActive returns a boolean if a field has been set.

### GetTotalUp

`func (o *HostTotals) GetTotalUp() int64`

GetTotalUp returns the TotalUp field if non-nil, zero value otherwise.

### GetTotalUpOk

`func (o *HostTotals) GetTotalUpOk() (*int64, bool)`

GetTotalUpOk returns a tuple with the TotalUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUp

`func (o *HostTotals) SetTotalUp(v int64)`

SetTotalUp sets TotalUp field to given value.

### HasTotalUp

`func (o *HostTotals) HasTotalUp() bool`

HasTotalUp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
