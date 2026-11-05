// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelModalities Supported model input and output modalities.
type AiGatewayModelModalities struct {
	//
	Input []string `json:"input,omitempty"`
	//
	Output []string `json:"output,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelModalities instantiates a new AiGatewayModelModalities object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelModalities() *AiGatewayModelModalities {
	this := AiGatewayModelModalities{}
	return &this
}

// NewAiGatewayModelModalitiesWithDefaults instantiates a new AiGatewayModelModalities object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelModalitiesWithDefaults() *AiGatewayModelModalities {
	this := AiGatewayModelModalities{}
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *AiGatewayModelModalities) GetInput() []string {
	if o == nil || o.Input == nil {
		var ret []string
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelModalities) GetInputOk() (*[]string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *AiGatewayModelModalities) HasInput() bool {
	return o != nil && o.Input != nil
}

// SetInput gets a reference to the given []string and assigns it to the Input field.
func (o *AiGatewayModelModalities) SetInput(v []string) {
	o.Input = v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *AiGatewayModelModalities) GetOutput() []string {
	if o == nil || o.Output == nil {
		var ret []string
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelModalities) GetOutputOk() (*[]string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *AiGatewayModelModalities) HasOutput() bool {
	return o != nil && o.Output != nil
}

// SetOutput gets a reference to the given []string and assigns it to the Output field.
func (o *AiGatewayModelModalities) SetOutput(v []string) {
	o.Output = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelModalities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelModalities) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Input  []string `json:"input,omitempty"`
		Output []string `json:"output,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"input", "output"})
	} else {
		return err
	}
	o.Input = all.Input
	o.Output = all.Output

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
