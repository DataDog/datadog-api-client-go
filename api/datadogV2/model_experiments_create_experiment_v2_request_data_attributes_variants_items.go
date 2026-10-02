// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems Variant selected for an experiment, with its identity and traffic allocation.
type ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems struct {
	// Stable backing flag variant ID. Required for Datadog-backed experiments and omitted for Warehouse-only experiments.
	FeatureFlagVariantId *uuid.UUID `json:"feature_flag_variant_id,omitempty"`
	// Whether this variant participates in the experiment. Responses always include this field. On writes an included variant defaults to active; omit its row to remove or unselect it. Explicit false supports sending an unchanged response back.
	IsActive *bool `json:"is_active,omitempty"`
	// Whether this is the single control variant.
	IsControl bool `json:"is_control"`
	// Assignment value. Datadog flag variant keys are server-owned.
	Key string `json:"key"`
	// Display name. Datadog flag variant names are server-owned.
	Name datadog.NullableString `json:"name,omitempty"`
	// Traffic allocation percentage. Existing variant weights cannot change through the public API after the experiment starts.
	Weight float64 `json:"weight"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentV2RequestDataAttributesVariantsItems instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentV2RequestDataAttributesVariantsItems(isControl bool, key string, weight float64) *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems{}
	this.IsControl = isControl
	this.Key = key
	this.Weight = weight
	return &this
}

// NewExperimentsCreateExperimentV2RequestDataAttributesVariantsItemsWithDefaults instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentV2RequestDataAttributesVariantsItemsWithDefaults() *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems{}
	return &this
}

// GetFeatureFlagVariantId returns the FeatureFlagVariantId field value if set, zero value otherwise.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetFeatureFlagVariantId() uuid.UUID {
	if o == nil || o.FeatureFlagVariantId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.FeatureFlagVariantId
}

// GetFeatureFlagVariantIdOk returns a tuple with the FeatureFlagVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetFeatureFlagVariantIdOk() (*uuid.UUID, bool) {
	if o == nil || o.FeatureFlagVariantId == nil {
		return nil, false
	}
	return o.FeatureFlagVariantId, true
}

// HasFeatureFlagVariantId returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) HasFeatureFlagVariantId() bool {
	return o != nil && o.FeatureFlagVariantId != nil
}

// SetFeatureFlagVariantId gets a reference to the given uuid.UUID and assigns it to the FeatureFlagVariantId field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) SetFeatureFlagVariantId(v uuid.UUID) {
	o.FeatureFlagVariantId = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) HasIsActive() bool {
	return o != nil && o.IsActive != nil
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsControl returns the IsControl field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetIsControl() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsControl
}

// GetIsControlOk returns a tuple with the IsControl field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetIsControlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsControl, true
}

// SetIsControl sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) SetIsControl(v bool) {
	o.IsControl = v
}

// GetKey returns the Key field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) UnsetName() {
	o.Name.Unset()
}

// GetWeight returns the Weight field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetWeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) SetWeight(v float64) {
	o.Weight = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FeatureFlagVariantId != nil {
		toSerialize["feature_flag_variant_id"] = o.FeatureFlagVariantId
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	toSerialize["is_control"] = o.IsControl
	toSerialize["key"] = o.Key
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["weight"] = o.Weight

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesVariantsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FeatureFlagVariantId *uuid.UUID             `json:"feature_flag_variant_id,omitempty"`
		IsActive             *bool                  `json:"is_active,omitempty"`
		IsControl            *bool                  `json:"is_control"`
		Key                  *string                `json:"key"`
		Name                 datadog.NullableString `json:"name,omitempty"`
		Weight               *float64               `json:"weight"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsControl == nil {
		return fmt.Errorf("required field is_control missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Weight == nil {
		return fmt.Errorf("required field weight missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"feature_flag_variant_id", "is_active", "is_control", "key", "name", "weight"})
	} else {
		return err
	}
	o.FeatureFlagVariantId = all.FeatureFlagVariantId
	o.IsActive = all.IsActive
	o.IsControl = *all.IsControl
	o.Key = *all.Key
	o.Name = all.Name
	o.Weight = *all.Weight

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
