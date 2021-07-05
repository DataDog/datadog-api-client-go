# UsageSummaryResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AgentHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all agent hosts over all hours in the current months for all organizations. | [optional] 
**ApmAzureAppServiceHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all Azure app services using APM over all hours in the current months all organizations. | [optional] 
**ApmHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current months for all organizations. | [optional] 
**AuditLogsLinesIndexedAggSum** | Pointer to **int64** | Shows the sum of all audit logs lines indexed over all hours in the current months for all organizations. | [optional] 
**AwsHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all AWS hosts over all hours in the current months for all organizations. | [optional] 
**AwsLambdaFuncCount** | Pointer to **int64** | Shows the average of the number of functions that executed 1 or more times each hour in the current months for all organizations. | [optional] 
**AwsLambdaInvocationsSum** | Pointer to **int64** | Shows the sum of all AWS Lambda invocations over all hours in the current months for all organizations. | [optional] 
**AzureAppServiceTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all Azure app services over all hours in the current months for all organizations. | [optional] 
**AzureHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all Azure hosts over all hours in the current months for all organizations. | [optional] 
**BillableIngestedBytesAggSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current months for all organizations. | [optional] 
**ContainerAvgSum** | Pointer to **int64** | Shows the average of all distinct containers over all hours in the current months for all organizations. | [optional] 
**ContainerHwmSum** | Pointer to **int64** | Shows the sum of the high-water marks of all distinct containers over all hours in the current months for all organizations. | [optional] 
**CspmContainerAvgSum** | Pointer to **int64** | Shows the average number of Cloud Security Posture Management containers over all hours in the current months for all organizations. | [optional] 
**CspmContainerHwmSum** | Pointer to **int64** | Shows the sum of the the high-water marks of Cloud Security Posture Management containers over all hours in the current months for all organizations. | [optional] 
**CspmHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all Cloud Security Posture Management hosts over all hours in the current months for all organizations. | [optional] 
**CustomTsSum** | Pointer to **int64** | Shows the average number of distinct custom metrics over all hours in the current months for all organizations. | [optional] 
**CwsContainersAvgSum** | Pointer to **int64** | Shows the average of all distinct Cloud Workload Security containers over all hours in the current months for all organizations. | [optional] 
**CwsHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all Cloud Workload Security hosts over all hours in the current months for all organizations. | [optional] 
**EndDate** | Pointer to **time.Time** | Shows the last date of usage in the current months for all organizations. | [optional] 
**FargateTasksCountAvgSum** | Pointer to **int64** | Shows the average of all Fargate tasks over all hours in the current months for all organizations. | [optional] 
**FargateTasksCountHwmSum** | Pointer to **int64** | Shows the sum of the high-water marks of all Fargate tasks over all hours in the current months for all organizations. | [optional] 
**GcpHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all GCP hosts over all hours in the current months for all organizations. | [optional] 
**HerokuHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all Heroku dynos over all hours in the current months for all organizations. | [optional] 
**IncidentManagementMonthlyActiveUsersHwmSum** | Pointer to **int64** | Shows sum of the the high-water marks of incident management monthly active users in the current months for all organizations. | [optional] 
**IndexedEventsCountAggSum** | Pointer to **int64** | Shows the sum of all log events indexed over all hours in the current months for all organizations. | [optional] 
**InfraHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current months for all organizations. | [optional] 
**IngestedEventsBytesAggSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current months for all organizations. | [optional] 
**IotDeviceAggSum** | Pointer to **int64** | Shows the sum of all IoT devices over all hours in the current months for all organizations. | [optional] 
**IotDeviceTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all IoT devices over all hours in the current months of all organizations. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Shows the the most recent hour in the current months for all organizations for which all usages were calculated. | [optional] 
**LiveIndexedEventsAggSum** | Pointer to **int64** | Shows the sum of all live logs indexed over all hours in the current months for all organizations (data available as of December 1, 2020). | [optional] 
**LiveIngestedBytesAggSum** | Pointer to **int64** | Shows the sum of all live logs bytes ingested over all hours in the current months for all organizations (data available as of December 1, 2020). | [optional] 
**LogsByRetention** | Pointer to [**LogsByRetention**](LogsByRetention.md) |  | [optional] 
**MobileRumSessionCountAggSum** | Pointer to **int64** | Shows the sum of all mobile RUM Sessions over all hours in the current months for all organizations. | [optional] 
**MobileRumSessionCountAndroidAggSum** | Pointer to **int64** | Shows the sum of all mobile RUM Sessions on Android over all hours in the current months for all organizations. | [optional] 
**MobileRumSessionCountIosAggSum** | Pointer to **int64** | Shows the sum of all mobile RUM Sessions on iOS over all hours in the current months for all organizations. | [optional] 
**NetflowIndexedEventsCountAggSum** | Pointer to **int64** | Shows the sum of all Network flows indexed over all hours in the current months for all organizations. | [optional] 
**NpmHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current months for all organizations. | [optional] 
**OpentelemetryHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current months for all organizations. | [optional] 
**ProfilingContainerAgentCountAvg** | Pointer to **int64** | Shows the average number of profiled containers over all hours in the current months for all organizations. | [optional] 
**ProfilingHostCountTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all profiled hosts over all hours in the current months for all organizations. | [optional] 
**RehydratedIndexedEventsAggSum** | Pointer to **int64** | Shows the sum of all rehydrated logs indexed over all hours in the current months for all organizations (data available as of December 1, 2020). | [optional] 
**RehydratedIngestedBytesAggSum** | Pointer to **int64** | Shows the sum of all rehydrated logs bytes ingested over all hours in the current months for all organizations (data available as of December 1, 2020). | [optional] 
**RumSessionCountAggSum** | Pointer to **int64** | Shows the sum of all browser RUM Sessions over all hours in the current months for all organizations. | [optional] 
**RumTotalSessionCountAggSum** | Pointer to **int64** | Shows the sum of RUM Sessions (browser and mobile) over all hours in the current months for all organizations. | [optional] 
**StartDate** | Pointer to **time.Time** | Shows the first date of usage in the current months for all organizations. | [optional] 
**SyntheticsBrowserCheckCallsCountAggSum** | Pointer to **int64** | Shows the sum of all Synthetic browser tests over all hours in the current months for all organizations. | [optional] 
**SyntheticsCheckCallsCountAggSum** | Pointer to **int64** | Shows the sum of all Synthetic API tests over all hours in the current months for all organizations. | [optional] 
**TraceSearchIndexedEventsCountAggSum** | Pointer to **int64** | Shows the sum of all Indexed Spans indexed over all hours in the current months for all organizations. | [optional] 
**TwolIngestedEventsBytesAggSum** | Pointer to **int64** | Shows the sum of all tracing without limits bytes ingested over all hours in the current months for all organizations. | [optional] 
**Usage** | Pointer to [**[]UsageSummaryDate**](UsageSummaryDate.md) | An array of objects regarding hourly usage. | [optional] 
**VsphereHostTop99pSum** | Pointer to **int64** | Shows the 99th percentile of all vSphere hosts over all hours in the current months for all organizations. | [optional] 

