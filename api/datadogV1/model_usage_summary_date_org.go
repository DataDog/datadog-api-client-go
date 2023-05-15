// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSummaryDateOrg Global hourly report of all data billed by Datadog for a given organization.
type UsageSummaryDateOrg struct {
	// Shows the 99th percentile of all agent hosts over all hours in the current date for the given org.
	AgentHostTop99p datadog.NullableInt64 `json:"agent_host_top99p,omitempty"`
	// Shows the 99th percentile of all Azure app services using APM over all hours in the current date for the given org.
	ApmAzureAppServiceHostTop99p datadog.NullableInt64 `json:"apm_azure_app_service_host_top99p,omitempty"`
	// Shows the average of all APM ECS Fargate tasks over all hours in the current months for the given org.
	ApmFargateCountAvg datadog.NullableInt64 `json:"apm_fargate_count_avg,omitempty"`
	// Shows the 99th percentile of all distinct APM hosts over all hours in the current date for the given org.
	ApmHostTop99p datadog.NullableInt64 `json:"apm_host_top99p,omitempty"`
	// Shows the average of all Application Security Monitoring ECS Fargate tasks over all hours in the current months for the given org.
	AppsecFargateCountAvg datadog.NullableInt64 `json:"appsec_fargate_count_avg,omitempty"`
	// Shows the sum of all audit logs lines indexed over all hours in the current date for the given org.
	// Deprecated
	AuditLogsLinesIndexedSum datadog.NullableInt64 `json:"audit_logs_lines_indexed_sum,omitempty"`
	// Shows whether Audit Trail is enabled for the current date for the given org.
	AuditTrailEnabledHwm datadog.NullableInt64 `json:"audit_trail_enabled_hwm,omitempty"`
	// The average profiled task count for Fargate Profiling.
	AvgProfiledFargateTasks datadog.NullableInt64 `json:"avg_profiled_fargate_tasks,omitempty"`
	// Shows the 99th percentile of all AWS hosts over all hours in the current date for the given org.
	AwsHostTop99p datadog.NullableInt64 `json:"aws_host_top99p,omitempty"`
	// Shows the sum of all AWS Lambda invocations over all hours in the current date for the given org.
	AwsLambdaFuncCount *int64 `json:"aws_lambda_func_count,omitempty"`
	// Shows the sum of all AWS Lambda invocations over all hours in the current date for the given org.
	AwsLambdaInvocationsSum datadog.NullableInt64 `json:"aws_lambda_invocations_sum,omitempty"`
	// Shows the 99th percentile of all Azure app services over all hours in the current date for the given org.
	AzureAppServiceTop99p datadog.NullableInt64 `json:"azure_app_service_top99p,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current date for the given org.
	BillableIngestedBytesSum datadog.NullableInt64 `json:"billable_ingested_bytes_sum,omitempty"`
	// Shows the sum of all browser lite sessions over all hours in the current date for the given org.
	BrowserRumLiteSessionCountSum datadog.NullableInt64 `json:"browser_rum_lite_session_count_sum,omitempty"`
	// Shows the sum of all browser replay sessions over all hours in the current date for the given org.
	BrowserRumReplaySessionCountSum datadog.NullableInt64 `json:"browser_rum_replay_session_count_sum,omitempty"`
	// Shows the sum of all browser RUM units over all hours in the current date for the given org.
	BrowserRumUnitsSum datadog.NullableInt64 `json:"browser_rum_units_sum,omitempty"`
	// Shows the sum of all CI pipeline indexed spans over all hours in the current date for the given org.
	CiPipelineIndexedSpansSum datadog.NullableInt64 `json:"ci_pipeline_indexed_spans_sum,omitempty"`
	// Shows the sum of all CI test indexed spans over all hours in the current date for the given org.
	CiTestIndexedSpansSum datadog.NullableInt64 `json:"ci_test_indexed_spans_sum,omitempty"`
	// Shows the high-water mark of all CI visibility pipeline committers over all hours in the current date for the given org.
	CiVisibilityPipelineCommittersHwm datadog.NullableInt64 `json:"ci_visibility_pipeline_committers_hwm,omitempty"`
	// Shows the high-water mark of all CI visibility test committers over all hours in the current date for the given org.
	CiVisibilityTestCommittersHwm datadog.NullableInt64 `json:"ci_visibility_test_committers_hwm,omitempty"`
	// Host count average of Cloud Cost Management for the given date and given org.
	CloudCostManagementHostCountAvg *int64 `json:"cloud_cost_management_host_count_avg,omitempty"`
	// Shows the average of all distinct containers over all hours in the current date for the given org.
	ContainerAvg datadog.NullableInt64 `json:"container_avg,omitempty"`
	// Shows the average of containers without the Datadog Agent over all hours in the current date for the given organization.
	ContainerExclAgentAvg datadog.NullableInt64 `json:"container_excl_agent_avg,omitempty"`
	// Shows the high-water mark of all distinct containers over all hours in the current date for the given org.
	ContainerHwm datadog.NullableInt64 `json:"container_hwm,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management Azure app services hosts over all hours in the current date for the given org.
	CspmAasHostTop99p datadog.NullableInt64 `json:"cspm_aas_host_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management AWS hosts over all hours in the current date for the given org.
	CspmAwsHostTop99p *int64 `json:"cspm_aws_host_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management Azure hosts over all hours in the current date for the given org.
	CspmAzureHostTop99p datadog.NullableInt64 `json:"cspm_azure_host_top99p,omitempty"`
	// Shows the average number of Cloud Security Posture Management containers over all hours in the current date for the given org.
	CspmContainerAvg datadog.NullableInt64 `json:"cspm_container_avg,omitempty"`
	// Shows the high-water mark of Cloud Security Posture Management containers over all hours in the current date for the given org.
	CspmContainerHwm datadog.NullableInt64 `json:"cspm_container_hwm,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management GCP hosts over all hours in the current date for the given org.
	CspmGcpHostTop99p datadog.NullableInt64 `json:"cspm_gcp_host_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management hosts over all hours in the current date for the given org.
	CspmHostTop99p datadog.NullableInt64 `json:"cspm_host_top99p,omitempty"`
	// Shows the average number of distinct custom metrics over all hours in the current date for the given org.
	CustomTsAvg datadog.NullableInt64 `json:"custom_ts_avg,omitempty"`
	// Shows the average of all distinct Cloud Workload Security containers over all hours in the current date for the given org.
	CwsContainerCountAvg datadog.NullableInt64 `json:"cws_container_count_avg,omitempty"`
	// Shows the 99th percentile of all Cloud Workload Security hosts over all hours in the current date for the given org.
	CwsHostTop99p datadog.NullableInt64 `json:"cws_host_top99p,omitempty"`
	// Shows the 99th percentile of all Database Monitoring hosts over all hours in the current month for the given org.
	DbmHostTop99pSum datadog.NullableInt64 `json:"dbm_host_top99p_sum,omitempty"`
	// Shows the average of all distinct Database Monitoring normalized queries over all hours in the current month for the given org.
	DbmQueriesAvgSum datadog.NullableInt64 `json:"dbm_queries_avg_sum,omitempty"`
	// The average task count for Fargate.
	FargateTasksCountAvg datadog.NullableInt64 `json:"fargate_tasks_count_avg,omitempty"`
	// Shows the high-water mark of all Fargate tasks over all hours in the current date for the given org.
	FargateTasksCountHwm datadog.NullableInt64 `json:"fargate_tasks_count_hwm,omitempty"`
	// Shows the sum of all log bytes forwarded over all hours in the current date for the given org.
	ForwardingEventsBytesSum datadog.NullableInt64 `json:"forwarding_events_bytes_sum,omitempty"`
	// Shows the 99th percentile of all GCP hosts over all hours in the current date for the given org.
	GcpHostTop99p datadog.NullableInt64 `json:"gcp_host_top99p,omitempty"`
	// Shows the 99th percentile of all Heroku dynos over all hours in the current date for the given org.
	HerokuHostTop99p datadog.NullableInt64 `json:"heroku_host_top99p,omitempty"`
	// The organization id.
	Id *string `json:"id,omitempty"`
	// Shows the high-water mark of incident management monthly active users over all hours in the current date for the given org.
	IncidentManagementMonthlyActiveUsersHwm datadog.NullableInt64 `json:"incident_management_monthly_active_users_hwm,omitempty"`
	// Shows the sum of all log events indexed over all hours in the current date for the given org.
	IndexedEventsCountSum datadog.NullableInt64 `json:"indexed_events_count_sum,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for the given org.
	InfraHostTop99p datadog.NullableInt64 `json:"infra_host_top99p,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current date for the given org.
	IngestedEventsBytesSum datadog.NullableInt64 `json:"ingested_events_bytes_sum,omitempty"`
	// Shows the sum of all IoT devices over all hours in the current date for the given org.
	IotDeviceAggSum datadog.NullableInt64 `json:"iot_device_agg_sum,omitempty"`
	// Shows the 99th percentile of all IoT devices over all hours in the current date for the given org.
	IotDeviceTop99pSum datadog.NullableInt64 `json:"iot_device_top99p_sum,omitempty"`
	// Shows the sum of all mobile lite sessions over all hours in the current date for the given org.
	MobileRumLiteSessionCountSum datadog.NullableInt64 `json:"mobile_rum_lite_session_count_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on Android over all hours in the current date for the given org.
	MobileRumSessionCountAndroidSum datadog.NullableInt64 `json:"mobile_rum_session_count_android_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on Flutter over all hours in the current date for the given org.
	MobileRumSessionCountFlutterSum datadog.NullableInt64 `json:"mobile_rum_session_count_flutter_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on iOS over all hours in the current date for the given org.
	MobileRumSessionCountIosSum datadog.NullableInt64 `json:"mobile_rum_session_count_ios_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on React Native over all hours in the current date for the given org.
	MobileRumSessionCountReactnativeSum datadog.NullableInt64 `json:"mobile_rum_session_count_reactnative_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions over all hours in the current date for the given org.
	MobileRumSessionCountSum datadog.NullableInt64 `json:"mobile_rum_session_count_sum,omitempty"`
	// Shows the sum of all mobile RUM units over all hours in the current date for the given org.
	MobileRumUnitsSum datadog.NullableInt64 `json:"mobile_rum_units_sum,omitempty"`
	// The organization name.
	Name *string `json:"name,omitempty"`
	// Shows the sum of all Network flows indexed over all hours in the current date for the given org.
	NetflowIndexedEventsCountSum datadog.NullableInt64 `json:"netflow_indexed_events_count_sum,omitempty"`
	// Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for the given org.
	NpmHostTop99p datadog.NullableInt64 `json:"npm_host_top99p,omitempty"`
	// Sum of all observability pipelines bytes processed over all hours in the current date for the given org.
	ObservabilityPipelinesBytesProcessedSum datadog.NullableInt64 `json:"observability_pipelines_bytes_processed_sum,omitempty"`
	// Sum of all online archived events over all hours in the current date for the given org.
	OnlineArchiveEventsCountSum datadog.NullableInt64 `json:"online_archive_events_count_sum,omitempty"`
	// Shows the 99th percentile of APM hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for the given org.
	OpentelemetryApmHostTop99p datadog.NullableInt64 `json:"opentelemetry_apm_host_top99p,omitempty"`
	// Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for the given org.
	OpentelemetryHostTop99p datadog.NullableInt64 `json:"opentelemetry_host_top99p,omitempty"`
	// Shows the 99th percentile of all profiled hosts over all hours in the current date for the given org.
	ProfilingHostTop99p datadog.NullableInt64 `json:"profiling_host_top99p,omitempty"`
	// The organization public id.
	PublicId *string `json:"public_id,omitempty"`
	// The region of the organization.
	Region *string `json:"region,omitempty"`
	// Shows the sum of all mobile sessions and all browser lite and legacy sessions over all hours in the current date for the given org.
	RumBrowserAndMobileSessionCount datadog.NullableInt64 `json:"rum_browser_and_mobile_session_count,omitempty"`
	// Shows the sum of all browser RUM Lite Sessions over all hours in the current date for the given org.
	RumSessionCountSum datadog.NullableInt64 `json:"rum_session_count_sum,omitempty"`
	// Shows the sum of RUM Sessions (browser and mobile) over all hours in the current date for the given org.
	RumTotalSessionCountSum datadog.NullableInt64 `json:"rum_total_session_count_sum,omitempty"`
	// Shows the sum of all browser and mobile RUM units over all hours in the current date for the given org.
	RumUnitsSum datadog.NullableInt64 `json:"rum_units_sum,omitempty"`
	// Sum of all APM bytes scanned with sensitive data scanner over all hours in the current date for the given org.
	SdsApmScannedBytesSum datadog.NullableInt64 `json:"sds_apm_scanned_bytes_sum,omitempty"`
	// Sum of all event stream events bytes scanned with sensitive data scanner over all hours in the current date for the given org.
	SdsEventsScannedBytesSum datadog.NullableInt64 `json:"sds_events_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned of logs usage by the Sensitive Data Scanner over all hours in the current month for the given org.
	SdsLogsScannedBytesSum datadog.NullableInt64 `json:"sds_logs_scanned_bytes_sum,omitempty"`
	// Sum of all RUM bytes scanned with sensitive data scanner over all hours in the current date for the given org.
	SdsRumScannedBytesSum datadog.NullableInt64 `json:"sds_rum_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned across all usage types by the Sensitive Data Scanner over all hours in the current month for the given org.
	SdsTotalScannedBytesSum datadog.NullableInt64 `json:"sds_total_scanned_bytes_sum,omitempty"`
	// Shows the sum of all Synthetic browser tests over all hours in the current date for the given org.
	SyntheticsBrowserCheckCallsCountSum datadog.NullableInt64 `json:"synthetics_browser_check_calls_count_sum,omitempty"`
	// Shows the sum of all Synthetic API tests over all hours in the current date for the given org.
	SyntheticsCheckCallsCountSum datadog.NullableInt64 `json:"synthetics_check_calls_count_sum,omitempty"`
	// Shows the high-water mark of used synthetics parallel testing slots over all hours in the current date for the given org.
	SyntheticsParallelTestingMaxSlotsHwm datadog.NullableInt64 `json:"synthetics_parallel_testing_max_slots_hwm,omitempty"`
	// Shows the sum of all Indexed Spans indexed over all hours in the current date for the given org.
	TraceSearchIndexedEventsCountSum datadog.NullableInt64 `json:"trace_search_indexed_events_count_sum,omitempty"`
	// Shows the sum of all ingested APM span bytes over all hours in the current date for the given org.
	TwolIngestedEventsBytesSum datadog.NullableInt64 `json:"twol_ingested_events_bytes_sum,omitempty"`
	// Shows the 99th percentile of all Universal Service Monitoring hosts over all hours in the current date for the given org.
	UniversalServiceMonitoringHostTop99p datadog.NullableInt64 `json:"universal_service_monitoring_host_top99p,omitempty"`
	// Shows the 99th percentile of all vSphere hosts over all hours in the current date for the given org.
	VsphereHostTop99p datadog.NullableInt64 `json:"vsphere_host_top99p,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageSummaryDateOrg instantiates a new UsageSummaryDateOrg object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSummaryDateOrg() *UsageSummaryDateOrg {
	this := UsageSummaryDateOrg{}
	return &this
}

// NewUsageSummaryDateOrgWithDefaults instantiates a new UsageSummaryDateOrg object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSummaryDateOrgWithDefaults() *UsageSummaryDateOrg {
	this := UsageSummaryDateOrg{}
	return &this
}

// GetAgentHostTop99p returns the AgentHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetAgentHostTop99p() int64 {
	if o == nil || o.AgentHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AgentHostTop99p.Get()
}

// GetAgentHostTop99pOk returns a tuple with the AgentHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetAgentHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentHostTop99p.Get(), o.AgentHostTop99p.IsSet()
}

// HasAgentHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasAgentHostTop99p() bool {
	return o != nil && o.AgentHostTop99p.IsSet()
}

// SetAgentHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the AgentHostTop99p field.
func (o *UsageSummaryDateOrg) SetAgentHostTop99p(v int64) {
	o.AgentHostTop99p.Set(&v)
}

// SetAgentHostTop99pNil sets the value for AgentHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetAgentHostTop99pNil() {
	o.AgentHostTop99p.Set(nil)
}

// UnsetAgentHostTop99p ensures that no value is present for AgentHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetAgentHostTop99p() {
	o.AgentHostTop99p.Unset()
}

// GetApmAzureAppServiceHostTop99p returns the ApmAzureAppServiceHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetApmAzureAppServiceHostTop99p() int64 {
	if o == nil || o.ApmAzureAppServiceHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmAzureAppServiceHostTop99p.Get()
}

// GetApmAzureAppServiceHostTop99pOk returns a tuple with the ApmAzureAppServiceHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetApmAzureAppServiceHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmAzureAppServiceHostTop99p.Get(), o.ApmAzureAppServiceHostTop99p.IsSet()
}

// HasApmAzureAppServiceHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasApmAzureAppServiceHostTop99p() bool {
	return o != nil && o.ApmAzureAppServiceHostTop99p.IsSet()
}

// SetApmAzureAppServiceHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the ApmAzureAppServiceHostTop99p field.
func (o *UsageSummaryDateOrg) SetApmAzureAppServiceHostTop99p(v int64) {
	o.ApmAzureAppServiceHostTop99p.Set(&v)
}

