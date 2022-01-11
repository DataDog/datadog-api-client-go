# UsageAttributionValues

## Properties

| Name                               | Type                   | Description                                                                                                                    | Notes      |
| ---------------------------------- | ---------------------- | ------------------------------------------------------------------------------------------------------------------------------ | ---------- |
| **ApiPercentage**                  | Pointer to **float64** | The percentage of synthetic API test usage by tag(s).                                                                          | [optional] |
| **ApiUsage**                       | Pointer to **float64** | The synthetic API test usage by tag(s).                                                                                        | [optional] |
| **ApmHostPercentage**              | Pointer to **float64** | The percentage of APM host usage by tag(s).                                                                                    | [optional] |
| **ApmHostUsage**                   | Pointer to **float64** | The APM host usage by tag(s).                                                                                                  | [optional] |
| **BrowserPercentage**              | Pointer to **float64** | The percentage of synthetic browser test usage by tag(s).                                                                      | [optional] |
| **BrowserUsage**                   | Pointer to **float64** | The synthetic browser test usage by tag(s).                                                                                    | [optional] |
| **ContainerPercentage**            | Pointer to **float64** | The percentage of container usage by tag(s).                                                                                   | [optional] |
| **ContainerUsage**                 | Pointer to **float64** | The container usage by tag(s).                                                                                                 | [optional] |
| **CspmContainerPercentage**        | Pointer to **float64** | The percentage of Cloud Security Posture Management container usage by tag(s)                                                  | [optional] |
| **CspmContainerUsage**             | Pointer to **float64** | The Cloud Security Posture Management container usage by tag(s)                                                                | [optional] |
| **CspmHostPercentage**             | Pointer to **float64** | The percentage of Cloud Security Posture Management host usage by tag(s)                                                       | [optional] |
| **CspmHostUsage**                  | Pointer to **float64** | The Cloud Security Posture Management host usage by tag(s)                                                                     | [optional] |
| **CustomTimeseriesPercentage**     | Pointer to **float64** | The percentage of custom metrics usage by tag(s).                                                                              | [optional] |
| **CustomTimeseriesUsage**          | Pointer to **float64** | The custom metrics usage by tag(s).                                                                                            | [optional] |
| **CwsContainerPercentage**         | Pointer to **float64** | The percentage of Cloud Workload Security container usage by tag(s)                                                            | [optional] |
| **CwsContainerUsage**              | Pointer to **float64** | The Cloud Workload Security container usage by tag(s)                                                                          | [optional] |
| **CwsHostPercentage**              | Pointer to **float64** | The percentage of Cloud Workload Security host usage by tag(s)                                                                 | [optional] |
| **CwsHostUsage**                   | Pointer to **float64** | The Cloud Workload Security host usage by tag(s)                                                                               | [optional] |
| **DbmHostsPercentage**             | Pointer to **float64** | The percentage of Database Monitoring host usage by tag(s).                                                                    | [optional] |
| **DbmHostsUsage**                  | Pointer to **float64** | The Database Monitoring host usage by tag(s).                                                                                  | [optional] |
| **DbmQueriesPercentage**           | Pointer to **float64** | The percentage of Database Monitoring normalized queries usage by tag(s).                                                      | [optional] |
| **DbmQueriesUsage**                | Pointer to **float64** | The Database Monitoring normalized queries usage by tag(s).                                                                    | [optional] |
| **EstimatedIndexedLogsPercentage** | Pointer to **float64** | The percentage of estimated live indexed logs usage by tag(s). Note this field is in private beta.                             | [optional] |
| **EstimatedIndexedLogsUsage**      | Pointer to **float64** | The estimated live indexed logs usage by tag(s). Note this field is in private beta.                                           | [optional] |
| **InfraHostPercentage**            | Pointer to **float64** | The percentage of infrastructure host usage by tag(s).                                                                         | [optional] |
| **InfraHostUsage**                 | Pointer to **float64** | The infrastructure host usage by tag(s).                                                                                       | [optional] |
| **LambdaFunctionsPercentage**      | Pointer to **float64** | The percentage of Lambda function usage by tag(s).                                                                             | [optional] |
| **LambdaFunctionsUsage**           | Pointer to **float64** | The Lambda function usage by tag(s).                                                                                           | [optional] |
| **LambdaInvocationsPercentage**    | Pointer to **float64** | The percentage of Lambda invocation usage by tag(s).                                                                           | [optional] |
| **LambdaInvocationsUsage**         | Pointer to **float64** | The Lambda invocation usage by tag(s).                                                                                         | [optional] |
| **LambdaPercentage**               | Pointer to **float64** | The percentage of Lambda function usage by tag(s). **Note** this field is deprecated. Use lambda_functions_percentage instead. | [optional] |
| **LambdaUsage**                    | Pointer to **float64** | The Lambda function usage by tag(s). **Note** this field is deprecated. Use lambda_functions_usage instead.                    | [optional] |
| **NpmHostPercentage**              | Pointer to **float64** | The percentage of network host usage by tag(s).                                                                                | [optional] |
| **NpmHostUsage**                   | Pointer to **float64** | The network host usage by tag(s).                                                                                              | [optional] |
| **ProfiledContainerPercentage**    | Pointer to **float64** | The percentage of profiled containers usage by tag(s).                                                                         | [optional] |
| **ProfiledContainerUsage**         | Pointer to **float64** | The profiled container usage by tag(s).                                                                                        | [optional] |
| **ProfiledHostsPercentage**        | Pointer to **float64** | The percentage of profiled hosts usage by tag(s).                                                                              | [optional] |
| **ProfiledHostsUsage**             | Pointer to **float64** | The profiled host usage by tag(s).                                                                                             | [optional] |
| **SnmpPercentage**                 | Pointer to **float64** | The percentage of network device usage by tag(s).                                                                              | [optional] |
| **SnmpUsage**                      | Pointer to **float64** | The network device usage by tag(s).                                                                                            | [optional] |

