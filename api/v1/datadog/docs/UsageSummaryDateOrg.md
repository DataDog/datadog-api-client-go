# UsageSummaryDateOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all agent hosts over all hours in the current date for the given org. | [optional] 
**ApmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current date for the given org. | [optional] 
**AwsHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all AWS hosts over all hours in the current date for the given org. | [optional] 
**AwsLambdaFuncCount** | Pointer to **int64** | Shows the sum of all AWS Labmda invocations over all hours in the current date for the given org. | [optional] 
**AwsLambdaInvocationsSum** | Pointer to **int64** | Shows the sum of all AWS Labmda invocations over all hours in the current date for the given org. | [optional] 
**BillableIngestedBytesSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current date for the given org. | [optional] 
**ContainerAvg** | Pointer to **int64** | Shows the average of all distinct containers over all hours in the current date for the given org. | [optional] 
**ContainerHwm** | Pointer to **int64** | Shows the high watermark of all distinct containers over all hours in the current date for the given org. | [optional] 
**CustomTsAvg** | Pointer to **int64** | Shows the average number of distinct custom metrics over all hours in the current date for the given org. | [optional] 
**FargateTasksCountAvg** | Pointer to **int64** |  | [optional] 
**FargateTasksCountHwm** | Pointer to **int64** | Shows the high watermark of all Fargate tasks over all hours in the current date for the given org. | [optional] 
**GcpHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all GCP hosts over all hours in the current date for the given org. | [optional] 
**Id** | Pointer to **string** | The organization id. | [optional] 
**IndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all log events indexed over all hours in the current date for the given org. | [optional] 
**InfraHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for the given org. | [optional] 
**IngestedEventsBytesSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current date for the given org. | [optional] 
**Name** | Pointer to **string** | The organization name. | [optional] 
**NetflowIndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all Network flows indexed over all hours in the current date for the given org. | [optional] 
**NpmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for the given org. | [optional] 
**PublicId** | Pointer to **string** | The organization public id. | [optional] 
**RumSessionCountSum** | Pointer to **int64** | Shows the sum of all RUM Sessions over all hours in the current date for the given org. | [optional] 
**SyntheticsBrowserCheckCallsCountSum** | Pointer to **int64** | Shows the sum of all Synthetic browser tests over all hours in the current date for the given org. | [optional] 
**SyntheticsCheckCallsCountSum** | Pointer to **int64** | Shows the sum of all Synthetic API tests over all hours in the current date for the given org. | [optional] 
**TraceSearchIndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all analyzed spans indexed over all hours in the current date for the given org. | [optional] 

## Methods

### NewUsageSummaryDateOrg

`func NewUsageSummaryDateOrg() *UsageSummaryDateOrg`

NewUsageSummaryDateOrg instantiates a new UsageSummaryDateOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSummaryDateOrgWithDefaults

`func NewUsageSummaryDateOrgWithDefaults() *UsageSummaryDateOrg`

NewUsageSummaryDateOrgWithDefaults instantiates a new UsageSummaryDateOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentHostTop99p

`func (o *UsageSummaryDateOrg) GetAgentHostTop99p() int64`

GetAgentHostTop99p returns the AgentHostTop99p field if non-nil, zero value otherwise.

### GetAgentHostTop99pOk

`func (o *UsageSummaryDateOrg) GetAgentHostTop99pOk() (*int64, bool)`

GetAgentHostTop99pOk returns a tuple with the AgentHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentHostTop99p

`func (o *UsageSummaryDateOrg) SetAgentHostTop99p(v int64)`

SetAgentHostTop99p sets AgentHostTop99p field to given value.

### HasAgentHostTop99p

`func (o *UsageSummaryDateOrg) HasAgentHostTop99p() bool`

HasAgentHostTop99p returns a boolean if a field has been set.

### GetApmHostTop99p

`func (o *UsageSummaryDateOrg) GetApmHostTop99p() int64`

