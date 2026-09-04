// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemVariantAttributes Attributes of a DEM journey variant.
type DemVariantAttributes struct {
	// An optional RUM query filter to scope this variant.
	Filter *string `json:"filter,omitempty"`
	// The name of the variant.
	Name string `json:"name"`
	// List of RUM journey steps.
	RumSteps []DemRumStep `json:"rum_steps"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemVariantAttributes instantiates a new DemVariantAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemVariantAttributes(name string, rumSteps []DemRumStep) *DemVariantAttributes {
	this := DemVariantAttributes{}
	this.Name = name
	this.RumSteps = rumSteps
	return &this
}

// NewDemVariantAttributesWithDefaults instantiates a new DemVariantAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemVariantAttributesWithDefaults() *DemVariantAttributes {
	this := DemVariantAttributes{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DemVariantAttributes) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemVariantAttributes) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DemVariantAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *DemVariantAttributes) SetFilter(v string) {
	o.Filter = &v
}

// GetName returns the Name field value.
func (o *DemVariantAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DemVariantAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DemVariantAttributes) SetName(v string) {
	o.Name = v
}

// GetRumSteps returns the RumSteps field value.
func (o *DemVariantAttributes) GetRumSteps() []DemRumStep {
	if o == nil {
		var ret []DemRumStep
		return ret
	}
	return o.RumSteps
}

// GetRumStepsOk returns a tuple with the RumSteps field value
// and a boolean to check if the value has been set.
func (o *DemVariantAttributes) GetRumStepsOk() (*[]DemRumStep, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RumSteps, true
}

// SetRumSteps sets field value.
func (o *DemVariantAttributes) SetRumSteps(v []DemRumStep) {
	o.RumSteps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemVariantAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["name"] = o.Name
	toSerialize["rum_steps"] = o.RumSteps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemVariantAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter   *string       `json:"filter,omitempty"`
		Name     *string       `json:"name"`
		RumSteps *[]DemRumStep `json:"rum_steps"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.RumSteps == nil {
		return fmt.Errorf("required field rum_steps missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "name", "rum_steps"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.Name = *all.Name
	o.RumSteps = *all.RumSteps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