## Methods

### NewUsageAttributionValues

`func NewUsageAttributionValues() *UsageAttributionValues`

NewUsageAttributionValues instantiates a new UsageAttributionValues object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageAttributionValuesWithDefaults

`func NewUsageAttributionValuesWithDefaults() *UsageAttributionValues`

NewUsageAttributionValuesWithDefaults instantiates a new UsageAttributionValues object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

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

### GetCspmContainerPercentage

`func (o *UsageAttributionValues) GetCspmContainerPercentage() float64`

GetCspmContainerPercentage returns the CspmContainerPercentage field if non-nil, zero value otherwise.

### GetCspmContainerPercentageOk

`func (o *UsageAttributionValues) GetCspmContainerPercentageOk() (*float64, bool)`

GetCspmContainerPercentageOk returns a tuple with the CspmContainerPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmContainerPercentage

`func (o *UsageAttributionValues) SetCspmContainerPercentage(v float64)`

SetCspmContainerPercentage sets CspmContainerPercentage field to given value.

### HasCspmContainerPercentage

`func (o *UsageAttributionValues) HasCspmContainerPercentage() bool`

HasCspmContainerPercentage returns a boolean if a field has been set.

### GetCspmContainerUsage

`func (o *UsageAttributionValues) GetCspmContainerUsage() float64`

GetCspmContainerUsage returns the CspmContainerUsage field if non-nil, zero value otherwise.

### GetCspmContainerUsageOk

`func (o *UsageAttributionValues) GetCspmContainerUsageOk() (*float64, bool)`

GetCspmContainerUsageOk returns a tuple with the CspmContainerUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmContainerUsage

`func (o *UsageAttributionValues) SetCspmContainerUsage(v float64)`

SetCspmContainerUsage sets CspmContainerUsage field to given value.

### HasCspmContainerUsage

`func (o *UsageAttributionValues) HasCspmContainerUsage() bool`

HasCspmContainerUsage returns a boolean if a field has been set.

### GetCspmHostPercentage

`func (o *UsageAttributionValues) GetCspmHostPercentage() float64`

GetCspmHostPercentage returns the CspmHostPercentage field if non-nil, zero value otherwise.

### GetCspmHostPercentageOk

