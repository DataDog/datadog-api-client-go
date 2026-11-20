// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisTransformation - Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisTransformation struct {
	StructuredAnalysisFilter_20270101           *StructuredAnalysisFilter
	StructuredAnalysisColumnExtraction_20270101 *StructuredAnalysisColumnExtraction
	StructuredAnalysisAggregation_20270101      *StructuredAnalysisAggregation
	StructuredAnalysisSort_20270101             *StructuredAnalysisSort
	StructuredAnalysisLimit_20270101            *StructuredAnalysisLimit
	StructuredAnalysisCalculatedField_20270101  *StructuredAnalysisCalculatedField
	StructuredAnalysisCast_20270101             *StructuredAnalysisCast
	StructuredAnalysisJoin_20270101             *StructuredAnalysisJoin
	StructuredAnalysisProjection_20270101       *StructuredAnalysisProjection
	StructuredAnalysisDestructure_20270101      *StructuredAnalysisDestructure

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StructuredAnalysisFilter_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisFilter wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisFilter_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisFilter) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisFilter_20270101: v}
}

// StructuredAnalysisColumnExtraction_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisColumnExtraction wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisColumnExtraction_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisColumnExtraction) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisColumnExtraction_20270101: v}
}

// StructuredAnalysisAggregation_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisAggregation wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisAggregation_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisAggregation) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisAggregation_20270101: v}
}

// StructuredAnalysisSort_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisSort wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisSort_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisSort) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisSort_20270101: v}
}

// StructuredAnalysisLimit_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisLimit wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisLimit_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisLimit) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisLimit_20270101: v}
}

// StructuredAnalysisCalculatedField_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisCalculatedField wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisCalculatedField_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisCalculatedField) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisCalculatedField_20270101: v}
}

// StructuredAnalysisCast_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisCast wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisCast_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisCast) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisCast_20270101: v}
}

// StructuredAnalysisJoin_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisJoin wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisJoin_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisJoin) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisJoin_20270101: v}
}

// StructuredAnalysisProjection_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisProjection wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisProjection_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisProjection) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisProjection_20270101: v}
}

