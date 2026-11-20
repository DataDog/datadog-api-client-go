// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamRequest - Test-only generated schema for the dashboard versioning proof of concept.
type ListStreamRequest struct {
	EventsListStreamRequest_20270101              *EventsListStreamRequest
	IssueStreamRequest_20270101                   *IssueStreamRequest
	CaseStreamRequest_20270101                    *CaseStreamRequest
	WatchdogAlertsRequest_20270101                *WatchdogAlertsRequest
	IncidentsStreamRequest_20270101               *IncidentsStreamRequest
	DatabaseRecommendationsStreamRequest_20270101 *DatabaseRecommendationsStreamRequest
	ApmRecommendationsStreamRequest_20270101      *ApmRecommendationsStreamRequest
	DetectionRuleStreamRequest_20270101           *DetectionRuleStreamRequest
	DataObservabilityAssetsRequest_20270101       *DataObservabilityAssetsRequest
	DoraListStreamRequest_20270101                *DoraListStreamRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// EventsListStreamRequest_20270101AsListStreamRequest is a convenience function that returns EventsListStreamRequest wrapped in ListStreamRequest.
func EventsListStreamRequest_20270101AsListStreamRequest(v *EventsListStreamRequest) ListStreamRequest {
	return ListStreamRequest{EventsListStreamRequest_20270101: v}
}

// IssueStreamRequest_20270101AsListStreamRequest is a convenience function that returns IssueStreamRequest wrapped in ListStreamRequest.
func IssueStreamRequest_20270101AsListStreamRequest(v *IssueStreamRequest) ListStreamRequest {
	return ListStreamRequest{IssueStreamRequest_20270101: v}
}

// CaseStreamRequest_20270101AsListStreamRequest is a convenience function that returns CaseStreamRequest wrapped in ListStreamRequest.
func CaseStreamRequest_20270101AsListStreamRequest(v *CaseStreamRequest) ListStreamRequest {
	return ListStreamRequest{CaseStreamRequest_20270101: v}
}

// WatchdogAlertsRequest_20270101AsListStreamRequest is a convenience function that returns WatchdogAlertsRequest wrapped in ListStreamRequest.
func WatchdogAlertsRequest_20270101AsListStreamRequest(v *WatchdogAlertsRequest) ListStreamRequest {
	return ListStreamRequest{WatchdogAlertsRequest_20270101: v}
}

// IncidentsStreamRequest_20270101AsListStreamRequest is a convenience function that returns IncidentsStreamRequest wrapped in ListStreamRequest.
func IncidentsStreamRequest_20270101AsListStreamRequest(v *IncidentsStreamRequest) ListStreamRequest {
	return ListStreamRequest{IncidentsStreamRequest_20270101: v}
}

// DatabaseRecommendationsStreamRequest_20270101AsListStreamRequest is a convenience function that returns DatabaseRecommendationsStreamRequest wrapped in ListStreamRequest.
func DatabaseRecommendationsStreamRequest_20270101AsListStreamRequest(v *DatabaseRecommendationsStreamRequest) ListStreamRequest {
	return ListStreamRequest{DatabaseRecommendationsStreamRequest_20270101: v}
}

// ApmRecommendationsStreamRequest_20270101AsListStreamRequest is a convenience function that returns ApmRecommendationsStreamRequest wrapped in ListStreamRequest.
func ApmRecommendationsStreamRequest_20270101AsListStreamRequest(v *ApmRecommendationsStreamRequest) ListStreamRequest {
	return ListStreamRequest{ApmRecommendationsStreamRequest_20270101: v}
}

// DetectionRuleStreamRequest_20270101AsListStreamRequest is a convenience function that returns DetectionRuleStreamRequest wrapped in ListStreamRequest.
func DetectionRuleStreamRequest_20270101AsListStreamRequest(v *DetectionRuleStreamRequest) ListStreamRequest {
	return ListStreamRequest{DetectionRuleStreamRequest_20270101: v}
}

// DataObservabilityAssetsRequest_20270101AsListStreamRequest is a convenience function that returns DataObservabilityAssetsRequest wrapped in ListStreamRequest.
func DataObservabilityAssetsRequest_20270101AsListStreamRequest(v *DataObservabilityAssetsRequest) ListStreamRequest {
	return ListStreamRequest{DataObservabilityAssetsRequest_20270101: v}
}

// DoraListStreamRequest_20270101AsListStreamRequest is a convenience function that returns DoraListStreamRequest wrapped in ListStreamRequest.
func DoraListStreamRequest_20270101AsListStreamRequest(v *DoraListStreamRequest) ListStreamRequest {
	return ListStreamRequest{DoraListStreamRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ListStreamRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EventsListStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.EventsListStreamRequest_20270101)
	if err == nil {
		if obj.EventsListStreamRequest_20270101 != nil && obj.EventsListStreamRequest_20270101.UnparsedObject == nil {
			jsonEventsListStreamRequest_20270101, _ := datadog.Marshal(obj.EventsListStreamRequest_20270101)
			if string(jsonEventsListStreamRequest_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.EventsListStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventsListStreamRequest_20270101 = nil
		}
	} else {
		obj.EventsListStreamRequest_20270101 = nil
	}

	// try to unmarshal data into IssueStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.IssueStreamRequest_20270101)
	if err == nil {
		if obj.IssueStreamRequest_20270101 != nil && obj.IssueStreamRequest_20270101.UnparsedObject == nil {
			jsonIssueStreamRequest_20270101, _ := datadog.Marshal(obj.IssueStreamRequest_20270101)
			if string(jsonIssueStreamRequest_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.IssueStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.IssueStreamRequest_20270101 = nil
		}
	} else {
		obj.IssueStreamRequest_20270101 = nil
	}

	// try to unmarshal data into CaseStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.CaseStreamRequest_20270101)
	if err == nil {
		if obj.CaseStreamRequest_20270101 != nil && obj.CaseStreamRequest_20270101.UnparsedObject == nil {
			jsonCaseStreamRequest_20270101, _ := datadog.Marshal(obj.CaseStreamRequest_20270101)
			if string(jsonCaseStreamRequest_20270101) == "{}" { // empty struct
				obj.CaseStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CaseStreamRequest_20270101 = nil
		}
	} else {
		obj.CaseStreamRequest_20270101 = nil
	}

	// try to unmarshal data into WatchdogAlertsRequest_20270101
	err = datadog.Unmarshal(data, &obj.WatchdogAlertsRequest_20270101)
	if err == nil {
		if obj.WatchdogAlertsRequest_20270101 != nil && obj.WatchdogAlertsRequest_20270101.UnparsedObject == nil {
			jsonWatchdogAlertsRequest_20270101, _ := datadog.Marshal(obj.WatchdogAlertsRequest_20270101)
			if string(jsonWatchdogAlertsRequest_20270101) == "{}" { // empty struct
				obj.WatchdogAlertsRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.WatchdogAlertsRequest_20270101 = nil
		}
	} else {
		obj.WatchdogAlertsRequest_20270101 = nil
	}

	// try to unmarshal data into IncidentsStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.IncidentsStreamRequest_20270101)
	if err == nil {
		if obj.IncidentsStreamRequest_20270101 != nil && obj.IncidentsStreamRequest_20270101.UnparsedObject == nil {
			jsonIncidentsStreamRequest_20270101, _ := datadog.Marshal(obj.IncidentsStreamRequest_20270101)
			if string(jsonIncidentsStreamRequest_20270101) == "{}" { // empty struct
				obj.IncidentsStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.IncidentsStreamRequest_20270101 = nil
		}
	} else {
		obj.IncidentsStreamRequest_20270101 = nil
	}

	// try to unmarshal data into DatabaseRecommendationsStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.DatabaseRecommendationsStreamRequest_20270101)
	if err == nil {
		if obj.DatabaseRecommendationsStreamRequest_20270101 != nil && obj.DatabaseRecommendationsStreamRequest_20270101.UnparsedObject == nil {
			jsonDatabaseRecommendationsStreamRequest_20270101, _ := datadog.Marshal(obj.DatabaseRecommendationsStreamRequest_20270101)
			if string(jsonDatabaseRecommendationsStreamRequest_20270101) == "{}" { // empty struct
				obj.DatabaseRecommendationsStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DatabaseRecommendationsStreamRequest_20270101 = nil
		}
	} else {
		obj.DatabaseRecommendationsStreamRequest_20270101 = nil
	}

	// try to unmarshal data into ApmRecommendationsStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.ApmRecommendationsStreamRequest_20270101)
	if err == nil {
		if obj.ApmRecommendationsStreamRequest_20270101 != nil && obj.ApmRecommendationsStreamRequest_20270101.UnparsedObject == nil {
			jsonApmRecommendationsStreamRequest_20270101, _ := datadog.Marshal(obj.ApmRecommendationsStreamRequest_20270101)
			if string(jsonApmRecommendationsStreamRequest_20270101) == "{}" { // empty struct
				obj.ApmRecommendationsStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ApmRecommendationsStreamRequest_20270101 = nil
		}
	} else {
		obj.ApmRecommendationsStreamRequest_20270101 = nil
	}

	// try to unmarshal data into DetectionRuleStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.DetectionRuleStreamRequest_20270101)
	if err == nil {
		if obj.DetectionRuleStreamRequest_20270101 != nil && obj.DetectionRuleStreamRequest_20270101.UnparsedObject == nil {
			jsonDetectionRuleStreamRequest_20270101, _ := datadog.Marshal(obj.DetectionRuleStreamRequest_20270101)
			if string(jsonDetectionRuleStreamRequest_20270101) == "{}" { // empty struct
				obj.DetectionRuleStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DetectionRuleStreamRequest_20270101 = nil
		}
	} else {
		obj.DetectionRuleStreamRequest_20270101 = nil
	}

	// try to unmarshal data into DataObservabilityAssetsRequest_20270101
	err = datadog.Unmarshal(data, &obj.DataObservabilityAssetsRequest_20270101)
	if err == nil {
		if obj.DataObservabilityAssetsRequest_20270101 != nil && obj.DataObservabilityAssetsRequest_20270101.UnparsedObject == nil {
			jsonDataObservabilityAssetsRequest_20270101, _ := datadog.Marshal(obj.DataObservabilityAssetsRequest_20270101)
			if string(jsonDataObservabilityAssetsRequest_20270101) == "{}" { // empty struct
				obj.DataObservabilityAssetsRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DataObservabilityAssetsRequest_20270101 = nil
		}
	} else {
		obj.DataObservabilityAssetsRequest_20270101 = nil
	}

	// try to unmarshal data into DoraListStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.DoraListStreamRequest_20270101)
	if err == nil {
		if obj.DoraListStreamRequest_20270101 != nil && obj.DoraListStreamRequest_20270101.UnparsedObject == nil {
			jsonDoraListStreamRequest_20270101, _ := datadog.Marshal(obj.DoraListStreamRequest_20270101)
			if string(jsonDoraListStreamRequest_20270101) == "{}" { // empty struct
				obj.DoraListStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DoraListStreamRequest_20270101 = nil
		}
	} else {
		obj.DoraListStreamRequest_20270101 = nil
	}

	if match > 1 {
		// more than one variant matched, so tell them apart by the items of the arrays they hold
		if obj.CaseStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.CaseStreamRequest_20270101.Columns) > 0
			for _, item := range obj.CaseStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.CaseStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.WatchdogAlertsRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.WatchdogAlertsRequest_20270101.Columns) > 0
			for _, item := range obj.WatchdogAlertsRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.WatchdogAlertsRequest_20270101 = nil
				match--
			}
		}
		if obj.IncidentsStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.IncidentsStreamRequest_20270101.Columns) > 0
			for _, item := range obj.IncidentsStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.IncidentsStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.DatabaseRecommendationsStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.DatabaseRecommendationsStreamRequest_20270101.Columns) > 0
			for _, item := range obj.DatabaseRecommendationsStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.DatabaseRecommendationsStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.ApmRecommendationsStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.ApmRecommendationsStreamRequest_20270101.Columns) > 0
			for _, item := range obj.ApmRecommendationsStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.ApmRecommendationsStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.DetectionRuleStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.DetectionRuleStreamRequest_20270101.Columns) > 0
			for _, item := range obj.DetectionRuleStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.DetectionRuleStreamRequest_20270101 = nil
				match--
			}
		}
		if obj.DataObservabilityAssetsRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.DataObservabilityAssetsRequest_20270101.Columns) > 0
			for _, item := range obj.DataObservabilityAssetsRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.DataObservabilityAssetsRequest_20270101 = nil
				match--
			}
		}
		if obj.DoraListStreamRequest_20270101 != nil {
			allColumnsUnparsed := len(obj.DoraListStreamRequest_20270101.Columns) > 0
			for _, item := range obj.DoraListStreamRequest_20270101.Columns {
				if item.UnparsedObject == nil {
					allColumnsUnparsed = false
					break
				}
			}
			if allColumnsUnparsed {
				obj.DoraListStreamRequest_20270101 = nil
				match--
			}
		}
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.EventsListStreamRequest_20270101 = nil
		obj.IssueStreamRequest_20270101 = nil
		obj.CaseStreamRequest_20270101 = nil
		obj.WatchdogAlertsRequest_20270101 = nil
		obj.IncidentsStreamRequest_20270101 = nil
		obj.DatabaseRecommendationsStreamRequest_20270101 = nil
		obj.ApmRecommendationsStreamRequest_20270101 = nil
		obj.DetectionRuleStreamRequest_20270101 = nil
		obj.DataObservabilityAssetsRequest_20270101 = nil
		obj.DoraListStreamRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ListStreamRequest) MarshalJSON() ([]byte, error) {
	if obj.EventsListStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.EventsListStreamRequest_20270101)
	}

	if obj.IssueStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.IssueStreamRequest_20270101)
	}

	if obj.CaseStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.CaseStreamRequest_20270101)
	}

	if obj.WatchdogAlertsRequest_20270101 != nil {
		return datadog.Marshal(&obj.WatchdogAlertsRequest_20270101)
	}

	if obj.IncidentsStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.IncidentsStreamRequest_20270101)
	}

	if obj.DatabaseRecommendationsStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.DatabaseRecommendationsStreamRequest_20270101)
	}

	if obj.ApmRecommendationsStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.ApmRecommendationsStreamRequest_20270101)
	}

	if obj.DetectionRuleStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.DetectionRuleStreamRequest_20270101)
	}

	if obj.DataObservabilityAssetsRequest_20270101 != nil {
		return datadog.Marshal(&obj.DataObservabilityAssetsRequest_20270101)
	}

	if obj.DoraListStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.DoraListStreamRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ListStreamRequest) GetActualInstance() interface{} {
	if obj.EventsListStreamRequest_20270101 != nil {
		return obj.EventsListStreamRequest_20270101
	}

	if obj.IssueStreamRequest_20270101 != nil {
		return obj.IssueStreamRequest_20270101
	}

	if obj.CaseStreamRequest_20270101 != nil {
		return obj.CaseStreamRequest_20270101
	}

	if obj.WatchdogAlertsRequest_20270101 != nil {
		return obj.WatchdogAlertsRequest_20270101
	}

	if obj.IncidentsStreamRequest_20270101 != nil {
		return obj.IncidentsStreamRequest_20270101
	}

	if obj.DatabaseRecommendationsStreamRequest_20270101 != nil {
		return obj.DatabaseRecommendationsStreamRequest_20270101
	}

	if obj.ApmRecommendationsStreamRequest_20270101 != nil {
		return obj.ApmRecommendationsStreamRequest_20270101
	}

	if obj.DetectionRuleStreamRequest_20270101 != nil {
		return obj.DetectionRuleStreamRequest_20270101
	}

	if obj.DataObservabilityAssetsRequest_20270101 != nil {
		return obj.DataObservabilityAssetsRequest_20270101
	}

	if obj.DoraListStreamRequest_20270101 != nil {
		return obj.DoraListStreamRequest_20270101
	}

	// all schemas are nil
	return nil
}
