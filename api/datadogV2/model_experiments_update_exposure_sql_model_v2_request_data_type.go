// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsUpdateExposureSQLModelV2RequestDataType Exposure SQL models resource type.
type ExperimentsUpdateExposureSQLModelV2RequestDataType string

// List of ExperimentsUpdateExposureSQLModelV2RequestDataType.
const (
	EXPERIMENTSUPDATEEXPOSURESQLMODELV2REQUESTDATATYPE_EXPOSURE_SQL_MODELS ExperimentsUpdateExposureSQLModelV2RequestDataType = "exposure-sql-models"
)

var allowedExperimentsUpdateExposureSQLModelV2RequestDataTypeEnumValues = []ExperimentsUpdateExposureSQLModelV2RequestDataType{
	EXPERIMENTSUPDATEEXPOSURESQLMODELV2REQUESTDATATYPE_EXPOSURE_SQL_MODELS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsUpdateExposureSQLModelV2RequestDataType) GetAllowedValues() []ExperimentsUpdateExposureSQLModelV2RequestDataType {
	return allowedExperimentsUpdateExposureSQLModelV2RequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsUpdateExposureSQLModelV2RequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsUpdateExposureSQLModelV2RequestDataType(value)
	return nil
}

// NewExperimentsUpdateExposureSQLModelV2RequestDataTypeFromValue returns a pointer to a valid ExperimentsUpdateExposureSQLModelV2RequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsUpdateExposureSQLModelV2RequestDataTypeFromValue(v string) (*ExperimentsUpdateExposureSQLModelV2RequestDataType, error) {
	ev := ExperimentsUpdateExposureSQLModelV2RequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsUpdateExposureSQLModelV2RequestDataType: valid values are %v", v, allowedExperimentsUpdateExposureSQLModelV2RequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsUpdateExposureSQLModelV2RequestDataType) IsValid() bool {
	for _, existing := range allowedExperimentsUpdateExposureSQLModelV2RequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsUpdateExposureSQLModelV2RequestDataType value.
func (v ExperimentsUpdateExposureSQLModelV2RequestDataType) Ptr() *ExperimentsUpdateExposureSQLModelV2RequestDataType {
	return &v
}
