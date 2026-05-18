// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotDataProjectionQueryStorage Storage tier to query.
type ScatterplotDataProjectionQueryStorage string

// List of ScatterplotDataProjectionQueryStorage.
const (
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_LIVE             ScatterplotDataProjectionQueryStorage = "live"
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_HOT              ScatterplotDataProjectionQueryStorage = "hot"
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_HABANERO         ScatterplotDataProjectionQueryStorage = "habanero"
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_ONLINE_ARCHIVES  ScatterplotDataProjectionQueryStorage = "online_archives"
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_DRIVELINE        ScatterplotDataProjectionQueryStorage = "driveline"
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_FLEX_TIER        ScatterplotDataProjectionQueryStorage = "flex_tier"
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_CASE_INSENSITIVE ScatterplotDataProjectionQueryStorage = "case_insensitive"
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_CLOUD_PREM       ScatterplotDataProjectionQueryStorage = "cloud_prem"
)

var allowedScatterplotDataProjectionQueryStorageEnumValues = []ScatterplotDataProjectionQueryStorage{
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_LIVE,
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_HOT,
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_HABANERO,
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_ONLINE_ARCHIVES,
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_DRIVELINE,
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_FLEX_TIER,
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_CASE_INSENSITIVE,
	SCATTERPLOTDATAPROJECTIONQUERYSTORAGE_CLOUD_PREM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScatterplotDataProjectionQueryStorage) GetAllowedValues() []ScatterplotDataProjectionQueryStorage {
	return allowedScatterplotDataProjectionQueryStorageEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScatterplotDataProjectionQueryStorage) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotDataProjectionQueryStorage(value)
	return nil
}

// NewScatterplotDataProjectionQueryStorageFromValue returns a pointer to a valid ScatterplotDataProjectionQueryStorage
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScatterplotDataProjectionQueryStorageFromValue(v string) (*ScatterplotDataProjectionQueryStorage, error) {
	ev := ScatterplotDataProjectionQueryStorage(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScatterplotDataProjectionQueryStorage: valid values are %v", v, allowedScatterplotDataProjectionQueryStorageEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScatterplotDataProjectionQueryStorage) IsValid() bool {
	for _, existing := range allowedScatterplotDataProjectionQueryStorageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotDataProjectionQueryStorage value.
func (v ScatterplotDataProjectionQueryStorage) Ptr() *ScatterplotDataProjectionQueryStorage {
	return &v
}
