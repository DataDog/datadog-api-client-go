# UsageSummaryDate

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AgentHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all agent hosts over all hours in the current date for all organizations. | [optional] 
**ApmAzureAppServiceHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all Azure app services using APM over all hours in the current date all organizations. | [optional] 
**ApmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current date for all organizations. | [optional] 
**AuditLogsLinesIndexedSum** | Pointer to **int64** | Shows the sum of audit logs lines indexed over all hours in the current date for all organizations. | [optional] 
**AwsHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all AWS hosts over all hours in the current date for all organizations. | [optional] 
**AwsLambdaFuncCount** | Pointer to **int64** | Shows the average of the number of functions that executed 1 or more times each hour in the current date for all organizations. | [optional] 
**AwsLambdaInvocationsSum** | Pointer to **int64** | Shows the sum of all AWS Lambda invocations over all hours in the current date for all organizations. | [optional] 
**AzureAppServiceTop99p** | Pointer to **int64** | Shows the 99th percentile of all Azure app services over all hours in the current date for all organizations. | [optional] 
**BillableIngestedBytesSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current date for all organizations. | [optional] 
**ContainerAvg** | Pointer to **int64** | Shows the average of all distinct containers over all hours in the current date for all organizations. | [optional] 
**ContainerHwm** | Pointer to **int64** | Shows the high-water mark of all distinct containers over all hours in the current date for all organizations. | [optional] 
**CspmContainerAvg** | Pointer to **int64** | Shows the average number of Cloud Security Posture Management containers over all hours in the current date for all organizations. | [optional] 
**CspmContainerHwm** | Pointer to **int64** | Shows the high-water mark of Cloud Security Posture Management containers over all hours in the current date for all organizations. | [optional] 
**CspmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all Cloud Security Posture Management hosts over all hours in the current date for all organizations. | [optional] 
**CustomTsAvg** | Pointer to **int64** | Shows the average number of distinct custom metrics over all hours in the current date for all organizations. | [optional] 
**CwsContainerCountAvg** | Pointer to **int64** | Shows the average of all distinct Cloud Workload Security containers over all hours in the current date for all organizations. | [optional] 
**CwsHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all Cloud Workload Security hosts over all hours in the current date for all organizations. | [optional] 
**Date** | Pointer to **time.Time** | The date for the usage. | [optional] 
**DbmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all Database Monitoring hosts over all hours in the current date for all organizations. | [optional] 
**DbmQueriesCountAvg** | Pointer to **int64** | Shows the average of all normalized Database Monitoring queries over all hours in the current date for all organizations. | [optional] 
**FargateTasksCountAvg** | Pointer to **int64** | Shows the high-watermark of all Fargate tasks over all hours in the current date for all organizations. | [optional] 
**FargateTasksCountHwm** | Pointer to **int64** | Shows the average of all Fargate tasks over all hours in the current date for all organizations. | [optional] 
**GcpHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all GCP hosts over all hours in the current date for all organizations. | [optional] 
**HerokuHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all Heroku dynos over all hours in the current date for all organizations. | [optional] 
**IncidentManagementMonthlyActiveUsersHwm** | Pointer to **int64** | Shows the high-water mark of incident management monthly active users over all hours in the current date for all organizations. | [optional] 
**IndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all log events indexed over all hours in the current date for all organizations. | [optional] 
**InfraHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for all organizations. | [optional] 
**IngestedEventsBytesSum** | Pointer to **int64** | Shows the sum of all log bytes ingested over all hours in the current date for all organizations. | [optional] 
**IotDeviceSum** | Pointer to **int64** | Shows the sum of all IoT devices over all hours in the current date for all organizations. | [optional] 
**IotDeviceTop99p** | Pointer to **int64** | Shows the 99th percentile of all IoT devices over all hours in the current date all organizations. | [optional] 
**MobileRumSessionCountAndroidSum** | Pointer to **int64** | Shows the sum of all mobile RUM Sessions on Android over all hours in the current date for all organizations. | [optional] 
**MobileRumSessionCountIosSum** | Pointer to **int64** | Shows the sum of all mobile RUM Sessions on iOS over all hours in the current date for all organizations. | [optional] 
**MobileRumSessionCountSum** | Pointer to **int64** | Shows the sum of all mobile RUM Sessions over all hours in the current date for all organizations | [optional] 
**NetflowIndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all Network flows indexed over all hours in the current date for all organizations. | [optional] 
**NpmHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for all organizations. | [optional] 
**OpentelemetryHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for all organizations. | [optional] 
**Orgs** | Pointer to [**[]UsageSummaryDateOrg**](UsageSummaryDateOrg.md) | Organizations associated with a user. | [optional] 
**ProfilingHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all profiled hosts over all hours in the current date for all organizations. | [optional] 
**RumSessionCountSum** | Pointer to **int64** | Shows the sum of all browser RUM Sessions over all hours in the current date for all organizations | [optional] 
**RumTotalSessionCountSum** | Pointer to **int64** | Shows the sum of RUM Sessions (browser and mobile) over all hours in the current date for all organizations. | [optional] 
**SyntheticsBrowserCheckCallsCountSum** | Pointer to **int64** | Shows the sum of all Synthetic browser tests over all hours in the current date for all organizations. | [optional] 
**SyntheticsCheckCallsCountSum** | Pointer to **int64** | Shows the sum of all Synthetic API tests over all hours in the current date for all organizations. | [optional] 
**TraceSearchIndexedEventsCountSum** | Pointer to **int64** | Shows the sum of all Indexed Spans indexed over all hours in the current date for all organizations. | [optional] 
**TwolIngestedEventsBytesSum** | Pointer to **int64** | Shows the sum of all tracing without limits bytes ingested over all hours in the current date for all organizations. | [optional] 
**VsphereHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all vSphere hosts over all hours in the current date for all organizations. | [optional] 

