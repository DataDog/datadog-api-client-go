// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountAttributes Attributes of an integration account. The `integration` field is a strongly-typed, per-integration union.
type IntegrationAccountAttributes struct {
	// Strongly-typed, per-integration configuration. Exactly one integration variant is set, selected by its `type`.
	Integration IntegrationAccountIntegration `json:"integration"`
	// Human-readable name of the account.
	Name string `json:"name"`
	// Read-only permission information for the account, derived from its restriction policy.
	Permissions *IntegrationAccountPermissions `json:"permissions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationAccountAttributes instantiates a new IntegrationAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationAccountAttributes(integration IntegrationAccountIntegration, name string) *IntegrationAccountAttributes {
	this := IntegrationAccountAttributes{}
	this.Integration = integration
	this.Name = name
	return &this
}

// NewIntegrationAccountAttributesWithDefaults instantiates a new IntegrationAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationAccountAttributesWithDefaults() *IntegrationAccountAttributes {
	this := IntegrationAccountAttributes{}
	return &this
}

// GetIntegration returns the Integration field value.
func (o *IntegrationAccountAttributes) GetIntegration() IntegrationAccountIntegration {
	if o == nil {
		var ret IntegrationAccountIntegration
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *IntegrationAccountAttributes) GetIntegrationOk() (*IntegrationAccountIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *IntegrationAccountAttributes) SetIntegration(v IntegrationAccountIntegration) {
	o.Integration = v
}

// GetName returns the Name field value.
func (o *IntegrationAccountAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntegrationAccountAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IntegrationAccountAttributes) SetName(v string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *IntegrationAccountAttributes) GetPermissions() IntegrationAccountPermissions {
	if o == nil || o.Permissions == nil {
		var ret IntegrationAccountPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAccountAttributes) GetPermissionsOk() (*IntegrationAccountPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IntegrationAccountAttributes) HasPermissions() bool {
	return o != nil && o.Permissions != nil
}

// SetPermissions gets a reference to the given IntegrationAccountPermissions and assigns it to the Permissions field.
func (o *IntegrationAccountAttributes) SetPermissions(v IntegrationAccountPermissions) {
	o.Permissions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["integration"] = o.Integration
	toSerialize["name"] = o.Name
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Integration *IntegrationAccountIntegration `json:"integration"`
		Name        *string                        `json:"name"`
		Permissions *IntegrationAccountPermissions `json:"permissions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Integration == nil {
		return fmt.Errorf("required field integration missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"integration", "name", "permissions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Integration = *all.Integration
	o.Name = *all.Name
	if all.Permissions != nil && all.Permissions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Permissions = all.Permissions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
