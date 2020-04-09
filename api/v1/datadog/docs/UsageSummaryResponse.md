# UsageSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all agent hosts over all hours in the current month(s) for all orgs. | [optional] 
**ApmHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current month(s) for all orgs. | [optional] 
**AwsHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all AWS hosts over all hours in the current month(s) for all orgs. | [optional] 
**AwsLambdaFuncCount** | Pointer to **int64** | Shows the average of the number of functions that executed 1 or more times each hour in the current month(s) for all orgs. | [optional] 
**AwsLambdaInvocationsSum** | Pointer to **int64** | Shows the sum of all AWS Labmda invocations over all hours in the current month(s) for all orgs. | [optional] 
**AzureHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all Azure hosts over all hours in the current month(s) for all orgs. | [optional] 
**BillableIngestedBytesAggSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current month(s) for all orgs. | [optional] 
**ContainerAvgSum** | Pointer to **int64** | Shows the average of all distinct containers over all hours in the current month(s) for all orgs. | [optional] 
**ContainerHwmSum** | Pointer to **int64** | Shows the high watermark of all distinct containers over all hours in the current month(s) for all orgs. | [optional] 
**CustomTsSum** | Pointer to **int64** | Shows the average number of distinct custom metrics over all hours in the current month(s) for all orgs. | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | Shows the last date of usage in the current month(s) for all orgs. | [optional] 
**FargateTasksCountAvgSum** | Pointer to **int64** | Shows the average of all Fargate tasks over all hours in the current month(s) for all orgs. | [optional] 
**FargateTasksCountHwmSum** | Pointer to **int64** | Shows the high watermark of all Fargate tasks over all hours in the current month(s) for all orgs. | [optional] 
**GcpHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all GCP hosts over all hours in the current month(s) for all orgs. | [optional] 
**IndexedEventsCountAggSum** | Pointer to **int64** | Shows the sum of all log events indexed over all hours in the current month(s) for all orgs. | [optional] 
**InfraHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current month(s) for all orgs. | [optional] 
**IngestedEventsBytesAggSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current month(s) for all orgs. | [optional] 
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) | Shows the the most recent hour in the current month(s) for all orgs for which all usages were calculated. | [optional] 
**NetflowIndexedEventsCountAggSum** | Pointer to **int64** | Shows the sum of all Network flows indexed over all hours in the current month(s) for all orgs. | [optional] 
**NpmHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current month(s) for all orgs. | [optional] 
**RumSessionCountAggSum** | Pointer to **int64** | Shows the sum of all RUM Sessions over all hours in the current month(s) for all orgs. | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Shows the first date of usage in the current month(s) for all orgs. | [optional] 
**SyntheticsBrowserCheckCallsCountAggSum** | Pointer to **int64** | Shows the sum of all Synthetic browser tests over all hours in the current month(s) for all orgs. | [optional] 
**SyntheticsCheckCallsCountAggSum** | Pointer to **int64** | Shows the sum of all Synthetic API tests over all hours in the current month(s) for all orgs. | [optional] 
**TraceSearchIndexedEventsCountAggSum** | Pointer to **int64** | Shows the sum of all analyzed spans indexed over all hours in the current month(s) for all orgs. | [optional] 
**Usage** | Pointer to [**[]UsageSummaryDate**](UsageSummaryDate.md) | TODO. | [optional] 

## Methods

### NewUsageSummaryResponse

`func NewUsageSummaryResponse() *UsageSummaryResponse`

NewUsageSummaryResponse instantiates a new UsageSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSummaryResponseWithDefaults

`func NewUsageSummaryResponseWithDefaults() *UsageSummaryResponse`

NewUsageSummaryResponseWithDefaults instantiates a new UsageSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentHostTop99pSum

`func (o *UsageSummaryResponse) GetAgentHostTop99pSum() int64`

GetAgentHostTop99pSum returns the AgentHostTop99pSum field if non-nil, zero value otherwise.

### GetAgentHostTop99pSumOk

