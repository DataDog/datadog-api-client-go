// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricMetadata Object with all metric related metadata.
type MetricMetadata struct {
	// Metric description.
	Description datadog.NullableString `json:"description,omitempty"`
	// Name of the integration that sent the metric if applicable.
	Integration datadog.NullableString `json:"integration,omitempty"`
	// Per unit of the metric such as `second` in `bytes per second`.
	PerUnit datadog.NullableString `json:"per_unit,omitempty"`
	// A more human-readable and abbreviated version of the metric name.
	ShortName datadog.NullableString `json:"short_name,omitempty"`
	// StatsD flush interval of the metric in seconds if applicable.
	StatsdInterval datadog.NullableInt64 `json:"statsd_interval,omitempty"`
	// Metric type such as `gauge` or `rate`.
	Type datadog.NullableString `json:"type,omitempty"`
	// Primary unit of the metric such as `byte` or `operation`.
	Unit datadog.NullableString `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricMetadata instantiates a new MetricMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricMetadata() *MetricMetadata {
	this := MetricMetadata{}
	return &this
}

// NewMetricMetadataWithDefaults instantiates a new MetricMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricMetadataWithDefaults() *MetricMetadata {
	this := MetricMetadata{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricMetadata) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricMetadata) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *MetricMetadata) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *MetricMetadata) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *MetricMetadata) UnsetDescription() {
	o.Description.Unset()
}

// GetIntegration returns the Integration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricMetadata) GetIntegration() string {
	if o == nil || o.Integration.Get() == nil {
		var ret string
		return ret
	}
	return *o.Integration.Get()
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadata) GetIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Integration.Get(), o.Integration.IsSet()
}

// HasIntegration returns a boolean if a field has been set.
func (o *MetricMetadata) HasIntegration() bool {
	return o != nil && o.Integration.IsSet()
}

// SetIntegration gets a reference to the given datadog.NullableString and assigns it to the Integration field.
func (o *MetricMetadata) SetIntegration(v string) {
	o.Integration.Set(&v)
}

// SetIntegrationNil sets the value for Integration to be an explicit nil.
func (o *MetricMetadata) SetIntegrationNil() {
	o.Integration.Set(nil)
}

// UnsetIntegration ensures that no value is present for Integration, not even an explicit nil.
func (o *MetricMetadata) UnsetIntegration() {
	o.Integration.Unset()
}

// GetPerUnit returns the PerUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricMetadata) GetPerUnit() string {
	if o == nil || o.PerUnit.Get() == nil {
		var ret string
		return ret
	}
	return *o.PerUnit.Get()
}

// GetPerUnitOk returns a tuple with the PerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadata) GetPerUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerUnit.Get(), o.PerUnit.IsSet()
}

// HasPerUnit returns a boolean if a field has been set.
func (o *MetricMetadata) HasPerUnit() bool {
	return o != nil && o.PerUnit.IsSet()
}

// SetPerUnit gets a reference to the given datadog.NullableString and assigns it to the PerUnit field.
func (o *MetricMetadata) SetPerUnit(v string) {
	o.PerUnit.Set(&v)
}

// SetPerUnitNil sets the value for PerUnit to be an explicit nil.
func (o *MetricMetadata) SetPerUnitNil() {
	o.PerUnit.Set(nil)
}

// UnsetPerUnit ensures that no value is present for PerUnit, not even an explicit nil.
func (o *MetricMetadata) UnsetPerUnit() {
	o.PerUnit.Unset()
}

// GetShortName returns the ShortName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricMetadata) GetShortName() string {
	if o == nil || o.ShortName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShortName.Get()
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadata) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortName.Get(), o.ShortName.IsSet()
}

// HasShortName returns a boolean if a field has been set.
func (o *MetricMetadata) HasShortName() bool {
	return o != nil && o.ShortName.IsSet()
}

// SetShortName gets a reference to the given datadog.NullableString and assigns it to the ShortName field.
func (o *MetricMetadata) SetShortName(v string) {
	o.ShortName.Set(&v)
}

// SetShortNameNil sets the value for ShortName to be an explicit nil.
func (o *MetricMetadata) SetShortNameNil() {
	o.ShortName.Set(nil)
}

// UnsetShortName ensures that no value is present for ShortName, not even an explicit nil.
func (o *MetricMetadata) UnsetShortName() {
	o.ShortName.Unset()
}

// GetStatsdInterval returns the StatsdInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricMetadata) GetStatsdInterval() int64 {
	if o == nil || o.StatsdInterval.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StatsdInterval.Get()
}

// GetStatsdIntervalOk returns a tuple with the StatsdInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadata) GetStatsdIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatsdInterval.Get(), o.StatsdInterval.IsSet()
}

// HasStatsdInterval returns a boolean if a field has been set.
func (o *MetricMetadata) HasStatsdInterval() bool {
	return o != nil && o.StatsdInterval.IsSet()
}

// SetStatsdInterval gets a reference to the given datadog.NullableInt64 and assigns it to the StatsdInterval field.
func (o *MetricMetadata) SetStatsdInterval(v int64) {
	o.StatsdInterval.Set(&v)
}

// SetStatsdIntervalNil sets the value for StatsdInterval to be an explicit nil.
func (o *MetricMetadata) SetStatsdIntervalNil() {
	o.StatsdInterval.Set(nil)
}

// UnsetStatsdInterval ensures that no value is present for StatsdInterval, not even an explicit nil.
func (o *MetricMetadata) UnsetStatsdInterval() {
	o.StatsdInterval.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricMetadata) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadata) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *MetricMetadata) HasType() bool {
	return o != nil && o.Type.IsSet()
}

// SetType gets a reference to the given datadog.NullableString and assigns it to the Type field.
func (o *MetricMetadata) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil.
func (o *MetricMetadata) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil.
func (o *MetricMetadata) UnsetType() {
	o.Type.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricMetadata) GetUnit() string {
	if o == nil || o.Unit.Get() == nil {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadata) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricMetadata) HasUnit() bool {
	return o != nil && o.Unit.IsSet()
}

// SetUnit gets a reference to the given datadog.NullableString and assigns it to the Unit field.
func (o *MetricMetadata) SetUnit(v string) {
	o.Unit.Set(&v)
}

// SetUnitNil sets the value for Unit to be an explicit nil.
func (o *MetricMetadata) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil.
func (o *MetricMetadata) UnsetUnit() {
	o.Unit.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Integration.IsSet() {
		toSerialize["integration"] = o.Integration.Get()
	}
	if o.PerUnit.IsSet() {
		toSerialize["per_unit"] = o.PerUnit.Get()
	}
	if o.ShortName.IsSet() {
		toSerialize["short_name"] = o.ShortName.Get()
	}
	if o.StatsdInterval.IsSet() {
		toSerialize["statsd_interval"] = o.StatsdInterval.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description    datadog.NullableString `json:"description,omitempty"`
		Integration    datadog.NullableString `json:"integration,omitempty"`
		PerUnit        datadog.NullableString `json:"per_unit,omitempty"`
		ShortName      datadog.NullableString `json:"short_name,omitempty"`
		StatsdInterval datadog.NullableInt64  `json:"statsd_interval,omitempty"`
		Type           datadog.NullableString `json:"type,omitempty"`
		Unit           datadog.NullableString `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "integration", "per_unit", "short_name", "statsd_interval", "type", "unit"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Integration = all.Integration
	o.PerUnit = all.PerUnit
	o.ShortName = all.ShortName
	o.StatsdInterval = all.StatsdInterval
	o.Type = all.Type
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
