// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountIntegration - Strongly-typed, per-integration configuration. Exactly one integration variant is set, selected by its `type`.
type IntegrationAccountIntegration struct {
	TwilioIntegration       *TwilioIntegration
	ElasticCloudIntegration *ElasticCloudIntegration

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TwilioIntegrationAsIntegrationAccountIntegration is a convenience function that returns TwilioIntegration wrapped in IntegrationAccountIntegration.
func TwilioIntegrationAsIntegrationAccountIntegration(v *TwilioIntegration) IntegrationAccountIntegration {
	return IntegrationAccountIntegration{TwilioIntegration: v}
}

// ElasticCloudIntegrationAsIntegrationAccountIntegration is a convenience function that returns ElasticCloudIntegration wrapped in IntegrationAccountIntegration.
func ElasticCloudIntegrationAsIntegrationAccountIntegration(v *ElasticCloudIntegration) IntegrationAccountIntegration {
	return IntegrationAccountIntegration{ElasticCloudIntegration: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IntegrationAccountIntegration) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TwilioIntegration
	err = datadog.Unmarshal(data, &obj.TwilioIntegration)
	if err == nil {
		if obj.TwilioIntegration != nil && obj.TwilioIntegration.UnparsedObject == nil {
			jsonTwilioIntegration, _ := datadog.Marshal(obj.TwilioIntegration)
			if string(jsonTwilioIntegration) == "{}" { // empty struct
				obj.TwilioIntegration = nil
			} else {
				match++
			}
		} else {
			obj.TwilioIntegration = nil
		}
	} else {
		obj.TwilioIntegration = nil
	}

	// try to unmarshal data into ElasticCloudIntegration
	err = datadog.Unmarshal(data, &obj.ElasticCloudIntegration)
	if err == nil {
		if obj.ElasticCloudIntegration != nil && obj.ElasticCloudIntegration.UnparsedObject == nil {
			jsonElasticCloudIntegration, _ := datadog.Marshal(obj.ElasticCloudIntegration)
			if string(jsonElasticCloudIntegration) == "{}" { // empty struct
				obj.ElasticCloudIntegration = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudIntegration = nil
		}
	} else {
		obj.ElasticCloudIntegration = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TwilioIntegration = nil
		obj.ElasticCloudIntegration = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IntegrationAccountIntegration) MarshalJSON() ([]byte, error) {
	if obj.TwilioIntegration != nil {
		return datadog.Marshal(&obj.TwilioIntegration)
	}

	if obj.ElasticCloudIntegration != nil {
		return datadog.Marshal(&obj.ElasticCloudIntegration)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IntegrationAccountIntegration) GetActualInstance() interface{} {
	if obj.TwilioIntegration != nil {
		return obj.TwilioIntegration
	}

	if obj.ElasticCloudIntegration != nil {
		return obj.ElasticCloudIntegration
	}

	// all schemas are nil
	return nil
}
