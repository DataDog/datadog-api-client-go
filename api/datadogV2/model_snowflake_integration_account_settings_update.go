// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountSettingsUpdate Settings for updating the Snowflake integration account. Only the fields provided are changed.
type SnowflakeIntegrationAccountSettingsUpdate struct {
	// Identifier of the Snowflake account to monitor, either as `organization-account` or as the legacy `account_name.region_id.cloud_provider` account locator. An account identifier can be configured once per Datadog organization; reusing one is rejected with a `422` response. Accounts reached through AWS PrivateLink are not supported.
	SnowflakeAccountIdentifier *string `json:"snowflake_account_identifier,omitempty"`
	// Snowflake user Datadog authenticates as. Create a dedicated user for Datadog and grant it a role with access to the data you want to collect.
	Username *string `json:"username,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationAccountSettingsUpdate instantiates a new SnowflakeIntegrationAccountSettingsUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationAccountSettingsUpdate() *SnowflakeIntegrationAccountSettingsUpdate {
	this := SnowflakeIntegrationAccountSettingsUpdate{}
	return &this
}

// NewSnowflakeIntegrationAccountSettingsUpdateWithDefaults instantiates a new SnowflakeIntegrationAccountSettingsUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationAccountSettingsUpdateWithDefaults() *SnowflakeIntegrationAccountSettingsUpdate {
	this := SnowflakeIntegrationAccountSettingsUpdate{}
	return &this
}

// GetSnowflakeAccountIdentifier returns the SnowflakeAccountIdentifier field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountSettingsUpdate) GetSnowflakeAccountIdentifier() string {
	if o == nil || o.SnowflakeAccountIdentifier == nil {
		var ret string
		return ret
	}
	return *o.SnowflakeAccountIdentifier
}

// GetSnowflakeAccountIdentifierOk returns a tuple with the SnowflakeAccountIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountSettingsUpdate) GetSnowflakeAccountIdentifierOk() (*string, bool) {
	if o == nil || o.SnowflakeAccountIdentifier == nil {
		return nil, false
	}
	return o.SnowflakeAccountIdentifier, true
}

// HasSnowflakeAccountIdentifier returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountSettingsUpdate) HasSnowflakeAccountIdentifier() bool {
	return o != nil && o.SnowflakeAccountIdentifier != nil
}

// SetSnowflakeAccountIdentifier gets a reference to the given string and assigns it to the SnowflakeAccountIdentifier field.
func (o *SnowflakeIntegrationAccountSettingsUpdate) SetSnowflakeAccountIdentifier(v string) {
	o.SnowflakeAccountIdentifier = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountSettingsUpdate) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountSettingsUpdate) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountSettingsUpdate) HasUsername() bool {
	return o != nil && o.Username != nil
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SnowflakeIntegrationAccountSettingsUpdate) SetUsername(v string) {
	o.Username = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationAccountSettingsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SnowflakeAccountIdentifier != nil {
		toSerialize["snowflake_account_identifier"] = o.SnowflakeAccountIdentifier
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeIntegrationAccountSettingsUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SnowflakeAccountIdentifier *string `json:"snowflake_account_identifier,omitempty"`
		Username                   *string `json:"username,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.SnowflakeAccountIdentifier = all.SnowflakeAccountIdentifier
	o.Username = all.Username

	return nil
}
