// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionReturnPeriodScope Narrows a retention query to a single return-period column.
type ProductAnalyticsRetentionReturnPeriodScope struct {
	// Selects a cohort or return period by its zero-based position in the grid.
	Target ProductAnalyticsRetentionIndexTarget `json:"target"`
	// The discriminator identifying a scope narrowed to one return period.
	Type ProductAnalyticsRetentionReturnPeriodScopeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionReturnPeriodScope instantiates a new ProductAnalyticsRetentionReturnPeriodScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionReturnPeriodScope(target ProductAnalyticsRetentionIndexTarget, typeVar ProductAnalyticsRetentionReturnPeriodScopeType) *ProductAnalyticsRetentionReturnPeriodScope {
	this := ProductAnalyticsRetentionReturnPeriodScope{}
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsRetentionReturnPeriodScopeWithDefaults instantiates a new ProductAnalyticsRetentionReturnPeriodScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionReturnPeriodScopeWithDefaults() *ProductAnalyticsRetentionReturnPeriodScope {
	this := ProductAnalyticsRetentionReturnPeriodScope{}
	return &this
}

// GetTarget returns the Target field value.
func (o *ProductAnalyticsRetentionReturnPeriodScope) GetTarget() ProductAnalyticsRetentionIndexTarget {
	if o == nil {
		var ret ProductAnalyticsRetentionIndexTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionReturnPeriodScope) GetTargetOk() (*ProductAnalyticsRetentionIndexTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ProductAnalyticsRetentionReturnPeriodScope) SetTarget(v ProductAnalyticsRetentionIndexTarget) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsRetentionReturnPeriodScope) GetType() ProductAnalyticsRetentionReturnPeriodScopeType {
	if o == nil {
		var ret ProductAnalyticsRetentionReturnPeriodScopeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionReturnPeriodScope) GetTypeOk() (*ProductAnalyticsRetentionReturnPeriodScopeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsRetentionReturnPeriodScope) SetType(v ProductAnalyticsRetentionReturnPeriodScopeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionReturnPeriodScope) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsRetentionReturnPeriodScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target *ProductAnalyticsRetentionIndexTarget           `json:"target"`
		Type   *ProductAnalyticsRetentionReturnPeriodScopeType `json:"type"`
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
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
