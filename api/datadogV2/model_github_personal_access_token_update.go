// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubPersonalAccessTokenUpdate Update to GitHub personal access token credentials.
type GithubPersonalAccessTokenUpdate struct {
	// GitHub personal access token.
	AccessToken *GithubPersonalAccessTokenUpdateAccessToken `json:"access_token,omitempty"`
	// Hostname of the GitHub Enterprise Server. Omit this field for github.com.
	Hostname *GithubPersonalAccessTokenUpdateHostname `json:"hostname,omitempty"`
	// The credential type for `GithubPersonalAccessToken`.
	Type GithubPersonalAccessTokenType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGithubPersonalAccessTokenUpdate instantiates a new GithubPersonalAccessTokenUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGithubPersonalAccessTokenUpdate(typeVar GithubPersonalAccessTokenType) *GithubPersonalAccessTokenUpdate {
	this := GithubPersonalAccessTokenUpdate{}
	this.Type = typeVar
	return &this
}

// NewGithubPersonalAccessTokenUpdateWithDefaults instantiates a new GithubPersonalAccessTokenUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGithubPersonalAccessTokenUpdateWithDefaults() *GithubPersonalAccessTokenUpdate {
	this := GithubPersonalAccessTokenUpdate{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *GithubPersonalAccessTokenUpdate) GetAccessToken() GithubPersonalAccessTokenUpdateAccessToken {
	if o == nil || o.AccessToken == nil {
		var ret GithubPersonalAccessTokenUpdateAccessToken
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubPersonalAccessTokenUpdate) GetAccessTokenOk() (*GithubPersonalAccessTokenUpdateAccessToken, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *GithubPersonalAccessTokenUpdate) HasAccessToken() bool {
	return o != nil && o.AccessToken != nil
}

// SetAccessToken gets a reference to the given GithubPersonalAccessTokenUpdateAccessToken and assigns it to the AccessToken field.
func (o *GithubPersonalAccessTokenUpdate) SetAccessToken(v GithubPersonalAccessTokenUpdateAccessToken) {
	o.AccessToken = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GithubPersonalAccessTokenUpdate) GetHostname() GithubPersonalAccessTokenUpdateHostname {
	if o == nil || o.Hostname == nil {
		var ret GithubPersonalAccessTokenUpdateHostname
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubPersonalAccessTokenUpdate) GetHostnameOk() (*GithubPersonalAccessTokenUpdateHostname, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GithubPersonalAccessTokenUpdate) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given GithubPersonalAccessTokenUpdateHostname and assigns it to the Hostname field.
func (o *GithubPersonalAccessTokenUpdate) SetHostname(v GithubPersonalAccessTokenUpdateHostname) {
	o.Hostname = &v
}

// GetType returns the Type field value.
func (o *GithubPersonalAccessTokenUpdate) GetType() GithubPersonalAccessTokenType {
	if o == nil {
		var ret GithubPersonalAccessTokenType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GithubPersonalAccessTokenUpdate) GetTypeOk() (*GithubPersonalAccessTokenType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GithubPersonalAccessTokenUpdate) SetType(v GithubPersonalAccessTokenType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GithubPersonalAccessTokenUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GithubPersonalAccessTokenUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessToken *GithubPersonalAccessTokenUpdateAccessToken `json:"access_token,omitempty"`
		Hostname    *GithubPersonalAccessTokenUpdateHostname    `json:"hostname,omitempty"`
		Type        *GithubPersonalAccessTokenType              `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_token", "hostname", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccessToken = all.AccessToken
	o.Hostname = all.Hostname
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