// SetApmAzureAppServiceHostTop99pNil sets the value for ApmAzureAppServiceHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetApmAzureAppServiceHostTop99pNil() {
	o.ApmAzureAppServiceHostTop99p.Set(nil)
}

// UnsetApmAzureAppServiceHostTop99p ensures that no value is present for ApmAzureAppServiceHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetApmAzureAppServiceHostTop99p() {
	o.ApmAzureAppServiceHostTop99p.Unset()
}

// GetApmFargateCountAvg returns the ApmFargateCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetApmFargateCountAvg() int64 {
	if o == nil || o.ApmFargateCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmFargateCountAvg.Get()
}

// GetApmFargateCountAvgOk returns a tuple with the ApmFargateCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetApmFargateCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmFargateCountAvg.Get(), o.ApmFargateCountAvg.IsSet()
}

// HasApmFargateCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasApmFargateCountAvg() bool {
	return o != nil && o.ApmFargateCountAvg.IsSet()
}

// SetApmFargateCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the ApmFargateCountAvg field.
func (o *UsageSummaryDateOrg) SetApmFargateCountAvg(v int64) {
	o.ApmFargateCountAvg.Set(&v)
}

// SetApmFargateCountAvgNil sets the value for ApmFargateCountAvg to be an explicit nil.
func (o *UsageSummaryDateOrg) SetApmFargateCountAvgNil() {
	o.ApmFargateCountAvg.Set(nil)
}

// UnsetApmFargateCountAvg ensures that no value is present for ApmFargateCountAvg, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetApmFargateCountAvg() {
	o.ApmFargateCountAvg.Unset()
}

// GetApmHostTop99p returns the ApmHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetApmHostTop99p() int64 {
	if o == nil || o.ApmHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmHostTop99p.Get()
}

// GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetApmHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmHostTop99p.Get(), o.ApmHostTop99p.IsSet()
}

// HasApmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasApmHostTop99p() bool {
	return o != nil && o.ApmHostTop99p.IsSet()
}

// SetApmHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the ApmHostTop99p field.
func (o *UsageSummaryDateOrg) SetApmHostTop99p(v int64) {
	o.ApmHostTop99p.Set(&v)
}

// SetApmHostTop99pNil sets the value for ApmHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetApmHostTop99pNil() {
	o.ApmHostTop99p.Set(nil)
}

// UnsetApmHostTop99p ensures that no value is present for ApmHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetApmHostTop99p() {
	o.ApmHostTop99p.Unset()
}

// GetAppsecFargateCountAvg returns the AppsecFargateCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetAppsecFargateCountAvg() int64 {
	if o == nil || o.AppsecFargateCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AppsecFargateCountAvg.Get()
}

// GetAppsecFargateCountAvgOk returns a tuple with the AppsecFargateCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetAppsecFargateCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecFargateCountAvg.Get(), o.AppsecFargateCountAvg.IsSet()
}

// HasAppsecFargateCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasAppsecFargateCountAvg() bool {
	return o != nil && o.AppsecFargateCountAvg.IsSet()
}

// SetAppsecFargateCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the AppsecFargateCountAvg field.
func (o *UsageSummaryDateOrg) SetAppsecFargateCountAvg(v int64) {
	o.AppsecFargateCountAvg.Set(&v)
}

// SetAppsecFargateCountAvgNil sets the value for AppsecFargateCountAvg to be an explicit nil.
func (o *UsageSummaryDateOrg) SetAppsecFargateCountAvgNil() {
	o.AppsecFargateCountAvg.Set(nil)
}

// UnsetAppsecFargateCountAvg ensures that no value is present for AppsecFargateCountAvg, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetAppsecFargateCountAvg() {
	o.AppsecFargateCountAvg.Unset()
}

// GetAuditLogsLinesIndexedSum returns the AuditLogsLinesIndexedSum field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *UsageSummaryDateOrg) GetAuditLogsLinesIndexedSum() int64 {
	if o == nil || o.AuditLogsLinesIndexedSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AuditLogsLinesIndexedSum.Get()
}

// GetAuditLogsLinesIndexedSumOk returns a tuple with the AuditLogsLinesIndexedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
// Deprecated
func (o *UsageSummaryDateOrg) GetAuditLogsLinesIndexedSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditLogsLinesIndexedSum.Get(), o.AuditLogsLinesIndexedSum.IsSet()
}

// HasAuditLogsLinesIndexedSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasAuditLogsLinesIndexedSum() bool {
	return o != nil && o.AuditLogsLinesIndexedSum.IsSet()
}

// SetAuditLogsLinesIndexedSum gets a reference to the given datadog.NullableInt64 and assigns it to the AuditLogsLinesIndexedSum field.
// Deprecated
func (o *UsageSummaryDateOrg) SetAuditLogsLinesIndexedSum(v int64) {
	o.AuditLogsLinesIndexedSum.Set(&v)
}

// SetAuditLogsLinesIndexedSumNil sets the value for AuditLogsLinesIndexedSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetAuditLogsLinesIndexedSumNil() {
	o.AuditLogsLinesIndexedSum.Set(nil)
}

// UnsetAuditLogsLinesIndexedSum ensures that no value is present for AuditLogsLinesIndexedSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetAuditLogsLinesIndexedSum() {
	o.AuditLogsLinesIndexedSum.Unset()
}

// GetAuditTrailEnabledHwm returns the AuditTrailEnabledHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetAuditTrailEnabledHwm() int64 {
	if o == nil || o.AuditTrailEnabledHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AuditTrailEnabledHwm.Get()
}

// GetAuditTrailEnabledHwmOk returns a tuple with the AuditTrailEnabledHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetAuditTrailEnabledHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditTrailEnabledHwm.Get(), o.AuditTrailEnabledHwm.IsSet()
}

// HasAuditTrailEnabledHwm returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasAuditTrailEnabledHwm() bool {
	return o != nil && o.AuditTrailEnabledHwm.IsSet()
}

// SetAuditTrailEnabledHwm gets a reference to the given datadog.NullableInt64 and assigns it to the AuditTrailEnabledHwm field.
func (o *UsageSummaryDateOrg) SetAuditTrailEnabledHwm(v int64) {
	o.AuditTrailEnabledHwm.Set(&v)
}

// SetAuditTrailEnabledHwmNil sets the value for AuditTrailEnabledHwm to be an explicit nil.
func (o *UsageSummaryDateOrg) SetAuditTrailEnabledHwmNil() {
	o.AuditTrailEnabledHwm.Set(nil)
}

// UnsetAuditTrailEnabledHwm ensures that no value is present for AuditTrailEnabledHwm, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetAuditTrailEnabledHwm() {
	o.AuditTrailEnabledHwm.Unset()
}

// GetAvgProfiledFargateTasks returns the AvgProfiledFargateTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetAvgProfiledFargateTasks() int64 {
	if o == nil || o.AvgProfiledFargateTasks.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AvgProfiledFargateTasks.Get()
}

// GetAvgProfiledFargateTasksOk returns a tuple with the AvgProfiledFargateTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetAvgProfiledFargateTasksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvgProfiledFargateTasks.Get(), o.AvgProfiledFargateTasks.IsSet()
}

// HasAvgProfiledFargateTasks returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasAvgProfiledFargateTasks() bool {
	return o != nil && o.AvgProfiledFargateTasks.IsSet()
}

// SetAvgProfiledFargateTasks gets a reference to the given datadog.NullableInt64 and assigns it to the AvgProfiledFargateTasks field.
func (o *UsageSummaryDateOrg) SetAvgProfiledFargateTasks(v int64) {
	o.AvgProfiledFargateTasks.Set(&v)
}

// SetAvgProfiledFargateTasksNil sets the value for AvgProfiledFargateTasks to be an explicit nil.
func (o *UsageSummaryDateOrg) SetAvgProfiledFargateTasksNil() {
	o.AvgProfiledFargateTasks.Set(nil)
}

// UnsetAvgProfiledFargateTasks ensures that no value is present for AvgProfiledFargateTasks, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetAvgProfiledFargateTasks() {
	o.AvgProfiledFargateTasks.Unset()
}

// GetAwsHostTop99p returns the AwsHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetAwsHostTop99p() int64 {
	if o == nil || o.AwsHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AwsHostTop99p.Get()
}

// GetAwsHostTop99pOk returns a tuple with the AwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetAwsHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsHostTop99p.Get(), o.AwsHostTop99p.IsSet()
}

// HasAwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasAwsHostTop99p() bool {
	return o != nil && o.AwsHostTop99p.IsSet()
}

// SetAwsHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the AwsHostTop99p field.
func (o *UsageSummaryDateOrg) SetAwsHostTop99p(v int64) {
	o.AwsHostTop99p.Set(&v)
}

// SetAwsHostTop99pNil sets the value for AwsHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetAwsHostTop99pNil() {
	o.AwsHostTop99p.Set(nil)
}

// UnsetAwsHostTop99p ensures that no value is present for AwsHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetAwsHostTop99p() {
	o.AwsHostTop99p.Unset()
}

// GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field value if set, zero value otherwise.
func (o *UsageSummaryDateOrg) GetAwsLambdaFuncCount() int64 {
	if o == nil || o.AwsLambdaFuncCount == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaFuncCount
}

// GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDateOrg) GetAwsLambdaFuncCountOk() (*int64, bool) {
	if o == nil || o.AwsLambdaFuncCount == nil {
		return nil, false
	}
	return o.AwsLambdaFuncCount, true
}

// HasAwsLambdaFuncCount returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasAwsLambdaFuncCount() bool {
	return o != nil && o.AwsLambdaFuncCount != nil
}

// SetAwsLambdaFuncCount gets a reference to the given int64 and assigns it to the AwsLambdaFuncCount field.
func (o *UsageSummaryDateOrg) SetAwsLambdaFuncCount(v int64) {
	o.AwsLambdaFuncCount = &v
}

// GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetAwsLambdaInvocationsSum() int64 {
	if o == nil || o.AwsLambdaInvocationsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaInvocationsSum.Get()
}

// GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetAwsLambdaInvocationsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsLambdaInvocationsSum.Get(), o.AwsLambdaInvocationsSum.IsSet()
}

// HasAwsLambdaInvocationsSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasAwsLambdaInvocationsSum() bool {
	return o != nil && o.AwsLambdaInvocationsSum.IsSet()
}

// SetAwsLambdaInvocationsSum gets a reference to the given datadog.NullableInt64 and assigns it to the AwsLambdaInvocationsSum field.
func (o *UsageSummaryDateOrg) SetAwsLambdaInvocationsSum(v int64) {
	o.AwsLambdaInvocationsSum.Set(&v)
}

// SetAwsLambdaInvocationsSumNil sets the value for AwsLambdaInvocationsSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetAwsLambdaInvocationsSumNil() {
	o.AwsLambdaInvocationsSum.Set(nil)
}

// UnsetAwsLambdaInvocationsSum ensures that no value is present for AwsLambdaInvocationsSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetAwsLambdaInvocationsSum() {
	o.AwsLambdaInvocationsSum.Unset()
}

// GetAzureAppServiceTop99p returns the AzureAppServiceTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetAzureAppServiceTop99p() int64 {
	if o == nil || o.AzureAppServiceTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AzureAppServiceTop99p.Get()
}

// GetAzureAppServiceTop99pOk returns a tuple with the AzureAppServiceTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetAzureAppServiceTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureAppServiceTop99p.Get(), o.AzureAppServiceTop99p.IsSet()
}

// HasAzureAppServiceTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasAzureAppServiceTop99p() bool {
	return o != nil && o.AzureAppServiceTop99p.IsSet()
}

// SetAzureAppServiceTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the AzureAppServiceTop99p field.
func (o *UsageSummaryDateOrg) SetAzureAppServiceTop99p(v int64) {
	o.AzureAppServiceTop99p.Set(&v)
}

// SetAzureAppServiceTop99pNil sets the value for AzureAppServiceTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetAzureAppServiceTop99pNil() {
	o.AzureAppServiceTop99p.Set(nil)
}

// UnsetAzureAppServiceTop99p ensures that no value is present for AzureAppServiceTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetAzureAppServiceTop99p() {
	o.AzureAppServiceTop99p.Unset()
}

// GetBillableIngestedBytesSum returns the BillableIngestedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetBillableIngestedBytesSum() int64 {
	if o == nil || o.BillableIngestedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BillableIngestedBytesSum.Get()
}

// GetBillableIngestedBytesSumOk returns a tuple with the BillableIngestedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetBillableIngestedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillableIngestedBytesSum.Get(), o.BillableIngestedBytesSum.IsSet()
}

// HasBillableIngestedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasBillableIngestedBytesSum() bool {
	return o != nil && o.BillableIngestedBytesSum.IsSet()
}

// SetBillableIngestedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the BillableIngestedBytesSum field.
func (o *UsageSummaryDateOrg) SetBillableIngestedBytesSum(v int64) {
	o.BillableIngestedBytesSum.Set(&v)
}

// SetBillableIngestedBytesSumNil sets the value for BillableIngestedBytesSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetBillableIngestedBytesSumNil() {
	o.BillableIngestedBytesSum.Set(nil)
}

// UnsetBillableIngestedBytesSum ensures that no value is present for BillableIngestedBytesSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetBillableIngestedBytesSum() {
	o.BillableIngestedBytesSum.Unset()
}

// GetBrowserRumLiteSessionCountSum returns the BrowserRumLiteSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetBrowserRumLiteSessionCountSum() int64 {
	if o == nil || o.BrowserRumLiteSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumLiteSessionCountSum.Get()
}

// GetBrowserRumLiteSessionCountSumOk returns a tuple with the BrowserRumLiteSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetBrowserRumLiteSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumLiteSessionCountSum.Get(), o.BrowserRumLiteSessionCountSum.IsSet()
}

// HasBrowserRumLiteSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasBrowserRumLiteSessionCountSum() bool {
	return o != nil && o.BrowserRumLiteSessionCountSum.IsSet()
}

// SetBrowserRumLiteSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumLiteSessionCountSum field.
func (o *UsageSummaryDateOrg) SetBrowserRumLiteSessionCountSum(v int64) {
	o.BrowserRumLiteSessionCountSum.Set(&v)
}

// SetBrowserRumLiteSessionCountSumNil sets the value for BrowserRumLiteSessionCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetBrowserRumLiteSessionCountSumNil() {
	o.BrowserRumLiteSessionCountSum.Set(nil)
}

// UnsetBrowserRumLiteSessionCountSum ensures that no value is present for BrowserRumLiteSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetBrowserRumLiteSessionCountSum() {
	o.BrowserRumLiteSessionCountSum.Unset()
}

// GetBrowserRumReplaySessionCountSum returns the BrowserRumReplaySessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetBrowserRumReplaySessionCountSum() int64 {
	if o == nil || o.BrowserRumReplaySessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumReplaySessionCountSum.Get()
}

// GetBrowserRumReplaySessionCountSumOk returns a tuple with the BrowserRumReplaySessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetBrowserRumReplaySessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumReplaySessionCountSum.Get(), o.BrowserRumReplaySessionCountSum.IsSet()
}

// HasBrowserRumReplaySessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasBrowserRumReplaySessionCountSum() bool {
	return o != nil && o.BrowserRumReplaySessionCountSum.IsSet()
}

// SetBrowserRumReplaySessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumReplaySessionCountSum field.
func (o *UsageSummaryDateOrg) SetBrowserRumReplaySessionCountSum(v int64) {
	o.BrowserRumReplaySessionCountSum.Set(&v)
}

// SetBrowserRumReplaySessionCountSumNil sets the value for BrowserRumReplaySessionCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetBrowserRumReplaySessionCountSumNil() {
	o.BrowserRumReplaySessionCountSum.Set(nil)
}

// UnsetBrowserRumReplaySessionCountSum ensures that no value is present for BrowserRumReplaySessionCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetBrowserRumReplaySessionCountSum() {
	o.BrowserRumReplaySessionCountSum.Unset()
}

// GetBrowserRumUnitsSum returns the BrowserRumUnitsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetBrowserRumUnitsSum() int64 {
	if o == nil || o.BrowserRumUnitsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumUnitsSum.Get()
}

// GetBrowserRumUnitsSumOk returns a tuple with the BrowserRumUnitsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetBrowserRumUnitsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumUnitsSum.Get(), o.BrowserRumUnitsSum.IsSet()
}

// HasBrowserRumUnitsSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasBrowserRumUnitsSum() bool {
	return o != nil && o.BrowserRumUnitsSum.IsSet()
}

// SetBrowserRumUnitsSum gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumUnitsSum field.
func (o *UsageSummaryDateOrg) SetBrowserRumUnitsSum(v int64) {
	o.BrowserRumUnitsSum.Set(&v)
}

// SetBrowserRumUnitsSumNil sets the value for BrowserRumUnitsSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetBrowserRumUnitsSumNil() {
	o.BrowserRumUnitsSum.Set(nil)
}

// UnsetBrowserRumUnitsSum ensures that no value is present for BrowserRumUnitsSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetBrowserRumUnitsSum() {
	o.BrowserRumUnitsSum.Unset()
}

// GetCiPipelineIndexedSpansSum returns the CiPipelineIndexedSpansSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCiPipelineIndexedSpansSum() int64 {
	if o == nil || o.CiPipelineIndexedSpansSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiPipelineIndexedSpansSum.Get()
}

// GetCiPipelineIndexedSpansSumOk returns a tuple with the CiPipelineIndexedSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCiPipelineIndexedSpansSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiPipelineIndexedSpansSum.Get(), o.CiPipelineIndexedSpansSum.IsSet()
}

// HasCiPipelineIndexedSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCiPipelineIndexedSpansSum() bool {
	return o != nil && o.CiPipelineIndexedSpansSum.IsSet()
}

// SetCiPipelineIndexedSpansSum gets a reference to the given datadog.NullableInt64 and assigns it to the CiPipelineIndexedSpansSum field.
func (o *UsageSummaryDateOrg) SetCiPipelineIndexedSpansSum(v int64) {
	o.CiPipelineIndexedSpansSum.Set(&v)
}

// SetCiPipelineIndexedSpansSumNil sets the value for CiPipelineIndexedSpansSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCiPipelineIndexedSpansSumNil() {
	o.CiPipelineIndexedSpansSum.Set(nil)
}

// UnsetCiPipelineIndexedSpansSum ensures that no value is present for CiPipelineIndexedSpansSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCiPipelineIndexedSpansSum() {
	o.CiPipelineIndexedSpansSum.Unset()
}

// GetCiTestIndexedSpansSum returns the CiTestIndexedSpansSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCiTestIndexedSpansSum() int64 {
	if o == nil || o.CiTestIndexedSpansSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiTestIndexedSpansSum.Get()
}

// GetCiTestIndexedSpansSumOk returns a tuple with the CiTestIndexedSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCiTestIndexedSpansSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiTestIndexedSpansSum.Get(), o.CiTestIndexedSpansSum.IsSet()
}

// HasCiTestIndexedSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCiTestIndexedSpansSum() bool {
	return o != nil && o.CiTestIndexedSpansSum.IsSet()
}

// SetCiTestIndexedSpansSum gets a reference to the given datadog.NullableInt64 and assigns it to the CiTestIndexedSpansSum field.
func (o *UsageSummaryDateOrg) SetCiTestIndexedSpansSum(v int64) {
	o.CiTestIndexedSpansSum.Set(&v)
}

// SetCiTestIndexedSpansSumNil sets the value for CiTestIndexedSpansSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCiTestIndexedSpansSumNil() {
	o.CiTestIndexedSpansSum.Set(nil)
}

// UnsetCiTestIndexedSpansSum ensures that no value is present for CiTestIndexedSpansSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCiTestIndexedSpansSum() {
	o.CiTestIndexedSpansSum.Unset()
}

// GetCiVisibilityPipelineCommittersHwm returns the CiVisibilityPipelineCommittersHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCiVisibilityPipelineCommittersHwm() int64 {
	if o == nil || o.CiVisibilityPipelineCommittersHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityPipelineCommittersHwm.Get()
}

// GetCiVisibilityPipelineCommittersHwmOk returns a tuple with the CiVisibilityPipelineCommittersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCiVisibilityPipelineCommittersHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiVisibilityPipelineCommittersHwm.Get(), o.CiVisibilityPipelineCommittersHwm.IsSet()
}

// HasCiVisibilityPipelineCommittersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCiVisibilityPipelineCommittersHwm() bool {
	return o != nil && o.CiVisibilityPipelineCommittersHwm.IsSet()
}

// SetCiVisibilityPipelineCommittersHwm gets a reference to the given datadog.NullableInt64 and assigns it to the CiVisibilityPipelineCommittersHwm field.
func (o *UsageSummaryDateOrg) SetCiVisibilityPipelineCommittersHwm(v int64) {
	o.CiVisibilityPipelineCommittersHwm.Set(&v)
}

// SetCiVisibilityPipelineCommittersHwmNil sets the value for CiVisibilityPipelineCommittersHwm to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCiVisibilityPipelineCommittersHwmNil() {
	o.CiVisibilityPipelineCommittersHwm.Set(nil)
}

// UnsetCiVisibilityPipelineCommittersHwm ensures that no value is present for CiVisibilityPipelineCommittersHwm, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCiVisibilityPipelineCommittersHwm() {
	o.CiVisibilityPipelineCommittersHwm.Unset()
}

// GetCiVisibilityTestCommittersHwm returns the CiVisibilityTestCommittersHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCiVisibilityTestCommittersHwm() int64 {
	if o == nil || o.CiVisibilityTestCommittersHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityTestCommittersHwm.Get()
}

// GetCiVisibilityTestCommittersHwmOk returns a tuple with the CiVisibilityTestCommittersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCiVisibilityTestCommittersHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiVisibilityTestCommittersHwm.Get(), o.CiVisibilityTestCommittersHwm.IsSet()
}

// HasCiVisibilityTestCommittersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCiVisibilityTestCommittersHwm() bool {
	return o != nil && o.CiVisibilityTestCommittersHwm.IsSet()
}

// SetCiVisibilityTestCommittersHwm gets a reference to the given datadog.NullableInt64 and assigns it to the CiVisibilityTestCommittersHwm field.
func (o *UsageSummaryDateOrg) SetCiVisibilityTestCommittersHwm(v int64) {
	o.CiVisibilityTestCommittersHwm.Set(&v)
}

// SetCiVisibilityTestCommittersHwmNil sets the value for CiVisibilityTestCommittersHwm to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCiVisibilityTestCommittersHwmNil() {
	o.CiVisibilityTestCommittersHwm.Set(nil)
}

// UnsetCiVisibilityTestCommittersHwm ensures that no value is present for CiVisibilityTestCommittersHwm, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCiVisibilityTestCommittersHwm() {
	o.CiVisibilityTestCommittersHwm.Unset()
}

// GetCloudCostManagementHostCountAvg returns the CloudCostManagementHostCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDateOrg) GetCloudCostManagementHostCountAvg() int64 {
	if o == nil || o.CloudCostManagementHostCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementHostCountAvg
}

// GetCloudCostManagementHostCountAvgOk returns a tuple with the CloudCostManagementHostCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDateOrg) GetCloudCostManagementHostCountAvgOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementHostCountAvg == nil {
		return nil, false
	}
	return o.CloudCostManagementHostCountAvg, true
}

// HasCloudCostManagementHostCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCloudCostManagementHostCountAvg() bool {
	return o != nil && o.CloudCostManagementHostCountAvg != nil
}

// SetCloudCostManagementHostCountAvg gets a reference to the given int64 and assigns it to the CloudCostManagementHostCountAvg field.
func (o *UsageSummaryDateOrg) SetCloudCostManagementHostCountAvg(v int64) {
	o.CloudCostManagementHostCountAvg = &v
}

// GetContainerAvg returns the ContainerAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetContainerAvg() int64 {
	if o == nil || o.ContainerAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerAvg.Get()
}

// GetContainerAvgOk returns a tuple with the ContainerAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetContainerAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerAvg.Get(), o.ContainerAvg.IsSet()
}

// HasContainerAvg returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasContainerAvg() bool {
	return o != nil && o.ContainerAvg.IsSet()
}

// SetContainerAvg gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerAvg field.
func (o *UsageSummaryDateOrg) SetContainerAvg(v int64) {
	o.ContainerAvg.Set(&v)
}

// SetContainerAvgNil sets the value for ContainerAvg to be an explicit nil.
func (o *UsageSummaryDateOrg) SetContainerAvgNil() {
	o.ContainerAvg.Set(nil)
}

// UnsetContainerAvg ensures that no value is present for ContainerAvg, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetContainerAvg() {
	o.ContainerAvg.Unset()
}

// GetContainerExclAgentAvg returns the ContainerExclAgentAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetContainerExclAgentAvg() int64 {
	if o == nil || o.ContainerExclAgentAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerExclAgentAvg.Get()
}

// GetContainerExclAgentAvgOk returns a tuple with the ContainerExclAgentAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetContainerExclAgentAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerExclAgentAvg.Get(), o.ContainerExclAgentAvg.IsSet()
}

// HasContainerExclAgentAvg returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasContainerExclAgentAvg() bool {
	return o != nil && o.ContainerExclAgentAvg.IsSet()
}

// SetContainerExclAgentAvg gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerExclAgentAvg field.
func (o *UsageSummaryDateOrg) SetContainerExclAgentAvg(v int64) {
	o.ContainerExclAgentAvg.Set(&v)
}

// SetContainerExclAgentAvgNil sets the value for ContainerExclAgentAvg to be an explicit nil.
func (o *UsageSummaryDateOrg) SetContainerExclAgentAvgNil() {
	o.ContainerExclAgentAvg.Set(nil)
}

// UnsetContainerExclAgentAvg ensures that no value is present for ContainerExclAgentAvg, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetContainerExclAgentAvg() {
	o.ContainerExclAgentAvg.Unset()
}

// GetContainerHwm returns the ContainerHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetContainerHwm() int64 {
	if o == nil || o.ContainerHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerHwm.Get()
}

// GetContainerHwmOk returns a tuple with the ContainerHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetContainerHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerHwm.Get(), o.ContainerHwm.IsSet()
}

// HasContainerHwm returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasContainerHwm() bool {
	return o != nil && o.ContainerHwm.IsSet()
}

// SetContainerHwm gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerHwm field.
func (o *UsageSummaryDateOrg) SetContainerHwm(v int64) {
	o.ContainerHwm.Set(&v)
}

// SetContainerHwmNil sets the value for ContainerHwm to be an explicit nil.
func (o *UsageSummaryDateOrg) SetContainerHwmNil() {
	o.ContainerHwm.Set(nil)
}

// UnsetContainerHwm ensures that no value is present for ContainerHwm, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetContainerHwm() {
	o.ContainerHwm.Unset()
}

// GetCspmAasHostTop99p returns the CspmAasHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCspmAasHostTop99p() int64 {
	if o == nil || o.CspmAasHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmAasHostTop99p.Get()
}

// GetCspmAasHostTop99pOk returns a tuple with the CspmAasHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCspmAasHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmAasHostTop99p.Get(), o.CspmAasHostTop99p.IsSet()
}

// HasCspmAasHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCspmAasHostTop99p() bool {
	return o != nil && o.CspmAasHostTop99p.IsSet()
}

// SetCspmAasHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CspmAasHostTop99p field.
func (o *UsageSummaryDateOrg) SetCspmAasHostTop99p(v int64) {
	o.CspmAasHostTop99p.Set(&v)
}

// SetCspmAasHostTop99pNil sets the value for CspmAasHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCspmAasHostTop99pNil() {
	o.CspmAasHostTop99p.Set(nil)
}

// UnsetCspmAasHostTop99p ensures that no value is present for CspmAasHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCspmAasHostTop99p() {
	o.CspmAasHostTop99p.Unset()
}

// GetCspmAwsHostTop99p returns the CspmAwsHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDateOrg) GetCspmAwsHostTop99p() int64 {
	if o == nil || o.CspmAwsHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CspmAwsHostTop99p
}

// GetCspmAwsHostTop99pOk returns a tuple with the CspmAwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDateOrg) GetCspmAwsHostTop99pOk() (*int64, bool) {
	if o == nil || o.CspmAwsHostTop99p == nil {
		return nil, false
	}
	return o.CspmAwsHostTop99p, true
}

// HasCspmAwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCspmAwsHostTop99p() bool {
	return o != nil && o.CspmAwsHostTop99p != nil
}

// SetCspmAwsHostTop99p gets a reference to the given int64 and assigns it to the CspmAwsHostTop99p field.
func (o *UsageSummaryDateOrg) SetCspmAwsHostTop99p(v int64) {
	o.CspmAwsHostTop99p = &v
}

// GetCspmAzureHostTop99p returns the CspmAzureHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCspmAzureHostTop99p() int64 {
	if o == nil || o.CspmAzureHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmAzureHostTop99p.Get()
}

// GetCspmAzureHostTop99pOk returns a tuple with the CspmAzureHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCspmAzureHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmAzureHostTop99p.Get(), o.CspmAzureHostTop99p.IsSet()
}

// HasCspmAzureHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCspmAzureHostTop99p() bool {
	return o != nil && o.CspmAzureHostTop99p.IsSet()
}

// SetCspmAzureHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CspmAzureHostTop99p field.
func (o *UsageSummaryDateOrg) SetCspmAzureHostTop99p(v int64) {
	o.CspmAzureHostTop99p.Set(&v)
}

// SetCspmAzureHostTop99pNil sets the value for CspmAzureHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCspmAzureHostTop99pNil() {
	o.CspmAzureHostTop99p.Set(nil)
}

// UnsetCspmAzureHostTop99p ensures that no value is present for CspmAzureHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCspmAzureHostTop99p() {
	o.CspmAzureHostTop99p.Unset()
}

// GetCspmContainerAvg returns the CspmContainerAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCspmContainerAvg() int64 {
	if o == nil || o.CspmContainerAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerAvg.Get()
}

// GetCspmContainerAvgOk returns a tuple with the CspmContainerAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCspmContainerAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainerAvg.Get(), o.CspmContainerAvg.IsSet()
}

// HasCspmContainerAvg returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCspmContainerAvg() bool {
	return o != nil && o.CspmContainerAvg.IsSet()
}

// SetCspmContainerAvg gets a reference to the given datadog.NullableInt64 and assigns it to the CspmContainerAvg field.
func (o *UsageSummaryDateOrg) SetCspmContainerAvg(v int64) {
	o.CspmContainerAvg.Set(&v)
}

// SetCspmContainerAvgNil sets the value for CspmContainerAvg to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCspmContainerAvgNil() {
	o.CspmContainerAvg.Set(nil)
}

// UnsetCspmContainerAvg ensures that no value is present for CspmContainerAvg, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCspmContainerAvg() {
	o.CspmContainerAvg.Unset()
}

// GetCspmContainerHwm returns the CspmContainerHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCspmContainerHwm() int64 {
	if o == nil || o.CspmContainerHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerHwm.Get()
}

// GetCspmContainerHwmOk returns a tuple with the CspmContainerHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCspmContainerHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainerHwm.Get(), o.CspmContainerHwm.IsSet()
}

// HasCspmContainerHwm returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCspmContainerHwm() bool {
	return o != nil && o.CspmContainerHwm.IsSet()
}

// SetCspmContainerHwm gets a reference to the given datadog.NullableInt64 and assigns it to the CspmContainerHwm field.
func (o *UsageSummaryDateOrg) SetCspmContainerHwm(v int64) {
	o.CspmContainerHwm.Set(&v)
}

// SetCspmContainerHwmNil sets the value for CspmContainerHwm to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCspmContainerHwmNil() {
	o.CspmContainerHwm.Set(nil)
}

// UnsetCspmContainerHwm ensures that no value is present for CspmContainerHwm, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCspmContainerHwm() {
	o.CspmContainerHwm.Unset()
}

// GetCspmGcpHostTop99p returns the CspmGcpHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCspmGcpHostTop99p() int64 {
	if o == nil || o.CspmGcpHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmGcpHostTop99p.Get()
}

// GetCspmGcpHostTop99pOk returns a tuple with the CspmGcpHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCspmGcpHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmGcpHostTop99p.Get(), o.CspmGcpHostTop99p.IsSet()
}

// HasCspmGcpHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCspmGcpHostTop99p() bool {
	return o != nil && o.CspmGcpHostTop99p.IsSet()
}

// SetCspmGcpHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CspmGcpHostTop99p field.
func (o *UsageSummaryDateOrg) SetCspmGcpHostTop99p(v int64) {
	o.CspmGcpHostTop99p.Set(&v)
}

// SetCspmGcpHostTop99pNil sets the value for CspmGcpHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCspmGcpHostTop99pNil() {
	o.CspmGcpHostTop99p.Set(nil)
}

// UnsetCspmGcpHostTop99p ensures that no value is present for CspmGcpHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCspmGcpHostTop99p() {
	o.CspmGcpHostTop99p.Unset()
}

// GetCspmHostTop99p returns the CspmHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCspmHostTop99p() int64 {
	if o == nil || o.CspmHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmHostTop99p.Get()
}

// GetCspmHostTop99pOk returns a tuple with the CspmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCspmHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmHostTop99p.Get(), o.CspmHostTop99p.IsSet()
}

// HasCspmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCspmHostTop99p() bool {
	return o != nil && o.CspmHostTop99p.IsSet()
}

// SetCspmHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CspmHostTop99p field.
func (o *UsageSummaryDateOrg) SetCspmHostTop99p(v int64) {
	o.CspmHostTop99p.Set(&v)
}

// SetCspmHostTop99pNil sets the value for CspmHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCspmHostTop99pNil() {
	o.CspmHostTop99p.Set(nil)
}

// UnsetCspmHostTop99p ensures that no value is present for CspmHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCspmHostTop99p() {
	o.CspmHostTop99p.Unset()
}

// GetCustomTsAvg returns the CustomTsAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCustomTsAvg() int64 {
	if o == nil || o.CustomTsAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CustomTsAvg.Get()
}

// GetCustomTsAvgOk returns a tuple with the CustomTsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCustomTsAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTsAvg.Get(), o.CustomTsAvg.IsSet()
}

// HasCustomTsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCustomTsAvg() bool {
	return o != nil && o.CustomTsAvg.IsSet()
}

// SetCustomTsAvg gets a reference to the given datadog.NullableInt64 and assigns it to the CustomTsAvg field.
func (o *UsageSummaryDateOrg) SetCustomTsAvg(v int64) {
	o.CustomTsAvg.Set(&v)
}

// SetCustomTsAvgNil sets the value for CustomTsAvg to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCustomTsAvgNil() {
	o.CustomTsAvg.Set(nil)
}

