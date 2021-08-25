# UsageSummaryDateOrg

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AgentHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all agent hosts over all hours in the current date for the given org. | [optional] 
**ApmAzureAppServiceHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all Azure app services using APM over all hours in the current date for the given org. | [optional] 
**ApmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current date for the given org. | [optional] 
**AuditLogsLinesIndexedSum** | Pointer to **int64** | Shows the sum of all audit logs lines indexed over all hours in the current date for the given org. | [optional] 
**AwsHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all AWS hosts over all hours in the current date for the given org. | [optional] 
**AwsLambdaFuncCount** | Pointer to **int64** | Shows the sum of all AWS Lambda invocations over all hours in the current date for the given org. | [optional] 
**AwsLambdaInvocationsSum** | Pointer to **int64** | Shows the sum of all AWS Lambda invocations over all hours in the current date for the given org. | [optional] 
**AzureAppServiceTop99p** | Pointer to **int64** | Shows the 99th percentile of all Azure app services over all hours in the current date for the given org. | [optional] 
**BillableIngestedBytesSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current date for the given org. | [optional] 
**ContainerAvg** | Pointer to **int64** | Shows the average of all distinct containers over all hours in the current date for the given org. | [optional] 
**ContainerHwm** | Pointer to **int64** | Shows the high-water mark of all distinct containers over all hours in the current date for the given org. | [optional] 
**CspmContainerAvg** | Pointer to **int64** | Shows the average number of Cloud Security Posture Management containers over all hours in the current date for the given org. | [optional] 
**CspmContainerHwm** | Pointer to **int64** | Shows the high-water mark of Cloud Security Posture Management containers over all hours in the current date for the given org. | [optional] 
**CspmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all Cloud Security Posture Management hosts over all hours in the current date for the given org. | [optional] 
**CustomTsAvg** | Pointer to **int64** | Shows the average number of distinct custom metrics over all hours in the current date for the given org. | [optional] 
**CwsContainerCountAvg** | Pointer to **int64** | Shows the average of all distinct Cloud Workload Security containers over all hours in the current date for the given org. | [optional] 
**CwsHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all Cloud Workload Security hosts over all hours in the current date for the given org. | [optional] 
**DbmHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all Database Monitoring hosts over all hours in the current month for all organizations. | [optional] 
**DbmQueriesAggSum** | Pointer to **int64** | Shows the sum of all distinct Database Monitoring normalized queries over all hours in the current month for all organizations. | [optional] 
**FargateTasksCountAvg** | Pointer to **int64** | The average task count for Fargate. | [optional] 
**FargateTasksCountHwm** | Pointer to **int64** | Shows the high-water mark of all Fargate tasks over all hours in the current date for the given org. | [optional] 
**GcpHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all GCP hosts over all hours in the current date for the given org. | [optional] 
**HerokuHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all Heroku dynos over all hours in the current date for the given org. | [optional] 
**Id** | Pointer to **string** | The organization id. | [optional] 
**IncidentManagementMonthlyActiveUsersHwm** | Pointer to **int64** | Shows the high-water mark of incident management monthly active users over all hours in the current date for the given org. | [optional] 
**IndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all log events indexed over all hours in the current date for the given org. | [optional] 
**InfraHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for the given org. | [optional] 
**IngestedEventsBytesSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current date for the given org. | [optional] 
**IotDeviceAggSum** | Pointer to **int64** | Shows the sum of all IoT devices over all hours in the current date for the given org. | [optional] 
**IotDeviceTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all IoT devices over all hours in the current date for the given org. | [optional] 
**MobileRumSessionCountAndroidSum** | Pointer to **int64** | Shows the sum of all mobile RUM Sessions on Android over all hours in the current date for the given org. | [optional] 
**MobileRumSessionCountIosSum** | Pointer to **int64** | Shows the sum of all mobile RUM Sessions on iOS over all hours in the current date for the given org. | [optional] 
**MobileRumSessionCountSum** | Pointer to **int64** | Shows the sum of all mobile RUM Sessions over all hours in the current date for the given org. | [optional] 
**Name** | Pointer to **string** | The organization name. | [optional] 
**NetflowIndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all Network flows indexed over all hours in the current date for the given org. | [optional] 
**NpmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for the given org. | [optional] 
**OpentelemetryHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for the given org. | [optional] 
**ProfilingHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all profiled hosts over all hours in the current date for the given org. | [optional] 
**PublicId** | Pointer to **string** | The organization public id. | [optional] 
**RumSessionCountSum** | Pointer to **int64** | Shows the sum of all browser RUM Sessions over all hours in the current date for the given org. | [optional] 
**RumTotalSessionCountSum** | Pointer to **int64** | Shows the sum of RUM Sessions (browser and mobile) over all hours in the current date for the given org. | [optional] 
**SyntheticsBrowserCheckCallsCountSum** | Pointer to **int64** | Shows the sum of all Synthetic browser tests over all hours in the current date for the given org. | [optional] 
**SyntheticsCheckCallsCountSum** | Pointer to **int64** | Shows the sum of all Synthetic API tests over all hours in the current date for the given org. | [optional] 
**TraceSearchIndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all Indexed Spans indexed over all hours in the current date for the given org. | [optional] 
**TwolIngestedEventsBytesSum** | Pointer to **int64** | Shows the sum of all tracing without limits bytes ingested over all hours in the current date for the given org. | [optional] 
**VsphereHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all vSphere hosts over all hours in the current date for the given org. | [optional] 

