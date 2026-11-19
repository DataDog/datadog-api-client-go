// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchExperimentStep Experiment linked to the launch. The draft fields are cleared when `experiment_id` is set.
type ProductAnalyticsLaunchExperimentStep struct {
	// UUID of the linked experiment.
	ExperimentId datadog.NullableString `json:"experiment_id,omitempty"`
	// Hypothesis of a draft experiment.
	Hypothesis datadog.NullableString `json:"hypothesis,omitempty"`
	// Name of a draft experiment.
	Name datadog.NullableString `json:"name,omitempty"`
	// Protocol ID of a draft experiment.
	ProtocolId datadog.NullableString `json:"protocol_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchExperimentStep instantiates a new ProductAnalyticsLaunchExperimentStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchExperimentStep() *ProductAnalyticsLaunchExperimentStep {
	this := ProductAnalyticsLaunchExperimentStep{}
	return &this
}

// NewProductAnalyticsLaunchExperimentStepWithDefaults instantiates a new ProductAnalyticsLaunchExperimentStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchExperimentStepWithDefaults() *ProductAnalyticsLaunchExperimentStep {
	this := ProductAnalyticsLaunchExperimentStep{}
	return &this
}

// GetExperimentId returns the ExperimentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchExperimentStep) GetExperimentId() string {
	if o == nil || o.ExperimentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExperimentId.Get()
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchExperimentStep) GetExperimentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperimentId.Get(), o.ExperimentId.IsSet()
}

// HasExperimentId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchExperimentStep) HasExperimentId() bool {
	return o != nil && o.ExperimentId.IsSet()
}

// SetExperimentId gets a reference to the given datadog.NullableString and assigns it to the ExperimentId field.
func (o *ProductAnalyticsLaunchExperimentStep) SetExperimentId(v string) {
	o.ExperimentId.Set(&v)
}

// SetExperimentIdNil sets the value for ExperimentId to be an explicit nil.
func (o *ProductAnalyticsLaunchExperimentStep) SetExperimentIdNil() {
	o.ExperimentId.Set(nil)
}

// UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil.
func (o *ProductAnalyticsLaunchExperimentStep) UnsetExperimentId() {
	o.ExperimentId.Unset()
}

// GetHypothesis returns the Hypothesis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchExperimentStep) GetHypothesis() string {
	if o == nil || o.Hypothesis.Get() == nil {
		var ret string
		return ret
	}
	return *o.Hypothesis.Get()
}

// GetHypothesisOk returns a tuple with the Hypothesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchExperimentStep) GetHypothesisOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hypothesis.Get(), o.Hypothesis.IsSet()
}

// HasHypothesis returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchExperimentStep) HasHypothesis() bool {
	return o != nil && o.Hypothesis.IsSet()
}

// SetHypothesis gets a reference to the given datadog.NullableString and assigns it to the Hypothesis field.
func (o *ProductAnalyticsLaunchExperimentStep) SetHypothesis(v string) {
	o.Hypothesis.Set(&v)
}

// SetHypothesisNil sets the value for Hypothesis to be an explicit nil.
func (o *ProductAnalyticsLaunchExperimentStep) SetHypothesisNil() {
	o.Hypothesis.Set(nil)
}

// UnsetHypothesis ensures that no value is present for Hypothesis, not even an explicit nil.
func (o *ProductAnalyticsLaunchExperimentStep) UnsetHypothesis() {
	o.Hypothesis.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchExperimentStep) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchExperimentStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchExperimentStep) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *ProductAnalyticsLaunchExperimentStep) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *ProductAnalyticsLaunchExperimentStep) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *ProductAnalyticsLaunchExperimentStep) UnsetName() {
	o.Name.Unset()
}

// GetProtocolId returns the ProtocolId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchExperimentStep) GetProtocolId() string {
	if o == nil || o.ProtocolId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtocolId.Get()
}

// GetProtocolIdOk returns a tuple with the ProtocolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchExperimentStep) GetProtocolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtocolId.Get(), o.ProtocolId.IsSet()
}

// HasProtocolId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchExperimentStep) HasProtocolId() bool {
	return o != nil && o.ProtocolId.IsSet()
}

// SetProtocolId gets a reference to the given datadog.NullableString and assigns it to the ProtocolId field.
func (o *ProductAnalyticsLaunchExperimentStep) SetProtocolId(v string) {
	o.ProtocolId.Set(&v)
}

// SetProtocolIdNil sets the value for ProtocolId to be an explicit nil.
func (o *ProductAnalyticsLaunchExperimentStep) SetProtocolIdNil() {
	o.ProtocolId.Set(nil)
}

// UnsetProtocolId ensures that no value is present for ProtocolId, not even an explicit nil.
func (o *ProductAnalyticsLaunchExperimentStep) UnsetProtocolId() {
	o.ProtocolId.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchExperimentStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExperimentId.IsSet() {
		toSerialize["experiment_id"] = o.ExperimentId.Get()
	}
	if o.Hypothesis.IsSet() {
		toSerialize["hypothesis"] = o.Hypothesis.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ProtocolId.IsSet() {
		toSerialize["protocol_id"] = o.ProtocolId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchExperimentStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExperimentId datadog.NullableString `json:"experiment_id,omitempty"`
		Hypothesis   datadog.NullableString `json:"hypothesis,omitempty"`
		Name         datadog.NullableString `json:"name,omitempty"`
		ProtocolId   datadog.NullableString `json:"protocol_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"experiment_id", "hypothesis", "name", "protocol_id"})
	} else {
		return err
	}
	o.ExperimentId = all.ExperimentId
	o.Hypothesis = all.Hypothesis
	o.Name = all.Name
	o.ProtocolId = all.ProtocolId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
