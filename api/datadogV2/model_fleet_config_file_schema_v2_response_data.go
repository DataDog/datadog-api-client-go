// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetConfigFileSchemaV2ResponseData - The schema resolved for the requested configuration file.
type FleetConfigFileSchemaV2ResponseData struct {
	FleetIntegrationSchemaDetailV2 *FleetIntegrationSchemaDetailV2
	FleetConfigFileSchemaV2        *FleetConfigFileSchemaV2

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FleetIntegrationSchemaDetailV2AsFleetConfigFileSchemaV2ResponseData is a convenience function that returns FleetIntegrationSchemaDetailV2 wrapped in FleetConfigFileSchemaV2ResponseData.
func FleetIntegrationSchemaDetailV2AsFleetConfigFileSchemaV2ResponseData(v *FleetIntegrationSchemaDetailV2) FleetConfigFileSchemaV2ResponseData {
	return FleetConfigFileSchemaV2ResponseData{FleetIntegrationSchemaDetailV2: v}
}

// FleetConfigFileSchemaV2AsFleetConfigFileSchemaV2ResponseData is a convenience function that returns FleetConfigFileSchemaV2 wrapped in FleetConfigFileSchemaV2ResponseData.
func FleetConfigFileSchemaV2AsFleetConfigFileSchemaV2ResponseData(v *FleetConfigFileSchemaV2) FleetConfigFileSchemaV2ResponseData {
	return FleetConfigFileSchemaV2ResponseData{FleetConfigFileSchemaV2: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FleetConfigFileSchemaV2ResponseData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FleetIntegrationSchemaDetailV2
	err = datadog.Unmarshal(data, &obj.FleetIntegrationSchemaDetailV2)
	if err == nil {
		if obj.FleetIntegrationSchemaDetailV2 != nil && obj.FleetIntegrationSchemaDetailV2.UnparsedObject == nil {
			jsonFleetIntegrationSchemaDetailV2, _ := datadog.Marshal(obj.FleetIntegrationSchemaDetailV2)
			if string(jsonFleetIntegrationSchemaDetailV2) == "{}" { // empty struct
				obj.FleetIntegrationSchemaDetailV2 = nil
			} else {
				match++
			}
		} else {
			obj.FleetIntegrationSchemaDetailV2 = nil
		}
	} else {
		obj.FleetIntegrationSchemaDetailV2 = nil
	}

	// try to unmarshal data into FleetConfigFileSchemaV2
	err = datadog.Unmarshal(data, &obj.FleetConfigFileSchemaV2)
	if err == nil {
		if obj.FleetConfigFileSchemaV2 != nil && obj.FleetConfigFileSchemaV2.UnparsedObject == nil {
			jsonFleetConfigFileSchemaV2, _ := datadog.Marshal(obj.FleetConfigFileSchemaV2)
			if string(jsonFleetConfigFileSchemaV2) == "{}" { // empty struct
				obj.FleetConfigFileSchemaV2 = nil
			} else {
				match++
			}
		} else {
			obj.FleetConfigFileSchemaV2 = nil
		}
	} else {
		obj.FleetConfigFileSchemaV2 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FleetIntegrationSchemaDetailV2 = nil
		obj.FleetConfigFileSchemaV2 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FleetConfigFileSchemaV2ResponseData) MarshalJSON() ([]byte, error) {
	if obj.FleetIntegrationSchemaDetailV2 != nil {
		return datadog.Marshal(&obj.FleetIntegrationSchemaDetailV2)
	}

	if obj.FleetConfigFileSchemaV2 != nil {
		return datadog.Marshal(&obj.FleetConfigFileSchemaV2)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FleetConfigFileSchemaV2ResponseData) GetActualInstance() interface{} {
	if obj.FleetIntegrationSchemaDetailV2 != nil {
		return obj.FleetIntegrationSchemaDetailV2
	}

	if obj.FleetConfigFileSchemaV2 != nil {
		return obj.FleetConfigFileSchemaV2
	}

	// all schemas are nil
	return nil
}