GetApmHostTop99p returns the ApmHostTop99p field if non-nil, zero value otherwise.

### GetApmHostTop99pOk

`func (o *UsageSummaryDateOrg) GetApmHostTop99pOk() (*int64, bool)`

GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostTop99p

`func (o *UsageSummaryDateOrg) SetApmHostTop99p(v int64)`

SetApmHostTop99p sets ApmHostTop99p field to given value.

### HasApmHostTop99p

`func (o *UsageSummaryDateOrg) HasApmHostTop99p() bool`

HasApmHostTop99p returns a boolean if a field has been set.

### GetAwsHostTop99p

`func (o *UsageSummaryDateOrg) GetAwsHostTop99p() int64`

GetAwsHostTop99p returns the AwsHostTop99p field if non-nil, zero value otherwise.

### GetAwsHostTop99pOk

`func (o *UsageSummaryDateOrg) GetAwsHostTop99pOk() (*int64, bool)`

GetAwsHostTop99pOk returns a tuple with the AwsHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostTop99p

`func (o *UsageSummaryDateOrg) SetAwsHostTop99p(v int64)`

SetAwsHostTop99p sets AwsHostTop99p field to given value.

### HasAwsHostTop99p

`func (o *UsageSummaryDateOrg) HasAwsHostTop99p() bool`

HasAwsHostTop99p returns a boolean if a field has been set.

### GetAwsLambdaFuncCount

`func (o *UsageSummaryDateOrg) GetAwsLambdaFuncCount() int64`

GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field if non-nil, zero value otherwise.

### GetAwsLambdaFuncCountOk

`func (o *UsageSummaryDateOrg) GetAwsLambdaFuncCountOk() (*int64, bool)`

GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaFuncCount

`func (o *UsageSummaryDateOrg) SetAwsLambdaFuncCount(v int64)`

SetAwsLambdaFuncCount sets AwsLambdaFuncCount field to given value.

### HasAwsLambdaFuncCount

`func (o *UsageSummaryDateOrg) HasAwsLambdaFuncCount() bool`

HasAwsLambdaFuncCount returns a boolean if a field has been set.

### GetAwsLambdaInvocationsSum

`func (o *UsageSummaryDateOrg) GetAwsLambdaInvocationsSum() int64`

GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field if non-nil, zero value otherwise.

### GetAwsLambdaInvocationsSumOk

`func (o *UsageSummaryDateOrg) GetAwsLambdaInvocationsSumOk() (*int64, bool)`

GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaInvocationsSum

`func (o *UsageSummaryDateOrg) SetAwsLambdaInvocationsSum(v int64)`

SetAwsLambdaInvocationsSum sets AwsLambdaInvocationsSum field to given value.

### HasAwsLambdaInvocationsSum

`func (o *UsageSummaryDateOrg) HasAwsLambdaInvocationsSum() bool`

HasAwsLambdaInvocationsSum returns a boolean if a field has been set.

### GetBillableIngestedBytesSum

`func (o *UsageSummaryDateOrg) GetBillableIngestedBytesSum() int64`

GetBillableIngestedBytesSum returns the BillableIngestedBytesSum field if non-nil, zero value otherwise.

### GetBillableIngestedBytesSumOk

`func (o *UsageSummaryDateOrg) GetBillableIngestedBytesSumOk() (*int64, bool)`

GetBillableIngestedBytesSumOk returns a tuple with the BillableIngestedBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableIngestedBytesSum

`func (o *UsageSummaryDateOrg) SetBillableIngestedBytesSum(v int64)`

SetBillableIngestedBytesSum sets BillableIngestedBytesSum field to given value.

### HasBillableIngestedBytesSum

`func (o *UsageSummaryDateOrg) HasBillableIngestedBytesSum() bool`

HasBillableIngestedBytesSum returns a boolean if a field has been set.

### GetContainerAvg

`func (o *UsageSummaryDateOrg) GetContainerAvg() int64`

