// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LegacyUserJourneySearchJoinKeys Test-only generated schema for the dashboard versioning proof of concept.
type LegacyUserJourneySearchJoinKeys struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Primary string `json:"primary"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Secondary []string `json:"secondary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLegacyUserJourneySearchJoinKeys instantiates a new LegacyUserJourneySearchJoinKeys object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLegacyUserJourneySearchJoinKeys(primary string) *LegacyUserJourneySearchJoinKeys {
	this := LegacyUserJourneySearchJoinKeys{}
	this.Primary = primary
	return &this
}

// NewLegacyUserJourneySearchJoinKeysWithDefaults instantiates a new LegacyUserJourneySearchJoinKeys object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLegacyUserJourneySearchJoinKeysWithDefaults() *LegacyUserJourneySearchJoinKeys {
	this := LegacyUserJourneySearchJoinKeys{}
	return &this
}

// GetPrimary returns the Primary field value.
func (o *LegacyUserJourneySearchJoinKeys) GetPrimary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *LegacyUserJourneySearchJoinKeys) GetPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value.
func (o *LegacyUserJourneySearchJoinKeys) SetPrimary(v string) {
	o.Primary = v
}

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *LegacyUserJourneySearchJoinKeys) GetSecondary() []string {
	if o == nil || o.Secondary == nil {
		var ret []string
		return ret
	}
	return o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyUserJourneySearchJoinKeys) GetSecondaryOk() (*[]string, bool) {
	if o == nil || o.Secondary == nil {
		return nil, false
	}
	return &o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *LegacyUserJourneySearchJoinKeys) HasSecondary() bool {
	return o != nil && o.Secondary != nil
}

// SetSecondary gets a reference to the given []string and assigns it to the Secondary field.
func (o *LegacyUserJourneySearchJoinKeys) SetSecondary(v []string) {
	o.Secondary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LegacyUserJourneySearchJoinKeys) MarshalJSON() ([]byte, error) {
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
func (o *LegacyUserJourneySearchJoinKeys) UnmarshalJSON(bytes []byte) (err error) {
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
