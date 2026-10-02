// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType Data type of a column in the SQL model.
type ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType string

// List of ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType.
const (
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_STRING    ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType = "STRING"
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_INTEGER   ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType = "INTEGER"
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_FLOAT     ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType = "FLOAT"
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_BOOLEAN   ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType = "BOOLEAN"
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_DATE      ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType = "DATE"
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_TIMESTAMP ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType = "TIMESTAMP"
)

var allowedExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnTypeEnumValues = []ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType{
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_STRING,
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_INTEGER,
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_FLOAT,
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_BOOLEAN,
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_DATE,
	EXPERIMENTSCREATEEXPOSURESQLMODELV2REQUESTDATAATTRIBUTESITEMSCOLUMNTYPE_TIMESTAMP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType) GetAllowedValues() []ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType {
	return allowedExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType(value)
	return nil
}

// NewExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnTypeFromValue returns a pointer to a valid ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnTypeFromValue(v string) (*ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType, error) {
	ev := ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType: valid values are %v", v, allowedExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType) IsValid() bool {
	for _, existing := range allowedExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType value.
func (v ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType) Ptr() *ExperimentsCreateExposureSQLModelV2RequestDataAttributesItemsColumnType {
	return &v
}
