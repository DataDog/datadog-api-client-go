// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionCellScope Narrows a retention query to a single cell, at the intersection of one cohort and one return period.
type ProductAnalyticsRetentionCellScope struct {
	// Selects a cohort, either by index or by the aggregation that rolls all cohorts together.
	CohortTarget ProductAnalyticsRetentionCohortTarget `json:"cohort_target"`
	// Selects a cohort or return period by its zero-based position in the grid.
	ReturnPeriodTarget ProductAnalyticsRetentionIndexTarget `json:"return_period_target"`
	// The discriminator identifying a scope narrowed to one grid cell.
	Type ProductAnalyticsRetentionCellScopeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionCellScope instantiates a new ProductAnalyticsRetentionCellScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionCellScope(cohortTarget ProductAnalyticsRetentionCohortTarget, returnPeriodTarget ProductAnalyticsRetentionIndexTarget, typeVar ProductAnalyticsRetentionCellScopeType) *ProductAnalyticsRetentionCellScope {
	this := ProductAnalyticsRetentionCellScope{}
	this.CohortTarget = cohortTarget
	this.ReturnPeriodTarget = returnPeriodTarget
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsRetentionCellScopeWithDefaults instantiates a new ProductAnalyticsRetentionCellScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionCellScopeWithDefaults() *ProductAnalyticsRetentionCellScope {
	this := ProductAnalyticsRetentionCellScope{}
	return &this
}

// GetCohortTarget returns the CohortTarget field value.
func (o *ProductAnalyticsRetentionCellScope) GetCohortTarget() ProductAnalyticsRetentionCohortTarget {
	if o == nil {
		var ret ProductAnalyticsRetentionCohortTarget
		return ret
	}
	return o.CohortTarget
}

// GetCohortTargetOk returns a tuple with the CohortTarget field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCellScope) GetCohortTargetOk() (*ProductAnalyticsRetentionCohortTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CohortTarget, true
}

// SetCohortTarget sets field value.
func (o *ProductAnalyticsRetentionCellScope) SetCohortTarget(v ProductAnalyticsRetentionCohortTarget) {
	o.CohortTarget = v
}

// GetReturnPeriodTarget returns the ReturnPeriodTarget field value.
func (o *ProductAnalyticsRetentionCellScope) GetReturnPeriodTarget() ProductAnalyticsRetentionIndexTarget {
	if o == nil {
		var ret ProductAnalyticsRetentionIndexTarget
		return ret
	}
	return o.ReturnPeriodTarget
}

// GetReturnPeriodTargetOk returns a tuple with the ReturnPeriodTarget field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCellScope) GetReturnPeriodTargetOk() (*ProductAnalyticsRetentionIndexTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnPeriodTarget, true
}

// SetReturnPeriodTarget sets field value.
func (o *ProductAnalyticsRetentionCellScope) SetReturnPeriodTarget(v ProductAnalyticsRetentionIndexTarget) {
	o.ReturnPeriodTarget = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsRetentionCellScope) GetType() ProductAnalyticsRetentionCellScopeType {
	if o == nil {
		var ret ProductAnalyticsRetentionCellScopeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCellScope) GetTypeOk() (*ProductAnalyticsRetentionCellScopeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsRetentionCellScope) SetType(v ProductAnalyticsRetentionCellScopeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionCellScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cohort_target"] = o.CohortTarget
	toSerialize["return_period_target"] = o.ReturnPeriodTarget
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionCellScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CohortTarget       *ProductAnalyticsRetentionCohortTarget  `json:"cohort_target"`
		ReturnPeriodTarget *ProductAnalyticsRetentionIndexTarget   `json:"return_period_target"`
		Type               *ProductAnalyticsRetentionCellScopeType `json:"type"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cohort_target", "return_period_target", "type"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
