// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationAccountAttributes Attributes of an Elastic Cloud integration account.
type ElasticCloudIntegrationAccountAttributes struct {
	// Elastic Cloud interface (source-type). Exactly one interface variant is set, selected by its `type`.
	Interface ElasticCloudInterface `json:"interface"`
	// Human-readable name of the account.
	Name string `json:"name"`
	// Read-only permission information for the account, derived from its restriction policy.
	Permissions *IntegrationAccountPermissions `json:"permissions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudIntegrationAccountAttributes instantiates a new ElasticCloudIntegrationAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudIntegrationAccountAttributes(interfaceVar ElasticCloudInterface, name string) *ElasticCloudIntegrationAccountAttributes {
	this := ElasticCloudIntegrationAccountAttributes{}
	this.Interface = interfaceVar
	this.Name = name
	return &this
}

// NewElasticCloudIntegrationAccountAttributesWithDefaults instantiates a new ElasticCloudIntegrationAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudIntegrationAccountAttributesWithDefaults() *ElasticCloudIntegrationAccountAttributes {
	this := ElasticCloudIntegrationAccountAttributes{}
	return &this
}

// GetInterface returns the Interface field value.
func (o *ElasticCloudIntegrationAccountAttributes) GetInterface() ElasticCloudInterface {
	if o == nil {
		var ret ElasticCloudInterface
		return ret
	}
	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationAccountAttributes) GetInterfaceOk() (*ElasticCloudInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value.
func (o *ElasticCloudIntegrationAccountAttributes) SetInterface(v ElasticCloudInterface) {
	o.Interface = v
}

// GetName returns the Name field value.
func (o *ElasticCloudIntegrationAccountAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationAccountAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ElasticCloudIntegrationAccountAttributes) SetName(v string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationAccountAttributes) GetPermissions() IntegrationAccountPermissions {
	if o == nil || o.Permissions == nil {
		var ret IntegrationAccountPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationAccountAttributes) GetPermissionsOk() (*IntegrationAccountPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationAccountAttributes) HasPermissions() bool {
	return o != nil && o.Permissions != nil
}

// SetPermissions gets a reference to the given IntegrationAccountPermissions and assigns it to the Permissions field.
func (o *ElasticCloudIntegrationAccountAttributes) SetPermissions(v IntegrationAccountPermissions) {
	o.Permissions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudIntegrationAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["interface"] = o.Interface
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
func (o *ElasticCloudIntegrationAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interface   *ElasticCloudInterface         `json:"interface"`
		Name        *string                        `json:"name"`
		Permissions *IntegrationAccountPermissions `json:"permissions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Interface == nil {
		return fmt.Errorf("required field interface missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interface", "name", "permissions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Interface = *all.Interface
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
