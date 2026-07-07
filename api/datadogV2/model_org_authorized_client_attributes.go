// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientAttributes Attributes of an org authorized client.
type OrgAuthorizedClientAttributes struct {
	// Whether the organization has disabled this client.
	Disabled bool `json:"disabled"`
	// The date and time this client was last exercised.
	LastExercised datadog.NullableTime `json:"last_exercised"`
	// The number of users in the organization who have authorized this client.
	UserCount int64 `json:"user_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgAuthorizedClientAttributes instantiates a new OrgAuthorizedClientAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgAuthorizedClientAttributes(disabled bool, lastExercised datadog.NullableTime, userCount int64) *OrgAuthorizedClientAttributes {
	this := OrgAuthorizedClientAttributes{}
	this.Disabled = disabled
	this.LastExercised = lastExercised
	this.UserCount = userCount
	return &this
}

// NewOrgAuthorizedClientAttributesWithDefaults instantiates a new OrgAuthorizedClientAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgAuthorizedClientAttributesWithDefaults() *OrgAuthorizedClientAttributes {
	this := OrgAuthorizedClientAttributes{}
	return &this
}

// GetDisabled returns the Disabled field value.
func (o *OrgAuthorizedClientAttributes) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *OrgAuthorizedClientAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value.
func (o *OrgAuthorizedClientAttributes) SetDisabled(v bool) {
	o.Disabled = v
}

// GetLastExercised returns the LastExercised field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *OrgAuthorizedClientAttributes) GetLastExercised() time.Time {
	if o == nil || o.LastExercised.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastExercised.Get()
}

// GetLastExercisedOk returns a tuple with the LastExercised field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OrgAuthorizedClientAttributes) GetLastExercisedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExercised.Get(), o.LastExercised.IsSet()
}

// SetLastExercised sets field value.
func (o *OrgAuthorizedClientAttributes) SetLastExercised(v time.Time) {
	o.LastExercised.Set(&v)
}

// GetUserCount returns the UserCount field value.
func (o *OrgAuthorizedClientAttributes) GetUserCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value
// and a boolean to check if the value has been set.
func (o *OrgAuthorizedClientAttributes) GetUserCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserCount, true
}

// SetUserCount sets field value.
func (o *OrgAuthorizedClientAttributes) SetUserCount(v int64) {
	o.UserCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgAuthorizedClientAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["disabled"] = o.Disabled
	toSerialize["last_exercised"] = o.LastExercised.Get()
	toSerialize["user_count"] = o.UserCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgAuthorizedClientAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Disabled      *bool                `json:"disabled"`
		LastExercised datadog.NullableTime `json:"last_exercised"`
		UserCount     *int64               `json:"user_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Disabled == nil {
		return fmt.Errorf("required field disabled missing")
	}
	if !all.LastExercised.IsSet() {
		return fmt.Errorf("required field last_exercised missing")
	}
	if all.UserCount == nil {
		return fmt.Errorf("required field user_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"disabled", "last_exercised", "user_count"})
	} else {
		return err
	}
	o.Disabled = *all.Disabled
	o.LastExercised = all.LastExercised
	o.UserCount = *all.UserCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