## Methods

### NewUsageSummaryResponse

`func NewUsageSummaryResponse() *UsageSummaryResponse`

NewUsageSummaryResponse instantiates a new UsageSummaryResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSummaryResponseWithDefaults

`func NewUsageSummaryResponseWithDefaults() *UsageSummaryResponse`

NewUsageSummaryResponseWithDefaults instantiates a new UsageSummaryResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

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

### GetApmAzureAppServiceHostTop99pSum

`func (o *UsageSummaryResponse) GetApmAzureAppServiceHostTop99pSum() int64`

GetApmAzureAppServiceHostTop99pSum returns the ApmAzureAppServiceHostTop99pSum field if non-nil, zero value otherwise.

### GetApmAzureAppServiceHostTop99pSumOk

`func (o *UsageSummaryResponse) GetApmAzureAppServiceHostTop99pSumOk() (*int64, bool)`

GetApmAzureAppServiceHostTop99pSumOk returns a tuple with the ApmAzureAppServiceHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmAzureAppServiceHostTop99pSum

`func (o *UsageSummaryResponse) SetApmAzureAppServiceHostTop99pSum(v int64)`

SetApmAzureAppServiceHostTop99pSum sets ApmAzureAppServiceHostTop99pSum field to given value.

### HasApmAzureAppServiceHostTop99pSum