`func (o *UsageSummaryResponse) GetAgentHostTop99pSumOk() (*int64, bool)`

GetAgentHostTop99pSumOk returns a tuple with the AgentHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentHostTop99pSum

`func (o *UsageSummaryResponse) SetAgentHostTop99pSum(v int64)`

SetAgentHostTop99pSum sets AgentHostTop99pSum field to given value.

### HasAgentHostTop99pSum

`func (o *UsageSummaryResponse) HasAgentHostTop99pSum() bool`

HasAgentHostTop99pSum returns a boolean if a field has been set.

### GetApmHostTop99pSum

`func (o *UsageSummaryResponse) GetApmHostTop99pSum() int64`

GetApmHostTop99pSum returns the ApmHostTop99pSum field if non-nil, zero value otherwise.

### GetApmHostTop99pSumOk

`func (o *UsageSummaryResponse) GetApmHostTop99pSumOk() (*int64, bool)`

GetApmHostTop99pSumOk returns a tuple with the ApmHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostTop99pSum

`func (o *UsageSummaryResponse) SetApmHostTop99pSum(v int64)`

SetApmHostTop99pSum sets ApmHostTop99pSum field to given value.

### HasApmHostTop99pSum

`func (o *UsageSummaryResponse) HasApmHostTop99pSum() bool`

HasApmHostTop99pSum returns a boolean if a field has been set.

### GetAwsHostTop99pSum

`func (o *UsageSummaryResponse) GetAwsHostTop99pSum() int64`

GetAwsHostTop99pSum returns the AwsHostTop99pSum field if non-nil, zero value otherwise.

### GetAwsHostTop99pSumOk

`func (o *UsageSummaryResponse) GetAwsHostTop99pSumOk() (*int64, bool)`

GetAwsHostTop99pSumOk returns a tuple with the AwsHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostTop99pSum

`func (o *UsageSummaryResponse) SetAwsHostTop99pSum(v int64)`

SetAwsHostTop99pSum sets AwsHostTop99pSum field to given value.

### HasAwsHostTop99pSum

`func (o *UsageSummaryResponse) HasAwsHostTop99pSum() bool`

HasAwsHostTop99pSum returns a boolean if a field has been set.

### GetAwsLambdaFuncCount

`func (o *UsageSummaryResponse) GetAwsLambdaFuncCount() int64`

GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field if non-nil, zero value otherwise.

### GetAwsLambdaFuncCountOk

`func (o *UsageSummaryResponse) GetAwsLambdaFuncCountOk() (*int64, bool)`

GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaFuncCount

`func (o *UsageSummaryResponse) SetAwsLambdaFuncCount(v int64)`

SetAwsLambdaFuncCount sets AwsLambdaFuncCount field to given value.

### HasAwsLambdaFuncCount

`func (o *UsageSummaryResponse) HasAwsLambdaFuncCount() bool`

HasAwsLambdaFuncCount returns a boolean if a field has been set.

### GetAwsLambdaInvocationsSum

`func (o *UsageSummaryResponse) GetAwsLambdaInvocationsSum() int64`

GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field if non-nil, zero value otherwise.

### GetAwsLambdaInvocationsSumOk

`func (o *UsageSummaryResponse) GetAwsLambdaInvocationsSumOk() (*int64, bool)`

GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaInvocationsSum

`func (o *UsageSummaryResponse) SetAwsLambdaInvocationsSum(v int64)`

SetAwsLambdaInvocationsSum sets AwsLambdaInvocationsSum field to given value.

### HasAwsLambdaInvocationsSum

`func (o *UsageSummaryResponse) HasAwsLambdaInvocationsSum() bool`

HasAwsLambdaInvocationsSum returns a boolean if a field has been set.

### GetAzureHostTop99pSum

`func (o *UsageSummaryResponse) GetAzureHostTop99pSum() int64`

GetAzureHostTop99pSum returns the AzureHostTop99pSum field if non-nil, zero value otherwise.

### GetAzureHostTop99pSumOk

