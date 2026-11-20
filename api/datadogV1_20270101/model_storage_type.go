// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StorageType Test-only generated schema for the dashboard versioning proof of concept.
type StorageType string

// List of StorageType.
const (
	STORAGETYPE_LIVE             StorageType = "live"
	STORAGETYPE_HOT              StorageType = "hot"
	STORAGETYPE_ONLINE_ARCHIVES  StorageType = "online_archives"
	STORAGETYPE_DRIVELINE        StorageType = "driveline"
	STORAGETYPE_FLEX_TIER        StorageType = "flex_tier"
	STORAGETYPE_CASE_INSENSITIVE StorageType = "case_insensitive"
	STORAGETYPE_CLOUD_PREM       StorageType = "cloud_prem"
)

var allowedStorageTypeEnumValues = []StorageType{
	STORAGETYPE_LIVE,
	STORAGETYPE_HOT,
	STORAGETYPE_ONLINE_ARCHIVES,
	STORAGETYPE_DRIVELINE,
	STORAGETYPE_FLEX_TIER,
	STORAGETYPE_CASE_INSENSITIVE,
	STORAGETYPE_CLOUD_PREM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StorageType) GetAllowedValues() []StorageType {
	return allowedStorageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StorageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StorageType(value)
	return nil
}

// NewStorageTypeFromValue returns a pointer to a valid StorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStorageTypeFromValue(v string) (*StorageType, error) {
	ev := StorageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StorageType: valid values are %v", v, allowedStorageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StorageType) IsValid() bool {
	for _, existing := range allowedStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageType value.
func (v StorageType) Ptr() *StorageType {
	return &v
}