`func (o *UsageSummaryResponse) HasApmAzureAppServiceHostTop99pSum() bool`

HasApmAzureAppServiceHostTop99pSum returns a boolean if a field has been set.

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

### GetAuditLogsLinesIndexedAggSum

`func (o *UsageSummaryResponse) GetAuditLogsLinesIndexedAggSum() int64`

GetAuditLogsLinesIndexedAggSum returns the AuditLogsLinesIndexedAggSum field if non-nil, zero value otherwise.

### GetAuditLogsLinesIndexedAggSumOk

`func (o *UsageSummaryResponse) GetAuditLogsLinesIndexedAggSumOk() (*int64, bool)`

GetAuditLogsLinesIndexedAggSumOk returns a tuple with the AuditLogsLinesIndexedAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsLinesIndexedAggSum

`func (o *UsageSummaryResponse) SetAuditLogsLinesIndexedAggSum(v int64)`

SetAuditLogsLinesIndexedAggSum sets AuditLogsLinesIndexedAggSum field to given value.

### HasAuditLogsLinesIndexedAggSum

`func (o *UsageSummaryResponse) HasAuditLogsLinesIndexedAggSum() bool`

HasAuditLogsLinesIndexedAggSum returns a boolean if a field has been set.

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

### GetAzureAppServiceTop99pSum

`func (o *UsageSummaryResponse) GetAzureAppServiceTop99pSum() int64`

GetAzureAppServiceTop99pSum returns the AzureAppServiceTop99pSum field if non-nil, zero value otherwise.

### GetAzureAppServiceTop99pSumOk

`func (o *UsageSummaryResponse) GetAzureAppServiceTop99pSumOk() (*int64, bool)`

GetAzureAppServiceTop99pSumOk returns a tuple with the AzureAppServiceTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppServiceTop99pSum

`func (o *UsageSummaryResponse) SetAzureAppServiceTop99pSum(v int64)`

SetAzureAppServiceTop99pSum sets AzureAppServiceTop99pSum field to given value.

### HasAzureAppServiceTop99pSum

`func (o *UsageSummaryResponse) HasAzureAppServiceTop99pSum() bool`

HasAzureAppServiceTop99pSum returns a boolean if a field has been set.

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

### GetCspmContainerAvgSum

`func (o *UsageSummaryResponse) GetCspmContainerAvgSum() int64`

GetCspmContainerAvgSum returns the CspmContainerAvgSum field if non-nil, zero value otherwise.

### GetCspmContainerAvgSumOk

`func (o *UsageSummaryResponse) GetCspmContainerAvgSumOk() (*int64, bool)`

GetCspmContainerAvgSumOk returns a tuple with the CspmContainerAvgSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmContainerAvgSum

`func (o *UsageSummaryResponse) SetCspmContainerAvgSum(v int64)`

SetCspmContainerAvgSum sets CspmContainerAvgSum field to given value.

### HasCspmContainerAvgSum

`func (o *UsageSummaryResponse) HasCspmContainerAvgSum() bool`

HasCspmContainerAvgSum returns a boolean if a field has been set.

### GetCspmContainerHwmSum

`func (o *UsageSummaryResponse) GetCspmContainerHwmSum() int64`

GetCspmContainerHwmSum returns the CspmContainerHwmSum field if non-nil, zero value otherwise.

### GetCspmContainerHwmSumOk

`func (o *UsageSummaryResponse) GetCspmContainerHwmSumOk() (*int64, bool)`

GetCspmContainerHwmSumOk returns a tuple with the CspmContainerHwmSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmContainerHwmSum

`func (o *UsageSummaryResponse) SetCspmContainerHwmSum(v int64)`

SetCspmContainerHwmSum sets CspmContainerHwmSum field to given value.

### HasCspmContainerHwmSum

`func (o *UsageSummaryResponse) HasCspmContainerHwmSum() bool`

HasCspmContainerHwmSum returns a boolean if a field has been set.

### GetCspmHostTop99pSum

`func (o *UsageSummaryResponse) GetCspmHostTop99pSum() int64`

GetCspmHostTop99pSum returns the CspmHostTop99pSum field if non-nil, zero value otherwise.

### GetCspmHostTop99pSumOk

`func (o *UsageSummaryResponse) GetCspmHostTop99pSumOk() (*int64, bool)`

GetCspmHostTop99pSumOk returns a tuple with the CspmHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmHostTop99pSum

`func (o *UsageSummaryResponse) SetCspmHostTop99pSum(v int64)`

SetCspmHostTop99pSum sets CspmHostTop99pSum field to given value.

### HasCspmHostTop99pSum

`func (o *UsageSummaryResponse) HasCspmHostTop99pSum() bool`

HasCspmHostTop99pSum returns a boolean if a field has been set.

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

### GetCwsContainersAvgSum

`func (o *UsageSummaryResponse) GetCwsContainersAvgSum() int64`

GetCwsContainersAvgSum returns the CwsContainersAvgSum field if non-nil, zero value otherwise.

### GetCwsContainersAvgSumOk

`func (o *UsageSummaryResponse) GetCwsContainersAvgSumOk() (*int64, bool)`

GetCwsContainersAvgSumOk returns a tuple with the CwsContainersAvgSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsContainersAvgSum

`func (o *UsageSummaryResponse) SetCwsContainersAvgSum(v int64)`

SetCwsContainersAvgSum sets CwsContainersAvgSum field to given value.

### HasCwsContainersAvgSum

`func (o *UsageSummaryResponse) HasCwsContainersAvgSum() bool`

HasCwsContainersAvgSum returns a boolean if a field has been set.

### GetCwsHostTop99pSum

`func (o *UsageSummaryResponse) GetCwsHostTop99pSum() int64`

GetCwsHostTop99pSum returns the CwsHostTop99pSum field if non-nil, zero value otherwise.

### GetCwsHostTop99pSumOk

`func (o *UsageSummaryResponse) GetCwsHostTop99pSumOk() (*int64, bool)`

GetCwsHostTop99pSumOk returns a tuple with the CwsHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsHostTop99pSum

`func (o *UsageSummaryResponse) SetCwsHostTop99pSum(v int64)`

SetCwsHostTop99pSum sets CwsHostTop99pSum field to given value.

### HasCwsHostTop99pSum

`func (o *UsageSummaryResponse) HasCwsHostTop99pSum() bool`

HasCwsHostTop99pSum returns a boolean if a field has been set.

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

### GetHerokuHostTop99pSum

`func (o *UsageSummaryResponse) GetHerokuHostTop99pSum() int64`

GetHerokuHostTop99pSum returns the HerokuHostTop99pSum field if non-nil, zero value otherwise.

### GetHerokuHostTop99pSumOk

`func (o *UsageSummaryResponse) GetHerokuHostTop99pSumOk() (*int64, bool)`

GetHerokuHostTop99pSumOk returns a tuple with the HerokuHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHerokuHostTop99pSum

`func (o *UsageSummaryResponse) SetHerokuHostTop99pSum(v int64)`

SetHerokuHostTop99pSum sets HerokuHostTop99pSum field to given value.

### HasHerokuHostTop99pSum

`func (o *UsageSummaryResponse) HasHerokuHostTop99pSum() bool`

HasHerokuHostTop99pSum returns a boolean if a field has been set.

### GetIncidentManagementMonthlyActiveUsersHwmSum

`func (o *UsageSummaryResponse) GetIncidentManagementMonthlyActiveUsersHwmSum() int64`

GetIncidentManagementMonthlyActiveUsersHwmSum returns the IncidentManagementMonthlyActiveUsersHwmSum field if non-nil, zero value otherwise.

