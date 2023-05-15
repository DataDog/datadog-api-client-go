// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyUsageAttributionValues Fields in Usage Summary by tag(s).
type MonthlyUsageAttributionValues struct {
	// The percentage of synthetic API test usage by tag(s).
	ApiPercentage datadog.NullableFloat64 `json:"api_percentage,omitempty"`
	// The synthetic API test usage by tag(s).
	ApiUsage datadog.NullableFloat64 `json:"api_usage,omitempty"`
	// The percentage of APM ECS Fargate task usage by tag(s).
	ApmFargatePercentage datadog.NullableFloat64 `json:"apm_fargate_percentage,omitempty"`
	// The APM ECS Fargate task usage by tag(s).
	ApmFargateUsage datadog.NullableFloat64 `json:"apm_fargate_usage,omitempty"`
	// The percentage of APM host usage by tag(s).
	ApmHostPercentage datadog.NullableFloat64 `json:"apm_host_percentage,omitempty"`
	// The APM host usage by tag(s).
	ApmHostUsage datadog.NullableFloat64 `json:"apm_host_usage,omitempty"`
	// The percentage of Application Security Monitoring ECS Fargate task usage by tag(s).
	AppsecFargatePercentage datadog.NullableFloat64 `json:"appsec_fargate_percentage,omitempty"`
	// The Application Security Monitoring ECS Fargate task usage by tag(s).
	AppsecFargateUsage datadog.NullableFloat64 `json:"appsec_fargate_usage,omitempty"`
	// The percentage of Application Security Monitoring host usage by tag(s).
	AppsecPercentage datadog.NullableFloat64 `json:"appsec_percentage,omitempty"`
	// The Application Security Monitoring host usage by tag(s).
	AppsecUsage datadog.NullableFloat64 `json:"appsec_usage,omitempty"`
	// The percentage of synthetic browser test usage by tag(s).
	BrowserPercentage datadog.NullableFloat64 `json:"browser_percentage,omitempty"`
	// The synthetic browser test usage by tag(s).
	BrowserUsage datadog.NullableFloat64 `json:"browser_usage,omitempty"`
	// The percentage of container usage without the Datadog Agent by tag(s).
	ContainerExclAgentPercentage datadog.NullableFloat64 `json:"container_excl_agent_percentage,omitempty"`
	// The container usage without the Datadog Agent by tag(s).
	ContainerExclAgentUsage datadog.NullableFloat64 `json:"container_excl_agent_usage,omitempty"`
	// The percentage of container usage by tag(s).
	ContainerPercentage datadog.NullableFloat64 `json:"container_percentage,omitempty"`
	// The container usage by tag(s).
	ContainerUsage datadog.NullableFloat64 `json:"container_usage,omitempty"`
	// The percentage of CSPM container usage by tag(s).
	CspmContainersPercentage datadog.NullableFloat64 `json:"cspm_containers_percentage,omitempty"`
	// The CSPM container usage by tag(s).
	CspmContainersUsage datadog.NullableFloat64 `json:"cspm_containers_usage,omitempty"`
	// The percentage of CSPM host usage by by tag(s).
	CspmHostsPercentage datadog.NullableFloat64 `json:"cspm_hosts_percentage,omitempty"`
	// The CSPM host usage by tag(s).
	CspmHostsUsage datadog.NullableFloat64 `json:"cspm_hosts_usage,omitempty"`
	// The percentage of ingested custom metrics usage by tag(s).
	CustomIngestedTimeseriesPercentage datadog.NullableFloat64 `json:"custom_ingested_timeseries_percentage,omitempty"`
	// The ingested custom metrics usage by tag(s).
	CustomIngestedTimeseriesUsage datadog.NullableFloat64 `json:"custom_ingested_timeseries_usage,omitempty"`
	// The percentage of indexed custom metrics usage by tag(s).
	CustomTimeseriesPercentage datadog.NullableFloat64 `json:"custom_timeseries_percentage,omitempty"`
	// The indexed custom metrics usage by tag(s).
	CustomTimeseriesUsage datadog.NullableFloat64 `json:"custom_timeseries_usage,omitempty"`
	// The percentage of Cloud Workload Security container usage by tag(s).
	CwsContainersPercentage datadog.NullableFloat64 `json:"cws_containers_percentage,omitempty"`
	// The Cloud Workload Security container usage by tag(s).
	CwsContainersUsage datadog.NullableFloat64 `json:"cws_containers_usage,omitempty"`
	// The percentage of Cloud Workload Security host usage by tag(s).
	CwsHostsPercentage datadog.NullableFloat64 `json:"cws_hosts_percentage,omitempty"`
	// The Cloud Workload Security host usage by tag(s).
	CwsHostsUsage datadog.NullableFloat64 `json:"cws_hosts_usage,omitempty"`
	// The percentage of Database Monitoring host usage by tag(s).
	DbmHostsPercentage datadog.NullableFloat64 `json:"dbm_hosts_percentage,omitempty"`
	// The Database Monitoring host usage by tag(s).
	DbmHostsUsage datadog.NullableFloat64 `json:"dbm_hosts_usage,omitempty"`
	// The percentage of Database Monitoring queries usage by tag(s).
	DbmQueriesPercentage datadog.NullableFloat64 `json:"dbm_queries_percentage,omitempty"`
	// The Database Monitoring queries usage by tag(s).
	DbmQueriesUsage datadog.NullableFloat64 `json:"dbm_queries_usage,omitempty"`
	// The percentage of estimated live indexed logs usage by tag(s).
	EstimatedIndexedLogsPercentage datadog.NullableFloat64 `json:"estimated_indexed_logs_percentage,omitempty"`
	// The estimated live indexed logs usage by tag(s).
	EstimatedIndexedLogsUsage datadog.NullableFloat64 `json:"estimated_indexed_logs_usage,omitempty"`
	// The percentage of estimated indexed spans usage by tag(s).
	EstimatedIndexedSpansPercentage datadog.NullableFloat64 `json:"estimated_indexed_spans_percentage,omitempty"`
	// The estimated indexed spans usage by tag(s).
	EstimatedIndexedSpansUsage datadog.NullableFloat64 `json:"estimated_indexed_spans_usage,omitempty"`
	// The percentage of estimated live ingested logs usage by tag(s).
	EstimatedIngestedLogsPercentage datadog.NullableFloat64 `json:"estimated_ingested_logs_percentage,omitempty"`
	// The estimated live ingested logs usage by tag(s).
	EstimatedIngestedLogsUsage datadog.NullableFloat64 `json:"estimated_ingested_logs_usage,omitempty"`
	// The percentage of estimated ingested spans usage by tag(s).
	EstimatedIngestedSpansPercentage datadog.NullableFloat64 `json:"estimated_ingested_spans_percentage,omitempty"`
	// The estimated ingested spans usage by tag(s).
	EstimatedIngestedSpansUsage datadog.NullableFloat64 `json:"estimated_ingested_spans_usage,omitempty"`
	// The percentage of estimated rum sessions usage by tag(s).
	EstimatedRumSessionsPercentage datadog.NullableFloat64 `json:"estimated_rum_sessions_percentage,omitempty"`
	// The estimated rum sessions usage by tag(s).
	EstimatedRumSessionsUsage datadog.NullableFloat64 `json:"estimated_rum_sessions_usage,omitempty"`
	// The percentage of Fargate usage by tags.
	FargatePercentage datadog.NullableFloat64 `json:"fargate_percentage,omitempty"`
	// The Fargate usage by tags.
	FargateUsage datadog.NullableFloat64 `json:"fargate_usage,omitempty"`
	// The percentage of Lambda function usage by tag(s).
	FunctionsPercentage datadog.NullableFloat64 `json:"functions_percentage,omitempty"`
	// The Lambda function usage by tag(s).
	FunctionsUsage datadog.NullableFloat64 `json:"functions_usage,omitempty"`
	// The percentage of infrastructure host usage by tag(s).
	InfraHostPercentage datadog.NullableFloat64 `json:"infra_host_percentage,omitempty"`
	// The infrastructure host usage by tag(s).
	InfraHostUsage datadog.NullableFloat64 `json:"infra_host_usage,omitempty"`
	// The percentage of Lambda invocation usage by tag(s).
	InvocationsPercentage datadog.NullableFloat64 `json:"invocations_percentage,omitempty"`
	// The Lambda invocation usage by tag(s).
	InvocationsUsage datadog.NullableFloat64 `json:"invocations_usage,omitempty"`
	// The percentage of network host usage by tag(s).
	NpmHostPercentage datadog.NullableFloat64 `json:"npm_host_percentage,omitempty"`
	// The network host usage by tag(s).
	NpmHostUsage datadog.NullableFloat64 `json:"npm_host_usage,omitempty"`
	// The percentage of profiled container usage by tag(s).
	ProfiledContainerPercentage datadog.NullableFloat64 `json:"profiled_container_percentage,omitempty"`
	// The profiled container usage by tag(s).
	ProfiledContainerUsage datadog.NullableFloat64 `json:"profiled_container_usage,omitempty"`
	// The percentage of profiled Fargate task usage by tag(s).
	ProfiledFargatePercentage datadog.NullableFloat64 `json:"profiled_fargate_percentage,omitempty"`
	// The profiled Fargate task usage by tag(s).
	ProfiledFargateUsage datadog.NullableFloat64 `json:"profiled_fargate_usage,omitempty"`
	// The percentage of profiled hosts usage by tag(s).
	ProfiledHostPercentage datadog.NullableFloat64 `json:"profiled_host_percentage,omitempty"`
	// The profiled hosts usage by tag(s).
	ProfiledHostUsage datadog.NullableFloat64 `json:"profiled_host_usage,omitempty"`
	// The percentage of network device usage by tag(s).
	SnmpPercentage datadog.NullableFloat64 `json:"snmp_percentage,omitempty"`
	// The network device usage by tag(s).
	SnmpUsage datadog.NullableFloat64 `json:"snmp_usage,omitempty"`
	// The percentage of universal service monitoring usage by tag(s).
	UniversalServiceMonitoringPercentage datadog.NullableFloat64 `json:"universal_service_monitoring_percentage,omitempty"`
	// The universal service monitoring usage by tag(s).
	UniversalServiceMonitoringUsage datadog.NullableFloat64 `json:"universal_service_monitoring_usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMonthlyUsageAttributionValues instantiates a new MonthlyUsageAttributionValues object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyUsageAttributionValues() *MonthlyUsageAttributionValues {
	this := MonthlyUsageAttributionValues{}
	return &this
}

// NewMonthlyUsageAttributionValuesWithDefaults instantiates a new MonthlyUsageAttributionValues object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyUsageAttributionValuesWithDefaults() *MonthlyUsageAttributionValues {
	this := MonthlyUsageAttributionValues{}
	return &this
}

// GetApiPercentage returns the ApiPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetApiPercentage() float64 {
	if o == nil || o.ApiPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApiPercentage.Get()
}

// GetApiPercentageOk returns a tuple with the ApiPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetApiPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiPercentage.Get(), o.ApiPercentage.IsSet()
}

// HasApiPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApiPercentage() bool {
	return o != nil && o.ApiPercentage.IsSet()
}

// SetApiPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApiPercentage field.
func (o *MonthlyUsageAttributionValues) SetApiPercentage(v float64) {
	o.ApiPercentage.Set(&v)
}

// SetApiPercentageNil sets the value for ApiPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetApiPercentageNil() {
	o.ApiPercentage.Set(nil)
}

// UnsetApiPercentage ensures that no value is present for ApiPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetApiPercentage() {
	o.ApiPercentage.Unset()
}

// GetApiUsage returns the ApiUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetApiUsage() float64 {
	if o == nil || o.ApiUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApiUsage.Get()
}

// GetApiUsageOk returns a tuple with the ApiUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetApiUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiUsage.Get(), o.ApiUsage.IsSet()
}

// HasApiUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApiUsage() bool {
	return o != nil && o.ApiUsage.IsSet()
}

// SetApiUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApiUsage field.
func (o *MonthlyUsageAttributionValues) SetApiUsage(v float64) {
	o.ApiUsage.Set(&v)
}

// SetApiUsageNil sets the value for ApiUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetApiUsageNil() {
	o.ApiUsage.Set(nil)
}

// UnsetApiUsage ensures that no value is present for ApiUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetApiUsage() {
	o.ApiUsage.Unset()
}

// GetApmFargatePercentage returns the ApmFargatePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetApmFargatePercentage() float64 {
	if o == nil || o.ApmFargatePercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApmFargatePercentage.Get()
}

// GetApmFargatePercentageOk returns a tuple with the ApmFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetApmFargatePercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmFargatePercentage.Get(), o.ApmFargatePercentage.IsSet()
}

// HasApmFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmFargatePercentage() bool {
	return o != nil && o.ApmFargatePercentage.IsSet()
}

// SetApmFargatePercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApmFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetApmFargatePercentage(v float64) {
	o.ApmFargatePercentage.Set(&v)
}

// SetApmFargatePercentageNil sets the value for ApmFargatePercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetApmFargatePercentageNil() {
	o.ApmFargatePercentage.Set(nil)
}

// UnsetApmFargatePercentage ensures that no value is present for ApmFargatePercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetApmFargatePercentage() {
	o.ApmFargatePercentage.Unset()
}

// GetApmFargateUsage returns the ApmFargateUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetApmFargateUsage() float64 {
	if o == nil || o.ApmFargateUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApmFargateUsage.Get()
}

// GetApmFargateUsageOk returns a tuple with the ApmFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetApmFargateUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmFargateUsage.Get(), o.ApmFargateUsage.IsSet()
}

// HasApmFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmFargateUsage() bool {
	return o != nil && o.ApmFargateUsage.IsSet()
}

// SetApmFargateUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApmFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetApmFargateUsage(v float64) {
	o.ApmFargateUsage.Set(&v)
}

// SetApmFargateUsageNil sets the value for ApmFargateUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetApmFargateUsageNil() {
	o.ApmFargateUsage.Set(nil)
}

// UnsetApmFargateUsage ensures that no value is present for ApmFargateUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetApmFargateUsage() {
	o.ApmFargateUsage.Unset()
}

// GetApmHostPercentage returns the ApmHostPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetApmHostPercentage() float64 {
	if o == nil || o.ApmHostPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApmHostPercentage.Get()
}

// GetApmHostPercentageOk returns a tuple with the ApmHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetApmHostPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmHostPercentage.Get(), o.ApmHostPercentage.IsSet()
}

// HasApmHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmHostPercentage() bool {
	return o != nil && o.ApmHostPercentage.IsSet()
}

// SetApmHostPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApmHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetApmHostPercentage(v float64) {
	o.ApmHostPercentage.Set(&v)
}

// SetApmHostPercentageNil sets the value for ApmHostPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetApmHostPercentageNil() {
	o.ApmHostPercentage.Set(nil)
}

// UnsetApmHostPercentage ensures that no value is present for ApmHostPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetApmHostPercentage() {
	o.ApmHostPercentage.Unset()
}

// GetApmHostUsage returns the ApmHostUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetApmHostUsage() float64 {
	if o == nil || o.ApmHostUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApmHostUsage.Get()
}

// GetApmHostUsageOk returns a tuple with the ApmHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetApmHostUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmHostUsage.Get(), o.ApmHostUsage.IsSet()
}

// HasApmHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmHostUsage() bool {
	return o != nil && o.ApmHostUsage.IsSet()
}

// SetApmHostUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApmHostUsage field.
func (o *MonthlyUsageAttributionValues) SetApmHostUsage(v float64) {
	o.ApmHostUsage.Set(&v)
}

// SetApmHostUsageNil sets the value for ApmHostUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetApmHostUsageNil() {
	o.ApmHostUsage.Set(nil)
}

// UnsetApmHostUsage ensures that no value is present for ApmHostUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetApmHostUsage() {
	o.ApmHostUsage.Unset()
}

// GetAppsecFargatePercentage returns the AppsecFargatePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetAppsecFargatePercentage() float64 {
	if o == nil || o.AppsecFargatePercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.AppsecFargatePercentage.Get()
}

// GetAppsecFargatePercentageOk returns a tuple with the AppsecFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetAppsecFargatePercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecFargatePercentage.Get(), o.AppsecFargatePercentage.IsSet()
}

// HasAppsecFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecFargatePercentage() bool {
	return o != nil && o.AppsecFargatePercentage.IsSet()
}

// SetAppsecFargatePercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the AppsecFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetAppsecFargatePercentage(v float64) {
	o.AppsecFargatePercentage.Set(&v)
}

// SetAppsecFargatePercentageNil sets the value for AppsecFargatePercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetAppsecFargatePercentageNil() {
	o.AppsecFargatePercentage.Set(nil)
}

// UnsetAppsecFargatePercentage ensures that no value is present for AppsecFargatePercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetAppsecFargatePercentage() {
	o.AppsecFargatePercentage.Unset()
}

// GetAppsecFargateUsage returns the AppsecFargateUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetAppsecFargateUsage() float64 {
	if o == nil || o.AppsecFargateUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.AppsecFargateUsage.Get()
}

// GetAppsecFargateUsageOk returns a tuple with the AppsecFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetAppsecFargateUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecFargateUsage.Get(), o.AppsecFargateUsage.IsSet()
}

// HasAppsecFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecFargateUsage() bool {
	return o != nil && o.AppsecFargateUsage.IsSet()
}

// SetAppsecFargateUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the AppsecFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetAppsecFargateUsage(v float64) {
	o.AppsecFargateUsage.Set(&v)
}

// SetAppsecFargateUsageNil sets the value for AppsecFargateUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetAppsecFargateUsageNil() {
	o.AppsecFargateUsage.Set(nil)
}

// UnsetAppsecFargateUsage ensures that no value is present for AppsecFargateUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetAppsecFargateUsage() {
	o.AppsecFargateUsage.Unset()
}

// GetAppsecPercentage returns the AppsecPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetAppsecPercentage() float64 {
	if o == nil || o.AppsecPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.AppsecPercentage.Get()
}

// GetAppsecPercentageOk returns a tuple with the AppsecPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetAppsecPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecPercentage.Get(), o.AppsecPercentage.IsSet()
}

// HasAppsecPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecPercentage() bool {
	return o != nil && o.AppsecPercentage.IsSet()
}

// SetAppsecPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the AppsecPercentage field.
func (o *MonthlyUsageAttributionValues) SetAppsecPercentage(v float64) {
	o.AppsecPercentage.Set(&v)
}

// SetAppsecPercentageNil sets the value for AppsecPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetAppsecPercentageNil() {
	o.AppsecPercentage.Set(nil)
}

// UnsetAppsecPercentage ensures that no value is present for AppsecPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetAppsecPercentage() {
	o.AppsecPercentage.Unset()
}

// GetAppsecUsage returns the AppsecUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetAppsecUsage() float64 {
	if o == nil || o.AppsecUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.AppsecUsage.Get()
}

// GetAppsecUsageOk returns a tuple with the AppsecUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetAppsecUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecUsage.Get(), o.AppsecUsage.IsSet()
}

// HasAppsecUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecUsage() bool {
	return o != nil && o.AppsecUsage.IsSet()
}

// SetAppsecUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the AppsecUsage field.
func (o *MonthlyUsageAttributionValues) SetAppsecUsage(v float64) {
	o.AppsecUsage.Set(&v)
}

// SetAppsecUsageNil sets the value for AppsecUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetAppsecUsageNil() {
	o.AppsecUsage.Set(nil)
}

// UnsetAppsecUsage ensures that no value is present for AppsecUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetAppsecUsage() {
	o.AppsecUsage.Unset()
}

// GetBrowserPercentage returns the BrowserPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetBrowserPercentage() float64 {
	if o == nil || o.BrowserPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.BrowserPercentage.Get()
}

// GetBrowserPercentageOk returns a tuple with the BrowserPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetBrowserPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserPercentage.Get(), o.BrowserPercentage.IsSet()
}

// HasBrowserPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasBrowserPercentage() bool {
	return o != nil && o.BrowserPercentage.IsSet()
}

// SetBrowserPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the BrowserPercentage field.
func (o *MonthlyUsageAttributionValues) SetBrowserPercentage(v float64) {
	o.BrowserPercentage.Set(&v)
}

// SetBrowserPercentageNil sets the value for BrowserPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetBrowserPercentageNil() {
	o.BrowserPercentage.Set(nil)
}

// UnsetBrowserPercentage ensures that no value is present for BrowserPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetBrowserPercentage() {
	o.BrowserPercentage.Unset()
}

// GetBrowserUsage returns the BrowserUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetBrowserUsage() float64 {
	if o == nil || o.BrowserUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.BrowserUsage.Get()
}

// GetBrowserUsageOk returns a tuple with the BrowserUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetBrowserUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserUsage.Get(), o.BrowserUsage.IsSet()
}

// HasBrowserUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasBrowserUsage() bool {
	return o != nil && o.BrowserUsage.IsSet()
}

// SetBrowserUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the BrowserUsage field.
func (o *MonthlyUsageAttributionValues) SetBrowserUsage(v float64) {
	o.BrowserUsage.Set(&v)
}

// SetBrowserUsageNil sets the value for BrowserUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetBrowserUsageNil() {
	o.BrowserUsage.Set(nil)
}

// UnsetBrowserUsage ensures that no value is present for BrowserUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetBrowserUsage() {
	o.BrowserUsage.Unset()
}

// GetContainerExclAgentPercentage returns the ContainerExclAgentPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentPercentage() float64 {
	if o == nil || o.ContainerExclAgentPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ContainerExclAgentPercentage.Get()
}

// GetContainerExclAgentPercentageOk returns a tuple with the ContainerExclAgentPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerExclAgentPercentage.Get(), o.ContainerExclAgentPercentage.IsSet()
}

// HasContainerExclAgentPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerExclAgentPercentage() bool {
	return o != nil && o.ContainerExclAgentPercentage.IsSet()
}

// SetContainerExclAgentPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ContainerExclAgentPercentage field.
func (o *MonthlyUsageAttributionValues) SetContainerExclAgentPercentage(v float64) {
	o.ContainerExclAgentPercentage.Set(&v)
}

// SetContainerExclAgentPercentageNil sets the value for ContainerExclAgentPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetContainerExclAgentPercentageNil() {
	o.ContainerExclAgentPercentage.Set(nil)
}

// UnsetContainerExclAgentPercentage ensures that no value is present for ContainerExclAgentPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetContainerExclAgentPercentage() {
	o.ContainerExclAgentPercentage.Unset()
}

// GetContainerExclAgentUsage returns the ContainerExclAgentUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentUsage() float64 {
	if o == nil || o.ContainerExclAgentUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ContainerExclAgentUsage.Get()
}

// GetContainerExclAgentUsageOk returns a tuple with the ContainerExclAgentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerExclAgentUsage.Get(), o.ContainerExclAgentUsage.IsSet()
}

// HasContainerExclAgentUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerExclAgentUsage() bool {
	return o != nil && o.ContainerExclAgentUsage.IsSet()
}

// SetContainerExclAgentUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ContainerExclAgentUsage field.
func (o *MonthlyUsageAttributionValues) SetContainerExclAgentUsage(v float64) {
	o.ContainerExclAgentUsage.Set(&v)
}

// SetContainerExclAgentUsageNil sets the value for ContainerExclAgentUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetContainerExclAgentUsageNil() {
	o.ContainerExclAgentUsage.Set(nil)
}

// UnsetContainerExclAgentUsage ensures that no value is present for ContainerExclAgentUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetContainerExclAgentUsage() {
	o.ContainerExclAgentUsage.Unset()
}

// GetContainerPercentage returns the ContainerPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetContainerPercentage() float64 {
	if o == nil || o.ContainerPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ContainerPercentage.Get()
}

// GetContainerPercentageOk returns a tuple with the ContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetContainerPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerPercentage.Get(), o.ContainerPercentage.IsSet()
}

// HasContainerPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerPercentage() bool {
	return o != nil && o.ContainerPercentage.IsSet()
}

// SetContainerPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ContainerPercentage field.
func (o *MonthlyUsageAttributionValues) SetContainerPercentage(v float64) {
	o.ContainerPercentage.Set(&v)
}

// SetContainerPercentageNil sets the value for ContainerPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetContainerPercentageNil() {
	o.ContainerPercentage.Set(nil)
}

// UnsetContainerPercentage ensures that no value is present for ContainerPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetContainerPercentage() {
	o.ContainerPercentage.Unset()
}

// GetContainerUsage returns the ContainerUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetContainerUsage() float64 {
	if o == nil || o.ContainerUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ContainerUsage.Get()
}

// GetContainerUsageOk returns a tuple with the ContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetContainerUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerUsage.Get(), o.ContainerUsage.IsSet()
}

// HasContainerUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerUsage() bool {
	return o != nil && o.ContainerUsage.IsSet()
}

// SetContainerUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ContainerUsage field.
func (o *MonthlyUsageAttributionValues) SetContainerUsage(v float64) {
	o.ContainerUsage.Set(&v)
}

// SetContainerUsageNil sets the value for ContainerUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetContainerUsageNil() {
	o.ContainerUsage.Set(nil)
}

// UnsetContainerUsage ensures that no value is present for ContainerUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetContainerUsage() {
	o.ContainerUsage.Unset()
}

// GetCspmContainersPercentage returns the CspmContainersPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCspmContainersPercentage() float64 {
	if o == nil || o.CspmContainersPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CspmContainersPercentage.Get()
}

// GetCspmContainersPercentageOk returns a tuple with the CspmContainersPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCspmContainersPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainersPercentage.Get(), o.CspmContainersPercentage.IsSet()
}

// HasCspmContainersPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmContainersPercentage() bool {
	return o != nil && o.CspmContainersPercentage.IsSet()
}

// SetCspmContainersPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CspmContainersPercentage field.
func (o *MonthlyUsageAttributionValues) SetCspmContainersPercentage(v float64) {
	o.CspmContainersPercentage.Set(&v)
}

// SetCspmContainersPercentageNil sets the value for CspmContainersPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCspmContainersPercentageNil() {
	o.CspmContainersPercentage.Set(nil)
}

// UnsetCspmContainersPercentage ensures that no value is present for CspmContainersPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCspmContainersPercentage() {
	o.CspmContainersPercentage.Unset()
}

// GetCspmContainersUsage returns the CspmContainersUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCspmContainersUsage() float64 {
	if o == nil || o.CspmContainersUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CspmContainersUsage.Get()
}

// GetCspmContainersUsageOk returns a tuple with the CspmContainersUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCspmContainersUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainersUsage.Get(), o.CspmContainersUsage.IsSet()
}

// HasCspmContainersUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmContainersUsage() bool {
	return o != nil && o.CspmContainersUsage.IsSet()
}

// SetCspmContainersUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CspmContainersUsage field.
func (o *MonthlyUsageAttributionValues) SetCspmContainersUsage(v float64) {
	o.CspmContainersUsage.Set(&v)
}

// SetCspmContainersUsageNil sets the value for CspmContainersUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCspmContainersUsageNil() {
	o.CspmContainersUsage.Set(nil)
}

// UnsetCspmContainersUsage ensures that no value is present for CspmContainersUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCspmContainersUsage() {
	o.CspmContainersUsage.Unset()
}

// GetCspmHostsPercentage returns the CspmHostsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCspmHostsPercentage() float64 {
	if o == nil || o.CspmHostsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CspmHostsPercentage.Get()
}

// GetCspmHostsPercentageOk returns a tuple with the CspmHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCspmHostsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmHostsPercentage.Get(), o.CspmHostsPercentage.IsSet()
}

// HasCspmHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmHostsPercentage() bool {
	return o != nil && o.CspmHostsPercentage.IsSet()
}

// SetCspmHostsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CspmHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetCspmHostsPercentage(v float64) {
	o.CspmHostsPercentage.Set(&v)
}

// SetCspmHostsPercentageNil sets the value for CspmHostsPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCspmHostsPercentageNil() {
	o.CspmHostsPercentage.Set(nil)
}

// UnsetCspmHostsPercentage ensures that no value is present for CspmHostsPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCspmHostsPercentage() {
	o.CspmHostsPercentage.Unset()
}

// GetCspmHostsUsage returns the CspmHostsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCspmHostsUsage() float64 {
	if o == nil || o.CspmHostsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CspmHostsUsage.Get()
}

// GetCspmHostsUsageOk returns a tuple with the CspmHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCspmHostsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmHostsUsage.Get(), o.CspmHostsUsage.IsSet()
}

// HasCspmHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmHostsUsage() bool {
	return o != nil && o.CspmHostsUsage.IsSet()
}

// SetCspmHostsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CspmHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetCspmHostsUsage(v float64) {
	o.CspmHostsUsage.Set(&v)
}

// SetCspmHostsUsageNil sets the value for CspmHostsUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCspmHostsUsageNil() {
	o.CspmHostsUsage.Set(nil)
}

// UnsetCspmHostsUsage ensures that no value is present for CspmHostsUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCspmHostsUsage() {
	o.CspmHostsUsage.Unset()
}

// GetCustomIngestedTimeseriesPercentage returns the CustomIngestedTimeseriesPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesPercentage() float64 {
	if o == nil || o.CustomIngestedTimeseriesPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CustomIngestedTimeseriesPercentage.Get()
}

// GetCustomIngestedTimeseriesPercentageOk returns a tuple with the CustomIngestedTimeseriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomIngestedTimeseriesPercentage.Get(), o.CustomIngestedTimeseriesPercentage.IsSet()
}

// HasCustomIngestedTimeseriesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomIngestedTimeseriesPercentage() bool {
	return o != nil && o.CustomIngestedTimeseriesPercentage.IsSet()
}

// SetCustomIngestedTimeseriesPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CustomIngestedTimeseriesPercentage field.
func (o *MonthlyUsageAttributionValues) SetCustomIngestedTimeseriesPercentage(v float64) {
	o.CustomIngestedTimeseriesPercentage.Set(&v)
}

// SetCustomIngestedTimeseriesPercentageNil sets the value for CustomIngestedTimeseriesPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCustomIngestedTimeseriesPercentageNil() {
	o.CustomIngestedTimeseriesPercentage.Set(nil)
}

// UnsetCustomIngestedTimeseriesPercentage ensures that no value is present for CustomIngestedTimeseriesPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCustomIngestedTimeseriesPercentage() {
	o.CustomIngestedTimeseriesPercentage.Unset()
}

// GetCustomIngestedTimeseriesUsage returns the CustomIngestedTimeseriesUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesUsage() float64 {
	if o == nil || o.CustomIngestedTimeseriesUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CustomIngestedTimeseriesUsage.Get()
}

// GetCustomIngestedTimeseriesUsageOk returns a tuple with the CustomIngestedTimeseriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomIngestedTimeseriesUsage.Get(), o.CustomIngestedTimeseriesUsage.IsSet()
}

// HasCustomIngestedTimeseriesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomIngestedTimeseriesUsage() bool {
	return o != nil && o.CustomIngestedTimeseriesUsage.IsSet()
}

// SetCustomIngestedTimeseriesUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CustomIngestedTimeseriesUsage field.
func (o *MonthlyUsageAttributionValues) SetCustomIngestedTimeseriesUsage(v float64) {
	o.CustomIngestedTimeseriesUsage.Set(&v)
}

// SetCustomIngestedTimeseriesUsageNil sets the value for CustomIngestedTimeseriesUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCustomIngestedTimeseriesUsageNil() {
	o.CustomIngestedTimeseriesUsage.Set(nil)
}

// UnsetCustomIngestedTimeseriesUsage ensures that no value is present for CustomIngestedTimeseriesUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCustomIngestedTimeseriesUsage() {
	o.CustomIngestedTimeseriesUsage.Unset()
}

// GetCustomTimeseriesPercentage returns the CustomTimeseriesPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesPercentage() float64 {
	if o == nil || o.CustomTimeseriesPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CustomTimeseriesPercentage.Get()
}

// GetCustomTimeseriesPercentageOk returns a tuple with the CustomTimeseriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTimeseriesPercentage.Get(), o.CustomTimeseriesPercentage.IsSet()
}

// HasCustomTimeseriesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomTimeseriesPercentage() bool {
	return o != nil && o.CustomTimeseriesPercentage.IsSet()
}

// SetCustomTimeseriesPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CustomTimeseriesPercentage field.
func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesPercentage(v float64) {
	o.CustomTimeseriesPercentage.Set(&v)
}

// SetCustomTimeseriesPercentageNil sets the value for CustomTimeseriesPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesPercentageNil() {
	o.CustomTimeseriesPercentage.Set(nil)
}

// UnsetCustomTimeseriesPercentage ensures that no value is present for CustomTimeseriesPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCustomTimeseriesPercentage() {
	o.CustomTimeseriesPercentage.Unset()
}

// GetCustomTimeseriesUsage returns the CustomTimeseriesUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesUsage() float64 {
	if o == nil || o.CustomTimeseriesUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CustomTimeseriesUsage.Get()
}

// GetCustomTimeseriesUsageOk returns a tuple with the CustomTimeseriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTimeseriesUsage.Get(), o.CustomTimeseriesUsage.IsSet()
}

// HasCustomTimeseriesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomTimeseriesUsage() bool {
	return o != nil && o.CustomTimeseriesUsage.IsSet()
}

// SetCustomTimeseriesUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CustomTimeseriesUsage field.
func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesUsage(v float64) {
	o.CustomTimeseriesUsage.Set(&v)
}

// SetCustomTimeseriesUsageNil sets the value for CustomTimeseriesUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesUsageNil() {
	o.CustomTimeseriesUsage.Set(nil)
}

// UnsetCustomTimeseriesUsage ensures that no value is present for CustomTimeseriesUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCustomTimeseriesUsage() {
	o.CustomTimeseriesUsage.Unset()
}

// GetCwsContainersPercentage returns the CwsContainersPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCwsContainersPercentage() float64 {
	if o == nil || o.CwsContainersPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CwsContainersPercentage.Get()
}

// GetCwsContainersPercentageOk returns a tuple with the CwsContainersPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCwsContainersPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsContainersPercentage.Get(), o.CwsContainersPercentage.IsSet()
}

// HasCwsContainersPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsContainersPercentage() bool {
	return o != nil && o.CwsContainersPercentage.IsSet()
}

// SetCwsContainersPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CwsContainersPercentage field.
func (o *MonthlyUsageAttributionValues) SetCwsContainersPercentage(v float64) {
	o.CwsContainersPercentage.Set(&v)
}

// SetCwsContainersPercentageNil sets the value for CwsContainersPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCwsContainersPercentageNil() {
	o.CwsContainersPercentage.Set(nil)
}

// UnsetCwsContainersPercentage ensures that no value is present for CwsContainersPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCwsContainersPercentage() {
	o.CwsContainersPercentage.Unset()
}

// GetCwsContainersUsage returns the CwsContainersUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCwsContainersUsage() float64 {
	if o == nil || o.CwsContainersUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CwsContainersUsage.Get()
}

// GetCwsContainersUsageOk returns a tuple with the CwsContainersUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCwsContainersUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsContainersUsage.Get(), o.CwsContainersUsage.IsSet()
}

// HasCwsContainersUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsContainersUsage() bool {
	return o != nil && o.CwsContainersUsage.IsSet()
}

// SetCwsContainersUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CwsContainersUsage field.
func (o *MonthlyUsageAttributionValues) SetCwsContainersUsage(v float64) {
	o.CwsContainersUsage.Set(&v)
}

// SetCwsContainersUsageNil sets the value for CwsContainersUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCwsContainersUsageNil() {
	o.CwsContainersUsage.Set(nil)
}

// UnsetCwsContainersUsage ensures that no value is present for CwsContainersUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCwsContainersUsage() {
	o.CwsContainersUsage.Unset()
}

// GetCwsHostsPercentage returns the CwsHostsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCwsHostsPercentage() float64 {
	if o == nil || o.CwsHostsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CwsHostsPercentage.Get()
}

// GetCwsHostsPercentageOk returns a tuple with the CwsHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCwsHostsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsHostsPercentage.Get(), o.CwsHostsPercentage.IsSet()
}

// HasCwsHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsHostsPercentage() bool {
	return o != nil && o.CwsHostsPercentage.IsSet()
}

// SetCwsHostsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CwsHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetCwsHostsPercentage(v float64) {
	o.CwsHostsPercentage.Set(&v)
}

// SetCwsHostsPercentageNil sets the value for CwsHostsPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCwsHostsPercentageNil() {
	o.CwsHostsPercentage.Set(nil)
}

// UnsetCwsHostsPercentage ensures that no value is present for CwsHostsPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCwsHostsPercentage() {
	o.CwsHostsPercentage.Unset()
}

// GetCwsHostsUsage returns the CwsHostsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetCwsHostsUsage() float64 {
	if o == nil || o.CwsHostsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CwsHostsUsage.Get()
}

// GetCwsHostsUsageOk returns a tuple with the CwsHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetCwsHostsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsHostsUsage.Get(), o.CwsHostsUsage.IsSet()
}

// HasCwsHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsHostsUsage() bool {
	return o != nil && o.CwsHostsUsage.IsSet()
}

// SetCwsHostsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CwsHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetCwsHostsUsage(v float64) {
	o.CwsHostsUsage.Set(&v)
}

// SetCwsHostsUsageNil sets the value for CwsHostsUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetCwsHostsUsageNil() {
	o.CwsHostsUsage.Set(nil)
}

// UnsetCwsHostsUsage ensures that no value is present for CwsHostsUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetCwsHostsUsage() {
	o.CwsHostsUsage.Unset()
}

// GetDbmHostsPercentage returns the DbmHostsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetDbmHostsPercentage() float64 {
	if o == nil || o.DbmHostsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DbmHostsPercentage.Get()
}

// GetDbmHostsPercentageOk returns a tuple with the DbmHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetDbmHostsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmHostsPercentage.Get(), o.DbmHostsPercentage.IsSet()
}

// HasDbmHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmHostsPercentage() bool {
	return o != nil && o.DbmHostsPercentage.IsSet()
}

// SetDbmHostsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the DbmHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetDbmHostsPercentage(v float64) {
	o.DbmHostsPercentage.Set(&v)
}

// SetDbmHostsPercentageNil sets the value for DbmHostsPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetDbmHostsPercentageNil() {
	o.DbmHostsPercentage.Set(nil)
}

// UnsetDbmHostsPercentage ensures that no value is present for DbmHostsPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetDbmHostsPercentage() {
	o.DbmHostsPercentage.Unset()
}

// GetDbmHostsUsage returns the DbmHostsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetDbmHostsUsage() float64 {
	if o == nil || o.DbmHostsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DbmHostsUsage.Get()
}

// GetDbmHostsUsageOk returns a tuple with the DbmHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetDbmHostsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmHostsUsage.Get(), o.DbmHostsUsage.IsSet()
}

// HasDbmHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmHostsUsage() bool {
	return o != nil && o.DbmHostsUsage.IsSet()
}

// SetDbmHostsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the DbmHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetDbmHostsUsage(v float64) {
	o.DbmHostsUsage.Set(&v)
}

// SetDbmHostsUsageNil sets the value for DbmHostsUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetDbmHostsUsageNil() {
	o.DbmHostsUsage.Set(nil)
}

// UnsetDbmHostsUsage ensures that no value is present for DbmHostsUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetDbmHostsUsage() {
	o.DbmHostsUsage.Unset()
}

// GetDbmQueriesPercentage returns the DbmQueriesPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetDbmQueriesPercentage() float64 {
	if o == nil || o.DbmQueriesPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DbmQueriesPercentage.Get()
}

// GetDbmQueriesPercentageOk returns a tuple with the DbmQueriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmQueriesPercentage.Get(), o.DbmQueriesPercentage.IsSet()
}

// HasDbmQueriesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmQueriesPercentage() bool {
	return o != nil && o.DbmQueriesPercentage.IsSet()
}

// SetDbmQueriesPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the DbmQueriesPercentage field.
func (o *MonthlyUsageAttributionValues) SetDbmQueriesPercentage(v float64) {
	o.DbmQueriesPercentage.Set(&v)
}

// SetDbmQueriesPercentageNil sets the value for DbmQueriesPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetDbmQueriesPercentageNil() {
	o.DbmQueriesPercentage.Set(nil)
}

// UnsetDbmQueriesPercentage ensures that no value is present for DbmQueriesPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetDbmQueriesPercentage() {
	o.DbmQueriesPercentage.Unset()
}

// GetDbmQueriesUsage returns the DbmQueriesUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetDbmQueriesUsage() float64 {
	if o == nil || o.DbmQueriesUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DbmQueriesUsage.Get()
}

// GetDbmQueriesUsageOk returns a tuple with the DbmQueriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmQueriesUsage.Get(), o.DbmQueriesUsage.IsSet()
}

// HasDbmQueriesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmQueriesUsage() bool {
	return o != nil && o.DbmQueriesUsage.IsSet()
}

// SetDbmQueriesUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the DbmQueriesUsage field.
func (o *MonthlyUsageAttributionValues) SetDbmQueriesUsage(v float64) {
	o.DbmQueriesUsage.Set(&v)
}

// SetDbmQueriesUsageNil sets the value for DbmQueriesUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetDbmQueriesUsageNil() {
	o.DbmQueriesUsage.Set(nil)
}

// UnsetDbmQueriesUsage ensures that no value is present for DbmQueriesUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetDbmQueriesUsage() {
	o.DbmQueriesUsage.Unset()
}

// GetEstimatedIndexedLogsPercentage returns the EstimatedIndexedLogsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedLogsPercentage() float64 {
	if o == nil || o.EstimatedIndexedLogsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedLogsPercentage.Get()
}

// GetEstimatedIndexedLogsPercentageOk returns a tuple with the EstimatedIndexedLogsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedLogsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIndexedLogsPercentage.Get(), o.EstimatedIndexedLogsPercentage.IsSet()
}

// HasEstimatedIndexedLogsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedLogsPercentage() bool {
	return o != nil && o.EstimatedIndexedLogsPercentage.IsSet()
}

// SetEstimatedIndexedLogsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIndexedLogsPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedLogsPercentage(v float64) {
	o.EstimatedIndexedLogsPercentage.Set(&v)
}

// SetEstimatedIndexedLogsPercentageNil sets the value for EstimatedIndexedLogsPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedLogsPercentageNil() {
	o.EstimatedIndexedLogsPercentage.Set(nil)
}

// UnsetEstimatedIndexedLogsPercentage ensures that no value is present for EstimatedIndexedLogsPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedIndexedLogsPercentage() {
	o.EstimatedIndexedLogsPercentage.Unset()
}

// GetEstimatedIndexedLogsUsage returns the EstimatedIndexedLogsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedLogsUsage() float64 {
	if o == nil || o.EstimatedIndexedLogsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedLogsUsage.Get()
}

// GetEstimatedIndexedLogsUsageOk returns a tuple with the EstimatedIndexedLogsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedLogsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIndexedLogsUsage.Get(), o.EstimatedIndexedLogsUsage.IsSet()
}

// HasEstimatedIndexedLogsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedLogsUsage() bool {
	return o != nil && o.EstimatedIndexedLogsUsage.IsSet()
}

// SetEstimatedIndexedLogsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIndexedLogsUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedLogsUsage(v float64) {
	o.EstimatedIndexedLogsUsage.Set(&v)
}

// SetEstimatedIndexedLogsUsageNil sets the value for EstimatedIndexedLogsUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedLogsUsageNil() {
	o.EstimatedIndexedLogsUsage.Set(nil)
}

// UnsetEstimatedIndexedLogsUsage ensures that no value is present for EstimatedIndexedLogsUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedIndexedLogsUsage() {
	o.EstimatedIndexedLogsUsage.Unset()
}

// GetEstimatedIndexedSpansPercentage returns the EstimatedIndexedSpansPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansPercentage() float64 {
	if o == nil || o.EstimatedIndexedSpansPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedSpansPercentage.Get()
}

// GetEstimatedIndexedSpansPercentageOk returns a tuple with the EstimatedIndexedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIndexedSpansPercentage.Get(), o.EstimatedIndexedSpansPercentage.IsSet()
}

// HasEstimatedIndexedSpansPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedSpansPercentage() bool {
	return o != nil && o.EstimatedIndexedSpansPercentage.IsSet()
}

// SetEstimatedIndexedSpansPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIndexedSpansPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedSpansPercentage(v float64) {
	o.EstimatedIndexedSpansPercentage.Set(&v)
}

// SetEstimatedIndexedSpansPercentageNil sets the value for EstimatedIndexedSpansPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedSpansPercentageNil() {
	o.EstimatedIndexedSpansPercentage.Set(nil)
}

// UnsetEstimatedIndexedSpansPercentage ensures that no value is present for EstimatedIndexedSpansPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedIndexedSpansPercentage() {
	o.EstimatedIndexedSpansPercentage.Unset()
}

// GetEstimatedIndexedSpansUsage returns the EstimatedIndexedSpansUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansUsage() float64 {
	if o == nil || o.EstimatedIndexedSpansUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedSpansUsage.Get()
}

// GetEstimatedIndexedSpansUsageOk returns a tuple with the EstimatedIndexedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIndexedSpansUsage.Get(), o.EstimatedIndexedSpansUsage.IsSet()
}

// HasEstimatedIndexedSpansUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedSpansUsage() bool {
	return o != nil && o.EstimatedIndexedSpansUsage.IsSet()
}

// SetEstimatedIndexedSpansUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIndexedSpansUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedSpansUsage(v float64) {
	o.EstimatedIndexedSpansUsage.Set(&v)
}

// SetEstimatedIndexedSpansUsageNil sets the value for EstimatedIndexedSpansUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedSpansUsageNil() {
	o.EstimatedIndexedSpansUsage.Set(nil)
}

// UnsetEstimatedIndexedSpansUsage ensures that no value is present for EstimatedIndexedSpansUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedIndexedSpansUsage() {
	o.EstimatedIndexedSpansUsage.Unset()
}

// GetEstimatedIngestedLogsPercentage returns the EstimatedIngestedLogsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedLogsPercentage() float64 {
	if o == nil || o.EstimatedIngestedLogsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedLogsPercentage.Get()
}

// GetEstimatedIngestedLogsPercentageOk returns a tuple with the EstimatedIngestedLogsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedLogsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIngestedLogsPercentage.Get(), o.EstimatedIngestedLogsPercentage.IsSet()
}

// HasEstimatedIngestedLogsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedLogsPercentage() bool {
	return o != nil && o.EstimatedIngestedLogsPercentage.IsSet()
}

// SetEstimatedIngestedLogsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIngestedLogsPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedLogsPercentage(v float64) {
	o.EstimatedIngestedLogsPercentage.Set(&v)
}

// SetEstimatedIngestedLogsPercentageNil sets the value for EstimatedIngestedLogsPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedLogsPercentageNil() {
	o.EstimatedIngestedLogsPercentage.Set(nil)
}

// UnsetEstimatedIngestedLogsPercentage ensures that no value is present for EstimatedIngestedLogsPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedIngestedLogsPercentage() {
	o.EstimatedIngestedLogsPercentage.Unset()
}

// GetEstimatedIngestedLogsUsage returns the EstimatedIngestedLogsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedLogsUsage() float64 {
	if o == nil || o.EstimatedIngestedLogsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedLogsUsage.Get()
}

// GetEstimatedIngestedLogsUsageOk returns a tuple with the EstimatedIngestedLogsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedLogsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIngestedLogsUsage.Get(), o.EstimatedIngestedLogsUsage.IsSet()
}

// HasEstimatedIngestedLogsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedLogsUsage() bool {
	return o != nil && o.EstimatedIngestedLogsUsage.IsSet()
}

// SetEstimatedIngestedLogsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIngestedLogsUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedLogsUsage(v float64) {
	o.EstimatedIngestedLogsUsage.Set(&v)
}

// SetEstimatedIngestedLogsUsageNil sets the value for EstimatedIngestedLogsUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedLogsUsageNil() {
	o.EstimatedIngestedLogsUsage.Set(nil)
}

// UnsetEstimatedIngestedLogsUsage ensures that no value is present for EstimatedIngestedLogsUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedIngestedLogsUsage() {
	o.EstimatedIngestedLogsUsage.Unset()
}

// GetEstimatedIngestedSpansPercentage returns the EstimatedIngestedSpansPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansPercentage() float64 {
	if o == nil || o.EstimatedIngestedSpansPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedSpansPercentage.Get()
}

// GetEstimatedIngestedSpansPercentageOk returns a tuple with the EstimatedIngestedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIngestedSpansPercentage.Get(), o.EstimatedIngestedSpansPercentage.IsSet()
}

// HasEstimatedIngestedSpansPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedSpansPercentage() bool {
	return o != nil && o.EstimatedIngestedSpansPercentage.IsSet()
}

// SetEstimatedIngestedSpansPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIngestedSpansPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedSpansPercentage(v float64) {
	o.EstimatedIngestedSpansPercentage.Set(&v)
}

// SetEstimatedIngestedSpansPercentageNil sets the value for EstimatedIngestedSpansPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedSpansPercentageNil() {
	o.EstimatedIngestedSpansPercentage.Set(nil)
}

// UnsetEstimatedIngestedSpansPercentage ensures that no value is present for EstimatedIngestedSpansPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedIngestedSpansPercentage() {
	o.EstimatedIngestedSpansPercentage.Unset()
}

// GetEstimatedIngestedSpansUsage returns the EstimatedIngestedSpansUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansUsage() float64 {
	if o == nil || o.EstimatedIngestedSpansUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedSpansUsage.Get()
}

// GetEstimatedIngestedSpansUsageOk returns a tuple with the EstimatedIngestedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIngestedSpansUsage.Get(), o.EstimatedIngestedSpansUsage.IsSet()
}

// HasEstimatedIngestedSpansUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedSpansUsage() bool {
	return o != nil && o.EstimatedIngestedSpansUsage.IsSet()
}

// SetEstimatedIngestedSpansUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIngestedSpansUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedSpansUsage(v float64) {
	o.EstimatedIngestedSpansUsage.Set(&v)
}

// SetEstimatedIngestedSpansUsageNil sets the value for EstimatedIngestedSpansUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedSpansUsageNil() {
	o.EstimatedIngestedSpansUsage.Set(nil)
}

// UnsetEstimatedIngestedSpansUsage ensures that no value is present for EstimatedIngestedSpansUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedIngestedSpansUsage() {
	o.EstimatedIngestedSpansUsage.Unset()
}

// GetEstimatedRumSessionsPercentage returns the EstimatedRumSessionsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedRumSessionsPercentage() float64 {
	if o == nil || o.EstimatedRumSessionsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedRumSessionsPercentage.Get()
}

// GetEstimatedRumSessionsPercentageOk returns a tuple with the EstimatedRumSessionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedRumSessionsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedRumSessionsPercentage.Get(), o.EstimatedRumSessionsPercentage.IsSet()
}

// HasEstimatedRumSessionsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedRumSessionsPercentage() bool {
	return o != nil && o.EstimatedRumSessionsPercentage.IsSet()
}

// SetEstimatedRumSessionsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedRumSessionsPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedRumSessionsPercentage(v float64) {
	o.EstimatedRumSessionsPercentage.Set(&v)
}

// SetEstimatedRumSessionsPercentageNil sets the value for EstimatedRumSessionsPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedRumSessionsPercentageNil() {
	o.EstimatedRumSessionsPercentage.Set(nil)
}

// UnsetEstimatedRumSessionsPercentage ensures that no value is present for EstimatedRumSessionsPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedRumSessionsPercentage() {
	o.EstimatedRumSessionsPercentage.Unset()
}

// GetEstimatedRumSessionsUsage returns the EstimatedRumSessionsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetEstimatedRumSessionsUsage() float64 {
	if o == nil || o.EstimatedRumSessionsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedRumSessionsUsage.Get()
}

// GetEstimatedRumSessionsUsageOk returns a tuple with the EstimatedRumSessionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetEstimatedRumSessionsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedRumSessionsUsage.Get(), o.EstimatedRumSessionsUsage.IsSet()
}

// HasEstimatedRumSessionsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedRumSessionsUsage() bool {
	return o != nil && o.EstimatedRumSessionsUsage.IsSet()
}

// SetEstimatedRumSessionsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedRumSessionsUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedRumSessionsUsage(v float64) {
	o.EstimatedRumSessionsUsage.Set(&v)
}

// SetEstimatedRumSessionsUsageNil sets the value for EstimatedRumSessionsUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetEstimatedRumSessionsUsageNil() {
	o.EstimatedRumSessionsUsage.Set(nil)
}

// UnsetEstimatedRumSessionsUsage ensures that no value is present for EstimatedRumSessionsUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetEstimatedRumSessionsUsage() {
	o.EstimatedRumSessionsUsage.Unset()
}

// GetFargatePercentage returns the FargatePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetFargatePercentage() float64 {
	if o == nil || o.FargatePercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FargatePercentage.Get()
}

// GetFargatePercentageOk returns a tuple with the FargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetFargatePercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FargatePercentage.Get(), o.FargatePercentage.IsSet()
}

// HasFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFargatePercentage() bool {
	return o != nil && o.FargatePercentage.IsSet()
}

// SetFargatePercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the FargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetFargatePercentage(v float64) {
	o.FargatePercentage.Set(&v)
}

// SetFargatePercentageNil sets the value for FargatePercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetFargatePercentageNil() {
	o.FargatePercentage.Set(nil)
}

// UnsetFargatePercentage ensures that no value is present for FargatePercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetFargatePercentage() {
	o.FargatePercentage.Unset()
}

// GetFargateUsage returns the FargateUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetFargateUsage() float64 {
	if o == nil || o.FargateUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FargateUsage.Get()
}

// GetFargateUsageOk returns a tuple with the FargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetFargateUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FargateUsage.Get(), o.FargateUsage.IsSet()
}

// HasFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFargateUsage() bool {
	return o != nil && o.FargateUsage.IsSet()
}

// SetFargateUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the FargateUsage field.
func (o *MonthlyUsageAttributionValues) SetFargateUsage(v float64) {
	o.FargateUsage.Set(&v)
}

// SetFargateUsageNil sets the value for FargateUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetFargateUsageNil() {
	o.FargateUsage.Set(nil)
}

// UnsetFargateUsage ensures that no value is present for FargateUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetFargateUsage() {
	o.FargateUsage.Unset()
}

// GetFunctionsPercentage returns the FunctionsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetFunctionsPercentage() float64 {
	if o == nil || o.FunctionsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FunctionsPercentage.Get()
}

// GetFunctionsPercentageOk returns a tuple with the FunctionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetFunctionsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FunctionsPercentage.Get(), o.FunctionsPercentage.IsSet()
}

// HasFunctionsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFunctionsPercentage() bool {
	return o != nil && o.FunctionsPercentage.IsSet()
}

// SetFunctionsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the FunctionsPercentage field.
func (o *MonthlyUsageAttributionValues) SetFunctionsPercentage(v float64) {
	o.FunctionsPercentage.Set(&v)
}

// SetFunctionsPercentageNil sets the value for FunctionsPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetFunctionsPercentageNil() {
	o.FunctionsPercentage.Set(nil)
}

// UnsetFunctionsPercentage ensures that no value is present for FunctionsPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetFunctionsPercentage() {
	o.FunctionsPercentage.Unset()
}

// GetFunctionsUsage returns the FunctionsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetFunctionsUsage() float64 {
	if o == nil || o.FunctionsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FunctionsUsage.Get()
}

// GetFunctionsUsageOk returns a tuple with the FunctionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetFunctionsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FunctionsUsage.Get(), o.FunctionsUsage.IsSet()
}

// HasFunctionsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFunctionsUsage() bool {
	return o != nil && o.FunctionsUsage.IsSet()
}

// SetFunctionsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the FunctionsUsage field.
func (o *MonthlyUsageAttributionValues) SetFunctionsUsage(v float64) {
	o.FunctionsUsage.Set(&v)
}

// SetFunctionsUsageNil sets the value for FunctionsUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetFunctionsUsageNil() {
	o.FunctionsUsage.Set(nil)
}

// UnsetFunctionsUsage ensures that no value is present for FunctionsUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetFunctionsUsage() {
	o.FunctionsUsage.Unset()
}

// GetInfraHostPercentage returns the InfraHostPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetInfraHostPercentage() float64 {
	if o == nil || o.InfraHostPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.InfraHostPercentage.Get()
}

// GetInfraHostPercentageOk returns a tuple with the InfraHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetInfraHostPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfraHostPercentage.Get(), o.InfraHostPercentage.IsSet()
}

// HasInfraHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInfraHostPercentage() bool {
	return o != nil && o.InfraHostPercentage.IsSet()
}

// SetInfraHostPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the InfraHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetInfraHostPercentage(v float64) {
	o.InfraHostPercentage.Set(&v)
}

// SetInfraHostPercentageNil sets the value for InfraHostPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetInfraHostPercentageNil() {
	o.InfraHostPercentage.Set(nil)
}

// UnsetInfraHostPercentage ensures that no value is present for InfraHostPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetInfraHostPercentage() {
	o.InfraHostPercentage.Unset()
}

// GetInfraHostUsage returns the InfraHostUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetInfraHostUsage() float64 {
	if o == nil || o.InfraHostUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.InfraHostUsage.Get()
}

// GetInfraHostUsageOk returns a tuple with the InfraHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetInfraHostUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfraHostUsage.Get(), o.InfraHostUsage.IsSet()
}

// HasInfraHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInfraHostUsage() bool {
	return o != nil && o.InfraHostUsage.IsSet()
}

// SetInfraHostUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the InfraHostUsage field.
func (o *MonthlyUsageAttributionValues) SetInfraHostUsage(v float64) {
	o.InfraHostUsage.Set(&v)
}

// SetInfraHostUsageNil sets the value for InfraHostUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetInfraHostUsageNil() {
	o.InfraHostUsage.Set(nil)
}

// UnsetInfraHostUsage ensures that no value is present for InfraHostUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetInfraHostUsage() {
	o.InfraHostUsage.Unset()
}

// GetInvocationsPercentage returns the InvocationsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetInvocationsPercentage() float64 {
	if o == nil || o.InvocationsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.InvocationsPercentage.Get()
}

// GetInvocationsPercentageOk returns a tuple with the InvocationsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetInvocationsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvocationsPercentage.Get(), o.InvocationsPercentage.IsSet()
}

// HasInvocationsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInvocationsPercentage() bool {
	return o != nil && o.InvocationsPercentage.IsSet()
}

// SetInvocationsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the InvocationsPercentage field.
func (o *MonthlyUsageAttributionValues) SetInvocationsPercentage(v float64) {
	o.InvocationsPercentage.Set(&v)
}

// SetInvocationsPercentageNil sets the value for InvocationsPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetInvocationsPercentageNil() {
	o.InvocationsPercentage.Set(nil)
}

// UnsetInvocationsPercentage ensures that no value is present for InvocationsPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetInvocationsPercentage() {
	o.InvocationsPercentage.Unset()
}

// GetInvocationsUsage returns the InvocationsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetInvocationsUsage() float64 {
	if o == nil || o.InvocationsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.InvocationsUsage.Get()
}

// GetInvocationsUsageOk returns a tuple with the InvocationsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetInvocationsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvocationsUsage.Get(), o.InvocationsUsage.IsSet()
}

// HasInvocationsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInvocationsUsage() bool {
	return o != nil && o.InvocationsUsage.IsSet()
}

// SetInvocationsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the InvocationsUsage field.
func (o *MonthlyUsageAttributionValues) SetInvocationsUsage(v float64) {
	o.InvocationsUsage.Set(&v)
}

// SetInvocationsUsageNil sets the value for InvocationsUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetInvocationsUsageNil() {
	o.InvocationsUsage.Set(nil)
}

// UnsetInvocationsUsage ensures that no value is present for InvocationsUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetInvocationsUsage() {
	o.InvocationsUsage.Unset()
}

// GetNpmHostPercentage returns the NpmHostPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetNpmHostPercentage() float64 {
	if o == nil || o.NpmHostPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NpmHostPercentage.Get()
}

// GetNpmHostPercentageOk returns a tuple with the NpmHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetNpmHostPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NpmHostPercentage.Get(), o.NpmHostPercentage.IsSet()
}

// HasNpmHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNpmHostPercentage() bool {
	return o != nil && o.NpmHostPercentage.IsSet()
}

// SetNpmHostPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the NpmHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetNpmHostPercentage(v float64) {
	o.NpmHostPercentage.Set(&v)
}

// SetNpmHostPercentageNil sets the value for NpmHostPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetNpmHostPercentageNil() {
	o.NpmHostPercentage.Set(nil)
}

// UnsetNpmHostPercentage ensures that no value is present for NpmHostPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetNpmHostPercentage() {
	o.NpmHostPercentage.Unset()
}

// GetNpmHostUsage returns the NpmHostUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetNpmHostUsage() float64 {
	if o == nil || o.NpmHostUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NpmHostUsage.Get()
}

// GetNpmHostUsageOk returns a tuple with the NpmHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetNpmHostUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NpmHostUsage.Get(), o.NpmHostUsage.IsSet()
}

// HasNpmHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNpmHostUsage() bool {
	return o != nil && o.NpmHostUsage.IsSet()
}

// SetNpmHostUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the NpmHostUsage field.
func (o *MonthlyUsageAttributionValues) SetNpmHostUsage(v float64) {
	o.NpmHostUsage.Set(&v)
}

// SetNpmHostUsageNil sets the value for NpmHostUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetNpmHostUsageNil() {
	o.NpmHostUsage.Set(nil)
}

// UnsetNpmHostUsage ensures that no value is present for NpmHostUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetNpmHostUsage() {
	o.NpmHostUsage.Unset()
}

// GetProfiledContainerPercentage returns the ProfiledContainerPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetProfiledContainerPercentage() float64 {
	if o == nil || o.ProfiledContainerPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledContainerPercentage.Get()
}

// GetProfiledContainerPercentageOk returns a tuple with the ProfiledContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledContainerPercentage.Get(), o.ProfiledContainerPercentage.IsSet()
}

// HasProfiledContainerPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledContainerPercentage() bool {
	return o != nil && o.ProfiledContainerPercentage.IsSet()
}

// SetProfiledContainerPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledContainerPercentage field.
func (o *MonthlyUsageAttributionValues) SetProfiledContainerPercentage(v float64) {
	o.ProfiledContainerPercentage.Set(&v)
}

// SetProfiledContainerPercentageNil sets the value for ProfiledContainerPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetProfiledContainerPercentageNil() {
	o.ProfiledContainerPercentage.Set(nil)
}

// UnsetProfiledContainerPercentage ensures that no value is present for ProfiledContainerPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetProfiledContainerPercentage() {
	o.ProfiledContainerPercentage.Unset()
}

// GetProfiledContainerUsage returns the ProfiledContainerUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetProfiledContainerUsage() float64 {
	if o == nil || o.ProfiledContainerUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledContainerUsage.Get()
}

// GetProfiledContainerUsageOk returns a tuple with the ProfiledContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledContainerUsage.Get(), o.ProfiledContainerUsage.IsSet()
}

// HasProfiledContainerUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledContainerUsage() bool {
	return o != nil && o.ProfiledContainerUsage.IsSet()
}

// SetProfiledContainerUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledContainerUsage field.
func (o *MonthlyUsageAttributionValues) SetProfiledContainerUsage(v float64) {
	o.ProfiledContainerUsage.Set(&v)
}

// SetProfiledContainerUsageNil sets the value for ProfiledContainerUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetProfiledContainerUsageNil() {
	o.ProfiledContainerUsage.Set(nil)
}

// UnsetProfiledContainerUsage ensures that no value is present for ProfiledContainerUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetProfiledContainerUsage() {
	o.ProfiledContainerUsage.Unset()
}

// GetProfiledFargatePercentage returns the ProfiledFargatePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetProfiledFargatePercentage() float64 {
	if o == nil || o.ProfiledFargatePercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledFargatePercentage.Get()
}

// GetProfiledFargatePercentageOk returns a tuple with the ProfiledFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetProfiledFargatePercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledFargatePercentage.Get(), o.ProfiledFargatePercentage.IsSet()
}

// HasProfiledFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledFargatePercentage() bool {
	return o != nil && o.ProfiledFargatePercentage.IsSet()
}

// SetProfiledFargatePercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetProfiledFargatePercentage(v float64) {
	o.ProfiledFargatePercentage.Set(&v)
}

// SetProfiledFargatePercentageNil sets the value for ProfiledFargatePercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetProfiledFargatePercentageNil() {
	o.ProfiledFargatePercentage.Set(nil)
}

// UnsetProfiledFargatePercentage ensures that no value is present for ProfiledFargatePercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetProfiledFargatePercentage() {
	o.ProfiledFargatePercentage.Unset()
}

// GetProfiledFargateUsage returns the ProfiledFargateUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetProfiledFargateUsage() float64 {
	if o == nil || o.ProfiledFargateUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledFargateUsage.Get()
}

// GetProfiledFargateUsageOk returns a tuple with the ProfiledFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetProfiledFargateUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledFargateUsage.Get(), o.ProfiledFargateUsage.IsSet()
}

// HasProfiledFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledFargateUsage() bool {
	return o != nil && o.ProfiledFargateUsage.IsSet()
}

// SetProfiledFargateUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetProfiledFargateUsage(v float64) {
	o.ProfiledFargateUsage.Set(&v)
}

// SetProfiledFargateUsageNil sets the value for ProfiledFargateUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetProfiledFargateUsageNil() {
	o.ProfiledFargateUsage.Set(nil)
}

// UnsetProfiledFargateUsage ensures that no value is present for ProfiledFargateUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetProfiledFargateUsage() {
	o.ProfiledFargateUsage.Unset()
}

// GetProfiledHostPercentage returns the ProfiledHostPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetProfiledHostPercentage() float64 {
	if o == nil || o.ProfiledHostPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledHostPercentage.Get()
}

// GetProfiledHostPercentageOk returns a tuple with the ProfiledHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetProfiledHostPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledHostPercentage.Get(), o.ProfiledHostPercentage.IsSet()
}

// HasProfiledHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledHostPercentage() bool {
	return o != nil && o.ProfiledHostPercentage.IsSet()
}

// SetProfiledHostPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetProfiledHostPercentage(v float64) {
	o.ProfiledHostPercentage.Set(&v)
}

// SetProfiledHostPercentageNil sets the value for ProfiledHostPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetProfiledHostPercentageNil() {
	o.ProfiledHostPercentage.Set(nil)
}

// UnsetProfiledHostPercentage ensures that no value is present for ProfiledHostPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetProfiledHostPercentage() {
	o.ProfiledHostPercentage.Unset()
}

// GetProfiledHostUsage returns the ProfiledHostUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetProfiledHostUsage() float64 {
	if o == nil || o.ProfiledHostUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledHostUsage.Get()
}

// GetProfiledHostUsageOk returns a tuple with the ProfiledHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetProfiledHostUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledHostUsage.Get(), o.ProfiledHostUsage.IsSet()
}

// HasProfiledHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledHostUsage() bool {
	return o != nil && o.ProfiledHostUsage.IsSet()
}

// SetProfiledHostUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledHostUsage field.
func (o *MonthlyUsageAttributionValues) SetProfiledHostUsage(v float64) {
	o.ProfiledHostUsage.Set(&v)
}

// SetProfiledHostUsageNil sets the value for ProfiledHostUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetProfiledHostUsageNil() {
	o.ProfiledHostUsage.Set(nil)
}

// UnsetProfiledHostUsage ensures that no value is present for ProfiledHostUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetProfiledHostUsage() {
	o.ProfiledHostUsage.Unset()
}

// GetSnmpPercentage returns the SnmpPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetSnmpPercentage() float64 {
	if o == nil || o.SnmpPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SnmpPercentage.Get()
}

// GetSnmpPercentageOk returns a tuple with the SnmpPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetSnmpPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnmpPercentage.Get(), o.SnmpPercentage.IsSet()
}

// HasSnmpPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSnmpPercentage() bool {
	return o != nil && o.SnmpPercentage.IsSet()
}

// SetSnmpPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the SnmpPercentage field.
func (o *MonthlyUsageAttributionValues) SetSnmpPercentage(v float64) {
	o.SnmpPercentage.Set(&v)
}

// SetSnmpPercentageNil sets the value for SnmpPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetSnmpPercentageNil() {
	o.SnmpPercentage.Set(nil)
}

// UnsetSnmpPercentage ensures that no value is present for SnmpPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetSnmpPercentage() {
	o.SnmpPercentage.Unset()
}

// GetSnmpUsage returns the SnmpUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetSnmpUsage() float64 {
	if o == nil || o.SnmpUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SnmpUsage.Get()
}

// GetSnmpUsageOk returns a tuple with the SnmpUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetSnmpUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnmpUsage.Get(), o.SnmpUsage.IsSet()
}

// HasSnmpUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSnmpUsage() bool {
	return o != nil && o.SnmpUsage.IsSet()
}

// SetSnmpUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the SnmpUsage field.
func (o *MonthlyUsageAttributionValues) SetSnmpUsage(v float64) {
	o.SnmpUsage.Set(&v)
}

// SetSnmpUsageNil sets the value for SnmpUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetSnmpUsageNil() {
	o.SnmpUsage.Set(nil)
}

// UnsetSnmpUsage ensures that no value is present for SnmpUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetSnmpUsage() {
	o.SnmpUsage.Unset()
}

// GetUniversalServiceMonitoringPercentage returns the UniversalServiceMonitoringPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringPercentage() float64 {
	if o == nil || o.UniversalServiceMonitoringPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.UniversalServiceMonitoringPercentage.Get()
}

// GetUniversalServiceMonitoringPercentageOk returns a tuple with the UniversalServiceMonitoringPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringPercentage.Get(), o.UniversalServiceMonitoringPercentage.IsSet()
}

// HasUniversalServiceMonitoringPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasUniversalServiceMonitoringPercentage() bool {
	return o != nil && o.UniversalServiceMonitoringPercentage.IsSet()
}

// SetUniversalServiceMonitoringPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the UniversalServiceMonitoringPercentage field.
func (o *MonthlyUsageAttributionValues) SetUniversalServiceMonitoringPercentage(v float64) {
	o.UniversalServiceMonitoringPercentage.Set(&v)
}

// SetUniversalServiceMonitoringPercentageNil sets the value for UniversalServiceMonitoringPercentage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetUniversalServiceMonitoringPercentageNil() {
	o.UniversalServiceMonitoringPercentage.Set(nil)
}

// UnsetUniversalServiceMonitoringPercentage ensures that no value is present for UniversalServiceMonitoringPercentage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetUniversalServiceMonitoringPercentage() {
	o.UniversalServiceMonitoringPercentage.Unset()
}

// GetUniversalServiceMonitoringUsage returns the UniversalServiceMonitoringUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringUsage() float64 {
	if o == nil || o.UniversalServiceMonitoringUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.UniversalServiceMonitoringUsage.Get()
}

// GetUniversalServiceMonitoringUsageOk returns a tuple with the UniversalServiceMonitoringUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringUsage.Get(), o.UniversalServiceMonitoringUsage.IsSet()
}

// HasUniversalServiceMonitoringUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasUniversalServiceMonitoringUsage() bool {
	return o != nil && o.UniversalServiceMonitoringUsage.IsSet()
}

// SetUniversalServiceMonitoringUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the UniversalServiceMonitoringUsage field.
func (o *MonthlyUsageAttributionValues) SetUniversalServiceMonitoringUsage(v float64) {
	o.UniversalServiceMonitoringUsage.Set(&v)
}

// SetUniversalServiceMonitoringUsageNil sets the value for UniversalServiceMonitoringUsage to be an explicit nil.
func (o *MonthlyUsageAttributionValues) SetUniversalServiceMonitoringUsageNil() {
	o.UniversalServiceMonitoringUsage.Set(nil)
}

// UnsetUniversalServiceMonitoringUsage ensures that no value is present for UniversalServiceMonitoringUsage, not even an explicit nil.
func (o *MonthlyUsageAttributionValues) UnsetUniversalServiceMonitoringUsage() {
	o.UniversalServiceMonitoringUsage.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyUsageAttributionValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ApiPercentage.IsSet() {
		toSerialize["api_percentage"] = o.ApiPercentage.Get()
	}
	if o.ApiUsage.IsSet() {
		toSerialize["api_usage"] = o.ApiUsage.Get()
	}
	if o.ApmFargatePercentage.IsSet() {
		toSerialize["apm_fargate_percentage"] = o.ApmFargatePercentage.Get()
	}
	if o.ApmFargateUsage.IsSet() {
		toSerialize["apm_fargate_usage"] = o.ApmFargateUsage.Get()
	}
	if o.ApmHostPercentage.IsSet() {
		toSerialize["apm_host_percentage"] = o.ApmHostPercentage.Get()
	}
	if o.ApmHostUsage.IsSet() {
		toSerialize["apm_host_usage"] = o.ApmHostUsage.Get()
	}
	if o.AppsecFargatePercentage.IsSet() {
		toSerialize["appsec_fargate_percentage"] = o.AppsecFargatePercentage.Get()
	}
	if o.AppsecFargateUsage.IsSet() {
		toSerialize["appsec_fargate_usage"] = o.AppsecFargateUsage.Get()
	}
	if o.AppsecPercentage.IsSet() {
		toSerialize["appsec_percentage"] = o.AppsecPercentage.Get()
	}
	if o.AppsecUsage.IsSet() {
		toSerialize["appsec_usage"] = o.AppsecUsage.Get()
	}
	if o.BrowserPercentage.IsSet() {
		toSerialize["browser_percentage"] = o.BrowserPercentage.Get()
	}
	if o.BrowserUsage.IsSet() {
		toSerialize["browser_usage"] = o.BrowserUsage.Get()
	}
	if o.ContainerExclAgentPercentage.IsSet() {
		toSerialize["container_excl_agent_percentage"] = o.ContainerExclAgentPercentage.Get()
	}
	if o.ContainerExclAgentUsage.IsSet() {
		toSerialize["container_excl_agent_usage"] = o.ContainerExclAgentUsage.Get()
	}
	if o.ContainerPercentage.IsSet() {
		toSerialize["container_percentage"] = o.ContainerPercentage.Get()
	}
	if o.ContainerUsage.IsSet() {
		toSerialize["container_usage"] = o.ContainerUsage.Get()
	}
	if o.CspmContainersPercentage.IsSet() {
		toSerialize["cspm_containers_percentage"] = o.CspmContainersPercentage.Get()
	}
	if o.CspmContainersUsage.IsSet() {
		toSerialize["cspm_containers_usage"] = o.CspmContainersUsage.Get()
	}
	if o.CspmHostsPercentage.IsSet() {
		toSerialize["cspm_hosts_percentage"] = o.CspmHostsPercentage.Get()
	}
	if o.CspmHostsUsage.IsSet() {
		toSerialize["cspm_hosts_usage"] = o.CspmHostsUsage.Get()
	}
	if o.CustomIngestedTimeseriesPercentage.IsSet() {
		toSerialize["custom_ingested_timeseries_percentage"] = o.CustomIngestedTimeseriesPercentage.Get()
	}
	if o.CustomIngestedTimeseriesUsage.IsSet() {
		toSerialize["custom_ingested_timeseries_usage"] = o.CustomIngestedTimeseriesUsage.Get()
	}
	if o.CustomTimeseriesPercentage.IsSet() {
		toSerialize["custom_timeseries_percentage"] = o.CustomTimeseriesPercentage.Get()
	}
	if o.CustomTimeseriesUsage.IsSet() {
		toSerialize["custom_timeseries_usage"] = o.CustomTimeseriesUsage.Get()
	}
	if o.CwsContainersPercentage.IsSet() {
		toSerialize["cws_containers_percentage"] = o.CwsContainersPercentage.Get()
	}
	if o.CwsContainersUsage.IsSet() {
		toSerialize["cws_containers_usage"] = o.CwsContainersUsage.Get()
	}
	if o.CwsHostsPercentage.IsSet() {
		toSerialize["cws_hosts_percentage"] = o.CwsHostsPercentage.Get()
	}
	if o.CwsHostsUsage.IsSet() {
		toSerialize["cws_hosts_usage"] = o.CwsHostsUsage.Get()
	}
	if o.DbmHostsPercentage.IsSet() {
		toSerialize["dbm_hosts_percentage"] = o.DbmHostsPercentage.Get()
	}
	if o.DbmHostsUsage.IsSet() {
		toSerialize["dbm_hosts_usage"] = o.DbmHostsUsage.Get()
	}
	if o.DbmQueriesPercentage.IsSet() {
		toSerialize["dbm_queries_percentage"] = o.DbmQueriesPercentage.Get()
	}
	if o.DbmQueriesUsage.IsSet() {
		toSerialize["dbm_queries_usage"] = o.DbmQueriesUsage.Get()
	}
	if o.EstimatedIndexedLogsPercentage.IsSet() {
		toSerialize["estimated_indexed_logs_percentage"] = o.EstimatedIndexedLogsPercentage.Get()
	}
	if o.EstimatedIndexedLogsUsage.IsSet() {
		toSerialize["estimated_indexed_logs_usage"] = o.EstimatedIndexedLogsUsage.Get()
	}
	if o.EstimatedIndexedSpansPercentage.IsSet() {
		toSerialize["estimated_indexed_spans_percentage"] = o.EstimatedIndexedSpansPercentage.Get()
	}
	if o.EstimatedIndexedSpansUsage.IsSet() {
		toSerialize["estimated_indexed_spans_usage"] = o.EstimatedIndexedSpansUsage.Get()
	}
	if o.EstimatedIngestedLogsPercentage.IsSet() {
		toSerialize["estimated_ingested_logs_percentage"] = o.EstimatedIngestedLogsPercentage.Get()
	}
	if o.EstimatedIngestedLogsUsage.IsSet() {
		toSerialize["estimated_ingested_logs_usage"] = o.EstimatedIngestedLogsUsage.Get()
	}
	if o.EstimatedIngestedSpansPercentage.IsSet() {
		toSerialize["estimated_ingested_spans_percentage"] = o.EstimatedIngestedSpansPercentage.Get()
	}
	if o.EstimatedIngestedSpansUsage.IsSet() {
		toSerialize["estimated_ingested_spans_usage"] = o.EstimatedIngestedSpansUsage.Get()
	}
	if o.EstimatedRumSessionsPercentage.IsSet() {
		toSerialize["estimated_rum_sessions_percentage"] = o.EstimatedRumSessionsPercentage.Get()
	}
	if o.EstimatedRumSessionsUsage.IsSet() {
		toSerialize["estimated_rum_sessions_usage"] = o.EstimatedRumSessionsUsage.Get()
	}
	if o.FargatePercentage.IsSet() {
		toSerialize["fargate_percentage"] = o.FargatePercentage.Get()
	}
	if o.FargateUsage.IsSet() {
		toSerialize["fargate_usage"] = o.FargateUsage.Get()
	}
	if o.FunctionsPercentage.IsSet() {
		toSerialize["functions_percentage"] = o.FunctionsPercentage.Get()
	}
	if o.FunctionsUsage.IsSet() {
		toSerialize["functions_usage"] = o.FunctionsUsage.Get()
	}
	if o.InfraHostPercentage.IsSet() {
		toSerialize["infra_host_percentage"] = o.InfraHostPercentage.Get()
	}
	if o.InfraHostUsage.IsSet() {
		toSerialize["infra_host_usage"] = o.InfraHostUsage.Get()
	}
	if o.InvocationsPercentage.IsSet() {
		toSerialize["invocations_percentage"] = o.InvocationsPercentage.Get()
	}
	if o.InvocationsUsage.IsSet() {
		toSerialize["invocations_usage"] = o.InvocationsUsage.Get()
	}
	if o.NpmHostPercentage.IsSet() {
		toSerialize["npm_host_percentage"] = o.NpmHostPercentage.Get()
	}
	if o.NpmHostUsage.IsSet() {
		toSerialize["npm_host_usage"] = o.NpmHostUsage.Get()
	}
	if o.ProfiledContainerPercentage.IsSet() {
		toSerialize["profiled_container_percentage"] = o.ProfiledContainerPercentage.Get()
	}
	if o.ProfiledContainerUsage.IsSet() {
		toSerialize["profiled_container_usage"] = o.ProfiledContainerUsage.Get()
	}
	if o.ProfiledFargatePercentage.IsSet() {
		toSerialize["profiled_fargate_percentage"] = o.ProfiledFargatePercentage.Get()
	}
	if o.ProfiledFargateUsage.IsSet() {
		toSerialize["profiled_fargate_usage"] = o.ProfiledFargateUsage.Get()
	}
	if o.ProfiledHostPercentage.IsSet() {
		toSerialize["profiled_host_percentage"] = o.ProfiledHostPercentage.Get()
	}
	if o.ProfiledHostUsage.IsSet() {
		toSerialize["profiled_host_usage"] = o.ProfiledHostUsage.Get()
	}
	if o.SnmpPercentage.IsSet() {
		toSerialize["snmp_percentage"] = o.SnmpPercentage.Get()
	}
	if o.SnmpUsage.IsSet() {
		toSerialize["snmp_usage"] = o.SnmpUsage.Get()
	}
	if o.UniversalServiceMonitoringPercentage.IsSet() {
		toSerialize["universal_service_monitoring_percentage"] = o.UniversalServiceMonitoringPercentage.Get()
	}
	if o.UniversalServiceMonitoringUsage.IsSet() {
		toSerialize["universal_service_monitoring_usage"] = o.UniversalServiceMonitoringUsage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonthlyUsageAttributionValues) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		ApiPercentage                        datadog.NullableFloat64 `json:"api_percentage,omitempty"`
		ApiUsage                             datadog.NullableFloat64 `json:"api_usage,omitempty"`
		ApmFargatePercentage                 datadog.NullableFloat64 `json:"apm_fargate_percentage,omitempty"`
		ApmFargateUsage                      datadog.NullableFloat64 `json:"apm_fargate_usage,omitempty"`
		ApmHostPercentage                    datadog.NullableFloat64 `json:"apm_host_percentage,omitempty"`
		ApmHostUsage                         datadog.NullableFloat64 `json:"apm_host_usage,omitempty"`
		AppsecFargatePercentage              datadog.NullableFloat64 `json:"appsec_fargate_percentage,omitempty"`
		AppsecFargateUsage                   datadog.NullableFloat64 `json:"appsec_fargate_usage,omitempty"`
		AppsecPercentage                     datadog.NullableFloat64 `json:"appsec_percentage,omitempty"`
		AppsecUsage                          datadog.NullableFloat64 `json:"appsec_usage,omitempty"`
		BrowserPercentage                    datadog.NullableFloat64 `json:"browser_percentage,omitempty"`
		BrowserUsage                         datadog.NullableFloat64 `json:"browser_usage,omitempty"`
		ContainerExclAgentPercentage         datadog.NullableFloat64 `json:"container_excl_agent_percentage,omitempty"`
		ContainerExclAgentUsage              datadog.NullableFloat64 `json:"container_excl_agent_usage,omitempty"`
		ContainerPercentage                  datadog.NullableFloat64 `json:"container_percentage,omitempty"`
		ContainerUsage                       datadog.NullableFloat64 `json:"container_usage,omitempty"`
		CspmContainersPercentage             datadog.NullableFloat64 `json:"cspm_containers_percentage,omitempty"`
		CspmContainersUsage                  datadog.NullableFloat64 `json:"cspm_containers_usage,omitempty"`
		CspmHostsPercentage                  datadog.NullableFloat64 `json:"cspm_hosts_percentage,omitempty"`
		CspmHostsUsage                       datadog.NullableFloat64 `json:"cspm_hosts_usage,omitempty"`
		CustomIngestedTimeseriesPercentage   datadog.NullableFloat64 `json:"custom_ingested_timeseries_percentage,omitempty"`
		CustomIngestedTimeseriesUsage        datadog.NullableFloat64 `json:"custom_ingested_timeseries_usage,omitempty"`
		CustomTimeseriesPercentage           datadog.NullableFloat64 `json:"custom_timeseries_percentage,omitempty"`
		CustomTimeseriesUsage                datadog.NullableFloat64 `json:"custom_timeseries_usage,omitempty"`
		CwsContainersPercentage              datadog.NullableFloat64 `json:"cws_containers_percentage,omitempty"`
		CwsContainersUsage                   datadog.NullableFloat64 `json:"cws_containers_usage,omitempty"`
		CwsHostsPercentage                   datadog.NullableFloat64 `json:"cws_hosts_percentage,omitempty"`
		CwsHostsUsage                        datadog.NullableFloat64 `json:"cws_hosts_usage,omitempty"`
		DbmHostsPercentage                   datadog.NullableFloat64 `json:"dbm_hosts_percentage,omitempty"`
		DbmHostsUsage                        datadog.NullableFloat64 `json:"dbm_hosts_usage,omitempty"`
		DbmQueriesPercentage                 datadog.NullableFloat64 `json:"dbm_queries_percentage,omitempty"`
		DbmQueriesUsage                      datadog.NullableFloat64 `json:"dbm_queries_usage,omitempty"`
		EstimatedIndexedLogsPercentage       datadog.NullableFloat64 `json:"estimated_indexed_logs_percentage,omitempty"`
		EstimatedIndexedLogsUsage            datadog.NullableFloat64 `json:"estimated_indexed_logs_usage,omitempty"`
		EstimatedIndexedSpansPercentage      datadog.NullableFloat64 `json:"estimated_indexed_spans_percentage,omitempty"`
		EstimatedIndexedSpansUsage           datadog.NullableFloat64 `json:"estimated_indexed_spans_usage,omitempty"`
		EstimatedIngestedLogsPercentage      datadog.NullableFloat64 `json:"estimated_ingested_logs_percentage,omitempty"`
		EstimatedIngestedLogsUsage           datadog.NullableFloat64 `json:"estimated_ingested_logs_usage,omitempty"`
		EstimatedIngestedSpansPercentage     datadog.NullableFloat64 `json:"estimated_ingested_spans_percentage,omitempty"`
		EstimatedIngestedSpansUsage          datadog.NullableFloat64 `json:"estimated_ingested_spans_usage,omitempty"`
		EstimatedRumSessionsPercentage       datadog.NullableFloat64 `json:"estimated_rum_sessions_percentage,omitempty"`
		EstimatedRumSessionsUsage            datadog.NullableFloat64 `json:"estimated_rum_sessions_usage,omitempty"`
		FargatePercentage                    datadog.NullableFloat64 `json:"fargate_percentage,omitempty"`
		FargateUsage                         datadog.NullableFloat64 `json:"fargate_usage,omitempty"`
		FunctionsPercentage                  datadog.NullableFloat64 `json:"functions_percentage,omitempty"`
		FunctionsUsage                       datadog.NullableFloat64 `json:"functions_usage,omitempty"`
		InfraHostPercentage                  datadog.NullableFloat64 `json:"infra_host_percentage,omitempty"`
		InfraHostUsage                       datadog.NullableFloat64 `json:"infra_host_usage,omitempty"`
		InvocationsPercentage                datadog.NullableFloat64 `json:"invocations_percentage,omitempty"`
		InvocationsUsage                     datadog.NullableFloat64 `json:"invocations_usage,omitempty"`
		NpmHostPercentage                    datadog.NullableFloat64 `json:"npm_host_percentage,omitempty"`
		NpmHostUsage                         datadog.NullableFloat64 `json:"npm_host_usage,omitempty"`
		ProfiledContainerPercentage          datadog.NullableFloat64 `json:"profiled_container_percentage,omitempty"`
		ProfiledContainerUsage               datadog.NullableFloat64 `json:"profiled_container_usage,omitempty"`
		ProfiledFargatePercentage            datadog.NullableFloat64 `json:"profiled_fargate_percentage,omitempty"`
		ProfiledFargateUsage                 datadog.NullableFloat64 `json:"profiled_fargate_usage,omitempty"`
		ProfiledHostPercentage               datadog.NullableFloat64 `json:"profiled_host_percentage,omitempty"`
		ProfiledHostUsage                    datadog.NullableFloat64 `json:"profiled_host_usage,omitempty"`
		SnmpPercentage                       datadog.NullableFloat64 `json:"snmp_percentage,omitempty"`
		SnmpUsage                            datadog.NullableFloat64 `json:"snmp_usage,omitempty"`
		UniversalServiceMonitoringPercentage datadog.NullableFloat64 `json:"universal_service_monitoring_percentage,omitempty"`
		UniversalServiceMonitoringUsage      datadog.NullableFloat64 `json:"universal_service_monitoring_usage,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"api_percentage", "api_usage", "apm_fargate_percentage", "apm_fargate_usage", "apm_host_percentage", "apm_host_usage", "appsec_fargate_percentage", "appsec_fargate_usage", "appsec_percentage", "appsec_usage", "browser_percentage", "browser_usage", "container_excl_agent_percentage", "container_excl_agent_usage", "container_percentage", "container_usage", "cspm_containers_percentage", "cspm_containers_usage", "cspm_hosts_percentage", "cspm_hosts_usage", "custom_ingested_timeseries_percentage", "custom_ingested_timeseries_usage", "custom_timeseries_percentage", "custom_timeseries_usage", "cws_containers_percentage", "cws_containers_usage", "cws_hosts_percentage", "cws_hosts_usage", "dbm_hosts_percentage", "dbm_hosts_usage", "dbm_queries_percentage", "dbm_queries_usage", "estimated_indexed_logs_percentage", "estimated_indexed_logs_usage", "estimated_indexed_spans_percentage", "estimated_indexed_spans_usage", "estimated_ingested_logs_percentage", "estimated_ingested_logs_usage", "estimated_ingested_spans_percentage", "estimated_ingested_spans_usage", "estimated_rum_sessions_percentage", "estimated_rum_sessions_usage", "fargate_percentage", "fargate_usage", "functions_percentage", "functions_usage", "infra_host_percentage", "infra_host_usage", "invocations_percentage", "invocations_usage", "npm_host_percentage", "npm_host_usage", "profiled_container_percentage", "profiled_container_usage", "profiled_fargate_percentage", "profiled_fargate_usage", "profiled_host_percentage", "profiled_host_usage", "snmp_percentage", "snmp_usage", "universal_service_monitoring_percentage", "universal_service_monitoring_usage"})
	} else {
		return err
	}
	o.ApiPercentage = all.ApiPercentage
	o.ApiUsage = all.ApiUsage
	o.ApmFargatePercentage = all.ApmFargatePercentage
	o.ApmFargateUsage = all.ApmFargateUsage
	o.ApmHostPercentage = all.ApmHostPercentage
	o.ApmHostUsage = all.ApmHostUsage
	o.AppsecFargatePercentage = all.AppsecFargatePercentage
	o.AppsecFargateUsage = all.AppsecFargateUsage
	o.AppsecPercentage = all.AppsecPercentage
	o.AppsecUsage = all.AppsecUsage
	o.BrowserPercentage = all.BrowserPercentage
	o.BrowserUsage = all.BrowserUsage
	o.ContainerExclAgentPercentage = all.ContainerExclAgentPercentage
	o.ContainerExclAgentUsage = all.ContainerExclAgentUsage
	o.ContainerPercentage = all.ContainerPercentage
	o.ContainerUsage = all.ContainerUsage
	o.CspmContainersPercentage = all.CspmContainersPercentage
	o.CspmContainersUsage = all.CspmContainersUsage
	o.CspmHostsPercentage = all.CspmHostsPercentage
	o.CspmHostsUsage = all.CspmHostsUsage
	o.CustomIngestedTimeseriesPercentage = all.CustomIngestedTimeseriesPercentage
	o.CustomIngestedTimeseriesUsage = all.CustomIngestedTimeseriesUsage
	o.CustomTimeseriesPercentage = all.CustomTimeseriesPercentage
	o.CustomTimeseriesUsage = all.CustomTimeseriesUsage
	o.CwsContainersPercentage = all.CwsContainersPercentage
	o.CwsContainersUsage = all.CwsContainersUsage
	o.CwsHostsPercentage = all.CwsHostsPercentage
	o.CwsHostsUsage = all.CwsHostsUsage
	o.DbmHostsPercentage = all.DbmHostsPercentage
	o.DbmHostsUsage = all.DbmHostsUsage
	o.DbmQueriesPercentage = all.DbmQueriesPercentage
	o.DbmQueriesUsage = all.DbmQueriesUsage
	o.EstimatedIndexedLogsPercentage = all.EstimatedIndexedLogsPercentage
	o.EstimatedIndexedLogsUsage = all.EstimatedIndexedLogsUsage
	o.EstimatedIndexedSpansPercentage = all.EstimatedIndexedSpansPercentage
	o.EstimatedIndexedSpansUsage = all.EstimatedIndexedSpansUsage
	o.EstimatedIngestedLogsPercentage = all.EstimatedIngestedLogsPercentage
	o.EstimatedIngestedLogsUsage = all.EstimatedIngestedLogsUsage
	o.EstimatedIngestedSpansPercentage = all.EstimatedIngestedSpansPercentage
	o.EstimatedIngestedSpansUsage = all.EstimatedIngestedSpansUsage
	o.EstimatedRumSessionsPercentage = all.EstimatedRumSessionsPercentage
	o.EstimatedRumSessionsUsage = all.EstimatedRumSessionsUsage
	o.FargatePercentage = all.FargatePercentage
	o.FargateUsage = all.FargateUsage
	o.FunctionsPercentage = all.FunctionsPercentage
	o.FunctionsUsage = all.FunctionsUsage
	o.InfraHostPercentage = all.InfraHostPercentage
	o.InfraHostUsage = all.InfraHostUsage
	o.InvocationsPercentage = all.InvocationsPercentage
	o.InvocationsUsage = all.InvocationsUsage
	o.NpmHostPercentage = all.NpmHostPercentage
	o.NpmHostUsage = all.NpmHostUsage
	o.ProfiledContainerPercentage = all.ProfiledContainerPercentage
	o.ProfiledContainerUsage = all.ProfiledContainerUsage
	o.ProfiledFargatePercentage = all.ProfiledFargatePercentage
	o.ProfiledFargateUsage = all.ProfiledFargateUsage
	o.ProfiledHostPercentage = all.ProfiledHostPercentage
	o.ProfiledHostUsage = all.ProfiledHostUsage
	o.SnmpPercentage = all.SnmpPercentage
	o.SnmpUsage = all.SnmpUsage
	o.UniversalServiceMonitoringPercentage = all.UniversalServiceMonitoringPercentage
	o.UniversalServiceMonitoringUsage = all.UniversalServiceMonitoringUsage
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
