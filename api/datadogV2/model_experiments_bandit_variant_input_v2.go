// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditVariantInputV2 A flag variant the bandit uses, named by its ID. The variant's `key` and `name` come from the flag and can't be set.
type ExperimentsBanditVariantInputV2 struct {
	// The variant's ID on the flag.
	FeatureFlagVariantId uuid.UUID `json:"feature_flag_variant_id"`
	// Whether the variant is one of the arms the bandit chooses between. Defaults to `true`.
	IsActive *bool `json:"is_active,omitempty"`
	// Whether the variant is served when the bandit reaches its scheduled end. At most one variant is the fallback, and it doesn't have to be an arm. Defaults to `false`.
	IsFallback *bool `json:"is_fallback,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditVariantInputV2 instantiates a new ExperimentsBanditVariantInputV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditVariantInputV2(featureFlagVariantId uuid.UUID) *ExperimentsBanditVariantInputV2 {
	this := ExperimentsBanditVariantInputV2{}
	this.FeatureFlagVariantId = featureFlagVariantId
	var isActive bool = true
	this.IsActive = &isActive
	var isFallback bool = false
	this.IsFallback = &isFallback
	return &this
}

// NewExperimentsBanditVariantInputV2WithDefaults instantiates a new ExperimentsBanditVariantInputV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditVariantInputV2WithDefaults() *ExperimentsBanditVariantInputV2 {
	this := ExperimentsBanditVariantInputV2{}
	var isActive bool = true
	this.IsActive = &isActive
	var isFallback bool = false
	this.IsFallback = &isFallback
	return &this
}

// GetFeatureFlagVariantId returns the FeatureFlagVariantId field value.
func (o *ExperimentsBanditVariantInputV2) GetFeatureFlagVariantId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagVariantId
}

// GetFeatureFlagVariantIdOk returns a tuple with the FeatureFlagVariantId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantInputV2) GetFeatureFlagVariantIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagVariantId, true
}

// SetFeatureFlagVariantId sets field value.
func (o *ExperimentsBanditVariantInputV2) SetFeatureFlagVariantId(v uuid.UUID) {
	o.FeatureFlagVariantId = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ExperimentsBanditVariantInputV2) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantInputV2) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ExperimentsBanditVariantInputV2) HasIsActive() bool {
	return o != nil && o.IsActive != nil
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ExperimentsBanditVariantInputV2) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsFallback returns the IsFallback field value if set, zero value otherwise.
func (o *ExperimentsBanditVariantInputV2) GetIsFallback() bool {
	if o == nil || o.IsFallback == nil {
		var ret bool
		return ret
	}
	return *o.IsFallback
}

// GetIsFallbackOk returns a tuple with the IsFallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantInputV2) GetIsFallbackOk() (*bool, bool) {
	if o == nil || o.IsFallback == nil {
		return nil, false
	}
	return o.IsFallback, true
}

// HasIsFallback returns a boolean if a field has been set.
func (o *ExperimentsBanditVariantInputV2) HasIsFallback() bool {
	return o != nil && o.IsFallback != nil
}

// SetIsFallback gets a reference to the given bool and assigns it to the IsFallback field.
func (o *ExperimentsBanditVariantInputV2) SetIsFallback(v bool) {
	o.IsFallback = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditVariantInputV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["feature_flag_variant_id"] = o.FeatureFlagVariantId
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.IsFallback != nil {
		toSerialize["is_fallback"] = o.IsFallback
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditVariantInputV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FeatureFlagVariantId *uuid.UUID `json:"feature_flag_variant_id"`
		IsActive             *bool      `json:"is_active,omitempty"`
		IsFallback           *bool      `json:"is_fallback,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FeatureFlagVariantId == nil {
		return fmt.Errorf("required field feature_flag_variant_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"feature_flag_variant_id", "is_active", "is_fallback"})
	} else {
		return err
	}
	o.FeatureFlagVariantId = *all.FeatureFlagVariantId
	o.IsActive = all.IsActive
	o.IsFallback = all.IsFallback

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