`func (o *UsageAttributionValues) GetCspmHostPercentageOk() (*float64, bool)`

GetCspmHostPercentageOk returns a tuple with the CspmHostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmHostPercentage

`func (o *UsageAttributionValues) SetCspmHostPercentage(v float64)`

SetCspmHostPercentage sets CspmHostPercentage field to given value.

### HasCspmHostPercentage

`func (o *UsageAttributionValues) HasCspmHostPercentage() bool`

HasCspmHostPercentage returns a boolean if a field has been set.

### GetCspmHostUsage

`func (o *UsageAttributionValues) GetCspmHostUsage() float64`

GetCspmHostUsage returns the CspmHostUsage field if non-nil, zero value otherwise.

### GetCspmHostUsageOk

`func (o *UsageAttributionValues) GetCspmHostUsageOk() (*float64, bool)`

GetCspmHostUsageOk returns a tuple with the CspmHostUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmHostUsage

`func (o *UsageAttributionValues) SetCspmHostUsage(v float64)`

SetCspmHostUsage sets CspmHostUsage field to given value.

### HasCspmHostUsage

`func (o *UsageAttributionValues) HasCspmHostUsage() bool`

HasCspmHostUsage returns a boolean if a field has been set.

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

### GetCwsContainerPercentage

`func (o *UsageAttributionValues) GetCwsContainerPercentage() float64`

GetCwsContainerPercentage returns the CwsContainerPercentage field if non-nil, zero value otherwise.

### GetCwsContainerPercentageOk

`func (o *UsageAttributionValues) GetCwsContainerPercentageOk() (*float64, bool)`

GetCwsContainerPercentageOk returns a tuple with the CwsContainerPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsContainerPercentage

`func (o *UsageAttributionValues) SetCwsContainerPercentage(v float64)`

SetCwsContainerPercentage sets CwsContainerPercentage field to given value.

### HasCwsContainerPercentage

`func (o *UsageAttributionValues) HasCwsContainerPercentage() bool`

HasCwsContainerPercentage returns a boolean if a field has been set.

### GetCwsContainerUsage

`func (o *UsageAttributionValues) GetCwsContainerUsage() float64`

GetCwsContainerUsage returns the CwsContainerUsage field if non-nil, zero value otherwise.

### GetCwsContainerUsageOk

`func (o *UsageAttributionValues) GetCwsContainerUsageOk() (*float64, bool)`

GetCwsContainerUsageOk returns a tuple with the CwsContainerUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsContainerUsage

`func (o *UsageAttributionValues) SetCwsContainerUsage(v float64)`

SetCwsContainerUsage sets CwsContainerUsage field to given value.

### HasCwsContainerUsage

`func (o *UsageAttributionValues) HasCwsContainerUsage() bool`

HasCwsContainerUsage returns a boolean if a field has been set.

### GetCwsHostPercentage

`func (o *UsageAttributionValues) GetCwsHostPercentage() float64`

GetCwsHostPercentage returns the CwsHostPercentage field if non-nil, zero value otherwise.

### GetCwsHostPercentageOk

`func (o *UsageAttributionValues) GetCwsHostPercentageOk() (*float64, bool)`

GetCwsHostPercentageOk returns a tuple with the CwsHostPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsHostPercentage

`func (o *UsageAttributionValues) SetCwsHostPercentage(v float64)`

SetCwsHostPercentage sets CwsHostPercentage field to given value.

### HasCwsHostPercentage

`func (o *UsageAttributionValues) HasCwsHostPercentage() bool`

HasCwsHostPercentage returns a boolean if a field has been set.

### GetCwsHostUsage

`func (o *UsageAttributionValues) GetCwsHostUsage() float64`

GetCwsHostUsage returns the CwsHostUsage field if non-nil, zero value otherwise.

### GetCwsHostUsageOk

`func (o *UsageAttributionValues) GetCwsHostUsageOk() (*float64, bool)`

GetCwsHostUsageOk returns a tuple with the CwsHostUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsHostUsage

`func (o *UsageAttributionValues) SetCwsHostUsage(v float64)`

