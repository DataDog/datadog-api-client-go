// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableQuery - Test-only generated schema for the dashboard versioning proof of concept.
type GuidedTableQuery struct {
	MetricsLikeQueryFragmentMetrics_20270101         *MetricsLikeQueryFragmentMetrics
	MetricsLikeQueryFragmentCloudCost_20270101       *MetricsLikeQueryFragmentCloudCost
	EventsQueryFragmentDefinition_20270101           *EventsQueryFragmentDefinition
	ReferentialQueryFragmentReferenceTable_20270101  *ReferentialQueryFragmentReferenceTable
	ReferentialQueryFragmentManagedResource_20270101 *ReferentialQueryFragmentManagedResource
	PublishedAnalysisQueryFragment_20270101          *PublishedAnalysisQueryFragment
	ApmMetricsQueryFragment_20270101                 *ApmMetricsQueryFragment

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricsLikeQueryFragmentMetrics_20270101AsGuidedTableQuery is a convenience function that returns MetricsLikeQueryFragmentMetrics wrapped in GuidedTableQuery.
func MetricsLikeQueryFragmentMetrics_20270101AsGuidedTableQuery(v *MetricsLikeQueryFragmentMetrics) GuidedTableQuery {
	return GuidedTableQuery{MetricsLikeQueryFragmentMetrics_20270101: v}
}

// MetricsLikeQueryFragmentCloudCost_20270101AsGuidedTableQuery is a convenience function that returns MetricsLikeQueryFragmentCloudCost wrapped in GuidedTableQuery.
func MetricsLikeQueryFragmentCloudCost_20270101AsGuidedTableQuery(v *MetricsLikeQueryFragmentCloudCost) GuidedTableQuery {
	return GuidedTableQuery{MetricsLikeQueryFragmentCloudCost_20270101: v}
}

// EventsQueryFragmentDefinition_20270101AsGuidedTableQuery is a convenience function that returns EventsQueryFragmentDefinition wrapped in GuidedTableQuery.
func EventsQueryFragmentDefinition_20270101AsGuidedTableQuery(v *EventsQueryFragmentDefinition) GuidedTableQuery {
	return GuidedTableQuery{EventsQueryFragmentDefinition_20270101: v}
}

// ReferentialQueryFragmentReferenceTable_20270101AsGuidedTableQuery is a convenience function that returns ReferentialQueryFragmentReferenceTable wrapped in GuidedTableQuery.
func ReferentialQueryFragmentReferenceTable_20270101AsGuidedTableQuery(v *ReferentialQueryFragmentReferenceTable) GuidedTableQuery {
	return GuidedTableQuery{ReferentialQueryFragmentReferenceTable_20270101: v}
}

// ReferentialQueryFragmentManagedResource_20270101AsGuidedTableQuery is a convenience function that returns ReferentialQueryFragmentManagedResource wrapped in GuidedTableQuery.
func ReferentialQueryFragmentManagedResource_20270101AsGuidedTableQuery(v *ReferentialQueryFragmentManagedResource) GuidedTableQuery {
	return GuidedTableQuery{ReferentialQueryFragmentManagedResource_20270101: v}
}

// PublishedAnalysisQueryFragment_20270101AsGuidedTableQuery is a convenience function that returns PublishedAnalysisQueryFragment wrapped in GuidedTableQuery.
func PublishedAnalysisQueryFragment_20270101AsGuidedTableQuery(v *PublishedAnalysisQueryFragment) GuidedTableQuery {
	return GuidedTableQuery{PublishedAnalysisQueryFragment_20270101: v}
}

// ApmMetricsQueryFragment_20270101AsGuidedTableQuery is a convenience function that returns ApmMetricsQueryFragment wrapped in GuidedTableQuery.
func ApmMetricsQueryFragment_20270101AsGuidedTableQuery(v *ApmMetricsQueryFragment) GuidedTableQuery {
	return GuidedTableQuery{ApmMetricsQueryFragment_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GuidedTableQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetricsLikeQueryFragmentMetrics_20270101
	err = datadog.Unmarshal(data, &obj.MetricsLikeQueryFragmentMetrics_20270101)
	if err == nil {
		if obj.MetricsLikeQueryFragmentMetrics_20270101 != nil && obj.MetricsLikeQueryFragmentMetrics_20270101.UnparsedObject == nil {
			jsonMetricsLikeQueryFragmentMetrics_20270101, _ := datadog.Marshal(obj.MetricsLikeQueryFragmentMetrics_20270101)
			if string(jsonMetricsLikeQueryFragmentMetrics_20270101) == "{}" { // empty struct
				obj.MetricsLikeQueryFragmentMetrics_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.MetricsLikeQueryFragmentMetrics_20270101 = nil
		}
	} else {
		obj.MetricsLikeQueryFragmentMetrics_20270101 = nil
	}

	// try to unmarshal data into MetricsLikeQueryFragmentCloudCost_20270101
	err = datadog.Unmarshal(data, &obj.MetricsLikeQueryFragmentCloudCost_20270101)
	if err == nil {
		if obj.MetricsLikeQueryFragmentCloudCost_20270101 != nil && obj.MetricsLikeQueryFragmentCloudCost_20270101.UnparsedObject == nil {
			jsonMetricsLikeQueryFragmentCloudCost_20270101, _ := datadog.Marshal(obj.MetricsLikeQueryFragmentCloudCost_20270101)
			if string(jsonMetricsLikeQueryFragmentCloudCost_20270101) == "{}" { // empty struct
				obj.MetricsLikeQueryFragmentCloudCost_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.MetricsLikeQueryFragmentCloudCost_20270101 = nil
		}
	} else {
		obj.MetricsLikeQueryFragmentCloudCost_20270101 = nil
	}

	// try to unmarshal data into EventsQueryFragmentDefinition_20270101
	err = datadog.Unmarshal(data, &obj.EventsQueryFragmentDefinition_20270101)
	if err == nil {
		if obj.EventsQueryFragmentDefinition_20270101 != nil && obj.EventsQueryFragmentDefinition_20270101.UnparsedObject == nil {
			jsonEventsQueryFragmentDefinition_20270101, _ := datadog.Marshal(obj.EventsQueryFragmentDefinition_20270101)
			if string(jsonEventsQueryFragmentDefinition_20270101) == "{}" { // empty struct
				obj.EventsQueryFragmentDefinition_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventsQueryFragmentDefinition_20270101 = nil
		}
	} else {
		obj.EventsQueryFragmentDefinition_20270101 = nil
	}

	// try to unmarshal data into ReferentialQueryFragmentReferenceTable_20270101
	err = datadog.Unmarshal(data, &obj.ReferentialQueryFragmentReferenceTable_20270101)
	if err == nil {
		if obj.ReferentialQueryFragmentReferenceTable_20270101 != nil && obj.ReferentialQueryFragmentReferenceTable_20270101.UnparsedObject == nil {
			jsonReferentialQueryFragmentReferenceTable_20270101, _ := datadog.Marshal(obj.ReferentialQueryFragmentReferenceTable_20270101)
			if string(jsonReferentialQueryFragmentReferenceTable_20270101) == "{}" { // empty struct
				obj.ReferentialQueryFragmentReferenceTable_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ReferentialQueryFragmentReferenceTable_20270101 = nil
		}
	} else {
		obj.ReferentialQueryFragmentReferenceTable_20270101 = nil
	}

	// try to unmarshal data into ReferentialQueryFragmentManagedResource_20270101
	err = datadog.Unmarshal(data, &obj.ReferentialQueryFragmentManagedResource_20270101)
	if err == nil {
		if obj.ReferentialQueryFragmentManagedResource_20270101 != nil && obj.ReferentialQueryFragmentManagedResource_20270101.UnparsedObject == nil {
			jsonReferentialQueryFragmentManagedResource_20270101, _ := datadog.Marshal(obj.ReferentialQueryFragmentManagedResource_20270101)
			if string(jsonReferentialQueryFragmentManagedResource_20270101) == "{}" { // empty struct
				obj.ReferentialQueryFragmentManagedResource_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ReferentialQueryFragmentManagedResource_20270101 = nil
		}
	} else {
		obj.ReferentialQueryFragmentManagedResource_20270101 = nil
	}

	// try to unmarshal data into PublishedAnalysisQueryFragment_20270101
	err = datadog.Unmarshal(data, &obj.PublishedAnalysisQueryFragment_20270101)
	if err == nil {
		if obj.PublishedAnalysisQueryFragment_20270101 != nil && obj.PublishedAnalysisQueryFragment_20270101.UnparsedObject == nil {
			jsonPublishedAnalysisQueryFragment_20270101, _ := datadog.Marshal(obj.PublishedAnalysisQueryFragment_20270101)
			if string(jsonPublishedAnalysisQueryFragment_20270101) == "{}" { // empty struct
				obj.PublishedAnalysisQueryFragment_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.PublishedAnalysisQueryFragment_20270101 = nil
		}
	} else {
		obj.PublishedAnalysisQueryFragment_20270101 = nil
	}

	// try to unmarshal data into ApmMetricsQueryFragment_20270101
	err = datadog.Unmarshal(data, &obj.ApmMetricsQueryFragment_20270101)
	if err == nil {
		if obj.ApmMetricsQueryFragment_20270101 != nil && obj.ApmMetricsQueryFragment_20270101.UnparsedObject == nil {
			jsonApmMetricsQueryFragment_20270101, _ := datadog.Marshal(obj.ApmMetricsQueryFragment_20270101)
			if string(jsonApmMetricsQueryFragment_20270101) == "{}" { // empty struct
				obj.ApmMetricsQueryFragment_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ApmMetricsQueryFragment_20270101 = nil
		}
	} else {
		obj.ApmMetricsQueryFragment_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MetricsLikeQueryFragmentMetrics_20270101 = nil
		obj.MetricsLikeQueryFragmentCloudCost_20270101 = nil
		obj.EventsQueryFragmentDefinition_20270101 = nil
		obj.ReferentialQueryFragmentReferenceTable_20270101 = nil
		obj.ReferentialQueryFragmentManagedResource_20270101 = nil
		obj.PublishedAnalysisQueryFragment_20270101 = nil
		obj.ApmMetricsQueryFragment_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GuidedTableQuery) MarshalJSON() ([]byte, error) {
	if obj.MetricsLikeQueryFragmentMetrics_20270101 != nil {
		return datadog.Marshal(&obj.MetricsLikeQueryFragmentMetrics_20270101)
	}

	if obj.MetricsLikeQueryFragmentCloudCost_20270101 != nil {
		return datadog.Marshal(&obj.MetricsLikeQueryFragmentCloudCost_20270101)
	}

	if obj.EventsQueryFragmentDefinition_20270101 != nil {
		return datadog.Marshal(&obj.EventsQueryFragmentDefinition_20270101)
	}

	if obj.ReferentialQueryFragmentReferenceTable_20270101 != nil {
		return datadog.Marshal(&obj.ReferentialQueryFragmentReferenceTable_20270101)
	}

	if obj.ReferentialQueryFragmentManagedResource_20270101 != nil {
		return datadog.Marshal(&obj.ReferentialQueryFragmentManagedResource_20270101)
	}

	if obj.PublishedAnalysisQueryFragment_20270101 != nil {
		return datadog.Marshal(&obj.PublishedAnalysisQueryFragment_20270101)
	}

	if obj.ApmMetricsQueryFragment_20270101 != nil {
		return datadog.Marshal(&obj.ApmMetricsQueryFragment_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GuidedTableQuery) GetActualInstance() interface{} {
	if obj.MetricsLikeQueryFragmentMetrics_20270101 != nil {
		return obj.MetricsLikeQueryFragmentMetrics_20270101
	}

	if obj.MetricsLikeQueryFragmentCloudCost_20270101 != nil {
		return obj.MetricsLikeQueryFragmentCloudCost_20270101
	}

	if obj.EventsQueryFragmentDefinition_20270101 != nil {
		return obj.EventsQueryFragmentDefinition_20270101
	}

	if obj.ReferentialQueryFragmentReferenceTable_20270101 != nil {
		return obj.ReferentialQueryFragmentReferenceTable_20270101
	}

	if obj.ReferentialQueryFragmentManagedResource_20270101 != nil {
		return obj.ReferentialQueryFragmentManagedResource_20270101
	}

	if obj.PublishedAnalysisQueryFragment_20270101 != nil {
		return obj.PublishedAnalysisQueryFragment_20270101
	}

	if obj.ApmMetricsQueryFragment_20270101 != nil {
		return obj.ApmMetricsQueryFragment_20270101
	}

	// all schemas are nil
	return nil
}
