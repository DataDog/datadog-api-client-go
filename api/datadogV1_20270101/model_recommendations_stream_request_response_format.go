// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationsStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type RecommendationsStreamRequestResponseFormat string

// List of RecommendationsStreamRequestResponseFormat.
const (
	RECOMMENDATIONSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST RecommendationsStreamRequestResponseFormat = "event_list"
)

var allowedRecommendationsStreamRequestResponseFormatEnumValues = []RecommendationsStreamRequestResponseFormat{
	RECOMMENDATIONSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RecommendationsStreamRequestResponseFormat) GetAllowedValues() []RecommendationsStreamRequestResponseFormat {
	return allowedRecommendationsStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RecommendationsStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RecommendationsStreamRequestResponseFormat(value)
	return nil
}

// NewRecommendationsStreamRequestResponseFormatFromValue returns a pointer to a valid RecommendationsStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRecommendationsStreamRequestResponseFormatFromValue(v string) (*RecommendationsStreamRequestResponseFormat, error) {
	ev := RecommendationsStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RecommendationsStreamRequestResponseFormat: valid values are %v", v, allowedRecommendationsStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RecommendationsStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedRecommendationsStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecommendationsStreamRequestResponseFormat value.
func (v RecommendationsStreamRequestResponseFormat) Ptr() *RecommendationsStreamRequestResponseFormat {
	return &v
}
