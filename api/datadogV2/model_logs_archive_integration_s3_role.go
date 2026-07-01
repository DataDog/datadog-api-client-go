// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveIntegrationS3Role The S3 Archive's integration destination using an IAM role.
type LogsArchiveIntegrationS3Role struct {
	// The account ID for the integration.
	AccountId string `json:"account_id"`
	// The name of the role to assume for the integration.
	RoleName string `json:"role_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArchiveIntegrationS3Role instantiates a new LogsArchiveIntegrationS3Role object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArchiveIntegrationS3Role(accountId string, roleName string) *LogsArchiveIntegrationS3Role {
	this := LogsArchiveIntegrationS3Role{}
	this.AccountId = accountId
	this.RoleName = roleName
	return &this
}

// NewLogsArchiveIntegrationS3RoleWithDefaults instantiates a new LogsArchiveIntegrationS3Role object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArchiveIntegrationS3RoleWithDefaults() *LogsArchiveIntegrationS3Role {
	this := LogsArchiveIntegrationS3Role{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *LogsArchiveIntegrationS3Role) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveIntegrationS3Role) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *LogsArchiveIntegrationS3Role) SetAccountId(v string) {
	o.AccountId = v
}

// GetRoleName returns the RoleName field value.
func (o *LogsArchiveIntegrationS3Role) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveIntegrationS3Role) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value.
func (o *LogsArchiveIntegrationS3Role) SetRoleName(v string) {
	o.RoleName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArchiveIntegrationS3Role) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["role_name"] = o.RoleName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArchiveIntegrationS3Role) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId *string `json:"account_id"`
		RoleName  *string `json:"role_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.RoleName == nil {
		return fmt.Errorf("required field role_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "role_name"})
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
