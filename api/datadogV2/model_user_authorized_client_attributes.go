// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserAuthorizedClientAttributes Attributes of a user authorized client.
type UserAuthorizedClientAttributes struct {
	// The date and time this authorization was created.
	CreatedAt time.Time `json:"created_at"`
	// Whether the user has disabled this authorization.
	Disabled bool `json:"disabled"`
	// The date and time this authorization was last exercised.
	LastExercised datadog.NullableTime `json:"last_exercised"`
	// The date and time this authorization was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// Whether the organization has disabled this authorization.
	OrgDisabled bool `json:"org_disabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserAuthorizedClientAttributes instantiates a new UserAuthorizedClientAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserAuthorizedClientAttributes(createdAt time.Time, disabled bool, lastExercised datadog.NullableTime, modifiedAt time.Time, orgDisabled bool) *UserAuthorizedClientAttributes {
	this := UserAuthorizedClientAttributes{}
	this.CreatedAt = createdAt
	this.Disabled = disabled
	this.LastExercised = lastExercised
	this.ModifiedAt = modifiedAt
	this.OrgDisabled = orgDisabled
	return &this
}

// NewUserAuthorizedClientAttributesWithDefaults instantiates a new UserAuthorizedClientAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserAuthorizedClientAttributesWithDefaults() *UserAuthorizedClientAttributes {
	this := UserAuthorizedClientAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *UserAuthorizedClientAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserAuthorizedClientAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *UserAuthorizedClientAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDisabled returns the Disabled field value.
func (o *UserAuthorizedClientAttributes) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *UserAuthorizedClientAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value.
func (o *UserAuthorizedClientAttributes) SetDisabled(v bool) {
	o.Disabled = v
}

// GetLastExercised returns the LastExercised field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *UserAuthorizedClientAttributes) GetLastExercised() time.Time {
	if o == nil || o.LastExercised.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastExercised.Get()
}

// GetLastExercisedOk returns a tuple with the LastExercised field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UserAuthorizedClientAttributes) GetLastExercisedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExercised.Get(), o.LastExercised.IsSet()
}

// SetLastExercised sets field value.
func (o *UserAuthorizedClientAttributes) SetLastExercised(v time.Time) {
	o.LastExercised.Set(&v)
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *UserAuthorizedClientAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *UserAuthorizedClientAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *UserAuthorizedClientAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetOrgDisabled returns the OrgDisabled field value.
func (o *UserAuthorizedClientAttributes) GetOrgDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.OrgDisabled
}

// GetOrgDisabledOk returns a tuple with the OrgDisabled field value
// and a boolean to check if the value has been set.
func (o *UserAuthorizedClientAttributes) GetOrgDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgDisabled, true
}

// SetOrgDisabled sets field value.
func (o *UserAuthorizedClientAttributes) SetOrgDisabled(v bool) {
	o.OrgDisabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserAuthorizedClientAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["disabled"] = o.Disabled
	toSerialize["last_exercised"] = o.LastExercised.Get()
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["org_disabled"] = o.OrgDisabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserAuthorizedClientAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time           `json:"created_at"`
		Disabled      *bool                `json:"disabled"`
		LastExercised datadog.NullableTime `json:"last_exercised"`
		ModifiedAt    *time.Time           `json:"modified_at"`
		OrgDisabled   *bool                `json:"org_disabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Disabled == nil {
		return fmt.Errorf("required field disabled missing")
	}
	if !all.LastExercised.IsSet() {
		return fmt.Errorf("required field last_exercised missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.OrgDisabled == nil {
		return fmt.Errorf("required field org_disabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "disabled", "last_exercised", "modified_at", "org_disabled"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.Disabled = *all.Disabled
	o.LastExercised = all.LastExercised
	o.ModifiedAt = *all.ModifiedAt
	o.OrgDisabled = *all.OrgDisabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
