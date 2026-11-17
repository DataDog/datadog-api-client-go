// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoreWeaveCloudIntegrationAccountAuthenticationUpdate - Authentication for updating the CoreWeave integration account. Exactly one method is set.
type CoreWeaveCloudIntegrationAccountAuthenticationUpdate struct {
	CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CoreWeaveCloudIntegrationAccountAPITokenAuthUpdateAsCoreWeaveCloudIntegrationAccountAuthenticationUpdate is a convenience function that returns CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate wrapped in CoreWeaveCloudIntegrationAccountAuthenticationUpdate.
func CoreWeaveCloudIntegrationAccountAPITokenAuthUpdateAsCoreWeaveCloudIntegrationAccountAuthenticationUpdate(v *CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) CoreWeaveCloudIntegrationAccountAuthenticationUpdate {
	return CoreWeaveCloudIntegrationAccountAuthenticationUpdate{CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CoreWeaveCloudIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate)
	if err == nil {
		if obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate != nil && obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate.UnparsedObject == nil {
			jsonCoreWeaveCloudIntegrationAccountAPITokenAuthUpdate, _ := datadog.Marshal(obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate)
			if string(jsonCoreWeaveCloudIntegrationAccountAPITokenAuthUpdate) == "{}" { // empty struct
				obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate = nil
		}
	} else {
		obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CoreWeaveCloudIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate != nil {
		return datadog.Marshal(&obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CoreWeaveCloudIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate != nil {
		return obj.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
