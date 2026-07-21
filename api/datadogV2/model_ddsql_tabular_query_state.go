// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DdsqlTabularQueryState Lifecycle state of a DDSQL tabular query response.
// `running` means the query is still executing and the client should poll
// the fetch endpoint with the returned `query_id`. `completed` means the
// result set is inlined in `columns` and no further polling is required.
type DdsqlTabularQueryState string

// List of DdsqlTabularQueryState.
const (
	DDSQLTABULARQUERYSTATE_RUNNING   DdsqlTabularQueryState = "running"
	DDSQLTABULARQUERYSTATE_COMPLETED DdsqlTabularQueryState = "completed"
)

var allowedDdsqlTabularQueryStateEnumValues = []DdsqlTabularQueryState{
	DDSQLTABULARQUERYSTATE_RUNNING,
	DDSQLTABULARQUERYSTATE_COMPLETED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DdsqlTabularQueryState) GetAllowedValues() []DdsqlTabularQueryState {
	return allowedDdsqlTabularQueryStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DdsqlTabularQueryState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DdsqlTabularQueryState(value)
	return nil
}

// NewDdsqlTabularQueryStateFromValue returns a pointer to a valid DdsqlTabularQueryState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDdsqlTabularQueryStateFromValue(v string) (*DdsqlTabularQueryState, error) {
	ev := DdsqlTabularQueryState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DdsqlTabularQueryState: valid values are %v", v, allowedDdsqlTabularQueryStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DdsqlTabularQueryState) IsValid() bool {
	for _, existing := range allowedDdsqlTabularQueryStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DdsqlTabularQueryState value.
func (v DdsqlTabularQueryState) Ptr() *DdsqlTabularQueryState {
	return &v
}
