// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedRoleDataAttributesRequest Attributes for creating an incident user-defined role.
type IncidentUserDefinedRoleDataAttributesRequest struct {
	// A description of the user-defined role.
	Description datadog.NullableString `json:"description,omitempty"`
	// The name of the user-defined role.
	Name string `json:"name"`
	// Policy configuration for a user-defined role.
	Policy *IncidentUserDefinedRolePolicy `json:"policy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedRoleDataAttributesRequest instantiates a new IncidentUserDefinedRoleDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedRoleDataAttributesRequest(name string) *IncidentUserDefinedRoleDataAttributesRequest {
	this := IncidentUserDefinedRoleDataAttributesRequest{}
	this.Name = name
	return &this
}

// NewIncidentUserDefinedRoleDataAttributesRequestWithDefaults instantiates a new IncidentUserDefinedRoleDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedRoleDataAttributesRequestWithDefaults() *IncidentUserDefinedRoleDataAttributesRequest {
	this := IncidentUserDefinedRoleDataAttributesRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedRoleDataAttributesRequest) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedRoleDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IncidentUserDefinedRoleDataAttributesRequest) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *IncidentUserDefinedRoleDataAttributesRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *IncidentUserDefinedRoleDataAttributesRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *IncidentUserDefinedRoleDataAttributesRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value.
func (o *IncidentUserDefinedRoleDataAttributesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentUserDefinedRoleDataAttributesRequest) SetName(v string) {
	o.Name = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *IncidentUserDefinedRoleDataAttributesRequest) GetPolicy() IncidentUserDefinedRolePolicy {
	if o == nil || o.Policy == nil {
		var ret IncidentUserDefinedRolePolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataAttributesRequest) GetPolicyOk() (*IncidentUserDefinedRolePolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *IncidentUserDefinedRoleDataAttributesRequest) HasPolicy() bool {
	return o != nil && o.Policy != nil
}

// SetPolicy gets a reference to the given IncidentUserDefinedRolePolicy and assigns it to the Policy field.
func (o *IncidentUserDefinedRoleDataAttributesRequest) SetPolicy(v IncidentUserDefinedRolePolicy) {
	o.Policy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedRoleDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedRoleDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description datadog.NullableString         `json:"description,omitempty"`
		Name        *string                        `json:"name"`
		Policy      *IncidentUserDefinedRolePolicy `json:"policy,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "policy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Name = *all.Name
	if all.Policy != nil && all.Policy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Policy = all.Policy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
