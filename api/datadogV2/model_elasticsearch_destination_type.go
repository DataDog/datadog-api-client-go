// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// ElasticsearchDestinationType The Elasticsearch destination type.
type ElasticsearchDestinationType string

// List of ElasticsearchDestinationType.
const (
	ELASTICSEARCHDESTINATIONTYPE_ELASTICSEARCH ElasticsearchDestinationType = "elasticsearch"
)

var allowedElasticsearchDestinationTypeEnumValues = []ElasticsearchDestinationType{
	ELASTICSEARCHDESTINATIONTYPE_ELASTICSEARCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ElasticsearchDestinationType) GetAllowedValues() []ElasticsearchDestinationType {
	return allowedElasticsearchDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticsearchDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticsearchDestinationType(value)
	return nil
}

// NewElasticsearchDestinationTypeFromValue returns a pointer to a valid ElasticsearchDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticsearchDestinationTypeFromValue(v string) (*ElasticsearchDestinationType, error) {
	ev := ElasticsearchDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticsearchDestinationType: valid values are %v", v, allowedElasticsearchDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticsearchDestinationType) IsValid() bool {
	for _, existing := range allowedElasticsearchDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticsearchDestinationType value.
func (v ElasticsearchDestinationType) Ptr() *ElasticsearchDestinationType {
	return &v
}
