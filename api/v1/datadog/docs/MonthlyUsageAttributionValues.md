# MonthlyUsageAttributionValues

## Properties

| Name                            | Type                   | Description                                               | Notes      |
| ------------------------------- | ---------------------- | --------------------------------------------------------- | ---------- |
| **ApiPercentage**               | Pointer to **float64** | The percentage of synthetic API test usage by tag(s).     | [optional] |
| **ApiUsage**                    | Pointer to **float64** | The synthetic API test usage by tag(s).                   | [optional] |
| **ApmHostPercentage**           | Pointer to **float64** | The percentage of APM host usage by tag(s).               | [optional] |
| **ApmHostUsage**                | Pointer to **float64** | The APM host usage by tag(s).                             | [optional] |
| **BrowserPercentage**           | Pointer to **float64** | The percentage of synthetic browser test usage by tag(s). | [optional] |
| **BrowserUsage**                | Pointer to **float64** | The synthetic browser test usage by tag(s).               | [optional] |
| **ContainerPercentage**         | Pointer to **float64** | The percentage of container usage by tag(s).              | [optional] |
| **ContainerUsage**              | Pointer to **float64** | The container usage by tag(s).                            | [optional] |
| **CustomTimeseriesPercentage**  | Pointer to **float64** | The percentage of custom metrics usage by tag(s).         | [optional] |
| **CustomTimeseriesUsage**       | Pointer to **float64** | The custom metrics usage by tag(s).                       | [optional] |
| **FargatePercentage**           | Pointer to **float64** | The percentage of Fargate usage by tags.                  | [optional] |
| **FargateUsage**                | Pointer to **float64** | The Fargate usage by tags.                                | [optional] |
| **FunctionsPercentage**         | Pointer to **float64** | The percentage of Lambda function usage by tag(s).        | [optional] |
| **FunctionsUsage**              | Pointer to **float64** | The Lambda function usage by tag(s).                      | [optional] |
| **IndexedLogsPercentage**       | Pointer to **float64** | The percentage of indexed logs usage by tags.             | [optional] |
| **IndexedLogsUsage**            | Pointer to **float64** | The indexed logs usage by tags.                           | [optional] |
| **InfraHostPercentage**         | Pointer to **float64** | The percentage of infrastructure host usage by tag(s).    | [optional] |
| **InfraHostUsage**              | Pointer to **float64** | The infrastructure host usage by tag(s).                  | [optional] |
| **InvocationsPercentage**       | Pointer to **float64** | The percentage of Lambda invocation usage by tag(s).      | [optional] |
| **InvocationsUsage**            | Pointer to **float64** | The Lambda invocation usage by tag(s).                    | [optional] |
| **NpmHostPercentage**           | Pointer to **float64** | The percentage of network host usage by tag(s).           | [optional] |
| **NpmHostUsage**                | Pointer to **float64** | The network host usage by tag(s).                         | [optional] |
| **ProfiledContainerPercentage** | Pointer to **float64** | The percentage of profiled container usage by tag(s).     | [optional] |
| **ProfiledContainerUsage**      | Pointer to **float64** | The profiled container usage by tag(s).                   | [optional] |
| **ProfiledHostPercentage**      | Pointer to **float64** | The percentage of profiled hosts usage by tag(s).         | [optional] |
| **ProfiledHostUsage**           | Pointer to **float64** | The profiled hosts usage by tag(s).                       | [optional] |
| **SnmpPercentage**              | Pointer to **float64** | The percentage of network device usage by tag(s).         | [optional] |
| **SnmpUsage**                   | Pointer to **float64** | The network device usage by tag(s).                       | [optional] |

## Methods

### NewMonthlyUsageAttributionValues

`func NewMonthlyUsageAttributionValues() *MonthlyUsageAttributionValues`

NewMonthlyUsageAttributionValues instantiates a new MonthlyUsageAttributionValues object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonthlyUsageAttributionValuesWithDefaults

`func NewMonthlyUsageAttributionValuesWithDefaults() *MonthlyUsageAttributionValues`

NewMonthlyUsageAttributionValuesWithDefaults instantiates a new MonthlyUsageAttributionValues object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetApiPercentage

`func (o *MonthlyUsageAttributionValues) GetApiPercentage() float64`

GetApiPercentage returns the ApiPercentage field if non-nil, zero value otherwise.