`func (o *UsageSummaryResponse) GetAzureHostTop99pSumOk() (*int64, bool)`

GetAzureHostTop99pSumOk returns a tuple with the AzureHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureHostTop99pSum

`func (o *UsageSummaryResponse) SetAzureHostTop99pSum(v int64)`

SetAzureHostTop99pSum sets AzureHostTop99pSum field to given value.

### HasAzureHostTop99pSum

`func (o *UsageSummaryResponse) HasAzureHostTop99pSum() bool`

HasAzureHostTop99pSum returns a boolean if a field has been set.

### GetBillableIngestedBytesAggSum

`func (o *UsageSummaryResponse) GetBillableIngestedBytesAggSum() int64`

GetBillableIngestedBytesAggSum returns the BillableIngestedBytesAggSum field if non-nil, zero value otherwise.

### GetBillableIngestedBytesAggSumOk

`func (o *UsageSummaryResponse) GetBillableIngestedBytesAggSumOk() (*int64, bool)`

GetBillableIngestedBytesAggSumOk returns a tuple with the BillableIngestedBytesAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableIngestedBytesAggSum

`func (o *UsageSummaryResponse) SetBillableIngestedBytesAggSum(v int64)`

SetBillableIngestedBytesAggSum sets BillableIngestedBytesAggSum field to given value.

### HasBillableIngestedBytesAggSum

`func (o *UsageSummaryResponse) HasBillableIngestedBytesAggSum() bool`

HasBillableIngestedBytesAggSum returns a boolean if a field has been set.

### GetContainerAvgSum

`func (o *UsageSummaryResponse) GetContainerAvgSum() int64`

GetContainerAvgSum returns the ContainerAvgSum field if non-nil, zero value otherwise.

### GetContainerAvgSumOk

`func (o *UsageSummaryResponse) GetContainerAvgSumOk() (*int64, bool)`

GetContainerAvgSumOk returns a tuple with the ContainerAvgSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerAvgSum

`func (o *UsageSummaryResponse) SetContainerAvgSum(v int64)`

SetContainerAvgSum sets ContainerAvgSum field to given value.

### HasContainerAvgSum

`func (o *UsageSummaryResponse) HasContainerAvgSum() bool`

HasContainerAvgSum returns a boolean if a field has been set.

### GetContainerHwmSum

`func (o *UsageSummaryResponse) GetContainerHwmSum() int64`

GetContainerHwmSum returns the ContainerHwmSum field if non-nil, zero value otherwise.

### GetContainerHwmSumOk

`func (o *UsageSummaryResponse) GetContainerHwmSumOk() (*int64, bool)`

GetContainerHwmSumOk returns a tuple with the ContainerHwmSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHwmSum

`func (o *UsageSummaryResponse) SetContainerHwmSum(v int64)`

SetContainerHwmSum sets ContainerHwmSum field to given value.

### HasContainerHwmSum

`func (o *UsageSummaryResponse) HasContainerHwmSum() bool`

HasContainerHwmSum returns a boolean if a field has been set.

### GetCustomTsSum

`func (o *UsageSummaryResponse) GetCustomTsSum() int64`

GetCustomTsSum returns the CustomTsSum field if non-nil, zero value otherwise.

### GetCustomTsSumOk

`func (o *UsageSummaryResponse) GetCustomTsSumOk() (*int64, bool)`

GetCustomTsSumOk returns a tuple with the CustomTsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTsSum

`func (o *UsageSummaryResponse) SetCustomTsSum(v int64)`

SetCustomTsSum sets CustomTsSum field to given value.

### HasCustomTsSum

`func (o *UsageSummaryResponse) HasCustomTsSum() bool`

HasCustomTsSum returns a boolean if a field has been set.

### GetEndDate

