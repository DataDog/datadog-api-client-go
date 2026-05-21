// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRoleAssignmentDataAttributesRequest Attributes for creating a role assignment.
type IncidentRoleAssignmentDataAttributesRequest struct {
	// The name of the role to assign.
	Role datadog.NullableString `json:"role,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRoleAssignmentDataAttributesRequest instantiates a new IncidentRoleAssignmentDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRoleAssignmentDataAttributesRequest() *IncidentRoleAssignmentDataAttributesRequest {
	this := IncidentRoleAssignmentDataAttributesRequest{}
	return &this
}

// NewIncidentRoleAssignmentDataAttributesRequestWithDefaults instantiates a new IncidentRoleAssignmentDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRoleAssignmentDataAttributesRequestWithDefaults() *IncidentRoleAssignmentDataAttributesRequest {
	this := IncidentRoleAssignmentDataAttributesRequest{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRoleAssignmentDataAttributesRequest) GetRole() string {
	if o == nil || o.Role.Get() == nil {
		var ret string
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRoleAssignmentDataAttributesRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *IncidentRoleAssignmentDataAttributesRequest) HasRole() bool {
	return o != nil && o.Role.IsSet()
}

// SetRole gets a reference to the given datadog.NullableString and assigns it to the Role field.
func (o *IncidentRoleAssignmentDataAttributesRequest) SetRole(v string) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil.
func (o *IncidentRoleAssignmentDataAttributesRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil.
func (o *IncidentRoleAssignmentDataAttributesRequest) UnsetRole() {
	o.Role.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRoleAssignmentDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRoleAssignmentDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role datadog.NullableString `json:"role,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"role"})
	} else {
		return err
	}
	o.Role = all.Role

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
