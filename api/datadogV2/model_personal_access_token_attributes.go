// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PersonalAccessTokenAttributes Attributes of a personal access token.
type PersonalAccessTokenAttributes struct {
	// Creation timestamp of the personal access token.
	CreatedAt time.Time `json:"created_at"`
	// Expiration timestamp of the personal access token.
	ExpiresAt time.Time `json:"expires_at"`
	// Last modification timestamp of the personal access token.
	ModifiedAt datadog.NullableTime `json:"modified_at,omitempty"`
	// Name of the personal access token.
	Name string `json:"name"`
	// Array of scopes granted to the personal access token. These define what
	// permissions the token has.
	Scopes []string `json:"scopes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPersonalAccessTokenAttributes instantiates a new PersonalAccessTokenAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersonalAccessTokenAttributes(createdAt time.Time, expiresAt time.Time, name string, scopes []string) *PersonalAccessTokenAttributes {
	this := PersonalAccessTokenAttributes{}
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.Name = name
	this.Scopes = scopes
	return &this
}

// NewPersonalAccessTokenAttributesWithDefaults instantiates a new PersonalAccessTokenAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPersonalAccessTokenAttributesWithDefaults() *PersonalAccessTokenAttributes {
	this := PersonalAccessTokenAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *PersonalAccessTokenAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *PersonalAccessTokenAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value.
func (o *PersonalAccessTokenAttributes) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value.
func (o *PersonalAccessTokenAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersonalAccessTokenAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersonalAccessTokenAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PersonalAccessTokenAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt.IsSet()
}

// SetModifiedAt gets a reference to the given datadog.NullableTime and assigns it to the ModifiedAt field.
func (o *PersonalAccessTokenAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil.
func (o *PersonalAccessTokenAttributes) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil.
func (o *PersonalAccessTokenAttributes) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

// GetName returns the Name field value.
func (o *PersonalAccessTokenAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PersonalAccessTokenAttributes) SetName(v string) {
	o.Name = v
}

// GetScopes returns the Scopes field value.
func (o *PersonalAccessTokenAttributes) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenAttributes) GetScopesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value.
func (o *PersonalAccessTokenAttributes) SetScopes(v []string) {
	o.Scopes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PersonalAccessTokenAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ExpiresAt.Nanosecond() == 0 {
		toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["scopes"] = o.Scopes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PersonalAccessTokenAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt  *time.Time           `json:"created_at"`
		ExpiresAt  *time.Time           `json:"expires_at"`
		ModifiedAt datadog.NullableTime `json:"modified_at,omitempty"`
		Name       *string              `json:"name"`
		Scopes     *[]string            `json:"scopes"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.ExpiresAt == nil {
		return fmt.Errorf("required field expires_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Scopes == nil {
		return fmt.Errorf("required field scopes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "expires_at", "modified_at", "name", "scopes"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.ExpiresAt = *all.ExpiresAt
	o.ModifiedAt = all.ModifiedAt
	o.Name = *all.Name
	o.Scopes = *all.Scopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
