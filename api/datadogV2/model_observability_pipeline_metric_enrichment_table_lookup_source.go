// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricEnrichmentTableLookupSource - Specifies the source of the key value used for metric enrichment table lookups.
// The lookup key can be either the metric name or a metric tag.
type ObservabilityPipelineMetricEnrichmentTableLookupSource struct {
	ObservabilityPipelineMetricEnrichmentTableMetricNameLookup *ObservabilityPipelineMetricEnrichmentTableMetricNameLookup
	ObservabilityPipelineMetricEnrichmentTableTagLookup        *ObservabilityPipelineMetricEnrichmentTableTagLookup

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineMetricEnrichmentTableMetricNameLookupAsObservabilityPipelineMetricEnrichmentTableLookupSource is a convenience function that returns ObservabilityPipelineMetricEnrichmentTableMetricNameLookup wrapped in ObservabilityPipelineMetricEnrichmentTableLookupSource.
func ObservabilityPipelineMetricEnrichmentTableMetricNameLookupAsObservabilityPipelineMetricEnrichmentTableLookupSource(v *ObservabilityPipelineMetricEnrichmentTableMetricNameLookup) ObservabilityPipelineMetricEnrichmentTableLookupSource {
	return ObservabilityPipelineMetricEnrichmentTableLookupSource{ObservabilityPipelineMetricEnrichmentTableMetricNameLookup: v}
}

// ObservabilityPipelineMetricEnrichmentTableTagLookupAsObservabilityPipelineMetricEnrichmentTableLookupSource is a convenience function that returns ObservabilityPipelineMetricEnrichmentTableTagLookup wrapped in ObservabilityPipelineMetricEnrichmentTableLookupSource.
func ObservabilityPipelineMetricEnrichmentTableTagLookupAsObservabilityPipelineMetricEnrichmentTableLookupSource(v *ObservabilityPipelineMetricEnrichmentTableTagLookup) ObservabilityPipelineMetricEnrichmentTableLookupSource {
	return ObservabilityPipelineMetricEnrichmentTableLookupSource{ObservabilityPipelineMetricEnrichmentTableTagLookup: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineMetricEnrichmentTableLookupSource) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineMetricEnrichmentTableMetricNameLookup
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup)
	if err == nil {
		if obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup != nil && obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup.UnparsedObject == nil {
			jsonObservabilityPipelineMetricEnrichmentTableMetricNameLookup, _ := datadog.Marshal(obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup)
			if string(jsonObservabilityPipelineMetricEnrichmentTableMetricNameLookup) == "{}" { // empty struct
				obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup = nil
		}
	} else {
		obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup = nil
	}

	// try to unmarshal data into ObservabilityPipelineMetricEnrichmentTableTagLookup
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineMetricEnrichmentTableTagLookup)
	if err == nil {
		if obj.ObservabilityPipelineMetricEnrichmentTableTagLookup != nil && obj.ObservabilityPipelineMetricEnrichmentTableTagLookup.UnparsedObject == nil {
			jsonObservabilityPipelineMetricEnrichmentTableTagLookup, _ := datadog.Marshal(obj.ObservabilityPipelineMetricEnrichmentTableTagLookup)
			if string(jsonObservabilityPipelineMetricEnrichmentTableTagLookup) == "{}" { // empty struct
				obj.ObservabilityPipelineMetricEnrichmentTableTagLookup = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineMetricEnrichmentTableTagLookup = nil
		}
	} else {
		obj.ObservabilityPipelineMetricEnrichmentTableTagLookup = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup = nil
		obj.ObservabilityPipelineMetricEnrichmentTableTagLookup = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineMetricEnrichmentTableLookupSource) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup)
	}

	if obj.ObservabilityPipelineMetricEnrichmentTableTagLookup != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineMetricEnrichmentTableTagLookup)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineMetricEnrichmentTableLookupSource) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup != nil {
		return obj.ObservabilityPipelineMetricEnrichmentTableMetricNameLookup
	}

	if obj.ObservabilityPipelineMetricEnrichmentTableTagLookup != nil {
		return obj.ObservabilityPipelineMetricEnrichmentTableTagLookup
	}

	// all schemas are nil
	return nil
}
