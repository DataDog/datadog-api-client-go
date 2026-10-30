// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigIntegrationProvider Name of the LLM integration provider. Use `deterministic` for a code evaluator, which runs `llm_judge_config.user_specified_json_post_processing_function` instead of calling an LLM. `insights_sponsored` is returned for evaluators that Datadog manages and cannot be set.
type LLMObsCustomEvalConfigIntegrationProvider string

// List of LLMObsCustomEvalConfigIntegrationProvider.
const (
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_OPENAI             LLMObsCustomEvalConfigIntegrationProvider = "openai"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_AMAZON_BEDROCK     LLMObsCustomEvalConfigIntegrationProvider = "amazon_bedrock"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_ANTHROPIC          LLMObsCustomEvalConfigIntegrationProvider = "anthropic"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_AZURE_OPENAI       LLMObsCustomEvalConfigIntegrationProvider = "azure_openai"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_VERTEX_AI          LLMObsCustomEvalConfigIntegrationProvider = "vertex_ai"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_LLM_PROXY          LLMObsCustomEvalConfigIntegrationProvider = "llmproxy"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_DETERMINISTIC      LLMObsCustomEvalConfigIntegrationProvider = "deterministic"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_TYPESAFE           LLMObsCustomEvalConfigIntegrationProvider = "typesafe"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_DATADOG            LLMObsCustomEvalConfigIntegrationProvider = "datadog"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_DATADOG_AI_GATEWAY LLMObsCustomEvalConfigIntegrationProvider = "datadog_ai_gateway"
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_INSIGHTS_SPONSORED LLMObsCustomEvalConfigIntegrationProvider = "insights_sponsored"
)

var allowedLLMObsCustomEvalConfigIntegrationProviderEnumValues = []LLMObsCustomEvalConfigIntegrationProvider{
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_OPENAI,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_AMAZON_BEDROCK,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_ANTHROPIC,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_AZURE_OPENAI,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_VERTEX_AI,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_LLM_PROXY,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_DETERMINISTIC,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_TYPESAFE,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_DATADOG,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_DATADOG_AI_GATEWAY,
	LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_INSIGHTS_SPONSORED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsCustomEvalConfigIntegrationProvider) GetAllowedValues() []LLMObsCustomEvalConfigIntegrationProvider {
	return allowedLLMObsCustomEvalConfigIntegrationProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsCustomEvalConfigIntegrationProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsCustomEvalConfigIntegrationProvider(value)
	return nil
}

// NewLLMObsCustomEvalConfigIntegrationProviderFromValue returns a pointer to a valid LLMObsCustomEvalConfigIntegrationProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsCustomEvalConfigIntegrationProviderFromValue(v string) (*LLMObsCustomEvalConfigIntegrationProvider, error) {
	ev := LLMObsCustomEvalConfigIntegrationProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsCustomEvalConfigIntegrationProvider: valid values are %v", v, allowedLLMObsCustomEvalConfigIntegrationProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsCustomEvalConfigIntegrationProvider) IsValid() bool {
	for _, existing := range allowedLLMObsCustomEvalConfigIntegrationProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsCustomEvalConfigIntegrationProvider value.
func (v LLMObsCustomEvalConfigIntegrationProvider) Ptr() *LLMObsCustomEvalConfigIntegrationProvider {
	return &v
}
