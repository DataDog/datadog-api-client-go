// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigLLMJudgeConfig LLM judge configuration for a custom evaluator.
type LLMObsCustomEvalConfigLLMJudgeConfig struct {
	// Criteria used to assess the pass/fail result of a custom evaluator.
	AssessmentCriteria *LLMObsCustomEvalConfigAssessmentCriteria `json:"assessment_criteria,omitempty"`
	// Query used to extract additional context for the evaluation.
	ContextQuery datadog.NullableString `json:"context_query,omitempty"`
	// LLM inference parameters for a custom evaluator.
	InferenceParams LLMObsCustomEvalConfigInferenceParams `json:"inference_params"`
	// Name of the last library prompt template used.
	LastUsedLibraryPromptTemplateName datadog.NullableString `json:"last_used_library_prompt_template_name,omitempty"`
	// Whether the library prompt template was modified.
	ModifiedLibraryPromptTemplate datadog.NullableBool `json:"modified_library_prompt_template,omitempty"`
	// JSON schema describing the expected output format of the LLM judge.
	OutputSchema map[string]interface{} `json:"output_schema,omitempty"`
	// Output parsing type for a custom LLM judge evaluator.
	ParsingType *LLMObsCustomEvalConfigParsingType `json:"parsing_type,omitempty"`
	// List of messages forming the LLM judge prompt template.
	PromptTemplate []LLMObsCustomEvalConfigPromptMessage `json:"prompt_template,omitempty"`
	// Query used to extract the target value to evaluate.
	TargetQuery datadog.NullableString `json:"target_query,omitempty"`
	// User-provided function applied to post-process the JSON output of the LLM judge.
	UserSpecifiedJsonPostProcessingFunction datadog.NullableString `json:"user_specified_json_post_processing_function,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigLLMJudgeConfig instantiates a new LLMObsCustomEvalConfigLLMJudgeConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigLLMJudgeConfig(inferenceParams LLMObsCustomEvalConfigInferenceParams) *LLMObsCustomEvalConfigLLMJudgeConfig {
	this := LLMObsCustomEvalConfigLLMJudgeConfig{}
	this.InferenceParams = inferenceParams
	return &this
}

// NewLLMObsCustomEvalConfigLLMJudgeConfigWithDefaults instantiates a new LLMObsCustomEvalConfigLLMJudgeConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigLLMJudgeConfigWithDefaults() *LLMObsCustomEvalConfigLLMJudgeConfig {
	this := LLMObsCustomEvalConfigLLMJudgeConfig{}
	return &this
}

// GetAssessmentCriteria returns the AssessmentCriteria field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetAssessmentCriteria() LLMObsCustomEvalConfigAssessmentCriteria {
	if o == nil || o.AssessmentCriteria == nil {
		var ret LLMObsCustomEvalConfigAssessmentCriteria
		return ret
	}
	return *o.AssessmentCriteria
}

// GetAssessmentCriteriaOk returns a tuple with the AssessmentCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetAssessmentCriteriaOk() (*LLMObsCustomEvalConfigAssessmentCriteria, bool) {
	if o == nil || o.AssessmentCriteria == nil {
		return nil, false
	}
	return o.AssessmentCriteria, true
}

// HasAssessmentCriteria returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) HasAssessmentCriteria() bool {
	return o != nil && o.AssessmentCriteria != nil
}

// SetAssessmentCriteria gets a reference to the given LLMObsCustomEvalConfigAssessmentCriteria and assigns it to the AssessmentCriteria field.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetAssessmentCriteria(v LLMObsCustomEvalConfigAssessmentCriteria) {
	o.AssessmentCriteria = &v
}

// GetContextQuery returns the ContextQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetContextQuery() string {
	if o == nil || o.ContextQuery.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContextQuery.Get()
}

// GetContextQueryOk returns a tuple with the ContextQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetContextQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContextQuery.Get(), o.ContextQuery.IsSet()
}

// HasContextQuery returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) HasContextQuery() bool {
	return o != nil && o.ContextQuery.IsSet()
}

// SetContextQuery gets a reference to the given datadog.NullableString and assigns it to the ContextQuery field.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetContextQuery(v string) {
	o.ContextQuery.Set(&v)
}

// SetContextQueryNil sets the value for ContextQuery to be an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetContextQueryNil() {
	o.ContextQuery.Set(nil)
}

// UnsetContextQuery ensures that no value is present for ContextQuery, not even an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) UnsetContextQuery() {
	o.ContextQuery.Unset()
}

// GetInferenceParams returns the InferenceParams field value.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetInferenceParams() LLMObsCustomEvalConfigInferenceParams {
	if o == nil {
		var ret LLMObsCustomEvalConfigInferenceParams
		return ret
	}
	return o.InferenceParams
}

// GetInferenceParamsOk returns a tuple with the InferenceParams field value
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetInferenceParamsOk() (*LLMObsCustomEvalConfigInferenceParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InferenceParams, true
}

// SetInferenceParams sets field value.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetInferenceParams(v LLMObsCustomEvalConfigInferenceParams) {
	o.InferenceParams = v
}

// GetLastUsedLibraryPromptTemplateName returns the LastUsedLibraryPromptTemplateName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetLastUsedLibraryPromptTemplateName() string {
	if o == nil || o.LastUsedLibraryPromptTemplateName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastUsedLibraryPromptTemplateName.Get()
}

// GetLastUsedLibraryPromptTemplateNameOk returns a tuple with the LastUsedLibraryPromptTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetLastUsedLibraryPromptTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsedLibraryPromptTemplateName.Get(), o.LastUsedLibraryPromptTemplateName.IsSet()
}

// HasLastUsedLibraryPromptTemplateName returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) HasLastUsedLibraryPromptTemplateName() bool {
	return o != nil && o.LastUsedLibraryPromptTemplateName.IsSet()
}

// SetLastUsedLibraryPromptTemplateName gets a reference to the given datadog.NullableString and assigns it to the LastUsedLibraryPromptTemplateName field.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetLastUsedLibraryPromptTemplateName(v string) {
	o.LastUsedLibraryPromptTemplateName.Set(&v)
}

// SetLastUsedLibraryPromptTemplateNameNil sets the value for LastUsedLibraryPromptTemplateName to be an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetLastUsedLibraryPromptTemplateNameNil() {
	o.LastUsedLibraryPromptTemplateName.Set(nil)
}

// UnsetLastUsedLibraryPromptTemplateName ensures that no value is present for LastUsedLibraryPromptTemplateName, not even an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) UnsetLastUsedLibraryPromptTemplateName() {
	o.LastUsedLibraryPromptTemplateName.Unset()
}

// GetModifiedLibraryPromptTemplate returns the ModifiedLibraryPromptTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetModifiedLibraryPromptTemplate() bool {
	if o == nil || o.ModifiedLibraryPromptTemplate.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ModifiedLibraryPromptTemplate.Get()
}

// GetModifiedLibraryPromptTemplateOk returns a tuple with the ModifiedLibraryPromptTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetModifiedLibraryPromptTemplateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedLibraryPromptTemplate.Get(), o.ModifiedLibraryPromptTemplate.IsSet()
}

// HasModifiedLibraryPromptTemplate returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) HasModifiedLibraryPromptTemplate() bool {
	return o != nil && o.ModifiedLibraryPromptTemplate.IsSet()
}

// SetModifiedLibraryPromptTemplate gets a reference to the given datadog.NullableBool and assigns it to the ModifiedLibraryPromptTemplate field.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetModifiedLibraryPromptTemplate(v bool) {
	o.ModifiedLibraryPromptTemplate.Set(&v)
}

// SetModifiedLibraryPromptTemplateNil sets the value for ModifiedLibraryPromptTemplate to be an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetModifiedLibraryPromptTemplateNil() {
	o.ModifiedLibraryPromptTemplate.Set(nil)
}

// UnsetModifiedLibraryPromptTemplate ensures that no value is present for ModifiedLibraryPromptTemplate, not even an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) UnsetModifiedLibraryPromptTemplate() {
	o.ModifiedLibraryPromptTemplate.Unset()
}

// GetOutputSchema returns the OutputSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetOutputSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.OutputSchema
}

// GetOutputSchemaOk returns a tuple with the OutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetOutputSchemaOk() (*map[string]interface{}, bool) {
	if o == nil || o.OutputSchema == nil {
		return nil, false
	}
	return &o.OutputSchema, true
}

// HasOutputSchema returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) HasOutputSchema() bool {
	return o != nil && o.OutputSchema != nil
}

// SetOutputSchema gets a reference to the given map[string]interface{} and assigns it to the OutputSchema field.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetOutputSchema(v map[string]interface{}) {
	o.OutputSchema = v
}

// GetParsingType returns the ParsingType field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetParsingType() LLMObsCustomEvalConfigParsingType {
	if o == nil || o.ParsingType == nil {
		var ret LLMObsCustomEvalConfigParsingType
		return ret
	}
	return *o.ParsingType
}

// GetParsingTypeOk returns a tuple with the ParsingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetParsingTypeOk() (*LLMObsCustomEvalConfigParsingType, bool) {
	if o == nil || o.ParsingType == nil {
		return nil, false
	}
	return o.ParsingType, true
}

// HasParsingType returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) HasParsingType() bool {
	return o != nil && o.ParsingType != nil
}

// SetParsingType gets a reference to the given LLMObsCustomEvalConfigParsingType and assigns it to the ParsingType field.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetParsingType(v LLMObsCustomEvalConfigParsingType) {
	o.ParsingType = &v
}

// GetPromptTemplate returns the PromptTemplate field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetPromptTemplate() []LLMObsCustomEvalConfigPromptMessage {
	if o == nil || o.PromptTemplate == nil {
		var ret []LLMObsCustomEvalConfigPromptMessage
		return ret
	}
	return o.PromptTemplate
}

// GetPromptTemplateOk returns a tuple with the PromptTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetPromptTemplateOk() (*[]LLMObsCustomEvalConfigPromptMessage, bool) {
	if o == nil || o.PromptTemplate == nil {
		return nil, false
	}
	return &o.PromptTemplate, true
}

// HasPromptTemplate returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) HasPromptTemplate() bool {
	return o != nil && o.PromptTemplate != nil
}

// SetPromptTemplate gets a reference to the given []LLMObsCustomEvalConfigPromptMessage and assigns it to the PromptTemplate field.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetPromptTemplate(v []LLMObsCustomEvalConfigPromptMessage) {
	o.PromptTemplate = v
}

// GetTargetQuery returns the TargetQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetTargetQuery() string {
	if o == nil || o.TargetQuery.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetQuery.Get()
}

// GetTargetQueryOk returns a tuple with the TargetQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetTargetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetQuery.Get(), o.TargetQuery.IsSet()
}

// HasTargetQuery returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) HasTargetQuery() bool {
	return o != nil && o.TargetQuery.IsSet()
}

// SetTargetQuery gets a reference to the given datadog.NullableString and assigns it to the TargetQuery field.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetTargetQuery(v string) {
	o.TargetQuery.Set(&v)
}

// SetTargetQueryNil sets the value for TargetQuery to be an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetTargetQueryNil() {
	o.TargetQuery.Set(nil)
}

// UnsetTargetQuery ensures that no value is present for TargetQuery, not even an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) UnsetTargetQuery() {
	o.TargetQuery.Unset()
}

// GetUserSpecifiedJsonPostProcessingFunction returns the UserSpecifiedJsonPostProcessingFunction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetUserSpecifiedJsonPostProcessingFunction() string {
	if o == nil || o.UserSpecifiedJsonPostProcessingFunction.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserSpecifiedJsonPostProcessingFunction.Get()
}

// GetUserSpecifiedJsonPostProcessingFunctionOk returns a tuple with the UserSpecifiedJsonPostProcessingFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) GetUserSpecifiedJsonPostProcessingFunctionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserSpecifiedJsonPostProcessingFunction.Get(), o.UserSpecifiedJsonPostProcessingFunction.IsSet()
}

// HasUserSpecifiedJsonPostProcessingFunction returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) HasUserSpecifiedJsonPostProcessingFunction() bool {
	return o != nil && o.UserSpecifiedJsonPostProcessingFunction.IsSet()
}

// SetUserSpecifiedJsonPostProcessingFunction gets a reference to the given datadog.NullableString and assigns it to the UserSpecifiedJsonPostProcessingFunction field.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetUserSpecifiedJsonPostProcessingFunction(v string) {
	o.UserSpecifiedJsonPostProcessingFunction.Set(&v)
}

// SetUserSpecifiedJsonPostProcessingFunctionNil sets the value for UserSpecifiedJsonPostProcessingFunction to be an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) SetUserSpecifiedJsonPostProcessingFunctionNil() {
	o.UserSpecifiedJsonPostProcessingFunction.Set(nil)
}

// UnsetUserSpecifiedJsonPostProcessingFunction ensures that no value is present for UserSpecifiedJsonPostProcessingFunction, not even an explicit nil.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) UnsetUserSpecifiedJsonPostProcessingFunction() {
	o.UserSpecifiedJsonPostProcessingFunction.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigLLMJudgeConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssessmentCriteria != nil {
		toSerialize["assessment_criteria"] = o.AssessmentCriteria
	}
	if o.ContextQuery.IsSet() {
		toSerialize["context_query"] = o.ContextQuery.Get()
	}
	toSerialize["inference_params"] = o.InferenceParams
	if o.LastUsedLibraryPromptTemplateName.IsSet() {
		toSerialize["last_used_library_prompt_template_name"] = o.LastUsedLibraryPromptTemplateName.Get()
	}
	if o.ModifiedLibraryPromptTemplate.IsSet() {
		toSerialize["modified_library_prompt_template"] = o.ModifiedLibraryPromptTemplate.Get()
	}
	if o.OutputSchema != nil {
		toSerialize["output_schema"] = o.OutputSchema
	}
	if o.ParsingType != nil {
		toSerialize["parsing_type"] = o.ParsingType
	}
	if o.PromptTemplate != nil {
		toSerialize["prompt_template"] = o.PromptTemplate
	}
	if o.TargetQuery.IsSet() {
		toSerialize["target_query"] = o.TargetQuery.Get()
	}
	if o.UserSpecifiedJsonPostProcessingFunction.IsSet() {
		toSerialize["user_specified_json_post_processing_function"] = o.UserSpecifiedJsonPostProcessingFunction.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigLLMJudgeConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssessmentCriteria                      *LLMObsCustomEvalConfigAssessmentCriteria `json:"assessment_criteria,omitempty"`
		ContextQuery                            datadog.NullableString                    `json:"context_query,omitempty"`
		InferenceParams                         *LLMObsCustomEvalConfigInferenceParams    `json:"inference_params"`
		LastUsedLibraryPromptTemplateName       datadog.NullableString                    `json:"last_used_library_prompt_template_name,omitempty"`
		ModifiedLibraryPromptTemplate           datadog.NullableBool                      `json:"modified_library_prompt_template,omitempty"`
		OutputSchema                            map[string]interface{}                    `json:"output_schema,omitempty"`
		ParsingType                             *LLMObsCustomEvalConfigParsingType        `json:"parsing_type,omitempty"`
		PromptTemplate                          []LLMObsCustomEvalConfigPromptMessage     `json:"prompt_template,omitempty"`
		TargetQuery                             datadog.NullableString                    `json:"target_query,omitempty"`
		UserSpecifiedJsonPostProcessingFunction datadog.NullableString                    `json:"user_specified_json_post_processing_function,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InferenceParams == nil {
		return fmt.Errorf("required field inference_params missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assessment_criteria", "context_query", "inference_params", "last_used_library_prompt_template_name", "modified_library_prompt_template", "output_schema", "parsing_type", "prompt_template", "target_query", "user_specified_json_post_processing_function"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AssessmentCriteria != nil && all.AssessmentCriteria.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AssessmentCriteria = all.AssessmentCriteria
	o.ContextQuery = all.ContextQuery
	if all.InferenceParams.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.InferenceParams = *all.InferenceParams
	o.LastUsedLibraryPromptTemplateName = all.LastUsedLibraryPromptTemplateName
	o.ModifiedLibraryPromptTemplate = all.ModifiedLibraryPromptTemplate
	o.OutputSchema = all.OutputSchema
	if all.ParsingType != nil && !all.ParsingType.IsValid() {
		hasInvalidField = true
	} else {
		o.ParsingType = all.ParsingType
	}
	o.PromptTemplate = all.PromptTemplate
	o.TargetQuery = all.TargetQuery
	o.UserSpecifiedJsonPostProcessingFunction = all.UserSpecifiedJsonPostProcessingFunction

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
