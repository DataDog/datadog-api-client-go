// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerAttributes - The trigger definition for starting an investigation.
type TriggerAttributes struct {
	MonitorAlertTrigger         *MonitorAlertTrigger
	GeneralInvestigationTrigger *GeneralInvestigationTrigger

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MonitorAlertTriggerAsTriggerAttributes is a convenience function that returns MonitorAlertTrigger wrapped in TriggerAttributes.
func MonitorAlertTriggerAsTriggerAttributes(v *MonitorAlertTrigger) TriggerAttributes {
	return TriggerAttributes{MonitorAlertTrigger: v}
}

// GeneralInvestigationTriggerAsTriggerAttributes is a convenience function that returns GeneralInvestigationTrigger wrapped in TriggerAttributes.
func GeneralInvestigationTriggerAsTriggerAttributes(v *GeneralInvestigationTrigger) TriggerAttributes {
	return TriggerAttributes{GeneralInvestigationTrigger: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TriggerAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MonitorAlertTrigger
	err = datadog.Unmarshal(data, &obj.MonitorAlertTrigger)
	if err == nil {
		if obj.MonitorAlertTrigger != nil && obj.MonitorAlertTrigger.UnparsedObject == nil {
			jsonMonitorAlertTrigger, _ := datadog.Marshal(obj.MonitorAlertTrigger)
			if string(jsonMonitorAlertTrigger) == "{}" { // empty struct
				obj.MonitorAlertTrigger = nil
			} else {
				match++
			}
		} else {
			obj.MonitorAlertTrigger = nil
		}
	} else {
		obj.MonitorAlertTrigger = nil
	}

	// try to unmarshal data into GeneralInvestigationTrigger
	err = datadog.Unmarshal(data, &obj.GeneralInvestigationTrigger)
	if err == nil {
		if obj.GeneralInvestigationTrigger != nil && obj.GeneralInvestigationTrigger.UnparsedObject == nil {
			jsonGeneralInvestigationTrigger, _ := datadog.Marshal(obj.GeneralInvestigationTrigger)
			if string(jsonGeneralInvestigationTrigger) == "{}" { // empty struct
				obj.GeneralInvestigationTrigger = nil
			} else {
				match++
			}
		} else {
			obj.GeneralInvestigationTrigger = nil
		}
	} else {
		obj.GeneralInvestigationTrigger = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MonitorAlertTrigger = nil
		obj.GeneralInvestigationTrigger = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TriggerAttributes) MarshalJSON() ([]byte, error) {
	if obj.MonitorAlertTrigger != nil {
		return datadog.Marshal(&obj.MonitorAlertTrigger)
	}

	if obj.GeneralInvestigationTrigger != nil {
		return datadog.Marshal(&obj.GeneralInvestigationTrigger)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TriggerAttributes) GetActualInstance() interface{} {
	if obj.MonitorAlertTrigger != nil {
		return obj.MonitorAlertTrigger
	}

	if obj.GeneralInvestigationTrigger != nil {
		return obj.GeneralInvestigationTrigger
	}

	// all schemas are nil
	return nil
}