SetCwsHostUsage sets CwsHostUsage field to given value.

### HasCwsHostUsage

`func (o *UsageAttributionValues) HasCwsHostUsage() bool`

HasCwsHostUsage returns a boolean if a field has been set.

### GetDbmHostsPercentage

`func (o *UsageAttributionValues) GetDbmHostsPercentage() float64`

GetDbmHostsPercentage returns the DbmHostsPercentage field if non-nil, zero value otherwise.

### GetDbmHostsPercentageOk

`func (o *UsageAttributionValues) GetDbmHostsPercentageOk() (*float64, bool)`

GetDbmHostsPercentageOk returns a tuple with the DbmHostsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmHostsPercentage

`func (o *UsageAttributionValues) SetDbmHostsPercentage(v float64)`

SetDbmHostsPercentage sets DbmHostsPercentage field to given value.

### HasDbmHostsPercentage

`func (o *UsageAttributionValues) HasDbmHostsPercentage() bool`

HasDbmHostsPercentage returns a boolean if a field has been set.

### GetDbmHostsUsage

`func (o *UsageAttributionValues) GetDbmHostsUsage() float64`

GetDbmHostsUsage returns the DbmHostsUsage field if non-nil, zero value otherwise.

### GetDbmHostsUsageOk

`func (o *UsageAttributionValues) GetDbmHostsUsageOk() (*float64, bool)`

GetDbmHostsUsageOk returns a tuple with the DbmHostsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmHostsUsage

`func (o *UsageAttributionValues) SetDbmHostsUsage(v float64)`

SetDbmHostsUsage sets DbmHostsUsage field to given value.

### HasDbmHostsUsage

`func (o *UsageAttributionValues) HasDbmHostsUsage() bool`

HasDbmHostsUsage returns a boolean if a field has been set.

### GetDbmQueriesPercentage

`func (o *UsageAttributionValues) GetDbmQueriesPercentage() float64`

GetDbmQueriesPercentage returns the DbmQueriesPercentage field if non-nil, zero value otherwise.

### GetDbmQueriesPercentageOk

`func (o *UsageAttributionValues) GetDbmQueriesPercentageOk() (*float64, bool)`

GetDbmQueriesPercentageOk returns a tuple with the DbmQueriesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmQueriesPercentage

`func (o *UsageAttributionValues) SetDbmQueriesPercentage(v float64)`

SetDbmQueriesPercentage sets DbmQueriesPercentage field to given value.

### HasDbmQueriesPercentage

`func (o *UsageAttributionValues) HasDbmQueriesPercentage() bool`

HasDbmQueriesPercentage returns a boolean if a field has been set.

### GetDbmQueriesUsage

`func (o *UsageAttributionValues) GetDbmQueriesUsage() float64`

GetDbmQueriesUsage returns the DbmQueriesUsage field if non-nil, zero value otherwise.

### GetDbmQueriesUsageOk

`func (o *UsageAttributionValues) GetDbmQueriesUsageOk() (*float64, bool)`

GetDbmQueriesUsageOk returns a tuple with the DbmQueriesUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmQueriesUsage

`func (o *UsageAttributionValues) SetDbmQueriesUsage(v float64)`

SetDbmQueriesUsage sets DbmQueriesUsage field to given value.

### HasDbmQueriesUsage

`func (o *UsageAttributionValues) HasDbmQueriesUsage() bool`

HasDbmQueriesUsage returns a boolean if a field has been set.

### GetEstimatedIndexedLogsPercentage

`func (o *UsageAttributionValues) GetEstimatedIndexedLogsPercentage() float64`

GetEstimatedIndexedLogsPercentage returns the EstimatedIndexedLogsPercentage field if non-nil, zero value otherwise.

### GetEstimatedIndexedLogsPercentageOk

`func (o *UsageAttributionValues) GetEstimatedIndexedLogsPercentageOk() (*float64, bool)`

GetEstimatedIndexedLogsPercentageOk returns a tuple with the EstimatedIndexedLogsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedIndexedLogsPercentage

`func (o *UsageAttributionValues) SetEstimatedIndexedLogsPercentage(v float64)`

