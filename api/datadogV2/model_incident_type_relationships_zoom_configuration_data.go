// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTypeRelationshipsZoomConfigurationData The Zoom configuration relationship data object.
type IncidentTypeRelationshipsZoomConfigurationData struct {
	// The unique identifier of the Zoom configuration.
	Id string `json:"id"`
	// The type of the Zoom configuration.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTypeRelationshipsZoomConfigurationData instantiates a new IncidentTypeRelationshipsZoomConfigurationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTypeRelationshipsZoomConfigurationData(id string, typeVar string) *IncidentTypeRelationshipsZoomConfigurationData {
	this := IncidentTypeRelationshipsZoomConfigurationData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentTypeRelationshipsZoomConfigurationDataWithDefaults instantiates a new IncidentTypeRelationshipsZoomConfigurationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTypeRelationshipsZoomConfigurationDataWithDefaults() *IncidentTypeRelationshipsZoomConfigurationData {
	this := IncidentTypeRelationshipsZoomConfigurationData{}
	return &this
}

// GetId returns the Id field value.
func (o *IncidentTypeRelationshipsZoomConfigurationData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentTypeRelationshipsZoomConfigurationData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentTypeRelationshipsZoomConfigurationData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *IncidentTypeRelationshipsZoomConfigurationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentTypeRelationshipsZoomConfigurationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentTypeRelationshipsZoomConfigurationData) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTypeRelationshipsZoomConfigurationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTypeRelationshipsZoomConfigurationData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string `json:"id"`
		Type *string `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableIncidentTypeRelationshipsZoomConfigurationData handles when a null is used for IncidentTypeRelationshipsZoomConfigurationData.
type NullableIncidentTypeRelationshipsZoomConfigurationData struct {
	value *IncidentTypeRelationshipsZoomConfigurationData
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentTypeRelationshipsZoomConfigurationData) Get() *IncidentTypeRelationshipsZoomConfigurationData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentTypeRelationshipsZoomConfigurationData) Set(val *IncidentTypeRelationshipsZoomConfigurationData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentTypeRelationshipsZoomConfigurationData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableIncidentTypeRelationshipsZoomConfigurationData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentTypeRelationshipsZoomConfigurationData initializes the struct as if Set has been called.
func NewNullableIncidentTypeRelationshipsZoomConfigurationData(val *IncidentTypeRelationshipsZoomConfigurationData) *NullableIncidentTypeRelationshipsZoomConfigurationData {
	return &NullableIncidentTypeRelationshipsZoomConfigurationData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentTypeRelationshipsZoomConfigurationData) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentTypeRelationshipsZoomConfigurationData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
