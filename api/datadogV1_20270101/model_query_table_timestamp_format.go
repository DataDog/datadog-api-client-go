// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableTimestampFormat Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableTimestampFormat string

// List of QueryTableTimestampFormat.
const (
	QUERYTABLETIMESTAMPFORMAT_MMM_DD_YYYY_HH_MM      QueryTableTimestampFormat = "MMM dd yyyy HH:mm"
	QUERYTABLETIMESTAMPFORMAT_MMM_DD_YYYY_HH_MM_SS   QueryTableTimestampFormat = "MMM dd yyyy HH:mm:ss"
	QUERYTABLETIMESTAMPFORMAT_MMM_DD_YYYY_H_MM_A     QueryTableTimestampFormat = "MMM dd yyyy h:mm a"
	QUERYTABLETIMESTAMPFORMAT_MMM_DD_YYYY_H_MM_SS_A  QueryTableTimestampFormat = "MMM dd yyyy h:mm:ss a"
	QUERYTABLETIMESTAMPFORMAT_MMM_D_HH_MM            QueryTableTimestampFormat = "MMM d HH:mm"
	QUERYTABLETIMESTAMPFORMAT_MMM_D_HH_MM_SS         QueryTableTimestampFormat = "MMM d HH:mm:ss"
	QUERYTABLETIMESTAMPFORMAT_MMM_D_H_MM_A           QueryTableTimestampFormat = "MMM d h:mm a"
	QUERYTABLETIMESTAMPFORMAT_MMM_D_H_MM_SS_A        QueryTableTimestampFormat = "MMM d h:mm:ss a"
	QUERYTABLETIMESTAMPFORMAT_MMM_D__YYYY_HH_MM      QueryTableTimestampFormat = "MMM d, yyyy HH:mm"
	QUERYTABLETIMESTAMPFORMAT_MMM_D__YYYY_HH_MM_SS   QueryTableTimestampFormat = "MMM d, yyyy HH:mm:ss"
	QUERYTABLETIMESTAMPFORMAT_MMM_D__YYYY_H_MM_A     QueryTableTimestampFormat = "MMM d, yyyy h:mm a"
	QUERYTABLETIMESTAMPFORMAT_MMM_D__YYYY_H_MM_SS_A  QueryTableTimestampFormat = "MMM d, yyyy h:mm:ss a"
	QUERYTABLETIMESTAMPFORMAT_MM_DD_YYYY_HH_MM       QueryTableTimestampFormat = "MM/dd/yyyy HH:mm"
	QUERYTABLETIMESTAMPFORMAT_MM_DD_YYYY_HH_MM_SS    QueryTableTimestampFormat = "MM/dd/yyyy HH:mm:ss"
	QUERYTABLETIMESTAMPFORMAT_MM_DD_YYYY_H_MM_A      QueryTableTimestampFormat = "MM/dd/yyyy h:mm a"
	QUERYTABLETIMESTAMPFORMAT_MM_DD_YYYY_H_MM_SS_A   QueryTableTimestampFormat = "MM/dd/yyyy h:mm:ss a"
	QUERYTABLETIMESTAMPFORMAT_DD_MM_YYYY_HH_MM       QueryTableTimestampFormat = "dd/MM/yyyy HH:mm"
	QUERYTABLETIMESTAMPFORMAT_DD_MM_YYYY_HH_MM_SS    QueryTableTimestampFormat = "dd/MM/yyyy HH:mm:ss"
	QUERYTABLETIMESTAMPFORMAT_DD_MM_YYYY_H_MM_A      QueryTableTimestampFormat = "dd/MM/yyyy h:mm a"
	QUERYTABLETIMESTAMPFORMAT_DD_MM_YYYY_H_MM_SS_A   QueryTableTimestampFormat = "dd/MM/yyyy h:mm:ss a"
	QUERYTABLETIMESTAMPFORMAT_MMMM_D__YYYY_HH_MM     QueryTableTimestampFormat = "MMMM d, yyyy HH:mm"
	QUERYTABLETIMESTAMPFORMAT_MMMM_D__YYYY_HH_MM_SS  QueryTableTimestampFormat = "MMMM d, yyyy HH:mm:ss"
	QUERYTABLETIMESTAMPFORMAT_MMMM_D__YYYY_H_MM_A    QueryTableTimestampFormat = "MMMM d, yyyy h:mm a"
	QUERYTABLETIMESTAMPFORMAT_MMMM_D__YYYY_H_MM_SS_A QueryTableTimestampFormat = "MMMM d, yyyy h:mm:ss a"
)