### GetIncidentManagementMonthlyActiveUsersHwmSumOk

`func (o *UsageSummaryResponse) GetIncidentManagementMonthlyActiveUsersHwmSumOk() (*int64, bool)`

GetIncidentManagementMonthlyActiveUsersHwmSumOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwmSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentManagementMonthlyActiveUsersHwmSum

`func (o *UsageSummaryResponse) SetIncidentManagementMonthlyActiveUsersHwmSum(v int64)`

SetIncidentManagementMonthlyActiveUsersHwmSum sets IncidentManagementMonthlyActiveUsersHwmSum field to given value.

### HasIncidentManagementMonthlyActiveUsersHwmSum

`func (o *UsageSummaryResponse) HasIncidentManagementMonthlyActiveUsersHwmSum() bool`

HasIncidentManagementMonthlyActiveUsersHwmSum returns a boolean if a field has been set.

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

### GetIotDeviceAggSum

`func (o *UsageSummaryResponse) GetIotDeviceAggSum() int64`

GetIotDeviceAggSum returns the IotDeviceAggSum field if non-nil, zero value otherwise.

### GetIotDeviceAggSumOk

`func (o *UsageSummaryResponse) GetIotDeviceAggSumOk() (*int64, bool)`

GetIotDeviceAggSumOk returns a tuple with the IotDeviceAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceAggSum

`func (o *UsageSummaryResponse) SetIotDeviceAggSum(v int64)`

SetIotDeviceAggSum sets IotDeviceAggSum field to given value.

### HasIotDeviceAggSum

`func (o *UsageSummaryResponse) HasIotDeviceAggSum() bool`

HasIotDeviceAggSum returns a boolean if a field has been set.

### GetIotDeviceTop99pSum

`func (o *UsageSummaryResponse) GetIotDeviceTop99pSum() int64`

GetIotDeviceTop99pSum returns the IotDeviceTop99pSum field if non-nil, zero value otherwise.

### GetIotDeviceTop99pSumOk

`func (o *UsageSummaryResponse) GetIotDeviceTop99pSumOk() (*int64, bool)`

GetIotDeviceTop99pSumOk returns a tuple with the IotDeviceTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceTop99pSum

`func (o *UsageSummaryResponse) SetIotDeviceTop99pSum(v int64)`

SetIotDeviceTop99pSum sets IotDeviceTop99pSum field to given value.

### HasIotDeviceTop99pSum

`func (o *UsageSummaryResponse) HasIotDeviceTop99pSum() bool`

HasIotDeviceTop99pSum returns a boolean if a field has been set.

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

### GetLiveIndexedEventsAggSum

`func (o *UsageSummaryResponse) GetLiveIndexedEventsAggSum() int64`

GetLiveIndexedEventsAggSum returns the LiveIndexedEventsAggSum field if non-nil, zero value otherwise.

### GetLiveIndexedEventsAggSumOk

`func (o *UsageSummaryResponse) GetLiveIndexedEventsAggSumOk() (*int64, bool)`

GetLiveIndexedEventsAggSumOk returns a tuple with the LiveIndexedEventsAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveIndexedEventsAggSum

`func (o *UsageSummaryResponse) SetLiveIndexedEventsAggSum(v int64)`

SetLiveIndexedEventsAggSum sets LiveIndexedEventsAggSum field to given value.

### HasLiveIndexedEventsAggSum

`func (o *UsageSummaryResponse) HasLiveIndexedEventsAggSum() bool`

HasLiveIndexedEventsAggSum returns a boolean if a field has been set.

### GetLiveIngestedBytesAggSum

`func (o *UsageSummaryResponse) GetLiveIngestedBytesAggSum() int64`

GetLiveIngestedBytesAggSum returns the LiveIngestedBytesAggSum field if non-nil, zero value otherwise.

### GetLiveIngestedBytesAggSumOk

`func (o *UsageSummaryResponse) GetLiveIngestedBytesAggSumOk() (*int64, bool)`

