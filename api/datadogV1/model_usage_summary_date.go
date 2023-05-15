// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSummaryDate Response with hourly report of all data billed by Datadog all organizations.
type UsageSummaryDate struct {
	// Shows the 99th percentile of all agent hosts over all hours in the current date for all organizations.
	AgentHostTop99p datadog.NullableInt64 `json:"agent_host_top99p,omitempty"`
	// Shows the 99th percentile of all Azure app services using APM over all hours in the current date all organizations.
	ApmAzureAppServiceHostTop99p datadog.NullableInt64 `json:"apm_azure_app_service_host_top99p,omitempty"`
	// Shows the average of all APM ECS Fargate tasks over all hours in the current date for all organizations.
	ApmFargateCountAvg datadog.NullableInt64 `json:"apm_fargate_count_avg,omitempty"`
	// Shows the 99th percentile of all distinct APM hosts over all hours in the current date for all organizations.
	ApmHostTop99p datadog.NullableInt64 `json:"apm_host_top99p,omitempty"`
	// Shows the average of all Application Security Monitoring ECS Fargate tasks over all hours in the current date for all organizations.
	AppsecFargateCountAvg datadog.NullableInt64 `json:"appsec_fargate_count_avg,omitempty"`
	// Shows the sum of audit logs lines indexed over all hours in the current date for all organizations.
	// Deprecated
	AuditLogsLinesIndexedSum datadog.NullableInt64 `json:"audit_logs_lines_indexed_sum,omitempty"`
	// Shows the number of organizations that had Audit Trail enabled in the current date.
	AuditTrailEnabledHwm datadog.NullableInt64 `json:"audit_trail_enabled_hwm,omitempty"`
	// The average profiled task count for Fargate Profiling.
	AvgProfiledFargateTasks datadog.NullableInt64 `json:"avg_profiled_fargate_tasks,omitempty"`
	// Shows the 99th percentile of all AWS hosts over all hours in the current date for all organizations.
	AwsHostTop99p datadog.NullableInt64 `json:"aws_host_top99p,omitempty"`
	// Shows the average of the number of functions that executed 1 or more times each hour in the current date for all organizations.
	AwsLambdaFuncCount datadog.NullableInt64 `json:"aws_lambda_func_count,omitempty"`
	// Shows the sum of all AWS Lambda invocations over all hours in the current date for all organizations.
	AwsLambdaInvocationsSum datadog.NullableInt64 `json:"aws_lambda_invocations_sum,omitempty"`
	// Shows the 99th percentile of all Azure app services over all hours in the current date for all organizations.
	AzureAppServiceTop99p datadog.NullableInt64 `json:"azure_app_service_top99p,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current date for all organizations.
	BillableIngestedBytesSum datadog.NullableInt64 `json:"billable_ingested_bytes_sum,omitempty"`
	// Shows the sum of all browser lite sessions over all hours in the current date for all organizations.
	BrowserRumLiteSessionCountSum datadog.NullableInt64 `json:"browser_rum_lite_session_count_sum,omitempty"`
	// Shows the sum of all browser replay sessions over all hours in the current date for all organizations.
	BrowserRumReplaySessionCountSum datadog.NullableInt64 `json:"browser_rum_replay_session_count_sum,omitempty"`
	// Shows the sum of all browser RUM units over all hours in the current date for all organizations.
	BrowserRumUnitsSum datadog.NullableInt64 `json:"browser_rum_units_sum,omitempty"`
	// Shows the sum of all CI pipeline indexed spans over all hours in the current month for all organizations.
	CiPipelineIndexedSpansSum datadog.NullableInt64 `json:"ci_pipeline_indexed_spans_sum,omitempty"`
	// Shows the sum of all CI test indexed spans over all hours in the current month for all organizations.
	CiTestIndexedSpansSum datadog.NullableInt64 `json:"ci_test_indexed_spans_sum,omitempty"`
	// Shows the high-water mark of all CI visibility pipeline committers over all hours in the current month for all organizations.
	CiVisibilityPipelineCommittersHwm datadog.NullableInt64 `json:"ci_visibility_pipeline_committers_hwm,omitempty"`
	// Shows the high-water mark of all CI visibility test committers over all hours in the current month for all organizations.
	CiVisibilityTestCommittersHwm datadog.NullableInt64 `json:"ci_visibility_test_committers_hwm,omitempty"`
	// Host count average of Cloud Cost Management for the given date and given organization.
	CloudCostManagementHostCountAvg datadog.NullableInt64 `json:"cloud_cost_management_host_count_avg,omitempty"`
	// Shows the average of all distinct containers over all hours in the current date for all organizations.
	ContainerAvg datadog.NullableInt64 `json:"container_avg,omitempty"`
	// Shows the average of containers without the Datadog Agent over all hours in the current date for all organizations.
	ContainerExclAgentAvg datadog.NullableInt64 `json:"container_excl_agent_avg,omitempty"`
	// Shows the high-water mark of all distinct containers over all hours in the current date for all organizations.
	ContainerHwm datadog.NullableInt64 `json:"container_hwm,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management Azure app services hosts over all hours in the current date for all organizations.
	CspmAasHostTop99p datadog.NullableInt64 `json:"cspm_aas_host_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management AWS hosts over all hours in the current date for all organizations.
	CspmAwsHostTop99p datadog.NullableInt64 `json:"cspm_aws_host_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management Azure hosts over all hours in the current date for all organizations.
	CspmAzureHostTop99p datadog.NullableInt64 `json:"cspm_azure_host_top99p,omitempty"`
	// Shows the average number of Cloud Security Posture Management containers over all hours in the current date for all organizations.
	CspmContainerAvg datadog.NullableInt64 `json:"cspm_container_avg,omitempty"`
	// Shows the high-water mark of Cloud Security Posture Management containers over all hours in the current date for all organizations.
	CspmContainerHwm datadog.NullableInt64 `json:"cspm_container_hwm,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management GCP hosts over all hours in the current date for all organizations.
	CspmGcpHostTop99p datadog.NullableInt64 `json:"cspm_gcp_host_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management hosts over all hours in the current date for all organizations.
	CspmHostTop99p datadog.NullableInt64 `json:"cspm_host_top99p,omitempty"`
	// Shows the average number of distinct custom metrics over all hours in the current date for all organizations.
	CustomTsAvg datadog.NullableInt64 `json:"custom_ts_avg,omitempty"`
	// Shows the average of all distinct Cloud Workload Security containers over all hours in the current date for all organizations.
	CwsContainerCountAvg datadog.NullableInt64 `json:"cws_container_count_avg,omitempty"`
	// Shows the 99th percentile of all Cloud Workload Security hosts over all hours in the current date for all organizations.
	CwsHostTop99p datadog.NullableInt64 `json:"cws_host_top99p,omitempty"`
	// The date for the usage.
	Date *time.Time `json:"date,omitempty"`
	// Shows the 99th percentile of all Database Monitoring hosts over all hours in the current date for all organizations.
	DbmHostTop99p datadog.NullableInt64 `json:"dbm_host_top99p,omitempty"`
	// Shows the average of all normalized Database Monitoring queries over all hours in the current date for all organizations.
	DbmQueriesCountAvg datadog.NullableInt64 `json:"dbm_queries_count_avg,omitempty"`
	// Shows the high-watermark of all Fargate tasks over all hours in the current date for all organizations.
	FargateTasksCountAvg datadog.NullableInt64 `json:"fargate_tasks_count_avg,omitempty"`
	// Shows the average of all Fargate tasks over all hours in the current date for all organizations.
	FargateTasksCountHwm datadog.NullableInt64 `json:"fargate_tasks_count_hwm,omitempty"`
	// Shows the sum of all log bytes forwarded over all hours in the current date for all organizations.
	ForwardingEventsBytesSum datadog.NullableInt64 `json:"forwarding_events_bytes_sum,omitempty"`
	// Shows the 99th percentile of all GCP hosts over all hours in the current date for all organizations.
	GcpHostTop99p datadog.NullableInt64 `json:"gcp_host_top99p,omitempty"`
	// Shows the 99th percentile of all Heroku dynos over all hours in the current date for all organizations.
	HerokuHostTop99p datadog.NullableInt64 `json:"heroku_host_top99p,omitempty"`
	// Shows the high-water mark of incident management monthly active users over all hours in the current date for all organizations.
	IncidentManagementMonthlyActiveUsersHwm datadog.NullableInt64 `json:"incident_management_monthly_active_users_hwm,omitempty"`
	// Shows the sum of all log events indexed over all hours in the current date for all organizations.
	IndexedEventsCountSum datadog.NullableInt64 `json:"indexed_events_count_sum,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for all organizations.
	InfraHostTop99p datadog.NullableInt64 `json:"infra_host_top99p,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current date for all organizations.
	IngestedEventsBytesSum datadog.NullableInt64 `json:"ingested_events_bytes_sum,omitempty"`
	// Shows the sum of all IoT devices over all hours in the current date for all organizations.
	IotDeviceSum datadog.NullableInt64 `json:"iot_device_sum,omitempty"`
	// Shows the 99th percentile of all IoT devices over all hours in the current date all organizations.
	IotDeviceTop99p datadog.NullableInt64 `json:"iot_device_top99p,omitempty"`
	// Shows the sum of all mobile lite sessions over all hours in the current date for all organizations.
	MobileRumLiteSessionCountSum datadog.NullableInt64 `json:"mobile_rum_lite_session_count_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on Android over all hours in the current date for all organizations.
	MobileRumSessionCountAndroidSum datadog.NullableInt64 `json:"mobile_rum_session_count_android_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on Flutter over all hours in the current date for all organizations.
	MobileRumSessionCountFlutterSum datadog.NullableInt64 `json:"mobile_rum_session_count_flutter_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on iOS over all hours in the current date for all organizations.
	MobileRumSessionCountIosSum datadog.NullableInt64 `json:"mobile_rum_session_count_ios_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on React Native over all hours in the current date for all organizations.
	MobileRumSessionCountReactnativeSum datadog.NullableInt64 `json:"mobile_rum_session_count_reactnative_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions over all hours in the current date for all organizations
	MobileRumSessionCountSum datadog.NullableInt64 `json:"mobile_rum_session_count_sum,omitempty"`
	// Shows the sum of all mobile RUM units over all hours in the current date for all organizations.
	MobileRumUnitsSum datadog.NullableInt64 `json:"mobile_rum_units_sum,omitempty"`
	// Shows the sum of all Network flows indexed over all hours in the current date for all organizations.
	NetflowIndexedEventsCountSum datadog.NullableInt64 `json:"netflow_indexed_events_count_sum,omitempty"`
	// Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for all organizations.
	NpmHostTop99p datadog.NullableInt64 `json:"npm_host_top99p,omitempty"`
	// Sum of all observability pipelines bytes processed over all hours in the current date for the given org.
	ObservabilityPipelinesBytesProcessedSum datadog.NullableInt64 `json:"observability_pipelines_bytes_processed_sum,omitempty"`
	// Sum of all online archived events over all hours in the current date for all organizations.
	OnlineArchiveEventsCountSum datadog.NullableInt64 `json:"online_archive_events_count_sum,omitempty"`
	// Shows the 99th percentile of APM hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for all organizations.
	OpentelemetryApmHostTop99p datadog.NullableInt64 `json:"opentelemetry_apm_host_top99p,omitempty"`
	// Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for all organizations.
	OpentelemetryHostTop99p datadog.NullableInt64 `json:"opentelemetry_host_top99p,omitempty"`
	// Organizations associated with a user.
	Orgs []UsageSummaryDateOrg `json:"orgs,omitempty"`
	// Shows the 99th percentile of all profiled hosts over all hours in the current date for all organizations.
	ProfilingHostTop99p datadog.NullableInt64 `json:"profiling_host_top99p,omitempty"`
	// Shows the sum of all mobile sessions and all browser lite and legacy sessions over all hours in the current month for all organizations.
	RumBrowserAndMobileSessionCount datadog.NullableInt64 `json:"rum_browser_and_mobile_session_count,omitempty"`
	// Shows the sum of all browser RUM Lite Sessions over all hours in the current date for all organizations
	RumSessionCountSum datadog.NullableInt64 `json:"rum_session_count_sum,omitempty"`
	// Shows the sum of RUM Sessions (browser and mobile) over all hours in the current date for all organizations.
	RumTotalSessionCountSum datadog.NullableInt64 `json:"rum_total_session_count_sum,omitempty"`
	// Shows the sum of all browser and mobile RUM units over all hours in the current date for all organizations.
	RumUnitsSum datadog.NullableInt64 `json:"rum_units_sum,omitempty"`
	// Sum of all APM bytes scanned with sensitive data scanner over all hours in the current date for all organizations.
	SdsApmScannedBytesSum datadog.NullableInt64 `json:"sds_apm_scanned_bytes_sum,omitempty"`
	// Sum of all event stream events bytes scanned with sensitive data scanner over all hours in the current date for all organizations.
	SdsEventsScannedBytesSum datadog.NullableInt64 `json:"sds_events_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned of logs usage by the Sensitive Data Scanner over all hours in the current month for all organizations.
	SdsLogsScannedBytesSum datadog.NullableInt64 `json:"sds_logs_scanned_bytes_sum,omitempty"`
	// Sum of all RUM bytes scanned with sensitive data scanner over all hours in the current date for all organizations.
	SdsRumScannedBytesSum datadog.NullableInt64 `json:"sds_rum_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned across all usage types by the Sensitive Data Scanner over all hours in the current month for all organizations.
	SdsTotalScannedBytesSum datadog.NullableInt64 `json:"sds_total_scanned_bytes_sum,omitempty"`
	// Shows the sum of all Synthetic browser tests over all hours in the current date for all organizations.
	SyntheticsBrowserCheckCallsCountSum datadog.NullableInt64 `json:"synthetics_browser_check_calls_count_sum,omitempty"`
	// Shows the sum of all Synthetic API tests over all hours in the current date for all organizations.
	SyntheticsCheckCallsCountSum datadog.NullableInt64 `json:"synthetics_check_calls_count_sum,omitempty"`
	// Shows the high-water mark of used synthetics parallel testing slots over all hours in the current date for all organizations.
	SyntheticsParallelTestingMaxSlotsHwm datadog.NullableInt64 `json:"synthetics_parallel_testing_max_slots_hwm,omitempty"`
	// Shows the sum of all Indexed Spans indexed over all hours in the current date for all organizations.
	TraceSearchIndexedEventsCountSum datadog.NullableInt64 `json:"trace_search_indexed_events_count_sum,omitempty"`
	// Shows the sum of all ingested APM span bytes over all hours in the current date for all organizations.
	TwolIngestedEventsBytesSum datadog.NullableInt64 `json:"twol_ingested_events_bytes_sum,omitempty"`
	// Shows the 99th percentile of all universal service management hosts over all hours in the current date for the given org.
	UniversalServiceMonitoringHostTop99p datadog.NullableInt64 `json:"universal_service_monitoring_host_top99p,omitempty"`
	// Shows the 99th percentile of all vSphere hosts over all hours in the current date for all organizations.
	VsphereHostTop99p datadog.NullableInt64 `json:"vsphere_host_top99p,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageSummaryDate instantiates a new UsageSummaryDate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSummaryDate() *UsageSummaryDate {
	this := UsageSummaryDate{}
	return &this
}