`func (o *UsageSummaryResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageSummaryResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageSummaryResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UsageSummaryResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFargateTasksCountAvgSum

`func (o *UsageSummaryResponse) GetFargateTasksCountAvgSum() int64`

GetFargateTasksCountAvgSum returns the FargateTasksCountAvgSum field if non-nil, zero value otherwise.

### GetFargateTasksCountAvgSumOk

`func (o *UsageSummaryResponse) GetFargateTasksCountAvgSumOk() (*int64, bool)`

GetFargateTasksCountAvgSumOk returns a tuple with the FargateTasksCountAvgSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountAvgSum

`func (o *UsageSummaryResponse) SetFargateTasksCountAvgSum(v int64)`

SetFargateTasksCountAvgSum sets FargateTasksCountAvgSum field to given value.

### HasFargateTasksCountAvgSum

`func (o *UsageSummaryResponse) HasFargateTasksCountAvgSum() bool`

HasFargateTasksCountAvgSum returns a boolean if a field has been set.

### GetFargateTasksCountHwmSum

`func (o *UsageSummaryResponse) GetFargateTasksCountHwmSum() int64`

GetFargateTasksCountHwmSum returns the FargateTasksCountHwmSum field if non-nil, zero value otherwise.

### GetFargateTasksCountHwmSumOk

`func (o *UsageSummaryResponse) GetFargateTasksCountHwmSumOk() (*int64, bool)`

GetFargateTasksCountHwmSumOk returns a tuple with the FargateTasksCountHwmSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountHwmSum

`func (o *UsageSummaryResponse) SetFargateTasksCountHwmSum(v int64)`

SetFargateTasksCountHwmSum sets FargateTasksCountHwmSum field to given value.

### HasFargateTasksCountHwmSum

`func (o *UsageSummaryResponse) HasFargateTasksCountHwmSum() bool`

HasFargateTasksCountHwmSum returns a boolean if a field has been set.

### GetGcpHostTop99pSum

`func (o *UsageSummaryResponse) GetGcpHostTop99pSum() int64`

GetGcpHostTop99pSum returns the GcpHostTop99pSum field if non-nil, zero value otherwise.

### GetGcpHostTop99pSumOk

`func (o *UsageSummaryResponse) GetGcpHostTop99pSumOk() (*int64, bool)`

GetGcpHostTop99pSumOk returns a tuple with the GcpHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpHostTop99pSum

`func (o *UsageSummaryResponse) SetGcpHostTop99pSum(v int64)`

SetGcpHostTop99pSum sets GcpHostTop99pSum field to given value.

### HasGcpHostTop99pSum

`func (o *UsageSummaryResponse) HasGcpHostTop99pSum() bool`

HasGcpHostTop99pSum returns a boolean if a field has been set.

### GetIndexedEventsCountAggSum

`func (o *UsageSummaryResponse) GetIndexedEventsCountAggSum() int64`

GetIndexedEventsCountAggSum returns the IndexedEventsCountAggSum field if non-nil, zero value otherwise.

### GetIndexedEventsCountAggSumOk

`func (o *UsageSummaryResponse) GetIndexedEventsCountAggSumOk() (*int64, bool)`

GetIndexedEventsCountAggSumOk returns a tuple with the IndexedEventsCountAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventsCountAggSum

`func (o *UsageSummaryResponse) SetIndexedEventsCountAggSum(v int64)`

SetIndexedEventsCountAggSum sets IndexedEventsCountAggSum field to given value.

### HasIndexedEventsCountAggSum

`func (o *UsageSummaryResponse) HasIndexedEventsCountAggSum() bool`

HasIndexedEventsCountAggSum returns a boolean if a field has been set.

### GetInfraHostTop99pSum

`func (o *UsageSummaryResponse) GetInfraHostTop99pSum() int64`

GetInfraHostTop99pSum returns the InfraHostTop99pSum field if non-nil, zero value otherwise.

### GetInfraHostTop99pSumOk

`func (o *UsageSummaryResponse) GetInfraHostTop99pSumOk() (*int64, bool)`

GetInfraHostTop99pSumOk returns a tuple with the InfraHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostTop99pSum

`func (o *UsageSummaryResponse) SetInfraHostTop99pSum(v int64)`

SetInfraHostTop99pSum sets InfraHostTop99pSum field to given value.

### HasInfraHostTop99pSum

`func (o *UsageSummaryResponse) HasInfraHostTop99pSum() bool`

HasInfraHostTop99pSum returns a boolean if a field has been set.

### GetIngestedEventsBytesAggSum

`func (o *UsageSummaryResponse) GetIngestedEventsBytesAggSum() int64`

GetIngestedEventsBytesAggSum returns the IngestedEventsBytesAggSum field if non-nil, zero value otherwise.

### GetIngestedEventsBytesAggSumOk

`func (o *UsageSummaryResponse) GetIngestedEventsBytesAggSumOk() (*int64, bool)`

GetIngestedEventsBytesAggSumOk returns a tuple with the IngestedEventsBytesAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEventsBytesAggSum

`func (o *UsageSummaryResponse) SetIngestedEventsBytesAggSum(v int64)`

SetIngestedEventsBytesAggSum sets IngestedEventsBytesAggSum field to given value.

### HasIngestedEventsBytesAggSum

`func (o *UsageSummaryResponse) HasIngestedEventsBytesAggSum() bool`

HasIngestedEventsBytesAggSum returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UsageSummaryResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UsageSummaryResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UsageSummaryResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UsageSummaryResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetNetflowIndexedEventsCountAggSum

`func (o *UsageSummaryResponse) GetNetflowIndexedEventsCountAggSum() int64`

GetNetflowIndexedEventsCountAggSum returns the NetflowIndexedEventsCountAggSum field if non-nil, zero value otherwise.

### GetNetflowIndexedEventsCountAggSumOk

`func (o *UsageSummaryResponse) GetNetflowIndexedEventsCountAggSumOk() (*int64, bool)`

GetNetflowIndexedEventsCountAggSumOk returns a tuple with the NetflowIndexedEventsCountAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetflowIndexedEventsCountAggSum

`func (o *UsageSummaryResponse) SetNetflowIndexedEventsCountAggSum(v int64)`

SetNetflowIndexedEventsCountAggSum sets NetflowIndexedEventsCountAggSum field to given value.

### HasNetflowIndexedEventsCountAggSum

`func (o *UsageSummaryResponse) HasNetflowIndexedEventsCountAggSum() bool`

HasNetflowIndexedEventsCountAggSum returns a boolean if a field has been set.

### GetNpmHostTop99pSum

`func (o *UsageSummaryResponse) GetNpmHostTop99pSum() int64`

GetNpmHostTop99pSum returns the NpmHostTop99pSum field if non-nil, zero value otherwise.

### GetNpmHostTop99pSumOk

`func (o *UsageSummaryResponse) GetNpmHostTop99pSumOk() (*int64, bool)`

GetNpmHostTop99pSumOk returns a tuple with the NpmHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostTop99pSum

`func (o *UsageSummaryResponse) SetNpmHostTop99pSum(v int64)`

SetNpmHostTop99pSum sets NpmHostTop99pSum field to given value.

### HasNpmHostTop99pSum

`func (o *UsageSummaryResponse) HasNpmHostTop99pSum() bool`

HasNpmHostTop99pSum returns a boolean if a field has been set.

### GetRumSessionCountAggSum

`func (o *UsageSummaryResponse) GetRumSessionCountAggSum() int64`

GetRumSessionCountAggSum returns the RumSessionCountAggSum field if non-nil, zero value otherwise.

### GetRumSessionCountAggSumOk

`func (o *UsageSummaryResponse) GetRumSessionCountAggSumOk() (*int64, bool)`

GetRumSessionCountAggSumOk returns a tuple with the RumSessionCountAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumSessionCountAggSum

`func (o *UsageSummaryResponse) SetRumSessionCountAggSum(v int64)`

SetRumSessionCountAggSum sets RumSessionCountAggSum field to given value.

### HasRumSessionCountAggSum

`func (o *UsageSummaryResponse) HasRumSessionCountAggSum() bool`

HasRumSessionCountAggSum returns a boolean if a field has been set.

### GetStartDate

`func (o *UsageSummaryResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageSummaryResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageSummaryResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UsageSummaryResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSyntheticsBrowserCheckCallsCountAggSum

`func (o *UsageSummaryResponse) GetSyntheticsBrowserCheckCallsCountAggSum() int64`

GetSyntheticsBrowserCheckCallsCountAggSum returns the SyntheticsBrowserCheckCallsCountAggSum field if non-nil, zero value otherwise.

### GetSyntheticsBrowserCheckCallsCountAggSumOk

`func (o *UsageSummaryResponse) GetSyntheticsBrowserCheckCallsCountAggSumOk() (*int64, bool)`

GetSyntheticsBrowserCheckCallsCountAggSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsBrowserCheckCallsCountAggSum

`func (o *UsageSummaryResponse) SetSyntheticsBrowserCheckCallsCountAggSum(v int64)`

SetSyntheticsBrowserCheckCallsCountAggSum sets SyntheticsBrowserCheckCallsCountAggSum field to given value.

### HasSyntheticsBrowserCheckCallsCountAggSum

`func (o *UsageSummaryResponse) HasSyntheticsBrowserCheckCallsCountAggSum() bool`

HasSyntheticsBrowserCheckCallsCountAggSum returns a boolean if a field has been set.

### GetSyntheticsCheckCallsCountAggSum

`func (o *UsageSummaryResponse) GetSyntheticsCheckCallsCountAggSum() int64`

GetSyntheticsCheckCallsCountAggSum returns the SyntheticsCheckCallsCountAggSum field if non-nil, zero value otherwise.

### GetSyntheticsCheckCallsCountAggSumOk

`func (o *UsageSummaryResponse) GetSyntheticsCheckCallsCountAggSumOk() (*int64, bool)`

GetSyntheticsCheckCallsCountAggSumOk returns a tuple with the SyntheticsCheckCallsCountAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsCheckCallsCountAggSum

`func (o *UsageSummaryResponse) SetSyntheticsCheckCallsCountAggSum(v int64)`

SetSyntheticsCheckCallsCountAggSum sets SyntheticsCheckCallsCountAggSum field to given value.

### HasSyntheticsCheckCallsCountAggSum

`func (o *UsageSummaryResponse) HasSyntheticsCheckCallsCountAggSum() bool`

HasSyntheticsCheckCallsCountAggSum returns a boolean if a field has been set.

### GetTraceSearchIndexedEventsCountAggSum

`func (o *UsageSummaryResponse) GetTraceSearchIndexedEventsCountAggSum() int64`

GetTraceSearchIndexedEventsCountAggSum returns the TraceSearchIndexedEventsCountAggSum field if non-nil, zero value otherwise.

### GetTraceSearchIndexedEventsCountAggSumOk

`func (o *UsageSummaryResponse) GetTraceSearchIndexedEventsCountAggSumOk() (*int64, bool)`

GetTraceSearchIndexedEventsCountAggSumOk returns a tuple with the TraceSearchIndexedEventsCountAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceSearchIndexedEventsCountAggSum

`func (o *UsageSummaryResponse) SetTraceSearchIndexedEventsCountAggSum(v int64)`

SetTraceSearchIndexedEventsCountAggSum sets TraceSearchIndexedEventsCountAggSum field to given value.

### HasTraceSearchIndexedEventsCountAggSum

`func (o *UsageSummaryResponse) HasTraceSearchIndexedEventsCountAggSum() bool`

HasTraceSearchIndexedEventsCountAggSum returns a boolean if a field has been set.

### GetUsage

`func (o *UsageSummaryResponse) GetUsage() []UsageSummaryDate`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageSummaryResponse) GetUsageOk() (*[]UsageSummaryDate, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageSummaryResponse) SetUsage(v []UsageSummaryDate)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageSummaryResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


