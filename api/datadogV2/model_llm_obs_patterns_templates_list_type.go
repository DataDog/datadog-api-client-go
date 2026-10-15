// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsTemplatesListType Resource type of a list of Agent Observability patterns templates.
type LLMObsPatternsTemplatesListType string

// List of LLMObsPatternsTemplatesListType.
const (
	LLMOBSPATTERNSTEMPLATESLISTTYPE_LIST_TOPIC_DISCOVERY_TEMPLATES_RESPONSE LLMObsPatternsTemplatesListType = "list_topic_discovery_templates_response"
)

var allowedLLMObsPatternsTemplatesListTypeEnumValues = []LLMObsPatternsTemplatesListType{
	LLMOBSPATTERNSTEMPLATESLISTTYPE_LIST_TOPIC_DISCOVERY_TEMPLATES_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsTemplatesListType) GetAllowedValues() []LLMObsPatternsTemplatesListType {
	return allowedLLMObsPatternsTemplatesListTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsTemplatesListType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsTemplatesListType(value)
	return nil
}

// NewLLMObsPatternsTemplatesListTypeFromValue returns a pointer to a valid LLMObsPatternsTemplatesListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsTemplatesListTypeFromValue(v string) (*LLMObsPatternsTemplatesListType, error) {
	ev := LLMObsPatternsTemplatesListType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsTemplatesListType: valid values are %v", v, allowedLLMObsPatternsTemplatesListTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsTemplatesListType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsTemplatesListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsTemplatesListType value.
func (v LLMObsPatternsTemplatesListType) Ptr() *LLMObsPatternsTemplatesListType {
	return &v
}
