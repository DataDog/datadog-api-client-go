// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudInterfaceUpdate - Partial Elastic Cloud interface for updates. Exactly one interface variant is set, selected by its `type`.
type ElasticCloudInterfaceUpdate struct {
	ElasticCloudMonitoringInterfaceUpdate *ElasticCloudMonitoringInterfaceUpdate
	ElasticCloudCcmInterfaceUpdate        *ElasticCloudCcmInterfaceUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ElasticCloudMonitoringInterfaceUpdateAsElasticCloudInterfaceUpdate is a convenience function that returns ElasticCloudMonitoringInterfaceUpdate wrapped in ElasticCloudInterfaceUpdate.
func ElasticCloudMonitoringInterfaceUpdateAsElasticCloudInterfaceUpdate(v *ElasticCloudMonitoringInterfaceUpdate) ElasticCloudInterfaceUpdate {
	return ElasticCloudInterfaceUpdate{ElasticCloudMonitoringInterfaceUpdate: v}
}

// ElasticCloudCcmInterfaceUpdateAsElasticCloudInterfaceUpdate is a convenience function that returns ElasticCloudCcmInterfaceUpdate wrapped in ElasticCloudInterfaceUpdate.
func ElasticCloudCcmInterfaceUpdateAsElasticCloudInterfaceUpdate(v *ElasticCloudCcmInterfaceUpdate) ElasticCloudInterfaceUpdate {
	return ElasticCloudInterfaceUpdate{ElasticCloudCcmInterfaceUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ElasticCloudInterfaceUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ElasticCloudMonitoringInterfaceUpdate
	err = datadog.Unmarshal(data, &obj.ElasticCloudMonitoringInterfaceUpdate)
	if err == nil {
		if obj.ElasticCloudMonitoringInterfaceUpdate != nil && obj.ElasticCloudMonitoringInterfaceUpdate.UnparsedObject == nil {
			jsonElasticCloudMonitoringInterfaceUpdate, _ := datadog.Marshal(obj.ElasticCloudMonitoringInterfaceUpdate)
			if string(jsonElasticCloudMonitoringInterfaceUpdate) == "{}" { // empty struct
				obj.ElasticCloudMonitoringInterfaceUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudMonitoringInterfaceUpdate = nil
		}
	} else {
		obj.ElasticCloudMonitoringInterfaceUpdate = nil
	}

	// try to unmarshal data into ElasticCloudCcmInterfaceUpdate
	err = datadog.Unmarshal(data, &obj.ElasticCloudCcmInterfaceUpdate)
	if err == nil {
		if obj.ElasticCloudCcmInterfaceUpdate != nil && obj.ElasticCloudCcmInterfaceUpdate.UnparsedObject == nil {
			jsonElasticCloudCcmInterfaceUpdate, _ := datadog.Marshal(obj.ElasticCloudCcmInterfaceUpdate)
			if string(jsonElasticCloudCcmInterfaceUpdate) == "{}" { // empty struct
				obj.ElasticCloudCcmInterfaceUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudCcmInterfaceUpdate = nil
		}
	} else {
		obj.ElasticCloudCcmInterfaceUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ElasticCloudMonitoringInterfaceUpdate = nil
		obj.ElasticCloudCcmInterfaceUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ElasticCloudInterfaceUpdate) MarshalJSON() ([]byte, error) {
	if obj.ElasticCloudMonitoringInterfaceUpdate != nil {
		return datadog.Marshal(&obj.ElasticCloudMonitoringInterfaceUpdate)
	}

	if obj.ElasticCloudCcmInterfaceUpdate != nil {
		return datadog.Marshal(&obj.ElasticCloudCcmInterfaceUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ElasticCloudInterfaceUpdate) GetActualInstance() interface{} {
	if obj.ElasticCloudMonitoringInterfaceUpdate != nil {
		return obj.ElasticCloudMonitoringInterfaceUpdate
	}

	if obj.ElasticCloudCcmInterfaceUpdate != nil {
		return obj.ElasticCloudCcmInterfaceUpdate
	}

	// all schemas are nil
	return nil
}
