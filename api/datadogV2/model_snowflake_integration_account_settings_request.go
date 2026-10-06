// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountSettingsRequest Settings for creating the Snowflake integration account.
type SnowflakeIntegrationAccountSettingsRequest struct {
	// Identifier of the Snowflake account to monitor, either as `organization-account` or as the legacy `account_name.region_id.cloud_provider` account locator. An account identifier can be configured once per Datadog organization; reusing one is rejected with a `422` response. Accounts reached through AWS PrivateLink are not supported.
	SnowflakeAccountIdentifier string `json:"snowflake_account_identifier"`
	// Snowflake user Datadog authenticates as. Create a dedicated user for Datadog and grant it a role with access to the data you want to collect.
	Username string `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationAccountSettingsRequest instantiates a new SnowflakeIntegrationAccountSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationAccountSettingsRequest(snowflakeAccountIdentifier string, username string) *SnowflakeIntegrationAccountSettingsRequest {
	this := SnowflakeIntegrationAccountSettingsRequest{}
	this.SnowflakeAccountIdentifier = snowflakeAccountIdentifier
	this.Username = username
	return &this
}

// NewSnowflakeIntegrationAccountSettingsRequestWithDefaults instantiates a new SnowflakeIntegrationAccountSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationAccountSettingsRequestWithDefaults() *SnowflakeIntegrationAccountSettingsRequest {
	this := SnowflakeIntegrationAccountSettingsRequest{}
	return &this
}

// GetSnowflakeAccountIdentifier returns the SnowflakeAccountIdentifier field value.
func (o *SnowflakeIntegrationAccountSettingsRequest) GetSnowflakeAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SnowflakeAccountIdentifier
}

// GetSnowflakeAccountIdentifierOk returns a tuple with the SnowflakeAccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountSettingsRequest) GetSnowflakeAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnowflakeAccountIdentifier, true
}

// SetSnowflakeAccountIdentifier sets field value.
func (o *SnowflakeIntegrationAccountSettingsRequest) SetSnowflakeAccountIdentifier(v string) {
	o.SnowflakeAccountIdentifier = v
}

// GetUsername returns the Username field value.
func (o *SnowflakeIntegrationAccountSettingsRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountSettingsRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *SnowflakeIntegrationAccountSettingsRequest) SetUsername(v string) {
	o.Username = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationAccountSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["snowflake_account_identifier"] = o.SnowflakeAccountIdentifier
	toSerialize["username"] = o.Username
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeIntegrationAccountSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SnowflakeAccountIdentifier *string `json:"snowflake_account_identifier"`
		Username                   *string `json:"username"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SnowflakeAccountIdentifier == nil {
		return fmt.Errorf("required field snowflake_account_identifier missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}
	o.SnowflakeAccountIdentifier = *all.SnowflakeAccountIdentifier
	o.Username = *all.Username

	return nil
}