// StructuredAnalysisDestructure_20270101AsStructuredAnalysisTransformation is a convenience function that returns StructuredAnalysisDestructure wrapped in StructuredAnalysisTransformation.
func StructuredAnalysisDestructure_20270101AsStructuredAnalysisTransformation(v *StructuredAnalysisDestructure) StructuredAnalysisTransformation {
	return StructuredAnalysisTransformation{StructuredAnalysisDestructure_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *StructuredAnalysisTransformation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StructuredAnalysisFilter_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisFilter_20270101)
	if err == nil {
		if obj.StructuredAnalysisFilter_20270101 != nil && obj.StructuredAnalysisFilter_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisFilter_20270101, _ := datadog.Marshal(obj.StructuredAnalysisFilter_20270101)
			if string(jsonStructuredAnalysisFilter_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisFilter_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisFilter_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisFilter_20270101 = nil
	}

	// try to unmarshal data into StructuredAnalysisColumnExtraction_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisColumnExtraction_20270101)
	if err == nil {
		if obj.StructuredAnalysisColumnExtraction_20270101 != nil && obj.StructuredAnalysisColumnExtraction_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisColumnExtraction_20270101, _ := datadog.Marshal(obj.StructuredAnalysisColumnExtraction_20270101)
			if string(jsonStructuredAnalysisColumnExtraction_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisColumnExtraction_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisColumnExtraction_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisColumnExtraction_20270101 = nil
	}

	// try to unmarshal data into StructuredAnalysisAggregation_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisAggregation_20270101)
	if err == nil {
		if obj.StructuredAnalysisAggregation_20270101 != nil && obj.StructuredAnalysisAggregation_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisAggregation_20270101, _ := datadog.Marshal(obj.StructuredAnalysisAggregation_20270101)
			if string(jsonStructuredAnalysisAggregation_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisAggregation_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisAggregation_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisAggregation_20270101 = nil
	}

	// try to unmarshal data into StructuredAnalysisSort_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisSort_20270101)
	if err == nil {
		if obj.StructuredAnalysisSort_20270101 != nil && obj.StructuredAnalysisSort_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisSort_20270101, _ := datadog.Marshal(obj.StructuredAnalysisSort_20270101)
			if string(jsonStructuredAnalysisSort_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisSort_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisSort_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisSort_20270101 = nil
	}

	// try to unmarshal data into StructuredAnalysisLimit_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisLimit_20270101)
	if err == nil {
		if obj.StructuredAnalysisLimit_20270101 != nil && obj.StructuredAnalysisLimit_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisLimit_20270101, _ := datadog.Marshal(obj.StructuredAnalysisLimit_20270101)
			if string(jsonStructuredAnalysisLimit_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisLimit_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisLimit_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisLimit_20270101 = nil
	}

	// try to unmarshal data into StructuredAnalysisCalculatedField_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisCalculatedField_20270101)
	if err == nil {
		if obj.StructuredAnalysisCalculatedField_20270101 != nil && obj.StructuredAnalysisCalculatedField_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisCalculatedField_20270101, _ := datadog.Marshal(obj.StructuredAnalysisCalculatedField_20270101)
			if string(jsonStructuredAnalysisCalculatedField_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisCalculatedField_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisCalculatedField_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisCalculatedField_20270101 = nil
	}

	// try to unmarshal data into StructuredAnalysisCast_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisCast_20270101)
	if err == nil {
		if obj.StructuredAnalysisCast_20270101 != nil && obj.StructuredAnalysisCast_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisCast_20270101, _ := datadog.Marshal(obj.StructuredAnalysisCast_20270101)
			if string(jsonStructuredAnalysisCast_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisCast_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisCast_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisCast_20270101 = nil
	}

	// try to unmarshal data into StructuredAnalysisJoin_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisJoin_20270101)
	if err == nil {
		if obj.StructuredAnalysisJoin_20270101 != nil && obj.StructuredAnalysisJoin_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisJoin_20270101, _ := datadog.Marshal(obj.StructuredAnalysisJoin_20270101)
			if string(jsonStructuredAnalysisJoin_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisJoin_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisJoin_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisJoin_20270101 = nil
	}

	// try to unmarshal data into StructuredAnalysisProjection_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisProjection_20270101)
	if err == nil {
		if obj.StructuredAnalysisProjection_20270101 != nil && obj.StructuredAnalysisProjection_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisProjection_20270101, _ := datadog.Marshal(obj.StructuredAnalysisProjection_20270101)
			if string(jsonStructuredAnalysisProjection_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisProjection_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisProjection_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisProjection_20270101 = nil
	}

	// try to unmarshal data into StructuredAnalysisDestructure_20270101
	err = datadog.Unmarshal(data, &obj.StructuredAnalysisDestructure_20270101)
	if err == nil {
		if obj.StructuredAnalysisDestructure_20270101 != nil && obj.StructuredAnalysisDestructure_20270101.UnparsedObject == nil {
			jsonStructuredAnalysisDestructure_20270101, _ := datadog.Marshal(obj.StructuredAnalysisDestructure_20270101)
			if string(jsonStructuredAnalysisDestructure_20270101) == "{}" { // empty struct
				obj.StructuredAnalysisDestructure_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StructuredAnalysisDestructure_20270101 = nil
		}
	} else {
		obj.StructuredAnalysisDestructure_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.StructuredAnalysisFilter_20270101 = nil
		obj.StructuredAnalysisColumnExtraction_20270101 = nil
		obj.StructuredAnalysisAggregation_20270101 = nil
		obj.StructuredAnalysisSort_20270101 = nil
		obj.StructuredAnalysisLimit_20270101 = nil
		obj.StructuredAnalysisCalculatedField_20270101 = nil
		obj.StructuredAnalysisCast_20270101 = nil
		obj.StructuredAnalysisJoin_20270101 = nil
		obj.StructuredAnalysisProjection_20270101 = nil
		obj.StructuredAnalysisDestructure_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj StructuredAnalysisTransformation) MarshalJSON() ([]byte, error) {
	if obj.StructuredAnalysisFilter_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisFilter_20270101)
	}

	if obj.StructuredAnalysisColumnExtraction_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisColumnExtraction_20270101)
	}

	if obj.StructuredAnalysisAggregation_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisAggregation_20270101)
	}

	if obj.StructuredAnalysisSort_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisSort_20270101)
	}

	if obj.StructuredAnalysisLimit_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisLimit_20270101)
	}

	if obj.StructuredAnalysisCalculatedField_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisCalculatedField_20270101)
	}

	if obj.StructuredAnalysisCast_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisCast_20270101)
	}

	if obj.StructuredAnalysisJoin_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisJoin_20270101)
	}

	if obj.StructuredAnalysisProjection_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisProjection_20270101)
	}

	if obj.StructuredAnalysisDestructure_20270101 != nil {
		return datadog.Marshal(&obj.StructuredAnalysisDestructure_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *StructuredAnalysisTransformation) GetActualInstance() interface{} {
	if obj.StructuredAnalysisFilter_20270101 != nil {
		return obj.StructuredAnalysisFilter_20270101
	}

	if obj.StructuredAnalysisColumnExtraction_20270101 != nil {
		return obj.StructuredAnalysisColumnExtraction_20270101
	}

	if obj.StructuredAnalysisAggregation_20270101 != nil {
		return obj.StructuredAnalysisAggregation_20270101
	}

	if obj.StructuredAnalysisSort_20270101 != nil {
		return obj.StructuredAnalysisSort_20270101
	}

	if obj.StructuredAnalysisLimit_20270101 != nil {
		return obj.StructuredAnalysisLimit_20270101
	}

	if obj.StructuredAnalysisCalculatedField_20270101 != nil {
		return obj.StructuredAnalysisCalculatedField_20270101
	}

	if obj.StructuredAnalysisCast_20270101 != nil {
		return obj.StructuredAnalysisCast_20270101
	}

	if obj.StructuredAnalysisJoin_20270101 != nil {
		return obj.StructuredAnalysisJoin_20270101
	}

	if obj.StructuredAnalysisProjection_20270101 != nil {
		return obj.StructuredAnalysisProjection_20270101
	}

	if obj.StructuredAnalysisDestructure_20270101 != nil {
		return obj.StructuredAnalysisDestructure_20270101
	}

	// all schemas are nil
	return nil
}