## Methods

### NewUsageSummaryDate

`func NewUsageSummaryDate() *UsageSummaryDate`

NewUsageSummaryDate instantiates a new UsageSummaryDate object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageSummaryDateWithDefaults

`func NewUsageSummaryDateWithDefaults() *UsageSummaryDate`

NewUsageSummaryDateWithDefaults instantiates a new UsageSummaryDate object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

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

### GetApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDate) GetApmAzureAppServiceHostTop99p() int64`

GetApmAzureAppServiceHostTop99p returns the ApmAzureAppServiceHostTop99p field if non-nil, zero value otherwise.

### GetApmAzureAppServiceHostTop99pOk

`func (o *UsageSummaryDate) GetApmAzureAppServiceHostTop99pOk() (*int64, bool)`

GetApmAzureAppServiceHostTop99pOk returns a tuple with the ApmAzureAppServiceHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDate) SetApmAzureAppServiceHostTop99p(v int64)`

SetApmAzureAppServiceHostTop99p sets ApmAzureAppServiceHostTop99p field to given value.

### HasApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDate) HasApmAzureAppServiceHostTop99p() bool`

HasApmAzureAppServiceHostTop99p returns a boolean if a field has been set.

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

### GetAuditLogsLinesIndexedSum

`func (o *UsageSummaryDate) GetAuditLogsLinesIndexedSum() int64`

GetAuditLogsLinesIndexedSum returns the AuditLogsLinesIndexedSum field if non-nil, zero value otherwise.

### GetAuditLogsLinesIndexedSumOk

`func (o *UsageSummaryDate) GetAuditLogsLinesIndexedSumOk() (*int64, bool)`

GetAuditLogsLinesIndexedSumOk returns a tuple with the AuditLogsLinesIndexedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsLinesIndexedSum

`func (o *UsageSummaryDate) SetAuditLogsLinesIndexedSum(v int64)`

SetAuditLogsLinesIndexedSum sets AuditLogsLinesIndexedSum field to given value.

### HasAuditLogsLinesIndexedSum

`func (o *UsageSummaryDate) HasAuditLogsLinesIndexedSum() bool`

HasAuditLogsLinesIndexedSum returns a boolean if a field has been set.

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

### GetAzureAppServiceTop99p

