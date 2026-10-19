// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DerivedMetricAttributes Attributes of a derived metric.
type DerivedMetricAttributes struct {
	// The definition of a derived metric as a formula over one or more queries.
	Definition DerivedMetricDefinition `json:"definition"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDerivedMetricAttributes instantiates a new DerivedMetricAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDerivedMetricAttributes(definition DerivedMetricDefinition) *DerivedMetricAttributes {
	this := DerivedMetricAttributes{}
	this.Definition = definition
	return &this
}

// NewDerivedMetricAttributesWithDefaults instantiates a new DerivedMetricAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDerivedMetricAttributesWithDefaults() *DerivedMetricAttributes {
	this := DerivedMetricAttributes{}
	return &this
}

// GetDefinition returns the Definition field value.
func (o *DerivedMetricAttributes) GetDefinition() DerivedMetricDefinition {
	if o == nil {
		var ret DerivedMetricDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *DerivedMetricAttributes) GetDefinitionOk() (*DerivedMetricDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *DerivedMetricAttributes) SetDefinition(v DerivedMetricDefinition) {
	o.Definition = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DerivedMetricAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["definition"] = o.Definition

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DerivedMetricAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition *DerivedMetricDefinition `json:"definition"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"definition"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
