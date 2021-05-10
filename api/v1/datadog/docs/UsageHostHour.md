# UsageHostHour

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AgentHostCount** | Pointer to **int64** | Contains the total number of infrastructure hosts reporting during a given hour that were running the Datadog Agent. | [optional] 
**AlibabaHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via Alibaba integration (and were NOT running the Datadog Agent). | [optional] 
**ApmAzureAppServiceHostCount** | Pointer to **int64** | Contains the total number of Azure App Services hosts using APM. | [optional] 
**ApmHostCount** | Pointer to **int64** | Shows the total number of hosts using APM during the hour, these are counted as billable (except during trial periods). | [optional] 
**AwsHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via the AWS integration (and were NOT running the Datadog Agent). | [optional] 
**AzureHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via Azure integration (and were NOT running the Datadog Agent). | [optional] 
**ContainerCount** | Pointer to **int64** | Shows the total number of containers reported by the Docker integration during the hour. | [optional] 
**GcpHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via the Google Cloud integration (and were NOT running the Datadog Agent). | [optional] 
**HerokuHostCount** | Pointer to **int64** | Contains the total number of Heroku dynos reported by the Datadog Agent. | [optional] 
**HostCount** | Pointer to **int64** | Contains the total number of billable infrastructure hosts reporting during a given hour. This is the sum of &#x60;agent_host_count&#x60;, &#x60;aws_host_count&#x60;, and &#x60;gcp_host_count&#x60;. | [optional] 
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 
**InfraAzureAppService** | Pointer to **int64** | Contains the total number of hosts that reported via the Azure App Services integration (and were NOT running the Datadog Agent). | [optional] 
**OpentelemetryHostCount** | Pointer to **int64** | Contains the total number of hosts reported by Datadog exporter for the OpenTelemetry Collector. | [optional] 
**VsphereHostCount** | Pointer to **int64** | Contains the total number of hosts that reported via vSphere integration (and were NOT running the Datadog Agent). | [optional] 

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

### GetApmAzureAppServiceHostCount

`func (o *UsageHostHour) GetApmAzureAppServiceHostCount() int64`

GetApmAzureAppServiceHostCount returns the ApmAzureAppServiceHostCount field if non-nil, zero value otherwise.

### GetApmAzureAppServiceHostCountOk

`func (o *UsageHostHour) GetApmAzureAppServiceHostCountOk() (*int64, bool)`

GetApmAzureAppServiceHostCountOk returns a tuple with the ApmAzureAppServiceHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmAzureAppServiceHostCount

`func (o *UsageHostHour) SetApmAzureAppServiceHostCount(v int64)`

SetApmAzureAppServiceHostCount sets ApmAzureAppServiceHostCount field to given value.

### HasApmAzureAppServiceHostCount

`func (o *UsageHostHour) HasApmAzureAppServiceHostCount() bool`

HasApmAzureAppServiceHostCount returns a boolean if a field has been set.

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

### GetHerokuHostCount

`func (o *UsageHostHour) GetHerokuHostCount() int64`

GetHerokuHostCount returns the HerokuHostCount field if non-nil, zero value otherwise.

### GetHerokuHostCountOk

`func (o *UsageHostHour) GetHerokuHostCountOk() (*int64, bool)`

GetHerokuHostCountOk returns a tuple with the HerokuHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHerokuHostCount

`func (o *UsageHostHour) SetHerokuHostCount(v int64)`

SetHerokuHostCount sets HerokuHostCount field to given value.

### HasHerokuHostCount

`func (o *UsageHostHour) HasHerokuHostCount() bool`

HasHerokuHostCount returns a boolean if a field has been set.

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

### GetInfraAzureAppService

`func (o *UsageHostHour) GetInfraAzureAppService() int64`

GetInfraAzureAppService returns the InfraAzureAppService field if non-nil, zero value otherwise.

### GetInfraAzureAppServiceOk

`func (o *UsageHostHour) GetInfraAzureAppServiceOk() (*int64, bool)`

GetInfraAzureAppServiceOk returns a tuple with the InfraAzureAppService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraAzureAppService

`func (o *UsageHostHour) SetInfraAzureAppService(v int64)`

SetInfraAzureAppService sets InfraAzureAppService field to given value.

### HasInfraAzureAppService

`func (o *UsageHostHour) HasInfraAzureAppService() bool`

HasInfraAzureAppService returns a boolean if a field has been set.

### GetOpentelemetryHostCount

`func (o *UsageHostHour) GetOpentelemetryHostCount() int64`

GetOpentelemetryHostCount returns the OpentelemetryHostCount field if non-nil, zero value otherwise.

### GetOpentelemetryHostCountOk

`func (o *UsageHostHour) GetOpentelemetryHostCountOk() (*int64, bool)`

GetOpentelemetryHostCountOk returns a tuple with the OpentelemetryHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpentelemetryHostCount

`func (o *UsageHostHour) SetOpentelemetryHostCount(v int64)`

SetOpentelemetryHostCount sets OpentelemetryHostCount field to given value.

### HasOpentelemetryHostCount

`func (o *UsageHostHour) HasOpentelemetryHostCount() bool`

HasOpentelemetryHostCount returns a boolean if a field has been set.

### GetVsphereHostCount

`func (o *UsageHostHour) GetVsphereHostCount() int64`

GetVsphereHostCount returns the VsphereHostCount field if non-nil, zero value otherwise.

### GetVsphereHostCountOk

`func (o *UsageHostHour) GetVsphereHostCountOk() (*int64, bool)`

GetVsphereHostCountOk returns a tuple with the VsphereHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereHostCount

`func (o *UsageHostHour) SetVsphereHostCount(v int64)`

SetVsphereHostCount sets VsphereHostCount field to given value.

### HasVsphereHostCount

`func (o *UsageHostHour) HasVsphereHostCount() bool`

HasVsphereHostCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


