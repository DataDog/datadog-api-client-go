// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DelegatedTokenAttributes Attributes of a delegated token.
type DelegatedTokenAttributes struct {
	// A short-lived JWT representing the authenticated Datadog user. Pass this as a bearer token in subsequent API calls.
	AccessToken string `json:"access_token"`
	// The expiry time of the token.
	Expires time.Time `json:"expires"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDelegatedTokenAttributes instantiates a new DelegatedTokenAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDelegatedTokenAttributes(accessToken string, expires time.Time) *DelegatedTokenAttributes {
	this := DelegatedTokenAttributes{}
	this.AccessToken = accessToken
	this.Expires = expires
	return &this
}

// NewDelegatedTokenAttributesWithDefaults instantiates a new DelegatedTokenAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDelegatedTokenAttributesWithDefaults() *DelegatedTokenAttributes {
	this := DelegatedTokenAttributes{}
	return &this
}

// GetAccessToken returns the AccessToken field value.
func (o *DelegatedTokenAttributes) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *DelegatedTokenAttributes) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value.
func (o *DelegatedTokenAttributes) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetExpires returns the Expires field value.
func (o *DelegatedTokenAttributes) GetExpires() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
func (o *DelegatedTokenAttributes) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expires, true
}

// SetExpires sets field value.
func (o *DelegatedTokenAttributes) SetExpires(v time.Time) {
	o.Expires = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DelegatedTokenAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["access_token"] = o.AccessToken
	if o.Expires.Nanosecond() == 0 {
		toSerialize["expires"] = o.Expires.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expires"] = o.Expires.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DelegatedTokenAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessToken *string    `json:"access_token"`
		Expires     *time.Time `json:"expires"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessToken == nil {
		return fmt.Errorf("required field access_token missing")
	}
	if all.Expires == nil {
		return fmt.Errorf("required field expires missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_token", "expires"})
	} else {
		return err
	}
	o.AccessToken = *all.AccessToken
	o.Expires = *all.Expires

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
