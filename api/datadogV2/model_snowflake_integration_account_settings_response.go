// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountSettingsResponse Settings configured on the Snowflake integration account.
type SnowflakeIntegrationAccountSettingsResponse struct {
	// Identifier of the Snowflake account being monitored.
	SnowflakeAccountIdentifier string `json:"snowflake_account_identifier"`
	// Snowflake user Datadog authenticates as.
	Username string `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationAccountSettingsResponse instantiates a new SnowflakeIntegrationAccountSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationAccountSettingsResponse(snowflakeAccountIdentifier string, username string) *SnowflakeIntegrationAccountSettingsResponse {
	this := SnowflakeIntegrationAccountSettingsResponse{}
	this.SnowflakeAccountIdentifier = snowflakeAccountIdentifier
	this.Username = username
	return &this
}

// NewSnowflakeIntegrationAccountSettingsResponseWithDefaults instantiates a new SnowflakeIntegrationAccountSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationAccountSettingsResponseWithDefaults() *SnowflakeIntegrationAccountSettingsResponse {
	this := SnowflakeIntegrationAccountSettingsResponse{}
	return &this
}

// GetSnowflakeAccountIdentifier returns the SnowflakeAccountIdentifier field value.
func (o *SnowflakeIntegrationAccountSettingsResponse) GetSnowflakeAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SnowflakeAccountIdentifier
}

// GetSnowflakeAccountIdentifierOk returns a tuple with the SnowflakeAccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountSettingsResponse) GetSnowflakeAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnowflakeAccountIdentifier, true
}

// SetSnowflakeAccountIdentifier sets field value.
func (o *SnowflakeIntegrationAccountSettingsResponse) SetSnowflakeAccountIdentifier(v string) {
	o.SnowflakeAccountIdentifier = v
}

// GetUsername returns the Username field value.
func (o *SnowflakeIntegrationAccountSettingsResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountSettingsResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *SnowflakeIntegrationAccountSettingsResponse) SetUsername(v string) {
	o.Username = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationAccountSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["snowflake_account_identifier"] = o.SnowflakeAccountIdentifier
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeIntegrationAccountSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"snowflake_account_identifier", "username"})
	} else {
		return err
	}
	o.SnowflakeAccountIdentifier = *all.SnowflakeAccountIdentifier
	o.Username = *all.Username

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
