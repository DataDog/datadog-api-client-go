// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageScheduledActionRelationships The relationships of a scheduled action.
type OnCallPageScheduledActionRelationships struct {
	// The schedules the action escalates to.
	Schedules *OnCallPageScheduledActionRelationshipsSchedules `json:"schedules,omitempty"`
	// The users the action escalates to.
	Users *OnCallPageScheduledActionRelationshipsUsers `json:"users,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageScheduledActionRelationships instantiates a new OnCallPageScheduledActionRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageScheduledActionRelationships() *OnCallPageScheduledActionRelationships {
	this := OnCallPageScheduledActionRelationships{}
	return &this
}

// NewOnCallPageScheduledActionRelationshipsWithDefaults instantiates a new OnCallPageScheduledActionRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageScheduledActionRelationshipsWithDefaults() *OnCallPageScheduledActionRelationships {
	this := OnCallPageScheduledActionRelationships{}
	return &this
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *OnCallPageScheduledActionRelationships) GetSchedules() OnCallPageScheduledActionRelationshipsSchedules {
	if o == nil || o.Schedules == nil {
		var ret OnCallPageScheduledActionRelationshipsSchedules
		return ret
	}
	return *o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageScheduledActionRelationships) GetSchedulesOk() (*OnCallPageScheduledActionRelationshipsSchedules, bool) {
	if o == nil || o.Schedules == nil {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *OnCallPageScheduledActionRelationships) HasSchedules() bool {
	return o != nil && o.Schedules != nil
}

// SetSchedules gets a reference to the given OnCallPageScheduledActionRelationshipsSchedules and assigns it to the Schedules field.
func (o *OnCallPageScheduledActionRelationships) SetSchedules(v OnCallPageScheduledActionRelationshipsSchedules) {
	o.Schedules = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *OnCallPageScheduledActionRelationships) GetUsers() OnCallPageScheduledActionRelationshipsUsers {
	if o == nil || o.Users == nil {
		var ret OnCallPageScheduledActionRelationshipsUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageScheduledActionRelationships) GetUsersOk() (*OnCallPageScheduledActionRelationshipsUsers, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *OnCallPageScheduledActionRelationships) HasUsers() bool {
	return o != nil && o.Users != nil
}

// SetUsers gets a reference to the given OnCallPageScheduledActionRelationshipsUsers and assigns it to the Users field.
func (o *OnCallPageScheduledActionRelationships) SetUsers(v OnCallPageScheduledActionRelationshipsUsers) {
	o.Users = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageScheduledActionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Schedules != nil {
		toSerialize["schedules"] = o.Schedules
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageScheduledActionRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Schedules *OnCallPageScheduledActionRelationshipsSchedules `json:"schedules,omitempty"`
		Users     *OnCallPageScheduledActionRelationshipsUsers     `json:"users,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"schedules", "users"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Schedules != nil && all.Schedules.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schedules = all.Schedules
	if all.Users != nil && all.Users.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Users = all.Users

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
