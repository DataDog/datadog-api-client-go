// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSMetricNameFilterPreviewType The `AWSMetricNameFilterPreviewResponseData` `type`.
type AWSMetricNameFilterPreviewType string

// List of AWSMetricNameFilterPreviewType.
const (
	AWSMETRICNAMEFILTERPREVIEWTYPE_METRIC_NAME_FILTER_PREVIEW AWSMetricNameFilterPreviewType = "metric_name_filter_preview"
)

var allowedAWSMetricNameFilterPreviewTypeEnumValues = []AWSMetricNameFilterPreviewType{
	AWSMETRICNAMEFILTERPREVIEWTYPE_METRIC_NAME_FILTER_PREVIEW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSMetricNameFilterPreviewType) GetAllowedValues() []AWSMetricNameFilterPreviewType {
	return allowedAWSMetricNameFilterPreviewTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSMetricNameFilterPreviewType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSMetricNameFilterPreviewType(value)
	return nil
}

// NewAWSMetricNameFilterPreviewTypeFromValue returns a pointer to a valid AWSMetricNameFilterPreviewType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSMetricNameFilterPreviewTypeFromValue(v string) (*AWSMetricNameFilterPreviewType, error) {
	ev := AWSMetricNameFilterPreviewType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSMetricNameFilterPreviewType: valid values are %v", v, allowedAWSMetricNameFilterPreviewTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSMetricNameFilterPreviewType) IsValid() bool {
	for _, existing := range allowedAWSMetricNameFilterPreviewTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSMetricNameFilterPreviewType value.
func (v AWSMetricNameFilterPreviewType) Ptr() *AWSMetricNameFilterPreviewType {
	return &v
}
