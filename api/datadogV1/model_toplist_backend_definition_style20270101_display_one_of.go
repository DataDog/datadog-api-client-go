// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendDefinitionStyle20270101DisplayOneOf Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendDefinitionStyle20270101DisplayOneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Legend *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend_20270101 `json:"legend,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType_20270101 `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewToplistBackendDefinitionStyle20270101DisplayOneOf instantiates a new ToplistBackendDefinitionStyle20270101DisplayOneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistBackendDefinitionStyle20270101DisplayOneOf(typeVar ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType_20270101) *ToplistBackendDefinitionStyle20270101DisplayOneOf {
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
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) GetLegend() ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend_20270101 {
	if o == nil || o.Legend == nil {
		var ret ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend_20270101
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) GetLegendOk() (*ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend_20270101, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend_20270101 and assigns it to the Legend field.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) SetLegend(v ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend_20270101) {
	o.Legend = &v
}

// GetType returns the Type field value.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) GetType() ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType_20270101 {
	if o == nil {
		var ret ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType_20270101
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) GetTypeOk() (*ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ToplistBackendDefinitionStyle20270101DisplayOneOf) SetType(v ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType_20270101) {
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
		Legend *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend_20270101 `json:"legend,omitempty"`
		Type   *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType_20270101   `json:"type"`
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
