// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold - The minimal threshold to apply to the matching partition.
type ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold struct {
	ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes      *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes
	ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesAsObservabilityPipelineSpikeProtectionProcessorMinimalThreshold is a convenience function that returns ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes wrapped in ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold.
func ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesAsObservabilityPipelineSpikeProtectionProcessorMinimalThreshold(v *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold {
	return ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold{ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes: v}
}

// ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepAsObservabilityPipelineSpikeProtectionProcessorMinimalThreshold is a convenience function that returns ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep wrapped in ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold.
func ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepAsObservabilityPipelineSpikeProtectionProcessorMinimalThreshold(v *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep) ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold {
	return ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold{ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes)
	if err == nil {
		if obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes != nil && obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes.UnparsedObject == nil {
			jsonObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes, _ := datadog.Marshal(obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes)
			if string(jsonObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes) == "{}" { // empty struct
				obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes = nil
		}
	} else {
		obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes = nil
	}

	// try to unmarshal data into ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep)
	if err == nil {
		if obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep != nil && obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep.UnparsedObject == nil {
			jsonObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep, _ := datadog.Marshal(obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep)
			if string(jsonObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep) == "{}" { // empty struct
				obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep = nil
		}
	} else {
		obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes = nil
		obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes)
	}

	if obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes != nil {
		return obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes
	}

	if obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep != nil {
		return obj.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep
	}

	// all schemas are nil
	return nil
}
