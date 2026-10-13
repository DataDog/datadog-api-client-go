// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateOverrideRequestRelationships Relationships to set when creating an on-call schedule override.
type CreateOverrideRequestRelationships struct {
	// Defines the relationship between an override and one of its associated users.
	OverriddenUser *OverrideRelationshipsUser `json:"overridden_user,omitempty"`
	// Defines the relationship between an override and one of its associated users.
	User *OverrideRelationshipsUser `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateOverrideRequestRelationships instantiates a new CreateOverrideRequestRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateOverrideRequestRelationships() *CreateOverrideRequestRelationships {
	this := CreateOverrideRequestRelationships{}
	return &this
}

// NewCreateOverrideRequestRelationshipsWithDefaults instantiates a new CreateOverrideRequestRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateOverrideRequestRelationshipsWithDefaults() *CreateOverrideRequestRelationships {
	this := CreateOverrideRequestRelationships{}
	return &this
}

// GetOverriddenUser returns the OverriddenUser field value if set, zero value otherwise.
func (o *CreateOverrideRequestRelationships) GetOverriddenUser() OverrideRelationshipsUser {
	if o == nil || o.OverriddenUser == nil {
		var ret OverrideRelationshipsUser
		return ret
	}
	return *o.OverriddenUser
}

// GetOverriddenUserOk returns a tuple with the OverriddenUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOverrideRequestRelationships) GetOverriddenUserOk() (*OverrideRelationshipsUser, bool) {
	if o == nil || o.OverriddenUser == nil {
		return nil, false
	}
	return o.OverriddenUser, true
}

// HasOverriddenUser returns a boolean if a field has been set.
func (o *CreateOverrideRequestRelationships) HasOverriddenUser() bool {
	return o != nil && o.OverriddenUser != nil
}

// SetOverriddenUser gets a reference to the given OverrideRelationshipsUser and assigns it to the OverriddenUser field.
func (o *CreateOverrideRequestRelationships) SetOverriddenUser(v OverrideRelationshipsUser) {
	o.OverriddenUser = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateOverrideRequestRelationships) GetUser() OverrideRelationshipsUser {
	if o == nil || o.User == nil {
		var ret OverrideRelationshipsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOverrideRequestRelationships) GetUserOk() (*OverrideRelationshipsUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateOverrideRequestRelationships) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given OverrideRelationshipsUser and assigns it to the User field.
func (o *CreateOverrideRequestRelationships) SetUser(v OverrideRelationshipsUser) {
	o.User = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateOverrideRequestRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OverriddenUser != nil {
		toSerialize["overridden_user"] = o.OverriddenUser
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateOverrideRequestRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OverriddenUser *OverrideRelationshipsUser `json:"overridden_user,omitempty"`
		User           *OverrideRelationshipsUser `json:"user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"overridden_user", "user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OverriddenUser != nil && all.OverriddenUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OverriddenUser = all.OverriddenUser
	if all.User != nil && all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.User = all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