GetLiveIngestedBytesAggSumOk returns a tuple with the LiveIngestedBytesAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveIngestedBytesAggSum

`func (o *UsageSummaryResponse) SetLiveIngestedBytesAggSum(v int64)`

SetLiveIngestedBytesAggSum sets LiveIngestedBytesAggSum field to given value.

### HasLiveIngestedBytesAggSum

`func (o *UsageSummaryResponse) HasLiveIngestedBytesAggSum() bool`

HasLiveIngestedBytesAggSum returns a boolean if a field has been set.

### GetLogsByRetention

`func (o *UsageSummaryResponse) GetLogsByRetention() LogsByRetention`

GetLogsByRetention returns the LogsByRetention field if non-nil, zero value otherwise.

### GetLogsByRetentionOk

`func (o *UsageSummaryResponse) GetLogsByRetentionOk() (*LogsByRetention, bool)`

GetLogsByRetentionOk returns a tuple with the LogsByRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsByRetention

`func (o *UsageSummaryResponse) SetLogsByRetention(v LogsByRetention)`

SetLogsByRetention sets LogsByRetention field to given value.

### HasLogsByRetention

`func (o *UsageSummaryResponse) HasLogsByRetention() bool`

HasLogsByRetention returns a boolean if a field has been set.

### GetMobileRumSessionCountAggSum

`func (o *UsageSummaryResponse) GetMobileRumSessionCountAggSum() int64`

GetMobileRumSessionCountAggSum returns the MobileRumSessionCountAggSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountAggSumOk

`func (o *UsageSummaryResponse) GetMobileRumSessionCountAggSumOk() (*int64, bool)`

GetMobileRumSessionCountAggSumOk returns a tuple with the MobileRumSessionCountAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountAggSum

`func (o *UsageSummaryResponse) SetMobileRumSessionCountAggSum(v int64)`

SetMobileRumSessionCountAggSum sets MobileRumSessionCountAggSum field to given value.

### HasMobileRumSessionCountAggSum

`func (o *UsageSummaryResponse) HasMobileRumSessionCountAggSum() bool`

HasMobileRumSessionCountAggSum returns a boolean if a field has been set.

### GetMobileRumSessionCountAndroidAggSum

`func (o *UsageSummaryResponse) GetMobileRumSessionCountAndroidAggSum() int64`

GetMobileRumSessionCountAndroidAggSum returns the MobileRumSessionCountAndroidAggSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountAndroidAggSumOk

`func (o *UsageSummaryResponse) GetMobileRumSessionCountAndroidAggSumOk() (*int64, bool)`

GetMobileRumSessionCountAndroidAggSumOk returns a tuple with the MobileRumSessionCountAndroidAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountAndroidAggSum

`func (o *UsageSummaryResponse) SetMobileRumSessionCountAndroidAggSum(v int64)`

SetMobileRumSessionCountAndroidAggSum sets MobileRumSessionCountAndroidAggSum field to given value.

### HasMobileRumSessionCountAndroidAggSum

`func (o *UsageSummaryResponse) HasMobileRumSessionCountAndroidAggSum() bool`

HasMobileRumSessionCountAndroidAggSum returns a boolean if a field has been set.

### GetMobileRumSessionCountIosAggSum

`func (o *UsageSummaryResponse) GetMobileRumSessionCountIosAggSum() int64`

GetMobileRumSessionCountIosAggSum returns the MobileRumSessionCountIosAggSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountIosAggSumOk

`func (o *UsageSummaryResponse) GetMobileRumSessionCountIosAggSumOk() (*int64, bool)`

GetMobileRumSessionCountIosAggSumOk returns a tuple with the MobileRumSessionCountIosAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountIosAggSum

`func (o *UsageSummaryResponse) SetMobileRumSessionCountIosAggSum(v int64)`

SetMobileRumSessionCountIosAggSum sets MobileRumSessionCountIosAggSum field to given value.

### HasMobileRumSessionCountIosAggSum

`func (o *UsageSummaryResponse) HasMobileRumSessionCountIosAggSum() bool`

