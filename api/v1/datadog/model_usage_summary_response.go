/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// UsageSummaryResponse Response with hourly report of all data billed by Datadog all organizations.
type UsageSummaryResponse struct {
	// Shows the 99th percentile of all agent hosts over all hours in the current month(s) for all orgs.
	AgentHostTop99pSum *int64 `json:"agent_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all distinct APM hosts over all hours in the current month(s) for all orgs.
	ApmHostTop99pSum *int64 `json:"apm_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all AWS hosts over all hours in the current month(s) for all orgs.
	AwsHostTop99pSum *int64 `json:"aws_host_top99p_sum,omitempty"`
	// Shows the average of the number of functions that executed 1 or more times each hour in the current month(s) for all orgs.
	AwsLambdaFuncCount *int64 `json:"aws_lambda_func_count,omitempty"`
	// Shows the sum of all AWS Labmda invocations over all hours in the current month(s) for all orgs.
	AwsLambdaInvocationsSum *int64 `json:"aws_lambda_invocations_sum,omitempty"`
	// Shows the 99th percentile of all Azure hosts over all hours in the current month(s) for all orgs.
	AzureHostTop99pSum *int64 `json:"azure_host_top99p_sum,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current month(s) for all orgs.
	BillableIngestedBytesAggSum *int64 `json:"billable_ingested_bytes_agg_sum,omitempty"`
	// Shows the average of all distinct containers over all hours in the current month(s) for all orgs.
	ContainerAvgSum *int64 `json:"container_avg_sum,omitempty"`
	// Shows the high watermark of all distinct containers over all hours in the current month(s) for all orgs.
	ContainerHwmSum *int64 `json:"container_hwm_sum,omitempty"`
	// Shows the average number of distinct custom metrics over all hours in the current month(s) for all orgs.
	CustomTsSum *int64 `json:"custom_ts_sum,omitempty"`
	// Shows the last date of usage in the current month(s) for all orgs.
	EndDate *time.Time `json:"end_date,omitempty"`
	// Shows the average of all Fargate tasks over all hours in the current month(s) for all orgs.
	FargateTasksCountAvgSum *int64 `json:"fargate_tasks_count_avg_sum,omitempty"`
	// Shows the high watermark of all Fargate tasks over all hours in the current month(s) for all orgs.
	FargateTasksCountHwmSum *int64 `json:"fargate_tasks_count_hwm_sum,omitempty"`
	// Shows the 99th percentile of all GCP hosts over all hours in the current month(s) for all orgs.
	GcpHostTop99pSum *int64 `json:"gcp_host_top99p_sum,omitempty"`
	// Shows the sum of all log events indexed over all hours in the current month(s) for all orgs.
	IndexedEventsCountAggSum *int64 `json:"indexed_events_count_agg_sum,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current month(s) for all orgs.
	InfraHostTop99pSum *int64 `json:"infra_host_top99p_sum,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current month(s) for all orgs.
	IngestedEventsBytesAggSum *int64 `json:"ingested_events_bytes_agg_sum,omitempty"`
	// Shows the the most recent hour in the current month(s) for all orgs for which all usages were calculated.
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Shows the sum of all Network flows indexed over all hours in the current month(s) for all orgs.
	NetflowIndexedEventsCountAggSum *int64 `json:"netflow_indexed_events_count_agg_sum,omitempty"`
	// Shows the 99th percentile of all distinct Networks hosts over all hours in the current month(s) for all orgs.
	NpmHostTop99pSum *int64 `json:"npm_host_top99p_sum,omitempty"`
	// Shows the sum of all RUM Sessions over all hours in the current month(s) for all orgs.
	RumSessionCountAggSum *int64 `json:"rum_session_count_agg_sum,omitempty"`
	// Shows the first date of usage in the current month(s) for all orgs.
	StartDate *time.Time `json:"start_date,omitempty"`
	// Shows the sum of all Synthetic browser tests over all hours in the current month(s) for all orgs.
	SyntheticsBrowserCheckCallsCountAggSum *int64 `json:"synthetics_browser_check_calls_count_agg_sum,omitempty"`
	// Shows the sum of all Synthetic API tests over all hours in the current month(s) for all orgs.
	SyntheticsCheckCallsCountAggSum *int64 `json:"synthetics_check_calls_count_agg_sum,omitempty"`
	// Shows the sum of all analyzed spans indexed over all hours in the current month(s) for all orgs.
	TraceSearchIndexedEventsCountAggSum *int64 `json:"trace_search_indexed_events_count_agg_sum,omitempty"`
	// An array of objects regarding hourly useage.
	Usage *[]UsageSummaryDate `json:"usage,omitempty"`
}

// NewUsageSummaryResponse instantiates a new UsageSummaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSummaryResponse() *UsageSummaryResponse {
	this := UsageSummaryResponse{}
	return &this
}

// NewUsageSummaryResponseWithDefaults instantiates a new UsageSummaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSummaryResponseWithDefaults() *UsageSummaryResponse {
	this := UsageSummaryResponse{}
	return &this
}

// GetAgentHostTop99pSum returns the AgentHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAgentHostTop99pSum() int64 {
	if o == nil || o.AgentHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.AgentHostTop99pSum
}

// GetAgentHostTop99pSumOk returns a tuple with the AgentHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAgentHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.AgentHostTop99pSum == nil {
		return nil, false
	}
	return o.AgentHostTop99pSum, true
}

// HasAgentHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAgentHostTop99pSum() bool {
	if o != nil && o.AgentHostTop99pSum != nil {
		return true
	}

	return false
}

// SetAgentHostTop99pSum gets a reference to the given int64 and assigns it to the AgentHostTop99pSum field.
func (o *UsageSummaryResponse) SetAgentHostTop99pSum(v int64) {
	o.AgentHostTop99pSum = &v
}

// GetApmHostTop99pSum returns the ApmHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetApmHostTop99pSum() int64 {
	if o == nil || o.ApmHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.ApmHostTop99pSum
}

// GetApmHostTop99pSumOk returns a tuple with the ApmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetApmHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.ApmHostTop99pSum == nil {
		return nil, false
	}
	return o.ApmHostTop99pSum, true
}

// HasApmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmHostTop99pSum() bool {
	if o != nil && o.ApmHostTop99pSum != nil {
		return true
	}

	return false
}

// SetApmHostTop99pSum gets a reference to the given int64 and assigns it to the ApmHostTop99pSum field.
func (o *UsageSummaryResponse) SetApmHostTop99pSum(v int64) {
	o.ApmHostTop99pSum = &v
}

// GetAwsHostTop99pSum returns the AwsHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAwsHostTop99pSum() int64 {
	if o == nil || o.AwsHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.AwsHostTop99pSum
}

// GetAwsHostTop99pSumOk returns a tuple with the AwsHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAwsHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.AwsHostTop99pSum == nil {
		return nil, false
	}
	return o.AwsHostTop99pSum, true
}

// HasAwsHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAwsHostTop99pSum() bool {
	if o != nil && o.AwsHostTop99pSum != nil {
		return true
	}

	return false
}

// SetAwsHostTop99pSum gets a reference to the given int64 and assigns it to the AwsHostTop99pSum field.
func (o *UsageSummaryResponse) SetAwsHostTop99pSum(v int64) {
	o.AwsHostTop99pSum = &v
}

// GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAwsLambdaFuncCount() int64 {
	if o == nil || o.AwsLambdaFuncCount == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaFuncCount
}

// GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAwsLambdaFuncCountOk() (*int64, bool) {
	if o == nil || o.AwsLambdaFuncCount == nil {
		return nil, false
	}
	return o.AwsLambdaFuncCount, true
}

// HasAwsLambdaFuncCount returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAwsLambdaFuncCount() bool {
	if o != nil && o.AwsLambdaFuncCount != nil {
		return true
	}

	return false
}

// SetAwsLambdaFuncCount gets a reference to the given int64 and assigns it to the AwsLambdaFuncCount field.
func (o *UsageSummaryResponse) SetAwsLambdaFuncCount(v int64) {
	o.AwsLambdaFuncCount = &v
}

// GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAwsLambdaInvocationsSum() int64 {
	if o == nil || o.AwsLambdaInvocationsSum == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaInvocationsSum
}

// GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAwsLambdaInvocationsSumOk() (*int64, bool) {
	if o == nil || o.AwsLambdaInvocationsSum == nil {
		return nil, false
	}
	return o.AwsLambdaInvocationsSum, true
}

// HasAwsLambdaInvocationsSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAwsLambdaInvocationsSum() bool {
	if o != nil && o.AwsLambdaInvocationsSum != nil {
		return true
	}

	return false
}

// SetAwsLambdaInvocationsSum gets a reference to the given int64 and assigns it to the AwsLambdaInvocationsSum field.
func (o *UsageSummaryResponse) SetAwsLambdaInvocationsSum(v int64) {
	o.AwsLambdaInvocationsSum = &v
}

// GetAzureHostTop99pSum returns the AzureHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAzureHostTop99pSum() int64 {
	if o == nil || o.AzureHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.AzureHostTop99pSum
}

// GetAzureHostTop99pSumOk returns a tuple with the AzureHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAzureHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.AzureHostTop99pSum == nil {
		return nil, false
	}
	return o.AzureHostTop99pSum, true
}

// HasAzureHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAzureHostTop99pSum() bool {
	if o != nil && o.AzureHostTop99pSum != nil {
		return true
	}

	return false
}

// SetAzureHostTop99pSum gets a reference to the given int64 and assigns it to the AzureHostTop99pSum field.
func (o *UsageSummaryResponse) SetAzureHostTop99pSum(v int64) {
	o.AzureHostTop99pSum = &v
}

// GetBillableIngestedBytesAggSum returns the BillableIngestedBytesAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetBillableIngestedBytesAggSum() int64 {
	if o == nil || o.BillableIngestedBytesAggSum == nil {
		var ret int64
		return ret
	}
	return *o.BillableIngestedBytesAggSum
}

// GetBillableIngestedBytesAggSumOk returns a tuple with the BillableIngestedBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetBillableIngestedBytesAggSumOk() (*int64, bool) {
	if o == nil || o.BillableIngestedBytesAggSum == nil {
		return nil, false
	}
	return o.BillableIngestedBytesAggSum, true
}

// HasBillableIngestedBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBillableIngestedBytesAggSum() bool {
	if o != nil && o.BillableIngestedBytesAggSum != nil {
		return true
	}

	return false
}

// SetBillableIngestedBytesAggSum gets a reference to the given int64 and assigns it to the BillableIngestedBytesAggSum field.
func (o *UsageSummaryResponse) SetBillableIngestedBytesAggSum(v int64) {
	o.BillableIngestedBytesAggSum = &v
}

// GetContainerAvgSum returns the ContainerAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetContainerAvgSum() int64 {
	if o == nil || o.ContainerAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ContainerAvgSum
}

// GetContainerAvgSumOk returns a tuple with the ContainerAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetContainerAvgSumOk() (*int64, bool) {
	if o == nil || o.ContainerAvgSum == nil {
		return nil, false
	}
	return o.ContainerAvgSum, true
}

// HasContainerAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasContainerAvgSum() bool {
	if o != nil && o.ContainerAvgSum != nil {
		return true
	}

	return false
}

// SetContainerAvgSum gets a reference to the given int64 and assigns it to the ContainerAvgSum field.
func (o *UsageSummaryResponse) SetContainerAvgSum(v int64) {
	o.ContainerAvgSum = &v
}

// GetContainerHwmSum returns the ContainerHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetContainerHwmSum() int64 {
	if o == nil || o.ContainerHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.ContainerHwmSum
}

// GetContainerHwmSumOk returns a tuple with the ContainerHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetContainerHwmSumOk() (*int64, bool) {
	if o == nil || o.ContainerHwmSum == nil {
		return nil, false
	}
	return o.ContainerHwmSum, true
}

// HasContainerHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasContainerHwmSum() bool {
	if o != nil && o.ContainerHwmSum != nil {
		return true
	}

	return false
}

// SetContainerHwmSum gets a reference to the given int64 and assigns it to the ContainerHwmSum field.
func (o *UsageSummaryResponse) SetContainerHwmSum(v int64) {
	o.ContainerHwmSum = &v
}

// GetCustomTsSum returns the CustomTsSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCustomTsSum() int64 {
	if o == nil || o.CustomTsSum == nil {
		var ret int64
		return ret
	}
	return *o.CustomTsSum
}

// GetCustomTsSumOk returns a tuple with the CustomTsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCustomTsSumOk() (*int64, bool) {
	if o == nil || o.CustomTsSum == nil {
		return nil, false
	}
	return o.CustomTsSum, true
}

// HasCustomTsSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCustomTsSum() bool {
	if o != nil && o.CustomTsSum != nil {
		return true
	}

	return false
}

// SetCustomTsSum gets a reference to the given int64 and assigns it to the CustomTsSum field.
func (o *UsageSummaryResponse) SetCustomTsSum(v int64) {
	o.CustomTsSum = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *UsageSummaryResponse) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetFargateTasksCountAvgSum returns the FargateTasksCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFargateTasksCountAvgSum() int64 {
	if o == nil || o.FargateTasksCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountAvgSum
}

// GetFargateTasksCountAvgSumOk returns a tuple with the FargateTasksCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFargateTasksCountAvgSumOk() (*int64, bool) {
	if o == nil || o.FargateTasksCountAvgSum == nil {
		return nil, false
	}
	return o.FargateTasksCountAvgSum, true
}

// HasFargateTasksCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFargateTasksCountAvgSum() bool {
	if o != nil && o.FargateTasksCountAvgSum != nil {
		return true
	}

	return false
}

// SetFargateTasksCountAvgSum gets a reference to the given int64 and assigns it to the FargateTasksCountAvgSum field.
func (o *UsageSummaryResponse) SetFargateTasksCountAvgSum(v int64) {
	o.FargateTasksCountAvgSum = &v
}

// GetFargateTasksCountHwmSum returns the FargateTasksCountHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFargateTasksCountHwmSum() int64 {
	if o == nil || o.FargateTasksCountHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountHwmSum
}

// GetFargateTasksCountHwmSumOk returns a tuple with the FargateTasksCountHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFargateTasksCountHwmSumOk() (*int64, bool) {
	if o == nil || o.FargateTasksCountHwmSum == nil {
		return nil, false
	}
	return o.FargateTasksCountHwmSum, true
}

// HasFargateTasksCountHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFargateTasksCountHwmSum() bool {
	if o != nil && o.FargateTasksCountHwmSum != nil {
		return true
	}

	return false
}

// SetFargateTasksCountHwmSum gets a reference to the given int64 and assigns it to the FargateTasksCountHwmSum field.
func (o *UsageSummaryResponse) SetFargateTasksCountHwmSum(v int64) {
	o.FargateTasksCountHwmSum = &v
}

// GetGcpHostTop99pSum returns the GcpHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetGcpHostTop99pSum() int64 {
	if o == nil || o.GcpHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.GcpHostTop99pSum
}

// GetGcpHostTop99pSumOk returns a tuple with the GcpHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetGcpHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.GcpHostTop99pSum == nil {
		return nil, false
	}
	return o.GcpHostTop99pSum, true
}

// HasGcpHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasGcpHostTop99pSum() bool {
	if o != nil && o.GcpHostTop99pSum != nil {
		return true
	}

	return false
}

// SetGcpHostTop99pSum gets a reference to the given int64 and assigns it to the GcpHostTop99pSum field.
func (o *UsageSummaryResponse) SetGcpHostTop99pSum(v int64) {
	o.GcpHostTop99pSum = &v
}

// GetIndexedEventsCountAggSum returns the IndexedEventsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetIndexedEventsCountAggSum() int64 {
	if o == nil || o.IndexedEventsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCountAggSum
}

// GetIndexedEventsCountAggSumOk returns a tuple with the IndexedEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetIndexedEventsCountAggSumOk() (*int64, bool) {
	if o == nil || o.IndexedEventsCountAggSum == nil {
		return nil, false
	}
	return o.IndexedEventsCountAggSum, true
}

// HasIndexedEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIndexedEventsCountAggSum() bool {
	if o != nil && o.IndexedEventsCountAggSum != nil {
		return true
	}

	return false
}

// SetIndexedEventsCountAggSum gets a reference to the given int64 and assigns it to the IndexedEventsCountAggSum field.
func (o *UsageSummaryResponse) SetIndexedEventsCountAggSum(v int64) {
	o.IndexedEventsCountAggSum = &v
}

// GetInfraHostTop99pSum returns the InfraHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetInfraHostTop99pSum() int64 {
	if o == nil || o.InfraHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostTop99pSum
}

// GetInfraHostTop99pSumOk returns a tuple with the InfraHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetInfraHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.InfraHostTop99pSum == nil {
		return nil, false
	}
	return o.InfraHostTop99pSum, true
}

// HasInfraHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasInfraHostTop99pSum() bool {
	if o != nil && o.InfraHostTop99pSum != nil {
		return true
	}

	return false
}

// SetInfraHostTop99pSum gets a reference to the given int64 and assigns it to the InfraHostTop99pSum field.
func (o *UsageSummaryResponse) SetInfraHostTop99pSum(v int64) {
	o.InfraHostTop99pSum = &v
}

// GetIngestedEventsBytesAggSum returns the IngestedEventsBytesAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetIngestedEventsBytesAggSum() int64 {
	if o == nil || o.IngestedEventsBytesAggSum == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytesAggSum
}

// GetIngestedEventsBytesAggSumOk returns a tuple with the IngestedEventsBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetIngestedEventsBytesAggSumOk() (*int64, bool) {
	if o == nil || o.IngestedEventsBytesAggSum == nil {
		return nil, false
	}
	return o.IngestedEventsBytesAggSum, true
}

// HasIngestedEventsBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIngestedEventsBytesAggSum() bool {
	if o != nil && o.IngestedEventsBytesAggSum != nil {
		return true
	}

	return false
}

// SetIngestedEventsBytesAggSum gets a reference to the given int64 and assigns it to the IngestedEventsBytesAggSum field.
func (o *UsageSummaryResponse) SetIngestedEventsBytesAggSum(v int64) {
	o.IngestedEventsBytesAggSum = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *UsageSummaryResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetNetflowIndexedEventsCountAggSum returns the NetflowIndexedEventsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetNetflowIndexedEventsCountAggSum() int64 {
	if o == nil || o.NetflowIndexedEventsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.NetflowIndexedEventsCountAggSum
}

// GetNetflowIndexedEventsCountAggSumOk returns a tuple with the NetflowIndexedEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetNetflowIndexedEventsCountAggSumOk() (*int64, bool) {
	if o == nil || o.NetflowIndexedEventsCountAggSum == nil {
		return nil, false
	}
	return o.NetflowIndexedEventsCountAggSum, true
}

// HasNetflowIndexedEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasNetflowIndexedEventsCountAggSum() bool {
	if o != nil && o.NetflowIndexedEventsCountAggSum != nil {
		return true
	}

	return false
}

// SetNetflowIndexedEventsCountAggSum gets a reference to the given int64 and assigns it to the NetflowIndexedEventsCountAggSum field.
func (o *UsageSummaryResponse) SetNetflowIndexedEventsCountAggSum(v int64) {
	o.NetflowIndexedEventsCountAggSum = &v
}

// GetNpmHostTop99pSum returns the NpmHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetNpmHostTop99pSum() int64 {
	if o == nil || o.NpmHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.NpmHostTop99pSum
}

// GetNpmHostTop99pSumOk returns a tuple with the NpmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetNpmHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.NpmHostTop99pSum == nil {
		return nil, false
	}
	return o.NpmHostTop99pSum, true
}

// HasNpmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasNpmHostTop99pSum() bool {
	if o != nil && o.NpmHostTop99pSum != nil {
		return true
	}

	return false
}

// SetNpmHostTop99pSum gets a reference to the given int64 and assigns it to the NpmHostTop99pSum field.
func (o *UsageSummaryResponse) SetNpmHostTop99pSum(v int64) {
	o.NpmHostTop99pSum = &v
}

// GetRumSessionCountAggSum returns the RumSessionCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumSessionCountAggSum() int64 {
	if o == nil || o.RumSessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumSessionCountAggSum
}

// GetRumSessionCountAggSumOk returns a tuple with the RumSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumSessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.RumSessionCountAggSum == nil {
		return nil, false
	}
	return o.RumSessionCountAggSum, true
}

// HasRumSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumSessionCountAggSum() bool {
	if o != nil && o.RumSessionCountAggSum != nil {
		return true
	}

	return false
}

// SetRumSessionCountAggSum gets a reference to the given int64 and assigns it to the RumSessionCountAggSum field.
func (o *UsageSummaryResponse) SetRumSessionCountAggSum(v int64) {
	o.RumSessionCountAggSum = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *UsageSummaryResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetSyntheticsBrowserCheckCallsCountAggSum returns the SyntheticsBrowserCheckCallsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSyntheticsBrowserCheckCallsCountAggSum() int64 {
	if o == nil || o.SyntheticsBrowserCheckCallsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsBrowserCheckCallsCountAggSum
}

// GetSyntheticsBrowserCheckCallsCountAggSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSyntheticsBrowserCheckCallsCountAggSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsBrowserCheckCallsCountAggSum == nil {
		return nil, false
	}
	return o.SyntheticsBrowserCheckCallsCountAggSum, true
}

// HasSyntheticsBrowserCheckCallsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSyntheticsBrowserCheckCallsCountAggSum() bool {
	if o != nil && o.SyntheticsBrowserCheckCallsCountAggSum != nil {
		return true
	}

	return false
}

// SetSyntheticsBrowserCheckCallsCountAggSum gets a reference to the given int64 and assigns it to the SyntheticsBrowserCheckCallsCountAggSum field.
func (o *UsageSummaryResponse) SetSyntheticsBrowserCheckCallsCountAggSum(v int64) {
	o.SyntheticsBrowserCheckCallsCountAggSum = &v
}

// GetSyntheticsCheckCallsCountAggSum returns the SyntheticsCheckCallsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSyntheticsCheckCallsCountAggSum() int64 {
	if o == nil || o.SyntheticsCheckCallsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsCheckCallsCountAggSum
}

// GetSyntheticsCheckCallsCountAggSumOk returns a tuple with the SyntheticsCheckCallsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSyntheticsCheckCallsCountAggSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsCheckCallsCountAggSum == nil {
		return nil, false
	}
	return o.SyntheticsCheckCallsCountAggSum, true
}

// HasSyntheticsCheckCallsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSyntheticsCheckCallsCountAggSum() bool {
	if o != nil && o.SyntheticsCheckCallsCountAggSum != nil {
		return true
	}

	return false
}

// SetSyntheticsCheckCallsCountAggSum gets a reference to the given int64 and assigns it to the SyntheticsCheckCallsCountAggSum field.
func (o *UsageSummaryResponse) SetSyntheticsCheckCallsCountAggSum(v int64) {
	o.SyntheticsCheckCallsCountAggSum = &v
}

// GetTraceSearchIndexedEventsCountAggSum returns the TraceSearchIndexedEventsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetTraceSearchIndexedEventsCountAggSum() int64 {
	if o == nil || o.TraceSearchIndexedEventsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.TraceSearchIndexedEventsCountAggSum
}

// GetTraceSearchIndexedEventsCountAggSumOk returns a tuple with the TraceSearchIndexedEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetTraceSearchIndexedEventsCountAggSumOk() (*int64, bool) {
	if o == nil || o.TraceSearchIndexedEventsCountAggSum == nil {
		return nil, false
	}
	return o.TraceSearchIndexedEventsCountAggSum, true
}

// HasTraceSearchIndexedEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasTraceSearchIndexedEventsCountAggSum() bool {
	if o != nil && o.TraceSearchIndexedEventsCountAggSum != nil {
		return true
	}

	return false
}

// SetTraceSearchIndexedEventsCountAggSum gets a reference to the given int64 and assigns it to the TraceSearchIndexedEventsCountAggSum field.
func (o *UsageSummaryResponse) SetTraceSearchIndexedEventsCountAggSum(v int64) {
	o.TraceSearchIndexedEventsCountAggSum = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetUsage() []UsageSummaryDate {
	if o == nil || o.Usage == nil {
		var ret []UsageSummaryDate
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetUsageOk() (*[]UsageSummaryDate, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageSummaryDate and assigns it to the Usage field.
func (o *UsageSummaryResponse) SetUsage(v []UsageSummaryDate) {
	o.Usage = &v
}

func (o UsageSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentHostTop99pSum != nil {
		toSerialize["agent_host_top99p_sum"] = o.AgentHostTop99pSum
	}
	if o.ApmHostTop99pSum != nil {
		toSerialize["apm_host_top99p_sum"] = o.ApmHostTop99pSum
	}
	if o.AwsHostTop99pSum != nil {
		toSerialize["aws_host_top99p_sum"] = o.AwsHostTop99pSum
	}
	if o.AwsLambdaFuncCount != nil {
		toSerialize["aws_lambda_func_count"] = o.AwsLambdaFuncCount
	}
	if o.AwsLambdaInvocationsSum != nil {
		toSerialize["aws_lambda_invocations_sum"] = o.AwsLambdaInvocationsSum
	}
	if o.AzureHostTop99pSum != nil {
		toSerialize["azure_host_top99p_sum"] = o.AzureHostTop99pSum
	}
	if o.BillableIngestedBytesAggSum != nil {
		toSerialize["billable_ingested_bytes_agg_sum"] = o.BillableIngestedBytesAggSum
	}
	if o.ContainerAvgSum != nil {
		toSerialize["container_avg_sum"] = o.ContainerAvgSum
	}
	if o.ContainerHwmSum != nil {
		toSerialize["container_hwm_sum"] = o.ContainerHwmSum
	}
	if o.CustomTsSum != nil {
		toSerialize["custom_ts_sum"] = o.CustomTsSum
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.FargateTasksCountAvgSum != nil {
		toSerialize["fargate_tasks_count_avg_sum"] = o.FargateTasksCountAvgSum
	}
	if o.FargateTasksCountHwmSum != nil {
		toSerialize["fargate_tasks_count_hwm_sum"] = o.FargateTasksCountHwmSum
	}
	if o.GcpHostTop99pSum != nil {
		toSerialize["gcp_host_top99p_sum"] = o.GcpHostTop99pSum
	}
	if o.IndexedEventsCountAggSum != nil {
		toSerialize["indexed_events_count_agg_sum"] = o.IndexedEventsCountAggSum
	}
	if o.InfraHostTop99pSum != nil {
		toSerialize["infra_host_top99p_sum"] = o.InfraHostTop99pSum
	}
	if o.IngestedEventsBytesAggSum != nil {
		toSerialize["ingested_events_bytes_agg_sum"] = o.IngestedEventsBytesAggSum
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.NetflowIndexedEventsCountAggSum != nil {
		toSerialize["netflow_indexed_events_count_agg_sum"] = o.NetflowIndexedEventsCountAggSum
	}
	if o.NpmHostTop99pSum != nil {
		toSerialize["npm_host_top99p_sum"] = o.NpmHostTop99pSum
	}
	if o.RumSessionCountAggSum != nil {
		toSerialize["rum_session_count_agg_sum"] = o.RumSessionCountAggSum
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.SyntheticsBrowserCheckCallsCountAggSum != nil {
		toSerialize["synthetics_browser_check_calls_count_agg_sum"] = o.SyntheticsBrowserCheckCallsCountAggSum
	}
	if o.SyntheticsCheckCallsCountAggSum != nil {
		toSerialize["synthetics_check_calls_count_agg_sum"] = o.SyntheticsCheckCallsCountAggSum
	}
	if o.TraceSearchIndexedEventsCountAggSum != nil {
		toSerialize["trace_search_indexed_events_count_agg_sum"] = o.TraceSearchIndexedEventsCountAggSum
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageSummaryResponse struct {
	value *UsageSummaryResponse
	isSet bool
}

func (v NullableUsageSummaryResponse) Get() *UsageSummaryResponse {
	return v.value
}

func (v *NullableUsageSummaryResponse) Set(val *UsageSummaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageSummaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageSummaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageSummaryResponse(val *UsageSummaryResponse) *NullableUsageSummaryResponse {
	return &NullableUsageSummaryResponse{value: val, isSet: true}
}

func (v NullableUsageSummaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageSummaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
