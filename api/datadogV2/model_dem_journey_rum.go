// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemJourneyRum The RUM definition for a DEM journey.
type DemJourneyRum struct {
	// An optional RUM query filter applied to the entire journey.
	Filter *string `json:"filter,omitempty"`
	// List of RUM journey steps.
	RumSteps []DemRumStep `json:"rum_steps"`
	// List of variants associated with a DEM journey.
	Variants []DemVariant `json:"variants,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemJourneyRum instantiates a new DemJourneyRum object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemJourneyRum(rumSteps []DemRumStep) *DemJourneyRum {
	this := DemJourneyRum{}
	this.RumSteps = rumSteps
	return &this
}

// NewDemJourneyRumWithDefaults instantiates a new DemJourneyRum object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemJourneyRumWithDefaults() *DemJourneyRum {
	this := DemJourneyRum{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DemJourneyRum) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemJourneyRum) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DemJourneyRum) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *DemJourneyRum) SetFilter(v string) {
	o.Filter = &v
}

// GetRumSteps returns the RumSteps field value.
func (o *DemJourneyRum) GetRumSteps() []DemRumStep {
	if o == nil {
		var ret []DemRumStep
		return ret
	}
	return o.RumSteps
}

// GetRumStepsOk returns a tuple with the RumSteps field value
// and a boolean to check if the value has been set.
func (o *DemJourneyRum) GetRumStepsOk() (*[]DemRumStep, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RumSteps, true
}

// SetRumSteps sets field value.
func (o *DemJourneyRum) SetRumSteps(v []DemRumStep) {
	o.RumSteps = v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *DemJourneyRum) GetVariants() []DemVariant {
	if o == nil || o.Variants == nil {
		var ret []DemVariant
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemJourneyRum) GetVariantsOk() (*[]DemVariant, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return &o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *DemJourneyRum) HasVariants() bool {
	return o != nil && o.Variants != nil
}

// SetVariants gets a reference to the given []DemVariant and assigns it to the Variants field.
func (o *DemJourneyRum) SetVariants(v []DemVariant) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemJourneyRum) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["rum_steps"] = o.RumSteps
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemJourneyRum) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter   *string       `json:"filter,omitempty"`
		RumSteps *[]DemRumStep `json:"rum_steps"`
		Variants []DemVariant  `json:"variants,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RumSteps == nil {
		return fmt.Errorf("required field rum_steps missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "rum_steps", "variants"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.RumSteps = *all.RumSteps
	o.Variants = all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