`func (o *UsageSummaryDate) GetAzureAppServiceTop99p() int64`

GetAzureAppServiceTop99p returns the AzureAppServiceTop99p field if non-nil, zero value otherwise.

### GetAzureAppServiceTop99pOk

`func (o *UsageSummaryDate) GetAzureAppServiceTop99pOk() (*int64, bool)`

GetAzureAppServiceTop99pOk returns a tuple with the AzureAppServiceTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppServiceTop99p

`func (o *UsageSummaryDate) SetAzureAppServiceTop99p(v int64)`

SetAzureAppServiceTop99p sets AzureAppServiceTop99p field to given value.

### HasAzureAppServiceTop99p

`func (o *UsageSummaryDate) HasAzureAppServiceTop99p() bool`

HasAzureAppServiceTop99p returns a boolean if a field has been set.

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

### GetCspmContainerAvg

`func (o *UsageSummaryDate) GetCspmContainerAvg() int64`

GetCspmContainerAvg returns the CspmContainerAvg field if non-nil, zero value otherwise.

### GetCspmContainerAvgOk

`func (o *UsageSummaryDate) GetCspmContainerAvgOk() (*int64, bool)`

GetCspmContainerAvgOk returns a tuple with the CspmContainerAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmContainerAvg

`func (o *UsageSummaryDate) SetCspmContainerAvg(v int64)`

SetCspmContainerAvg sets CspmContainerAvg field to given value.

### HasCspmContainerAvg

`func (o *UsageSummaryDate) HasCspmContainerAvg() bool`

HasCspmContainerAvg returns a boolean if a field has been set.

### GetCspmContainerHwm

`func (o *UsageSummaryDate) GetCspmContainerHwm() int64`

GetCspmContainerHwm returns the CspmContainerHwm field if non-nil, zero value otherwise.

### GetCspmContainerHwmOk

`func (o *UsageSummaryDate) GetCspmContainerHwmOk() (*int64, bool)`

GetCspmContainerHwmOk returns a tuple with the CspmContainerHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmContainerHwm

`func (o *UsageSummaryDate) SetCspmContainerHwm(v int64)`

SetCspmContainerHwm sets CspmContainerHwm field to given value.

### HasCspmContainerHwm

`func (o *UsageSummaryDate) HasCspmContainerHwm() bool`

HasCspmContainerHwm returns a boolean if a field has been set.

### GetCspmHostTop99p

`func (o *UsageSummaryDate) GetCspmHostTop99p() int64`

GetCspmHostTop99p returns the CspmHostTop99p field if non-nil, zero value otherwise.

### GetCspmHostTop99pOk

`func (o *UsageSummaryDate) GetCspmHostTop99pOk() (*int64, bool)`

GetCspmHostTop99pOk returns a tuple with the CspmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspmHostTop99p

`func (o *UsageSummaryDate) SetCspmHostTop99p(v int64)`

SetCspmHostTop99p sets CspmHostTop99p field to given value.

### HasCspmHostTop99p

`func (o *UsageSummaryDate) HasCspmHostTop99p() bool`

HasCspmHostTop99p returns a boolean if a field has been set.

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

### GetCwsContainerCountAvg

`func (o *UsageSummaryDate) GetCwsContainerCountAvg() int64`

GetCwsContainerCountAvg returns the CwsContainerCountAvg field if non-nil, zero value otherwise.

### GetCwsContainerCountAvgOk

`func (o *UsageSummaryDate) GetCwsContainerCountAvgOk() (*int64, bool)`

GetCwsContainerCountAvgOk returns a tuple with the CwsContainerCountAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsContainerCountAvg

`func (o *UsageSummaryDate) SetCwsContainerCountAvg(v int64)`

SetCwsContainerCountAvg sets CwsContainerCountAvg field to given value.

### HasCwsContainerCountAvg

`func (o *UsageSummaryDate) HasCwsContainerCountAvg() bool`

HasCwsContainerCountAvg returns a boolean if a field has been set.

### GetCwsHostTop99p

`func (o *UsageSummaryDate) GetCwsHostTop99p() int64`

