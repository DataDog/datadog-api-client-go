// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AdvancedQueryDataType Test-only generated schema for the dashboard versioning proof of concept.
type AdvancedQueryDataType string

// List of AdvancedQueryDataType.
const (
	ADVANCEDQUERYDATATYPE_STRING             AdvancedQueryDataType = "string"
	ADVANCEDQUERYDATATYPE_GROUP              AdvancedQueryDataType = "group"
	ADVANCEDQUERYDATATYPE_INET               AdvancedQueryDataType = "inet"
	ADVANCEDQUERYDATATYPE_INT64              AdvancedQueryDataType = "int64"
	ADVANCEDQUERYDATATYPE_FLOAT64            AdvancedQueryDataType = "float64"
	ADVANCEDQUERYDATATYPE_TIMESTAMP          AdvancedQueryDataType = "timestamp"
	ADVANCEDQUERYDATATYPE_DATE               AdvancedQueryDataType = "date"
	ADVANCEDQUERYDATATYPE_INTERVAL           AdvancedQueryDataType = "interval"
	ADVANCEDQUERYDATATYPE_BOOL               AdvancedQueryDataType = "bool"
	ADVANCEDQUERYDATATYPE_JSON               AdvancedQueryDataType = "json"
	ADVANCEDQUERYDATATYPE_HSTORE             AdvancedQueryDataType = "hstore"
	ADVANCEDQUERYDATATYPE_ARRAYLTSTRINGGT    AdvancedQueryDataType = "array<string>"
	ADVANCEDQUERYDATATYPE_ARRAYLTINT64GT     AdvancedQueryDataType = "array<int64>"
	ADVANCEDQUERYDATATYPE_ARRAYLTFLOAT64GT   AdvancedQueryDataType = "array<float64>"
	ADVANCEDQUERYDATATYPE_ARRAYLTTIMESTAMPGT AdvancedQueryDataType = "array<timestamp>"
	ADVANCEDQUERYDATATYPE_ARRAYLTINTERVALGT  AdvancedQueryDataType = "array<interval>"
	ADVANCEDQUERYDATATYPE_ARRAYLTBOOLGT      AdvancedQueryDataType = "array<bool>"
	ADVANCEDQUERYDATATYPE_ARRAYLTJSONGT      AdvancedQueryDataType = "array<json>"
)

var allowedAdvancedQueryDataTypeEnumValues = []AdvancedQueryDataType{
	ADVANCEDQUERYDATATYPE_STRING,
	ADVANCEDQUERYDATATYPE_GROUP,
	ADVANCEDQUERYDATATYPE_INET,
	ADVANCEDQUERYDATATYPE_INT64,
	ADVANCEDQUERYDATATYPE_FLOAT64,
	ADVANCEDQUERYDATATYPE_TIMESTAMP,
	ADVANCEDQUERYDATATYPE_DATE,
	ADVANCEDQUERYDATATYPE_INTERVAL,
	ADVANCEDQUERYDATATYPE_BOOL,
	ADVANCEDQUERYDATATYPE_JSON,
	ADVANCEDQUERYDATATYPE_HSTORE,
	ADVANCEDQUERYDATATYPE_ARRAYLTSTRINGGT,
	ADVANCEDQUERYDATATYPE_ARRAYLTINT64GT,
	ADVANCEDQUERYDATATYPE_ARRAYLTFLOAT64GT,
	ADVANCEDQUERYDATATYPE_ARRAYLTTIMESTAMPGT,
	ADVANCEDQUERYDATATYPE_ARRAYLTINTERVALGT,
	ADVANCEDQUERYDATATYPE_ARRAYLTBOOLGT,
	ADVANCEDQUERYDATATYPE_ARRAYLTJSONGT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AdvancedQueryDataType) GetAllowedValues() []AdvancedQueryDataType {
	return allowedAdvancedQueryDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AdvancedQueryDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AdvancedQueryDataType(value)
	return nil
}

// NewAdvancedQueryDataTypeFromValue returns a pointer to a valid AdvancedQueryDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAdvancedQueryDataTypeFromValue(v string) (*AdvancedQueryDataType, error) {
	ev := AdvancedQueryDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AdvancedQueryDataType: valid values are %v", v, allowedAdvancedQueryDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AdvancedQueryDataType) IsValid() bool {
	for _, existing := range allowedAdvancedQueryDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdvancedQueryDataType value.
func (v AdvancedQueryDataType) Ptr() *AdvancedQueryDataType {
	return &v
}
