// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventFlag The feature flag the change applies to.
type FeatureFlagChangeEventFlag struct {
	// The ID of the feature flag.
	Id uuid.UUID `json:"id"`
	// The key of the feature flag.
	Key string `json:"key"`
	// The name of the feature flag.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagChangeEventFlag instantiates a new FeatureFlagChangeEventFlag object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagChangeEventFlag(id uuid.UUID, key string, name string) *FeatureFlagChangeEventFlag {
	this := FeatureFlagChangeEventFlag{}
	this.Id = id
	this.Key = key
	this.Name = name
	return &this
}

// NewFeatureFlagChangeEventFlagWithDefaults instantiates a new FeatureFlagChangeEventFlag object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagChangeEventFlagWithDefaults() *FeatureFlagChangeEventFlag {
	this := FeatureFlagChangeEventFlag{}
	return &this
}

// GetId returns the Id field value.
func (o *FeatureFlagChangeEventFlag) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventFlag) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *FeatureFlagChangeEventFlag) SetId(v uuid.UUID) {
	o.Id = v
}

// GetKey returns the Key field value.
func (o *FeatureFlagChangeEventFlag) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventFlag) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *FeatureFlagChangeEventFlag) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value.
func (o *FeatureFlagChangeEventFlag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventFlag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FeatureFlagChangeEventFlag) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagChangeEventFlag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagChangeEventFlag) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *uuid.UUID `json:"id"`
		Key  *string    `json:"key"`
		Name *string    `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "key", "name"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Key = *all.Key
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