// NewUsageSummaryDateWithDefaults instantiates a new UsageSummaryDate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSummaryDateWithDefaults() *UsageSummaryDate {
	this := UsageSummaryDate{}
	return &this
}

// GetAgentHostTop99p returns the AgentHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetAgentHostTop99p() int64 {
	if o == nil || o.AgentHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AgentHostTop99p.Get()
}

// GetAgentHostTop99pOk returns a tuple with the AgentHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetAgentHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentHostTop99p.Get(), o.AgentHostTop99p.IsSet()
}

// HasAgentHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAgentHostTop99p() bool {
	return o != nil && o.AgentHostTop99p.IsSet()
}

// SetAgentHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the AgentHostTop99p field.
func (o *UsageSummaryDate) SetAgentHostTop99p(v int64) {
	o.AgentHostTop99p.Set(&v)
}

// SetAgentHostTop99pNil sets the value for AgentHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetAgentHostTop99pNil() {
	o.AgentHostTop99p.Set(nil)
}

// UnsetAgentHostTop99p ensures that no value is present for AgentHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetAgentHostTop99p() {
	o.AgentHostTop99p.Unset()
}

// GetApmAzureAppServiceHostTop99p returns the ApmAzureAppServiceHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetApmAzureAppServiceHostTop99p() int64 {
	if o == nil || o.ApmAzureAppServiceHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmAzureAppServiceHostTop99p.Get()
}

// GetApmAzureAppServiceHostTop99pOk returns a tuple with the ApmAzureAppServiceHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetApmAzureAppServiceHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmAzureAppServiceHostTop99p.Get(), o.ApmAzureAppServiceHostTop99p.IsSet()
}

// HasApmAzureAppServiceHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmAzureAppServiceHostTop99p() bool {
	return o != nil && o.ApmAzureAppServiceHostTop99p.IsSet()
}

// SetApmAzureAppServiceHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the ApmAzureAppServiceHostTop99p field.
func (o *UsageSummaryDate) SetApmAzureAppServiceHostTop99p(v int64) {
	o.ApmAzureAppServiceHostTop99p.Set(&v)
}

// SetApmAzureAppServiceHostTop99pNil sets the value for ApmAzureAppServiceHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetApmAzureAppServiceHostTop99pNil() {
	o.ApmAzureAppServiceHostTop99p.Set(nil)
}

// UnsetApmAzureAppServiceHostTop99p ensures that no value is present for ApmAzureAppServiceHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetApmAzureAppServiceHostTop99p() {
	o.ApmAzureAppServiceHostTop99p.Unset()
}

// GetApmFargateCountAvg returns the ApmFargateCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetApmFargateCountAvg() int64 {
	if o == nil || o.ApmFargateCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmFargateCountAvg.Get()
}

// GetApmFargateCountAvgOk returns a tuple with the ApmFargateCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetApmFargateCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmFargateCountAvg.Get(), o.ApmFargateCountAvg.IsSet()
}

// HasApmFargateCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmFargateCountAvg() bool {
	return o != nil && o.ApmFargateCountAvg.IsSet()
}

// SetApmFargateCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the ApmFargateCountAvg field.
func (o *UsageSummaryDate) SetApmFargateCountAvg(v int64) {
	o.ApmFargateCountAvg.Set(&v)
}

// SetApmFargateCountAvgNil sets the value for ApmFargateCountAvg to be an explicit nil.
func (o *UsageSummaryDate) SetApmFargateCountAvgNil() {
	o.ApmFargateCountAvg.Set(nil)
}

// UnsetApmFargateCountAvg ensures that no value is present for ApmFargateCountAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetApmFargateCountAvg() {
	o.ApmFargateCountAvg.Unset()
}

// GetApmHostTop99p returns the ApmHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetApmHostTop99p() int64 {
	if o == nil || o.ApmHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmHostTop99p.Get()
}

// GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetApmHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmHostTop99p.Get(), o.ApmHostTop99p.IsSet()
}

// HasApmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmHostTop99p() bool {
	return o != nil && o.ApmHostTop99p.IsSet()
}

// SetApmHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the ApmHostTop99p field.
func (o *UsageSummaryDate) SetApmHostTop99p(v int64) {
	o.ApmHostTop99p.Set(&v)
}

// SetApmHostTop99pNil sets the value for ApmHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetApmHostTop99pNil() {
	o.ApmHostTop99p.Set(nil)
}

// UnsetApmHostTop99p ensures that no value is present for ApmHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetApmHostTop99p() {
	o.ApmHostTop99p.Unset()
}

// GetAppsecFargateCountAvg returns the AppsecFargateCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetAppsecFargateCountAvg() int64 {
	if o == nil || o.AppsecFargateCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AppsecFargateCountAvg.Get()
}

// GetAppsecFargateCountAvgOk returns a tuple with the AppsecFargateCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetAppsecFargateCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecFargateCountAvg.Get(), o.AppsecFargateCountAvg.IsSet()
}

// HasAppsecFargateCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAppsecFargateCountAvg() bool {
	return o != nil && o.AppsecFargateCountAvg.IsSet()
}

// SetAppsecFargateCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the AppsecFargateCountAvg field.
func (o *UsageSummaryDate) SetAppsecFargateCountAvg(v int64) {
	o.AppsecFargateCountAvg.Set(&v)
}

// SetAppsecFargateCountAvgNil sets the value for AppsecFargateCountAvg to be an explicit nil.
func (o *UsageSummaryDate) SetAppsecFargateCountAvgNil() {
	o.AppsecFargateCountAvg.Set(nil)
}

// UnsetAppsecFargateCountAvg ensures that no value is present for AppsecFargateCountAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetAppsecFargateCountAvg() {
	o.AppsecFargateCountAvg.Unset()
}

// GetAuditLogsLinesIndexedSum returns the AuditLogsLinesIndexedSum field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *UsageSummaryDate) GetAuditLogsLinesIndexedSum() int64 {
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
func (o *UsageSummaryDate) GetAuditLogsLinesIndexedSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditLogsLinesIndexedSum.Get(), o.AuditLogsLinesIndexedSum.IsSet()
}

// HasAuditLogsLinesIndexedSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAuditLogsLinesIndexedSum() bool {
	return o != nil && o.AuditLogsLinesIndexedSum.IsSet()
}

// SetAuditLogsLinesIndexedSum gets a reference to the given datadog.NullableInt64 and assigns it to the AuditLogsLinesIndexedSum field.
// Deprecated
func (o *UsageSummaryDate) SetAuditLogsLinesIndexedSum(v int64) {
	o.AuditLogsLinesIndexedSum.Set(&v)
}

// SetAuditLogsLinesIndexedSumNil sets the value for AuditLogsLinesIndexedSum to be an explicit nil.
func (o *UsageSummaryDate) SetAuditLogsLinesIndexedSumNil() {
	o.AuditLogsLinesIndexedSum.Set(nil)
}

// UnsetAuditLogsLinesIndexedSum ensures that no value is present for AuditLogsLinesIndexedSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetAuditLogsLinesIndexedSum() {
	o.AuditLogsLinesIndexedSum.Unset()
}

// GetAuditTrailEnabledHwm returns the AuditTrailEnabledHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetAuditTrailEnabledHwm() int64 {
	if o == nil || o.AuditTrailEnabledHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AuditTrailEnabledHwm.Get()
}

// GetAuditTrailEnabledHwmOk returns a tuple with the AuditTrailEnabledHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetAuditTrailEnabledHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditTrailEnabledHwm.Get(), o.AuditTrailEnabledHwm.IsSet()
}

// HasAuditTrailEnabledHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAuditTrailEnabledHwm() bool {
	return o != nil && o.AuditTrailEnabledHwm.IsSet()
}

// SetAuditTrailEnabledHwm gets a reference to the given datadog.NullableInt64 and assigns it to the AuditTrailEnabledHwm field.
func (o *UsageSummaryDate) SetAuditTrailEnabledHwm(v int64) {
	o.AuditTrailEnabledHwm.Set(&v)
}

// SetAuditTrailEnabledHwmNil sets the value for AuditTrailEnabledHwm to be an explicit nil.
func (o *UsageSummaryDate) SetAuditTrailEnabledHwmNil() {
	o.AuditTrailEnabledHwm.Set(nil)
}

// UnsetAuditTrailEnabledHwm ensures that no value is present for AuditTrailEnabledHwm, not even an explicit nil.
func (o *UsageSummaryDate) UnsetAuditTrailEnabledHwm() {
	o.AuditTrailEnabledHwm.Unset()
}

// GetAvgProfiledFargateTasks returns the AvgProfiledFargateTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetAvgProfiledFargateTasks() int64 {
	if o == nil || o.AvgProfiledFargateTasks.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AvgProfiledFargateTasks.Get()
}

// GetAvgProfiledFargateTasksOk returns a tuple with the AvgProfiledFargateTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetAvgProfiledFargateTasksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvgProfiledFargateTasks.Get(), o.AvgProfiledFargateTasks.IsSet()
}

// HasAvgProfiledFargateTasks returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAvgProfiledFargateTasks() bool {
	return o != nil && o.AvgProfiledFargateTasks.IsSet()
}

// SetAvgProfiledFargateTasks gets a reference to the given datadog.NullableInt64 and assigns it to the AvgProfiledFargateTasks field.
func (o *UsageSummaryDate) SetAvgProfiledFargateTasks(v int64) {
	o.AvgProfiledFargateTasks.Set(&v)
}

// SetAvgProfiledFargateTasksNil sets the value for AvgProfiledFargateTasks to be an explicit nil.
func (o *UsageSummaryDate) SetAvgProfiledFargateTasksNil() {
	o.AvgProfiledFargateTasks.Set(nil)
}

// UnsetAvgProfiledFargateTasks ensures that no value is present for AvgProfiledFargateTasks, not even an explicit nil.
func (o *UsageSummaryDate) UnsetAvgProfiledFargateTasks() {
	o.AvgProfiledFargateTasks.Unset()
}

// GetAwsHostTop99p returns the AwsHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetAwsHostTop99p() int64 {
	if o == nil || o.AwsHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AwsHostTop99p.Get()
}

// GetAwsHostTop99pOk returns a tuple with the AwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetAwsHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsHostTop99p.Get(), o.AwsHostTop99p.IsSet()
}

// HasAwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAwsHostTop99p() bool {
	return o != nil && o.AwsHostTop99p.IsSet()
}

// SetAwsHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the AwsHostTop99p field.
func (o *UsageSummaryDate) SetAwsHostTop99p(v int64) {
	o.AwsHostTop99p.Set(&v)
}

// SetAwsHostTop99pNil sets the value for AwsHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetAwsHostTop99pNil() {
	o.AwsHostTop99p.Set(nil)
}

// UnsetAwsHostTop99p ensures that no value is present for AwsHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetAwsHostTop99p() {
	o.AwsHostTop99p.Unset()
}

// GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetAwsLambdaFuncCount() int64 {
	if o == nil || o.AwsLambdaFuncCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaFuncCount.Get()
}

// GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetAwsLambdaFuncCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsLambdaFuncCount.Get(), o.AwsLambdaFuncCount.IsSet()
}

// HasAwsLambdaFuncCount returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAwsLambdaFuncCount() bool {
	return o != nil && o.AwsLambdaFuncCount.IsSet()
}

// SetAwsLambdaFuncCount gets a reference to the given datadog.NullableInt64 and assigns it to the AwsLambdaFuncCount field.
func (o *UsageSummaryDate) SetAwsLambdaFuncCount(v int64) {
	o.AwsLambdaFuncCount.Set(&v)
}

// SetAwsLambdaFuncCountNil sets the value for AwsLambdaFuncCount to be an explicit nil.
func (o *UsageSummaryDate) SetAwsLambdaFuncCountNil() {
	o.AwsLambdaFuncCount.Set(nil)
}

// UnsetAwsLambdaFuncCount ensures that no value is present for AwsLambdaFuncCount, not even an explicit nil.
func (o *UsageSummaryDate) UnsetAwsLambdaFuncCount() {
	o.AwsLambdaFuncCount.Unset()
}

// GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetAwsLambdaInvocationsSum() int64 {
	if o == nil || o.AwsLambdaInvocationsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaInvocationsSum.Get()
}

// GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetAwsLambdaInvocationsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsLambdaInvocationsSum.Get(), o.AwsLambdaInvocationsSum.IsSet()
}

// HasAwsLambdaInvocationsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAwsLambdaInvocationsSum() bool {
	return o != nil && o.AwsLambdaInvocationsSum.IsSet()
}

// SetAwsLambdaInvocationsSum gets a reference to the given datadog.NullableInt64 and assigns it to the AwsLambdaInvocationsSum field.
func (o *UsageSummaryDate) SetAwsLambdaInvocationsSum(v int64) {
	o.AwsLambdaInvocationsSum.Set(&v)
}

// SetAwsLambdaInvocationsSumNil sets the value for AwsLambdaInvocationsSum to be an explicit nil.
func (o *UsageSummaryDate) SetAwsLambdaInvocationsSumNil() {
	o.AwsLambdaInvocationsSum.Set(nil)
}

// UnsetAwsLambdaInvocationsSum ensures that no value is present for AwsLambdaInvocationsSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetAwsLambdaInvocationsSum() {
	o.AwsLambdaInvocationsSum.Unset()
}

// GetAzureAppServiceTop99p returns the AzureAppServiceTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetAzureAppServiceTop99p() int64 {
	if o == nil || o.AzureAppServiceTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AzureAppServiceTop99p.Get()
}

// GetAzureAppServiceTop99pOk returns a tuple with the AzureAppServiceTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetAzureAppServiceTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureAppServiceTop99p.Get(), o.AzureAppServiceTop99p.IsSet()
}

// HasAzureAppServiceTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAzureAppServiceTop99p() bool {
	return o != nil && o.AzureAppServiceTop99p.IsSet()
}

// SetAzureAppServiceTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the AzureAppServiceTop99p field.
func (o *UsageSummaryDate) SetAzureAppServiceTop99p(v int64) {
	o.AzureAppServiceTop99p.Set(&v)
}

// SetAzureAppServiceTop99pNil sets the value for AzureAppServiceTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetAzureAppServiceTop99pNil() {
	o.AzureAppServiceTop99p.Set(nil)
}

// UnsetAzureAppServiceTop99p ensures that no value is present for AzureAppServiceTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetAzureAppServiceTop99p() {
	o.AzureAppServiceTop99p.Unset()
}

// GetBillableIngestedBytesSum returns the BillableIngestedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetBillableIngestedBytesSum() int64 {
	if o == nil || o.BillableIngestedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BillableIngestedBytesSum.Get()
}

// GetBillableIngestedBytesSumOk returns a tuple with the BillableIngestedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetBillableIngestedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillableIngestedBytesSum.Get(), o.BillableIngestedBytesSum.IsSet()
}

// HasBillableIngestedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBillableIngestedBytesSum() bool {
	return o != nil && o.BillableIngestedBytesSum.IsSet()
}

// SetBillableIngestedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the BillableIngestedBytesSum field.
func (o *UsageSummaryDate) SetBillableIngestedBytesSum(v int64) {
	o.BillableIngestedBytesSum.Set(&v)
}

// SetBillableIngestedBytesSumNil sets the value for BillableIngestedBytesSum to be an explicit nil.
func (o *UsageSummaryDate) SetBillableIngestedBytesSumNil() {
	o.BillableIngestedBytesSum.Set(nil)
}

// UnsetBillableIngestedBytesSum ensures that no value is present for BillableIngestedBytesSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetBillableIngestedBytesSum() {
	o.BillableIngestedBytesSum.Unset()
}

// GetBrowserRumLiteSessionCountSum returns the BrowserRumLiteSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetBrowserRumLiteSessionCountSum() int64 {
	if o == nil || o.BrowserRumLiteSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumLiteSessionCountSum.Get()
}

// GetBrowserRumLiteSessionCountSumOk returns a tuple with the BrowserRumLiteSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetBrowserRumLiteSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumLiteSessionCountSum.Get(), o.BrowserRumLiteSessionCountSum.IsSet()
}

// HasBrowserRumLiteSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBrowserRumLiteSessionCountSum() bool {
	return o != nil && o.BrowserRumLiteSessionCountSum.IsSet()
}

// SetBrowserRumLiteSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumLiteSessionCountSum field.
func (o *UsageSummaryDate) SetBrowserRumLiteSessionCountSum(v int64) {
	o.BrowserRumLiteSessionCountSum.Set(&v)
}

// SetBrowserRumLiteSessionCountSumNil sets the value for BrowserRumLiteSessionCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetBrowserRumLiteSessionCountSumNil() {
	o.BrowserRumLiteSessionCountSum.Set(nil)
}

// UnsetBrowserRumLiteSessionCountSum ensures that no value is present for BrowserRumLiteSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetBrowserRumLiteSessionCountSum() {
	o.BrowserRumLiteSessionCountSum.Unset()
}

// GetBrowserRumReplaySessionCountSum returns the BrowserRumReplaySessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetBrowserRumReplaySessionCountSum() int64 {
	if o == nil || o.BrowserRumReplaySessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumReplaySessionCountSum.Get()
}

// GetBrowserRumReplaySessionCountSumOk returns a tuple with the BrowserRumReplaySessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetBrowserRumReplaySessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumReplaySessionCountSum.Get(), o.BrowserRumReplaySessionCountSum.IsSet()
}

// HasBrowserRumReplaySessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBrowserRumReplaySessionCountSum() bool {
	return o != nil && o.BrowserRumReplaySessionCountSum.IsSet()
}

// SetBrowserRumReplaySessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumReplaySessionCountSum field.
func (o *UsageSummaryDate) SetBrowserRumReplaySessionCountSum(v int64) {
	o.BrowserRumReplaySessionCountSum.Set(&v)
}

// SetBrowserRumReplaySessionCountSumNil sets the value for BrowserRumReplaySessionCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetBrowserRumReplaySessionCountSumNil() {
	o.BrowserRumReplaySessionCountSum.Set(nil)
}

// UnsetBrowserRumReplaySessionCountSum ensures that no value is present for BrowserRumReplaySessionCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetBrowserRumReplaySessionCountSum() {
	o.BrowserRumReplaySessionCountSum.Unset()
}

// GetBrowserRumUnitsSum returns the BrowserRumUnitsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetBrowserRumUnitsSum() int64 {
	if o == nil || o.BrowserRumUnitsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumUnitsSum.Get()
}

// GetBrowserRumUnitsSumOk returns a tuple with the BrowserRumUnitsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetBrowserRumUnitsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumUnitsSum.Get(), o.BrowserRumUnitsSum.IsSet()
}

// HasBrowserRumUnitsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBrowserRumUnitsSum() bool {
	return o != nil && o.BrowserRumUnitsSum.IsSet()
}

// SetBrowserRumUnitsSum gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumUnitsSum field.
func (o *UsageSummaryDate) SetBrowserRumUnitsSum(v int64) {
	o.BrowserRumUnitsSum.Set(&v)
}

// SetBrowserRumUnitsSumNil sets the value for BrowserRumUnitsSum to be an explicit nil.
func (o *UsageSummaryDate) SetBrowserRumUnitsSumNil() {
	o.BrowserRumUnitsSum.Set(nil)
}

// UnsetBrowserRumUnitsSum ensures that no value is present for BrowserRumUnitsSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetBrowserRumUnitsSum() {
	o.BrowserRumUnitsSum.Unset()
}

// GetCiPipelineIndexedSpansSum returns the CiPipelineIndexedSpansSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCiPipelineIndexedSpansSum() int64 {
	if o == nil || o.CiPipelineIndexedSpansSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiPipelineIndexedSpansSum.Get()
}

// GetCiPipelineIndexedSpansSumOk returns a tuple with the CiPipelineIndexedSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCiPipelineIndexedSpansSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiPipelineIndexedSpansSum.Get(), o.CiPipelineIndexedSpansSum.IsSet()
}

// HasCiPipelineIndexedSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCiPipelineIndexedSpansSum() bool {
	return o != nil && o.CiPipelineIndexedSpansSum.IsSet()
}

// SetCiPipelineIndexedSpansSum gets a reference to the given datadog.NullableInt64 and assigns it to the CiPipelineIndexedSpansSum field.
func (o *UsageSummaryDate) SetCiPipelineIndexedSpansSum(v int64) {
	o.CiPipelineIndexedSpansSum.Set(&v)
}

// SetCiPipelineIndexedSpansSumNil sets the value for CiPipelineIndexedSpansSum to be an explicit nil.
func (o *UsageSummaryDate) SetCiPipelineIndexedSpansSumNil() {
	o.CiPipelineIndexedSpansSum.Set(nil)
}

// UnsetCiPipelineIndexedSpansSum ensures that no value is present for CiPipelineIndexedSpansSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCiPipelineIndexedSpansSum() {
	o.CiPipelineIndexedSpansSum.Unset()
}

// GetCiTestIndexedSpansSum returns the CiTestIndexedSpansSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCiTestIndexedSpansSum() int64 {
	if o == nil || o.CiTestIndexedSpansSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiTestIndexedSpansSum.Get()
}

// GetCiTestIndexedSpansSumOk returns a tuple with the CiTestIndexedSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCiTestIndexedSpansSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiTestIndexedSpansSum.Get(), o.CiTestIndexedSpansSum.IsSet()
}

// HasCiTestIndexedSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCiTestIndexedSpansSum() bool {
	return o != nil && o.CiTestIndexedSpansSum.IsSet()
}

// SetCiTestIndexedSpansSum gets a reference to the given datadog.NullableInt64 and assigns it to the CiTestIndexedSpansSum field.
func (o *UsageSummaryDate) SetCiTestIndexedSpansSum(v int64) {
	o.CiTestIndexedSpansSum.Set(&v)
}

// SetCiTestIndexedSpansSumNil sets the value for CiTestIndexedSpansSum to be an explicit nil.
func (o *UsageSummaryDate) SetCiTestIndexedSpansSumNil() {
	o.CiTestIndexedSpansSum.Set(nil)
}

// UnsetCiTestIndexedSpansSum ensures that no value is present for CiTestIndexedSpansSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCiTestIndexedSpansSum() {
	o.CiTestIndexedSpansSum.Unset()
}

// GetCiVisibilityPipelineCommittersHwm returns the CiVisibilityPipelineCommittersHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCiVisibilityPipelineCommittersHwm() int64 {
	if o == nil || o.CiVisibilityPipelineCommittersHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityPipelineCommittersHwm.Get()
}

// GetCiVisibilityPipelineCommittersHwmOk returns a tuple with the CiVisibilityPipelineCommittersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCiVisibilityPipelineCommittersHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiVisibilityPipelineCommittersHwm.Get(), o.CiVisibilityPipelineCommittersHwm.IsSet()
}

// HasCiVisibilityPipelineCommittersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCiVisibilityPipelineCommittersHwm() bool {
	return o != nil && o.CiVisibilityPipelineCommittersHwm.IsSet()
}

// SetCiVisibilityPipelineCommittersHwm gets a reference to the given datadog.NullableInt64 and assigns it to the CiVisibilityPipelineCommittersHwm field.
func (o *UsageSummaryDate) SetCiVisibilityPipelineCommittersHwm(v int64) {
	o.CiVisibilityPipelineCommittersHwm.Set(&v)
}

// SetCiVisibilityPipelineCommittersHwmNil sets the value for CiVisibilityPipelineCommittersHwm to be an explicit nil.
func (o *UsageSummaryDate) SetCiVisibilityPipelineCommittersHwmNil() {
	o.CiVisibilityPipelineCommittersHwm.Set(nil)
}

// UnsetCiVisibilityPipelineCommittersHwm ensures that no value is present for CiVisibilityPipelineCommittersHwm, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCiVisibilityPipelineCommittersHwm() {
	o.CiVisibilityPipelineCommittersHwm.Unset()
}

// GetCiVisibilityTestCommittersHwm returns the CiVisibilityTestCommittersHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCiVisibilityTestCommittersHwm() int64 {
	if o == nil || o.CiVisibilityTestCommittersHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityTestCommittersHwm.Get()
}

// GetCiVisibilityTestCommittersHwmOk returns a tuple with the CiVisibilityTestCommittersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCiVisibilityTestCommittersHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiVisibilityTestCommittersHwm.Get(), o.CiVisibilityTestCommittersHwm.IsSet()
}

// HasCiVisibilityTestCommittersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCiVisibilityTestCommittersHwm() bool {
	return o != nil && o.CiVisibilityTestCommittersHwm.IsSet()
}

// SetCiVisibilityTestCommittersHwm gets a reference to the given datadog.NullableInt64 and assigns it to the CiVisibilityTestCommittersHwm field.
func (o *UsageSummaryDate) SetCiVisibilityTestCommittersHwm(v int64) {
	o.CiVisibilityTestCommittersHwm.Set(&v)
}

// SetCiVisibilityTestCommittersHwmNil sets the value for CiVisibilityTestCommittersHwm to be an explicit nil.
func (o *UsageSummaryDate) SetCiVisibilityTestCommittersHwmNil() {
	o.CiVisibilityTestCommittersHwm.Set(nil)
}

// UnsetCiVisibilityTestCommittersHwm ensures that no value is present for CiVisibilityTestCommittersHwm, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCiVisibilityTestCommittersHwm() {
	o.CiVisibilityTestCommittersHwm.Unset()
}

// GetCloudCostManagementHostCountAvg returns the CloudCostManagementHostCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCloudCostManagementHostCountAvg() int64 {
	if o == nil || o.CloudCostManagementHostCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementHostCountAvg.Get()
}

// GetCloudCostManagementHostCountAvgOk returns a tuple with the CloudCostManagementHostCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCloudCostManagementHostCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudCostManagementHostCountAvg.Get(), o.CloudCostManagementHostCountAvg.IsSet()
}

// HasCloudCostManagementHostCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCloudCostManagementHostCountAvg() bool {
	return o != nil && o.CloudCostManagementHostCountAvg.IsSet()
}

// SetCloudCostManagementHostCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the CloudCostManagementHostCountAvg field.
func (o *UsageSummaryDate) SetCloudCostManagementHostCountAvg(v int64) {
	o.CloudCostManagementHostCountAvg.Set(&v)
}

// SetCloudCostManagementHostCountAvgNil sets the value for CloudCostManagementHostCountAvg to be an explicit nil.
func (o *UsageSummaryDate) SetCloudCostManagementHostCountAvgNil() {
	o.CloudCostManagementHostCountAvg.Set(nil)
}

// UnsetCloudCostManagementHostCountAvg ensures that no value is present for CloudCostManagementHostCountAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCloudCostManagementHostCountAvg() {
	o.CloudCostManagementHostCountAvg.Unset()
}

// GetContainerAvg returns the ContainerAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetContainerAvg() int64 {
	if o == nil || o.ContainerAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerAvg.Get()
}

// GetContainerAvgOk returns a tuple with the ContainerAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetContainerAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerAvg.Get(), o.ContainerAvg.IsSet()
}

// HasContainerAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasContainerAvg() bool {
	return o != nil && o.ContainerAvg.IsSet()
}

// SetContainerAvg gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerAvg field.
func (o *UsageSummaryDate) SetContainerAvg(v int64) {
	o.ContainerAvg.Set(&v)
}

// SetContainerAvgNil sets the value for ContainerAvg to be an explicit nil.
func (o *UsageSummaryDate) SetContainerAvgNil() {
	o.ContainerAvg.Set(nil)
}

// UnsetContainerAvg ensures that no value is present for ContainerAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetContainerAvg() {
	o.ContainerAvg.Unset()
}

// GetContainerExclAgentAvg returns the ContainerExclAgentAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetContainerExclAgentAvg() int64 {
	if o == nil || o.ContainerExclAgentAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerExclAgentAvg.Get()
}

// GetContainerExclAgentAvgOk returns a tuple with the ContainerExclAgentAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetContainerExclAgentAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerExclAgentAvg.Get(), o.ContainerExclAgentAvg.IsSet()
}

// HasContainerExclAgentAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasContainerExclAgentAvg() bool {
	return o != nil && o.ContainerExclAgentAvg.IsSet()
}

// SetContainerExclAgentAvg gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerExclAgentAvg field.
func (o *UsageSummaryDate) SetContainerExclAgentAvg(v int64) {
	o.ContainerExclAgentAvg.Set(&v)
}

// SetContainerExclAgentAvgNil sets the value for ContainerExclAgentAvg to be an explicit nil.
func (o *UsageSummaryDate) SetContainerExclAgentAvgNil() {
	o.ContainerExclAgentAvg.Set(nil)
}

// UnsetContainerExclAgentAvg ensures that no value is present for ContainerExclAgentAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetContainerExclAgentAvg() {
	o.ContainerExclAgentAvg.Unset()
}

// GetContainerHwm returns the ContainerHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetContainerHwm() int64 {
	if o == nil || o.ContainerHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerHwm.Get()
}

// GetContainerHwmOk returns a tuple with the ContainerHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetContainerHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerHwm.Get(), o.ContainerHwm.IsSet()
}

// HasContainerHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasContainerHwm() bool {
	return o != nil && o.ContainerHwm.IsSet()
}

// SetContainerHwm gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerHwm field.
func (o *UsageSummaryDate) SetContainerHwm(v int64) {
	o.ContainerHwm.Set(&v)
}

// SetContainerHwmNil sets the value for ContainerHwm to be an explicit nil.
func (o *UsageSummaryDate) SetContainerHwmNil() {
	o.ContainerHwm.Set(nil)
}

// UnsetContainerHwm ensures that no value is present for ContainerHwm, not even an explicit nil.
func (o *UsageSummaryDate) UnsetContainerHwm() {
	o.ContainerHwm.Unset()
}

// GetCspmAasHostTop99p returns the CspmAasHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCspmAasHostTop99p() int64 {
	if o == nil || o.CspmAasHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmAasHostTop99p.Get()
}

// GetCspmAasHostTop99pOk returns a tuple with the CspmAasHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCspmAasHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmAasHostTop99p.Get(), o.CspmAasHostTop99p.IsSet()
}

// HasCspmAasHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmAasHostTop99p() bool {
	return o != nil && o.CspmAasHostTop99p.IsSet()
}

// SetCspmAasHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CspmAasHostTop99p field.
func (o *UsageSummaryDate) SetCspmAasHostTop99p(v int64) {
	o.CspmAasHostTop99p.Set(&v)
}

// SetCspmAasHostTop99pNil sets the value for CspmAasHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetCspmAasHostTop99pNil() {
	o.CspmAasHostTop99p.Set(nil)
}

// UnsetCspmAasHostTop99p ensures that no value is present for CspmAasHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCspmAasHostTop99p() {
	o.CspmAasHostTop99p.Unset()
}

// GetCspmAwsHostTop99p returns the CspmAwsHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCspmAwsHostTop99p() int64 {
	if o == nil || o.CspmAwsHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmAwsHostTop99p.Get()
}

// GetCspmAwsHostTop99pOk returns a tuple with the CspmAwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCspmAwsHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmAwsHostTop99p.Get(), o.CspmAwsHostTop99p.IsSet()
}

// HasCspmAwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmAwsHostTop99p() bool {
	return o != nil && o.CspmAwsHostTop99p.IsSet()
}

// SetCspmAwsHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CspmAwsHostTop99p field.
func (o *UsageSummaryDate) SetCspmAwsHostTop99p(v int64) {
	o.CspmAwsHostTop99p.Set(&v)
}

// SetCspmAwsHostTop99pNil sets the value for CspmAwsHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetCspmAwsHostTop99pNil() {
	o.CspmAwsHostTop99p.Set(nil)
}

// UnsetCspmAwsHostTop99p ensures that no value is present for CspmAwsHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCspmAwsHostTop99p() {
	o.CspmAwsHostTop99p.Unset()
}

// GetCspmAzureHostTop99p returns the CspmAzureHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCspmAzureHostTop99p() int64 {
	if o == nil || o.CspmAzureHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmAzureHostTop99p.Get()
}

// GetCspmAzureHostTop99pOk returns a tuple with the CspmAzureHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCspmAzureHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmAzureHostTop99p.Get(), o.CspmAzureHostTop99p.IsSet()
}

// HasCspmAzureHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmAzureHostTop99p() bool {
	return o != nil && o.CspmAzureHostTop99p.IsSet()
}

// SetCspmAzureHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CspmAzureHostTop99p field.
func (o *UsageSummaryDate) SetCspmAzureHostTop99p(v int64) {
	o.CspmAzureHostTop99p.Set(&v)
}

// SetCspmAzureHostTop99pNil sets the value for CspmAzureHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetCspmAzureHostTop99pNil() {
	o.CspmAzureHostTop99p.Set(nil)
}

// UnsetCspmAzureHostTop99p ensures that no value is present for CspmAzureHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCspmAzureHostTop99p() {
	o.CspmAzureHostTop99p.Unset()
}

// GetCspmContainerAvg returns the CspmContainerAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCspmContainerAvg() int64 {
	if o == nil || o.CspmContainerAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerAvg.Get()
}

// GetCspmContainerAvgOk returns a tuple with the CspmContainerAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCspmContainerAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainerAvg.Get(), o.CspmContainerAvg.IsSet()
}

// HasCspmContainerAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmContainerAvg() bool {
	return o != nil && o.CspmContainerAvg.IsSet()
}

// SetCspmContainerAvg gets a reference to the given datadog.NullableInt64 and assigns it to the CspmContainerAvg field.
func (o *UsageSummaryDate) SetCspmContainerAvg(v int64) {
	o.CspmContainerAvg.Set(&v)
}

// SetCspmContainerAvgNil sets the value for CspmContainerAvg to be an explicit nil.
func (o *UsageSummaryDate) SetCspmContainerAvgNil() {
	o.CspmContainerAvg.Set(nil)
}

// UnsetCspmContainerAvg ensures that no value is present for CspmContainerAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCspmContainerAvg() {
	o.CspmContainerAvg.Unset()
}

// GetCspmContainerHwm returns the CspmContainerHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCspmContainerHwm() int64 {
	if o == nil || o.CspmContainerHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerHwm.Get()
}

// GetCspmContainerHwmOk returns a tuple with the CspmContainerHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCspmContainerHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainerHwm.Get(), o.CspmContainerHwm.IsSet()
}

// HasCspmContainerHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmContainerHwm() bool {
	return o != nil && o.CspmContainerHwm.IsSet()
}

// SetCspmContainerHwm gets a reference to the given datadog.NullableInt64 and assigns it to the CspmContainerHwm field.
func (o *UsageSummaryDate) SetCspmContainerHwm(v int64) {
	o.CspmContainerHwm.Set(&v)
}

// SetCspmContainerHwmNil sets the value for CspmContainerHwm to be an explicit nil.
func (o *UsageSummaryDate) SetCspmContainerHwmNil() {
	o.CspmContainerHwm.Set(nil)
}

// UnsetCspmContainerHwm ensures that no value is present for CspmContainerHwm, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCspmContainerHwm() {
	o.CspmContainerHwm.Unset()
}

// GetCspmGcpHostTop99p returns the CspmGcpHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCspmGcpHostTop99p() int64 {
	if o == nil || o.CspmGcpHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmGcpHostTop99p.Get()
}

// GetCspmGcpHostTop99pOk returns a tuple with the CspmGcpHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCspmGcpHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmGcpHostTop99p.Get(), o.CspmGcpHostTop99p.IsSet()
}

// HasCspmGcpHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmGcpHostTop99p() bool {
	return o != nil && o.CspmGcpHostTop99p.IsSet()
}

// SetCspmGcpHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CspmGcpHostTop99p field.
func (o *UsageSummaryDate) SetCspmGcpHostTop99p(v int64) {
	o.CspmGcpHostTop99p.Set(&v)
}

// SetCspmGcpHostTop99pNil sets the value for CspmGcpHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetCspmGcpHostTop99pNil() {
	o.CspmGcpHostTop99p.Set(nil)
}

// UnsetCspmGcpHostTop99p ensures that no value is present for CspmGcpHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCspmGcpHostTop99p() {
	o.CspmGcpHostTop99p.Unset()
}

// GetCspmHostTop99p returns the CspmHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCspmHostTop99p() int64 {
	if o == nil || o.CspmHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmHostTop99p.Get()
}

// GetCspmHostTop99pOk returns a tuple with the CspmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCspmHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmHostTop99p.Get(), o.CspmHostTop99p.IsSet()
}

// HasCspmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmHostTop99p() bool {
	return o != nil && o.CspmHostTop99p.IsSet()
}

// SetCspmHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CspmHostTop99p field.
func (o *UsageSummaryDate) SetCspmHostTop99p(v int64) {
	o.CspmHostTop99p.Set(&v)
}

// SetCspmHostTop99pNil sets the value for CspmHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetCspmHostTop99pNil() {
	o.CspmHostTop99p.Set(nil)
}

// UnsetCspmHostTop99p ensures that no value is present for CspmHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCspmHostTop99p() {
	o.CspmHostTop99p.Unset()
}

// GetCustomTsAvg returns the CustomTsAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCustomTsAvg() int64 {
	if o == nil || o.CustomTsAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CustomTsAvg.Get()
}

// GetCustomTsAvgOk returns a tuple with the CustomTsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCustomTsAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTsAvg.Get(), o.CustomTsAvg.IsSet()
}

// HasCustomTsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCustomTsAvg() bool {
	return o != nil && o.CustomTsAvg.IsSet()
}

// SetCustomTsAvg gets a reference to the given datadog.NullableInt64 and assigns it to the CustomTsAvg field.
func (o *UsageSummaryDate) SetCustomTsAvg(v int64) {
	o.CustomTsAvg.Set(&v)
}

// SetCustomTsAvgNil sets the value for CustomTsAvg to be an explicit nil.
func (o *UsageSummaryDate) SetCustomTsAvgNil() {
	o.CustomTsAvg.Set(nil)
}

// UnsetCustomTsAvg ensures that no value is present for CustomTsAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCustomTsAvg() {
	o.CustomTsAvg.Unset()
}

// GetCwsContainerCountAvg returns the CwsContainerCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCwsContainerCountAvg() int64 {
	if o == nil || o.CwsContainerCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CwsContainerCountAvg.Get()
}

// GetCwsContainerCountAvgOk returns a tuple with the CwsContainerCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCwsContainerCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsContainerCountAvg.Get(), o.CwsContainerCountAvg.IsSet()
}

// HasCwsContainerCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCwsContainerCountAvg() bool {
	return o != nil && o.CwsContainerCountAvg.IsSet()
}

// SetCwsContainerCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the CwsContainerCountAvg field.
func (o *UsageSummaryDate) SetCwsContainerCountAvg(v int64) {
	o.CwsContainerCountAvg.Set(&v)
}

// SetCwsContainerCountAvgNil sets the value for CwsContainerCountAvg to be an explicit nil.
func (o *UsageSummaryDate) SetCwsContainerCountAvgNil() {
	o.CwsContainerCountAvg.Set(nil)
}

// UnsetCwsContainerCountAvg ensures that no value is present for CwsContainerCountAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCwsContainerCountAvg() {
	o.CwsContainerCountAvg.Unset()
}

// GetCwsHostTop99p returns the CwsHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetCwsHostTop99p() int64 {
	if o == nil || o.CwsHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CwsHostTop99p.Get()
}

// GetCwsHostTop99pOk returns a tuple with the CwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetCwsHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsHostTop99p.Get(), o.CwsHostTop99p.IsSet()
}

// HasCwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCwsHostTop99p() bool {
	return o != nil && o.CwsHostTop99p.IsSet()
}

// SetCwsHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the CwsHostTop99p field.
func (o *UsageSummaryDate) SetCwsHostTop99p(v int64) {
	o.CwsHostTop99p.Set(&v)
}

// SetCwsHostTop99pNil sets the value for CwsHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetCwsHostTop99pNil() {
	o.CwsHostTop99p.Set(nil)
}

// UnsetCwsHostTop99p ensures that no value is present for CwsHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetCwsHostTop99p() {
	o.CwsHostTop99p.Unset()
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
	return o != nil && o.Date != nil
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *UsageSummaryDate) SetDate(v time.Time) {
	o.Date = &v
}

// GetDbmHostTop99p returns the DbmHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetDbmHostTop99p() int64 {
	if o == nil || o.DbmHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DbmHostTop99p.Get()
}

// GetDbmHostTop99pOk returns a tuple with the DbmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetDbmHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmHostTop99p.Get(), o.DbmHostTop99p.IsSet()
}

// HasDbmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDbmHostTop99p() bool {
	return o != nil && o.DbmHostTop99p.IsSet()
}

// SetDbmHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the DbmHostTop99p field.
func (o *UsageSummaryDate) SetDbmHostTop99p(v int64) {
	o.DbmHostTop99p.Set(&v)
}

// SetDbmHostTop99pNil sets the value for DbmHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetDbmHostTop99pNil() {
	o.DbmHostTop99p.Set(nil)
}

// UnsetDbmHostTop99p ensures that no value is present for DbmHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetDbmHostTop99p() {
	o.DbmHostTop99p.Unset()
}

// GetDbmQueriesCountAvg returns the DbmQueriesCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetDbmQueriesCountAvg() int64 {
	if o == nil || o.DbmQueriesCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DbmQueriesCountAvg.Get()
}

// GetDbmQueriesCountAvgOk returns a tuple with the DbmQueriesCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetDbmQueriesCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmQueriesCountAvg.Get(), o.DbmQueriesCountAvg.IsSet()
}

// HasDbmQueriesCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDbmQueriesCountAvg() bool {
	return o != nil && o.DbmQueriesCountAvg.IsSet()
}

// SetDbmQueriesCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the DbmQueriesCountAvg field.
func (o *UsageSummaryDate) SetDbmQueriesCountAvg(v int64) {
	o.DbmQueriesCountAvg.Set(&v)
}

// SetDbmQueriesCountAvgNil sets the value for DbmQueriesCountAvg to be an explicit nil.
func (o *UsageSummaryDate) SetDbmQueriesCountAvgNil() {
	o.DbmQueriesCountAvg.Set(nil)
}

// UnsetDbmQueriesCountAvg ensures that no value is present for DbmQueriesCountAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetDbmQueriesCountAvg() {
	o.DbmQueriesCountAvg.Unset()
}

// GetFargateTasksCountAvg returns the FargateTasksCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetFargateTasksCountAvg() int64 {
	if o == nil || o.FargateTasksCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountAvg.Get()
}

// GetFargateTasksCountAvgOk returns a tuple with the FargateTasksCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetFargateTasksCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FargateTasksCountAvg.Get(), o.FargateTasksCountAvg.IsSet()
}

// HasFargateTasksCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFargateTasksCountAvg() bool {
	return o != nil && o.FargateTasksCountAvg.IsSet()
}

// SetFargateTasksCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the FargateTasksCountAvg field.
func (o *UsageSummaryDate) SetFargateTasksCountAvg(v int64) {
	o.FargateTasksCountAvg.Set(&v)
}

// SetFargateTasksCountAvgNil sets the value for FargateTasksCountAvg to be an explicit nil.
func (o *UsageSummaryDate) SetFargateTasksCountAvgNil() {
	o.FargateTasksCountAvg.Set(nil)
}

// UnsetFargateTasksCountAvg ensures that no value is present for FargateTasksCountAvg, not even an explicit nil.
func (o *UsageSummaryDate) UnsetFargateTasksCountAvg() {
	o.FargateTasksCountAvg.Unset()
}

// GetFargateTasksCountHwm returns the FargateTasksCountHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetFargateTasksCountHwm() int64 {
	if o == nil || o.FargateTasksCountHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountHwm.Get()
}

// GetFargateTasksCountHwmOk returns a tuple with the FargateTasksCountHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetFargateTasksCountHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FargateTasksCountHwm.Get(), o.FargateTasksCountHwm.IsSet()
}

// HasFargateTasksCountHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFargateTasksCountHwm() bool {
	return o != nil && o.FargateTasksCountHwm.IsSet()
}

// SetFargateTasksCountHwm gets a reference to the given datadog.NullableInt64 and assigns it to the FargateTasksCountHwm field.
func (o *UsageSummaryDate) SetFargateTasksCountHwm(v int64) {
	o.FargateTasksCountHwm.Set(&v)
}

// SetFargateTasksCountHwmNil sets the value for FargateTasksCountHwm to be an explicit nil.
func (o *UsageSummaryDate) SetFargateTasksCountHwmNil() {
	o.FargateTasksCountHwm.Set(nil)
}

// UnsetFargateTasksCountHwm ensures that no value is present for FargateTasksCountHwm, not even an explicit nil.
func (o *UsageSummaryDate) UnsetFargateTasksCountHwm() {
	o.FargateTasksCountHwm.Unset()
}

// GetForwardingEventsBytesSum returns the ForwardingEventsBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetForwardingEventsBytesSum() int64 {
	if o == nil || o.ForwardingEventsBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ForwardingEventsBytesSum.Get()
}

// GetForwardingEventsBytesSumOk returns a tuple with the ForwardingEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetForwardingEventsBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForwardingEventsBytesSum.Get(), o.ForwardingEventsBytesSum.IsSet()
}

// HasForwardingEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasForwardingEventsBytesSum() bool {
	return o != nil && o.ForwardingEventsBytesSum.IsSet()
}

// SetForwardingEventsBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the ForwardingEventsBytesSum field.
func (o *UsageSummaryDate) SetForwardingEventsBytesSum(v int64) {
	o.ForwardingEventsBytesSum.Set(&v)
}

// SetForwardingEventsBytesSumNil sets the value for ForwardingEventsBytesSum to be an explicit nil.
func (o *UsageSummaryDate) SetForwardingEventsBytesSumNil() {
	o.ForwardingEventsBytesSum.Set(nil)
}

// UnsetForwardingEventsBytesSum ensures that no value is present for ForwardingEventsBytesSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetForwardingEventsBytesSum() {
	o.ForwardingEventsBytesSum.Unset()
}

// GetGcpHostTop99p returns the GcpHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetGcpHostTop99p() int64 {
	if o == nil || o.GcpHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.GcpHostTop99p.Get()
}

// GetGcpHostTop99pOk returns a tuple with the GcpHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetGcpHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GcpHostTop99p.Get(), o.GcpHostTop99p.IsSet()
}

// HasGcpHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasGcpHostTop99p() bool {
	return o != nil && o.GcpHostTop99p.IsSet()
}

// SetGcpHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the GcpHostTop99p field.
func (o *UsageSummaryDate) SetGcpHostTop99p(v int64) {
	o.GcpHostTop99p.Set(&v)
}

// SetGcpHostTop99pNil sets the value for GcpHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetGcpHostTop99pNil() {
	o.GcpHostTop99p.Set(nil)
}

// UnsetGcpHostTop99p ensures that no value is present for GcpHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetGcpHostTop99p() {
	o.GcpHostTop99p.Unset()
}

// GetHerokuHostTop99p returns the HerokuHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetHerokuHostTop99p() int64 {
	if o == nil || o.HerokuHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HerokuHostTop99p.Get()
}

// GetHerokuHostTop99pOk returns a tuple with the HerokuHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetHerokuHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HerokuHostTop99p.Get(), o.HerokuHostTop99p.IsSet()
}

// HasHerokuHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasHerokuHostTop99p() bool {
	return o != nil && o.HerokuHostTop99p.IsSet()
}

// SetHerokuHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the HerokuHostTop99p field.
func (o *UsageSummaryDate) SetHerokuHostTop99p(v int64) {
	o.HerokuHostTop99p.Set(&v)
}

// SetHerokuHostTop99pNil sets the value for HerokuHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetHerokuHostTop99pNil() {
	o.HerokuHostTop99p.Set(nil)
}

// UnsetHerokuHostTop99p ensures that no value is present for HerokuHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetHerokuHostTop99p() {
	o.HerokuHostTop99p.Unset()
}

// GetIncidentManagementMonthlyActiveUsersHwm returns the IncidentManagementMonthlyActiveUsersHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetIncidentManagementMonthlyActiveUsersHwm() int64 {
	if o == nil || o.IncidentManagementMonthlyActiveUsersHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IncidentManagementMonthlyActiveUsersHwm.Get()
}

// GetIncidentManagementMonthlyActiveUsersHwmOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetIncidentManagementMonthlyActiveUsersHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentManagementMonthlyActiveUsersHwm.Get(), o.IncidentManagementMonthlyActiveUsersHwm.IsSet()
}

// HasIncidentManagementMonthlyActiveUsersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIncidentManagementMonthlyActiveUsersHwm() bool {
	return o != nil && o.IncidentManagementMonthlyActiveUsersHwm.IsSet()
}

// SetIncidentManagementMonthlyActiveUsersHwm gets a reference to the given datadog.NullableInt64 and assigns it to the IncidentManagementMonthlyActiveUsersHwm field.
func (o *UsageSummaryDate) SetIncidentManagementMonthlyActiveUsersHwm(v int64) {
	o.IncidentManagementMonthlyActiveUsersHwm.Set(&v)
}

// SetIncidentManagementMonthlyActiveUsersHwmNil sets the value for IncidentManagementMonthlyActiveUsersHwm to be an explicit nil.
func (o *UsageSummaryDate) SetIncidentManagementMonthlyActiveUsersHwmNil() {
	o.IncidentManagementMonthlyActiveUsersHwm.Set(nil)
}

// UnsetIncidentManagementMonthlyActiveUsersHwm ensures that no value is present for IncidentManagementMonthlyActiveUsersHwm, not even an explicit nil.
func (o *UsageSummaryDate) UnsetIncidentManagementMonthlyActiveUsersHwm() {
	o.IncidentManagementMonthlyActiveUsersHwm.Unset()
}

// GetIndexedEventsCountSum returns the IndexedEventsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetIndexedEventsCountSum() int64 {
	if o == nil || o.IndexedEventsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCountSum.Get()
}

// GetIndexedEventsCountSumOk returns a tuple with the IndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexedEventsCountSum.Get(), o.IndexedEventsCountSum.IsSet()
}

// HasIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIndexedEventsCountSum() bool {
	return o != nil && o.IndexedEventsCountSum.IsSet()
}

// SetIndexedEventsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the IndexedEventsCountSum field.
func (o *UsageSummaryDate) SetIndexedEventsCountSum(v int64) {
	o.IndexedEventsCountSum.Set(&v)
}

// SetIndexedEventsCountSumNil sets the value for IndexedEventsCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetIndexedEventsCountSumNil() {
	o.IndexedEventsCountSum.Set(nil)
}

// UnsetIndexedEventsCountSum ensures that no value is present for IndexedEventsCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetIndexedEventsCountSum() {
	o.IndexedEventsCountSum.Unset()
}

// GetInfraHostTop99p returns the InfraHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetInfraHostTop99p() int64 {
	if o == nil || o.InfraHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostTop99p.Get()
}

// GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetInfraHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfraHostTop99p.Get(), o.InfraHostTop99p.IsSet()
}

// HasInfraHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraHostTop99p() bool {
	return o != nil && o.InfraHostTop99p.IsSet()
}

// SetInfraHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the InfraHostTop99p field.
func (o *UsageSummaryDate) SetInfraHostTop99p(v int64) {
	o.InfraHostTop99p.Set(&v)
}

// SetInfraHostTop99pNil sets the value for InfraHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetInfraHostTop99pNil() {
	o.InfraHostTop99p.Set(nil)
}

// UnsetInfraHostTop99p ensures that no value is present for InfraHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetInfraHostTop99p() {
	o.InfraHostTop99p.Unset()
}

// GetIngestedEventsBytesSum returns the IngestedEventsBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetIngestedEventsBytesSum() int64 {
	if o == nil || o.IngestedEventsBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytesSum.Get()
}

// GetIngestedEventsBytesSumOk returns a tuple with the IngestedEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetIngestedEventsBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IngestedEventsBytesSum.Get(), o.IngestedEventsBytesSum.IsSet()
}

// HasIngestedEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIngestedEventsBytesSum() bool {
	return o != nil && o.IngestedEventsBytesSum.IsSet()
}

// SetIngestedEventsBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the IngestedEventsBytesSum field.
func (o *UsageSummaryDate) SetIngestedEventsBytesSum(v int64) {
	o.IngestedEventsBytesSum.Set(&v)
}

// SetIngestedEventsBytesSumNil sets the value for IngestedEventsBytesSum to be an explicit nil.
func (o *UsageSummaryDate) SetIngestedEventsBytesSumNil() {
	o.IngestedEventsBytesSum.Set(nil)
}

// UnsetIngestedEventsBytesSum ensures that no value is present for IngestedEventsBytesSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetIngestedEventsBytesSum() {
	o.IngestedEventsBytesSum.Unset()
}

// GetIotDeviceSum returns the IotDeviceSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetIotDeviceSum() int64 {
	if o == nil || o.IotDeviceSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceSum.Get()
}

// GetIotDeviceSumOk returns a tuple with the IotDeviceSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetIotDeviceSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IotDeviceSum.Get(), o.IotDeviceSum.IsSet()
}

// HasIotDeviceSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIotDeviceSum() bool {
	return o != nil && o.IotDeviceSum.IsSet()
}

// SetIotDeviceSum gets a reference to the given datadog.NullableInt64 and assigns it to the IotDeviceSum field.
func (o *UsageSummaryDate) SetIotDeviceSum(v int64) {
	o.IotDeviceSum.Set(&v)
}

// SetIotDeviceSumNil sets the value for IotDeviceSum to be an explicit nil.
func (o *UsageSummaryDate) SetIotDeviceSumNil() {
	o.IotDeviceSum.Set(nil)
}

// UnsetIotDeviceSum ensures that no value is present for IotDeviceSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetIotDeviceSum() {
	o.IotDeviceSum.Unset()
}

// GetIotDeviceTop99p returns the IotDeviceTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetIotDeviceTop99p() int64 {
	if o == nil || o.IotDeviceTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceTop99p.Get()
}

// GetIotDeviceTop99pOk returns a tuple with the IotDeviceTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetIotDeviceTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IotDeviceTop99p.Get(), o.IotDeviceTop99p.IsSet()
}

// HasIotDeviceTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIotDeviceTop99p() bool {
	return o != nil && o.IotDeviceTop99p.IsSet()
}

// SetIotDeviceTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the IotDeviceTop99p field.
func (o *UsageSummaryDate) SetIotDeviceTop99p(v int64) {
	o.IotDeviceTop99p.Set(&v)
}

// SetIotDeviceTop99pNil sets the value for IotDeviceTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetIotDeviceTop99pNil() {
	o.IotDeviceTop99p.Set(nil)
}

// UnsetIotDeviceTop99p ensures that no value is present for IotDeviceTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetIotDeviceTop99p() {
	o.IotDeviceTop99p.Unset()
}

// GetMobileRumLiteSessionCountSum returns the MobileRumLiteSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetMobileRumLiteSessionCountSum() int64 {
	if o == nil || o.MobileRumLiteSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumLiteSessionCountSum.Get()
}

// GetMobileRumLiteSessionCountSumOk returns a tuple with the MobileRumLiteSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetMobileRumLiteSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumLiteSessionCountSum.Get(), o.MobileRumLiteSessionCountSum.IsSet()
}

// HasMobileRumLiteSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumLiteSessionCountSum() bool {
	return o != nil && o.MobileRumLiteSessionCountSum.IsSet()
}

// SetMobileRumLiteSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumLiteSessionCountSum field.
func (o *UsageSummaryDate) SetMobileRumLiteSessionCountSum(v int64) {
	o.MobileRumLiteSessionCountSum.Set(&v)
}

// SetMobileRumLiteSessionCountSumNil sets the value for MobileRumLiteSessionCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetMobileRumLiteSessionCountSumNil() {
	o.MobileRumLiteSessionCountSum.Set(nil)
}

// UnsetMobileRumLiteSessionCountSum ensures that no value is present for MobileRumLiteSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetMobileRumLiteSessionCountSum() {
	o.MobileRumLiteSessionCountSum.Unset()
}

// GetMobileRumSessionCountAndroidSum returns the MobileRumSessionCountAndroidSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetMobileRumSessionCountAndroidSum() int64 {
	if o == nil || o.MobileRumSessionCountAndroidSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountAndroidSum.Get()
}

// GetMobileRumSessionCountAndroidSumOk returns a tuple with the MobileRumSessionCountAndroidSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetMobileRumSessionCountAndroidSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountAndroidSum.Get(), o.MobileRumSessionCountAndroidSum.IsSet()
}

// HasMobileRumSessionCountAndroidSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountAndroidSum() bool {
	return o != nil && o.MobileRumSessionCountAndroidSum.IsSet()
}

// SetMobileRumSessionCountAndroidSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountAndroidSum field.
func (o *UsageSummaryDate) SetMobileRumSessionCountAndroidSum(v int64) {
	o.MobileRumSessionCountAndroidSum.Set(&v)
}

// SetMobileRumSessionCountAndroidSumNil sets the value for MobileRumSessionCountAndroidSum to be an explicit nil.
func (o *UsageSummaryDate) SetMobileRumSessionCountAndroidSumNil() {
	o.MobileRumSessionCountAndroidSum.Set(nil)
}

// UnsetMobileRumSessionCountAndroidSum ensures that no value is present for MobileRumSessionCountAndroidSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetMobileRumSessionCountAndroidSum() {
	o.MobileRumSessionCountAndroidSum.Unset()
}

// GetMobileRumSessionCountFlutterSum returns the MobileRumSessionCountFlutterSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetMobileRumSessionCountFlutterSum() int64 {
	if o == nil || o.MobileRumSessionCountFlutterSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountFlutterSum.Get()
}

// GetMobileRumSessionCountFlutterSumOk returns a tuple with the MobileRumSessionCountFlutterSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetMobileRumSessionCountFlutterSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountFlutterSum.Get(), o.MobileRumSessionCountFlutterSum.IsSet()
}

// HasMobileRumSessionCountFlutterSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountFlutterSum() bool {
	return o != nil && o.MobileRumSessionCountFlutterSum.IsSet()
}

// SetMobileRumSessionCountFlutterSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountFlutterSum field.
func (o *UsageSummaryDate) SetMobileRumSessionCountFlutterSum(v int64) {
	o.MobileRumSessionCountFlutterSum.Set(&v)
}

// SetMobileRumSessionCountFlutterSumNil sets the value for MobileRumSessionCountFlutterSum to be an explicit nil.
func (o *UsageSummaryDate) SetMobileRumSessionCountFlutterSumNil() {
	o.MobileRumSessionCountFlutterSum.Set(nil)
}

// UnsetMobileRumSessionCountFlutterSum ensures that no value is present for MobileRumSessionCountFlutterSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetMobileRumSessionCountFlutterSum() {
	o.MobileRumSessionCountFlutterSum.Unset()
}

// GetMobileRumSessionCountIosSum returns the MobileRumSessionCountIosSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetMobileRumSessionCountIosSum() int64 {
	if o == nil || o.MobileRumSessionCountIosSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountIosSum.Get()
}

// GetMobileRumSessionCountIosSumOk returns a tuple with the MobileRumSessionCountIosSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetMobileRumSessionCountIosSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountIosSum.Get(), o.MobileRumSessionCountIosSum.IsSet()
}

// HasMobileRumSessionCountIosSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountIosSum() bool {
	return o != nil && o.MobileRumSessionCountIosSum.IsSet()
}

// SetMobileRumSessionCountIosSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountIosSum field.
func (o *UsageSummaryDate) SetMobileRumSessionCountIosSum(v int64) {
	o.MobileRumSessionCountIosSum.Set(&v)
}

// SetMobileRumSessionCountIosSumNil sets the value for MobileRumSessionCountIosSum to be an explicit nil.
func (o *UsageSummaryDate) SetMobileRumSessionCountIosSumNil() {
	o.MobileRumSessionCountIosSum.Set(nil)
}

// UnsetMobileRumSessionCountIosSum ensures that no value is present for MobileRumSessionCountIosSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetMobileRumSessionCountIosSum() {
	o.MobileRumSessionCountIosSum.Unset()
}