## Methods

### NewUsageSummaryDateOrg

`func NewUsageSummaryDateOrg() *UsageSummaryDateOrg`

NewUsageSummaryDateOrg instantiates a new UsageSummaryDateOrg object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSummaryDateOrgWithDefaults

`func NewUsageSummaryDateOrgWithDefaults() *UsageSummaryDateOrg`

NewUsageSummaryDateOrgWithDefaults instantiates a new UsageSummaryDateOrg object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

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

### GetApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDateOrg) GetApmAzureAppServiceHostTop99p() int64`

GetApmAzureAppServiceHostTop99p returns the ApmAzureAppServiceHostTop99p field if non-nil, zero value otherwise.

### GetApmAzureAppServiceHostTop99pOk

`func (o *UsageSummaryDateOrg) GetApmAzureAppServiceHostTop99pOk() (*int64, bool)`

GetApmAzureAppServiceHostTop99pOk returns a tuple with the ApmAzureAppServiceHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDateOrg) SetApmAzureAppServiceHostTop99p(v int64)`

SetApmAzureAppServiceHostTop99p sets ApmAzureAppServiceHostTop99p field to given value.

### HasApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDateOrg) HasApmAzureAppServiceHostTop99p() bool`

HasApmAzureAppServiceHostTop99p returns a boolean if a field has been set.

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

### GetAuditLogsLinesIndexedSum

`func (o *UsageSummaryDateOrg) GetAuditLogsLinesIndexedSum() int64`

GetAuditLogsLinesIndexedSum returns the AuditLogsLinesIndexedSum field if non-nil, zero value otherwise.

### GetAuditLogsLinesIndexedSumOk

`func (o *UsageSummaryDateOrg) GetAuditLogsLinesIndexedSumOk() (*int64, bool)`

GetAuditLogsLinesIndexedSumOk returns a tuple with the AuditLogsLinesIndexedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsLinesIndexedSum

`func (o *UsageSummaryDateOrg) SetAuditLogsLinesIndexedSum(v int64)`

SetAuditLogsLinesIndexedSum sets AuditLogsLinesIndexedSum field to given value.

### HasAuditLogsLinesIndexedSum

`func (o *UsageSummaryDateOrg) HasAuditLogsLinesIndexedSum() bool`

HasAuditLogsLinesIndexedSum returns a boolean if a field has been set.

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

### GetAzureAppServiceTop99p

`func (o *UsageSummaryDateOrg) GetAzureAppServiceTop99p() int64`

GetAzureAppServiceTop99p returns the AzureAppServiceTop99p field if non-nil, zero value otherwise.

### GetAzureAppServiceTop99pOk

