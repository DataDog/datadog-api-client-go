// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPBasicAuthUpdate The definition of the `HTTPBasicAuth` object.
type HTTPBasicAuthUpdate struct {
	// Password used for authentication. Saved in a secret store
	Password *string `json:"password,omitempty"`
	// The definition of the `HTTPBasicAuth` object.
	Type HTTPBasicAuthType `json:"type"`
	// Username used for authentication.
	Username *string `json:"username,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHTTPBasicAuthUpdate instantiates a new HTTPBasicAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPBasicAuthUpdate(typeVar HTTPBasicAuthType) *HTTPBasicAuthUpdate {
	this := HTTPBasicAuthUpdate{}
	this.Type = typeVar
	return &this
}

// NewHTTPBasicAuthUpdateWithDefaults instantiates a new HTTPBasicAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPBasicAuthUpdateWithDefaults() *HTTPBasicAuthUpdate {
	this := HTTPBasicAuthUpdate{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *HTTPBasicAuthUpdate) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPBasicAuthUpdate) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *HTTPBasicAuthUpdate) HasPassword() bool {
	return o != nil && o.Password != nil
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *HTTPBasicAuthUpdate) SetPassword(v string) {
	o.Password = &v
}

// GetType returns the Type field value.
func (o *HTTPBasicAuthUpdate) GetType() HTTPBasicAuthType {
	if o == nil {
		var ret HTTPBasicAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HTTPBasicAuthUpdate) GetTypeOk() (*HTTPBasicAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HTTPBasicAuthUpdate) SetType(v HTTPBasicAuthType) {
	o.Type = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *HTTPBasicAuthUpdate) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPBasicAuthUpdate) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *HTTPBasicAuthUpdate) HasUsername() bool {
	return o != nil && o.Username != nil
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *HTTPBasicAuthUpdate) SetUsername(v string) {
	o.Username = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPBasicAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	toSerialize["type"] = o.Type
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HTTPBasicAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Password *string            `json:"password,omitempty"`
		Type     *HTTPBasicAuthType `json:"type"`
		Username *string            `json:"username,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"password", "type", "username"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Password = all.Password
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Username = all.Username

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