SetEstimatedIndexedLogsPercentage sets EstimatedIndexedLogsPercentage field to given value.

### HasEstimatedIndexedLogsPercentage

`func (o *UsageAttributionValues) HasEstimatedIndexedLogsPercentage() bool`

HasEstimatedIndexedLogsPercentage returns a boolean if a field has been set.

### GetEstimatedIndexedLogsUsage

`func (o *UsageAttributionValues) GetEstimatedIndexedLogsUsage() float64`

GetEstimatedIndexedLogsUsage returns the EstimatedIndexedLogsUsage field if non-nil, zero value otherwise.

### GetEstimatedIndexedLogsUsageOk

`func (o *UsageAttributionValues) GetEstimatedIndexedLogsUsageOk() (*float64, bool)`

GetEstimatedIndexedLogsUsageOk returns a tuple with the EstimatedIndexedLogsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedIndexedLogsUsage

`func (o *UsageAttributionValues) SetEstimatedIndexedLogsUsage(v float64)`

SetEstimatedIndexedLogsUsage sets EstimatedIndexedLogsUsage field to given value.

### HasEstimatedIndexedLogsUsage

`func (o *UsageAttributionValues) HasEstimatedIndexedLogsUsage() bool`

HasEstimatedIndexedLogsUsage returns a boolean if a field has been set.

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

### GetLambdaFunctionsPercentage

`func (o *UsageAttributionValues) GetLambdaFunctionsPercentage() float64`

GetLambdaFunctionsPercentage returns the LambdaFunctionsPercentage field if non-nil, zero value otherwise.

### GetLambdaFunctionsPercentageOk

`func (o *UsageAttributionValues) GetLambdaFunctionsPercentageOk() (*float64, bool)`

GetLambdaFunctionsPercentageOk returns a tuple with the LambdaFunctionsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaFunctionsPercentage

`func (o *UsageAttributionValues) SetLambdaFunctionsPercentage(v float64)`

SetLambdaFunctionsPercentage sets LambdaFunctionsPercentage field to given value.

### HasLambdaFunctionsPercentage

`func (o *UsageAttributionValues) HasLambdaFunctionsPercentage() bool`

HasLambdaFunctionsPercentage returns a boolean if a field has been set.

### GetLambdaFunctionsUsage

`func (o *UsageAttributionValues) GetLambdaFunctionsUsage() float64`

GetLambdaFunctionsUsage returns the LambdaFunctionsUsage field if non-nil, zero value otherwise.

### GetLambdaFunctionsUsageOk

`func (o *UsageAttributionValues) GetLambdaFunctionsUsageOk() (*float64, bool)`

GetLambdaFunctionsUsageOk returns a tuple with the LambdaFunctionsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaFunctionsUsage

`func (o *UsageAttributionValues) SetLambdaFunctionsUsage(v float64)`

SetLambdaFunctionsUsage sets LambdaFunctionsUsage field to given value.

### HasLambdaFunctionsUsage

`func (o *UsageAttributionValues) HasLambdaFunctionsUsage() bool`

HasLambdaFunctionsUsage returns a boolean if a field has been set.

### GetLambdaInvocationsPercentage

`func (o *UsageAttributionValues) GetLambdaInvocationsPercentage() float64`

GetLambdaInvocationsPercentage returns the LambdaInvocationsPercentage field if non-nil, zero value otherwise.

### GetLambdaInvocationsPercentageOk

`func (o *UsageAttributionValues) GetLambdaInvocationsPercentageOk() (*float64, bool)`

GetLambdaInvocationsPercentageOk returns a tuple with the LambdaInvocationsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaInvocationsPercentage

`func (o *UsageAttributionValues) SetLambdaInvocationsPercentage(v float64)`

SetLambdaInvocationsPercentage sets LambdaInvocationsPercentage field to given value.

### HasLambdaInvocationsPercentage

`func (o *UsageAttributionValues) HasLambdaInvocationsPercentage() bool`

HasLambdaInvocationsPercentage returns a boolean if a field has been set.

### GetLambdaInvocationsUsage

`func (o *UsageAttributionValues) GetLambdaInvocationsUsage() float64`

