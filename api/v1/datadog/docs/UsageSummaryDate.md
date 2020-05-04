# UsageSummaryDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all agent hosts over all hours in the current date for all organizations. | [optional] 
**ApmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current date for all organizations. | [optional] 
**AwsHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all AWS hosts over all hours in the current date for all organizations. | [optional] 
**AwsLambdaFuncCount** | Pointer to **int64** | Shows the average of the number of functions that executed 1 or more times each hour in the current date for all organizations. | [optional] 
**AwsLambdaInvocationsSum** | Pointer to **int64** | Shows the sum of all AWS Lambda invocations over all hours in the current date for all organizations. | [optional] 
**BillableIngestedBytesSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current date for all organizations. | [optional] 
**ContainerAvg** | Pointer to **int64** | Shows the average of all distinct containers over all hours in the current date for all organizations. | [optional] 
**ContainerHwm** | Pointer to **int64** | Shows the high watermark of all distinct containers over all hours in the current date for all organizations. | [optional] 
**CustomTsAvg** | Pointer to **int64** | Shows the average number of distinct custom metrics over all hours in the current date for all organizations. | [optional] 
**Date** | Pointer to [**time.Time**](time.Time.md) | The date for the usage. | [optional] 
**FargateTasksCountAvg** | Pointer to **int64** | Shows the high watermark of all Fargate tasks over all hours in the current date for all organizations. | [optional] 
**FargateTasksCountHwm** | Pointer to **int64** | Shows the average of all Fargate tasks over all hours in the current date for all organizations. | [optional] 
**GcpHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all GCP hosts over all hours in the current date for all organizations. | [optional] 
**IndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all log events indexed over all hours in the current date for all organizations. | [optional] 
**InfraHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for all organizations. | [optional] 
**IngestedEventsBytesSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current date for all organizations. | [optional] 
**NetflowIndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all Network flows indexed over all hours in the current date for all organizations. | [optional] 
**NpmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for all organizations. | [optional] 
**Orgs** | Pointer to [**[]UsageSummaryDateOrg**](UsageSummaryDateOrg.md) | Organizations associated with a user. | [optional] 
**RumSessionCountSum** | Pointer to **int64** | Shows the sum of all RUM Sessions over all hours in the current date for all organizations | [optional] 
**SyntheticsBrowserCheckCallsCountSum** | Pointer to **int64** | Shows the sum of all Synthetic browser tests over all hours in the current date for all organizations. | [optional] 
**SyntheticsCheckCallsCountSum** | Pointer to **int64** | Shows the sum of all Synthetic API tests over all hours in the current date for all organizations. | [optional] 
**TraceSearchIndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all analyzed spans indexed over all hours in the current date for all organizations. | [optional] 

## Methods

### NewUsageSummaryDate

`func NewUsageSummaryDate() *UsageSummaryDate`

NewUsageSummaryDate instantiates a new UsageSummaryDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSummaryDateWithDefaults

`func NewUsageSummaryDateWithDefaults() *UsageSummaryDate`

NewUsageSummaryDateWithDefaults instantiates a new UsageSummaryDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentHostTop99p

`func (o *UsageSummaryDate) GetAgentHostTop99p() int64`

GetAgentHostTop99p returns the AgentHostTop99p field if non-nil, zero value otherwise.

### GetAgentHostTop99pOk

`func (o *UsageSummaryDate) GetAgentHostTop99pOk() (*int64, bool)`

GetAgentHostTop99pOk returns a tuple with the AgentHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentHostTop99p

`func (o *UsageSummaryDate) SetAgentHostTop99p(v int64)`

SetAgentHostTop99p sets AgentHostTop99p field to given value.

### HasAgentHostTop99p

`func (o *UsageSummaryDate) HasAgentHostTop99p() bool`

HasAgentHostTop99p returns a boolean if a field has been set.

### GetApmHostTop99p

`func (o *UsageSummaryDate) GetApmHostTop99p() int64`

GetApmHostTop99p returns the ApmHostTop99p field if non-nil, zero value otherwise.

### GetApmHostTop99pOk

`func (o *UsageSummaryDate) GetApmHostTop99pOk() (*int64, bool)`

GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostTop99p

`func (o *UsageSummaryDate) SetApmHostTop99p(v int64)`

SetApmHostTop99p sets ApmHostTop99p field to given value.

### HasApmHostTop99p

`func (o *UsageSummaryDate) HasApmHostTop99p() bool`

HasApmHostTop99p returns a boolean if a field has been set.

### GetAwsHostTop99p

`func (o *UsageSummaryDate) GetAwsHostTop99p() int64`

GetAwsHostTop99p returns the AwsHostTop99p field if non-nil, zero value otherwise.

### GetAwsHostTop99pOk

`func (o *UsageSummaryDate) GetAwsHostTop99pOk() (*int64, bool)`