`func (o *UsageSummaryDateOrg) GetAzureAppServiceTop99pOk() (*int64, bool)`

GetAzureAppServiceTop99pOk returns a tuple with the AzureAppServiceTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppServiceTop99p

`func (o *UsageSummaryDateOrg) SetAzureAppServiceTop99p(v int64)`

SetAzureAppServiceTop99p sets AzureAppServiceTop99p field to given value.

### HasAzureAppServiceTop99p

`func (o *UsageSummaryDateOrg) HasAzureAppServiceTop99p() bool`

HasAzureAppServiceTop99p returns a boolean if a field has been set.

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

### GetCspmContainerAvg

`func (o *UsageSummaryDateOrg) GetCspmContainerAvg() int64`

GetCspmContainerAvg returns the CspmContainerAvg field if non-nil, zero value otherwise.

### GetCspmContainerAvgOk

`func (o *UsageSummaryDateOrg) GetCspmContainerAvgOk() (*int64, bool)`

GetCspmContainerAvgOk returns a tuple with the CspmContainerAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmContainerAvg

`func (o *UsageSummaryDateOrg) SetCspmContainerAvg(v int64)`

SetCspmContainerAvg sets CspmContainerAvg field to given value.

### HasCspmContainerAvg

`func (o *UsageSummaryDateOrg) HasCspmContainerAvg() bool`

HasCspmContainerAvg returns a boolean if a field has been set.

### GetCspmContainerHwm

`func (o *UsageSummaryDateOrg) GetCspmContainerHwm() int64`

GetCspmContainerHwm returns the CspmContainerHwm field if non-nil, zero value otherwise.

### GetCspmContainerHwmOk

`func (o *UsageSummaryDateOrg) GetCspmContainerHwmOk() (*int64, bool)`

GetCspmContainerHwmOk returns a tuple with the CspmContainerHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmContainerHwm

`func (o *UsageSummaryDateOrg) SetCspmContainerHwm(v int64)`

SetCspmContainerHwm sets CspmContainerHwm field to given value.

### HasCspmContainerHwm

`func (o *UsageSummaryDateOrg) HasCspmContainerHwm() bool`

HasCspmContainerHwm returns a boolean if a field has been set.

### GetCspmHostTop99p

`func (o *UsageSummaryDateOrg) GetCspmHostTop99p() int64`

GetCspmHostTop99p returns the CspmHostTop99p field if non-nil, zero value otherwise.

### GetCspmHostTop99pOk

`func (o *UsageSummaryDateOrg) GetCspmHostTop99pOk() (*int64, bool)`

GetCspmHostTop99pOk returns a tuple with the CspmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmHostTop99p

`func (o *UsageSummaryDateOrg) SetCspmHostTop99p(v int64)`

SetCspmHostTop99p sets CspmHostTop99p field to given value.

### HasCspmHostTop99p

`func (o *UsageSummaryDateOrg) HasCspmHostTop99p() bool`

HasCspmHostTop99p returns a boolean if a field has been set.

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

### GetCwsContainerCountAvg

`func (o *UsageSummaryDateOrg) GetCwsContainerCountAvg() int64`

GetCwsContainerCountAvg returns the CwsContainerCountAvg field if non-nil, zero value otherwise.

### GetCwsContainerCountAvgOk

`func (o *UsageSummaryDateOrg) GetCwsContainerCountAvgOk() (*int64, bool)`

GetCwsContainerCountAvgOk returns a tuple with the CwsContainerCountAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsContainerCountAvg

`func (o *UsageSummaryDateOrg) SetCwsContainerCountAvg(v int64)`

SetCwsContainerCountAvg sets CwsContainerCountAvg field to given value.

### HasCwsContainerCountAvg

`func (o *UsageSummaryDateOrg) HasCwsContainerCountAvg() bool`

HasCwsContainerCountAvg returns a boolean if a field has been set.

### GetCwsHostTop99p

`func (o *UsageSummaryDateOrg) GetCwsHostTop99p() int64`

GetCwsHostTop99p returns the CwsHostTop99p field if non-nil, zero value otherwise.

