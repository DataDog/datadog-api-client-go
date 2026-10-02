// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentV2DTODataAttributesVariantsItems Variant in an experiment, with its identity and traffic allocation.
type ExperimentsExperimentV2DTODataAttributesVariantsItems struct {
	// Backing feature flag variant ID. Present for Datadog feature flag experiments and omitted for Warehouse experiments.
	FeatureFlagVariantId *string `json:"feature_flag_variant_id,omitempty"`
	// Whether this variant participates in the experiment.
	IsActive bool `json:"is_active"`
	// Whether this is the single control variant.
	IsControl bool `json:"is_control"`
	// Value recorded in exposure data for this variant.
	Key string `json:"key"`
	// Display name of the experiment variant.
	Name datadog.NullableString `json:"name,omitempty"`
	// Traffic allocation percentage.
	Weight float64 `json:"weight"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentV2DTODataAttributesVariantsItems instantiates a new ExperimentsExperimentV2DTODataAttributesVariantsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentV2DTODataAttributesVariantsItems(isActive bool, isControl bool, key string, weight float64) *ExperimentsExperimentV2DTODataAttributesVariantsItems {
	this := ExperimentsExperimentV2DTODataAttributesVariantsItems{}
	this.IsActive = isActive
	this.IsControl = isControl
	this.Key = key
	this.Weight = weight
	return &this
}

// NewExperimentsExperimentV2DTODataAttributesVariantsItemsWithDefaults instantiates a new ExperimentsExperimentV2DTODataAttributesVariantsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentV2DTODataAttributesVariantsItemsWithDefaults() *ExperimentsExperimentV2DTODataAttributesVariantsItems {
	this := ExperimentsExperimentV2DTODataAttributesVariantsItems{}
	return &this
}

// GetFeatureFlagVariantId returns the FeatureFlagVariantId field value if set, zero value otherwise.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetFeatureFlagVariantId() string {
	if o == nil || o.FeatureFlagVariantId == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlagVariantId
}

// GetFeatureFlagVariantIdOk returns a tuple with the FeatureFlagVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetFeatureFlagVariantIdOk() (*string, bool) {
	if o == nil || o.FeatureFlagVariantId == nil {
		return nil, false
	}
	return o.FeatureFlagVariantId, true
}

// HasFeatureFlagVariantId returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) HasFeatureFlagVariantId() bool {
	return o != nil && o.FeatureFlagVariantId != nil
}

// SetFeatureFlagVariantId gets a reference to the given string and assigns it to the FeatureFlagVariantId field.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) SetFeatureFlagVariantId(v string) {
	o.FeatureFlagVariantId = &v
}

// GetIsActive returns the IsActive field value.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsControl returns the IsControl field value.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetIsControl() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsControl
}

// GetIsControlOk returns a tuple with the IsControl field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetIsControlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsControl, true
}

// SetIsControl sets field value.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) SetIsControl(v bool) {
	o.IsControl = v
}

// GetKey returns the Key field value.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) UnsetName() {
	o.Name.Unset()
}

// GetWeight returns the Weight field value.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetWeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value.
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) SetWeight(v float64) {
	o.Weight = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentV2DTODataAttributesVariantsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FeatureFlagVariantId != nil {
		toSerialize["feature_flag_variant_id"] = o.FeatureFlagVariantId
	}
	toSerialize["is_active"] = o.IsActive
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
func (o *ExperimentsExperimentV2DTODataAttributesVariantsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FeatureFlagVariantId *string                `json:"feature_flag_variant_id,omitempty"`
		IsActive             *bool                  `json:"is_active"`
		IsControl            *bool                  `json:"is_control"`
		Key                  *string                `json:"key"`
		Name                 datadog.NullableString `json:"name,omitempty"`
		Weight               *float64               `json:"weight"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsActive == nil {
		return fmt.Errorf("required field is_active missing")
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
	o.IsActive = *all.IsActive
	o.IsControl = *all.IsControl
	o.Key = *all.Key
	o.Name = all.Name
	o.Weight = *all.Weight

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
