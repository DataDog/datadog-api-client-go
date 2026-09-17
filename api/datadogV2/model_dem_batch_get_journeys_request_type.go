// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemBatchGetJourneysRequestType The resource type for a request to retrieve DEM journeys by test suite IDs.
type DemBatchGetJourneysRequestType string

// List of DemBatchGetJourneysRequestType.
const (
	DEMBATCHGETJOURNEYSREQUESTTYPE_BATCH_GET_JOURNEYS_BY_TEST_SUITE_IDS_REQUEST DemBatchGetJourneysRequestType = "batch_get_journeys_by_test_suite_ids_request"
)

var allowedDemBatchGetJourneysRequestTypeEnumValues = []DemBatchGetJourneysRequestType{
	DEMBATCHGETJOURNEYSREQUESTTYPE_BATCH_GET_JOURNEYS_BY_TEST_SUITE_IDS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DemBatchGetJourneysRequestType) GetAllowedValues() []DemBatchGetJourneysRequestType {
	return allowedDemBatchGetJourneysRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DemBatchGetJourneysRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DemBatchGetJourneysRequestType(value)
	return nil
}

// NewDemBatchGetJourneysRequestTypeFromValue returns a pointer to a valid DemBatchGetJourneysRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDemBatchGetJourneysRequestTypeFromValue(v string) (*DemBatchGetJourneysRequestType, error) {
	ev := DemBatchGetJourneysRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DemBatchGetJourneysRequestType: valid values are %v", v, allowedDemBatchGetJourneysRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DemBatchGetJourneysRequestType) IsValid() bool {
	for _, existing := range allowedDemBatchGetJourneysRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DemBatchGetJourneysRequestType value.
func (v DemBatchGetJourneysRequestType) Ptr() *DemBatchGetJourneysRequestType {
	return &v
}
