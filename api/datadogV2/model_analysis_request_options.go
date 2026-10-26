// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisRequestOptions Analyzer options sent by browser-based clients. Accepted but ignored: the API always runs
// the analysis with tree-sitter enabled and rule output logged.
type AnalysisRequestOptions struct {
	// Whether to capture the output logged by rules during execution. Ignored by this operation.
	LogOutput *bool `json:"log_output,omitempty"`
	// Whether to use tree-sitter to parse the source code. Ignored by this operation.
	UseTreeSitter *bool `json:"use_tree_sitter,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisRequestOptions instantiates a new AnalysisRequestOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisRequestOptions() *AnalysisRequestOptions {
	this := AnalysisRequestOptions{}
	return &this
}

// NewAnalysisRequestOptionsWithDefaults instantiates a new AnalysisRequestOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisRequestOptionsWithDefaults() *AnalysisRequestOptions {
	this := AnalysisRequestOptions{}
	return &this
}

// GetLogOutput returns the LogOutput field value if set, zero value otherwise.
func (o *AnalysisRequestOptions) GetLogOutput() bool {
	if o == nil || o.LogOutput == nil {
		var ret bool
		return ret
	}
	return *o.LogOutput
}

// GetLogOutputOk returns a tuple with the LogOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestOptions) GetLogOutputOk() (*bool, bool) {
	if o == nil || o.LogOutput == nil {
		return nil, false
	}
	return o.LogOutput, true
}

// HasLogOutput returns a boolean if a field has been set.
func (o *AnalysisRequestOptions) HasLogOutput() bool {
	return o != nil && o.LogOutput != nil
}

// SetLogOutput gets a reference to the given bool and assigns it to the LogOutput field.
func (o *AnalysisRequestOptions) SetLogOutput(v bool) {
	o.LogOutput = &v
}

// GetUseTreeSitter returns the UseTreeSitter field value if set, zero value otherwise.
func (o *AnalysisRequestOptions) GetUseTreeSitter() bool {
	if o == nil || o.UseTreeSitter == nil {
		var ret bool
		return ret
	}
	return *o.UseTreeSitter
}

// GetUseTreeSitterOk returns a tuple with the UseTreeSitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestOptions) GetUseTreeSitterOk() (*bool, bool) {
	if o == nil || o.UseTreeSitter == nil {
		return nil, false
	}
	return o.UseTreeSitter, true
}

// HasUseTreeSitter returns a boolean if a field has been set.
func (o *AnalysisRequestOptions) HasUseTreeSitter() bool {
	return o != nil && o.UseTreeSitter != nil
}

// SetUseTreeSitter gets a reference to the given bool and assigns it to the UseTreeSitter field.
func (o *AnalysisRequestOptions) SetUseTreeSitter(v bool) {
	o.UseTreeSitter = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LogOutput != nil {
		toSerialize["log_output"] = o.LogOutput
	}
	if o.UseTreeSitter != nil {
		toSerialize["use_tree_sitter"] = o.UseTreeSitter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisRequestOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LogOutput     *bool `json:"log_output,omitempty"`
		UseTreeSitter *bool `json:"use_tree_sitter,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"log_output", "use_tree_sitter"})
	} else {
		return err
	}
	o.LogOutput = all.LogOutput
	o.UseTreeSitter = all.UseTreeSitter

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