GetAwsHostTop99pOk returns a tuple with the AwsHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostTop99p

`func (o *UsageSummaryDate) SetAwsHostTop99p(v int64)`

SetAwsHostTop99p sets AwsHostTop99p field to given value.

### HasAwsHostTop99p

`func (o *UsageSummaryDate) HasAwsHostTop99p() bool`

HasAwsHostTop99p returns a boolean if a field has been set.

### GetAwsLambdaFuncCount

`func (o *UsageSummaryDate) GetAwsLambdaFuncCount() int64`

GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field if non-nil, zero value otherwise.

### GetAwsLambdaFuncCountOk

`func (o *UsageSummaryDate) GetAwsLambdaFuncCountOk() (*int64, bool)`

GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaFuncCount

`func (o *UsageSummaryDate) SetAwsLambdaFuncCount(v int64)`

SetAwsLambdaFuncCount sets AwsLambdaFuncCount field to given value.

### HasAwsLambdaFuncCount

`func (o *UsageSummaryDate) HasAwsLambdaFuncCount() bool`

HasAwsLambdaFuncCount returns a boolean if a field has been set.

### GetAwsLambdaInvocationsSum

`func (o *UsageSummaryDate) GetAwsLambdaInvocationsSum() int64`

GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field if non-nil, zero value otherwise.

### GetAwsLambdaInvocationsSumOk

`func (o *UsageSummaryDate) GetAwsLambdaInvocationsSumOk() (*int64, bool)`

GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaInvocationsSum

`func (o *UsageSummaryDate) SetAwsLambdaInvocationsSum(v int64)`

SetAwsLambdaInvocationsSum sets AwsLambdaInvocationsSum field to given value.

### HasAwsLambdaInvocationsSum

`func (o *UsageSummaryDate) HasAwsLambdaInvocationsSum() bool`

HasAwsLambdaInvocationsSum returns a boolean if a field has been set.

### GetBillableIngestedBytesSum

`func (o *UsageSummaryDate) GetBillableIngestedBytesSum() int64`

GetBillableIngestedBytesSum returns the BillableIngestedBytesSum field if non-nil, zero value otherwise.

### GetBillableIngestedBytesSumOk

`func (o *UsageSummaryDate) GetBillableIngestedBytesSumOk() (*int64, bool)`

GetBillableIngestedBytesSumOk returns a tuple with the BillableIngestedBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableIngestedBytesSum

`func (o *UsageSummaryDate) SetBillableIngestedBytesSum(v int64)`

SetBillableIngestedBytesSum sets BillableIngestedBytesSum field to given value.

### HasBillableIngestedBytesSum

`func (o *UsageSummaryDate) HasBillableIngestedBytesSum() bool`

HasBillableIngestedBytesSum returns a boolean if a field has been set.

### GetContainerAvg

`func (o *UsageSummaryDate) GetContainerAvg() int64`

GetContainerAvg returns the ContainerAvg field if non-nil, zero value otherwise.

### GetContainerAvgOk

`func (o *UsageSummaryDate) GetContainerAvgOk() (*int64, bool)`

GetContainerAvgOk returns a tuple with the ContainerAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerAvg

`func (o *UsageSummaryDate) SetContainerAvg(v int64)`

SetContainerAvg sets ContainerAvg field to given value.

### HasContainerAvg

`func (o *UsageSummaryDate) HasContainerAvg() bool`

HasContainerAvg returns a boolean if a field has been set.

### GetContainerHwm

`func (o *UsageSummaryDate) GetContainerHwm() int64`

GetContainerHwm returns the ContainerHwm field if non-nil, zero value otherwise.

### GetContainerHwmOk

`func (o *UsageSummaryDate) GetContainerHwmOk() (*int64, bool)`

GetContainerHwmOk returns a tuple with the ContainerHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHwm

`func (o *UsageSummaryDate) SetContainerHwm(v int64)`

SetContainerHwm sets ContainerHwm field to given value.

### HasContainerHwm

`func (o *UsageSummaryDate) HasContainerHwm() bool`

HasContainerHwm returns a boolean if a field has been set.

### GetCustomTsAvg

`func (o *UsageSummaryDate) GetCustomTsAvg() int64`

GetCustomTsAvg returns the CustomTsAvg field if non-nil, zero value otherwise.

### GetCustomTsAvgOk

`func (o *UsageSummaryDate) GetCustomTsAvgOk() (*int64, bool)`

GetCustomTsAvgOk returns a tuple with the CustomTsAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTsAvg

`func (o *UsageSummaryDate) SetCustomTsAvg(v int64)`

SetCustomTsAvg sets CustomTsAvg field to given value.

### HasCustomTsAvg

`func (o *UsageSummaryDate) HasCustomTsAvg() bool`

HasCustomTsAvg returns a boolean if a field has been set.

