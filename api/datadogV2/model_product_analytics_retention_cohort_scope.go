// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionCohortScope Narrows a retention query to a single cohort row.
type ProductAnalyticsRetentionCohortScope struct {
	// Selects a cohort, either by index or by the aggregation that rolls all cohorts together.
	Target ProductAnalyticsRetentionCohortTarget `json:"target"`
	// The discriminator identifying a scope narrowed to one cohort.
	Type ProductAnalyticsRetentionCohortScopeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionCohortScope instantiates a new ProductAnalyticsRetentionCohortScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionCohortScope(target ProductAnalyticsRetentionCohortTarget, typeVar ProductAnalyticsRetentionCohortScopeType) *ProductAnalyticsRetentionCohortScope {
	this := ProductAnalyticsRetentionCohortScope{}
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsRetentionCohortScopeWithDefaults instantiates a new ProductAnalyticsRetentionCohortScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionCohortScopeWithDefaults() *ProductAnalyticsRetentionCohortScope {
	this := ProductAnalyticsRetentionCohortScope{}
	return &this
}

// GetTarget returns the Target field value.
func (o *ProductAnalyticsRetentionCohortScope) GetTarget() ProductAnalyticsRetentionCohortTarget {
	if o == nil {
		var ret ProductAnalyticsRetentionCohortTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCohortScope) GetTargetOk() (*ProductAnalyticsRetentionCohortTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ProductAnalyticsRetentionCohortScope) SetTarget(v ProductAnalyticsRetentionCohortTarget) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsRetentionCohortScope) GetType() ProductAnalyticsRetentionCohortScopeType {
	if o == nil {
		var ret ProductAnalyticsRetentionCohortScopeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCohortScope) GetTypeOk() (*ProductAnalyticsRetentionCohortScopeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsRetentionCohortScope) SetType(v ProductAnalyticsRetentionCohortScopeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionCohortScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionCohortScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target *ProductAnalyticsRetentionCohortTarget    `json:"target"`
		Type   *ProductAnalyticsRetentionCohortScopeType `json:"type"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Target = *all.Target
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