// UnsetCustomTsAvg ensures that no value is present for CustomTsAvg, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCustomTsAvg() {
	o.CustomTsAvg.Unset()
}

// GetCwsContainerCountAvg returns the CwsContainerCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCwsContainerCountAvg() int64 {
	if o == nil || o.CwsContainerCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CwsContainerCountAvg.Get()
}

// GetCwsContainerCountAvgOk returns a tuple with the CwsContainerCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCwsContainerCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsContainerCountAvg.Get(), o.CwsContainerCountAvg.IsSet()
}

// HasCwsContainerCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCwsContainerCountAvg() bool {
	return o != nil && o.CwsContainerCountAvg.IsSet()
}

// SetCwsContainerCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the CwsContainerCountAvg field.
func (o *UsageSummaryDateOrg) SetCwsContainerCountAvg(v int64) {
	o.CwsContainerCountAvg.Set(&v)
}

// SetCwsContainerCountAvgNil sets the value for CwsContainerCountAvg to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCwsContainerCountAvgNil() {
	o.CwsContainerCountAvg.Set(nil)
}

// UnsetCwsContainerCountAvg ensures that no value is present for CwsContainerCountAvg, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCwsContainerCountAvg() {
	o.CwsContainerCountAvg.Unset()
}

// GetCwsHostTop99p returns the CwsHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetCwsHostTop99p() int64 {
	if o == nil || o.CwsHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CwsHostTop99p.Get()
}

// GetCwsHostTop99pOk returns a tuple with the CwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetCwsHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsHostTop99p.Get(), o.CwsHostTop99p.IsSet()
}

// HasCwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasCwsHostTop99p() bool {
	return o != nil && o.CwsHostTop99p.IsSet()
}

// SetCwsHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CwsHostTop99p field.
func (o *UsageSummaryDateOrg) SetCwsHostTop99p(v int64) {
	o.CwsHostTop99p.Set(&v)
}

// SetCwsHostTop99pNil sets the value for CwsHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetCwsHostTop99pNil() {
	o.CwsHostTop99p.Set(nil)
}

// UnsetCwsHostTop99p ensures that no value is present for CwsHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetCwsHostTop99p() {
	o.CwsHostTop99p.Unset()
}

// GetDbmHostTop99pSum returns the DbmHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetDbmHostTop99pSum() int64 {
	if o == nil || o.DbmHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DbmHostTop99pSum.Get()
}

// GetDbmHostTop99pSumOk returns a tuple with the DbmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetDbmHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmHostTop99pSum.Get(), o.DbmHostTop99pSum.IsSet()
}

// HasDbmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasDbmHostTop99pSum() bool {
	return o != nil && o.DbmHostTop99pSum.IsSet()
}

// SetDbmHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the DbmHostTop99pSum field.
func (o *UsageSummaryDateOrg) SetDbmHostTop99pSum(v int64) {
	o.DbmHostTop99pSum.Set(&v)
}

// SetDbmHostTop99pSumNil sets the value for DbmHostTop99pSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetDbmHostTop99pSumNil() {
	o.DbmHostTop99pSum.Set(nil)
}

// UnsetDbmHostTop99pSum ensures that no value is present for DbmHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetDbmHostTop99pSum() {
	o.DbmHostTop99pSum.Unset()
}

// GetDbmQueriesAvgSum returns the DbmQueriesAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetDbmQueriesAvgSum() int64 {
	if o == nil || o.DbmQueriesAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DbmQueriesAvgSum.Get()
}

// GetDbmQueriesAvgSumOk returns a tuple with the DbmQueriesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetDbmQueriesAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmQueriesAvgSum.Get(), o.DbmQueriesAvgSum.IsSet()
}

// HasDbmQueriesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasDbmQueriesAvgSum() bool {
	return o != nil && o.DbmQueriesAvgSum.IsSet()
}

// SetDbmQueriesAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the DbmQueriesAvgSum field.
func (o *UsageSummaryDateOrg) SetDbmQueriesAvgSum(v int64) {
	o.DbmQueriesAvgSum.Set(&v)
}

// SetDbmQueriesAvgSumNil sets the value for DbmQueriesAvgSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetDbmQueriesAvgSumNil() {
	o.DbmQueriesAvgSum.Set(nil)
}

// UnsetDbmQueriesAvgSum ensures that no value is present for DbmQueriesAvgSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetDbmQueriesAvgSum() {
	o.DbmQueriesAvgSum.Unset()
}

// GetFargateTasksCountAvg returns the FargateTasksCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetFargateTasksCountAvg() int64 {
	if o == nil || o.FargateTasksCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountAvg.Get()
}

// GetFargateTasksCountAvgOk returns a tuple with the FargateTasksCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetFargateTasksCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FargateTasksCountAvg.Get(), o.FargateTasksCountAvg.IsSet()
}

// HasFargateTasksCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasFargateTasksCountAvg() bool {
	return o != nil && o.FargateTasksCountAvg.IsSet()
}

// SetFargateTasksCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the FargateTasksCountAvg field.
func (o *UsageSummaryDateOrg) SetFargateTasksCountAvg(v int64) {
	o.FargateTasksCountAvg.Set(&v)
}

// SetFargateTasksCountAvgNil sets the value for FargateTasksCountAvg to be an explicit nil.
func (o *UsageSummaryDateOrg) SetFargateTasksCountAvgNil() {
	o.FargateTasksCountAvg.Set(nil)
}

// UnsetFargateTasksCountAvg ensures that no value is present for FargateTasksCountAvg, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetFargateTasksCountAvg() {
	o.FargateTasksCountAvg.Unset()
}

// GetFargateTasksCountHwm returns the FargateTasksCountHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetFargateTasksCountHwm() int64 {
	if o == nil || o.FargateTasksCountHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountHwm.Get()
}

// GetFargateTasksCountHwmOk returns a tuple with the FargateTasksCountHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetFargateTasksCountHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FargateTasksCountHwm.Get(), o.FargateTasksCountHwm.IsSet()
}

// HasFargateTasksCountHwm returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasFargateTasksCountHwm() bool {
	return o != nil && o.FargateTasksCountHwm.IsSet()
}

// SetFargateTasksCountHwm gets a reference to the given datadog.NullableInt64 and assigns it to the FargateTasksCountHwm field.
func (o *UsageSummaryDateOrg) SetFargateTasksCountHwm(v int64) {
	o.FargateTasksCountHwm.Set(&v)
}

// SetFargateTasksCountHwmNil sets the value for FargateTasksCountHwm to be an explicit nil.
func (o *UsageSummaryDateOrg) SetFargateTasksCountHwmNil() {
	o.FargateTasksCountHwm.Set(nil)
}

// UnsetFargateTasksCountHwm ensures that no value is present for FargateTasksCountHwm, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetFargateTasksCountHwm() {
	o.FargateTasksCountHwm.Unset()
}

// GetForwardingEventsBytesSum returns the ForwardingEventsBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetForwardingEventsBytesSum() int64 {
	if o == nil || o.ForwardingEventsBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ForwardingEventsBytesSum.Get()
}

// GetForwardingEventsBytesSumOk returns a tuple with the ForwardingEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetForwardingEventsBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForwardingEventsBytesSum.Get(), o.ForwardingEventsBytesSum.IsSet()
}

// HasForwardingEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasForwardingEventsBytesSum() bool {
	return o != nil && o.ForwardingEventsBytesSum.IsSet()
}

// SetForwardingEventsBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the ForwardingEventsBytesSum field.
func (o *UsageSummaryDateOrg) SetForwardingEventsBytesSum(v int64) {
	o.ForwardingEventsBytesSum.Set(&v)
}

// SetForwardingEventsBytesSumNil sets the value for ForwardingEventsBytesSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetForwardingEventsBytesSumNil() {
	o.ForwardingEventsBytesSum.Set(nil)
}

// UnsetForwardingEventsBytesSum ensures that no value is present for ForwardingEventsBytesSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetForwardingEventsBytesSum() {
	o.ForwardingEventsBytesSum.Unset()
}

// GetGcpHostTop99p returns the GcpHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetGcpHostTop99p() int64 {
	if o == nil || o.GcpHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.GcpHostTop99p.Get()
}

// GetGcpHostTop99pOk returns a tuple with the GcpHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetGcpHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GcpHostTop99p.Get(), o.GcpHostTop99p.IsSet()
}

// HasGcpHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasGcpHostTop99p() bool {
	return o != nil && o.GcpHostTop99p.IsSet()
}

// SetGcpHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the GcpHostTop99p field.
func (o *UsageSummaryDateOrg) SetGcpHostTop99p(v int64) {
	o.GcpHostTop99p.Set(&v)
}

// SetGcpHostTop99pNil sets the value for GcpHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetGcpHostTop99pNil() {
	o.GcpHostTop99p.Set(nil)
}

// UnsetGcpHostTop99p ensures that no value is present for GcpHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetGcpHostTop99p() {
	o.GcpHostTop99p.Unset()
}

// GetHerokuHostTop99p returns the HerokuHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetHerokuHostTop99p() int64 {
	if o == nil || o.HerokuHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HerokuHostTop99p.Get()
}

// GetHerokuHostTop99pOk returns a tuple with the HerokuHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetHerokuHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HerokuHostTop99p.Get(), o.HerokuHostTop99p.IsSet()
}

// HasHerokuHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasHerokuHostTop99p() bool {
	return o != nil && o.HerokuHostTop99p.IsSet()
}

// SetHerokuHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the HerokuHostTop99p field.
func (o *UsageSummaryDateOrg) SetHerokuHostTop99p(v int64) {
	o.HerokuHostTop99p.Set(&v)
}

// SetHerokuHostTop99pNil sets the value for HerokuHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetHerokuHostTop99pNil() {
	o.HerokuHostTop99p.Set(nil)
}

// UnsetHerokuHostTop99p ensures that no value is present for HerokuHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetHerokuHostTop99p() {
	o.HerokuHostTop99p.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UsageSummaryDateOrg) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDateOrg) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UsageSummaryDateOrg) SetId(v string) {
	o.Id = &v
}

// GetIncidentManagementMonthlyActiveUsersHwm returns the IncidentManagementMonthlyActiveUsersHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetIncidentManagementMonthlyActiveUsersHwm() int64 {
	if o == nil || o.IncidentManagementMonthlyActiveUsersHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IncidentManagementMonthlyActiveUsersHwm.Get()
}

// GetIncidentManagementMonthlyActiveUsersHwmOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetIncidentManagementMonthlyActiveUsersHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentManagementMonthlyActiveUsersHwm.Get(), o.IncidentManagementMonthlyActiveUsersHwm.IsSet()
}

// HasIncidentManagementMonthlyActiveUsersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasIncidentManagementMonthlyActiveUsersHwm() bool {
	return o != nil && o.IncidentManagementMonthlyActiveUsersHwm.IsSet()
}

// SetIncidentManagementMonthlyActiveUsersHwm gets a reference to the given datadog.NullableInt64 and assigns it to the IncidentManagementMonthlyActiveUsersHwm field.
func (o *UsageSummaryDateOrg) SetIncidentManagementMonthlyActiveUsersHwm(v int64) {
	o.IncidentManagementMonthlyActiveUsersHwm.Set(&v)
}

// SetIncidentManagementMonthlyActiveUsersHwmNil sets the value for IncidentManagementMonthlyActiveUsersHwm to be an explicit nil.
func (o *UsageSummaryDateOrg) SetIncidentManagementMonthlyActiveUsersHwmNil() {
	o.IncidentManagementMonthlyActiveUsersHwm.Set(nil)
}

// UnsetIncidentManagementMonthlyActiveUsersHwm ensures that no value is present for IncidentManagementMonthlyActiveUsersHwm, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetIncidentManagementMonthlyActiveUsersHwm() {
	o.IncidentManagementMonthlyActiveUsersHwm.Unset()
}

// GetIndexedEventsCountSum returns the IndexedEventsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetIndexedEventsCountSum() int64 {
	if o == nil || o.IndexedEventsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCountSum.Get()
}

// GetIndexedEventsCountSumOk returns a tuple with the IndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexedEventsCountSum.Get(), o.IndexedEventsCountSum.IsSet()
}

// HasIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasIndexedEventsCountSum() bool {
	return o != nil && o.IndexedEventsCountSum.IsSet()
}

// SetIndexedEventsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the IndexedEventsCountSum field.
func (o *UsageSummaryDateOrg) SetIndexedEventsCountSum(v int64) {
	o.IndexedEventsCountSum.Set(&v)
}

// SetIndexedEventsCountSumNil sets the value for IndexedEventsCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetIndexedEventsCountSumNil() {
	o.IndexedEventsCountSum.Set(nil)
}

// UnsetIndexedEventsCountSum ensures that no value is present for IndexedEventsCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetIndexedEventsCountSum() {
	o.IndexedEventsCountSum.Unset()
}

// GetInfraHostTop99p returns the InfraHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetInfraHostTop99p() int64 {
	if o == nil || o.InfraHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostTop99p.Get()
}

// GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetInfraHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfraHostTop99p.Get(), o.InfraHostTop99p.IsSet()
}

// HasInfraHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasInfraHostTop99p() bool {
	return o != nil && o.InfraHostTop99p.IsSet()
}

// SetInfraHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the InfraHostTop99p field.
func (o *UsageSummaryDateOrg) SetInfraHostTop99p(v int64) {
	o.InfraHostTop99p.Set(&v)
}

// SetInfraHostTop99pNil sets the value for InfraHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetInfraHostTop99pNil() {
	o.InfraHostTop99p.Set(nil)
}

// UnsetInfraHostTop99p ensures that no value is present for InfraHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetInfraHostTop99p() {
	o.InfraHostTop99p.Unset()
}

// GetIngestedEventsBytesSum returns the IngestedEventsBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetIngestedEventsBytesSum() int64 {
	if o == nil || o.IngestedEventsBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytesSum.Get()
}

// GetIngestedEventsBytesSumOk returns a tuple with the IngestedEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetIngestedEventsBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IngestedEventsBytesSum.Get(), o.IngestedEventsBytesSum.IsSet()
}

// HasIngestedEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasIngestedEventsBytesSum() bool {
	return o != nil && o.IngestedEventsBytesSum.IsSet()
}

// SetIngestedEventsBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the IngestedEventsBytesSum field.
func (o *UsageSummaryDateOrg) SetIngestedEventsBytesSum(v int64) {
	o.IngestedEventsBytesSum.Set(&v)
}

// SetIngestedEventsBytesSumNil sets the value for IngestedEventsBytesSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetIngestedEventsBytesSumNil() {
	o.IngestedEventsBytesSum.Set(nil)
}

// UnsetIngestedEventsBytesSum ensures that no value is present for IngestedEventsBytesSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetIngestedEventsBytesSum() {
	o.IngestedEventsBytesSum.Unset()
}

// GetIotDeviceAggSum returns the IotDeviceAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetIotDeviceAggSum() int64 {
	if o == nil || o.IotDeviceAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceAggSum.Get()
}

// GetIotDeviceAggSumOk returns a tuple with the IotDeviceAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetIotDeviceAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IotDeviceAggSum.Get(), o.IotDeviceAggSum.IsSet()
}

// HasIotDeviceAggSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasIotDeviceAggSum() bool {
	return o != nil && o.IotDeviceAggSum.IsSet()
}

// SetIotDeviceAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the IotDeviceAggSum field.
func (o *UsageSummaryDateOrg) SetIotDeviceAggSum(v int64) {
	o.IotDeviceAggSum.Set(&v)
}

// SetIotDeviceAggSumNil sets the value for IotDeviceAggSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetIotDeviceAggSumNil() {
	o.IotDeviceAggSum.Set(nil)
}

// UnsetIotDeviceAggSum ensures that no value is present for IotDeviceAggSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetIotDeviceAggSum() {
	o.IotDeviceAggSum.Unset()
}

// GetIotDeviceTop99pSum returns the IotDeviceTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetIotDeviceTop99pSum() int64 {
	if o == nil || o.IotDeviceTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceTop99pSum.Get()
}

// GetIotDeviceTop99pSumOk returns a tuple with the IotDeviceTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetIotDeviceTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IotDeviceTop99pSum.Get(), o.IotDeviceTop99pSum.IsSet()
}

// HasIotDeviceTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasIotDeviceTop99pSum() bool {
	return o != nil && o.IotDeviceTop99pSum.IsSet()
}

// SetIotDeviceTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the IotDeviceTop99pSum field.
func (o *UsageSummaryDateOrg) SetIotDeviceTop99pSum(v int64) {
	o.IotDeviceTop99pSum.Set(&v)
}

// SetIotDeviceTop99pSumNil sets the value for IotDeviceTop99pSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetIotDeviceTop99pSumNil() {
	o.IotDeviceTop99pSum.Set(nil)
}

// UnsetIotDeviceTop99pSum ensures that no value is present for IotDeviceTop99pSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetIotDeviceTop99pSum() {
	o.IotDeviceTop99pSum.Unset()
}

// GetMobileRumLiteSessionCountSum returns the MobileRumLiteSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetMobileRumLiteSessionCountSum() int64 {
	if o == nil || o.MobileRumLiteSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumLiteSessionCountSum.Get()
}

// GetMobileRumLiteSessionCountSumOk returns a tuple with the MobileRumLiteSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetMobileRumLiteSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumLiteSessionCountSum.Get(), o.MobileRumLiteSessionCountSum.IsSet()
}

// HasMobileRumLiteSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasMobileRumLiteSessionCountSum() bool {
	return o != nil && o.MobileRumLiteSessionCountSum.IsSet()
}

// SetMobileRumLiteSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumLiteSessionCountSum field.
func (o *UsageSummaryDateOrg) SetMobileRumLiteSessionCountSum(v int64) {
	o.MobileRumLiteSessionCountSum.Set(&v)
}

// SetMobileRumLiteSessionCountSumNil sets the value for MobileRumLiteSessionCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetMobileRumLiteSessionCountSumNil() {
	o.MobileRumLiteSessionCountSum.Set(nil)
}

// UnsetMobileRumLiteSessionCountSum ensures that no value is present for MobileRumLiteSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetMobileRumLiteSessionCountSum() {
	o.MobileRumLiteSessionCountSum.Unset()
}

// GetMobileRumSessionCountAndroidSum returns the MobileRumSessionCountAndroidSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountAndroidSum() int64 {
	if o == nil || o.MobileRumSessionCountAndroidSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountAndroidSum.Get()
}

// GetMobileRumSessionCountAndroidSumOk returns a tuple with the MobileRumSessionCountAndroidSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountAndroidSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountAndroidSum.Get(), o.MobileRumSessionCountAndroidSum.IsSet()
}

// HasMobileRumSessionCountAndroidSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasMobileRumSessionCountAndroidSum() bool {
	return o != nil && o.MobileRumSessionCountAndroidSum.IsSet()
}

// SetMobileRumSessionCountAndroidSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountAndroidSum field.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountAndroidSum(v int64) {
	o.MobileRumSessionCountAndroidSum.Set(&v)
}

// SetMobileRumSessionCountAndroidSumNil sets the value for MobileRumSessionCountAndroidSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountAndroidSumNil() {
	o.MobileRumSessionCountAndroidSum.Set(nil)
}

// UnsetMobileRumSessionCountAndroidSum ensures that no value is present for MobileRumSessionCountAndroidSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetMobileRumSessionCountAndroidSum() {
	o.MobileRumSessionCountAndroidSum.Unset()
}

// GetMobileRumSessionCountFlutterSum returns the MobileRumSessionCountFlutterSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountFlutterSum() int64 {
	if o == nil || o.MobileRumSessionCountFlutterSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountFlutterSum.Get()
}

// GetMobileRumSessionCountFlutterSumOk returns a tuple with the MobileRumSessionCountFlutterSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountFlutterSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountFlutterSum.Get(), o.MobileRumSessionCountFlutterSum.IsSet()
}

// HasMobileRumSessionCountFlutterSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasMobileRumSessionCountFlutterSum() bool {
	return o != nil && o.MobileRumSessionCountFlutterSum.IsSet()
}

// SetMobileRumSessionCountFlutterSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountFlutterSum field.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountFlutterSum(v int64) {
	o.MobileRumSessionCountFlutterSum.Set(&v)
}

// SetMobileRumSessionCountFlutterSumNil sets the value for MobileRumSessionCountFlutterSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountFlutterSumNil() {
	o.MobileRumSessionCountFlutterSum.Set(nil)
}

// UnsetMobileRumSessionCountFlutterSum ensures that no value is present for MobileRumSessionCountFlutterSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetMobileRumSessionCountFlutterSum() {
	o.MobileRumSessionCountFlutterSum.Unset()
}

// GetMobileRumSessionCountIosSum returns the MobileRumSessionCountIosSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountIosSum() int64 {
	if o == nil || o.MobileRumSessionCountIosSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountIosSum.Get()
}

// GetMobileRumSessionCountIosSumOk returns a tuple with the MobileRumSessionCountIosSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountIosSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountIosSum.Get(), o.MobileRumSessionCountIosSum.IsSet()
}

// HasMobileRumSessionCountIosSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasMobileRumSessionCountIosSum() bool {
	return o != nil && o.MobileRumSessionCountIosSum.IsSet()
}

// SetMobileRumSessionCountIosSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountIosSum field.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountIosSum(v int64) {
	o.MobileRumSessionCountIosSum.Set(&v)
}

// SetMobileRumSessionCountIosSumNil sets the value for MobileRumSessionCountIosSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountIosSumNil() {
	o.MobileRumSessionCountIosSum.Set(nil)
}

// UnsetMobileRumSessionCountIosSum ensures that no value is present for MobileRumSessionCountIosSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetMobileRumSessionCountIosSum() {
	o.MobileRumSessionCountIosSum.Unset()
}

// GetMobileRumSessionCountReactnativeSum returns the MobileRumSessionCountReactnativeSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountReactnativeSum() int64 {
	if o == nil || o.MobileRumSessionCountReactnativeSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountReactnativeSum.Get()
}

// GetMobileRumSessionCountReactnativeSumOk returns a tuple with the MobileRumSessionCountReactnativeSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountReactnativeSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountReactnativeSum.Get(), o.MobileRumSessionCountReactnativeSum.IsSet()
}

// HasMobileRumSessionCountReactnativeSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasMobileRumSessionCountReactnativeSum() bool {
	return o != nil && o.MobileRumSessionCountReactnativeSum.IsSet()
}

// SetMobileRumSessionCountReactnativeSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountReactnativeSum field.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountReactnativeSum(v int64) {
	o.MobileRumSessionCountReactnativeSum.Set(&v)
}

// SetMobileRumSessionCountReactnativeSumNil sets the value for MobileRumSessionCountReactnativeSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountReactnativeSumNil() {
	o.MobileRumSessionCountReactnativeSum.Set(nil)
}

// UnsetMobileRumSessionCountReactnativeSum ensures that no value is present for MobileRumSessionCountReactnativeSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetMobileRumSessionCountReactnativeSum() {
	o.MobileRumSessionCountReactnativeSum.Unset()
}

// GetMobileRumSessionCountSum returns the MobileRumSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountSum() int64 {
	if o == nil || o.MobileRumSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountSum.Get()
}

// GetMobileRumSessionCountSumOk returns a tuple with the MobileRumSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetMobileRumSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountSum.Get(), o.MobileRumSessionCountSum.IsSet()
}

// HasMobileRumSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasMobileRumSessionCountSum() bool {
	return o != nil && o.MobileRumSessionCountSum.IsSet()
}

// SetMobileRumSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountSum field.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountSum(v int64) {
	o.MobileRumSessionCountSum.Set(&v)
}

// SetMobileRumSessionCountSumNil sets the value for MobileRumSessionCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetMobileRumSessionCountSumNil() {
	o.MobileRumSessionCountSum.Set(nil)
}

// UnsetMobileRumSessionCountSum ensures that no value is present for MobileRumSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetMobileRumSessionCountSum() {
	o.MobileRumSessionCountSum.Unset()
}

// GetMobileRumUnitsSum returns the MobileRumUnitsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetMobileRumUnitsSum() int64 {
	if o == nil || o.MobileRumUnitsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumUnitsSum.Get()
}

// GetMobileRumUnitsSumOk returns a tuple with the MobileRumUnitsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetMobileRumUnitsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumUnitsSum.Get(), o.MobileRumUnitsSum.IsSet()
}

// HasMobileRumUnitsSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasMobileRumUnitsSum() bool {
	return o != nil && o.MobileRumUnitsSum.IsSet()
}

// SetMobileRumUnitsSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumUnitsSum field.
func (o *UsageSummaryDateOrg) SetMobileRumUnitsSum(v int64) {
	o.MobileRumUnitsSum.Set(&v)
}

// SetMobileRumUnitsSumNil sets the value for MobileRumUnitsSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetMobileRumUnitsSumNil() {
	o.MobileRumUnitsSum.Set(nil)
}

// UnsetMobileRumUnitsSum ensures that no value is present for MobileRumUnitsSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetMobileRumUnitsSum() {
	o.MobileRumUnitsSum.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UsageSummaryDateOrg) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDateOrg) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UsageSummaryDateOrg) SetName(v string) {
	o.Name = &v
}

// GetNetflowIndexedEventsCountSum returns the NetflowIndexedEventsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetNetflowIndexedEventsCountSum() int64 {
	if o == nil || o.NetflowIndexedEventsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NetflowIndexedEventsCountSum.Get()
}

// GetNetflowIndexedEventsCountSumOk returns a tuple with the NetflowIndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetNetflowIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetflowIndexedEventsCountSum.Get(), o.NetflowIndexedEventsCountSum.IsSet()
}

// HasNetflowIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasNetflowIndexedEventsCountSum() bool {
	return o != nil && o.NetflowIndexedEventsCountSum.IsSet()
}

// SetNetflowIndexedEventsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the NetflowIndexedEventsCountSum field.
func (o *UsageSummaryDateOrg) SetNetflowIndexedEventsCountSum(v int64) {
	o.NetflowIndexedEventsCountSum.Set(&v)
}

// SetNetflowIndexedEventsCountSumNil sets the value for NetflowIndexedEventsCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetNetflowIndexedEventsCountSumNil() {
	o.NetflowIndexedEventsCountSum.Set(nil)
}

// UnsetNetflowIndexedEventsCountSum ensures that no value is present for NetflowIndexedEventsCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetNetflowIndexedEventsCountSum() {
	o.NetflowIndexedEventsCountSum.Unset()
}

// GetNpmHostTop99p returns the NpmHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetNpmHostTop99p() int64 {
	if o == nil || o.NpmHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NpmHostTop99p.Get()
}

// GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetNpmHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NpmHostTop99p.Get(), o.NpmHostTop99p.IsSet()
}

// HasNpmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasNpmHostTop99p() bool {
	return o != nil && o.NpmHostTop99p.IsSet()
}

// SetNpmHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the NpmHostTop99p field.
func (o *UsageSummaryDateOrg) SetNpmHostTop99p(v int64) {
	o.NpmHostTop99p.Set(&v)
}

// SetNpmHostTop99pNil sets the value for NpmHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetNpmHostTop99pNil() {
	o.NpmHostTop99p.Set(nil)
}

// UnsetNpmHostTop99p ensures that no value is present for NpmHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetNpmHostTop99p() {
	o.NpmHostTop99p.Unset()
}

// GetObservabilityPipelinesBytesProcessedSum returns the ObservabilityPipelinesBytesProcessedSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetObservabilityPipelinesBytesProcessedSum() int64 {
	if o == nil || o.ObservabilityPipelinesBytesProcessedSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ObservabilityPipelinesBytesProcessedSum.Get()
}

// GetObservabilityPipelinesBytesProcessedSumOk returns a tuple with the ObservabilityPipelinesBytesProcessedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetObservabilityPipelinesBytesProcessedSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObservabilityPipelinesBytesProcessedSum.Get(), o.ObservabilityPipelinesBytesProcessedSum.IsSet()
}

// HasObservabilityPipelinesBytesProcessedSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasObservabilityPipelinesBytesProcessedSum() bool {
	return o != nil && o.ObservabilityPipelinesBytesProcessedSum.IsSet()
}

// SetObservabilityPipelinesBytesProcessedSum gets a reference to the given datadog.NullableInt64 and assigns it to the ObservabilityPipelinesBytesProcessedSum field.
func (o *UsageSummaryDateOrg) SetObservabilityPipelinesBytesProcessedSum(v int64) {
	o.ObservabilityPipelinesBytesProcessedSum.Set(&v)
}

// SetObservabilityPipelinesBytesProcessedSumNil sets the value for ObservabilityPipelinesBytesProcessedSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetObservabilityPipelinesBytesProcessedSumNil() {
	o.ObservabilityPipelinesBytesProcessedSum.Set(nil)
}

// UnsetObservabilityPipelinesBytesProcessedSum ensures that no value is present for ObservabilityPipelinesBytesProcessedSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetObservabilityPipelinesBytesProcessedSum() {
	o.ObservabilityPipelinesBytesProcessedSum.Unset()
}

// GetOnlineArchiveEventsCountSum returns the OnlineArchiveEventsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetOnlineArchiveEventsCountSum() int64 {
	if o == nil || o.OnlineArchiveEventsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OnlineArchiveEventsCountSum.Get()
}

// GetOnlineArchiveEventsCountSumOk returns a tuple with the OnlineArchiveEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetOnlineArchiveEventsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlineArchiveEventsCountSum.Get(), o.OnlineArchiveEventsCountSum.IsSet()
}

// HasOnlineArchiveEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasOnlineArchiveEventsCountSum() bool {
	return o != nil && o.OnlineArchiveEventsCountSum.IsSet()
}

// SetOnlineArchiveEventsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the OnlineArchiveEventsCountSum field.
func (o *UsageSummaryDateOrg) SetOnlineArchiveEventsCountSum(v int64) {
	o.OnlineArchiveEventsCountSum.Set(&v)
}

// SetOnlineArchiveEventsCountSumNil sets the value for OnlineArchiveEventsCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetOnlineArchiveEventsCountSumNil() {
	o.OnlineArchiveEventsCountSum.Set(nil)
}

// UnsetOnlineArchiveEventsCountSum ensures that no value is present for OnlineArchiveEventsCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetOnlineArchiveEventsCountSum() {
	o.OnlineArchiveEventsCountSum.Unset()
}

// GetOpentelemetryApmHostTop99p returns the OpentelemetryApmHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetOpentelemetryApmHostTop99p() int64 {
	if o == nil || o.OpentelemetryApmHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryApmHostTop99p.Get()
}

// GetOpentelemetryApmHostTop99pOk returns a tuple with the OpentelemetryApmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetOpentelemetryApmHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpentelemetryApmHostTop99p.Get(), o.OpentelemetryApmHostTop99p.IsSet()
}

// HasOpentelemetryApmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasOpentelemetryApmHostTop99p() bool {
	return o != nil && o.OpentelemetryApmHostTop99p.IsSet()
}

// SetOpentelemetryApmHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the OpentelemetryApmHostTop99p field.
func (o *UsageSummaryDateOrg) SetOpentelemetryApmHostTop99p(v int64) {
	o.OpentelemetryApmHostTop99p.Set(&v)
}

// SetOpentelemetryApmHostTop99pNil sets the value for OpentelemetryApmHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetOpentelemetryApmHostTop99pNil() {
	o.OpentelemetryApmHostTop99p.Set(nil)
}

// UnsetOpentelemetryApmHostTop99p ensures that no value is present for OpentelemetryApmHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetOpentelemetryApmHostTop99p() {
	o.OpentelemetryApmHostTop99p.Unset()
}

// GetOpentelemetryHostTop99p returns the OpentelemetryHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetOpentelemetryHostTop99p() int64 {
	if o == nil || o.OpentelemetryHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryHostTop99p.Get()
}

// GetOpentelemetryHostTop99pOk returns a tuple with the OpentelemetryHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetOpentelemetryHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpentelemetryHostTop99p.Get(), o.OpentelemetryHostTop99p.IsSet()
}

// HasOpentelemetryHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasOpentelemetryHostTop99p() bool {
	return o != nil && o.OpentelemetryHostTop99p.IsSet()
}

// SetOpentelemetryHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the OpentelemetryHostTop99p field.
func (o *UsageSummaryDateOrg) SetOpentelemetryHostTop99p(v int64) {
	o.OpentelemetryHostTop99p.Set(&v)
}

// SetOpentelemetryHostTop99pNil sets the value for OpentelemetryHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetOpentelemetryHostTop99pNil() {
	o.OpentelemetryHostTop99p.Set(nil)
}

// UnsetOpentelemetryHostTop99p ensures that no value is present for OpentelemetryHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetOpentelemetryHostTop99p() {
	o.OpentelemetryHostTop99p.Unset()
}

// GetProfilingHostTop99p returns the ProfilingHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetProfilingHostTop99p() int64 {
	if o == nil || o.ProfilingHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ProfilingHostTop99p.Get()
}

// GetProfilingHostTop99pOk returns a tuple with the ProfilingHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetProfilingHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilingHostTop99p.Get(), o.ProfilingHostTop99p.IsSet()
}

// HasProfilingHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasProfilingHostTop99p() bool {
	return o != nil && o.ProfilingHostTop99p.IsSet()
}

// SetProfilingHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the ProfilingHostTop99p field.
func (o *UsageSummaryDateOrg) SetProfilingHostTop99p(v int64) {
	o.ProfilingHostTop99p.Set(&v)
}

// SetProfilingHostTop99pNil sets the value for ProfilingHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetProfilingHostTop99pNil() {
	o.ProfilingHostTop99p.Set(nil)
}

// UnsetProfilingHostTop99p ensures that no value is present for ProfilingHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetProfilingHostTop99p() {
	o.ProfilingHostTop99p.Unset()
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageSummaryDateOrg) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDateOrg) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageSummaryDateOrg) SetPublicId(v string) {
	o.PublicId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *UsageSummaryDateOrg) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDateOrg) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *UsageSummaryDateOrg) SetRegion(v string) {
	o.Region = &v
}

// GetRumBrowserAndMobileSessionCount returns the RumBrowserAndMobileSessionCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetRumBrowserAndMobileSessionCount() int64 {
	if o == nil || o.RumBrowserAndMobileSessionCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserAndMobileSessionCount.Get()
}

// GetRumBrowserAndMobileSessionCountOk returns a tuple with the RumBrowserAndMobileSessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetRumBrowserAndMobileSessionCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumBrowserAndMobileSessionCount.Get(), o.RumBrowserAndMobileSessionCount.IsSet()
}

