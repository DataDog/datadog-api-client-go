// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumHardcodedRetentionFilterAttributes The attributes of a hardcoded retention filter.
type RumHardcodedRetentionFilterAttributes struct {
	// Cross-product retention settings for a hardcoded retention filter.
	CrossProductSampling *RumHardcodedCrossProductSampling `json:"cross_product_sampling,omitempty"`
	// Flags indicating which `cross_product_sampling` fields can be updated. Read-only.
	CrossProductSamplingEditability *RumHardcodedCrossProductSamplingEditability `json:"cross_product_sampling_editability,omitempty"`
	// Indicates whether the hardcoded retention filter is active. Read-only.
	Enabled *bool `json:"enabled,omitempty"`
	// The type of RUM events the hardcoded filter applies to. Read-only.
	EventType *RumHardcodedRetentionFilterEventType `json:"event_type,omitempty"`
	// The name of the hardcoded retention filter. Read-only.
	Name *string `json:"name,omitempty"`
	// The query string for the hardcoded retention filter. Read-only.
	Query *string `json:"query,omitempty"`
	// The retention sample rate (0–100) for the hardcoded filter. Read-only.
	SampleRate *float64 `json:"sample_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumHardcodedRetentionFilterAttributes instantiates a new RumHardcodedRetentionFilterAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumHardcodedRetentionFilterAttributes() *RumHardcodedRetentionFilterAttributes {
	this := RumHardcodedRetentionFilterAttributes{}
	return &this
}

// NewRumHardcodedRetentionFilterAttributesWithDefaults instantiates a new RumHardcodedRetentionFilterAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumHardcodedRetentionFilterAttributesWithDefaults() *RumHardcodedRetentionFilterAttributes {
	this := RumHardcodedRetentionFilterAttributes{}
	return &this
}

// GetCrossProductSampling returns the CrossProductSampling field value if set, zero value otherwise.
func (o *RumHardcodedRetentionFilterAttributes) GetCrossProductSampling() RumHardcodedCrossProductSampling {
	if o == nil || o.CrossProductSampling == nil {
		var ret RumHardcodedCrossProductSampling
		return ret
	}
	return *o.CrossProductSampling
}

// GetCrossProductSamplingOk returns a tuple with the CrossProductSampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedRetentionFilterAttributes) GetCrossProductSamplingOk() (*RumHardcodedCrossProductSampling, bool) {
	if o == nil || o.CrossProductSampling == nil {
		return nil, false
	}
	return o.CrossProductSampling, true
}

// HasCrossProductSampling returns a boolean if a field has been set.
func (o *RumHardcodedRetentionFilterAttributes) HasCrossProductSampling() bool {
	return o != nil && o.CrossProductSampling != nil
}

// SetCrossProductSampling gets a reference to the given RumHardcodedCrossProductSampling and assigns it to the CrossProductSampling field.
func (o *RumHardcodedRetentionFilterAttributes) SetCrossProductSampling(v RumHardcodedCrossProductSampling) {
	o.CrossProductSampling = &v
}

// GetCrossProductSamplingEditability returns the CrossProductSamplingEditability field value if set, zero value otherwise.
func (o *RumHardcodedRetentionFilterAttributes) GetCrossProductSamplingEditability() RumHardcodedCrossProductSamplingEditability {
	if o == nil || o.CrossProductSamplingEditability == nil {
		var ret RumHardcodedCrossProductSamplingEditability
		return ret
	}
	return *o.CrossProductSamplingEditability
}

// GetCrossProductSamplingEditabilityOk returns a tuple with the CrossProductSamplingEditability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedRetentionFilterAttributes) GetCrossProductSamplingEditabilityOk() (*RumHardcodedCrossProductSamplingEditability, bool) {
	if o == nil || o.CrossProductSamplingEditability == nil {
		return nil, false
	}
	return o.CrossProductSamplingEditability, true
}

// HasCrossProductSamplingEditability returns a boolean if a field has been set.
func (o *RumHardcodedRetentionFilterAttributes) HasCrossProductSamplingEditability() bool {
	return o != nil && o.CrossProductSamplingEditability != nil
}

// SetCrossProductSamplingEditability gets a reference to the given RumHardcodedCrossProductSamplingEditability and assigns it to the CrossProductSamplingEditability field.
func (o *RumHardcodedRetentionFilterAttributes) SetCrossProductSamplingEditability(v RumHardcodedCrossProductSamplingEditability) {
	o.CrossProductSamplingEditability = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RumHardcodedRetentionFilterAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedRetentionFilterAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RumHardcodedRetentionFilterAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RumHardcodedRetentionFilterAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RumHardcodedRetentionFilterAttributes) GetEventType() RumHardcodedRetentionFilterEventType {
	if o == nil || o.EventType == nil {
		var ret RumHardcodedRetentionFilterEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedRetentionFilterAttributes) GetEventTypeOk() (*RumHardcodedRetentionFilterEventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RumHardcodedRetentionFilterAttributes) HasEventType() bool {
	return o != nil && o.EventType != nil
}

// SetEventType gets a reference to the given RumHardcodedRetentionFilterEventType and assigns it to the EventType field.
func (o *RumHardcodedRetentionFilterAttributes) SetEventType(v RumHardcodedRetentionFilterEventType) {
	o.EventType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RumHardcodedRetentionFilterAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedRetentionFilterAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RumHardcodedRetentionFilterAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RumHardcodedRetentionFilterAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *RumHardcodedRetentionFilterAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedRetentionFilterAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *RumHardcodedRetentionFilterAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *RumHardcodedRetentionFilterAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetSampleRate returns the SampleRate field value if set, zero value otherwise.
func (o *RumHardcodedRetentionFilterAttributes) GetSampleRate() float64 {
	if o == nil || o.SampleRate == nil {
		var ret float64
		return ret
	}
	return *o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedRetentionFilterAttributes) GetSampleRateOk() (*float64, bool) {
	if o == nil || o.SampleRate == nil {
		return nil, false
	}
	return o.SampleRate, true
}

// HasSampleRate returns a boolean if a field has been set.
func (o *RumHardcodedRetentionFilterAttributes) HasSampleRate() bool {
	return o != nil && o.SampleRate != nil
}

// SetSampleRate gets a reference to the given float64 and assigns it to the SampleRate field.
func (o *RumHardcodedRetentionFilterAttributes) SetSampleRate(v float64) {
	o.SampleRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumHardcodedRetentionFilterAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossProductSampling != nil {
		toSerialize["cross_product_sampling"] = o.CrossProductSampling
	}
	if o.CrossProductSamplingEditability != nil {
		toSerialize["cross_product_sampling_editability"] = o.CrossProductSamplingEditability
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.EventType != nil {
		toSerialize["event_type"] = o.EventType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.SampleRate != nil {
		toSerialize["sample_rate"] = o.SampleRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumHardcodedRetentionFilterAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossProductSampling            *RumHardcodedCrossProductSampling            `json:"cross_product_sampling,omitempty"`
		CrossProductSamplingEditability *RumHardcodedCrossProductSamplingEditability `json:"cross_product_sampling_editability,omitempty"`
		Enabled                         *bool                                        `json:"enabled,omitempty"`
		EventType                       *RumHardcodedRetentionFilterEventType        `json:"event_type,omitempty"`
		Name                            *string                                      `json:"name,omitempty"`
		Query                           *string                                      `json:"query,omitempty"`
		SampleRate                      *float64                                     `json:"sample_rate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cross_product_sampling", "cross_product_sampling_editability", "enabled", "event_type", "name", "query", "sample_rate"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CrossProductSampling != nil && all.CrossProductSampling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CrossProductSampling = all.CrossProductSampling
	if all.CrossProductSamplingEditability != nil && all.CrossProductSamplingEditability.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CrossProductSamplingEditability = all.CrossProductSamplingEditability
	o.Enabled = all.Enabled
	if all.EventType != nil && !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = all.EventType
	}
	o.Name = all.Name
	o.Query = all.Query
	o.SampleRate = all.SampleRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
