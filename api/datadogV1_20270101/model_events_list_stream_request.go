// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsListStreamRequest - Test-only generated schema for the dashboard versioning proof of concept.
type EventsListStreamRequest struct {
	EventPlatformStreamRequest_20270101          *EventPlatformStreamRequest
	LogsStreamRequest_20270101                   *LogsStreamRequest
	LogsPatternStreamRequest_20270101            *LogsPatternStreamRequest
	LogsTransactionStreamRequest_20270101        *LogsTransactionStreamRequest
	AuditStreamRequest_20270101                  *AuditStreamRequest
	TraceStreamRequest_20270101                  *TraceStreamRequest
	ErrorsStreamRequest_20270101                 *ErrorsStreamRequest
	EventStreamRequest_20270101                  *EventStreamRequest
	EventPatternStreamRequest_20270101           *EventPatternStreamRequest
	SecurityRuntimePatternStreamRequest_20270101 *SecurityRuntimePatternStreamRequest
	RumStreamRequest_20270101                    *RumStreamRequest
	RecommendationsStreamRequest_20270101        *RecommendationsStreamRequest
	SecurityRuntimeStreamRequest_20270101        *SecurityRuntimeStreamRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// EventPlatformStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns EventPlatformStreamRequest wrapped in EventsListStreamRequest.
func EventPlatformStreamRequest_20270101AsEventsListStreamRequest(v *EventPlatformStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{EventPlatformStreamRequest_20270101: v}
}

// LogsStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns LogsStreamRequest wrapped in EventsListStreamRequest.
func LogsStreamRequest_20270101AsEventsListStreamRequest(v *LogsStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{LogsStreamRequest_20270101: v}
}

// LogsPatternStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns LogsPatternStreamRequest wrapped in EventsListStreamRequest.
func LogsPatternStreamRequest_20270101AsEventsListStreamRequest(v *LogsPatternStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{LogsPatternStreamRequest_20270101: v}
}

// LogsTransactionStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns LogsTransactionStreamRequest wrapped in EventsListStreamRequest.
func LogsTransactionStreamRequest_20270101AsEventsListStreamRequest(v *LogsTransactionStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{LogsTransactionStreamRequest_20270101: v}
}

// AuditStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns AuditStreamRequest wrapped in EventsListStreamRequest.
func AuditStreamRequest_20270101AsEventsListStreamRequest(v *AuditStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{AuditStreamRequest_20270101: v}
}

// TraceStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns TraceStreamRequest wrapped in EventsListStreamRequest.
func TraceStreamRequest_20270101AsEventsListStreamRequest(v *TraceStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{TraceStreamRequest_20270101: v}
}

// ErrorsStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns ErrorsStreamRequest wrapped in EventsListStreamRequest.
func ErrorsStreamRequest_20270101AsEventsListStreamRequest(v *ErrorsStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{ErrorsStreamRequest_20270101: v}
}

// EventStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns EventStreamRequest wrapped in EventsListStreamRequest.
func EventStreamRequest_20270101AsEventsListStreamRequest(v *EventStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{EventStreamRequest_20270101: v}
}

// EventPatternStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns EventPatternStreamRequest wrapped in EventsListStreamRequest.
func EventPatternStreamRequest_20270101AsEventsListStreamRequest(v *EventPatternStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{EventPatternStreamRequest_20270101: v}
}

// SecurityRuntimePatternStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns SecurityRuntimePatternStreamRequest wrapped in EventsListStreamRequest.
func SecurityRuntimePatternStreamRequest_20270101AsEventsListStreamRequest(v *SecurityRuntimePatternStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{SecurityRuntimePatternStreamRequest_20270101: v}
}

// RumStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns RumStreamRequest wrapped in EventsListStreamRequest.
func RumStreamRequest_20270101AsEventsListStreamRequest(v *RumStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{RumStreamRequest_20270101: v}
}

// RecommendationsStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns RecommendationsStreamRequest wrapped in EventsListStreamRequest.
func RecommendationsStreamRequest_20270101AsEventsListStreamRequest(v *RecommendationsStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{RecommendationsStreamRequest_20270101: v}
}

// SecurityRuntimeStreamRequest_20270101AsEventsListStreamRequest is a convenience function that returns SecurityRuntimeStreamRequest wrapped in EventsListStreamRequest.
func SecurityRuntimeStreamRequest_20270101AsEventsListStreamRequest(v *SecurityRuntimeStreamRequest) EventsListStreamRequest {
	return EventsListStreamRequest{SecurityRuntimeStreamRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EventsListStreamRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EventPlatformStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.EventPlatformStreamRequest_20270101)
	if err == nil {
		if obj.EventPlatformStreamRequest_20270101 != nil && obj.EventPlatformStreamRequest_20270101.UnparsedObject == nil {
			jsonEventPlatformStreamRequest_20270101, _ := datadog.Marshal(obj.EventPlatformStreamRequest_20270101)
			if string(jsonEventPlatformStreamRequest_20270101) == "{}" { // empty struct
				obj.EventPlatformStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventPlatformStreamRequest_20270101 = nil
		}
	} else {
		obj.EventPlatformStreamRequest_20270101 = nil
	}

	// try to unmarshal data into LogsStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.LogsStreamRequest_20270101)
	if err == nil {
		if obj.LogsStreamRequest_20270101 != nil && obj.LogsStreamRequest_20270101.UnparsedObject == nil {
			jsonLogsStreamRequest_20270101, _ := datadog.Marshal(obj.LogsStreamRequest_20270101)
			if string(jsonLogsStreamRequest_20270101) == "{}" { // empty struct
				obj.LogsStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.LogsStreamRequest_20270101 = nil
		}
	} else {
		obj.LogsStreamRequest_20270101 = nil
	}

	// try to unmarshal data into LogsPatternStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.LogsPatternStreamRequest_20270101)
	if err == nil {
		if obj.LogsPatternStreamRequest_20270101 != nil && obj.LogsPatternStreamRequest_20270101.UnparsedObject == nil {
			jsonLogsPatternStreamRequest_20270101, _ := datadog.Marshal(obj.LogsPatternStreamRequest_20270101)
			if string(jsonLogsPatternStreamRequest_20270101) == "{}" { // empty struct
				obj.LogsPatternStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.LogsPatternStreamRequest_20270101 = nil
		}
	} else {
		obj.LogsPatternStreamRequest_20270101 = nil
	}

	// try to unmarshal data into LogsTransactionStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.LogsTransactionStreamRequest_20270101)
	if err == nil {
		if obj.LogsTransactionStreamRequest_20270101 != nil && obj.LogsTransactionStreamRequest_20270101.UnparsedObject == nil {
			jsonLogsTransactionStreamRequest_20270101, _ := datadog.Marshal(obj.LogsTransactionStreamRequest_20270101)
			if string(jsonLogsTransactionStreamRequest_20270101) == "{}" { // empty struct
				obj.LogsTransactionStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.LogsTransactionStreamRequest_20270101 = nil
		}
	} else {
		obj.LogsTransactionStreamRequest_20270101 = nil
	}

	// try to unmarshal data into AuditStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.AuditStreamRequest_20270101)
	if err == nil {
		if obj.AuditStreamRequest_20270101 != nil && obj.AuditStreamRequest_20270101.UnparsedObject == nil {
			jsonAuditStreamRequest_20270101, _ := datadog.Marshal(obj.AuditStreamRequest_20270101)
			if string(jsonAuditStreamRequest_20270101) == "{}" { // empty struct
				obj.AuditStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.AuditStreamRequest_20270101 = nil
		}
	} else {
		obj.AuditStreamRequest_20270101 = nil
	}

	// try to unmarshal data into TraceStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.TraceStreamRequest_20270101)
	if err == nil {
		if obj.TraceStreamRequest_20270101 != nil && obj.TraceStreamRequest_20270101.UnparsedObject == nil {
			jsonTraceStreamRequest_20270101, _ := datadog.Marshal(obj.TraceStreamRequest_20270101)
			if string(jsonTraceStreamRequest_20270101) == "{}" { // empty struct
				obj.TraceStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TraceStreamRequest_20270101 = nil
		}
	} else {
		obj.TraceStreamRequest_20270101 = nil
	}

	// try to unmarshal data into ErrorsStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.ErrorsStreamRequest_20270101)
	if err == nil {
		if obj.ErrorsStreamRequest_20270101 != nil && obj.ErrorsStreamRequest_20270101.UnparsedObject == nil {
			jsonErrorsStreamRequest_20270101, _ := datadog.Marshal(obj.ErrorsStreamRequest_20270101)
			if string(jsonErrorsStreamRequest_20270101) == "{}" { // empty struct
				obj.ErrorsStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ErrorsStreamRequest_20270101 = nil
		}
	} else {
		obj.ErrorsStreamRequest_20270101 = nil
	}

	// try to unmarshal data into EventStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.EventStreamRequest_20270101)
	if err == nil {
		if obj.EventStreamRequest_20270101 != nil && obj.EventStreamRequest_20270101.UnparsedObject == nil {
			jsonEventStreamRequest_20270101, _ := datadog.Marshal(obj.EventStreamRequest_20270101)
			if string(jsonEventStreamRequest_20270101) == "{}" { // empty struct
				obj.EventStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventStreamRequest_20270101 = nil
		}
	} else {
		obj.EventStreamRequest_20270101 = nil
	}

	// try to unmarshal data into EventPatternStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.EventPatternStreamRequest_20270101)
	if err == nil {
		if obj.EventPatternStreamRequest_20270101 != nil && obj.EventPatternStreamRequest_20270101.UnparsedObject == nil {
			jsonEventPatternStreamRequest_20270101, _ := datadog.Marshal(obj.EventPatternStreamRequest_20270101)
			if string(jsonEventPatternStreamRequest_20270101) == "{}" { // empty struct
				obj.EventPatternStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventPatternStreamRequest_20270101 = nil
		}
	} else {
		obj.EventPatternStreamRequest_20270101 = nil
	}

	// try to unmarshal data into SecurityRuntimePatternStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.SecurityRuntimePatternStreamRequest_20270101)
	if err == nil {
		if obj.SecurityRuntimePatternStreamRequest_20270101 != nil && obj.SecurityRuntimePatternStreamRequest_20270101.UnparsedObject == nil {
			jsonSecurityRuntimePatternStreamRequest_20270101, _ := datadog.Marshal(obj.SecurityRuntimePatternStreamRequest_20270101)
			if string(jsonSecurityRuntimePatternStreamRequest_20270101) == "{}" { // empty struct
				obj.SecurityRuntimePatternStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SecurityRuntimePatternStreamRequest_20270101 = nil
		}
	} else {
		obj.SecurityRuntimePatternStreamRequest_20270101 = nil
	}

	// try to unmarshal data into RumStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.RumStreamRequest_20270101)
	if err == nil {
		if obj.RumStreamRequest_20270101 != nil && obj.RumStreamRequest_20270101.UnparsedObject == nil {
			jsonRumStreamRequest_20270101, _ := datadog.Marshal(obj.RumStreamRequest_20270101)
			if string(jsonRumStreamRequest_20270101) == "{}" { // empty struct
				obj.RumStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RumStreamRequest_20270101 = nil
		}
	} else {
		obj.RumStreamRequest_20270101 = nil
	}

	// try to unmarshal data into RecommendationsStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.RecommendationsStreamRequest_20270101)
	if err == nil {
		if obj.RecommendationsStreamRequest_20270101 != nil && obj.RecommendationsStreamRequest_20270101.UnparsedObject == nil {
			jsonRecommendationsStreamRequest_20270101, _ := datadog.Marshal(obj.RecommendationsStreamRequest_20270101)
			if string(jsonRecommendationsStreamRequest_20270101) == "{}" { // empty struct
				obj.RecommendationsStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RecommendationsStreamRequest_20270101 = nil
		}
	} else {
		obj.RecommendationsStreamRequest_20270101 = nil
	}

	// try to unmarshal data into SecurityRuntimeStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.SecurityRuntimeStreamRequest_20270101)
	if err == nil {
		if obj.SecurityRuntimeStreamRequest_20270101 != nil && obj.SecurityRuntimeStreamRequest_20270101.UnparsedObject == nil {
			jsonSecurityRuntimeStreamRequest_20270101, _ := datadog.Marshal(obj.SecurityRuntimeStreamRequest_20270101)
			if string(jsonSecurityRuntimeStreamRequest_20270101) == "{}" { // empty struct
				obj.SecurityRuntimeStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SecurityRuntimeStreamRequest_20270101 = nil
		}
	} else {
		obj.SecurityRuntimeStreamRequest_20270101 = nil
	}

	if match > 1 {
		// more than one variant matched, so tell them apart by the items of the arrays they hold
		if obj.EventPlatformStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.EventPlatformStreamRequest_20270101.Columns) > 0
			for _, item := range obj.EventPlatformStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.EventPlatformStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.LogsStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.LogsStreamRequest_20270101.Columns) > 0
			for _, item := range obj.LogsStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.LogsStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.LogsPatternStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.LogsPatternStreamRequest_20270101.Columns) > 0
			for _, item := range obj.LogsPatternStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.LogsPatternStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.LogsTransactionStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.LogsTransactionStreamRequest_20270101.Columns) > 0
			for _, item := range obj.LogsTransactionStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.LogsTransactionStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.AuditStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.AuditStreamRequest_20270101.Columns) > 0
			for _, item := range obj.AuditStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.AuditStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.TraceStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.TraceStreamRequest_20270101.Columns) > 0
			for _, item := range obj.TraceStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.TraceStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.ErrorsStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.ErrorsStreamRequest_20270101.Columns) > 0
			for _, item := range obj.ErrorsStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.ErrorsStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.EventStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.EventStreamRequest_20270101.Columns) > 0
			for _, item := range obj.EventStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.EventStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.EventPatternStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.EventPatternStreamRequest_20270101.Columns) > 0
			for _, item := range obj.EventPatternStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.EventPatternStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.SecurityRuntimePatternStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.SecurityRuntimePatternStreamRequest_20270101.Columns) > 0
			for _, item := range obj.SecurityRuntimePatternStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.SecurityRuntimePatternStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.RumStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.RumStreamRequest_20270101.Columns) > 0
			for _, item := range obj.RumStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.RumStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.RecommendationsStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.RecommendationsStreamRequest_20270101.Columns) > 0
			for _, item := range obj.RecommendationsStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.RecommendationsStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.SecurityRuntimeStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.SecurityRuntimeStreamRequest_20270101.Columns) > 0
			for _, item := range obj.SecurityRuntimeStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.SecurityRuntimeStreamRequest_20270101 = nil
				match--
			}
		}
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.EventPlatformStreamRequest_20270101 = nil
		obj.LogsStreamRequest_20270101 = nil
		obj.LogsPatternStreamRequest_20270101 = nil
		obj.LogsTransactionStreamRequest_20270101 = nil
		obj.AuditStreamRequest_20270101 = nil
		obj.TraceStreamRequest_20270101 = nil
		obj.ErrorsStreamRequest_20270101 = nil
		obj.EventStreamRequest_20270101 = nil
		obj.EventPatternStreamRequest_20270101 = nil
		obj.SecurityRuntimePatternStreamRequest_20270101 = nil
		obj.RumStreamRequest_20270101 = nil
		obj.RecommendationsStreamRequest_20270101 = nil
		obj.SecurityRuntimeStreamRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EventsListStreamRequest) MarshalJSON() ([]byte, error) {
	if obj.EventPlatformStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.EventPlatformStreamRequest_20270101)
	}

	if obj.LogsStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.LogsStreamRequest_20270101)
	}

	if obj.LogsPatternStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.LogsPatternStreamRequest_20270101)
	}

	if obj.LogsTransactionStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.LogsTransactionStreamRequest_20270101)
	}

	if obj.AuditStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.AuditStreamRequest_20270101)
	}

	if obj.TraceStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.TraceStreamRequest_20270101)
	}

	if obj.ErrorsStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.ErrorsStreamRequest_20270101)
	}

	if obj.EventStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.EventStreamRequest_20270101)
	}

	if obj.EventPatternStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.EventPatternStreamRequest_20270101)
	}

	if obj.SecurityRuntimePatternStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.SecurityRuntimePatternStreamRequest_20270101)
	}

	if obj.RumStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.RumStreamRequest_20270101)
	}

	if obj.RecommendationsStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.RecommendationsStreamRequest_20270101)
	}

	if obj.SecurityRuntimeStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.SecurityRuntimeStreamRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EventsListStreamRequest) GetActualInstance() interface{} {
	if obj.EventPlatformStreamRequest_20270101 != nil {
		return obj.EventPlatformStreamRequest_20270101
	}

	if obj.LogsStreamRequest_20270101 != nil {
		return obj.LogsStreamRequest_20270101
	}

	if obj.LogsPatternStreamRequest_20270101 != nil {
		return obj.LogsPatternStreamRequest_20270101
	}

	if obj.LogsTransactionStreamRequest_20270101 != nil {
		return obj.LogsTransactionStreamRequest_20270101
	}

	if obj.AuditStreamRequest_20270101 != nil {
		return obj.AuditStreamRequest_20270101
	}

	if obj.TraceStreamRequest_20270101 != nil {
		return obj.TraceStreamRequest_20270101
	}

	if obj.ErrorsStreamRequest_20270101 != nil {
		return obj.ErrorsStreamRequest_20270101
	}

	if obj.EventStreamRequest_20270101 != nil {
		return obj.EventStreamRequest_20270101
	}

	if obj.EventPatternStreamRequest_20270101 != nil {
		return obj.EventPatternStreamRequest_20270101
	}

	if obj.SecurityRuntimePatternStreamRequest_20270101 != nil {
		return obj.SecurityRuntimePatternStreamRequest_20270101
	}

	if obj.RumStreamRequest_20270101 != nil {
		return obj.RumStreamRequest_20270101
	}

	if obj.RecommendationsStreamRequest_20270101 != nil {
		return obj.RecommendationsStreamRequest_20270101
	}

	if obj.SecurityRuntimeStreamRequest_20270101 != nil {
		return obj.SecurityRuntimeStreamRequest_20270101
	}

	// all schemas are nil
	return nil
}
