// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsWifPersonaMappingAttributes Attributes of an AWS WIF persona mapping.
type AwsWifPersonaMappingAttributes struct {
	// The Datadog user handle (email address) to map the AWS principal to.
	AccountIdentifier string `json:"account_identifier"`
	// The Datadog user UUID corresponding to `account_identifier`. Read-only — set by the server.
	AccountUuid *uuid.UUID `json:"account_uuid,omitempty"`
	// The AWS IAM ARN pattern identifying the role or user that will be mapped.
	// Supports wildcards (`*`) to match multiple principals within an account.
	ArnPattern string `json:"arn_pattern"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsWifPersonaMappingAttributes instantiates a new AwsWifPersonaMappingAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsWifPersonaMappingAttributes(accountIdentifier string, arnPattern string) *AwsWifPersonaMappingAttributes {
	this := AwsWifPersonaMappingAttributes{}
	this.AccountIdentifier = accountIdentifier
	this.ArnPattern = arnPattern
	return &this
}

// NewAwsWifPersonaMappingAttributesWithDefaults instantiates a new AwsWifPersonaMappingAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsWifPersonaMappingAttributesWithDefaults() *AwsWifPersonaMappingAttributes {
	this := AwsWifPersonaMappingAttributes{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value.
func (o *AwsWifPersonaMappingAttributes) GetAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *AwsWifPersonaMappingAttributes) GetAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifier, true
}

// SetAccountIdentifier sets field value.
func (o *AwsWifPersonaMappingAttributes) SetAccountIdentifier(v string) {
	o.AccountIdentifier = v
}

// GetAccountUuid returns the AccountUuid field value if set, zero value otherwise.
func (o *AwsWifPersonaMappingAttributes) GetAccountUuid() uuid.UUID {
	if o == nil || o.AccountUuid == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.AccountUuid
}

// GetAccountUuidOk returns a tuple with the AccountUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsWifPersonaMappingAttributes) GetAccountUuidOk() (*uuid.UUID, bool) {
	if o == nil || o.AccountUuid == nil {
		return nil, false
	}
	return o.AccountUuid, true
}

// HasAccountUuid returns a boolean if a field has been set.
func (o *AwsWifPersonaMappingAttributes) HasAccountUuid() bool {
	return o != nil && o.AccountUuid != nil
}

// SetAccountUuid gets a reference to the given uuid.UUID and assigns it to the AccountUuid field.
func (o *AwsWifPersonaMappingAttributes) SetAccountUuid(v uuid.UUID) {
	o.AccountUuid = &v
}

// GetArnPattern returns the ArnPattern field value.
func (o *AwsWifPersonaMappingAttributes) GetArnPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArnPattern
}

// GetArnPatternOk returns a tuple with the ArnPattern field value
// and a boolean to check if the value has been set.
func (o *AwsWifPersonaMappingAttributes) GetArnPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArnPattern, true
}

// SetArnPattern sets field value.
func (o *AwsWifPersonaMappingAttributes) SetArnPattern(v string) {
	o.ArnPattern = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsWifPersonaMappingAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_identifier"] = o.AccountIdentifier
	if o.AccountUuid != nil {
		toSerialize["account_uuid"] = o.AccountUuid
	}
	toSerialize["arn_pattern"] = o.ArnPattern

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsWifPersonaMappingAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountIdentifier *string    `json:"account_identifier"`
		AccountUuid       *uuid.UUID `json:"account_uuid,omitempty"`
		ArnPattern        *string    `json:"arn_pattern"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"account_identifier", "account_uuid", "arn_pattern"})
	} else {
		return err
	}
	o.AccountIdentifier = *all.AccountIdentifier
	o.AccountUuid = all.AccountUuid
	o.ArnPattern = *all.ArnPattern

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
