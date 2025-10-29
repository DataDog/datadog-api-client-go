// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelResponseDataAttributesFunnelStepsItems
type FunnelResponseDataAttributesFunnelStepsItems struct {
	//
	ElapsedTimeToNextStep *FunnelResponseElapsedTime `json:"elapsed_time_to_next_step,omitempty"`
	//
	Label *string `json:"label,omitempty"`
	//
	Value *int64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelResponseDataAttributesFunnelStepsItems instantiates a new FunnelResponseDataAttributesFunnelStepsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelResponseDataAttributesFunnelStepsItems() *FunnelResponseDataAttributesFunnelStepsItems {
	this := FunnelResponseDataAttributesFunnelStepsItems{}
	return &this
}

// NewFunnelResponseDataAttributesFunnelStepsItemsWithDefaults instantiates a new FunnelResponseDataAttributesFunnelStepsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelResponseDataAttributesFunnelStepsItemsWithDefaults() *FunnelResponseDataAttributesFunnelStepsItems {
	this := FunnelResponseDataAttributesFunnelStepsItems{}
	return &this
}

// GetElapsedTimeToNextStep returns the ElapsedTimeToNextStep field value if set, zero value otherwise.
func (o *FunnelResponseDataAttributesFunnelStepsItems) GetElapsedTimeToNextStep() FunnelResponseElapsedTime {
	if o == nil || o.ElapsedTimeToNextStep == nil {
		var ret FunnelResponseElapsedTime
		return ret
	}
	return *o.ElapsedTimeToNextStep
}

// GetElapsedTimeToNextStepOk returns a tuple with the ElapsedTimeToNextStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseDataAttributesFunnelStepsItems) GetElapsedTimeToNextStepOk() (*FunnelResponseElapsedTime, bool) {
	if o == nil || o.ElapsedTimeToNextStep == nil {
		return nil, false
	}
	return o.ElapsedTimeToNextStep, true
}

// HasElapsedTimeToNextStep returns a boolean if a field has been set.
func (o *FunnelResponseDataAttributesFunnelStepsItems) HasElapsedTimeToNextStep() bool {
	return o != nil && o.ElapsedTimeToNextStep != nil
}

// SetElapsedTimeToNextStep gets a reference to the given FunnelResponseElapsedTime and assigns it to the ElapsedTimeToNextStep field.
func (o *FunnelResponseDataAttributesFunnelStepsItems) SetElapsedTimeToNextStep(v FunnelResponseElapsedTime) {
	o.ElapsedTimeToNextStep = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FunnelResponseDataAttributesFunnelStepsItems) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseDataAttributesFunnelStepsItems) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FunnelResponseDataAttributesFunnelStepsItems) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FunnelResponseDataAttributesFunnelStepsItems) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FunnelResponseDataAttributesFunnelStepsItems) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelResponseDataAttributesFunnelStepsItems) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FunnelResponseDataAttributesFunnelStepsItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *FunnelResponseDataAttributesFunnelStepsItems) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelResponseDataAttributesFunnelStepsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ElapsedTimeToNextStep != nil {
		toSerialize["elapsed_time_to_next_step"] = o.ElapsedTimeToNextStep
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelResponseDataAttributesFunnelStepsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ElapsedTimeToNextStep *FunnelResponseElapsedTime `json:"elapsed_time_to_next_step,omitempty"`
		Label                 *string                    `json:"label,omitempty"`
		Value                 *int64                     `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"elapsed_time_to_next_step", "label", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ElapsedTimeToNextStep != nil && all.ElapsedTimeToNextStep.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElapsedTimeToNextStep = all.ElapsedTimeToNextStep
	o.Label = all.Label
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
