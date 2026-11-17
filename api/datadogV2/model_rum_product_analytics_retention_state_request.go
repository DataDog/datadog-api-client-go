// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMProductAnalyticsRetentionStateRequest Retention policy to set for Product Analytics data derived from RUM events. `PREVIEW` is the initial state of a new application and is reported as `NONE` in responses. Setting `MAX` requires the `product_analytics_apps_write` permission.
type RUMProductAnalyticsRetentionStateRequest string

// List of RUMProductAnalyticsRetentionStateRequest.
const (
	RUMPRODUCTANALYTICSRETENTIONSTATEREQUEST_MAX     RUMProductAnalyticsRetentionStateRequest = "MAX"
	RUMPRODUCTANALYTICSRETENTIONSTATEREQUEST_NONE    RUMProductAnalyticsRetentionStateRequest = "NONE"
	RUMPRODUCTANALYTICSRETENTIONSTATEREQUEST_PREVIEW RUMProductAnalyticsRetentionStateRequest = "PREVIEW"
)

var allowedRUMProductAnalyticsRetentionStateRequestEnumValues = []RUMProductAnalyticsRetentionStateRequest{
	RUMPRODUCTANALYTICSRETENTIONSTATEREQUEST_MAX,
	RUMPRODUCTANALYTICSRETENTIONSTATEREQUEST_NONE,
	RUMPRODUCTANALYTICSRETENTIONSTATEREQUEST_PREVIEW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMProductAnalyticsRetentionStateRequest) GetAllowedValues() []RUMProductAnalyticsRetentionStateRequest {
	return allowedRUMProductAnalyticsRetentionStateRequestEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMProductAnalyticsRetentionStateRequest) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMProductAnalyticsRetentionStateRequest(value)
	return nil
}

// NewRUMProductAnalyticsRetentionStateRequestFromValue returns a pointer to a valid RUMProductAnalyticsRetentionStateRequest
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMProductAnalyticsRetentionStateRequestFromValue(v string) (*RUMProductAnalyticsRetentionStateRequest, error) {
	ev := RUMProductAnalyticsRetentionStateRequest(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMProductAnalyticsRetentionStateRequest: valid values are %v", v, allowedRUMProductAnalyticsRetentionStateRequestEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMProductAnalyticsRetentionStateRequest) IsValid() bool {
	for _, existing := range allowedRUMProductAnalyticsRetentionStateRequestEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMProductAnalyticsRetentionStateRequest value.
func (v RUMProductAnalyticsRetentionStateRequest) Ptr() *RUMProductAnalyticsRetentionStateRequest {
	return &v
}
