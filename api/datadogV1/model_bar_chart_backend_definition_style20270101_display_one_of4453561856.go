// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544_20270101 `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewBarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 instantiates a new BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBarChartBackendDefinitionStyle20270101DisplayOneOf4453561856(typeVar BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544_20270101) *BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 {
	this := BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856{}
	this.Type = typeVar
	return &this
}

// NewBarChartBackendDefinitionStyle20270101DisplayOneOf4453561856WithDefaults instantiates a new BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBarChartBackendDefinitionStyle20270101DisplayOneOf4453561856WithDefaults() *BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856 {
	this := BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856{}
	return &this
}

// GetType returns the Type field value.
func (o *BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856) GetType() BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544_20270101 {
	if o == nil {
		var ret BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544_20270101
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856) GetTypeOk() (*BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856) SetType(v BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544_20270101) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BarChartBackendDefinitionStyle20270101DisplayOneOf4453561856) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544_20270101 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
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