// HasRumBrowserAndMobileSessionCount returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasRumBrowserAndMobileSessionCount() bool {
	return o != nil && o.RumBrowserAndMobileSessionCount.IsSet()
}

// SetRumBrowserAndMobileSessionCount gets a reference to the given datadog.NullableInt64 and assigns it to the RumBrowserAndMobileSessionCount field.
func (o *UsageSummaryDateOrg) SetRumBrowserAndMobileSessionCount(v int64) {
	o.RumBrowserAndMobileSessionCount.Set(&v)
}

// SetRumBrowserAndMobileSessionCountNil sets the value for RumBrowserAndMobileSessionCount to be an explicit nil.
func (o *UsageSummaryDateOrg) SetRumBrowserAndMobileSessionCountNil() {
	o.RumBrowserAndMobileSessionCount.Set(nil)
}

// UnsetRumBrowserAndMobileSessionCount ensures that no value is present for RumBrowserAndMobileSessionCount, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetRumBrowserAndMobileSessionCount() {
	o.RumBrowserAndMobileSessionCount.Unset()
}

// GetRumSessionCountSum returns the RumSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetRumSessionCountSum() int64 {
	if o == nil || o.RumSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumSessionCountSum.Get()
}

// GetRumSessionCountSumOk returns a tuple with the RumSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetRumSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumSessionCountSum.Get(), o.RumSessionCountSum.IsSet()
}

// HasRumSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasRumSessionCountSum() bool {
	return o != nil && o.RumSessionCountSum.IsSet()
}

// SetRumSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the RumSessionCountSum field.
func (o *UsageSummaryDateOrg) SetRumSessionCountSum(v int64) {
	o.RumSessionCountSum.Set(&v)
}

// SetRumSessionCountSumNil sets the value for RumSessionCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetRumSessionCountSumNil() {
	o.RumSessionCountSum.Set(nil)
}

// UnsetRumSessionCountSum ensures that no value is present for RumSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetRumSessionCountSum() {
	o.RumSessionCountSum.Unset()
}

// GetRumTotalSessionCountSum returns the RumTotalSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetRumTotalSessionCountSum() int64 {
	if o == nil || o.RumTotalSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumTotalSessionCountSum.Get()
}

// GetRumTotalSessionCountSumOk returns a tuple with the RumTotalSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetRumTotalSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumTotalSessionCountSum.Get(), o.RumTotalSessionCountSum.IsSet()
}

// HasRumTotalSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasRumTotalSessionCountSum() bool {
	return o != nil && o.RumTotalSessionCountSum.IsSet()
}

// SetRumTotalSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the RumTotalSessionCountSum field.
func (o *UsageSummaryDateOrg) SetRumTotalSessionCountSum(v int64) {
	o.RumTotalSessionCountSum.Set(&v)
}

// SetRumTotalSessionCountSumNil sets the value for RumTotalSessionCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetRumTotalSessionCountSumNil() {
	o.RumTotalSessionCountSum.Set(nil)
}

// UnsetRumTotalSessionCountSum ensures that no value is present for RumTotalSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetRumTotalSessionCountSum() {
	o.RumTotalSessionCountSum.Unset()
}

// GetRumUnitsSum returns the RumUnitsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetRumUnitsSum() int64 {
	if o == nil || o.RumUnitsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumUnitsSum.Get()
}

// GetRumUnitsSumOk returns a tuple with the RumUnitsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetRumUnitsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumUnitsSum.Get(), o.RumUnitsSum.IsSet()
}

// HasRumUnitsSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasRumUnitsSum() bool {
	return o != nil && o.RumUnitsSum.IsSet()
}

// SetRumUnitsSum gets a reference to the given datadog.NullableInt64 and assigns it to the RumUnitsSum field.
func (o *UsageSummaryDateOrg) SetRumUnitsSum(v int64) {
	o.RumUnitsSum.Set(&v)
}

// SetRumUnitsSumNil sets the value for RumUnitsSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetRumUnitsSumNil() {
	o.RumUnitsSum.Set(nil)
}

// UnsetRumUnitsSum ensures that no value is present for RumUnitsSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetRumUnitsSum() {
	o.RumUnitsSum.Unset()
}

// GetSdsApmScannedBytesSum returns the SdsApmScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetSdsApmScannedBytesSum() int64 {
	if o == nil || o.SdsApmScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsApmScannedBytesSum.Get()
}

// GetSdsApmScannedBytesSumOk returns a tuple with the SdsApmScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetSdsApmScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsApmScannedBytesSum.Get(), o.SdsApmScannedBytesSum.IsSet()
}

// HasSdsApmScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasSdsApmScannedBytesSum() bool {
	return o != nil && o.SdsApmScannedBytesSum.IsSet()
}

// SetSdsApmScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsApmScannedBytesSum field.
func (o *UsageSummaryDateOrg) SetSdsApmScannedBytesSum(v int64) {
	o.SdsApmScannedBytesSum.Set(&v)
}

// SetSdsApmScannedBytesSumNil sets the value for SdsApmScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetSdsApmScannedBytesSumNil() {
	o.SdsApmScannedBytesSum.Set(nil)
}

// UnsetSdsApmScannedBytesSum ensures that no value is present for SdsApmScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetSdsApmScannedBytesSum() {
	o.SdsApmScannedBytesSum.Unset()
}

// GetSdsEventsScannedBytesSum returns the SdsEventsScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetSdsEventsScannedBytesSum() int64 {
	if o == nil || o.SdsEventsScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsEventsScannedBytesSum.Get()
}

// GetSdsEventsScannedBytesSumOk returns a tuple with the SdsEventsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetSdsEventsScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsEventsScannedBytesSum.Get(), o.SdsEventsScannedBytesSum.IsSet()
}

// HasSdsEventsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasSdsEventsScannedBytesSum() bool {
	return o != nil && o.SdsEventsScannedBytesSum.IsSet()
}

// SetSdsEventsScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsEventsScannedBytesSum field.
func (o *UsageSummaryDateOrg) SetSdsEventsScannedBytesSum(v int64) {
	o.SdsEventsScannedBytesSum.Set(&v)
}

// SetSdsEventsScannedBytesSumNil sets the value for SdsEventsScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetSdsEventsScannedBytesSumNil() {
	o.SdsEventsScannedBytesSum.Set(nil)
}

// UnsetSdsEventsScannedBytesSum ensures that no value is present for SdsEventsScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetSdsEventsScannedBytesSum() {
	o.SdsEventsScannedBytesSum.Unset()
}

// GetSdsLogsScannedBytesSum returns the SdsLogsScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetSdsLogsScannedBytesSum() int64 {
	if o == nil || o.SdsLogsScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsLogsScannedBytesSum.Get()
}

// GetSdsLogsScannedBytesSumOk returns a tuple with the SdsLogsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetSdsLogsScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsLogsScannedBytesSum.Get(), o.SdsLogsScannedBytesSum.IsSet()
}

// HasSdsLogsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasSdsLogsScannedBytesSum() bool {
	return o != nil && o.SdsLogsScannedBytesSum.IsSet()
}

// SetSdsLogsScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsLogsScannedBytesSum field.
func (o *UsageSummaryDateOrg) SetSdsLogsScannedBytesSum(v int64) {
	o.SdsLogsScannedBytesSum.Set(&v)
}

// SetSdsLogsScannedBytesSumNil sets the value for SdsLogsScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetSdsLogsScannedBytesSumNil() {
	o.SdsLogsScannedBytesSum.Set(nil)
}

// UnsetSdsLogsScannedBytesSum ensures that no value is present for SdsLogsScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetSdsLogsScannedBytesSum() {
	o.SdsLogsScannedBytesSum.Unset()
}

// GetSdsRumScannedBytesSum returns the SdsRumScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetSdsRumScannedBytesSum() int64 {
	if o == nil || o.SdsRumScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsRumScannedBytesSum.Get()
}

// GetSdsRumScannedBytesSumOk returns a tuple with the SdsRumScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetSdsRumScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsRumScannedBytesSum.Get(), o.SdsRumScannedBytesSum.IsSet()
}

// HasSdsRumScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasSdsRumScannedBytesSum() bool {
	return o != nil && o.SdsRumScannedBytesSum.IsSet()
}

// SetSdsRumScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsRumScannedBytesSum field.
func (o *UsageSummaryDateOrg) SetSdsRumScannedBytesSum(v int64) {
	o.SdsRumScannedBytesSum.Set(&v)
}

// SetSdsRumScannedBytesSumNil sets the value for SdsRumScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetSdsRumScannedBytesSumNil() {
	o.SdsRumScannedBytesSum.Set(nil)
}

// UnsetSdsRumScannedBytesSum ensures that no value is present for SdsRumScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetSdsRumScannedBytesSum() {
	o.SdsRumScannedBytesSum.Unset()
}

// GetSdsTotalScannedBytesSum returns the SdsTotalScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetSdsTotalScannedBytesSum() int64 {
	if o == nil || o.SdsTotalScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsTotalScannedBytesSum.Get()
}

// GetSdsTotalScannedBytesSumOk returns a tuple with the SdsTotalScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetSdsTotalScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsTotalScannedBytesSum.Get(), o.SdsTotalScannedBytesSum.IsSet()
}

// HasSdsTotalScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasSdsTotalScannedBytesSum() bool {
	return o != nil && o.SdsTotalScannedBytesSum.IsSet()
}

// SetSdsTotalScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsTotalScannedBytesSum field.
func (o *UsageSummaryDateOrg) SetSdsTotalScannedBytesSum(v int64) {
	o.SdsTotalScannedBytesSum.Set(&v)
}

// SetSdsTotalScannedBytesSumNil sets the value for SdsTotalScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetSdsTotalScannedBytesSumNil() {
	o.SdsTotalScannedBytesSum.Set(nil)
}

// UnsetSdsTotalScannedBytesSum ensures that no value is present for SdsTotalScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetSdsTotalScannedBytesSum() {
	o.SdsTotalScannedBytesSum.Unset()
}

// GetSyntheticsBrowserCheckCallsCountSum returns the SyntheticsBrowserCheckCallsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetSyntheticsBrowserCheckCallsCountSum() int64 {
	if o == nil || o.SyntheticsBrowserCheckCallsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsBrowserCheckCallsCountSum.Get()
}

// GetSyntheticsBrowserCheckCallsCountSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetSyntheticsBrowserCheckCallsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyntheticsBrowserCheckCallsCountSum.Get(), o.SyntheticsBrowserCheckCallsCountSum.IsSet()
}

// HasSyntheticsBrowserCheckCallsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasSyntheticsBrowserCheckCallsCountSum() bool {
	return o != nil && o.SyntheticsBrowserCheckCallsCountSum.IsSet()
}

// SetSyntheticsBrowserCheckCallsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the SyntheticsBrowserCheckCallsCountSum field.
func (o *UsageSummaryDateOrg) SetSyntheticsBrowserCheckCallsCountSum(v int64) {
	o.SyntheticsBrowserCheckCallsCountSum.Set(&v)
}

// SetSyntheticsBrowserCheckCallsCountSumNil sets the value for SyntheticsBrowserCheckCallsCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetSyntheticsBrowserCheckCallsCountSumNil() {
	o.SyntheticsBrowserCheckCallsCountSum.Set(nil)
}

// UnsetSyntheticsBrowserCheckCallsCountSum ensures that no value is present for SyntheticsBrowserCheckCallsCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetSyntheticsBrowserCheckCallsCountSum() {
	o.SyntheticsBrowserCheckCallsCountSum.Unset()
}

// GetSyntheticsCheckCallsCountSum returns the SyntheticsCheckCallsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetSyntheticsCheckCallsCountSum() int64 {
	if o == nil || o.SyntheticsCheckCallsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsCheckCallsCountSum.Get()
}

// GetSyntheticsCheckCallsCountSumOk returns a tuple with the SyntheticsCheckCallsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetSyntheticsCheckCallsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyntheticsCheckCallsCountSum.Get(), o.SyntheticsCheckCallsCountSum.IsSet()
}

// HasSyntheticsCheckCallsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasSyntheticsCheckCallsCountSum() bool {
	return o != nil && o.SyntheticsCheckCallsCountSum.IsSet()
}

// SetSyntheticsCheckCallsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the SyntheticsCheckCallsCountSum field.
func (o *UsageSummaryDateOrg) SetSyntheticsCheckCallsCountSum(v int64) {
	o.SyntheticsCheckCallsCountSum.Set(&v)
}

// SetSyntheticsCheckCallsCountSumNil sets the value for SyntheticsCheckCallsCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetSyntheticsCheckCallsCountSumNil() {
	o.SyntheticsCheckCallsCountSum.Set(nil)
}

// UnsetSyntheticsCheckCallsCountSum ensures that no value is present for SyntheticsCheckCallsCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetSyntheticsCheckCallsCountSum() {
	o.SyntheticsCheckCallsCountSum.Unset()
}

// GetSyntheticsParallelTestingMaxSlotsHwm returns the SyntheticsParallelTestingMaxSlotsHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetSyntheticsParallelTestingMaxSlotsHwm() int64 {
	if o == nil || o.SyntheticsParallelTestingMaxSlotsHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsParallelTestingMaxSlotsHwm.Get()
}

// GetSyntheticsParallelTestingMaxSlotsHwmOk returns a tuple with the SyntheticsParallelTestingMaxSlotsHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetSyntheticsParallelTestingMaxSlotsHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyntheticsParallelTestingMaxSlotsHwm.Get(), o.SyntheticsParallelTestingMaxSlotsHwm.IsSet()
}

// HasSyntheticsParallelTestingMaxSlotsHwm returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasSyntheticsParallelTestingMaxSlotsHwm() bool {
	return o != nil && o.SyntheticsParallelTestingMaxSlotsHwm.IsSet()
}

// SetSyntheticsParallelTestingMaxSlotsHwm gets a reference to the given datadog.NullableInt64 and assigns it to the SyntheticsParallelTestingMaxSlotsHwm field.
func (o *UsageSummaryDateOrg) SetSyntheticsParallelTestingMaxSlotsHwm(v int64) {
	o.SyntheticsParallelTestingMaxSlotsHwm.Set(&v)
}

// SetSyntheticsParallelTestingMaxSlotsHwmNil sets the value for SyntheticsParallelTestingMaxSlotsHwm to be an explicit nil.
func (o *UsageSummaryDateOrg) SetSyntheticsParallelTestingMaxSlotsHwmNil() {
	o.SyntheticsParallelTestingMaxSlotsHwm.Set(nil)
}

// UnsetSyntheticsParallelTestingMaxSlotsHwm ensures that no value is present for SyntheticsParallelTestingMaxSlotsHwm, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetSyntheticsParallelTestingMaxSlotsHwm() {
	o.SyntheticsParallelTestingMaxSlotsHwm.Unset()
}

// GetTraceSearchIndexedEventsCountSum returns the TraceSearchIndexedEventsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetTraceSearchIndexedEventsCountSum() int64 {
	if o == nil || o.TraceSearchIndexedEventsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TraceSearchIndexedEventsCountSum.Get()
}

// GetTraceSearchIndexedEventsCountSumOk returns a tuple with the TraceSearchIndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetTraceSearchIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceSearchIndexedEventsCountSum.Get(), o.TraceSearchIndexedEventsCountSum.IsSet()
}

// HasTraceSearchIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasTraceSearchIndexedEventsCountSum() bool {
	return o != nil && o.TraceSearchIndexedEventsCountSum.IsSet()
}

// SetTraceSearchIndexedEventsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the TraceSearchIndexedEventsCountSum field.
func (o *UsageSummaryDateOrg) SetTraceSearchIndexedEventsCountSum(v int64) {
	o.TraceSearchIndexedEventsCountSum.Set(&v)
}

// SetTraceSearchIndexedEventsCountSumNil sets the value for TraceSearchIndexedEventsCountSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetTraceSearchIndexedEventsCountSumNil() {
	o.TraceSearchIndexedEventsCountSum.Set(nil)
}

// UnsetTraceSearchIndexedEventsCountSum ensures that no value is present for TraceSearchIndexedEventsCountSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetTraceSearchIndexedEventsCountSum() {
	o.TraceSearchIndexedEventsCountSum.Unset()
}

// GetTwolIngestedEventsBytesSum returns the TwolIngestedEventsBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetTwolIngestedEventsBytesSum() int64 {
	if o == nil || o.TwolIngestedEventsBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TwolIngestedEventsBytesSum.Get()
}

// GetTwolIngestedEventsBytesSumOk returns a tuple with the TwolIngestedEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetTwolIngestedEventsBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwolIngestedEventsBytesSum.Get(), o.TwolIngestedEventsBytesSum.IsSet()
}

// HasTwolIngestedEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasTwolIngestedEventsBytesSum() bool {
	return o != nil && o.TwolIngestedEventsBytesSum.IsSet()
}

// SetTwolIngestedEventsBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the TwolIngestedEventsBytesSum field.
func (o *UsageSummaryDateOrg) SetTwolIngestedEventsBytesSum(v int64) {
	o.TwolIngestedEventsBytesSum.Set(&v)
}

// SetTwolIngestedEventsBytesSumNil sets the value for TwolIngestedEventsBytesSum to be an explicit nil.
func (o *UsageSummaryDateOrg) SetTwolIngestedEventsBytesSumNil() {
	o.TwolIngestedEventsBytesSum.Set(nil)
}

