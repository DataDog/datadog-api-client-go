// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueInteractionItem - A single interaction to add to an annotation queue.
type LLMObsAnnotationQueueInteractionItem struct {
	LLMObsTraceInteractionItem        *LLMObsTraceInteractionItem
	LLMObsDisplayBlockInteractionItem *LLMObsDisplayBlockInteractionItem
	LLMObsFrontendInteractionItem     *LLMObsFrontendInteractionItem

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LLMObsTraceInteractionItemAsLLMObsAnnotationQueueInteractionItem is a convenience function that returns LLMObsTraceInteractionItem wrapped in LLMObsAnnotationQueueInteractionItem.
func LLMObsTraceInteractionItemAsLLMObsAnnotationQueueInteractionItem(v *LLMObsTraceInteractionItem) LLMObsAnnotationQueueInteractionItem {
	return LLMObsAnnotationQueueInteractionItem{LLMObsTraceInteractionItem: v}
}

// LLMObsDisplayBlockInteractionItemAsLLMObsAnnotationQueueInteractionItem is a convenience function that returns LLMObsDisplayBlockInteractionItem wrapped in LLMObsAnnotationQueueInteractionItem.
func LLMObsDisplayBlockInteractionItemAsLLMObsAnnotationQueueInteractionItem(v *LLMObsDisplayBlockInteractionItem) LLMObsAnnotationQueueInteractionItem {
	return LLMObsAnnotationQueueInteractionItem{LLMObsDisplayBlockInteractionItem: v}
}

// LLMObsFrontendInteractionItemAsLLMObsAnnotationQueueInteractionItem is a convenience function that returns LLMObsFrontendInteractionItem wrapped in LLMObsAnnotationQueueInteractionItem.
func LLMObsFrontendInteractionItemAsLLMObsAnnotationQueueInteractionItem(v *LLMObsFrontendInteractionItem) LLMObsAnnotationQueueInteractionItem {
	return LLMObsAnnotationQueueInteractionItem{LLMObsFrontendInteractionItem: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LLMObsAnnotationQueueInteractionItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LLMObsTraceInteractionItem
	err = datadog.Unmarshal(data, &obj.LLMObsTraceInteractionItem)
	if err == nil {
		if obj.LLMObsTraceInteractionItem != nil && obj.LLMObsTraceInteractionItem.UnparsedObject == nil {
			jsonLLMObsTraceInteractionItem, _ := datadog.Marshal(obj.LLMObsTraceInteractionItem)
			if string(jsonLLMObsTraceInteractionItem) == "{}" { // empty struct
				obj.LLMObsTraceInteractionItem = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsTraceInteractionItem = nil
		}
	} else {
		obj.LLMObsTraceInteractionItem = nil
	}

	// try to unmarshal data into LLMObsDisplayBlockInteractionItem
	err = datadog.Unmarshal(data, &obj.LLMObsDisplayBlockInteractionItem)
	if err == nil {
		if obj.LLMObsDisplayBlockInteractionItem != nil && obj.LLMObsDisplayBlockInteractionItem.UnparsedObject == nil {
			jsonLLMObsDisplayBlockInteractionItem, _ := datadog.Marshal(obj.LLMObsDisplayBlockInteractionItem)
			if string(jsonLLMObsDisplayBlockInteractionItem) == "{}" { // empty struct
				obj.LLMObsDisplayBlockInteractionItem = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsDisplayBlockInteractionItem = nil
		}
	} else {
		obj.LLMObsDisplayBlockInteractionItem = nil
	}

	// try to unmarshal data into LLMObsFrontendInteractionItem
	err = datadog.Unmarshal(data, &obj.LLMObsFrontendInteractionItem)
	if err == nil {
		if obj.LLMObsFrontendInteractionItem != nil && obj.LLMObsFrontendInteractionItem.UnparsedObject == nil {
			jsonLLMObsFrontendInteractionItem, _ := datadog.Marshal(obj.LLMObsFrontendInteractionItem)
			if string(jsonLLMObsFrontendInteractionItem) == "{}" { // empty struct
				obj.LLMObsFrontendInteractionItem = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsFrontendInteractionItem = nil
		}
	} else {
		obj.LLMObsFrontendInteractionItem = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LLMObsTraceInteractionItem = nil
		obj.LLMObsDisplayBlockInteractionItem = nil
		obj.LLMObsFrontendInteractionItem = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LLMObsAnnotationQueueInteractionItem) MarshalJSON() ([]byte, error) {
	if obj.LLMObsTraceInteractionItem != nil {
		return datadog.Marshal(&obj.LLMObsTraceInteractionItem)
	}

	if obj.LLMObsDisplayBlockInteractionItem != nil {
		return datadog.Marshal(&obj.LLMObsDisplayBlockInteractionItem)
	}

	if obj.LLMObsFrontendInteractionItem != nil {
		return datadog.Marshal(&obj.LLMObsFrontendInteractionItem)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LLMObsAnnotationQueueInteractionItem) GetActualInstance() interface{} {
	if obj.LLMObsTraceInteractionItem != nil {
		return obj.LLMObsTraceInteractionItem
	}

	if obj.LLMObsDisplayBlockInteractionItem != nil {
		return obj.LLMObsDisplayBlockInteractionItem
	}

	if obj.LLMObsFrontendInteractionItem != nil {
		return obj.LLMObsFrontendInteractionItem
	}

	// all schemas are nil
	return nil
}