### GetApiPercentageOk

`func (o *MonthlyUsageAttributionValues) GetApiPercentageOk() (*float64, bool)`

GetApiPercentageOk returns a tuple with the ApiPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPercentage

`func (o *MonthlyUsageAttributionValues) SetApiPercentage(v float64)`

SetApiPercentage sets ApiPercentage field to given value.

### HasApiPercentage

`func (o *MonthlyUsageAttributionValues) HasApiPercentage() bool`

HasApiPercentage returns a boolean if a field has been set.

### GetApiUsage

`func (o *MonthlyUsageAttributionValues) GetApiUsage() float64`

GetApiUsage returns the ApiUsage field if non-nil, zero value otherwise.

### GetApiUsageOk

`func (o *MonthlyUsageAttributionValues) GetApiUsageOk() (*float64, bool)`

GetApiUsageOk returns a tuple with the ApiUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsage

`func (o *MonthlyUsageAttributionValues) SetApiUsage(v float64)`

SetApiUsage sets ApiUsage field to given value.

### HasApiUsage

`func (o *MonthlyUsageAttributionValues) HasApiUsage() bool`

HasApiUsage returns a boolean if a field has been set.

### GetApmHostPercentage

`func (o *MonthlyUsageAttributionValues) GetApmHostPercentage() float64`

GetApmHostPercentage returns the ApmHostPercentage field if non-nil, zero value otherwise.

### GetApmHostPercentageOk

`func (o *MonthlyUsageAttributionValues) GetApmHostPercentageOk() (*float64, bool)`

GetApmHostPercentageOk returns a tuple with the ApmHostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostPercentage

`func (o *MonthlyUsageAttributionValues) SetApmHostPercentage(v float64)`

SetApmHostPercentage sets ApmHostPercentage field to given value.

### HasApmHostPercentage

`func (o *MonthlyUsageAttributionValues) HasApmHostPercentage() bool`

HasApmHostPercentage returns a boolean if a field has been set.

### GetApmHostUsage

`func (o *MonthlyUsageAttributionValues) GetApmHostUsage() float64`

GetApmHostUsage returns the ApmHostUsage field if non-nil, zero value otherwise.

### GetApmHostUsageOk

`func (o *MonthlyUsageAttributionValues) GetApmHostUsageOk() (*float64, bool)`

GetApmHostUsageOk returns a tuple with the ApmHostUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostUsage

`func (o *MonthlyUsageAttributionValues) SetApmHostUsage(v float64)`

SetApmHostUsage sets ApmHostUsage field to given value.

### HasApmHostUsage

`func (o *MonthlyUsageAttributionValues) HasApmHostUsage() bool`

HasApmHostUsage returns a boolean if a field has been set.

### GetBrowserPercentage

`func (o *MonthlyUsageAttributionValues) GetBrowserPercentage() float64`

GetBrowserPercentage returns the BrowserPercentage field if non-nil, zero value otherwise.

### GetBrowserPercentageOk

`func (o *MonthlyUsageAttributionValues) GetBrowserPercentageOk() (*float64, bool)`

GetBrowserPercentageOk returns a tuple with the BrowserPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserPercentage

`func (o *MonthlyUsageAttributionValues) SetBrowserPercentage(v float64)`

SetBrowserPercentage sets BrowserPercentage field to given value.

### HasBrowserPercentage

`func (o *MonthlyUsageAttributionValues) HasBrowserPercentage() bool`

HasBrowserPercentage returns a boolean if a field has been set.

### GetBrowserUsage

`func (o *MonthlyUsageAttributionValues) GetBrowserUsage() float64`

GetBrowserUsage returns the BrowserUsage field if non-nil, zero value otherwise.

### GetBrowserUsageOk

`func (o *MonthlyUsageAttributionValues) GetBrowserUsageOk() (*float64, bool)`

GetBrowserUsageOk returns a tuple with the BrowserUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUsage

`func (o *MonthlyUsageAttributionValues) SetBrowserUsage(v float64)`

SetBrowserUsage sets BrowserUsage field to given value.

### HasBrowserUsage

`func (o *MonthlyUsageAttributionValues) HasBrowserUsage() bool`

HasBrowserUsage returns a boolean if a field has been set.

### GetContainerPercentage

`func (o *MonthlyUsageAttributionValues) GetContainerPercentage() float64`

