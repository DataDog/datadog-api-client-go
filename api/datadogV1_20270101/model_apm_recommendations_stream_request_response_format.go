// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmRecommendationsStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type ApmRecommendationsStreamRequestResponseFormat string

// List of ApmRecommendationsStreamRequestResponseFormat.
const (
	APMRECOMMENDATIONSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST ApmRecommendationsStreamRequestResponseFormat = "event_list"
)

var allowedApmRecommendationsStreamRequestResponseFormatEnumValues = []ApmRecommendationsStreamRequestResponseFormat{
	APMRECOMMENDATIONSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmRecommendationsStreamRequestResponseFormat) GetAllowedValues() []ApmRecommendationsStreamRequestResponseFormat {
	return allowedApmRecommendationsStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmRecommendationsStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmRecommendationsStreamRequestResponseFormat(value)
	return nil
}

// NewApmRecommendationsStreamRequestResponseFormatFromValue returns a pointer to a valid ApmRecommendationsStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmRecommendationsStreamRequestResponseFormatFromValue(v string) (*ApmRecommendationsStreamRequestResponseFormat, error) {
	ev := ApmRecommendationsStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmRecommendationsStreamRequestResponseFormat: valid values are %v", v, allowedApmRecommendationsStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmRecommendationsStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedApmRecommendationsStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmRecommendationsStreamRequestResponseFormat value.
func (v ApmRecommendationsStreamRequestResponseFormat) Ptr() *ApmRecommendationsStreamRequestResponseFormat {
	return &v
}