### GetDate

`func (o *UsageSummaryDate) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UsageSummaryDate) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UsageSummaryDate) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *UsageSummaryDate) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFargateTasksCountAvg

`func (o *UsageSummaryDate) GetFargateTasksCountAvg() int64`

GetFargateTasksCountAvg returns the FargateTasksCountAvg field if non-nil, zero value otherwise.

### GetFargateTasksCountAvgOk

`func (o *UsageSummaryDate) GetFargateTasksCountAvgOk() (*int64, bool)`

GetFargateTasksCountAvgOk returns a tuple with the FargateTasksCountAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountAvg

`func (o *UsageSummaryDate) SetFargateTasksCountAvg(v int64)`

SetFargateTasksCountAvg sets FargateTasksCountAvg field to given value.

### HasFargateTasksCountAvg

`func (o *UsageSummaryDate) HasFargateTasksCountAvg() bool`

HasFargateTasksCountAvg returns a boolean if a field has been set.

### GetFargateTasksCountHwm

`func (o *UsageSummaryDate) GetFargateTasksCountHwm() int64`

GetFargateTasksCountHwm returns the FargateTasksCountHwm field if non-nil, zero value otherwise.

### GetFargateTasksCountHwmOk

`func (o *UsageSummaryDate) GetFargateTasksCountHwmOk() (*int64, bool)`

GetFargateTasksCountHwmOk returns a tuple with the FargateTasksCountHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountHwm

`func (o *UsageSummaryDate) SetFargateTasksCountHwm(v int64)`

SetFargateTasksCountHwm sets FargateTasksCountHwm field to given value.

### HasFargateTasksCountHwm

`func (o *UsageSummaryDate) HasFargateTasksCountHwm() bool`

HasFargateTasksCountHwm returns a boolean if a field has been set.

### GetGcpHostTop99p

`func (o *UsageSummaryDate) GetGcpHostTop99p() int64`

GetGcpHostTop99p returns the GcpHostTop99p field if non-nil, zero value otherwise.

### GetGcpHostTop99pOk

`func (o *UsageSummaryDate) GetGcpHostTop99pOk() (*int64, bool)`

GetGcpHostTop99pOk returns a tuple with the GcpHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpHostTop99p

`func (o *UsageSummaryDate) SetGcpHostTop99p(v int64)`

SetGcpHostTop99p sets GcpHostTop99p field to given value.

### HasGcpHostTop99p

`func (o *UsageSummaryDate) HasGcpHostTop99p() bool`

HasGcpHostTop99p returns a boolean if a field has been set.

### GetIndexedEventsCountSum

`func (o *UsageSummaryDate) GetIndexedEventsCountSum() int64`

GetIndexedEventsCountSum returns the IndexedEventsCountSum field if non-nil, zero value otherwise.

### GetIndexedEventsCountSumOk

`func (o *UsageSummaryDate) GetIndexedEventsCountSumOk() (*int64, bool)`

GetIndexedEventsCountSumOk returns a tuple with the IndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventsCountSum

`func (o *UsageSummaryDate) SetIndexedEventsCountSum(v int64)`

SetIndexedEventsCountSum sets IndexedEventsCountSum field to given value.

### HasIndexedEventsCountSum

`func (o *UsageSummaryDate) HasIndexedEventsCountSum() bool`

HasIndexedEventsCountSum returns a boolean if a field has been set.

### GetInfraHostTop99p

`func (o *UsageSummaryDate) GetInfraHostTop99p() int64`

GetInfraHostTop99p returns the InfraHostTop99p field if non-nil, zero value otherwise.

### GetInfraHostTop99pOk

`func (o *UsageSummaryDate) GetInfraHostTop99pOk() (*int64, bool)`

GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostTop99p

`func (o *UsageSummaryDate) SetInfraHostTop99p(v int64)`

SetInfraHostTop99p sets InfraHostTop99p field to given value.

### HasInfraHostTop99p

`func (o *UsageSummaryDate) HasInfraHostTop99p() bool`

HasInfraHostTop99p returns a boolean if a field has been set.

### GetIngestedEventsBytesSum

`func (o *UsageSummaryDate) GetIngestedEventsBytesSum() int64`

GetIngestedEventsBytesSum returns the IngestedEventsBytesSum field if non-nil, zero value otherwise.

### GetIngestedEventsBytesSumOk

`func (o *UsageSummaryDate) GetIngestedEventsBytesSumOk() (*int64, bool)`

GetIngestedEventsBytesSumOk returns a tuple with the IngestedEventsBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEventsBytesSum

`func (o *UsageSummaryDate) SetIngestedEventsBytesSum(v int64)`

SetIngestedEventsBytesSum sets IngestedEventsBytesSum field to given value.

### HasIngestedEventsBytesSum

`func (o *UsageSummaryDate) HasIngestedEventsBytesSum() bool`

HasIngestedEventsBytesSum returns a boolean if a field has been set.

### GetNetflowIndexedEventsCountSum

`func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSum() int64`

GetNetflowIndexedEventsCountSum returns the NetflowIndexedEventsCountSum field if non-nil, zero value otherwise.

### GetNetflowIndexedEventsCountSumOk

`func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSumOk() (*int64, bool)`

GetNetflowIndexedEventsCountSumOk returns a tuple with the NetflowIndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetflowIndexedEventsCountSum

`func (o *UsageSummaryDate) SetNetflowIndexedEventsCountSum(v int64)`

SetNetflowIndexedEventsCountSum sets NetflowIndexedEventsCountSum field to given value.

### HasNetflowIndexedEventsCountSum

`func (o *UsageSummaryDate) HasNetflowIndexedEventsCountSum() bool`

HasNetflowIndexedEventsCountSum returns a boolean if a field has been set.

### GetNpmHostTop99p

`func (o *UsageSummaryDate) GetNpmHostTop99p() int64`

GetNpmHostTop99p returns the NpmHostTop99p field if non-nil, zero value otherwise.

### GetNpmHostTop99pOk

`func (o *UsageSummaryDate) GetNpmHostTop99pOk() (*int64, bool)`

GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostTop99p

`func (o *UsageSummaryDate) SetNpmHostTop99p(v int64)`

SetNpmHostTop99p sets NpmHostTop99p field to given value.

### HasNpmHostTop99p

`func (o *UsageSummaryDate) HasNpmHostTop99p() bool`

HasNpmHostTop99p returns a boolean if a field has been set.

### GetOrgs

`func (o *UsageSummaryDate) GetOrgs() []UsageSummaryDateOrg`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *UsageSummaryDate) GetOrgsOk() (*[]UsageSummaryDateOrg, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *UsageSummaryDate) SetOrgs(v []UsageSummaryDateOrg)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *UsageSummaryDate) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetRumSessionCountSum

`func (o *UsageSummaryDate) GetRumSessionCountSum() int64`

GetRumSessionCountSum returns the RumSessionCountSum field if non-nil, zero value otherwise.

### GetRumSessionCountSumOk

`func (o *UsageSummaryDate) GetRumSessionCountSumOk() (*int64, bool)`

GetRumSessionCountSumOk returns a tuple with the RumSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumSessionCountSum

`func (o *UsageSummaryDate) SetRumSessionCountSum(v int64)`

SetRumSessionCountSum sets RumSessionCountSum field to given value.

### HasRumSessionCountSum

`func (o *UsageSummaryDate) HasRumSessionCountSum() bool`

HasRumSessionCountSum returns a boolean if a field has been set.

### GetSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSum() int64`

