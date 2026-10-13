// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OverrideRelationships Relationships for an on-call schedule override.
type OverrideRelationships struct {
	// Defines the relationship between an override and one of its associated users.
	OverriddenUser *OverrideRelationshipsUser `json:"overridden_user,omitempty"`
	// Defines the relationship between an override and the schedule it belongs to.
	Schedule *OverrideRelationshipsSchedule `json:"schedule,omitempty"`
	// Defines the relationship between an override and one of its associated users.
	User *OverrideRelationshipsUser `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOverrideRelationships instantiates a new OverrideRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOverrideRelationships() *OverrideRelationships {
	this := OverrideRelationships{}
	return &this
}

// NewOverrideRelationshipsWithDefaults instantiates a new OverrideRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOverrideRelationshipsWithDefaults() *OverrideRelationships {
	this := OverrideRelationships{}
	return &this
}

// GetOverriddenUser returns the OverriddenUser field value if set, zero value otherwise.
func (o *OverrideRelationships) GetOverriddenUser() OverrideRelationshipsUser {
	if o == nil || o.OverriddenUser == nil {
		var ret OverrideRelationshipsUser
		return ret
	}
	return *o.OverriddenUser
}

// GetOverriddenUserOk returns a tuple with the OverriddenUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideRelationships) GetOverriddenUserOk() (*OverrideRelationshipsUser, bool) {
	if o == nil || o.OverriddenUser == nil {
		return nil, false
	}
	return o.OverriddenUser, true
}

// HasOverriddenUser returns a boolean if a field has been set.
func (o *OverrideRelationships) HasOverriddenUser() bool {
	return o != nil && o.OverriddenUser != nil
}

// SetOverriddenUser gets a reference to the given OverrideRelationshipsUser and assigns it to the OverriddenUser field.
func (o *OverrideRelationships) SetOverriddenUser(v OverrideRelationshipsUser) {
	o.OverriddenUser = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *OverrideRelationships) GetSchedule() OverrideRelationshipsSchedule {
	if o == nil || o.Schedule == nil {
		var ret OverrideRelationshipsSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideRelationships) GetScheduleOk() (*OverrideRelationshipsSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *OverrideRelationships) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given OverrideRelationshipsSchedule and assigns it to the Schedule field.
func (o *OverrideRelationships) SetSchedule(v OverrideRelationshipsSchedule) {
	o.Schedule = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OverrideRelationships) GetUser() OverrideRelationshipsUser {
	if o == nil || o.User == nil {
		var ret OverrideRelationshipsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideRelationships) GetUserOk() (*OverrideRelationshipsUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OverrideRelationships) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given OverrideRelationshipsUser and assigns it to the User field.
func (o *OverrideRelationships) SetUser(v OverrideRelationshipsUser) {
	o.User = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OverrideRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OverriddenUser != nil {
		toSerialize["overridden_user"] = o.OverriddenUser
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
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
func (o *OverrideRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OverriddenUser *OverrideRelationshipsUser     `json:"overridden_user,omitempty"`
		Schedule       *OverrideRelationshipsSchedule `json:"schedule,omitempty"`
		User           *OverrideRelationshipsUser     `json:"user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"overridden_user", "schedule", "user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OverriddenUser != nil && all.OverriddenUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OverriddenUser = all.OverriddenUser
	if all.Schedule != nil && all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schedule = all.Schedule
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