GetContainerAvg returns the ContainerAvg field if non-nil, zero value otherwise.

### GetContainerAvgOk

`func (o *UsageSummaryDateOrg) GetContainerAvgOk() (*int64, bool)`

GetContainerAvgOk returns a tuple with the ContainerAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerAvg

`func (o *UsageSummaryDateOrg) SetContainerAvg(v int64)`

SetContainerAvg sets ContainerAvg field to given value.

### HasContainerAvg

`func (o *UsageSummaryDateOrg) HasContainerAvg() bool`

HasContainerAvg returns a boolean if a field has been set.

### GetContainerHwm

`func (o *UsageSummaryDateOrg) GetContainerHwm() int64`

GetContainerHwm returns the ContainerHwm field if non-nil, zero value otherwise.

### GetContainerHwmOk

`func (o *UsageSummaryDateOrg) GetContainerHwmOk() (*int64, bool)`

GetContainerHwmOk returns a tuple with the ContainerHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHwm

`func (o *UsageSummaryDateOrg) SetContainerHwm(v int64)`

SetContainerHwm sets ContainerHwm field to given value.

### HasContainerHwm

`func (o *UsageSummaryDateOrg) HasContainerHwm() bool`

HasContainerHwm returns a boolean if a field has been set.

### GetCustomTsAvg

`func (o *UsageSummaryDateOrg) GetCustomTsAvg() int64`

GetCustomTsAvg returns the CustomTsAvg field if non-nil, zero value otherwise.

### GetCustomTsAvgOk

`func (o *UsageSummaryDateOrg) GetCustomTsAvgOk() (*int64, bool)`

GetCustomTsAvgOk returns a tuple with the CustomTsAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTsAvg

`func (o *UsageSummaryDateOrg) SetCustomTsAvg(v int64)`

SetCustomTsAvg sets CustomTsAvg field to given value.

### HasCustomTsAvg

`func (o *UsageSummaryDateOrg) HasCustomTsAvg() bool`

HasCustomTsAvg returns a boolean if a field has been set.

### GetFargateTasksCountAvg

`func (o *UsageSummaryDateOrg) GetFargateTasksCountAvg() int64`

GetFargateTasksCountAvg returns the FargateTasksCountAvg field if non-nil, zero value otherwise.

### GetFargateTasksCountAvgOk

`func (o *UsageSummaryDateOrg) GetFargateTasksCountAvgOk() (*int64, bool)`

GetFargateTasksCountAvgOk returns a tuple with the FargateTasksCountAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountAvg

`func (o *UsageSummaryDateOrg) SetFargateTasksCountAvg(v int64)`

SetFargateTasksCountAvg sets FargateTasksCountAvg field to given value.

### HasFargateTasksCountAvg

`func (o *UsageSummaryDateOrg) HasFargateTasksCountAvg() bool`

HasFargateTasksCountAvg returns a boolean if a field has been set.

### GetFargateTasksCountHwm

`func (o *UsageSummaryDateOrg) GetFargateTasksCountHwm() int64`

GetFargateTasksCountHwm returns the FargateTasksCountHwm field if non-nil, zero value otherwise.

### GetFargateTasksCountHwmOk

`func (o *UsageSummaryDateOrg) GetFargateTasksCountHwmOk() (*int64, bool)`

GetFargateTasksCountHwmOk returns a tuple with the FargateTasksCountHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountHwm

`func (o *UsageSummaryDateOrg) SetFargateTasksCountHwm(v int64)`

SetFargateTasksCountHwm sets FargateTasksCountHwm field to given value.

### HasFargateTasksCountHwm

`func (o *UsageSummaryDateOrg) HasFargateTasksCountHwm() bool`

HasFargateTasksCountHwm returns a boolean if a field has been set.

### GetGcpHostTop99p

`func (o *UsageSummaryDateOrg) GetGcpHostTop99p() int64`

GetGcpHostTop99p returns the GcpHostTop99p field if non-nil, zero value otherwise.

