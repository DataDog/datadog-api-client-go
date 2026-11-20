// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AdvancedQueryDataset - Test-only generated schema for the dashboard versioning proof of concept.
type AdvancedQueryDataset struct {
	EventPlatformDatasetDefinition_20270101       *EventPlatformDatasetDefinition
	ReferentialDatasetReferenceTable_20270101     *ReferentialDatasetReferenceTable
	ReferentialDatasetManagedResource_20270101    *ReferentialDatasetManagedResource
	ActionsDatastoreDataset_20270101              *ActionsDatastoreDataset
	MetricDatasetCloudCostScalar_20270101         *MetricDatasetCloudCostScalar
	MetricDatasetMetricScalar_20270101            *MetricDatasetMetricScalar
	FormulaDataset_20270101                       *FormulaDataset
	AnalysisDataset_20270101                      *AnalysisDataset
	PublishedAnalysis_20270101                    *PublishedAnalysis
	SnowflakeQueryDataset_20270101                *SnowflakeQueryDataset
	BigQueryDataset_20270101                      *BigQueryDataset
	ProductAnalyticsJourneyDataset_20270101       *ProductAnalyticsJourneyDataset
	HistoricalDiscoveredProcessesDataset_20270101 *HistoricalDiscoveredProcessesDataset

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// EventPlatformDatasetDefinition_20270101AsAdvancedQueryDataset is a convenience function that returns EventPlatformDatasetDefinition wrapped in AdvancedQueryDataset.
func EventPlatformDatasetDefinition_20270101AsAdvancedQueryDataset(v *EventPlatformDatasetDefinition) AdvancedQueryDataset {
	return AdvancedQueryDataset{EventPlatformDatasetDefinition_20270101: v}
}

// ReferentialDatasetReferenceTable_20270101AsAdvancedQueryDataset is a convenience function that returns ReferentialDatasetReferenceTable wrapped in AdvancedQueryDataset.
func ReferentialDatasetReferenceTable_20270101AsAdvancedQueryDataset(v *ReferentialDatasetReferenceTable) AdvancedQueryDataset {
	return AdvancedQueryDataset{ReferentialDatasetReferenceTable_20270101: v}
}

// ReferentialDatasetManagedResource_20270101AsAdvancedQueryDataset is a convenience function that returns ReferentialDatasetManagedResource wrapped in AdvancedQueryDataset.
func ReferentialDatasetManagedResource_20270101AsAdvancedQueryDataset(v *ReferentialDatasetManagedResource) AdvancedQueryDataset {
	return AdvancedQueryDataset{ReferentialDatasetManagedResource_20270101: v}
}

// ActionsDatastoreDataset_20270101AsAdvancedQueryDataset is a convenience function that returns ActionsDatastoreDataset wrapped in AdvancedQueryDataset.
func ActionsDatastoreDataset_20270101AsAdvancedQueryDataset(v *ActionsDatastoreDataset) AdvancedQueryDataset {
	return AdvancedQueryDataset{ActionsDatastoreDataset_20270101: v}
}

// MetricDatasetCloudCostScalar_20270101AsAdvancedQueryDataset is a convenience function that returns MetricDatasetCloudCostScalar wrapped in AdvancedQueryDataset.
func MetricDatasetCloudCostScalar_20270101AsAdvancedQueryDataset(v *MetricDatasetCloudCostScalar) AdvancedQueryDataset {
	return AdvancedQueryDataset{MetricDatasetCloudCostScalar_20270101: v}
}

// MetricDatasetMetricScalar_20270101AsAdvancedQueryDataset is a convenience function that returns MetricDatasetMetricScalar wrapped in AdvancedQueryDataset.
func MetricDatasetMetricScalar_20270101AsAdvancedQueryDataset(v *MetricDatasetMetricScalar) AdvancedQueryDataset {
	return AdvancedQueryDataset{MetricDatasetMetricScalar_20270101: v}
}

// FormulaDataset_20270101AsAdvancedQueryDataset is a convenience function that returns FormulaDataset wrapped in AdvancedQueryDataset.
func FormulaDataset_20270101AsAdvancedQueryDataset(v *FormulaDataset) AdvancedQueryDataset {
	return AdvancedQueryDataset{FormulaDataset_20270101: v}
}

// AnalysisDataset_20270101AsAdvancedQueryDataset is a convenience function that returns AnalysisDataset wrapped in AdvancedQueryDataset.
func AnalysisDataset_20270101AsAdvancedQueryDataset(v *AnalysisDataset) AdvancedQueryDataset {
	return AdvancedQueryDataset{AnalysisDataset_20270101: v}
}

// PublishedAnalysis_20270101AsAdvancedQueryDataset is a convenience function that returns PublishedAnalysis wrapped in AdvancedQueryDataset.
func PublishedAnalysis_20270101AsAdvancedQueryDataset(v *PublishedAnalysis) AdvancedQueryDataset {
	return AdvancedQueryDataset{PublishedAnalysis_20270101: v}
}

// SnowflakeQueryDataset_20270101AsAdvancedQueryDataset is a convenience function that returns SnowflakeQueryDataset wrapped in AdvancedQueryDataset.
func SnowflakeQueryDataset_20270101AsAdvancedQueryDataset(v *SnowflakeQueryDataset) AdvancedQueryDataset {
	return AdvancedQueryDataset{SnowflakeQueryDataset_20270101: v}
}

// BigQueryDataset_20270101AsAdvancedQueryDataset is a convenience function that returns BigQueryDataset wrapped in AdvancedQueryDataset.
func BigQueryDataset_20270101AsAdvancedQueryDataset(v *BigQueryDataset) AdvancedQueryDataset {
	return AdvancedQueryDataset{BigQueryDataset_20270101: v}
}

// ProductAnalyticsJourneyDataset_20270101AsAdvancedQueryDataset is a convenience function that returns ProductAnalyticsJourneyDataset wrapped in AdvancedQueryDataset.
func ProductAnalyticsJourneyDataset_20270101AsAdvancedQueryDataset(v *ProductAnalyticsJourneyDataset) AdvancedQueryDataset {
	return AdvancedQueryDataset{ProductAnalyticsJourneyDataset_20270101: v}
}

// HistoricalDiscoveredProcessesDataset_20270101AsAdvancedQueryDataset is a convenience function that returns HistoricalDiscoveredProcessesDataset wrapped in AdvancedQueryDataset.
func HistoricalDiscoveredProcessesDataset_20270101AsAdvancedQueryDataset(v *HistoricalDiscoveredProcessesDataset) AdvancedQueryDataset {
	return AdvancedQueryDataset{HistoricalDiscoveredProcessesDataset_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AdvancedQueryDataset) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EventPlatformDatasetDefinition_20270101
	err = datadog.Unmarshal(data, &obj.EventPlatformDatasetDefinition_20270101)
	if err == nil {
		if obj.EventPlatformDatasetDefinition_20270101 != nil && obj.EventPlatformDatasetDefinition_20270101.UnparsedObject == nil {
			jsonEventPlatformDatasetDefinition_20270101, _ := datadog.Marshal(obj.EventPlatformDatasetDefinition_20270101)
			if string(jsonEventPlatformDatasetDefinition_20270101) == "{}" { // empty struct
				obj.EventPlatformDatasetDefinition_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventPlatformDatasetDefinition_20270101 = nil
		}
	} else {
		obj.EventPlatformDatasetDefinition_20270101 = nil
	}

	// try to unmarshal data into ReferentialDatasetReferenceTable_20270101
	err = datadog.Unmarshal(data, &obj.ReferentialDatasetReferenceTable_20270101)
	if err == nil {
		if obj.ReferentialDatasetReferenceTable_20270101 != nil && obj.ReferentialDatasetReferenceTable_20270101.UnparsedObject == nil {
			jsonReferentialDatasetReferenceTable_20270101, _ := datadog.Marshal(obj.ReferentialDatasetReferenceTable_20270101)
			if string(jsonReferentialDatasetReferenceTable_20270101) == "{}" { // empty struct
				obj.ReferentialDatasetReferenceTable_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ReferentialDatasetReferenceTable_20270101 = nil
		}
	} else {
		obj.ReferentialDatasetReferenceTable_20270101 = nil
	}

	// try to unmarshal data into ReferentialDatasetManagedResource_20270101
	err = datadog.Unmarshal(data, &obj.ReferentialDatasetManagedResource_20270101)
	if err == nil {
		if obj.ReferentialDatasetManagedResource_20270101 != nil && obj.ReferentialDatasetManagedResource_20270101.UnparsedObject == nil {
			jsonReferentialDatasetManagedResource_20270101, _ := datadog.Marshal(obj.ReferentialDatasetManagedResource_20270101)
			if string(jsonReferentialDatasetManagedResource_20270101) == "{}" { // empty struct
				obj.ReferentialDatasetManagedResource_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ReferentialDatasetManagedResource_20270101 = nil
		}
	} else {
		obj.ReferentialDatasetManagedResource_20270101 = nil
	}

	// try to unmarshal data into ActionsDatastoreDataset_20270101
	err = datadog.Unmarshal(data, &obj.ActionsDatastoreDataset_20270101)
	if err == nil {
		if obj.ActionsDatastoreDataset_20270101 != nil && obj.ActionsDatastoreDataset_20270101.UnparsedObject == nil {
			jsonActionsDatastoreDataset_20270101, _ := datadog.Marshal(obj.ActionsDatastoreDataset_20270101)
			if string(jsonActionsDatastoreDataset_20270101) == "{}" { // empty struct
				obj.ActionsDatastoreDataset_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ActionsDatastoreDataset_20270101 = nil
		}
	} else {
		obj.ActionsDatastoreDataset_20270101 = nil
	}

	// try to unmarshal data into MetricDatasetCloudCostScalar_20270101
	err = datadog.Unmarshal(data, &obj.MetricDatasetCloudCostScalar_20270101)
	if err == nil {
		if obj.MetricDatasetCloudCostScalar_20270101 != nil && obj.MetricDatasetCloudCostScalar_20270101.UnparsedObject == nil {
			jsonMetricDatasetCloudCostScalar_20270101, _ := datadog.Marshal(obj.MetricDatasetCloudCostScalar_20270101)
			if string(jsonMetricDatasetCloudCostScalar_20270101) == "{}" { // empty struct
				obj.MetricDatasetCloudCostScalar_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.MetricDatasetCloudCostScalar_20270101 = nil
		}
	} else {
		obj.MetricDatasetCloudCostScalar_20270101 = nil
	}

	// try to unmarshal data into MetricDatasetMetricScalar_20270101
	err = datadog.Unmarshal(data, &obj.MetricDatasetMetricScalar_20270101)
	if err == nil {
		if obj.MetricDatasetMetricScalar_20270101 != nil && obj.MetricDatasetMetricScalar_20270101.UnparsedObject == nil {
			jsonMetricDatasetMetricScalar_20270101, _ := datadog.Marshal(obj.MetricDatasetMetricScalar_20270101)
			if string(jsonMetricDatasetMetricScalar_20270101) == "{}" { // empty struct
				obj.MetricDatasetMetricScalar_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.MetricDatasetMetricScalar_20270101 = nil
		}
	} else {
		obj.MetricDatasetMetricScalar_20270101 = nil
	}

	// try to unmarshal data into FormulaDataset_20270101
	err = datadog.Unmarshal(data, &obj.FormulaDataset_20270101)
	if err == nil {
		if obj.FormulaDataset_20270101 != nil && obj.FormulaDataset_20270101.UnparsedObject == nil {
			jsonFormulaDataset_20270101, _ := datadog.Marshal(obj.FormulaDataset_20270101)
			if string(jsonFormulaDataset_20270101) == "{}" { // empty struct
				obj.FormulaDataset_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FormulaDataset_20270101 = nil
		}
	} else {
		obj.FormulaDataset_20270101 = nil
	}

	// try to unmarshal data into AnalysisDataset_20270101
	err = datadog.Unmarshal(data, &obj.AnalysisDataset_20270101)
	if err == nil {
		if obj.AnalysisDataset_20270101 != nil && obj.AnalysisDataset_20270101.UnparsedObject == nil {
			jsonAnalysisDataset_20270101, _ := datadog.Marshal(obj.AnalysisDataset_20270101)
			if string(jsonAnalysisDataset_20270101) == "{}" { // empty struct
				obj.AnalysisDataset_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.AnalysisDataset_20270101 = nil
		}
	} else {
		obj.AnalysisDataset_20270101 = nil
	}

	// try to unmarshal data into PublishedAnalysis_20270101
	err = datadog.Unmarshal(data, &obj.PublishedAnalysis_20270101)
	if err == nil {
		if obj.PublishedAnalysis_20270101 != nil && obj.PublishedAnalysis_20270101.UnparsedObject == nil {
			jsonPublishedAnalysis_20270101, _ := datadog.Marshal(obj.PublishedAnalysis_20270101)
			if string(jsonPublishedAnalysis_20270101) == "{}" { // empty struct
				obj.PublishedAnalysis_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.PublishedAnalysis_20270101 = nil
		}
	} else {
		obj.PublishedAnalysis_20270101 = nil
	}

	// try to unmarshal data into SnowflakeQueryDataset_20270101
	err = datadog.Unmarshal(data, &obj.SnowflakeQueryDataset_20270101)
	if err == nil {
		if obj.SnowflakeQueryDataset_20270101 != nil && obj.SnowflakeQueryDataset_20270101.UnparsedObject == nil {
			jsonSnowflakeQueryDataset_20270101, _ := datadog.Marshal(obj.SnowflakeQueryDataset_20270101)
			if string(jsonSnowflakeQueryDataset_20270101) == "{}" { // empty struct
				obj.SnowflakeQueryDataset_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SnowflakeQueryDataset_20270101 = nil
		}
	} else {
		obj.SnowflakeQueryDataset_20270101 = nil
	}

	// try to unmarshal data into BigQueryDataset_20270101
	err = datadog.Unmarshal(data, &obj.BigQueryDataset_20270101)
	if err == nil {
		if obj.BigQueryDataset_20270101 != nil && obj.BigQueryDataset_20270101.UnparsedObject == nil {
			jsonBigQueryDataset_20270101, _ := datadog.Marshal(obj.BigQueryDataset_20270101)
			if string(jsonBigQueryDataset_20270101) == "{}" { // empty struct
				obj.BigQueryDataset_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.BigQueryDataset_20270101 = nil
		}
	} else {
		obj.BigQueryDataset_20270101 = nil
	}

	// try to unmarshal data into ProductAnalyticsJourneyDataset_20270101
	err = datadog.Unmarshal(data, &obj.ProductAnalyticsJourneyDataset_20270101)
	if err == nil {
		if obj.ProductAnalyticsJourneyDataset_20270101 != nil && obj.ProductAnalyticsJourneyDataset_20270101.UnparsedObject == nil {
			jsonProductAnalyticsJourneyDataset_20270101, _ := datadog.Marshal(obj.ProductAnalyticsJourneyDataset_20270101)
			if string(jsonProductAnalyticsJourneyDataset_20270101) == "{}" { // empty struct
				obj.ProductAnalyticsJourneyDataset_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ProductAnalyticsJourneyDataset_20270101 = nil
		}
	} else {
		obj.ProductAnalyticsJourneyDataset_20270101 = nil
	}

	// try to unmarshal data into HistoricalDiscoveredProcessesDataset_20270101
	err = datadog.Unmarshal(data, &obj.HistoricalDiscoveredProcessesDataset_20270101)
	if err == nil {
		if obj.HistoricalDiscoveredProcessesDataset_20270101 != nil && obj.HistoricalDiscoveredProcessesDataset_20270101.UnparsedObject == nil {
			jsonHistoricalDiscoveredProcessesDataset_20270101, _ := datadog.Marshal(obj.HistoricalDiscoveredProcessesDataset_20270101)
			if string(jsonHistoricalDiscoveredProcessesDataset_20270101) == "{}" { // empty struct
				obj.HistoricalDiscoveredProcessesDataset_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.HistoricalDiscoveredProcessesDataset_20270101 = nil
		}
	} else {
		obj.HistoricalDiscoveredProcessesDataset_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.EventPlatformDatasetDefinition_20270101 = nil
		obj.ReferentialDatasetReferenceTable_20270101 = nil
		obj.ReferentialDatasetManagedResource_20270101 = nil
		obj.ActionsDatastoreDataset_20270101 = nil
		obj.MetricDatasetCloudCostScalar_20270101 = nil
		obj.MetricDatasetMetricScalar_20270101 = nil
		obj.FormulaDataset_20270101 = nil
		obj.AnalysisDataset_20270101 = nil
		obj.PublishedAnalysis_20270101 = nil
		obj.SnowflakeQueryDataset_20270101 = nil
		obj.BigQueryDataset_20270101 = nil
		obj.ProductAnalyticsJourneyDataset_20270101 = nil
		obj.HistoricalDiscoveredProcessesDataset_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AdvancedQueryDataset) MarshalJSON() ([]byte, error) {
	if obj.EventPlatformDatasetDefinition_20270101 != nil {
		return datadog.Marshal(&obj.EventPlatformDatasetDefinition_20270101)
	}

	if obj.ReferentialDatasetReferenceTable_20270101 != nil {
		return datadog.Marshal(&obj.ReferentialDatasetReferenceTable_20270101)
	}

	if obj.ReferentialDatasetManagedResource_20270101 != nil {
		return datadog.Marshal(&obj.ReferentialDatasetManagedResource_20270101)
	}

	if obj.ActionsDatastoreDataset_20270101 != nil {
		return datadog.Marshal(&obj.ActionsDatastoreDataset_20270101)
	}

	if obj.MetricDatasetCloudCostScalar_20270101 != nil {
		return datadog.Marshal(&obj.MetricDatasetCloudCostScalar_20270101)
	}

	if obj.MetricDatasetMetricScalar_20270101 != nil {
		return datadog.Marshal(&obj.MetricDatasetMetricScalar_20270101)
	}

	if obj.FormulaDataset_20270101 != nil {
		return datadog.Marshal(&obj.FormulaDataset_20270101)
	}

	if obj.AnalysisDataset_20270101 != nil {
		return datadog.Marshal(&obj.AnalysisDataset_20270101)
	}

	if obj.PublishedAnalysis_20270101 != nil {
		return datadog.Marshal(&obj.PublishedAnalysis_20270101)
	}

	if obj.SnowflakeQueryDataset_20270101 != nil {
		return datadog.Marshal(&obj.SnowflakeQueryDataset_20270101)
	}

	if obj.BigQueryDataset_20270101 != nil {
		return datadog.Marshal(&obj.BigQueryDataset_20270101)
	}

	if obj.ProductAnalyticsJourneyDataset_20270101 != nil {
		return datadog.Marshal(&obj.ProductAnalyticsJourneyDataset_20270101)
	}

	if obj.HistoricalDiscoveredProcessesDataset_20270101 != nil {
		return datadog.Marshal(&obj.HistoricalDiscoveredProcessesDataset_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AdvancedQueryDataset) GetActualInstance() interface{} {
	if obj.EventPlatformDatasetDefinition_20270101 != nil {
		return obj.EventPlatformDatasetDefinition_20270101
	}

	if obj.ReferentialDatasetReferenceTable_20270101 != nil {
		return obj.ReferentialDatasetReferenceTable_20270101
	}

	if obj.ReferentialDatasetManagedResource_20270101 != nil {
		return obj.ReferentialDatasetManagedResource_20270101
	}

	if obj.ActionsDatastoreDataset_20270101 != nil {
		return obj.ActionsDatastoreDataset_20270101
	}

	if obj.MetricDatasetCloudCostScalar_20270101 != nil {
		return obj.MetricDatasetCloudCostScalar_20270101
	}

	if obj.MetricDatasetMetricScalar_20270101 != nil {
		return obj.MetricDatasetMetricScalar_20270101
	}

	if obj.FormulaDataset_20270101 != nil {
		return obj.FormulaDataset_20270101
	}

	if obj.AnalysisDataset_20270101 != nil {
		return obj.AnalysisDataset_20270101
	}

	if obj.PublishedAnalysis_20270101 != nil {
		return obj.PublishedAnalysis_20270101
	}

	if obj.SnowflakeQueryDataset_20270101 != nil {
		return obj.SnowflakeQueryDataset_20270101
	}

	if obj.BigQueryDataset_20270101 != nil {
		return obj.BigQueryDataset_20270101
	}

	if obj.ProductAnalyticsJourneyDataset_20270101 != nil {
		return obj.ProductAnalyticsJourneyDataset_20270101
	}

	if obj.HistoricalDiscoveredProcessesDataset_20270101 != nil {
		return obj.HistoricalDiscoveredProcessesDataset_20270101
	}

	// all schemas are nil
	return nil
}