### GetCwsHostTop99pOk

`func (o *UsageSummaryDateOrg) GetCwsHostTop99pOk() (*int64, bool)`

GetCwsHostTop99pOk returns a tuple with the CwsHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsHostTop99p

`func (o *UsageSummaryDateOrg) SetCwsHostTop99p(v int64)`

SetCwsHostTop99p sets CwsHostTop99p field to given value.

### HasCwsHostTop99p

`func (o *UsageSummaryDateOrg) HasCwsHostTop99p() bool`

HasCwsHostTop99p returns a boolean if a field has been set.

### GetDbmHostTop99pSum

`func (o *UsageSummaryDateOrg) GetDbmHostTop99pSum() int64`

GetDbmHostTop99pSum returns the DbmHostTop99pSum field if non-nil, zero value otherwise.

### GetDbmHostTop99pSumOk

`func (o *UsageSummaryDateOrg) GetDbmHostTop99pSumOk() (*int64, bool)`

GetDbmHostTop99pSumOk returns a tuple with the DbmHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmHostTop99pSum

`func (o *UsageSummaryDateOrg) SetDbmHostTop99pSum(v int64)`

SetDbmHostTop99pSum sets DbmHostTop99pSum field to given value.

### HasDbmHostTop99pSum

`func (o *UsageSummaryDateOrg) HasDbmHostTop99pSum() bool`

HasDbmHostTop99pSum returns a boolean if a field has been set.

### GetDbmQueriesAggSum

`func (o *UsageSummaryDateOrg) GetDbmQueriesAggSum() int64`

GetDbmQueriesAggSum returns the DbmQueriesAggSum field if non-nil, zero value otherwise.

### GetDbmQueriesAggSumOk

`func (o *UsageSummaryDateOrg) GetDbmQueriesAggSumOk() (*int64, bool)`

GetDbmQueriesAggSumOk returns a tuple with the DbmQueriesAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmQueriesAggSum

`func (o *UsageSummaryDateOrg) SetDbmQueriesAggSum(v int64)`

SetDbmQueriesAggSum sets DbmQueriesAggSum field to given value.

### HasDbmQueriesAggSum

`func (o *UsageSummaryDateOrg) HasDbmQueriesAggSum() bool`

HasDbmQueriesAggSum returns a boolean if a field has been set.

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

### GetHerokuHostTop99p

`func (o *UsageSummaryDateOrg) GetHerokuHostTop99p() int64`

GetHerokuHostTop99p returns the HerokuHostTop99p field if non-nil, zero value otherwise.

### GetHerokuHostTop99pOk

`func (o *UsageSummaryDateOrg) GetHerokuHostTop99pOk() (*int64, bool)`

GetHerokuHostTop99pOk returns a tuple with the HerokuHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHerokuHostTop99p

`func (o *UsageSummaryDateOrg) SetHerokuHostTop99p(v int64)`

SetHerokuHostTop99p sets HerokuHostTop99p field to given value.

### HasHerokuHostTop99p

`func (o *UsageSummaryDateOrg) HasHerokuHostTop99p() bool`

HasHerokuHostTop99p returns a boolean if a field has been set.

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

### GetIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDateOrg) GetIncidentManagementMonthlyActiveUsersHwm() int64`

GetIncidentManagementMonthlyActiveUsersHwm returns the IncidentManagementMonthlyActiveUsersHwm field if non-nil, zero value otherwise.

### GetIncidentManagementMonthlyActiveUsersHwmOk

`func (o *UsageSummaryDateOrg) GetIncidentManagementMonthlyActiveUsersHwmOk() (*int64, bool)`

GetIncidentManagementMonthlyActiveUsersHwmOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDateOrg) SetIncidentManagementMonthlyActiveUsersHwm(v int64)`

SetIncidentManagementMonthlyActiveUsersHwm sets IncidentManagementMonthlyActiveUsersHwm field to given value.

### HasIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDateOrg) HasIncidentManagementMonthlyActiveUsersHwm() bool`

HasIncidentManagementMonthlyActiveUsersHwm returns a boolean if a field has been set.

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

