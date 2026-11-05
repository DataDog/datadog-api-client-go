// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionAWSSecretsManagerReference Reference to a string value in AWS Secrets Manager. The connection must also specify `external_secrets_manager` settings. AWS Secrets Manager references are supported for token-based credentials and HTTP authentication tokens on `DATACENTER` connections.
type ActionConnectionAWSSecretsManagerReference struct {
	// Top-level JSON key to extract from the secret string. Omit this field to use the entire secret string.
	Key *string `json:"key,omitempty"`
	// Name or ARN of the AWS Secrets Manager secret.
	SecretId string `json:"secret_id"`
	// AWS Secrets Manager secret backend.
	Type ActionConnectionAWSSecretsManagerType `json:"type"`
	// ID of the secret version to retrieve. Cannot be used together with `version_stage`. If both are omitted, the current version is used.
	VersionId *string `json:"version_id,omitempty"`
	// Staging label of the secret version to retrieve. Cannot be used together with `version_id`. If both are omitted, the current version is used.
	VersionStage *string `json:"version_stage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionConnectionAWSSecretsManagerReference instantiates a new ActionConnectionAWSSecretsManagerReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionConnectionAWSSecretsManagerReference(secretId string, typeVar ActionConnectionAWSSecretsManagerType) *ActionConnectionAWSSecretsManagerReference {
	this := ActionConnectionAWSSecretsManagerReference{}
	this.SecretId = secretId
	this.Type = typeVar
	return &this
}

// NewActionConnectionAWSSecretsManagerReferenceWithDefaults instantiates a new ActionConnectionAWSSecretsManagerReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionConnectionAWSSecretsManagerReferenceWithDefaults() *ActionConnectionAWSSecretsManagerReference {
	this := ActionConnectionAWSSecretsManagerReference{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ActionConnectionAWSSecretsManagerReference) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionConnectionAWSSecretsManagerReference) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ActionConnectionAWSSecretsManagerReference) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ActionConnectionAWSSecretsManagerReference) SetKey(v string) {
	o.Key = &v
}

// GetSecretId returns the SecretId field value.
func (o *ActionConnectionAWSSecretsManagerReference) GetSecretId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value
// and a boolean to check if the value has been set.
func (o *ActionConnectionAWSSecretsManagerReference) GetSecretIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretId, true
}

// SetSecretId sets field value.
func (o *ActionConnectionAWSSecretsManagerReference) SetSecretId(v string) {
	o.SecretId = v
}

// GetType returns the Type field value.
func (o *ActionConnectionAWSSecretsManagerReference) GetType() ActionConnectionAWSSecretsManagerType {
	if o == nil {
		var ret ActionConnectionAWSSecretsManagerType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActionConnectionAWSSecretsManagerReference) GetTypeOk() (*ActionConnectionAWSSecretsManagerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ActionConnectionAWSSecretsManagerReference) SetType(v ActionConnectionAWSSecretsManagerType) {
	o.Type = v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *ActionConnectionAWSSecretsManagerReference) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionConnectionAWSSecretsManagerReference) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *ActionConnectionAWSSecretsManagerReference) HasVersionId() bool {
	return o != nil && o.VersionId != nil
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *ActionConnectionAWSSecretsManagerReference) SetVersionId(v string) {
	o.VersionId = &v
}

// GetVersionStage returns the VersionStage field value if set, zero value otherwise.
func (o *ActionConnectionAWSSecretsManagerReference) GetVersionStage() string {
	if o == nil || o.VersionStage == nil {
		var ret string
		return ret
	}
	return *o.VersionStage
}

// GetVersionStageOk returns a tuple with the VersionStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionConnectionAWSSecretsManagerReference) GetVersionStageOk() (*string, bool) {
	if o == nil || o.VersionStage == nil {
		return nil, false
	}
	return o.VersionStage, true
}

// HasVersionStage returns a boolean if a field has been set.
func (o *ActionConnectionAWSSecretsManagerReference) HasVersionStage() bool {
	return o != nil && o.VersionStage != nil
}

// SetVersionStage gets a reference to the given string and assigns it to the VersionStage field.
func (o *ActionConnectionAWSSecretsManagerReference) SetVersionStage(v string) {
	o.VersionStage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionConnectionAWSSecretsManagerReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	toSerialize["secret_id"] = o.SecretId
	toSerialize["type"] = o.Type
	if o.VersionId != nil {
		toSerialize["version_id"] = o.VersionId
	}
	if o.VersionStage != nil {
		toSerialize["version_stage"] = o.VersionStage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionConnectionAWSSecretsManagerReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key          *string                                `json:"key,omitempty"`
		SecretId     *string                                `json:"secret_id"`
		Type         *ActionConnectionAWSSecretsManagerType `json:"type"`
		VersionId    *string                                `json:"version_id,omitempty"`
		VersionStage *string                                `json:"version_stage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SecretId == nil {
		return fmt.Errorf("required field secret_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "secret_id", "type", "version_id", "version_stage"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Key = all.Key
	o.SecretId = *all.SecretId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.VersionId = all.VersionId
	o.VersionStage = all.VersionStage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
