// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSummaryResponse Response summarizing all usage aggregated across the months in the request for all organizations, and broken down by month and by organization.
type UsageSummaryResponse struct {
	// Shows the 99th percentile of all agent hosts over all hours in the current months for all organizations.
	AgentHostTop99pSum datadog.NullableInt64 `json:"agent_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Azure app services using APM over all hours in the current months all organizations.
	ApmAzureAppServiceHostTop99pSum datadog.NullableInt64 `json:"apm_azure_app_service_host_top99p_sum,omitempty"`
	// Shows the average of all APM ECS Fargate tasks over all hours in the current months for all organizations.
	ApmFargateCountAvgSum datadog.NullableInt64 `json:"apm_fargate_count_avg_sum,omitempty"`
	// Shows the 99th percentile of all distinct APM hosts over all hours in the current months for all organizations.
	ApmHostTop99pSum datadog.NullableInt64 `json:"apm_host_top99p_sum,omitempty"`
	// Shows the average of all Application Security Monitoring ECS Fargate tasks over all hours in the current months for all organizations.
	AppsecFargateCountAvgSum datadog.NullableInt64 `json:"appsec_fargate_count_avg_sum,omitempty"`
	// Shows the sum of all audit logs lines indexed over all hours in the current months for all organizations.
	// Deprecated
	AuditLogsLinesIndexedAggSum datadog.NullableInt64 `json:"audit_logs_lines_indexed_agg_sum,omitempty"`
	// Shows the total number of organizations that had Audit Trail enabled over a specific number of months.
	AuditTrailEnabledHwmSum datadog.NullableInt64 `json:"audit_trail_enabled_hwm_sum,omitempty"`
	// Shows the average of all profiled Fargate tasks over all hours in the current months for all organizations.
	AvgProfiledFargateTasksSum datadog.NullableInt64 `json:"avg_profiled_fargate_tasks_sum,omitempty"`
	// Shows the 99th percentile of all AWS hosts over all hours in the current months for all organizations.
	AwsHostTop99pSum datadog.NullableInt64 `json:"aws_host_top99p_sum,omitempty"`
	// Shows the average of the number of functions that executed 1 or more times each hour in the current months for all organizations.
	AwsLambdaFuncCount datadog.NullableInt64 `json:"aws_lambda_func_count,omitempty"`
	// Shows the sum of all AWS Lambda invocations over all hours in the current months for all organizations.
	AwsLambdaInvocationsSum datadog.NullableInt64 `json:"aws_lambda_invocations_sum,omitempty"`
	// Shows the 99th percentile of all Azure app services over all hours in the current months for all organizations.
	AzureAppServiceTop99pSum datadog.NullableInt64 `json:"azure_app_service_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Azure hosts over all hours in the current months for all organizations.
	AzureHostTop99pSum datadog.NullableInt64 `json:"azure_host_top99p_sum,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current months for all organizations.
	BillableIngestedBytesAggSum datadog.NullableInt64 `json:"billable_ingested_bytes_agg_sum,omitempty"`
	// Shows the sum of all browser lite sessions over all hours in the current months for all organizations.
	BrowserRumLiteSessionCountAggSum datadog.NullableInt64 `json:"browser_rum_lite_session_count_agg_sum,omitempty"`
	// Shows the sum of all browser replay sessions over all hours in the current months for all organizations.
	BrowserRumReplaySessionCountAggSum datadog.NullableInt64 `json:"browser_rum_replay_session_count_agg_sum,omitempty"`
	// Shows the sum of all browser RUM units over all hours in the current months for all organizations.
	BrowserRumUnitsAggSum datadog.NullableInt64 `json:"browser_rum_units_agg_sum,omitempty"`
	// Shows the sum of all CI pipeline indexed spans over all hours in the current months for all organizations.
	CiPipelineIndexedSpansAggSum datadog.NullableInt64 `json:"ci_pipeline_indexed_spans_agg_sum,omitempty"`
	// Shows the sum of all CI test indexed spans over all hours in the current months for all organizations.
	CiTestIndexedSpansAggSum datadog.NullableInt64 `json:"ci_test_indexed_spans_agg_sum,omitempty"`
	// Shows the high-water mark of all CI visibility pipeline committers over all hours in the current months for all organizations.
	CiVisibilityPipelineCommittersHwmSum datadog.NullableInt64 `json:"ci_visibility_pipeline_committers_hwm_sum,omitempty"`
	// Shows the high-water mark of all CI visibility test committers over all hours in the current months for all organizations.
	CiVisibilityTestCommittersHwmSum datadog.NullableInt64 `json:"ci_visibility_test_committers_hwm_sum,omitempty"`
	// Sum of the host count average for Cloud Cost Management.
	CloudCostManagementHostCountAvgSum datadog.NullableInt64 `json:"cloud_cost_management_host_count_avg_sum,omitempty"`
	// Shows the average of all distinct containers over all hours in the current months for all organizations.
	ContainerAvgSum datadog.NullableInt64 `json:"container_avg_sum,omitempty"`
	// Shows the average of the containers without the Datadog Agent over all hours in the current month for all organizations.
	ContainerExclAgentAvgSum datadog.NullableInt64 `json:"container_excl_agent_avg_sum,omitempty"`
	// Shows the sum of the high-water marks of all distinct containers over all hours in the current months for all organizations.
	ContainerHwmSum datadog.NullableInt64 `json:"container_hwm_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management Azure app services hosts over all hours in the current months for all organizations.
	CspmAasHostTop99pSum datadog.NullableInt64 `json:"cspm_aas_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management AWS hosts over all hours in the current months for all organizations.
	CspmAwsHostTop99pSum datadog.NullableInt64 `json:"cspm_aws_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management Azure hosts over all hours in the current months for all organizations.
	CspmAzureHostTop99pSum datadog.NullableInt64 `json:"cspm_azure_host_top99p_sum,omitempty"`
	// Shows the average number of Cloud Security Posture Management containers over all hours in the current months for all organizations.
	CspmContainerAvgSum datadog.NullableInt64 `json:"cspm_container_avg_sum,omitempty"`
	// Shows the sum of the the high-water marks of Cloud Security Posture Management containers over all hours in the current months for all organizations.
	CspmContainerHwmSum datadog.NullableInt64 `json:"cspm_container_hwm_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management GCP hosts over all hours in the current months for all organizations.
	CspmGcpHostTop99pSum datadog.NullableInt64 `json:"cspm_gcp_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Posture Management hosts over all hours in the current months for all organizations.
	CspmHostTop99pSum datadog.NullableInt64 `json:"cspm_host_top99p_sum,omitempty"`
	// Shows the average number of distinct custom metrics over all hours in the current months for all organizations.
	CustomTsSum datadog.NullableInt64 `json:"custom_ts_sum,omitempty"`
	// Shows the average of all distinct Cloud Workload Security containers over all hours in the current months for all organizations.
	CwsContainersAvgSum datadog.NullableInt64 `json:"cws_containers_avg_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Workload Security hosts over all hours in the current months for all organizations.
	CwsHostTop99pSum datadog.NullableInt64 `json:"cws_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Database Monitoring hosts over all hours in the current month for all organizations.
	DbmHostTop99pSum datadog.NullableInt64 `json:"dbm_host_top99p_sum,omitempty"`
	// Shows the average of all distinct Database Monitoring Normalized Queries over all hours in the current month for all organizations.
	DbmQueriesAvgSum datadog.NullableInt64 `json:"dbm_queries_avg_sum,omitempty"`
	// Shows the last date of usage in the current months for all organizations.
	EndDate *time.Time `json:"end_date,omitempty"`
	// Shows the average of all Fargate tasks over all hours in the current months for all organizations.
	FargateTasksCountAvgSum datadog.NullableInt64 `json:"fargate_tasks_count_avg_sum,omitempty"`
	// Shows the sum of the high-water marks of all Fargate tasks over all hours in the current months for all organizations.
	FargateTasksCountHwmSum datadog.NullableInt64 `json:"fargate_tasks_count_hwm_sum,omitempty"`
	// Shows the sum of all logs forwarding bytes over all hours in the current months for all organizations (data available as of April 1, 2023)
	ForwardingEventsBytesAggSum datadog.NullableInt64 `json:"forwarding_events_bytes_agg_sum,omitempty"`
	// Shows the 99th percentile of all GCP hosts over all hours in the current months for all organizations.
	GcpHostTop99pSum datadog.NullableInt64 `json:"gcp_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Heroku dynos over all hours in the current months for all organizations.
	HerokuHostTop99pSum datadog.NullableInt64 `json:"heroku_host_top99p_sum,omitempty"`
	// Shows sum of the the high-water marks of incident management monthly active users in the current months for all organizations.
	IncidentManagementMonthlyActiveUsersHwmSum datadog.NullableInt64 `json:"incident_management_monthly_active_users_hwm_sum,omitempty"`
	// Shows the sum of all log events indexed over all hours in the current months for all organizations.
	IndexedEventsCountAggSum datadog.NullableInt64 `json:"indexed_events_count_agg_sum,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current months for all organizations.
	InfraHostTop99pSum datadog.NullableInt64 `json:"infra_host_top99p_sum,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current months for all organizations.
	IngestedEventsBytesAggSum datadog.NullableInt64 `json:"ingested_events_bytes_agg_sum,omitempty"`
	// Shows the sum of all IoT devices over all hours in the current months for all organizations.
	IotDeviceAggSum datadog.NullableInt64 `json:"iot_device_agg_sum,omitempty"`
	// Shows the 99th percentile of all IoT devices over all hours in the current months of all organizations.
	IotDeviceTop99pSum datadog.NullableInt64 `json:"iot_device_top99p_sum,omitempty"`
	// Shows the the most recent hour in the current months for all organizations for which all usages were calculated.
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Shows the sum of all live logs indexed over all hours in the current months for all organizations (data available as of December 1, 2020).
	LiveIndexedEventsAggSum datadog.NullableInt64 `json:"live_indexed_events_agg_sum,omitempty"`
	// Shows the sum of all live logs bytes ingested over all hours in the current months for all organizations (data available as of December 1, 2020).
	LiveIngestedBytesAggSum datadog.NullableInt64 `json:"live_ingested_bytes_agg_sum,omitempty"`
	// Object containing logs usage data broken down by retention period.
	LogsByRetention *LogsByRetention `json:"logs_by_retention,omitempty"`
	// Shows the sum of all mobile lite sessions over all hours in the current months for all organizations.
	MobileRumLiteSessionCountAggSum datadog.NullableInt64 `json:"mobile_rum_lite_session_count_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions over all hours in the current months for all organizations.
	MobileRumSessionCountAggSum datadog.NullableInt64 `json:"mobile_rum_session_count_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on Android over all hours in the current months for all organizations.
	MobileRumSessionCountAndroidAggSum datadog.NullableInt64 `json:"mobile_rum_session_count_android_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on Flutter over all hours in the current months for all organizations.
	MobileRumSessionCountFlutterAggSum datadog.NullableInt64 `json:"mobile_rum_session_count_flutter_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on iOS over all hours in the current months for all organizations.
	MobileRumSessionCountIosAggSum datadog.NullableInt64 `json:"mobile_rum_session_count_ios_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM Sessions on React Native over all hours in the current months for all organizations.
	MobileRumSessionCountReactnativeAggSum datadog.NullableInt64 `json:"mobile_rum_session_count_reactnative_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM units over all hours in the current months for all organizations.
	MobileRumUnitsAggSum datadog.NullableInt64 `json:"mobile_rum_units_agg_sum,omitempty"`
	// Shows the sum of all Network flows indexed over all hours in the current months for all organizations.
	NetflowIndexedEventsCountAggSum datadog.NullableInt64 `json:"netflow_indexed_events_count_agg_sum,omitempty"`
	// Shows the 99th percentile of all distinct Networks hosts over all hours in the current months for all organizations.
	NpmHostTop99pSum datadog.NullableInt64 `json:"npm_host_top99p_sum,omitempty"`
	// Sum of all observability pipelines bytes processed over all hours in the current months for all organizations.
	ObservabilityPipelinesBytesProcessedAggSum datadog.NullableInt64 `json:"observability_pipelines_bytes_processed_agg_sum,omitempty"`
	// Sum of all online archived events over all hours in the current months for all organizations.
	OnlineArchiveEventsCountAggSum datadog.NullableInt64 `json:"online_archive_events_count_agg_sum,omitempty"`
	// Shows the 99th percentile of APM hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current months for all organizations.
	OpentelemetryApmHostTop99pSum datadog.NullableInt64 `json:"opentelemetry_apm_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current months for all organizations.
	OpentelemetryHostTop99pSum datadog.NullableInt64 `json:"opentelemetry_host_top99p_sum,omitempty"`
	// Shows the average number of profiled containers over all hours in the current months for all organizations.
	ProfilingContainerAgentCountAvg datadog.NullableInt64 `json:"profiling_container_agent_count_avg,omitempty"`
	// Shows the 99th percentile of all profiled hosts over all hours in the current months for all organizations.
	ProfilingHostCountTop99pSum datadog.NullableInt64 `json:"profiling_host_count_top99p_sum,omitempty"`
	// Shows the sum of all rehydrated logs indexed over all hours in the current months for all organizations (data available as of December 1, 2020).
	RehydratedIndexedEventsAggSum datadog.NullableInt64 `json:"rehydrated_indexed_events_agg_sum,omitempty"`
	// Shows the sum of all rehydrated logs bytes ingested over all hours in the current months for all organizations (data available as of December 1, 2020).
	RehydratedIngestedBytesAggSum datadog.NullableInt64 `json:"rehydrated_ingested_bytes_agg_sum,omitempty"`
	// Shows the sum of all mobile sessions and all browser lite and legacy sessions over all hours in the current month for all organizations.
	RumBrowserAndMobileSessionCount datadog.NullableInt64 `json:"rum_browser_and_mobile_session_count,omitempty"`
	// Shows the sum of all browser RUM Lite Sessions over all hours in the current months for all organizations.
	RumSessionCountAggSum datadog.NullableInt64 `json:"rum_session_count_agg_sum,omitempty"`
	// Shows the sum of RUM Sessions (browser and mobile) over all hours in the current months for all organizations.
	RumTotalSessionCountAggSum datadog.NullableInt64 `json:"rum_total_session_count_agg_sum,omitempty"`
	// Shows the sum of all browser and mobile RUM units over all hours in the current months for all organizations.
	RumUnitsAggSum datadog.NullableInt64 `json:"rum_units_agg_sum,omitempty"`
	// Sum of all APM bytes scanned with sensitive data scanner in the current months for all organizations.
	SdsApmScannedBytesSum datadog.NullableInt64 `json:"sds_apm_scanned_bytes_sum,omitempty"`
	// Sum of all event stream events bytes scanned with sensitive data scanner in the current months for all organizations.
	SdsEventsScannedBytesSum datadog.NullableInt64 `json:"sds_events_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned of logs usage by the Sensitive Data Scanner over all hours in the current month for all organizations.
	SdsLogsScannedBytesSum datadog.NullableInt64 `json:"sds_logs_scanned_bytes_sum,omitempty"`
	// Sum of all RUM bytes scanned with sensitive data scanner in the current months for all organizations.
	SdsRumScannedBytesSum datadog.NullableInt64 `json:"sds_rum_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned across all usage types by the Sensitive Data Scanner over all hours in the current month for all organizations.
	SdsTotalScannedBytesSum datadog.NullableInt64 `json:"sds_total_scanned_bytes_sum,omitempty"`
	// Shows the first date of usage in the current months for all organizations.
	StartDate *time.Time `json:"start_date,omitempty"`
	// Shows the sum of all Synthetic browser tests over all hours in the current months for all organizations.
	SyntheticsBrowserCheckCallsCountAggSum datadog.NullableInt64 `json:"synthetics_browser_check_calls_count_agg_sum,omitempty"`
	// Shows the sum of all Synthetic API tests over all hours in the current months for all organizations.
	SyntheticsCheckCallsCountAggSum datadog.NullableInt64 `json:"synthetics_check_calls_count_agg_sum,omitempty"`
	// Shows the sum of the high-water marks of used synthetics parallel testing slots over all hours in the current month for all organizations.
	SyntheticsParallelTestingMaxSlotsHwmSum datadog.NullableInt64 `json:"synthetics_parallel_testing_max_slots_hwm_sum,omitempty"`
	// Shows the sum of all Indexed Spans indexed over all hours in the current months for all organizations.
	TraceSearchIndexedEventsCountAggSum datadog.NullableInt64 `json:"trace_search_indexed_events_count_agg_sum,omitempty"`
	// Shows the sum of all ingested APM span bytes over all hours in the current months for all organizations.
	TwolIngestedEventsBytesAggSum datadog.NullableInt64 `json:"twol_ingested_events_bytes_agg_sum,omitempty"`
	// Shows the 99th percentile of all Universal Service Monitoring hosts over all hours in the current months for all organizations.
	UniversalServiceMonitoringHostTop99pSum datadog.NullableInt64 `json:"universal_service_monitoring_host_top99p_sum,omitempty"`
	// An array of objects regarding hourly usage.
	Usage []UsageSummaryDate `json:"usage,omitempty"`
	// Shows the 99th percentile of all vSphere hosts over all hours in the current months for all organizations.
	VsphereHostTop99pSum datadog.NullableInt64 `json:"vsphere_host_top99p_sum,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageSummaryResponse instantiates a new UsageSummaryResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSummaryResponse() *UsageSummaryResponse {
	this := UsageSummaryResponse{}
	return &this
}

// NewUsageSummaryResponseWithDefaults instantiates a new UsageSummaryResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSummaryResponseWithDefaults() *UsageSummaryResponse {
	this := UsageSummaryResponse{}
	return &this
}

// GetAgentHostTop99pSum returns the AgentHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetAgentHostTop99pSum() int64 {
	if o == nil || o.AgentHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AgentHostTop99pSum.Get()
}

// GetAgentHostTop99pSumOk returns a tuple with the AgentHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetAgentHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentHostTop99pSum.Get(), o.AgentHostTop99pSum.IsSet()
}

// HasAgentHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAgentHostTop99pSum() bool {
	return o != nil && o.AgentHostTop99pSum.IsSet()
}

// SetAgentHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the AgentHostTop99pSum field.
func (o *UsageSummaryResponse) SetAgentHostTop99pSum(v int64) {
	o.AgentHostTop99pSum.Set(&v)
}

// SetAgentHostTop99pSumNil sets the value for AgentHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetAgentHostTop99pSumNil() {
	o.AgentHostTop99pSum.Set(nil)
}

// UnsetAgentHostTop99pSum ensures that no value is present for AgentHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAgentHostTop99pSum() {
	o.AgentHostTop99pSum.Unset()
}

// GetApmAzureAppServiceHostTop99pSum returns the ApmAzureAppServiceHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetApmAzureAppServiceHostTop99pSum() int64 {
	if o == nil || o.ApmAzureAppServiceHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmAzureAppServiceHostTop99pSum.Get()
}

// GetApmAzureAppServiceHostTop99pSumOk returns a tuple with the ApmAzureAppServiceHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetApmAzureAppServiceHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmAzureAppServiceHostTop99pSum.Get(), o.ApmAzureAppServiceHostTop99pSum.IsSet()
}

// HasApmAzureAppServiceHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmAzureAppServiceHostTop99pSum() bool {
	return o != nil && o.ApmAzureAppServiceHostTop99pSum.IsSet()
}

// SetApmAzureAppServiceHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the ApmAzureAppServiceHostTop99pSum field.
func (o *UsageSummaryResponse) SetApmAzureAppServiceHostTop99pSum(v int64) {
	o.ApmAzureAppServiceHostTop99pSum.Set(&v)
}

// SetApmAzureAppServiceHostTop99pSumNil sets the value for ApmAzureAppServiceHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetApmAzureAppServiceHostTop99pSumNil() {
	o.ApmAzureAppServiceHostTop99pSum.Set(nil)
}

// UnsetApmAzureAppServiceHostTop99pSum ensures that no value is present for ApmAzureAppServiceHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetApmAzureAppServiceHostTop99pSum() {
	o.ApmAzureAppServiceHostTop99pSum.Unset()
}

// GetApmFargateCountAvgSum returns the ApmFargateCountAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetApmFargateCountAvgSum() int64 {
	if o == nil || o.ApmFargateCountAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmFargateCountAvgSum.Get()
}

// GetApmFargateCountAvgSumOk returns a tuple with the ApmFargateCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetApmFargateCountAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmFargateCountAvgSum.Get(), o.ApmFargateCountAvgSum.IsSet()
}

// HasApmFargateCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmFargateCountAvgSum() bool {
	return o != nil && o.ApmFargateCountAvgSum.IsSet()
}

// SetApmFargateCountAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the ApmFargateCountAvgSum field.
func (o *UsageSummaryResponse) SetApmFargateCountAvgSum(v int64) {
	o.ApmFargateCountAvgSum.Set(&v)
}

// SetApmFargateCountAvgSumNil sets the value for ApmFargateCountAvgSum to be an explicit nil.
func (o *UsageSummaryResponse) SetApmFargateCountAvgSumNil() {
	o.ApmFargateCountAvgSum.Set(nil)
}

// UnsetApmFargateCountAvgSum ensures that no value is present for ApmFargateCountAvgSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetApmFargateCountAvgSum() {
	o.ApmFargateCountAvgSum.Unset()
}

// GetApmHostTop99pSum returns the ApmHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetApmHostTop99pSum() int64 {
	if o == nil || o.ApmHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmHostTop99pSum.Get()
}

// GetApmHostTop99pSumOk returns a tuple with the ApmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetApmHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmHostTop99pSum.Get(), o.ApmHostTop99pSum.IsSet()
}

// HasApmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmHostTop99pSum() bool {
	return o != nil && o.ApmHostTop99pSum.IsSet()
}

// SetApmHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the ApmHostTop99pSum field.
func (o *UsageSummaryResponse) SetApmHostTop99pSum(v int64) {
	o.ApmHostTop99pSum.Set(&v)
}

// SetApmHostTop99pSumNil sets the value for ApmHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetApmHostTop99pSumNil() {
	o.ApmHostTop99pSum.Set(nil)
}

// UnsetApmHostTop99pSum ensures that no value is present for ApmHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetApmHostTop99pSum() {
	o.ApmHostTop99pSum.Unset()
}

// GetAppsecFargateCountAvgSum returns the AppsecFargateCountAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetAppsecFargateCountAvgSum() int64 {
	if o == nil || o.AppsecFargateCountAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AppsecFargateCountAvgSum.Get()
}

// GetAppsecFargateCountAvgSumOk returns a tuple with the AppsecFargateCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetAppsecFargateCountAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecFargateCountAvgSum.Get(), o.AppsecFargateCountAvgSum.IsSet()
}

// HasAppsecFargateCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAppsecFargateCountAvgSum() bool {
	return o != nil && o.AppsecFargateCountAvgSum.IsSet()
}

// SetAppsecFargateCountAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the AppsecFargateCountAvgSum field.
func (o *UsageSummaryResponse) SetAppsecFargateCountAvgSum(v int64) {
	o.AppsecFargateCountAvgSum.Set(&v)
}

// SetAppsecFargateCountAvgSumNil sets the value for AppsecFargateCountAvgSum to be an explicit nil.
func (o *UsageSummaryResponse) SetAppsecFargateCountAvgSumNil() {
	o.AppsecFargateCountAvgSum.Set(nil)
}

// UnsetAppsecFargateCountAvgSum ensures that no value is present for AppsecFargateCountAvgSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAppsecFargateCountAvgSum() {
	o.AppsecFargateCountAvgSum.Unset()
}

// GetAuditLogsLinesIndexedAggSum returns the AuditLogsLinesIndexedAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *UsageSummaryResponse) GetAuditLogsLinesIndexedAggSum() int64 {
	if o == nil || o.AuditLogsLinesIndexedAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AuditLogsLinesIndexedAggSum.Get()
}

// GetAuditLogsLinesIndexedAggSumOk returns a tuple with the AuditLogsLinesIndexedAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
// Deprecated
func (o *UsageSummaryResponse) GetAuditLogsLinesIndexedAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditLogsLinesIndexedAggSum.Get(), o.AuditLogsLinesIndexedAggSum.IsSet()
}

// HasAuditLogsLinesIndexedAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAuditLogsLinesIndexedAggSum() bool {
	return o != nil && o.AuditLogsLinesIndexedAggSum.IsSet()
}

// SetAuditLogsLinesIndexedAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the AuditLogsLinesIndexedAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetAuditLogsLinesIndexedAggSum(v int64) {
	o.AuditLogsLinesIndexedAggSum.Set(&v)
}

// SetAuditLogsLinesIndexedAggSumNil sets the value for AuditLogsLinesIndexedAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetAuditLogsLinesIndexedAggSumNil() {
	o.AuditLogsLinesIndexedAggSum.Set(nil)
}

// UnsetAuditLogsLinesIndexedAggSum ensures that no value is present for AuditLogsLinesIndexedAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAuditLogsLinesIndexedAggSum() {
	o.AuditLogsLinesIndexedAggSum.Unset()
}

// GetAuditTrailEnabledHwmSum returns the AuditTrailEnabledHwmSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetAuditTrailEnabledHwmSum() int64 {
	if o == nil || o.AuditTrailEnabledHwmSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AuditTrailEnabledHwmSum.Get()
}

// GetAuditTrailEnabledHwmSumOk returns a tuple with the AuditTrailEnabledHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetAuditTrailEnabledHwmSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditTrailEnabledHwmSum.Get(), o.AuditTrailEnabledHwmSum.IsSet()
}

// HasAuditTrailEnabledHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAuditTrailEnabledHwmSum() bool {
	return o != nil && o.AuditTrailEnabledHwmSum.IsSet()
}

// SetAuditTrailEnabledHwmSum gets a reference to the given datadog.NullableInt64 and assigns it to the AuditTrailEnabledHwmSum field.
func (o *UsageSummaryResponse) SetAuditTrailEnabledHwmSum(v int64) {
	o.AuditTrailEnabledHwmSum.Set(&v)
}

// SetAuditTrailEnabledHwmSumNil sets the value for AuditTrailEnabledHwmSum to be an explicit nil.
func (o *UsageSummaryResponse) SetAuditTrailEnabledHwmSumNil() {
	o.AuditTrailEnabledHwmSum.Set(nil)
}

// UnsetAuditTrailEnabledHwmSum ensures that no value is present for AuditTrailEnabledHwmSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAuditTrailEnabledHwmSum() {
	o.AuditTrailEnabledHwmSum.Unset()
}

// GetAvgProfiledFargateTasksSum returns the AvgProfiledFargateTasksSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetAvgProfiledFargateTasksSum() int64 {
	if o == nil || o.AvgProfiledFargateTasksSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AvgProfiledFargateTasksSum.Get()
}

// GetAvgProfiledFargateTasksSumOk returns a tuple with the AvgProfiledFargateTasksSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetAvgProfiledFargateTasksSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvgProfiledFargateTasksSum.Get(), o.AvgProfiledFargateTasksSum.IsSet()
}

// HasAvgProfiledFargateTasksSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAvgProfiledFargateTasksSum() bool {
	return o != nil && o.AvgProfiledFargateTasksSum.IsSet()
}

// SetAvgProfiledFargateTasksSum gets a reference to the given datadog.NullableInt64 and assigns it to the AvgProfiledFargateTasksSum field.
func (o *UsageSummaryResponse) SetAvgProfiledFargateTasksSum(v int64) {
	o.AvgProfiledFargateTasksSum.Set(&v)
}

// SetAvgProfiledFargateTasksSumNil sets the value for AvgProfiledFargateTasksSum to be an explicit nil.
func (o *UsageSummaryResponse) SetAvgProfiledFargateTasksSumNil() {
	o.AvgProfiledFargateTasksSum.Set(nil)
}

// UnsetAvgProfiledFargateTasksSum ensures that no value is present for AvgProfiledFargateTasksSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAvgProfiledFargateTasksSum() {
	o.AvgProfiledFargateTasksSum.Unset()
}

// GetAwsHostTop99pSum returns the AwsHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetAwsHostTop99pSum() int64 {
	if o == nil || o.AwsHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AwsHostTop99pSum.Get()
}

// GetAwsHostTop99pSumOk returns a tuple with the AwsHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetAwsHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsHostTop99pSum.Get(), o.AwsHostTop99pSum.IsSet()
}

// HasAwsHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAwsHostTop99pSum() bool {
	return o != nil && o.AwsHostTop99pSum.IsSet()
}

// SetAwsHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the AwsHostTop99pSum field.
func (o *UsageSummaryResponse) SetAwsHostTop99pSum(v int64) {
	o.AwsHostTop99pSum.Set(&v)
}

// SetAwsHostTop99pSumNil sets the value for AwsHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetAwsHostTop99pSumNil() {
	o.AwsHostTop99pSum.Set(nil)
}

// UnsetAwsHostTop99pSum ensures that no value is present for AwsHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAwsHostTop99pSum() {
	o.AwsHostTop99pSum.Unset()
}

// GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetAwsLambdaFuncCount() int64 {
	if o == nil || o.AwsLambdaFuncCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaFuncCount.Get()
}

// GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetAwsLambdaFuncCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsLambdaFuncCount.Get(), o.AwsLambdaFuncCount.IsSet()
}

// HasAwsLambdaFuncCount returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAwsLambdaFuncCount() bool {
	return o != nil && o.AwsLambdaFuncCount.IsSet()
}

// SetAwsLambdaFuncCount gets a reference to the given datadog.NullableInt64 and assigns it to the AwsLambdaFuncCount field.
func (o *UsageSummaryResponse) SetAwsLambdaFuncCount(v int64) {
	o.AwsLambdaFuncCount.Set(&v)
}

// SetAwsLambdaFuncCountNil sets the value for AwsLambdaFuncCount to be an explicit nil.
func (o *UsageSummaryResponse) SetAwsLambdaFuncCountNil() {
	o.AwsLambdaFuncCount.Set(nil)
}

// UnsetAwsLambdaFuncCount ensures that no value is present for AwsLambdaFuncCount, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAwsLambdaFuncCount() {
	o.AwsLambdaFuncCount.Unset()
}

// GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetAwsLambdaInvocationsSum() int64 {
	if o == nil || o.AwsLambdaInvocationsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaInvocationsSum.Get()
}

// GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetAwsLambdaInvocationsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsLambdaInvocationsSum.Get(), o.AwsLambdaInvocationsSum.IsSet()
}

// HasAwsLambdaInvocationsSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAwsLambdaInvocationsSum() bool {
	return o != nil && o.AwsLambdaInvocationsSum.IsSet()
}

// SetAwsLambdaInvocationsSum gets a reference to the given datadog.NullableInt64 and assigns it to the AwsLambdaInvocationsSum field.
func (o *UsageSummaryResponse) SetAwsLambdaInvocationsSum(v int64) {
	o.AwsLambdaInvocationsSum.Set(&v)
}

// SetAwsLambdaInvocationsSumNil sets the value for AwsLambdaInvocationsSum to be an explicit nil.
func (o *UsageSummaryResponse) SetAwsLambdaInvocationsSumNil() {
	o.AwsLambdaInvocationsSum.Set(nil)
}

// UnsetAwsLambdaInvocationsSum ensures that no value is present for AwsLambdaInvocationsSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAwsLambdaInvocationsSum() {
	o.AwsLambdaInvocationsSum.Unset()
}

// GetAzureAppServiceTop99pSum returns the AzureAppServiceTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetAzureAppServiceTop99pSum() int64 {
	if o == nil || o.AzureAppServiceTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AzureAppServiceTop99pSum.Get()
}

// GetAzureAppServiceTop99pSumOk returns a tuple with the AzureAppServiceTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetAzureAppServiceTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureAppServiceTop99pSum.Get(), o.AzureAppServiceTop99pSum.IsSet()
}

// HasAzureAppServiceTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAzureAppServiceTop99pSum() bool {
	return o != nil && o.AzureAppServiceTop99pSum.IsSet()
}

// SetAzureAppServiceTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the AzureAppServiceTop99pSum field.
func (o *UsageSummaryResponse) SetAzureAppServiceTop99pSum(v int64) {
	o.AzureAppServiceTop99pSum.Set(&v)
}

// SetAzureAppServiceTop99pSumNil sets the value for AzureAppServiceTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetAzureAppServiceTop99pSumNil() {
	o.AzureAppServiceTop99pSum.Set(nil)
}

// UnsetAzureAppServiceTop99pSum ensures that no value is present for AzureAppServiceTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAzureAppServiceTop99pSum() {
	o.AzureAppServiceTop99pSum.Unset()
}

// GetAzureHostTop99pSum returns the AzureHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetAzureHostTop99pSum() int64 {
	if o == nil || o.AzureHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AzureHostTop99pSum.Get()
}

// GetAzureHostTop99pSumOk returns a tuple with the AzureHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetAzureHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureHostTop99pSum.Get(), o.AzureHostTop99pSum.IsSet()
}

// HasAzureHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAzureHostTop99pSum() bool {
	return o != nil && o.AzureHostTop99pSum.IsSet()
}

// SetAzureHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the AzureHostTop99pSum field.
func (o *UsageSummaryResponse) SetAzureHostTop99pSum(v int64) {
	o.AzureHostTop99pSum.Set(&v)
}

// SetAzureHostTop99pSumNil sets the value for AzureHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetAzureHostTop99pSumNil() {
	o.AzureHostTop99pSum.Set(nil)
}

// UnsetAzureHostTop99pSum ensures that no value is present for AzureHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetAzureHostTop99pSum() {
	o.AzureHostTop99pSum.Unset()
}

// GetBillableIngestedBytesAggSum returns the BillableIngestedBytesAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetBillableIngestedBytesAggSum() int64 {
	if o == nil || o.BillableIngestedBytesAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BillableIngestedBytesAggSum.Get()
}

// GetBillableIngestedBytesAggSumOk returns a tuple with the BillableIngestedBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetBillableIngestedBytesAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillableIngestedBytesAggSum.Get(), o.BillableIngestedBytesAggSum.IsSet()
}

// HasBillableIngestedBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBillableIngestedBytesAggSum() bool {
	return o != nil && o.BillableIngestedBytesAggSum.IsSet()
}

// SetBillableIngestedBytesAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the BillableIngestedBytesAggSum field.
func (o *UsageSummaryResponse) SetBillableIngestedBytesAggSum(v int64) {
	o.BillableIngestedBytesAggSum.Set(&v)
}

// SetBillableIngestedBytesAggSumNil sets the value for BillableIngestedBytesAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetBillableIngestedBytesAggSumNil() {
	o.BillableIngestedBytesAggSum.Set(nil)
}

// UnsetBillableIngestedBytesAggSum ensures that no value is present for BillableIngestedBytesAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetBillableIngestedBytesAggSum() {
	o.BillableIngestedBytesAggSum.Unset()
}

// GetBrowserRumLiteSessionCountAggSum returns the BrowserRumLiteSessionCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetBrowserRumLiteSessionCountAggSum() int64 {
	if o == nil || o.BrowserRumLiteSessionCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumLiteSessionCountAggSum.Get()
}

// GetBrowserRumLiteSessionCountAggSumOk returns a tuple with the BrowserRumLiteSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetBrowserRumLiteSessionCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumLiteSessionCountAggSum.Get(), o.BrowserRumLiteSessionCountAggSum.IsSet()
}

// HasBrowserRumLiteSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBrowserRumLiteSessionCountAggSum() bool {
	return o != nil && o.BrowserRumLiteSessionCountAggSum.IsSet()
}

// SetBrowserRumLiteSessionCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumLiteSessionCountAggSum field.
func (o *UsageSummaryResponse) SetBrowserRumLiteSessionCountAggSum(v int64) {
	o.BrowserRumLiteSessionCountAggSum.Set(&v)
}

// SetBrowserRumLiteSessionCountAggSumNil sets the value for BrowserRumLiteSessionCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetBrowserRumLiteSessionCountAggSumNil() {
	o.BrowserRumLiteSessionCountAggSum.Set(nil)
}

// UnsetBrowserRumLiteSessionCountAggSum ensures that no value is present for BrowserRumLiteSessionCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetBrowserRumLiteSessionCountAggSum() {
	o.BrowserRumLiteSessionCountAggSum.Unset()
}

// GetBrowserRumReplaySessionCountAggSum returns the BrowserRumReplaySessionCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetBrowserRumReplaySessionCountAggSum() int64 {
	if o == nil || o.BrowserRumReplaySessionCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumReplaySessionCountAggSum.Get()
}

// GetBrowserRumReplaySessionCountAggSumOk returns a tuple with the BrowserRumReplaySessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetBrowserRumReplaySessionCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumReplaySessionCountAggSum.Get(), o.BrowserRumReplaySessionCountAggSum.IsSet()
}

// HasBrowserRumReplaySessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBrowserRumReplaySessionCountAggSum() bool {
	return o != nil && o.BrowserRumReplaySessionCountAggSum.IsSet()
}

// SetBrowserRumReplaySessionCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumReplaySessionCountAggSum field.
func (o *UsageSummaryResponse) SetBrowserRumReplaySessionCountAggSum(v int64) {
	o.BrowserRumReplaySessionCountAggSum.Set(&v)
}

// SetBrowserRumReplaySessionCountAggSumNil sets the value for BrowserRumReplaySessionCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetBrowserRumReplaySessionCountAggSumNil() {
	o.BrowserRumReplaySessionCountAggSum.Set(nil)
}

// UnsetBrowserRumReplaySessionCountAggSum ensures that no value is present for BrowserRumReplaySessionCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetBrowserRumReplaySessionCountAggSum() {
	o.BrowserRumReplaySessionCountAggSum.Unset()
}

// GetBrowserRumUnitsAggSum returns the BrowserRumUnitsAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetBrowserRumUnitsAggSum() int64 {
	if o == nil || o.BrowserRumUnitsAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumUnitsAggSum.Get()
}

// GetBrowserRumUnitsAggSumOk returns a tuple with the BrowserRumUnitsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetBrowserRumUnitsAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumUnitsAggSum.Get(), o.BrowserRumUnitsAggSum.IsSet()
}

// HasBrowserRumUnitsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBrowserRumUnitsAggSum() bool {
	return o != nil && o.BrowserRumUnitsAggSum.IsSet()
}

// SetBrowserRumUnitsAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumUnitsAggSum field.
func (o *UsageSummaryResponse) SetBrowserRumUnitsAggSum(v int64) {
	o.BrowserRumUnitsAggSum.Set(&v)
}

// SetBrowserRumUnitsAggSumNil sets the value for BrowserRumUnitsAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetBrowserRumUnitsAggSumNil() {
	o.BrowserRumUnitsAggSum.Set(nil)
}

// UnsetBrowserRumUnitsAggSum ensures that no value is present for BrowserRumUnitsAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetBrowserRumUnitsAggSum() {
	o.BrowserRumUnitsAggSum.Unset()
}

// GetCiPipelineIndexedSpansAggSum returns the CiPipelineIndexedSpansAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCiPipelineIndexedSpansAggSum() int64 {
	if o == nil || o.CiPipelineIndexedSpansAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiPipelineIndexedSpansAggSum.Get()
}

// GetCiPipelineIndexedSpansAggSumOk returns a tuple with the CiPipelineIndexedSpansAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCiPipelineIndexedSpansAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiPipelineIndexedSpansAggSum.Get(), o.CiPipelineIndexedSpansAggSum.IsSet()
}

// HasCiPipelineIndexedSpansAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCiPipelineIndexedSpansAggSum() bool {
	return o != nil && o.CiPipelineIndexedSpansAggSum.IsSet()
}

// SetCiPipelineIndexedSpansAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the CiPipelineIndexedSpansAggSum field.
func (o *UsageSummaryResponse) SetCiPipelineIndexedSpansAggSum(v int64) {
	o.CiPipelineIndexedSpansAggSum.Set(&v)
}

// SetCiPipelineIndexedSpansAggSumNil sets the value for CiPipelineIndexedSpansAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCiPipelineIndexedSpansAggSumNil() {
	o.CiPipelineIndexedSpansAggSum.Set(nil)
}

// UnsetCiPipelineIndexedSpansAggSum ensures that no value is present for CiPipelineIndexedSpansAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCiPipelineIndexedSpansAggSum() {
	o.CiPipelineIndexedSpansAggSum.Unset()
}

// GetCiTestIndexedSpansAggSum returns the CiTestIndexedSpansAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCiTestIndexedSpansAggSum() int64 {
	if o == nil || o.CiTestIndexedSpansAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiTestIndexedSpansAggSum.Get()
}

// GetCiTestIndexedSpansAggSumOk returns a tuple with the CiTestIndexedSpansAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCiTestIndexedSpansAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiTestIndexedSpansAggSum.Get(), o.CiTestIndexedSpansAggSum.IsSet()
}

// HasCiTestIndexedSpansAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCiTestIndexedSpansAggSum() bool {
	return o != nil && o.CiTestIndexedSpansAggSum.IsSet()
}

// SetCiTestIndexedSpansAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the CiTestIndexedSpansAggSum field.
func (o *UsageSummaryResponse) SetCiTestIndexedSpansAggSum(v int64) {
	o.CiTestIndexedSpansAggSum.Set(&v)
}

// SetCiTestIndexedSpansAggSumNil sets the value for CiTestIndexedSpansAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCiTestIndexedSpansAggSumNil() {
	o.CiTestIndexedSpansAggSum.Set(nil)
}

// UnsetCiTestIndexedSpansAggSum ensures that no value is present for CiTestIndexedSpansAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCiTestIndexedSpansAggSum() {
	o.CiTestIndexedSpansAggSum.Unset()
}

// GetCiVisibilityPipelineCommittersHwmSum returns the CiVisibilityPipelineCommittersHwmSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCiVisibilityPipelineCommittersHwmSum() int64 {
	if o == nil || o.CiVisibilityPipelineCommittersHwmSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityPipelineCommittersHwmSum.Get()
}

// GetCiVisibilityPipelineCommittersHwmSumOk returns a tuple with the CiVisibilityPipelineCommittersHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCiVisibilityPipelineCommittersHwmSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiVisibilityPipelineCommittersHwmSum.Get(), o.CiVisibilityPipelineCommittersHwmSum.IsSet()
}

// HasCiVisibilityPipelineCommittersHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCiVisibilityPipelineCommittersHwmSum() bool {
	return o != nil && o.CiVisibilityPipelineCommittersHwmSum.IsSet()
}

// SetCiVisibilityPipelineCommittersHwmSum gets a reference to the given datadog.NullableInt64 and assigns it to the CiVisibilityPipelineCommittersHwmSum field.
func (o *UsageSummaryResponse) SetCiVisibilityPipelineCommittersHwmSum(v int64) {
	o.CiVisibilityPipelineCommittersHwmSum.Set(&v)
}

// SetCiVisibilityPipelineCommittersHwmSumNil sets the value for CiVisibilityPipelineCommittersHwmSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCiVisibilityPipelineCommittersHwmSumNil() {
	o.CiVisibilityPipelineCommittersHwmSum.Set(nil)
}

// UnsetCiVisibilityPipelineCommittersHwmSum ensures that no value is present for CiVisibilityPipelineCommittersHwmSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCiVisibilityPipelineCommittersHwmSum() {
	o.CiVisibilityPipelineCommittersHwmSum.Unset()
}

// GetCiVisibilityTestCommittersHwmSum returns the CiVisibilityTestCommittersHwmSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCiVisibilityTestCommittersHwmSum() int64 {
	if o == nil || o.CiVisibilityTestCommittersHwmSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityTestCommittersHwmSum.Get()
}

// GetCiVisibilityTestCommittersHwmSumOk returns a tuple with the CiVisibilityTestCommittersHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCiVisibilityTestCommittersHwmSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiVisibilityTestCommittersHwmSum.Get(), o.CiVisibilityTestCommittersHwmSum.IsSet()
}

// HasCiVisibilityTestCommittersHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCiVisibilityTestCommittersHwmSum() bool {
	return o != nil && o.CiVisibilityTestCommittersHwmSum.IsSet()
}

// SetCiVisibilityTestCommittersHwmSum gets a reference to the given datadog.NullableInt64 and assigns it to the CiVisibilityTestCommittersHwmSum field.
func (o *UsageSummaryResponse) SetCiVisibilityTestCommittersHwmSum(v int64) {
	o.CiVisibilityTestCommittersHwmSum.Set(&v)
}

// SetCiVisibilityTestCommittersHwmSumNil sets the value for CiVisibilityTestCommittersHwmSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCiVisibilityTestCommittersHwmSumNil() {
	o.CiVisibilityTestCommittersHwmSum.Set(nil)
}

// UnsetCiVisibilityTestCommittersHwmSum ensures that no value is present for CiVisibilityTestCommittersHwmSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCiVisibilityTestCommittersHwmSum() {
	o.CiVisibilityTestCommittersHwmSum.Unset()
}

// GetCloudCostManagementHostCountAvgSum returns the CloudCostManagementHostCountAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCloudCostManagementHostCountAvgSum() int64 {
	if o == nil || o.CloudCostManagementHostCountAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementHostCountAvgSum.Get()
}

// GetCloudCostManagementHostCountAvgSumOk returns a tuple with the CloudCostManagementHostCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCloudCostManagementHostCountAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudCostManagementHostCountAvgSum.Get(), o.CloudCostManagementHostCountAvgSum.IsSet()
}

// HasCloudCostManagementHostCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCloudCostManagementHostCountAvgSum() bool {
	return o != nil && o.CloudCostManagementHostCountAvgSum.IsSet()
}

// SetCloudCostManagementHostCountAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the CloudCostManagementHostCountAvgSum field.
func (o *UsageSummaryResponse) SetCloudCostManagementHostCountAvgSum(v int64) {
	o.CloudCostManagementHostCountAvgSum.Set(&v)
}

// SetCloudCostManagementHostCountAvgSumNil sets the value for CloudCostManagementHostCountAvgSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCloudCostManagementHostCountAvgSumNil() {
	o.CloudCostManagementHostCountAvgSum.Set(nil)
}

// UnsetCloudCostManagementHostCountAvgSum ensures that no value is present for CloudCostManagementHostCountAvgSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCloudCostManagementHostCountAvgSum() {
	o.CloudCostManagementHostCountAvgSum.Unset()
}

// GetContainerAvgSum returns the ContainerAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetContainerAvgSum() int64 {
	if o == nil || o.ContainerAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerAvgSum.Get()
}

// GetContainerAvgSumOk returns a tuple with the ContainerAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetContainerAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerAvgSum.Get(), o.ContainerAvgSum.IsSet()
}

// HasContainerAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasContainerAvgSum() bool {
	return o != nil && o.ContainerAvgSum.IsSet()
}

// SetContainerAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerAvgSum field.
func (o *UsageSummaryResponse) SetContainerAvgSum(v int64) {
	o.ContainerAvgSum.Set(&v)
}

// SetContainerAvgSumNil sets the value for ContainerAvgSum to be an explicit nil.
func (o *UsageSummaryResponse) SetContainerAvgSumNil() {
	o.ContainerAvgSum.Set(nil)
}

// UnsetContainerAvgSum ensures that no value is present for ContainerAvgSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetContainerAvgSum() {
	o.ContainerAvgSum.Unset()
}

// GetContainerExclAgentAvgSum returns the ContainerExclAgentAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetContainerExclAgentAvgSum() int64 {
	if o == nil || o.ContainerExclAgentAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerExclAgentAvgSum.Get()
}

// GetContainerExclAgentAvgSumOk returns a tuple with the ContainerExclAgentAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetContainerExclAgentAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerExclAgentAvgSum.Get(), o.ContainerExclAgentAvgSum.IsSet()
}

// HasContainerExclAgentAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasContainerExclAgentAvgSum() bool {
	return o != nil && o.ContainerExclAgentAvgSum.IsSet()
}

// SetContainerExclAgentAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerExclAgentAvgSum field.
func (o *UsageSummaryResponse) SetContainerExclAgentAvgSum(v int64) {
	o.ContainerExclAgentAvgSum.Set(&v)
}

// SetContainerExclAgentAvgSumNil sets the value for ContainerExclAgentAvgSum to be an explicit nil.
func (o *UsageSummaryResponse) SetContainerExclAgentAvgSumNil() {
	o.ContainerExclAgentAvgSum.Set(nil)
}

// UnsetContainerExclAgentAvgSum ensures that no value is present for ContainerExclAgentAvgSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetContainerExclAgentAvgSum() {
	o.ContainerExclAgentAvgSum.Unset()
}

// GetContainerHwmSum returns the ContainerHwmSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetContainerHwmSum() int64 {
	if o == nil || o.ContainerHwmSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerHwmSum.Get()
}

// GetContainerHwmSumOk returns a tuple with the ContainerHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetContainerHwmSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerHwmSum.Get(), o.ContainerHwmSum.IsSet()
}

// HasContainerHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasContainerHwmSum() bool {
	return o != nil && o.ContainerHwmSum.IsSet()
}

// SetContainerHwmSum gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerHwmSum field.
func (o *UsageSummaryResponse) SetContainerHwmSum(v int64) {
	o.ContainerHwmSum.Set(&v)
}

// SetContainerHwmSumNil sets the value for ContainerHwmSum to be an explicit nil.
func (o *UsageSummaryResponse) SetContainerHwmSumNil() {
	o.ContainerHwmSum.Set(nil)
}

// UnsetContainerHwmSum ensures that no value is present for ContainerHwmSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetContainerHwmSum() {
	o.ContainerHwmSum.Unset()
}

// GetCspmAasHostTop99pSum returns the CspmAasHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCspmAasHostTop99pSum() int64 {
	if o == nil || o.CspmAasHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmAasHostTop99pSum.Get()
}

// GetCspmAasHostTop99pSumOk returns a tuple with the CspmAasHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCspmAasHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmAasHostTop99pSum.Get(), o.CspmAasHostTop99pSum.IsSet()
}

// HasCspmAasHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmAasHostTop99pSum() bool {
	return o != nil && o.CspmAasHostTop99pSum.IsSet()
}

// SetCspmAasHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the CspmAasHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmAasHostTop99pSum(v int64) {
	o.CspmAasHostTop99pSum.Set(&v)
}

// SetCspmAasHostTop99pSumNil sets the value for CspmAasHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCspmAasHostTop99pSumNil() {
	o.CspmAasHostTop99pSum.Set(nil)
}

// UnsetCspmAasHostTop99pSum ensures that no value is present for CspmAasHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCspmAasHostTop99pSum() {
	o.CspmAasHostTop99pSum.Unset()
}

// GetCspmAwsHostTop99pSum returns the CspmAwsHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCspmAwsHostTop99pSum() int64 {
	if o == nil || o.CspmAwsHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmAwsHostTop99pSum.Get()
}

// GetCspmAwsHostTop99pSumOk returns a tuple with the CspmAwsHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCspmAwsHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmAwsHostTop99pSum.Get(), o.CspmAwsHostTop99pSum.IsSet()
}

// HasCspmAwsHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmAwsHostTop99pSum() bool {
	return o != nil && o.CspmAwsHostTop99pSum.IsSet()
}

// SetCspmAwsHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the CspmAwsHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmAwsHostTop99pSum(v int64) {
	o.CspmAwsHostTop99pSum.Set(&v)
}

// SetCspmAwsHostTop99pSumNil sets the value for CspmAwsHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCspmAwsHostTop99pSumNil() {
	o.CspmAwsHostTop99pSum.Set(nil)
}

// UnsetCspmAwsHostTop99pSum ensures that no value is present for CspmAwsHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCspmAwsHostTop99pSum() {
	o.CspmAwsHostTop99pSum.Unset()
}

// GetCspmAzureHostTop99pSum returns the CspmAzureHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCspmAzureHostTop99pSum() int64 {
	if o == nil || o.CspmAzureHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmAzureHostTop99pSum.Get()
}

// GetCspmAzureHostTop99pSumOk returns a tuple with the CspmAzureHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCspmAzureHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmAzureHostTop99pSum.Get(), o.CspmAzureHostTop99pSum.IsSet()
}

// HasCspmAzureHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmAzureHostTop99pSum() bool {
	return o != nil && o.CspmAzureHostTop99pSum.IsSet()
}

// SetCspmAzureHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the CspmAzureHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmAzureHostTop99pSum(v int64) {
	o.CspmAzureHostTop99pSum.Set(&v)
}

// SetCspmAzureHostTop99pSumNil sets the value for CspmAzureHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCspmAzureHostTop99pSumNil() {
	o.CspmAzureHostTop99pSum.Set(nil)
}

// UnsetCspmAzureHostTop99pSum ensures that no value is present for CspmAzureHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCspmAzureHostTop99pSum() {
	o.CspmAzureHostTop99pSum.Unset()
}

// GetCspmContainerAvgSum returns the CspmContainerAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCspmContainerAvgSum() int64 {
	if o == nil || o.CspmContainerAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerAvgSum.Get()
}

// GetCspmContainerAvgSumOk returns a tuple with the CspmContainerAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCspmContainerAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainerAvgSum.Get(), o.CspmContainerAvgSum.IsSet()
}

// HasCspmContainerAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmContainerAvgSum() bool {
	return o != nil && o.CspmContainerAvgSum.IsSet()
}

// SetCspmContainerAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the CspmContainerAvgSum field.
func (o *UsageSummaryResponse) SetCspmContainerAvgSum(v int64) {
	o.CspmContainerAvgSum.Set(&v)
}

// SetCspmContainerAvgSumNil sets the value for CspmContainerAvgSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCspmContainerAvgSumNil() {
	o.CspmContainerAvgSum.Set(nil)
}

// UnsetCspmContainerAvgSum ensures that no value is present for CspmContainerAvgSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCspmContainerAvgSum() {
	o.CspmContainerAvgSum.Unset()
}

// GetCspmContainerHwmSum returns the CspmContainerHwmSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCspmContainerHwmSum() int64 {
	if o == nil || o.CspmContainerHwmSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerHwmSum.Get()
}

// GetCspmContainerHwmSumOk returns a tuple with the CspmContainerHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCspmContainerHwmSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainerHwmSum.Get(), o.CspmContainerHwmSum.IsSet()
}

// HasCspmContainerHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmContainerHwmSum() bool {
	return o != nil && o.CspmContainerHwmSum.IsSet()
}

// SetCspmContainerHwmSum gets a reference to the given datadog.NullableInt64 and assigns it to the CspmContainerHwmSum field.
func (o *UsageSummaryResponse) SetCspmContainerHwmSum(v int64) {
	o.CspmContainerHwmSum.Set(&v)
}

// SetCspmContainerHwmSumNil sets the value for CspmContainerHwmSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCspmContainerHwmSumNil() {
	o.CspmContainerHwmSum.Set(nil)
}

// UnsetCspmContainerHwmSum ensures that no value is present for CspmContainerHwmSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCspmContainerHwmSum() {
	o.CspmContainerHwmSum.Unset()
}

// GetCspmGcpHostTop99pSum returns the CspmGcpHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCspmGcpHostTop99pSum() int64 {
	if o == nil || o.CspmGcpHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmGcpHostTop99pSum.Get()
}

// GetCspmGcpHostTop99pSumOk returns a tuple with the CspmGcpHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCspmGcpHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmGcpHostTop99pSum.Get(), o.CspmGcpHostTop99pSum.IsSet()
}

// HasCspmGcpHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmGcpHostTop99pSum() bool {
	return o != nil && o.CspmGcpHostTop99pSum.IsSet()
}

// SetCspmGcpHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the CspmGcpHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmGcpHostTop99pSum(v int64) {
	o.CspmGcpHostTop99pSum.Set(&v)
}

// SetCspmGcpHostTop99pSumNil sets the value for CspmGcpHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCspmGcpHostTop99pSumNil() {
	o.CspmGcpHostTop99pSum.Set(nil)
}

// UnsetCspmGcpHostTop99pSum ensures that no value is present for CspmGcpHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCspmGcpHostTop99pSum() {
	o.CspmGcpHostTop99pSum.Unset()
}

// GetCspmHostTop99pSum returns the CspmHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCspmHostTop99pSum() int64 {
	if o == nil || o.CspmHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CspmHostTop99pSum.Get()
}

// GetCspmHostTop99pSumOk returns a tuple with the CspmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCspmHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmHostTop99pSum.Get(), o.CspmHostTop99pSum.IsSet()
}

// HasCspmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmHostTop99pSum() bool {
	return o != nil && o.CspmHostTop99pSum.IsSet()
}

// SetCspmHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the CspmHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmHostTop99pSum(v int64) {
	o.CspmHostTop99pSum.Set(&v)
}

// SetCspmHostTop99pSumNil sets the value for CspmHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCspmHostTop99pSumNil() {
	o.CspmHostTop99pSum.Set(nil)
}

// UnsetCspmHostTop99pSum ensures that no value is present for CspmHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCspmHostTop99pSum() {
	o.CspmHostTop99pSum.Unset()
}

// GetCustomTsSum returns the CustomTsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCustomTsSum() int64 {
	if o == nil || o.CustomTsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CustomTsSum.Get()
}

// GetCustomTsSumOk returns a tuple with the CustomTsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCustomTsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTsSum.Get(), o.CustomTsSum.IsSet()
}

// HasCustomTsSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCustomTsSum() bool {
	return o != nil && o.CustomTsSum.IsSet()
}

// SetCustomTsSum gets a reference to the given datadog.NullableInt64 and assigns it to the CustomTsSum field.
func (o *UsageSummaryResponse) SetCustomTsSum(v int64) {
	o.CustomTsSum.Set(&v)
}

// SetCustomTsSumNil sets the value for CustomTsSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCustomTsSumNil() {
	o.CustomTsSum.Set(nil)
}

// UnsetCustomTsSum ensures that no value is present for CustomTsSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCustomTsSum() {
	o.CustomTsSum.Unset()
}

// GetCwsContainersAvgSum returns the CwsContainersAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCwsContainersAvgSum() int64 {
	if o == nil || o.CwsContainersAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CwsContainersAvgSum.Get()
}

// GetCwsContainersAvgSumOk returns a tuple with the CwsContainersAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCwsContainersAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsContainersAvgSum.Get(), o.CwsContainersAvgSum.IsSet()
}

// HasCwsContainersAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCwsContainersAvgSum() bool {
	return o != nil && o.CwsContainersAvgSum.IsSet()
}

// SetCwsContainersAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the CwsContainersAvgSum field.
func (o *UsageSummaryResponse) SetCwsContainersAvgSum(v int64) {
	o.CwsContainersAvgSum.Set(&v)
}

// SetCwsContainersAvgSumNil sets the value for CwsContainersAvgSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCwsContainersAvgSumNil() {
	o.CwsContainersAvgSum.Set(nil)
}

// UnsetCwsContainersAvgSum ensures that no value is present for CwsContainersAvgSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCwsContainersAvgSum() {
	o.CwsContainersAvgSum.Unset()
}

// GetCwsHostTop99pSum returns the CwsHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetCwsHostTop99pSum() int64 {
	if o == nil || o.CwsHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CwsHostTop99pSum.Get()
}

// GetCwsHostTop99pSumOk returns a tuple with the CwsHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetCwsHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsHostTop99pSum.Get(), o.CwsHostTop99pSum.IsSet()
}

// HasCwsHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCwsHostTop99pSum() bool {
	return o != nil && o.CwsHostTop99pSum.IsSet()
}

// SetCwsHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the CwsHostTop99pSum field.
func (o *UsageSummaryResponse) SetCwsHostTop99pSum(v int64) {
	o.CwsHostTop99pSum.Set(&v)
}

// SetCwsHostTop99pSumNil sets the value for CwsHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetCwsHostTop99pSumNil() {
	o.CwsHostTop99pSum.Set(nil)
}

// UnsetCwsHostTop99pSum ensures that no value is present for CwsHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetCwsHostTop99pSum() {
	o.CwsHostTop99pSum.Unset()
}

// GetDbmHostTop99pSum returns the DbmHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetDbmHostTop99pSum() int64 {
	if o == nil || o.DbmHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DbmHostTop99pSum.Get()
}

// GetDbmHostTop99pSumOk returns a tuple with the DbmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetDbmHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmHostTop99pSum.Get(), o.DbmHostTop99pSum.IsSet()
}

// HasDbmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasDbmHostTop99pSum() bool {
	return o != nil && o.DbmHostTop99pSum.IsSet()
}

// SetDbmHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the DbmHostTop99pSum field.
func (o *UsageSummaryResponse) SetDbmHostTop99pSum(v int64) {
	o.DbmHostTop99pSum.Set(&v)
}

// SetDbmHostTop99pSumNil sets the value for DbmHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetDbmHostTop99pSumNil() {
	o.DbmHostTop99pSum.Set(nil)
}

// UnsetDbmHostTop99pSum ensures that no value is present for DbmHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetDbmHostTop99pSum() {
	o.DbmHostTop99pSum.Unset()
}

// GetDbmQueriesAvgSum returns the DbmQueriesAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetDbmQueriesAvgSum() int64 {
	if o == nil || o.DbmQueriesAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DbmQueriesAvgSum.Get()
}

// GetDbmQueriesAvgSumOk returns a tuple with the DbmQueriesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetDbmQueriesAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmQueriesAvgSum.Get(), o.DbmQueriesAvgSum.IsSet()
}

// HasDbmQueriesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasDbmQueriesAvgSum() bool {
	return o != nil && o.DbmQueriesAvgSum.IsSet()
}

// SetDbmQueriesAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the DbmQueriesAvgSum field.
func (o *UsageSummaryResponse) SetDbmQueriesAvgSum(v int64) {
	o.DbmQueriesAvgSum.Set(&v)
}

// SetDbmQueriesAvgSumNil sets the value for DbmQueriesAvgSum to be an explicit nil.
func (o *UsageSummaryResponse) SetDbmQueriesAvgSumNil() {
	o.DbmQueriesAvgSum.Set(nil)
}

// UnsetDbmQueriesAvgSum ensures that no value is present for DbmQueriesAvgSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetDbmQueriesAvgSum() {
	o.DbmQueriesAvgSum.Unset()
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
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *UsageSummaryResponse) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetFargateTasksCountAvgSum returns the FargateTasksCountAvgSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetFargateTasksCountAvgSum() int64 {
	if o == nil || o.FargateTasksCountAvgSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountAvgSum.Get()
}

// GetFargateTasksCountAvgSumOk returns a tuple with the FargateTasksCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetFargateTasksCountAvgSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FargateTasksCountAvgSum.Get(), o.FargateTasksCountAvgSum.IsSet()
}

// HasFargateTasksCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFargateTasksCountAvgSum() bool {
	return o != nil && o.FargateTasksCountAvgSum.IsSet()
}

// SetFargateTasksCountAvgSum gets a reference to the given datadog.NullableInt64 and assigns it to the FargateTasksCountAvgSum field.
func (o *UsageSummaryResponse) SetFargateTasksCountAvgSum(v int64) {
	o.FargateTasksCountAvgSum.Set(&v)
}

// SetFargateTasksCountAvgSumNil sets the value for FargateTasksCountAvgSum to be an explicit nil.
func (o *UsageSummaryResponse) SetFargateTasksCountAvgSumNil() {
	o.FargateTasksCountAvgSum.Set(nil)
}

// UnsetFargateTasksCountAvgSum ensures that no value is present for FargateTasksCountAvgSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetFargateTasksCountAvgSum() {
	o.FargateTasksCountAvgSum.Unset()
}

// GetFargateTasksCountHwmSum returns the FargateTasksCountHwmSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetFargateTasksCountHwmSum() int64 {
	if o == nil || o.FargateTasksCountHwmSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountHwmSum.Get()
}

// GetFargateTasksCountHwmSumOk returns a tuple with the FargateTasksCountHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetFargateTasksCountHwmSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FargateTasksCountHwmSum.Get(), o.FargateTasksCountHwmSum.IsSet()
}

// HasFargateTasksCountHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFargateTasksCountHwmSum() bool {
	return o != nil && o.FargateTasksCountHwmSum.IsSet()
}

// SetFargateTasksCountHwmSum gets a reference to the given datadog.NullableInt64 and assigns it to the FargateTasksCountHwmSum field.
func (o *UsageSummaryResponse) SetFargateTasksCountHwmSum(v int64) {
	o.FargateTasksCountHwmSum.Set(&v)
}

// SetFargateTasksCountHwmSumNil sets the value for FargateTasksCountHwmSum to be an explicit nil.
func (o *UsageSummaryResponse) SetFargateTasksCountHwmSumNil() {
	o.FargateTasksCountHwmSum.Set(nil)
}

// UnsetFargateTasksCountHwmSum ensures that no value is present for FargateTasksCountHwmSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetFargateTasksCountHwmSum() {
	o.FargateTasksCountHwmSum.Unset()
}

// GetForwardingEventsBytesAggSum returns the ForwardingEventsBytesAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetForwardingEventsBytesAggSum() int64 {
	if o == nil || o.ForwardingEventsBytesAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ForwardingEventsBytesAggSum.Get()
}

// GetForwardingEventsBytesAggSumOk returns a tuple with the ForwardingEventsBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetForwardingEventsBytesAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForwardingEventsBytesAggSum.Get(), o.ForwardingEventsBytesAggSum.IsSet()
}

// HasForwardingEventsBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasForwardingEventsBytesAggSum() bool {
	return o != nil && o.ForwardingEventsBytesAggSum.IsSet()
}

// SetForwardingEventsBytesAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the ForwardingEventsBytesAggSum field.
func (o *UsageSummaryResponse) SetForwardingEventsBytesAggSum(v int64) {
	o.ForwardingEventsBytesAggSum.Set(&v)
}

// SetForwardingEventsBytesAggSumNil sets the value for ForwardingEventsBytesAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetForwardingEventsBytesAggSumNil() {
	o.ForwardingEventsBytesAggSum.Set(nil)
}

// UnsetForwardingEventsBytesAggSum ensures that no value is present for ForwardingEventsBytesAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetForwardingEventsBytesAggSum() {
	o.ForwardingEventsBytesAggSum.Unset()
}

// GetGcpHostTop99pSum returns the GcpHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetGcpHostTop99pSum() int64 {
	if o == nil || o.GcpHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.GcpHostTop99pSum.Get()
}

// GetGcpHostTop99pSumOk returns a tuple with the GcpHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetGcpHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GcpHostTop99pSum.Get(), o.GcpHostTop99pSum.IsSet()
}

// HasGcpHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasGcpHostTop99pSum() bool {
	return o != nil && o.GcpHostTop99pSum.IsSet()
}

// SetGcpHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the GcpHostTop99pSum field.
func (o *UsageSummaryResponse) SetGcpHostTop99pSum(v int64) {
	o.GcpHostTop99pSum.Set(&v)
}

// SetGcpHostTop99pSumNil sets the value for GcpHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetGcpHostTop99pSumNil() {
	o.GcpHostTop99pSum.Set(nil)
}

// UnsetGcpHostTop99pSum ensures that no value is present for GcpHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetGcpHostTop99pSum() {
	o.GcpHostTop99pSum.Unset()
}

// GetHerokuHostTop99pSum returns the HerokuHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetHerokuHostTop99pSum() int64 {
	if o == nil || o.HerokuHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HerokuHostTop99pSum.Get()
}

// GetHerokuHostTop99pSumOk returns a tuple with the HerokuHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetHerokuHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HerokuHostTop99pSum.Get(), o.HerokuHostTop99pSum.IsSet()
}

// HasHerokuHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasHerokuHostTop99pSum() bool {
	return o != nil && o.HerokuHostTop99pSum.IsSet()
}

// SetHerokuHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the HerokuHostTop99pSum field.
func (o *UsageSummaryResponse) SetHerokuHostTop99pSum(v int64) {
	o.HerokuHostTop99pSum.Set(&v)
}

// SetHerokuHostTop99pSumNil sets the value for HerokuHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetHerokuHostTop99pSumNil() {
	o.HerokuHostTop99pSum.Set(nil)
}

// UnsetHerokuHostTop99pSum ensures that no value is present for HerokuHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetHerokuHostTop99pSum() {
	o.HerokuHostTop99pSum.Unset()
}

// GetIncidentManagementMonthlyActiveUsersHwmSum returns the IncidentManagementMonthlyActiveUsersHwmSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetIncidentManagementMonthlyActiveUsersHwmSum() int64 {
	if o == nil || o.IncidentManagementMonthlyActiveUsersHwmSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IncidentManagementMonthlyActiveUsersHwmSum.Get()
}

// GetIncidentManagementMonthlyActiveUsersHwmSumOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetIncidentManagementMonthlyActiveUsersHwmSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncidentManagementMonthlyActiveUsersHwmSum.Get(), o.IncidentManagementMonthlyActiveUsersHwmSum.IsSet()
}

// HasIncidentManagementMonthlyActiveUsersHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIncidentManagementMonthlyActiveUsersHwmSum() bool {
	return o != nil && o.IncidentManagementMonthlyActiveUsersHwmSum.IsSet()
}

// SetIncidentManagementMonthlyActiveUsersHwmSum gets a reference to the given datadog.NullableInt64 and assigns it to the IncidentManagementMonthlyActiveUsersHwmSum field.
func (o *UsageSummaryResponse) SetIncidentManagementMonthlyActiveUsersHwmSum(v int64) {
	o.IncidentManagementMonthlyActiveUsersHwmSum.Set(&v)
}

// SetIncidentManagementMonthlyActiveUsersHwmSumNil sets the value for IncidentManagementMonthlyActiveUsersHwmSum to be an explicit nil.
func (o *UsageSummaryResponse) SetIncidentManagementMonthlyActiveUsersHwmSumNil() {
	o.IncidentManagementMonthlyActiveUsersHwmSum.Set(nil)
}

// UnsetIncidentManagementMonthlyActiveUsersHwmSum ensures that no value is present for IncidentManagementMonthlyActiveUsersHwmSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetIncidentManagementMonthlyActiveUsersHwmSum() {
	o.IncidentManagementMonthlyActiveUsersHwmSum.Unset()
}

// GetIndexedEventsCountAggSum returns the IndexedEventsCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetIndexedEventsCountAggSum() int64 {
	if o == nil || o.IndexedEventsCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCountAggSum.Get()
}

// GetIndexedEventsCountAggSumOk returns a tuple with the IndexedEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetIndexedEventsCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexedEventsCountAggSum.Get(), o.IndexedEventsCountAggSum.IsSet()
}

// HasIndexedEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIndexedEventsCountAggSum() bool {
	return o != nil && o.IndexedEventsCountAggSum.IsSet()
}

// SetIndexedEventsCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the IndexedEventsCountAggSum field.
func (o *UsageSummaryResponse) SetIndexedEventsCountAggSum(v int64) {
	o.IndexedEventsCountAggSum.Set(&v)
}

// SetIndexedEventsCountAggSumNil sets the value for IndexedEventsCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetIndexedEventsCountAggSumNil() {
	o.IndexedEventsCountAggSum.Set(nil)
}

// UnsetIndexedEventsCountAggSum ensures that no value is present for IndexedEventsCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetIndexedEventsCountAggSum() {
	o.IndexedEventsCountAggSum.Unset()
}

// GetInfraHostTop99pSum returns the InfraHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetInfraHostTop99pSum() int64 {
	if o == nil || o.InfraHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostTop99pSum.Get()
}

// GetInfraHostTop99pSumOk returns a tuple with the InfraHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetInfraHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfraHostTop99pSum.Get(), o.InfraHostTop99pSum.IsSet()
}

// HasInfraHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasInfraHostTop99pSum() bool {
	return o != nil && o.InfraHostTop99pSum.IsSet()
}

// SetInfraHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the InfraHostTop99pSum field.
func (o *UsageSummaryResponse) SetInfraHostTop99pSum(v int64) {
	o.InfraHostTop99pSum.Set(&v)
}

// SetInfraHostTop99pSumNil sets the value for InfraHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetInfraHostTop99pSumNil() {
	o.InfraHostTop99pSum.Set(nil)
}

// UnsetInfraHostTop99pSum ensures that no value is present for InfraHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetInfraHostTop99pSum() {
	o.InfraHostTop99pSum.Unset()
}

// GetIngestedEventsBytesAggSum returns the IngestedEventsBytesAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetIngestedEventsBytesAggSum() int64 {
	if o == nil || o.IngestedEventsBytesAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytesAggSum.Get()
}

// GetIngestedEventsBytesAggSumOk returns a tuple with the IngestedEventsBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetIngestedEventsBytesAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IngestedEventsBytesAggSum.Get(), o.IngestedEventsBytesAggSum.IsSet()
}

// HasIngestedEventsBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIngestedEventsBytesAggSum() bool {
	return o != nil && o.IngestedEventsBytesAggSum.IsSet()
}

// SetIngestedEventsBytesAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the IngestedEventsBytesAggSum field.
func (o *UsageSummaryResponse) SetIngestedEventsBytesAggSum(v int64) {
	o.IngestedEventsBytesAggSum.Set(&v)
}

// SetIngestedEventsBytesAggSumNil sets the value for IngestedEventsBytesAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetIngestedEventsBytesAggSumNil() {
	o.IngestedEventsBytesAggSum.Set(nil)
}

// UnsetIngestedEventsBytesAggSum ensures that no value is present for IngestedEventsBytesAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetIngestedEventsBytesAggSum() {
	o.IngestedEventsBytesAggSum.Unset()
}

// GetIotDeviceAggSum returns the IotDeviceAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetIotDeviceAggSum() int64 {
	if o == nil || o.IotDeviceAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceAggSum.Get()
}

// GetIotDeviceAggSumOk returns a tuple with the IotDeviceAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetIotDeviceAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IotDeviceAggSum.Get(), o.IotDeviceAggSum.IsSet()
}

// HasIotDeviceAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIotDeviceAggSum() bool {
	return o != nil && o.IotDeviceAggSum.IsSet()
}

// SetIotDeviceAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the IotDeviceAggSum field.
func (o *UsageSummaryResponse) SetIotDeviceAggSum(v int64) {
	o.IotDeviceAggSum.Set(&v)
}

// SetIotDeviceAggSumNil sets the value for IotDeviceAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetIotDeviceAggSumNil() {
	o.IotDeviceAggSum.Set(nil)
}

// UnsetIotDeviceAggSum ensures that no value is present for IotDeviceAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetIotDeviceAggSum() {
	o.IotDeviceAggSum.Unset()
}

// GetIotDeviceTop99pSum returns the IotDeviceTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetIotDeviceTop99pSum() int64 {
	if o == nil || o.IotDeviceTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceTop99pSum.Get()
}

// GetIotDeviceTop99pSumOk returns a tuple with the IotDeviceTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetIotDeviceTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IotDeviceTop99pSum.Get(), o.IotDeviceTop99pSum.IsSet()
}

// HasIotDeviceTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIotDeviceTop99pSum() bool {
	return o != nil && o.IotDeviceTop99pSum.IsSet()
}

// SetIotDeviceTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the IotDeviceTop99pSum field.
func (o *UsageSummaryResponse) SetIotDeviceTop99pSum(v int64) {
	o.IotDeviceTop99pSum.Set(&v)
}

// SetIotDeviceTop99pSumNil sets the value for IotDeviceTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetIotDeviceTop99pSumNil() {
	o.IotDeviceTop99pSum.Set(nil)
}

// UnsetIotDeviceTop99pSum ensures that no value is present for IotDeviceTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetIotDeviceTop99pSum() {
	o.IotDeviceTop99pSum.Unset()
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
	return o != nil && o.LastUpdated != nil
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *UsageSummaryResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLiveIndexedEventsAggSum returns the LiveIndexedEventsAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetLiveIndexedEventsAggSum() int64 {
	if o == nil || o.LiveIndexedEventsAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LiveIndexedEventsAggSum.Get()
}

// GetLiveIndexedEventsAggSumOk returns a tuple with the LiveIndexedEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetLiveIndexedEventsAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveIndexedEventsAggSum.Get(), o.LiveIndexedEventsAggSum.IsSet()
}

// HasLiveIndexedEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLiveIndexedEventsAggSum() bool {
	return o != nil && o.LiveIndexedEventsAggSum.IsSet()
}

// SetLiveIndexedEventsAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the LiveIndexedEventsAggSum field.
func (o *UsageSummaryResponse) SetLiveIndexedEventsAggSum(v int64) {
	o.LiveIndexedEventsAggSum.Set(&v)
}

// SetLiveIndexedEventsAggSumNil sets the value for LiveIndexedEventsAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetLiveIndexedEventsAggSumNil() {
	o.LiveIndexedEventsAggSum.Set(nil)
}

// UnsetLiveIndexedEventsAggSum ensures that no value is present for LiveIndexedEventsAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetLiveIndexedEventsAggSum() {
	o.LiveIndexedEventsAggSum.Unset()
}

// GetLiveIngestedBytesAggSum returns the LiveIngestedBytesAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetLiveIngestedBytesAggSum() int64 {
	if o == nil || o.LiveIngestedBytesAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LiveIngestedBytesAggSum.Get()
}

// GetLiveIngestedBytesAggSumOk returns a tuple with the LiveIngestedBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetLiveIngestedBytesAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveIngestedBytesAggSum.Get(), o.LiveIngestedBytesAggSum.IsSet()
}

// HasLiveIngestedBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLiveIngestedBytesAggSum() bool {
	return o != nil && o.LiveIngestedBytesAggSum.IsSet()
}

// SetLiveIngestedBytesAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the LiveIngestedBytesAggSum field.
func (o *UsageSummaryResponse) SetLiveIngestedBytesAggSum(v int64) {
	o.LiveIngestedBytesAggSum.Set(&v)
}

// SetLiveIngestedBytesAggSumNil sets the value for LiveIngestedBytesAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetLiveIngestedBytesAggSumNil() {
	o.LiveIngestedBytesAggSum.Set(nil)
}

// UnsetLiveIngestedBytesAggSum ensures that no value is present for LiveIngestedBytesAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetLiveIngestedBytesAggSum() {
	o.LiveIngestedBytesAggSum.Unset()
}

// GetLogsByRetention returns the LogsByRetention field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetLogsByRetention() LogsByRetention {
	if o == nil || o.LogsByRetention == nil {
		var ret LogsByRetention
		return ret
	}
	return *o.LogsByRetention
}

// GetLogsByRetentionOk returns a tuple with the LogsByRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetLogsByRetentionOk() (*LogsByRetention, bool) {
	if o == nil || o.LogsByRetention == nil {
		return nil, false
	}
	return o.LogsByRetention, true
}

// HasLogsByRetention returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLogsByRetention() bool {
	return o != nil && o.LogsByRetention != nil
}

// SetLogsByRetention gets a reference to the given LogsByRetention and assigns it to the LogsByRetention field.
func (o *UsageSummaryResponse) SetLogsByRetention(v LogsByRetention) {
	o.LogsByRetention = &v
}

// GetMobileRumLiteSessionCountAggSum returns the MobileRumLiteSessionCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetMobileRumLiteSessionCountAggSum() int64 {
	if o == nil || o.MobileRumLiteSessionCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumLiteSessionCountAggSum.Get()
}

// GetMobileRumLiteSessionCountAggSumOk returns a tuple with the MobileRumLiteSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetMobileRumLiteSessionCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumLiteSessionCountAggSum.Get(), o.MobileRumLiteSessionCountAggSum.IsSet()
}

// HasMobileRumLiteSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumLiteSessionCountAggSum() bool {
	return o != nil && o.MobileRumLiteSessionCountAggSum.IsSet()
}

// SetMobileRumLiteSessionCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumLiteSessionCountAggSum field.
func (o *UsageSummaryResponse) SetMobileRumLiteSessionCountAggSum(v int64) {
	o.MobileRumLiteSessionCountAggSum.Set(&v)
}

// SetMobileRumLiteSessionCountAggSumNil sets the value for MobileRumLiteSessionCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetMobileRumLiteSessionCountAggSumNil() {
	o.MobileRumLiteSessionCountAggSum.Set(nil)
}

// UnsetMobileRumLiteSessionCountAggSum ensures that no value is present for MobileRumLiteSessionCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetMobileRumLiteSessionCountAggSum() {
	o.MobileRumLiteSessionCountAggSum.Unset()
}

// GetMobileRumSessionCountAggSum returns the MobileRumSessionCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetMobileRumSessionCountAggSum() int64 {
	if o == nil || o.MobileRumSessionCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountAggSum.Get()
}

// GetMobileRumSessionCountAggSumOk returns a tuple with the MobileRumSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetMobileRumSessionCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountAggSum.Get(), o.MobileRumSessionCountAggSum.IsSet()
}

// HasMobileRumSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountAggSum() bool {
	return o != nil && o.MobileRumSessionCountAggSum.IsSet()
}

// SetMobileRumSessionCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountAggSum field.
func (o *UsageSummaryResponse) SetMobileRumSessionCountAggSum(v int64) {
	o.MobileRumSessionCountAggSum.Set(&v)
}

// SetMobileRumSessionCountAggSumNil sets the value for MobileRumSessionCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetMobileRumSessionCountAggSumNil() {
	o.MobileRumSessionCountAggSum.Set(nil)
}

// UnsetMobileRumSessionCountAggSum ensures that no value is present for MobileRumSessionCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetMobileRumSessionCountAggSum() {
	o.MobileRumSessionCountAggSum.Unset()
}

// GetMobileRumSessionCountAndroidAggSum returns the MobileRumSessionCountAndroidAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetMobileRumSessionCountAndroidAggSum() int64 {
	if o == nil || o.MobileRumSessionCountAndroidAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountAndroidAggSum.Get()
}

// GetMobileRumSessionCountAndroidAggSumOk returns a tuple with the MobileRumSessionCountAndroidAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetMobileRumSessionCountAndroidAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountAndroidAggSum.Get(), o.MobileRumSessionCountAndroidAggSum.IsSet()
}

// HasMobileRumSessionCountAndroidAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountAndroidAggSum() bool {
	return o != nil && o.MobileRumSessionCountAndroidAggSum.IsSet()
}

// SetMobileRumSessionCountAndroidAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountAndroidAggSum field.
func (o *UsageSummaryResponse) SetMobileRumSessionCountAndroidAggSum(v int64) {
	o.MobileRumSessionCountAndroidAggSum.Set(&v)
}

// SetMobileRumSessionCountAndroidAggSumNil sets the value for MobileRumSessionCountAndroidAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetMobileRumSessionCountAndroidAggSumNil() {
	o.MobileRumSessionCountAndroidAggSum.Set(nil)
}

// UnsetMobileRumSessionCountAndroidAggSum ensures that no value is present for MobileRumSessionCountAndroidAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetMobileRumSessionCountAndroidAggSum() {
	o.MobileRumSessionCountAndroidAggSum.Unset()
}

// GetMobileRumSessionCountFlutterAggSum returns the MobileRumSessionCountFlutterAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetMobileRumSessionCountFlutterAggSum() int64 {
	if o == nil || o.MobileRumSessionCountFlutterAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountFlutterAggSum.Get()
}

// GetMobileRumSessionCountFlutterAggSumOk returns a tuple with the MobileRumSessionCountFlutterAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetMobileRumSessionCountFlutterAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountFlutterAggSum.Get(), o.MobileRumSessionCountFlutterAggSum.IsSet()
}

// HasMobileRumSessionCountFlutterAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountFlutterAggSum() bool {
	return o != nil && o.MobileRumSessionCountFlutterAggSum.IsSet()
}

// SetMobileRumSessionCountFlutterAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountFlutterAggSum field.
func (o *UsageSummaryResponse) SetMobileRumSessionCountFlutterAggSum(v int64) {
	o.MobileRumSessionCountFlutterAggSum.Set(&v)
}

// SetMobileRumSessionCountFlutterAggSumNil sets the value for MobileRumSessionCountFlutterAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetMobileRumSessionCountFlutterAggSumNil() {
	o.MobileRumSessionCountFlutterAggSum.Set(nil)
}

// UnsetMobileRumSessionCountFlutterAggSum ensures that no value is present for MobileRumSessionCountFlutterAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetMobileRumSessionCountFlutterAggSum() {
	o.MobileRumSessionCountFlutterAggSum.Unset()
}

// GetMobileRumSessionCountIosAggSum returns the MobileRumSessionCountIosAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetMobileRumSessionCountIosAggSum() int64 {
	if o == nil || o.MobileRumSessionCountIosAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountIosAggSum.Get()
}

// GetMobileRumSessionCountIosAggSumOk returns a tuple with the MobileRumSessionCountIosAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetMobileRumSessionCountIosAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountIosAggSum.Get(), o.MobileRumSessionCountIosAggSum.IsSet()
}

// HasMobileRumSessionCountIosAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountIosAggSum() bool {
	return o != nil && o.MobileRumSessionCountIosAggSum.IsSet()
}

// SetMobileRumSessionCountIosAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountIosAggSum field.
func (o *UsageSummaryResponse) SetMobileRumSessionCountIosAggSum(v int64) {
	o.MobileRumSessionCountIosAggSum.Set(&v)
}

// SetMobileRumSessionCountIosAggSumNil sets the value for MobileRumSessionCountIosAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetMobileRumSessionCountIosAggSumNil() {
	o.MobileRumSessionCountIosAggSum.Set(nil)
}

// UnsetMobileRumSessionCountIosAggSum ensures that no value is present for MobileRumSessionCountIosAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetMobileRumSessionCountIosAggSum() {
	o.MobileRumSessionCountIosAggSum.Unset()
}

// GetMobileRumSessionCountReactnativeAggSum returns the MobileRumSessionCountReactnativeAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetMobileRumSessionCountReactnativeAggSum() int64 {
	if o == nil || o.MobileRumSessionCountReactnativeAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountReactnativeAggSum.Get()
}

// GetMobileRumSessionCountReactnativeAggSumOk returns a tuple with the MobileRumSessionCountReactnativeAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetMobileRumSessionCountReactnativeAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumSessionCountReactnativeAggSum.Get(), o.MobileRumSessionCountReactnativeAggSum.IsSet()
}

// HasMobileRumSessionCountReactnativeAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountReactnativeAggSum() bool {
	return o != nil && o.MobileRumSessionCountReactnativeAggSum.IsSet()
}

// SetMobileRumSessionCountReactnativeAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumSessionCountReactnativeAggSum field.
func (o *UsageSummaryResponse) SetMobileRumSessionCountReactnativeAggSum(v int64) {
	o.MobileRumSessionCountReactnativeAggSum.Set(&v)
}

// SetMobileRumSessionCountReactnativeAggSumNil sets the value for MobileRumSessionCountReactnativeAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetMobileRumSessionCountReactnativeAggSumNil() {
	o.MobileRumSessionCountReactnativeAggSum.Set(nil)
}

// UnsetMobileRumSessionCountReactnativeAggSum ensures that no value is present for MobileRumSessionCountReactnativeAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetMobileRumSessionCountReactnativeAggSum() {
	o.MobileRumSessionCountReactnativeAggSum.Unset()
}

// GetMobileRumUnitsAggSum returns the MobileRumUnitsAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetMobileRumUnitsAggSum() int64 {
	if o == nil || o.MobileRumUnitsAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumUnitsAggSum.Get()
}

// GetMobileRumUnitsAggSumOk returns a tuple with the MobileRumUnitsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetMobileRumUnitsAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumUnitsAggSum.Get(), o.MobileRumUnitsAggSum.IsSet()
}

// HasMobileRumUnitsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumUnitsAggSum() bool {
	return o != nil && o.MobileRumUnitsAggSum.IsSet()
}

// SetMobileRumUnitsAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumUnitsAggSum field.
func (o *UsageSummaryResponse) SetMobileRumUnitsAggSum(v int64) {
	o.MobileRumUnitsAggSum.Set(&v)
}

// SetMobileRumUnitsAggSumNil sets the value for MobileRumUnitsAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetMobileRumUnitsAggSumNil() {
	o.MobileRumUnitsAggSum.Set(nil)
}

// UnsetMobileRumUnitsAggSum ensures that no value is present for MobileRumUnitsAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetMobileRumUnitsAggSum() {
	o.MobileRumUnitsAggSum.Unset()
}

// GetNetflowIndexedEventsCountAggSum returns the NetflowIndexedEventsCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetNetflowIndexedEventsCountAggSum() int64 {
	if o == nil || o.NetflowIndexedEventsCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NetflowIndexedEventsCountAggSum.Get()
}

// GetNetflowIndexedEventsCountAggSumOk returns a tuple with the NetflowIndexedEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetNetflowIndexedEventsCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetflowIndexedEventsCountAggSum.Get(), o.NetflowIndexedEventsCountAggSum.IsSet()
}

// HasNetflowIndexedEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasNetflowIndexedEventsCountAggSum() bool {
	return o != nil && o.NetflowIndexedEventsCountAggSum.IsSet()
}

// SetNetflowIndexedEventsCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the NetflowIndexedEventsCountAggSum field.
func (o *UsageSummaryResponse) SetNetflowIndexedEventsCountAggSum(v int64) {
	o.NetflowIndexedEventsCountAggSum.Set(&v)
}

// SetNetflowIndexedEventsCountAggSumNil sets the value for NetflowIndexedEventsCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetNetflowIndexedEventsCountAggSumNil() {
	o.NetflowIndexedEventsCountAggSum.Set(nil)
}

// UnsetNetflowIndexedEventsCountAggSum ensures that no value is present for NetflowIndexedEventsCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetNetflowIndexedEventsCountAggSum() {
	o.NetflowIndexedEventsCountAggSum.Unset()
}

// GetNpmHostTop99pSum returns the NpmHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetNpmHostTop99pSum() int64 {
	if o == nil || o.NpmHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NpmHostTop99pSum.Get()
}

// GetNpmHostTop99pSumOk returns a tuple with the NpmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetNpmHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NpmHostTop99pSum.Get(), o.NpmHostTop99pSum.IsSet()
}

// HasNpmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasNpmHostTop99pSum() bool {
	return o != nil && o.NpmHostTop99pSum.IsSet()
}

// SetNpmHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the NpmHostTop99pSum field.
func (o *UsageSummaryResponse) SetNpmHostTop99pSum(v int64) {
	o.NpmHostTop99pSum.Set(&v)
}

// SetNpmHostTop99pSumNil sets the value for NpmHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetNpmHostTop99pSumNil() {
	o.NpmHostTop99pSum.Set(nil)
}

// UnsetNpmHostTop99pSum ensures that no value is present for NpmHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetNpmHostTop99pSum() {
	o.NpmHostTop99pSum.Unset()
}

// GetObservabilityPipelinesBytesProcessedAggSum returns the ObservabilityPipelinesBytesProcessedAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetObservabilityPipelinesBytesProcessedAggSum() int64 {
	if o == nil || o.ObservabilityPipelinesBytesProcessedAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ObservabilityPipelinesBytesProcessedAggSum.Get()
}

// GetObservabilityPipelinesBytesProcessedAggSumOk returns a tuple with the ObservabilityPipelinesBytesProcessedAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetObservabilityPipelinesBytesProcessedAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObservabilityPipelinesBytesProcessedAggSum.Get(), o.ObservabilityPipelinesBytesProcessedAggSum.IsSet()
}

// HasObservabilityPipelinesBytesProcessedAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasObservabilityPipelinesBytesProcessedAggSum() bool {
	return o != nil && o.ObservabilityPipelinesBytesProcessedAggSum.IsSet()
}

// SetObservabilityPipelinesBytesProcessedAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the ObservabilityPipelinesBytesProcessedAggSum field.
func (o *UsageSummaryResponse) SetObservabilityPipelinesBytesProcessedAggSum(v int64) {
	o.ObservabilityPipelinesBytesProcessedAggSum.Set(&v)
}

// SetObservabilityPipelinesBytesProcessedAggSumNil sets the value for ObservabilityPipelinesBytesProcessedAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetObservabilityPipelinesBytesProcessedAggSumNil() {
	o.ObservabilityPipelinesBytesProcessedAggSum.Set(nil)
}

// UnsetObservabilityPipelinesBytesProcessedAggSum ensures that no value is present for ObservabilityPipelinesBytesProcessedAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetObservabilityPipelinesBytesProcessedAggSum() {
	o.ObservabilityPipelinesBytesProcessedAggSum.Unset()
}

// GetOnlineArchiveEventsCountAggSum returns the OnlineArchiveEventsCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetOnlineArchiveEventsCountAggSum() int64 {
	if o == nil || o.OnlineArchiveEventsCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OnlineArchiveEventsCountAggSum.Get()
}

// GetOnlineArchiveEventsCountAggSumOk returns a tuple with the OnlineArchiveEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetOnlineArchiveEventsCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlineArchiveEventsCountAggSum.Get(), o.OnlineArchiveEventsCountAggSum.IsSet()
}

// HasOnlineArchiveEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasOnlineArchiveEventsCountAggSum() bool {
	return o != nil && o.OnlineArchiveEventsCountAggSum.IsSet()
}

// SetOnlineArchiveEventsCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the OnlineArchiveEventsCountAggSum field.
func (o *UsageSummaryResponse) SetOnlineArchiveEventsCountAggSum(v int64) {
	o.OnlineArchiveEventsCountAggSum.Set(&v)
}

// SetOnlineArchiveEventsCountAggSumNil sets the value for OnlineArchiveEventsCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetOnlineArchiveEventsCountAggSumNil() {
	o.OnlineArchiveEventsCountAggSum.Set(nil)
}

// UnsetOnlineArchiveEventsCountAggSum ensures that no value is present for OnlineArchiveEventsCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetOnlineArchiveEventsCountAggSum() {
	o.OnlineArchiveEventsCountAggSum.Unset()
}

// GetOpentelemetryApmHostTop99pSum returns the OpentelemetryApmHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetOpentelemetryApmHostTop99pSum() int64 {
	if o == nil || o.OpentelemetryApmHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryApmHostTop99pSum.Get()
}

// GetOpentelemetryApmHostTop99pSumOk returns a tuple with the OpentelemetryApmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetOpentelemetryApmHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpentelemetryApmHostTop99pSum.Get(), o.OpentelemetryApmHostTop99pSum.IsSet()
}

// HasOpentelemetryApmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasOpentelemetryApmHostTop99pSum() bool {
	return o != nil && o.OpentelemetryApmHostTop99pSum.IsSet()
}

// SetOpentelemetryApmHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the OpentelemetryApmHostTop99pSum field.
func (o *UsageSummaryResponse) SetOpentelemetryApmHostTop99pSum(v int64) {
	o.OpentelemetryApmHostTop99pSum.Set(&v)
}

// SetOpentelemetryApmHostTop99pSumNil sets the value for OpentelemetryApmHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetOpentelemetryApmHostTop99pSumNil() {
	o.OpentelemetryApmHostTop99pSum.Set(nil)
}

// UnsetOpentelemetryApmHostTop99pSum ensures that no value is present for OpentelemetryApmHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetOpentelemetryApmHostTop99pSum() {
	o.OpentelemetryApmHostTop99pSum.Unset()
}

// GetOpentelemetryHostTop99pSum returns the OpentelemetryHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetOpentelemetryHostTop99pSum() int64 {
	if o == nil || o.OpentelemetryHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryHostTop99pSum.Get()
}

// GetOpentelemetryHostTop99pSumOk returns a tuple with the OpentelemetryHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetOpentelemetryHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpentelemetryHostTop99pSum.Get(), o.OpentelemetryHostTop99pSum.IsSet()
}

// HasOpentelemetryHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasOpentelemetryHostTop99pSum() bool {
	return o != nil && o.OpentelemetryHostTop99pSum.IsSet()
}

// SetOpentelemetryHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the OpentelemetryHostTop99pSum field.
func (o *UsageSummaryResponse) SetOpentelemetryHostTop99pSum(v int64) {
	o.OpentelemetryHostTop99pSum.Set(&v)
}

// SetOpentelemetryHostTop99pSumNil sets the value for OpentelemetryHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetOpentelemetryHostTop99pSumNil() {
	o.OpentelemetryHostTop99pSum.Set(nil)
}

// UnsetOpentelemetryHostTop99pSum ensures that no value is present for OpentelemetryHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetOpentelemetryHostTop99pSum() {
	o.OpentelemetryHostTop99pSum.Unset()
}

// GetProfilingContainerAgentCountAvg returns the ProfilingContainerAgentCountAvg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetProfilingContainerAgentCountAvg() int64 {
	if o == nil || o.ProfilingContainerAgentCountAvg.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ProfilingContainerAgentCountAvg.Get()
}

// GetProfilingContainerAgentCountAvgOk returns a tuple with the ProfilingContainerAgentCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetProfilingContainerAgentCountAvgOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilingContainerAgentCountAvg.Get(), o.ProfilingContainerAgentCountAvg.IsSet()
}

// HasProfilingContainerAgentCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasProfilingContainerAgentCountAvg() bool {
	return o != nil && o.ProfilingContainerAgentCountAvg.IsSet()
}

// SetProfilingContainerAgentCountAvg gets a reference to the given datadog.NullableInt64 and assigns it to the ProfilingContainerAgentCountAvg field.
func (o *UsageSummaryResponse) SetProfilingContainerAgentCountAvg(v int64) {
	o.ProfilingContainerAgentCountAvg.Set(&v)
}

// SetProfilingContainerAgentCountAvgNil sets the value for ProfilingContainerAgentCountAvg to be an explicit nil.
func (o *UsageSummaryResponse) SetProfilingContainerAgentCountAvgNil() {
	o.ProfilingContainerAgentCountAvg.Set(nil)
}

// UnsetProfilingContainerAgentCountAvg ensures that no value is present for ProfilingContainerAgentCountAvg, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetProfilingContainerAgentCountAvg() {
	o.ProfilingContainerAgentCountAvg.Unset()
}

// GetProfilingHostCountTop99pSum returns the ProfilingHostCountTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetProfilingHostCountTop99pSum() int64 {
	if o == nil || o.ProfilingHostCountTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ProfilingHostCountTop99pSum.Get()
}

// GetProfilingHostCountTop99pSumOk returns a tuple with the ProfilingHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetProfilingHostCountTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilingHostCountTop99pSum.Get(), o.ProfilingHostCountTop99pSum.IsSet()
}

// HasProfilingHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasProfilingHostCountTop99pSum() bool {
	return o != nil && o.ProfilingHostCountTop99pSum.IsSet()
}

// SetProfilingHostCountTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the ProfilingHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetProfilingHostCountTop99pSum(v int64) {
	o.ProfilingHostCountTop99pSum.Set(&v)
}

// SetProfilingHostCountTop99pSumNil sets the value for ProfilingHostCountTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetProfilingHostCountTop99pSumNil() {
	o.ProfilingHostCountTop99pSum.Set(nil)
}

// UnsetProfilingHostCountTop99pSum ensures that no value is present for ProfilingHostCountTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetProfilingHostCountTop99pSum() {
	o.ProfilingHostCountTop99pSum.Unset()
}

// GetRehydratedIndexedEventsAggSum returns the RehydratedIndexedEventsAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetRehydratedIndexedEventsAggSum() int64 {
	if o == nil || o.RehydratedIndexedEventsAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RehydratedIndexedEventsAggSum.Get()
}

// GetRehydratedIndexedEventsAggSumOk returns a tuple with the RehydratedIndexedEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetRehydratedIndexedEventsAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RehydratedIndexedEventsAggSum.Get(), o.RehydratedIndexedEventsAggSum.IsSet()
}

// HasRehydratedIndexedEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRehydratedIndexedEventsAggSum() bool {
	return o != nil && o.RehydratedIndexedEventsAggSum.IsSet()
}

// SetRehydratedIndexedEventsAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the RehydratedIndexedEventsAggSum field.
func (o *UsageSummaryResponse) SetRehydratedIndexedEventsAggSum(v int64) {
	o.RehydratedIndexedEventsAggSum.Set(&v)
}

// SetRehydratedIndexedEventsAggSumNil sets the value for RehydratedIndexedEventsAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetRehydratedIndexedEventsAggSumNil() {
	o.RehydratedIndexedEventsAggSum.Set(nil)
}

// UnsetRehydratedIndexedEventsAggSum ensures that no value is present for RehydratedIndexedEventsAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetRehydratedIndexedEventsAggSum() {
	o.RehydratedIndexedEventsAggSum.Unset()
}

// GetRehydratedIngestedBytesAggSum returns the RehydratedIngestedBytesAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetRehydratedIngestedBytesAggSum() int64 {
	if o == nil || o.RehydratedIngestedBytesAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RehydratedIngestedBytesAggSum.Get()
}

// GetRehydratedIngestedBytesAggSumOk returns a tuple with the RehydratedIngestedBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetRehydratedIngestedBytesAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RehydratedIngestedBytesAggSum.Get(), o.RehydratedIngestedBytesAggSum.IsSet()
}

// HasRehydratedIngestedBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRehydratedIngestedBytesAggSum() bool {
	return o != nil && o.RehydratedIngestedBytesAggSum.IsSet()
}

// SetRehydratedIngestedBytesAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the RehydratedIngestedBytesAggSum field.
func (o *UsageSummaryResponse) SetRehydratedIngestedBytesAggSum(v int64) {
	o.RehydratedIngestedBytesAggSum.Set(&v)
}

// SetRehydratedIngestedBytesAggSumNil sets the value for RehydratedIngestedBytesAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetRehydratedIngestedBytesAggSumNil() {
	o.RehydratedIngestedBytesAggSum.Set(nil)
}

// UnsetRehydratedIngestedBytesAggSum ensures that no value is present for RehydratedIngestedBytesAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetRehydratedIngestedBytesAggSum() {
	o.RehydratedIngestedBytesAggSum.Unset()
}

// GetRumBrowserAndMobileSessionCount returns the RumBrowserAndMobileSessionCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetRumBrowserAndMobileSessionCount() int64 {
	if o == nil || o.RumBrowserAndMobileSessionCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserAndMobileSessionCount.Get()
}

// GetRumBrowserAndMobileSessionCountOk returns a tuple with the RumBrowserAndMobileSessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetRumBrowserAndMobileSessionCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumBrowserAndMobileSessionCount.Get(), o.RumBrowserAndMobileSessionCount.IsSet()
}

// HasRumBrowserAndMobileSessionCount returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumBrowserAndMobileSessionCount() bool {
	return o != nil && o.RumBrowserAndMobileSessionCount.IsSet()
}

// SetRumBrowserAndMobileSessionCount gets a reference to the given datadog.NullableInt64 and assigns it to the RumBrowserAndMobileSessionCount field.
func (o *UsageSummaryResponse) SetRumBrowserAndMobileSessionCount(v int64) {
	o.RumBrowserAndMobileSessionCount.Set(&v)
}

// SetRumBrowserAndMobileSessionCountNil sets the value for RumBrowserAndMobileSessionCount to be an explicit nil.
func (o *UsageSummaryResponse) SetRumBrowserAndMobileSessionCountNil() {
	o.RumBrowserAndMobileSessionCount.Set(nil)
}

// UnsetRumBrowserAndMobileSessionCount ensures that no value is present for RumBrowserAndMobileSessionCount, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetRumBrowserAndMobileSessionCount() {
	o.RumBrowserAndMobileSessionCount.Unset()
}

// GetRumSessionCountAggSum returns the RumSessionCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetRumSessionCountAggSum() int64 {
	if o == nil || o.RumSessionCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumSessionCountAggSum.Get()
}

// GetRumSessionCountAggSumOk returns a tuple with the RumSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetRumSessionCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumSessionCountAggSum.Get(), o.RumSessionCountAggSum.IsSet()
}

// HasRumSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumSessionCountAggSum() bool {
	return o != nil && o.RumSessionCountAggSum.IsSet()
}

// SetRumSessionCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the RumSessionCountAggSum field.
func (o *UsageSummaryResponse) SetRumSessionCountAggSum(v int64) {
	o.RumSessionCountAggSum.Set(&v)
}

// SetRumSessionCountAggSumNil sets the value for RumSessionCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetRumSessionCountAggSumNil() {
	o.RumSessionCountAggSum.Set(nil)
}

// UnsetRumSessionCountAggSum ensures that no value is present for RumSessionCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetRumSessionCountAggSum() {
	o.RumSessionCountAggSum.Unset()
}

// GetRumTotalSessionCountAggSum returns the RumTotalSessionCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetRumTotalSessionCountAggSum() int64 {
	if o == nil || o.RumTotalSessionCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumTotalSessionCountAggSum.Get()
}

// GetRumTotalSessionCountAggSumOk returns a tuple with the RumTotalSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetRumTotalSessionCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumTotalSessionCountAggSum.Get(), o.RumTotalSessionCountAggSum.IsSet()
}

// HasRumTotalSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumTotalSessionCountAggSum() bool {
	return o != nil && o.RumTotalSessionCountAggSum.IsSet()
}

// SetRumTotalSessionCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the RumTotalSessionCountAggSum field.
func (o *UsageSummaryResponse) SetRumTotalSessionCountAggSum(v int64) {
	o.RumTotalSessionCountAggSum.Set(&v)
}

// SetRumTotalSessionCountAggSumNil sets the value for RumTotalSessionCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetRumTotalSessionCountAggSumNil() {
	o.RumTotalSessionCountAggSum.Set(nil)
}

// UnsetRumTotalSessionCountAggSum ensures that no value is present for RumTotalSessionCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetRumTotalSessionCountAggSum() {
	o.RumTotalSessionCountAggSum.Unset()
}

// GetRumUnitsAggSum returns the RumUnitsAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetRumUnitsAggSum() int64 {
	if o == nil || o.RumUnitsAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumUnitsAggSum.Get()
}

// GetRumUnitsAggSumOk returns a tuple with the RumUnitsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetRumUnitsAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumUnitsAggSum.Get(), o.RumUnitsAggSum.IsSet()
}

// HasRumUnitsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumUnitsAggSum() bool {
	return o != nil && o.RumUnitsAggSum.IsSet()
}

// SetRumUnitsAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the RumUnitsAggSum field.
func (o *UsageSummaryResponse) SetRumUnitsAggSum(v int64) {
	o.RumUnitsAggSum.Set(&v)
}

// SetRumUnitsAggSumNil sets the value for RumUnitsAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetRumUnitsAggSumNil() {
	o.RumUnitsAggSum.Set(nil)
}

// UnsetRumUnitsAggSum ensures that no value is present for RumUnitsAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetRumUnitsAggSum() {
	o.RumUnitsAggSum.Unset()
}

// GetSdsApmScannedBytesSum returns the SdsApmScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetSdsApmScannedBytesSum() int64 {
	if o == nil || o.SdsApmScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsApmScannedBytesSum.Get()
}

// GetSdsApmScannedBytesSumOk returns a tuple with the SdsApmScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetSdsApmScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsApmScannedBytesSum.Get(), o.SdsApmScannedBytesSum.IsSet()
}

// HasSdsApmScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsApmScannedBytesSum() bool {
	return o != nil && o.SdsApmScannedBytesSum.IsSet()
}

// SetSdsApmScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsApmScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsApmScannedBytesSum(v int64) {
	o.SdsApmScannedBytesSum.Set(&v)
}

// SetSdsApmScannedBytesSumNil sets the value for SdsApmScannedBytesSum to be an explicit nil.
func (o *UsageSummaryResponse) SetSdsApmScannedBytesSumNil() {
	o.SdsApmScannedBytesSum.Set(nil)
}

// UnsetSdsApmScannedBytesSum ensures that no value is present for SdsApmScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetSdsApmScannedBytesSum() {
	o.SdsApmScannedBytesSum.Unset()
}

// GetSdsEventsScannedBytesSum returns the SdsEventsScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetSdsEventsScannedBytesSum() int64 {
	if o == nil || o.SdsEventsScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsEventsScannedBytesSum.Get()
}

// GetSdsEventsScannedBytesSumOk returns a tuple with the SdsEventsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetSdsEventsScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsEventsScannedBytesSum.Get(), o.SdsEventsScannedBytesSum.IsSet()
}

// HasSdsEventsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsEventsScannedBytesSum() bool {
	return o != nil && o.SdsEventsScannedBytesSum.IsSet()
}

// SetSdsEventsScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsEventsScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsEventsScannedBytesSum(v int64) {
	o.SdsEventsScannedBytesSum.Set(&v)
}

// SetSdsEventsScannedBytesSumNil sets the value for SdsEventsScannedBytesSum to be an explicit nil.
func (o *UsageSummaryResponse) SetSdsEventsScannedBytesSumNil() {
	o.SdsEventsScannedBytesSum.Set(nil)
}

// UnsetSdsEventsScannedBytesSum ensures that no value is present for SdsEventsScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetSdsEventsScannedBytesSum() {
	o.SdsEventsScannedBytesSum.Unset()
}

// GetSdsLogsScannedBytesSum returns the SdsLogsScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetSdsLogsScannedBytesSum() int64 {
	if o == nil || o.SdsLogsScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsLogsScannedBytesSum.Get()
}

// GetSdsLogsScannedBytesSumOk returns a tuple with the SdsLogsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetSdsLogsScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsLogsScannedBytesSum.Get(), o.SdsLogsScannedBytesSum.IsSet()
}

// HasSdsLogsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsLogsScannedBytesSum() bool {
	return o != nil && o.SdsLogsScannedBytesSum.IsSet()
}

// SetSdsLogsScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsLogsScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsLogsScannedBytesSum(v int64) {
	o.SdsLogsScannedBytesSum.Set(&v)
}

// SetSdsLogsScannedBytesSumNil sets the value for SdsLogsScannedBytesSum to be an explicit nil.
func (o *UsageSummaryResponse) SetSdsLogsScannedBytesSumNil() {
	o.SdsLogsScannedBytesSum.Set(nil)
}

// UnsetSdsLogsScannedBytesSum ensures that no value is present for SdsLogsScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetSdsLogsScannedBytesSum() {
	o.SdsLogsScannedBytesSum.Unset()
}

// GetSdsRumScannedBytesSum returns the SdsRumScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetSdsRumScannedBytesSum() int64 {
	if o == nil || o.SdsRumScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsRumScannedBytesSum.Get()
}

// GetSdsRumScannedBytesSumOk returns a tuple with the SdsRumScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetSdsRumScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsRumScannedBytesSum.Get(), o.SdsRumScannedBytesSum.IsSet()
}

// HasSdsRumScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsRumScannedBytesSum() bool {
	return o != nil && o.SdsRumScannedBytesSum.IsSet()
}

// SetSdsRumScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsRumScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsRumScannedBytesSum(v int64) {
	o.SdsRumScannedBytesSum.Set(&v)
}

// SetSdsRumScannedBytesSumNil sets the value for SdsRumScannedBytesSum to be an explicit nil.
func (o *UsageSummaryResponse) SetSdsRumScannedBytesSumNil() {
	o.SdsRumScannedBytesSum.Set(nil)
}

// UnsetSdsRumScannedBytesSum ensures that no value is present for SdsRumScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetSdsRumScannedBytesSum() {
	o.SdsRumScannedBytesSum.Unset()
}

// GetSdsTotalScannedBytesSum returns the SdsTotalScannedBytesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetSdsTotalScannedBytesSum() int64 {
	if o == nil || o.SdsTotalScannedBytesSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SdsTotalScannedBytesSum.Get()
}

// GetSdsTotalScannedBytesSumOk returns a tuple with the SdsTotalScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetSdsTotalScannedBytesSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdsTotalScannedBytesSum.Get(), o.SdsTotalScannedBytesSum.IsSet()
}

// HasSdsTotalScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsTotalScannedBytesSum() bool {
	return o != nil && o.SdsTotalScannedBytesSum.IsSet()
}

// SetSdsTotalScannedBytesSum gets a reference to the given datadog.NullableInt64 and assigns it to the SdsTotalScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsTotalScannedBytesSum(v int64) {
	o.SdsTotalScannedBytesSum.Set(&v)
}

// SetSdsTotalScannedBytesSumNil sets the value for SdsTotalScannedBytesSum to be an explicit nil.
func (o *UsageSummaryResponse) SetSdsTotalScannedBytesSumNil() {
	o.SdsTotalScannedBytesSum.Set(nil)
}

// UnsetSdsTotalScannedBytesSum ensures that no value is present for SdsTotalScannedBytesSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetSdsTotalScannedBytesSum() {
	o.SdsTotalScannedBytesSum.Unset()
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
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *UsageSummaryResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetSyntheticsBrowserCheckCallsCountAggSum returns the SyntheticsBrowserCheckCallsCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetSyntheticsBrowserCheckCallsCountAggSum() int64 {
	if o == nil || o.SyntheticsBrowserCheckCallsCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsBrowserCheckCallsCountAggSum.Get()
}

// GetSyntheticsBrowserCheckCallsCountAggSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetSyntheticsBrowserCheckCallsCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyntheticsBrowserCheckCallsCountAggSum.Get(), o.SyntheticsBrowserCheckCallsCountAggSum.IsSet()
}

// HasSyntheticsBrowserCheckCallsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSyntheticsBrowserCheckCallsCountAggSum() bool {
	return o != nil && o.SyntheticsBrowserCheckCallsCountAggSum.IsSet()
}

// SetSyntheticsBrowserCheckCallsCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the SyntheticsBrowserCheckCallsCountAggSum field.
func (o *UsageSummaryResponse) SetSyntheticsBrowserCheckCallsCountAggSum(v int64) {
	o.SyntheticsBrowserCheckCallsCountAggSum.Set(&v)
}

// SetSyntheticsBrowserCheckCallsCountAggSumNil sets the value for SyntheticsBrowserCheckCallsCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetSyntheticsBrowserCheckCallsCountAggSumNil() {
	o.SyntheticsBrowserCheckCallsCountAggSum.Set(nil)
}

// UnsetSyntheticsBrowserCheckCallsCountAggSum ensures that no value is present for SyntheticsBrowserCheckCallsCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetSyntheticsBrowserCheckCallsCountAggSum() {
	o.SyntheticsBrowserCheckCallsCountAggSum.Unset()
}

// GetSyntheticsCheckCallsCountAggSum returns the SyntheticsCheckCallsCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetSyntheticsCheckCallsCountAggSum() int64 {
	if o == nil || o.SyntheticsCheckCallsCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsCheckCallsCountAggSum.Get()
}

// GetSyntheticsCheckCallsCountAggSumOk returns a tuple with the SyntheticsCheckCallsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetSyntheticsCheckCallsCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyntheticsCheckCallsCountAggSum.Get(), o.SyntheticsCheckCallsCountAggSum.IsSet()
}

// HasSyntheticsCheckCallsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSyntheticsCheckCallsCountAggSum() bool {
	return o != nil && o.SyntheticsCheckCallsCountAggSum.IsSet()
}

// SetSyntheticsCheckCallsCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the SyntheticsCheckCallsCountAggSum field.
func (o *UsageSummaryResponse) SetSyntheticsCheckCallsCountAggSum(v int64) {
	o.SyntheticsCheckCallsCountAggSum.Set(&v)
}

// SetSyntheticsCheckCallsCountAggSumNil sets the value for SyntheticsCheckCallsCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetSyntheticsCheckCallsCountAggSumNil() {
	o.SyntheticsCheckCallsCountAggSum.Set(nil)
}

// UnsetSyntheticsCheckCallsCountAggSum ensures that no value is present for SyntheticsCheckCallsCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetSyntheticsCheckCallsCountAggSum() {
	o.SyntheticsCheckCallsCountAggSum.Unset()
}

// GetSyntheticsParallelTestingMaxSlotsHwmSum returns the SyntheticsParallelTestingMaxSlotsHwmSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetSyntheticsParallelTestingMaxSlotsHwmSum() int64 {
	if o == nil || o.SyntheticsParallelTestingMaxSlotsHwmSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsParallelTestingMaxSlotsHwmSum.Get()
}

// GetSyntheticsParallelTestingMaxSlotsHwmSumOk returns a tuple with the SyntheticsParallelTestingMaxSlotsHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetSyntheticsParallelTestingMaxSlotsHwmSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyntheticsParallelTestingMaxSlotsHwmSum.Get(), o.SyntheticsParallelTestingMaxSlotsHwmSum.IsSet()
}

// HasSyntheticsParallelTestingMaxSlotsHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSyntheticsParallelTestingMaxSlotsHwmSum() bool {
	return o != nil && o.SyntheticsParallelTestingMaxSlotsHwmSum.IsSet()
}

// SetSyntheticsParallelTestingMaxSlotsHwmSum gets a reference to the given datadog.NullableInt64 and assigns it to the SyntheticsParallelTestingMaxSlotsHwmSum field.
func (o *UsageSummaryResponse) SetSyntheticsParallelTestingMaxSlotsHwmSum(v int64) {
	o.SyntheticsParallelTestingMaxSlotsHwmSum.Set(&v)
}

// SetSyntheticsParallelTestingMaxSlotsHwmSumNil sets the value for SyntheticsParallelTestingMaxSlotsHwmSum to be an explicit nil.
func (o *UsageSummaryResponse) SetSyntheticsParallelTestingMaxSlotsHwmSumNil() {
	o.SyntheticsParallelTestingMaxSlotsHwmSum.Set(nil)
}

// UnsetSyntheticsParallelTestingMaxSlotsHwmSum ensures that no value is present for SyntheticsParallelTestingMaxSlotsHwmSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetSyntheticsParallelTestingMaxSlotsHwmSum() {
	o.SyntheticsParallelTestingMaxSlotsHwmSum.Unset()
}

// GetTraceSearchIndexedEventsCountAggSum returns the TraceSearchIndexedEventsCountAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetTraceSearchIndexedEventsCountAggSum() int64 {
	if o == nil || o.TraceSearchIndexedEventsCountAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TraceSearchIndexedEventsCountAggSum.Get()
}

// GetTraceSearchIndexedEventsCountAggSumOk returns a tuple with the TraceSearchIndexedEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetTraceSearchIndexedEventsCountAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceSearchIndexedEventsCountAggSum.Get(), o.TraceSearchIndexedEventsCountAggSum.IsSet()
}

// HasTraceSearchIndexedEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasTraceSearchIndexedEventsCountAggSum() bool {
	return o != nil && o.TraceSearchIndexedEventsCountAggSum.IsSet()
}

// SetTraceSearchIndexedEventsCountAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the TraceSearchIndexedEventsCountAggSum field.
func (o *UsageSummaryResponse) SetTraceSearchIndexedEventsCountAggSum(v int64) {
	o.TraceSearchIndexedEventsCountAggSum.Set(&v)
}

// SetTraceSearchIndexedEventsCountAggSumNil sets the value for TraceSearchIndexedEventsCountAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetTraceSearchIndexedEventsCountAggSumNil() {
	o.TraceSearchIndexedEventsCountAggSum.Set(nil)
}

// UnsetTraceSearchIndexedEventsCountAggSum ensures that no value is present for TraceSearchIndexedEventsCountAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetTraceSearchIndexedEventsCountAggSum() {
	o.TraceSearchIndexedEventsCountAggSum.Unset()
}

// GetTwolIngestedEventsBytesAggSum returns the TwolIngestedEventsBytesAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetTwolIngestedEventsBytesAggSum() int64 {
	if o == nil || o.TwolIngestedEventsBytesAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TwolIngestedEventsBytesAggSum.Get()
}

// GetTwolIngestedEventsBytesAggSumOk returns a tuple with the TwolIngestedEventsBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetTwolIngestedEventsBytesAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwolIngestedEventsBytesAggSum.Get(), o.TwolIngestedEventsBytesAggSum.IsSet()
}

// HasTwolIngestedEventsBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasTwolIngestedEventsBytesAggSum() bool {
	return o != nil && o.TwolIngestedEventsBytesAggSum.IsSet()
}

// SetTwolIngestedEventsBytesAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the TwolIngestedEventsBytesAggSum field.
func (o *UsageSummaryResponse) SetTwolIngestedEventsBytesAggSum(v int64) {
	o.TwolIngestedEventsBytesAggSum.Set(&v)
}

// SetTwolIngestedEventsBytesAggSumNil sets the value for TwolIngestedEventsBytesAggSum to be an explicit nil.
func (o *UsageSummaryResponse) SetTwolIngestedEventsBytesAggSumNil() {
	o.TwolIngestedEventsBytesAggSum.Set(nil)
}

// UnsetTwolIngestedEventsBytesAggSum ensures that no value is present for TwolIngestedEventsBytesAggSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetTwolIngestedEventsBytesAggSum() {
	o.TwolIngestedEventsBytesAggSum.Unset()
}

// GetUniversalServiceMonitoringHostTop99pSum returns the UniversalServiceMonitoringHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetUniversalServiceMonitoringHostTop99pSum() int64 {
	if o == nil || o.UniversalServiceMonitoringHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UniversalServiceMonitoringHostTop99pSum.Get()
}

// GetUniversalServiceMonitoringHostTop99pSumOk returns a tuple with the UniversalServiceMonitoringHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetUniversalServiceMonitoringHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringHostTop99pSum.Get(), o.UniversalServiceMonitoringHostTop99pSum.IsSet()
}

// HasUniversalServiceMonitoringHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasUniversalServiceMonitoringHostTop99pSum() bool {
	return o != nil && o.UniversalServiceMonitoringHostTop99pSum.IsSet()
}

// SetUniversalServiceMonitoringHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the UniversalServiceMonitoringHostTop99pSum field.
func (o *UsageSummaryResponse) SetUniversalServiceMonitoringHostTop99pSum(v int64) {
	o.UniversalServiceMonitoringHostTop99pSum.Set(&v)
}

// SetUniversalServiceMonitoringHostTop99pSumNil sets the value for UniversalServiceMonitoringHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetUniversalServiceMonitoringHostTop99pSumNil() {
	o.UniversalServiceMonitoringHostTop99pSum.Set(nil)
}

// UnsetUniversalServiceMonitoringHostTop99pSum ensures that no value is present for UniversalServiceMonitoringHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetUniversalServiceMonitoringHostTop99pSum() {
	o.UniversalServiceMonitoringHostTop99pSum.Unset()
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetUsage() []UsageSummaryDate {
	if o == nil || o.Usage == nil {
		var ret []UsageSummaryDate
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetUsageOk() (*[]UsageSummaryDate, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given []UsageSummaryDate and assigns it to the Usage field.
func (o *UsageSummaryResponse) SetUsage(v []UsageSummaryDate) {
	o.Usage = v
}

// GetVsphereHostTop99pSum returns the VsphereHostTop99pSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSummaryResponse) GetVsphereHostTop99pSum() int64 {
	if o == nil || o.VsphereHostTop99pSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.VsphereHostTop99pSum.Get()
}

// GetVsphereHostTop99pSumOk returns a tuple with the VsphereHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSummaryResponse) GetVsphereHostTop99pSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VsphereHostTop99pSum.Get(), o.VsphereHostTop99pSum.IsSet()
}

// HasVsphereHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasVsphereHostTop99pSum() bool {
	return o != nil && o.VsphereHostTop99pSum.IsSet()
}

// SetVsphereHostTop99pSum gets a reference to the given datadog.NullableInt64 and assigns it to the VsphereHostTop99pSum field.
func (o *UsageSummaryResponse) SetVsphereHostTop99pSum(v int64) {
	o.VsphereHostTop99pSum.Set(&v)
}

// SetVsphereHostTop99pSumNil sets the value for VsphereHostTop99pSum to be an explicit nil.
func (o *UsageSummaryResponse) SetVsphereHostTop99pSumNil() {
	o.VsphereHostTop99pSum.Set(nil)
}

// UnsetVsphereHostTop99pSum ensures that no value is present for VsphereHostTop99pSum, not even an explicit nil.
func (o *UsageSummaryResponse) UnsetVsphereHostTop99pSum() {
	o.VsphereHostTop99pSum.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AgentHostTop99pSum.IsSet() {
		toSerialize["agent_host_top99p_sum"] = o.AgentHostTop99pSum.Get()
	}
	if o.ApmAzureAppServiceHostTop99pSum.IsSet() {
		toSerialize["apm_azure_app_service_host_top99p_sum"] = o.ApmAzureAppServiceHostTop99pSum.Get()
	}
	if o.ApmFargateCountAvgSum.IsSet() {
		toSerialize["apm_fargate_count_avg_sum"] = o.ApmFargateCountAvgSum.Get()
	}
	if o.ApmHostTop99pSum.IsSet() {
		toSerialize["apm_host_top99p_sum"] = o.ApmHostTop99pSum.Get()
	}
	if o.AppsecFargateCountAvgSum.IsSet() {
		toSerialize["appsec_fargate_count_avg_sum"] = o.AppsecFargateCountAvgSum.Get()
	}
	if o.AuditLogsLinesIndexedAggSum.IsSet() {
		toSerialize["audit_logs_lines_indexed_agg_sum"] = o.AuditLogsLinesIndexedAggSum.Get()
	}
	if o.AuditTrailEnabledHwmSum.IsSet() {
		toSerialize["audit_trail_enabled_hwm_sum"] = o.AuditTrailEnabledHwmSum.Get()
	}
	if o.AvgProfiledFargateTasksSum.IsSet() {
		toSerialize["avg_profiled_fargate_tasks_sum"] = o.AvgProfiledFargateTasksSum.Get()
	}
	if o.AwsHostTop99pSum.IsSet() {
		toSerialize["aws_host_top99p_sum"] = o.AwsHostTop99pSum.Get()
	}
	if o.AwsLambdaFuncCount.IsSet() {
		toSerialize["aws_lambda_func_count"] = o.AwsLambdaFuncCount.Get()
	}
	if o.AwsLambdaInvocationsSum.IsSet() {
		toSerialize["aws_lambda_invocations_sum"] = o.AwsLambdaInvocationsSum.Get()
	}
	if o.AzureAppServiceTop99pSum.IsSet() {
		toSerialize["azure_app_service_top99p_sum"] = o.AzureAppServiceTop99pSum.Get()
	}
	if o.AzureHostTop99pSum.IsSet() {
		toSerialize["azure_host_top99p_sum"] = o.AzureHostTop99pSum.Get()
	}
	if o.BillableIngestedBytesAggSum.IsSet() {
		toSerialize["billable_ingested_bytes_agg_sum"] = o.BillableIngestedBytesAggSum.Get()
	}
	if o.BrowserRumLiteSessionCountAggSum.IsSet() {
		toSerialize["browser_rum_lite_session_count_agg_sum"] = o.BrowserRumLiteSessionCountAggSum.Get()
	}
	if o.BrowserRumReplaySessionCountAggSum.IsSet() {
		toSerialize["browser_rum_replay_session_count_agg_sum"] = o.BrowserRumReplaySessionCountAggSum.Get()
	}
	if o.BrowserRumUnitsAggSum.IsSet() {
		toSerialize["browser_rum_units_agg_sum"] = o.BrowserRumUnitsAggSum.Get()
	}
	if o.CiPipelineIndexedSpansAggSum.IsSet() {
		toSerialize["ci_pipeline_indexed_spans_agg_sum"] = o.CiPipelineIndexedSpansAggSum.Get()
	}
	if o.CiTestIndexedSpansAggSum.IsSet() {
		toSerialize["ci_test_indexed_spans_agg_sum"] = o.CiTestIndexedSpansAggSum.Get()
	}
	if o.CiVisibilityPipelineCommittersHwmSum.IsSet() {
		toSerialize["ci_visibility_pipeline_committers_hwm_sum"] = o.CiVisibilityPipelineCommittersHwmSum.Get()
	}
	if o.CiVisibilityTestCommittersHwmSum.IsSet() {
		toSerialize["ci_visibility_test_committers_hwm_sum"] = o.CiVisibilityTestCommittersHwmSum.Get()
	}
	if o.CloudCostManagementHostCountAvgSum.IsSet() {
		toSerialize["cloud_cost_management_host_count_avg_sum"] = o.CloudCostManagementHostCountAvgSum.Get()
	}
	if o.ContainerAvgSum.IsSet() {
		toSerialize["container_avg_sum"] = o.ContainerAvgSum.Get()
	}
	if o.ContainerExclAgentAvgSum.IsSet() {
		toSerialize["container_excl_agent_avg_sum"] = o.ContainerExclAgentAvgSum.Get()
	}
	if o.ContainerHwmSum.IsSet() {
		toSerialize["container_hwm_sum"] = o.ContainerHwmSum.Get()
	}
	if o.CspmAasHostTop99pSum.IsSet() {
		toSerialize["cspm_aas_host_top99p_sum"] = o.CspmAasHostTop99pSum.Get()
	}
	if o.CspmAwsHostTop99pSum.IsSet() {
		toSerialize["cspm_aws_host_top99p_sum"] = o.CspmAwsHostTop99pSum.Get()
	}
	if o.CspmAzureHostTop99pSum.IsSet() {
		toSerialize["cspm_azure_host_top99p_sum"] = o.CspmAzureHostTop99pSum.Get()
	}
	if o.CspmContainerAvgSum.IsSet() {
		toSerialize["cspm_container_avg_sum"] = o.CspmContainerAvgSum.Get()
	}
	if o.CspmContainerHwmSum.IsSet() {
		toSerialize["cspm_container_hwm_sum"] = o.CspmContainerHwmSum.Get()
	}
	if o.CspmGcpHostTop99pSum.IsSet() {
		toSerialize["cspm_gcp_host_top99p_sum"] = o.CspmGcpHostTop99pSum.Get()
	}
	if o.CspmHostTop99pSum.IsSet() {
		toSerialize["cspm_host_top99p_sum"] = o.CspmHostTop99pSum.Get()
	}
	if o.CustomTsSum.IsSet() {
		toSerialize["custom_ts_sum"] = o.CustomTsSum.Get()
	}
	if o.CwsContainersAvgSum.IsSet() {
		toSerialize["cws_containers_avg_sum"] = o.CwsContainersAvgSum.Get()
	}
	if o.CwsHostTop99pSum.IsSet() {
		toSerialize["cws_host_top99p_sum"] = o.CwsHostTop99pSum.Get()
	}
	if o.DbmHostTop99pSum.IsSet() {
		toSerialize["dbm_host_top99p_sum"] = o.DbmHostTop99pSum.Get()
	}
	if o.DbmQueriesAvgSum.IsSet() {
		toSerialize["dbm_queries_avg_sum"] = o.DbmQueriesAvgSum.Get()
	}
	if o.EndDate != nil {
		if o.EndDate.Nanosecond() == 0 {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.FargateTasksCountAvgSum.IsSet() {
		toSerialize["fargate_tasks_count_avg_sum"] = o.FargateTasksCountAvgSum.Get()
	}
	if o.FargateTasksCountHwmSum.IsSet() {
		toSerialize["fargate_tasks_count_hwm_sum"] = o.FargateTasksCountHwmSum.Get()
	}
	if o.ForwardingEventsBytesAggSum.IsSet() {
		toSerialize["forwarding_events_bytes_agg_sum"] = o.ForwardingEventsBytesAggSum.Get()
	}
	if o.GcpHostTop99pSum.IsSet() {
		toSerialize["gcp_host_top99p_sum"] = o.GcpHostTop99pSum.Get()
	}
	if o.HerokuHostTop99pSum.IsSet() {
		toSerialize["heroku_host_top99p_sum"] = o.HerokuHostTop99pSum.Get()
	}
	if o.IncidentManagementMonthlyActiveUsersHwmSum.IsSet() {
		toSerialize["incident_management_monthly_active_users_hwm_sum"] = o.IncidentManagementMonthlyActiveUsersHwmSum.Get()
	}
	if o.IndexedEventsCountAggSum.IsSet() {
		toSerialize["indexed_events_count_agg_sum"] = o.IndexedEventsCountAggSum.Get()
	}
	if o.InfraHostTop99pSum.IsSet() {
		toSerialize["infra_host_top99p_sum"] = o.InfraHostTop99pSum.Get()
	}
	if o.IngestedEventsBytesAggSum.IsSet() {
		toSerialize["ingested_events_bytes_agg_sum"] = o.IngestedEventsBytesAggSum.Get()
	}
	if o.IotDeviceAggSum.IsSet() {
		toSerialize["iot_device_agg_sum"] = o.IotDeviceAggSum.Get()
	}
	if o.IotDeviceTop99pSum.IsSet() {
		toSerialize["iot_device_top99p_sum"] = o.IotDeviceTop99pSum.Get()
	}
	if o.LastUpdated != nil {
		if o.LastUpdated.Nanosecond() == 0 {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LiveIndexedEventsAggSum.IsSet() {
		toSerialize["live_indexed_events_agg_sum"] = o.LiveIndexedEventsAggSum.Get()
	}
	if o.LiveIngestedBytesAggSum.IsSet() {
		toSerialize["live_ingested_bytes_agg_sum"] = o.LiveIngestedBytesAggSum.Get()
	}
	if o.LogsByRetention != nil {
		toSerialize["logs_by_retention"] = o.LogsByRetention
	}
	if o.MobileRumLiteSessionCountAggSum.IsSet() {
		toSerialize["mobile_rum_lite_session_count_agg_sum"] = o.MobileRumLiteSessionCountAggSum.Get()
	}
	if o.MobileRumSessionCountAggSum.IsSet() {
		toSerialize["mobile_rum_session_count_agg_sum"] = o.MobileRumSessionCountAggSum.Get()
	}
	if o.MobileRumSessionCountAndroidAggSum.IsSet() {
		toSerialize["mobile_rum_session_count_android_agg_sum"] = o.MobileRumSessionCountAndroidAggSum.Get()
	}
	if o.MobileRumSessionCountFlutterAggSum.IsSet() {
		toSerialize["mobile_rum_session_count_flutter_agg_sum"] = o.MobileRumSessionCountFlutterAggSum.Get()
	}
	if o.MobileRumSessionCountIosAggSum.IsSet() {
		toSerialize["mobile_rum_session_count_ios_agg_sum"] = o.MobileRumSessionCountIosAggSum.Get()
	}
	if o.MobileRumSessionCountReactnativeAggSum.IsSet() {
		toSerialize["mobile_rum_session_count_reactnative_agg_sum"] = o.MobileRumSessionCountReactnativeAggSum.Get()
	}
	if o.MobileRumUnitsAggSum.IsSet() {
		toSerialize["mobile_rum_units_agg_sum"] = o.MobileRumUnitsAggSum.Get()
	}
	if o.NetflowIndexedEventsCountAggSum.IsSet() {
		toSerialize["netflow_indexed_events_count_agg_sum"] = o.NetflowIndexedEventsCountAggSum.Get()
	}
	if o.NpmHostTop99pSum.IsSet() {
		toSerialize["npm_host_top99p_sum"] = o.NpmHostTop99pSum.Get()
	}
	if o.ObservabilityPipelinesBytesProcessedAggSum.IsSet() {
		toSerialize["observability_pipelines_bytes_processed_agg_sum"] = o.ObservabilityPipelinesBytesProcessedAggSum.Get()
	}
	if o.OnlineArchiveEventsCountAggSum.IsSet() {
		toSerialize["online_archive_events_count_agg_sum"] = o.OnlineArchiveEventsCountAggSum.Get()
	}
	if o.OpentelemetryApmHostTop99pSum.IsSet() {
		toSerialize["opentelemetry_apm_host_top99p_sum"] = o.OpentelemetryApmHostTop99pSum.Get()
	}
	if o.OpentelemetryHostTop99pSum.IsSet() {
		toSerialize["opentelemetry_host_top99p_sum"] = o.OpentelemetryHostTop99pSum.Get()
	}
	if o.ProfilingContainerAgentCountAvg.IsSet() {
		toSerialize["profiling_container_agent_count_avg"] = o.ProfilingContainerAgentCountAvg.Get()
	}
	if o.ProfilingHostCountTop99pSum.IsSet() {
		toSerialize["profiling_host_count_top99p_sum"] = o.ProfilingHostCountTop99pSum.Get()
	}
	if o.RehydratedIndexedEventsAggSum.IsSet() {
		toSerialize["rehydrated_indexed_events_agg_sum"] = o.RehydratedIndexedEventsAggSum.Get()
	}
	if o.RehydratedIngestedBytesAggSum.IsSet() {
		toSerialize["rehydrated_ingested_bytes_agg_sum"] = o.RehydratedIngestedBytesAggSum.Get()
	}
	if o.RumBrowserAndMobileSessionCount.IsSet() {
		toSerialize["rum_browser_and_mobile_session_count"] = o.RumBrowserAndMobileSessionCount.Get()
	}
	if o.RumSessionCountAggSum.IsSet() {
		toSerialize["rum_session_count_agg_sum"] = o.RumSessionCountAggSum.Get()
	}
	if o.RumTotalSessionCountAggSum.IsSet() {
		toSerialize["rum_total_session_count_agg_sum"] = o.RumTotalSessionCountAggSum.Get()
	}
	if o.RumUnitsAggSum.IsSet() {
		toSerialize["rum_units_agg_sum"] = o.RumUnitsAggSum.Get()
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
	if o.StartDate != nil {
		if o.StartDate.Nanosecond() == 0 {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.SyntheticsBrowserCheckCallsCountAggSum.IsSet() {
		toSerialize["synthetics_browser_check_calls_count_agg_sum"] = o.SyntheticsBrowserCheckCallsCountAggSum.Get()
	}
	if o.SyntheticsCheckCallsCountAggSum.IsSet() {
		toSerialize["synthetics_check_calls_count_agg_sum"] = o.SyntheticsCheckCallsCountAggSum.Get()
	}
	if o.SyntheticsParallelTestingMaxSlotsHwmSum.IsSet() {
		toSerialize["synthetics_parallel_testing_max_slots_hwm_sum"] = o.SyntheticsParallelTestingMaxSlotsHwmSum.Get()
	}
	if o.TraceSearchIndexedEventsCountAggSum.IsSet() {
		toSerialize["trace_search_indexed_events_count_agg_sum"] = o.TraceSearchIndexedEventsCountAggSum.Get()
	}
	if o.TwolIngestedEventsBytesAggSum.IsSet() {
		toSerialize["twol_ingested_events_bytes_agg_sum"] = o.TwolIngestedEventsBytesAggSum.Get()
	}
	if o.UniversalServiceMonitoringHostTop99pSum.IsSet() {
		toSerialize["universal_service_monitoring_host_top99p_sum"] = o.UniversalServiceMonitoringHostTop99pSum.Get()
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	if o.VsphereHostTop99pSum.IsSet() {
		toSerialize["vsphere_host_top99p_sum"] = o.VsphereHostTop99pSum.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageSummaryResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AgentHostTop99pSum                         datadog.NullableInt64 `json:"agent_host_top99p_sum,omitempty"`
		ApmAzureAppServiceHostTop99pSum            datadog.NullableInt64 `json:"apm_azure_app_service_host_top99p_sum,omitempty"`
		ApmFargateCountAvgSum                      datadog.NullableInt64 `json:"apm_fargate_count_avg_sum,omitempty"`
		ApmHostTop99pSum                           datadog.NullableInt64 `json:"apm_host_top99p_sum,omitempty"`
		AppsecFargateCountAvgSum                   datadog.NullableInt64 `json:"appsec_fargate_count_avg_sum,omitempty"`
		AuditLogsLinesIndexedAggSum                datadog.NullableInt64 `json:"audit_logs_lines_indexed_agg_sum,omitempty"`
		AuditTrailEnabledHwmSum                    datadog.NullableInt64 `json:"audit_trail_enabled_hwm_sum,omitempty"`
		AvgProfiledFargateTasksSum                 datadog.NullableInt64 `json:"avg_profiled_fargate_tasks_sum,omitempty"`
		AwsHostTop99pSum                           datadog.NullableInt64 `json:"aws_host_top99p_sum,omitempty"`
		AwsLambdaFuncCount                         datadog.NullableInt64 `json:"aws_lambda_func_count,omitempty"`
		AwsLambdaInvocationsSum                    datadog.NullableInt64 `json:"aws_lambda_invocations_sum,omitempty"`
		AzureAppServiceTop99pSum                   datadog.NullableInt64 `json:"azure_app_service_top99p_sum,omitempty"`
		AzureHostTop99pSum                         datadog.NullableInt64 `json:"azure_host_top99p_sum,omitempty"`
		BillableIngestedBytesAggSum                datadog.NullableInt64 `json:"billable_ingested_bytes_agg_sum,omitempty"`
		BrowserRumLiteSessionCountAggSum           datadog.NullableInt64 `json:"browser_rum_lite_session_count_agg_sum,omitempty"`
		BrowserRumReplaySessionCountAggSum         datadog.NullableInt64 `json:"browser_rum_replay_session_count_agg_sum,omitempty"`
		BrowserRumUnitsAggSum                      datadog.NullableInt64 `json:"browser_rum_units_agg_sum,omitempty"`
		CiPipelineIndexedSpansAggSum               datadog.NullableInt64 `json:"ci_pipeline_indexed_spans_agg_sum,omitempty"`
		CiTestIndexedSpansAggSum                   datadog.NullableInt64 `json:"ci_test_indexed_spans_agg_sum,omitempty"`
		CiVisibilityPipelineCommittersHwmSum       datadog.NullableInt64 `json:"ci_visibility_pipeline_committers_hwm_sum,omitempty"`
		CiVisibilityTestCommittersHwmSum           datadog.NullableInt64 `json:"ci_visibility_test_committers_hwm_sum,omitempty"`
		CloudCostManagementHostCountAvgSum         datadog.NullableInt64 `json:"cloud_cost_management_host_count_avg_sum,omitempty"`
		ContainerAvgSum                            datadog.NullableInt64 `json:"container_avg_sum,omitempty"`
		ContainerExclAgentAvgSum                   datadog.NullableInt64 `json:"container_excl_agent_avg_sum,omitempty"`
		ContainerHwmSum                            datadog.NullableInt64 `json:"container_hwm_sum,omitempty"`
		CspmAasHostTop99pSum                       datadog.NullableInt64 `json:"cspm_aas_host_top99p_sum,omitempty"`
		CspmAwsHostTop99pSum                       datadog.NullableInt64 `json:"cspm_aws_host_top99p_sum,omitempty"`
		CspmAzureHostTop99pSum                     datadog.NullableInt64 `json:"cspm_azure_host_top99p_sum,omitempty"`
		CspmContainerAvgSum                        datadog.NullableInt64 `json:"cspm_container_avg_sum,omitempty"`
		CspmContainerHwmSum                        datadog.NullableInt64 `json:"cspm_container_hwm_sum,omitempty"`
		CspmGcpHostTop99pSum                       datadog.NullableInt64 `json:"cspm_gcp_host_top99p_sum,omitempty"`
		CspmHostTop99pSum                          datadog.NullableInt64 `json:"cspm_host_top99p_sum,omitempty"`
		CustomTsSum                                datadog.NullableInt64 `json:"custom_ts_sum,omitempty"`
		CwsContainersAvgSum                        datadog.NullableInt64 `json:"cws_containers_avg_sum,omitempty"`
		CwsHostTop99pSum                           datadog.NullableInt64 `json:"cws_host_top99p_sum,omitempty"`
		DbmHostTop99pSum                           datadog.NullableInt64 `json:"dbm_host_top99p_sum,omitempty"`
		DbmQueriesAvgSum                           datadog.NullableInt64 `json:"dbm_queries_avg_sum,omitempty"`
		EndDate                                    *time.Time            `json:"end_date,omitempty"`
		FargateTasksCountAvgSum                    datadog.NullableInt64 `json:"fargate_tasks_count_avg_sum,omitempty"`
		FargateTasksCountHwmSum                    datadog.NullableInt64 `json:"fargate_tasks_count_hwm_sum,omitempty"`
		ForwardingEventsBytesAggSum                datadog.NullableInt64 `json:"forwarding_events_bytes_agg_sum,omitempty"`
		GcpHostTop99pSum                           datadog.NullableInt64 `json:"gcp_host_top99p_sum,omitempty"`
		HerokuHostTop99pSum                        datadog.NullableInt64 `json:"heroku_host_top99p_sum,omitempty"`
		IncidentManagementMonthlyActiveUsersHwmSum datadog.NullableInt64 `json:"incident_management_monthly_active_users_hwm_sum,omitempty"`
		IndexedEventsCountAggSum                   datadog.NullableInt64 `json:"indexed_events_count_agg_sum,omitempty"`
		InfraHostTop99pSum                         datadog.NullableInt64 `json:"infra_host_top99p_sum,omitempty"`
		IngestedEventsBytesAggSum                  datadog.NullableInt64 `json:"ingested_events_bytes_agg_sum,omitempty"`
		IotDeviceAggSum                            datadog.NullableInt64 `json:"iot_device_agg_sum,omitempty"`
		IotDeviceTop99pSum                         datadog.NullableInt64 `json:"iot_device_top99p_sum,omitempty"`
		LastUpdated                                *time.Time            `json:"last_updated,omitempty"`
		LiveIndexedEventsAggSum                    datadog.NullableInt64 `json:"live_indexed_events_agg_sum,omitempty"`
		LiveIngestedBytesAggSum                    datadog.NullableInt64 `json:"live_ingested_bytes_agg_sum,omitempty"`
		LogsByRetention                            *LogsByRetention      `json:"logs_by_retention,omitempty"`
		MobileRumLiteSessionCountAggSum            datadog.NullableInt64 `json:"mobile_rum_lite_session_count_agg_sum,omitempty"`
		MobileRumSessionCountAggSum                datadog.NullableInt64 `json:"mobile_rum_session_count_agg_sum,omitempty"`
		MobileRumSessionCountAndroidAggSum         datadog.NullableInt64 `json:"mobile_rum_session_count_android_agg_sum,omitempty"`
		MobileRumSessionCountFlutterAggSum         datadog.NullableInt64 `json:"mobile_rum_session_count_flutter_agg_sum,omitempty"`
		MobileRumSessionCountIosAggSum             datadog.NullableInt64 `json:"mobile_rum_session_count_ios_agg_sum,omitempty"`
		MobileRumSessionCountReactnativeAggSum     datadog.NullableInt64 `json:"mobile_rum_session_count_reactnative_agg_sum,omitempty"`
		MobileRumUnitsAggSum                       datadog.NullableInt64 `json:"mobile_rum_units_agg_sum,omitempty"`
		NetflowIndexedEventsCountAggSum            datadog.NullableInt64 `json:"netflow_indexed_events_count_agg_sum,omitempty"`
		NpmHostTop99pSum                           datadog.NullableInt64 `json:"npm_host_top99p_sum,omitempty"`
		ObservabilityPipelinesBytesProcessedAggSum datadog.NullableInt64 `json:"observability_pipelines_bytes_processed_agg_sum,omitempty"`
		OnlineArchiveEventsCountAggSum             datadog.NullableInt64 `json:"online_archive_events_count_agg_sum,omitempty"`
		OpentelemetryApmHostTop99pSum              datadog.NullableInt64 `json:"opentelemetry_apm_host_top99p_sum,omitempty"`
		OpentelemetryHostTop99pSum                 datadog.NullableInt64 `json:"opentelemetry_host_top99p_sum,omitempty"`
		ProfilingContainerAgentCountAvg            datadog.NullableInt64 `json:"profiling_container_agent_count_avg,omitempty"`
		ProfilingHostCountTop99pSum                datadog.NullableInt64 `json:"profiling_host_count_top99p_sum,omitempty"`
		RehydratedIndexedEventsAggSum              datadog.NullableInt64 `json:"rehydrated_indexed_events_agg_sum,omitempty"`
		RehydratedIngestedBytesAggSum              datadog.NullableInt64 `json:"rehydrated_ingested_bytes_agg_sum,omitempty"`
		RumBrowserAndMobileSessionCount            datadog.NullableInt64 `json:"rum_browser_and_mobile_session_count,omitempty"`
		RumSessionCountAggSum                      datadog.NullableInt64 `json:"rum_session_count_agg_sum,omitempty"`
		RumTotalSessionCountAggSum                 datadog.NullableInt64 `json:"rum_total_session_count_agg_sum,omitempty"`
		RumUnitsAggSum                             datadog.NullableInt64 `json:"rum_units_agg_sum,omitempty"`
		SdsApmScannedBytesSum                      datadog.NullableInt64 `json:"sds_apm_scanned_bytes_sum,omitempty"`
		SdsEventsScannedBytesSum                   datadog.NullableInt64 `json:"sds_events_scanned_bytes_sum,omitempty"`
		SdsLogsScannedBytesSum                     datadog.NullableInt64 `json:"sds_logs_scanned_bytes_sum,omitempty"`
		SdsRumScannedBytesSum                      datadog.NullableInt64 `json:"sds_rum_scanned_bytes_sum,omitempty"`
		SdsTotalScannedBytesSum                    datadog.NullableInt64 `json:"sds_total_scanned_bytes_sum,omitempty"`
		StartDate                                  *time.Time            `json:"start_date,omitempty"`
		SyntheticsBrowserCheckCallsCountAggSum     datadog.NullableInt64 `json:"synthetics_browser_check_calls_count_agg_sum,omitempty"`
		SyntheticsCheckCallsCountAggSum            datadog.NullableInt64 `json:"synthetics_check_calls_count_agg_sum,omitempty"`
		SyntheticsParallelTestingMaxSlotsHwmSum    datadog.NullableInt64 `json:"synthetics_parallel_testing_max_slots_hwm_sum,omitempty"`
		TraceSearchIndexedEventsCountAggSum        datadog.NullableInt64 `json:"trace_search_indexed_events_count_agg_sum,omitempty"`
		TwolIngestedEventsBytesAggSum              datadog.NullableInt64 `json:"twol_ingested_events_bytes_agg_sum,omitempty"`
		UniversalServiceMonitoringHostTop99pSum    datadog.NullableInt64 `json:"universal_service_monitoring_host_top99p_sum,omitempty"`
		Usage                                      []UsageSummaryDate    `json:"usage,omitempty"`
		VsphereHostTop99pSum                       datadog.NullableInt64 `json:"vsphere_host_top99p_sum,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_host_top99p_sum", "apm_azure_app_service_host_top99p_sum", "apm_fargate_count_avg_sum", "apm_host_top99p_sum", "appsec_fargate_count_avg_sum", "audit_logs_lines_indexed_agg_sum", "audit_trail_enabled_hwm_sum", "avg_profiled_fargate_tasks_sum", "aws_host_top99p_sum", "aws_lambda_func_count", "aws_lambda_invocations_sum", "azure_app_service_top99p_sum", "azure_host_top99p_sum", "billable_ingested_bytes_agg_sum", "browser_rum_lite_session_count_agg_sum", "browser_rum_replay_session_count_agg_sum", "browser_rum_units_agg_sum", "ci_pipeline_indexed_spans_agg_sum", "ci_test_indexed_spans_agg_sum", "ci_visibility_pipeline_committers_hwm_sum", "ci_visibility_test_committers_hwm_sum", "cloud_cost_management_host_count_avg_sum", "container_avg_sum", "container_excl_agent_avg_sum", "container_hwm_sum", "cspm_aas_host_top99p_sum", "cspm_aws_host_top99p_sum", "cspm_azure_host_top99p_sum", "cspm_container_avg_sum", "cspm_container_hwm_sum", "cspm_gcp_host_top99p_sum", "cspm_host_top99p_sum", "custom_ts_sum", "cws_containers_avg_sum", "cws_host_top99p_sum", "dbm_host_top99p_sum", "dbm_queries_avg_sum", "end_date", "fargate_tasks_count_avg_sum", "fargate_tasks_count_hwm_sum", "forwarding_events_bytes_agg_sum", "gcp_host_top99p_sum", "heroku_host_top99p_sum", "incident_management_monthly_active_users_hwm_sum", "indexed_events_count_agg_sum", "infra_host_top99p_sum", "ingested_events_bytes_agg_sum", "iot_device_agg_sum", "iot_device_top99p_sum", "last_updated", "live_indexed_events_agg_sum", "live_ingested_bytes_agg_sum", "logs_by_retention", "mobile_rum_lite_session_count_agg_sum", "mobile_rum_session_count_agg_sum", "mobile_rum_session_count_android_agg_sum", "mobile_rum_session_count_flutter_agg_sum", "mobile_rum_session_count_ios_agg_sum", "mobile_rum_session_count_reactnative_agg_sum", "mobile_rum_units_agg_sum", "netflow_indexed_events_count_agg_sum", "npm_host_top99p_sum", "observability_pipelines_bytes_processed_agg_sum", "online_archive_events_count_agg_sum", "opentelemetry_apm_host_top99p_sum", "opentelemetry_host_top99p_sum", "profiling_container_agent_count_avg", "profiling_host_count_top99p_sum", "rehydrated_indexed_events_agg_sum", "rehydrated_ingested_bytes_agg_sum", "rum_browser_and_mobile_session_count", "rum_session_count_agg_sum", "rum_total_session_count_agg_sum", "rum_units_agg_sum", "sds_apm_scanned_bytes_sum", "sds_events_scanned_bytes_sum", "sds_logs_scanned_bytes_sum", "sds_rum_scanned_bytes_sum", "sds_total_scanned_bytes_sum", "start_date", "synthetics_browser_check_calls_count_agg_sum", "synthetics_check_calls_count_agg_sum", "synthetics_parallel_testing_max_slots_hwm_sum", "trace_search_indexed_events_count_agg_sum", "twol_ingested_events_bytes_agg_sum", "universal_service_monitoring_host_top99p_sum", "usage", "vsphere_host_top99p_sum"})
	} else {
		return err
	}
	o.AgentHostTop99pSum = all.AgentHostTop99pSum
	o.ApmAzureAppServiceHostTop99pSum = all.ApmAzureAppServiceHostTop99pSum
	o.ApmFargateCountAvgSum = all.ApmFargateCountAvgSum
	o.ApmHostTop99pSum = all.ApmHostTop99pSum
	o.AppsecFargateCountAvgSum = all.AppsecFargateCountAvgSum
	o.AuditLogsLinesIndexedAggSum = all.AuditLogsLinesIndexedAggSum
	o.AuditTrailEnabledHwmSum = all.AuditTrailEnabledHwmSum
	o.AvgProfiledFargateTasksSum = all.AvgProfiledFargateTasksSum
	o.AwsHostTop99pSum = all.AwsHostTop99pSum
	o.AwsLambdaFuncCount = all.AwsLambdaFuncCount
	o.AwsLambdaInvocationsSum = all.AwsLambdaInvocationsSum
	o.AzureAppServiceTop99pSum = all.AzureAppServiceTop99pSum
	o.AzureHostTop99pSum = all.AzureHostTop99pSum
	o.BillableIngestedBytesAggSum = all.BillableIngestedBytesAggSum
	o.BrowserRumLiteSessionCountAggSum = all.BrowserRumLiteSessionCountAggSum
	o.BrowserRumReplaySessionCountAggSum = all.BrowserRumReplaySessionCountAggSum
	o.BrowserRumUnitsAggSum = all.BrowserRumUnitsAggSum
	o.CiPipelineIndexedSpansAggSum = all.CiPipelineIndexedSpansAggSum
	o.CiTestIndexedSpansAggSum = all.CiTestIndexedSpansAggSum
	o.CiVisibilityPipelineCommittersHwmSum = all.CiVisibilityPipelineCommittersHwmSum
	o.CiVisibilityTestCommittersHwmSum = all.CiVisibilityTestCommittersHwmSum
	o.CloudCostManagementHostCountAvgSum = all.CloudCostManagementHostCountAvgSum
	o.ContainerAvgSum = all.ContainerAvgSum
	o.ContainerExclAgentAvgSum = all.ContainerExclAgentAvgSum
	o.ContainerHwmSum = all.ContainerHwmSum
	o.CspmAasHostTop99pSum = all.CspmAasHostTop99pSum
	o.CspmAwsHostTop99pSum = all.CspmAwsHostTop99pSum
	o.CspmAzureHostTop99pSum = all.CspmAzureHostTop99pSum
	o.CspmContainerAvgSum = all.CspmContainerAvgSum
	o.CspmContainerHwmSum = all.CspmContainerHwmSum
	o.CspmGcpHostTop99pSum = all.CspmGcpHostTop99pSum
	o.CspmHostTop99pSum = all.CspmHostTop99pSum
	o.CustomTsSum = all.CustomTsSum
	o.CwsContainersAvgSum = all.CwsContainersAvgSum
	o.CwsHostTop99pSum = all.CwsHostTop99pSum
	o.DbmHostTop99pSum = all.DbmHostTop99pSum
	o.DbmQueriesAvgSum = all.DbmQueriesAvgSum
	o.EndDate = all.EndDate
	o.FargateTasksCountAvgSum = all.FargateTasksCountAvgSum
	o.FargateTasksCountHwmSum = all.FargateTasksCountHwmSum
	o.ForwardingEventsBytesAggSum = all.ForwardingEventsBytesAggSum
	o.GcpHostTop99pSum = all.GcpHostTop99pSum
	o.HerokuHostTop99pSum = all.HerokuHostTop99pSum
	o.IncidentManagementMonthlyActiveUsersHwmSum = all.IncidentManagementMonthlyActiveUsersHwmSum
	o.IndexedEventsCountAggSum = all.IndexedEventsCountAggSum
	o.InfraHostTop99pSum = all.InfraHostTop99pSum
	o.IngestedEventsBytesAggSum = all.IngestedEventsBytesAggSum
	o.IotDeviceAggSum = all.IotDeviceAggSum
	o.IotDeviceTop99pSum = all.IotDeviceTop99pSum
	o.LastUpdated = all.LastUpdated
	o.LiveIndexedEventsAggSum = all.LiveIndexedEventsAggSum
	o.LiveIngestedBytesAggSum = all.LiveIngestedBytesAggSum
	if all.LogsByRetention != nil && all.LogsByRetention.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsByRetention = all.LogsByRetention
	o.MobileRumLiteSessionCountAggSum = all.MobileRumLiteSessionCountAggSum
	o.MobileRumSessionCountAggSum = all.MobileRumSessionCountAggSum
	o.MobileRumSessionCountAndroidAggSum = all.MobileRumSessionCountAndroidAggSum
	o.MobileRumSessionCountFlutterAggSum = all.MobileRumSessionCountFlutterAggSum
	o.MobileRumSessionCountIosAggSum = all.MobileRumSessionCountIosAggSum
	o.MobileRumSessionCountReactnativeAggSum = all.MobileRumSessionCountReactnativeAggSum
	o.MobileRumUnitsAggSum = all.MobileRumUnitsAggSum
	o.NetflowIndexedEventsCountAggSum = all.NetflowIndexedEventsCountAggSum
	o.NpmHostTop99pSum = all.NpmHostTop99pSum
	o.ObservabilityPipelinesBytesProcessedAggSum = all.ObservabilityPipelinesBytesProcessedAggSum
	o.OnlineArchiveEventsCountAggSum = all.OnlineArchiveEventsCountAggSum
	o.OpentelemetryApmHostTop99pSum = all.OpentelemetryApmHostTop99pSum
	o.OpentelemetryHostTop99pSum = all.OpentelemetryHostTop99pSum
	o.ProfilingContainerAgentCountAvg = all.ProfilingContainerAgentCountAvg
	o.ProfilingHostCountTop99pSum = all.ProfilingHostCountTop99pSum
	o.RehydratedIndexedEventsAggSum = all.RehydratedIndexedEventsAggSum
	o.RehydratedIngestedBytesAggSum = all.RehydratedIngestedBytesAggSum
	o.RumBrowserAndMobileSessionCount = all.RumBrowserAndMobileSessionCount
	o.RumSessionCountAggSum = all.RumSessionCountAggSum
	o.RumTotalSessionCountAggSum = all.RumTotalSessionCountAggSum
	o.RumUnitsAggSum = all.RumUnitsAggSum
	o.SdsApmScannedBytesSum = all.SdsApmScannedBytesSum
	o.SdsEventsScannedBytesSum = all.SdsEventsScannedBytesSum
	o.SdsLogsScannedBytesSum = all.SdsLogsScannedBytesSum
	o.SdsRumScannedBytesSum = all.SdsRumScannedBytesSum
	o.SdsTotalScannedBytesSum = all.SdsTotalScannedBytesSum
	o.StartDate = all.StartDate
	o.SyntheticsBrowserCheckCallsCountAggSum = all.SyntheticsBrowserCheckCallsCountAggSum
	o.SyntheticsCheckCallsCountAggSum = all.SyntheticsCheckCallsCountAggSum
	o.SyntheticsParallelTestingMaxSlotsHwmSum = all.SyntheticsParallelTestingMaxSlotsHwmSum
	o.TraceSearchIndexedEventsCountAggSum = all.TraceSearchIndexedEventsCountAggSum
	o.TwolIngestedEventsBytesAggSum = all.TwolIngestedEventsBytesAggSum
	o.UniversalServiceMonitoringHostTop99pSum = all.UniversalServiceMonitoringHostTop99pSum
	o.Usage = all.Usage
	o.VsphereHostTop99pSum = all.VsphereHostTop99pSum
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
