// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationAccountAuthenticationUpdate - Authentication for updating the Elastic Cloud integration account. Exactly one method is set.
type ElasticCloudIntegrationAccountAuthenticationUpdate struct {
	ElasticCloudIntegrationAccountBasicAuthUpdate *ElasticCloudIntegrationAccountBasicAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ElasticCloudIntegrationAccountBasicAuthUpdateAsElasticCloudIntegrationAccountAuthenticationUpdate is a convenience function that returns ElasticCloudIntegrationAccountBasicAuthUpdate wrapped in ElasticCloudIntegrationAccountAuthenticationUpdate.
func ElasticCloudIntegrationAccountBasicAuthUpdateAsElasticCloudIntegrationAccountAuthenticationUpdate(v *ElasticCloudIntegrationAccountBasicAuthUpdate) ElasticCloudIntegrationAccountAuthenticationUpdate {
	return ElasticCloudIntegrationAccountAuthenticationUpdate{ElasticCloudIntegrationAccountBasicAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ElasticCloudIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ElasticCloudIntegrationAccountBasicAuthUpdate
	err = datadog.Unmarshal(data, &obj.ElasticCloudIntegrationAccountBasicAuthUpdate)
	if err == nil {
		if obj.ElasticCloudIntegrationAccountBasicAuthUpdate != nil && obj.ElasticCloudIntegrationAccountBasicAuthUpdate.UnparsedObject == nil {
			jsonElasticCloudIntegrationAccountBasicAuthUpdate, _ := datadog.Marshal(obj.ElasticCloudIntegrationAccountBasicAuthUpdate)
			if string(jsonElasticCloudIntegrationAccountBasicAuthUpdate) == "{}" { // empty struct
				obj.ElasticCloudIntegrationAccountBasicAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudIntegrationAccountBasicAuthUpdate = nil
		}
	} else {
		obj.ElasticCloudIntegrationAccountBasicAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ElasticCloudIntegrationAccountBasicAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ElasticCloudIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.ElasticCloudIntegrationAccountBasicAuthUpdate != nil {
		return datadog.Marshal(&obj.ElasticCloudIntegrationAccountBasicAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ElasticCloudIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.ElasticCloudIntegrationAccountBasicAuthUpdate != nil {
		return obj.ElasticCloudIntegrationAccountBasicAuthUpdate
	}

	// all schemas are nil
	return nil
}
