// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeatUserDataAttributes
type SeatUserDataAttributes struct {
	// The date and time the seat was assigned.
	AssignedAt *time.Time `json:"assigned_at,omitempty"`
	// The email of the user.
	Email *string `json:"email,omitempty"`
	// The name of the user.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeatUserDataAttributes instantiates a new SeatUserDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeatUserDataAttributes() *SeatUserDataAttributes {
	this := SeatUserDataAttributes{}
	return &this
}

// NewSeatUserDataAttributesWithDefaults instantiates a new SeatUserDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeatUserDataAttributesWithDefaults() *SeatUserDataAttributes {
	this := SeatUserDataAttributes{}
	return &this
}

// GetAssignedAt returns the AssignedAt field value if set, zero value otherwise.
func (o *SeatUserDataAttributes) GetAssignedAt() time.Time {
	if o == nil || o.AssignedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignedAt
}

// GetAssignedAtOk returns a tuple with the AssignedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeatUserDataAttributes) GetAssignedAtOk() (*time.Time, bool) {
	if o == nil || o.AssignedAt == nil {
		return nil, false
	}
	return o.AssignedAt, true
}

// HasAssignedAt returns a boolean if a field has been set.
func (o *SeatUserDataAttributes) HasAssignedAt() bool {
	return o != nil && o.AssignedAt != nil
}

// SetAssignedAt gets a reference to the given time.Time and assigns it to the AssignedAt field.
func (o *SeatUserDataAttributes) SetAssignedAt(v time.Time) {
	o.AssignedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SeatUserDataAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeatUserDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SeatUserDataAttributes) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SeatUserDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SeatUserDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeatUserDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SeatUserDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SeatUserDataAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeatUserDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignedAt != nil {
		if o.AssignedAt.Nanosecond() == 0 {
			toSerialize["assigned_at"] = o.AssignedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["assigned_at"] = o.AssignedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
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
func (o *SeatUserDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignedAt *time.Time `json:"assigned_at,omitempty"`
		Email      *string    `json:"email,omitempty"`
		Name       *string    `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assigned_at", "email", "name"})
	} else {
		return err
	}
	o.AssignedAt = all.AssignedAt
	o.Email = all.Email
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
