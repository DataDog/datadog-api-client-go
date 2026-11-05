// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelMetadataSourceMatchedBy Identifier strategy used to match the gateway catalog entry. Date-suffix and shared canonical-model fallback matches produce partial metadata.
type AiGatewayModelMetadataSourceMatchedBy string

// List of AiGatewayModelMetadataSourceMatchedBy.
const (
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_OVERRIDE            AiGatewayModelMetadataSourceMatchedBy = "override"
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_WIRE_ID             AiGatewayModelMetadataSourceMatchedBy = "wire_id"
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_CANONICAL_ID        AiGatewayModelMetadataSourceMatchedBy = "canonical_id"
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_WIRE_ID_DATE_SUFFIX AiGatewayModelMetadataSourceMatchedBy = "wire_id_date_suffix"
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_CANONICAL_MODEL     AiGatewayModelMetadataSourceMatchedBy = "canonical_model"
)

var allowedAiGatewayModelMetadataSourceMatchedByEnumValues = []AiGatewayModelMetadataSourceMatchedBy{
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_OVERRIDE,
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_WIRE_ID,
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_CANONICAL_ID,
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_WIRE_ID_DATE_SUFFIX,
	AIGATEWAYMODELMETADATASOURCEMATCHEDBY_CANONICAL_MODEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelMetadataSourceMatchedBy) GetAllowedValues() []AiGatewayModelMetadataSourceMatchedBy {
	return allowedAiGatewayModelMetadataSourceMatchedByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelMetadataSourceMatchedBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelMetadataSourceMatchedBy(value)
	return nil
}

// NewAiGatewayModelMetadataSourceMatchedByFromValue returns a pointer to a valid AiGatewayModelMetadataSourceMatchedBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelMetadataSourceMatchedByFromValue(v string) (*AiGatewayModelMetadataSourceMatchedBy, error) {
	ev := AiGatewayModelMetadataSourceMatchedBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelMetadataSourceMatchedBy: valid values are %v", v, allowedAiGatewayModelMetadataSourceMatchedByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelMetadataSourceMatchedBy) IsValid() bool {
	for _, existing := range allowedAiGatewayModelMetadataSourceMatchedByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelMetadataSourceMatchedBy value.
func (v AiGatewayModelMetadataSourceMatchedBy) Ptr() *AiGatewayModelMetadataSourceMatchedBy {
	return &v
}
