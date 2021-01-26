# UsageAttributionValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiPercentage** | Pointer to **float64** | The percentage of synthetic API test usage by tag(s). | [optional] 
**ApiUsage** | Pointer to **float64** | The synthetic API test usage by tag(s). | [optional] 
**ApmHostPercentage** | Pointer to **float64** | The percentage of APM host usage by tag(s). | [optional] 
**ApmHostUsage** | Pointer to **float64** | The APM host usage by tag(s). | [optional] 
**BrowserPercentage** | Pointer to **float64** | The percentage of synthetic browser test usage by tag(s). | [optional] 
**BrowserUsage** | Pointer to **float64** | The synthetic browser test usage by tag(s). | [optional] 
**ContainerPercentage** | Pointer to **float64** | The percentage of container usage by tag(s). | [optional] 
**ContainerUsage** | Pointer to **float64** | The container usage by tag(s). | [optional] 
**CustomTimeseriesPercentage** | Pointer to **float64** | The percentage of custom metrics usage by tag(s). | [optional] 
**CustomTimeseriesUsage** | Pointer to **float64** | The custom metrics usage by tag(s). | [optional] 
**InfraHostPercentage** | Pointer to **float64** | The percentage of infrastructure host usage by tag(s). | [optional] 
**InfraHostUsage** | Pointer to **float64** | The infrastructure host usage by tag(s). | [optional] 
**LambdaPercentage** | Pointer to **float64** | The percentage of lambda function usage by tag(s). | [optional] 
**LambdaUsage** | Pointer to **float64** | The lambda function usage by tag(s). | [optional] 
**NpmHostPercentage** | Pointer to **float64** | The percentage of network host usage by tag(s). | [optional] 
**NpmHostUsage** | Pointer to **float64** | The network host usage by tag(s). | [optional] 
**SnmpPercentage** | Pointer to **float64** | The percentage of network device usage by tag(s). | [optional] 
**SnmpUsage** | Pointer to **float64** | The network device usage by tag(s). | [optional] 

## Methods

### NewUsageAttributionValues

`func NewUsageAttributionValues() *UsageAttributionValues`

NewUsageAttributionValues instantiates a new UsageAttributionValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAttributionValuesWithDefaults

`func NewUsageAttributionValuesWithDefaults() *UsageAttributionValues`

NewUsageAttributionValuesWithDefaults instantiates a new UsageAttributionValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiPercentage

`func (o *UsageAttributionValues) GetApiPercentage() float64`

GetApiPercentage returns the ApiPercentage field if non-nil, zero value otherwise.

### GetApiPercentageOk

`func (o *UsageAttributionValues) GetApiPercentageOk() (*float64, bool)`

GetApiPercentageOk returns a tuple with the ApiPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPercentage

`func (o *UsageAttributionValues) SetApiPercentage(v float64)`

SetApiPercentage sets ApiPercentage field to given value.

### HasApiPercentage

`func (o *UsageAttributionValues) HasApiPercentage() bool`

HasApiPercentage returns a boolean if a field has been set.

### GetApiUsage

`func (o *UsageAttributionValues) GetApiUsage() float64`

GetApiUsage returns the ApiUsage field if non-nil, zero value otherwise.

### GetApiUsageOk

`func (o *UsageAttributionValues) GetApiUsageOk() (*float64, bool)`

GetApiUsageOk returns a tuple with the ApiUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsage

`func (o *UsageAttributionValues) SetApiUsage(v float64)`

SetApiUsage sets ApiUsage field to given value.

### HasApiUsage

`func (o *UsageAttributionValues) HasApiUsage() bool`

HasApiUsage returns a boolean if a field has been set.

### GetApmHostPercentage

`func (o *UsageAttributionValues) GetApmHostPercentage() float64`

GetApmHostPercentage returns the ApmHostPercentage field if non-nil, zero value otherwise.

### GetApmHostPercentageOk

`func (o *UsageAttributionValues) GetApmHostPercentageOk() (*float64, bool)`

GetApmHostPercentageOk returns a tuple with the ApmHostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostPercentage

`func (o *UsageAttributionValues) SetApmHostPercentage(v float64)`

SetApmHostPercentage sets ApmHostPercentage field to given value.

### HasApmHostPercentage

`func (o *UsageAttributionValues) HasApmHostPercentage() bool`

HasApmHostPercentage returns a boolean if a field has been set.

### GetApmHostUsage

`func (o *UsageAttributionValues) GetApmHostUsage() float64`

GetApmHostUsage returns the ApmHostUsage field if non-nil, zero value otherwise.

### GetApmHostUsageOk

`func (o *UsageAttributionValues) GetApmHostUsageOk() (*float64, bool)`

GetApmHostUsageOk returns a tuple with the ApmHostUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostUsage

`func (o *UsageAttributionValues) SetApmHostUsage(v float64)`

SetApmHostUsage sets ApmHostUsage field to given value.

### HasApmHostUsage

`func (o *UsageAttributionValues) HasApmHostUsage() bool`

HasApmHostUsage returns a boolean if a field has been set.

### GetBrowserPercentage