GetCwsHostTop99p returns the CwsHostTop99p field if non-nil, zero value otherwise.

### GetCwsHostTop99pOk

`func (o *UsageSummaryDate) GetCwsHostTop99pOk() (*int64, bool)`

GetCwsHostTop99pOk returns a tuple with the CwsHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwsHostTop99p

`func (o *UsageSummaryDate) SetCwsHostTop99p(v int64)`

SetCwsHostTop99p sets CwsHostTop99p field to given value.

### HasCwsHostTop99p

`func (o *UsageSummaryDate) HasCwsHostTop99p() bool`

HasCwsHostTop99p returns a boolean if a field has been set.

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

### GetDbmHostTop99p

`func (o *UsageSummaryDate) GetDbmHostTop99p() int64`

GetDbmHostTop99p returns the DbmHostTop99p field if non-nil, zero value otherwise.

### GetDbmHostTop99pOk

`func (o *UsageSummaryDate) GetDbmHostTop99pOk() (*int64, bool)`

GetDbmHostTop99pOk returns a tuple with the DbmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmHostTop99p

`func (o *UsageSummaryDate) SetDbmHostTop99p(v int64)`

SetDbmHostTop99p sets DbmHostTop99p field to given value.

### HasDbmHostTop99p

`func (o *UsageSummaryDate) HasDbmHostTop99p() bool`

HasDbmHostTop99p returns a boolean if a field has been set.

### GetDbmQueriesCountAvg

`func (o *UsageSummaryDate) GetDbmQueriesCountAvg() int64`

GetDbmQueriesCountAvg returns the DbmQueriesCountAvg field if non-nil, zero value otherwise.

### GetDbmQueriesCountAvgOk

`func (o *UsageSummaryDate) GetDbmQueriesCountAvgOk() (*int64, bool)`

GetDbmQueriesCountAvgOk returns a tuple with the DbmQueriesCountAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmQueriesCountAvg

`func (o *UsageSummaryDate) SetDbmQueriesCountAvg(v int64)`

SetDbmQueriesCountAvg sets DbmQueriesCountAvg field to given value.

### HasDbmQueriesCountAvg

`func (o *UsageSummaryDate) HasDbmQueriesCountAvg() bool`

HasDbmQueriesCountAvg returns a boolean if a field has been set.

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

### GetHerokuHostTop99p

`func (o *UsageSummaryDate) GetHerokuHostTop99p() int64`

GetHerokuHostTop99p returns the HerokuHostTop99p field if non-nil, zero value otherwise.

### GetHerokuHostTop99pOk

`func (o *UsageSummaryDate) GetHerokuHostTop99pOk() (*int64, bool)`

GetHerokuHostTop99pOk returns a tuple with the HerokuHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHerokuHostTop99p

`func (o *UsageSummaryDate) SetHerokuHostTop99p(v int64)`

SetHerokuHostTop99p sets HerokuHostTop99p field to given value.

### HasHerokuHostTop99p

`func (o *UsageSummaryDate) HasHerokuHostTop99p() bool`

HasHerokuHostTop99p returns a boolean if a field has been set.

### GetIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDate) GetIncidentManagementMonthlyActiveUsersHwm() int64`

GetIncidentManagementMonthlyActiveUsersHwm returns the IncidentManagementMonthlyActiveUsersHwm field if non-nil, zero value otherwise.

### GetIncidentManagementMonthlyActiveUsersHwmOk

`func (o *UsageSummaryDate) GetIncidentManagementMonthlyActiveUsersHwmOk() (*int64, bool)`

GetIncidentManagementMonthlyActiveUsersHwmOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDate) SetIncidentManagementMonthlyActiveUsersHwm(v int64)`

SetIncidentManagementMonthlyActiveUsersHwm sets IncidentManagementMonthlyActiveUsersHwm field to given value.

### HasIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDate) HasIncidentManagementMonthlyActiveUsersHwm() bool`

HasIncidentManagementMonthlyActiveUsersHwm returns a boolean if a field has been set.

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

### GetIotDeviceSum

`func (o *UsageSummaryDate) GetIotDeviceSum() int64`

