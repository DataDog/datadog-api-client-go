// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WatcherResourceType Watcher resource type
type WatcherResourceType string

// List of WatcherResourceType.
const (
	WATCHERRESOURCETYPE_WATCHER WatcherResourceType = "watcher"
)

var allowedWatcherResourceTypeEnumValues = []WatcherResourceType{
	WATCHERRESOURCETYPE_WATCHER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WatcherResourceType) GetAllowedValues() []WatcherResourceType {
	return allowedWatcherResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WatcherResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WatcherResourceType(value)
	return nil
}

// NewWatcherResourceTypeFromValue returns a pointer to a valid WatcherResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWatcherResourceTypeFromValue(v string) (*WatcherResourceType, error) {
	ev := WatcherResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WatcherResourceType: valid values are %v", v, allowedWatcherResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WatcherResourceType) IsValid() bool {
	for _, existing := range allowedWatcherResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatcherResourceType value.
func (v WatcherResourceType) Ptr() *WatcherResourceType {
	return &v
}
