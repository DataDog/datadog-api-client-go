// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorOptions Options of a Product Analytics monitor.
type ProductAnalyticsMonitorOptions struct {
	// UUID of the chart the monitor was created from.
	ChartUuid *string `json:"chart_uuid,omitempty"`
	// Metric of the chart that the monitor evaluates.
	MetricSelector *string `json:"metric_selector,omitempty"`
	// Index of the chart query that the monitor evaluates.
	QueryIndex *int64 `json:"query_index,omitempty"`
	// Number of minutes after the last notification before the monitor notifies again on the current status.
	// `0` means it never notifies again.
	RenotifyInterval *int64 `json:"renotify_interval,omitempty"`
	// Alert thresholds of a Product Analytics monitor.
	// Keep `critical` in sync with the comparison in the monitor query; the query is not derived from these values.
	Thresholds *ProductAnalyticsMonitorThresholds `json:"thresholds,omitempty"`
	// Time frame of the chart that the monitor evaluates.
	Timeframe *string `json:"timeframe,omitempty"`
	// Query variables referenced by the monitor query.
	Variables []map[string]interface{} `json:"variables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsMonitorOptions instantiates a new ProductAnalyticsMonitorOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsMonitorOptions() *ProductAnalyticsMonitorOptions {
	this := ProductAnalyticsMonitorOptions{}
	return &this
}

// NewProductAnalyticsMonitorOptionsWithDefaults instantiates a new ProductAnalyticsMonitorOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsMonitorOptionsWithDefaults() *ProductAnalyticsMonitorOptions {
	this := ProductAnalyticsMonitorOptions{}
	return &this
}

// GetChartUuid returns the ChartUuid field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorOptions) GetChartUuid() string {
	if o == nil || o.ChartUuid == nil {
		var ret string
		return ret
	}
	return *o.ChartUuid
}

// GetChartUuidOk returns a tuple with the ChartUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorOptions) GetChartUuidOk() (*string, bool) {
	if o == nil || o.ChartUuid == nil {
		return nil, false
	}
	return o.ChartUuid, true
}

// HasChartUuid returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorOptions) HasChartUuid() bool {
	return o != nil && o.ChartUuid != nil
}

// SetChartUuid gets a reference to the given string and assigns it to the ChartUuid field.
func (o *ProductAnalyticsMonitorOptions) SetChartUuid(v string) {
	o.ChartUuid = &v
}

// GetMetricSelector returns the MetricSelector field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorOptions) GetMetricSelector() string {
	if o == nil || o.MetricSelector == nil {
		var ret string
		return ret
	}
	return *o.MetricSelector
}

// GetMetricSelectorOk returns a tuple with the MetricSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorOptions) GetMetricSelectorOk() (*string, bool) {
	if o == nil || o.MetricSelector == nil {
		return nil, false
	}
	return o.MetricSelector, true
}

// HasMetricSelector returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorOptions) HasMetricSelector() bool {
	return o != nil && o.MetricSelector != nil
}

// SetMetricSelector gets a reference to the given string and assigns it to the MetricSelector field.
func (o *ProductAnalyticsMonitorOptions) SetMetricSelector(v string) {
	o.MetricSelector = &v
}

// GetQueryIndex returns the QueryIndex field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorOptions) GetQueryIndex() int64 {
	if o == nil || o.QueryIndex == nil {
		var ret int64
		return ret
	}
	return *o.QueryIndex
}

// GetQueryIndexOk returns a tuple with the QueryIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorOptions) GetQueryIndexOk() (*int64, bool) {
	if o == nil || o.QueryIndex == nil {
		return nil, false
	}
	return o.QueryIndex, true
}

// HasQueryIndex returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorOptions) HasQueryIndex() bool {
	return o != nil && o.QueryIndex != nil
}

// SetQueryIndex gets a reference to the given int64 and assigns it to the QueryIndex field.
func (o *ProductAnalyticsMonitorOptions) SetQueryIndex(v int64) {
	o.QueryIndex = &v
}

// GetRenotifyInterval returns the RenotifyInterval field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorOptions) GetRenotifyInterval() int64 {
	if o == nil || o.RenotifyInterval == nil {
		var ret int64
		return ret
	}
	return *o.RenotifyInterval
}

// GetRenotifyIntervalOk returns a tuple with the RenotifyInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorOptions) GetRenotifyIntervalOk() (*int64, bool) {
	if o == nil || o.RenotifyInterval == nil {
		return nil, false
	}
	return o.RenotifyInterval, true
}

// HasRenotifyInterval returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorOptions) HasRenotifyInterval() bool {
	return o != nil && o.RenotifyInterval != nil
}

// SetRenotifyInterval gets a reference to the given int64 and assigns it to the RenotifyInterval field.
func (o *ProductAnalyticsMonitorOptions) SetRenotifyInterval(v int64) {
	o.RenotifyInterval = &v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorOptions) GetThresholds() ProductAnalyticsMonitorThresholds {
	if o == nil || o.Thresholds == nil {
		var ret ProductAnalyticsMonitorThresholds
		return ret
	}
	return *o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorOptions) GetThresholdsOk() (*ProductAnalyticsMonitorThresholds, bool) {
	if o == nil || o.Thresholds == nil {
		return nil, false
	}
	return o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorOptions) HasThresholds() bool {
	return o != nil && o.Thresholds != nil
}

// SetThresholds gets a reference to the given ProductAnalyticsMonitorThresholds and assigns it to the Thresholds field.
func (o *ProductAnalyticsMonitorOptions) SetThresholds(v ProductAnalyticsMonitorThresholds) {
	o.Thresholds = &v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorOptions) GetTimeframe() string {
	if o == nil || o.Timeframe == nil {
		var ret string
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorOptions) GetTimeframeOk() (*string, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorOptions) HasTimeframe() bool {
	return o != nil && o.Timeframe != nil
}

// SetTimeframe gets a reference to the given string and assigns it to the Timeframe field.
func (o *ProductAnalyticsMonitorOptions) SetTimeframe(v string) {
	o.Timeframe = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorOptions) GetVariables() []map[string]interface{} {
	if o == nil || o.Variables == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorOptions) GetVariablesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorOptions) HasVariables() bool {
	return o != nil && o.Variables != nil
}

// SetVariables gets a reference to the given []map[string]interface{} and assigns it to the Variables field.
func (o *ProductAnalyticsMonitorOptions) SetVariables(v []map[string]interface{}) {
	o.Variables = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsMonitorOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChartUuid != nil {
		toSerialize["chart_uuid"] = o.ChartUuid
	}
	if o.MetricSelector != nil {
		toSerialize["metric_selector"] = o.MetricSelector
	}
	if o.QueryIndex != nil {
		toSerialize["query_index"] = o.QueryIndex
	}
	if o.RenotifyInterval != nil {
		toSerialize["renotify_interval"] = o.RenotifyInterval
	}
	if o.Thresholds != nil {
		toSerialize["thresholds"] = o.Thresholds
	}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsMonitorOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChartUuid        *string                            `json:"chart_uuid,omitempty"`
		MetricSelector   *string                            `json:"metric_selector,omitempty"`
		QueryIndex       *int64                             `json:"query_index,omitempty"`
		RenotifyInterval *int64                             `json:"renotify_interval,omitempty"`
		Thresholds       *ProductAnalyticsMonitorThresholds `json:"thresholds,omitempty"`
		Timeframe        *string                            `json:"timeframe,omitempty"`
		Variables        []map[string]interface{}           `json:"variables,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chart_uuid", "metric_selector", "query_index", "renotify_interval", "thresholds", "timeframe", "variables"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChartUuid = all.ChartUuid
	o.MetricSelector = all.MetricSelector
	o.QueryIndex = all.QueryIndex
	o.RenotifyInterval = all.RenotifyInterval
	if all.Thresholds != nil && all.Thresholds.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Thresholds = all.Thresholds
	o.Timeframe = all.Timeframe
	o.Variables = all.Variables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