`func (o *UsageAttributionValues) GetBrowserPercentage() float64`

GetBrowserPercentage returns the BrowserPercentage field if non-nil, zero value otherwise.

### GetBrowserPercentageOk

`func (o *UsageAttributionValues) GetBrowserPercentageOk() (*float64, bool)`

GetBrowserPercentageOk returns a tuple with the BrowserPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserPercentage

`func (o *UsageAttributionValues) SetBrowserPercentage(v float64)`

SetBrowserPercentage sets BrowserPercentage field to given value.

### HasBrowserPercentage

`func (o *UsageAttributionValues) HasBrowserPercentage() bool`

HasBrowserPercentage returns a boolean if a field has been set.

### GetBrowserUsage

`func (o *UsageAttributionValues) GetBrowserUsage() float64`

GetBrowserUsage returns the BrowserUsage field if non-nil, zero value otherwise.

### GetBrowserUsageOk

`func (o *UsageAttributionValues) GetBrowserUsageOk() (*float64, bool)`

GetBrowserUsageOk returns a tuple with the BrowserUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUsage

`func (o *UsageAttributionValues) SetBrowserUsage(v float64)`

SetBrowserUsage sets BrowserUsage field to given value.

### HasBrowserUsage

`func (o *UsageAttributionValues) HasBrowserUsage() bool`

HasBrowserUsage returns a boolean if a field has been set.

### GetContainerPercentage

`func (o *UsageAttributionValues) GetContainerPercentage() float64`

GetContainerPercentage returns the ContainerPercentage field if non-nil, zero value otherwise.

### GetContainerPercentageOk

`func (o *UsageAttributionValues) GetContainerPercentageOk() (*float64, bool)`

GetContainerPercentageOk returns a tuple with the ContainerPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPercentage

`func (o *UsageAttributionValues) SetContainerPercentage(v float64)`

SetContainerPercentage sets ContainerPercentage field to given value.

### HasContainerPercentage

`func (o *UsageAttributionValues) HasContainerPercentage() bool`

HasContainerPercentage returns a boolean if a field has been set.

### GetContainerUsage

`func (o *UsageAttributionValues) GetContainerUsage() float64`

GetContainerUsage returns the ContainerUsage field if non-nil, zero value otherwise.

### GetContainerUsageOk

`func (o *UsageAttributionValues) GetContainerUsageOk() (*float64, bool)`

GetContainerUsageOk returns a tuple with the ContainerUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerUsage

`func (o *UsageAttributionValues) SetContainerUsage(v float64)`

SetContainerUsage sets ContainerUsage field to given value.

### HasContainerUsage

`func (o *UsageAttributionValues) HasContainerUsage() bool`

HasContainerUsage returns a boolean if a field has been set.

### GetCustomTimeseriesPercentage

`func (o *UsageAttributionValues) GetCustomTimeseriesPercentage() float64`

GetCustomTimeseriesPercentage returns the CustomTimeseriesPercentage field if non-nil, zero value otherwise.

### GetCustomTimeseriesPercentageOk

`func (o *UsageAttributionValues) GetCustomTimeseriesPercentageOk() (*float64, bool)`

GetCustomTimeseriesPercentageOk returns a tuple with the CustomTimeseriesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeseriesPercentage

`func (o *UsageAttributionValues) SetCustomTimeseriesPercentage(v float64)`

SetCustomTimeseriesPercentage sets CustomTimeseriesPercentage field to given value.

### HasCustomTimeseriesPercentage

`func (o *UsageAttributionValues) HasCustomTimeseriesPercentage() bool`

HasCustomTimeseriesPercentage returns a boolean if a field has been set.

### GetCustomTimeseriesUsage

`func (o *UsageAttributionValues) GetCustomTimeseriesUsage() float64`

GetCustomTimeseriesUsage returns the CustomTimeseriesUsage field if non-nil, zero value otherwise.

### GetCustomTimeseriesUsageOk

`func (o *UsageAttributionValues) GetCustomTimeseriesUsageOk() (*float64, bool)`

GetCustomTimeseriesUsageOk returns a tuple with the CustomTimeseriesUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeseriesUsage

`func (o *UsageAttributionValues) SetCustomTimeseriesUsage(v float64)`

SetCustomTimeseriesUsage sets CustomTimeseriesUsage field to given value.

### HasCustomTimeseriesUsage

`func (o *UsageAttributionValues) HasCustomTimeseriesUsage() bool`

HasCustomTimeseriesUsage returns a boolean if a field has been set.

### GetInfraHostPercentage

`func (o *UsageAttributionValues) GetInfraHostPercentage() float64`

GetInfraHostPercentage returns the InfraHostPercentage field if non-nil, zero value otherwise.

### GetInfraHostPercentageOk

`func (o *UsageAttributionValues) GetInfraHostPercentageOk() (*float64, bool)`

GetInfraHostPercentageOk returns a tuple with the InfraHostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostPercentage

`func (o *UsageAttributionValues) SetInfraHostPercentage(v float64)`

SetInfraHostPercentage sets InfraHostPercentage field to given value.

### HasInfraHostPercentage

