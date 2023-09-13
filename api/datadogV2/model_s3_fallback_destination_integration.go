// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// S3FallbackDestinationIntegration The S3 Archive's integration destination.
type S3FallbackDestinationIntegration struct {
	// The account ID for the integration.
	AccountId string `json:"accountId"`
	// The path of the integration.
	RoleName string `json:"roleName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewS3FallbackDestinationIntegration instantiates a new S3FallbackDestinationIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewS3FallbackDestinationIntegration(accountId string, roleName string) *S3FallbackDestinationIntegration {
	this := S3FallbackDestinationIntegration{}
	this.AccountId = accountId
	this.RoleName = roleName
	return &this
}

// NewS3FallbackDestinationIntegrationWithDefaults instantiates a new S3FallbackDestinationIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewS3FallbackDestinationIntegrationWithDefaults() *S3FallbackDestinationIntegration {
	this := S3FallbackDestinationIntegration{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *S3FallbackDestinationIntegration) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *S3FallbackDestinationIntegration) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *S3FallbackDestinationIntegration) SetAccountId(v string) {
	o.AccountId = v
}

// GetRoleName returns the RoleName field value.
func (o *S3FallbackDestinationIntegration) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *S3FallbackDestinationIntegration) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value.
func (o *S3FallbackDestinationIntegration) SetRoleName(v string) {
	o.RoleName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o S3FallbackDestinationIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["accountId"] = o.AccountId
	toSerialize["roleName"] = o.RoleName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *S3FallbackDestinationIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId *string `json:"accountId"`
		RoleName  *string `json:"roleName"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field accountId missing")
	}
	if all.RoleName == nil {
		return fmt.Errorf("required field roleName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accountId", "roleName"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.RoleName = *all.RoleName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
