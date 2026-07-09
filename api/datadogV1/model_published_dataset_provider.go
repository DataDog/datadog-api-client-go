// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PublishedDatasetProvider Product page that published the dataset queried by a `DatasetListQuery`. `ddsql_query` is the only provider currently supported for host map widgets.
type PublishedDatasetProvider string

// List of PublishedDatasetProvider.
const (
	PUBLISHEDDATASETPROVIDER_DDSQL_QUERY PublishedDatasetProvider = "ddsql_query"
)

var allowedPublishedDatasetProviderEnumValues = []PublishedDatasetProvider{
	PUBLISHEDDATASETPROVIDER_DDSQL_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PublishedDatasetProvider) GetAllowedValues() []PublishedDatasetProvider {
	return allowedPublishedDatasetProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PublishedDatasetProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PublishedDatasetProvider(value)
	return nil
}

// NewPublishedDatasetProviderFromValue returns a pointer to a valid PublishedDatasetProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPublishedDatasetProviderFromValue(v string) (*PublishedDatasetProvider, error) {
	ev := PublishedDatasetProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PublishedDatasetProvider: valid values are %v", v, allowedPublishedDatasetProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PublishedDatasetProvider) IsValid() bool {
	for _, existing := range allowedPublishedDatasetProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PublishedDatasetProvider value.
func (v PublishedDatasetProvider) Ptr() *PublishedDatasetProvider {
	return &v
}