GetContainerPercentage returns the ContainerPercentage field if non-nil, zero value otherwise.

### GetContainerPercentageOk

`func (o *MonthlyUsageAttributionValues) GetContainerPercentageOk() (*float64, bool)`

GetContainerPercentageOk returns a tuple with the ContainerPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPercentage

`func (o *MonthlyUsageAttributionValues) SetContainerPercentage(v float64)`

SetContainerPercentage sets ContainerPercentage field to given value.

### HasContainerPercentage

`func (o *MonthlyUsageAttributionValues) HasContainerPercentage() bool`

HasContainerPercentage returns a boolean if a field has been set.

### GetContainerUsage

`func (o *MonthlyUsageAttributionValues) GetContainerUsage() float64`

GetContainerUsage returns the ContainerUsage field if non-nil, zero value otherwise.

### GetContainerUsageOk

`func (o *MonthlyUsageAttributionValues) GetContainerUsageOk() (*float64, bool)`

GetContainerUsageOk returns a tuple with the ContainerUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerUsage

`func (o *MonthlyUsageAttributionValues) SetContainerUsage(v float64)`

SetContainerUsage sets ContainerUsage field to given value.

### HasContainerUsage

`func (o *MonthlyUsageAttributionValues) HasContainerUsage() bool`

HasContainerUsage returns a boolean if a field has been set.

### GetCustomTimeseriesPercentage

`func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesPercentage() float64`

GetCustomTimeseriesPercentage returns the CustomTimeseriesPercentage field if non-nil, zero value otherwise.

### GetCustomTimeseriesPercentageOk

`func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesPercentageOk() (*float64, bool)`

GetCustomTimeseriesPercentageOk returns a tuple with the CustomTimeseriesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeseriesPercentage

`func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesPercentage(v float64)`

SetCustomTimeseriesPercentage sets CustomTimeseriesPercentage field to given value.

### HasCustomTimeseriesPercentage

`func (o *MonthlyUsageAttributionValues) HasCustomTimeseriesPercentage() bool`

HasCustomTimeseriesPercentage returns a boolean if a field has been set.

### GetCustomTimeseriesUsage

`func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesUsage() float64`

GetCustomTimeseriesUsage returns the CustomTimeseriesUsage field if non-nil, zero value otherwise.

### GetCustomTimeseriesUsageOk

`func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesUsageOk() (*float64, bool)`

GetCustomTimeseriesUsageOk returns a tuple with the CustomTimeseriesUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeseriesUsage

`func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesUsage(v float64)`

SetCustomTimeseriesUsage sets CustomTimeseriesUsage field to given value.

### HasCustomTimeseriesUsage

`func (o *MonthlyUsageAttributionValues) HasCustomTimeseriesUsage() bool`

HasCustomTimeseriesUsage returns a boolean if a field has been set.

### GetFargatePercentage

`func (o *MonthlyUsageAttributionValues) GetFargatePercentage() float64`

GetFargatePercentage returns the FargatePercentage field if non-nil, zero value otherwise.

### GetFargatePercentageOk

`func (o *MonthlyUsageAttributionValues) GetFargatePercentageOk() (*float64, bool)`

GetFargatePercentageOk returns a tuple with the FargatePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargatePercentage

`func (o *MonthlyUsageAttributionValues) SetFargatePercentage(v float64)`

SetFargatePercentage sets FargatePercentage field to given value.

### HasFargatePercentage

`func (o *MonthlyUsageAttributionValues) HasFargatePercentage() bool`

HasFargatePercentage returns a boolean if a field has been set.

### GetFargateUsage

`func (o *MonthlyUsageAttributionValues) GetFargateUsage() float64`

GetFargateUsage returns the FargateUsage field if non-nil, zero value otherwise.

### GetFargateUsageOk

`func (o *MonthlyUsageAttributionValues) GetFargateUsageOk() (*float64, bool)`

GetFargateUsageOk returns a tuple with the FargateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateUsage

`func (o *MonthlyUsageAttributionValues) SetFargateUsage(v float64)`

SetFargateUsage sets FargateUsage field to given value.

### HasFargateUsage

`func (o *MonthlyUsageAttributionValues) HasFargateUsage() bool`

HasFargateUsage returns a boolean if a field has been set.

### GetFunctionsPercentage

