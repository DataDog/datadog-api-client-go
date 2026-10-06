// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricEnrichmentTableProcessor - The `enrichment_table` processor enriches metrics with tags from a static CSV file or a Datadog reference table.
// It looks up a row using the metric name or a metric tag value. It then adds each column of the matching row as a
// metric tag, overwriting any existing tag with the same key. Exactly one of `file` or `reference_table` must be
// configured.
//
// **Supported pipeline types:** metrics
type ObservabilityPipelineMetricEnrichmentTableProcessor struct {
	ObservabilityPipelineMetricEnrichmentTableFileProcessor           *ObservabilityPipelineMetricEnrichmentTableFileProcessor
	ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor *ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineMetricEnrichmentTableFileProcessorAsObservabilityPipelineMetricEnrichmentTableProcessor is a convenience function that returns ObservabilityPipelineMetricEnrichmentTableFileProcessor wrapped in ObservabilityPipelineMetricEnrichmentTableProcessor.
func ObservabilityPipelineMetricEnrichmentTableFileProcessorAsObservabilityPipelineMetricEnrichmentTableProcessor(v *ObservabilityPipelineMetricEnrichmentTableFileProcessor) ObservabilityPipelineMetricEnrichmentTableProcessor {
	return ObservabilityPipelineMetricEnrichmentTableProcessor{ObservabilityPipelineMetricEnrichmentTableFileProcessor: v}
}

// ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessorAsObservabilityPipelineMetricEnrichmentTableProcessor is a convenience function that returns ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor wrapped in ObservabilityPipelineMetricEnrichmentTableProcessor.
func ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessorAsObservabilityPipelineMetricEnrichmentTableProcessor(v *ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor) ObservabilityPipelineMetricEnrichmentTableProcessor {
	return ObservabilityPipelineMetricEnrichmentTableProcessor{ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineMetricEnrichmentTableProcessor) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineMetricEnrichmentTableFileProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor)
	if err == nil {
		if obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor != nil && obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineMetricEnrichmentTableFileProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor)
			if string(jsonObservabilityPipelineMetricEnrichmentTableFileProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor)
	if err == nil {
		if obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor != nil && obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor)
			if string(jsonObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor = nil
		obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineMetricEnrichmentTableProcessor) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor)
	}

	if obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineMetricEnrichmentTableProcessor) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor != nil {
		return obj.ObservabilityPipelineMetricEnrichmentTableFileProcessor
	}

	if obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor != nil {
		return obj.ObservabilityPipelineMetricEnrichmentTableReferenceTableProcessor
	}

	// all schemas are nil
	return nil
}
