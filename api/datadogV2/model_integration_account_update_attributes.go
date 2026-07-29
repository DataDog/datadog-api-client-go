// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountUpdateAttributes Updatable attributes of an integration account. Every field is optional; only the fields provided are changed.
type IntegrationAccountUpdateAttributes struct {
	// Strongly-typed, per-integration partial configuration. Exactly one integration variant is set, selected by its `type`.
	Integration *IntegrationAccountIntegrationUpdate `json:"integration,omitempty"`
	// Human-readable name of the account.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationAccountUpdateAttributes instantiates a new IntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationAccountUpdateAttributes() *IntegrationAccountUpdateAttributes {
	this := IntegrationAccountUpdateAttributes{}
	return &this
}

// NewIntegrationAccountUpdateAttributesWithDefaults instantiates a new IntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationAccountUpdateAttributesWithDefaults() *IntegrationAccountUpdateAttributes {
	this := IntegrationAccountUpdateAttributes{}
	return &this
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *IntegrationAccountUpdateAttributes) GetIntegration() IntegrationAccountIntegrationUpdate {
	if o == nil || o.Integration == nil {
		var ret IntegrationAccountIntegrationUpdate
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAccountUpdateAttributes) GetIntegrationOk() (*IntegrationAccountIntegrationUpdate, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *IntegrationAccountUpdateAttributes) HasIntegration() bool {
	return o != nil && o.Integration != nil
}

// SetIntegration gets a reference to the given IntegrationAccountIntegrationUpdate and assigns it to the Integration field.
func (o *IntegrationAccountUpdateAttributes) SetIntegration(v IntegrationAccountIntegrationUpdate) {
	o.Integration = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Integration *IntegrationAccountIntegrationUpdate `json:"integration,omitempty"`
		Name        *string                              `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"integration", "name"})
	} else {
		return err
	}
	o.Integration = all.Integration
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
