// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigProcessorItem - A processor for the pipeline.
type ObservabilityPipelineConfigProcessorItem struct {
	ObservabilityPipelineFilterProcessor          *ObservabilityPipelineFilterProcessor
	ObservabilityPipelineParseJSONProcessor       *ObservabilityPipelineParseJSONProcessor
	ObservabilityPipelineQuotaProcessor           *ObservabilityPipelineQuotaProcessor
	ObservabilityPipelineAddFieldsProcessor       *ObservabilityPipelineAddFieldsProcessor
	ObservabilityPipelineRemoveFieldsProcessor    *ObservabilityPipelineRemoveFieldsProcessor
	ObservabilityPipelineRenameFieldsProcessor    *ObservabilityPipelineRenameFieldsProcessor
	ObservabilityPipelineOcsfMapperProcessor      *ObservabilityPipelineOcsfMapperProcessor
	ObservabilityPipelineAddEnvVarsProcessor      *ObservabilityPipelineAddEnvVarsProcessor
	ObservabilityPipelineDedupeProcessor          *ObservabilityPipelineDedupeProcessor
	ObservabilityPipelineEnrichmentTableProcessor *ObservabilityPipelineEnrichmentTableProcessor
	ObservabilityPipelineReduceProcessor          *ObservabilityPipelineReduceProcessor
	ObservabilityPipelineThrottleProcessor        *ObservabilityPipelineThrottleProcessor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineFilterProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineFilterProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineFilterProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineFilterProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineFilterProcessor: v}
}

// ObservabilityPipelineParseJSONProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineParseJSONProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineParseJSONProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineParseJSONProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineParseJSONProcessor: v}
}

// ObservabilityPipelineQuotaProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineQuotaProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineQuotaProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineQuotaProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineQuotaProcessor: v}
}

// ObservabilityPipelineAddFieldsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineAddFieldsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineAddFieldsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineAddFieldsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineAddFieldsProcessor: v}
}

// ObservabilityPipelineRemoveFieldsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineRemoveFieldsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineRemoveFieldsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineRemoveFieldsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineRemoveFieldsProcessor: v}
}

// ObservabilityPipelineRenameFieldsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineRenameFieldsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineRenameFieldsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineRenameFieldsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineRenameFieldsProcessor: v}
}

// ObservabilityPipelineOcsfMapperProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineOcsfMapperProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineOcsfMapperProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineOcsfMapperProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineOcsfMapperProcessor: v}
}

// ObservabilityPipelineAddEnvVarsProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineAddEnvVarsProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineAddEnvVarsProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineAddEnvVarsProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineAddEnvVarsProcessor: v}
}

// ObservabilityPipelineDedupeProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineDedupeProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineDedupeProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineDedupeProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineDedupeProcessor: v}
}

// ObservabilityPipelineEnrichmentTableProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineEnrichmentTableProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineEnrichmentTableProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineEnrichmentTableProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineEnrichmentTableProcessor: v}
}

// ObservabilityPipelineReduceProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineReduceProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineReduceProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineReduceProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineReduceProcessor: v}
}

