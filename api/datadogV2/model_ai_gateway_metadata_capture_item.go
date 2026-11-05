// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayMetadataCaptureItem An item Metadata-only capture can add to LLM Observability spans. Each is structure or a number, never conversation text. `tool_usage` adds tool call and result counts; `tool_output_size` adds estimated tool output tokens, total and average; `tool_names` adds each tool call and result with its name, ID and type; `conversation_shape` adds every message with its role and no text; `input_size_by_role` adds estimated input tokens per role; `available_tools` adds the names of the tools the request offered; `generation_settings` adds numeric and enumerated settings such as temperature and the finish reason.
type AiGatewayMetadataCaptureItem string

// List of AiGatewayMetadataCaptureItem.
const (
	AIGATEWAYMETADATACAPTUREITEM_TOOL_USAGE          AiGatewayMetadataCaptureItem = "tool_usage"
	AIGATEWAYMETADATACAPTUREITEM_TOOL_OUTPUT_SIZE    AiGatewayMetadataCaptureItem = "tool_output_size"
	AIGATEWAYMETADATACAPTUREITEM_TOOL_NAMES          AiGatewayMetadataCaptureItem = "tool_names"
	AIGATEWAYMETADATACAPTUREITEM_CONVERSATION_SHAPE  AiGatewayMetadataCaptureItem = "conversation_shape"
	AIGATEWAYMETADATACAPTUREITEM_INPUT_SIZE_BY_ROLE  AiGatewayMetadataCaptureItem = "input_size_by_role"
	AIGATEWAYMETADATACAPTUREITEM_AVAILABLE_TOOLS     AiGatewayMetadataCaptureItem = "available_tools"
	AIGATEWAYMETADATACAPTUREITEM_GENERATION_SETTINGS AiGatewayMetadataCaptureItem = "generation_settings"
)

var allowedAiGatewayMetadataCaptureItemEnumValues = []AiGatewayMetadataCaptureItem{
	AIGATEWAYMETADATACAPTUREITEM_TOOL_USAGE,
	AIGATEWAYMETADATACAPTUREITEM_TOOL_OUTPUT_SIZE,
	AIGATEWAYMETADATACAPTUREITEM_TOOL_NAMES,
	AIGATEWAYMETADATACAPTUREITEM_CONVERSATION_SHAPE,
	AIGATEWAYMETADATACAPTUREITEM_INPUT_SIZE_BY_ROLE,
	AIGATEWAYMETADATACAPTUREITEM_AVAILABLE_TOOLS,
	AIGATEWAYMETADATACAPTUREITEM_GENERATION_SETTINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayMetadataCaptureItem) GetAllowedValues() []AiGatewayMetadataCaptureItem {
	return allowedAiGatewayMetadataCaptureItemEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayMetadataCaptureItem) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayMetadataCaptureItem(value)
	return nil
}

// NewAiGatewayMetadataCaptureItemFromValue returns a pointer to a valid AiGatewayMetadataCaptureItem
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayMetadataCaptureItemFromValue(v string) (*AiGatewayMetadataCaptureItem, error) {
	ev := AiGatewayMetadataCaptureItem(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayMetadataCaptureItem: valid values are %v", v, allowedAiGatewayMetadataCaptureItemEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayMetadataCaptureItem) IsValid() bool {
	for _, existing := range allowedAiGatewayMetadataCaptureItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayMetadataCaptureItem value.
func (v AiGatewayMetadataCaptureItem) Ptr() *AiGatewayMetadataCaptureItem {
	return &v
}