### GetIotDeviceAggSum

`func (o *UsageSummaryDateOrg) GetIotDeviceAggSum() int64`

GetIotDeviceAggSum returns the IotDeviceAggSum field if non-nil, zero value otherwise.

### GetIotDeviceAggSumOk

`func (o *UsageSummaryDateOrg) GetIotDeviceAggSumOk() (*int64, bool)`

GetIotDeviceAggSumOk returns a tuple with the IotDeviceAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceAggSum

`func (o *UsageSummaryDateOrg) SetIotDeviceAggSum(v int64)`

SetIotDeviceAggSum sets IotDeviceAggSum field to given value.

### HasIotDeviceAggSum

`func (o *UsageSummaryDateOrg) HasIotDeviceAggSum() bool`

HasIotDeviceAggSum returns a boolean if a field has been set.

### GetIotDeviceTop99pSum

`func (o *UsageSummaryDateOrg) GetIotDeviceTop99pSum() int64`

GetIotDeviceTop99pSum returns the IotDeviceTop99pSum field if non-nil, zero value otherwise.

### GetIotDeviceTop99pSumOk

`func (o *UsageSummaryDateOrg) GetIotDeviceTop99pSumOk() (*int64, bool)`

GetIotDeviceTop99pSumOk returns a tuple with the IotDeviceTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceTop99pSum

`func (o *UsageSummaryDateOrg) SetIotDeviceTop99pSum(v int64)`

SetIotDeviceTop99pSum sets IotDeviceTop99pSum field to given value.

### HasIotDeviceTop99pSum

`func (o *UsageSummaryDateOrg) HasIotDeviceTop99pSum() bool`

HasIotDeviceTop99pSum returns a boolean if a field has been set.

### GetMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountAndroidSum() int64`

GetMobileRumSessionCountAndroidSum returns the MobileRumSessionCountAndroidSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountAndroidSumOk

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountAndroidSumOk() (*int64, bool)`

GetMobileRumSessionCountAndroidSumOk returns a tuple with the MobileRumSessionCountAndroidSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDateOrg) SetMobileRumSessionCountAndroidSum(v int64)`

SetMobileRumSessionCountAndroidSum sets MobileRumSessionCountAndroidSum field to given value.

### HasMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDateOrg) HasMobileRumSessionCountAndroidSum() bool`

HasMobileRumSessionCountAndroidSum returns a boolean if a field has been set.

### GetMobileRumSessionCountIosSum

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountIosSum() int64`

GetMobileRumSessionCountIosSum returns the MobileRumSessionCountIosSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountIosSumOk

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountIosSumOk() (*int64, bool)`

GetMobileRumSessionCountIosSumOk returns a tuple with the MobileRumSessionCountIosSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountIosSum

`func (o *UsageSummaryDateOrg) SetMobileRumSessionCountIosSum(v int64)`

SetMobileRumSessionCountIosSum sets MobileRumSessionCountIosSum field to given value.

### HasMobileRumSessionCountIosSum

`func (o *UsageSummaryDateOrg) HasMobileRumSessionCountIosSum() bool`

HasMobileRumSessionCountIosSum returns a boolean if a field has been set.

### GetMobileRumSessionCountSum

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountSum() int64`

GetMobileRumSessionCountSum returns the MobileRumSessionCountSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountSumOk

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountSumOk() (*int64, bool)`

GetMobileRumSessionCountSumOk returns a tuple with the MobileRumSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountSum

`func (o *UsageSummaryDateOrg) SetMobileRumSessionCountSum(v int64)`

SetMobileRumSessionCountSum sets MobileRumSessionCountSum field to given value.

### HasMobileRumSessionCountSum

`func (o *UsageSummaryDateOrg) HasMobileRumSessionCountSum() bool`

HasMobileRumSessionCountSum returns a boolean if a field has been set.

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

### GetOpentelemetryHostTop99p

`func (o *UsageSummaryDateOrg) GetOpentelemetryHostTop99p() int64`

GetOpentelemetryHostTop99p returns the OpentelemetryHostTop99p field if non-nil, zero value otherwise.