HasMobileRumSessionCountIosAggSum returns a boolean if a field has been set.

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

### GetOpentelemetryHostTop99pSum

`func (o *UsageSummaryResponse) GetOpentelemetryHostTop99pSum() int64`

GetOpentelemetryHostTop99pSum returns the OpentelemetryHostTop99pSum field if non-nil, zero value otherwise.

### GetOpentelemetryHostTop99pSumOk

`func (o *UsageSummaryResponse) GetOpentelemetryHostTop99pSumOk() (*int64, bool)`

GetOpentelemetryHostTop99pSumOk returns a tuple with the OpentelemetryHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpentelemetryHostTop99pSum

`func (o *UsageSummaryResponse) SetOpentelemetryHostTop99pSum(v int64)`

SetOpentelemetryHostTop99pSum sets OpentelemetryHostTop99pSum field to given value.

### HasOpentelemetryHostTop99pSum

`func (o *UsageSummaryResponse) HasOpentelemetryHostTop99pSum() bool`

HasOpentelemetryHostTop99pSum returns a boolean if a field has been set.

### GetProfilingContainerAgentCountAvg

`func (o *UsageSummaryResponse) GetProfilingContainerAgentCountAvg() int64`

GetProfilingContainerAgentCountAvg returns the ProfilingContainerAgentCountAvg field if non-nil, zero value otherwise.

### GetProfilingContainerAgentCountAvgOk

`func (o *UsageSummaryResponse) GetProfilingContainerAgentCountAvgOk() (*int64, bool)`

GetProfilingContainerAgentCountAvgOk returns a tuple with the ProfilingContainerAgentCountAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilingContainerAgentCountAvg

`func (o *UsageSummaryResponse) SetProfilingContainerAgentCountAvg(v int64)`

SetProfilingContainerAgentCountAvg sets ProfilingContainerAgentCountAvg field to given value.

### HasProfilingContainerAgentCountAvg

`func (o *UsageSummaryResponse) HasProfilingContainerAgentCountAvg() bool`

HasProfilingContainerAgentCountAvg returns a boolean if a field has been set.

### GetProfilingHostCountTop99pSum

`func (o *UsageSummaryResponse) GetProfilingHostCountTop99pSum() int64`

GetProfilingHostCountTop99pSum returns the ProfilingHostCountTop99pSum field if non-nil, zero value otherwise.

### GetProfilingHostCountTop99pSumOk

`func (o *UsageSummaryResponse) GetProfilingHostCountTop99pSumOk() (*int64, bool)`

GetProfilingHostCountTop99pSumOk returns a tuple with the ProfilingHostCountTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilingHostCountTop99pSum

`func (o *UsageSummaryResponse) SetProfilingHostCountTop99pSum(v int64)`

SetProfilingHostCountTop99pSum sets ProfilingHostCountTop99pSum field to given value.

### HasProfilingHostCountTop99pSum

`func (o *UsageSummaryResponse) HasProfilingHostCountTop99pSum() bool`

HasProfilingHostCountTop99pSum returns a boolean if a field has been set.

### GetRehydratedIndexedEventsAggSum

`func (o *UsageSummaryResponse) GetRehydratedIndexedEventsAggSum() int64`

GetRehydratedIndexedEventsAggSum returns the RehydratedIndexedEventsAggSum field if non-nil, zero value otherwise.

### GetRehydratedIndexedEventsAggSumOk

`func (o *UsageSummaryResponse) GetRehydratedIndexedEventsAggSumOk() (*int64, bool)`

GetRehydratedIndexedEventsAggSumOk returns a tuple with the RehydratedIndexedEventsAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRehydratedIndexedEventsAggSum

`func (o *UsageSummaryResponse) SetRehydratedIndexedEventsAggSum(v int64)`

SetRehydratedIndexedEventsAggSum sets RehydratedIndexedEventsAggSum field to given value.

### HasRehydratedIndexedEventsAggSum

`func (o *UsageSummaryResponse) HasRehydratedIndexedEventsAggSum() bool`