// GetMobileRumSessionCountReactnativeSum returns the MobileRumSessionCountReactnativeSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetMobileRumSessionCountReactnativeSum() int64 {
	if o == nil || o.MobileRumSessionCountReactnativeSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountReactnativeSum.Get()
}

// GetMobileRumSessionCountReactnativeSumOk returns a tuple with the MobileRumSessionCountReactnativeSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetMobileRumSessionCountReactnativeSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountReactnativeSum.Get(), o.MobileRumSessionCountReactnativeSum.IsSet()
}

// HasMobileRumSessionCountReactnativeSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountReactnativeSum() bool {
	return o != nil && o.MobileRumSessionCountReactnativeSum.IsSet()
}

// SetMobileRumSessionCountReactnativeSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountReactnativeSum field.
func (o *UsageSummaryDate) SetMobileRumSessionCountReactnativeSum(v int64) {
	o.MobileRumSessionCountReactnativeSum.Set(&v)
}

// SetMobileRumSessionCountReactnativeSumNil sets the value for MobileRumSessionCountReactnativeSum to be an explicit nil.
func (o *UsageSummaryDate) SetMobileRumSessionCountReactnativeSumNil() {
	o.MobileRumSessionCountReactnativeSum.Set(nil)
}

// UnsetMobileRumSessionCountReactnativeSum ensures that no value is present for MobileRumSessionCountReactnativeSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetMobileRumSessionCountReactnativeSum() {
	o.MobileRumSessionCountReactnativeSum.Unset()
}

// GetMobileRumSessionCountSum returns the MobileRumSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetMobileRumSessionCountSum() int64 {
	if o == nil || o.MobileRumSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountSum.Get()
}

// GetMobileRumSessionCountSumOk returns a tuple with the MobileRumSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetMobileRumSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountSum.Get(), o.MobileRumSessionCountSum.IsSet()
}

// HasMobileRumSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountSum() bool {
	return o != nil && o.MobileRumSessionCountSum.IsSet()
}

// SetMobileRumSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountSum field.
func (o *UsageSummaryDate) SetMobileRumSessionCountSum(v int64) {
	o.MobileRumSessionCountSum.Set(&v)
}

// SetMobileRumSessionCountSumNil sets the value for MobileRumSessionCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetMobileRumSessionCountSumNil() {
	o.MobileRumSessionCountSum.Set(nil)
}

// UnsetMobileRumSessionCountSum ensures that no value is present for MobileRumSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetMobileRumSessionCountSum() {
	o.MobileRumSessionCountSum.Unset()
}

// GetMobileRumUnitsSum returns the MobileRumUnitsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetMobileRumUnitsSum() int64 {
	if o == nil || o.MobileRumUnitsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumUnitsSum.Get()
}

// GetMobileRumUnitsSumOk returns a tuple with the MobileRumUnitsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetMobileRumUnitsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumUnitsSum.Get(), o.MobileRumUnitsSum.IsSet()
}

// HasMobileRumUnitsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumUnitsSum() bool {
	return o != nil && o.MobileRumUnitsSum.IsSet()
}

// SetMobileRumUnitsSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumUnitsSum field.
func (o *UsageSummaryDate) SetMobileRumUnitsSum(v int64) {
	o.MobileRumUnitsSum.Set(&v)
}

// SetMobileRumUnitsSumNil sets the value for MobileRumUnitsSum to be an explicit nil.
func (o *UsageSummaryDate) SetMobileRumUnitsSumNil() {
	o.MobileRumUnitsSum.Set(nil)
}

// UnsetMobileRumUnitsSum ensures that no value is present for MobileRumUnitsSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetMobileRumUnitsSum() {
	o.MobileRumUnitsSum.Unset()
}

// GetNetflowIndexedEventsCountSum returns the NetflowIndexedEventsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSum() int64 {
	if o == nil || o.NetflowIndexedEventsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NetflowIndexedEventsCountSum.Get()
}

// GetNetflowIndexedEventsCountSumOk returns a tuple with the NetflowIndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetflowIndexedEventsCountSum.Get(), o.NetflowIndexedEventsCountSum.IsSet()
}

// HasNetflowIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasNetflowIndexedEventsCountSum() bool {
	return o != nil && o.NetflowIndexedEventsCountSum.IsSet()
}

// SetNetflowIndexedEventsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the NetflowIndexedEventsCountSum field.
func (o *UsageSummaryDate) SetNetflowIndexedEventsCountSum(v int64) {
	o.NetflowIndexedEventsCountSum.Set(&v)
}

// SetNetflowIndexedEventsCountSumNil sets the value for NetflowIndexedEventsCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetNetflowIndexedEventsCountSumNil() {
	o.NetflowIndexedEventsCountSum.Set(nil)
}

// UnsetNetflowIndexedEventsCountSum ensures that no value is present for NetflowIndexedEventsCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetNetflowIndexedEventsCountSum() {
	o.NetflowIndexedEventsCountSum.Unset()
}

// GetNpmHostTop99p returns the NpmHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetNpmHostTop99p() int64 {
	if o == nil || o.NpmHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NpmHostTop99p.Get()
}

// GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetNpmHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NpmHostTop99p.Get(), o.NpmHostTop99p.IsSet()
}

// HasNpmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasNpmHostTop99p() bool {
	return o != nil && o.NpmHostTop99p.IsSet()
}

// SetNpmHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the NpmHostTop99p field.
func (o *UsageSummaryDate) SetNpmHostTop99p(v int64) {
	o.NpmHostTop99p.Set(&v)
}

// SetNpmHostTop99pNil sets the value for NpmHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetNpmHostTop99pNil() {
	o.NpmHostTop99p.Set(nil)
}

// UnsetNpmHostTop99p ensures that no value is present for NpmHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetNpmHostTop99p() {
	o.NpmHostTop99p.Unset()
}

// GetObservabilityPipelinesBytesProcessedSum returns the ObservabilityPipelinesBytesProcessedSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetObservabilityPipelinesBytesProcessedSum() int64 {
	if o == nil || o.ObservabilityPipelinesBytesProcessedSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ObservabilityPipelinesBytesProcessedSum.Get()
}

// GetObservabilityPipelinesBytesProcessedSumOk returns a tuple with the ObservabilityPipelinesBytesProcessedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetObservabilityPipelinesBytesProcessedSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObservabilityPipelinesBytesProcessedSum.Get(), o.ObservabilityPipelinesBytesProcessedSum.IsSet()
}

// HasObservabilityPipelinesBytesProcessedSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasObservabilityPipelinesBytesProcessedSum() bool {
	return o != nil && o.ObservabilityPipelinesBytesProcessedSum.IsSet()
}

// SetObservabilityPipelinesBytesProcessedSum gets a reference to the given datadog.NullableInt64 and assigns it to the ObservabilityPipelinesBytesProcessedSum field.
func (o *UsageSummaryDate) SetObservabilityPipelinesBytesProcessedSum(v int64) {
	o.ObservabilityPipelinesBytesProcessedSum.Set(&v)
}

// SetObservabilityPipelinesBytesProcessedSumNil sets the value for ObservabilityPipelinesBytesProcessedSum to be an explicit nil.
func (o *UsageSummaryDate) SetObservabilityPipelinesBytesProcessedSumNil() {
	o.ObservabilityPipelinesBytesProcessedSum.Set(nil)
}

// UnsetObservabilityPipelinesBytesProcessedSum ensures that no value is present for ObservabilityPipelinesBytesProcessedSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetObservabilityPipelinesBytesProcessedSum() {
	o.ObservabilityPipelinesBytesProcessedSum.Unset()
}

// GetOnlineArchiveEventsCountSum returns the OnlineArchiveEventsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetOnlineArchiveEventsCountSum() int64 {
	if o == nil || o.OnlineArchiveEventsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OnlineArchiveEventsCountSum.Get()
}

// GetOnlineArchiveEventsCountSumOk returns a tuple with the OnlineArchiveEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetOnlineArchiveEventsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlineArchiveEventsCountSum.Get(), o.OnlineArchiveEventsCountSum.IsSet()
}

// HasOnlineArchiveEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOnlineArchiveEventsCountSum() bool {
	return o != nil && o.OnlineArchiveEventsCountSum.IsSet()
}

// SetOnlineArchiveEventsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the OnlineArchiveEventsCountSum field.
func (o *UsageSummaryDate) SetOnlineArchiveEventsCountSum(v int64) {
	o.OnlineArchiveEventsCountSum.Set(&v)
}

// SetOnlineArchiveEventsCountSumNil sets the value for OnlineArchiveEventsCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetOnlineArchiveEventsCountSumNil() {
	o.OnlineArchiveEventsCountSum.Set(nil)
}

// UnsetOnlineArchiveEventsCountSum ensures that no value is present for OnlineArchiveEventsCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetOnlineArchiveEventsCountSum() {
	o.OnlineArchiveEventsCountSum.Unset()
}

// GetOpentelemetryApmHostTop99p returns the OpentelemetryApmHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetOpentelemetryApmHostTop99p() int64 {
	if o == nil || o.OpentelemetryApmHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryApmHostTop99p.Get()
}

// GetOpentelemetryApmHostTop99pOk returns a tuple with the OpentelemetryApmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetOpentelemetryApmHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpentelemetryApmHostTop99p.Get(), o.OpentelemetryApmHostTop99p.IsSet()
}

// HasOpentelemetryApmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOpentelemetryApmHostTop99p() bool {
	return o != nil && o.OpentelemetryApmHostTop99p.IsSet()
}

// SetOpentelemetryApmHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the OpentelemetryApmHostTop99p field.
func (o *UsageSummaryDate) SetOpentelemetryApmHostTop99p(v int64) {
	o.OpentelemetryApmHostTop99p.Set(&v)
}

// SetOpentelemetryApmHostTop99pNil sets the value for OpentelemetryApmHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetOpentelemetryApmHostTop99pNil() {
	o.OpentelemetryApmHostTop99p.Set(nil)
}

// UnsetOpentelemetryApmHostTop99p ensures that no value is present for OpentelemetryApmHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetOpentelemetryApmHostTop99p() {
	o.OpentelemetryApmHostTop99p.Unset()
}

// GetOpentelemetryHostTop99p returns the OpentelemetryHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetOpentelemetryHostTop99p() int64 {
	if o == nil || o.OpentelemetryHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryHostTop99p.Get()
}

// GetOpentelemetryHostTop99pOk returns a tuple with the OpentelemetryHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetOpentelemetryHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpentelemetryHostTop99p.Get(), o.OpentelemetryHostTop99p.IsSet()
}

// HasOpentelemetryHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOpentelemetryHostTop99p() bool {
	return o != nil && o.OpentelemetryHostTop99p.IsSet()
}

// SetOpentelemetryHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the OpentelemetryHostTop99p field.
func (o *UsageSummaryDate) SetOpentelemetryHostTop99p(v int64) {
	o.OpentelemetryHostTop99p.Set(&v)
}

// SetOpentelemetryHostTop99pNil sets the value for OpentelemetryHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetOpentelemetryHostTop99pNil() {
	o.OpentelemetryHostTop99p.Set(nil)
}

// UnsetOpentelemetryHostTop99p ensures that no value is present for OpentelemetryHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetOpentelemetryHostTop99p() {
	o.OpentelemetryHostTop99p.Unset()
}

// GetOrgs returns the Orgs field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetOrgs() []UsageSummaryDateOrg {
	if o == nil || o.Orgs == nil {
		var ret []UsageSummaryDateOrg
		return ret
	}
	return o.Orgs
}

// GetOrgsOk returns a tuple with the Orgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetOrgsOk() (*[]UsageSummaryDateOrg, bool) {
	if o == nil || o.Orgs == nil {
		return nil, false
	}
	return &o.Orgs, true
}

// HasOrgs returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOrgs() bool {
	return o != nil && o.Orgs != nil
}

// SetOrgs gets a reference to the given []UsageSummaryDateOrg and assigns it to the Orgs field.
func (o *UsageSummaryDate) SetOrgs(v []UsageSummaryDateOrg) {
	o.Orgs = v
}

// GetProfilingHostTop99p returns the ProfilingHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetProfilingHostTop99p() int64 {
	if o == nil || o.ProfilingHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ProfilingHostTop99p.Get()
}

// GetProfilingHostTop99pOk returns a tuple with the ProfilingHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetProfilingHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilingHostTop99p.Get(), o.ProfilingHostTop99p.IsSet()
}

// HasProfilingHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasProfilingHostTop99p() bool {
	return o != nil && o.ProfilingHostTop99p.IsSet()
}

// SetProfilingHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the ProfilingHostTop99p field.
func (o *UsageSummaryDate) SetProfilingHostTop99p(v int64) {
	o.ProfilingHostTop99p.Set(&v)
}

// SetProfilingHostTop99pNil sets the value for ProfilingHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetProfilingHostTop99pNil() {
	o.ProfilingHostTop99p.Set(nil)
}

// UnsetProfilingHostTop99p ensures that no value is present for ProfilingHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetProfilingHostTop99p() {
	o.ProfilingHostTop99p.Unset()
}

// GetRumBrowserAndMobileSessionCount returns the RumBrowserAndMobileSessionCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetRumBrowserAndMobileSessionCount() int64 {
	if o == nil || o.RumBrowserAndMobileSessionCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserAndMobileSessionCount.Get()
}

// GetRumBrowserAndMobileSessionCountOk returns a tuple with the RumBrowserAndMobileSessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetRumBrowserAndMobileSessionCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumBrowserAndMobileSessionCount.Get(), o.RumBrowserAndMobileSessionCount.IsSet()
}

// HasRumBrowserAndMobileSessionCount returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumBrowserAndMobileSessionCount() bool {
	return o != nil && o.RumBrowserAndMobileSessionCount.IsSet()
}

// SetRumBrowserAndMobileSessionCount gets a reference to the given datadog.NullableInt64 and assigns it to the RumBrowserAndMobileSessionCount field.
func (o *UsageSummaryDate) SetRumBrowserAndMobileSessionCount(v int64) {
	o.RumBrowserAndMobileSessionCount.Set(&v)
}

// SetRumBrowserAndMobileSessionCountNil sets the value for RumBrowserAndMobileSessionCount to be an explicit nil.
func (o *UsageSummaryDate) SetRumBrowserAndMobileSessionCountNil() {
	o.RumBrowserAndMobileSessionCount.Set(nil)
}

// UnsetRumBrowserAndMobileSessionCount ensures that no value is present for RumBrowserAndMobileSessionCount, not even an explicit nil.
func (o *UsageSummaryDate) UnsetRumBrowserAndMobileSessionCount() {
	o.RumBrowserAndMobileSessionCount.Unset()
}

// GetRumSessionCountSum returns the RumSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetRumSessionCountSum() int64 {
	if o == nil || o.RumSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumSessionCountSum.Get()
}

// GetRumSessionCountSumOk returns a tuple with the RumSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetRumSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumSessionCountSum.Get(), o.RumSessionCountSum.IsSet()
}

// HasRumSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumSessionCountSum() bool {
	return o != nil && o.RumSessionCountSum.IsSet()
}

// SetRumSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the RumSessionCountSum field.
func (o *UsageSummaryDate) SetRumSessionCountSum(v int64) {
	o.RumSessionCountSum.Set(&v)
}