GetSyntheticsBrowserCheckCallsCountSum returns the SyntheticsBrowserCheckCallsCountSum field if non-nil, zero value otherwise.

### GetSyntheticsBrowserCheckCallsCountSumOk

`func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSumOk() (*int64, bool)`

GetSyntheticsBrowserCheckCallsCountSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDate) SetSyntheticsBrowserCheckCallsCountSum(v int64)`

SetSyntheticsBrowserCheckCallsCountSum sets SyntheticsBrowserCheckCallsCountSum field to given value.

### HasSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDate) HasSyntheticsBrowserCheckCallsCountSum() bool`

HasSyntheticsBrowserCheckCallsCountSum returns a boolean if a field has been set.

### GetSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSum() int64`

GetSyntheticsCheckCallsCountSum returns the SyntheticsCheckCallsCountSum field if non-nil, zero value otherwise.

### GetSyntheticsCheckCallsCountSumOk

`func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSumOk() (*int64, bool)`

GetSyntheticsCheckCallsCountSumOk returns a tuple with the SyntheticsCheckCallsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDate) SetSyntheticsCheckCallsCountSum(v int64)`

SetSyntheticsCheckCallsCountSum sets SyntheticsCheckCallsCountSum field to given value.

### HasSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDate) HasSyntheticsCheckCallsCountSum() bool`

HasSyntheticsCheckCallsCountSum returns a boolean if a field has been set.

### GetTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSum() int64`

GetTraceSearchIndexedEventsCountSum returns the TraceSearchIndexedEventsCountSum field if non-nil, zero value otherwise.

### GetTraceSearchIndexedEventsCountSumOk

`func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSumOk() (*int64, bool)`

GetTraceSearchIndexedEventsCountSumOk returns a tuple with the TraceSearchIndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDate) SetTraceSearchIndexedEventsCountSum(v int64)`

SetTraceSearchIndexedEventsCountSum sets TraceSearchIndexedEventsCountSum field to given value.

### HasTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDate) HasTraceSearchIndexedEventsCountSum() bool`

HasTraceSearchIndexedEventsCountSum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


