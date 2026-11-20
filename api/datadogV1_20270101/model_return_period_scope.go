// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReturnPeriodScope Test-only generated schema for the dashboard versioning proof of concept.
type ReturnPeriodScope struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Target ReturnPeriodScopeTarget `json:"target"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ReturnPeriodScopeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewReturnPeriodScope instantiates a new ReturnPeriodScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReturnPeriodScope(target ReturnPeriodScopeTarget, typeVar ReturnPeriodScopeType) *ReturnPeriodScope {
	this := ReturnPeriodScope{}
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewReturnPeriodScopeWithDefaults instantiates a new ReturnPeriodScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReturnPeriodScopeWithDefaults() *ReturnPeriodScope {
	this := ReturnPeriodScope{}
	return &this
}

// GetTarget returns the Target field value.
func (o *ReturnPeriodScope) GetTarget() ReturnPeriodScopeTarget {
	if o == nil {
		var ret ReturnPeriodScopeTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ReturnPeriodScope) GetTargetOk() (*ReturnPeriodScopeTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ReturnPeriodScope) SetTarget(v ReturnPeriodScopeTarget) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *ReturnPeriodScope) GetType() ReturnPeriodScopeType {
	if o == nil {
		var ret ReturnPeriodScopeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReturnPeriodScope) GetTypeOk() (*ReturnPeriodScopeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ReturnPeriodScope) SetType(v ReturnPeriodScopeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReturnPeriodScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReturnPeriodScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target *ReturnPeriodScopeTarget `json:"target"`
		Type   *ReturnPeriodScopeType   `json:"type"`
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
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
