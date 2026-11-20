// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeSourceBaseQuery - Test-only generated schema for the dashboard versioning proof of concept.
type CompositeSourceBaseQuery struct {
	FormulaMetricsQueryMetrics_20270101             *FormulaMetricsQueryMetrics
	FormulaMetricsQueryCloudCost_20270101           *FormulaMetricsQueryCloudCost
	FormulaMetricsQueryDatadogUsage_20270101        *FormulaMetricsQueryDatadogUsage
	CompositeEventsPlatformQueryDefinition_20270101 *CompositeEventsPlatformQueryDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaMetricsQueryMetrics_20270101AsCompositeSourceBaseQuery is a convenience function that returns FormulaMetricsQueryMetrics wrapped in CompositeSourceBaseQuery.
func FormulaMetricsQueryMetrics_20270101AsCompositeSourceBaseQuery(v *FormulaMetricsQueryMetrics) CompositeSourceBaseQuery {
	return CompositeSourceBaseQuery{FormulaMetricsQueryMetrics_20270101: v}
}

// FormulaMetricsQueryCloudCost_20270101AsCompositeSourceBaseQuery is a convenience function that returns FormulaMetricsQueryCloudCost wrapped in CompositeSourceBaseQuery.
func FormulaMetricsQueryCloudCost_20270101AsCompositeSourceBaseQuery(v *FormulaMetricsQueryCloudCost) CompositeSourceBaseQuery {
	return CompositeSourceBaseQuery{FormulaMetricsQueryCloudCost_20270101: v}
}

// FormulaMetricsQueryDatadogUsage_20270101AsCompositeSourceBaseQuery is a convenience function that returns FormulaMetricsQueryDatadogUsage wrapped in CompositeSourceBaseQuery.
func FormulaMetricsQueryDatadogUsage_20270101AsCompositeSourceBaseQuery(v *FormulaMetricsQueryDatadogUsage) CompositeSourceBaseQuery {
	return CompositeSourceBaseQuery{FormulaMetricsQueryDatadogUsage_20270101: v}
}

// CompositeEventsPlatformQueryDefinition_20270101AsCompositeSourceBaseQuery is a convenience function that returns CompositeEventsPlatformQueryDefinition wrapped in CompositeSourceBaseQuery.
func CompositeEventsPlatformQueryDefinition_20270101AsCompositeSourceBaseQuery(v *CompositeEventsPlatformQueryDefinition) CompositeSourceBaseQuery {
	return CompositeSourceBaseQuery{CompositeEventsPlatformQueryDefinition_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CompositeSourceBaseQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FormulaMetricsQueryMetrics_20270101
	err = datadog.Unmarshal(data, &obj.FormulaMetricsQueryMetrics_20270101)
	if err == nil {
		if obj.FormulaMetricsQueryMetrics_20270101 != nil && obj.FormulaMetricsQueryMetrics_20270101.UnparsedObject == nil {
			jsonFormulaMetricsQueryMetrics_20270101, _ := datadog.Marshal(obj.FormulaMetricsQueryMetrics_20270101)
			if string(jsonFormulaMetricsQueryMetrics_20270101) == "{}" { // empty struct
				obj.FormulaMetricsQueryMetrics_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaMetricsQueryMetrics_20270101 = nil
		}
	} else {
		obj.FormulaMetricsQueryMetrics_20270101 = nil
	}

	// try to unmarshal data into FormulaMetricsQueryCloudCost_20270101
	err = datadog.Unmarshal(data, &obj.FormulaMetricsQueryCloudCost_20270101)
	if err == nil {
		if obj.FormulaMetricsQueryCloudCost_20270101 != nil && obj.FormulaMetricsQueryCloudCost_20270101.UnparsedObject == nil {
			jsonFormulaMetricsQueryCloudCost_20270101, _ := datadog.Marshal(obj.FormulaMetricsQueryCloudCost_20270101)
			if string(jsonFormulaMetricsQueryCloudCost_20270101) == "{}" { // empty struct
				obj.FormulaMetricsQueryCloudCost_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaMetricsQueryCloudCost_20270101 = nil
		}
	} else {
		obj.FormulaMetricsQueryCloudCost_20270101 = nil
	}

	// try to unmarshal data into FormulaMetricsQueryDatadogUsage_20270101
	err = datadog.Unmarshal(data, &obj.FormulaMetricsQueryDatadogUsage_20270101)
	if err == nil {
		if obj.FormulaMetricsQueryDatadogUsage_20270101 != nil && obj.FormulaMetricsQueryDatadogUsage_20270101.UnparsedObject == nil {
			jsonFormulaMetricsQueryDatadogUsage_20270101, _ := datadog.Marshal(obj.FormulaMetricsQueryDatadogUsage_20270101)
			if string(jsonFormulaMetricsQueryDatadogUsage_20270101) == "{}" { // empty struct
				obj.FormulaMetricsQueryDatadogUsage_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaMetricsQueryDatadogUsage_20270101 = nil
		}
	} else {
		obj.FormulaMetricsQueryDatadogUsage_20270101 = nil
	}

	// try to unmarshal data into CompositeEventsPlatformQueryDefinition_20270101
	err = datadog.Unmarshal(data, &obj.CompositeEventsPlatformQueryDefinition_20270101)
	if err == nil {
		if obj.CompositeEventsPlatformQueryDefinition_20270101 != nil && obj.CompositeEventsPlatformQueryDefinition_20270101.UnparsedObject == nil {
			jsonCompositeEventsPlatformQueryDefinition_20270101, _ := datadog.Marshal(obj.CompositeEventsPlatformQueryDefinition_20270101)
			if string(jsonCompositeEventsPlatformQueryDefinition_20270101) == "{}" { // empty struct
				obj.CompositeEventsPlatformQueryDefinition_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CompositeEventsPlatformQueryDefinition_20270101 = nil
		}
	} else {
		obj.CompositeEventsPlatformQueryDefinition_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaMetricsQueryMetrics_20270101 = nil
		obj.FormulaMetricsQueryCloudCost_20270101 = nil
		obj.FormulaMetricsQueryDatadogUsage_20270101 = nil
		obj.CompositeEventsPlatformQueryDefinition_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CompositeSourceBaseQuery) MarshalJSON() ([]byte, error) {
	if obj.FormulaMetricsQueryMetrics_20270101 != nil {
		return datadog.Marshal(&obj.FormulaMetricsQueryMetrics_20270101)
	}

	if obj.FormulaMetricsQueryCloudCost_20270101 != nil {
		return datadog.Marshal(&obj.FormulaMetricsQueryCloudCost_20270101)
	}

	if obj.FormulaMetricsQueryDatadogUsage_20270101 != nil {
		return datadog.Marshal(&obj.FormulaMetricsQueryDatadogUsage_20270101)
	}

	if obj.CompositeEventsPlatformQueryDefinition_20270101 != nil {
		return datadog.Marshal(&obj.CompositeEventsPlatformQueryDefinition_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CompositeSourceBaseQuery) GetActualInstance() interface{} {
	if obj.FormulaMetricsQueryMetrics_20270101 != nil {
		return obj.FormulaMetricsQueryMetrics_20270101
	}

	if obj.FormulaMetricsQueryCloudCost_20270101 != nil {
		return obj.FormulaMetricsQueryCloudCost_20270101
	}

	if obj.FormulaMetricsQueryDatadogUsage_20270101 != nil {
		return obj.FormulaMetricsQueryDatadogUsage_20270101
	}

	if obj.CompositeEventsPlatformQueryDefinition_20270101 != nil {
		return obj.CompositeEventsPlatformQueryDefinition_20270101
	}

	// all schemas are nil
	return nil
}
