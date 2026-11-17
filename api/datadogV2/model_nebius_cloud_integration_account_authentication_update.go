// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountAuthenticationUpdate - Authentication for updating the Nebius Cloud integration account. Exactly one method is set.
type NebiusCloudIntegrationAccountAuthenticationUpdate struct {
	NebiusCloudIntegrationAccountServiceAccountAuthUpdate *NebiusCloudIntegrationAccountServiceAccountAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NebiusCloudIntegrationAccountServiceAccountAuthUpdateAsNebiusCloudIntegrationAccountAuthenticationUpdate is a convenience function that returns NebiusCloudIntegrationAccountServiceAccountAuthUpdate wrapped in NebiusCloudIntegrationAccountAuthenticationUpdate.
func NebiusCloudIntegrationAccountServiceAccountAuthUpdateAsNebiusCloudIntegrationAccountAuthenticationUpdate(v *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) NebiusCloudIntegrationAccountAuthenticationUpdate {
	return NebiusCloudIntegrationAccountAuthenticationUpdate{NebiusCloudIntegrationAccountServiceAccountAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NebiusCloudIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NebiusCloudIntegrationAccountServiceAccountAuthUpdate
	err = datadog.Unmarshal(data, &obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate)
	if err == nil {
		if obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate != nil && obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate.UnparsedObject == nil {
			jsonNebiusCloudIntegrationAccountServiceAccountAuthUpdate, _ := datadog.Marshal(obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate)
			if string(jsonNebiusCloudIntegrationAccountServiceAccountAuthUpdate) == "{}" { // empty struct
				obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate = nil
		}
	} else {
		obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NebiusCloudIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate != nil {
		return datadog.Marshal(&obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NebiusCloudIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate != nil {
		return obj.NebiusCloudIntegrationAccountServiceAccountAuthUpdate
	}

	// all schemas are nil
	return nil
}
