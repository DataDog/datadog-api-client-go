// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsWifPersonaMappingCreateAttributes Attributes for creating an AWS WIF persona mapping.
type AwsWifPersonaMappingCreateAttributes struct {
	// The Datadog user handle (email address) to map the AWS principal to.
	AccountIdentifier string `json:"account_identifier"`
	// The AWS IAM ARN pattern identifying the role or user that will be mapped.
	// Supports wildcards (`*`) to match multiple principals within an account.
	ArnPattern string `json:"arn_pattern"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsWifPersonaMappingCreateAttributes instantiates a new AwsWifPersonaMappingCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsWifPersonaMappingCreateAttributes(accountIdentifier string, arnPattern string) *AwsWifPersonaMappingCreateAttributes {
	this := AwsWifPersonaMappingCreateAttributes{}
	this.AccountIdentifier = accountIdentifier
	this.ArnPattern = arnPattern
	return &this
}

// NewAwsWifPersonaMappingCreateAttributesWithDefaults instantiates a new AwsWifPersonaMappingCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsWifPersonaMappingCreateAttributesWithDefaults() *AwsWifPersonaMappingCreateAttributes {
	this := AwsWifPersonaMappingCreateAttributes{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value.
func (o *AwsWifPersonaMappingCreateAttributes) GetAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *AwsWifPersonaMappingCreateAttributes) GetAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifier, true
}

// SetAccountIdentifier sets field value.
func (o *AwsWifPersonaMappingCreateAttributes) SetAccountIdentifier(v string) {
	o.AccountIdentifier = v
}

// GetArnPattern returns the ArnPattern field value.
func (o *AwsWifPersonaMappingCreateAttributes) GetArnPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArnPattern
}

// GetArnPatternOk returns a tuple with the ArnPattern field value
// and a boolean to check if the value has been set.
func (o *AwsWifPersonaMappingCreateAttributes) GetArnPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArnPattern, true
}

// SetArnPattern sets field value.
func (o *AwsWifPersonaMappingCreateAttributes) SetArnPattern(v string) {
	o.ArnPattern = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsWifPersonaMappingCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_identifier"] = o.AccountIdentifier
	toSerialize["arn_pattern"] = o.ArnPattern

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsWifPersonaMappingCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountIdentifier *string `json:"account_identifier"`
		ArnPattern        *string `json:"arn_pattern"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountIdentifier == nil {
		return fmt.Errorf("required field account_identifier missing")
	}
	if all.ArnPattern == nil {
		return fmt.Errorf("required field arn_pattern missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_identifier", "arn_pattern"})
	} else {
		return err
	}
	o.AccountIdentifier = *all.AccountIdentifier
	o.ArnPattern = *all.ArnPattern

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
