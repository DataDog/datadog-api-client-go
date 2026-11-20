// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PublishedAnalysisDataSource Test-only generated schema for the dashboard versioning proof of concept.
type PublishedAnalysisDataSource string

// List of PublishedAnalysisDataSource.
const (
	PUBLISHEDANALYSISDATASOURCE_PUBLISHED_ANALYSIS PublishedAnalysisDataSource = "published_analysis"
)

var allowedPublishedAnalysisDataSourceEnumValues = []PublishedAnalysisDataSource{
	PUBLISHEDANALYSISDATASOURCE_PUBLISHED_ANALYSIS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PublishedAnalysisDataSource) GetAllowedValues() []PublishedAnalysisDataSource {
	return allowedPublishedAnalysisDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PublishedAnalysisDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PublishedAnalysisDataSource(value)
	return nil
}

// NewPublishedAnalysisDataSourceFromValue returns a pointer to a valid PublishedAnalysisDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPublishedAnalysisDataSourceFromValue(v string) (*PublishedAnalysisDataSource, error) {
	ev := PublishedAnalysisDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PublishedAnalysisDataSource: valid values are %v", v, allowedPublishedAnalysisDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PublishedAnalysisDataSource) IsValid() bool {
	for _, existing := range allowedPublishedAnalysisDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PublishedAnalysisDataSource value.
func (v PublishedAnalysisDataSource) Ptr() *PublishedAnalysisDataSource {
	return &v
}
