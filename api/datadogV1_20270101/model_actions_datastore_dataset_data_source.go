// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionsDatastoreDatasetDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ActionsDatastoreDatasetDataSource string

// List of ActionsDatastoreDatasetDataSource.
const (
	ACTIONSDATASTOREDATASETDATASOURCE_ACTIONS_DATASTORE ActionsDatastoreDatasetDataSource = "actions_datastore"
)

var allowedActionsDatastoreDatasetDataSourceEnumValues = []ActionsDatastoreDatasetDataSource{
	ACTIONSDATASTOREDATASETDATASOURCE_ACTIONS_DATASTORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ActionsDatastoreDatasetDataSource) GetAllowedValues() []ActionsDatastoreDatasetDataSource {
	return allowedActionsDatastoreDatasetDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ActionsDatastoreDatasetDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ActionsDatastoreDatasetDataSource(value)
	return nil
}

// NewActionsDatastoreDatasetDataSourceFromValue returns a pointer to a valid ActionsDatastoreDatasetDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewActionsDatastoreDatasetDataSourceFromValue(v string) (*ActionsDatastoreDatasetDataSource, error) {
	ev := ActionsDatastoreDatasetDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ActionsDatastoreDatasetDataSource: valid values are %v", v, allowedActionsDatastoreDatasetDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ActionsDatastoreDatasetDataSource) IsValid() bool {
	for _, existing := range allowedActionsDatastoreDatasetDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionsDatastoreDatasetDataSource value.
func (v ActionsDatastoreDatasetDataSource) Ptr() *ActionsDatastoreDatasetDataSource {
	return &v
}
