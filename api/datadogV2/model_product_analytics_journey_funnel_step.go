// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyFunnelStep A single step of the funnel with its conversion counts and timings.
type ProductAnalyticsJourneyFunnelStep struct {
	// Elapsed time statistics (min/max/avg in milliseconds).
	ElapsedTimeToNextStep ProductAnalyticsElapsedTime `json:"elapsed_time_to_next_step"`
	// Breakdown of this step by the requested group-by facets. Omitted when the query has no group-by.
	Groups []ProductAnalyticsJourneyFunnelStepGroup `json:"groups,omitempty"`
	// Label of the step, derived from the node alias.
	Label string `json:"label"`
	// The entity counted in `value`, such as `session` or `user`, set by the join key.
	Unit string `json:"unit"`
	// Value of the computed metric at this step.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyFunnelStep instantiates a new ProductAnalyticsJourneyFunnelStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyFunnelStep(elapsedTimeToNextStep ProductAnalyticsElapsedTime, label string, unit string, value float64) *ProductAnalyticsJourneyFunnelStep {
	this := ProductAnalyticsJourneyFunnelStep{}
	this.ElapsedTimeToNextStep = elapsedTimeToNextStep
	this.Label = label
	this.Unit = unit
	this.Value = value
	return &this
}

// NewProductAnalyticsJourneyFunnelStepWithDefaults instantiates a new ProductAnalyticsJourneyFunnelStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyFunnelStepWithDefaults() *ProductAnalyticsJourneyFunnelStep {
	this := ProductAnalyticsJourneyFunnelStep{}
	return &this
}

// GetElapsedTimeToNextStep returns the ElapsedTimeToNextStep field value.
func (o *ProductAnalyticsJourneyFunnelStep) GetElapsedTimeToNextStep() ProductAnalyticsElapsedTime {
	if o == nil {
		var ret ProductAnalyticsElapsedTime
		return ret
	}
	return o.ElapsedTimeToNextStep
}

// GetElapsedTimeToNextStepOk returns a tuple with the ElapsedTimeToNextStep field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelStep) GetElapsedTimeToNextStepOk() (*ProductAnalyticsElapsedTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ElapsedTimeToNextStep, true
}

// SetElapsedTimeToNextStep sets field value.
func (o *ProductAnalyticsJourneyFunnelStep) SetElapsedTimeToNextStep(v ProductAnalyticsElapsedTime) {
	o.ElapsedTimeToNextStep = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyFunnelStep) GetGroups() []ProductAnalyticsJourneyFunnelStepGroup {
	if o == nil || o.Groups == nil {
		var ret []ProductAnalyticsJourneyFunnelStepGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelStep) GetGroupsOk() (*[]ProductAnalyticsJourneyFunnelStepGroup, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyFunnelStep) HasGroups() bool {
	return o != nil && o.Groups != nil
}

// SetGroups gets a reference to the given []ProductAnalyticsJourneyFunnelStepGroup and assigns it to the Groups field.
func (o *ProductAnalyticsJourneyFunnelStep) SetGroups(v []ProductAnalyticsJourneyFunnelStepGroup) {
	o.Groups = v
}

// GetLabel returns the Label field value.
func (o *ProductAnalyticsJourneyFunnelStep) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelStep) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *ProductAnalyticsJourneyFunnelStep) SetLabel(v string) {
	o.Label = v
}

// GetUnit returns the Unit field value.
func (o *ProductAnalyticsJourneyFunnelStep) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelStep) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *ProductAnalyticsJourneyFunnelStep) SetUnit(v string) {
	o.Unit = v
}

// GetValue returns the Value field value.
func (o *ProductAnalyticsJourneyFunnelStep) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelStep) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ProductAnalyticsJourneyFunnelStep) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyFunnelStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["elapsed_time_to_next_step"] = o.ElapsedTimeToNextStep
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	toSerialize["label"] = o.Label
	toSerialize["unit"] = o.Unit
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyFunnelStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ElapsedTimeToNextStep *ProductAnalyticsElapsedTime             `json:"elapsed_time_to_next_step"`
		Groups                []ProductAnalyticsJourneyFunnelStepGroup `json:"groups,omitempty"`
		Label                 *string                                  `json:"label"`
		Unit                  *string                                  `json:"unit"`
		Value                 *float64                                 `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ElapsedTimeToNextStep == nil {
		return fmt.Errorf("required field elapsed_time_to_next_step missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"elapsed_time_to_next_step", "groups", "label", "unit", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ElapsedTimeToNextStep.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElapsedTimeToNextStep = *all.ElapsedTimeToNextStep
	o.Groups = all.Groups
	o.Label = *all.Label
	o.Unit = *all.Unit
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
