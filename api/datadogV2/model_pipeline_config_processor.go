// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineConfigProcessor - A processor for the pipeline.
type PipelineConfigProcessor struct {
	PipelineFilterProcessor       *PipelineFilterProcessor
	ParseJSONProcessor            *ParseJSONProcessor
	PipelineQuotaProcessor        *PipelineQuotaProcessor
	PipelineAddFieldsProcessor    *PipelineAddFieldsProcessor
	PipelineRemoveFieldsProcessor *PipelineRemoveFieldsProcessor
	PipelineRenameFieldsProcessor *PipelineRenameFieldsProcessor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// PipelineFilterProcessorAsPipelineConfigProcessor is a convenience function that returns PipelineFilterProcessor wrapped in PipelineConfigProcessor.
func PipelineFilterProcessorAsPipelineConfigProcessor(v *PipelineFilterProcessor) PipelineConfigProcessor {
	return PipelineConfigProcessor{PipelineFilterProcessor: v}
}

// ParseJSONProcessorAsPipelineConfigProcessor is a convenience function that returns ParseJSONProcessor wrapped in PipelineConfigProcessor.
func ParseJSONProcessorAsPipelineConfigProcessor(v *ParseJSONProcessor) PipelineConfigProcessor {
	return PipelineConfigProcessor{ParseJSONProcessor: v}
}

// PipelineQuotaProcessorAsPipelineConfigProcessor is a convenience function that returns PipelineQuotaProcessor wrapped in PipelineConfigProcessor.
func PipelineQuotaProcessorAsPipelineConfigProcessor(v *PipelineQuotaProcessor) PipelineConfigProcessor {
	return PipelineConfigProcessor{PipelineQuotaProcessor: v}
}

// PipelineAddFieldsProcessorAsPipelineConfigProcessor is a convenience function that returns PipelineAddFieldsProcessor wrapped in PipelineConfigProcessor.
func PipelineAddFieldsProcessorAsPipelineConfigProcessor(v *PipelineAddFieldsProcessor) PipelineConfigProcessor {
	return PipelineConfigProcessor{PipelineAddFieldsProcessor: v}
}

// PipelineRemoveFieldsProcessorAsPipelineConfigProcessor is a convenience function that returns PipelineRemoveFieldsProcessor wrapped in PipelineConfigProcessor.
func PipelineRemoveFieldsProcessorAsPipelineConfigProcessor(v *PipelineRemoveFieldsProcessor) PipelineConfigProcessor {
	return PipelineConfigProcessor{PipelineRemoveFieldsProcessor: v}
}

// PipelineRenameFieldsProcessorAsPipelineConfigProcessor is a convenience function that returns PipelineRenameFieldsProcessor wrapped in PipelineConfigProcessor.
func PipelineRenameFieldsProcessorAsPipelineConfigProcessor(v *PipelineRenameFieldsProcessor) PipelineConfigProcessor {
	return PipelineConfigProcessor{PipelineRenameFieldsProcessor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *PipelineConfigProcessor) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PipelineFilterProcessor
	err = datadog.Unmarshal(data, &obj.PipelineFilterProcessor)
	if err == nil {
		if obj.PipelineFilterProcessor != nil && obj.PipelineFilterProcessor.UnparsedObject == nil {
			jsonPipelineFilterProcessor, _ := datadog.Marshal(obj.PipelineFilterProcessor)
			if string(jsonPipelineFilterProcessor) == "{}" { // empty struct
				obj.PipelineFilterProcessor = nil
			} else {
				match++
			}
		} else {
			obj.PipelineFilterProcessor = nil
		}
	} else {
		obj.PipelineFilterProcessor = nil
	}

	// try to unmarshal data into ParseJSONProcessor
	err = datadog.Unmarshal(data, &obj.ParseJSONProcessor)
	if err == nil {
		if obj.ParseJSONProcessor != nil && obj.ParseJSONProcessor.UnparsedObject == nil {
			jsonParseJSONProcessor, _ := datadog.Marshal(obj.ParseJSONProcessor)
			if string(jsonParseJSONProcessor) == "{}" { // empty struct
				obj.ParseJSONProcessor = nil
			} else {
				match++
			}
		} else {
			obj.ParseJSONProcessor = nil
		}
	} else {
		obj.ParseJSONProcessor = nil
	}

	// try to unmarshal data into PipelineQuotaProcessor
	err = datadog.Unmarshal(data, &obj.PipelineQuotaProcessor)
	if err == nil {
		if obj.PipelineQuotaProcessor != nil && obj.PipelineQuotaProcessor.UnparsedObject == nil {
			jsonPipelineQuotaProcessor, _ := datadog.Marshal(obj.PipelineQuotaProcessor)
			if string(jsonPipelineQuotaProcessor) == "{}" { // empty struct
				obj.PipelineQuotaProcessor = nil
			} else {
				match++
			}
		} else {
			obj.PipelineQuotaProcessor = nil
		}
	} else {
		obj.PipelineQuotaProcessor = nil
	}

	// try to unmarshal data into PipelineAddFieldsProcessor
	err = datadog.Unmarshal(data, &obj.PipelineAddFieldsProcessor)
	if err == nil {
		if obj.PipelineAddFieldsProcessor != nil && obj.PipelineAddFieldsProcessor.UnparsedObject == nil {
			jsonPipelineAddFieldsProcessor, _ := datadog.Marshal(obj.PipelineAddFieldsProcessor)
			if string(jsonPipelineAddFieldsProcessor) == "{}" { // empty struct
				obj.PipelineAddFieldsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.PipelineAddFieldsProcessor = nil
		}
	} else {
		obj.PipelineAddFieldsProcessor = nil
	}

	// try to unmarshal data into PipelineRemoveFieldsProcessor
	err = datadog.Unmarshal(data, &obj.PipelineRemoveFieldsProcessor)
	if err == nil {
		if obj.PipelineRemoveFieldsProcessor != nil && obj.PipelineRemoveFieldsProcessor.UnparsedObject == nil {
			jsonPipelineRemoveFieldsProcessor, _ := datadog.Marshal(obj.PipelineRemoveFieldsProcessor)
			if string(jsonPipelineRemoveFieldsProcessor) == "{}" { // empty struct
				obj.PipelineRemoveFieldsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.PipelineRemoveFieldsProcessor = nil
		}
	} else {
		obj.PipelineRemoveFieldsProcessor = nil
	}

	// try to unmarshal data into PipelineRenameFieldsProcessor
	err = datadog.Unmarshal(data, &obj.PipelineRenameFieldsProcessor)
	if err == nil {
		if obj.PipelineRenameFieldsProcessor != nil && obj.PipelineRenameFieldsProcessor.UnparsedObject == nil {
			jsonPipelineRenameFieldsProcessor, _ := datadog.Marshal(obj.PipelineRenameFieldsProcessor)
			if string(jsonPipelineRenameFieldsProcessor) == "{}" { // empty struct
				obj.PipelineRenameFieldsProcessor = nil
			} else {
				match++
			}
		} else {
			obj.PipelineRenameFieldsProcessor = nil
		}
	} else {
		obj.PipelineRenameFieldsProcessor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.PipelineFilterProcessor = nil
		obj.ParseJSONProcessor = nil
		obj.PipelineQuotaProcessor = nil
		obj.PipelineAddFieldsProcessor = nil
		obj.PipelineRemoveFieldsProcessor = nil
		obj.PipelineRenameFieldsProcessor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj PipelineConfigProcessor) MarshalJSON() ([]byte, error) {
	if obj.PipelineFilterProcessor != nil {
		return datadog.Marshal(&obj.PipelineFilterProcessor)
	}

	if obj.ParseJSONProcessor != nil {
		return datadog.Marshal(&obj.ParseJSONProcessor)
	}

	if obj.PipelineQuotaProcessor != nil {
		return datadog.Marshal(&obj.PipelineQuotaProcessor)
	}

	if obj.PipelineAddFieldsProcessor != nil {
		return datadog.Marshal(&obj.PipelineAddFieldsProcessor)
	}

	if obj.PipelineRemoveFieldsProcessor != nil {
		return datadog.Marshal(&obj.PipelineRemoveFieldsProcessor)
	}

	if obj.PipelineRenameFieldsProcessor != nil {
		return datadog.Marshal(&obj.PipelineRenameFieldsProcessor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *PipelineConfigProcessor) GetActualInstance() interface{} {
	if obj.PipelineFilterProcessor != nil {
		return obj.PipelineFilterProcessor
	}

	if obj.ParseJSONProcessor != nil {
		return obj.ParseJSONProcessor
	}

	if obj.PipelineQuotaProcessor != nil {
		return obj.PipelineQuotaProcessor
	}

	if obj.PipelineAddFieldsProcessor != nil {
		return obj.PipelineAddFieldsProcessor
	}

	if obj.PipelineRemoveFieldsProcessor != nil {
		return obj.PipelineRemoveFieldsProcessor
	}

	if obj.PipelineRenameFieldsProcessor != nil {
		return obj.PipelineRenameFieldsProcessor
	}

	// all schemas are nil
	return nil
}
