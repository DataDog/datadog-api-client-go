// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountAuthenticationRequest - Authentication for creating the Nebius Cloud integration account. Exactly one method is set.
type NebiusCloudIntegrationAccountAuthenticationRequest struct {
	NebiusCloudIntegrationAccountServiceAccountAuthRequest *NebiusCloudIntegrationAccountServiceAccountAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NebiusCloudIntegrationAccountServiceAccountAuthRequestAsNebiusCloudIntegrationAccountAuthenticationRequest is a convenience function that returns NebiusCloudIntegrationAccountServiceAccountAuthRequest wrapped in NebiusCloudIntegrationAccountAuthenticationRequest.
func NebiusCloudIntegrationAccountServiceAccountAuthRequestAsNebiusCloudIntegrationAccountAuthenticationRequest(v *NebiusCloudIntegrationAccountServiceAccountAuthRequest) NebiusCloudIntegrationAccountAuthenticationRequest {
	return NebiusCloudIntegrationAccountAuthenticationRequest{NebiusCloudIntegrationAccountServiceAccountAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NebiusCloudIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NebiusCloudIntegrationAccountServiceAccountAuthRequest
	err = datadog.Unmarshal(data, &obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest)
	if err == nil {
		if obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest != nil && obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest.UnparsedObject == nil {
			jsonNebiusCloudIntegrationAccountServiceAccountAuthRequest, _ := datadog.Marshal(obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest)
			if string(jsonNebiusCloudIntegrationAccountServiceAccountAuthRequest) == "{}" { // empty struct
				obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest = nil
		}
	} else {
		obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NebiusCloudIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest != nil {
		return datadog.Marshal(&obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NebiusCloudIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest != nil {
		return obj.NebiusCloudIntegrationAccountServiceAccountAuthRequest
	}

	// all schemas are nil
	return nil
}
