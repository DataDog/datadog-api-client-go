// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyDefinition Sankey visualization definition. Set either `source` or `target`, not both.
// Use `source` for forward flow (where do users go after this page?) or
// `target` for backward flow (where did users come from?).
type ProductAnalyticsSankeyDefinition struct {
	// Number of page entries per step. Default 5, max 10.
	EntriesPerStep *int64 `json:"entries_per_step,omitempty"`
	// Number of steps in the flow. Default 5, max 10.
	NumberOfSteps *int64 `json:"number_of_steps,omitempty"`
	// The source page for forward flow analysis. Use "*" for all pages.
	Source *string `json:"source,omitempty"`
	// The target page for backward flow analysis.
	Target *string `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsSankeyDefinition instantiates a new ProductAnalyticsSankeyDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsSankeyDefinition() *ProductAnalyticsSankeyDefinition {
	this := ProductAnalyticsSankeyDefinition{}
	var entriesPerStep int64 = 5
	this.EntriesPerStep = &entriesPerStep
	var numberOfSteps int64 = 5
	this.NumberOfSteps = &numberOfSteps
	return &this
}

// NewProductAnalyticsSankeyDefinitionWithDefaults instantiates a new ProductAnalyticsSankeyDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsSankeyDefinitionWithDefaults() *ProductAnalyticsSankeyDefinition {
	this := ProductAnalyticsSankeyDefinition{}
	var entriesPerStep int64 = 5
	this.EntriesPerStep = &entriesPerStep
	var numberOfSteps int64 = 5
	this.NumberOfSteps = &numberOfSteps
	return &this
}

// GetEntriesPerStep returns the EntriesPerStep field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyDefinition) GetEntriesPerStep() int64 {
	if o == nil || o.EntriesPerStep == nil {
		var ret int64
		return ret
	}
	return *o.EntriesPerStep
}

// GetEntriesPerStepOk returns a tuple with the EntriesPerStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyDefinition) GetEntriesPerStepOk() (*int64, bool) {
	if o == nil || o.EntriesPerStep == nil {
		return nil, false
	}
	return o.EntriesPerStep, true
}

// HasEntriesPerStep returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyDefinition) HasEntriesPerStep() bool {
	return o != nil && o.EntriesPerStep != nil
}

// SetEntriesPerStep gets a reference to the given int64 and assigns it to the EntriesPerStep field.
func (o *ProductAnalyticsSankeyDefinition) SetEntriesPerStep(v int64) {
	o.EntriesPerStep = &v
}

// GetNumberOfSteps returns the NumberOfSteps field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyDefinition) GetNumberOfSteps() int64 {
	if o == nil || o.NumberOfSteps == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSteps
}

// GetNumberOfStepsOk returns a tuple with the NumberOfSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyDefinition) GetNumberOfStepsOk() (*int64, bool) {
	if o == nil || o.NumberOfSteps == nil {
		return nil, false
	}
	return o.NumberOfSteps, true
}

// HasNumberOfSteps returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyDefinition) HasNumberOfSteps() bool {
	return o != nil && o.NumberOfSteps != nil
}

// SetNumberOfSteps gets a reference to the given int64 and assigns it to the NumberOfSteps field.
func (o *ProductAnalyticsSankeyDefinition) SetNumberOfSteps(v int64) {
	o.NumberOfSteps = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyDefinition) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyDefinition) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyDefinition) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ProductAnalyticsSankeyDefinition) SetSource(v string) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyDefinition) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyDefinition) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyDefinition) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ProductAnalyticsSankeyDefinition) SetTarget(v string) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsSankeyDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EntriesPerStep != nil {
		toSerialize["entries_per_step"] = o.EntriesPerStep
	}
	if o.NumberOfSteps != nil {
		toSerialize["number_of_steps"] = o.NumberOfSteps
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsSankeyDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EntriesPerStep *int64  `json:"entries_per_step,omitempty"`
		NumberOfSteps  *int64  `json:"number_of_steps,omitempty"`
		Source         *string `json:"source,omitempty"`
		Target         *string `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entries_per_step", "number_of_steps", "source", "target"})
	} else {
		return err
	}
	o.EntriesPerStep = all.EntriesPerStep
	o.NumberOfSteps = all.NumberOfSteps
	o.Source = all.Source
	o.Target = all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
