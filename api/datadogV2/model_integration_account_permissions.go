// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountPermissions Read-only permission information for the account, derived from its restriction policy.
type IntegrationAccountPermissions struct {
	// Restriction-policy resource identifier of this account.
	ResourceId *string `json:"resource_id,omitempty"`
	// Whether the requesting user may edit this account.
	UserCanEdit *bool `json:"user_can_edit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationAccountPermissions instantiates a new IntegrationAccountPermissions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationAccountPermissions() *IntegrationAccountPermissions {
	this := IntegrationAccountPermissions{}
	return &this
}

// NewIntegrationAccountPermissionsWithDefaults instantiates a new IntegrationAccountPermissions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationAccountPermissionsWithDefaults() *IntegrationAccountPermissions {
	this := IntegrationAccountPermissions{}
	return &this
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *IntegrationAccountPermissions) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAccountPermissions) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *IntegrationAccountPermissions) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *IntegrationAccountPermissions) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetUserCanEdit returns the UserCanEdit field value if set, zero value otherwise.
func (o *IntegrationAccountPermissions) GetUserCanEdit() bool {
	if o == nil || o.UserCanEdit == nil {
		var ret bool
		return ret
	}
	return *o.UserCanEdit
}

// GetUserCanEditOk returns a tuple with the UserCanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAccountPermissions) GetUserCanEditOk() (*bool, bool) {
	if o == nil || o.UserCanEdit == nil {
		return nil, false
	}
	return o.UserCanEdit, true
}

// HasUserCanEdit returns a boolean if a field has been set.
func (o *IntegrationAccountPermissions) HasUserCanEdit() bool {
	return o != nil && o.UserCanEdit != nil
}

// SetUserCanEdit gets a reference to the given bool and assigns it to the UserCanEdit field.
func (o *IntegrationAccountPermissions) SetUserCanEdit(v bool) {
	o.UserCanEdit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationAccountPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.UserCanEdit != nil {
		toSerialize["user_can_edit"] = o.UserCanEdit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationAccountPermissions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResourceId  *string `json:"resource_id,omitempty"`
		UserCanEdit *bool   `json:"user_can_edit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"resource_id", "user_can_edit"})
	} else {
		return err
	}
	o.ResourceId = all.ResourceId
	o.UserCanEdit = all.UserCanEdit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
