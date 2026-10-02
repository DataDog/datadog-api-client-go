// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateMetricV2RequestDataAttributes - Configuration for the new metric. Supply either numerator_aggregation or percentile_aggregation. A denominator_aggregation requires numerator_aggregation. Omit unused aggregation fields; do not send them as null.
type ExperimentsCreateMetricV2RequestDataAttributes struct {
	ExperimentsCreateMetricNumeratorAttributes  *ExperimentsCreateMetricNumeratorAttributes
	ExperimentsCreateMetricPercentileAttributes *ExperimentsCreateMetricPercentileAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ExperimentsCreateMetricNumeratorAttributesAsExperimentsCreateMetricV2RequestDataAttributes is a convenience function that returns ExperimentsCreateMetricNumeratorAttributes wrapped in ExperimentsCreateMetricV2RequestDataAttributes.
func ExperimentsCreateMetricNumeratorAttributesAsExperimentsCreateMetricV2RequestDataAttributes(v *ExperimentsCreateMetricNumeratorAttributes) ExperimentsCreateMetricV2RequestDataAttributes {
	return ExperimentsCreateMetricV2RequestDataAttributes{ExperimentsCreateMetricNumeratorAttributes: v}
}

// ExperimentsCreateMetricPercentileAttributesAsExperimentsCreateMetricV2RequestDataAttributes is a convenience function that returns ExperimentsCreateMetricPercentileAttributes wrapped in ExperimentsCreateMetricV2RequestDataAttributes.
func ExperimentsCreateMetricPercentileAttributesAsExperimentsCreateMetricV2RequestDataAttributes(v *ExperimentsCreateMetricPercentileAttributes) ExperimentsCreateMetricV2RequestDataAttributes {
	return ExperimentsCreateMetricV2RequestDataAttributes{ExperimentsCreateMetricPercentileAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ExperimentsCreateMetricV2RequestDataAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExperimentsCreateMetricNumeratorAttributes
	err = datadog.Unmarshal(data, &obj.ExperimentsCreateMetricNumeratorAttributes)
	if err == nil {
		if obj.ExperimentsCreateMetricNumeratorAttributes != nil && obj.ExperimentsCreateMetricNumeratorAttributes.UnparsedObject == nil {
			jsonExperimentsCreateMetricNumeratorAttributes, _ := datadog.Marshal(obj.ExperimentsCreateMetricNumeratorAttributes)
			if string(jsonExperimentsCreateMetricNumeratorAttributes) == "{}" { // empty struct
				obj.ExperimentsCreateMetricNumeratorAttributes = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsCreateMetricNumeratorAttributes = nil
		}
	} else {
		obj.ExperimentsCreateMetricNumeratorAttributes = nil
	}

	// try to unmarshal data into ExperimentsCreateMetricPercentileAttributes
	err = datadog.Unmarshal(data, &obj.ExperimentsCreateMetricPercentileAttributes)
	if err == nil {
		if obj.ExperimentsCreateMetricPercentileAttributes != nil && obj.ExperimentsCreateMetricPercentileAttributes.UnparsedObject == nil {
			jsonExperimentsCreateMetricPercentileAttributes, _ := datadog.Marshal(obj.ExperimentsCreateMetricPercentileAttributes)
			if string(jsonExperimentsCreateMetricPercentileAttributes) == "{}" { // empty struct
				obj.ExperimentsCreateMetricPercentileAttributes = nil
			} else {
				match++
			}
		} else {
			obj.ExperimentsCreateMetricPercentileAttributes = nil
		}
	} else {
		obj.ExperimentsCreateMetricPercentileAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ExperimentsCreateMetricNumeratorAttributes = nil
		obj.ExperimentsCreateMetricPercentileAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ExperimentsCreateMetricV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	if obj.ExperimentsCreateMetricNumeratorAttributes != nil {
		return datadog.Marshal(&obj.ExperimentsCreateMetricNumeratorAttributes)
	}

	if obj.ExperimentsCreateMetricPercentileAttributes != nil {
		return datadog.Marshal(&obj.ExperimentsCreateMetricPercentileAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ExperimentsCreateMetricV2RequestDataAttributes) GetActualInstance() interface{} {
	if obj.ExperimentsCreateMetricNumeratorAttributes != nil {
		return obj.ExperimentsCreateMetricNumeratorAttributes
	}

	if obj.ExperimentsCreateMetricPercentileAttributes != nil {
		return obj.ExperimentsCreateMetricPercentileAttributes
	}

	// all schemas are nil
	return nil
}