// SetRumSessionCountSumNil sets the value for RumSessionCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetRumSessionCountSumNil() {
	o.RumSessionCountSum.Set(nil)
}

// UnsetRumSessionCountSum ensures that no value is present for RumSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetRumSessionCountSum() {
	o.RumSessionCountSum.Unset()
}

// GetRumTotalSessionCountSum returns the RumTotalSessionCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetRumTotalSessionCountSum() int64 {
	if o == nil || o.RumTotalSessionCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumTotalSessionCountSum.Get()
}

// GetRumTotalSessionCountSumOk returns a tuple with the RumTotalSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetRumTotalSessionCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumTotalSessionCountSum.Get(), o.RumTotalSessionCountSum.IsSet()
}

// HasRumTotalSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumTotalSessionCountSum() bool {
	return o != nil && o.RumTotalSessionCountSum.IsSet()
}

// SetRumTotalSessionCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the RumTotalSessionCountSum field.
func (o *UsageSummaryDate) SetRumTotalSessionCountSum(v int64) {
	o.RumTotalSessionCountSum.Set(&v)
}

// SetRumTotalSessionCountSumNil sets the value for RumTotalSessionCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetRumTotalSessionCountSumNil() {
	o.RumTotalSessionCountSum.Set(nil)
}

// UnsetRumTotalSessionCountSum ensures that no value is present for RumTotalSessionCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetRumTotalSessionCountSum() {
	o.RumTotalSessionCountSum.Unset()
}

// GetRumUnitsSum returns the RumUnitsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetRumUnitsSum() int64 {
	if o == nil || o.RumUnitsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumUnitsSum.Get()
}

// GetRumUnitsSumOk returns a tuple with the RumUnitsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetRumUnitsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumUnitsSum.Get(), o.RumUnitsSum.IsSet()
}

// HasRumUnitsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumUnitsSum() bool {
	return o != nil && o.RumUnitsSum.IsSet()
}

// SetRumUnitsSum gets a reference to the given datadog.NullableInt64 and assigns it to the RumUnitsSum field.
func (o *UsageSummaryDate) SetRumUnitsSum(v int64) {
	o.RumUnitsSum.Set(&v)
}

// SetRumUnitsSumNil sets the value for RumUnitsSum to be an explicit nil.
func (o *UsageSummaryDate) SetRumUnitsSumNil() {
	o.RumUnitsSum.Set(nil)
}

// UnsetRumUnitsSum ensures that no value is present for RumUnitsSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetRumUnitsSum() {
	o.RumUnitsSum.Unset()
}

// GetSdsApmScannedBytesSum returns the SdsApmScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetSdsApmScannedBytesSum() int64 {
	if o == nil || o.SdsApmScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsApmScannedBytesSum.Get()
}

// GetSdsApmScannedBytesSumOk returns a tuple with the SdsApmScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetSdsApmScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsApmScannedBytesSum.Get(), o.SdsApmScannedBytesSum.IsSet()
}

// HasSdsApmScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsApmScannedBytesSum() bool {
	return o != nil && o.SdsApmScannedBytesSum.IsSet()
}

// SetSdsApmScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsApmScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsApmScannedBytesSum(v int64) {
	o.SdsApmScannedBytesSum.Set(&v)
}

// SetSdsApmScannedBytesSumNil sets the value for SdsApmScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDate) SetSdsApmScannedBytesSumNil() {
	o.SdsApmScannedBytesSum.Set(nil)
}

// UnsetSdsApmScannedBytesSum ensures that no value is present for SdsApmScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetSdsApmScannedBytesSum() {
	o.SdsApmScannedBytesSum.Unset()
}

// GetSdsEventsScannedBytesSum returns the SdsEventsScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetSdsEventsScannedBytesSum() int64 {
	if o == nil || o.SdsEventsScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsEventsScannedBytesSum.Get()
}

// GetSdsEventsScannedBytesSumOk returns a tuple with the SdsEventsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetSdsEventsScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsEventsScannedBytesSum.Get(), o.SdsEventsScannedBytesSum.IsSet()
}

// HasSdsEventsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsEventsScannedBytesSum() bool {
	return o != nil && o.SdsEventsScannedBytesSum.IsSet()
}

// SetSdsEventsScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsEventsScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsEventsScannedBytesSum(v int64) {
	o.SdsEventsScannedBytesSum.Set(&v)
}

// SetSdsEventsScannedBytesSumNil sets the value for SdsEventsScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDate) SetSdsEventsScannedBytesSumNil() {
	o.SdsEventsScannedBytesSum.Set(nil)
}

// UnsetSdsEventsScannedBytesSum ensures that no value is present for SdsEventsScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetSdsEventsScannedBytesSum() {
	o.SdsEventsScannedBytesSum.Unset()
}

// GetSdsLogsScannedBytesSum returns the SdsLogsScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetSdsLogsScannedBytesSum() int64 {
	if o == nil || o.SdsLogsScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsLogsScannedBytesSum.Get()
}

// GetSdsLogsScannedBytesSumOk returns a tuple with the SdsLogsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetSdsLogsScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsLogsScannedBytesSum.Get(), o.SdsLogsScannedBytesSum.IsSet()
}

// HasSdsLogsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsLogsScannedBytesSum() bool {
	return o != nil && o.SdsLogsScannedBytesSum.IsSet()
}

// SetSdsLogsScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsLogsScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsLogsScannedBytesSum(v int64) {
	o.SdsLogsScannedBytesSum.Set(&v)
}

// SetSdsLogsScannedBytesSumNil sets the value for SdsLogsScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDate) SetSdsLogsScannedBytesSumNil() {
	o.SdsLogsScannedBytesSum.Set(nil)
}

// UnsetSdsLogsScannedBytesSum ensures that no value is present for SdsLogsScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetSdsLogsScannedBytesSum() {
	o.SdsLogsScannedBytesSum.Unset()
}

// GetSdsRumScannedBytesSum returns the SdsRumScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetSdsRumScannedBytesSum() int64 {
	if o == nil || o.SdsRumScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsRumScannedBytesSum.Get()
}

// GetSdsRumScannedBytesSumOk returns a tuple with the SdsRumScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetSdsRumScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsRumScannedBytesSum.Get(), o.SdsRumScannedBytesSum.IsSet()
}

// HasSdsRumScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsRumScannedBytesSum() bool {
	return o != nil && o.SdsRumScannedBytesSum.IsSet()
}

// SetSdsRumScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsRumScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsRumScannedBytesSum(v int64) {
	o.SdsRumScannedBytesSum.Set(&v)
}

// SetSdsRumScannedBytesSumNil sets the value for SdsRumScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDate) SetSdsRumScannedBytesSumNil() {
	o.SdsRumScannedBytesSum.Set(nil)
}

// UnsetSdsRumScannedBytesSum ensures that no value is present for SdsRumScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetSdsRumScannedBytesSum() {
	o.SdsRumScannedBytesSum.Unset()
}

// GetSdsTotalScannedBytesSum returns the SdsTotalScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetSdsTotalScannedBytesSum() int64 {
	if o == nil || o.SdsTotalScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsTotalScannedBytesSum.Get()
}

// GetSdsTotalScannedBytesSumOk returns a tuple with the SdsTotalScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetSdsTotalScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsTotalScannedBytesSum.Get(), o.SdsTotalScannedBytesSum.IsSet()
}

// HasSdsTotalScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsTotalScannedBytesSum() bool {
	return o != nil && o.SdsTotalScannedBytesSum.IsSet()
}

// SetSdsTotalScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsTotalScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsTotalScannedBytesSum(v int64) {
	o.SdsTotalScannedBytesSum.Set(&v)
}

// SetSdsTotalScannedBytesSumNil sets the value for SdsTotalScannedBytesSum to be an explicit nil.
func (o *UsageSummaryDate) SetSdsTotalScannedBytesSumNil() {
	o.SdsTotalScannedBytesSum.Set(nil)
}

// UnsetSdsTotalScannedBytesSum ensures that no value is present for SdsTotalScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetSdsTotalScannedBytesSum() {
	o.SdsTotalScannedBytesSum.Unset()
}

// GetSyntheticsBrowserCheckCallsCountSum returns the SyntheticsBrowserCheckCallsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSum() int64 {
	if o == nil || o.SyntheticsBrowserCheckCallsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsBrowserCheckCallsCountSum.Get()
}

// GetSyntheticsBrowserCheckCallsCountSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyntheticsBrowserCheckCallsCountSum.Get(), o.SyntheticsBrowserCheckCallsCountSum.IsSet()
}

// HasSyntheticsBrowserCheckCallsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSyntheticsBrowserCheckCallsCountSum() bool {
	return o != nil && o.SyntheticsBrowserCheckCallsCountSum.IsSet()
}

// SetSyntheticsBrowserCheckCallsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the SyntheticsBrowserCheckCallsCountSum field.
func (o *UsageSummaryDate) SetSyntheticsBrowserCheckCallsCountSum(v int64) {
	o.SyntheticsBrowserCheckCallsCountSum.Set(&v)
}

// SetSyntheticsBrowserCheckCallsCountSumNil sets the value for SyntheticsBrowserCheckCallsCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetSyntheticsBrowserCheckCallsCountSumNil() {
	o.SyntheticsBrowserCheckCallsCountSum.Set(nil)
}

// UnsetSyntheticsBrowserCheckCallsCountSum ensures that no value is present for SyntheticsBrowserCheckCallsCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetSyntheticsBrowserCheckCallsCountSum() {
	o.SyntheticsBrowserCheckCallsCountSum.Unset()
}

// GetSyntheticsCheckCallsCountSum returns the SyntheticsCheckCallsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSum() int64 {
	if o == nil || o.SyntheticsCheckCallsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsCheckCallsCountSum.Get()
}

// GetSyntheticsCheckCallsCountSumOk returns a tuple with the SyntheticsCheckCallsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyntheticsCheckCallsCountSum.Get(), o.SyntheticsCheckCallsCountSum.IsSet()
}

// HasSyntheticsCheckCallsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSyntheticsCheckCallsCountSum() bool {
	return o != nil && o.SyntheticsCheckCallsCountSum.IsSet()
}

// SetSyntheticsCheckCallsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the SyntheticsCheckCallsCountSum field.
func (o *UsageSummaryDate) SetSyntheticsCheckCallsCountSum(v int64) {
	o.SyntheticsCheckCallsCountSum.Set(&v)
}

// SetSyntheticsCheckCallsCountSumNil sets the value for SyntheticsCheckCallsCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetSyntheticsCheckCallsCountSumNil() {
	o.SyntheticsCheckCallsCountSum.Set(nil)
}

// UnsetSyntheticsCheckCallsCountSum ensures that no value is present for SyntheticsCheckCallsCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetSyntheticsCheckCallsCountSum() {
	o.SyntheticsCheckCallsCountSum.Unset()
}

// GetSyntheticsParallelTestingMaxSlotsHwm returns the SyntheticsParallelTestingMaxSlotsHwm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetSyntheticsParallelTestingMaxSlotsHwm() int64 {
	if o == nil || o.SyntheticsParallelTestingMaxSlotsHwm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsParallelTestingMaxSlotsHwm.Get()
}

// GetSyntheticsParallelTestingMaxSlotsHwmOk returns a tuple with the SyntheticsParallelTestingMaxSlotsHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetSyntheticsParallelTestingMaxSlotsHwmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyntheticsParallelTestingMaxSlotsHwm.Get(), o.SyntheticsParallelTestingMaxSlotsHwm.IsSet()
}

// HasSyntheticsParallelTestingMaxSlotsHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSyntheticsParallelTestingMaxSlotsHwm() bool {
	return o != nil && o.SyntheticsParallelTestingMaxSlotsHwm.IsSet()
}

// SetSyntheticsParallelTestingMaxSlotsHwm gets a reference to the given datadog.NullableInt64 and assigns it to the SyntheticsParallelTestingMaxSlotsHwm field.
func (o *UsageSummaryDate) SetSyntheticsParallelTestingMaxSlotsHwm(v int64) {
	o.SyntheticsParallelTestingMaxSlotsHwm.Set(&v)
}

// SetSyntheticsParallelTestingMaxSlotsHwmNil sets the value for SyntheticsParallelTestingMaxSlotsHwm to be an explicit nil.
func (o *UsageSummaryDate) SetSyntheticsParallelTestingMaxSlotsHwmNil() {
	o.SyntheticsParallelTestingMaxSlotsHwm.Set(nil)
}

// UnsetSyntheticsParallelTestingMaxSlotsHwm ensures that no value is present for SyntheticsParallelTestingMaxSlotsHwm, not even an explicit nil.
func (o *UsageSummaryDate) UnsetSyntheticsParallelTestingMaxSlotsHwm() {
	o.SyntheticsParallelTestingMaxSlotsHwm.Unset()
}

// GetTraceSearchIndexedEventsCountSum returns the TraceSearchIndexedEventsCountSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSum() int64 {
	if o == nil || o.TraceSearchIndexedEventsCountSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TraceSearchIndexedEventsCountSum.Get()
}

// GetTraceSearchIndexedEventsCountSumOk returns a tuple with the TraceSearchIndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceSearchIndexedEventsCountSum.Get(), o.TraceSearchIndexedEventsCountSum.IsSet()
}

// HasTraceSearchIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasTraceSearchIndexedEventsCountSum() bool {
	return o != nil && o.TraceSearchIndexedEventsCountSum.IsSet()
}

// SetTraceSearchIndexedEventsCountSum gets a reference to the given datadog.NullableInt64 and assigns it to the TraceSearchIndexedEventsCountSum field.
func (o *UsageSummaryDate) SetTraceSearchIndexedEventsCountSum(v int64) {
	o.TraceSearchIndexedEventsCountSum.Set(&v)
}

// SetTraceSearchIndexedEventsCountSumNil sets the value for TraceSearchIndexedEventsCountSum to be an explicit nil.
func (o *UsageSummaryDate) SetTraceSearchIndexedEventsCountSumNil() {
	o.TraceSearchIndexedEventsCountSum.Set(nil)
}

// UnsetTraceSearchIndexedEventsCountSum ensures that no value is present for TraceSearchIndexedEventsCountSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetTraceSearchIndexedEventsCountSum() {
	o.TraceSearchIndexedEventsCountSum.Unset()
}

// GetTwolIngestedEventsBytesSum returns the TwolIngestedEventsBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetTwolIngestedEventsBytesSum() int64 {
	if o == nil || o.TwolIngestedEventsBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TwolIngestedEventsBytesSum.Get()
}

// GetTwolIngestedEventsBytesSumOk returns a tuple with the TwolIngestedEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetTwolIngestedEventsBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwolIngestedEventsBytesSum.Get(), o.TwolIngestedEventsBytesSum.IsSet()
}

// HasTwolIngestedEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasTwolIngestedEventsBytesSum() bool {
	return o != nil && o.TwolIngestedEventsBytesSum.IsSet()
}

// SetTwolIngestedEventsBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the TwolIngestedEventsBytesSum field.
func (o *UsageSummaryDate) SetTwolIngestedEventsBytesSum(v int64) {
	o.TwolIngestedEventsBytesSum.Set(&v)
}

