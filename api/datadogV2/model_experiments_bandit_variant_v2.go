// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditVariantV2 A flag variant the bandit uses.
type ExperimentsBanditVariantV2 struct {
	// The variant's ID on the flag.
	FeatureFlagVariantId uuid.UUID `json:"feature_flag_variant_id"`
	// Whether the variant is one of the arms the bandit chooses between.
	IsActive bool `json:"is_active"`
	// Whether the variant is served when the bandit reaches its scheduled end. At most one variant is the fallback.
	IsFallback bool `json:"is_fallback"`
	// The variant's key. An arm keeps the key it had when it was added, because exposures are matched on it. Null only for a fallback deleted from the flag.
	Key datadog.NullableString `json:"key"`
	// The variant's name on the flag. Null when it is no longer on the flag.
	Name datadog.NullableString `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditVariantV2 instantiates a new ExperimentsBanditVariantV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditVariantV2(featureFlagVariantId uuid.UUID, isActive bool, isFallback bool, key datadog.NullableString, name datadog.NullableString) *ExperimentsBanditVariantV2 {
	this := ExperimentsBanditVariantV2{}
	this.FeatureFlagVariantId = featureFlagVariantId
	this.IsActive = isActive
	this.IsFallback = isFallback
	this.Key = key
	this.Name = name
	return &this
}

// NewExperimentsBanditVariantV2WithDefaults instantiates a new ExperimentsBanditVariantV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditVariantV2WithDefaults() *ExperimentsBanditVariantV2 {
	this := ExperimentsBanditVariantV2{}
	return &this
}

// GetFeatureFlagVariantId returns the FeatureFlagVariantId field value.
func (o *ExperimentsBanditVariantV2) GetFeatureFlagVariantId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagVariantId
}

// GetFeatureFlagVariantIdOk returns a tuple with the FeatureFlagVariantId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantV2) GetFeatureFlagVariantIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagVariantId, true
}

// SetFeatureFlagVariantId sets field value.
func (o *ExperimentsBanditVariantV2) SetFeatureFlagVariantId(v uuid.UUID) {
	o.FeatureFlagVariantId = v
}

// GetIsActive returns the IsActive field value.
func (o *ExperimentsBanditVariantV2) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantV2) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value.
func (o *ExperimentsBanditVariantV2) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsFallback returns the IsFallback field value.
func (o *ExperimentsBanditVariantV2) GetIsFallback() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsFallback
}

// GetIsFallbackOk returns a tuple with the IsFallback field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditVariantV2) GetIsFallbackOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFallback, true
}

// SetIsFallback sets field value.
func (o *ExperimentsBanditVariantV2) SetIsFallback(v bool) {
	o.IsFallback = v
}

// GetKey returns the Key field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditVariantV2) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditVariantV2) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value.
func (o *ExperimentsBanditVariantV2) SetKey(v string) {
	o.Key.Set(&v)
}

// GetName returns the Name field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ExperimentsBanditVariantV2) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditVariantV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value.
func (o *ExperimentsBanditVariantV2) SetName(v string) {
	o.Name.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditVariantV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["feature_flag_variant_id"] = o.FeatureFlagVariantId
	toSerialize["is_active"] = o.IsActive
	toSerialize["is_fallback"] = o.IsFallback
	toSerialize["key"] = o.Key.Get()
	toSerialize["name"] = o.Name.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditVariantV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FeatureFlagVariantId *uuid.UUID             `json:"feature_flag_variant_id"`
		IsActive             *bool                  `json:"is_active"`
		IsFallback           *bool                  `json:"is_fallback"`
		Key                  datadog.NullableString `json:"key"`
		Name                 datadog.NullableString `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FeatureFlagVariantId == nil {
		return fmt.Errorf("required field feature_flag_variant_id missing")
	}
	if all.IsActive == nil {
		return fmt.Errorf("required field is_active missing")
	}
	if all.IsFallback == nil {
		return fmt.Errorf("required field is_fallback missing")
	}
	if !all.Key.IsSet() {
		return fmt.Errorf("required field key missing")
	}
	if !all.Name.IsSet() {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"feature_flag_variant_id", "is_active", "is_fallback", "key", "name"})
	} else {
		return err
	}
	o.FeatureFlagVariantId = *all.FeatureFlagVariantId
	o.IsActive = *all.IsActive
	o.IsFallback = *all.IsFallback
	o.Key = all.Key
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
