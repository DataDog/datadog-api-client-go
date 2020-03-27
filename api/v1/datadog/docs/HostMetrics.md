# HostMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **float64** | The percent of CPU used (everything but idle). | [optional] 
**Iowait** | Pointer to **float64** | The percent of CPU spent waiting on the IO (not reported for all platforms). | [optional] 
**Load** | Pointer to **float64** | The system load over the last 15 minutes. | [optional] 

## Methods

### NewHostMetrics

`func NewHostMetrics() *HostMetrics`

NewHostMetrics instantiates a new HostMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMetricsWithDefaults

`func NewHostMetricsWithDefaults() *HostMetrics`

NewHostMetricsWithDefaults instantiates a new HostMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *HostMetrics) GetCpu() float64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *HostMetrics) GetCpuOk() (*float64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *HostMetrics) SetCpu(v float64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *HostMetrics) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetIowait

`func (o *HostMetrics) GetIowait() float64`

GetIowait returns the Iowait field if non-nil, zero value otherwise.

### GetIowaitOk

`func (o *HostMetrics) GetIowaitOk() (*float64, bool)`

GetIowaitOk returns a tuple with the Iowait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIowait

`func (o *HostMetrics) SetIowait(v float64)`

SetIowait sets Iowait field to given value.

### HasIowait

`func (o *HostMetrics) HasIowait() bool`

HasIowait returns a boolean if a field has been set.

### GetLoad

`func (o *HostMetrics) GetLoad() float64`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *HostMetrics) GetLoadOk() (*float64, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *HostMetrics) SetLoad(v float64)`

SetLoad sets Load field to given value.

### HasLoad

`func (o *HostMetrics) HasLoad() bool`

HasLoad returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


