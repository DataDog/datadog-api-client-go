// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CohortScope Test-only generated schema for the dashboard versioning proof of concept.
type CohortScope struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Target CohortScope20270101Target `json:"target"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type CohortScopeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCohortScope instantiates a new CohortScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCohortScope(target CohortScope20270101Target, typeVar CohortScopeType) *CohortScope {
	this := CohortScope{}
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewCohortScopeWithDefaults instantiates a new CohortScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCohortScopeWithDefaults() *CohortScope {
	this := CohortScope{}
	return &this
}

// GetTarget returns the Target field value.
func (o *CohortScope) GetTarget() CohortScope20270101Target {
	if o == nil {
		var ret CohortScope20270101Target
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *CohortScope) GetTargetOk() (*CohortScope20270101Target, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *CohortScope) SetTarget(v CohortScope20270101Target) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *CohortScope) GetType() CohortScopeType {
	if o == nil {
		var ret CohortScopeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CohortScope) GetTypeOk() (*CohortScopeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CohortScope) SetType(v CohortScopeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CohortScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CohortScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target *CohortScope20270101Target `json:"target"`
		Type   *CohortScopeType           `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Target = *all.Target
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
