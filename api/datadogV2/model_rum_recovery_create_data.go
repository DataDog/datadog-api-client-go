// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryCreateData JSON:API resource.
type RumRecoveryCreateData struct {
	// Targeted recovery (the default) or query-based batch recovery.
	Attributes RumRecoveryCreateAttributes `json:"attributes"`
	// JSON:API resource type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoveryCreateData instantiates a new RumRecoveryCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoveryCreateData(attributes RumRecoveryCreateAttributes) *RumRecoveryCreateData {
	this := RumRecoveryCreateData{}
	this.Attributes = attributes
	return &this
}

// NewRumRecoveryCreateDataWithDefaults instantiates a new RumRecoveryCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoveryCreateDataWithDefaults() *RumRecoveryCreateData {
	this := RumRecoveryCreateData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *RumRecoveryCreateData) GetAttributes() RumRecoveryCreateAttributes {
	if o == nil {
		var ret RumRecoveryCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryCreateData) GetAttributesOk() (*RumRecoveryCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *RumRecoveryCreateData) SetAttributes(v RumRecoveryCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RumRecoveryCreateData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryCreateData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RumRecoveryCreateData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RumRecoveryCreateData) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoveryCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRecoveryCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RumRecoveryCreateAttributes `json:"attributes"`
		Type       *string                      `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}
	o.Attributes = *all.Attributes
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
