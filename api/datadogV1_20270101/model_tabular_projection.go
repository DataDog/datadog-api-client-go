// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TabularProjection - Test-only generated schema for the dashboard versioning proof of concept.
type TabularProjection struct {
	ScatterplotProjection_20270101      *ScatterplotProjection
	TreeDataProjectionBarChart_20270101 *TreeDataProjectionBarChart
	TreeDataProjectionSunburst_20270101 *TreeDataProjectionSunburst
	TreeDataProjectionToplist_20270101  *TreeDataProjectionToplist
	TreeDataProjectionTreemap_20270101  *TreeDataProjectionTreemap
	QueryValueProjection_20270101       *QueryValueProjection
	QueryTableProjection_20270101       *QueryTableProjection
	RunChartProjection_20270101         *RunChartProjection

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ScatterplotProjection_20270101AsTabularProjection is a convenience function that returns ScatterplotProjection wrapped in TabularProjection.
func ScatterplotProjection_20270101AsTabularProjection(v *ScatterplotProjection) TabularProjection {
	return TabularProjection{ScatterplotProjection_20270101: v}
}

// TreeDataProjectionBarChart_20270101AsTabularProjection is a convenience function that returns TreeDataProjectionBarChart wrapped in TabularProjection.
func TreeDataProjectionBarChart_20270101AsTabularProjection(v *TreeDataProjectionBarChart) TabularProjection {
	return TabularProjection{TreeDataProjectionBarChart_20270101: v}
}

// TreeDataProjectionSunburst_20270101AsTabularProjection is a convenience function that returns TreeDataProjectionSunburst wrapped in TabularProjection.
func TreeDataProjectionSunburst_20270101AsTabularProjection(v *TreeDataProjectionSunburst) TabularProjection {
	return TabularProjection{TreeDataProjectionSunburst_20270101: v}
}

// TreeDataProjectionToplist_20270101AsTabularProjection is a convenience function that returns TreeDataProjectionToplist wrapped in TabularProjection.
func TreeDataProjectionToplist_20270101AsTabularProjection(v *TreeDataProjectionToplist) TabularProjection {
	return TabularProjection{TreeDataProjectionToplist_20270101: v}
}

// TreeDataProjectionTreemap_20270101AsTabularProjection is a convenience function that returns TreeDataProjectionTreemap wrapped in TabularProjection.
func TreeDataProjectionTreemap_20270101AsTabularProjection(v *TreeDataProjectionTreemap) TabularProjection {
	return TabularProjection{TreeDataProjectionTreemap_20270101: v}
}

// QueryValueProjection_20270101AsTabularProjection is a convenience function that returns QueryValueProjection wrapped in TabularProjection.
func QueryValueProjection_20270101AsTabularProjection(v *QueryValueProjection) TabularProjection {
	return TabularProjection{QueryValueProjection_20270101: v}
}

// QueryTableProjection_20270101AsTabularProjection is a convenience function that returns QueryTableProjection wrapped in TabularProjection.
func QueryTableProjection_20270101AsTabularProjection(v *QueryTableProjection) TabularProjection {
	return TabularProjection{QueryTableProjection_20270101: v}
}

// RunChartProjection_20270101AsTabularProjection is a convenience function that returns RunChartProjection wrapped in TabularProjection.
func RunChartProjection_20270101AsTabularProjection(v *RunChartProjection) TabularProjection {
	return TabularProjection{RunChartProjection_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TabularProjection) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ScatterplotProjection_20270101
	err = datadog.Unmarshal(data, &obj.ScatterplotProjection_20270101)
	if err == nil {
		if obj.ScatterplotProjection_20270101 != nil && obj.ScatterplotProjection_20270101.UnparsedObject == nil {
			jsonScatterplotProjection_20270101, _ := datadog.Marshal(obj.ScatterplotProjection_20270101)
			if string(jsonScatterplotProjection_20270101) == "{}" { // empty struct
				obj.ScatterplotProjection_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotProjection_20270101 = nil
		}
	} else {
		obj.ScatterplotProjection_20270101 = nil
	}

	// try to unmarshal data into TreeDataProjectionBarChart_20270101
	err = datadog.Unmarshal(data, &obj.TreeDataProjectionBarChart_20270101)
	if err == nil {
		if obj.TreeDataProjectionBarChart_20270101 != nil && obj.TreeDataProjectionBarChart_20270101.UnparsedObject == nil {
			jsonTreeDataProjectionBarChart_20270101, _ := datadog.Marshal(obj.TreeDataProjectionBarChart_20270101)
			if string(jsonTreeDataProjectionBarChart_20270101) == "{}" { // empty struct
				obj.TreeDataProjectionBarChart_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TreeDataProjectionBarChart_20270101 = nil
		}
	} else {
		obj.TreeDataProjectionBarChart_20270101 = nil
	}

	// try to unmarshal data into TreeDataProjectionSunburst_20270101
	err = datadog.Unmarshal(data, &obj.TreeDataProjectionSunburst_20270101)
	if err == nil {
		if obj.TreeDataProjectionSunburst_20270101 != nil && obj.TreeDataProjectionSunburst_20270101.UnparsedObject == nil {
			jsonTreeDataProjectionSunburst_20270101, _ := datadog.Marshal(obj.TreeDataProjectionSunburst_20270101)
			if string(jsonTreeDataProjectionSunburst_20270101) == "{}" { // empty struct
				obj.TreeDataProjectionSunburst_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TreeDataProjectionSunburst_20270101 = nil
		}
	} else {
		obj.TreeDataProjectionSunburst_20270101 = nil
	}

	// try to unmarshal data into TreeDataProjectionToplist_20270101
	err = datadog.Unmarshal(data, &obj.TreeDataProjectionToplist_20270101)
	if err == nil {
		if obj.TreeDataProjectionToplist_20270101 != nil && obj.TreeDataProjectionToplist_20270101.UnparsedObject == nil {
			jsonTreeDataProjectionToplist_20270101, _ := datadog.Marshal(obj.TreeDataProjectionToplist_20270101)
			if string(jsonTreeDataProjectionToplist_20270101) == "{}" { // empty struct
				obj.TreeDataProjectionToplist_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TreeDataProjectionToplist_20270101 = nil
		}
	} else {
		obj.TreeDataProjectionToplist_20270101 = nil
	}

	// try to unmarshal data into TreeDataProjectionTreemap_20270101
	err = datadog.Unmarshal(data, &obj.TreeDataProjectionTreemap_20270101)
	if err == nil {
		if obj.TreeDataProjectionTreemap_20270101 != nil && obj.TreeDataProjectionTreemap_20270101.UnparsedObject == nil {
			jsonTreeDataProjectionTreemap_20270101, _ := datadog.Marshal(obj.TreeDataProjectionTreemap_20270101)
			if string(jsonTreeDataProjectionTreemap_20270101) == "{}" { // empty struct
				obj.TreeDataProjectionTreemap_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TreeDataProjectionTreemap_20270101 = nil
		}
	} else {
		obj.TreeDataProjectionTreemap_20270101 = nil
	}

	// try to unmarshal data into QueryValueProjection_20270101
	err = datadog.Unmarshal(data, &obj.QueryValueProjection_20270101)
	if err == nil {
		if obj.QueryValueProjection_20270101 != nil && obj.QueryValueProjection_20270101.UnparsedObject == nil {
			jsonQueryValueProjection_20270101, _ := datadog.Marshal(obj.QueryValueProjection_20270101)
			if string(jsonQueryValueProjection_20270101) == "{}" { // empty struct
				obj.QueryValueProjection_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.QueryValueProjection_20270101 = nil
		}
	} else {
		obj.QueryValueProjection_20270101 = nil
	}

	// try to unmarshal data into QueryTableProjection_20270101
	err = datadog.Unmarshal(data, &obj.QueryTableProjection_20270101)
	if err == nil {
		if obj.QueryTableProjection_20270101 != nil && obj.QueryTableProjection_20270101.UnparsedObject == nil {
			jsonQueryTableProjection_20270101, _ := datadog.Marshal(obj.QueryTableProjection_20270101)
			if string(jsonQueryTableProjection_20270101) == "{}" { // empty struct
				obj.QueryTableProjection_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.QueryTableProjection_20270101 = nil
		}
	} else {
		obj.QueryTableProjection_20270101 = nil
	}

	// try to unmarshal data into RunChartProjection_20270101
	err = datadog.Unmarshal(data, &obj.RunChartProjection_20270101)
	if err == nil {
		if obj.RunChartProjection_20270101 != nil && obj.RunChartProjection_20270101.UnparsedObject == nil {
			jsonRunChartProjection_20270101, _ := datadog.Marshal(obj.RunChartProjection_20270101)
			if string(jsonRunChartProjection_20270101) == "{}" { // empty struct
				obj.RunChartProjection_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RunChartProjection_20270101 = nil
		}
	} else {
		obj.RunChartProjection_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ScatterplotProjection_20270101 = nil
		obj.TreeDataProjectionBarChart_20270101 = nil
		obj.TreeDataProjectionSunburst_20270101 = nil
		obj.TreeDataProjectionToplist_20270101 = nil
		obj.TreeDataProjectionTreemap_20270101 = nil
		obj.QueryValueProjection_20270101 = nil
		obj.QueryTableProjection_20270101 = nil
		obj.RunChartProjection_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TabularProjection) MarshalJSON() ([]byte, error) {
	if obj.ScatterplotProjection_20270101 != nil {
		return datadog.Marshal(&obj.ScatterplotProjection_20270101)
	}

	if obj.TreeDataProjectionBarChart_20270101 != nil {
		return datadog.Marshal(&obj.TreeDataProjectionBarChart_20270101)
	}

	if obj.TreeDataProjectionSunburst_20270101 != nil {
		return datadog.Marshal(&obj.TreeDataProjectionSunburst_20270101)
	}

	if obj.TreeDataProjectionToplist_20270101 != nil {
		return datadog.Marshal(&obj.TreeDataProjectionToplist_20270101)
	}

	if obj.TreeDataProjectionTreemap_20270101 != nil {
		return datadog.Marshal(&obj.TreeDataProjectionTreemap_20270101)
	}

	if obj.QueryValueProjection_20270101 != nil {
		return datadog.Marshal(&obj.QueryValueProjection_20270101)
	}

	if obj.QueryTableProjection_20270101 != nil {
		return datadog.Marshal(&obj.QueryTableProjection_20270101)
	}

	if obj.RunChartProjection_20270101 != nil {
		return datadog.Marshal(&obj.RunChartProjection_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TabularProjection) GetActualInstance() interface{} {
	if obj.ScatterplotProjection_20270101 != nil {
		return obj.ScatterplotProjection_20270101
	}

	if obj.TreeDataProjectionBarChart_20270101 != nil {
		return obj.TreeDataProjectionBarChart_20270101
	}

	if obj.TreeDataProjectionSunburst_20270101 != nil {
		return obj.TreeDataProjectionSunburst_20270101
	}

	if obj.TreeDataProjectionToplist_20270101 != nil {
		return obj.TreeDataProjectionToplist_20270101
	}

	if obj.TreeDataProjectionTreemap_20270101 != nil {
		return obj.TreeDataProjectionTreemap_20270101
	}

	if obj.QueryValueProjection_20270101 != nil {
		return obj.QueryValueProjection_20270101
	}

	if obj.QueryTableProjection_20270101 != nil {
		return obj.QueryTableProjection_20270101
	}

	if obj.RunChartProjection_20270101 != nil {
		return obj.RunChartProjection_20270101
	}

	// all schemas are nil
	return nil
}
