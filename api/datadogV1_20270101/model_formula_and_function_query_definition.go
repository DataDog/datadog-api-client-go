// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionQueryDefinition - Test-only generated schema for the dashboard versioning proof of concept.
type FormulaAndFunctionQueryDefinition struct {
	FormulaMetricsQueryMetrics_20270101               *FormulaMetricsQueryMetrics
	FormulaMetricsQueryCloudCost_20270101             *FormulaMetricsQueryCloudCost
	FormulaMetricsQueryDatadogUsage_20270101          *FormulaMetricsQueryDatadogUsage
	FormulaAndFunctionEventQueryDefinition_20270101   *FormulaAndFunctionEventQueryDefinition
	FormulaAndFunctionProcessQueryDefinition_20270101 *FormulaAndFunctionProcessQueryDefinition
	FormulaApmDependencyStatsQuery_20270101           *FormulaApmDependencyStatsQuery
	FormulaApmResourceStatsQuery_20270101             *FormulaApmResourceStatsQuery
	FormulaApmMetricsQuery_20270101                   *FormulaApmMetricsQuery
	FormulaIncidentsQuery_20270101                    *FormulaIncidentsQuery
	FormulaSloQuery_20270101                          *FormulaSloQuery
	FormulaCaseQuery_20270101                         *FormulaCaseQuery
	FormulaGraphQuery_20270101                        *FormulaGraphQuery
	FormulaUserJourneyQuery_20270101                  *FormulaUserJourneyQuery
	FormulaDoraQuery_20270101                         *FormulaDoraQuery
	FormulaChangeTrackingQuery_20270101               *FormulaChangeTrackingQuery
	FormulaDetectionRulesQuery_20270101               *FormulaDetectionRulesQuery
	AggregateFilteredQuery_20270101                   *AggregateFilteredQuery
	AggregateAugmentedQuery_20270101                  *AggregateAugmentedQuery
	AggregateSegmentQuery_20270101                    *AggregateSegmentQuery
	FormulaDatasetQuery_20270101                      *FormulaDatasetQuery
	FormulaManagedResourceQuery_20270101              *FormulaManagedResourceQuery
	FormulaReferenceTableQuery_20270101               *FormulaReferenceTableQuery
	FormulaRetentionQuery_20270101                    *FormulaRetentionQuery
	FormulaDataQualityMetricsQuery_20270101           *FormulaDataQualityMetricsQuery
	FormulaProductAnalyticsExtendedQuery_20270101     *FormulaProductAnalyticsExtendedQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaMetricsQueryMetrics_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaMetricsQueryMetrics wrapped in FormulaAndFunctionQueryDefinition.
func FormulaMetricsQueryMetrics_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaMetricsQueryMetrics) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaMetricsQueryMetrics_20270101: v}
}

// FormulaMetricsQueryCloudCost_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaMetricsQueryCloudCost wrapped in FormulaAndFunctionQueryDefinition.
func FormulaMetricsQueryCloudCost_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaMetricsQueryCloudCost) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaMetricsQueryCloudCost_20270101: v}
}

// FormulaMetricsQueryDatadogUsage_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaMetricsQueryDatadogUsage wrapped in FormulaAndFunctionQueryDefinition.
func FormulaMetricsQueryDatadogUsage_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaMetricsQueryDatadogUsage) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaMetricsQueryDatadogUsage_20270101: v}
}

// FormulaAndFunctionEventQueryDefinition_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionEventQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionEventQueryDefinition_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionEventQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionEventQueryDefinition_20270101: v}
}

// FormulaAndFunctionProcessQueryDefinition_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionProcessQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionProcessQueryDefinition_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionProcessQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionProcessQueryDefinition_20270101: v}
}

// FormulaApmDependencyStatsQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaApmDependencyStatsQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaApmDependencyStatsQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaApmDependencyStatsQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaApmDependencyStatsQuery_20270101: v}
}

// FormulaApmResourceStatsQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaApmResourceStatsQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaApmResourceStatsQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaApmResourceStatsQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaApmResourceStatsQuery_20270101: v}
}

// FormulaApmMetricsQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaApmMetricsQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaApmMetricsQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaApmMetricsQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaApmMetricsQuery_20270101: v}
}

// FormulaIncidentsQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaIncidentsQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaIncidentsQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaIncidentsQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaIncidentsQuery_20270101: v}
}

// FormulaSloQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaSloQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaSloQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaSloQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaSloQuery_20270101: v}
}

// FormulaCaseQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaCaseQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaCaseQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaCaseQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaCaseQuery_20270101: v}
}

// FormulaGraphQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaGraphQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaGraphQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaGraphQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaGraphQuery_20270101: v}
}

// FormulaUserJourneyQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaUserJourneyQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaUserJourneyQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaUserJourneyQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaUserJourneyQuery_20270101: v}
}

// FormulaDoraQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaDoraQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaDoraQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaDoraQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaDoraQuery_20270101: v}
}

// FormulaChangeTrackingQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaChangeTrackingQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaChangeTrackingQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaChangeTrackingQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaChangeTrackingQuery_20270101: v}
}

// FormulaDetectionRulesQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaDetectionRulesQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaDetectionRulesQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaDetectionRulesQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaDetectionRulesQuery_20270101: v}
}

// AggregateFilteredQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns AggregateFilteredQuery wrapped in FormulaAndFunctionQueryDefinition.
func AggregateFilteredQuery_20270101AsFormulaAndFunctionQueryDefinition(v *AggregateFilteredQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{AggregateFilteredQuery_20270101: v}
}

// AggregateAugmentedQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns AggregateAugmentedQuery wrapped in FormulaAndFunctionQueryDefinition.
func AggregateAugmentedQuery_20270101AsFormulaAndFunctionQueryDefinition(v *AggregateAugmentedQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{AggregateAugmentedQuery_20270101: v}
}

// AggregateSegmentQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns AggregateSegmentQuery wrapped in FormulaAndFunctionQueryDefinition.
func AggregateSegmentQuery_20270101AsFormulaAndFunctionQueryDefinition(v *AggregateSegmentQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{AggregateSegmentQuery_20270101: v}
}

// FormulaDatasetQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaDatasetQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaDatasetQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaDatasetQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaDatasetQuery_20270101: v}
}

// FormulaManagedResourceQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaManagedResourceQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaManagedResourceQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaManagedResourceQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaManagedResourceQuery_20270101: v}
}

// FormulaReferenceTableQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaReferenceTableQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaReferenceTableQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaReferenceTableQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaReferenceTableQuery_20270101: v}
}

// FormulaRetentionQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaRetentionQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaRetentionQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaRetentionQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaRetentionQuery_20270101: v}
}

// FormulaDataQualityMetricsQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaDataQualityMetricsQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaDataQualityMetricsQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaDataQualityMetricsQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaDataQualityMetricsQuery_20270101: v}
}

