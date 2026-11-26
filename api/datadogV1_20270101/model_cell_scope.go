// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CellScope Test-only generated schema for the dashboard versioning proof of concept.
type CellScope struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CohortTarget CellScope20270101CohortTarget `json:"cohort_target"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ReturnPeriodTarget CellScopeReturnPeriodTarget `json:"return_period_target"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type CellScopeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCellScope instantiates a new CellScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCellScope(cohortTarget CellScope20270101CohortTarget, returnPeriodTarget CellScopeReturnPeriodTarget, typeVar CellScopeType) *CellScope {
	this := CellScope{}
	this.CohortTarget = cohortTarget
	this.ReturnPeriodTarget = returnPeriodTarget
	this.Type = typeVar
	return &this
}

// NewCellScopeWithDefaults instantiates a new CellScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCellScopeWithDefaults() *CellScope {
	this := CellScope{}
	return &this
}

// GetCohortTarget returns the CohortTarget field value.
func (o *CellScope) GetCohortTarget() CellScope20270101CohortTarget {
	if o == nil {
		var ret CellScope20270101CohortTarget
		return ret
	}
	return o.CohortTarget
}

// GetCohortTargetOk returns a tuple with the CohortTarget field value
// and a boolean to check if the value has been set.
func (o *CellScope) GetCohortTargetOk() (*CellScope20270101CohortTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CohortTarget, true
}

// SetCohortTarget sets field value.
func (o *CellScope) SetCohortTarget(v CellScope20270101CohortTarget) {
	o.CohortTarget = v
}

// GetReturnPeriodTarget returns the ReturnPeriodTarget field value.
func (o *CellScope) GetReturnPeriodTarget() CellScopeReturnPeriodTarget {
	if o == nil {
		var ret CellScopeReturnPeriodTarget
		return ret
	}
	return o.ReturnPeriodTarget
}

// GetReturnPeriodTargetOk returns a tuple with the ReturnPeriodTarget field value
// and a boolean to check if the value has been set.
func (o *CellScope) GetReturnPeriodTargetOk() (*CellScopeReturnPeriodTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnPeriodTarget, true
}

// SetReturnPeriodTarget sets field value.
func (o *CellScope) SetReturnPeriodTarget(v CellScopeReturnPeriodTarget) {
	o.ReturnPeriodTarget = v
}

// GetType returns the Type field value.
func (o *CellScope) GetType() CellScopeType {
	if o == nil {
		var ret CellScopeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CellScope) GetTypeOk() (*CellScopeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CellScope) SetType(v CellScopeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CellScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cohort_target"] = o.CohortTarget
	toSerialize["return_period_target"] = o.ReturnPeriodTarget
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CellScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CohortTarget       *CellScope20270101CohortTarget `json:"cohort_target"`
		ReturnPeriodTarget *CellScopeReturnPeriodTarget   `json:"return_period_target"`
		Type               *CellScopeType                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CohortTarget == nil {
		return fmt.Errorf("required field cohort_target missing")
	}
	if all.ReturnPeriodTarget == nil {
		return fmt.Errorf("required field return_period_target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.CohortTarget = *all.CohortTarget
	if all.ReturnPeriodTarget.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReturnPeriodTarget = *all.ReturnPeriodTarget
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
