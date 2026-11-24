// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionFilterCreateAttributes The object describing attributes of a RUM retention filter to create.
type RUMRetentionFilterCreateAttributes struct {
	// The configuration for cross-product retention filters.
	CrossProductSampling *RUMCrossProductSamplingCreate `json:"cross_product_sampling,omitempty"`
	// Whether the retention filter is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The type of RUM events to filter on.
	EventType RUMRetentionFilterEventType `json:"event_type"`
	// The name of a RUM retention filter.
	Name string `json:"name"`
	// The query string for a RUM retention filter.
	Query *string `json:"query,omitempty"`
	// The sample rate for a RUM retention filter, between 0.1 and 100. Values are truncated to one decimal place.
	SampleRate float64 `json:"sample_rate"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMRetentionFilterCreateAttributes instantiates a new RUMRetentionFilterCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMRetentionFilterCreateAttributes(eventType RUMRetentionFilterEventType, name string, sampleRate float64) *RUMRetentionFilterCreateAttributes {
	this := RUMRetentionFilterCreateAttributes{}
	this.EventType = eventType
	this.Name = name
	this.SampleRate = sampleRate
	return &this
}

// NewRUMRetentionFilterCreateAttributesWithDefaults instantiates a new RUMRetentionFilterCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMRetentionFilterCreateAttributesWithDefaults() *RUMRetentionFilterCreateAttributes {
	this := RUMRetentionFilterCreateAttributes{}
	return &this
}

// GetCrossProductSampling returns the CrossProductSampling field value if set, zero value otherwise.
func (o *RUMRetentionFilterCreateAttributes) GetCrossProductSampling() RUMCrossProductSamplingCreate {
	if o == nil || o.CrossProductSampling == nil {
		var ret RUMCrossProductSamplingCreate
		return ret
	}
	return *o.CrossProductSampling
}

// GetCrossProductSamplingOk returns a tuple with the CrossProductSampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterCreateAttributes) GetCrossProductSamplingOk() (*RUMCrossProductSamplingCreate, bool) {
	if o == nil || o.CrossProductSampling == nil {
		return nil, false
	}
	return o.CrossProductSampling, true
}

// HasCrossProductSampling returns a boolean if a field has been set.
func (o *RUMRetentionFilterCreateAttributes) HasCrossProductSampling() bool {
	return o != nil && o.CrossProductSampling != nil
}

// SetCrossProductSampling gets a reference to the given RUMCrossProductSamplingCreate and assigns it to the CrossProductSampling field.
func (o *RUMRetentionFilterCreateAttributes) SetCrossProductSampling(v RUMCrossProductSamplingCreate) {
	o.CrossProductSampling = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RUMRetentionFilterCreateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RUMRetentionFilterCreateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RUMRetentionFilterCreateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventType returns the EventType field value.
func (o *RUMRetentionFilterCreateAttributes) GetEventType() RUMRetentionFilterEventType {
	if o == nil {
		var ret RUMRetentionFilterEventType
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterCreateAttributes) GetEventTypeOk() (*RUMRetentionFilterEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *RUMRetentionFilterCreateAttributes) SetEventType(v RUMRetentionFilterEventType) {
	o.EventType = v
}

// GetName returns the Name field value.
func (o *RUMRetentionFilterCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RUMRetentionFilterCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *RUMRetentionFilterCreateAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterCreateAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *RUMRetentionFilterCreateAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *RUMRetentionFilterCreateAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetSampleRate returns the SampleRate field value.
func (o *RUMRetentionFilterCreateAttributes) GetSampleRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value
// and a boolean to check if the value has been set.
func (o *RUMRetentionFilterCreateAttributes) GetSampleRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleRate, true
}

// SetSampleRate sets field value.
func (o *RUMRetentionFilterCreateAttributes) SetSampleRate(v float64) {
	o.SampleRate = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMRetentionFilterCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossProductSampling != nil {
		toSerialize["cross_product_sampling"] = o.CrossProductSampling
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["event_type"] = o.EventType
	toSerialize["name"] = o.Name
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	toSerialize["sample_rate"] = o.SampleRate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMRetentionFilterCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossProductSampling *RUMCrossProductSamplingCreate `json:"cross_product_sampling,omitempty"`
		Enabled              *bool                          `json:"enabled,omitempty"`
		EventType            *RUMRetentionFilterEventType   `json:"event_type"`
		Name                 *string                        `json:"name"`
		Query                *string                        `json:"query,omitempty"`
		SampleRate           *float64                       `json:"sample_rate"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EventType == nil {
		return fmt.Errorf("required field event_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SampleRate == nil {
		return fmt.Errorf("required field sample_rate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cross_product_sampling", "enabled", "event_type", "name", "query", "sample_rate"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CrossProductSampling != nil && all.CrossProductSampling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CrossProductSampling = all.CrossProductSampling
	o.Enabled = all.Enabled
	if !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = *all.EventType
	}
	o.Name = *all.Name
	o.Query = all.Query
	o.SampleRate = *all.SampleRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