// FormulaProductAnalyticsExtendedQuery_20270101AsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaProductAnalyticsExtendedQuery wrapped in FormulaAndFunctionQueryDefinition.
func FormulaProductAnalyticsExtendedQuery_20270101AsFormulaAndFunctionQueryDefinition(v *FormulaProductAnalyticsExtendedQuery) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaProductAnalyticsExtendedQuery_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FormulaAndFunctionQueryDefinition) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into FormulaAndFunctionEventQueryDefinition_20270101
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionEventQueryDefinition_20270101)
	if err == nil {
		if obj.FormulaAndFunctionEventQueryDefinition_20270101 != nil && obj.FormulaAndFunctionEventQueryDefinition_20270101.UnparsedObject == nil {
			jsonFormulaAndFunctionEventQueryDefinition_20270101, _ := datadog.Marshal(obj.FormulaAndFunctionEventQueryDefinition_20270101)
			if string(jsonFormulaAndFunctionEventQueryDefinition_20270101) == "{}" { // empty struct
				obj.FormulaAndFunctionEventQueryDefinition_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionEventQueryDefinition_20270101 = nil
		}
	} else {
		obj.FormulaAndFunctionEventQueryDefinition_20270101 = nil
	}

	// try to unmarshal data into FormulaAndFunctionProcessQueryDefinition_20270101
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionProcessQueryDefinition_20270101)
	if err == nil {
		if obj.FormulaAndFunctionProcessQueryDefinition_20270101 != nil && obj.FormulaAndFunctionProcessQueryDefinition_20270101.UnparsedObject == nil {
			jsonFormulaAndFunctionProcessQueryDefinition_20270101, _ := datadog.Marshal(obj.FormulaAndFunctionProcessQueryDefinition_20270101)
			if string(jsonFormulaAndFunctionProcessQueryDefinition_20270101) == "{}" { // empty struct
				obj.FormulaAndFunctionProcessQueryDefinition_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionProcessQueryDefinition_20270101 = nil
		}
	} else {
		obj.FormulaAndFunctionProcessQueryDefinition_20270101 = nil
	}

	// try to unmarshal data into FormulaApmDependencyStatsQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaApmDependencyStatsQuery_20270101)
	if err == nil {
		if obj.FormulaApmDependencyStatsQuery_20270101 != nil && obj.FormulaApmDependencyStatsQuery_20270101.UnparsedObject == nil {
			jsonFormulaApmDependencyStatsQuery_20270101, _ := datadog.Marshal(obj.FormulaApmDependencyStatsQuery_20270101)
			if string(jsonFormulaApmDependencyStatsQuery_20270101) == "{}" { // empty struct
				obj.FormulaApmDependencyStatsQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaApmDependencyStatsQuery_20270101 = nil
		}
	} else {
		obj.FormulaApmDependencyStatsQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaApmResourceStatsQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaApmResourceStatsQuery_20270101)
	if err == nil {
		if obj.FormulaApmResourceStatsQuery_20270101 != nil && obj.FormulaApmResourceStatsQuery_20270101.UnparsedObject == nil {
			jsonFormulaApmResourceStatsQuery_20270101, _ := datadog.Marshal(obj.FormulaApmResourceStatsQuery_20270101)
			if string(jsonFormulaApmResourceStatsQuery_20270101) == "{}" { // empty struct
				obj.FormulaApmResourceStatsQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaApmResourceStatsQuery_20270101 = nil
		}
	} else {
		obj.FormulaApmResourceStatsQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaApmMetricsQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaApmMetricsQuery_20270101)
	if err == nil {
		if obj.FormulaApmMetricsQuery_20270101 != nil && obj.FormulaApmMetricsQuery_20270101.UnparsedObject == nil {
			jsonFormulaApmMetricsQuery_20270101, _ := datadog.Marshal(obj.FormulaApmMetricsQuery_20270101)
			if string(jsonFormulaApmMetricsQuery_20270101) == "{}" { // empty struct
				obj.FormulaApmMetricsQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaApmMetricsQuery_20270101 = nil
		}
	} else {
		obj.FormulaApmMetricsQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaIncidentsQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaIncidentsQuery_20270101)
	if err == nil {
		if obj.FormulaIncidentsQuery_20270101 != nil && obj.FormulaIncidentsQuery_20270101.UnparsedObject == nil {
			jsonFormulaIncidentsQuery_20270101, _ := datadog.Marshal(obj.FormulaIncidentsQuery_20270101)
			if string(jsonFormulaIncidentsQuery_20270101) == "{}" { // empty struct
				obj.FormulaIncidentsQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaIncidentsQuery_20270101 = nil
		}
	} else {
		obj.FormulaIncidentsQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaSloQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaSloQuery_20270101)
	if err == nil {
		if obj.FormulaSloQuery_20270101 != nil && obj.FormulaSloQuery_20270101.UnparsedObject == nil {
			jsonFormulaSloQuery_20270101, _ := datadog.Marshal(obj.FormulaSloQuery_20270101)
			if string(jsonFormulaSloQuery_20270101) == "{}" { // empty struct
				obj.FormulaSloQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaSloQuery_20270101 = nil
		}
	} else {
		obj.FormulaSloQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaCaseQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaCaseQuery_20270101)
	if err == nil {
		if obj.FormulaCaseQuery_20270101 != nil && obj.FormulaCaseQuery_20270101.UnparsedObject == nil {
			jsonFormulaCaseQuery_20270101, _ := datadog.Marshal(obj.FormulaCaseQuery_20270101)
			if string(jsonFormulaCaseQuery_20270101) == "{}" { // empty struct
				obj.FormulaCaseQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaCaseQuery_20270101 = nil
		}
	} else {
		obj.FormulaCaseQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaGraphQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaGraphQuery_20270101)
	if err == nil {
		if obj.FormulaGraphQuery_20270101 != nil && obj.FormulaGraphQuery_20270101.UnparsedObject == nil {
			jsonFormulaGraphQuery_20270101, _ := datadog.Marshal(obj.FormulaGraphQuery_20270101)
			if string(jsonFormulaGraphQuery_20270101) == "{}" { // empty struct
				obj.FormulaGraphQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaGraphQuery_20270101 = nil
		}
	} else {
		obj.FormulaGraphQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaUserJourneyQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaUserJourneyQuery_20270101)
	if err == nil {
		if obj.FormulaUserJourneyQuery_20270101 != nil && obj.FormulaUserJourneyQuery_20270101.UnparsedObject == nil {
			jsonFormulaUserJourneyQuery_20270101, _ := datadog.Marshal(obj.FormulaUserJourneyQuery_20270101)
			if string(jsonFormulaUserJourneyQuery_20270101) == "{}" { // empty struct
				obj.FormulaUserJourneyQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaUserJourneyQuery_20270101 = nil
		}
	} else {
		obj.FormulaUserJourneyQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaDoraQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaDoraQuery_20270101)
	if err == nil {
		if obj.FormulaDoraQuery_20270101 != nil && obj.FormulaDoraQuery_20270101.UnparsedObject == nil {
			jsonFormulaDoraQuery_20270101, _ := datadog.Marshal(obj.FormulaDoraQuery_20270101)
			if string(jsonFormulaDoraQuery_20270101) == "{}" { // empty struct
				obj.FormulaDoraQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaDoraQuery_20270101 = nil
		}
	} else {
		obj.FormulaDoraQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaChangeTrackingQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaChangeTrackingQuery_20270101)
	if err == nil {
		if obj.FormulaChangeTrackingQuery_20270101 != nil && obj.FormulaChangeTrackingQuery_20270101.UnparsedObject == nil {
			jsonFormulaChangeTrackingQuery_20270101, _ := datadog.Marshal(obj.FormulaChangeTrackingQuery_20270101)
			if string(jsonFormulaChangeTrackingQuery_20270101) == "{}" { // empty struct
				obj.FormulaChangeTrackingQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaChangeTrackingQuery_20270101 = nil
		}
	} else {
		obj.FormulaChangeTrackingQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaDetectionRulesQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaDetectionRulesQuery_20270101)
	if err == nil {
		if obj.FormulaDetectionRulesQuery_20270101 != nil && obj.FormulaDetectionRulesQuery_20270101.UnparsedObject == nil {
			jsonFormulaDetectionRulesQuery_20270101, _ := datadog.Marshal(obj.FormulaDetectionRulesQuery_20270101)
			if string(jsonFormulaDetectionRulesQuery_20270101) == "{}" { // empty struct
				obj.FormulaDetectionRulesQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaDetectionRulesQuery_20270101 = nil
		}
	} else {
		obj.FormulaDetectionRulesQuery_20270101 = nil
	}

	// try to unmarshal data into AggregateFilteredQuery_20270101
	err = datadog.Unmarshal(data, &obj.AggregateFilteredQuery_20270101)
	if err == nil {
		if obj.AggregateFilteredQuery_20270101 != nil && obj.AggregateFilteredQuery_20270101.UnparsedObject == nil {
			jsonAggregateFilteredQuery_20270101, _ := datadog.Marshal(obj.AggregateFilteredQuery_20270101)
			if string(jsonAggregateFilteredQuery_20270101) == "{}" { // empty struct
				obj.AggregateFilteredQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.AggregateFilteredQuery_20270101 = nil
		}
	} else {
		obj.AggregateFilteredQuery_20270101 = nil
	}

	// try to unmarshal data into AggregateAugmentedQuery_20270101
	err = datadog.Unmarshal(data, &obj.AggregateAugmentedQuery_20270101)
	if err == nil {
		if obj.AggregateAugmentedQuery_20270101 != nil && obj.AggregateAugmentedQuery_20270101.UnparsedObject == nil {
			jsonAggregateAugmentedQuery_20270101, _ := datadog.Marshal(obj.AggregateAugmentedQuery_20270101)
			if string(jsonAggregateAugmentedQuery_20270101) == "{}" { // empty struct
				obj.AggregateAugmentedQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.AggregateAugmentedQuery_20270101 = nil
		}
	} else {
		obj.AggregateAugmentedQuery_20270101 = nil
	}

	// try to unmarshal data into AggregateSegmentQuery_20270101
	err = datadog.Unmarshal(data, &obj.AggregateSegmentQuery_20270101)
	if err == nil {
		if obj.AggregateSegmentQuery_20270101 != nil && obj.AggregateSegmentQuery_20270101.UnparsedObject == nil {
			jsonAggregateSegmentQuery_20270101, _ := datadog.Marshal(obj.AggregateSegmentQuery_20270101)
			if string(jsonAggregateSegmentQuery_20270101) == "{}" { // empty struct
				obj.AggregateSegmentQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.AggregateSegmentQuery_20270101 = nil
		}
	} else {
		obj.AggregateSegmentQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaDatasetQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaDatasetQuery_20270101)
	if err == nil {
		if obj.FormulaDatasetQuery_20270101 != nil && obj.FormulaDatasetQuery_20270101.UnparsedObject == nil {
			jsonFormulaDatasetQuery_20270101, _ := datadog.Marshal(obj.FormulaDatasetQuery_20270101)
			if string(jsonFormulaDatasetQuery_20270101) == "{}" { // empty struct
				obj.FormulaDatasetQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaDatasetQuery_20270101 = nil
		}
	} else {
		obj.FormulaDatasetQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaManagedResourceQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaManagedResourceQuery_20270101)
	if err == nil {
		if obj.FormulaManagedResourceQuery_20270101 != nil && obj.FormulaManagedResourceQuery_20270101.UnparsedObject == nil {
			jsonFormulaManagedResourceQuery_20270101, _ := datadog.Marshal(obj.FormulaManagedResourceQuery_20270101)
			if string(jsonFormulaManagedResourceQuery_20270101) == "{}" { // empty struct
				obj.FormulaManagedResourceQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaManagedResourceQuery_20270101 = nil
		}
	} else {
		obj.FormulaManagedResourceQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaReferenceTableQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaReferenceTableQuery_20270101)
	if err == nil {
		if obj.FormulaReferenceTableQuery_20270101 != nil && obj.FormulaReferenceTableQuery_20270101.UnparsedObject == nil {
			jsonFormulaReferenceTableQuery_20270101, _ := datadog.Marshal(obj.FormulaReferenceTableQuery_20270101)
			if string(jsonFormulaReferenceTableQuery_20270101) == "{}" { // empty struct
				obj.FormulaReferenceTableQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaReferenceTableQuery_20270101 = nil
		}
	} else {
		obj.FormulaReferenceTableQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaRetentionQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaRetentionQuery_20270101)
	if err == nil {
		if obj.FormulaRetentionQuery_20270101 != nil && obj.FormulaRetentionQuery_20270101.UnparsedObject == nil {
			jsonFormulaRetentionQuery_20270101, _ := datadog.Marshal(obj.FormulaRetentionQuery_20270101)
			if string(jsonFormulaRetentionQuery_20270101) == "{}" { // empty struct
				obj.FormulaRetentionQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaRetentionQuery_20270101 = nil
		}
	} else {
		obj.FormulaRetentionQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaDataQualityMetricsQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaDataQualityMetricsQuery_20270101)
	if err == nil {
		if obj.FormulaDataQualityMetricsQuery_20270101 != nil && obj.FormulaDataQualityMetricsQuery_20270101.UnparsedObject == nil {
			jsonFormulaDataQualityMetricsQuery_20270101, _ := datadog.Marshal(obj.FormulaDataQualityMetricsQuery_20270101)
			if string(jsonFormulaDataQualityMetricsQuery_20270101) == "{}" { // empty struct
				obj.FormulaDataQualityMetricsQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaDataQualityMetricsQuery_20270101 = nil
		}
	} else {
		obj.FormulaDataQualityMetricsQuery_20270101 = nil
	}

	// try to unmarshal data into FormulaProductAnalyticsExtendedQuery_20270101
	err = datadog.Unmarshal(data, &obj.FormulaProductAnalyticsExtendedQuery_20270101)
	if err == nil {
		if obj.FormulaProductAnalyticsExtendedQuery_20270101 != nil && obj.FormulaProductAnalyticsExtendedQuery_20270101.UnparsedObject == nil {
			jsonFormulaProductAnalyticsExtendedQuery_20270101, _ := datadog.Marshal(obj.FormulaProductAnalyticsExtendedQuery_20270101)
			if string(jsonFormulaProductAnalyticsExtendedQuery_20270101) == "{}" { // empty struct
				obj.FormulaProductAnalyticsExtendedQuery_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaProductAnalyticsExtendedQuery_20270101 = nil
		}
	} else {
		obj.FormulaProductAnalyticsExtendedQuery_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaMetricsQueryMetrics_20270101 = nil
		obj.FormulaMetricsQueryCloudCost_20270101 = nil
		obj.FormulaMetricsQueryDatadogUsage_20270101 = nil
		obj.FormulaAndFunctionEventQueryDefinition_20270101 = nil
		obj.FormulaAndFunctionProcessQueryDefinition_20270101 = nil
		obj.FormulaApmDependencyStatsQuery_20270101 = nil
		obj.FormulaApmResourceStatsQuery_20270101 = nil
		obj.FormulaApmMetricsQuery_20270101 = nil
		obj.FormulaIncidentsQuery_20270101 = nil
		obj.FormulaSloQuery_20270101 = nil
		obj.FormulaCaseQuery_20270101 = nil
		obj.FormulaGraphQuery_20270101 = nil
		obj.FormulaUserJourneyQuery_20270101 = nil
		obj.FormulaDoraQuery_20270101 = nil
		obj.FormulaChangeTrackingQuery_20270101 = nil
		obj.FormulaDetectionRulesQuery_20270101 = nil
		obj.AggregateFilteredQuery_20270101 = nil
		obj.AggregateAugmentedQuery_20270101 = nil
		obj.AggregateSegmentQuery_20270101 = nil
		obj.FormulaDatasetQuery_20270101 = nil
		obj.FormulaManagedResourceQuery_20270101 = nil
		obj.FormulaReferenceTableQuery_20270101 = nil
		obj.FormulaRetentionQuery_20270101 = nil
		obj.FormulaDataQualityMetricsQuery_20270101 = nil
		obj.FormulaProductAnalyticsExtendedQuery_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FormulaAndFunctionQueryDefinition) MarshalJSON() ([]byte, error) {
	if obj.FormulaMetricsQueryMetrics_20270101 != nil {
		return datadog.Marshal(&obj.FormulaMetricsQueryMetrics_20270101)
	}

	if obj.FormulaMetricsQueryCloudCost_20270101 != nil {
		return datadog.Marshal(&obj.FormulaMetricsQueryCloudCost_20270101)
	}

	if obj.FormulaMetricsQueryDatadogUsage_20270101 != nil {
		return datadog.Marshal(&obj.FormulaMetricsQueryDatadogUsage_20270101)
	}

	if obj.FormulaAndFunctionEventQueryDefinition_20270101 != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionEventQueryDefinition_20270101)
	}

	if obj.FormulaAndFunctionProcessQueryDefinition_20270101 != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionProcessQueryDefinition_20270101)
	}

	if obj.FormulaApmDependencyStatsQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaApmDependencyStatsQuery_20270101)
	}

	if obj.FormulaApmResourceStatsQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaApmResourceStatsQuery_20270101)
	}

	if obj.FormulaApmMetricsQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaApmMetricsQuery_20270101)
	}

	if obj.FormulaIncidentsQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaIncidentsQuery_20270101)
	}

	if obj.FormulaSloQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaSloQuery_20270101)
	}

	if obj.FormulaCaseQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaCaseQuery_20270101)
	}

	if obj.FormulaGraphQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaGraphQuery_20270101)
	}

	if obj.FormulaUserJourneyQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaUserJourneyQuery_20270101)
	}

	if obj.FormulaDoraQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaDoraQuery_20270101)
	}

	if obj.FormulaChangeTrackingQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaChangeTrackingQuery_20270101)
	}

	if obj.FormulaDetectionRulesQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaDetectionRulesQuery_20270101)
	}

	if obj.AggregateFilteredQuery_20270101 != nil {
		return datadog.Marshal(&obj.AggregateFilteredQuery_20270101)
	}

	if obj.AggregateAugmentedQuery_20270101 != nil {
		return datadog.Marshal(&obj.AggregateAugmentedQuery_20270101)
	}

	if obj.AggregateSegmentQuery_20270101 != nil {
		return datadog.Marshal(&obj.AggregateSegmentQuery_20270101)
	}

	if obj.FormulaDatasetQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaDatasetQuery_20270101)
	}

	if obj.FormulaManagedResourceQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaManagedResourceQuery_20270101)
	}

	if obj.FormulaReferenceTableQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaReferenceTableQuery_20270101)
	}

	if obj.FormulaRetentionQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaRetentionQuery_20270101)
	}

	if obj.FormulaDataQualityMetricsQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaDataQualityMetricsQuery_20270101)
	}

	if obj.FormulaProductAnalyticsExtendedQuery_20270101 != nil {
		return datadog.Marshal(&obj.FormulaProductAnalyticsExtendedQuery_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FormulaAndFunctionQueryDefinition) GetActualInstance() interface{} {
	if obj.FormulaMetricsQueryMetrics_20270101 != nil {
		return obj.FormulaMetricsQueryMetrics_20270101
	}

	if obj.FormulaMetricsQueryCloudCost_20270101 != nil {
		return obj.FormulaMetricsQueryCloudCost_20270101
	}

	if obj.FormulaMetricsQueryDatadogUsage_20270101 != nil {
		return obj.FormulaMetricsQueryDatadogUsage_20270101
	}

	if obj.FormulaAndFunctionEventQueryDefinition_20270101 != nil {
		return obj.FormulaAndFunctionEventQueryDefinition_20270101
	}

	if obj.FormulaAndFunctionProcessQueryDefinition_20270101 != nil {
		return obj.FormulaAndFunctionProcessQueryDefinition_20270101
	}

	if obj.FormulaApmDependencyStatsQuery_20270101 != nil {
		return obj.FormulaApmDependencyStatsQuery_20270101
	}

	if obj.FormulaApmResourceStatsQuery_20270101 != nil {
		return obj.FormulaApmResourceStatsQuery_20270101
	}

	if obj.FormulaApmMetricsQuery_20270101 != nil {
		return obj.FormulaApmMetricsQuery_20270101
	}

	if obj.FormulaIncidentsQuery_20270101 != nil {
		return obj.FormulaIncidentsQuery_20270101
	}

	if obj.FormulaSloQuery_20270101 != nil {
		return obj.FormulaSloQuery_20270101
	}

	if obj.FormulaCaseQuery_20270101 != nil {
		return obj.FormulaCaseQuery_20270101
	}

	if obj.FormulaGraphQuery_20270101 != nil {
		return obj.FormulaGraphQuery_20270101
	}

	if obj.FormulaUserJourneyQuery_20270101 != nil {
		return obj.FormulaUserJourneyQuery_20270101
	}

	if obj.FormulaDoraQuery_20270101 != nil {
		return obj.FormulaDoraQuery_20270101
	}

	if obj.FormulaChangeTrackingQuery_20270101 != nil {
		return obj.FormulaChangeTrackingQuery_20270101
	}

	if obj.FormulaDetectionRulesQuery_20270101 != nil {
		return obj.FormulaDetectionRulesQuery_20270101
	}

	if obj.AggregateFilteredQuery_20270101 != nil {
		return obj.AggregateFilteredQuery_20270101
	}

	if obj.AggregateAugmentedQuery_20270101 != nil {
		return obj.AggregateAugmentedQuery_20270101
	}

	if obj.AggregateSegmentQuery_20270101 != nil {
		return obj.AggregateSegmentQuery_20270101
	}

	if obj.FormulaDatasetQuery_20270101 != nil {
		return obj.FormulaDatasetQuery_20270101
	}

	if obj.FormulaManagedResourceQuery_20270101 != nil {
		return obj.FormulaManagedResourceQuery_20270101
	}

	if obj.FormulaReferenceTableQuery_20270101 != nil {
		return obj.FormulaReferenceTableQuery_20270101
	}

	if obj.FormulaRetentionQuery_20270101 != nil {
		return obj.FormulaRetentionQuery_20270101
	}

	if obj.FormulaDataQualityMetricsQuery_20270101 != nil {
		return obj.FormulaDataQualityMetricsQuery_20270101
	}

	if obj.FormulaProductAnalyticsExtendedQuery_20270101 != nil {
		return obj.FormulaProductAnalyticsExtendedQuery_20270101
	}

	// all schemas are nil
	return nil
}
