// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricMetadataAttributes Metadata for one metric. Every attribute is always present. A value that is unknown or not set is `null`.
type MetricMetadataAttributes struct {
	// Description of the metric, or `null` if none is set.
	Description datadog.NullableString `json:"description"`
	// Whether historical metrics ingestion is enabled for the metric.
	HistoricalMetricsEnabled bool `json:"historical_metrics_enabled"`
	// Whether percentile aggregations are enabled for the metric. Percentile aggregations are available only for distribution metrics.
	IncludePercentiles bool `json:"include_percentiles"`
	// Name of the integration that sent the metric, or `null` if none is set. This is always `null` for
	// out-of-the-box RUM metrics.
	Integration datadog.NullableString `json:"integration"`
	// When the metric's tag or percentile configuration was last changed, or `null` if it has never been configured.
	LastConfiguredAt datadog.NullableTime `json:"last_configured_at"`
	// Metric type, or `null` if it isn't known. Known values are `gauge`, `count`, `rate`, `distribution`,
	// `explicit_histogram`, and `exponential_histogram`, plus the legacy values `counter` and `hll`. New values may
	// be added. For histograms and sketches, this can differ from the type returned by
	// `GET /api/v1/metrics/{metric_name}`.
	MetricType datadog.NullableString `json:"metric_type"`
	// Where the metric comes from. Each value is the display label shown on the Metrics Summary page. The set of labels can grow as Datadog adds products. Each field is `null` when it isn't known.
	Origin MetricMetadataOrigin `json:"origin"`
	// A unit of measure from the Datadog unit catalog, or `null` when no unit is set.
	PerUnit NullableMetricMetadataUnit `json:"per_unit"`
	// Short display name of the metric, or `null` if none is set.
	ShortName datadog.NullableString `json:"short_name"`
	// How the metric reaches Datadog, or `null` if it isn't known. Known values are `intake`, `api_default`, and
	// `intake_serverless` for submitted metrics, and `api_logs`, `api_processes`, `api_apm_spans`, and `api_rum`
	// for metrics generated from logs, processes, APM spans, and RUM events. New values may be added.
	Source datadog.NullableString `json:"source"`
	// StatsD flush interval of the metric in seconds, or `null` if none is set.
	StatsdInterval datadog.NullableInt64 `json:"statsd_interval"`
	// A unit of measure from the Datadog unit catalog, or `null` when no unit is set.
	Unit NullableMetricMetadataUnit `json:"unit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricMetadataAttributes instantiates a new MetricMetadataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricMetadataAttributes(description datadog.NullableString, historicalMetricsEnabled bool, includePercentiles bool, integration datadog.NullableString, lastConfiguredAt datadog.NullableTime, metricType datadog.NullableString, origin MetricMetadataOrigin, perUnit NullableMetricMetadataUnit, shortName datadog.NullableString, source datadog.NullableString, statsdInterval datadog.NullableInt64, unit NullableMetricMetadataUnit) *MetricMetadataAttributes {
	this := MetricMetadataAttributes{}
	this.Description = description
	this.HistoricalMetricsEnabled = historicalMetricsEnabled
	this.IncludePercentiles = includePercentiles
	this.Integration = integration
	this.LastConfiguredAt = lastConfiguredAt
	this.MetricType = metricType
	this.Origin = origin
	this.PerUnit = perUnit
	this.ShortName = shortName
	this.Source = source
	this.StatsdInterval = statsdInterval
	this.Unit = unit
	return &this
}

// NewMetricMetadataAttributesWithDefaults instantiates a new MetricMetadataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricMetadataAttributesWithDefaults() *MetricMetadataAttributes {
	this := MetricMetadataAttributes{}
	return &this
}

// GetDescription returns the Description field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MetricMetadataAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value.
func (o *MetricMetadataAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetHistoricalMetricsEnabled returns the HistoricalMetricsEnabled field value.
func (o *MetricMetadataAttributes) GetHistoricalMetricsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HistoricalMetricsEnabled
}

// GetHistoricalMetricsEnabledOk returns a tuple with the HistoricalMetricsEnabled field value
// and a boolean to check if the value has been set.
func (o *MetricMetadataAttributes) GetHistoricalMetricsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HistoricalMetricsEnabled, true
}

// SetHistoricalMetricsEnabled sets field value.
func (o *MetricMetadataAttributes) SetHistoricalMetricsEnabled(v bool) {
	o.HistoricalMetricsEnabled = v
}

// GetIncludePercentiles returns the IncludePercentiles field value.
func (o *MetricMetadataAttributes) GetIncludePercentiles() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IncludePercentiles
}

// GetIncludePercentilesOk returns a tuple with the IncludePercentiles field value
// and a boolean to check if the value has been set.
func (o *MetricMetadataAttributes) GetIncludePercentilesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludePercentiles, true
}

// SetIncludePercentiles sets field value.
func (o *MetricMetadataAttributes) SetIncludePercentiles(v bool) {
	o.IncludePercentiles = v
}

// GetIntegration returns the Integration field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MetricMetadataAttributes) GetIntegration() string {
	if o == nil || o.Integration.Get() == nil {
		var ret string
		return ret
	}
	return *o.Integration.Get()
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataAttributes) GetIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Integration.Get(), o.Integration.IsSet()
}

// SetIntegration sets field value.
func (o *MetricMetadataAttributes) SetIntegration(v string) {
	o.Integration.Set(&v)
}

// GetLastConfiguredAt returns the LastConfiguredAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *MetricMetadataAttributes) GetLastConfiguredAt() time.Time {
	if o == nil || o.LastConfiguredAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastConfiguredAt.Get()
}

// GetLastConfiguredAtOk returns a tuple with the LastConfiguredAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataAttributes) GetLastConfiguredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastConfiguredAt.Get(), o.LastConfiguredAt.IsSet()
}

// SetLastConfiguredAt sets field value.
func (o *MetricMetadataAttributes) SetLastConfiguredAt(v time.Time) {
	o.LastConfiguredAt.Set(&v)
}

// GetMetricType returns the MetricType field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MetricMetadataAttributes) GetMetricType() string {
	if o == nil || o.MetricType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MetricType.Get()
}

// GetMetricTypeOk returns a tuple with the MetricType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataAttributes) GetMetricTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricType.Get(), o.MetricType.IsSet()
}

// SetMetricType sets field value.
func (o *MetricMetadataAttributes) SetMetricType(v string) {
	o.MetricType.Set(&v)
}

// GetOrigin returns the Origin field value.
func (o *MetricMetadataAttributes) GetOrigin() MetricMetadataOrigin {
	if o == nil {
		var ret MetricMetadataOrigin
		return ret
	}
	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *MetricMetadataAttributes) GetOriginOk() (*MetricMetadataOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value.
func (o *MetricMetadataAttributes) SetOrigin(v MetricMetadataOrigin) {
	o.Origin = v
}

// GetPerUnit returns the PerUnit field value.
// If the value is explicit nil, the zero value for MetricMetadataUnit will be returned.
func (o *MetricMetadataAttributes) GetPerUnit() MetricMetadataUnit {
	if o == nil || o.PerUnit.Get() == nil {
		var ret MetricMetadataUnit
		return ret
	}
	return *o.PerUnit.Get()
}

// GetPerUnitOk returns a tuple with the PerUnit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataAttributes) GetPerUnitOk() (*MetricMetadataUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerUnit.Get(), o.PerUnit.IsSet()
}

// SetPerUnit sets field value.
func (o *MetricMetadataAttributes) SetPerUnit(v MetricMetadataUnit) {
	o.PerUnit.Set(&v)
}

// GetShortName returns the ShortName field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MetricMetadataAttributes) GetShortName() string {
	if o == nil || o.ShortName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShortName.Get()
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataAttributes) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortName.Get(), o.ShortName.IsSet()
}

// SetShortName sets field value.
func (o *MetricMetadataAttributes) SetShortName(v string) {
	o.ShortName.Set(&v)
}

// GetSource returns the Source field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MetricMetadataAttributes) GetSource() string {
	if o == nil || o.Source.Get() == nil {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataAttributes) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// SetSource sets field value.
func (o *MetricMetadataAttributes) SetSource(v string) {
	o.Source.Set(&v)
}

// GetStatsdInterval returns the StatsdInterval field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *MetricMetadataAttributes) GetStatsdInterval() int64 {
	if o == nil || o.StatsdInterval.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StatsdInterval.Get()
}

// GetStatsdIntervalOk returns a tuple with the StatsdInterval field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataAttributes) GetStatsdIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatsdInterval.Get(), o.StatsdInterval.IsSet()
}

// SetStatsdInterval sets field value.
func (o *MetricMetadataAttributes) SetStatsdInterval(v int64) {
	o.StatsdInterval.Set(&v)
}

// GetUnit returns the Unit field value.
// If the value is explicit nil, the zero value for MetricMetadataUnit will be returned.
func (o *MetricMetadataAttributes) GetUnit() MetricMetadataUnit {
	if o == nil || o.Unit.Get() == nil {
		var ret MetricMetadataUnit
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataAttributes) GetUnitOk() (*MetricMetadataUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// SetUnit sets field value.
func (o *MetricMetadataAttributes) SetUnit(v MetricMetadataUnit) {
	o.Unit.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricMetadataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description.Get()
	toSerialize["historical_metrics_enabled"] = o.HistoricalMetricsEnabled
	toSerialize["include_percentiles"] = o.IncludePercentiles
	toSerialize["integration"] = o.Integration.Get()
	toSerialize["last_configured_at"] = o.LastConfiguredAt.Get()
	toSerialize["metric_type"] = o.MetricType.Get()
	toSerialize["origin"] = o.Origin
	toSerialize["per_unit"] = o.PerUnit.Get()
	toSerialize["short_name"] = o.ShortName.Get()
	toSerialize["source"] = o.Source.Get()
	toSerialize["statsd_interval"] = o.StatsdInterval.Get()
	toSerialize["unit"] = o.Unit.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricMetadataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description              datadog.NullableString     `json:"description"`
		HistoricalMetricsEnabled *bool                      `json:"historical_metrics_enabled"`
		IncludePercentiles       *bool                      `json:"include_percentiles"`
		Integration              datadog.NullableString     `json:"integration"`
		LastConfiguredAt         datadog.NullableTime       `json:"last_configured_at"`
		MetricType               datadog.NullableString     `json:"metric_type"`
		Origin                   *MetricMetadataOrigin      `json:"origin"`
		PerUnit                  NullableMetricMetadataUnit `json:"per_unit"`
		ShortName                datadog.NullableString     `json:"short_name"`
		Source                   datadog.NullableString     `json:"source"`
		StatsdInterval           datadog.NullableInt64      `json:"statsd_interval"`
		Unit                     NullableMetricMetadataUnit `json:"unit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Description.IsSet() {
		return fmt.Errorf("required field description missing")
	}
	if all.HistoricalMetricsEnabled == nil {
		return fmt.Errorf("required field historical_metrics_enabled missing")
	}
	if all.IncludePercentiles == nil {
		return fmt.Errorf("required field include_percentiles missing")
	}
	if !all.Integration.IsSet() {
		return fmt.Errorf("required field integration missing")
	}
	if !all.LastConfiguredAt.IsSet() {
		return fmt.Errorf("required field last_configured_at missing")
	}
	if !all.MetricType.IsSet() {
		return fmt.Errorf("required field metric_type missing")
	}
	if all.Origin == nil {
		return fmt.Errorf("required field origin missing")
	}
	if !all.PerUnit.IsSet() {
		return fmt.Errorf("required field per_unit missing")
	}
	if !all.ShortName.IsSet() {
		return fmt.Errorf("required field short_name missing")
	}
	if !all.Source.IsSet() {
		return fmt.Errorf("required field source missing")
	}
	if !all.StatsdInterval.IsSet() {
		return fmt.Errorf("required field statsd_interval missing")
	}
	if !all.Unit.IsSet() {
		return fmt.Errorf("required field unit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "historical_metrics_enabled", "include_percentiles", "integration", "last_configured_at", "metric_type", "origin", "per_unit", "short_name", "source", "statsd_interval", "unit"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.HistoricalMetricsEnabled = *all.HistoricalMetricsEnabled
	o.IncludePercentiles = *all.IncludePercentiles
	o.Integration = all.Integration
	o.LastConfiguredAt = all.LastConfiguredAt
	o.MetricType = all.MetricType
	if all.Origin.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Origin = *all.Origin
	o.PerUnit = all.PerUnit
	o.ShortName = all.ShortName
	o.Source = all.Source
	o.StatsdInterval = all.StatsdInterval
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