`func (o *UsageAttributionValues) HasInfraHostPercentage() bool`

HasInfraHostPercentage returns a boolean if a field has been set.

### GetInfraHostUsage

`func (o *UsageAttributionValues) GetInfraHostUsage() float64`

GetInfraHostUsage returns the InfraHostUsage field if non-nil, zero value otherwise.

### GetInfraHostUsageOk

`func (o *UsageAttributionValues) GetInfraHostUsageOk() (*float64, bool)`

GetInfraHostUsageOk returns a tuple with the InfraHostUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostUsage

`func (o *UsageAttributionValues) SetInfraHostUsage(v float64)`

SetInfraHostUsage sets InfraHostUsage field to given value.

### HasInfraHostUsage

`func (o *UsageAttributionValues) HasInfraHostUsage() bool`

HasInfraHostUsage returns a boolean if a field has been set.

### GetLambdaPercentage

`func (o *UsageAttributionValues) GetLambdaPercentage() float64`

GetLambdaPercentage returns the LambdaPercentage field if non-nil, zero value otherwise.

### GetLambdaPercentageOk

`func (o *UsageAttributionValues) GetLambdaPercentageOk() (*float64, bool)`

GetLambdaPercentageOk returns a tuple with the LambdaPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaPercentage

`func (o *UsageAttributionValues) SetLambdaPercentage(v float64)`

SetLambdaPercentage sets LambdaPercentage field to given value.

### HasLambdaPercentage

`func (o *UsageAttributionValues) HasLambdaPercentage() bool`

HasLambdaPercentage returns a boolean if a field has been set.

### GetLambdaUsage

`func (o *UsageAttributionValues) GetLambdaUsage() float64`

GetLambdaUsage returns the LambdaUsage field if non-nil, zero value otherwise.

### GetLambdaUsageOk

`func (o *UsageAttributionValues) GetLambdaUsageOk() (*float64, bool)`

GetLambdaUsageOk returns a tuple with the LambdaUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaUsage

`func (o *UsageAttributionValues) SetLambdaUsage(v float64)`

SetLambdaUsage sets LambdaUsage field to given value.

### HasLambdaUsage

`func (o *UsageAttributionValues) HasLambdaUsage() bool`

HasLambdaUsage returns a boolean if a field has been set.

### GetNpmHostPercentage

`func (o *UsageAttributionValues) GetNpmHostPercentage() float64`

GetNpmHostPercentage returns the NpmHostPercentage field if non-nil, zero value otherwise.

### GetNpmHostPercentageOk

`func (o *UsageAttributionValues) GetNpmHostPercentageOk() (*float64, bool)`

GetNpmHostPercentageOk returns a tuple with the NpmHostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostPercentage

`func (o *UsageAttributionValues) SetNpmHostPercentage(v float64)`

SetNpmHostPercentage sets NpmHostPercentage field to given value.

### HasNpmHostPercentage

`func (o *UsageAttributionValues) HasNpmHostPercentage() bool`

HasNpmHostPercentage returns a boolean if a field has been set.

### GetNpmHostUsage

`func (o *UsageAttributionValues) GetNpmHostUsage() float64`

GetNpmHostUsage returns the NpmHostUsage field if non-nil, zero value otherwise.

### GetNpmHostUsageOk

`func (o *UsageAttributionValues) GetNpmHostUsageOk() (*float64, bool)`

GetNpmHostUsageOk returns a tuple with the NpmHostUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostUsage

`func (o *UsageAttributionValues) SetNpmHostUsage(v float64)`

SetNpmHostUsage sets NpmHostUsage field to given value.

### HasNpmHostUsage

`func (o *UsageAttributionValues) HasNpmHostUsage() bool`

HasNpmHostUsage returns a boolean if a field has been set.

### GetSnmpPercentage

`func (o *UsageAttributionValues) GetSnmpPercentage() float64`

GetSnmpPercentage returns the SnmpPercentage field if non-nil, zero value otherwise.

### GetSnmpPercentageOk

`func (o *UsageAttributionValues) GetSnmpPercentageOk() (*float64, bool)`

GetSnmpPercentageOk returns a tuple with the SnmpPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpPercentage

`func (o *UsageAttributionValues) SetSnmpPercentage(v float64)`

SetSnmpPercentage sets SnmpPercentage field to given value.

### HasSnmpPercentage

`func (o *UsageAttributionValues) HasSnmpPercentage() bool`

HasSnmpPercentage returns a boolean if a field has been set.

### GetSnmpUsage

`func (o *UsageAttributionValues) GetSnmpUsage() float64`

GetSnmpUsage returns the SnmpUsage field if non-nil, zero value otherwise.

### GetSnmpUsageOk

`func (o *UsageAttributionValues) GetSnmpUsageOk() (*float64, bool)`

GetSnmpUsageOk returns a tuple with the SnmpUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpUsage

`func (o *UsageAttributionValues) SetSnmpUsage(v float64)`

SetSnmpUsage sets SnmpUsage field to given value.

### HasSnmpUsage

`func (o *UsageAttributionValues) HasSnmpUsage() bool`

HasSnmpUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


