// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse - Authentication configured on the Temporal Cloud OpenMetrics integration account.
type TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse struct {
	TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponseAsTemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse is a convenience function that returns TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse wrapped in TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse.
func TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponseAsTemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse(v *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse) TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse {
	return TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse{TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse != nil && obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonTemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse)
			if string(jsonTemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}
