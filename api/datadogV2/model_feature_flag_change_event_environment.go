// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventEnvironment The environment the change applies to. Null for a change to the flag itself rather than to one environment.
type FeatureFlagChangeEventEnvironment struct {
	// The ID of the environment.
	Id uuid.UUID `json:"id"`
	// Whether the environment is a production environment.
	IsProduction bool `json:"is_production"`
	// The name of the environment.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagChangeEventEnvironment instantiates a new FeatureFlagChangeEventEnvironment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagChangeEventEnvironment(id uuid.UUID, isProduction bool, name string) *FeatureFlagChangeEventEnvironment {
	this := FeatureFlagChangeEventEnvironment{}
	this.Id = id
	this.IsProduction = isProduction
	this.Name = name
	return &this
}

// NewFeatureFlagChangeEventEnvironmentWithDefaults instantiates a new FeatureFlagChangeEventEnvironment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagChangeEventEnvironmentWithDefaults() *FeatureFlagChangeEventEnvironment {
	this := FeatureFlagChangeEventEnvironment{}
	return &this
}

// GetId returns the Id field value.
func (o *FeatureFlagChangeEventEnvironment) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventEnvironment) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *FeatureFlagChangeEventEnvironment) SetId(v uuid.UUID) {
	o.Id = v
}

// GetIsProduction returns the IsProduction field value.
func (o *FeatureFlagChangeEventEnvironment) GetIsProduction() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventEnvironment) GetIsProductionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsProduction, true
}

// SetIsProduction sets field value.
func (o *FeatureFlagChangeEventEnvironment) SetIsProduction(v bool) {
	o.IsProduction = v
}

// GetName returns the Name field value.
func (o *FeatureFlagChangeEventEnvironment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventEnvironment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FeatureFlagChangeEventEnvironment) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagChangeEventEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["is_production"] = o.IsProduction
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagChangeEventEnvironment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id           *uuid.UUID `json:"id"`
		IsProduction *bool      `json:"is_production"`
		Name         *string    `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.IsProduction == nil {
		return fmt.Errorf("required field is_production missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "is_production", "name"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.IsProduction = *all.IsProduction
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableFeatureFlagChangeEventEnvironment handles when a null is used for FeatureFlagChangeEventEnvironment.
type NullableFeatureFlagChangeEventEnvironment struct {
	value *FeatureFlagChangeEventEnvironment
	isSet bool
}

// Get returns the associated value.
func (v NullableFeatureFlagChangeEventEnvironment) Get() *FeatureFlagChangeEventEnvironment {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableFeatureFlagChangeEventEnvironment) Set(val *FeatureFlagChangeEventEnvironment) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableFeatureFlagChangeEventEnvironment) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableFeatureFlagChangeEventEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFeatureFlagChangeEventEnvironment initializes the struct as if Set has been called.
func NewNullableFeatureFlagChangeEventEnvironment(val *FeatureFlagChangeEventEnvironment) *NullableFeatureFlagChangeEventEnvironment {
	return &NullableFeatureFlagChangeEventEnvironment{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableFeatureFlagChangeEventEnvironment) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableFeatureFlagChangeEventEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
