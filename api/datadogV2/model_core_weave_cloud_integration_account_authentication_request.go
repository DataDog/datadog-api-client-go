// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoreWeaveCloudIntegrationAccountAuthenticationRequest - Authentication for creating the CoreWeave integration account. Exactly one method is set.
type CoreWeaveCloudIntegrationAccountAuthenticationRequest struct {
	CoreWeaveCloudIntegrationAccountAPITokenAuthRequest *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CoreWeaveCloudIntegrationAccountAPITokenAuthRequestAsCoreWeaveCloudIntegrationAccountAuthenticationRequest is a convenience function that returns CoreWeaveCloudIntegrationAccountAPITokenAuthRequest wrapped in CoreWeaveCloudIntegrationAccountAuthenticationRequest.
func CoreWeaveCloudIntegrationAccountAPITokenAuthRequestAsCoreWeaveCloudIntegrationAccountAuthenticationRequest(v *CoreWeaveCloudIntegrationAccountAPITokenAuthRequest) CoreWeaveCloudIntegrationAccountAuthenticationRequest {
	return CoreWeaveCloudIntegrationAccountAuthenticationRequest{CoreWeaveCloudIntegrationAccountAPITokenAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CoreWeaveCloudIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CoreWeaveCloudIntegrationAccountAPITokenAuthRequest
	err = datadog.Unmarshal(data, &obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest)
	if err == nil {
		if obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest != nil && obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest.UnparsedObject == nil {
			jsonCoreWeaveCloudIntegrationAccountAPITokenAuthRequest, _ := datadog.Marshal(obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest)
			if string(jsonCoreWeaveCloudIntegrationAccountAPITokenAuthRequest) == "{}" { // empty struct
				obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest = nil
		}
	} else {
		obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CoreWeaveCloudIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest != nil {
		return datadog.Marshal(&obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CoreWeaveCloudIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest != nil {
		return obj.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest
	}

	// all schemas are nil
	return nil
}
