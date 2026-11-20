// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendDefinitionStyle20270101DisplayOneOf Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendDefinitionStyle20270101DisplayOneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Legend *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend `json:"legend,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewToplistBackendDefinitionStyle20270101DisplayOneOf instantiates a new ToplistBackendDefinitionStyle20270101DisplayOneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistBackendDefinitionStyle20270101DisplayOneOf(typeVar ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType) *ToplistBackendDefinitionStyle20270101DisplayOneOf {
	this := ToplistBackendDefinitionStyle20270101DisplayOneOf{}
	this.Type = typeVar
	return &this
}

// NewToplistBackendDefinitionStyle20270101DisplayOneOfWithDefaults instantiates a new ToplistBackendDefinitionStyle20270101DisplayOneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistBackendDefinitionStyle20270101DisplayOneOfWithDefaults() *ToplistBackendDefinitionStyle20270101DisplayOneOf {
	this := ToplistBackendDefinitionStyle20270101DisplayOneOf{}
	return &this
}

// GetLegend returns the Legend field value if set, zero value otherwise.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) GetLegend() ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend {
	if o == nil || o.Legend == nil {
		var ret ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) GetLegendOk() (*ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend and assigns it to the Legend field.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) SetLegend(v ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend) {
	o.Legend = &v
}

// GetType returns the Type field value.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) GetType() ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType {
	if o == nil {
		var ret ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) GetTypeOk() (*ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) SetType(v ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistBackendDefinitionStyle20270101DisplayOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Legend != nil {
		toSerialize["legend"] = o.Legend
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Legend *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend `json:"legend,omitempty"`
		Type   *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if all.Legend != nil && !all.Legend.IsValid() {
		hasInvalidField = true
	} else {
		o.Legend = all.Legend
	}
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
