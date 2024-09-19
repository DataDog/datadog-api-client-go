// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// MetricMetadata Object with all metric related metadata.
type MetricMetadata struct {
	// Metric description.
	Description *string `json:"description,omitempty"`
	// Name of the integration that sent the metric if applicable.
	Integration *string `json:"integration,omitempty"`
	// Per unit of the metric such as `second` in `bytes per second`.
	PerUnit *string `json:"per_unit,omitempty"`
	// A more human-readable and abbreviated version of the metric name.
	ShortName *string `json:"short_name,omitempty"`
	// StatsD flush interval of the metric in seconds if applicable.
	StatsdInterval *int64 `json:"statsd_interval,omitempty"`
	// Metric type such as `gauge` or `rate`.
	Type *string `json:"type,omitempty"`
	// Primary unit of the metric such as `byte` or `operation`.
	Unit *string `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
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
// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetricMetadata) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricMetadata) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetricMetadata) SetDescription(v string) {
	o.Description = &v
}


// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *MetricMetadata) GetIntegration() string {
	if o == nil || o.Integration == nil {
		var ret string
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadata) GetIntegrationOk() (*string, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *MetricMetadata) HasIntegration() bool {
	return o != nil && o.Integration != nil
}

// SetIntegration gets a reference to the given string and assigns it to the Integration field.
func (o *MetricMetadata) SetIntegration(v string) {
	o.Integration = &v
}


// GetPerUnit returns the PerUnit field value if set, zero value otherwise.
func (o *MetricMetadata) GetPerUnit() string {
	if o == nil || o.PerUnit == nil {
		var ret string
		return ret
	}
	return *o.PerUnit
}

// GetPerUnitOk returns a tuple with the PerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadata) GetPerUnitOk() (*string, bool) {
	if o == nil || o.PerUnit == nil {
		return nil, false
	}
	return o.PerUnit, true
}

// HasPerUnit returns a boolean if a field has been set.
func (o *MetricMetadata) HasPerUnit() bool {
	return o != nil && o.PerUnit != nil
}

// SetPerUnit gets a reference to the given string and assigns it to the PerUnit field.
func (o *MetricMetadata) SetPerUnit(v string) {
	o.PerUnit = &v
}


// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *MetricMetadata) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadata) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *MetricMetadata) HasShortName() bool {
	return o != nil && o.ShortName != nil
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *MetricMetadata) SetShortName(v string) {
	o.ShortName = &v
}


// GetStatsdInterval returns the StatsdInterval field value if set, zero value otherwise.
func (o *MetricMetadata) GetStatsdInterval() int64 {
	if o == nil || o.StatsdInterval == nil {
		var ret int64
		return ret
	}
	return *o.StatsdInterval
}

// GetStatsdIntervalOk returns a tuple with the StatsdInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadata) GetStatsdIntervalOk() (*int64, bool) {
	if o == nil || o.StatsdInterval == nil {
		return nil, false
	}
	return o.StatsdInterval, true
}

// HasStatsdInterval returns a boolean if a field has been set.
func (o *MetricMetadata) HasStatsdInterval() bool {
	return o != nil && o.StatsdInterval != nil
}

// SetStatsdInterval gets a reference to the given int64 and assigns it to the StatsdInterval field.
func (o *MetricMetadata) SetStatsdInterval(v int64) {
	o.StatsdInterval = &v
}


// GetType returns the Type field value if set, zero value otherwise.
func (o *MetricMetadata) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadata) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetricMetadata) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MetricMetadata) SetType(v string) {
	o.Type = &v
}


// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MetricMetadata) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadata) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricMetadata) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *MetricMetadata) SetUnit(v string) {
	o.Unit = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o MetricMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	if o.PerUnit != nil {
		toSerialize["per_unit"] = o.PerUnit
	}
	if o.ShortName != nil {
		toSerialize["short_name"] = o.ShortName
	}
	if o.StatsdInterval != nil {
		toSerialize["statsd_interval"] = o.StatsdInterval
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string `json:"description,omitempty"`
		Integration *string `json:"integration,omitempty"`
		PerUnit *string `json:"per_unit,omitempty"`
		ShortName *string `json:"short_name,omitempty"`
		StatsdInterval *int64 `json:"statsd_interval,omitempty"`
		Type *string `json:"type,omitempty"`
		Unit *string `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "description", "integration", "per_unit", "short_name", "statsd_interval", "type", "unit",  })
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
