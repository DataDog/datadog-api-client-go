// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIImpactUserActivityData A single daily AI tool activity entry.
type AIImpactUserActivityData struct {
	// Daily AI coding tool activity for a single user. Each entry reports whether the user was
	// active on a given day and which AI tools and models they used.
	Attributes AIImpactUserActivityAttributes `json:"attributes"`
	// JSON:API type for AI Impact user activity entries.
	Type AIImpactUserActivityType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAIImpactUserActivityData instantiates a new AIImpactUserActivityData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIImpactUserActivityData(attributes AIImpactUserActivityAttributes, typeVar AIImpactUserActivityType) *AIImpactUserActivityData {
	this := AIImpactUserActivityData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewAIImpactUserActivityDataWithDefaults instantiates a new AIImpactUserActivityData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIImpactUserActivityDataWithDefaults() *AIImpactUserActivityData {
	this := AIImpactUserActivityData{}
	var typeVar AIImpactUserActivityType = AIIMPACTUSERACTIVITYTYPE_AI_IMPACT_USER_ACTIVITY
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AIImpactUserActivityData) GetAttributes() AIImpactUserActivityAttributes {
	if o == nil {
		var ret AIImpactUserActivityAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AIImpactUserActivityData) GetAttributesOk() (*AIImpactUserActivityAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AIImpactUserActivityData) SetAttributes(v AIImpactUserActivityAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *AIImpactUserActivityData) GetType() AIImpactUserActivityType {
	if o == nil {
		var ret AIImpactUserActivityType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AIImpactUserActivityData) GetTypeOk() (*AIImpactUserActivityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AIImpactUserActivityData) SetType(v AIImpactUserActivityType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIImpactUserActivityData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIImpactUserActivityData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AIImpactUserActivityAttributes `json:"attributes"`
		Type       *AIImpactUserActivityType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