`func (o *MonthlyUsageAttributionValues) GetFunctionsPercentage() float64`

GetFunctionsPercentage returns the FunctionsPercentage field if non-nil, zero value otherwise.

### GetFunctionsPercentageOk

`func (o *MonthlyUsageAttributionValues) GetFunctionsPercentageOk() (*float64, bool)`

GetFunctionsPercentageOk returns a tuple with the FunctionsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsPercentage

`func (o *MonthlyUsageAttributionValues) SetFunctionsPercentage(v float64)`

SetFunctionsPercentage sets FunctionsPercentage field to given value.

### HasFunctionsPercentage

`func (o *MonthlyUsageAttributionValues) HasFunctionsPercentage() bool`

HasFunctionsPercentage returns a boolean if a field has been set.

### GetFunctionsUsage

`func (o *MonthlyUsageAttributionValues) GetFunctionsUsage() float64`

GetFunctionsUsage returns the FunctionsUsage field if non-nil, zero value otherwise.

### GetFunctionsUsageOk

`func (o *MonthlyUsageAttributionValues) GetFunctionsUsageOk() (*float64, bool)`

GetFunctionsUsageOk returns a tuple with the FunctionsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsUsage

`func (o *MonthlyUsageAttributionValues) SetFunctionsUsage(v float64)`

SetFunctionsUsage sets FunctionsUsage field to given value.

### HasFunctionsUsage

`func (o *MonthlyUsageAttributionValues) HasFunctionsUsage() bool`

HasFunctionsUsage returns a boolean if a field has been set.

### GetIndexedLogsPercentage

`func (o *MonthlyUsageAttributionValues) GetIndexedLogsPercentage() float64`

GetIndexedLogsPercentage returns the IndexedLogsPercentage field if non-nil, zero value otherwise.

### GetIndexedLogsPercentageOk

`func (o *MonthlyUsageAttributionValues) GetIndexedLogsPercentageOk() (*float64, bool)`

GetIndexedLogsPercentageOk returns a tuple with the IndexedLogsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedLogsPercentage

`func (o *MonthlyUsageAttributionValues) SetIndexedLogsPercentage(v float64)`

SetIndexedLogsPercentage sets IndexedLogsPercentage field to given value.

### HasIndexedLogsPercentage

`func (o *MonthlyUsageAttributionValues) HasIndexedLogsPercentage() bool`

HasIndexedLogsPercentage returns a boolean if a field has been set.

### GetIndexedLogsUsage

`func (o *MonthlyUsageAttributionValues) GetIndexedLogsUsage() float64`

GetIndexedLogsUsage returns the IndexedLogsUsage field if non-nil, zero value otherwise.

### GetIndexedLogsUsageOk

`func (o *MonthlyUsageAttributionValues) GetIndexedLogsUsageOk() (*float64, bool)`

GetIndexedLogsUsageOk returns a tuple with the IndexedLogsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedLogsUsage

`func (o *MonthlyUsageAttributionValues) SetIndexedLogsUsage(v float64)`

SetIndexedLogsUsage sets IndexedLogsUsage field to given value.

### HasIndexedLogsUsage

`func (o *MonthlyUsageAttributionValues) HasIndexedLogsUsage() bool`

HasIndexedLogsUsage returns a boolean if a field has been set.

### GetInfraHostPercentage

`func (o *MonthlyUsageAttributionValues) GetInfraHostPercentage() float64`

GetInfraHostPercentage returns the InfraHostPercentage field if non-nil, zero value otherwise.

### GetInfraHostPercentageOk

`func (o *MonthlyUsageAttributionValues) GetInfraHostPercentageOk() (*float64, bool)`

GetInfraHostPercentageOk returns a tuple with the InfraHostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostPercentage

`func (o *MonthlyUsageAttributionValues) SetInfraHostPercentage(v float64)`

SetInfraHostPercentage sets InfraHostPercentage field to given value.

### HasInfraHostPercentage

`func (o *MonthlyUsageAttributionValues) HasInfraHostPercentage() bool`

HasInfraHostPercentage returns a boolean if a field has been set.

### GetInfraHostUsage

`func (o *MonthlyUsageAttributionValues) GetInfraHostUsage() float64`

GetInfraHostUsage returns the InfraHostUsage field if non-nil, zero value otherwise.

### GetInfraHostUsageOk

