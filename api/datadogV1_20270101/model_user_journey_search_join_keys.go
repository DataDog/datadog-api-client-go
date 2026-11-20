// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserJourneySearchJoinKeys Test-only generated schema for the dashboard versioning proof of concept.
type UserJourneySearchJoinKeys struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Primary string `json:"primary"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Secondary []string `json:"secondary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewUserJourneySearchJoinKeys instantiates a new UserJourneySearchJoinKeys object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserJourneySearchJoinKeys(primary string) *UserJourneySearchJoinKeys {
	this := UserJourneySearchJoinKeys{}
	this.Primary = primary
	return &this
}

// NewUserJourneySearchJoinKeysWithDefaults instantiates a new UserJourneySearchJoinKeys object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserJourneySearchJoinKeysWithDefaults() *UserJourneySearchJoinKeys {
	this := UserJourneySearchJoinKeys{}
	return &this
}

// GetPrimary returns the Primary field value.
func (o *UserJourneySearchJoinKeys) GetPrimary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *UserJourneySearchJoinKeys) GetPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value.
func (o *UserJourneySearchJoinKeys) SetPrimary(v string) {
	o.Primary = v
}

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *UserJourneySearchJoinKeys) GetSecondary() []string {
	if o == nil || o.Secondary == nil {
		var ret []string
		return ret
	}
	return o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearchJoinKeys) GetSecondaryOk() (*[]string, bool) {
	if o == nil || o.Secondary == nil {
		return nil, false
	}
	return &o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *UserJourneySearchJoinKeys) HasSecondary() bool {
	return o != nil && o.Secondary != nil
}

// SetSecondary gets a reference to the given []string and assigns it to the Secondary field.
func (o *UserJourneySearchJoinKeys) SetSecondary(v []string) {
	o.Secondary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserJourneySearchJoinKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["primary"] = o.Primary
	if o.Secondary != nil {
		toSerialize["secondary"] = o.Secondary
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserJourneySearchJoinKeys) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Primary   *string  `json:"primary"`
		Secondary []string `json:"secondary,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Primary == nil {
		return fmt.Errorf("required field primary missing")
	}
	o.Primary = *all.Primary
	o.Secondary = all.Secondary

	return nil
}