// SetTwolIngestedEventsBytesSumNil sets the value for TwolIngestedEventsBytesSum to be an explicit nil.
func (o *UsageSummaryDate) SetTwolIngestedEventsBytesSumNil() {
	o.TwolIngestedEventsBytesSum.Set(nil)
}

// UnsetTwolIngestedEventsBytesSum ensures that no value is present for TwolIngestedEventsBytesSum, not even an explicit nil.
func (o *UsageSummaryDate) UnsetTwolIngestedEventsBytesSum() {
	o.TwolIngestedEventsBytesSum.Unset()
}

// GetUniversalServiceMonitoringHostTop99p returns the UniversalServiceMonitoringHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetUniversalServiceMonitoringHostTop99p() int64 {
	if o == nil || o.UniversalServiceMonitoringHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UniversalServiceMonitoringHostTop99p.Get()
}

// GetUniversalServiceMonitoringHostTop99pOk returns a tuple with the UniversalServiceMonitoringHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetUniversalServiceMonitoringHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringHostTop99p.Get(), o.UniversalServiceMonitoringHostTop99p.IsSet()
}

// HasUniversalServiceMonitoringHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasUniversalServiceMonitoringHostTop99p() bool {
	return o != nil && o.UniversalServiceMonitoringHostTop99p.IsSet()
}

// SetUniversalServiceMonitoringHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the UniversalServiceMonitoringHostTop99p field.
func (o *UsageSummaryDate) SetUniversalServiceMonitoringHostTop99p(v int64) {
	o.UniversalServiceMonitoringHostTop99p.Set(&v)
}

// SetUniversalServiceMonitoringHostTop99pNil sets the value for UniversalServiceMonitoringHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetUniversalServiceMonitoringHostTop99pNil() {
	o.UniversalServiceMonitoringHostTop99p.Set(nil)
}

// UnsetUniversalServiceMonitoringHostTop99p ensures that no value is present for UniversalServiceMonitoringHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetUniversalServiceMonitoringHostTop99p() {
	o.UniversalServiceMonitoringHostTop99p.Unset()
}

// GetVsphereHostTop99p returns the VsphereHostTop99p field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryDate) GetVsphereHostTop99p() int64 {
	if o == nil || o.VsphereHostTop99p.Get() == nil {
		var ret int64
		return ret
	}
	return *o.VsphereHostTop99p.Get()
}

// GetVsphereHostTop99pOk returns a tuple with the VsphereHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryDate) GetVsphereHostTop99pOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VsphereHostTop99p.Get(), o.VsphereHostTop99p.IsSet()
}

// HasVsphereHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasVsphereHostTop99p() bool {
	return o != nil && o.VsphereHostTop99p.IsSet()
}

// SetVsphereHostTop99p gets a reference to the given datadog.NullableInt64 and assigns it to the VsphereHostTop99p field.
func (o *UsageSummaryDate) SetVsphereHostTop99p(v int64) {
	o.VsphereHostTop99p.Set(&v)
}

// SetVsphereHostTop99pNil sets the value for VsphereHostTop99p to be an explicit nil.
func (o *UsageSummaryDate) SetVsphereHostTop99pNil() {
	o.VsphereHostTop99p.Set(nil)
}

// UnsetVsphereHostTop99p ensures that no value is present for VsphereHostTop99p, not even an explicit nil.
func (o *UsageSummaryDate) UnsetVsphereHostTop99p() {
	o.VsphereHostTop99p.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSummaryDate) MarshalJSON() ([]byte, error) {
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
	if o.AwsLambdaFuncCount.IsSet() {
		toSerialize["aws_lambda_func_count"] = o.AwsLambdaFuncCount.Get()
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
	if o.CloudCostManagementHostCountAvg.IsSet() {
		toSerialize["cloud_cost_management_host_count_avg"] = o.CloudCostManagementHostCountAvg.Get()
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
	if o.CspmAwsHostTop99p.IsSet() {
		toSerialize["cspm_aws_host_top99p"] = o.CspmAwsHostTop99p.Get()
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
	if o.Date != nil {
		if o.Date.Nanosecond() == 0 {
			toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DbmHostTop99p.IsSet() {
		toSerialize["dbm_host_top99p"] = o.DbmHostTop99p.Get()
	}
	if o.DbmQueriesCountAvg.IsSet() {
		toSerialize["dbm_queries_count_avg"] = o.DbmQueriesCountAvg.Get()
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
	if o.IotDeviceSum.IsSet() {
		toSerialize["iot_device_sum"] = o.IotDeviceSum.Get()
	}
	if o.IotDeviceTop99p.IsSet() {
		toSerialize["iot_device_top99p"] = o.IotDeviceTop99p.Get()
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
	if o.Orgs != nil {
		toSerialize["orgs"] = o.Orgs
	}
	if o.ProfilingHostTop99p.IsSet() {
		toSerialize["profiling_host_top99p"] = o.ProfilingHostTop99p.Get()
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
func (o *UsageSummaryDate) UnmarshalJSON(bytes []byte) (err error) {
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
		AwsLambdaFuncCount                      datadog.NullableInt64 `json:"aws_lambda_func_count,omitempty"`
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
		CloudCostManagementHostCountAvg         datadog.NullableInt64 `json:"cloud_cost_management_host_count_avg,omitempty"`
		ContainerAvg                            datadog.NullableInt64 `json:"container_avg,omitempty"`
		ContainerExclAgentAvg                   datadog.NullableInt64 `json:"container_excl_agent_avg,omitempty"`
		ContainerHwm                            datadog.NullableInt64 `json:"container_hwm,omitempty"`
		CspmAasHostTop99p                       datadog.NullableInt64 `json:"cspm_aas_host_top99p,omitempty"`
		CspmAwsHostTop99p                       datadog.NullableInt64 `json:"cspm_aws_host_top99p,omitempty"`
		CspmAzureHostTop99p                     datadog.NullableInt64 `json:"cspm_azure_host_top99p,omitempty"`
		CspmContainerAvg                        datadog.NullableInt64 `json:"cspm_container_avg,omitempty"`
		CspmContainerHwm                        datadog.NullableInt64 `json:"cspm_container_hwm,omitempty"`
		CspmGcpHostTop99p                       datadog.NullableInt64 `json:"cspm_gcp_host_top99p,omitempty"`
		CspmHostTop99p                          datadog.NullableInt64 `json:"cspm_host_top99p,omitempty"`
		CustomTsAvg                             datadog.NullableInt64 `json:"custom_ts_avg,omitempty"`
		CwsContainerCountAvg                    datadog.NullableInt64 `json:"cws_container_count_avg,omitempty"`
		CwsHostTop99p                           datadog.NullableInt64 `json:"cws_host_top99p,omitempty"`
		Date                                    *time.Time            `json:"date,omitempty"`
		DbmHostTop99p                           datadog.NullableInt64 `json:"dbm_host_top99p,omitempty"`
		DbmQueriesCountAvg                      datadog.NullableInt64 `json:"dbm_queries_count_avg,omitempty"`
		FargateTasksCountAvg                    datadog.NullableInt64 `json:"fargate_tasks_count_avg,omitempty"`
		FargateTasksCountHwm                    datadog.NullableInt64 `json:"fargate_tasks_count_hwm,omitempty"`
		ForwardingEventsBytesSum                datadog.NullableInt64 `json:"forwarding_events_bytes_sum,omitempty"`
		GcpHostTop99p                           datadog.NullableInt64 `json:"gcp_host_top99p,omitempty"`
		HerokuHostTop99p                        datadog.NullableInt64 `json:"heroku_host_top99p,omitempty"`
		IncidentManagementMonthlyActiveUsersHwm datadog.NullableInt64 `json:"incident_management_monthly_active_users_hwm,omitempty"`
		IndexedEventsCountSum                   datadog.NullableInt64 `json:"indexed_events_count_sum,omitempty"`
		InfraHostTop99p                         datadog.NullableInt64 `json:"infra_host_top99p,omitempty"`
		IngestedEventsBytesSum                  datadog.NullableInt64 `json:"ingested_events_bytes_sum,omitempty"`
		IotDeviceSum                            datadog.NullableInt64 `json:"iot_device_sum,omitempty"`
		IotDeviceTop99p                         datadog.NullableInt64 `json:"iot_device_top99p,omitempty"`
		MobileRumLiteSessionCountSum            datadog.NullableInt64 `json:"mobile_rum_lite_session_count_sum,omitempty"`
		MobileRumSessionCountAndroidSum         datadog.NullableInt64 `json:"mobile_rum_session_count_android_sum,omitempty"`
		MobileRumSessionCountFlutterSum         datadog.NullableInt64 `json:"mobile_rum_session_count_flutter_sum,omitempty"`
		MobileRumSessionCountIosSum             datadog.NullableInt64 `json:"mobile_rum_session_count_ios_sum,omitempty"`
		MobileRumSessionCountReactnativeSum     datadog.NullableInt64 `json:"mobile_rum_session_count_reactnative_sum,omitempty"`
		MobileRumSessionCountSum                datadog.NullableInt64 `json:"mobile_rum_session_count_sum,omitempty"`
		MobileRumUnitsSum                       datadog.NullableInt64 `json:"mobile_rum_units_sum,omitempty"`
		NetflowIndexedEventsCountSum            datadog.NullableInt64 `json:"netflow_indexed_events_count_sum,omitempty"`
		NpmHostTop99p                           datadog.NullableInt64 `json:"npm_host_top99p,omitempty"`
		ObservabilityPipelinesBytesProcessedSum datadog.NullableInt64 `json:"observability_pipelines_bytes_processed_sum,omitempty"`
		OnlineArchiveEventsCountSum             datadog.NullableInt64 `json:"online_archive_events_count_sum,omitempty"`
		OpentelemetryApmHostTop99p              datadog.NullableInt64 `json:"opentelemetry_apm_host_top99p,omitempty"`
		OpentelemetryHostTop99p                 datadog.NullableInt64 `json:"opentelemetry_host_top99p,omitempty"`
		Orgs                                    []UsageSummaryDateOrg `json:"orgs,omitempty"`
		ProfilingHostTop99p                     datadog.NullableInt64 `json:"profiling_host_top99p,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_host_top99p", "apm_azure_app_service_host_top99p", "apm_fargate_count_avg", "apm_host_top99p", "appsec_fargate_count_avg", "audit_logs_lines_indexed_sum", "audit_trail_enabled_hwm", "avg_profiled_fargate_tasks", "aws_host_top99p", "aws_lambda_func_count", "aws_lambda_invocations_sum", "azure_app_service_top99p", "billable_ingested_bytes_sum", "browser_rum_lite_session_count_sum", "browser_rum_replay_session_count_sum", "browser_rum_units_sum", "ci_pipeline_indexed_spans_sum", "ci_test_indexed_spans_sum", "ci_visibility_pipeline_committers_hwm", "ci_visibility_test_committers_hwm", "cloud_cost_management_host_count_avg", "container_avg", "container_excl_agent_avg", "container_hwm", "cspm_aas_host_top99p", "cspm_aws_host_top99p", "cspm_azure_host_top99p", "cspm_container_avg", "cspm_container_hwm", "cspm_gcp_host_top99p", "cspm_host_top99p", "custom_ts_avg", "cws_container_count_avg", "cws_host_top99p", "date", "dbm_host_top99p", "dbm_queries_count_avg", "fargate_tasks_count_avg", "fargate_tasks_count_hwm", "forwarding_events_bytes_sum", "gcp_host_top99p", "heroku_host_top99p", "incident_management_monthly_active_users_hwm", "indexed_events_count_sum", "infra_host_top99p", "ingested_events_bytes_sum", "iot_device_sum", "iot_device_top99p", "mobile_rum_lite_session_count_sum", "mobile_rum_session_count_android_sum", "mobile_rum_session_count_flutter_sum", "mobile_rum_session_count_ios_sum", "mobile_rum_session_count_reactnative_sum", "mobile_rum_session_count_sum", "mobile_rum_units_sum", "netflow_indexed_events_count_sum", "npm_host_top99p", "observability_pipelines_bytes_processed_sum", "online_archive_events_count_sum", "opentelemetry_apm_host_top99p", "opentelemetry_host_top99p", "orgs", "profiling_host_top99p", "rum_browser_and_mobile_session_count", "rum_session_count_sum", "rum_total_session_count_sum", "rum_units_sum", "sds_apm_scanned_bytes_sum", "sds_events_scanned_bytes_sum", "sds_logs_scanned_bytes_sum", "sds_rum_scanned_bytes_sum", "sds_total_scanned_bytes_sum", "synthetics_browser_check_calls_count_sum", "synthetics_check_calls_count_sum", "synthetics_parallel_testing_max_slots_hwm", "trace_search_indexed_events_count_sum", "twol_ingested_events_bytes_sum", "universal_service_monitoring_host_top99p", "vsphere_host_top99p"})
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
	o.Date = all.Date
	o.DbmHostTop99p = all.DbmHostTop99p
	o.DbmQueriesCountAvg = all.DbmQueriesCountAvg
	o.FargateTasksCountAvg = all.FargateTasksCountAvg
	o.FargateTasksCountHwm = all.FargateTasksCountHwm
	o.ForwardingEventsBytesSum = all.ForwardingEventsBytesSum
	o.GcpHostTop99p = all.GcpHostTop99p
	o.HerokuHostTop99p = all.HerokuHostTop99p
	o.IncidentManagementMonthlyActiveUsersHwm = all.IncidentManagementMonthlyActiveUsersHwm
	o.IndexedEventsCountSum = all.IndexedEventsCountSum
	o.InfraHostTop99p = all.InfraHostTop99p
	o.IngestedEventsBytesSum = all.IngestedEventsBytesSum
	o.IotDeviceSum = all.IotDeviceSum
	o.IotDeviceTop99p = all.IotDeviceTop99p
	o.MobileRumLiteSessionCountSum = all.MobileRumLiteSessionCountSum
	o.MobileRumSessionCountAndroidSum = all.MobileRumSessionCountAndroidSum
	o.MobileRumSessionCountFlutterSum = all.MobileRumSessionCountFlutterSum
	o.MobileRumSessionCountIosSum = all.MobileRumSessionCountIosSum
	o.MobileRumSessionCountReactnativeSum = all.MobileRumSessionCountReactnativeSum
	o.MobileRumSessionCountSum = all.MobileRumSessionCountSum
	o.MobileRumUnitsSum = all.MobileRumUnitsSum
	o.NetflowIndexedEventsCountSum = all.NetflowIndexedEventsCountSum
	o.NpmHostTop99p = all.NpmHostTop99p
	o.ObservabilityPipelinesBytesProcessedSum = all.ObservabilityPipelinesBytesProcessedSum
	o.OnlineArchiveEventsCountSum = all.OnlineArchiveEventsCountSum
	o.OpentelemetryApmHostTop99p = all.OpentelemetryApmHostTop99p
	o.OpentelemetryHostTop99p = all.OpentelemetryHostTop99p
	o.Orgs = all.Orgs
	o.ProfilingHostTop99p = all.ProfilingHostTop99p
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