GetLambdaInvocationsUsage returns the LambdaInvocationsUsage field if non-nil, zero value otherwise.

### GetLambdaInvocationsUsageOk

`func (o *UsageAttributionValues) GetLambdaInvocationsUsageOk() (*float64, bool)`

GetLambdaInvocationsUsageOk returns a tuple with the LambdaInvocationsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaInvocationsUsage

`func (o *UsageAttributionValues) SetLambdaInvocationsUsage(v float64)`

SetLambdaInvocationsUsage sets LambdaInvocationsUsage field to given value.

### HasLambdaInvocationsUsage

`func (o *UsageAttributionValues) HasLambdaInvocationsUsage() bool`

HasLambdaInvocationsUsage returns a boolean if a field has been set.

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

### GetProfiledContainerPercentage

`func (o *UsageAttributionValues) GetProfiledContainerPercentage() float64`

GetProfiledContainerPercentage returns the ProfiledContainerPercentage field if non-nil, zero value otherwise.

### GetProfiledContainerPercentageOk

`func (o *UsageAttributionValues) GetProfiledContainerPercentageOk() (*float64, bool)`

GetProfiledContainerPercentageOk returns a tuple with the ProfiledContainerPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiledContainerPercentage

`func (o *UsageAttributionValues) SetProfiledContainerPercentage(v float64)`

SetProfiledContainerPercentage sets ProfiledContainerPercentage field to given value.

### HasProfiledContainerPercentage

`func (o *UsageAttributionValues) HasProfiledContainerPercentage() bool`

HasProfiledContainerPercentage returns a boolean if a field has been set.

### GetProfiledContainerUsage

`func (o *UsageAttributionValues) GetProfiledContainerUsage() float64`

GetProfiledContainerUsage returns the ProfiledContainerUsage field if non-nil, zero value otherwise.

### GetProfiledContainerUsageOk

`func (o *UsageAttributionValues) GetProfiledContainerUsageOk() (*float64, bool)`

GetProfiledContainerUsageOk returns a tuple with the ProfiledContainerUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiledContainerUsage

`func (o *UsageAttributionValues) SetProfiledContainerUsage(v float64)`

SetProfiledContainerUsage sets ProfiledContainerUsage field to given value.

### HasProfiledContainerUsage

`func (o *UsageAttributionValues) HasProfiledContainerUsage() bool`

HasProfiledContainerUsage returns a boolean if a field has been set.

### GetProfiledHostsPercentage

`func (o *UsageAttributionValues) GetProfiledHostsPercentage() float64`

GetProfiledHostsPercentage returns the ProfiledHostsPercentage field if non-nil, zero value otherwise.

### GetProfiledHostsPercentageOk

`func (o *UsageAttributionValues) GetProfiledHostsPercentageOk() (*float64, bool)`

GetProfiledHostsPercentageOk returns a tuple with the ProfiledHostsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiledHostsPercentage

`func (o *UsageAttributionValues) SetProfiledHostsPercentage(v float64)`

SetProfiledHostsPercentage sets ProfiledHostsPercentage field to given value.

### HasProfiledHostsPercentage

`func (o *UsageAttributionValues) HasProfiledHostsPercentage() bool`

HasProfiledHostsPercentage returns a boolean if a field has been set.

### GetProfiledHostsUsage

`func (o *UsageAttributionValues) GetProfiledHostsUsage() float64`

GetProfiledHostsUsage returns the ProfiledHostsUsage field if non-nil, zero value otherwise.

### GetProfiledHostsUsageOk

`func (o *UsageAttributionValues) GetProfiledHostsUsageOk() (*float64, bool)`

GetProfiledHostsUsageOk returns a tuple with the ProfiledHostsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiledHostsUsage

`func (o *UsageAttributionValues) SetProfiledHostsUsage(v float64)`

SetProfiledHostsUsage sets ProfiledHostsUsage field to given value.

### HasProfiledHostsUsage

`func (o *UsageAttributionValues) HasProfiledHostsUsage() bool`

HasProfiledHostsUsage returns a boolean if a field has been set.

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