### GetGcpHostTop99pOk

`func (o *UsageSummaryDateOrg) GetGcpHostTop99pOk() (*int64, bool)`

GetGcpHostTop99pOk returns a tuple with the GcpHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpHostTop99p

`func (o *UsageSummaryDateOrg) SetGcpHostTop99p(v int64)`

SetGcpHostTop99p sets GcpHostTop99p field to given value.

### HasGcpHostTop99p

`func (o *UsageSummaryDateOrg) HasGcpHostTop99p() bool`

HasGcpHostTop99p returns a boolean if a field has been set.

### GetId

`func (o *UsageSummaryDateOrg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageSummaryDateOrg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageSummaryDateOrg) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsageSummaryDateOrg) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) GetIndexedEventsCountSum() int64`

GetIndexedEventsCountSum returns the IndexedEventsCountSum field if non-nil, zero value otherwise.

### GetIndexedEventsCountSumOk

`func (o *UsageSummaryDateOrg) GetIndexedEventsCountSumOk() (*int64, bool)`

GetIndexedEventsCountSumOk returns a tuple with the IndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) SetIndexedEventsCountSum(v int64)`

SetIndexedEventsCountSum sets IndexedEventsCountSum field to given value.

### HasIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) HasIndexedEventsCountSum() bool`

HasIndexedEventsCountSum returns a boolean if a field has been set.

### GetInfraHostTop99p

`func (o *UsageSummaryDateOrg) GetInfraHostTop99p() int64`

GetInfraHostTop99p returns the InfraHostTop99p field if non-nil, zero value otherwise.

### GetInfraHostTop99pOk

`func (o *UsageSummaryDateOrg) GetInfraHostTop99pOk() (*int64, bool)`

GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostTop99p

`func (o *UsageSummaryDateOrg) SetInfraHostTop99p(v int64)`

SetInfraHostTop99p sets InfraHostTop99p field to given value.

### HasInfraHostTop99p

`func (o *UsageSummaryDateOrg) HasInfraHostTop99p() bool`

HasInfraHostTop99p returns a boolean if a field has been set.

### GetIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) GetIngestedEventsBytesSum() int64`

GetIngestedEventsBytesSum returns the IngestedEventsBytesSum field if non-nil, zero value otherwise.

### GetIngestedEventsBytesSumOk

`func (o *UsageSummaryDateOrg) GetIngestedEventsBytesSumOk() (*int64, bool)`

GetIngestedEventsBytesSumOk returns a tuple with the IngestedEventsBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) SetIngestedEventsBytesSum(v int64)`

SetIngestedEventsBytesSum sets IngestedEventsBytesSum field to given value.

### HasIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) HasIngestedEventsBytesSum() bool`

HasIngestedEventsBytesSum returns a boolean if a field has been set.

### GetName

`func (o *UsageSummaryDateOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsageSummaryDateOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsageSummaryDateOrg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsageSummaryDateOrg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetflowIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) GetNetflowIndexedEventsCountSum() int64`

GetNetflowIndexedEventsCountSum returns the NetflowIndexedEventsCountSum field if non-nil, zero value otherwise.

### GetNetflowIndexedEventsCountSumOk

`func (o *UsageSummaryDateOrg) GetNetflowIndexedEventsCountSumOk() (*int64, bool)`

GetNetflowIndexedEventsCountSumOk returns a tuple with the NetflowIndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetflowIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) SetNetflowIndexedEventsCountSum(v int64)`

SetNetflowIndexedEventsCountSum sets NetflowIndexedEventsCountSum field to given value.

### HasNetflowIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) HasNetflowIndexedEventsCountSum() bool`

HasNetflowIndexedEventsCountSum returns a boolean if a field has been set.

### GetNpmHostTop99p

`func (o *UsageSummaryDateOrg) GetNpmHostTop99p() int64`

GetNpmHostTop99p returns the NpmHostTop99p field if non-nil, zero value otherwise.

