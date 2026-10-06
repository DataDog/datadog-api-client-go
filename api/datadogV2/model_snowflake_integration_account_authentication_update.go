// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountAuthenticationUpdate - Authentication for updating the Snowflake integration account. Exactly one method is set.
type SnowflakeIntegrationAccountAuthenticationUpdate struct {
	SnowflakeIntegrationAccountPrivateKeyAuthUpdate *SnowflakeIntegrationAccountPrivateKeyAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SnowflakeIntegrationAccountPrivateKeyAuthUpdateAsSnowflakeIntegrationAccountAuthenticationUpdate is a convenience function that returns SnowflakeIntegrationAccountPrivateKeyAuthUpdate wrapped in SnowflakeIntegrationAccountAuthenticationUpdate.
func SnowflakeIntegrationAccountPrivateKeyAuthUpdateAsSnowflakeIntegrationAccountAuthenticationUpdate(v *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) SnowflakeIntegrationAccountAuthenticationUpdate {
	return SnowflakeIntegrationAccountAuthenticationUpdate{SnowflakeIntegrationAccountPrivateKeyAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SnowflakeIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SnowflakeIntegrationAccountPrivateKeyAuthUpdate
	err = datadog.Unmarshal(data, &obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate)
	if err == nil {
		if obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate != nil && obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate.UnparsedObject == nil {
			jsonSnowflakeIntegrationAccountPrivateKeyAuthUpdate, _ := datadog.Marshal(obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate)
			if string(jsonSnowflakeIntegrationAccountPrivateKeyAuthUpdate) == "{}" { // empty struct
				obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate = nil
		}
	} else {
		obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SnowflakeIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate != nil {
		return datadog.Marshal(&obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SnowflakeIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate != nil {
		return obj.SnowflakeIntegrationAccountPrivateKeyAuthUpdate
	}

	// all schemas are nil
	return nil
}