GetIotDeviceSum returns the IotDeviceSum field if non-nil, zero value otherwise.

### GetIotDeviceSumOk

`func (o *UsageSummaryDate) GetIotDeviceSumOk() (*int64, bool)`

GetIotDeviceSumOk returns a tuple with the IotDeviceSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceSum

`func (o *UsageSummaryDate) SetIotDeviceSum(v int64)`

SetIotDeviceSum sets IotDeviceSum field to given value.

### HasIotDeviceSum

`func (o *UsageSummaryDate) HasIotDeviceSum() bool`

HasIotDeviceSum returns a boolean if a field has been set.

### GetIotDeviceTop99p

`func (o *UsageSummaryDate) GetIotDeviceTop99p() int64`

GetIotDeviceTop99p returns the IotDeviceTop99p field if non-nil, zero value otherwise.

### GetIotDeviceTop99pOk

`func (o *UsageSummaryDate) GetIotDeviceTop99pOk() (*int64, bool)`

GetIotDeviceTop99pOk returns a tuple with the IotDeviceTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceTop99p

`func (o *UsageSummaryDate) SetIotDeviceTop99p(v int64)`

SetIotDeviceTop99p sets IotDeviceTop99p field to given value.

### HasIotDeviceTop99p

`func (o *UsageSummaryDate) HasIotDeviceTop99p() bool`

HasIotDeviceTop99p returns a boolean if a field has been set.

### GetMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDate) GetMobileRumSessionCountAndroidSum() int64`

GetMobileRumSessionCountAndroidSum returns the MobileRumSessionCountAndroidSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountAndroidSumOk

`func (o *UsageSummaryDate) GetMobileRumSessionCountAndroidSumOk() (*int64, bool)`

GetMobileRumSessionCountAndroidSumOk returns a tuple with the MobileRumSessionCountAndroidSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDate) SetMobileRumSessionCountAndroidSum(v int64)`

SetMobileRumSessionCountAndroidSum sets MobileRumSessionCountAndroidSum field to given value.

### HasMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDate) HasMobileRumSessionCountAndroidSum() bool`

HasMobileRumSessionCountAndroidSum returns a boolean if a field has been set.

### GetMobileRumSessionCountIosSum

`func (o *UsageSummaryDate) GetMobileRumSessionCountIosSum() int64`

GetMobileRumSessionCountIosSum returns the MobileRumSessionCountIosSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountIosSumOk

`func (o *UsageSummaryDate) GetMobileRumSessionCountIosSumOk() (*int64, bool)`

GetMobileRumSessionCountIosSumOk returns a tuple with the MobileRumSessionCountIosSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountIosSum

`func (o *UsageSummaryDate) SetMobileRumSessionCountIosSum(v int64)`

SetMobileRumSessionCountIosSum sets MobileRumSessionCountIosSum field to given value.

### HasMobileRumSessionCountIosSum

`func (o *UsageSummaryDate) HasMobileRumSessionCountIosSum() bool`

HasMobileRumSessionCountIosSum returns a boolean if a field has been set.

### GetMobileRumSessionCountSum

`func (o *UsageSummaryDate) GetMobileRumSessionCountSum() int64`

GetMobileRumSessionCountSum returns the MobileRumSessionCountSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountSumOk

`func (o *UsageSummaryDate) GetMobileRumSessionCountSumOk() (*int64, bool)`

GetMobileRumSessionCountSumOk returns a tuple with the MobileRumSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountSum

`func (o *UsageSummaryDate) SetMobileRumSessionCountSum(v int64)`

SetMobileRumSessionCountSum sets MobileRumSessionCountSum field to given value.

### HasMobileRumSessionCountSum

`func (o *UsageSummaryDate) HasMobileRumSessionCountSum() bool`

HasMobileRumSessionCountSum returns a boolean if a field has been set.

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

### GetOpentelemetryHostTop99p

`func (o *UsageSummaryDate) GetOpentelemetryHostTop99p() int64`

GetOpentelemetryHostTop99p returns the OpentelemetryHostTop99p field if non-nil, zero value otherwise.