`func (o *MonthlyUsageAttributionValues) GetInfraHostUsageOk() (*float64, bool)`

GetInfraHostUsageOk returns a tuple with the InfraHostUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostUsage

`func (o *MonthlyUsageAttributionValues) SetInfraHostUsage(v float64)`

SetInfraHostUsage sets InfraHostUsage field to given value.

### HasInfraHostUsage

`func (o *MonthlyUsageAttributionValues) HasInfraHostUsage() bool`

HasInfraHostUsage returns a boolean if a field has been set.

### GetInvocationsPercentage

`func (o *MonthlyUsageAttributionValues) GetInvocationsPercentage() float64`

GetInvocationsPercentage returns the InvocationsPercentage field if non-nil, zero value otherwise.

### GetInvocationsPercentageOk

`func (o *MonthlyUsageAttributionValues) GetInvocationsPercentageOk() (*float64, bool)`

GetInvocationsPercentageOk returns a tuple with the InvocationsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationsPercentage

`func (o *MonthlyUsageAttributionValues) SetInvocationsPercentage(v float64)`

SetInvocationsPercentage sets InvocationsPercentage field to given value.

### HasInvocationsPercentage

`func (o *MonthlyUsageAttributionValues) HasInvocationsPercentage() bool`

HasInvocationsPercentage returns a boolean if a field has been set.

### GetInvocationsUsage

`func (o *MonthlyUsageAttributionValues) GetInvocationsUsage() float64`

GetInvocationsUsage returns the InvocationsUsage field if non-nil, zero value otherwise.

### GetInvocationsUsageOk

`func (o *MonthlyUsageAttributionValues) GetInvocationsUsageOk() (*float64, bool)`

GetInvocationsUsageOk returns a tuple with the InvocationsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationsUsage

`func (o *MonthlyUsageAttributionValues) SetInvocationsUsage(v float64)`

SetInvocationsUsage sets InvocationsUsage field to given value.

### HasInvocationsUsage

`func (o *MonthlyUsageAttributionValues) HasInvocationsUsage() bool`

HasInvocationsUsage returns a boolean if a field has been set.

### GetNpmHostPercentage

`func (o *MonthlyUsageAttributionValues) GetNpmHostPercentage() float64`

GetNpmHostPercentage returns the NpmHostPercentage field if non-nil, zero value otherwise.

### GetNpmHostPercentageOk

`func (o *MonthlyUsageAttributionValues) GetNpmHostPercentageOk() (*float64, bool)`

GetNpmHostPercentageOk returns a tuple with the NpmHostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostPercentage

`func (o *MonthlyUsageAttributionValues) SetNpmHostPercentage(v float64)`

SetNpmHostPercentage sets NpmHostPercentage field to given value.

### HasNpmHostPercentage

`func (o *MonthlyUsageAttributionValues) HasNpmHostPercentage() bool`

HasNpmHostPercentage returns a boolean if a field has been set.

### GetNpmHostUsage

`func (o *MonthlyUsageAttributionValues) GetNpmHostUsage() float64`

GetNpmHostUsage returns the NpmHostUsage field if non-nil, zero value otherwise.

### GetNpmHostUsageOk

`func (o *MonthlyUsageAttributionValues) GetNpmHostUsageOk() (*float64, bool)`

GetNpmHostUsageOk returns a tuple with the NpmHostUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostUsage

`func (o *MonthlyUsageAttributionValues) SetNpmHostUsage(v float64)`

SetNpmHostUsage sets NpmHostUsage field to given value.

### HasNpmHostUsage

`func (o *MonthlyUsageAttributionValues) HasNpmHostUsage() bool`

HasNpmHostUsage returns a boolean if a field has been set.

### GetProfiledContainerPercentage

`func (o *MonthlyUsageAttributionValues) GetProfiledContainerPercentage() float64`

GetProfiledContainerPercentage returns the ProfiledContainerPercentage field if non-nil, zero value otherwise.

### GetProfiledContainerPercentageOk

`func (o *MonthlyUsageAttributionValues) GetProfiledContainerPercentageOk() (*float64, bool)`

GetProfiledContainerPercentageOk returns a tuple with the ProfiledContainerPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiledContainerPercentage

`func (o *MonthlyUsageAttributionValues) SetProfiledContainerPercentage(v float64)`

