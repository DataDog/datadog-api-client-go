// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SqlAnalysisQueryType Test-only generated schema for the dashboard versioning proof of concept.
type SqlAnalysisQueryType string

// List of SqlAnalysisQueryType.
const (
	SQLANALYSISQUERYTYPE_SQL_ANALYSIS SqlAnalysisQueryType = "sql_analysis"
)

var allowedSqlAnalysisQueryTypeEnumValues = []SqlAnalysisQueryType{
	SQLANALYSISQUERYTYPE_SQL_ANALYSIS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SqlAnalysisQueryType) GetAllowedValues() []SqlAnalysisQueryType {
	return allowedSqlAnalysisQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SqlAnalysisQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SqlAnalysisQueryType(value)
	return nil
}

// NewSqlAnalysisQueryTypeFromValue returns a pointer to a valid SqlAnalysisQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSqlAnalysisQueryTypeFromValue(v string) (*SqlAnalysisQueryType, error) {
	ev := SqlAnalysisQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SqlAnalysisQueryType: valid values are %v", v, allowedSqlAnalysisQueryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SqlAnalysisQueryType) IsValid() bool {
	for _, existing := range allowedSqlAnalysisQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SqlAnalysisQueryType value.
func (v SqlAnalysisQueryType) Ptr() *SqlAnalysisQueryType {
	return &v
}