### GetOpentelemetryHostTop99pOk

`func (o *UsageSummaryDate) GetOpentelemetryHostTop99pOk() (*int64, bool)`

GetOpentelemetryHostTop99pOk returns a tuple with the OpentelemetryHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpentelemetryHostTop99p

`func (o *UsageSummaryDate) SetOpentelemetryHostTop99p(v int64)`

SetOpentelemetryHostTop99p sets OpentelemetryHostTop99p field to given value.

### HasOpentelemetryHostTop99p

`func (o *UsageSummaryDate) HasOpentelemetryHostTop99p() bool`

HasOpentelemetryHostTop99p returns a boolean if a field has been set.

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

### GetProfilingHostTop99p

`func (o *UsageSummaryDate) GetProfilingHostTop99p() int64`

GetProfilingHostTop99p returns the ProfilingHostTop99p field if non-nil, zero value otherwise.

### GetProfilingHostTop99pOk

`func (o *UsageSummaryDate) GetProfilingHostTop99pOk() (*int64, bool)`

GetProfilingHostTop99pOk returns a tuple with the ProfilingHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilingHostTop99p

`func (o *UsageSummaryDate) SetProfilingHostTop99p(v int64)`

SetProfilingHostTop99p sets ProfilingHostTop99p field to given value.

### HasProfilingHostTop99p

`func (o *UsageSummaryDate) HasProfilingHostTop99p() bool`

HasProfilingHostTop99p returns a boolean if a field has been set.

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

### GetRumTotalSessionCountSum

`func (o *UsageSummaryDate) GetRumTotalSessionCountSum() int64`

GetRumTotalSessionCountSum returns the RumTotalSessionCountSum field if non-nil, zero value otherwise.

### GetRumTotalSessionCountSumOk

`func (o *UsageSummaryDate) GetRumTotalSessionCountSumOk() (*int64, bool)`

GetRumTotalSessionCountSumOk returns a tuple with the RumTotalSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumTotalSessionCountSum

`func (o *UsageSummaryDate) SetRumTotalSessionCountSum(v int64)`

SetRumTotalSessionCountSum sets RumTotalSessionCountSum field to given value.

### HasRumTotalSessionCountSum

`func (o *UsageSummaryDate) HasRumTotalSessionCountSum() bool`

HasRumTotalSessionCountSum returns a boolean if a field has been set.

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

### GetTwolIngestedEventsBytesSum

`func (o *UsageSummaryDate) GetTwolIngestedEventsBytesSum() int64`

GetTwolIngestedEventsBytesSum returns the TwolIngestedEventsBytesSum field if non-nil, zero value otherwise.

### GetTwolIngestedEventsBytesSumOk

`func (o *UsageSummaryDate) GetTwolIngestedEventsBytesSumOk() (*int64, bool)`

GetTwolIngestedEventsBytesSumOk returns a tuple with the TwolIngestedEventsBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwolIngestedEventsBytesSum

`func (o *UsageSummaryDate) SetTwolIngestedEventsBytesSum(v int64)`

SetTwolIngestedEventsBytesSum sets TwolIngestedEventsBytesSum field to given value.

### HasTwolIngestedEventsBytesSum

`func (o *UsageSummaryDate) HasTwolIngestedEventsBytesSum() bool`

HasTwolIngestedEventsBytesSum returns a boolean if a field has been set.

### GetVsphereHostTop99p

`func (o *UsageSummaryDate) GetVsphereHostTop99p() int64`

GetVsphereHostTop99p returns the VsphereHostTop99p field if non-nil, zero value otherwise.

### GetVsphereHostTop99pOk

`func (o *UsageSummaryDate) GetVsphereHostTop99pOk() (*int64, bool)`

GetVsphereHostTop99pOk returns a tuple with the VsphereHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereHostTop99p

`func (o *UsageSummaryDate) SetVsphereHostTop99p(v int64)`

SetVsphereHostTop99p sets VsphereHostTop99p field to given value.

### HasVsphereHostTop99p

`func (o *UsageSummaryDate) HasVsphereHostTop99p() bool`

HasVsphereHostTop99p returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