SetProfiledContainerPercentage sets ProfiledContainerPercentage field to given value.

### HasProfiledContainerPercentage

`func (o *MonthlyUsageAttributionValues) HasProfiledContainerPercentage() bool`

HasProfiledContainerPercentage returns a boolean if a field has been set.

### GetProfiledContainerUsage

`func (o *MonthlyUsageAttributionValues) GetProfiledContainerUsage() float64`

GetProfiledContainerUsage returns the ProfiledContainerUsage field if non-nil, zero value otherwise.

### GetProfiledContainerUsageOk

`func (o *MonthlyUsageAttributionValues) GetProfiledContainerUsageOk() (*float64, bool)`

GetProfiledContainerUsageOk returns a tuple with the ProfiledContainerUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiledContainerUsage

`func (o *MonthlyUsageAttributionValues) SetProfiledContainerUsage(v float64)`

SetProfiledContainerUsage sets ProfiledContainerUsage field to given value.

### HasProfiledContainerUsage

`func (o *MonthlyUsageAttributionValues) HasProfiledContainerUsage() bool`

HasProfiledContainerUsage returns a boolean if a field has been set.

### GetProfiledHostPercentage

`func (o *MonthlyUsageAttributionValues) GetProfiledHostPercentage() float64`

GetProfiledHostPercentage returns the ProfiledHostPercentage field if non-nil, zero value otherwise.

### GetProfiledHostPercentageOk

`func (o *MonthlyUsageAttributionValues) GetProfiledHostPercentageOk() (*float64, bool)`

GetProfiledHostPercentageOk returns a tuple with the ProfiledHostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiledHostPercentage

`func (o *MonthlyUsageAttributionValues) SetProfiledHostPercentage(v float64)`

SetProfiledHostPercentage sets ProfiledHostPercentage field to given value.

### HasProfiledHostPercentage

`func (o *MonthlyUsageAttributionValues) HasProfiledHostPercentage() bool`

HasProfiledHostPercentage returns a boolean if a field has been set.

### GetProfiledHostUsage

`func (o *MonthlyUsageAttributionValues) GetProfiledHostUsage() float64`

GetProfiledHostUsage returns the ProfiledHostUsage field if non-nil, zero value otherwise.

### GetProfiledHostUsageOk

`func (o *MonthlyUsageAttributionValues) GetProfiledHostUsageOk() (*float64, bool)`

GetProfiledHostUsageOk returns a tuple with the ProfiledHostUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiledHostUsage

`func (o *MonthlyUsageAttributionValues) SetProfiledHostUsage(v float64)`

SetProfiledHostUsage sets ProfiledHostUsage field to given value.

### HasProfiledHostUsage

`func (o *MonthlyUsageAttributionValues) HasProfiledHostUsage() bool`

HasProfiledHostUsage returns a boolean if a field has been set.

### GetSnmpPercentage

`func (o *MonthlyUsageAttributionValues) GetSnmpPercentage() float64`

GetSnmpPercentage returns the SnmpPercentage field if non-nil, zero value otherwise.

### GetSnmpPercentageOk

`func (o *MonthlyUsageAttributionValues) GetSnmpPercentageOk() (*float64, bool)`

GetSnmpPercentageOk returns a tuple with the SnmpPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpPercentage

`func (o *MonthlyUsageAttributionValues) SetSnmpPercentage(v float64)`

SetSnmpPercentage sets SnmpPercentage field to given value.

### HasSnmpPercentage

`func (o *MonthlyUsageAttributionValues) HasSnmpPercentage() bool`

HasSnmpPercentage returns a boolean if a field has been set.

### GetSnmpUsage

`func (o *MonthlyUsageAttributionValues) GetSnmpUsage() float64`

GetSnmpUsage returns the SnmpUsage field if non-nil, zero value otherwise.

### GetSnmpUsageOk

`func (o *MonthlyUsageAttributionValues) GetSnmpUsageOk() (*float64, bool)`

GetSnmpUsageOk returns a tuple with the SnmpUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpUsage

`func (o *MonthlyUsageAttributionValues) SetSnmpUsage(v float64)`

SetSnmpUsage sets SnmpUsage field to given value.

### HasSnmpUsage

`func (o *MonthlyUsageAttributionValues) HasSnmpUsage() bool`

HasSnmpUsage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
