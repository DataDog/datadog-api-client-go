// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumStaticSegmentJourneyFilter A filter within a journey query node.
type RumStaticSegmentJourneyFilter struct {
	// The attribute to filter on.
	Attribute string `json:"attribute"`
	// The value to match.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumStaticSegmentJourneyFilter instantiates a new RumStaticSegmentJourneyFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumStaticSegmentJourneyFilter(attribute string, value string) *RumStaticSegmentJourneyFilter {
	this := RumStaticSegmentJourneyFilter{}
	this.Attribute = attribute
	this.Value = value
	return &this
}

// NewRumStaticSegmentJourneyFilterWithDefaults instantiates a new RumStaticSegmentJourneyFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumStaticSegmentJourneyFilterWithDefaults() *RumStaticSegmentJourneyFilter {
	this := RumStaticSegmentJourneyFilter{}
	return &this
}

// GetAttribute returns the Attribute field value.
func (o *RumStaticSegmentJourneyFilter) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *RumStaticSegmentJourneyFilter) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value.
func (o *RumStaticSegmentJourneyFilter) SetAttribute(v string) {
	o.Attribute = v
}

// GetValue returns the Value field value.
func (o *RumStaticSegmentJourneyFilter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RumStaticSegmentJourneyFilter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *RumStaticSegmentJourneyFilter) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumStaticSegmentJourneyFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attribute"] = o.Attribute
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumStaticSegmentJourneyFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute *string `json:"attribute"`
		Value     *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attribute == nil {
		return fmt.Errorf("required field attribute missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "value"})
	} else {
		return err
	}
	o.Attribute = *all.Attribute
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
