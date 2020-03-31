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

// UsageSummaryDate struct for UsageSummaryDate
type UsageSummaryDate struct {
	// Shows the 99th percentile of all agent hosts over all hours in the current date for all orgs.
	AgentHostTop99p *int64 `json:"agent_host_top99p,omitempty"`
	// Shows the 99th percentile of all distinct APM hosts over all hours in the current date for all orgs.
	ApmHostTop99p *int64 `json:"apm_host_top99p,omitempty"`
	// Shows the 99th percentile of all AWS hosts over all hours in the current date for all orgs.
	AwsHostTop99p *int64 `json:"aws_host_top99p,omitempty"`
	// Shows the average of the number of functions that executed 1 or more times each hour in the current date for all orgs.
	AwsLambdaFuncCount *int64 `json:"aws_lambda_func_count,omitempty"`
	// Shows the sum of all AWS Labmda invocations over all hours in the current date for all orgs.
	AwsLambdaInvocationsSum *int64 `json:"aws_lambda_invocations_sum,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current date for all orgs.
	BillableIngestedBytesSum *int64 `json:"billable_ingested_bytes_sum,omitempty"`
	// Shows the average of all distinct containers over all hours in the current date for all orgs.
	ContainerAvg *int64 `json:"container_avg,omitempty"`
	// Shows the high watermark of all distinct containers over all hours in the current date for all orgs.
	ContainerHwm *int64 `json:"container_hwm,omitempty"`
	// Shows the average number of distinct custom metrics over all hours in the current date for all orgs.
	CustomTsAvg *int64 `json:"custom_ts_avg,omitempty"`
	// The date for the usage.
	Date *time.Time `json:"date,omitempty"`
	// Shows the high watermark of all Fargate tasks over all hours in the current date for all orgs.
	FargateTasksCountAvg *int64 `json:"fargate_tasks_count_avg,omitempty"`
	// Shows the average of all Fargate tasks over all hours in the current date for all orgs.
	FargateTasksCountHwm *int64 `json:"fargate_tasks_count_hwm,omitempty"`
	// Shows the 99th percentile of all GCP hosts over all hours in the current date for all orgs.
	GcpHostTop99p *int64 `json:"gcp_host_top99p,omitempty"`
	// Shows the sum of all log events indexed over all hours in the current date for all orgs.
	IndexedEventsCountSum *int64 `json:"indexed_events_count_sum,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for all orgs.
	InfraHostTop99p *int64 `json:"infra_host_top99p,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current date for all orgs.
	IngestedEventsBytesSum *int64 `json:"ingested_events_bytes_sum,omitempty"`
	// Shows the sum of all Network flows indexed over all hours in the current date for all orgs.
	NetflowIndexedEventsCountSum *int64 `json:"netflow_indexed_events_count_sum,omitempty"`
	// Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for all orgs.
	NpmHostTop99p *int64                 `json:"npm_host_top99p,omitempty"`
	Orgs          *[]UsageSummaryDateOrg `json:"orgs,omitempty"`
	// Shows the sum of all RUM Sessions over all hours in the current date for all orgs
	RumSessionCountSum *int64 `json:"rum_session_count_sum,omitempty"`
	// Shows the sum of all Synthetic browser tests over all hours in the current date for all orgs.
	SyntheticsBrowserCheckCallsCountSum *int64 `json:"synthetics_browser_check_calls_count_sum,omitempty"`
	// Shows the sum of all Synthetic API tests over all hours in the current date for all orgs.
	SyntheticsCheckCallsCountSum *int64 `json:"synthetics_check_calls_count_sum,omitempty"`
	// Shows the sum of all analyzed spans indexed over all hours in the current date for all orgs.
	TraceSearchIndexedEventsCountSum *int64 `json:"trace_search_indexed_events_count_sum,omitempty"`
}

// NewUsageSummaryDate instantiates a new UsageSummaryDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSummaryDate() *UsageSummaryDate {
	this := UsageSummaryDate{}
	return &this
}

// NewUsageSummaryDateWithDefaults instantiates a new UsageSummaryDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSummaryDateWithDefaults() *UsageSummaryDate {
	this := UsageSummaryDate{}
	return &this
}

// GetAgentHostTop99p returns the AgentHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAgentHostTop99p() int64 {
	if o == nil || o.AgentHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.AgentHostTop99p
}

// GetAgentHostTop99pOk returns a tuple with the AgentHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAgentHostTop99pOk() (*int64, bool) {
	if o == nil || o.AgentHostTop99p == nil {
		return nil, false
	}
	return o.AgentHostTop99p, true
}

// HasAgentHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAgentHostTop99p() bool {
	if o != nil && o.AgentHostTop99p != nil {
		return true
	}

	return false
}

// SetAgentHostTop99p gets a reference to the given int64 and assigns it to the AgentHostTop99p field.
func (o *UsageSummaryDate) SetAgentHostTop99p(v int64) {
	o.AgentHostTop99p = &v
}

// GetApmHostTop99p returns the ApmHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetApmHostTop99p() int64 {
	if o == nil || o.ApmHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.ApmHostTop99p
}

// GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetApmHostTop99pOk() (*int64, bool) {
	if o == nil || o.ApmHostTop99p == nil {
		return nil, false
	}
	return o.ApmHostTop99p, true
}

// HasApmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmHostTop99p() bool {
	if o != nil && o.ApmHostTop99p != nil {
		return true
	}

	return false
}

// SetApmHostTop99p gets a reference to the given int64 and assigns it to the ApmHostTop99p field.
func (o *UsageSummaryDate) SetApmHostTop99p(v int64) {
	o.ApmHostTop99p = &v
}

// GetAwsHostTop99p returns the AwsHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAwsHostTop99p() int64 {
	if o == nil || o.AwsHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.AwsHostTop99p
}

// GetAwsHostTop99pOk returns a tuple with the AwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAwsHostTop99pOk() (*int64, bool) {
	if o == nil || o.AwsHostTop99p == nil {
		return nil, false
	}
	return o.AwsHostTop99p, true
}

// HasAwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAwsHostTop99p() bool {
	if o != nil && o.AwsHostTop99p != nil {
		return true
	}

	return false
}

// SetAwsHostTop99p gets a reference to the given int64 and assigns it to the AwsHostTop99p field.
func (o *UsageSummaryDate) SetAwsHostTop99p(v int64) {
	o.AwsHostTop99p = &v
}

// GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAwsLambdaFuncCount() int64 {
	if o == nil || o.AwsLambdaFuncCount == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaFuncCount
}

// GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAwsLambdaFuncCountOk() (*int64, bool) {
	if o == nil || o.AwsLambdaFuncCount == nil {
		return nil, false
	}
	return o.AwsLambdaFuncCount, true
}

// HasAwsLambdaFuncCount returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAwsLambdaFuncCount() bool {
	if o != nil && o.AwsLambdaFuncCount != nil {
		return true
	}

	return false
}

// SetAwsLambdaFuncCount gets a reference to the given int64 and assigns it to the AwsLambdaFuncCount field.
func (o *UsageSummaryDate) SetAwsLambdaFuncCount(v int64) {
	o.AwsLambdaFuncCount = &v
}

// GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAwsLambdaInvocationsSum() int64 {
	if o == nil || o.AwsLambdaInvocationsSum == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaInvocationsSum
}

// GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAwsLambdaInvocationsSumOk() (*int64, bool) {
	if o == nil || o.AwsLambdaInvocationsSum == nil {
		return nil, false
	}
	return o.AwsLambdaInvocationsSum, true
}

// HasAwsLambdaInvocationsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAwsLambdaInvocationsSum() bool {
	if o != nil && o.AwsLambdaInvocationsSum != nil {
		return true
	}

	return false
}

// SetAwsLambdaInvocationsSum gets a reference to the given int64 and assigns it to the AwsLambdaInvocationsSum field.
func (o *UsageSummaryDate) SetAwsLambdaInvocationsSum(v int64) {
	o.AwsLambdaInvocationsSum = &v
}

// GetBillableIngestedBytesSum returns the BillableIngestedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetBillableIngestedBytesSum() int64 {
	if o == nil || o.BillableIngestedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.BillableIngestedBytesSum
}

// GetBillableIngestedBytesSumOk returns a tuple with the BillableIngestedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetBillableIngestedBytesSumOk() (*int64, bool) {
	if o == nil || o.BillableIngestedBytesSum == nil {
		return nil, false
	}
	return o.BillableIngestedBytesSum, true
}

// HasBillableIngestedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBillableIngestedBytesSum() bool {
	if o != nil && o.BillableIngestedBytesSum != nil {
		return true
	}

	return false
}

// SetBillableIngestedBytesSum gets a reference to the given int64 and assigns it to the BillableIngestedBytesSum field.
func (o *UsageSummaryDate) SetBillableIngestedBytesSum(v int64) {
	o.BillableIngestedBytesSum = &v
}

// GetContainerAvg returns the ContainerAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetContainerAvg() int64 {
	if o == nil || o.ContainerAvg == nil {
		var ret int64
		return ret
	}
	return *o.ContainerAvg
}

// GetContainerAvgOk returns a tuple with the ContainerAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetContainerAvgOk() (*int64, bool) {
	if o == nil || o.ContainerAvg == nil {
		return nil, false
	}
	return o.ContainerAvg, true
}

// HasContainerAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasContainerAvg() bool {
	if o != nil && o.ContainerAvg != nil {
		return true
	}

	return false
}

// SetContainerAvg gets a reference to the given int64 and assigns it to the ContainerAvg field.
func (o *UsageSummaryDate) SetContainerAvg(v int64) {
	o.ContainerAvg = &v
}

// GetContainerHwm returns the ContainerHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetContainerHwm() int64 {
	if o == nil || o.ContainerHwm == nil {
		var ret int64
		return ret
	}
	return *o.ContainerHwm
}

// GetContainerHwmOk returns a tuple with the ContainerHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetContainerHwmOk() (*int64, bool) {
	if o == nil || o.ContainerHwm == nil {
		return nil, false
	}
	return o.ContainerHwm, true
}

// HasContainerHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasContainerHwm() bool {
	if o != nil && o.ContainerHwm != nil {
		return true
	}

	return false
}

// SetContainerHwm gets a reference to the given int64 and assigns it to the ContainerHwm field.
func (o *UsageSummaryDate) SetContainerHwm(v int64) {
	o.ContainerHwm = &v
}

// GetCustomTsAvg returns the CustomTsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCustomTsAvg() int64 {
	if o == nil || o.CustomTsAvg == nil {
		var ret int64
		return ret
	}
	return *o.CustomTsAvg
}

// GetCustomTsAvgOk returns a tuple with the CustomTsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCustomTsAvgOk() (*int64, bool) {
	if o == nil || o.CustomTsAvg == nil {
		return nil, false
	}
	return o.CustomTsAvg, true
}

// HasCustomTsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCustomTsAvg() bool {
	if o != nil && o.CustomTsAvg != nil {
		return true
	}

	return false
}

// SetCustomTsAvg gets a reference to the given int64 and assigns it to the CustomTsAvg field.
func (o *UsageSummaryDate) SetCustomTsAvg(v int64) {
	o.CustomTsAvg = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *UsageSummaryDate) SetDate(v time.Time) {
	o.Date = &v
}

// GetFargateTasksCountAvg returns the FargateTasksCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFargateTasksCountAvg() int64 {
	if o == nil || o.FargateTasksCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountAvg
}

// GetFargateTasksCountAvgOk returns a tuple with the FargateTasksCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFargateTasksCountAvgOk() (*int64, bool) {
	if o == nil || o.FargateTasksCountAvg == nil {
		return nil, false
	}
	return o.FargateTasksCountAvg, true
}

// HasFargateTasksCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFargateTasksCountAvg() bool {
	if o != nil && o.FargateTasksCountAvg != nil {
		return true
	}

	return false
}

// SetFargateTasksCountAvg gets a reference to the given int64 and assigns it to the FargateTasksCountAvg field.
func (o *UsageSummaryDate) SetFargateTasksCountAvg(v int64) {
	o.FargateTasksCountAvg = &v
}

// GetFargateTasksCountHwm returns the FargateTasksCountHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFargateTasksCountHwm() int64 {
	if o == nil || o.FargateTasksCountHwm == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountHwm
}

// GetFargateTasksCountHwmOk returns a tuple with the FargateTasksCountHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFargateTasksCountHwmOk() (*int64, bool) {
	if o == nil || o.FargateTasksCountHwm == nil {
		return nil, false
	}
	return o.FargateTasksCountHwm, true
}

// HasFargateTasksCountHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFargateTasksCountHwm() bool {
	if o != nil && o.FargateTasksCountHwm != nil {
		return true
	}

	return false
}

// SetFargateTasksCountHwm gets a reference to the given int64 and assigns it to the FargateTasksCountHwm field.
func (o *UsageSummaryDate) SetFargateTasksCountHwm(v int64) {
	o.FargateTasksCountHwm = &v
}

// GetGcpHostTop99p returns the GcpHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetGcpHostTop99p() int64 {
	if o == nil || o.GcpHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.GcpHostTop99p
}

// GetGcpHostTop99pOk returns a tuple with the GcpHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetGcpHostTop99pOk() (*int64, bool) {
	if o == nil || o.GcpHostTop99p == nil {
		return nil, false
	}
	return o.GcpHostTop99p, true
}

// HasGcpHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasGcpHostTop99p() bool {
	if o != nil && o.GcpHostTop99p != nil {
		return true
	}

	return false
}

// SetGcpHostTop99p gets a reference to the given int64 and assigns it to the GcpHostTop99p field.
func (o *UsageSummaryDate) SetGcpHostTop99p(v int64) {
	o.GcpHostTop99p = &v
}

// GetIndexedEventsCountSum returns the IndexedEventsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIndexedEventsCountSum() int64 {
	if o == nil || o.IndexedEventsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCountSum
}

// GetIndexedEventsCountSumOk returns a tuple with the IndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil || o.IndexedEventsCountSum == nil {
		return nil, false
	}
	return o.IndexedEventsCountSum, true
}

// HasIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIndexedEventsCountSum() bool {
	if o != nil && o.IndexedEventsCountSum != nil {
		return true
	}

	return false
}

// SetIndexedEventsCountSum gets a reference to the given int64 and assigns it to the IndexedEventsCountSum field.
func (o *UsageSummaryDate) SetIndexedEventsCountSum(v int64) {
	o.IndexedEventsCountSum = &v
}

// GetInfraHostTop99p returns the InfraHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraHostTop99p() int64 {
	if o == nil || o.InfraHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostTop99p
}

// GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraHostTop99pOk() (*int64, bool) {
	if o == nil || o.InfraHostTop99p == nil {
		return nil, false
	}
	return o.InfraHostTop99p, true
}

// HasInfraHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraHostTop99p() bool {
	if o != nil && o.InfraHostTop99p != nil {
		return true
	}

	return false
}

// SetInfraHostTop99p gets a reference to the given int64 and assigns it to the InfraHostTop99p field.
func (o *UsageSummaryDate) SetInfraHostTop99p(v int64) {
	o.InfraHostTop99p = &v
}

// GetIngestedEventsBytesSum returns the IngestedEventsBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIngestedEventsBytesSum() int64 {
	if o == nil || o.IngestedEventsBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytesSum
}

// GetIngestedEventsBytesSumOk returns a tuple with the IngestedEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIngestedEventsBytesSumOk() (*int64, bool) {
	if o == nil || o.IngestedEventsBytesSum == nil {
		return nil, false
	}
	return o.IngestedEventsBytesSum, true
}

// HasIngestedEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIngestedEventsBytesSum() bool {
	if o != nil && o.IngestedEventsBytesSum != nil {
		return true
	}

	return false
}

// SetIngestedEventsBytesSum gets a reference to the given int64 and assigns it to the IngestedEventsBytesSum field.
func (o *UsageSummaryDate) SetIngestedEventsBytesSum(v int64) {
	o.IngestedEventsBytesSum = &v
}

// GetNetflowIndexedEventsCountSum returns the NetflowIndexedEventsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSum() int64 {
	if o == nil || o.NetflowIndexedEventsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.NetflowIndexedEventsCountSum
}

// GetNetflowIndexedEventsCountSumOk returns a tuple with the NetflowIndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil || o.NetflowIndexedEventsCountSum == nil {
		return nil, false
	}
	return o.NetflowIndexedEventsCountSum, true
}

// HasNetflowIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasNetflowIndexedEventsCountSum() bool {
	if o != nil && o.NetflowIndexedEventsCountSum != nil {
		return true
	}

	return false
}

// SetNetflowIndexedEventsCountSum gets a reference to the given int64 and assigns it to the NetflowIndexedEventsCountSum field.
func (o *UsageSummaryDate) SetNetflowIndexedEventsCountSum(v int64) {
	o.NetflowIndexedEventsCountSum = &v
}

// GetNpmHostTop99p returns the NpmHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetNpmHostTop99p() int64 {
	if o == nil || o.NpmHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.NpmHostTop99p
}

// GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetNpmHostTop99pOk() (*int64, bool) {
	if o == nil || o.NpmHostTop99p == nil {
		return nil, false
	}
	return o.NpmHostTop99p, true
}

// HasNpmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasNpmHostTop99p() bool {
	if o != nil && o.NpmHostTop99p != nil {
		return true
	}

	return false
}

// SetNpmHostTop99p gets a reference to the given int64 and assigns it to the NpmHostTop99p field.
func (o *UsageSummaryDate) SetNpmHostTop99p(v int64) {
	o.NpmHostTop99p = &v
}

// GetOrgs returns the Orgs field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetOrgs() []UsageSummaryDateOrg {
	if o == nil || o.Orgs == nil {
		var ret []UsageSummaryDateOrg
		return ret
	}
	return *o.Orgs
}

// GetOrgsOk returns a tuple with the Orgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetOrgsOk() (*[]UsageSummaryDateOrg, bool) {
	if o == nil || o.Orgs == nil {
		return nil, false
	}
	return o.Orgs, true
}

// HasOrgs returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOrgs() bool {
	if o != nil && o.Orgs != nil {
		return true
	}

	return false
}

// SetOrgs gets a reference to the given []UsageSummaryDateOrg and assigns it to the Orgs field.
func (o *UsageSummaryDate) SetOrgs(v []UsageSummaryDateOrg) {
	o.Orgs = &v
}

// GetRumSessionCountSum returns the RumSessionCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumSessionCountSum() int64 {
	if o == nil || o.RumSessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.RumSessionCountSum
}

// GetRumSessionCountSumOk returns a tuple with the RumSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumSessionCountSumOk() (*int64, bool) {
	if o == nil || o.RumSessionCountSum == nil {
		return nil, false
	}
	return o.RumSessionCountSum, true
}

// HasRumSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumSessionCountSum() bool {
	if o != nil && o.RumSessionCountSum != nil {
		return true
	}

	return false
}

// SetRumSessionCountSum gets a reference to the given int64 and assigns it to the RumSessionCountSum field.
func (o *UsageSummaryDate) SetRumSessionCountSum(v int64) {
	o.RumSessionCountSum = &v
}

// GetSyntheticsBrowserCheckCallsCountSum returns the SyntheticsBrowserCheckCallsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSum() int64 {
	if o == nil || o.SyntheticsBrowserCheckCallsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsBrowserCheckCallsCountSum
}

// GetSyntheticsBrowserCheckCallsCountSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsBrowserCheckCallsCountSum == nil {
		return nil, false
	}
	return o.SyntheticsBrowserCheckCallsCountSum, true
}

// HasSyntheticsBrowserCheckCallsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSyntheticsBrowserCheckCallsCountSum() bool {
	if o != nil && o.SyntheticsBrowserCheckCallsCountSum != nil {
		return true
	}

	return false
}

// SetSyntheticsBrowserCheckCallsCountSum gets a reference to the given int64 and assigns it to the SyntheticsBrowserCheckCallsCountSum field.
func (o *UsageSummaryDate) SetSyntheticsBrowserCheckCallsCountSum(v int64) {
	o.SyntheticsBrowserCheckCallsCountSum = &v
}

// GetSyntheticsCheckCallsCountSum returns the SyntheticsCheckCallsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSum() int64 {
	if o == nil || o.SyntheticsCheckCallsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsCheckCallsCountSum
}

// GetSyntheticsCheckCallsCountSumOk returns a tuple with the SyntheticsCheckCallsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsCheckCallsCountSum == nil {
		return nil, false
	}
	return o.SyntheticsCheckCallsCountSum, true
}

// HasSyntheticsCheckCallsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSyntheticsCheckCallsCountSum() bool {
	if o != nil && o.SyntheticsCheckCallsCountSum != nil {
		return true
	}

	return false
}

// SetSyntheticsCheckCallsCountSum gets a reference to the given int64 and assigns it to the SyntheticsCheckCallsCountSum field.
func (o *UsageSummaryDate) SetSyntheticsCheckCallsCountSum(v int64) {
	o.SyntheticsCheckCallsCountSum = &v
}

// GetTraceSearchIndexedEventsCountSum returns the TraceSearchIndexedEventsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSum() int64 {
	if o == nil || o.TraceSearchIndexedEventsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.TraceSearchIndexedEventsCountSum
}

// GetTraceSearchIndexedEventsCountSumOk returns a tuple with the TraceSearchIndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil || o.TraceSearchIndexedEventsCountSum == nil {
		return nil, false
	}
	return o.TraceSearchIndexedEventsCountSum, true
}

// HasTraceSearchIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasTraceSearchIndexedEventsCountSum() bool {
	if o != nil && o.TraceSearchIndexedEventsCountSum != nil {
		return true
	}

	return false
}

// SetTraceSearchIndexedEventsCountSum gets a reference to the given int64 and assigns it to the TraceSearchIndexedEventsCountSum field.
func (o *UsageSummaryDate) SetTraceSearchIndexedEventsCountSum(v int64) {
	o.TraceSearchIndexedEventsCountSum = &v
}

func (o UsageSummaryDate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentHostTop99p != nil {
		toSerialize["agent_host_top99p"] = o.AgentHostTop99p
	}
	if o.ApmHostTop99p != nil {
		toSerialize["apm_host_top99p"] = o.ApmHostTop99p
	}
	if o.AwsHostTop99p != nil {
		toSerialize["aws_host_top99p"] = o.AwsHostTop99p
	}
	if o.AwsLambdaFuncCount != nil {
		toSerialize["aws_lambda_func_count"] = o.AwsLambdaFuncCount
	}
	if o.AwsLambdaInvocationsSum != nil {
		toSerialize["aws_lambda_invocations_sum"] = o.AwsLambdaInvocationsSum
	}
	if o.BillableIngestedBytesSum != nil {
		toSerialize["billable_ingested_bytes_sum"] = o.BillableIngestedBytesSum
	}
	if o.ContainerAvg != nil {
		toSerialize["container_avg"] = o.ContainerAvg
	}
	if o.ContainerHwm != nil {
		toSerialize["container_hwm"] = o.ContainerHwm
	}
	if o.CustomTsAvg != nil {
		toSerialize["custom_ts_avg"] = o.CustomTsAvg
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.FargateTasksCountAvg != nil {
		toSerialize["fargate_tasks_count_avg"] = o.FargateTasksCountAvg
	}
	if o.FargateTasksCountHwm != nil {
		toSerialize["fargate_tasks_count_hwm"] = o.FargateTasksCountHwm
	}
	if o.GcpHostTop99p != nil {
		toSerialize["gcp_host_top99p"] = o.GcpHostTop99p
	}
	if o.IndexedEventsCountSum != nil {
		toSerialize["indexed_events_count_sum"] = o.IndexedEventsCountSum
	}
	if o.InfraHostTop99p != nil {
		toSerialize["infra_host_top99p"] = o.InfraHostTop99p
	}
	if o.IngestedEventsBytesSum != nil {
		toSerialize["ingested_events_bytes_sum"] = o.IngestedEventsBytesSum
	}
	if o.NetflowIndexedEventsCountSum != nil {
		toSerialize["netflow_indexed_events_count_sum"] = o.NetflowIndexedEventsCountSum
	}
	if o.NpmHostTop99p != nil {
		toSerialize["npm_host_top99p"] = o.NpmHostTop99p
	}
	if o.Orgs != nil {
		toSerialize["orgs"] = o.Orgs
	}
	if o.RumSessionCountSum != nil {
		toSerialize["rum_session_count_sum"] = o.RumSessionCountSum
	}
	if o.SyntheticsBrowserCheckCallsCountSum != nil {
		toSerialize["synthetics_browser_check_calls_count_sum"] = o.SyntheticsBrowserCheckCallsCountSum
	}
	if o.SyntheticsCheckCallsCountSum != nil {
		toSerialize["synthetics_check_calls_count_sum"] = o.SyntheticsCheckCallsCountSum
	}
	if o.TraceSearchIndexedEventsCountSum != nil {
		toSerialize["trace_search_indexed_events_count_sum"] = o.TraceSearchIndexedEventsCountSum
	}
	return json.Marshal(toSerialize)
}

type NullableUsageSummaryDate struct {
	value *UsageSummaryDate
	isSet bool
}

func (v NullableUsageSummaryDate) Get() *UsageSummaryDate {
	return v.value
}

func (v *NullableUsageSummaryDate) Set(val *UsageSummaryDate) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageSummaryDate) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageSummaryDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageSummaryDate(val *UsageSummaryDate) *NullableUsageSummaryDate {
	return &NullableUsageSummaryDate{value: val, isSet: true}
}

func (v NullableUsageSummaryDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageSummaryDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
