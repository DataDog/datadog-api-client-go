// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemJourneyCreateAttributes Attributes for creating or updating a DEM journey.
type DemJourneyCreateAttributes struct {
	// An optional human-readable description of the journey.
	Description *string `json:"description,omitempty"`
	// The RUM definition for a DEM journey.
	JourneyRum DemJourneyRum `json:"journey_rum"`
	// The name of the DEM journey.
	Name string `json:"name"`
	// List of tags associated with a DEM resource.
	Tags []string `json:"tags"`
	// List of variants associated with a DEM journey.
	Variants []DemVariant `json:"variants,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemJourneyCreateAttributes instantiates a new DemJourneyCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemJourneyCreateAttributes(journeyRum DemJourneyRum, name string, tags []string) *DemJourneyCreateAttributes {
	this := DemJourneyCreateAttributes{}
	this.JourneyRum = journeyRum
	this.Name = name
	this.Tags = tags
	return &this
}

// NewDemJourneyCreateAttributesWithDefaults instantiates a new DemJourneyCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemJourneyCreateAttributesWithDefaults() *DemJourneyCreateAttributes {
	this := DemJourneyCreateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DemJourneyCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemJourneyCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DemJourneyCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DemJourneyCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetJourneyRum returns the JourneyRum field value.
func (o *DemJourneyCreateAttributes) GetJourneyRum() DemJourneyRum {
	if o == nil {
		var ret DemJourneyRum
		return ret
	}
	return o.JourneyRum
}

// GetJourneyRumOk returns a tuple with the JourneyRum field value
// and a boolean to check if the value has been set.
func (o *DemJourneyCreateAttributes) GetJourneyRumOk() (*DemJourneyRum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JourneyRum, true
}

// SetJourneyRum sets field value.
func (o *DemJourneyCreateAttributes) SetJourneyRum(v DemJourneyRum) {
	o.JourneyRum = v
}

// GetName returns the Name field value.
func (o *DemJourneyCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DemJourneyCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DemJourneyCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value.
func (o *DemJourneyCreateAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *DemJourneyCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *DemJourneyCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *DemJourneyCreateAttributes) GetVariants() []DemVariant {
	if o == nil || o.Variants == nil {
		var ret []DemVariant
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemJourneyCreateAttributes) GetVariantsOk() (*[]DemVariant, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return &o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *DemJourneyCreateAttributes) HasVariants() bool {
	return o != nil && o.Variants != nil
}

// SetVariants gets a reference to the given []DemVariant and assigns it to the Variants field.
func (o *DemJourneyCreateAttributes) SetVariants(v []DemVariant) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemJourneyCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["journey_rum"] = o.JourneyRum
	toSerialize["name"] = o.Name
	toSerialize["tags"] = o.Tags
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemJourneyCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string        `json:"description,omitempty"`
		JourneyRum  *DemJourneyRum `json:"journey_rum"`
		Name        *string        `json:"name"`
		Tags        *[]string      `json:"tags"`
		Variants    []DemVariant   `json:"variants,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JourneyRum == nil {
		return fmt.Errorf("required field journey_rum missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "journey_rum", "name", "tags", "variants"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	if all.JourneyRum.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JourneyRum = *all.JourneyRum
	o.Name = *all.Name
	o.Tags = *all.Tags
	o.Variants = all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
