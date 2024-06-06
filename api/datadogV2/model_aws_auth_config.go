// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAuthConfig AWS Authentication config
type AWSAuthConfig struct {
	// AWS Access Key ID
	AccessKeyId *string `json:"access_key_id,omitempty"`
	// AWS IAM External ID for associated role
	ExternalId *string `json:"external_id,omitempty"`
	// AWS IAM Role name
	RoleName *string `json:"role_name,omitempty"`
	// AWS Secret Access Key
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSAuthConfig instantiates a new AWSAuthConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAuthConfig() *AWSAuthConfig {
	this := AWSAuthConfig{}
	return &this
}

// NewAWSAuthConfigWithDefaults instantiates a new AWSAuthConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAuthConfigWithDefaults() *AWSAuthConfig {
	this := AWSAuthConfig{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AWSAuthConfig) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAuthConfig) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AWSAuthConfig) HasAccessKeyId() bool {
	return o != nil && o.AccessKeyId != nil
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AWSAuthConfig) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AWSAuthConfig) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAuthConfig) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AWSAuthConfig) HasExternalId() bool {
	return o != nil && o.ExternalId != nil
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AWSAuthConfig) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *AWSAuthConfig) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAuthConfig) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *AWSAuthConfig) HasRoleName() bool {
	return o != nil && o.RoleName != nil
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *AWSAuthConfig) SetRoleName(v string) {
	o.RoleName = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *AWSAuthConfig) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAuthConfig) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *AWSAuthConfig) HasSecretAccessKey() bool {
	return o != nil && o.SecretAccessKey != nil
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *AWSAuthConfig) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAuthConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessKeyId != nil {
		toSerialize["access_key_id"] = o.AccessKeyId
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	if o.SecretAccessKey != nil {
		toSerialize["secret_access_key"] = o.SecretAccessKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSAuthConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessKeyId     *string `json:"access_key_id,omitempty"`
		ExternalId      *string `json:"external_id,omitempty"`
		RoleName        *string `json:"role_name,omitempty"`
		SecretAccessKey *string `json:"secret_access_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_key_id", "external_id", "role_name", "secret_access_key"})
	} else {
		return err
	}
	o.AccessKeyId = all.AccessKeyId
	o.ExternalId = all.ExternalId
	o.RoleName = all.RoleName
	o.SecretAccessKey = all.SecretAccessKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