### GetOpentelemetryHostTop99pOk

`func (o *UsageSummaryDateOrg) GetOpentelemetryHostTop99pOk() (*int64, bool)`

GetOpentelemetryHostTop99pOk returns a tuple with the OpentelemetryHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpentelemetryHostTop99p

`func (o *UsageSummaryDateOrg) SetOpentelemetryHostTop99p(v int64)`

SetOpentelemetryHostTop99p sets OpentelemetryHostTop99p field to given value.

### HasOpentelemetryHostTop99p

`func (o *UsageSummaryDateOrg) HasOpentelemetryHostTop99p() bool`

HasOpentelemetryHostTop99p returns a boolean if a field has been set.

### GetProfilingHostTop99p

`func (o *UsageSummaryDateOrg) GetProfilingHostTop99p() int64`

GetProfilingHostTop99p returns the ProfilingHostTop99p field if non-nil, zero value otherwise.

### GetProfilingHostTop99pOk

`func (o *UsageSummaryDateOrg) GetProfilingHostTop99pOk() (*int64, bool)`

GetProfilingHostTop99pOk returns a tuple with the ProfilingHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilingHostTop99p

`func (o *UsageSummaryDateOrg) SetProfilingHostTop99p(v int64)`

SetProfilingHostTop99p sets ProfilingHostTop99p field to given value.

### HasProfilingHostTop99p

`func (o *UsageSummaryDateOrg) HasProfilingHostTop99p() bool`

HasProfilingHostTop99p returns a boolean if a field has been set.

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

### GetRumTotalSessionCountSum

`func (o *UsageSummaryDateOrg) GetRumTotalSessionCountSum() int64`

GetRumTotalSessionCountSum returns the RumTotalSessionCountSum field if non-nil, zero value otherwise.

### GetRumTotalSessionCountSumOk

`func (o *UsageSummaryDateOrg) GetRumTotalSessionCountSumOk() (*int64, bool)`

GetRumTotalSessionCountSumOk returns a tuple with the RumTotalSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumTotalSessionCountSum

`func (o *UsageSummaryDateOrg) SetRumTotalSessionCountSum(v int64)`

SetRumTotalSessionCountSum sets RumTotalSessionCountSum field to given value.

### HasRumTotalSessionCountSum

`func (o *UsageSummaryDateOrg) HasRumTotalSessionCountSum() bool`

HasRumTotalSessionCountSum returns a boolean if a field has been set.

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

### GetTwolIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) GetTwolIngestedEventsBytesSum() int64`

GetTwolIngestedEventsBytesSum returns the TwolIngestedEventsBytesSum field if non-nil, zero value otherwise.

### GetTwolIngestedEventsBytesSumOk

`func (o *UsageSummaryDateOrg) GetTwolIngestedEventsBytesSumOk() (*int64, bool)`

GetTwolIngestedEventsBytesSumOk returns a tuple with the TwolIngestedEventsBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwolIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) SetTwolIngestedEventsBytesSum(v int64)`

SetTwolIngestedEventsBytesSum sets TwolIngestedEventsBytesSum field to given value.

### HasTwolIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) HasTwolIngestedEventsBytesSum() bool`

HasTwolIngestedEventsBytesSum returns a boolean if a field has been set.

### GetVsphereHostTop99p

`func (o *UsageSummaryDateOrg) GetVsphereHostTop99p() int64`

GetVsphereHostTop99p returns the VsphereHostTop99p field if non-nil, zero value otherwise.

### GetVsphereHostTop99pOk

`func (o *UsageSummaryDateOrg) GetVsphereHostTop99pOk() (*int64, bool)`

GetVsphereHostTop99pOk returns a tuple with the VsphereHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereHostTop99p

`func (o *UsageSummaryDateOrg) SetVsphereHostTop99p(v int64)`

SetVsphereHostTop99p sets VsphereHostTop99p field to given value.

### HasVsphereHostTop99p

`func (o *UsageSummaryDateOrg) HasVsphereHostTop99p() bool`

HasVsphereHostTop99p returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


