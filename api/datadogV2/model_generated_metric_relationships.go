// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeneratedMetricRelationships Relationships for a generated metric.
type GeneratedMetricRelationships struct {
	// Relationship to the attributes of a generated metric.
	GeneratedMetricAttributes *GeneratedMetricAttributesRelationship `json:"generated_metric_attributes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGeneratedMetricRelationships instantiates a new GeneratedMetricRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeneratedMetricRelationships() *GeneratedMetricRelationships {
	this := GeneratedMetricRelationships{}
	return &this
}

// NewGeneratedMetricRelationshipsWithDefaults instantiates a new GeneratedMetricRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeneratedMetricRelationshipsWithDefaults() *GeneratedMetricRelationships {
	this := GeneratedMetricRelationships{}
	return &this
}

// GetGeneratedMetricAttributes returns the GeneratedMetricAttributes field value if set, zero value otherwise.
func (o *GeneratedMetricRelationships) GetGeneratedMetricAttributes() GeneratedMetricAttributesRelationship {
	if o == nil || o.GeneratedMetricAttributes == nil {
		var ret GeneratedMetricAttributesRelationship
		return ret
	}
	return *o.GeneratedMetricAttributes
}

// GetGeneratedMetricAttributesOk returns a tuple with the GeneratedMetricAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratedMetricRelationships) GetGeneratedMetricAttributesOk() (*GeneratedMetricAttributesRelationship, bool) {
	if o == nil || o.GeneratedMetricAttributes == nil {
		return nil, false
	}
	return o.GeneratedMetricAttributes, true
}

// HasGeneratedMetricAttributes returns a boolean if a field has been set.
func (o *GeneratedMetricRelationships) HasGeneratedMetricAttributes() bool {
	return o != nil && o.GeneratedMetricAttributes != nil
}

// SetGeneratedMetricAttributes gets a reference to the given GeneratedMetricAttributesRelationship and assigns it to the GeneratedMetricAttributes field.
func (o *GeneratedMetricRelationships) SetGeneratedMetricAttributes(v GeneratedMetricAttributesRelationship) {
	o.GeneratedMetricAttributes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeneratedMetricRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GeneratedMetricAttributes != nil {
		toSerialize["generated_metric_attributes"] = o.GeneratedMetricAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeneratedMetricRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GeneratedMetricAttributes *GeneratedMetricAttributesRelationship `json:"generated_metric_attributes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"generated_metric_attributes"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GeneratedMetricAttributes != nil && all.GeneratedMetricAttributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GeneratedMetricAttributes = all.GeneratedMetricAttributes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