// UnsetTwolIngestedEventsBytesSum ensures that no value is present for TwolIngestedEventsBytesSum, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetTwolIngestedEventsBytesSum() {
	o.TwolIngestedEventsBytesSum.Unset()
}

// GetUniversalServiceMonitoringHostTop99p returns the UniversalServiceMonitoringHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetUniversalServiceMonitoringHostTop99p() int64 {
	if o == nil || o.UniversalServiceMonitoringHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UniversalServiceMonitoringHostTop99p.Get()
}

// GetUniversalServiceMonitoringHostTop99pOk returns a tuple with the UniversalServiceMonitoringHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetUniversalServiceMonitoringHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringHostTop99p.Get(), o.UniversalServiceMonitoringHostTop99p.IsSet()
}

// HasUniversalServiceMonitoringHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasUniversalServiceMonitoringHostTop99p() bool {
	return o != nil && o.UniversalServiceMonitoringHostTop99p.IsSet()
}

// SetUniversalServiceMonitoringHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the UniversalServiceMonitoringHostTop99p field.
func (o *UsageSummaryDateOrg) SetUniversalServiceMonitoringHostTop99p(v int64) {
	o.UniversalServiceMonitoringHostTop99p.Set(&v)
}

// SetUniversalServiceMonitoringHostTop99pNil sets the value for UniversalServiceMonitoringHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetUniversalServiceMonitoringHostTop99pNil() {
	o.UniversalServiceMonitoringHostTop99p.Set(nil)
}

// UnsetUniversalServiceMonitoringHostTop99p ensures that no value is present for UniversalServiceMonitoringHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetUniversalServiceMonitoringHostTop99p() {
	o.UniversalServiceMonitoringHostTop99p.Unset()
}

// GetVsphereHostTop99p returns the VsphereHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDateOrg) GetVsphereHostTop99p() int64 {
	if o == nil || o.VsphereHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.VsphereHostTop99p.Get()
}

// GetVsphereHostTop99pOk returns a tuple with the VsphereHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDateOrg) GetVsphereHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VsphereHostTop99p.Get(), o.VsphereHostTop99p.IsSet()
}

// HasVsphereHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDateOrg) HasVsphereHostTop99p() bool {
	return o != nil && o.VsphereHostTop99p.IsSet()
}

// SetVsphereHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the VsphereHostTop99p field.
func (o *UsageSummaryDateOrg) SetVsphereHostTop99p(v int64) {
	o.VsphereHostTop99p.Set(&v)
}

// SetVsphereHostTop99pNil sets the value for VsphereHostTop99p to be an explicit nil.
func (o *UsageSummaryDateOrg) SetVsphereHostTop99pNil() {
	o.VsphereHostTop99p.Set(nil)
}

