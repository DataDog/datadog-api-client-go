// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountIntegrationUpdate - Strongly-typed, per-integration partial configuration. Exactly one integration variant is set, selected by its `type`.
type IntegrationAccountIntegrationUpdate struct {
	TwilioIntegrationUpdate       *TwilioIntegrationUpdate
	ElasticCloudIntegrationUpdate *ElasticCloudIntegrationUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TwilioIntegrationUpdateAsIntegrationAccountIntegrationUpdate is a convenience function that returns TwilioIntegrationUpdate wrapped in IntegrationAccountIntegrationUpdate.
func TwilioIntegrationUpdateAsIntegrationAccountIntegrationUpdate(v *TwilioIntegrationUpdate) IntegrationAccountIntegrationUpdate {
	return IntegrationAccountIntegrationUpdate{TwilioIntegrationUpdate: v}
}

// ElasticCloudIntegrationUpdateAsIntegrationAccountIntegrationUpdate is a convenience function that returns ElasticCloudIntegrationUpdate wrapped in IntegrationAccountIntegrationUpdate.
func ElasticCloudIntegrationUpdateAsIntegrationAccountIntegrationUpdate(v *ElasticCloudIntegrationUpdate) IntegrationAccountIntegrationUpdate {
	return IntegrationAccountIntegrationUpdate{ElasticCloudIntegrationUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IntegrationAccountIntegrationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TwilioIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.TwilioIntegrationUpdate)
	if err == nil {
		if obj.TwilioIntegrationUpdate != nil && obj.TwilioIntegrationUpdate.UnparsedObject == nil {
			jsonTwilioIntegrationUpdate, _ := datadog.Marshal(obj.TwilioIntegrationUpdate)
			if string(jsonTwilioIntegrationUpdate) == "{}" { // empty struct
				obj.TwilioIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.TwilioIntegrationUpdate = nil
		}
	} else {
		obj.TwilioIntegrationUpdate = nil
	}

	// try to unmarshal data into ElasticCloudIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.ElasticCloudIntegrationUpdate)
	if err == nil {
		if obj.ElasticCloudIntegrationUpdate != nil && obj.ElasticCloudIntegrationUpdate.UnparsedObject == nil {
			jsonElasticCloudIntegrationUpdate, _ := datadog.Marshal(obj.ElasticCloudIntegrationUpdate)
			if string(jsonElasticCloudIntegrationUpdate) == "{}" { // empty struct
				obj.ElasticCloudIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudIntegrationUpdate = nil
		}
	} else {
		obj.ElasticCloudIntegrationUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TwilioIntegrationUpdate = nil
		obj.ElasticCloudIntegrationUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IntegrationAccountIntegrationUpdate) MarshalJSON() ([]byte, error) {
	if obj.TwilioIntegrationUpdate != nil {
		return datadog.Marshal(&obj.TwilioIntegrationUpdate)
	}

	if obj.ElasticCloudIntegrationUpdate != nil {
		return datadog.Marshal(&obj.ElasticCloudIntegrationUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IntegrationAccountIntegrationUpdate) GetActualInstance() interface{} {
	if obj.TwilioIntegrationUpdate != nil {
		return obj.TwilioIntegrationUpdate
	}

	if obj.ElasticCloudIntegrationUpdate != nil {
		return obj.ElasticCloudIntegrationUpdate
	}

	// all schemas are nil
	return nil
}
