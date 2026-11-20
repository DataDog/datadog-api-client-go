// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabaseRecommendationsStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type DatabaseRecommendationsStreamRequestResponseFormat string

// List of DatabaseRecommendationsStreamRequestResponseFormat.
const (
	DATABASERECOMMENDATIONSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST DatabaseRecommendationsStreamRequestResponseFormat = "event_list"
)

var allowedDatabaseRecommendationsStreamRequestResponseFormatEnumValues = []DatabaseRecommendationsStreamRequestResponseFormat{
	DATABASERECOMMENDATIONSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatabaseRecommendationsStreamRequestResponseFormat) GetAllowedValues() []DatabaseRecommendationsStreamRequestResponseFormat {
	return allowedDatabaseRecommendationsStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatabaseRecommendationsStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatabaseRecommendationsStreamRequestResponseFormat(value)
	return nil
}

// NewDatabaseRecommendationsStreamRequestResponseFormatFromValue returns a pointer to a valid DatabaseRecommendationsStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatabaseRecommendationsStreamRequestResponseFormatFromValue(v string) (*DatabaseRecommendationsStreamRequestResponseFormat, error) {
	ev := DatabaseRecommendationsStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatabaseRecommendationsStreamRequestResponseFormat: valid values are %v", v, allowedDatabaseRecommendationsStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatabaseRecommendationsStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedDatabaseRecommendationsStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabaseRecommendationsStreamRequestResponseFormat value.
func (v DatabaseRecommendationsStreamRequestResponseFormat) Ptr() *DatabaseRecommendationsStreamRequestResponseFormat {
	return &v
}