// ObservabilityPipelineThrottleProcessorAsObservabilityPipelineConfigProcessorItem is a convenience function that returns ObservabilityPipelineThrottleProcessor wrapped in ObservabilityPipelineConfigProcessorItem.
func ObservabilityPipelineThrottleProcessorAsObservabilityPipelineConfigProcessorItem(v *ObservabilityPipelineThrottleProcessor) ObservabilityPipelineConfigProcessorItem {
	return ObservabilityPipelineConfigProcessorItem{ObservabilityPipelineThrottleProcessor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigProcessorItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineFilterProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineFilterProcessor)
	if err == nil {
		if obj.ObservabilityPipelineFilterProcessor != nil && obj.ObservabilityPipelineFilterProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineFilterProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineFilterProcessor)
			if string(jsonObservabilityPipelineFilterProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineFilterProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineFilterProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineFilterProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineParseJSONProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineParseJSONProcessor)
	if err == nil {
		if obj.ObservabilityPipelineParseJSONProcessor != nil && obj.ObservabilityPipelineParseJSONProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineParseJSONProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineParseJSONProcessor)
			if string(jsonObservabilityPipelineParseJSONProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineParseJSONProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineParseJSONProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineParseJSONProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineQuotaProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineQuotaProcessor)
	if err == nil {
		if obj.ObservabilityPipelineQuotaProcessor != nil && obj.ObservabilityPipelineQuotaProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineQuotaProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineQuotaProcessor)
			if string(jsonObservabilityPipelineQuotaProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineQuotaProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineQuotaProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineQuotaProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineAddFieldsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAddFieldsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineAddFieldsProcessor != nil && obj.ObservabilityPipelineAddFieldsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineAddFieldsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineAddFieldsProcessor)
			if string(jsonObservabilityPipelineAddFieldsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineAddFieldsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAddFieldsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineAddFieldsProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineRemoveFieldsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineRemoveFieldsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineRemoveFieldsProcessor != nil && obj.ObservabilityPipelineRemoveFieldsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineRemoveFieldsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineRemoveFieldsProcessor)
			if string(jsonObservabilityPipelineRemoveFieldsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineRemoveFieldsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineRemoveFieldsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineRemoveFieldsProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineRenameFieldsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineRenameFieldsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineRenameFieldsProcessor != nil && obj.ObservabilityPipelineRenameFieldsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineRenameFieldsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineRenameFieldsProcessor)
			if string(jsonObservabilityPipelineRenameFieldsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineRenameFieldsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineRenameFieldsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineRenameFieldsProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineOcsfMapperProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineOcsfMapperProcessor)
	if err == nil {
		if obj.ObservabilityPipelineOcsfMapperProcessor != nil && obj.ObservabilityPipelineOcsfMapperProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineOcsfMapperProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineOcsfMapperProcessor)
			if string(jsonObservabilityPipelineOcsfMapperProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineOcsfMapperProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineOcsfMapperProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineOcsfMapperProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineAddEnvVarsProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAddEnvVarsProcessor)
	if err == nil {
		if obj.ObservabilityPipelineAddEnvVarsProcessor != nil && obj.ObservabilityPipelineAddEnvVarsProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineAddEnvVarsProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineAddEnvVarsProcessor)
			if string(jsonObservabilityPipelineAddEnvVarsProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineAddEnvVarsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAddEnvVarsProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineAddEnvVarsProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineDedupeProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDedupeProcessor)
	if err == nil {
		if obj.ObservabilityPipelineDedupeProcessor != nil && obj.ObservabilityPipelineDedupeProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineDedupeProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineDedupeProcessor)
			if string(jsonObservabilityPipelineDedupeProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineDedupeProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDedupeProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineDedupeProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineEnrichmentTableProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineEnrichmentTableProcessor)
	if err == nil {
		if obj.ObservabilityPipelineEnrichmentTableProcessor != nil && obj.ObservabilityPipelineEnrichmentTableProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineEnrichmentTableProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineEnrichmentTableProcessor)
			if string(jsonObservabilityPipelineEnrichmentTableProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineEnrichmentTableProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineEnrichmentTableProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineEnrichmentTableProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineReduceProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineReduceProcessor)
	if err == nil {
		if obj.ObservabilityPipelineReduceProcessor != nil && obj.ObservabilityPipelineReduceProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineReduceProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineReduceProcessor)
			if string(jsonObservabilityPipelineReduceProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineReduceProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineReduceProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineReduceProcessor = nil
	}

	// try to unmarshal data into ObservabilityPipelineThrottleProcessor
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineThrottleProcessor)
	if err == nil {
		if obj.ObservabilityPipelineThrottleProcessor != nil && obj.ObservabilityPipelineThrottleProcessor.UnparsedObject == nil {
			jsonObservabilityPipelineThrottleProcessor, _ := datadog.Marshal(obj.ObservabilityPipelineThrottleProcessor)
			if string(jsonObservabilityPipelineThrottleProcessor) == "{}" { // empty struct
				obj.ObservabilityPipelineThrottleProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineThrottleProcessor = nil
		}
	} else {
		obj.ObservabilityPipelineThrottleProcessor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineFilterProcessor = nil
		obj.ObservabilityPipelineParseJSONProcessor = nil
		obj.ObservabilityPipelineQuotaProcessor = nil
		obj.ObservabilityPipelineAddFieldsProcessor = nil
		obj.ObservabilityPipelineRemoveFieldsProcessor = nil
		obj.ObservabilityPipelineRenameFieldsProcessor = nil
		obj.ObservabilityPipelineOcsfMapperProcessor = nil
		obj.ObservabilityPipelineAddEnvVarsProcessor = nil
		obj.ObservabilityPipelineDedupeProcessor = nil
		obj.ObservabilityPipelineEnrichmentTableProcessor = nil
		obj.ObservabilityPipelineReduceProcessor = nil
		obj.ObservabilityPipelineThrottleProcessor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigProcessorItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineFilterProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineFilterProcessor)
	}

	if obj.ObservabilityPipelineParseJSONProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineParseJSONProcessor)
	}

	if obj.ObservabilityPipelineQuotaProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineQuotaProcessor)
	}

	if obj.ObservabilityPipelineAddFieldsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAddFieldsProcessor)
	}

	if obj.ObservabilityPipelineRemoveFieldsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineRemoveFieldsProcessor)
	}

	if obj.ObservabilityPipelineRenameFieldsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineRenameFieldsProcessor)
	}

	if obj.ObservabilityPipelineOcsfMapperProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineOcsfMapperProcessor)
	}

	if obj.ObservabilityPipelineAddEnvVarsProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAddEnvVarsProcessor)
	}

	if obj.ObservabilityPipelineDedupeProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDedupeProcessor)
	}

	if obj.ObservabilityPipelineEnrichmentTableProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineEnrichmentTableProcessor)
	}

	if obj.ObservabilityPipelineReduceProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineReduceProcessor)
	}

	if obj.ObservabilityPipelineThrottleProcessor != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineThrottleProcessor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigProcessorItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineFilterProcessor != nil {
		return obj.ObservabilityPipelineFilterProcessor
	}

	if obj.ObservabilityPipelineParseJSONProcessor != nil {
		return obj.ObservabilityPipelineParseJSONProcessor
	}

	if obj.ObservabilityPipelineQuotaProcessor != nil {
		return obj.ObservabilityPipelineQuotaProcessor
	}

	if obj.ObservabilityPipelineAddFieldsProcessor != nil {
		return obj.ObservabilityPipelineAddFieldsProcessor
	}

	if obj.ObservabilityPipelineRemoveFieldsProcessor != nil {
		return obj.ObservabilityPipelineRemoveFieldsProcessor
	}

	if obj.ObservabilityPipelineRenameFieldsProcessor != nil {
		return obj.ObservabilityPipelineRenameFieldsProcessor
	}

	if obj.ObservabilityPipelineOcsfMapperProcessor != nil {
		return obj.ObservabilityPipelineOcsfMapperProcessor
	}

	if obj.ObservabilityPipelineAddEnvVarsProcessor != nil {
		return obj.ObservabilityPipelineAddEnvVarsProcessor
	}

	if obj.ObservabilityPipelineDedupeProcessor != nil {
		return obj.ObservabilityPipelineDedupeProcessor
	}

	if obj.ObservabilityPipelineEnrichmentTableProcessor != nil {
		return obj.ObservabilityPipelineEnrichmentTableProcessor
	}

	if obj.ObservabilityPipelineReduceProcessor != nil {
		return obj.ObservabilityPipelineReduceProcessor
	}

	if obj.ObservabilityPipelineThrottleProcessor != nil {
		return obj.ObservabilityPipelineThrottleProcessor
	}

	// all schemas are nil
	return nil
}
