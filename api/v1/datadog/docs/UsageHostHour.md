# UsageHostHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostCount** | Pointer to **int64** | Contains the total number of infrastructure hosts reporting during a given hour that were running the Datadog Agent. | [optional] 
**AlibabaHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via Alibaba integration (and were NOT running the Datadog Agent). | [optional] 
**ApmHostCount** | Pointer to **int64** | Shows the total number of hosts using APM during the hour, these are counted as billable (except during trial periods). | [optional] 
**AwsHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via the AWS integration (and were NOT running the Datadog Agent). | [optional] 
**AzureHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via Azure integration (and were NOT running the Datadog Agent). | [optional] 
**ContainerCount** | Pointer to **int64** | Contains the total number of billable infrastructure hosts reporting during a given hour. This is the sum of &#x60;agent_host_count&#x60;, &#x60;aws_host_count&#x60;, and &#x60;gcp_host_count&#x60;. | [optional] 
**GcpHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via the Google Cloud integration (and were NOT running the Datadog Agent). | [optional] 
**HostCount** | Pointer to **int64** | Shows the total number of containers reporting via the Docker integration during the hour. | [optional] 
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 

## Methods

### NewUsageHostHour

`func NewUsageHostHour() *UsageHostHour`

NewUsageHostHour instantiates a new UsageHostHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageHostHourWithDefaults

`func NewUsageHostHourWithDefaults() *UsageHostHour`

NewUsageHostHourWithDefaults instantiates a new UsageHostHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentHostCount

`func (o *UsageHostHour) GetAgentHostCount() int64`

GetAgentHostCount returns the AgentHostCount field if non-nil, zero value otherwise.

### GetAgentHostCountOk

`func (o *UsageHostHour) GetAgentHostCountOk() (*int64, bool)`

GetAgentHostCountOk returns a tuple with the AgentHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentHostCount

`func (o *UsageHostHour) SetAgentHostCount(v int64)`

SetAgentHostCount sets AgentHostCount field to given value.

### HasAgentHostCount

`func (o *UsageHostHour) HasAgentHostCount() bool`

HasAgentHostCount returns a boolean if a field has been set.

### GetAlibabaHostCount

`func (o *UsageHostHour) GetAlibabaHostCount() int64`

GetAlibabaHostCount returns the AlibabaHostCount field if non-nil, zero value otherwise.

### GetAlibabaHostCountOk

`func (o *UsageHostHour) GetAlibabaHostCountOk() (*int64, bool)`

GetAlibabaHostCountOk returns a tuple with the AlibabaHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlibabaHostCount

`func (o *UsageHostHour) SetAlibabaHostCount(v int64)`

SetAlibabaHostCount sets AlibabaHostCount field to given value.

### HasAlibabaHostCount

`func (o *UsageHostHour) HasAlibabaHostCount() bool`

HasAlibabaHostCount returns a boolean if a field has been set.

### GetApmHostCount

`func (o *UsageHostHour) GetApmHostCount() int64`

GetApmHostCount returns the ApmHostCount field if non-nil, zero value otherwise.

### GetApmHostCountOk

`func (o *UsageHostHour) GetApmHostCountOk() (*int64, bool)`

GetApmHostCountOk returns a tuple with the ApmHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostCount

`func (o *UsageHostHour) SetApmHostCount(v int64)`

SetApmHostCount sets ApmHostCount field to given value.

### HasApmHostCount

`func (o *UsageHostHour) HasApmHostCount() bool`

HasApmHostCount returns a boolean if a field has been set.

### GetAwsHostCount

`func (o *UsageHostHour) GetAwsHostCount() int64`

GetAwsHostCount returns the AwsHostCount field if non-nil, zero value otherwise.

### GetAwsHostCountOk

`func (o *UsageHostHour) GetAwsHostCountOk() (*int64, bool)`

GetAwsHostCountOk returns a tuple with the AwsHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostCount

`func (o *UsageHostHour) SetAwsHostCount(v int64)`

SetAwsHostCount sets AwsHostCount field to given value.

### HasAwsHostCount

`func (o *UsageHostHour) HasAwsHostCount() bool`

HasAwsHostCount returns a boolean if a field has been set.

### GetAzureHostCount

`func (o *UsageHostHour) GetAzureHostCount() int64`

GetAzureHostCount returns the AzureHostCount field if non-nil, zero value otherwise.

### GetAzureHostCountOk

`func (o *UsageHostHour) GetAzureHostCountOk() (*int64, bool)`

GetAzureHostCountOk returns a tuple with the AzureHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureHostCount

`func (o *UsageHostHour) SetAzureHostCount(v int64)`

SetAzureHostCount sets AzureHostCount field to given value.

### HasAzureHostCount

`func (o *UsageHostHour) HasAzureHostCount() bool`

HasAzureHostCount returns a boolean if a field has been set.

### GetContainerCount

`func (o *UsageHostHour) GetContainerCount() int64`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *UsageHostHour) GetContainerCountOk() (*int64, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *UsageHostHour) SetContainerCount(v int64)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *UsageHostHour) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetGcpHostCount

`func (o *UsageHostHour) GetGcpHostCount() int64`

GetGcpHostCount returns the GcpHostCount field if non-nil, zero value otherwise.

### GetGcpHostCountOk

`func (o *UsageHostHour) GetGcpHostCountOk() (*int64, bool)`

GetGcpHostCountOk returns a tuple with the GcpHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpHostCount

`func (o *UsageHostHour) SetGcpHostCount(v int64)`

SetGcpHostCount sets GcpHostCount field to given value.

### HasGcpHostCount

`func (o *UsageHostHour) HasGcpHostCount() bool`

HasGcpHostCount returns a boolean if a field has been set.

### GetHostCount

`func (o *UsageHostHour) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *UsageHostHour) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *UsageHostHour) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *UsageHostHour) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageHostHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageHostHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageHostHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageHostHour) HasHour() bool`

HasHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


