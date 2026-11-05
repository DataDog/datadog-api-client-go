// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate - Authentication for updating the Temporal Cloud OpenMetrics integration account. Exactly one method is set.
type TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate struct {
	TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdateAsTemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate is a convenience function that returns TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate wrapped in TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate.
func TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdateAsTemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate(v *TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate {
	return TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate{TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate)
	if err == nil {
		if obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate != nil && obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate.UnparsedObject == nil {
			jsonTemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate, _ := datadog.Marshal(obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate)
			if string(jsonTemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate) == "{}" { // empty struct
				obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate = nil
		}
	} else {
		obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate != nil {
		return obj.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
