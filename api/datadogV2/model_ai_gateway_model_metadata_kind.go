// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelMetadataKind Primary model kind. Decision models return a choice or score instead of generated text.
type AiGatewayModelMetadataKind string

// List of AiGatewayModelMetadataKind.
const (
	AIGATEWAYMODELMETADATAKIND_LANGUAGE  AiGatewayModelMetadataKind = "language"
	AIGATEWAYMODELMETADATAKIND_EMBEDDING AiGatewayModelMetadataKind = "embedding"
	AIGATEWAYMODELMETADATAKIND_IMAGE     AiGatewayModelMetadataKind = "image"
	AIGATEWAYMODELMETADATAKIND_AUDIO     AiGatewayModelMetadataKind = "audio"
	AIGATEWAYMODELMETADATAKIND_DECISION  AiGatewayModelMetadataKind = "decision"
	AIGATEWAYMODELMETADATAKIND_RERANKING AiGatewayModelMetadataKind = "reranking"
)

var allowedAiGatewayModelMetadataKindEnumValues = []AiGatewayModelMetadataKind{
	AIGATEWAYMODELMETADATAKIND_LANGUAGE,
	AIGATEWAYMODELMETADATAKIND_EMBEDDING,
	AIGATEWAYMODELMETADATAKIND_IMAGE,
	AIGATEWAYMODELMETADATAKIND_AUDIO,
	AIGATEWAYMODELMETADATAKIND_DECISION,
	AIGATEWAYMODELMETADATAKIND_RERANKING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelMetadataKind) GetAllowedValues() []AiGatewayModelMetadataKind {
	return allowedAiGatewayModelMetadataKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelMetadataKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelMetadataKind(value)
	return nil
}

// NewAiGatewayModelMetadataKindFromValue returns a pointer to a valid AiGatewayModelMetadataKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelMetadataKindFromValue(v string) (*AiGatewayModelMetadataKind, error) {
	ev := AiGatewayModelMetadataKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelMetadataKind: valid values are %v", v, allowedAiGatewayModelMetadataKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelMetadataKind) IsValid() bool {
	for _, existing := range allowedAiGatewayModelMetadataKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelMetadataKind value.
func (v AiGatewayModelMetadataKind) Ptr() *AiGatewayModelMetadataKind {
	return &v
}
