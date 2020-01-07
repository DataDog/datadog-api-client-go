# UsageHostHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostCount** | Pointer to **int64** | Contains the total number of infrastructure hosts reporting during a given hour that were running the Datadog Agent. | [optional] 
**ApmHostCount** | Pointer to **int64** | Shows the total number of hosts using APM during the hour, these are counted as billable (except during trial periods). | [optional] 
**AwsHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via the AWS integration (and were NOT running the Datadog Agent). When AWS or GCP hosts are also running the Datadog Agent, they are counted as Agent hosts, NOT as AWS or GCP. | [optional] 
**ContainerCount** | Pointer to **int64** | Contains the total number of billable infrastructure hosts reporting during a given hour. This is the sum of &#x60;agent_host_count&#x60;, &#x60;aws_host_count&#x60;, and &#x60;gcp_host_count&#x60;. | [optional] 
**GcpHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via the Google Cloud integration (and were NOT running the Datadog Agent). | [optional] 
**HostCount** | Pointer to **int64** | Shows the total number of containers reporting via the Docker integration during the hour. | [optional] 
**Hour** | Pointer to [**time.Time**](time.Time.md) | The hour for the usage. | [optional] 

## Methods

### GetAgentHostCount

`func (o *UsageHostHour) GetAgentHostCount() int64`

GetAgentHostCount returns the AgentHostCount field if non-nil, zero value otherwise.

### GetAgentHostCountOk

`func (o *UsageHostHour) GetAgentHostCountOk() (int64, bool)`

GetAgentHostCountOk returns a tuple with the AgentHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgentHostCount

`func (o *UsageHostHour) HasAgentHostCount() bool`

HasAgentHostCount returns a boolean if a field has been set.

### SetAgentHostCount

`func (o *UsageHostHour) SetAgentHostCount(v int64)`

SetAgentHostCount gets a reference to the given int64 and assigns it to the AgentHostCount field.

### GetApmHostCount

`func (o *UsageHostHour) GetApmHostCount() int64`

GetApmHostCount returns the ApmHostCount field if non-nil, zero value otherwise.

### GetApmHostCountOk

`func (o *UsageHostHour) GetApmHostCountOk() (int64, bool)`

GetApmHostCountOk returns a tuple with the ApmHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApmHostCount

`func (o *UsageHostHour) HasApmHostCount() bool`

HasApmHostCount returns a boolean if a field has been set.

### SetApmHostCount

`func (o *UsageHostHour) SetApmHostCount(v int64)`

SetApmHostCount gets a reference to the given int64 and assigns it to the ApmHostCount field.

### GetAwsHostCount

`func (o *UsageHostHour) GetAwsHostCount() int64`

GetAwsHostCount returns the AwsHostCount field if non-nil, zero value otherwise.

### GetAwsHostCountOk

`func (o *UsageHostHour) GetAwsHostCountOk() (int64, bool)`

GetAwsHostCountOk returns a tuple with the AwsHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAwsHostCount

`func (o *UsageHostHour) HasAwsHostCount() bool`

HasAwsHostCount returns a boolean if a field has been set.

### SetAwsHostCount

`func (o *UsageHostHour) SetAwsHostCount(v int64)`

SetAwsHostCount gets a reference to the given int64 and assigns it to the AwsHostCount field.

### GetContainerCount

`func (o *UsageHostHour) GetContainerCount() int64`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *UsageHostHour) GetContainerCountOk() (int64, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContainerCount

`func (o *UsageHostHour) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### SetContainerCount

`func (o *UsageHostHour) SetContainerCount(v int64)`

SetContainerCount gets a reference to the given int64 and assigns it to the ContainerCount field.

### GetGcpHostCount

`func (o *UsageHostHour) GetGcpHostCount() int64`

GetGcpHostCount returns the GcpHostCount field if non-nil, zero value otherwise.

### GetGcpHostCountOk

`func (o *UsageHostHour) GetGcpHostCountOk() (int64, bool)`

GetGcpHostCountOk returns a tuple with the GcpHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGcpHostCount

`func (o *UsageHostHour) HasGcpHostCount() bool`

HasGcpHostCount returns a boolean if a field has been set.

### SetGcpHostCount

`func (o *UsageHostHour) SetGcpHostCount(v int64)`

SetGcpHostCount gets a reference to the given int64 and assigns it to the GcpHostCount field.

### GetHostCount

`func (o *UsageHostHour) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *UsageHostHour) GetHostCountOk() (int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostCount

`func (o *UsageHostHour) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### SetHostCount

`func (o *UsageHostHour) SetHostCount(v int64)`

SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.

### GetHour

`func (o *UsageHostHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageHostHour) GetHourOk() (time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHour

`func (o *UsageHostHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### SetHour

`func (o *UsageHostHour) SetHour(v time.Time)`

SetHour gets a reference to the given time.Time and assigns it to the Hour field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


