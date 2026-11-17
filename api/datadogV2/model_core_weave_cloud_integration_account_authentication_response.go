// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoreWeaveCloudIntegrationAccountAuthenticationResponse - Authentication configured on the CoreWeave integration account.
type CoreWeaveCloudIntegrationAccountAuthenticationResponse struct {
	CoreWeaveCloudIntegrationAccountAPITokenAuthResponse *CoreWeaveCloudIntegrationAccountAPITokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CoreWeaveCloudIntegrationAccountAPITokenAuthResponseAsCoreWeaveCloudIntegrationAccountAuthenticationResponse is a convenience function that returns CoreWeaveCloudIntegrationAccountAPITokenAuthResponse wrapped in CoreWeaveCloudIntegrationAccountAuthenticationResponse.
func CoreWeaveCloudIntegrationAccountAPITokenAuthResponseAsCoreWeaveCloudIntegrationAccountAuthenticationResponse(v *CoreWeaveCloudIntegrationAccountAPITokenAuthResponse) CoreWeaveCloudIntegrationAccountAuthenticationResponse {
	return CoreWeaveCloudIntegrationAccountAuthenticationResponse{CoreWeaveCloudIntegrationAccountAPITokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CoreWeaveCloudIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CoreWeaveCloudIntegrationAccountAPITokenAuthResponse
	err = datadog.Unmarshal(data, &obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse)
	if err == nil {
		if obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse != nil && obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse.UnparsedObject == nil {
			jsonCoreWeaveCloudIntegrationAccountAPITokenAuthResponse, _ := datadog.Marshal(obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse)
			if string(jsonCoreWeaveCloudIntegrationAccountAPITokenAuthResponse) == "{}" { // empty struct
				obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse = nil
		}
	} else {
		obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CoreWeaveCloudIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse != nil {
		return datadog.Marshal(&obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CoreWeaveCloudIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse != nil {
		return obj.CoreWeaveCloudIntegrationAccountAPITokenAuthResponse
	}

	// all schemas are nil
	return nil
}
