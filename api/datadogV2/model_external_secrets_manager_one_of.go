// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExternalSecretsManagerOneOf
type ExternalSecretsManagerOneOf struct {
	// AWS Secrets Manager configuration.
	Aws AWSSecretManager `json:"aws"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExternalSecretsManagerOneOf instantiates a new ExternalSecretsManagerOneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExternalSecretsManagerOneOf(aws AWSSecretManager) *ExternalSecretsManagerOneOf {
	this := ExternalSecretsManagerOneOf{}
	this.Aws = aws
	return &this
}

// NewExternalSecretsManagerOneOfWithDefaults instantiates a new ExternalSecretsManagerOneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExternalSecretsManagerOneOfWithDefaults() *ExternalSecretsManagerOneOf {
	this := ExternalSecretsManagerOneOf{}
	return &this
}

// GetAws returns the Aws field value.
func (o *ExternalSecretsManagerOneOf) GetAws() AWSSecretManager {
	if o == nil {
		var ret AWSSecretManager
		return ret
	}
	return o.Aws
}

// GetAwsOk returns a tuple with the Aws field value
// and a boolean to check if the value has been set.
func (o *ExternalSecretsManagerOneOf) GetAwsOk() (*AWSSecretManager, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aws, true
}

// SetAws sets field value.
func (o *ExternalSecretsManagerOneOf) SetAws(v AWSSecretManager) {
	o.Aws = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExternalSecretsManagerOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aws"] = o.Aws

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExternalSecretsManagerOneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aws *AWSSecretManager `json:"aws"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aws == nil {
		return fmt.Errorf("required field aws missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aws"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aws.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Aws = *all.Aws

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
