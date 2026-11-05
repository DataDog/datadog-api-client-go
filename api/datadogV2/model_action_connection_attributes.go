// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionAttributes The definition of `ActionConnectionAttributes` object.
type ActionConnectionAttributes struct {
	// AWS Secrets Manager settings shared by the credential references on a `DATACENTER` connection. These settings are required whenever a field uses an AWS Secrets Manager reference.
	ExternalSecretsManager NullableActionConnectionAWSSecretsManager `json:"external_secrets_manager,omitempty"`
	// The definition of `ActionConnectionIntegration` object.
	Integration ActionConnectionIntegration `json:"integration"`
	// Name of the connection
	Name string `json:"name"`
	// Tags associated with the connection. Each tag must follow the `key:value` format.
	// The `default` tag key is reserved.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionConnectionAttributes instantiates a new ActionConnectionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionConnectionAttributes(integration ActionConnectionIntegration, name string) *ActionConnectionAttributes {
	this := ActionConnectionAttributes{}
	this.Integration = integration
	this.Name = name
	return &this
}

// NewActionConnectionAttributesWithDefaults instantiates a new ActionConnectionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionConnectionAttributesWithDefaults() *ActionConnectionAttributes {
	this := ActionConnectionAttributes{}
	return &this
}

// GetExternalSecretsManager returns the ExternalSecretsManager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionConnectionAttributes) GetExternalSecretsManager() ActionConnectionAWSSecretsManager {
	if o == nil || o.ExternalSecretsManager.Get() == nil {
		var ret ActionConnectionAWSSecretsManager
		return ret
	}
	return *o.ExternalSecretsManager.Get()
}

// GetExternalSecretsManagerOk returns a tuple with the ExternalSecretsManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ActionConnectionAttributes) GetExternalSecretsManagerOk() (*ActionConnectionAWSSecretsManager, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalSecretsManager.Get(), o.ExternalSecretsManager.IsSet()
}

// HasExternalSecretsManager returns a boolean if a field has been set.
func (o *ActionConnectionAttributes) HasExternalSecretsManager() bool {
	return o != nil && o.ExternalSecretsManager.IsSet()
}

// SetExternalSecretsManager gets a reference to the given NullableActionConnectionAWSSecretsManager and assigns it to the ExternalSecretsManager field.
func (o *ActionConnectionAttributes) SetExternalSecretsManager(v ActionConnectionAWSSecretsManager) {
	o.ExternalSecretsManager.Set(&v)
}

// SetExternalSecretsManagerNil sets the value for ExternalSecretsManager to be an explicit nil.
func (o *ActionConnectionAttributes) SetExternalSecretsManagerNil() {
	o.ExternalSecretsManager.Set(nil)
}

// UnsetExternalSecretsManager ensures that no value is present for ExternalSecretsManager, not even an explicit nil.
func (o *ActionConnectionAttributes) UnsetExternalSecretsManager() {
	o.ExternalSecretsManager.Unset()
}

// GetIntegration returns the Integration field value.
func (o *ActionConnectionAttributes) GetIntegration() ActionConnectionIntegration {
	if o == nil {
		var ret ActionConnectionIntegration
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *ActionConnectionAttributes) GetIntegrationOk() (*ActionConnectionIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *ActionConnectionAttributes) SetIntegration(v ActionConnectionIntegration) {
	o.Integration = v
}

// GetName returns the Name field value.
func (o *ActionConnectionAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActionConnectionAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ActionConnectionAttributes) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ActionConnectionAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionConnectionAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ActionConnectionAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ActionConnectionAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionConnectionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExternalSecretsManager.IsSet() {
		toSerialize["external_secrets_manager"] = o.ExternalSecretsManager.Get()
	}
	toSerialize["integration"] = o.Integration
	toSerialize["name"] = o.Name
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionConnectionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExternalSecretsManager NullableActionConnectionAWSSecretsManager `json:"external_secrets_manager,omitempty"`
		Integration            *ActionConnectionIntegration              `json:"integration"`
		Name                   *string                                   `json:"name"`
		Tags                   []string                                  `json:"tags,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"external_secrets_manager", "integration", "name", "tags"})
	} else {
		return err
	}
	o.ExternalSecretsManager = all.ExternalSecretsManager
	o.Integration = *all.Integration
	o.Name = *all.Name
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