// UnsetVsphereHostTop99p ensures that no value is present for VsphereHostTop99p, not even an explicit nil.
func (o *UsageSummaryDateOrg) UnsetVsphereHostTop99p() {
	o.VsphereHostTop99p.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSummaryDateOrg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AgentHostTop99p.IsSet() {
		toSerialize["agent_host_top99p"] = o.AgentHostTop99p.Get()
	}
	if o.ApmAzureAppServiceHostTop99p.IsSet() {
		toSerialize["apm_azure_app_service_host_top99p"] = o.ApmAzureAppServiceHostTop99p.Get()
	}
	if o.ApmFargateCountAvg.IsSet() {
		toSerialize["apm_fargate_count_avg"] = o.ApmFargateCountAvg.Get()
	}
	if o.ApmHostTop99p.IsSet() {
		toSerialize["apm_host_top99p"] = o.ApmHostTop99p.Get()
	}
	if o.AppsecFargateCountAvg.IsSet() {
		toSerialize["appsec_fargate_count_avg"] = o.AppsecFargateCountAvg.Get()
	}
	if o.AuditLogsLinesIndexedSum.IsSet() {
		toSerialize["audit_logs_lines_indexed_sum"] = o.AuditLogsLinesIndexedSum.Get()
	}
	if o.AuditTrailEnabledHwm.IsSet() {
		toSerialize["audit_trail_enabled_hwm"] = o.AuditTrailEnabledHwm.Get()
	}
	if o.AvgProfiledFargateTasks.IsSet() {
		toSerialize["avg_profiled_fargate_tasks"] = o.AvgProfiledFargateTasks.Get()
	}
	if o.AwsHostTop99p.IsSet() {
		toSerialize["aws_host_top99p"] = o.AwsHostTop99p.Get()
	}
	if o.AwsLambdaFuncCount != nil {
		toSerialize["aws_lambda_func_count"] = o.AwsLambdaFuncCount
	}
	if o.AwsLambdaInvocationsSum.IsSet() {
		toSerialize["aws_lambda_invocations_sum"] = o.AwsLambdaInvocationsSum.Get()
	}
	if o.AzureAppServiceTop99p.IsSet() {
		toSerialize["azure_app_service_top99p"] = o.AzureAppServiceTop99p.Get()
	}
	if o.BillableIngestedBytesSum.IsSet() {
		toSerialize["billable_ingested_bytes_sum"] = o.BillableIngestedBytesSum.Get()
	}
	if o.BrowserRumLiteSessionCountSum.IsSet() {
		toSerialize["browser_rum_lite_session_count_sum"] = o.BrowserRumLiteSessionCountSum.Get()
	}
	if o.BrowserRumReplaySessionCountSum.IsSet() {
		toSerialize["browser_rum_replay_session_count_sum"] = o.BrowserRumReplaySessionCountSum.Get()
	}
	if o.BrowserRumUnitsSum.IsSet() {
		toSerialize["browser_rum_units_sum"] = o.BrowserRumUnitsSum.Get()
	}
	if o.CiPipelineIndexedSpansSum.IsSet() {
		toSerialize["ci_pipeline_indexed_spans_sum"] = o.CiPipelineIndexedSpansSum.Get()
	}
	if o.CiTestIndexedSpansSum.IsSet() {
		toSerialize["ci_test_indexed_spans_sum"] = o.CiTestIndexedSpansSum.Get()
	}
	if o.CiVisibilityPipelineCommittersHwm.IsSet() {
		toSerialize["ci_visibility_pipeline_committers_hwm"] = o.CiVisibilityPipelineCommittersHwm.Get()
	}
	if o.CiVisibilityTestCommittersHwm.IsSet() {
		toSerialize["ci_visibility_test_committers_hwm"] = o.CiVisibilityTestCommittersHwm.Get()
	}
	if o.CloudCostManagementHostCountAvg != nil {
		toSerialize["cloud_cost_management_host_count_avg"] = o.CloudCostManagementHostCountAvg
	}
	if o.ContainerAvg.IsSet() {
		toSerialize["container_avg"] = o.ContainerAvg.Get()
	}
	if o.ContainerExclAgentAvg.IsSet() {
		toSerialize["container_excl_agent_avg"] = o.ContainerExclAgentAvg.Get()
	}
	if o.ContainerHwm.IsSet() {
		toSerialize["container_hwm"] = o.ContainerHwm.Get()
	}
	if o.CspmAasHostTop99p.IsSet() {
		toSerialize["cspm_aas_host_top99p"] = o.CspmAasHostTop99p.Get()
	}
	if o.CspmAwsHostTop99p != nil {
		toSerialize["cspm_aws_host_top99p"] = o.CspmAwsHostTop99p
	}
	if o.CspmAzureHostTop99p.IsSet() {
		toSerialize["cspm_azure_host_top99p"] = o.CspmAzureHostTop99p.Get()
	}
	if o.CspmContainerAvg.IsSet() {
		toSerialize["cspm_container_avg"] = o.CspmContainerAvg.Get()
	}
	if o.CspmContainerHwm.IsSet() {
		toSerialize["cspm_container_hwm"] = o.CspmContainerHwm.Get()
	}
	if o.CspmGcpHostTop99p.IsSet() {
		toSerialize["cspm_gcp_host_top99p"] = o.CspmGcpHostTop99p.Get()
	}
	if o.CspmHostTop99p.IsSet() {
		toSerialize["cspm_host_top99p"] = o.CspmHostTop99p.Get()
	}
	if o.CustomTsAvg.IsSet() {
		toSerialize["custom_ts_avg"] = o.CustomTsAvg.Get()
	}
	if o.CwsContainerCountAvg.IsSet() {
		toSerialize["cws_container_count_avg"] = o.CwsContainerCountAvg.Get()
	}
	if o.CwsHostTop99p.IsSet() {
		toSerialize["cws_host_top99p"] = o.CwsHostTop99p.Get()
	}
	if o.DbmHostTop99pSum.IsSet() {
		toSerialize["dbm_host_top99p_sum"] = o.DbmHostTop99pSum.Get()
	}
	if o.DbmQueriesAvgSum.IsSet() {
		toSerialize["dbm_queries_avg_sum"] = o.DbmQueriesAvgSum.Get()
	}
	if o.FargateTasksCountAvg.IsSet() {
		toSerialize["fargate_tasks_count_avg"] = o.FargateTasksCountAvg.Get()
	}
	if o.FargateTasksCountHwm.IsSet() {
		toSerialize["fargate_tasks_count_hwm"] = o.FargateTasksCountHwm.Get()
	}
	if o.ForwardingEventsBytesSum.IsSet() {
		toSerialize["forwarding_events_bytes_sum"] = o.ForwardingEventsBytesSum.Get()
	}
	if o.GcpHostTop99p.IsSet() {
		toSerialize["gcp_host_top99p"] = o.GcpHostTop99p.Get()
	}
	if o.HerokuHostTop99p.IsSet() {
		toSerialize["heroku_host_top99p"] = o.HerokuHostTop99p.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IncidentManagementMonthlyActiveUsersHwm.IsSet() {
		toSerialize["incident_management_monthly_active_users_hwm"] = o.IncidentManagementMonthlyActiveUsersHwm.Get()
	}
	if o.IndexedEventsCountSum.IsSet() {
		toSerialize["indexed_events_count_sum"] = o.IndexedEventsCountSum.Get()
	}
	if o.InfraHostTop99p.IsSet() {
		toSerialize["infra_host_top99p"] = o.InfraHostTop99p.Get()
	}
	if o.IngestedEventsBytesSum.IsSet() {
		toSerialize["ingested_events_bytes_sum"] = o.IngestedEventsBytesSum.Get()
	}
	if o.IotDeviceAggSum.IsSet() {
		toSerialize["iot_device_agg_sum"] = o.IotDeviceAggSum.Get()
	}
	if o.IotDeviceTop99pSum.IsSet() {
		toSerialize["iot_device_top99p_sum"] = o.IotDeviceTop99pSum.Get()
	}
	if o.MobileRumLiteSessionCountSum.IsSet() {
		toSerialize["mobile_rum_lite_session_count_sum"] = o.MobileRumLiteSessionCountSum.Get()
	}
	if o.MobileRumSessionCountAndroidSum.IsSet() {
		toSerialize["mobile_rum_session_count_android_sum"] = o.MobileRumSessionCountAndroidSum.Get()
	}
	if o.MobileRumSessionCountFlutterSum.IsSet() {
		toSerialize["mobile_rum_session_count_flutter_sum"] = o.MobileRumSessionCountFlutterSum.Get()
	}
	if o.MobileRumSessionCountIosSum.IsSet() {
		toSerialize["mobile_rum_session_count_ios_sum"] = o.MobileRumSessionCountIosSum.Get()
	}
	if o.MobileRumSessionCountReactnativeSum.IsSet() {
		toSerialize["mobile_rum_session_count_reactnative_sum"] = o.MobileRumSessionCountReactnativeSum.Get()
	}
	if o.MobileRumSessionCountSum.IsSet() {
		toSerialize["mobile_rum_session_count_sum"] = o.MobileRumSessionCountSum.Get()
	}
	if o.MobileRumUnitsSum.IsSet() {
		toSerialize["mobile_rum_units_sum"] = o.MobileRumUnitsSum.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NetflowIndexedEventsCountSum.IsSet() {
		toSerialize["netflow_indexed_events_count_sum"] = o.NetflowIndexedEventsCountSum.Get()
	}
	if o.NpmHostTop99p.IsSet() {
		toSerialize["npm_host_top99p"] = o.NpmHostTop99p.Get()
	}
	if o.ObservabilityPipelinesBytesProcessedSum.IsSet() {
		toSerialize["observability_pipelines_bytes_processed_sum"] = o.ObservabilityPipelinesBytesProcessedSum.Get()
	}
	if o.OnlineArchiveEventsCountSum.IsSet() {
		toSerialize["online_archive_events_count_sum"] = o.OnlineArchiveEventsCountSum.Get()
	}
	if o.OpentelemetryApmHostTop99p.IsSet() {
		toSerialize["opentelemetry_apm_host_top99p"] = o.OpentelemetryApmHostTop99p.Get()
	}
	if o.OpentelemetryHostTop99p.IsSet() {
		toSerialize["opentelemetry_host_top99p"] = o.OpentelemetryHostTop99p.Get()
	}
	if o.ProfilingHostTop99p.IsSet() {
		toSerialize["profiling_host_top99p"] = o.ProfilingHostTop99p.Get()
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.RumBrowserAndMobileSessionCount.IsSet() {
		toSerialize["rum_browser_and_mobile_session_count"] = o.RumBrowserAndMobileSessionCount.Get()
	}
	if o.RumSessionCountSum.IsSet() {
		toSerialize["rum_session_count_sum"] = o.RumSessionCountSum.Get()
	}
	if o.RumTotalSessionCountSum.IsSet() {
		toSerialize["rum_total_session_count_sum"] = o.RumTotalSessionCountSum.Get()
	}
	if o.RumUnitsSum.IsSet() {
		toSerialize["rum_units_sum"] = o.RumUnitsSum.Get()
	}
	if o.SdsApmScannedBytesSum.IsSet() {
		toSerialize["sds_apm_scanned_bytes_sum"] = o.SdsApmScannedBytesSum.Get()
	}
	if o.SdsEventsScannedBytesSum.IsSet() {
		toSerialize["sds_events_scanned_bytes_sum"] = o.SdsEventsScannedBytesSum.Get()
	}
	if o.SdsLogsScannedBytesSum.IsSet() {
		toSerialize["sds_logs_scanned_bytes_sum"] = o.SdsLogsScannedBytesSum.Get()
	}
	if o.SdsRumScannedBytesSum.IsSet() {
		toSerialize["sds_rum_scanned_bytes_sum"] = o.SdsRumScannedBytesSum.Get()
	}
	if o.SdsTotalScannedBytesSum.IsSet() {
		toSerialize["sds_total_scanned_bytes_sum"] = o.SdsTotalScannedBytesSum.Get()
	}
	if o.SyntheticsBrowserCheckCallsCountSum.IsSet() {
		toSerialize["synthetics_browser_check_calls_count_sum"] = o.SyntheticsBrowserCheckCallsCountSum.Get()
	}
	if o.SyntheticsCheckCallsCountSum.IsSet() {
		toSerialize["synthetics_check_calls_count_sum"] = o.SyntheticsCheckCallsCountSum.Get()
	}
	if o.SyntheticsParallelTestingMaxSlotsHwm.IsSet() {
		toSerialize["synthetics_parallel_testing_max_slots_hwm"] = o.SyntheticsParallelTestingMaxSlotsHwm.Get()
	}
	if o.TraceSearchIndexedEventsCountSum.IsSet() {
		toSerialize["trace_search_indexed_events_count_sum"] = o.TraceSearchIndexedEventsCountSum.Get()
	}
	if o.TwolIngestedEventsBytesSum.IsSet() {
		toSerialize["twol_ingested_events_bytes_sum"] = o.TwolIngestedEventsBytesSum.Get()
	}
	if o.UniversalServiceMonitoringHostTop99p.IsSet() {
		toSerialize["universal_service_monitoring_host_top99p"] = o.UniversalServiceMonitoringHostTop99p.Get()
	}
	if o.VsphereHostTop99p.IsSet() {
		toSerialize["vsphere_host_top99p"] = o.VsphereHostTop99p.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageSummaryDateOrg) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AgentHostTop99p                         datadog.NullableInt64 `json:"agent_host_top99p,omitempty"`
		ApmAzureAppServiceHostTop99p            datadog.NullableInt64 `json:"apm_azure_app_service_host_top99p,omitempty"`
		ApmFargateCountAvg                      datadog.NullableInt64 `json:"apm_fargate_count_avg,omitempty"`
		ApmHostTop99p                           datadog.NullableInt64 `json:"apm_host_top99p,omitempty"`
		AppsecFargateCountAvg                   datadog.NullableInt64 `json:"appsec_fargate_count_avg,omitempty"`
		AuditLogsLinesIndexedSum                datadog.NullableInt64 `json:"audit_logs_lines_indexed_sum,omitempty"`
		AuditTrailEnabledHwm                    datadog.NullableInt64 `json:"audit_trail_enabled_hwm,omitempty"`
		AvgProfiledFargateTasks                 datadog.NullableInt64 `json:"avg_profiled_fargate_tasks,omitempty"`
		AwsHostTop99p                           datadog.NullableInt64 `json:"aws_host_top99p,omitempty"`
		AwsLambdaFuncCount                      *int64                `json:"aws_lambda_func_count,omitempty"`
		AwsLambdaInvocationsSum                 datadog.NullableInt64 `json:"aws_lambda_invocations_sum,omitempty"`
		AzureAppServiceTop99p                   datadog.NullableInt64 `json:"azure_app_service_top99p,omitempty"`
		BillableIngestedBytesSum                datadog.NullableInt64 `json:"billable_ingested_bytes_sum,omitempty"`
		BrowserRumLiteSessionCountSum           datadog.NullableInt64 `json:"browser_rum_lite_session_count_sum,omitempty"`
		BrowserRumReplaySessionCountSum         datadog.NullableInt64 `json:"browser_rum_replay_session_count_sum,omitempty"`
		BrowserRumUnitsSum                      datadog.NullableInt64 `json:"browser_rum_units_sum,omitempty"`
		CiPipelineIndexedSpansSum               datadog.NullableInt64 `json:"ci_pipeline_indexed_spans_sum,omitempty"`
		CiTestIndexedSpansSum                   datadog.NullableInt64 `json:"ci_test_indexed_spans_sum,omitempty"`
		CiVisibilityPipelineCommittersHwm       datadog.NullableInt64 `json:"ci_visibility_pipeline_committers_hwm,omitempty"`
		CiVisibilityTestCommittersHwm           datadog.NullableInt64 `json:"ci_visibility_test_committers_hwm,omitempty"`
		CloudCostManagementHostCountAvg         *int64                `json:"cloud_cost_management_host_count_avg,omitempty"`
		ContainerAvg                            datadog.NullableInt64 `json:"container_avg,omitempty"`
		ContainerExclAgentAvg                   datadog.NullableInt64 `json:"container_excl_agent_avg,omitempty"`
		ContainerHwm                            datadog.NullableInt64 `json:"container_hwm,omitempty"`
		CspmAasHostTop99p                       datadog.NullableInt64 `json:"cspm_aas_host_top99p,omitempty"`
		CspmAwsHostTop99p                       *int64                `json:"cspm_aws_host_top99p,omitempty"`
		CspmAzureHostTop99p                     datadog.NullableInt64 `json:"cspm_azure_host_top99p,omitempty"`
		CspmContainerAvg                        datadog.NullableInt64 `json:"cspm_container_avg,omitempty"`
		CspmContainerHwm                        datadog.NullableInt64 `json:"cspm_container_hwm,omitempty"`
		CspmGcpHostTop99p                       datadog.NullableInt64 `json:"cspm_gcp_host_top99p,omitempty"`
		CspmHostTop99p                          datadog.NullableInt64 `json:"cspm_host_top99p,omitempty"`
		CustomTsAvg                             datadog.NullableInt64 `json:"custom_ts_avg,omitempty"`
		CwsContainerCountAvg                    datadog.NullableInt64 `json:"cws_container_count_avg,omitempty"`
		CwsHostTop99p                           datadog.NullableInt64 `json:"cws_host_top99p,omitempty"`
		DbmHostTop99pSum                        datadog.NullableInt64 `json:"dbm_host_top99p_sum,omitempty"`
		DbmQueriesAvgSum                        datadog.NullableInt64 `json:"dbm_queries_avg_sum,omitempty"`
		FargateTasksCountAvg                    datadog.NullableInt64 `json:"fargate_tasks_count_avg,omitempty"`
		FargateTasksCountHwm                    datadog.NullableInt64 `json:"fargate_tasks_count_hwm,omitempty"`
		ForwardingEventsBytesSum                datadog.NullableInt64 `json:"forwarding_events_bytes_sum,omitempty"`
		GcpHostTop99p                           datadog.NullableInt64 `json:"gcp_host_top99p,omitempty"`
		HerokuHostTop99p                        datadog.NullableInt64 `json:"heroku_host_top99p,omitempty"`
		Id                                      *string               `json:"id,omitempty"`
		IncidentManagementMonthlyActiveUsersHwm datadog.NullableInt64 `json:"incident_management_monthly_active_users_hwm,omitempty"`
		IndexedEventsCountSum                   datadog.NullableInt64 `json:"indexed_events_count_sum,omitempty"`
		InfraHostTop99p                         datadog.NullableInt64 `json:"infra_host_top99p,omitempty"`
		IngestedEventsBytesSum                  datadog.NullableInt64 `json:"ingested_events_bytes_sum,omitempty"`
		IotDeviceAggSum                         datadog.NullableInt64 `json:"iot_device_agg_sum,omitempty"`
		IotDeviceTop99pSum                      datadog.NullableInt64 `json:"iot_device_top99p_sum,omitempty"`
		MobileRumLiteSessionCountSum            datadog.NullableInt64 `json:"mobile_rum_lite_session_count_sum,omitempty"`
		MobileRumSessionCountAndroidSum         datadog.NullableInt64 `json:"mobile_rum_session_count_android_sum,omitempty"`
		MobileRumSessionCountFlutterSum         datadog.NullableInt64 `json:"mobile_rum_session_count_flutter_sum,omitempty"`
		MobileRumSessionCountIosSum             datadog.NullableInt64 `json:"mobile_rum_session_count_ios_sum,omitempty"`
		MobileRumSessionCountReactnativeSum     datadog.NullableInt64 `json:"mobile_rum_session_count_reactnative_sum,omitempty"`
		MobileRumSessionCountSum                datadog.NullableInt64 `json:"mobile_rum_session_count_sum,omitempty"`
		MobileRumUnitsSum                       datadog.NullableInt64 `json:"mobile_rum_units_sum,omitempty"`
		Name                                    *string               `json:"name,omitempty"`
		NetflowIndexedEventsCountSum            datadog.NullableInt64 `json:"netflow_indexed_events_count_sum,omitempty"`
		NpmHostTop99p                           datadog.NullableInt64 `json:"npm_host_top99p,omitempty"`
		ObservabilityPipelinesBytesProcessedSum datadog.NullableInt64 `json:"observability_pipelines_bytes_processed_sum,omitempty"`
		OnlineArchiveEventsCountSum             datadog.NullableInt64 `json:"online_archive_events_count_sum,omitempty"`
		OpentelemetryApmHostTop99p              datadog.NullableInt64 `json:"opentelemetry_apm_host_top99p,omitempty"`
		OpentelemetryHostTop99p                 datadog.NullableInt64 `json:"opentelemetry_host_top99p,omitempty"`
		ProfilingHostTop99p                     datadog.NullableInt64 `json:"profiling_host_top99p,omitempty"`
		PublicId                                *string               `json:"public_id,omitempty"`
		Region                                  *string               `json:"region,omitempty"`
		RumBrowserAndMobileSessionCount         datadog.NullableInt64 `json:"rum_browser_and_mobile_session_count,omitempty"`
		RumSessionCountSum                      datadog.NullableInt64 `json:"rum_session_count_sum,omitempty"`
		RumTotalSessionCountSum                 datadog.NullableInt64 `json:"rum_total_session_count_sum,omitempty"`
		RumUnitsSum                             datadog.NullableInt64 `json:"rum_units_sum,omitempty"`
		SdsApmScannedBytesSum                   datadog.NullableInt64 `json:"sds_apm_scanned_bytes_sum,omitempty"`
		SdsEventsScannedBytesSum                datadog.NullableInt64 `json:"sds_events_scanned_bytes_sum,omitempty"`
		SdsLogsScannedBytesSum                  datadog.NullableInt64 `json:"sds_logs_scanned_bytes_sum,omitempty"`
		SdsRumScannedBytesSum                   datadog.NullableInt64 `json:"sds_rum_scanned_bytes_sum,omitempty"`
		SdsTotalScannedBytesSum                 datadog.NullableInt64 `json:"sds_total_scanned_bytes_sum,omitempty"`
		SyntheticsBrowserCheckCallsCountSum     datadog.NullableInt64 `json:"synthetics_browser_check_calls_count_sum,omitempty"`
		SyntheticsCheckCallsCountSum            datadog.NullableInt64 `json:"synthetics_check_calls_count_sum,omitempty"`
		SyntheticsParallelTestingMaxSlotsHwm    datadog.NullableInt64 `json:"synthetics_parallel_testing_max_slots_hwm,omitempty"`
		TraceSearchIndexedEventsCountSum        datadog.NullableInt64 `json:"trace_search_indexed_events_count_sum,omitempty"`
		TwolIngestedEventsBytesSum              datadog.NullableInt64 `json:"twol_ingested_events_bytes_sum,omitempty"`
		UniversalServiceMonitoringHostTop99p    datadog.NullableInt64 `json:"universal_service_monitoring_host_top99p,omitempty"`
		VsphereHostTop99p                       datadog.NullableInt64 `json:"vsphere_host_top99p,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_host_top99p", "apm_azure_app_service_host_top99p", "apm_fargate_count_avg", "apm_host_top99p", "appsec_fargate_count_avg", "audit_logs_lines_indexed_sum", "audit_trail_enabled_hwm", "avg_profiled_fargate_tasks", "aws_host_top99p", "aws_lambda_func_count", "aws_lambda_invocations_sum", "azure_app_service_top99p", "billable_ingested_bytes_sum", "browser_rum_lite_session_count_sum", "browser_rum_replay_session_count_sum", "browser_rum_units_sum", "ci_pipeline_indexed_spans_sum", "ci_test_indexed_spans_sum", "ci_visibility_pipeline_committers_hwm", "ci_visibility_test_committers_hwm", "cloud_cost_management_host_count_avg", "container_avg", "container_excl_agent_avg", "container_hwm", "cspm_aas_host_top99p", "cspm_aws_host_top99p", "cspm_azure_host_top99p", "cspm_container_avg", "cspm_container_hwm", "cspm_gcp_host_top99p", "cspm_host_top99p", "custom_ts_avg", "cws_container_count_avg", "cws_host_top99p", "dbm_host_top99p_sum", "dbm_queries_avg_sum", "fargate_tasks_count_avg", "fargate_tasks_count_hwm", "forwarding_events_bytes_sum", "gcp_host_top99p", "heroku_host_top99p", "id", "incident_management_monthly_active_users_hwm", "indexed_events_count_sum", "infra_host_top99p", "ingested_events_bytes_sum", "iot_device_agg_sum", "iot_device_top99p_sum", "mobile_rum_lite_session_count_sum", "mobile_rum_session_count_android_sum", "mobile_rum_session_count_flutter_sum", "mobile_rum_session_count_ios_sum", "mobile_rum_session_count_reactnative_sum", "mobile_rum_session_count_sum", "mobile_rum_units_sum", "name", "netflow_indexed_events_count_sum", "npm_host_top99p", "observability_pipelines_bytes_processed_sum", "online_archive_events_count_sum", "opentelemetry_apm_host_top99p", "opentelemetry_host_top99p", "profiling_host_top99p", "public_id", "region", "rum_browser_and_mobile_session_count", "rum_session_count_sum", "rum_total_session_count_sum", "rum_units_sum", "sds_apm_scanned_bytes_sum", "sds_events_scanned_bytes_sum", "sds_logs_scanned_bytes_sum", "sds_rum_scanned_bytes_sum", "sds_total_scanned_bytes_sum", "synthetics_browser_check_calls_count_sum", "synthetics_check_calls_count_sum", "synthetics_parallel_testing_max_slots_hwm", "trace_search_indexed_events_count_sum", "twol_ingested_events_bytes_sum", "universal_service_monitoring_host_top99p", "vsphere_host_top99p"})
	} else {
		return err
	}
	o.AgentHostTop99p = all.AgentHostTop99p
	o.ApmAzureAppServiceHostTop99p = all.ApmAzureAppServiceHostTop99p
	o.ApmFargateCountAvg = all.ApmFargateCountAvg
	o.ApmHostTop99p = all.ApmHostTop99p
	o.AppsecFargateCountAvg = all.AppsecFargateCountAvg
	o.AuditLogsLinesIndexedSum = all.AuditLogsLinesIndexedSum
	o.AuditTrailEnabledHwm = all.AuditTrailEnabledHwm
	o.AvgProfiledFargateTasks = all.AvgProfiledFargateTasks
	o.AwsHostTop99p = all.AwsHostTop99p
	o.AwsLambdaFuncCount = all.AwsLambdaFuncCount
	o.AwsLambdaInvocationsSum = all.AwsLambdaInvocationsSum
	o.AzureAppServiceTop99p = all.AzureAppServiceTop99p
	o.BillableIngestedBytesSum = all.BillableIngestedBytesSum
	o.BrowserRumLiteSessionCountSum = all.BrowserRumLiteSessionCountSum
	o.BrowserRumReplaySessionCountSum = all.BrowserRumReplaySessionCountSum
	o.BrowserRumUnitsSum = all.BrowserRumUnitsSum
	o.CiPipelineIndexedSpansSum = all.CiPipelineIndexedSpansSum
	o.CiTestIndexedSpansSum = all.CiTestIndexedSpansSum
	o.CiVisibilityPipelineCommittersHwm = all.CiVisibilityPipelineCommittersHwm
	o.CiVisibilityTestCommittersHwm = all.CiVisibilityTestCommittersHwm
	o.CloudCostManagementHostCountAvg = all.CloudCostManagementHostCountAvg
	o.ContainerAvg = all.ContainerAvg
	o.ContainerExclAgentAvg = all.ContainerExclAgentAvg
	o.ContainerHwm = all.ContainerHwm
	o.CspmAasHostTop99p = all.CspmAasHostTop99p
	o.CspmAwsHostTop99p = all.CspmAwsHostTop99p
	o.CspmAzureHostTop99p = all.CspmAzureHostTop99p
	o.CspmContainerAvg = all.CspmContainerAvg
	o.CspmContainerHwm = all.CspmContainerHwm
	o.CspmGcpHostTop99p = all.CspmGcpHostTop99p
	o.CspmHostTop99p = all.CspmHostTop99p
	o.CustomTsAvg = all.CustomTsAvg
	o.CwsContainerCountAvg = all.CwsContainerCountAvg
	o.CwsHostTop99p = all.CwsHostTop99p
	o.DbmHostTop99pSum = all.DbmHostTop99pSum
	o.DbmQueriesAvgSum = all.DbmQueriesAvgSum
	o.FargateTasksCountAvg = all.FargateTasksCountAvg
	o.FargateTasksCountHwm = all.FargateTasksCountHwm
	o.ForwardingEventsBytesSum = all.ForwardingEventsBytesSum
	o.GcpHostTop99p = all.GcpHostTop99p
	o.HerokuHostTop99p = all.HerokuHostTop99p
	o.Id = all.Id
	o.IncidentManagementMonthlyActiveUsersHwm = all.IncidentManagementMonthlyActiveUsersHwm
	o.IndexedEventsCountSum = all.IndexedEventsCountSum
	o.InfraHostTop99p = all.InfraHostTop99p
	o.IngestedEventsBytesSum = all.IngestedEventsBytesSum
	o.IotDeviceAggSum = all.IotDeviceAggSum
	o.IotDeviceTop99pSum = all.IotDeviceTop99pSum
	o.MobileRumLiteSessionCountSum = all.MobileRumLiteSessionCountSum
	o.MobileRumSessionCountAndroidSum = all.MobileRumSessionCountAndroidSum
	o.MobileRumSessionCountFlutterSum = all.MobileRumSessionCountFlutterSum
	o.MobileRumSessionCountIosSum = all.MobileRumSessionCountIosSum
	o.MobileRumSessionCountReactnativeSum = all.MobileRumSessionCountReactnativeSum
	o.MobileRumSessionCountSum = all.MobileRumSessionCountSum
	o.MobileRumUnitsSum = all.MobileRumUnitsSum
	o.Name = all.Name
	o.NetflowIndexedEventsCountSum = all.NetflowIndexedEventsCountSum
	o.NpmHostTop99p = all.NpmHostTop99p
	o.ObservabilityPipelinesBytesProcessedSum = all.ObservabilityPipelinesBytesProcessedSum
	o.OnlineArchiveEventsCountSum = all.OnlineArchiveEventsCountSum
	o.OpentelemetryApmHostTop99p = all.OpentelemetryApmHostTop99p
	o.OpentelemetryHostTop99p = all.OpentelemetryHostTop99p
	o.ProfilingHostTop99p = all.ProfilingHostTop99p
	o.PublicId = all.PublicId
	o.Region = all.Region
	o.RumBrowserAndMobileSessionCount = all.RumBrowserAndMobileSessionCount
	o.RumSessionCountSum = all.RumSessionCountSum
	o.RumTotalSessionCountSum = all.RumTotalSessionCountSum
	o.RumUnitsSum = all.RumUnitsSum
	o.SdsApmScannedBytesSum = all.SdsApmScannedBytesSum
	o.SdsEventsScannedBytesSum = all.SdsEventsScannedBytesSum
	o.SdsLogsScannedBytesSum = all.SdsLogsScannedBytesSum
	o.SdsRumScannedBytesSum = all.SdsRumScannedBytesSum
	o.SdsTotalScannedBytesSum = all.SdsTotalScannedBytesSum
	o.SyntheticsBrowserCheckCallsCountSum = all.SyntheticsBrowserCheckCallsCountSum
	o.SyntheticsCheckCallsCountSum = all.SyntheticsCheckCallsCountSum
	o.SyntheticsParallelTestingMaxSlotsHwm = all.SyntheticsParallelTestingMaxSlotsHwm
	o.TraceSearchIndexedEventsCountSum = all.TraceSearchIndexedEventsCountSum
	o.TwolIngestedEventsBytesSum = all.TwolIngestedEventsBytesSum
	o.UniversalServiceMonitoringHostTop99p = all.UniversalServiceMonitoringHostTop99p
	o.VsphereHostTop99p = all.VsphereHostTop99p
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
