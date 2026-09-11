// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountPrivateKeyAuthType The authentication method type.
type SnowflakeIntegrationAccountPrivateKeyAuthType string

// List of SnowflakeIntegrationAccountPrivateKeyAuthType.
const (
	SNOWFLAKEINTEGRATIONACCOUNTPRIVATEKEYAUTHTYPE_SNOWFLAKE_PRIVATE_KEY SnowflakeIntegrationAccountPrivateKeyAuthType = "snowflake-private-key"
)

var allowedSnowflakeIntegrationAccountPrivateKeyAuthTypeEnumValues = []SnowflakeIntegrationAccountPrivateKeyAuthType{
	SNOWFLAKEINTEGRATIONACCOUNTPRIVATEKEYAUTHTYPE_SNOWFLAKE_PRIVATE_KEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SnowflakeIntegrationAccountPrivateKeyAuthType) GetAllowedValues() []SnowflakeIntegrationAccountPrivateKeyAuthType {
	return allowedSnowflakeIntegrationAccountPrivateKeyAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SnowflakeIntegrationAccountPrivateKeyAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SnowflakeIntegrationAccountPrivateKeyAuthType(value)
	return nil
}

// NewSnowflakeIntegrationAccountPrivateKeyAuthTypeFromValue returns a pointer to a valid SnowflakeIntegrationAccountPrivateKeyAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSnowflakeIntegrationAccountPrivateKeyAuthTypeFromValue(v string) (*SnowflakeIntegrationAccountPrivateKeyAuthType, error) {
	ev := SnowflakeIntegrationAccountPrivateKeyAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SnowflakeIntegrationAccountPrivateKeyAuthType: valid values are %v", v, allowedSnowflakeIntegrationAccountPrivateKeyAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SnowflakeIntegrationAccountPrivateKeyAuthType) IsValid() bool {
	for _, existing := range allowedSnowflakeIntegrationAccountPrivateKeyAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SnowflakeIntegrationAccountPrivateKeyAuthType value.
func (v SnowflakeIntegrationAccountPrivateKeyAuthType) Ptr() *SnowflakeIntegrationAccountPrivateKeyAuthType {
	return &v
}
