// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PublishedAnalysisQueryFragmentDataSource Test-only generated schema for the dashboard versioning proof of concept.
type PublishedAnalysisQueryFragmentDataSource string

// List of PublishedAnalysisQueryFragmentDataSource.
const (
	PUBLISHEDANALYSISQUERYFRAGMENTDATASOURCE_PUBLISHED_ANALYSIS PublishedAnalysisQueryFragmentDataSource = "published_analysis"
)

var allowedPublishedAnalysisQueryFragmentDataSourceEnumValues = []PublishedAnalysisQueryFragmentDataSource{
	PUBLISHEDANALYSISQUERYFRAGMENTDATASOURCE_PUBLISHED_ANALYSIS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PublishedAnalysisQueryFragmentDataSource) GetAllowedValues() []PublishedAnalysisQueryFragmentDataSource {
	return allowedPublishedAnalysisQueryFragmentDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PublishedAnalysisQueryFragmentDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PublishedAnalysisQueryFragmentDataSource(value)
	return nil
}

// NewPublishedAnalysisQueryFragmentDataSourceFromValue returns a pointer to a valid PublishedAnalysisQueryFragmentDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPublishedAnalysisQueryFragmentDataSourceFromValue(v string) (*PublishedAnalysisQueryFragmentDataSource, error) {
	ev := PublishedAnalysisQueryFragmentDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PublishedAnalysisQueryFragmentDataSource: valid values are %v", v, allowedPublishedAnalysisQueryFragmentDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PublishedAnalysisQueryFragmentDataSource) IsValid() bool {
	for _, existing := range allowedPublishedAnalysisQueryFragmentDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PublishedAnalysisQueryFragmentDataSource value.
func (v PublishedAnalysisQueryFragmentDataSource) Ptr() *PublishedAnalysisQueryFragmentDataSource {
	return &v
}