### GetNpmHostTop99pOk

`func (o *UsageSummaryDateOrg) GetNpmHostTop99pOk() (*int64, bool)`

GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostTop99p

`func (o *UsageSummaryDateOrg) SetNpmHostTop99p(v int64)`

SetNpmHostTop99p sets NpmHostTop99p field to given value.

### HasNpmHostTop99p

`func (o *UsageSummaryDateOrg) HasNpmHostTop99p() bool`

HasNpmHostTop99p returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageSummaryDateOrg) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageSummaryDateOrg) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageSummaryDateOrg) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageSummaryDateOrg) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetRumSessionCountSum

`func (o *UsageSummaryDateOrg) GetRumSessionCountSum() int64`

GetRumSessionCountSum returns the RumSessionCountSum field if non-nil, zero value otherwise.

### GetRumSessionCountSumOk

`func (o *UsageSummaryDateOrg) GetRumSessionCountSumOk() (*int64, bool)`

GetRumSessionCountSumOk returns a tuple with the RumSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumSessionCountSum

`func (o *UsageSummaryDateOrg) SetRumSessionCountSum(v int64)`

SetRumSessionCountSum sets RumSessionCountSum field to given value.

### HasRumSessionCountSum

`func (o *UsageSummaryDateOrg) HasRumSessionCountSum() bool`

HasRumSessionCountSum returns a boolean if a field has been set.

### GetSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDateOrg) GetSyntheticsBrowserCheckCallsCountSum() int64`

GetSyntheticsBrowserCheckCallsCountSum returns the SyntheticsBrowserCheckCallsCountSum field if non-nil, zero value otherwise.

### GetSyntheticsBrowserCheckCallsCountSumOk

`func (o *UsageSummaryDateOrg) GetSyntheticsBrowserCheckCallsCountSumOk() (*int64, bool)`

GetSyntheticsBrowserCheckCallsCountSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDateOrg) SetSyntheticsBrowserCheckCallsCountSum(v int64)`

SetSyntheticsBrowserCheckCallsCountSum sets SyntheticsBrowserCheckCallsCountSum field to given value.

### HasSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDateOrg) HasSyntheticsBrowserCheckCallsCountSum() bool`

HasSyntheticsBrowserCheckCallsCountSum returns a boolean if a field has been set.

### GetSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDateOrg) GetSyntheticsCheckCallsCountSum() int64`

GetSyntheticsCheckCallsCountSum returns the SyntheticsCheckCallsCountSum field if non-nil, zero value otherwise.

### GetSyntheticsCheckCallsCountSumOk

`func (o *UsageSummaryDateOrg) GetSyntheticsCheckCallsCountSumOk() (*int64, bool)`

GetSyntheticsCheckCallsCountSumOk returns a tuple with the SyntheticsCheckCallsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDateOrg) SetSyntheticsCheckCallsCountSum(v int64)`

SetSyntheticsCheckCallsCountSum sets SyntheticsCheckCallsCountSum field to given value.

### HasSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDateOrg) HasSyntheticsCheckCallsCountSum() bool`

HasSyntheticsCheckCallsCountSum returns a boolean if a field has been set.

### GetTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) GetTraceSearchIndexedEventsCountSum() int64`

GetTraceSearchIndexedEventsCountSum returns the TraceSearchIndexedEventsCountSum field if non-nil, zero value otherwise.

### GetTraceSearchIndexedEventsCountSumOk

`func (o *UsageSummaryDateOrg) GetTraceSearchIndexedEventsCountSumOk() (*int64, bool)`

GetTraceSearchIndexedEventsCountSumOk returns a tuple with the TraceSearchIndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) SetTraceSearchIndexedEventsCountSum(v int64)`

SetTraceSearchIndexedEventsCountSum sets TraceSearchIndexedEventsCountSum field to given value.

### HasTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) HasTraceSearchIndexedEventsCountSum() bool`

HasTraceSearchIndexedEventsCountSum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


