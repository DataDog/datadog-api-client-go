// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelSuggestionRequestDataAttributesSearchStepsItems
type FunnelSuggestionRequestDataAttributesSearchStepsItems struct {
	//
	Facet *string `json:"facet,omitempty"`
	//
	StepFilter *string `json:"step_filter,omitempty"`
	//
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelSuggestionRequestDataAttributesSearchStepsItems instantiates a new FunnelSuggestionRequestDataAttributesSearchStepsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelSuggestionRequestDataAttributesSearchStepsItems() *FunnelSuggestionRequestDataAttributesSearchStepsItems {
	this := FunnelSuggestionRequestDataAttributesSearchStepsItems{}
	return &this
}

// NewFunnelSuggestionRequestDataAttributesSearchStepsItemsWithDefaults instantiates a new FunnelSuggestionRequestDataAttributesSearchStepsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelSuggestionRequestDataAttributesSearchStepsItemsWithDefaults() *FunnelSuggestionRequestDataAttributesSearchStepsItems {
	this := FunnelSuggestionRequestDataAttributesSearchStepsItems{}
	return &this
}

// GetFacet returns the Facet field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) GetFacet() string {
	if o == nil || o.Facet == nil {
		var ret string
		return ret
	}
	return *o.Facet
}

// GetFacetOk returns a tuple with the Facet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) GetFacetOk() (*string, bool) {
	if o == nil || o.Facet == nil {
		return nil, false
	}
	return o.Facet, true
}

// HasFacet returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) HasFacet() bool {
	return o != nil && o.Facet != nil
}

// SetFacet gets a reference to the given string and assigns it to the Facet field.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) SetFacet(v string) {
	o.Facet = &v
}

// GetStepFilter returns the StepFilter field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) GetStepFilter() string {
	if o == nil || o.StepFilter == nil {
		var ret string
		return ret
	}
	return *o.StepFilter
}

// GetStepFilterOk returns a tuple with the StepFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) GetStepFilterOk() (*string, bool) {
	if o == nil || o.StepFilter == nil {
		return nil, false
	}
	return o.StepFilter, true
}

// HasStepFilter returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) HasStepFilter() bool {
	return o != nil && o.StepFilter != nil
}

// SetStepFilter gets a reference to the given string and assigns it to the StepFilter field.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) SetStepFilter(v string) {
	o.StepFilter = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelSuggestionRequestDataAttributesSearchStepsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Facet != nil {
		toSerialize["facet"] = o.Facet
	}
	if o.StepFilter != nil {
		toSerialize["step_filter"] = o.StepFilter
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
func (o *FunnelSuggestionRequestDataAttributesSearchStepsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet      *string `json:"facet,omitempty"`
		StepFilter *string `json:"step_filter,omitempty"`
		Value      *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "step_filter", "value"})
	} else {
		return err
	}
	o.Facet = all.Facet
	o.StepFilter = all.StepFilter
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
