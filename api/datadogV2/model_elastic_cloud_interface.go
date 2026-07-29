// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudInterface - Elastic Cloud interface (source-type). Exactly one interface variant is set, selected by its `type`.
type ElasticCloudInterface struct {
	ElasticCloudMonitoringInterface *ElasticCloudMonitoringInterface
	ElasticCloudCcmInterface        *ElasticCloudCcmInterface

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ElasticCloudMonitoringInterfaceAsElasticCloudInterface is a convenience function that returns ElasticCloudMonitoringInterface wrapped in ElasticCloudInterface.
func ElasticCloudMonitoringInterfaceAsElasticCloudInterface(v *ElasticCloudMonitoringInterface) ElasticCloudInterface {
	return ElasticCloudInterface{ElasticCloudMonitoringInterface: v}
}

// ElasticCloudCcmInterfaceAsElasticCloudInterface is a convenience function that returns ElasticCloudCcmInterface wrapped in ElasticCloudInterface.
func ElasticCloudCcmInterfaceAsElasticCloudInterface(v *ElasticCloudCcmInterface) ElasticCloudInterface {
	return ElasticCloudInterface{ElasticCloudCcmInterface: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ElasticCloudInterface) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ElasticCloudMonitoringInterface
	err = datadog.Unmarshal(data, &obj.ElasticCloudMonitoringInterface)
	if err == nil {
		if obj.ElasticCloudMonitoringInterface != nil && obj.ElasticCloudMonitoringInterface.UnparsedObject == nil {
			jsonElasticCloudMonitoringInterface, _ := datadog.Marshal(obj.ElasticCloudMonitoringInterface)
			if string(jsonElasticCloudMonitoringInterface) == "{}" { // empty struct
				obj.ElasticCloudMonitoringInterface = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudMonitoringInterface = nil
		}
	} else {
		obj.ElasticCloudMonitoringInterface = nil
	}

	// try to unmarshal data into ElasticCloudCcmInterface
	err = datadog.Unmarshal(data, &obj.ElasticCloudCcmInterface)
	if err == nil {
		if obj.ElasticCloudCcmInterface != nil && obj.ElasticCloudCcmInterface.UnparsedObject == nil {
			jsonElasticCloudCcmInterface, _ := datadog.Marshal(obj.ElasticCloudCcmInterface)
			if string(jsonElasticCloudCcmInterface) == "{}" { // empty struct
				obj.ElasticCloudCcmInterface = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudCcmInterface = nil
		}
	} else {
		obj.ElasticCloudCcmInterface = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ElasticCloudMonitoringInterface = nil
		obj.ElasticCloudCcmInterface = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ElasticCloudInterface) MarshalJSON() ([]byte, error) {
	if obj.ElasticCloudMonitoringInterface != nil {
		return datadog.Marshal(&obj.ElasticCloudMonitoringInterface)
	}

	if obj.ElasticCloudCcmInterface != nil {
		return datadog.Marshal(&obj.ElasticCloudCcmInterface)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ElasticCloudInterface) GetActualInstance() interface{} {
	if obj.ElasticCloudMonitoringInterface != nil {
		return obj.ElasticCloudMonitoringInterface
	}

	if obj.ElasticCloudCcmInterface != nil {
		return obj.ElasticCloudCcmInterface
	}

	// all schemas are nil
	return nil
}