HasRehydratedIndexedEventsAggSum returns a boolean if a field has been set.

### GetRehydratedIngestedBytesAggSum

`func (o *UsageSummaryResponse) GetRehydratedIngestedBytesAggSum() int64`

GetRehydratedIngestedBytesAggSum returns the RehydratedIngestedBytesAggSum field if non-nil, zero value otherwise.

### GetRehydratedIngestedBytesAggSumOk

`func (o *UsageSummaryResponse) GetRehydratedIngestedBytesAggSumOk() (*int64, bool)`

GetRehydratedIngestedBytesAggSumOk returns a tuple with the RehydratedIngestedBytesAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRehydratedIngestedBytesAggSum

`func (o *UsageSummaryResponse) SetRehydratedIngestedBytesAggSum(v int64)`

SetRehydratedIngestedBytesAggSum sets RehydratedIngestedBytesAggSum field to given value.

### HasRehydratedIngestedBytesAggSum

`func (o *UsageSummaryResponse) HasRehydratedIngestedBytesAggSum() bool`

HasRehydratedIngestedBytesAggSum returns a boolean if a field has been set.

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

### GetRumTotalSessionCountAggSum

`func (o *UsageSummaryResponse) GetRumTotalSessionCountAggSum() int64`

GetRumTotalSessionCountAggSum returns the RumTotalSessionCountAggSum field if non-nil, zero value otherwise.

### GetRumTotalSessionCountAggSumOk

`func (o *UsageSummaryResponse) GetRumTotalSessionCountAggSumOk() (*int64, bool)`

GetRumTotalSessionCountAggSumOk returns a tuple with the RumTotalSessionCountAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumTotalSessionCountAggSum

`func (o *UsageSummaryResponse) SetRumTotalSessionCountAggSum(v int64)`

SetRumTotalSessionCountAggSum sets RumTotalSessionCountAggSum field to given value.

### HasRumTotalSessionCountAggSum

`func (o *UsageSummaryResponse) HasRumTotalSessionCountAggSum() bool`

HasRumTotalSessionCountAggSum returns a boolean if a field has been set.

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

### GetTwolIngestedEventsBytesAggSum

`func (o *UsageSummaryResponse) GetTwolIngestedEventsBytesAggSum() int64`

GetTwolIngestedEventsBytesAggSum returns the TwolIngestedEventsBytesAggSum field if non-nil, zero value otherwise.

### GetTwolIngestedEventsBytesAggSumOk

`func (o *UsageSummaryResponse) GetTwolIngestedEventsBytesAggSumOk() (*int64, bool)`

GetTwolIngestedEventsBytesAggSumOk returns a tuple with the TwolIngestedEventsBytesAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwolIngestedEventsBytesAggSum

`func (o *UsageSummaryResponse) SetTwolIngestedEventsBytesAggSum(v int64)`

SetTwolIngestedEventsBytesAggSum sets TwolIngestedEventsBytesAggSum field to given value.

### HasTwolIngestedEventsBytesAggSum

`func (o *UsageSummaryResponse) HasTwolIngestedEventsBytesAggSum() bool`

HasTwolIngestedEventsBytesAggSum returns a boolean if a field has been set.

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

### GetVsphereHostTop99pSum

`func (o *UsageSummaryResponse) GetVsphereHostTop99pSum() int64`

GetVsphereHostTop99pSum returns the VsphereHostTop99pSum field if non-nil, zero value otherwise.

### GetVsphereHostTop99pSumOk

`func (o *UsageSummaryResponse) GetVsphereHostTop99pSumOk() (*int64, bool)`

GetVsphereHostTop99pSumOk returns a tuple with the VsphereHostTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereHostTop99pSum

`func (o *UsageSummaryResponse) SetVsphereHostTop99pSum(v int64)`

SetVsphereHostTop99pSum sets VsphereHostTop99pSum field to given value.

### HasVsphereHostTop99pSum

`func (o *UsageSummaryResponse) HasVsphereHostTop99pSum() bool`

HasVsphereHostTop99pSum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


