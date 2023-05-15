// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageAttributionValues Fields in Usage Summary by tag(s).
type UsageAttributionValues struct {
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
	// The percentage of container usage by tag(s).
	ContainerPercentage datadog.NullableFloat64 `json:"container_percentage,omitempty"`
	// The container usage by tag(s).
	ContainerUsage datadog.NullableFloat64 `json:"container_usage,omitempty"`
	// The percentage of Cloud Security Posture Management container usage by tag(s)
	CspmContainerPercentage datadog.NullableFloat64 `json:"cspm_container_percentage,omitempty"`
	// The Cloud Security Posture Management container usage by tag(s)
	CspmContainerUsage datadog.NullableFloat64 `json:"cspm_container_usage,omitempty"`
	// The percentage of Cloud Security Posture Management host usage by tag(s)
	CspmHostPercentage datadog.NullableFloat64 `json:"cspm_host_percentage,omitempty"`
	// The Cloud Security Posture Management host usage by tag(s)
	CspmHostUsage datadog.NullableFloat64 `json:"cspm_host_usage,omitempty"`
	// The percentage of custom metrics usage by tag(s).
	CustomTimeseriesPercentage datadog.NullableFloat64 `json:"custom_timeseries_percentage,omitempty"`
	// The custom metrics usage by tag(s).
	CustomTimeseriesUsage datadog.NullableFloat64 `json:"custom_timeseries_usage,omitempty"`
	// The percentage of Cloud Workload Security container usage by tag(s)
	CwsContainerPercentage datadog.NullableFloat64 `json:"cws_container_percentage,omitempty"`
	// The Cloud Workload Security container usage by tag(s)
	CwsContainerUsage datadog.NullableFloat64 `json:"cws_container_usage,omitempty"`
	// The percentage of Cloud Workload Security host usage by tag(s)
	CwsHostPercentage datadog.NullableFloat64 `json:"cws_host_percentage,omitempty"`
	// The Cloud Workload Security host usage by tag(s)
	CwsHostUsage datadog.NullableFloat64 `json:"cws_host_usage,omitempty"`
	// The percentage of Database Monitoring host usage by tag(s).
	DbmHostsPercentage datadog.NullableFloat64 `json:"dbm_hosts_percentage,omitempty"`
	// The Database Monitoring host usage by tag(s).
	DbmHostsUsage datadog.NullableFloat64 `json:"dbm_hosts_usage,omitempty"`
	// The percentage of Database Monitoring normalized queries usage by tag(s).
	DbmQueriesPercentage datadog.NullableFloat64 `json:"dbm_queries_percentage,omitempty"`
	// The Database Monitoring normalized queries usage by tag(s).
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
	// The percentage of infrastructure host usage by tag(s).
	InfraHostPercentage datadog.NullableFloat64 `json:"infra_host_percentage,omitempty"`
	// The infrastructure host usage by tag(s).
	InfraHostUsage datadog.NullableFloat64 `json:"infra_host_usage,omitempty"`
	// The percentage of Lambda function usage by tag(s).
	LambdaFunctionsPercentage datadog.NullableFloat64 `json:"lambda_functions_percentage,omitempty"`
	// The Lambda function usage by tag(s).
	LambdaFunctionsUsage datadog.NullableFloat64 `json:"lambda_functions_usage,omitempty"`
	// The percentage of Lambda invocation usage by tag(s).
	LambdaInvocationsPercentage datadog.NullableFloat64 `json:"lambda_invocations_percentage,omitempty"`
	// The Lambda invocation usage by tag(s).
	LambdaInvocationsUsage datadog.NullableFloat64 `json:"lambda_invocations_usage,omitempty"`
	// The percentage of network host usage by tag(s).
	NpmHostPercentage datadog.NullableFloat64 `json:"npm_host_percentage,omitempty"`
	// The network host usage by tag(s).
	NpmHostUsage datadog.NullableFloat64 `json:"npm_host_usage,omitempty"`
	// The percentage of profiled containers usage by tag(s).
	ProfiledContainerPercentage datadog.NullableFloat64 `json:"profiled_container_percentage,omitempty"`
	// The profiled container usage by tag(s).
	ProfiledContainerUsage datadog.NullableFloat64 `json:"profiled_container_usage,omitempty"`
	// The percentage of profiled hosts usage by tag(s).
	ProfiledHostsPercentage datadog.NullableFloat64 `json:"profiled_hosts_percentage,omitempty"`
	// The profiled host usage by tag(s).
	ProfiledHostsUsage datadog.NullableFloat64 `json:"profiled_hosts_usage,omitempty"`
	// The percentage of network device usage by tag(s).
	SnmpPercentage datadog.NullableFloat64 `json:"snmp_percentage,omitempty"`
	// The network device usage by tag(s).
	SnmpUsage datadog.NullableFloat64 `json:"snmp_usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageAttributionValues instantiates a new UsageAttributionValues object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageAttributionValues() *UsageAttributionValues {
	this := UsageAttributionValues{}
	return &this
}

// NewUsageAttributionValuesWithDefaults instantiates a new UsageAttributionValues object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageAttributionValuesWithDefaults() *UsageAttributionValues {
	this := UsageAttributionValues{}
	return &this
}

// GetApiPercentage returns the ApiPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetApiPercentage() float64 {
	if o == nil || o.ApiPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApiPercentage.Get()
}

// GetApiPercentageOk returns a tuple with the ApiPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetApiPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiPercentage.Get(), o.ApiPercentage.IsSet()
}

// HasApiPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasApiPercentage() bool {
	return o != nil && o.ApiPercentage.IsSet()
}

// SetApiPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApiPercentage field.
func (o *UsageAttributionValues) SetApiPercentage(v float64) {
	o.ApiPercentage.Set(&v)
}

// SetApiPercentageNil sets the value for ApiPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetApiPercentageNil() {
	o.ApiPercentage.Set(nil)
}

// UnsetApiPercentage ensures that no value is present for ApiPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetApiPercentage() {
	o.ApiPercentage.Unset()
}

// GetApiUsage returns the ApiUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetApiUsage() float64 {
	if o == nil || o.ApiUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApiUsage.Get()
}

// GetApiUsageOk returns a tuple with the ApiUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetApiUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiUsage.Get(), o.ApiUsage.IsSet()
}

// HasApiUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasApiUsage() bool {
	return o != nil && o.ApiUsage.IsSet()
}

// SetApiUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApiUsage field.
func (o *UsageAttributionValues) SetApiUsage(v float64) {
	o.ApiUsage.Set(&v)
}

// SetApiUsageNil sets the value for ApiUsage to be an explicit nil.
func (o *UsageAttributionValues) SetApiUsageNil() {
	o.ApiUsage.Set(nil)
}

// UnsetApiUsage ensures that no value is present for ApiUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetApiUsage() {
	o.ApiUsage.Unset()
}

// GetApmFargatePercentage returns the ApmFargatePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetApmFargatePercentage() float64 {
	if o == nil || o.ApmFargatePercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApmFargatePercentage.Get()
}

// GetApmFargatePercentageOk returns a tuple with the ApmFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetApmFargatePercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmFargatePercentage.Get(), o.ApmFargatePercentage.IsSet()
}

// HasApmFargatePercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasApmFargatePercentage() bool {
	return o != nil && o.ApmFargatePercentage.IsSet()
}

// SetApmFargatePercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApmFargatePercentage field.
func (o *UsageAttributionValues) SetApmFargatePercentage(v float64) {
	o.ApmFargatePercentage.Set(&v)
}

// SetApmFargatePercentageNil sets the value for ApmFargatePercentage to be an explicit nil.
func (o *UsageAttributionValues) SetApmFargatePercentageNil() {
	o.ApmFargatePercentage.Set(nil)
}

// UnsetApmFargatePercentage ensures that no value is present for ApmFargatePercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetApmFargatePercentage() {
	o.ApmFargatePercentage.Unset()
}

// GetApmFargateUsage returns the ApmFargateUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetApmFargateUsage() float64 {
	if o == nil || o.ApmFargateUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApmFargateUsage.Get()
}

// GetApmFargateUsageOk returns a tuple with the ApmFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetApmFargateUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmFargateUsage.Get(), o.ApmFargateUsage.IsSet()
}

// HasApmFargateUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasApmFargateUsage() bool {
	return o != nil && o.ApmFargateUsage.IsSet()
}

// SetApmFargateUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApmFargateUsage field.
func (o *UsageAttributionValues) SetApmFargateUsage(v float64) {
	o.ApmFargateUsage.Set(&v)
}

// SetApmFargateUsageNil sets the value for ApmFargateUsage to be an explicit nil.
func (o *UsageAttributionValues) SetApmFargateUsageNil() {
	o.ApmFargateUsage.Set(nil)
}

// UnsetApmFargateUsage ensures that no value is present for ApmFargateUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetApmFargateUsage() {
	o.ApmFargateUsage.Unset()
}

// GetApmHostPercentage returns the ApmHostPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetApmHostPercentage() float64 {
	if o == nil || o.ApmHostPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApmHostPercentage.Get()
}

// GetApmHostPercentageOk returns a tuple with the ApmHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetApmHostPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmHostPercentage.Get(), o.ApmHostPercentage.IsSet()
}

// HasApmHostPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasApmHostPercentage() bool {
	return o != nil && o.ApmHostPercentage.IsSet()
}

// SetApmHostPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApmHostPercentage field.
func (o *UsageAttributionValues) SetApmHostPercentage(v float64) {
	o.ApmHostPercentage.Set(&v)
}

// SetApmHostPercentageNil sets the value for ApmHostPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetApmHostPercentageNil() {
	o.ApmHostPercentage.Set(nil)
}

// UnsetApmHostPercentage ensures that no value is present for ApmHostPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetApmHostPercentage() {
	o.ApmHostPercentage.Unset()
}

// GetApmHostUsage returns the ApmHostUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetApmHostUsage() float64 {
	if o == nil || o.ApmHostUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ApmHostUsage.Get()
}

// GetApmHostUsageOk returns a tuple with the ApmHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetApmHostUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmHostUsage.Get(), o.ApmHostUsage.IsSet()
}

// HasApmHostUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasApmHostUsage() bool {
	return o != nil && o.ApmHostUsage.IsSet()
}

// SetApmHostUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ApmHostUsage field.
func (o *UsageAttributionValues) SetApmHostUsage(v float64) {
	o.ApmHostUsage.Set(&v)
}

// SetApmHostUsageNil sets the value for ApmHostUsage to be an explicit nil.
func (o *UsageAttributionValues) SetApmHostUsageNil() {
	o.ApmHostUsage.Set(nil)
}

// UnsetApmHostUsage ensures that no value is present for ApmHostUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetApmHostUsage() {
	o.ApmHostUsage.Unset()
}

// GetAppsecFargatePercentage returns the AppsecFargatePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetAppsecFargatePercentage() float64 {
	if o == nil || o.AppsecFargatePercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.AppsecFargatePercentage.Get()
}

// GetAppsecFargatePercentageOk returns a tuple with the AppsecFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetAppsecFargatePercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecFargatePercentage.Get(), o.AppsecFargatePercentage.IsSet()
}

// HasAppsecFargatePercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasAppsecFargatePercentage() bool {
	return o != nil && o.AppsecFargatePercentage.IsSet()
}

// SetAppsecFargatePercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the AppsecFargatePercentage field.
func (o *UsageAttributionValues) SetAppsecFargatePercentage(v float64) {
	o.AppsecFargatePercentage.Set(&v)
}

// SetAppsecFargatePercentageNil sets the value for AppsecFargatePercentage to be an explicit nil.
func (o *UsageAttributionValues) SetAppsecFargatePercentageNil() {
	o.AppsecFargatePercentage.Set(nil)
}

// UnsetAppsecFargatePercentage ensures that no value is present for AppsecFargatePercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetAppsecFargatePercentage() {
	o.AppsecFargatePercentage.Unset()
}

// GetAppsecFargateUsage returns the AppsecFargateUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetAppsecFargateUsage() float64 {
	if o == nil || o.AppsecFargateUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.AppsecFargateUsage.Get()
}

// GetAppsecFargateUsageOk returns a tuple with the AppsecFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetAppsecFargateUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecFargateUsage.Get(), o.AppsecFargateUsage.IsSet()
}

// HasAppsecFargateUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasAppsecFargateUsage() bool {
	return o != nil && o.AppsecFargateUsage.IsSet()
}

// SetAppsecFargateUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the AppsecFargateUsage field.
func (o *UsageAttributionValues) SetAppsecFargateUsage(v float64) {
	o.AppsecFargateUsage.Set(&v)
}

// SetAppsecFargateUsageNil sets the value for AppsecFargateUsage to be an explicit nil.
func (o *UsageAttributionValues) SetAppsecFargateUsageNil() {
	o.AppsecFargateUsage.Set(nil)
}

// UnsetAppsecFargateUsage ensures that no value is present for AppsecFargateUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetAppsecFargateUsage() {
	o.AppsecFargateUsage.Unset()
}

// GetAppsecPercentage returns the AppsecPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetAppsecPercentage() float64 {
	if o == nil || o.AppsecPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.AppsecPercentage.Get()
}

// GetAppsecPercentageOk returns a tuple with the AppsecPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetAppsecPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecPercentage.Get(), o.AppsecPercentage.IsSet()
}

// HasAppsecPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasAppsecPercentage() bool {
	return o != nil && o.AppsecPercentage.IsSet()
}

// SetAppsecPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the AppsecPercentage field.
func (o *UsageAttributionValues) SetAppsecPercentage(v float64) {
	o.AppsecPercentage.Set(&v)
}

// SetAppsecPercentageNil sets the value for AppsecPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetAppsecPercentageNil() {
	o.AppsecPercentage.Set(nil)
}

// UnsetAppsecPercentage ensures that no value is present for AppsecPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetAppsecPercentage() {
	o.AppsecPercentage.Unset()
}

// GetAppsecUsage returns the AppsecUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetAppsecUsage() float64 {
	if o == nil || o.AppsecUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.AppsecUsage.Get()
}

// GetAppsecUsageOk returns a tuple with the AppsecUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetAppsecUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppsecUsage.Get(), o.AppsecUsage.IsSet()
}

// HasAppsecUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasAppsecUsage() bool {
	return o != nil && o.AppsecUsage.IsSet()
}

// SetAppsecUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the AppsecUsage field.
func (o *UsageAttributionValues) SetAppsecUsage(v float64) {
	o.AppsecUsage.Set(&v)
}

// SetAppsecUsageNil sets the value for AppsecUsage to be an explicit nil.
func (o *UsageAttributionValues) SetAppsecUsageNil() {
	o.AppsecUsage.Set(nil)
}

// UnsetAppsecUsage ensures that no value is present for AppsecUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetAppsecUsage() {
	o.AppsecUsage.Unset()
}

// GetBrowserPercentage returns the BrowserPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetBrowserPercentage() float64 {
	if o == nil || o.BrowserPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.BrowserPercentage.Get()
}

// GetBrowserPercentageOk returns a tuple with the BrowserPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetBrowserPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserPercentage.Get(), o.BrowserPercentage.IsSet()
}

// HasBrowserPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasBrowserPercentage() bool {
	return o != nil && o.BrowserPercentage.IsSet()
}

// SetBrowserPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the BrowserPercentage field.
func (o *UsageAttributionValues) SetBrowserPercentage(v float64) {
	o.BrowserPercentage.Set(&v)
}

// SetBrowserPercentageNil sets the value for BrowserPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetBrowserPercentageNil() {
	o.BrowserPercentage.Set(nil)
}

// UnsetBrowserPercentage ensures that no value is present for BrowserPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetBrowserPercentage() {
	o.BrowserPercentage.Unset()
}

// GetBrowserUsage returns the BrowserUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetBrowserUsage() float64 {
	if o == nil || o.BrowserUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.BrowserUsage.Get()
}

// GetBrowserUsageOk returns a tuple with the BrowserUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetBrowserUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserUsage.Get(), o.BrowserUsage.IsSet()
}

// HasBrowserUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasBrowserUsage() bool {
	return o != nil && o.BrowserUsage.IsSet()
}

// SetBrowserUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the BrowserUsage field.
func (o *UsageAttributionValues) SetBrowserUsage(v float64) {
	o.BrowserUsage.Set(&v)
}

// SetBrowserUsageNil sets the value for BrowserUsage to be an explicit nil.
func (o *UsageAttributionValues) SetBrowserUsageNil() {
	o.BrowserUsage.Set(nil)
}

// UnsetBrowserUsage ensures that no value is present for BrowserUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetBrowserUsage() {
	o.BrowserUsage.Unset()
}

// GetContainerPercentage returns the ContainerPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetContainerPercentage() float64 {
	if o == nil || o.ContainerPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ContainerPercentage.Get()
}

// GetContainerPercentageOk returns a tuple with the ContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetContainerPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerPercentage.Get(), o.ContainerPercentage.IsSet()
}

// HasContainerPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasContainerPercentage() bool {
	return o != nil && o.ContainerPercentage.IsSet()
}

// SetContainerPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ContainerPercentage field.
func (o *UsageAttributionValues) SetContainerPercentage(v float64) {
	o.ContainerPercentage.Set(&v)
}

// SetContainerPercentageNil sets the value for ContainerPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetContainerPercentageNil() {
	o.ContainerPercentage.Set(nil)
}

// UnsetContainerPercentage ensures that no value is present for ContainerPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetContainerPercentage() {
	o.ContainerPercentage.Unset()
}

// GetContainerUsage returns the ContainerUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetContainerUsage() float64 {
	if o == nil || o.ContainerUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ContainerUsage.Get()
}

// GetContainerUsageOk returns a tuple with the ContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetContainerUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerUsage.Get(), o.ContainerUsage.IsSet()
}

// HasContainerUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasContainerUsage() bool {
	return o != nil && o.ContainerUsage.IsSet()
}

// SetContainerUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ContainerUsage field.
func (o *UsageAttributionValues) SetContainerUsage(v float64) {
	o.ContainerUsage.Set(&v)
}

// SetContainerUsageNil sets the value for ContainerUsage to be an explicit nil.
func (o *UsageAttributionValues) SetContainerUsageNil() {
	o.ContainerUsage.Set(nil)
}

// UnsetContainerUsage ensures that no value is present for ContainerUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetContainerUsage() {
	o.ContainerUsage.Unset()
}

// GetCspmContainerPercentage returns the CspmContainerPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCspmContainerPercentage() float64 {
	if o == nil || o.CspmContainerPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CspmContainerPercentage.Get()
}

// GetCspmContainerPercentageOk returns a tuple with the CspmContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCspmContainerPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainerPercentage.Get(), o.CspmContainerPercentage.IsSet()
}

// HasCspmContainerPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCspmContainerPercentage() bool {
	return o != nil && o.CspmContainerPercentage.IsSet()
}

// SetCspmContainerPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CspmContainerPercentage field.
func (o *UsageAttributionValues) SetCspmContainerPercentage(v float64) {
	o.CspmContainerPercentage.Set(&v)
}

// SetCspmContainerPercentageNil sets the value for CspmContainerPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetCspmContainerPercentageNil() {
	o.CspmContainerPercentage.Set(nil)
}

// UnsetCspmContainerPercentage ensures that no value is present for CspmContainerPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCspmContainerPercentage() {
	o.CspmContainerPercentage.Unset()
}

// GetCspmContainerUsage returns the CspmContainerUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCspmContainerUsage() float64 {
	if o == nil || o.CspmContainerUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CspmContainerUsage.Get()
}

// GetCspmContainerUsageOk returns a tuple with the CspmContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCspmContainerUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmContainerUsage.Get(), o.CspmContainerUsage.IsSet()
}

// HasCspmContainerUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCspmContainerUsage() bool {
	return o != nil && o.CspmContainerUsage.IsSet()
}

// SetCspmContainerUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CspmContainerUsage field.
func (o *UsageAttributionValues) SetCspmContainerUsage(v float64) {
	o.CspmContainerUsage.Set(&v)
}

// SetCspmContainerUsageNil sets the value for CspmContainerUsage to be an explicit nil.
func (o *UsageAttributionValues) SetCspmContainerUsageNil() {
	o.CspmContainerUsage.Set(nil)
}

// UnsetCspmContainerUsage ensures that no value is present for CspmContainerUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCspmContainerUsage() {
	o.CspmContainerUsage.Unset()
}

// GetCspmHostPercentage returns the CspmHostPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCspmHostPercentage() float64 {
	if o == nil || o.CspmHostPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CspmHostPercentage.Get()
}

// GetCspmHostPercentageOk returns a tuple with the CspmHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCspmHostPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmHostPercentage.Get(), o.CspmHostPercentage.IsSet()
}

// HasCspmHostPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCspmHostPercentage() bool {
	return o != nil && o.CspmHostPercentage.IsSet()
}

// SetCspmHostPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CspmHostPercentage field.
func (o *UsageAttributionValues) SetCspmHostPercentage(v float64) {
	o.CspmHostPercentage.Set(&v)
}

// SetCspmHostPercentageNil sets the value for CspmHostPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetCspmHostPercentageNil() {
	o.CspmHostPercentage.Set(nil)
}

// UnsetCspmHostPercentage ensures that no value is present for CspmHostPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCspmHostPercentage() {
	o.CspmHostPercentage.Unset()
}

// GetCspmHostUsage returns the CspmHostUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCspmHostUsage() float64 {
	if o == nil || o.CspmHostUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CspmHostUsage.Get()
}

// GetCspmHostUsageOk returns a tuple with the CspmHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCspmHostUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CspmHostUsage.Get(), o.CspmHostUsage.IsSet()
}

// HasCspmHostUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCspmHostUsage() bool {
	return o != nil && o.CspmHostUsage.IsSet()
}

// SetCspmHostUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CspmHostUsage field.
func (o *UsageAttributionValues) SetCspmHostUsage(v float64) {
	o.CspmHostUsage.Set(&v)
}

// SetCspmHostUsageNil sets the value for CspmHostUsage to be an explicit nil.
func (o *UsageAttributionValues) SetCspmHostUsageNil() {
	o.CspmHostUsage.Set(nil)
}

// UnsetCspmHostUsage ensures that no value is present for CspmHostUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCspmHostUsage() {
	o.CspmHostUsage.Unset()
}

// GetCustomTimeseriesPercentage returns the CustomTimeseriesPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCustomTimeseriesPercentage() float64 {
	if o == nil || o.CustomTimeseriesPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CustomTimeseriesPercentage.Get()
}

// GetCustomTimeseriesPercentageOk returns a tuple with the CustomTimeseriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCustomTimeseriesPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTimeseriesPercentage.Get(), o.CustomTimeseriesPercentage.IsSet()
}

// HasCustomTimeseriesPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCustomTimeseriesPercentage() bool {
	return o != nil && o.CustomTimeseriesPercentage.IsSet()
}

// SetCustomTimeseriesPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CustomTimeseriesPercentage field.
func (o *UsageAttributionValues) SetCustomTimeseriesPercentage(v float64) {
	o.CustomTimeseriesPercentage.Set(&v)
}

// SetCustomTimeseriesPercentageNil sets the value for CustomTimeseriesPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetCustomTimeseriesPercentageNil() {
	o.CustomTimeseriesPercentage.Set(nil)
}

// UnsetCustomTimeseriesPercentage ensures that no value is present for CustomTimeseriesPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCustomTimeseriesPercentage() {
	o.CustomTimeseriesPercentage.Unset()
}

// GetCustomTimeseriesUsage returns the CustomTimeseriesUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCustomTimeseriesUsage() float64 {
	if o == nil || o.CustomTimeseriesUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CustomTimeseriesUsage.Get()
}

// GetCustomTimeseriesUsageOk returns a tuple with the CustomTimeseriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCustomTimeseriesUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTimeseriesUsage.Get(), o.CustomTimeseriesUsage.IsSet()
}

// HasCustomTimeseriesUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCustomTimeseriesUsage() bool {
	return o != nil && o.CustomTimeseriesUsage.IsSet()
}

// SetCustomTimeseriesUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CustomTimeseriesUsage field.
func (o *UsageAttributionValues) SetCustomTimeseriesUsage(v float64) {
	o.CustomTimeseriesUsage.Set(&v)
}

// SetCustomTimeseriesUsageNil sets the value for CustomTimeseriesUsage to be an explicit nil.
func (o *UsageAttributionValues) SetCustomTimeseriesUsageNil() {
	o.CustomTimeseriesUsage.Set(nil)
}

// UnsetCustomTimeseriesUsage ensures that no value is present for CustomTimeseriesUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCustomTimeseriesUsage() {
	o.CustomTimeseriesUsage.Unset()
}

// GetCwsContainerPercentage returns the CwsContainerPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCwsContainerPercentage() float64 {
	if o == nil || o.CwsContainerPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CwsContainerPercentage.Get()
}

// GetCwsContainerPercentageOk returns a tuple with the CwsContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCwsContainerPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsContainerPercentage.Get(), o.CwsContainerPercentage.IsSet()
}

// HasCwsContainerPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCwsContainerPercentage() bool {
	return o != nil && o.CwsContainerPercentage.IsSet()
}

// SetCwsContainerPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CwsContainerPercentage field.
func (o *UsageAttributionValues) SetCwsContainerPercentage(v float64) {
	o.CwsContainerPercentage.Set(&v)
}

// SetCwsContainerPercentageNil sets the value for CwsContainerPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetCwsContainerPercentageNil() {
	o.CwsContainerPercentage.Set(nil)
}

// UnsetCwsContainerPercentage ensures that no value is present for CwsContainerPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCwsContainerPercentage() {
	o.CwsContainerPercentage.Unset()
}

// GetCwsContainerUsage returns the CwsContainerUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCwsContainerUsage() float64 {
	if o == nil || o.CwsContainerUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CwsContainerUsage.Get()
}

// GetCwsContainerUsageOk returns a tuple with the CwsContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCwsContainerUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsContainerUsage.Get(), o.CwsContainerUsage.IsSet()
}

// HasCwsContainerUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCwsContainerUsage() bool {
	return o != nil && o.CwsContainerUsage.IsSet()
}

// SetCwsContainerUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CwsContainerUsage field.
func (o *UsageAttributionValues) SetCwsContainerUsage(v float64) {
	o.CwsContainerUsage.Set(&v)
}

// SetCwsContainerUsageNil sets the value for CwsContainerUsage to be an explicit nil.
func (o *UsageAttributionValues) SetCwsContainerUsageNil() {
	o.CwsContainerUsage.Set(nil)
}

// UnsetCwsContainerUsage ensures that no value is present for CwsContainerUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCwsContainerUsage() {
	o.CwsContainerUsage.Unset()
}

// GetCwsHostPercentage returns the CwsHostPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCwsHostPercentage() float64 {
	if o == nil || o.CwsHostPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CwsHostPercentage.Get()
}

// GetCwsHostPercentageOk returns a tuple with the CwsHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCwsHostPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsHostPercentage.Get(), o.CwsHostPercentage.IsSet()
}

// HasCwsHostPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCwsHostPercentage() bool {
	return o != nil && o.CwsHostPercentage.IsSet()
}

// SetCwsHostPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the CwsHostPercentage field.
func (o *UsageAttributionValues) SetCwsHostPercentage(v float64) {
	o.CwsHostPercentage.Set(&v)
}

// SetCwsHostPercentageNil sets the value for CwsHostPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetCwsHostPercentageNil() {
	o.CwsHostPercentage.Set(nil)
}

// UnsetCwsHostPercentage ensures that no value is present for CwsHostPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCwsHostPercentage() {
	o.CwsHostPercentage.Unset()
}

// GetCwsHostUsage returns the CwsHostUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetCwsHostUsage() float64 {
	if o == nil || o.CwsHostUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CwsHostUsage.Get()
}

// GetCwsHostUsageOk returns a tuple with the CwsHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetCwsHostUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsHostUsage.Get(), o.CwsHostUsage.IsSet()
}

// HasCwsHostUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasCwsHostUsage() bool {
	return o != nil && o.CwsHostUsage.IsSet()
}

// SetCwsHostUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the CwsHostUsage field.
func (o *UsageAttributionValues) SetCwsHostUsage(v float64) {
	o.CwsHostUsage.Set(&v)
}

// SetCwsHostUsageNil sets the value for CwsHostUsage to be an explicit nil.
func (o *UsageAttributionValues) SetCwsHostUsageNil() {
	o.CwsHostUsage.Set(nil)
}

// UnsetCwsHostUsage ensures that no value is present for CwsHostUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetCwsHostUsage() {
	o.CwsHostUsage.Unset()
}

// GetDbmHostsPercentage returns the DbmHostsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetDbmHostsPercentage() float64 {
	if o == nil || o.DbmHostsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DbmHostsPercentage.Get()
}

// GetDbmHostsPercentageOk returns a tuple with the DbmHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetDbmHostsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmHostsPercentage.Get(), o.DbmHostsPercentage.IsSet()
}

// HasDbmHostsPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasDbmHostsPercentage() bool {
	return o != nil && o.DbmHostsPercentage.IsSet()
}

// SetDbmHostsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the DbmHostsPercentage field.
func (o *UsageAttributionValues) SetDbmHostsPercentage(v float64) {
	o.DbmHostsPercentage.Set(&v)
}

// SetDbmHostsPercentageNil sets the value for DbmHostsPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetDbmHostsPercentageNil() {
	o.DbmHostsPercentage.Set(nil)
}

// UnsetDbmHostsPercentage ensures that no value is present for DbmHostsPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetDbmHostsPercentage() {
	o.DbmHostsPercentage.Unset()
}

// GetDbmHostsUsage returns the DbmHostsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetDbmHostsUsage() float64 {
	if o == nil || o.DbmHostsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DbmHostsUsage.Get()
}

// GetDbmHostsUsageOk returns a tuple with the DbmHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetDbmHostsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmHostsUsage.Get(), o.DbmHostsUsage.IsSet()
}

// HasDbmHostsUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasDbmHostsUsage() bool {
	return o != nil && o.DbmHostsUsage.IsSet()
}

// SetDbmHostsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the DbmHostsUsage field.
func (o *UsageAttributionValues) SetDbmHostsUsage(v float64) {
	o.DbmHostsUsage.Set(&v)
}

// SetDbmHostsUsageNil sets the value for DbmHostsUsage to be an explicit nil.
func (o *UsageAttributionValues) SetDbmHostsUsageNil() {
	o.DbmHostsUsage.Set(nil)
}

// UnsetDbmHostsUsage ensures that no value is present for DbmHostsUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetDbmHostsUsage() {
	o.DbmHostsUsage.Unset()
}

// GetDbmQueriesPercentage returns the DbmQueriesPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetDbmQueriesPercentage() float64 {
	if o == nil || o.DbmQueriesPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DbmQueriesPercentage.Get()
}

// GetDbmQueriesPercentageOk returns a tuple with the DbmQueriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetDbmQueriesPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmQueriesPercentage.Get(), o.DbmQueriesPercentage.IsSet()
}

// HasDbmQueriesPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasDbmQueriesPercentage() bool {
	return o != nil && o.DbmQueriesPercentage.IsSet()
}

// SetDbmQueriesPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the DbmQueriesPercentage field.
func (o *UsageAttributionValues) SetDbmQueriesPercentage(v float64) {
	o.DbmQueriesPercentage.Set(&v)
}

// SetDbmQueriesPercentageNil sets the value for DbmQueriesPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetDbmQueriesPercentageNil() {
	o.DbmQueriesPercentage.Set(nil)
}

// UnsetDbmQueriesPercentage ensures that no value is present for DbmQueriesPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetDbmQueriesPercentage() {
	o.DbmQueriesPercentage.Unset()
}

// GetDbmQueriesUsage returns the DbmQueriesUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetDbmQueriesUsage() float64 {
	if o == nil || o.DbmQueriesUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DbmQueriesUsage.Get()
}

// GetDbmQueriesUsageOk returns a tuple with the DbmQueriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetDbmQueriesUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbmQueriesUsage.Get(), o.DbmQueriesUsage.IsSet()
}

// HasDbmQueriesUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasDbmQueriesUsage() bool {
	return o != nil && o.DbmQueriesUsage.IsSet()
}

// SetDbmQueriesUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the DbmQueriesUsage field.
func (o *UsageAttributionValues) SetDbmQueriesUsage(v float64) {
	o.DbmQueriesUsage.Set(&v)
}

// SetDbmQueriesUsageNil sets the value for DbmQueriesUsage to be an explicit nil.
func (o *UsageAttributionValues) SetDbmQueriesUsageNil() {
	o.DbmQueriesUsage.Set(nil)
}

// UnsetDbmQueriesUsage ensures that no value is present for DbmQueriesUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetDbmQueriesUsage() {
	o.DbmQueriesUsage.Unset()
}

// GetEstimatedIndexedLogsPercentage returns the EstimatedIndexedLogsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedIndexedLogsPercentage() float64 {
	if o == nil || o.EstimatedIndexedLogsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedLogsPercentage.Get()
}

// GetEstimatedIndexedLogsPercentageOk returns a tuple with the EstimatedIndexedLogsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedIndexedLogsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIndexedLogsPercentage.Get(), o.EstimatedIndexedLogsPercentage.IsSet()
}

// HasEstimatedIndexedLogsPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedIndexedLogsPercentage() bool {
	return o != nil && o.EstimatedIndexedLogsPercentage.IsSet()
}

// SetEstimatedIndexedLogsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIndexedLogsPercentage field.
func (o *UsageAttributionValues) SetEstimatedIndexedLogsPercentage(v float64) {
	o.EstimatedIndexedLogsPercentage.Set(&v)
}

// SetEstimatedIndexedLogsPercentageNil sets the value for EstimatedIndexedLogsPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedIndexedLogsPercentageNil() {
	o.EstimatedIndexedLogsPercentage.Set(nil)
}

// UnsetEstimatedIndexedLogsPercentage ensures that no value is present for EstimatedIndexedLogsPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedIndexedLogsPercentage() {
	o.EstimatedIndexedLogsPercentage.Unset()
}

// GetEstimatedIndexedLogsUsage returns the EstimatedIndexedLogsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedIndexedLogsUsage() float64 {
	if o == nil || o.EstimatedIndexedLogsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedLogsUsage.Get()
}

// GetEstimatedIndexedLogsUsageOk returns a tuple with the EstimatedIndexedLogsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedIndexedLogsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIndexedLogsUsage.Get(), o.EstimatedIndexedLogsUsage.IsSet()
}

// HasEstimatedIndexedLogsUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedIndexedLogsUsage() bool {
	return o != nil && o.EstimatedIndexedLogsUsage.IsSet()
}

// SetEstimatedIndexedLogsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIndexedLogsUsage field.
func (o *UsageAttributionValues) SetEstimatedIndexedLogsUsage(v float64) {
	o.EstimatedIndexedLogsUsage.Set(&v)
}

// SetEstimatedIndexedLogsUsageNil sets the value for EstimatedIndexedLogsUsage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedIndexedLogsUsageNil() {
	o.EstimatedIndexedLogsUsage.Set(nil)
}

// UnsetEstimatedIndexedLogsUsage ensures that no value is present for EstimatedIndexedLogsUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedIndexedLogsUsage() {
	o.EstimatedIndexedLogsUsage.Unset()
}

// GetEstimatedIndexedSpansPercentage returns the EstimatedIndexedSpansPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedIndexedSpansPercentage() float64 {
	if o == nil || o.EstimatedIndexedSpansPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedSpansPercentage.Get()
}

// GetEstimatedIndexedSpansPercentageOk returns a tuple with the EstimatedIndexedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedIndexedSpansPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIndexedSpansPercentage.Get(), o.EstimatedIndexedSpansPercentage.IsSet()
}

// HasEstimatedIndexedSpansPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedIndexedSpansPercentage() bool {
	return o != nil && o.EstimatedIndexedSpansPercentage.IsSet()
}

// SetEstimatedIndexedSpansPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIndexedSpansPercentage field.
func (o *UsageAttributionValues) SetEstimatedIndexedSpansPercentage(v float64) {
	o.EstimatedIndexedSpansPercentage.Set(&v)
}

// SetEstimatedIndexedSpansPercentageNil sets the value for EstimatedIndexedSpansPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedIndexedSpansPercentageNil() {
	o.EstimatedIndexedSpansPercentage.Set(nil)
}

// UnsetEstimatedIndexedSpansPercentage ensures that no value is present for EstimatedIndexedSpansPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedIndexedSpansPercentage() {
	o.EstimatedIndexedSpansPercentage.Unset()
}

// GetEstimatedIndexedSpansUsage returns the EstimatedIndexedSpansUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedIndexedSpansUsage() float64 {
	if o == nil || o.EstimatedIndexedSpansUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedSpansUsage.Get()
}

// GetEstimatedIndexedSpansUsageOk returns a tuple with the EstimatedIndexedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedIndexedSpansUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIndexedSpansUsage.Get(), o.EstimatedIndexedSpansUsage.IsSet()
}

// HasEstimatedIndexedSpansUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedIndexedSpansUsage() bool {
	return o != nil && o.EstimatedIndexedSpansUsage.IsSet()
}

// SetEstimatedIndexedSpansUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIndexedSpansUsage field.
func (o *UsageAttributionValues) SetEstimatedIndexedSpansUsage(v float64) {
	o.EstimatedIndexedSpansUsage.Set(&v)
}

// SetEstimatedIndexedSpansUsageNil sets the value for EstimatedIndexedSpansUsage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedIndexedSpansUsageNil() {
	o.EstimatedIndexedSpansUsage.Set(nil)
}

// UnsetEstimatedIndexedSpansUsage ensures that no value is present for EstimatedIndexedSpansUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedIndexedSpansUsage() {
	o.EstimatedIndexedSpansUsage.Unset()
}

// GetEstimatedIngestedLogsPercentage returns the EstimatedIngestedLogsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedIngestedLogsPercentage() float64 {
	if o == nil || o.EstimatedIngestedLogsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedLogsPercentage.Get()
}

// GetEstimatedIngestedLogsPercentageOk returns a tuple with the EstimatedIngestedLogsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedIngestedLogsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIngestedLogsPercentage.Get(), o.EstimatedIngestedLogsPercentage.IsSet()
}

// HasEstimatedIngestedLogsPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedIngestedLogsPercentage() bool {
	return o != nil && o.EstimatedIngestedLogsPercentage.IsSet()
}

// SetEstimatedIngestedLogsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIngestedLogsPercentage field.
func (o *UsageAttributionValues) SetEstimatedIngestedLogsPercentage(v float64) {
	o.EstimatedIngestedLogsPercentage.Set(&v)
}

// SetEstimatedIngestedLogsPercentageNil sets the value for EstimatedIngestedLogsPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedIngestedLogsPercentageNil() {
	o.EstimatedIngestedLogsPercentage.Set(nil)
}

// UnsetEstimatedIngestedLogsPercentage ensures that no value is present for EstimatedIngestedLogsPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedIngestedLogsPercentage() {
	o.EstimatedIngestedLogsPercentage.Unset()
}

// GetEstimatedIngestedLogsUsage returns the EstimatedIngestedLogsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedIngestedLogsUsage() float64 {
	if o == nil || o.EstimatedIngestedLogsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedLogsUsage.Get()
}

// GetEstimatedIngestedLogsUsageOk returns a tuple with the EstimatedIngestedLogsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedIngestedLogsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIngestedLogsUsage.Get(), o.EstimatedIngestedLogsUsage.IsSet()
}

// HasEstimatedIngestedLogsUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedIngestedLogsUsage() bool {
	return o != nil && o.EstimatedIngestedLogsUsage.IsSet()
}

// SetEstimatedIngestedLogsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIngestedLogsUsage field.
func (o *UsageAttributionValues) SetEstimatedIngestedLogsUsage(v float64) {
	o.EstimatedIngestedLogsUsage.Set(&v)
}

// SetEstimatedIngestedLogsUsageNil sets the value for EstimatedIngestedLogsUsage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedIngestedLogsUsageNil() {
	o.EstimatedIngestedLogsUsage.Set(nil)
}

// UnsetEstimatedIngestedLogsUsage ensures that no value is present for EstimatedIngestedLogsUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedIngestedLogsUsage() {
	o.EstimatedIngestedLogsUsage.Unset()
}

// GetEstimatedIngestedSpansPercentage returns the EstimatedIngestedSpansPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedIngestedSpansPercentage() float64 {
	if o == nil || o.EstimatedIngestedSpansPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedSpansPercentage.Get()
}

// GetEstimatedIngestedSpansPercentageOk returns a tuple with the EstimatedIngestedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedIngestedSpansPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIngestedSpansPercentage.Get(), o.EstimatedIngestedSpansPercentage.IsSet()
}

// HasEstimatedIngestedSpansPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedIngestedSpansPercentage() bool {
	return o != nil && o.EstimatedIngestedSpansPercentage.IsSet()
}

// SetEstimatedIngestedSpansPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIngestedSpansPercentage field.
func (o *UsageAttributionValues) SetEstimatedIngestedSpansPercentage(v float64) {
	o.EstimatedIngestedSpansPercentage.Set(&v)
}

// SetEstimatedIngestedSpansPercentageNil sets the value for EstimatedIngestedSpansPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedIngestedSpansPercentageNil() {
	o.EstimatedIngestedSpansPercentage.Set(nil)
}

// UnsetEstimatedIngestedSpansPercentage ensures that no value is present for EstimatedIngestedSpansPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedIngestedSpansPercentage() {
	o.EstimatedIngestedSpansPercentage.Unset()
}

// GetEstimatedIngestedSpansUsage returns the EstimatedIngestedSpansUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedIngestedSpansUsage() float64 {
	if o == nil || o.EstimatedIngestedSpansUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedSpansUsage.Get()
}

// GetEstimatedIngestedSpansUsageOk returns a tuple with the EstimatedIngestedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedIngestedSpansUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedIngestedSpansUsage.Get(), o.EstimatedIngestedSpansUsage.IsSet()
}

// HasEstimatedIngestedSpansUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedIngestedSpansUsage() bool {
	return o != nil && o.EstimatedIngestedSpansUsage.IsSet()
}

// SetEstimatedIngestedSpansUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedIngestedSpansUsage field.
func (o *UsageAttributionValues) SetEstimatedIngestedSpansUsage(v float64) {
	o.EstimatedIngestedSpansUsage.Set(&v)
}

// SetEstimatedIngestedSpansUsageNil sets the value for EstimatedIngestedSpansUsage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedIngestedSpansUsageNil() {
	o.EstimatedIngestedSpansUsage.Set(nil)
}

// UnsetEstimatedIngestedSpansUsage ensures that no value is present for EstimatedIngestedSpansUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedIngestedSpansUsage() {
	o.EstimatedIngestedSpansUsage.Unset()
}

// GetEstimatedRumSessionsPercentage returns the EstimatedRumSessionsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedRumSessionsPercentage() float64 {
	if o == nil || o.EstimatedRumSessionsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedRumSessionsPercentage.Get()
}

// GetEstimatedRumSessionsPercentageOk returns a tuple with the EstimatedRumSessionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedRumSessionsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedRumSessionsPercentage.Get(), o.EstimatedRumSessionsPercentage.IsSet()
}

// HasEstimatedRumSessionsPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedRumSessionsPercentage() bool {
	return o != nil && o.EstimatedRumSessionsPercentage.IsSet()
}

// SetEstimatedRumSessionsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedRumSessionsPercentage field.
func (o *UsageAttributionValues) SetEstimatedRumSessionsPercentage(v float64) {
	o.EstimatedRumSessionsPercentage.Set(&v)
}

// SetEstimatedRumSessionsPercentageNil sets the value for EstimatedRumSessionsPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedRumSessionsPercentageNil() {
	o.EstimatedRumSessionsPercentage.Set(nil)
}

// UnsetEstimatedRumSessionsPercentage ensures that no value is present for EstimatedRumSessionsPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedRumSessionsPercentage() {
	o.EstimatedRumSessionsPercentage.Unset()
}

// GetEstimatedRumSessionsUsage returns the EstimatedRumSessionsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetEstimatedRumSessionsUsage() float64 {
	if o == nil || o.EstimatedRumSessionsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedRumSessionsUsage.Get()
}

// GetEstimatedRumSessionsUsageOk returns a tuple with the EstimatedRumSessionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetEstimatedRumSessionsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedRumSessionsUsage.Get(), o.EstimatedRumSessionsUsage.IsSet()
}

// HasEstimatedRumSessionsUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasEstimatedRumSessionsUsage() bool {
	return o != nil && o.EstimatedRumSessionsUsage.IsSet()
}

// SetEstimatedRumSessionsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the EstimatedRumSessionsUsage field.
func (o *UsageAttributionValues) SetEstimatedRumSessionsUsage(v float64) {
	o.EstimatedRumSessionsUsage.Set(&v)
}

// SetEstimatedRumSessionsUsageNil sets the value for EstimatedRumSessionsUsage to be an explicit nil.
func (o *UsageAttributionValues) SetEstimatedRumSessionsUsageNil() {
	o.EstimatedRumSessionsUsage.Set(nil)
}

// UnsetEstimatedRumSessionsUsage ensures that no value is present for EstimatedRumSessionsUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetEstimatedRumSessionsUsage() {
	o.EstimatedRumSessionsUsage.Unset()
}

// GetInfraHostPercentage returns the InfraHostPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetInfraHostPercentage() float64 {
	if o == nil || o.InfraHostPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.InfraHostPercentage.Get()
}

// GetInfraHostPercentageOk returns a tuple with the InfraHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetInfraHostPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfraHostPercentage.Get(), o.InfraHostPercentage.IsSet()
}

// HasInfraHostPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasInfraHostPercentage() bool {
	return o != nil && o.InfraHostPercentage.IsSet()
}

// SetInfraHostPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the InfraHostPercentage field.
func (o *UsageAttributionValues) SetInfraHostPercentage(v float64) {
	o.InfraHostPercentage.Set(&v)
}

// SetInfraHostPercentageNil sets the value for InfraHostPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetInfraHostPercentageNil() {
	o.InfraHostPercentage.Set(nil)
}

// UnsetInfraHostPercentage ensures that no value is present for InfraHostPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetInfraHostPercentage() {
	o.InfraHostPercentage.Unset()
}

// GetInfraHostUsage returns the InfraHostUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetInfraHostUsage() float64 {
	if o == nil || o.InfraHostUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.InfraHostUsage.Get()
}

// GetInfraHostUsageOk returns a tuple with the InfraHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetInfraHostUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfraHostUsage.Get(), o.InfraHostUsage.IsSet()
}

// HasInfraHostUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasInfraHostUsage() bool {
	return o != nil && o.InfraHostUsage.IsSet()
}

// SetInfraHostUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the InfraHostUsage field.
func (o *UsageAttributionValues) SetInfraHostUsage(v float64) {
	o.InfraHostUsage.Set(&v)
}

// SetInfraHostUsageNil sets the value for InfraHostUsage to be an explicit nil.
func (o *UsageAttributionValues) SetInfraHostUsageNil() {
	o.InfraHostUsage.Set(nil)
}

// UnsetInfraHostUsage ensures that no value is present for InfraHostUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetInfraHostUsage() {
	o.InfraHostUsage.Unset()
}

// GetLambdaFunctionsPercentage returns the LambdaFunctionsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetLambdaFunctionsPercentage() float64 {
	if o == nil || o.LambdaFunctionsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LambdaFunctionsPercentage.Get()
}

// GetLambdaFunctionsPercentageOk returns a tuple with the LambdaFunctionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetLambdaFunctionsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LambdaFunctionsPercentage.Get(), o.LambdaFunctionsPercentage.IsSet()
}

// HasLambdaFunctionsPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasLambdaFunctionsPercentage() bool {
	return o != nil && o.LambdaFunctionsPercentage.IsSet()
}

// SetLambdaFunctionsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the LambdaFunctionsPercentage field.
func (o *UsageAttributionValues) SetLambdaFunctionsPercentage(v float64) {
	o.LambdaFunctionsPercentage.Set(&v)
}

// SetLambdaFunctionsPercentageNil sets the value for LambdaFunctionsPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetLambdaFunctionsPercentageNil() {
	o.LambdaFunctionsPercentage.Set(nil)
}

// UnsetLambdaFunctionsPercentage ensures that no value is present for LambdaFunctionsPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetLambdaFunctionsPercentage() {
	o.LambdaFunctionsPercentage.Unset()
}

// GetLambdaFunctionsUsage returns the LambdaFunctionsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetLambdaFunctionsUsage() float64 {
	if o == nil || o.LambdaFunctionsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LambdaFunctionsUsage.Get()
}

// GetLambdaFunctionsUsageOk returns a tuple with the LambdaFunctionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetLambdaFunctionsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LambdaFunctionsUsage.Get(), o.LambdaFunctionsUsage.IsSet()
}

// HasLambdaFunctionsUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasLambdaFunctionsUsage() bool {
	return o != nil && o.LambdaFunctionsUsage.IsSet()
}

// SetLambdaFunctionsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the LambdaFunctionsUsage field.
func (o *UsageAttributionValues) SetLambdaFunctionsUsage(v float64) {
	o.LambdaFunctionsUsage.Set(&v)
}

// SetLambdaFunctionsUsageNil sets the value for LambdaFunctionsUsage to be an explicit nil.
func (o *UsageAttributionValues) SetLambdaFunctionsUsageNil() {
	o.LambdaFunctionsUsage.Set(nil)
}

// UnsetLambdaFunctionsUsage ensures that no value is present for LambdaFunctionsUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetLambdaFunctionsUsage() {
	o.LambdaFunctionsUsage.Unset()
}

// GetLambdaInvocationsPercentage returns the LambdaInvocationsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetLambdaInvocationsPercentage() float64 {
	if o == nil || o.LambdaInvocationsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LambdaInvocationsPercentage.Get()
}

// GetLambdaInvocationsPercentageOk returns a tuple with the LambdaInvocationsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetLambdaInvocationsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LambdaInvocationsPercentage.Get(), o.LambdaInvocationsPercentage.IsSet()
}

// HasLambdaInvocationsPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasLambdaInvocationsPercentage() bool {
	return o != nil && o.LambdaInvocationsPercentage.IsSet()
}

// SetLambdaInvocationsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the LambdaInvocationsPercentage field.
func (o *UsageAttributionValues) SetLambdaInvocationsPercentage(v float64) {
	o.LambdaInvocationsPercentage.Set(&v)
}

// SetLambdaInvocationsPercentageNil sets the value for LambdaInvocationsPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetLambdaInvocationsPercentageNil() {
	o.LambdaInvocationsPercentage.Set(nil)
}

// UnsetLambdaInvocationsPercentage ensures that no value is present for LambdaInvocationsPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetLambdaInvocationsPercentage() {
	o.LambdaInvocationsPercentage.Unset()
}

// GetLambdaInvocationsUsage returns the LambdaInvocationsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetLambdaInvocationsUsage() float64 {
	if o == nil || o.LambdaInvocationsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LambdaInvocationsUsage.Get()
}

// GetLambdaInvocationsUsageOk returns a tuple with the LambdaInvocationsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetLambdaInvocationsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LambdaInvocationsUsage.Get(), o.LambdaInvocationsUsage.IsSet()
}

// HasLambdaInvocationsUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasLambdaInvocationsUsage() bool {
	return o != nil && o.LambdaInvocationsUsage.IsSet()
}

// SetLambdaInvocationsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the LambdaInvocationsUsage field.
func (o *UsageAttributionValues) SetLambdaInvocationsUsage(v float64) {
	o.LambdaInvocationsUsage.Set(&v)
}

// SetLambdaInvocationsUsageNil sets the value for LambdaInvocationsUsage to be an explicit nil.
func (o *UsageAttributionValues) SetLambdaInvocationsUsageNil() {
	o.LambdaInvocationsUsage.Set(nil)
}

// UnsetLambdaInvocationsUsage ensures that no value is present for LambdaInvocationsUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetLambdaInvocationsUsage() {
	o.LambdaInvocationsUsage.Unset()
}

// GetNpmHostPercentage returns the NpmHostPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetNpmHostPercentage() float64 {
	if o == nil || o.NpmHostPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NpmHostPercentage.Get()
}

// GetNpmHostPercentageOk returns a tuple with the NpmHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetNpmHostPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NpmHostPercentage.Get(), o.NpmHostPercentage.IsSet()
}

// HasNpmHostPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasNpmHostPercentage() bool {
	return o != nil && o.NpmHostPercentage.IsSet()
}

// SetNpmHostPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the NpmHostPercentage field.
func (o *UsageAttributionValues) SetNpmHostPercentage(v float64) {
	o.NpmHostPercentage.Set(&v)
}

// SetNpmHostPercentageNil sets the value for NpmHostPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetNpmHostPercentageNil() {
	o.NpmHostPercentage.Set(nil)
}

// UnsetNpmHostPercentage ensures that no value is present for NpmHostPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetNpmHostPercentage() {
	o.NpmHostPercentage.Unset()
}

// GetNpmHostUsage returns the NpmHostUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetNpmHostUsage() float64 {
	if o == nil || o.NpmHostUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NpmHostUsage.Get()
}

// GetNpmHostUsageOk returns a tuple with the NpmHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetNpmHostUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NpmHostUsage.Get(), o.NpmHostUsage.IsSet()
}

// HasNpmHostUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasNpmHostUsage() bool {
	return o != nil && o.NpmHostUsage.IsSet()
}

// SetNpmHostUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the NpmHostUsage field.
func (o *UsageAttributionValues) SetNpmHostUsage(v float64) {
	o.NpmHostUsage.Set(&v)
}

// SetNpmHostUsageNil sets the value for NpmHostUsage to be an explicit nil.
func (o *UsageAttributionValues) SetNpmHostUsageNil() {
	o.NpmHostUsage.Set(nil)
}

// UnsetNpmHostUsage ensures that no value is present for NpmHostUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetNpmHostUsage() {
	o.NpmHostUsage.Unset()
}

// GetProfiledContainerPercentage returns the ProfiledContainerPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetProfiledContainerPercentage() float64 {
	if o == nil || o.ProfiledContainerPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledContainerPercentage.Get()
}

// GetProfiledContainerPercentageOk returns a tuple with the ProfiledContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetProfiledContainerPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledContainerPercentage.Get(), o.ProfiledContainerPercentage.IsSet()
}

// HasProfiledContainerPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasProfiledContainerPercentage() bool {
	return o != nil && o.ProfiledContainerPercentage.IsSet()
}

// SetProfiledContainerPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledContainerPercentage field.
func (o *UsageAttributionValues) SetProfiledContainerPercentage(v float64) {
	o.ProfiledContainerPercentage.Set(&v)
}

// SetProfiledContainerPercentageNil sets the value for ProfiledContainerPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetProfiledContainerPercentageNil() {
	o.ProfiledContainerPercentage.Set(nil)
}

// UnsetProfiledContainerPercentage ensures that no value is present for ProfiledContainerPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetProfiledContainerPercentage() {
	o.ProfiledContainerPercentage.Unset()
}

// GetProfiledContainerUsage returns the ProfiledContainerUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetProfiledContainerUsage() float64 {
	if o == nil || o.ProfiledContainerUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledContainerUsage.Get()
}

// GetProfiledContainerUsageOk returns a tuple with the ProfiledContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetProfiledContainerUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledContainerUsage.Get(), o.ProfiledContainerUsage.IsSet()
}

// HasProfiledContainerUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasProfiledContainerUsage() bool {
	return o != nil && o.ProfiledContainerUsage.IsSet()
}

// SetProfiledContainerUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledContainerUsage field.
func (o *UsageAttributionValues) SetProfiledContainerUsage(v float64) {
	o.ProfiledContainerUsage.Set(&v)
}

// SetProfiledContainerUsageNil sets the value for ProfiledContainerUsage to be an explicit nil.
func (o *UsageAttributionValues) SetProfiledContainerUsageNil() {
	o.ProfiledContainerUsage.Set(nil)
}

// UnsetProfiledContainerUsage ensures that no value is present for ProfiledContainerUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetProfiledContainerUsage() {
	o.ProfiledContainerUsage.Unset()
}

// GetProfiledHostsPercentage returns the ProfiledHostsPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetProfiledHostsPercentage() float64 {
	if o == nil || o.ProfiledHostsPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledHostsPercentage.Get()
}

// GetProfiledHostsPercentageOk returns a tuple with the ProfiledHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetProfiledHostsPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledHostsPercentage.Get(), o.ProfiledHostsPercentage.IsSet()
}

// HasProfiledHostsPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasProfiledHostsPercentage() bool {
	return o != nil && o.ProfiledHostsPercentage.IsSet()
}

// SetProfiledHostsPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledHostsPercentage field.
func (o *UsageAttributionValues) SetProfiledHostsPercentage(v float64) {
	o.ProfiledHostsPercentage.Set(&v)
}

// SetProfiledHostsPercentageNil sets the value for ProfiledHostsPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetProfiledHostsPercentageNil() {
	o.ProfiledHostsPercentage.Set(nil)
}

// UnsetProfiledHostsPercentage ensures that no value is present for ProfiledHostsPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetProfiledHostsPercentage() {
	o.ProfiledHostsPercentage.Unset()
}

// GetProfiledHostsUsage returns the ProfiledHostsUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetProfiledHostsUsage() float64 {
	if o == nil || o.ProfiledHostsUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledHostsUsage.Get()
}

// GetProfiledHostsUsageOk returns a tuple with the ProfiledHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetProfiledHostsUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfiledHostsUsage.Get(), o.ProfiledHostsUsage.IsSet()
}

// HasProfiledHostsUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasProfiledHostsUsage() bool {
	return o != nil && o.ProfiledHostsUsage.IsSet()
}

// SetProfiledHostsUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the ProfiledHostsUsage field.
func (o *UsageAttributionValues) SetProfiledHostsUsage(v float64) {
	o.ProfiledHostsUsage.Set(&v)
}

// SetProfiledHostsUsageNil sets the value for ProfiledHostsUsage to be an explicit nil.
func (o *UsageAttributionValues) SetProfiledHostsUsageNil() {
	o.ProfiledHostsUsage.Set(nil)
}

// UnsetProfiledHostsUsage ensures that no value is present for ProfiledHostsUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetProfiledHostsUsage() {
	o.ProfiledHostsUsage.Unset()
}

// GetSnmpPercentage returns the SnmpPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetSnmpPercentage() float64 {
	if o == nil || o.SnmpPercentage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SnmpPercentage.Get()
}

// GetSnmpPercentageOk returns a tuple with the SnmpPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetSnmpPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnmpPercentage.Get(), o.SnmpPercentage.IsSet()
}

// HasSnmpPercentage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasSnmpPercentage() bool {
	return o != nil && o.SnmpPercentage.IsSet()
}

// SetSnmpPercentage gets a reference to the given datadog.NullableFloat64 and assigns it to the SnmpPercentage field.
func (o *UsageAttributionValues) SetSnmpPercentage(v float64) {
	o.SnmpPercentage.Set(&v)
}

// SetSnmpPercentageNil sets the value for SnmpPercentage to be an explicit nil.
func (o *UsageAttributionValues) SetSnmpPercentageNil() {
	o.SnmpPercentage.Set(nil)
}

// UnsetSnmpPercentage ensures that no value is present for SnmpPercentage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetSnmpPercentage() {
	o.SnmpPercentage.Unset()
}

// GetSnmpUsage returns the SnmpUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAttributionValues) GetSnmpUsage() float64 {
	if o == nil || o.SnmpUsage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SnmpUsage.Get()
}

// GetSnmpUsageOk returns a tuple with the SnmpUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAttributionValues) GetSnmpUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnmpUsage.Get(), o.SnmpUsage.IsSet()
}

// HasSnmpUsage returns a boolean if a field has been set.
func (o *UsageAttributionValues) HasSnmpUsage() bool {
	return o != nil && o.SnmpUsage.IsSet()
}

// SetSnmpUsage gets a reference to the given datadog.NullableFloat64 and assigns it to the SnmpUsage field.
func (o *UsageAttributionValues) SetSnmpUsage(v float64) {
	o.SnmpUsage.Set(&v)
}

// SetSnmpUsageNil sets the value for SnmpUsage to be an explicit nil.
func (o *UsageAttributionValues) SetSnmpUsageNil() {
	o.SnmpUsage.Set(nil)
}

// UnsetSnmpUsage ensures that no value is present for SnmpUsage, not even an explicit nil.
func (o *UsageAttributionValues) UnsetSnmpUsage() {
	o.SnmpUsage.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageAttributionValues) MarshalJSON() ([]byte, error) {
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
	if o.ContainerPercentage.IsSet() {
		toSerialize["container_percentage"] = o.ContainerPercentage.Get()
	}
	if o.ContainerUsage.IsSet() {
		toSerialize["container_usage"] = o.ContainerUsage.Get()
	}
	if o.CspmContainerPercentage.IsSet() {
		toSerialize["cspm_container_percentage"] = o.CspmContainerPercentage.Get()
	}
	if o.CspmContainerUsage.IsSet() {
		toSerialize["cspm_container_usage"] = o.CspmContainerUsage.Get()
	}
	if o.CspmHostPercentage.IsSet() {
		toSerialize["cspm_host_percentage"] = o.CspmHostPercentage.Get()
	}
	if o.CspmHostUsage.IsSet() {
		toSerialize["cspm_host_usage"] = o.CspmHostUsage.Get()
	}
	if o.CustomTimeseriesPercentage.IsSet() {
		toSerialize["custom_timeseries_percentage"] = o.CustomTimeseriesPercentage.Get()
	}
	if o.CustomTimeseriesUsage.IsSet() {
		toSerialize["custom_timeseries_usage"] = o.CustomTimeseriesUsage.Get()
	}
	if o.CwsContainerPercentage.IsSet() {
		toSerialize["cws_container_percentage"] = o.CwsContainerPercentage.Get()
	}
	if o.CwsContainerUsage.IsSet() {
		toSerialize["cws_container_usage"] = o.CwsContainerUsage.Get()
	}
	if o.CwsHostPercentage.IsSet() {
		toSerialize["cws_host_percentage"] = o.CwsHostPercentage.Get()
	}
	if o.CwsHostUsage.IsSet() {
		toSerialize["cws_host_usage"] = o.CwsHostUsage.Get()
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
	if o.InfraHostPercentage.IsSet() {
		toSerialize["infra_host_percentage"] = o.InfraHostPercentage.Get()
	}
	if o.InfraHostUsage.IsSet() {
		toSerialize["infra_host_usage"] = o.InfraHostUsage.Get()
	}
	if o.LambdaFunctionsPercentage.IsSet() {
		toSerialize["lambda_functions_percentage"] = o.LambdaFunctionsPercentage.Get()
	}
	if o.LambdaFunctionsUsage.IsSet() {
		toSerialize["lambda_functions_usage"] = o.LambdaFunctionsUsage.Get()
	}
	if o.LambdaInvocationsPercentage.IsSet() {
		toSerialize["lambda_invocations_percentage"] = o.LambdaInvocationsPercentage.Get()
	}
	if o.LambdaInvocationsUsage.IsSet() {
		toSerialize["lambda_invocations_usage"] = o.LambdaInvocationsUsage.Get()
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
	if o.ProfiledHostsPercentage.IsSet() {
		toSerialize["profiled_hosts_percentage"] = o.ProfiledHostsPercentage.Get()
	}
	if o.ProfiledHostsUsage.IsSet() {
		toSerialize["profiled_hosts_usage"] = o.ProfiledHostsUsage.Get()
	}
	if o.SnmpPercentage.IsSet() {
		toSerialize["snmp_percentage"] = o.SnmpPercentage.Get()
	}
	if o.SnmpUsage.IsSet() {
		toSerialize["snmp_usage"] = o.SnmpUsage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageAttributionValues) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		ApiPercentage                    datadog.NullableFloat64 `json:"api_percentage,omitempty"`
		ApiUsage                         datadog.NullableFloat64 `json:"api_usage,omitempty"`
		ApmFargatePercentage             datadog.NullableFloat64 `json:"apm_fargate_percentage,omitempty"`
		ApmFargateUsage                  datadog.NullableFloat64 `json:"apm_fargate_usage,omitempty"`
		ApmHostPercentage                datadog.NullableFloat64 `json:"apm_host_percentage,omitempty"`
		ApmHostUsage                     datadog.NullableFloat64 `json:"apm_host_usage,omitempty"`
		AppsecFargatePercentage          datadog.NullableFloat64 `json:"appsec_fargate_percentage,omitempty"`
		AppsecFargateUsage               datadog.NullableFloat64 `json:"appsec_fargate_usage,omitempty"`
		AppsecPercentage                 datadog.NullableFloat64 `json:"appsec_percentage,omitempty"`
		AppsecUsage                      datadog.NullableFloat64 `json:"appsec_usage,omitempty"`
		BrowserPercentage                datadog.NullableFloat64 `json:"browser_percentage,omitempty"`
		BrowserUsage                     datadog.NullableFloat64 `json:"browser_usage,omitempty"`
		ContainerPercentage              datadog.NullableFloat64 `json:"container_percentage,omitempty"`
		ContainerUsage                   datadog.NullableFloat64 `json:"container_usage,omitempty"`
		CspmContainerPercentage          datadog.NullableFloat64 `json:"cspm_container_percentage,omitempty"`
		CspmContainerUsage               datadog.NullableFloat64 `json:"cspm_container_usage,omitempty"`
		CspmHostPercentage               datadog.NullableFloat64 `json:"cspm_host_percentage,omitempty"`
		CspmHostUsage                    datadog.NullableFloat64 `json:"cspm_host_usage,omitempty"`
		CustomTimeseriesPercentage       datadog.NullableFloat64 `json:"custom_timeseries_percentage,omitempty"`
		CustomTimeseriesUsage            datadog.NullableFloat64 `json:"custom_timeseries_usage,omitempty"`
		CwsContainerPercentage           datadog.NullableFloat64 `json:"cws_container_percentage,omitempty"`
		CwsContainerUsage                datadog.NullableFloat64 `json:"cws_container_usage,omitempty"`
		CwsHostPercentage                datadog.NullableFloat64 `json:"cws_host_percentage,omitempty"`
		CwsHostUsage                     datadog.NullableFloat64 `json:"cws_host_usage,omitempty"`
		DbmHostsPercentage               datadog.NullableFloat64 `json:"dbm_hosts_percentage,omitempty"`
		DbmHostsUsage                    datadog.NullableFloat64 `json:"dbm_hosts_usage,omitempty"`
		DbmQueriesPercentage             datadog.NullableFloat64 `json:"dbm_queries_percentage,omitempty"`
		DbmQueriesUsage                  datadog.NullableFloat64 `json:"dbm_queries_usage,omitempty"`
		EstimatedIndexedLogsPercentage   datadog.NullableFloat64 `json:"estimated_indexed_logs_percentage,omitempty"`
		EstimatedIndexedLogsUsage        datadog.NullableFloat64 `json:"estimated_indexed_logs_usage,omitempty"`
		EstimatedIndexedSpansPercentage  datadog.NullableFloat64 `json:"estimated_indexed_spans_percentage,omitempty"`
		EstimatedIndexedSpansUsage       datadog.NullableFloat64 `json:"estimated_indexed_spans_usage,omitempty"`
		EstimatedIngestedLogsPercentage  datadog.NullableFloat64 `json:"estimated_ingested_logs_percentage,omitempty"`
		EstimatedIngestedLogsUsage       datadog.NullableFloat64 `json:"estimated_ingested_logs_usage,omitempty"`
		EstimatedIngestedSpansPercentage datadog.NullableFloat64 `json:"estimated_ingested_spans_percentage,omitempty"`
		EstimatedIngestedSpansUsage      datadog.NullableFloat64 `json:"estimated_ingested_spans_usage,omitempty"`
		EstimatedRumSessionsPercentage   datadog.NullableFloat64 `json:"estimated_rum_sessions_percentage,omitempty"`
		EstimatedRumSessionsUsage        datadog.NullableFloat64 `json:"estimated_rum_sessions_usage,omitempty"`
		InfraHostPercentage              datadog.NullableFloat64 `json:"infra_host_percentage,omitempty"`
		InfraHostUsage                   datadog.NullableFloat64 `json:"infra_host_usage,omitempty"`
		LambdaFunctionsPercentage        datadog.NullableFloat64 `json:"lambda_functions_percentage,omitempty"`
		LambdaFunctionsUsage             datadog.NullableFloat64 `json:"lambda_functions_usage,omitempty"`
		LambdaInvocationsPercentage      datadog.NullableFloat64 `json:"lambda_invocations_percentage,omitempty"`
		LambdaInvocationsUsage           datadog.NullableFloat64 `json:"lambda_invocations_usage,omitempty"`
		NpmHostPercentage                datadog.NullableFloat64 `json:"npm_host_percentage,omitempty"`
		NpmHostUsage                     datadog.NullableFloat64 `json:"npm_host_usage,omitempty"`
		ProfiledContainerPercentage      datadog.NullableFloat64 `json:"profiled_container_percentage,omitempty"`
		ProfiledContainerUsage           datadog.NullableFloat64 `json:"profiled_container_usage,omitempty"`
		ProfiledHostsPercentage          datadog.NullableFloat64 `json:"profiled_hosts_percentage,omitempty"`
		ProfiledHostsUsage               datadog.NullableFloat64 `json:"profiled_hosts_usage,omitempty"`
		SnmpPercentage                   datadog.NullableFloat64 `json:"snmp_percentage,omitempty"`
		SnmpUsage                        datadog.NullableFloat64 `json:"snmp_usage,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"api_percentage", "api_usage", "apm_fargate_percentage", "apm_fargate_usage", "apm_host_percentage", "apm_host_usage", "appsec_fargate_percentage", "appsec_fargate_usage", "appsec_percentage", "appsec_usage", "browser_percentage", "browser_usage", "container_percentage", "container_usage", "cspm_container_percentage", "cspm_container_usage", "cspm_host_percentage", "cspm_host_usage", "custom_timeseries_percentage", "custom_timeseries_usage", "cws_container_percentage", "cws_container_usage", "cws_host_percentage", "cws_host_usage", "dbm_hosts_percentage", "dbm_hosts_usage", "dbm_queries_percentage", "dbm_queries_usage", "estimated_indexed_logs_percentage", "estimated_indexed_logs_usage", "estimated_indexed_spans_percentage", "estimated_indexed_spans_usage", "estimated_ingested_logs_percentage", "estimated_ingested_logs_usage", "estimated_ingested_spans_percentage", "estimated_ingested_spans_usage", "estimated_rum_sessions_percentage", "estimated_rum_sessions_usage", "infra_host_percentage", "infra_host_usage", "lambda_functions_percentage", "lambda_functions_usage", "lambda_invocations_percentage", "lambda_invocations_usage", "npm_host_percentage", "npm_host_usage", "profiled_container_percentage", "profiled_container_usage", "profiled_hosts_percentage", "profiled_hosts_usage", "snmp_percentage", "snmp_usage"})
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
	o.ContainerPercentage = all.ContainerPercentage
	o.ContainerUsage = all.ContainerUsage
	o.CspmContainerPercentage = all.CspmContainerPercentage
	o.CspmContainerUsage = all.CspmContainerUsage
	o.CspmHostPercentage = all.CspmHostPercentage
	o.CspmHostUsage = all.CspmHostUsage
	o.CustomTimeseriesPercentage = all.CustomTimeseriesPercentage
	o.CustomTimeseriesUsage = all.CustomTimeseriesUsage
	o.CwsContainerPercentage = all.CwsContainerPercentage
	o.CwsContainerUsage = all.CwsContainerUsage
	o.CwsHostPercentage = all.CwsHostPercentage
	o.CwsHostUsage = all.CwsHostUsage
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
	o.InfraHostPercentage = all.InfraHostPercentage
	o.InfraHostUsage = all.InfraHostUsage
	o.LambdaFunctionsPercentage = all.LambdaFunctionsPercentage
	o.LambdaFunctionsUsage = all.LambdaFunctionsUsage
	o.LambdaInvocationsPercentage = all.LambdaInvocationsPercentage
	o.LambdaInvocationsUsage = all.LambdaInvocationsUsage
	o.NpmHostPercentage = all.NpmHostPercentage
	o.NpmHostUsage = all.NpmHostUsage
	o.ProfiledContainerPercentage = all.ProfiledContainerPercentage
	o.ProfiledContainerUsage = all.ProfiledContainerUsage
	o.ProfiledHostsPercentage = all.ProfiledHostsPercentage
	o.ProfiledHostsUsage = all.ProfiledHostsUsage
	o.SnmpPercentage = all.SnmpPercentage
	o.SnmpUsage = all.SnmpUsage
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