var allowedQueryTableTimestampFormatEnumValues = []QueryTableTimestampFormat{
	QUERYTABLETIMESTAMPFORMAT_MMM_DD_YYYY_HH_MM,
	QUERYTABLETIMESTAMPFORMAT_MMM_DD_YYYY_HH_MM_SS,
	QUERYTABLETIMESTAMPFORMAT_MMM_DD_YYYY_H_MM_A,
	QUERYTABLETIMESTAMPFORMAT_MMM_DD_YYYY_H_MM_SS_A,
	QUERYTABLETIMESTAMPFORMAT_MMM_D_HH_MM,
	QUERYTABLETIMESTAMPFORMAT_MMM_D_HH_MM_SS,
	QUERYTABLETIMESTAMPFORMAT_MMM_D_H_MM_A,
	QUERYTABLETIMESTAMPFORMAT_MMM_D_H_MM_SS_A,
	QUERYTABLETIMESTAMPFORMAT_MMM_D__YYYY_HH_MM,
	QUERYTABLETIMESTAMPFORMAT_MMM_D__YYYY_HH_MM_SS,
	QUERYTABLETIMESTAMPFORMAT_MMM_D__YYYY_H_MM_A,
	QUERYTABLETIMESTAMPFORMAT_MMM_D__YYYY_H_MM_SS_A,
	QUERYTABLETIMESTAMPFORMAT_MM_DD_YYYY_HH_MM,
	QUERYTABLETIMESTAMPFORMAT_MM_DD_YYYY_HH_MM_SS,
	QUERYTABLETIMESTAMPFORMAT_MM_DD_YYYY_H_MM_A,
	QUERYTABLETIMESTAMPFORMAT_MM_DD_YYYY_H_MM_SS_A,
	QUERYTABLETIMESTAMPFORMAT_DD_MM_YYYY_HH_MM,
	QUERYTABLETIMESTAMPFORMAT_DD_MM_YYYY_HH_MM_SS,
	QUERYTABLETIMESTAMPFORMAT_DD_MM_YYYY_H_MM_A,
	QUERYTABLETIMESTAMPFORMAT_DD_MM_YYYY_H_MM_SS_A,
	QUERYTABLETIMESTAMPFORMAT_MMMM_D__YYYY_HH_MM,
	QUERYTABLETIMESTAMPFORMAT_MMMM_D__YYYY_HH_MM_SS,
	QUERYTABLETIMESTAMPFORMAT_MMMM_D__YYYY_H_MM_A,
	QUERYTABLETIMESTAMPFORMAT_MMMM_D__YYYY_H_MM_SS_A,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableTimestampFormat) GetAllowedValues() []QueryTableTimestampFormat {
	return allowedQueryTableTimestampFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableTimestampFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableTimestampFormat(value)
	return nil
}

// NewQueryTableTimestampFormatFromValue returns a pointer to a valid QueryTableTimestampFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableTimestampFormatFromValue(v string) (*QueryTableTimestampFormat, error) {
	ev := QueryTableTimestampFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableTimestampFormat: valid values are %v", v, allowedQueryTableTimestampFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableTimestampFormat) IsValid() bool {
	for _, existing := range allowedQueryTableTimestampFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableTimestampFormat value.
func (v QueryTableTimestampFormat) Ptr() *QueryTableTimestampFormat {
	return &v
}
